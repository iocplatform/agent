package cmd

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/iocplatform/agent/cli/agent/internal/descriptor"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// -----------------------------------------------------------------------------

var descriptorFile string

// PullCmd is a cobra command builder for "pull" command
func PullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pull",
		Aliases: []string{"p"},
		Short:   "Pull observables using descriptor strategy",
		Run:     pullRun,
	}

	// Add flags
	cmd.Flags().StringVar(&descriptorFile, "file", "", "defines agent descriptor file path")

	return cmd
}

func pullRun(cmd *cobra.Command, args []string) {
	// initialize config
	initConfig()

	// Initialize application context
	ctx := context.Background()

	// Loading agent definition
	configFileName, _ := filepath.Abs(descriptorFile)
	yamlFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}

	// Deserialize descriptor
	var agentConfig descriptor.Agent
	err = yaml.Unmarshal(yamlFile, &agentConfig)
	if err != nil {
		panic(err)
	}

	// Build agent from descriptor
	agent, err := descriptor.Build(agentConfig)
	if err != nil {
		panic(err)
	}

	if err := descriptor.Run(ctx, agent); err != nil {
		panic(err)
	}
}
