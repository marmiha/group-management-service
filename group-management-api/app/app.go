// Package that defines our entry point into microservice execution.
package app

import (
	"github.com/joeshaw/envdecode"
	"github.com/sirupsen/logrus"
	"group-management-api/app/config"
	"group-management-api/app/config/impl"
	"group-management-api/app/container"
	"group-management-api/app/factory"
	"group-management-api/app/logger"
)

func InitApp() *container.Container{
	// Read configurations from environment.
	cfg := InitConfig()
	// New application container.
	c := container.New(cfg)
	// Logger initialization.
	InitLogger(c)
	err := factory.InitAdapter(c)
	if err != nil {
		logger.Log.WithField("err", err).Fatal("Something went wrong initializing the application.")
	}
	return c
}

func InitConfig() *config.AppConfig{
	cfg := config.AppConfig{}

	if err := envdecode.Decode(&cfg); err != nil {
		logger.Log.
			WithField("err", err).
			Fatal("Unable decode environment variables into struct.")
	}

	return &cfg
}

func InitLogger(c *container.Container) {
	// Shut down all shutdownables when calling Log.Fatal().
	logrus.RegisterExitHandler(func() {
		c.ShutdownAll()
	})

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
		logger.Log.WithField("LogLevel", logLevel).Fatal("Unknown logging level in config")
	}

	// Set the logging level provided.
	logger.Log.SetLevel(logrusLevel)
	logger.Log.WithFields(logrus.Fields{
		"LogLevel": logLevel,
	}).Info("Initialized Logrus Logger")
}