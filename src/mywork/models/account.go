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
	db         = Database
	collection = "accounts"
)

func (m *Account) IsExist(query interface{}) bool {
	return IsExist(db, collection, query)
}

func (m *Account) Insert(docs interface{}) error {
	return Insert(db, collection, docs)
}

func (m *Account) FindOne(query, selector, result interface{}) error {
	return FindOne(db, collection, query, selector, result)
}

func (m *Account) FindById(id string) (Account, error) {
	var result Account
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (m *Account) FindAllAccount() ([]Account, error) {
	var result []Account
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (m *Account) RemoveAccount(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}

func (m *Account) UpdateAccount(accont Account) error {
	return Update(db, collection, bson.M{"_id": accont.Id}, accont)
}
