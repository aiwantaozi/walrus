package walrus

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/seal-io/utils/pools/gopool"
	"github.com/seal-io/utils/stringx"
	"golang.org/x/crypto/bcrypt"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/kubemeta"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemkuberes"
	"github.com/seal-io/walrus/pkg/systemmeta"
)

// SubjectHandler handles v1.Subject objects.
//
// SubjectHandler maps the v1.Subject object to a Kubernetes ServiceAccount resource,
// which is named as "walrus-subject-<name>".
type SubjectHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CurdOperations

	Client    ctrlcli.Client
	APIReader ctrlcli.Reader
}

func (h *SubjectHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Configure field indexer.
	fi := opts.Manager.GetFieldIndexer()
	err = fi.IndexField(ctx, &core.ServiceAccount{}, "metadata.name",
		func(obj ctrlcli.Object) []string {
			if obj == nil {
				return nil
			}
			return []string{obj.GetName()}
		})
	if err != nil {
		return
	}

	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("subjects")

	// As storage.
	h.ObjectInfo = &walrus.Subject{}
	h.CurdOperations = extensionapi.WithCurd(nil, h)

	// Set client.
	h.Client = opts.Manager.GetClient()
	h.APIReader = opts.Manager.GetAPIReader()

	// Create subresource handlers.
	srs = map[string]rest.Storage{
		"login": newSubjectLoginHandler(opts),
		"token": newSubjectTokenHandler(opts),
	}

	return
}

var (
	_ rest.Storage           = (*SubjectHandler)(nil)
	_ rest.Creater           = (*SubjectHandler)(nil)
	_ rest.Lister            = (*SubjectHandler)(nil)
	_ rest.Watcher           = (*SubjectHandler)(nil)
	_ rest.Getter            = (*SubjectHandler)(nil)
	_ rest.Updater           = (*SubjectHandler)(nil)
	_ rest.Patcher           = (*SubjectHandler)(nil)
	_ rest.GracefulDeleter   = (*SubjectHandler)(nil)
	_ rest.CollectionDeleter = (*SubjectHandler)(nil)
)

func (h *SubjectHandler) New() runtime.Object {
	return &walrus.Subject{}
}

func (h *SubjectHandler) Destroy() {
}

func (h *SubjectHandler) OnCreate(ctx context.Context, obj runtime.Object, opts ctrlcli.CreateOptions) (runtime.Object, error) {
	// Validate.
	subj := obj.(*walrus.Subject)
	{
		var errs field.ErrorList
		if subj.Namespace != systemkuberes.SystemNamespaceName {
			errs = append(errs, field.Invalid(
				field.NewPath("metadata.namespace"), subj.Namespace, "subject namespace must be "+systemkuberes.SystemNamespaceName))
		}
		if stringx.StringWidth(subj.Name) > 30 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("metadata.name"), stringx.StringWidth(subj.Name), 30))
		}
		if subj.Spec.Provider == systemkuberes.DefaultSubjectProviderName && strings.Contains(subj.Name, ".") {
			errs = append(errs, field.Invalid(
				field.NewPath("metadata.name"), subj.Name, "name must not contain '.'"))
		}
		if subj.Spec.Provider == "" {
			errs = append(errs, field.Required(
				field.NewPath("spec.provider"), "provider must be specified"))
		}
		if err := subj.Spec.Role.Validate(); err != nil {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.role"), subj.Spec.Role, err.Error()))
		}
		if stringx.StringWidth(subj.Spec.DisplayName) > 30 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.displayName"), stringx.StringWidth(subj.Spec.DisplayName), 30))
		}
		if stringx.StringWidth(subj.Spec.Description) > 50 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.description"), stringx.StringWidth(subj.Spec.Description), 50))
		}
		if subj.Spec.Email == "" {
			errs = append(errs, field.Required(
				field.NewPath("spec.email"), "email must be specified"))
		}
		switch {
		case subj.Spec.Credential == nil:
			errs = append(errs, field.Required(
				field.NewPath("spec.credential"), "credential must be specified"))
		case subj.Spec.Provider == systemkuberes.DefaultSubjectProviderName:
			cred := ptr.Deref(subj.Spec.Credential, "")
			switch {
			case cred == "":
				errs = append(errs, field.Required(
					field.NewPath("spec.credential"), "credential must be specified"),
				)
			case len(cred) < 10:
				errs = append(errs, field.Forbidden(
					field.NewPath("spec.credential"), "credential must be at least 10 characters"),
				)
			case len(cred) > 72:
				errs = append(errs, field.Forbidden(
					field.NewPath("spec.credential"), "credential must be at most 72 characters"),
				)
			}
		}

		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjects"), subj.Name, errs)
		}
	}

	// Validate with provider.
	var subjProv *walrus.SubjectProvider
	{
		subjProv = &walrus.SubjectProvider{
			ObjectMeta: meta.ObjectMeta{
				Namespace: subj.Namespace,
				Name:      subj.Spec.Provider,
			},
		}
		err := h.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(subjProv), subjProv)
		if err != nil {
			errs := field.ErrorList{
				field.Invalid(
					field.NewPath("spec.provider"), subj.Spec.Provider, err.Error()),
			}
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjects"), subj.Name, errs)
		}
	}

	// NB(thxCode): the permission grant is handled by SubjectAuthzReconciler.

	// Create.
	{
		sa, err := convertServiceAccountFromSubject(subj)
		if err != nil {
			return nil, kerrors.NewInternalError(err)
		}
		kubemeta.ControlOn(sa, subjProv, walrus.SchemeGroupVersionKind("SubjectProvider"))
		err = h.Client.Create(ctx, sa, &opts)
		if err != nil {
			return nil, err
		}
		subj = convertSubjectFromServiceAccount(sa)
	}

	return subj, nil
}

func (h *SubjectHandler) NewList() runtime.Object {
	return &walrus.SubjectList{}
}

func (h *SubjectHandler) OnList(ctx context.Context, opts ctrlcli.ListOptions) (runtime.Object, error) {
	// List.
	saList := new(core.ServiceAccountList)
	err := h.APIReader.List(ctx, saList,
		convertServiceAccountListOptsFromSubjectListOpts(opts))
	if err != nil {
		return nil, err
	}

	// Convert.
	sList := convertSubjectListFromServiceAccountList(saList, opts)
	return sList, nil
}

func (h *SubjectHandler) OnWatch(ctx context.Context, opts ctrlcli.ListOptions) (watch.Interface, error) {
	// Watch.
	uw, err := h.Client.(ctrlcli.WithWatch).Watch(ctx, new(core.ServiceAccountList),
		convertServiceAccountListOptsFromSubjectListOpts(opts))
	if err != nil {
		return nil, err
	}

	c := make(chan watch.Event)
	dw := watch.NewProxyWatcher(c)
	gopool.Go(func() {
		defer close(c)
		defer uw.Stop()

		for {
			select {
			case <-ctx.Done():
				// Cancel by context.
				return
			case <-dw.StopChan():
				// Stop by downstream.
				return
			case e, ok := <-uw.ResultChan():
				if !ok {
					// Close by upstream.
					return
				}

				// Nothing to do.
				if e.Object == nil {
					c <- e
					continue
				}

				// Type assert.
				sa, ok := e.Object.(*core.ServiceAccount)
				if !ok {
					c <- e
					continue
				}

				// Process bookmark.
				if e.Type == watch.Bookmark {
					e.Object = &walrus.Subject{ObjectMeta: sa.ObjectMeta}
					c <- e
					continue
				}

				// Convert.
				subj := safeConvertSubjectFromServiceAccount(sa, opts.Namespace)
				if subj == nil {
					continue
				}

				// Dispatch.
				e.Object = subj
				c <- e
			}
		}
	})

	return dw, nil
}

func (h *SubjectHandler) OnGet(ctx context.Context, key types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	// Get.
	sa := &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Namespace: key.Namespace,
			Name:      systemauthz.ConvertServiceAccountNameFromSubjectName(key.Name),
		},
	}
	err := h.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(sa), sa, &opts)
	if err != nil {
		return nil, err
	}

	// Convert.
	subj := convertSubjectFromServiceAccount(sa)
	if subj == nil {
		return nil, kerrors.NewNotFound(walrus.SchemeResource("subjects"), key.Name)
	}
	return subj, nil
}

func (h *SubjectHandler) OnUpdate(ctx context.Context, obj, oldObj runtime.Object, opts ctrlcli.UpdateOptions) (runtime.Object, error) {
	// Validate.
	subj, oldSubj := obj.(*walrus.Subject), oldObj.(*walrus.Subject)
	{
		var errs field.ErrorList
		if subj.Spec.Provider != oldSubj.Spec.Provider {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.provider"), subj.Spec.Provider, "provider is immutable"))
		}
		if subj.Spec.Email == "" {
			errs = append(errs, field.Required(
				field.NewPath("spec.email"), "email must be specified"))
		}
		if subj.Spec.Provider == systemkuberes.DefaultSubjectProviderName && subj.Spec.Credential != nil {
			cred := ptr.Deref(subj.Spec.Credential, "")
			switch {
			case cred == "":
				errs = append(errs, field.Required(
					field.NewPath("spec.credential"), "credential must be specified"),
				)
			case len(cred) < 8:
				errs = append(errs, field.Invalid(
					field.NewPath("spec.credential"), cred, "credential must be at least 8 characters"),
				)
			case len(cred) > 72:
				errs = append(errs, field.Invalid(
					field.NewPath("spec.credential"), cred, "credential must be at most 72 characters"),
				)
			}
		}
		if err := subj.Spec.Role.Validate(); err != nil {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.role"), subj.Spec.Role, err.Error()))
		}
		if stringx.StringWidth(subj.Spec.DisplayName) > 30 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.displayName"), stringx.StringWidth(subj.Spec.DisplayName), 30))
		}
		if stringx.StringWidth(subj.Spec.Description) > 50 {
			errs = append(errs, field.TooLongMaxLength(
				field.NewPath("spec.description"), stringx.StringWidth(subj.Spec.Description), 50))
		}
		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjects"), subj.Name, errs)
		}
	}

	// NB(thxCode): we revoke the stale permission here if the role has changed,
	// and the new permission grant is handled by SubjectAuthzReconciler.
	if subj.Spec.Role != oldSubj.Spec.Role {
		err := systemauthz.RevokeSubject(ctx, h.Client, oldSubj)
		if err != nil {
			return nil, kerrors.NewInternalError(err)
		}
	}

	// Update.
	{
		sa, err := convertServiceAccountFromSubject(subj)
		if err != nil {
			return nil, kerrors.NewInternalError(err)
		}
		err = h.Client.Update(ctx, sa, &opts)
		if err != nil {
			return nil, err
		}
		subj = convertSubjectFromServiceAccount(sa)
	}

	return subj, nil
}

func (h *SubjectHandler) OnDelete(ctx context.Context, obj runtime.Object, opts ctrlcli.DeleteOptions) error {
	subj := obj.(*walrus.Subject)

	// Validate.
	{
		// Prevent deleting default subject provider.
		if subj.Name == systemkuberes.AdminSubjectName {
			return kerrors.NewBadRequest("admin subject is reserved")
		}
	}

	// Delete.
	sa, _ := convertServiceAccountFromSubject(subj)
	return h.Client.Delete(ctx, sa, &opts)
}

func convertServiceAccountListOptsFromSubjectListOpts(in ctrlcli.ListOptions) (out *ctrlcli.ListOptions) {
	if in.Namespace != systemkuberes.SystemNamespaceName {
		return &in
	}

	// Ignore name selector
	if in.FieldSelector != nil {
		reqs := slices.DeleteFunc(in.FieldSelector.Requirements(), func(req fields.Requirement) bool {
			return req.Field == "metadata.name"
		})
		if len(reqs) == 0 {
			in.FieldSelector = nil
		} else {
			in.FieldSelector = kubemeta.FieldSelectorFromRequirements(reqs)
		}
	}

	// Add necessary label selector.
	if lbs := systemmeta.GetResourcesLabelSelectorOfType("subjects"); in.LabelSelector == nil {
		in.LabelSelector = lbs
	} else {
		reqs, _ := lbs.Requirements()
		in.LabelSelector = in.LabelSelector.DeepCopySelector().Add(reqs...)
	}

	return &in
}

func convertServiceAccountFromSubject(subj *walrus.Subject) (*core.ServiceAccount, error) {
	sa := &core.ServiceAccount{
		ObjectMeta: subj.ObjectMeta,
	}
	sa.Name = systemauthz.ConvertServiceAccountNameFromSubjectName(subj.Name)

	notes := map[string]string{
		"provider":    subj.Spec.Provider,
		"role":        subj.Spec.Role.String(),
		"displayName": subj.Spec.DisplayName,
		"description": subj.Spec.Description,
		"groups":      strings.Join(subj.Spec.Groups, ","),
		"email":       subj.Spec.Email,
	}
	if subj.Spec.Credential != nil {
		// Armor.
		bs, err := bcrypt.GenerateFromPassword(stringx.ToBytes(subj.Spec.Credential), 12)
		if err != nil {
			return sa, fmt.Errorf("armor credential: %w", err)
		}
		notes["armorCredential"] = stringx.FromBytes(&bs)
	}
	systemmeta.NoteResource(sa, "subjects", notes)

	kubemeta.SanitizeLastAppliedAnnotation(sa)
	return sa, nil
}

func convertSubjectFromServiceAccount(sa *core.ServiceAccount) *walrus.Subject {
	if sa == nil {
		return nil
	}

	resType, notes := systemmeta.UnnoteResource(sa)
	if resType != "subjects" {
		return nil
	}
	subjName := systemauthz.ConvertSubjectNameFromServiceAccountName(sa.Name)
	if subjName == "" {
		return nil
	}

	subj := &walrus.Subject{
		ObjectMeta: sa.ObjectMeta,
		Spec: walrus.SubjectSpec{
			Provider:    notes["provider"],
			Role:        walrus.SubjectRole(notes["role"]),
			DisplayName: notes["displayName"],
			Description: notes["description"],
			Email:       notes["email"],
		},
	}
	if v := notes["groups"]; v != "" {
		subj.Spec.Groups = strings.Split(v, ",")
	}
	subj.Name = subjName

	kubemeta.OverwriteLastAppliedAnnotation(subj)
	return subj
}

func safeConvertSubjectFromServiceAccount(sa *core.ServiceAccount, reqNamespace string) *walrus.Subject {
	subj := convertSubjectFromServiceAccount(sa)
	if subj != nil && reqNamespace != "" && reqNamespace != subj.Namespace {
		// NB(thxCode): sanitize if the subject's namespace doesn't match requested namespace.
		subj = nil
	}
	return subj
}

func convertSubjectListFromServiceAccountList(saList *core.ServiceAccountList, opts ctrlcli.ListOptions) *walrus.SubjectList {
	if saList == nil {
		return &walrus.SubjectList{}
	}

	// Sort by resource version.
	sort.SliceStable(saList.Items, func(i, j int) bool {
		l, r := saList.Items[i].ResourceVersion, saList.Items[j].ResourceVersion
		return len(l) < len(r) ||
			(len(l) == len(r) && l < r)
	})

	sList := &walrus.SubjectList{
		Items: make([]walrus.Subject, 0, len(saList.Items)),
	}

	for i := range saList.Items {
		subj := safeConvertSubjectFromServiceAccount(&saList.Items[i], opts.Namespace)
		if subj == nil {
			continue
		}
		// Ignore if not be selected by `kubectl get --field-selector=metadata.namespace=...`.
		if fs := opts.FieldSelector; fs != nil &&
			!fs.Matches(fields.Set{"metadata.namespace": subj.Namespace, "metadata.name": subj.Name}) {
			continue
		}
		sList.Items = append(sList.Items, *subj)
	}

	return sList
}
