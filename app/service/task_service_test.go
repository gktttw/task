package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http/httptest"
	"strings"
	"task/app/model"
	"task/app/repository"
	"task/mocks"
	"testing"
)

type TaskServiceTestSuite struct {
	suite.Suite
	service TaskService
	repo    repository.TaskRepository
	context *gin.Context
}

func (suite *TaskServiceTestSuite) SetupTest() {
	suite.repo = &mocks.TaskRepository{}
	suite.service = TaskServiceInit(suite.repo)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	suite.context = c
}

func (suite *TaskServiceTestSuite) TestGetAllTask() {
	suite.repo.(*mocks.TaskRepository).On("GetAllTask").Return([]model.Task{
		{
			ID:     1,
			Name:   "Task 1",
			Status: 1,
		},
	}, nil)
	suite.service.GetAllTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "GetAllTask")
}

func (suite *TaskServiceTestSuite) TestCreateTaskSuccess() {
	suite.repo.(*mocks.TaskRepository).On("CreateTask", mock.Anything, mock.Anything).Return(model.Task{
		ID:     1,
		Name:   "Task 1",
		Status: 1,
	}, nil)

	request := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{
		"name": "Task 1",
		"status": 1
	}`))
	request.Header.Add("Content-Type", "application/json")
	suite.context.Request = request

	suite.service.CreateTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "CreateTask", mock.Anything, mock.Anything)
	assert.Equal(suite.T(), 200, suite.context.Writer.Status())
}

func (suite *TaskServiceTestSuite) TestCreateTaskFailUnknownError() {
	suite.repo.(*mocks.TaskRepository).On("CreateTask", mock.Anything, mock.Anything).Return(model.Task{}, errors.New("some error"))

	request := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{
		"name": "Task 1",
		"status": 1
	}`))
	request.Header.Add("Content-Type", "application/json")
	suite.context.Request = request

	suite.service.CreateTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "CreateTask", mock.Anything, mock.Anything)
	assert.Equal(suite.T(), 422, suite.context.Writer.Status())
}

func (suite *TaskServiceTestSuite) TestCreateTaskFailInvalidDataError() {
	suite.repo.(*mocks.TaskRepository).On("CreateTask", mock.Anything, mock.Anything).Return(model.Task{}, gorm.ErrInvalidData)

	request := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{
		"name": "Task 1",
		"status": 1
	}`))
	request.Header.Add("Content-Type", "application/json")
	suite.context.Request = request

	suite.service.CreateTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "CreateTask", mock.Anything, mock.Anything)
	assert.Equal(suite.T(), 400, suite.context.Writer.Status())
}

func (suite *TaskServiceTestSuite) TestDeleteTaskSuccess() {
	suite.repo.(*mocks.TaskRepository).On("DeleteTask", mock.Anything).Return(nil)

	request := httptest.NewRequest("DELETE", "/tasks", nil)
	suite.context.Params = gin.Params{{Key: "taskID", Value: "3"}}
	suite.context.Request = request

	suite.service.DeleteTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "DeleteTask", 3)
	assert.Equal(suite.T(), 200, suite.context.Writer.Status())
}

func (suite *TaskServiceTestSuite) TestDeleteTaskFail() {
	suite.repo.(*mocks.TaskRepository).On("DeleteTask", mock.Anything).Return(errors.New("some error"))

	request := httptest.NewRequest("DELETE", "/tasks", nil)
	suite.context.Params = gin.Params{{Key: "taskID", Value: "3"}}
	suite.context.Request = request

	suite.service.DeleteTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "DeleteTask", 3)
	assert.Equal(suite.T(), 422, suite.context.Writer.Status())
}

func (suite *TaskServiceTestSuite) TestUpdateTaskSuccess() {
	suite.repo.(*mocks.TaskRepository).On("UpdateTask", mock.Anything, mock.Anything, mock.Anything).Return(&model.Task{}, nil)

	request := httptest.NewRequest("PUT", "/tasks", strings.NewReader(`{
			"name": "Task 3",
			"status": 1
		}`))
	request.Header.Add("Content-Type", "application/json")
	suite.context.Params = gin.Params{{Key: "taskID", Value: "3"}}
	suite.context.Request = request
	suite.service.UpdateTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "UpdateTask", 3, "Task 3", 1)

	assert.Equal(suite.T(), 200, suite.context.Writer.Status())
}

func (suite *TaskServiceTestSuite) TestUpdateTaskFail() {
	suite.repo.(*mocks.TaskRepository).On("UpdateTask", mock.Anything, mock.Anything, mock.Anything).Return(&model.Task{}, errors.New("some error"))

	request := httptest.NewRequest("PUT", "/tasks", strings.NewReader(`{
			"name": "Task 3",
			"status": 1
		}`))
	request.Header.Add("Content-Type", "application/json")
	suite.context.Params = gin.Params{{Key: "taskID", Value: "3"}}
	suite.context.Request = request
	suite.service.UpdateTask(suite.context)
	suite.repo.(*mocks.TaskRepository).AssertCalled(suite.T(), "UpdateTask", 3, "Task 3", 1)

	assert.Equal(suite.T(), 400, suite.context.Writer.Status())
}

func TestTaskServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TaskServiceTestSuite))
}
