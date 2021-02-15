package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UsersCollection = "users"
	TasksCollection = "tasks"
)

type (
	User struct {
		Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		FirstName    string             `bson:"firstName" json:"firstName"`
		LastName     string             `bson:"LastName" json:"lastName"`
		Email        string             `bson:"email" json:"email"`
		Password     string             `bson:"password,omitempty" json:"password"`
		HashPassword []byte             `bson:"hashPassword,omitempty" json:"hashPassword"`
	}
	Task struct {
		Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		CreatedBy   string             `bson:"createdBy" json:"createdBy"`
		UpdatedBy   string             `bson:"updatedBy" json:"updatedBy"`
		Name        string             `bson:"name" json:"name"`
		Description string             `bson:"description" json:"description"`
		CreatedOn   int64              `bson:"createdOn" json:"createdOn"`
		UpdatedOn   int64              `bson:"updatedOn" json:"updatedOn"`
		DueOn       int64              `bson:"dueOn" json:"dueOn"`
		Status      string             `bson:"status" json:"status"`
		Tags        []string           `bson:"tags" json:"tags"`
	}
	TaskNote struct {
		Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		TaskId      primitive.ObjectID `bson:"taskId" json:"taskId"`
		CreatedBy   string             `bson:"createdBy" json:"createdBy"`
		UpdatedBy   string             `bson:"updatedBy" json:"updatedBy"`
		Description string             `bson:"description" json:"description"`
		CreatedOn   int64              `bson:"createdOn" json:"createdOn"`
		UpdatedOn   int64              `bson:"updatedOn" json:"updatedOn"`
	}
)
