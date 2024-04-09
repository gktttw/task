package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"task/app/service"
	"task/mocks"
	"testing"
)

type TaskControllerTestSuite struct {
	suite.Suite
	service    service.TaskService
	controller TaskController
	context    *gin.Context
}

func (suite *TaskControllerTestSuite) SetupTest() {
	suite.service = &mocks.TaskService{}
	suite.controller = TaskControllerInit(suite.service)
}

func (suite *TaskControllerTestSuite) TestGetAllTaskData() {
	suite.service.(*mocks.TaskService).On("GetAllTask", suite.context)
	suite.controller.GetAllTaskData(suite.context)

	suite.service.(*mocks.TaskService).AssertCalled(suite.T(), "GetAllTask", suite.context)
}

func (suite *TaskControllerTestSuite) TestCreateTask() {
	suite.service.(*mocks.TaskService).On("CreateTask", suite.context)
	suite.controller.CreateTask(suite.context)

	suite.service.(*mocks.TaskService).AssertCalled(suite.T(), "CreateTask", suite.context)
}

func (suite *TaskControllerTestSuite) TestUpdateTask() {
	suite.service.(*mocks.TaskService).On("UpdateTask", suite.context)
	suite.controller.UpdateTask(suite.context)

	suite.service.(*mocks.TaskService).AssertCalled(suite.T(), "UpdateTask", suite.context)
}

func (suite *TaskControllerTestSuite) TestDeleteTask() {
	suite.service.(*mocks.TaskService).On("DeleteTask", suite.context)
	suite.controller.DeleteTask(suite.context)

	suite.service.(*mocks.TaskService).AssertCalled(suite.T(), "DeleteTask", suite.context)
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
