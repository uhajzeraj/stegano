package main

import (
	"context"

	_ "crypto/sha256"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Returns a mongo client to use for interactions with the database
func mongoConnect() (*mongo.Client, error) {
	// Connect to OpenStack remote MongoDB
	conn, err := mongo.Connect(context.Background(), "mongodb://admin:adminPass321@10.212.138.222/stegano", nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func storeImage(client *mongo.Client, encodedImg string) error {

	coll := client.Database("stegano").Collection("images") // `images` collection, `stegano` database

	_, err := coll.InsertOne(context.Background(), // Insert base64 encoding of image in the database
		bson.NewDocument(
			bson.EC.String("imgEncoding", encodedImg),
		))

	if err != nil {
		return err
	}

	return nil
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
