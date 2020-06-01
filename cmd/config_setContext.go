package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk"
)

// getContextCmd represents the getContext command
var setContextCmd = &cobra.Command{
	Use:   "set-context",
	Short: "Configure a new or update a context",
	Long:  `Configure a context for future use, `,
	Run: func(cmd *cobra.Command, args []string) {
		sdk.SetContext(contextName, token, gitlabURL)
	},
}
var contextName string
var token string
var gitlabURL string

func init() {
	configCmd.AddCommand(setContextCmd)

	setContextCmd.Flags().StringVar(&gitlabURL, "gitlabURL", "", "gitlab connection URL")
	setContextCmd.Flags().StringVar(&contextName, "contextName", "", "name for configuration context")
	setContextCmd.Flags().StringVar(&token, "token", "", "gitlab token")
	setContextCmd.MarkFlagRequired("gitlabURL")
	setContextCmd.MarkFlagRequired("contextName")
	setContextCmd.MarkFlagRequired("token")
}
