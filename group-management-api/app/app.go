package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"group-management-api/app/config"
	"group-management-api/app/config/impl"
	"group-management-api/app/container"
	"group-management-api/app/logger"
)

func InitApp() *container.Container{

	// Read configurations from environment.
	cfg := InitConfig()

	// New application container.
	con := container.New(cfg)

	// Logger initialization.
	InitLogger(con)


	return con
}

func InitConfig() *config.AppConfig{
	var appConfig config.AppConfig
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		logger.Log.
			WithField("err", err).
			Fatal("Unable decode environment variables into struct.")
	}
	return &appConfig
}

func InitLogger(c *container.Container) {
	logLevel := c.AppConfig.LoggerConfig.LogLevel

	var logrusLevel logrus.Level

	switch logLevel {
	case impl.TraceLevel:
		logrusLevel = logrus.TraceLevel
	case impl.DebugLevel:
		logrusLevel = logrus.DebugLevel
	case impl.InfoLevel:
		logrusLevel = logrus.InfoLevel
	case impl.WarnLevel:
		logrusLevel = logrus.WarnLevel
	case impl.ErrorLevel:
		logrusLevel = logrus.ErrorLevel
	case impl.FatalLevel:
		logrusLevel = logrus.FatalLevel
	case impl.PanicLevel:
		logrusLevel = logrus.PanicLevel
	default:
		logger.Log.WithField("LogLevel", logLevel).Panic("Unknown logging level in config")
	}

	logger.Log.SetLevel(logrusLevel)
	logger.Log.WithFields(logrus.Fields{
		"LogLevel": logLevel,
	}).Info("Initialized Logrus Logger")
}