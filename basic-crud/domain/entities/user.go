package entities

import (
	"basic-crud/storage"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Repository interface {
	Add([]User) error
	GetAll() []User
	Delete()
}

type repository struct {
	db storage.Database
}

func NewRepository(db storage.Database) Repository {
	return &repository{db}
}

func (repo repository) DB() *gorm.DB {
	return repo.db.Connection()
}

func (repo repository) Add(users []User) error {

	for _, u := range users {

		var user User
		result := repo.DB().Where(User{ID: u.ID}).Assign(u).FirstOrCreate(&user)
		if err := result.Error; err != nil {
			log.Println(err)
		}

		log.Println(user)
	}
	return nil
}

func (repo repository) GetAll() []User {
	var users []User
	result := repo.DB().Find(&users)
	if err := result.Error; err != nil {
		log.Println(err)
		return nil
	}
	return users
}

func (repo repository) Delete() {
	repo.DB().Delete(User{})
}
