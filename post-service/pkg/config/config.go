package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port               string `mapstructure:"port" yaml:"port"`
	JwtSecret          string `mapstructure:"jwtsecret" yaml:"jwtsecret"`
	DbConnectionString string `mapstructure:"dbconnectionstring" yaml:"dbconnectionstring"`
}

var Config *AppConfig

func Read() *AppConfig {
	viper.SetConfigName("config")      // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$PWD/config") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	viper.AddConfigPath("/config")     // optionally look for config in the working directory
	viper.AddConfigPath("./config")    // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}
	Config = &appConfig
	return &appConfig
}
