// Configuration struct for all our different pluggable modules.
package config

import (
	"group-management-api/app/config/impl"
)

type AppConfig struct {
	ExposedPort string // At which port the Application is accessible.

	// Logger configuration.
	LoggerConfig struct {
		LogLevel impl.LoggerLevel
	}

	// Containing configuration for DataService modules (Databases or data services).
	DataServiceConfig struct {
		UserDataServiceConfig struct {
			// Chosen UserDataService Implementation.
			Impl impl.DataServiceImpl
		}

		GroupDataServiceConfig struct {
			// Chosen GroupDataService Implementation
			Impl impl.DataServiceImpl
		}

		// DataService Implementations.
		PostgresDataServiceConfig struct {
			PostgresHost     string
			PostgresUser     string
			PostgresPassword string
		}
	}

	// Containing configuration for UseCase modules (Business logic).
	UseCaseConfig struct {
		// At this moment non of our business logic use cases don't need configuration.
	}

	// Containing configuration for Adapter modules (API).
	AdapterConfig struct {
		// Chosen Adapter Implementation Selection.
		Impl impl.AdapterImpl
		// Adapter Implementations.
		RestConfig struct {
			JwtKey string // Used for singing our tokens.
		}
	}
}
