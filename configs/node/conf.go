package node

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	ProxyAddr string `toml:"ProxyAddr"`
	Name      string `toml:"Name"`
	Hooks     []Hook `toml:"Hook"`
}

type Hook struct {
	Homepage  string `toml:"Homepage"`
	Branch    string `toml:"Branch"`
	ShellPath string `toml:"ShellPath"`
}

var config *Config

func LoadConfig(path string) error {
	config = &Config{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		return err
	}
	fmt.Println(config)
	return nil
}

func GetConfig() *Config {
	return config
}
