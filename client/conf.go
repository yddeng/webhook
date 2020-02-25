package client

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	NetAddr   string `toml:"NetAddr"`
	Homepage  string `toml:"Homepage"`
	Branch    string `toml:"Branch"`
	ShellPath string `toml:"ShellPath"`
}

var config *Config

func LoadConfig(path string) {
	config = &Config{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

}

func GetConfig() *Config {
	return config
}
