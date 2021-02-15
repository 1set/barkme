package cmd

import (
	"fmt"
	"strconv"

	"github.com/1set/barkme/bark"
	"github.com/1set/barkme/util"
	"github.com/spf13/cobra"
)

// listRingtoneCmd represents the ringtone command
var listRingtoneCmd = &cobra.Command{
	Use:     "ringtone",
	Aliases: []string{"ring", "tone"},
	Short:   "List available ringtones",
	Long:    `List all available built-in ringtones supported by Bark app.`,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			header    = []string{"Num", "Ringtone"}
			ringtones [][]string
		)
		for i, ring := range bark.AllRingtones {
			ringtones = append(ringtones, []string{
				strconv.Itoa(i + 1),
				string(ring),
			})
		}
		fmt.Println(util.RenderTableString(header, ringtones))
		return nil
	},
}

func init() {
	listCmd.AddCommand(listRingtoneCmd)
}
