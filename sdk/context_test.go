package sdk

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHome(t *testing.T) {
	t.Run("test getHome return string", func(t *testing.T) {
		home, err := getHome()
		require.NoError(t, err)
		require.Equal(t, reflect.TypeOf(home).Kind(), reflect.String)
	})
	t.Run("test getHome return error on no HOME var", func(t *testing.T) {
		old := os.Getenv("HOME")
		os.Unsetenv("HOME")
		_, err := getHome()
		require.NotNil(t, err)
		os.Setenv("HOME", old)
	})
}

func getHomeMockFail() (string, error) {
	return "", errors.New("Mock Error")
}

func getHomeMock() (string, error) {
	return "/home/test", nil
}

func TestGetConfigFileName(t *testing.T) {
	t.Run("test with filename empty", func(t *testing.T) {
		configFileName = ""
		cfn, err := getConfigFileName(getHomeMock)
		require.NoError(t, err)
		require.Equal(t, cfn, "/home/test/.gitlabctl")
	})
	t.Run("test with filename filled", func(t *testing.T) {
		configFileName = "/home/test/.gitlabctl"
		cfn, err := getConfigFileName(getHomeMock)
		require.NoError(t, err)
		require.Equal(t, cfn, "/home/test/.gitlabctl")
	})
	t.Run("test with failing getHome", func(t *testing.T) {
		configFileName = ""
		cfn, err := getConfigFileName(getHomeMockFail)
		require.Error(t, err)
		require.Empty(t, cfn)
	})
}

var mockContext = context{Name: "Name", Token: "dsajkhkg", GitlabURL: "https://example.com"}
var mockCfg *configFile = &configFile{CurrentContext: "Name", Contexts: []context{mockContext}}

func readCfgMock() (*configFile, error) {
	return mockCfg, nil
}
func readCfgMockFail() (*configFile, error) {
	return nil, errors.New("Failed reafcfg")
}

func TestGetConfig(t *testing.T) {
	t.Run("test with confFile nil", func(t *testing.T) {
		confFile = nil
		cfg, err := getConfig(readCfgMock)
		require.NoError(t, err)
		require.Equal(t, cfg, mockCfg)
	})
	t.Run("test with confFile filled", func(t *testing.T) {
		confFile = mockCfg
		cfg, err := getConfig(readCfgMock)
		require.NoError(t, err)
		require.Equal(t, cfg, mockCfg)
	})
	t.Run("test with failing readCfg", func(t *testing.T) {
		confFile = nil
		cfg, err := getConfig(readCfgMockFail)
		require.Error(t, err)
		require.Nil(t, cfg)
	})
}
