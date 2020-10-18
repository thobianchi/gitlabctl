package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk/project"
)

var proj string
var projectCmd *cobra.Command

func ProjectSubcommand(cmd *cobra.Command) {
	projectCmd = &cobra.Command{
		Use:     "project",
		Aliases: []string{"projects"},
		Short:   "Interact with Gitlab Projects",
		Long:    `Execute commands on Projects, like get remote environment`,
		Run: func(cmd *cobra.Command, args []string) {
			// Root command does nothing
			cmd.Help()
			os.Exit(1)
		},
	}

	projectCmd.PersistentFlags().StringVar(&proj, "id", "", "project id")
	projectCmd.MarkPersistentFlagRequired("project")

	cmd.AddCommand(projectCmd)
}

func ProjectGetEnvSubcommand(cmd *cobra.Command) {
	projectGetEnvCmd := &cobra.Command{
		Use:   "get-env",
		Short: "Get remote project environment",
		Long:  `Fetch remote environment variables and print out export statement`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return project.GetEnv(proj)
		},
	}
	projectCmd.AddCommand(projectGetEnvCmd)
}
