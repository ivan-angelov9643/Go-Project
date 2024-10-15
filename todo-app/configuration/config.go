package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode config, %v", err)
	}
	return
}
