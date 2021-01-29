package postgresds

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"group-management-api/service/dataservice/postgresds/pgmodel"
)

func NewDatabaseConnection(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}

// Creates Schema From the models
func CreateSchema(db *pg.DB, options *orm.CreateTableOptions) error {
	// Our database models/structs.
	models := []interface{} {
		(*pgmodel.User)(nil),
		(*pgmodel.Group)(nil),
		// More can be added if needed.
	}

	// Create table for each of these models.
	for _, model := range models {
		err := db.Model(model).CreateTable(options)
		if err != nil {
			return err
		}
	}
	return nil
}