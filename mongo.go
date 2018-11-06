package main

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Returns a mongo client to interact with the database
func mongoConnect() (*mongo.Client, error) {
	// Connect to OpenStack remote MongoDB
	conn, err := mongo.Connect(context.Background(), "mongodb://admin:connecttome123@ds151533.mlab.com:51533/stegano", nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func storeImage(encodedImg []byte) error {
	con, err := mongoConnect()
	if err != nil {
		return err
	}

	coll := con.Database("stegano").Collection("images") // `images` collection, `stegano` database

	// Insert image into collection
	_, err = coll.InsertOne(context.Background(),
		bson.NewDocument(bson.EC.Binary("imgEncoding", encodedImg)),
	)
	if err != nil {
		return err
	}

	return nil
}

func addUser(user, email, passHash string) error {
	conn, err := mongoConnect()
	if err != nil {
		return err
	}

	coll := conn.Database("stegano").Collection("users")
	_, err = coll.InsertOne(context.Background(),
		bson.NewDocument(
			bson.EC.String("user", user),
			bson.EC.String("email", email),
			bson.EC.String("passHash", passHash),
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func entryExists(entry string, value string, collection string) (bool, error) {
	conn, err := mongoConnect()
	if err != nil {
		return false, err
	}

	coll := conn.Database("stegano").Collection(collection)
	cur, err := coll.Find(context.Background(), bson.NewDocument(bson.EC.String(entry, value)))
	if err != nil {
		return false, err
	}

	var holder map[string]interface{}
	for cur.Next(context.Background()) {
		err := cur.Decode(&holder)
		if err != nil {
			return false, err
		}
	}

	if _, ok := holder[entry]; ok {
		return true, nil
	}
	return false, nil
}

func getImage(client *mongo.Client) (map[string]interface{}, error) {

	coll := client.Database("stegano").Collection("images") // `images` collection, `stegano` database

	cur, err := coll.Find(context.Background(), nil) // Find all occurences
	if err != nil {
		return nil, err
	}

	var img map[string]interface{} // Here we'll store fetched images

	for cur.Next(context.Background()) { // Iterate the cursor
		err := cur.Decode(&img) // Store fetched images
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}
