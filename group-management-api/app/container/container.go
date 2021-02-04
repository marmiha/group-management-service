package container

import (
	"group-management-api/app/config"
	"group-management-api/app/logger"
)

// Handles all the things that need to be shut down.
type Container struct {
	AppConfig *config.AppConfig
	Shutdownables []Shutdownable
}

func New(appConfig *config.AppConfig) *Container{
	return &Container{AppConfig: appConfig}
}

func (c Container) ShutdownAll() {
	for _, s := range c.Shutdownables {
		err := s.Shutdown()

		if err != nil {
			logger.Log.
				WithField("err", err).
				Errorf("Shutdownable %v had an error shutting down. ", s.GetName())
			continue
		}

		logger.Log.Infof("Shutdownable %v shut down successful", s.GetName())
	}
}

func (c Container) AddShutdownable(s Shutdownable) {
	c.Shutdownables = append(c.Shutdownables, s)
}

type Shutdownable interface {
	GetName() string
	Shutdown() error
}