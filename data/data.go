package data

import (
	"fmt"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/sergiorb/gotapult/data/config"
)

var Store *Datastore

func Init(config config.Conf) {

	fmt.Sprintf("mongodb://%v:%f/", config.Host, config.Port)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%d/", config.Host, config.Port)))

	// Store, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	Store = &Datastore{
		HttpEvents:	client.Database(config.Database).Collection("events_http"),
		MqttEvents:	client.Database(config.Database).Collection("events_mqtt"),
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
