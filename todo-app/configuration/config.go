package configuration

import (
	"awesomeProject/todo-app/global_constants"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func (config Config) LogDebugConfigAttributes() {
	log.Debug("Configuration:")
	log.Debug("		port: " + config.Port)
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{"8080"} // set default value for port

	if _, err := os.Stat("./" + global_constants.ConfigFileName); errors.Is(err, os.ErrNotExist) {
		log.Error(global_constants.ConfigFileName + " does not exist")
		return config, nil
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Error reading config file, %s", err)
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Error("Unable to decode config, %v", err)
		return nil, err
	}

	log.Info("Successfully loaded config file")
	config.LogDebugConfigAttributes()
	return config, nil
}
