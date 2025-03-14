package utils

import (
	"errors"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Rpc struct {
		user string `mapstructure:"user"`
		farm string `mapstructure:"farm"`
	} `mapstructure:"rpc"`

	JWT struct {
		Secret string `mapstructure:"secret"`
		Expire int    `mapstructure:"expire"`
	} `mapstructure:"jwt"`
}

var Config ServerConfig

func InitConfig(configFile string, configType string) error {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("Failed to read config " + err.Error())
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return errors.New("Failed to Unmarshal config " + err.Error())
	}

	return nil
}
