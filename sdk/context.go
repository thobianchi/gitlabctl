package sdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	configFileName string      = getHome() + string(os.PathSeparator) + ".gitlabctl"
	confFile       *configFile = nil
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

func getHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Can't find user home directory:", err)
	}
	return home
}

func getConfig() *configFile {
	if confFile == nil {
		confFile = readConfig()
	}
	return confFile
}

func readConfig() *configFile {
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
	return &cf
}

func writeConfig(cf *configFile) {
	d, err := yaml.Marshal(&cf)
	if err != nil {
		log.Fatalf("marshal configFile error: %v", err)
	}
	if err := ioutil.WriteFile(configFileName, d, 0644); err != nil {
		log.Fatalf("Error writing config file %v: %v", configFileName, err)
	}
}

func getCurrentContext() (context, error) {
	cf := getConfig()
	cur := cf.CurrentContext
	if cur == "" {
		return context{}, errors.New("Current context not set")
	}
	for _, ctx := range cf.Contexts {
		if ctx.Name == cur {
			return ctx, nil
		}
	}
	return context{}, errors.New("Current context not found in configFile")
}

func SetContext(name, token, url string) {
	cf := getConfig()
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

func UseContext(name string) {
	cf := getConfig()
	var found bool = false
	for _, ctx := range cf.Contexts {
		if ctx.Name == name {
			found = true
		}
	}
	if found {
		cf.CurrentContext = name
		writeConfig(cf)
	} else {
		fmt.Fprintf(os.Stderr, "Context %v: not found in config file\n", name)
	}
}

func GetContexts() {
	cf := getConfig()
	for _, ctx := range cf.Contexts {
		fmt.Println(ctx.Name)
	}
}

func CurrentContext() {
	cf := getConfig()
	cc := cf.CurrentContext
	if cc != "" {
		fmt.Println(cf.CurrentContext)
	} else {
		fmt.Fprintln(os.Stderr, "current-context is not set")
	}
}
