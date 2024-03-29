package utils

import (
	"strings"

	"github.com/seal-io/walrus/pkg/types"
)

const (
	// WalrusResourceDefinitionRepositoryTopic indicates the repository stores a resource definition template.
	WalrusResourceDefinitionRepositoryTopic = "walrus-resource-definition"
)

func GenWalrusBuiltinLabels(topics []string) map[string]string {
	labels := make(map[string]string)

	for _, topic := range topics {
		switch {
		case topic == WalrusResourceDefinitionRepositoryTopic:
			labels[types.LabelWalrusResourceDefinition] = "true"
		case strings.HasPrefix(topic, "c-"):
			labels[types.LabelWalrusConnectorType] = strings.TrimPrefix(topic, "c-")
		case strings.HasPrefix(topic, "t-"):
			labels[types.LabelWalrusResourceType] = strings.TrimPrefix(topic, "t-")
		}
	}

	return labels
}
