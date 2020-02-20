package conf

import "github.com/BurntSushi/toml"

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
