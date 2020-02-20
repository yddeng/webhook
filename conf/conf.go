package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	NetAddr string   `toml:"NetAddr"`
	Access  *Access  `toml:"Access"`
	Command []string `toml:"Command"`
	Robot   []*Robot
}

type Access struct {
	AccessIP    []string `toml:"AccessIp"`
	AccessToken string   `toml:"AccessToken"`
}

type Robot struct {
	Name string `toml:"Name"`
	Url  string `toml:"Url"`
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
