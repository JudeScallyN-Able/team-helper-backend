package graph_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"team-helper-backend/graph"
	"team-helper-backend/graph/model"
	"testing"
)

type MockJson struct {
	mock.Mock
}

type MockFileReader struct {
	mock.Mock
}

type SchemaHelperTest struct {
	suite.Suite

	mockFileReader *MockFileReader
	mockJson       *MockJson
}

func (fr *MockFileReader) ReadFile(filePath string) ([]byte, error) {
	args := fr.Called(filePath)
	return args.Get(0).([]byte), args.Error(1)
}

func (j *MockJson) Unmarshal(data []byte, v any) error {
	args := j.Called(data, v)
	return args.Error(0)
}

func TestSchemaHelpersTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaHelperTest))
}

func (s *SchemaHelperTest) TestGetAllTasksReturnsExpectedTasks() {
	mockJsonBlock := `[{
	"id": "1",
	"title": "mock title",
	"description": "mock desc",
	"status": "TODO"
	}]`
	s.mockFileReader.On("ReadFile", "tasks.json").Return(mockJsonBlock, nil)
	s.mockJson.On("Unmarshal", mockJsonBlock, []*model.Task{}).Return(nil)

	tasks, err := graph.GetAllTasks()

	s.NoError(err)
	s.Equal(tasks[0].ID, "1")
	s.Equal(tasks[0].Title, "mock title")
	s.Equal(tasks[0].Description, "mock desc")
	s.Equal(tasks[0].Status, "TODO")
}
