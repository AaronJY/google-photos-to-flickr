package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

// Config represents the app's configuration structure and
// contains instructions on how to parse config values from
// config.yml and the environment (via environment variables)
type Config struct {
	Google struct {
		ClientID     string `yaml:"clientid" envconfig:"GOOGLE_CLIENTID"`
		ClientSecret string `yaml:"clientsecret" envconfig:"GOOGLE_CLIENTSECRET"`
		RedirectURI  string `yaml:"redirecturi" envconfig:"GOOGLE_REDIRECTURI"`
	} `yaml:"google"`
}

// ReadConfig loads config into the given config instance
// from the config.yml file
func ReadConfig(cfg *Config) {
	f, err := os.Open("../config.yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

// ReadConfig loads config into the given config instance
// from environment variables
func ReadEnv(cfg *Config) {
	err := envconfig.Process("", &cfg)
	if err != nil {
		processError(err)
	}
}

// processError processes errors from config.go
func processError(err error) {
	fmt.Errorf(err.Error())
	os.Exit(2)
}
