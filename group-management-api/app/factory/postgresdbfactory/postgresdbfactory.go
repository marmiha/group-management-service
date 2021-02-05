package postgresdbfactory

import (
	"github.com/go-pg/pg/v10"
	"group-management-api/app/container"
	"group-management-api/app/logger"
	"group-management-api/service/dataservice/postgresds"
)

var pgdb *pg.DB = nil
func GetPostgresDB(c *container.Container) (*pg.DB, error) {
	if pgdb == nil {
		// Get the config variables.
		postgresUser := c.AppConfig.DataServiceConfig.PostgresDataServiceConfig.PostgresUser
		postgresPassword := c.AppConfig.DataServiceConfig.PostgresDataServiceConfig.PostgresPassword
		postgresHost := c.AppConfig.DataServiceConfig.PostgresDataServiceConfig.PostgresHost
		postgresConnPoolSize := c.AppConfig.DataServiceConfig.PostgresDataServiceConfig.ConnPoolSize
		postgresDropTable := c.AppConfig.DataServiceConfig.PostgresDataServiceConfig.DropTableOnConn

		// Log important information.
		if postgresDropTable {
			logger.Log.Info("Enabled dropping tables on the database")
		}

		// Create a new database connection.
		tempdb, err := postgresds.NewConnectionAndSchemaCreation(&pg.Options{
			User:     postgresUser,
			Password: postgresPassword,
			Addr:     postgresHost,
			PoolSize: postgresConnPoolSize,
		}, postgresDropTable)


		if err != nil {
			return nil, err
		}

		// Set the singleton.
		pgdb = tempdb

		// Closing connection should be added to the shutdownables.
		c.AddShutdownable(PostgresDBClose{
			DB: tempdb,
		})
	}

	return pgdb, nil
}


// For closing the database
type PostgresDBClose struct {
	*pg.DB
}

func (p PostgresDBClose) GetName() string {
	return "PostgresDatabase"
}

func (p PostgresDBClose) Shutdown() error {
	return p.DB.Close()
}
