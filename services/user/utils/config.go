package utils

import (
	"errors"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Mode string `mapstructure:"mode"`
	} `mapstructure:"server"`

	MySQL struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DBName       string `mapstructure:"dbname"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
	} `mapstructure:"mysql"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
	} `mapstructure:"redis"`

	JWT struct {
		Secret string `mapstructure:"secret"`
		Expire int    `mapstructure:"expire"`
	} `mapstructure:"jwt"`

	Log struct {
		Level      string `mapstructure:"level"`
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxAge     int    `mapstructure:"max_age"`
		MaxBackups int    `mapstructure:"max_backups"`
	} `mapstructure:"log"`

	Mnemonic struct {
		Key  string `mapstructure:"key"`
		Path string `mapstructure:"path"`
	} `mapstructure:"mnemonic"`

	Bsc struct {
		Api string `mapstructure:"api"`
	} `mapstructure:"bsc"`
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
