package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	repo TaskRepository
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (suite *TaskRepositoryTestSuite) SetupTest() {
	mockDb, mock, _ := sqlmock.New()
	suite.mock = mock
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	suite.db = db
	suite.repo = TaskRepositoryInit(db)
}

func (suite *TaskRepositoryTestSuite) TestGetAllTask() {
	suite.mock.ExpectQuery(`SELECT`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "status"}).AddRow(1, "Task 1", 1))
	tasks, _ := suite.repo.GetAllTask()

	suite.Equal(1, len(tasks))
	suite.Equal(1, tasks[0].ID)
	suite.Equal("Task 1", tasks[0].Name)
	suite.Equal(1, tasks[0].Status)
}

func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(`INSERT`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "status"}).AddRow(9, "Task 9", 1))
	suite.mock.ExpectCommit()
	task, _ := suite.repo.CreateTask("Task 1", 1)

	suite.Equal(9, task.ID)
	suite.Equal("Task 9", task.Name)
	suite.Equal(1, task.Status)
}

func (suite *TaskRepositoryTestSuite) TestDeleteTask() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tasks" WHERE "tasks"."id" = $1`)).WithArgs(6).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mock.ExpectCommit()
	err := suite.repo.DeleteTask(6)

	assert.NoError(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask() {

	suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "status"}).AddRow(6, "Task 666", 0))
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tasks" SET `)).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mock.ExpectCommit()
	task, err := suite.repo.UpdateTask(6, "Task 6", 1)

	suite.Equal(6, task.ID)
	suite.Equal("Task 6", task.Name)
	suite.Equal(1, task.Status)
	assert.NoError(suite.T(), err)
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
