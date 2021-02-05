package postgresds

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"group-management-api/service/dataservice/postgresds/modelpg"
)

func NewConnectionAndSchemaCreation(opts *pg.Options, dropTable bool) (*pg.DB, error) {
	db := NewDatabaseConnection(opts)
	err := CreateSchema(db, &orm.CreateTableOptions{
		IfNotExists: true,
	}, dropTable)

	if err != nil{
		// Close the connection as we will not need it.
		defer db.Close()
		return nil, err
	}

	return db, nil
}

func NewDatabaseConnection(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}

// Creates Schema From the models
func CreateSchema(db *pg.DB, options *orm.CreateTableOptions, dropTable bool) error {
	// Our database models/structs.
	models := []interface{} {
		(*modelpg.User)(nil),
		(*modelpg.Group)(nil),
		// More can be added if needed.
	}

	// Create table for each of these models.
	for _, model := range models {
		// Drop tables.
		exists, err := db.Model(model).Exists()

		// Drop the table if it exists.
		if exists {
			err = db.Model(model).DropTable(&orm.DropTableOptions{
				IfExists: dropTable,
				Cascade:  dropTable,
			})

			if err != nil {
				return err
			}
		}

		// Make Tables.
		err = db.Model(model).CreateTable(options)
		if err != nil {
			return err
		}
	}
	return nil
}