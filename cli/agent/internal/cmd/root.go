package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	iconfig "github.com/iocplatform/agent/cli/agent/internal/config"

	"go.zenithar.org/pkg/config"
	cmdcfg "go.zenithar.org/pkg/config/cmd"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
)

// -----------------------------------------------------------------------------

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "agent",
	Short: "Agent feed processor",
}

func init() {
	mainCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (config.toml)")
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(cmdcfg.NewConfigCommand(conf, "IOCP"))
	mainCmd.AddCommand(serverCmd)
}

// -----------------------------------------------------------------------------

// Execute main command
func Execute() error {
	feature.DefaultMutableGate.AddFlag(mainCmd.Flags())
	return mainCmd.Execute()
}

// -----------------------------------------------------------------------------

var (
	cfgFile string
	conf    = &iconfig.Configuration{}
)

// -----------------------------------------------------------------------------

func initConfig() {
	if err := config.Load(conf, "IOCP", cfgFile); err != nil {
		log.Bg().Fatal("Unable load config", zap.Error(err))
	}
}
