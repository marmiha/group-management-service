// Contains everything needed for our application buildup and starting and shutting down the application.
package container

import (
	"group-management-api/app/config"
	"group-management-api/app/logger"
	"os"
	"os/signal"
	"syscall"
)

// Handles all the things that need to be shut down.
type Container struct {
	Startable Startable
	AppConfig *config.AppConfig
	Shutdownables []Shutdownable
}

func New(appConfig *config.AppConfig) *Container{
	return &Container{AppConfig: appConfig}
}

func (c Container) ShutdownAll() {
	logger.Log.Info("Shutting down Application Container Shutdownables...")
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
	logger.Log.Info("... completed.")
}
func (c *Container) StartApp() {
	if c.Startable == nil {
		logger.Log.Info("Nothing was started :,(. Something was not configured right.")
		return
	}

	// In case of SIG-TERM we will be notified.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// Run thread for shutting the container application shutdownables down.
	go func() {
		<-signals
		c.ShutdownAll()
	}()

	c.Startable()
}

func (c Container) AddShutdownable(s Shutdownable) {
	c.Shutdownables = append(c.Shutdownables, s)
}

// An operation that will perform on application shutdown.
type Shutdownable interface {
	GetName() string
	Shutdown() error
}

// An operation that will perform on application start.
type Startable func()