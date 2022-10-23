package utils

import (
	"github.com/spf13/viper"
)

type ConfigStruct struct {
	PORT              string
	DYNAMO_TABLE_NAME string
}

func LoadConfig(configPath, configName, configType string) *ConfigStruct {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		panic("Need environment variables file")
	}

	return &ConfigStruct{
		PORT:              viper.GetString("PORT"),
		DYNAMO_TABLE_NAME: viper.GetString("DYNAMO_TABLE_NAME"),
	}
}
