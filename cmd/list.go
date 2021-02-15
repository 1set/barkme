package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List items in config",
	Long:    `List devices in config as table.`,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
