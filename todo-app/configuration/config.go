package configuration

import (
	"awesomeProject/todo-app/global"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	LogFormat   string `mapstructure:"LOG_FORMAT"`
	LogSeverity string `mapstructure:"LOG_SEVERITY"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func (config Config) setLogFormat() error {
	switch config.LogFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		return fmt.Errorf("unrecognized log format: %s", config.LogFormat)
	}
	return nil
}

func (config Config) setLogSeverity() error {
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
		return fmt.Errorf("unrecognized log severity: %s", config.LogSeverity)
	}
	return nil
}

func (config Config) logDebugConfigAttributes() {
	log.Debug("[logDebugConfigAttributes] Configuration:")
	log.Debug("[logDebugConfigAttributes]		port: " + config.Port)
	log.Debug("[logDebugConfigAttributes]		log format: " + config.LogFormat)
	log.Debug("[logDebugConfigAttributes]		log severity: " + config.LogSeverity)
}

func validateDbEnvVars(config *Config) error {
	//log
	field := ""
	if config.DBHost == "" {
		field = "DB_HOST"
	} else if config.DBPort == "" {
		field = "DB_PORT"
	} else if config.DBUser == "" {
		field = "DB_USER"
	} else if config.DBPassword == "" {
		field = "DB_PASSWORD"
	} else if config.DBName == "" {
		field = "DB_NAME"
	}
	if field != "" {
		return fmt.Errorf("[validateDbEnvVars] %s is not present in %s", field, global.ConfigFileName)
	}
	return nil
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{
		Port:        "8080",
		LogFormat:   "text",
		LogSeverity: "debug",
	}
	if _, err := os.Stat("./" + global.ConfigFileName); errors.Is(err, os.ErrNotExist) {
		log.Error("[LoadConfig] " + global.ConfigFileName + " does not exist")
		log.Info("[LoadConfig] Using default configuration")
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigName("config")
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Errorf("[LoadConfig] Error reading config file, %v", err)
			return nil, err
		}

		err = viper.Unmarshal(config)
		if err != nil {
			log.Errorf("[LoadConfig] Unable to decode config, %v", err)
			return nil, err
		}

		err = validateDbEnvVars(config)
		if err != nil {
			log.Errorf("[LoadConfig] Error parsing env vars, %v", err)
			return nil, err
		}

		log.Info("[LoadConfig] Successfully loaded config file")
	}

	err := config.setLogFormat()
	if err != nil {
		log.Errorf("[LoadConfig] Unable to se log format, %v", err)
		return nil, err
	}
	err = config.setLogSeverity()
	if err != nil {
		log.Errorf("[LoadConfig] Unable to set log severity, %v", err)
		return nil, err
	}
	config.logDebugConfigAttributes()

	return config, nil
}
