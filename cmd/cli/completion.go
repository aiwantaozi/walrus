package main

import (
	"github.com/spf13/cobra"

	"github.com/seal-io/walrus/pkg/cli/api"
)

// SetCompletionAnn set annotation to completion command.
func SetCompletionAnn(root *cobra.Command) {
	root.InitDefaultCompletionCmd()

	compCmdName := "completion"

	for i := range root.Commands() {
		name := root.Commands()[i].Name()

		if name == compCmdName {
			// Completion command.
			if root.Commands()[i].Annotations == nil {
				root.Commands()[i].Annotations = make(map[string]string)
			}
			root.Commands()[i].Annotations[api.AnnResourceName] = compCmdName

			// Completion sub commands.
			for j := range root.Commands()[i].Commands() {
				if root.Commands()[i].Commands()[j].Annotations == nil {
					root.Commands()[i].Commands()[j].Annotations = make(map[string]string)
				}

				root.Commands()[i].Commands()[j].Annotations[api.AnnResourceName] = compCmdName
			}
		}

	}
}
