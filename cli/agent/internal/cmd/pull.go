package cmd

import "github.com/spf13/cobra"

// -----------------------------------------------------------------------------

// PullCmd is a cobra command builder for "pull" command
func PullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pull",
		Aliases: []string{"p"},
		Short:   "Pull observables using descriptor strategy",
		Run: func(cmd *cobra.Command, args []string) {
			// initialize config
			initConfig()
		},
	}

	// Add flags
	cmd.Flags().String("file", "", "defines agent descriptor file path")

	return cmd
}
