package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk"
)

// getContextCmd represents the getContext command
var currentContextCmd = &cobra.Command{
	Use:   "current-context",
	Short: "Get current defined context",
	Long: `Get currently defined context, all calls will refer to this gitlab server, with the configured
  secret and config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := sdk.CurrentContext()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(currentContextCmd)
}
