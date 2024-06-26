// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"github.com/google/wire"
	"task/app/controller"
	"task/app/repository"
	"task/app/service"
	"task/db"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := db.Connect()
	taskRepositoryImpl := repository.TaskRepositoryInit(gormDB)
	taskServiceImpl := service.TaskServiceInit(taskRepositoryImpl)
	taskControllerImpl := controller.TaskControllerInit(taskServiceImpl)
	initialization := NewInitialization(taskRepositoryImpl, taskServiceImpl, taskControllerImpl)
	return initialization
}

// injector.go:

var database = wire.NewSet(db.Connect)

var taskServiceSet = wire.NewSet(service.TaskServiceInit, wire.Bind(new(service.TaskService), new(*service.TaskServiceImpl)))

var taskRepoSet = wire.NewSet(repository.TaskRepositoryInit, wire.Bind(new(repository.TaskRepository), new(*repository.TaskRepositoryImpl)))

var taskCtrlSet = wire.NewSet(controller.TaskControllerInit, wire.Bind(new(controller.TaskController), new(*controller.TaskControllerImpl)))
