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
	Images []Image `bson:"images"`
}

// Image struct for storig one image
type Image struct {
	Name string `bson:"name"`
	Img  []byte `bson:"img"`
}

func storeImage(user, imageName string, encodedImg []byte) error {

	coll := conn.DB("stegano").C("users") // `users` collection, `stegano` database

	// Insert image into collection
	err := coll.Update(
		bson.M{"user": user},
		bson.M{"$push": bson.M{"images": bson.M{"name": imageName, "img": encodedImg}}},
	)
	if err != nil {
		return err
	}

	return nil
}

func addUser(user, email, passHash string) error {

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

func getImages(user string) ([]Image, error) {

	coll := conn.DB("stegano").C("users") // `users` collection, `stegano` database

	var img Images

	err := coll.Find(
		bson.M{"user": user},
	).Select(
		bson.M{
			"images": 1,
		},
	).One(&img)

	if err != nil {
		return nil, err
	}

	return img.Images, nil
}

func removeImage(user, imageName string) error {

	coll := conn.DB("stegano").C("users") // `users` collection, `stegano` database

	// Remove the image
	err := coll.Update(
		bson.M{"user": user},
		bson.M{"$pull": bson.M{"images": bson.M{"name": imageName}}},
	)
	if err != nil {
		return err
	}

	return nil
}
