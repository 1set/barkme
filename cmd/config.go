package cmd

import (
	"fmt"

	"github.com/1set/barkme/bark"
	"github.com/1set/gut/ystring"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// FullConfig represents structure of .barkme config file.
type FullConfig struct {
	DefaultName string                  `mapstructure:"default"`
	DeviceMap   map[string]DeviceConfig `mapstructure:"device"`
}

// DeviceConfig represents config for a registered device.
type DeviceConfig struct {
	URL string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

var (
	config FullConfig
)

// GetDefault returns the default bark.Device client.
func (c FullConfig) GetDefault() (*bark.Device, error) {
	if ystring.IsBlank(c.DefaultName) {
		return nil, fmt.Errorf("no default device")
	}
	return c.GetDevice(c.DefaultName)
}

// GetDevice returns the specific bark.Device client by device name.
func (c FullConfig) GetDevice(name string) (*bark.Device, error) {
	var (
		dev DeviceConfig
		ok  bool
	)
	if dev, ok = c.DeviceMap[name]; !ok {
		return nil, fmt.Errorf("missing device: %s", name)
	}
	return bark.New(dev.URL, dev.Key), nil
}

func loadConfig() {
	if err := viper.ReadInConfig(); err == nil {
		log.Debugw("use config file", "path", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalw("fail to load config", zap.Error(err))
		}
	} else {
		log.Debugw("found no config file", zap.Error(err))
	}

	if config.DeviceMap == nil {
		config.DeviceMap = make(map[string]DeviceConfig)
	}
}

func saveConfig() {
	viper.Set("default", config.DefaultName)
	viper.Set("device", config.DeviceMap)

	if err := viper.WriteConfig(); err != nil {
		// retry if the config file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.WriteConfigAs(fallbackConfigFile)
		}

		if err != nil {
			log.Fatalw("fail to save config file", zap.Error(err))
		} else {
			log.Debugw("save as new config file", "path", fallbackConfigFile)
		}
	} else {
		log.Debugw("save as config file", "path", viper.ConfigFileUsed())
	}
}
