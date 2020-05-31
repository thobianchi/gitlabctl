package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(projectCmd)

	projectCmd.PersistentFlags().StringVar(&project, "id", "", "project id")
	projectCmd.MarkPersistentFlagRequired("project")
}

var project string

var projectCmd = &cobra.Command{
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
