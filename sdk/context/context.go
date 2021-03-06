package context

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	configFileName string
	confFile       *configFile
)

type Context struct {
	Name      string `yaml:""`
	Token     string `yaml:""`
	GitlabURL string `yaml:""`
}

type configFile struct {
	CurrentContext string    `yaml:",omitempty"`
	Contexts       []Context `yaml:",omitempty"`
}

type homeGetter func() (string, error)
type configReader func() (*configFile, error)

func getConfigFileName(getHomeFunc homeGetter) (string, error) {
	if configFileName == "" {
		cf, err := getHomeFunc()
		if err != nil {
			return "", err
		}
		configFileName = cf + string(os.PathSeparator) + ".gitlabctl"
	}
	return configFileName, nil
}

func getHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}

func getConfig(readCfg configReader) (*configFile, error) {
	if confFile == nil {
		cf, err := readCfg()
		if err != nil {
			return nil, err
		}
		confFile = cf
	}
	return confFile, nil
}

func readConfig() (*configFile, error) {
	cf := configFile{}
	cfName, err := getConfigFileName(getHome)
	if err != nil {
		return nil, err
	}
	yamlFile, err := ioutil.ReadFile(cfName)
	if errors.Is(err, os.ErrNotExist) {
		yamlFile = []byte{}
	} else if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &cf)
	if err != nil {
		return nil, err
	}
	return &cf, nil
}

func writeConfig(cf *configFile) error {
	d, err := yaml.Marshal(&cf)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(configFileName, d, 0644); err != nil {
		return err
	}
	return nil
}

// GetCurrentContext returns context struct obj for accessing connection informations
func GetCurrentContext() (*Context, error) {
	cf, err := getConfig(readConfig)
	if err != nil {
		return nil, err
	}
	cur := cf.CurrentContext
	if cur == "" {
		return nil, errors.New("Current context not set")
	}
	for _, ctx := range cf.Contexts {
		if ctx.Name == cur {
			return &ctx, nil
		}
	}
	return nil, errors.New("Current context not found in configFile")
}

// SetContext main function of set-context command
func SetContext(name, token, url string) error {
	cf, err := getConfig(readConfig)
	if err != nil {
		return err
	}
	cf.CurrentContext = name
	newConfig := Context{
		Name:      name,
		Token:     token,
		GitlabURL: url,
	}
	for i, ctx := range cf.Contexts {
		if name == ctx.Name {
			cf.Contexts[i] = newConfig
			if err := writeConfig(cf); err != nil {
				return err
			}
			return nil
		}
	}
	cf.Contexts = append(cf.Contexts, newConfig)
	if err := writeConfig(cf); err != nil {
		return err
	}
	return nil
}

// UseContext main function of use-context command
func UseContext(name string) error {
	cf, err := getConfig(readConfig)
	if err != nil {
		return err
	}
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
	return nil
}

// GetContexts main function of get-contexts command
func GetContexts() error {
	cf, err := getConfig(readConfig)
	if err != nil {
		return err
	}
	for _, ctx := range cf.Contexts {
		fmt.Println(ctx.Name)
	}
	return nil
}

// CurrentContext main function of current-context command
func CurrentContext() error {
	cf, err := getConfig(readConfig)
	if err != nil {
		return err
	}
	cc := cf.CurrentContext
	if cc != "" {
		fmt.Println(cf.CurrentContext)
	} else {
		fmt.Fprintln(os.Stderr, "current-context is not set")
	}
	return nil // TODO return strings print in another func
}
