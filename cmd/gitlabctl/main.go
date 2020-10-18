package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/internal/cli"
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

func main() {
	rootCmd := &cobra.Command{
		Use:   "gitlabctl",
		Short: "gitlabctl CLI",
		Long: `gitlabctl is a Command line utility to interacto with Gitlab
    is allows you to fetch remote project environemnt, launch pipeline and much more`,
		BashCompletionFunction: customCompletionFunction,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
  }

  // cli.AddGlobalFlags(rootCmd)
  cli.ProjectSubcommand(rootCmd)
	cli.ProjectGetEnvSubcommand(rootCmd)
	cli.ConfigSubcommand(rootCmd)
	cli.ConfigCurrCtxSubcommand(rootCmd)
	cli.ConfigGetCtxSubcommand(rootCmd)
	cli.ConfigSetCtxSubcommand(rootCmd)
	cli.ConfigUseCtxSubcommand(rootCmd)
	cli.BackupSubcommand(rootCmd)
	cli.CleanSubcommand(rootCmd)
	cli.AutocompSubcommand(rootCmd)
	cli.VersionSubcommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
