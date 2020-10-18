package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk/context"
)

var configCmd *cobra.Command

func ConfigSubcommand(cmd *cobra.Command) {
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure environments",
		Long:  `Like kubectl this command allows to create configurations and reuse them`,
		Run: func(cmd *cobra.Command, args []string) {
			// Root command does nothing
			cmd.Help()
			os.Exit(1)
		},
	}
	cmd.AddCommand(configCmd)
}

func ConfigCurrCtxSubcommand(cmd *cobra.Command) {
	currentContextCmd := &cobra.Command{
		Use:   "current-context",
		Short: "Get current defined context",
		Long: `Get currently defined context, all calls will refer to this gitlab server, with the configured
		secret and config`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return context.CurrentContext()
		},
	}
	configCmd.AddCommand(currentContextCmd)
}

func ConfigGetCtxSubcommand(cmd *cobra.Command) {
	getContextCmd := &cobra.Command{
		Use:     "get-contexts",
		Aliases: []string{"get-context"},
		Short:   "Get all configuration contexts",
		Long:    `Retrieve a list of contexts configured`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return context.GetContexts()
		},
	}
	configCmd.AddCommand(getContextCmd)
}

func ConfigSetCtxSubcommand(cmd *cobra.Command) {
	var contextName, token, gitlabURL string

	setContextCmd := &cobra.Command{
		Use:   "set-context",
		Short: "Configure a new or update a context",
		Long:  `Configure a context for future use, `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return context.SetContext(contextName, token, gitlabURL)
		},
	}
	configCmd.AddCommand(setContextCmd)

	flags := configCmd.PersistentFlags()
	flags.StringVarP(&gitlabURL, "gitlabURL", "", "", "gitlab connection URL")
	flags.StringVarP(&contextName, "contextName", "", "", "name for configuration context")
	flags.StringVarP(&token, "token", "", "", "gitlab token")
	setContextCmd.MarkFlagRequired("gitlabURL")
	setContextCmd.MarkFlagRequired("contextName")
	setContextCmd.MarkFlagRequired("token")
}

func ConfigUseCtxSubcommand(cmd *cobra.Command) {
	useContextCmd := &cobra.Command{
		Use:   "use-context",
		Short: "Set a config defined context to current",
		Long:  `Set the context passed as the one to connect to`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("Pass to use-context exaclty one parameter")
			}
			return context.UseContext(args[0])
		},
	}
	configCmd.AddCommand(useContextCmd)
}
