package walrus

import (
	"context"

	"github.com/seal-io/utils/funcx"
	"github.com/seal-io/utils/stringx"
	"golang.org/x/crypto/bcrypt"
	authentication "k8s.io/api/authentication/v1"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

// SubjectLoginHandler is a handler for v1.SubjectLogin objects,
// which is a subresource of v1.Subject objects.
type SubjectLoginHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CreateOperation

	Client ctrlcli.Client
}

func newSubjectLoginHandler(opts extensionapi.SetupOptions) *SubjectLoginHandler {
	h := &SubjectLoginHandler{}

	// As storage.
	h.ObjectInfo = &walrus.SubjectLogin{}
	h.CreateOperation = extensionapi.WithCreate(h)

	// Set client.
	h.Client = opts.Manager.GetClient()

	return h
}

var (
	_ rest.Storage = (*SubjectLoginHandler)(nil)
	_ rest.Creater = (*SubjectLoginHandler)(nil)
)

func (h *SubjectLoginHandler) New() runtime.Object {
	return &walrus.SubjectLogin{}
}

func (h *SubjectLoginHandler) Destroy() {}

func (h *SubjectLoginHandler) OnCreate(ctx context.Context, obj runtime.Object, _ ctrlcli.CreateOptions) (runtime.Object, error) {
	subjl := obj.(*walrus.SubjectLogin)

	// Validate.
	{
		if subjl.Spec.Credential == "" {
			errs := field.ErrorList{field.Required(
				field.NewPath("spec.credential"), "credential is required")}
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjectlogins"), subjl.Name, errs)
		}
	}

	// Get the parent Subject object via underlay resource.
	sa := &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Namespace: subjl.Namespace,
			Name:      systemauthz.ConvertServiceAccountNameFromSubjectName(subjl.Name),
		},
	}
	err := h.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(sa), sa)
	if err != nil {
		return nil, kerrors.NewBadRequest("subject is not found")
	}

	resType, notes := systemmeta.DescribeResource(sa)
	if resType != "subjects" {
		return nil, kerrors.NewBadRequest("subject is not found")
	}

	armorCredential := stringx.ToBytes(ptr.To(notes["armorCredential"]))
	if bcrypt.CompareHashAndPassword(armorCredential, stringx.ToBytes(&subjl.Spec.Credential)) != nil {
		return nil, kerrors.NewBadRequest("credential is mismatched")
	}

	// Generate a token.
	es := funcx.NoError(systemsetting.SubjectLoginExpirationSeconds.ValueInt64(ctx))
	if es < 3600 {
		es = 3600
	}
	tr := &authentication.TokenRequest{
		Spec: authentication.TokenRequestSpec{
			ExpirationSeconds: ptr.To[int64](es),
		},
	}
	err = h.Client.SubResource("token").Create(ctx, sa, tr)
	if err != nil {
		return nil, kerrors.NewInternalError(err)
	}

	subjl.ResourceVersion = tr.ResourceVersion
	subjl.Status.Token = tr.Status.Token
	return subjl, nil
}
