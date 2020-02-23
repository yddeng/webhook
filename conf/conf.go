package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	NetAddr     string   `toml:"NetAddr"`
	AccessIP    []string `toml:"AccessIp"`
	AccessToken string   `toml:"AccessToken"`
	Robot       []*Robot
}

type Robot struct {
	RobotType string   `toml:"RobotType"`
	Homepage  string   `toml:"Homepage"`
	RobotUrl  string   `toml:"RobotUrl"`
	NotifyCmd []string `toml:"NotifyCmd"`
}

var config *Config

func LoadConfig(path string) {
	config = &Config{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	for _, r := range config.Robot {
		fmt.Println("Robot", r)
	}
}

func GetConfig() *Config {
	return config
}
