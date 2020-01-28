package main

import (
	"fmt"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/network/service"
	"log"
	"net/http"
	"os"

	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/http/rest"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/registry"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/storage/sqlite"
)

func main() {
	var port string = fmt.Sprintf(":%d", 5000)
	fmt.Printf("Server has been started on port: %s\n", port)

	config := GetConfig(os.Getenv("YOUTISE_ENV"))

	sqlite.Migrate()

	db, _ := sqlite.GetDatabase()
	defer db.Close()

	client := service.NewHTTPClient(service.NewClient(), config.YoutiseURL)

	r := registry.New(db.Connection(), client)
	controller := r.NewAppController()

	log.Fatal(http.ListenAndServe(port, rest.Router(controller)))
}
