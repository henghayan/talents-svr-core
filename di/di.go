package di

import (
	"core/config"
	"core/db"
	"core/pkg/Infrastructure/repository"
	"core/pkg/application"
	"sync"
)

type ServiceLocator struct {
	userService *application.UserService
	once        sync.Once
}

var instance *ServiceLocator
var once sync.Once

func GetServiceLocator(cfg *config.Config) *ServiceLocator {
	once.Do(func() {
		if cfg == nil {
			panic("init service, cfg nil!")
		}
		instance = &ServiceLocator{}
		instance.init(cfg)
	})
	return instance
}

func (sl *ServiceLocator) init(cfg *config.Config) {
	newDB, err := initializeDatabase(cfg)
	if err != nil {
		panic("db not init!")
	}
	userRepo := repository.NewUserRepository(newDB.NewTable("user"))
	sl.userService = application.NewUserService(userRepo)
}

func (sl *ServiceLocator) UserService() *application.UserService {
	return sl.userService
}

func initializeDatabase(cfg *config.Config) (*db.Database, error) {
	newDB, err := db.NewDatabase(cfg.DBPassword)
	if err != nil {
		return nil, err
	}
	return newDB, nil
}
