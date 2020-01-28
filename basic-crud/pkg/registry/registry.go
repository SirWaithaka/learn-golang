package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/network/service"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/storage"
	"github.com/sirwaithaka/learn-golang/basic-crud/pkg/usecase"
)

func New(db *gorm.DB, hc *service.HTTPClient) Registry {
	return &registry{db, hc}
}

type registry struct {
	db         *gorm.DB
	httpClient *service.HTTPClient
}

type Registry interface {
	NewAppController() usecase.App
	NewStorageRepository() storage.Storage
	NewInteractor() usecase.AppInteractor
}

func (r *registry) NewStorageRepository() storage.Storage {
	return storage.GetStorageRepository(r.db)
}

func (r *registry) NewApiService() service.ApiService {
	return *service.NewApiService(r.httpClient)
}

func (r *registry) NewInteractor() usecase.AppInteractor {
	return usecase.NewAppInteractor(r.NewStorageRepository())
}

func (r *registry) NewAppController() usecase.App {
	return usecase.NewAppController((*r).NewInteractor(), (*r).NewApiService())
}
