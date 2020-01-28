package rest

import (
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/usecase"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router(app usecase.App) http.Handler {
	router := httprouter.New()

	//router.GET("/device", getActivePlayer(app.DeviceController()))
	//router.GET("/adverts", getDevice(app.DeviceController()))
	//router.GET("/advert", getDevice(app.DeviceController()))

	//router.GET("/conversion", getDevice(app.DeviceController()))

	router.GET("/player", getActivePlayer(app.DeviceController()))
	router.PUT("/player", updatePlayer(app.DeviceController()))
	router.DELETE("/player", deactivatePlayer(app.DeviceController()))

	router.POST("/player/authentication", authenticate(app.DeviceController()))
	//router.GET("/device", getDevice(app.DeviceController()))
	//router.GET("/device", getDevice(app.DeviceController()))

	return router
}
