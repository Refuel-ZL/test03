package models

import "github.com/globalsign/mgo/bson"

type Address struct {
	Id     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Street string        `bson:"street" json:"street"`
}

type Account struct {
	Id       bson.ObjectId `bson:"_id" json:"id"`
	Username string        `bson:"username" json:"username"`
	Name     string        `bson:"name" json:"name"`
	Phone    string        `bson:"phone" json:"phone"`
	Password string        `bson:"_password" json:"_password"`
	Role     []string      `bson:"role" json:"role"`
	Address  []Address     `bson:"address" json:"address"`
}

const (
	db         = "test"
	collection = "accounts"
)

func (m *Account) FindAllAccount() ([]Account, error) {
	var result []Account
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}
func (m *Account)IsExist( query interface{}) bool {
	return IsExist(db, collection, query)
}
