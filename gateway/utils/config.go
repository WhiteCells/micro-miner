package utils

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Rpc struct {
		user string `mapstructure:"user"`
		farm string `mapstructure:"farm"`
	} `mapstructure:"rpc-server"`

	JWT struct {
		Secret string `mapstructure:"secret"`
		Expire int    `mapstructure:"expire"`
	} `mapstructure:"jwt"`

	Redis redisConfig `mapstructure:"redis"`
}

type redisConfig struct {
	Address []string `mapstructure:"address"`
}

var Config ServerConfig

func InitConfig(configFile string, configType string) error {
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("failed to read unmarshal config: %v", err)
	}

	return nil
}
