package envset

import (
	"context"
	"time"

	"envv/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnvSetStore struct {
	coll *mongo.Collection
}

func NewEnvSetStore(db *mongo.Database) *EnvSetStore {
	return &EnvSetStore{
		coll: db.Collection("envsets"),
	}
}


