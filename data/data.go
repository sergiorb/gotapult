package data

import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Store *Datastore

func Init() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017/"))

	// Store, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	Store = &Datastore{
		HttpEvents:	client.Database("gotapult").Collection("events_http"),
		MqttEvents:	client.Database("gotapult").Collection("events_mqtt"),
		Ctx:		func() context.Context {

			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

			return ctx
		},
	}
}

type Datastore struct {
	Ctx			func() context.Context
	HttpEvents	*mongo.Collection
	MqttEvents	*mongo.Collection
}
