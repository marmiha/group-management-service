// Package that defines our entry point into microservice execution.
package app

import (
	"errors"
	"github.com/joeshaw/envdecode"
	"github.com/sirupsen/logrus"
	"group-management-api/app/config"
	"group-management-api/app/config/impl"
	"group-management-api/app/container"
	"group-management-api/app/factory"
	"group-management-api/app/logger"
)

func InitApp() (*container.Container, error){
	// Read configurations from environment.
	cfg, err := initConfig()
	if err != nil {
		return nil, err
	}
	// New application container.
	c := container.New(cfg)
	// Logger initialization.
	if err = initLogger(c); err != nil {
		return nil, err
	}

	// Adapter initialization. This is the top most layer.
	if err = factory.InitAdapter(c); err != nil {
		return nil, err
	}

	return c, nil
}

func initConfig() (*config.AppConfig, error){
	cfg := config.AppConfig{}

	if err := envdecode.Decode(&cfg); err != nil {
		logger.Log.
			WithField("err", err).
			Info("Unable decode environment variables into struct.")
		return nil, err
	}

	return &cfg, nil
}

func initLogger(c *container.Container) error {
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
		return errors.New("Unknown logging level in config " + string(logLevel))
	}

	// Set the logging level provided.
	logger.Log.SetLevel(logrusLevel)
	logger.Log.WithFields(logrus.Fields{
		"LogLevel": logLevel,
	}).Info("Initialized Logrus Logger")
	return nil
}