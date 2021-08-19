package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generates spartan components",
	Long:    "Generates diferent spartan components, like resources, migrations, handlers",
	Example: `Can use:
	spartan generate
	or
	spartan -g`,
	Aliases: []string{"g"},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
