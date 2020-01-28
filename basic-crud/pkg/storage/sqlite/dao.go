package sqlite

import (
	"github.com/jinzhu/gorm"
)

type deviceDao struct {
	db *gorm.DB
}

func GetDeviceDao(d *gorm.DB) *deviceDao {
	return &deviceDao{db: d}
}

func (dao deviceDao) Add(device Device) {
	dao.db.Create(&device)
}

func (dao deviceDao) Get(uuid string) Device {
	var device Device
	dao.db.Where("UUID = ?", uuid).First(&device)
	return device
}

func (dao deviceDao) GetActive() (Device, error) {
	var d Device
	dao.db.Where(Device{Active: true}).First(&d)
	return d, nil
}

func (dao deviceDao) GetAll() []Device {
	var devices []Device
	dao.db.Find(&devices)
	return devices
}

func (dao deviceDao) Update(device Device) {
	return
}

func (dao deviceDao) Upsert(device Device) {
	return
}

func (dao deviceDao) Delete(device Device) {
	return
}
