package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(autocompCmd)

	autocompCmd.Flags().StringVar(&autocompshell, "shell", "", "shell type")
}

var autocompshell string = "zsh"

var autocompCmd = &cobra.Command{
	Use:   "autocomp",
	Short: "Generate autocompletion for zsh, bash, fish. Default: zsh",
	Long:  `Generates autocompletion function for zsh, bash, fish`,
	Run: func(cmd *cobra.Command, args []string) {
		// Root command does nothing
		switch autocompshell {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		default:
			rootCmd.GenZshCompletion(os.Stdout)
		}
	},
}
