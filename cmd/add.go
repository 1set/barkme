package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Append item into config",
	Long:    `Verify new device and append into config.`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
