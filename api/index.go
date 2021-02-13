package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"vercel-api-demo/app"
	"vercel-api-demo/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MONGO_URI = os.Getenv("MONGODB_URI")
	client    *mongo.Client
	repo      *db.Repository
)

func init() {
	if client == nil {
		client = db.NewMongo(MONGO_URI)
		repo = db.NewRepository(client)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		if err := client.Connect(context.TODO()); err != nil {
			log.Fatalf("[ERROR] can not connect to mongo, %v", err)
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	repo.Ping()
	app.ResponseSuccess(w, "Hello world!")

	defer repo.Disconnect()
}
