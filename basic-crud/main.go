package main

import (
	"basic-crud/storage/sqlite"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"basic-crud/domain/entities"
)

func main() {
	jsonStr :=
		`[{	"name": "mary",	"id": 11},
		{ "name": "kenn", "id": 12}]`

	var users []entities.User
	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&users)
	if err != nil {
		fmt.Printf("Error occured decoding %v", err)
	}

	db, err := sqlite.GetDatabase()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	sqlite.Migrate()

	var userRepo entities.Repository = entities.NewRepository(db)
	err = userRepo.Add(users)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("users after add %v\n", userRepo.GetAll())

	//userRepo.Delete()

	//fmt.Printf("users after delete %v\n", userRepo.GetAll())
}
