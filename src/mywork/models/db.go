package models

import (
	"fmt"
	"github.com/globalsign/mgo"
	"log"
)

var globalS *mgo.Session

const (
	host      = "127.0.0.1:27017"
	source    = "admin"
	Database  = "mywork"
	poolLimit = 4096
	user      = "user"
	pass      = "123456"
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{host},
		Source:    source,
		Database:  Database,
		PoolLimit: poolLimit,
		//Username: user,
		//Password: pass,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	} else {
		fmt.Println("create session OK ")
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func IsExist(db, collection string, query interface{}) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}
