package postgresds

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"group-management-api/service/dataservice/postgresds/modelpg"
)

func NewDatabaseConnection(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}

// Creates Schema From the models
func CreateSchema(db *pg.DB, options *orm.CreateTableOptions) error {
	// Our database models/structs.
	models := []interface{} {
		(*modelpg.User)(nil),
		(*modelpg.Group)(nil),
		// More can be added if needed.
	}

	// Create table for each of these models.
	for _, model := range models {
		// Drop tables.
		err := db.Model(model).DropTable(&orm.DropTableOptions{
			IfExists: true,
			Cascade:  true,
		})
		if err != nil {
			return err
		}
		// Make Tables.
		err = db.Model(model).CreateTable(options)
		if err != nil {
			return err
		}
	}
	return nil
}