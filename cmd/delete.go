package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "d"},
	Short:   "Delete item from config",
	Long:    `Delete registered devices from config.`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
