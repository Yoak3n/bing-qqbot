package config

import (
	"Bing-QQBot/model"
	"fmt"
	"github.com/spf13/viper"
)

var MyConfig model.PreConfig

func LoadConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		MyConfig.Bridge = viper.GetString("pre_config.ports.bridge")
		MyConfig.Bot = viper.GetString("pre_config.ports.bot")
		MyConfig.ID = viper.GetString("pre_config.account.id")
		MyConfig.Password = viper.GetString("pre_config.account.password")
	}

}
