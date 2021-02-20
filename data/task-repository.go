package data

import (
	"github.com/lambda-0/base-common/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TaskRepository struct {
	Col *mgo.Collection
}

func (rep *TaskRepository) Create(task *models.Task) error  {
	obj_id := primitive.NewObjectID()
	task.Id = obj_id
	task.CreatedOn = time.Now().UnixNano()/int64(time.Millisecond)
	task.Status = "Created"
	err := rep.Col.Insert(&task)
	if err != nil {
		return err
	}
	return nil
}

func (rep *TaskRepository) Update(task *models.Task) error  {
	err := rep.Col.Update(bson.M{"_id": task.Id},
		bson.M{"$set": bson.M{
			"name" : task.Name,
			"description": task.Description,
			"due" : task.DueOn,
			"status" : task.Status,
			"tags" : task.Tags,
		}})
	if err != nil {
		return err
	}
	return nil
}

func (rep *TaskRepository) Delete(id *primitive.ObjectID) error {
	err := rep.Col.Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (rep *TaskRepository) GetAll() []models.Task {
	var tasks []models.Task
	iter := rep.Col.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

