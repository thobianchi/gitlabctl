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

	variables, err := client.ListVariables(projectID)
	if err != nil {
		return err
	}

	for _, v := range variables {
		// fmt.Printf("Key: %v, Value: %v, Kind %v\n", v.Key, v.Value, v.VariableType)
		exportEnv(v.Key, v.Value, v.VariableType)
	}

	return nil
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
