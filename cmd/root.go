package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gitlabctl",
		Short: "gitlabctl CLI",
		Long: `gitlabctl is a Command line utility to interacto with Gitlab
    is allows you to fetch remote project environemnt, launch pipeline and much more`,
		Run: func(cmd *cobra.Command, args []string) {
			if gitlabToken == "" {
				fmt.Println("GITLAB_TOKEN not set or empty")
				os.Exit(2)
			}
		},
	}
	project     string
	gitlabURL   string
	gitlabToken string
)

// Execute cobra execute CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	gitlabToken = os.Getenv("GITLAB_TOKEN")
}
