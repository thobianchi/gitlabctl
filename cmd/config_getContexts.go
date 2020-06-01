package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk"
)

// getContextCmd represents the getContext command
var getContextCmd = &cobra.Command{
	Use:     "get-contexts",
	Aliases: []string{"get-context"},
	Short:   "Get all configuration contexts",
	Long:    `Retrieve a list of contexts configured`,
	Run: func(cmd *cobra.Command, args []string) {
		sdk.GetContexts()
	},
}

func init() {
	configCmd.AddCommand(getContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
