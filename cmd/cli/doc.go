package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/seal-io/walrus/pkg/cli/doc"
)

func NewDocCmd(root *cobra.Command) *cobra.Command {
	var dir string
	cmd := &cobra.Command{
		Use:    "doc",
		Short:  "Generate CLI documentation",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			err := doc.GenMarkdownTree(root, dir)
			if err != nil {
				fmt.Printf("Failed to generate Markdown documentation: %s\n", err.Error())
				os.Exit(1)
			}

			fmt.Println("Documentation generated successfully!")
		},
	}
	cmd.Flags().StringVar(&dir, "dir", "./docs", "Directory for generated docs")

	return cmd
}
