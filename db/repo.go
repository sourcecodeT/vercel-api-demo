package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository struct {
	c *mongo.Client
}

func NewRepository(c *mongo.Client) *Repository {
	return &Repository{c}
}

func (r *Repository) Ping() {
	if err := r.c.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Printf("[ERROR] can not ping to mongo, %v \n", err)
		return
	}

	log.Printf("[INFO] Ping to mongo ok!")
}

func (r *Repository) Disconnect() {
	if err := r.c.Disconnect(context.TODO()); err != nil {
		log.Printf("[ERROR] can not disconnect to mongo, %v", err)
		return
	}

	log.Printf("[INFO] Disconnected to mongo!")
}
