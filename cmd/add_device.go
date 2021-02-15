package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/1set/barkme/bark"
	"github.com/1set/gut/ystring"
	"github.com/spf13/cobra"
)

var (
	fullExampleURL   string
	endpointURL      string
	deviceKey        string
	shouldSetDefault bool
	shouldVerify     bool
)

var (
	defaultEndpointURL = "https://api.day.app"
	exampleURLPattern  = regexp.MustCompile(`(?P<url>https?://[/.\-\w]+)/(?P<key>[\w]{20,})/?`)
)

// addDeviceCmd represents the device command
var addDeviceCmd = &cobra.Command{
	Use:     "device",
	Aliases: []string{"dev"},
	Short:   "Add new registered device",
	Long:    `Verify if the given device is registered, and append into config.`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceName := strings.TrimSpace(args[0])
		log.Debugw("attempt to register new device", "name", deviceName, "url", endpointURL, "key", deviceKey)
		if ystring.IsEmpty(deviceName) {
			return errors.New("invalid device name")
		}

		if _, found := config.DeviceMap[deviceName]; found {
			return fmt.Errorf("duplicate device name: %s", deviceName)
		}

		// parse url and key from full example url
		if ystring.IsNotBlank(fullExampleURL) {
			items := reSubMatchMap(exampleURLPattern, fullExampleURL)
			log.Debugw("parse full example url", "raw_url", fullExampleURL, "result", items)
			endpointURL = items["url"]
			deviceKey = items["key"]
		}

		if ystring.IsBlank(endpointURL) {
			return fmt.Errorf("invalid endpoint url: %q", endpointURL)
		}
		if ystring.IsBlank(deviceKey) {
			return fmt.Errorf("invalid device key: %q", deviceKey)
		}

		// send ping to device to verify
		if shouldVerify {
			host, _ := os.Hostname()
			title := "[barkme] Verification"
			body := fmt.Sprintf("Verify registered device from %s at %s", host, time.Now().Format("2006-01-02T15:04:05-0700"))
			device := bark.New(endpointURL, deviceKey)

			if err := device.SendMessage(title, body, bark.Options{
				Ringtone:     bark.RingtoneSilence,
				ForceArchive: false,
			}); err != nil {
				return fmt.Errorf("fail to verify device: %w", err)
			}
		}

		config.DeviceMap[deviceName] = DeviceConfig{
			URL: endpointURL,
			Key: deviceKey,
		}

		if shouldSetDefault || len(config.DeviceMap) <= 1 {
			config.DefaultName = deviceName
		}

		saveConfig()
		log.Infow("add registered device", "name", deviceName, "url", endpointURL, "key", deviceKey, "is_verified", shouldVerify, "is_default", config.DefaultName == deviceName)

		return nil
	},
}

func init() {
	addCmd.AddCommand(addDeviceCmd)

	addDeviceCmd.Flags().BoolVarP(&shouldSetDefault, "default", "d", false, "Set this device as default")
	addDeviceCmd.Flags().BoolVarP(&shouldVerify, "verify", "v", false, "Verify this device is registered first")
	addDeviceCmd.Flags().StringVarP(&endpointURL, "url", "u", defaultEndpointURL, "Endpoint URL of Bark server")
	addDeviceCmd.Flags().StringVarP(&deviceKey, "key", "k", "", "Key of the registered device")
	addDeviceCmd.Flags().StringVarP(&fullExampleURL, "full", "f", "", "Full example URL copied from Bark app")
}

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	if len(match) > 0 {
		for i, name := range r.SubexpNames() {
			if i != 0 && ystring.IsNotEmpty(name) {
				subMatchMap[name] = match[i]
			}
		}
	}
	return subMatchMap
}
