package helm

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"helm.sh/helm/v3/pkg/action"

	"github.com/seal-io/utils/stringx"
	walrusv1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/templates/ensurer/common"
	"github.com/seal-io/walrus/pkg/templates/loader"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Ensurer struct {
}

func New() *Ensurer {
	return &Ensurer{}
}

// Ensure complete the template info and save it.
func (l *Ensurer) Ensure(ctx context.Context, t *walrusv1.Template) error {
	for i := range t.Status.Versions {
		tv, err := genTemplateVersion(ctx, t, t.Status.Versions[i])
		if err != nil {
			return err
		}

		t.Status.Versions[i] = *tv
	}

	return common.UpsertTemplate(ctx, t)
}

// getTemplateVersion retrieves template versions from a helm chart.
func genTemplateVersion(
	ctx context.Context,
	t *walrusv1.Template,
	tv mytypes.TemplateVersion,
) (*mytypes.TemplateVersion, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	var (
		username, password        string
		caFile, certFile, keyFile string
		insecureSkipTLSverify     bool
	)

	if auth := t.Spec.Auth; auth != nil && auth.AuthSecretRef != nil {
		secret, err := loopbackKubeCli.CoreV1().Secrets(t.Namespace).Get(
			ctx,
			auth.AuthSecretRef.Name,
			meta.GetOptions{},
		)
		if err != nil {
			return nil, err
		}

		switch auth.AuthType {
		case mytypes.AuthTypeToken:
			return nil, fmt.Errorf("auth type token not supported when using format helm")
		case mytypes.AuthTypeBasic:
			username = secret.StringData["username"]
			if username == "" {
				return nil, fmt.Errorf("username not found in secret %s", auth.AuthSecretRef.Name)
			}

			password = secret.StringData["password"]
			if password == "" {
				return nil, fmt.Errorf("password not found in secret %s", auth.AuthSecretRef.Name)
			}
		}

		if auth.OCIRegistryAuth != nil {
			// TODO(michelia): oci
		}
	}

	if advance := t.Spec.Advance; advance != nil {
		secret, err := loopbackKubeCli.CoreV1().Secrets(t.Namespace).Get(
			ctx,
			advance.CertSecretRef.Name,
			meta.GetOptions{},
		)
		if err != nil {
			return nil, err
		}

		if ca, ok := secret.Data["ca.crt"]; ok {
			caFile, err = createTempFile(ca, fmt.Sprintf("seal-template-helm-chart-ca-%s", t.Name))
			if err != nil {
				return nil, err
			}
			defer os.Remove(caFile)
		}

		if cert, ok := secret.Data["tls.crt"]; ok {
			certFile, err = createTempFile(cert, fmt.Sprintf("seal-template-helm-chart-cert-%s", t.Name))
			if err != nil {
				return nil, err
			}
			defer os.Remove(certFile)
		}

		if key, ok := secret.Data["tls.key"]; ok {
			keyFile, err = createTempFile(key, fmt.Sprintf("seal-template-helm-chart-key-%s", t.Name))
			if err != nil {
				return nil, err
			}
			defer os.Remove(keyFile)
		}

		insecureSkipTLSverify = advance.SkipTLSVerification
	}

	tempDir := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("seal-template-%s-%s-"+t.Name, stringx.Dasherize(tv.Version), stringx.String(10)))
	defer os.RemoveAll(tempDir)

	pullCli := action.NewPull()
	pullCli.InsecureSkipTLSverify = insecureSkipTLSverify
	pullCli.DestDir = tempDir
	pullCli.Untar = true
	pullCli.PassCredentialsAll = true
	pullCli.Username = username
	pullCli.Password = password
	pullCli.CaFile = caFile
	pullCli.CertFile = certFile
	pullCli.KeyFile = keyFile

	_, err := pullCli.Run(tv.URL)
	if err != nil {
		return nil, err
	}

	sg, err := loader.LoadSchema(pullCli.DestDir, t.Name)
	if err != nil {
		return nil, err
	}

	return common.GenTemplateVersion(ctx, tv.URL, tv.Version, t, sg)
}

func createTempFile(data []byte, pattern string) (string, error) {
	tempCert, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", err
	}

	_, err = tempCert.Write(data)
	if err != nil {
		return "", err
	}

	return tempCert.Name(), nil
}
