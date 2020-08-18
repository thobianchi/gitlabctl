package sdk

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

// GetEnv get environment
func GetEnv(projectID string) {

	git, err := getGitClient()
	if err != nil {
		log.Fatalf("Failed to get git client: %v", err)
	}

	variables, _, err := git.ProjectVariables.ListVariables(projectID, nil)
	if err != nil {
		fmt.Printf("Unable to fetch the variables:  %v\n", err)
		os.Exit(1)
	}

	for _, v := range variables {
		// fmt.Printf("Key: %v, Value: %v, Kind %v\n", v.Key, v.Value, v.VariableType)
		exportEnv(v.Key, v.Value, v.VariableType)
	}
}

func exportEnv(varName, value string, kind gitlab.VariableTypeValue) {
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
