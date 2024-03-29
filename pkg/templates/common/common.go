package common

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/seal-io/walrus/pkg/system"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getSecret(ctx context.Context, namespace, name string) (*corev1.Secret, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	return loopbackKubeCli.CoreV1().Secrets(namespace).Get(ctx, name, meta.GetOptions{})
}
