package gitlab

import (
	"github.com/thobianchi/gitlabctl/sdk/context"
	"github.com/xanzy/go-gitlab"
)

type IClient interface {
	SearchGroup(groupname string) ([]*gitlab.Group, error)
	ListGroupProjects(groupID int) ([]*gitlab.Project, error)
	ListSubgroups(groupID int) ([]*gitlab.Group, error)
	ListVariables(projectID string) ([]*gitlab.ProjectVariable, error)
}

// Client gitlab client struct
type Client struct {
	client *gitlab.Client
}

// NewClient create gitlab client
func NewClient(ctx *context.Context) (*Client, error) {
	client, err := gitlab.NewClient(ctx.Token, gitlab.WithBaseURL(ctx.GitlabURL))
	if err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

func (c Client) SearchGroup(groupname string) ([]*gitlab.Group, error) {
	groups, _, err := c.client.Groups.SearchGroup(groupname)
	return groups, err
}

func (c Client) ListGroupProjects(groupID int) ([]*gitlab.Project, error) {
	projects, _, err := c.client.Groups.ListGroupProjects(groupID, nil)
	return projects, err
}

func (c Client) ListSubgroups(groupID int) ([]*gitlab.Group, error) {
	groups, _, err := c.client.Groups.ListSubgroups(groupID, nil)
	return groups, err
}

func (c Client) ListVariables(projectID string) ([]*gitlab.ProjectVariable, error){
	variables, _, err := c.client.ProjectVariables.ListVariables(projectID, nil)
	return variables, err
}
