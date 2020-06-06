package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk"
)

// getContextCmd represents the getContext command
var useContextCmd = &cobra.Command{
	Use:   "use-context",
	Short: "Set a config defined context to current",
	Long:  `Set the context passed as the one to connect to`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Pass to use-context exaclty one parameter")
		}
		err := sdk.UseContext(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(useContextCmd)
}
