package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/api"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete downloaded files",
	Long:  `Delete files created by for example project get-env`,
	Run: func(cmd *cobra.Command, args []string) {
		// Root command does nothing
		api.Clean()
	},
}
