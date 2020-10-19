package project

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/thobianchi/gitlabctl/internal/gitlab"
	"github.com/thobianchi/gitlabctl/sdk/context"
	gogitlab "github.com/xanzy/go-gitlab"
)

// GetEnv get environment
func GetEnv(projectID string) error {
	ctx, err := context.GetCurrentContext()
	if err != nil {
		return err
	}

	client, err := gitlab.NewClient(ctx)
	if err != nil {
		return err
	}

	variables, err := listVariables(client, projectID)
	if err != nil {
		return err
	}

	for _, variab := range variables {
		switch v := variab.(type) {
		case gogitlab.ProjectVariable:
			exportEnv(v.Key, v.Value, v.VariableType)
		case gogitlab.GroupVariable:
			exportEnv(v.Key, v.Value, v.VariableType)
		}
	}

	return nil
}

func listVariables(gitclient gitlab.IClient, projectID string) ([]gitlab.EnvVariable, error) {
	prj, err := gitclient.GetProject(projectID)
	if err != nil {
		return nil, err
	}

	if prj.Namespace.Kind != "group" {
		return nil, fmt.Errorf("Project parent namespace kind: %s not supported", prj.Namespace.Kind)
	}

	grp, err := gitclient.GetGroup(prj.Namespace.ID)
	if err != nil {
		return nil, err
	}

	grpVars, err := gitclient.GetGrpVariables(prj.Namespace.ID)
	if err != nil {
		return nil, err
	}

	for grp.ParentID != 0 {
		NewgrpVars, err := gitclient.GetGrpVariables(grp.ParentID)
		if err != nil {
			return nil, err
		}
		grpVars = append(grpVars, NewgrpVars...)
		grp, err = gitclient.GetGroup(grp.ParentID)
		if err != nil {
			return nil, err
		}
	}

	prjVars, err := gitclient.GetPrjVariables(projectID)
	if err != nil {
		return nil, err
	}

	var variables []gitlab.EnvVariable
	for _, x := range grpVars {
		variables = append(variables, *x)
	}
	for _, y := range prjVars {
		variables = append(variables, *y)
	}

	return variables, nil
}

func exportEnv(varName, value string, kind gogitlab.VariableTypeValue) {
	switch kind {
	case "env_var":
		exportVar(varName, value)
	case "file":
		exportFileVar(varName, value)
	}
}

func exportVar(varName, value string) {
	fmt.Printf("export %v='%v'\n", varName, value)
}

func exportFileVar(varName, value string) {
	file, err := ioutil.TempFile("/tmp", "GetGit_"+varName[0:8]+"_")
	if err != nil {
		log.Fatalf("Failed to create temp var file: %v", err)
	}
	file.WriteString(value)
	fmt.Printf("export %v='%v'\n", varName, file.Name())
}
