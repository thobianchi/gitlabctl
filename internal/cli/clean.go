package cli

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk/clean"
)

func CleanSubcommand(cmd *cobra.Command) {
	cleanCmd := &cobra.Command{
		Use:   "clean",
		Short: "Delete downloaded files",
		Long:  `Delete files created by for example project get-env`,
		Run: func(cmd *cobra.Command, args []string) {
			// Root command does nothing
			clean.Clean()
		},
	}
	cmd.AddCommand(cleanCmd)
}
