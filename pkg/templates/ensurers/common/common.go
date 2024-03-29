package common

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/set"

	"github.com/seal-io/utils/json"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/templates/utils"
	"github.com/seal-io/walrus/pkg/types"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func UpsertTemplate(ctx context.Context, t *walruscore.Template) error {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	// Set of version configmaps to be deleted.A
	deleteSet := set.Set[string]{}
	existed, err := loopbackKubeCli.WalruscoreV1().Templates(t.Namespace).Get(ctx, t.Name, meta.GetOptions{})
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return err
		}

		_, err = loopbackKubeCli.WalruscoreV1().Templates(t.Namespace).Create(ctx, t, meta.CreateOptions{})
		if err != nil && !apierrors.IsAlreadyExists(err) {
			return err
		}

		if err == nil {
			return nil
		}
	}

	newSet := set.Set[string]{}
	for _, v := range t.Status.Versions {
		newSet.Insert(v.Version)
	}

	for _, v := range existed.Status.Versions {
		if !newSet.Has(v.Version) {
			deleteSet.Insert(v.Version)
		}
	}

	_, err = loopbackKubeCli.WalruscoreV1().Templates(t.Namespace).UpdateStatus(ctx, t, meta.UpdateOptions{})
	if err != nil {
		return err
	}

	for _, v := range deleteSet.UnsortedList() {
		err = loopbackKubeCli.CoreV1().ConfigMaps(t.Namespace).Delete(ctx, v, meta.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

// GenTemplateVersion generates a template version with schema.
func GenTemplateVersion(
	ctx context.Context,
	url, version string,
	t *walruscore.Template,
	sg *types.SchemaGroup,
) (*walruscore.TemplateVersion, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	uiSchema, err := applyUserEditedUISchema(ctx, *t, version, sg.UISchema)
	if err != nil {
		return nil, err
	}

	uisb, err := json.Marshal(uiSchema)
	if err != nil {
		return nil, err
	}

	sb, err := json.Marshal(sg.Schema)
	if err != nil {
		return nil, err
	}

	var (
		schemaConfigMapName           = utils.NormalizeTemplateVersionConfigmapName(t.Name, version, "schema")
		originalUISchemaConfigMapName = utils.NormalizeTemplateVersionConfigmapName(t.Name, version, "original-ui-schema")
		uiSchemaConfigMapName         = utils.NormalizeTemplateVersionConfigmapName(t.Name, version, "ui-schema")
		schemaData                    = map[string][]byte{"data": sb}
		uiSchemaData                  = map[string][]byte{"data": uisb}
	)

	err = createOrUpdateConfigmap(ctx, loopbackKubeCli, t.Namespace, schemaConfigMapName, schemaData)
	if err != nil {
		return nil, err
	}

	err = createOrUpdateConfigmap(ctx, loopbackKubeCli, t.Namespace, originalUISchemaConfigMapName, uiSchemaData)
	if err != nil {
		return nil, err
	}

	err = createOrUpdateConfigmap(ctx, loopbackKubeCli, t.Namespace, uiSchemaConfigMapName, uiSchemaData)

	// Generate template version.
	tv := &walruscore.TemplateVersion{
		Version: version,
		URL:     url,
		SchemaRef: &walruscore.LocalObjectReference{
			Name: schemaConfigMapName,
		},
		OriginalUISchemaRef: &walruscore.LocalObjectReference{
			Name: originalUISchemaConfigMapName,
		},
		UISchemaRef: &walruscore.LocalObjectReference{
			Name: uiSchemaConfigMapName,
		},
	}

	return tv, nil
}

func createOrUpdateConfigmap(ctx context.Context, loopbackKubeCli clientset.Interface, namespace, name string, data map[string][]byte) error {
	cm, err := loopbackKubeCli.CoreV1().ConfigMaps(namespace).Get(ctx, name, meta.GetOptions{})

	if err != nil {
		if !apierrors.IsNotFound(err) {
			return err
		}

		cm = &corev1.ConfigMap{
			ObjectMeta: meta.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			BinaryData: data,
		}

		_, err = loopbackKubeCli.CoreV1().ConfigMaps(namespace).Create(ctx, cm, meta.CreateOptions{})
		if err == nil {
			return nil
		}

		if !apierrors.IsAlreadyExists(err) {
			return err
		}
	}

	cm.BinaryData = data
	_, err = loopbackKubeCli.CoreV1().ConfigMaps(namespace).Update(ctx, cm, meta.UpdateOptions{})
	return err
}

// applySchemaDefault applies exist user edit schema to ui schema.
func applyUserEditedUISchema(ctx context.Context, t walruscore.Template, version string, s *types.UISchema) (*types.UISchema, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	name := utils.NormalizeTemplateVersionConfigmapName(t.Name, version, "ui-schema")

	cm, err := loopbackKubeCli.CoreV1().ConfigMaps(t.Namespace).Get(ctx, name, meta.GetOptions{})
	if err != nil {
		return nil, err
	}

	b, ok := cm.BinaryData["data"]
	if !ok {
		return nil, fmt.Errorf("could not found data in configmap %s", name)
	}

	var uiSchema types.UISchema
	err = json.Unmarshal(b, &uiSchema)
	if err != nil {
		return nil, err
	}

	if uiSchema.IsUserEdited() {
		return &uiSchema, nil
	}
	return s, nil
}
