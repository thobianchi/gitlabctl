package gitlab

import (
	"github.com/thobianchi/gitlabctl/sdk/context"
	"github.com/xanzy/go-gitlab"
)

type EnvVariable interface {
	String() string
}

type IClient interface {
	SearchGroup(groupname string) ([]*gitlab.Group, error)
	ListGroupProjects(groupID int) ([]*gitlab.Project, error)
	ListSubgroups(groupID int) ([]*gitlab.Group, error)
	GetPrjVariables(projectID string) ([]*gitlab.ProjectVariable, error)
	GetGrpVariables(groupID int) ([]*gitlab.GroupVariable, error)
	GetProject(projectID string) (*gitlab.Project, error)
	GetGroup(groupID int) (*gitlab.Group, error)
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

func (c Client) GetPrjVariables(projectID string) ([]*gitlab.ProjectVariable, error) {
	variables, response, err := c.client.ProjectVariables.ListVariables(projectID, nil)
	if err != nil {
		return nil, err
	}
	for response.NextPage != 0 {
		var newvars []*gitlab.ProjectVariable
		newvars, response, err = c.client.ProjectVariables.ListVariables(projectID, &gitlab.ListProjectVariablesOptions{Page: response.NextPage})
		if err != nil {
			return nil, err
		}
		variables = append(variables, newvars...)
	}
	return variables, err
}

func (c Client) GetGrpVariables(groupID int) ([]*gitlab.GroupVariable, error) {
	variables, response, err := c.client.GroupVariables.ListVariables(groupID, nil)
	if err != nil {
		return nil, err
	}
	for response.NextPage != 0 {
		var newvars []*gitlab.GroupVariable
		newvars, response, err = c.client.GroupVariables.ListVariables(groupID, &gitlab.ListGroupVariablesOptions{Page: response.NextPage})
		if err != nil {
			return nil, err
		}
		variables = append(variables, newvars...)
	}
	return variables, err
}

func (c Client) GetProject(projectID string) (*gitlab.Project, error) {
	prj, _, err := c.client.Projects.GetProject(projectID, nil)
	return prj, err
}

func (c Client) GetGroup(groupID int) (*gitlab.Group, error) {
	grp, _, err := c.client.Groups.GetGroup(groupID, nil)
	return grp, err
}
