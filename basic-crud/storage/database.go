package storage

import "github.com/jinzhu/gorm"

// Database Connection
type Database interface {
	Connection() *gorm.DB
	Close()
}
