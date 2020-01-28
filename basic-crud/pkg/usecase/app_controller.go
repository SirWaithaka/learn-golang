package usecase

import (
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/network/service"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/usecase/device"
)

func NewAppController(i AppInteractor, s service.ApiService) App {
	return &app{i, s}
}

type App interface {
	DeviceController() device.Controller
}

type app struct {
	interactor AppInteractor
	service    service.ApiService
}

func (a *app) DeviceController() device.Controller {
	return device.NewController((*a).interactor.DeviceInteractor(), (*a).service)
}
