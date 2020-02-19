package conf

import "github.com/BurntSushi/toml"

type Config struct {
	NetAddr string `toml:"NetAddr"`
	Token   string `toml:"Token"`
	Robot   []*Robot
}

type Robot struct {
	Url string `toml:"Url"`
}

var config *Config

func LoadConfig(path string) {
	config = &Config{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}
