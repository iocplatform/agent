package cmd

import (
	"fmt"

	"github.com/iocplatform/agent/internal/version"
	"github.com/spf13/cobra"
)

// -----------------------------------------------------------------------------

var displayAsJSON bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display service version",
	Run: func(cmd *cobra.Command, args []string) {
		if displayAsJSON {
			fmt.Printf("%s", version.JSON())
		} else {
			fmt.Printf("%s", version.Full())
		}
	},
}

func init() {
	versionCmd.Flags().BoolVar(&displayAsJSON, "json", false, "Display build info as json")
}
