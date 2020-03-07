package sqlite

import (
	"basic-crud/domain/entities"
	"log"
)

// Migrate handles all migrations
func Migrate() {

	database, err := GetDatabase()
	if err != nil {
		print("Database object not created")
		return
	}

	(*database).conn.AutoMigrate(entities.User{})
	log.Println("migrated")
}
