package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk"
)

// getContextCmd represents the getContext command
var useContextCmd = &cobra.Command{
	Use:   "use-context",
	Short: "Set a config defined context to current",
	Long:  `Set the context passed as the one to connect to`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Pass to use-context exaclty one parameter")
		}
		sdk.UseContext(args[0])
	},
}

func init() {
	configCmd.AddCommand(useContextCmd)
}
