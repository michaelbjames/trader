package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config Type.
type Config struct {
	Brokerages struct {
		Analysis struct {
			Initial int `yaml:"initialValue"`
		}
	}
}

var config Config

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	data, err := ioutil.ReadFile("config/base.yaml")
	check(err)
	err = yaml.Unmarshal(data, &config)
	check(err)
}

// Get gives you a reference to our config object.
func Get() *Config {
	return &config
}
