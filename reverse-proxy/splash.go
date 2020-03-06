package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


// Get the port to listen on
func listenAddress() string {
	port := "2709"
	return ":" + port
}


func handleRequests(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s : %s from %s\n", r.Method, r.URL, r.RemoteAddr)

	w.Write([]byte("Request recieved\n"))
}

func main() {

	router := mux.NewRouter()

	router.PathPrefix("/").HandlerFunc(handleRequests)

	if err := http.ListenAndServe(listenAddress(), router); err != nil {
		panic(err)
	}
}
