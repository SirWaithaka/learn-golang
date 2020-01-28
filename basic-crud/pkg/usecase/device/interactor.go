package device

import (
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/domain/entity"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/storage"
)

func NewInteractor(dr storage.DeviceDao) Interactor {
	return &deviceInteractor{
		DeviceRepository: dr,
	}
}

type Interactor interface {
	GetActive() entity.Player
	CreateDevice(player entity.Player) error
}

type deviceInteractor struct {
	DeviceRepository storage.DeviceDao
}

func (di *deviceInteractor) GetActive() entity.Player {
	d, _ := (*di).DeviceRepository.GetActive()

	device := entity.Player{
		Active:      d.Active,
		AuthStatus:  d.AuthStatus,
		Token:       d.Token,
		PlaceId:     d.Token,
		PlaceName:   d.PlaceName,
		LastUpdated: d.LastUpdated,
	}
	return device
}

func (di *deviceInteractor) AddPlayer(player entity.Player) {

}

func (di *deviceInteractor) CreateDevice(player entity.Player) error {

	return nil
}
