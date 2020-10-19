package cli

import (
	"github.com/spf13/cobra"
	"github.com/thobianchi/gitlabctl/sdk/backup"
)

func BackupSubcommand(cmd *cobra.Command) {
	var groupName, clonePath string
	var groupID int

	backupCmd := &cobra.Command{
		Use: "backup",
		// Aliases: []string{""},
		Short: "Backup recursively one group",
		Long: `Make git clone on every repository in the provided group, maintaining the structure.
	A search for the group name ( search in the path also) provided will be made,
	if multiple groups are returned the group ID is needed or modify the search term.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return backup.Backup(groupName, groupID, clonePath)
		},
	}
	cmd.AddCommand(backupCmd)
	flags := backupCmd.PersistentFlags()
	flags.StringVarP(&groupName, "groupName", "", "", "Group Name")
	flags.StringVarP(&clonePath, "clonePath", "", ".", "[Optional] Clone Path, if not provided group will be cloned in PWD")
	flags.IntVarP(&groupID, "groupID", "", -1, "[Optional] Group ID")
	configCmd.MarkPersistentFlagRequired("groupName")
}
