package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `yaml:"app_name" default:"grpc"`
	Port    int32  `yaml:"port" default:"8081"`
	DB      struct {
		Use      string `yaml:"use" default:"postgresql"`
		Postgres []struct {
			Enabled  bool   `yaml:"enabled" default:"true"`
			Host     string `yaml:"host" default:"localhost"`
			Port     string `yaml:"port" default:"5432"`
			UserName string `yaml:"username" default:"youruser"`
			Password string `yaml:"password" default:"yourpassword"`
			Database string `yaml:"database" default:"yourdbname"`
		} `yaml:"postgres"`
	} `yaml:"db"`
}

func (c *Config) NewConfig() (*Config, error) {
	err := configor.Load(c, "config.yaml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
