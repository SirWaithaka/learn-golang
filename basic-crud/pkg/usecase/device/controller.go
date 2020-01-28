package device

import (
	"fmt"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/domain/entity"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/network/service"
)

func NewController(i Interactor, s service.ApiService) Controller {
	return &deviceController{i, s}
}

type Controller interface {
	GetActiveDevice() entity.Player
	Authenticate(map[string]string) error
}

type deviceController struct {
	deviceInteractor Interactor
	apiservice       service.ApiService
}

func (dc *deviceController) GetActiveDevice() entity.Player {
	return (*dc).deviceInteractor.GetActive()
}

func (dc *deviceController) Authenticate(params map[string]string) error {
	placeId := params["placeId"]
	password := params["password"]

	byteResp, err := dc.apiservice.Authenticate(placeId, password)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Println(string(byteResp))

	return nil
}
