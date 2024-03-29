package common

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/system"
	"k8s.io/apimachinery/pkg/types"
)

func GetBasicAuth(ctx context.Context, namespace string, secretRef *walruscore.LocalObjectReference) (*walruscore.BasicAuth, error) {
	loopbackCtrlCli := system.LoopbackCtrlClient.Get()

	var (
		namespacedName = types.NamespacedName{
			Name:      secretRef.Name,
			Namespace: namespace,
		}
		obj = &corev1.Secret{}
	)

	err := loopbackCtrlCli.Get(context.Background(), namespacedName, obj)
	if err != nil {
		return nil, err
	}

	username := obj.StringData["token"]
	if username == "" {
		return nil, fmt.Errorf("basic auth username from %s:%s is empty", namespace, secretRef.Name)
	}

	password := obj.StringData["password"]
	if password == "" {
		return nil, fmt.Errorf("basic auth password from %s:%s is empty", namespace, secretRef.Name)
	}

	return &walruscore.BasicAuth{
		Username:  username,
		Password:  password,
		SecretRef: secretRef,
	}, nil
}

func GetToken(ctx context.Context, namespace string, secretRef *walruscore.LocalObjectReference) (*walruscore.TokenAuth, error) {
	loopbackCtrlCli := system.LoopbackCtrlClient.Get()

	var (
		namespacedName = types.NamespacedName{
			Name:      secretRef.Name,
			Namespace: namespace,
		}
		obj = &corev1.Secret{}
	)
	err := loopbackCtrlCli.Get(ctx, namespacedName, obj)
	if err != nil {
		return nil, err
	}

	token := obj.StringData["token"]
	if token == "" {
		return nil, fmt.Errorf("auth token from %s:%s is empty", namespace, secretRef.Name)
	}

	return &walruscore.TokenAuth{
		Token:     token,
		SecretRef: secretRef,
	}, nil
}
