package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/usecase/device"
)

func getActivePlayer(controller device.Controller) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(controller.GetActiveDevice())
	}
}

func updatePlayer(controller device.Controller) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	}
}

func deactivatePlayer(controller device.Controller) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	}
}

func authenticate(controller device.Controller) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var p map[string]string

		// Populate the post parameters
		json.NewDecoder(r.Body).Decode(&p)

		err := controller.Authenticate(p)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}

	}
}

func getPlayerAttribute(controller device.Controller) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	}
}
