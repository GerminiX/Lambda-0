package controllers

import (
	base_common "github.com/lambda-0/base-common/base-common"
	"gopkg.in/mgo.v2"
)

type Context struct {
	MongoDBSession *mgo.Session
}

func NewContext() *Context {
	session := base_common.GetSession().Copy()
	context := &Context{
		MongoDBSession: session,
	}
	return context
}

func (cont *Context) Collection(name string) *mgo.Collection  {
	return cont.MongoDBSession.DB(base_common.AppConf.Database).C(name)
}

func (cont *Context) Close()  {
	cont.MongoDBSession.Close()
}
