package factory

import (
	"errors"
	"group-management-api/app/config/impl"
	"group-management-api/app/container"
	"group-management-api/app/factory/postgresdbfactory"
	"group-management-api/service/dataservice"
	"group-management-api/service/dataservice/postgresds"
)

// Singleton pattern factory for our data service factories.
// These don't need to be thread safe as the constructions will be called on single thread.
// Based on the config parameters we construct specific data service interfaces implementations.

var udds dataservice.UserDataInterface
func GetUserDataService(c *container.Container) (dataservice.UserDataInterface, error){
	if udds == nil {
		// Fetch the specific implementation tag from the config.
		udsi := c.AppConfig.DataServiceConfig.GroupDataServiceConfig.Impl

		switch udsi {
		case impl.DataServicePostgres:
			// Get the needed underlying postgres database connection.
			tempdb, err := postgresdbfactory.GetPostgresDB(c)
			if err != nil {
				return nil, err
			}

			x := &postgresds.UserData{
				DB: tempdb,
			}

			// If all good set IT!
			udds = x
		// Multiple different implementations can be added which will be pinpointed by configuration.
		default:
			return nil, errors.New("unknown group data service implementation " + string(udsi))
		}
	}

	return udds, nil
}

var ugds dataservice.GroupDataInterface
func GetGroupDataService(c *container.Container) (dataservice.GroupDataInterface, error){
	if ugds == nil {
		// Fetch the specific implementation tag from the config.
		gdsi := c.AppConfig.DataServiceConfig.GroupDataServiceConfig.Impl

		switch gdsi {
		// Postgres implementation of dataservice.GroupDataServiceInterface.
		case impl.DataServicePostgres:
			tempdb, err := postgresdbfactory.GetPostgresDB(c)

			if err != nil {
				return nil, err
			}

			// Set the group data interface to postgres implementation.
			ugds = postgresds.GroupData{
				DB: tempdb,
			}
		// Multiple different implementations can be added which will be pinpointed by configuration.
		default:
			return nil, errors.New("unknown group data service implementation " + string(gdsi))
		}
	}
	return ugds, nil
}