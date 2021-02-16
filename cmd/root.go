package cmd

import (
	"errors"
	"fmt"
	"os"

	cl "bitbucket.org/ai69/so-colorful/colorlogo"
	"github.com/1set/barkme/util"
	"github.com/1set/gut/yos"
	"github.com/1set/gut/ystring"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configFile         string
	fallbackConfigFile string

	debugMode bool
	logLevel  string
	logFile   string
	log       *zap.SugaredLogger
)

var (
	errNoConfigDevice = errors.New("no registered devices")
)

var (
	logoArt = `
  ██████╗  █████╗ ██████╗ ██╗  ██╗███╗   ███╗███████╗
  ██╔══██╗██╔══██╗██╔══██╗██║ ██╔╝████╗ ████║██╔════╝
  ██████╔╝███████║██████╔╝█████╔╝ ██╔████╔██║█████╗  
  ██╔══██╗██╔══██║██╔══██╗██╔═██╗ ██║╚██╔╝██║██╔══╝  
  ██████╔╝██║  ██║██║  ██║██║  ██╗██║ ╚═╝ ██║███████╗
  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
`
	colorLogoArt = cl.RoseWaterByColumn(logoArt)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "barkme",
	Short: "Push notifications to iOS devices with Bark",
	Long: colorLogoArt + ystring.NewLine +
		`A simple tool to manage registered iOS devices with Bark app installed,
and push notifications to the iOS devices via Bark server.

- Bark Client: https://github.com/Finb/Bark
- Bark Server: https://github.com/Finb/bark-server`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.barkme.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log", "info", "minimum enabled logging level")
	rootCmd.PersistentFlags().StringVar(&logFile, "logfile", "", "log file path")
	rootCmd.PersistentFlags().BoolVar(&debugMode, "debug", false, "debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Set loggers
	logger := util.NewLogger(logFile, debugMode)
	logger.SetLogLevel(logLevel)
	log = logger.LoggerSugared().With(zap.Int("pid", os.Getpid()))

	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalw("fail to find homedir", zap.Error(err))
		}

		// Create config directory under user's home directory
		dirName := yos.JoinPath(home, ".config", "barkme")
		if err := yos.MakeDir(dirName); err != nil {
			log.Fatalw("fail to create config dir", zap.Error(err))
		}

		// Search config in current directory and config directory (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(dirName)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		// Create if config file doesn't exist
		fallbackConfigFile = yos.JoinPath(dirName, "config.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in, or initialize default one
	loadConfig()
}
