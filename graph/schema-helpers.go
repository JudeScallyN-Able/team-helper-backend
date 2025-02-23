package graph

import (
	"encoding/json"
	"errors"
	"os"
	"team-helper-backend/graph/model"
)

func GetAllTasks() ([]*model.Task, error) {
	jsonTasks, err := os.ReadFile("graph/task-data-store.json")
	if err != nil {
		return nil, errors.New("failed to read tasks")
	}

	var tasks []*model.Task

	err = json.Unmarshal(jsonTasks, &tasks)

	if err != nil {
		return nil, errors.New("failed to unmarshall json")
	}
	return tasks, nil
}
