package cmd

import (
	"fmt"

	"github.com/1set/barkme/bark"
	"github.com/1set/gut/ystring"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send notification to devices",
	Long:  `Send custom notification to registered devices.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(config.DeviceMap) == 0 {
			log.Warnw("found no registered devices in config", "path", viper.ConfigFileUsed())
			return errNoConfigDevice
		}

		var (
			err     error
			dev     *bark.Device
			devices []*bark.Device
		)
		if len(args) == 0 {
			// for no arguments, use default device if exists
			log.Debugw("got no args as device name", "default", config.DefaultName)
			if dev, err = config.GetDefault(); err != nil {
				return err
			}
			devices = append(devices, dev)
		} else {
			// handle each argument as device name
			log.Debugw("got args as device name", "count", len(args), "devices", args)
			devSet := make(map[string]bool)
			for _, name := range args {
				if _, found := devSet[name]; found {
					continue
				}
				devSet[name] = true

				if dev, err = config.GetDevice(name); err != nil {
					return err
				}
				devices = append(devices, dev)
			}
		}

		isRingFound := false
		for _, r := range bark.AllRingtones {
			if string(r) == ringtone {
				isRingFound = true
				break
			}
		}
		if !isRingFound {
			return fmt.Errorf("invalid ringtone: %v", ringtone)
		}

		opts := bark.Options{
			Ringtone:     bark.RingtoneName(ringtone),
			OpenURL:      openURL,
			CopyText:     copyText,
			ForceArchive: forceArchive,
			ForceCopy:    forceCopy,
		}

		for i, dev := range devices {
			l := log.With("num", i+1, "device", dev, "option", opts)
			switch {
			case ystring.IsBlank(title) && ystring.IsBlank(body):
				if err := dev.Ping(opts); err != nil {
					l.Warnw("fail to ping", zap.Error(err))
				} else {
					l.Infow("ping device")
				}
			case ystring.IsBlank(title) && ystring.IsNotBlank(body):
				if err := dev.SendShortMessage(body, opts); err != nil {
					l.Warnw("fail to send short message", "body", body, zap.Error(err))
				} else {
					l.Infow("send short message", "body", body)
				}
			case ystring.IsNotBlank(title) && ystring.IsNotBlank(body):
				if err := dev.SendMessage(title, body, opts); err != nil {
					l.Warnw("fail to send message", "title", title, "body", body, zap.Error(err))
				} else {
					l.Infow("send message", "title", title, "body", body)
				}
			default:
				l.Warnw("not supported message", "title", title, "body", body)
			}
		}
		return nil
	},
}

var (
	title        string
	body         string
	ringtone     string
	copyText     string
	openURL      string
	forceArchive bool
	forceCopy    bool
)

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&title, "title", "t", "", "Title of notification")
	sendCmd.Flags().StringVarP(&body, "body", "b", "", "Body of notification")
	sendCmd.Flags().StringVarP(&ringtone, "ringtone", "r", "", "Ringtone of notification")
	sendCmd.Flags().StringVarP(&copyText, "copy", "c", "", "Text to copy")
	sendCmd.Flags().StringVarP(&openURL, "url", "u", "", "URL to open")
	sendCmd.Flags().BoolVar(&forceArchive, "force-archive", false, "Force to archive notification")
	sendCmd.Flags().BoolVar(&forceCopy, "force-copy", false, "Force to copy text")
	//_ = sendCmd.MarkFlagRequired("body")
}
