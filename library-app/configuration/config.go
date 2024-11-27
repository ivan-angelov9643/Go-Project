package configuration

import (
	"awesomeProject/library-app/global"
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

	POSTGRESHost     string `mapstructure:"POSTGRES_HOST"`
	POSTGRESPort     string `mapstructure:"POSTGRES_PORT"`
	POSTGRESUser     string `mapstructure:"POSTGRES_USER"`
	POSTGRESPassword string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRESDb       string `mapstructure:"POSTGRES_DB"`
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
	case "db_error":
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
	if config.POSTGRESHost == "" {
		field = "POSTGRES_HOST"
	} else if config.Port == "" {
		field = "POSTGRES_PORT"
	} else if config.POSTGRESUser == "" {
		field = "POSTGRES_USER"
	} else if config.POSTGRESPassword == "" {
		field = "POSTGRES_PASSWORD"
	} else if config.POSTGRESDb == "" {
		field = "POSTGRES_DB"
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

	_, err := os.Stat(path + "/" + global.ConfigFileName)
	if errors.Is(err, os.ErrNotExist) {
		log.Error("[LoadConfig] " + global.ConfigFileName + " does not exist")
		log.Info("[LoadConfig] Using default configuration")
	} else {
		viper.SetConfigName(global.ConfigFileName)
		viper.SetConfigType("env")
		viper.AddConfigPath(path)
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

	err = config.setLogFormat()
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
