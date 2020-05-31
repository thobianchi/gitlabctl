package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/api"
)

func init() {
	projectCmd.AddCommand(projectGetEnvCmd)
}

var projectGetEnvCmd = &cobra.Command{
	Use:   "get-env",
	Short: "Get remote project environment",
	Long:  `Fetch remote environment variables and print out export statement`,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetEnv(project)
	},
}
