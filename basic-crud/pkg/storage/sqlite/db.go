package sqlite

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // import sqlite
)

const (
	dbName = "youtise.db"
)

var db *database
var once sync.Once

// database object for the db
type database struct {
	conn *gorm.DB
}

// GetDatabase creates a new Database object
func GetDatabase() (*database, error) {
	var err error

	once.Do(func() {
		db = new(database)
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Dir(filename)
		var p string = filepath.Join(dir, dbName)

		db.conn, err = gorm.Open("sqlite3", p)
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// return the database connection property
func (d *database) Connection() *gorm.DB {
	return d.conn
}

// Close the Database connection
func (d *database) Close() {
	err := d.conn.Close()
	if err != nil {
		fmt.Printf("Error closing db: %v\n", err)
	}
}
