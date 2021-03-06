package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfigFromFile(t *testing.T) {
	// SetDefaultConfiguration must be called BEFORE LoadConfigFromFile.
	InitializeConfiguration("docker-config", "./")

	Address := viper.GetString("address")
	if Address == "" {
		t.Error("TestLoadConfigFromFile: Can't get Address!")
	}

	Port := viper.GetInt("port")
	if Port == 0 {
		t.Error("TestLoadConfigFromFile: Can't get Port!")
	}

	ConfigPath := viper.GetString("configPath")
	if ConfigPath == "" {
		t.Error("TestLoadConfigFromFile: Can't get ConfigPath!")
	}

	DirDepth := viper.GetInt("maxDirDepth")
	if DirDepth == 0 {
		t.Error("TestLoadConfigFromFile: Can't get DirDepth!")
	}

	BasePath := viper.GetString("absoluteServePath")
	if BasePath == "" {
		t.Error("TestLoadConfigFromFile: Can't get AbsoluteServePath!")
	}
}
