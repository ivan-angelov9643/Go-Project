package configuration

import (
	"awesomeProject/todo-app/global_constants"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	LogFormat   string `mapstructure:"LOG_FORMAT"`
	LogSeverity string `mapstructure:"LOG_SEVERITY"`
}

func (config Config) SetLogFormat() {
	switch config.LogFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		log.Error("Unrecognized log format: %s", config.LogFormat)
	}
}

func (config Config) SetLogSeverity() {
	switch config.LogSeverity {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.Error("Unrecognized log severity: %s", config.LogSeverity)
	}
}

func (config Config) LogDebugConfigAttributes() {
	log.Debug("Configuration:")
	log.Debug("		port: " + config.Port)
	log.Debug("		log format: " + config.LogFormat)
	log.Debug("		log severity: " + config.LogSeverity)
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{"8080", "text", "debug"} // set default value for port

	if _, err := os.Stat("./" + global_constants.ConfigFileName); errors.Is(err, os.ErrNotExist) {
		log.Error(global_constants.ConfigFileName + " does not exist")
		log.Info("Using default configuration")
	} else {
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
	}

	config.SetLogFormat()
	config.SetLogSeverity()
	config.LogDebugConfigAttributes()

	return config, nil
}
