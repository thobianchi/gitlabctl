package backup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/thobianchi/gitlabctl/internal/gitlab"
	"github.com/thobianchi/gitlabctl/sdk/context"
	gogitlab "github.com/xanzy/go-gitlab"
)

var (
	privateSSHKey string = fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
)

func printGroupsFound(groupsFound *[]string) {
	fmt.Println("Found multiple groups:")
	for _, g := range *groupsFound {
		fmt.Println(g)
	}
}

func getGroupBySearch(gitClient gitlab.IClient, groupName string, ID int) (*gogitlab.Group, error) {
	groups, err := gitClient.SearchGroup(groupName)
	if err != nil {
		return nil, err
	}
	if len(groups) == 1 {
		return groups[0], nil
	}

	if len(groups) == 0 {
		return nil, errors.New("no groups for querystring")
	}

	groupsFound := make([]string, 0)
	for _, g := range groups {
		if g.ID == ID {
			return g, nil
		}
		groupsFound = append(groupsFound, fmt.Sprintf("Group Name: %s | FullPath: %s | ID: %d", g.Name, g.FullPath, g.ID))
	}
	printGroupsFound(&groupsFound)

	if ID == -1 {
		return nil, errors.New("multiple groups found and id not specified")
	}
	return nil, errors.New("group with provided ID not found")
}

func createGroupDir(rootDir string, groupFullPath string) error {
	fullPath := path.Join(rootDir, groupFullPath)
	return os.MkdirAll(fullPath, 0755)
}

func cloneRepo(path string, url string) error {
	sshKey, err := ioutil.ReadFile(privateSSHKey)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		return err
	}
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:  url,
		Auth: publicKey,
	})
	return err
}

func traverseGroup(gitClient gitlab.IClient, group *gogitlab.Group, clonePath string) error {
	err := createGroupDir(clonePath, group.FullPath)
	if err != nil {
		return err
	}
	repos, err := gitClient.ListGroupProjects(group.ID)
	if err != nil {
		return err
	}
	for _, repo := range repos {
		repoPath := path.Join(clonePath, group.FullPath, repo.Path)
		fmt.Printf("Cloning %s/%s in %s\n", group.FullPath, repo.Name, repoPath)
		err := cloneRepo(repoPath, repo.SSHURLToRepo)
		if err != nil {
			return err
		}
	}
	subGroups, err := gitClient.ListSubgroups(group.ID)
	if err != nil {
		return err
	}
	for _, subGroup := range subGroups {
		traverseGroup(gitClient, subGroup, clonePath)
	}
	return nil
}
func ensureClonePath(dirName string) error {
	err := os.Mkdir(dirName, os.ModeDir|0755)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

// Backup main function of backup command
func Backup(groupName string, groupID int, clonePath string) error {
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}

	client, err := gitlab.NewClient(ctx)
	if err != nil {
		return err
	}

	group, err := getGroupBySearch(client, groupName, groupID)
	if err != nil {
		return err
	}
	err = ensureClonePath(clonePath)
	if err != nil {
		return err
	}
	return traverseGroup(client, group, clonePath)
}
