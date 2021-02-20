package controllers

import "github.com/lambda-0/base-common/models"

//user resource
type UserResource struct {
	Data models.User `json:"data"`
}

type LoginResource struct {
	Data LoginModel `json:"data"`
}

type AuthUserResource struct {
	Data AuthUserModel `json:"data"`
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserModel struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

//task resource
type TaskResource struct {
	Data models.Task `json:"data"`
}

type TasksResource struct {
	Data []models.Task `json:"data"`
}

//note resource
type NoteResource struct {
	Data NoteModel `json:"data"`
}

type NotesResource struct {
	Data []models.TaskNote `json:"data"`
}

type NoteModel struct {
	TaskId      string `json:"taskId"`
	Description string `json:"description"`
}
