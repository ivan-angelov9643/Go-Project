package server

import (
	"awesomeProject/library-app/global"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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

	AuthURL          string `mapstructure:"AUTH_URL"`
	AuthRealm        string `mapstructure:"AUTH_REALM"`
	AuthClientID     string `mapstructure:"AUTH_CLIENT_ID"`
	AuthClientSecret string `mapstructure:"AUTH_CLIENT_SECRET"`
}

func NewConfig() *Config {
	return &Config{
		Port:        global.DefaultPort,
		LogFormat:   global.DefaultLogFormat,
		LogSeverity: global.DefaultLogSeverity,
	}
}

func (config *Config) setLogFormat() error {
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

func (config *Config) setLogSeverity() error {
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

func (config *Config) logDebugConfigAttributes() {
	log.Debug("[Config.logDebugConfigAttributes] Configuration:")
	log.Debug("[Config.logDebugConfigAttributes]		port: " + config.Port)
	log.Debug("[Config.logDebugConfigAttributes]		log format: " + config.LogFormat)
	log.Debug("[Config.logDebugConfigAttributes]		log severity: " + config.LogSeverity)
}

func (config *Config) validateDBEnvVars() error {
	log.Info("[Config.validateDBEnvVars] Validating database environment variables...")
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
		log.Infof("[Config.validateDBEnvVars] %s is not present in configuraion file", field)
		return fmt.Errorf("[Config.validateDBEnvVars] %s is not present in configuraion file", field)
	}
	log.Info("[Config.validateDBEnvVars] Successfully validated database environment variables")
	return nil
}

func (config *Config) Load(path string) error {

	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		log.Error("[Config.Load] " + path + " does not exist")
		log.Info("[Config.Load] Using default configuration")
	} else {
		viper.SetConfigName(filepath.Base(path))
		viper.SetConfigType("env")
		viper.AddConfigPath(filepath.Dir(path))
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Errorf("[Config.Load] Error reading config file, %v", err)
			return err
		}

		err = viper.Unmarshal(config)
		if err != nil {
			log.Errorf("[Config.Load] Unable to decode config, %v", err)
			return err
		}

		err = config.validateDBEnvVars()
		if err != nil {
			log.Errorf("[Config.Load] Error parsing env vars, %v", err)
			return err
		}

		log.Info("[Config.Load] Successfully loaded config file")
	}

	err = config.setLogFormat()
	if err != nil {
		log.Errorf("[Config.Load] Unable to se log format, %v", err)
		return err
	}
	err = config.setLogSeverity()
	if err != nil {
		log.Errorf("[Config.Load] Unable to set log severity, %v", err)
		return err
	}
	config.logDebugConfigAttributes()

	return nil
}

func (server *Server) InitializeConfig(path string) {
	server.Config = NewConfig()
	err := server.Config.Load(path)
	if err != nil {
		log.Fatal("[Server.InitializeConfig] Cannot load configuration: ", err)
	}
}
