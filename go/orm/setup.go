// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"

	"gongd3/go/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(stage *models.StageStruct, logMode bool, filepath string) *gorm.DB {
	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gongd3_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filepath), gormConfig)

	if err != nil {
		panic("Failed to connect to database!")
	}

	AutoMigrate(stage, db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(stage *models.StageStruct, db *gorm.DB) {
	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "gongd3_go_", // table name prefix
	}

	err := db.AutoMigrate( // insertion point for reference to structs
		&BarDB{},
		&KeyDB{},
		&PieDB{},
		&ScatterDB{},
		&SerieDB{},
		&ValueDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package gongd3/go")
	}
	// log.Printf("Database Migration of package gongd3/go is OK")

	BackRepo.init(stage, db)
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&BarDB{})
	db.Delete(&KeyDB{})
	db.Delete(&PieDB{})
	db.Delete(&ScatterDB{})
	db.Delete(&SerieDB{})
	db.Delete(&ValueDB{})
}
