package main

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	conn, _ = mgo.Dial("mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano")
)

// Images struct for storing fetched images
type Images struct {
	Images []int
}

// Returns a mongo client to interact with the database
// func mongoConnect() (*mongo.Client, error) {
// 	// Connect to OpenStack remote MongoDB
// 	conn, err := mongo.Connect(context.Background(), "mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano", nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return conn, nil
// }

func storeImage(encodedImg []byte) error {

	coll := conn.DB("stegano").C("images") // `images` collection, `stegano` database

	// Insert image into collection
	err := coll.Insert(
		bson.M{"imgEncoding": encodedImg},
	)
	if err != nil {
		return err
	}

	return nil
}

func addUser(user, email, passHash string) error {
	// conn, err := mongoConnect()
	// if err != nil {
	// 	return err
	// }

	coll := conn.DB("stegano").C("users")
	err := coll.Insert(
		bson.M{
			"user":     user,
			"email":    email,
			"passHash": passHash},
	)
	if err != nil {
		return err
	}

	return nil
}

func entryExists(entry string, value string, collection string) (bool, error) {

	coll := conn.DB("stegano").C(collection)
	count, err := coll.Find(bson.M{entry: value}).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func getImages(user string) ([]int, error) {

	coll := conn.DB("stegano").C("users") // `images` collection, `stegano` database

	var img Images

	err := coll.Find(
		bson.M{"user": user},
	).Select(
		bson.M{"images": 1, "_id": 0},
	).One(&img)
	if err != nil {
		return nil, err
	}

	return img.Images, nil
}
