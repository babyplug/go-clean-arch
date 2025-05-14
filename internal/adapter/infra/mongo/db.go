package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx    = context.Background()
	client *mongo.Client
	cOnce  = sync.Once{}
)

func New(ctx context.Context, uri string) (*mongo.Client, error) {
	var err error
	cOnce.Do(func() {

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			return
		}

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return
		}
	})
	return client, err
}
