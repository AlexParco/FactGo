package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Database string `yaml:"database,omitempty"`
	Server   string `yaml:"server,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Password string `yaml:"password,omitempty"`
	User     string `yaml:"user,omitempty"`
}

func ParseConf(path string) *Conf {
	var config Conf

	fb, err := os.ReadFile(path)
	if err != nil {
		log.Printf("ReadFile: %v", err)
		panic(err)
	}

	err = yaml.Unmarshal(fb, &config)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		panic(err)
	}

	return &config
}
