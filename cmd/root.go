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
			cmd.Help()
		},
	}
)

// Execute cobra execute CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
