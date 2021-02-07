// Configuration struct for all our different pluggable modules.
package config

import (
	"group-management-api/app/config/impl"
)

type AppConfig struct {
	// At which port the Application is accessible.
	ExposedPort string `env:"ExposedPort"`

	// Logger configuration.
	LoggerConfig struct {
		LogLevel impl.LoggerLevel `env:"LogLevel"`
	}

	// Containing configuration for DataService modules (Databases or data services).
	DataServiceConfig struct {
		UserDataServiceConfig struct {
			// Chosen UserDataService Implementation.
			Impl impl.DataServiceImpl `env:"DataService.User.Impl"`
		}

		GroupDataServiceConfig struct {
			// Chosen GroupDataService Implementation
			Impl impl.DataServiceImpl `env:"DataService.Group.Impl"`
		}

		// DataService Implementations.
		PostgresDataServiceConfig struct {
			PostgresHost     string `env:"DataService.Postgres.Host"`
			PostgresUser     string `env:"DataService.Postgres.User"`
			PostgresPassword string `env:"DataService.Postgres.Password"`
			ConnPoolSize		int `env:"DataService.Postgres.ConnPoolSize"`
			DropTableOnConn	bool `env:"Dataservice.Postgres.DropTableOnConn"`
		}
	}

	// Containing configuration for UseCase modules (Business logic).
	UseCaseConfig struct {
		// At this moment non of our business logic use cases don't need configuration.
	}

	// Containing configuration for Adapter modules (API).
	AdapterConfig struct {
		// Chosen Adapter Implementation Selection.
		Impl impl.AdapterImpl `env:"Adapter.Impl"`

		// Adapter Implementations.
		RestConfig struct {
			// Used for singing our tokens.
			JwtKey string `env:"Adapter.Rest.JwtKey"`
		}
	}
}
