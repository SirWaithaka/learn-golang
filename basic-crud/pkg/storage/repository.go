package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/storage/sqlite"
)

type repository struct {
	db *gorm.DB
}

func GetStorageRepository(db *gorm.DB) Storage {
	return &repository{ db: db }
}

// GetDeviceDao returns the concrete implementation of the
// Data Access Object(DAO) used for the specific database.
func (r *repository) GetDeviceDao() DeviceDao {
	return sqlite.GetDeviceDao((*r).db)
}