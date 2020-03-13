package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	HookAddr    string   `toml:"HookAddr"`
	TcpAddr     string   `toml:"TcpAddr"`
	AccessIP    []string `toml:"AccessIp"`
	AccessToken string   `toml:"AccessToken"`
	Robots      []Robot  `toml:"Robot"`
}

type Robot struct {
	RobotType string   `toml:"RobotType"`
	Homepage  string   `toml:"Homepage"`
	RobotUrl  string   `toml:"RobotUrl"`
	NotifyCmd []string `toml:"NotifyCmd"`
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
