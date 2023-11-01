package configmanager

import (
	"encoding/json"
	"errors"
	"os"
)

var Config *Configuration
var configFile *string

type Configuration struct {
	IPFS IPFSConfig `json:"ipfs"`
}

type IPFSConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func GetConfig() Configuration {
	return *Config
}

// LoadConfiguration loads configuration from file
// decodes the json config file into an instance of application config
// if the decoded config is valid it is set as config
func LoadConfiguration(configFile *string) error {
	if configFile == nil {
		return errors.New("config not initialized")
	}
	config := new(Configuration)
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return err
	}
	Config = config
	return nil
}
