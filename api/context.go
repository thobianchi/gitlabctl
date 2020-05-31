package api

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFileName string = "/home/thobianchi/.gitlabctl"
)

type context struct {
	Name      string `yaml:""`
	Token     string `yaml:""`
	GitlabURL string `yaml:""`
}

type configFile struct {
	CurrentContext string    `yaml:",omitempty"`
	Contexts       []context `yaml:",omitempty"`
}

func readConfig() configFile {
	cf := configFile{}
	yamlFile, err := ioutil.ReadFile(configFileName)
	if errors.Is(err, os.ErrNotExist) {
		yamlFile = []byte{}
	} else if err != nil {
		log.Fatalf("Failed open file %v: %v", configFileName, err)
	}
	err = yaml.Unmarshal(yamlFile, &cf)
	if err != nil {
		log.Fatalf("Error unmarshaling config yaml: %v", err)
	}
	return cf
}

func writeConfig(cf configFile) {
	d, err := yaml.Marshal(&cf)
	if err != nil {
		log.Fatalf("marshal configFile error: %v", err)
	}
	err = ioutil.WriteFile(configFileName, d, 0644)
	if err != nil {
		log.Fatalf("Error writing config file %v: %v", configFileName, err)
	}
}

func SetContext(name, token, url string) {
	cf := readConfig()
	cf.CurrentContext = name
	newConfig := context{
		Name:      name,
		Token:     token,
		GitlabURL: url,
	}
	for i, ctx := range cf.Contexts {
		if name == ctx.Name {
			cf.Contexts[i] = newConfig
			writeConfig(cf)
			return
		}
	}
	cf.Contexts = append(cf.Contexts, newConfig)
	writeConfig(cf)
}
