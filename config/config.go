package config

import (
	"github.com/BurntSushi/toml"
)

var (
	config *Config
)

type DbConfig struct {
	Engine   string
	Host     string
	Port     uint
	Database string
	Username string
	Password string
	Charset  string
}

type AppConfig struct {
	Port       uint
	Debug      bool
	AllowHost  string
	AppName    string
	Owner      string
	AdminEmail []string
	Charset    string
}

type Config struct {
	DB  DbConfig
	App AppConfig
}

func Get() *Config {

	if config != nil {
		return config
	}
	var c *Config
	if _, err := toml.DecodeFile("./config/conf.toml", &c); err != nil {
		panic(err)
	}
	return config
}
