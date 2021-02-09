package base_common

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var session *mgo.Session

func addIndexesToDataBase()  {
	var err error
	session := GetSession().Copy()
	defer session.Close()
	userIndex := mgo.Index{
		Key: []string{"email"},
		Unique: true,
		Background: true,
		Sparse: true,
	}
	taskIndex := mgo.Index{
		Key: []string{"createdBy"},
		Unique: false,
		Background: true,
		Sparse: true,
	}
	noteIndex := mgo.Index{
		Key: []string{"taskId"},
		Unique: false,
		Background: true,
		Sparse: true,
	}

	userCol := session.DB(AppConf.Database).C("users")
	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addUserIndexes]: %s\n", err)
	}
	taskCol := session.DB(AppConf.Database).C("task")
	err = taskCol.EnsureIndex(taskIndex)
	if err != nil {
		log.Fatalf("[addTaskIndexes]: %s\n", err)
	}
	noteCol := session.DB(AppConf.Database).C("notes")
	err = noteCol.EnsureIndex(noteIndex)
	if err != nil {
		log.Fatalf("[addNoteIndexes]: %s\n", err)
	}

}

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConf.MongoDBHost},
			Username: AppConf.DBUser,
			Password: AppConf.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConf.MongoDBHost},
		Username: AppConf.DBUser,
		Password: AppConf.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}
