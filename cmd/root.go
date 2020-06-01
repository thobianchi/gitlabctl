package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	customCompletionFunction = `
    __get_ctx_sugg() {
      local gitlabctl_get_ctx_out
      if gitlabctl_get_ctx_out=$(gitlabctl config get-contexts 2>/dev/null); then
        COMPREPLY=( $( compgen -W "${gitlabctl_get_ctx_out[*]}" -- "$cur" ) )
      fi
    }
    __custom_func() {
      case ${last_command} in
        gitlabctl_config_use-context)
            __get_ctx_sugg
              return
              ;;
          *)
              ;;
      esac
  }
  `
)

var (
	rootCmd = &cobra.Command{
		Use:   "gitlabctl",
		Short: "gitlabctl CLI",
		Long: `gitlabctl is a Command line utility to interacto with Gitlab
    is allows you to fetch remote project environemnt, launch pipeline and much more`,
		BashCompletionFunction: customCompletionFunction,
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
