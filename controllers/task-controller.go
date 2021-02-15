package controllers

import (
	"encoding/json"
	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/lambda-0/base-common/data"
	"github.com/lambda-0/base-common/models"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		base_common.DisplayAppError(w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	context := NewContext()
	defer context.Close()
	taskCol := context.Collection(models.TasksCollection)
	repo := &data.TaskRepository{taskCol}
	repo.Create(task)
	if resp, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		base_common.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
	}
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	taskCol := context.Collection(models.TasksCollection)
	repo := &data.TaskRepository{taskCol}
	tasks := repo.GetAll()
	if resp, err := json.Marshal(TasksResource{Data: tasks}); err != nil {
		base_common.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {

}

func GetTaskByUserHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}

