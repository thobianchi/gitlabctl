package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk"
)

var groupName string
var groupID int
var clonePath string

var backupCmd = &cobra.Command{
	Use: "backup",
	// Aliases: []string{""},
	Short: "Backup recursively one group",
	Long: `Make git clone on every repository in the provided group, maintaining the structure.
A search for the group name ( search in the path also) provided will be made, 
if multiple groups are returned the group ID is needed or modify the search term.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return sdk.Backup(groupName, groupID, clonePath)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.PersistentFlags().StringVar(&groupName, "groupName", "", "Group Name")
	backupCmd.PersistentFlags().IntVar(&groupID, "groupID", -1, "[Optional] Group ID")
	backupCmd.PersistentFlags().StringVar(&clonePath, "clonePath", ".", "[Optional] Clone Path, if not provided group will be cloned in PWD")
	backupCmd.MarkPersistentFlagRequired("groupName")
}
