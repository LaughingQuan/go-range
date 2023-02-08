package nosql

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"xmirror.cn/iast/goat/config"
)

var (
	mongoClient *mongo.Client
	mongoDB     *mongo.Database
	collection  *mongo.Collection
	ctx         context.Context
	cancel      context.CancelFunc
)

func initData() {
	mogConfig, err := config.GetConfMog()
	if err != nil {
		log.Printf("Could not get mongo config: err = %s", err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", mogConfig.Username, mogConfig.Password, mogConfig.Host, mogConfig.Port)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("Could not connect the Mongo client: err = %s", err)
	}
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("not ping mongo server err = %s", err)
		return
	}
	mongoDB = mongoClient.Database(mogConfig.Database)
	collection = mongoDB.Collection(mogConfig.Collection)

	opts := options.InsertMany().SetOrdered(false)
	docs := []interface{}{
		bson.D{{Key: "name", Value: RandString(10)}},
		bson.D{{Key: "name", Value: RandString(10)}},
	}

	_, _ = collection.InsertMany(context.TODO(), docs, opts)
}

//MongoKill cleans up the mongo database by dropping the data and
// disconnecting the client
func MongoKill() {
	_ = collection.Drop(ctx)
	_ = mongoDB.Drop(ctx)
	_ = mongoClient.Disconnect(ctx)
	cancel()
}

func RandString(len int) string {
	a := make([]rune, len)
	for i := range a {
		a[i] = rune(RandInt(19968, 40869))
	}
	return string(a)
}

func RandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}
