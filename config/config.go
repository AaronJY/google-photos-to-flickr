package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Google struct {
		ClientID     string `yaml:"clientid" envconfig:"GOOGLE_CLIENTID"`
		ClientSecret string `yaml:"clientsecret" envconfig:"GOOGLE_CLIENTSECRET"`
	} `yaml:"google"`
	Server struct {
		Port int `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
}

func ReadConfig(cfg *Config) {
	if _, err := os.Stat("config.yml"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Before continuing, please take a copy of 'config.template.yml', rename it as 'config.yml' and set your config.")
			os.Exit(2)
		}
	}
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", &cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err.Error())
	os.Exit(2)
}
