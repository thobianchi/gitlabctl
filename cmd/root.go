package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thobianchi/getGitlabEnv/api"
)

var (
	rootCmd = &cobra.Command{
		Use:   "getGitlabEnv",
		Short: "getGitlabEnv is a Gitlab Environment Importer",
		Long: `getGitlabEnv is a Gitlab Environment Importer
  done because pipelines are wonderful but my machine is better.
  set GITLAB_TOKEN environment variable`,
		Run: func(cmd *cobra.Command, args []string) {
			if gitlabToken == "" {
				fmt.Println("GITLAB_TOKEN not set or empty")
				os.Exit(2)
			}
			api.GetEnv(gitlabToken, project, gitlabURL)
		},
	}
	project     string
	gitlabURL   string
	gitlabToken string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	gitlabToken = os.Getenv("GITLAB_TOKEN")
	rootCmd.PersistentFlags().StringVar(&project, "project", "", "gitlab project in form of <group/project>")
	rootCmd.PersistentFlags().StringVar(&gitlabURL, "gitlabURL", "", "complete gitlab url ex: <https://gitlab.com>")
	rootCmd.MarkPersistentFlagRequired("project")
	rootCmd.MarkPersistentFlagRequired("gitlabURL")
}
