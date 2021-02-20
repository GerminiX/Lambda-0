package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	base_common "github.com/lambda-0/base-common/base-common"
	"github.com/lambda-0/base-common/data"
	"github.com/lambda-0/base-common/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
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
	repo := &data.TaskRepository{Col: taskCol}
	_ = repo.Create(task)
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
	repo := &data.TaskRepository{Col: taskCol}
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

func GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, _ := primitive.ObjectIDFromHex(vars["id"])
	context := NewContext()
	defer context.Close()
	taskCol := context.Collection(models.TasksCollection)
	repo := &data.TaskRepository{Col: taskCol}
	task, err := repo.GetById(taskId)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			base_common.DisplayAppError(w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
	}
	if resp, err := json.Marshal(task); err != nil {
		base_common.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func GetTaskByUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	context := NewContext()
	defer context.Close()
	taskCol := context.Collection(models.TasksCollection)
	repo := &data.TaskRepository{taskCol}
	tasks := repo.GetByUser(userId)
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

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	taskId, _ := primitive.ObjectIDFromHex(vars["id"])
	context := NewContext()
	defer context.Close()
	taskCol := context.Collection(models.TasksCollection)
	repo := &data.TaskRepository{Col: taskCol}
	err := repo.Delete(&taskId)
	if err != nil {
		base_common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, _ := primitive.ObjectIDFromHex(vars["id"])
	var dataResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		base_common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	task.Id = taskId
	context := NewContext()
	defer context.Close()
	taskCol := context.Collection(models.TasksCollection)
	repo := &data.TaskRepository{taskCol}
	if err := repo.Update(task); err != nil {
		base_common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}else {
		w.WriteHeader(http.StatusNoContent)
	}
}
