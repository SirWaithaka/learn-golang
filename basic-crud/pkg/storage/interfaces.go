package storage

import (
	models "github.com/sirwaithaka/learn-golang/basic-crud/pkg/storage/sqlite"
)

// Storage behaviour definition: how a repository should behave
type Storage interface {
	GetDeviceDao() DeviceDao
}

// DeviceDao: should implement these specific methods
type DeviceDao interface {
	Add(device models.Device)
	//Get() models.Device
	GetActive() (models.Device, error)
	GetAll() []models.Device
	Update(device models.Device)
	Upsert(device models.Device)
	Delete(device models.Device)
}
