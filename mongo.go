package main

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func mongoConnect() (*mongo.Client, error) {
	// Connect to OpenStack remote MongoDB
	conn, err := mongo.Connect(context.Background(), "mongodb://admin:adminPass321@10.212.138.222/stegano", nil)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return conn, nil
}
