package main

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type MongoDB struct {
	DatabaseURL    string
	DatabaseName   string
	CollectionName string
}

type Users struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	PassHash string        `json:"passHash"`
}

func setupDB(t *testing.T) *MongoDB {
	db := MongoDB{
		"mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano", // place your mLabs URL here for testing
		"stegano",
		"test_users",
	}

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}
	return &db
}
func (db *MongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// handle to "db"
	count, err := session.DB(db.DatabaseName).C(db.CollectionName).Count()
	if err != nil {
		fmt.Printf("error in Count(): %v", err.Error())
		return -1
	}

	return count
}

/*
Init initializes the mongo storage.
*/
func (db *MongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"passHash"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.CollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func tearDownDB(t *testing.T, db *MongoDB) {
	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()
	if err != nil {
		t.Error(err)
	}

	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

// Testing the Upsert statement
func TestMongo_Upsert(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()

	user := Users{Name: "Etnik", Email: "etnikg@stud.ntnu.no", PassHash: "afierfrieogrfe5623regwtr56g26"}

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var resultUser Users
	upsertInfo, err := session.DB(db.DatabaseName).C(db.CollectionName).Upsert(user, user)

	if err != nil {
		fmt.Printf("error in FindId(): %v", err.Error())
		return
	}

	id := upsertInfo.UpsertedId
	err = session.DB(db.DatabaseName).C(db.CollectionName).FindId(id).One(&resultUser)

	if err != nil {
		fmt.Printf("error in FindId(): %v", err.Error())
		return
	}

	// coll := conn.DB("stegano").C("users")
	// count, err := coll.Find(bson.M{"users": "Etnik"}).Count()
	// if err != nil {
	// 	t.Errorf("Error: %s", err)
	// }

	// if count != 1 {
	// 	t.Error("adding new user failed.")
	// }
}

// Testing the Insert and FindId statements
func TestMongoDB_Insert(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("database not properly initialized. student count() should be 0.")
		return
	}

	user := Users{Name: "Etnik", Email: "entikg@stud.ntnu.no", PassHash: "psaferwfref156213vre", Id: bson.NewObjectId()}
	addUser(user.Name, user.Email, user.PassHash)

	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var resultUser Users
	err = session.DB(db.DatabaseName).C(db.CollectionName).FindId(user.Id).One(&resultUser)

	if err != nil {
		fmt.Printf("error in FindId(): %v", err.Error())
		return
	}

	if db.Count() != 1 {
		t.Error("adding new student failed.")
	}
}

func Test_connection(t *testing.T) {
	_, err := mgo.Dial("mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano")
	if err != nil {
		t.Errorf("Not connected, : %s", err)
	}
}

func Test_existsDB(t *testing.T) {
	coll := conn.DB("stegano").C("users")
	count, err := coll.Find(bson.M{"users": "Etnik"}).Count()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if count > 0 {
		fmt.Println("There are data in DB")
		return
	}

}
