package cmd

import "github.com/spf13/cobra"

// -----------------------------------------------------------------------------

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Starts a service dispatcher",
	PreRun: func(cmd *cobra.Command, args []string) {
		// initialize config
		initConfig()
	},
}
