//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"task/app/controller"
	"task/app/repository"
	"task/app/service"
	"task/db"
)

var database = wire.NewSet(db.Connect)

var taskServiceSet = wire.NewSet(service.TaskServiceInit,
	wire.Bind(new(service.TaskService), new(*service.TaskServiceImpl)),
)

var taskRepoSet = wire.NewSet(repository.TaskRepositoryInit,
	wire.Bind(new(repository.TaskRepository), new(*repository.TaskRepositoryImpl)),
)

var taskCtrlSet = wire.NewSet(controller.TaskControllerInit,
	wire.Bind(new(controller.TaskController), new(*controller.TaskControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, database, taskRepoSet, taskServiceSet, taskCtrlSet)
	return nil
}
