package usecase

import (
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/storage"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/usecase/device"
)

func NewAppInteractor(repository storage.Storage) AppInteractor {
	return &appInteractor{repository}
}

type appInteractor struct {
	repository storage.Storage
}

type AppInteractor interface {
	DeviceInteractor() device.Interactor
}

func (ai *appInteractor) DeviceInteractor() device.Interactor {
	return device.NewInteractor((*ai).repository.GetDeviceDao())
}
