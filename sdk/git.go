package sdk

import (
	"github.com/xanzy/go-gitlab"
)

var (
	gitClient *gitlab.Client
)

func getGitClient() (*gitlab.Client, error) {
	if gitClient == nil {
		ctx, err := getCurrentContext()
		if err != nil {
			return nil, err
		}
		return gitlab.NewClient(ctx.Token, gitlab.WithBaseURL(ctx.GitlabURL))
	}
	return gitClient, nil
}
