package cmd

import (
	"fmt"
	"sort"

	"github.com/1set/barkme/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listDeviceCmd represents the device command
var listDeviceCmd = &cobra.Command{
	Use:     "device",
	Aliases: []string{"dev"},
	Short:   "List registered devices",
	Long:    `List all registered devices in config and render as table.`,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(config.DeviceMap) == 0 {
			log.Warnw("found no registered devices in config", "path", viper.ConfigFileUsed())
			return errNoConfigDevice
		}

		var (
			header  = []string{"Name", "URL", "Key", "Default"}
			devices [][]string
		)
		for name, dev := range config.DeviceMap {
			var mark string
			if name == config.DefaultName {
				mark = checkMark
			}
			devices = append(devices, []string{
				name,
				dev.URL,
				dev.Key,
				mark,
			})
		}
		sort.SliceStable(devices, func(i, j int) bool {
			return devices[i][0] < devices[j][0]
		})

		fmt.Println(util.RenderTableString(header, devices))
		return nil
	},
}

func init() {
	listCmd.AddCommand(listDeviceCmd)
}
