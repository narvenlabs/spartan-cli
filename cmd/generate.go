package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates components",
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
