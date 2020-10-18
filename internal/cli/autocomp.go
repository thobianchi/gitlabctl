package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func AutocompSubcommand(cmd *cobra.Command) {
	var autocompshell string

	autocompCmd := &cobra.Command{
		Use:   "autocomp",
		Short: "Generate autocompletion for zsh, bash, fish. Default: zsh",
		Long:  `Generates autocompletion function for zsh, bash, fish`,
		Run: func(cmd *cobra.Command, args []string) {
			// Root command does nothing
			switch autocompshell {
			case "bash":
				cmd.GenBashCompletion(os.Stdout)
			case "fish":
				cmd.GenFishCompletion(os.Stdout, true)
			default:
				cmd.GenZshCompletion(os.Stdout)
			}
		},
	}
	configCmd.AddCommand(autocompCmd)
	flags := configCmd.PersistentFlags()
	flags.StringVarP(&autocompshell, "shell", "", "zsh", "shell type")
}
