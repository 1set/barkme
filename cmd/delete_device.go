package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteDeviceCmd represents the device command
var deleteDeviceCmd = &cobra.Command{
	Use:     "device",
	Aliases: []string{"dev"},
	Short:   "Delete registered devices",
	Long:    `Delete registered devices from config.`,
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(config.DeviceMap) == 0 {
			log.Warnw("found no registered devices in config", "path", viper.ConfigFileUsed())
			return errNoConfigDevice
		}

		for _, n := range args {
			dev := strings.TrimSpace(n)
			if _, found := config.DeviceMap[dev]; found {
				delete(config.DeviceMap, dev)
				if dev == config.DefaultName {
					config.DefaultName = ""
				}
			} else {
				return fmt.Errorf("missing device name: %s", dev)
			}
		}

		saveConfig()
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteDeviceCmd)
}
