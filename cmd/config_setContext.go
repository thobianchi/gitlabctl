/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/api"
)

// getContextCmd represents the getContext command
var setContextCmd = &cobra.Command{
	Use:   "set-context",
	Short: "Configure a new or update a context",
	Long:  `Configure a context for future use, `,
	Run: func(cmd *cobra.Command, args []string) {
		api.SetContext(contextName, token, gitlabURL)
	},
}
var contextName string
var token string

// var gitlabURL string

func init() {
	configCmd.AddCommand(setContextCmd)

	setContextCmd.Flags().StringVar(&gitlabURL, "gitlabURL", "", "gitlab connection URL")
	setContextCmd.Flags().StringVar(&contextName, "contextName", "", "name for configuration context")
	setContextCmd.Flags().StringVar(&token, "token", "", "gitlab token")
	setContextCmd.MarkFlagRequired("gitlabURL")
	setContextCmd.MarkFlagRequired("contextName")
	setContextCmd.MarkFlagRequired("token")
}
