package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database db `mapstructure:"db"`
}

type db struct {
	Port    string `mapstructure:"port"`
	Address string `mapstructure:"address"`
	DBName  string `mapstructure:"db_name"`
	User    string `mapstructure:"user"`
	Passwd  string `mapstructure:"passwd"`
}

func GetConfig() (config Config, err error) {
	viper.SetConfigFile("config.toml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return config, err
}
