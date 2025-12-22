package store

import "time"
import "context"
import "os"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import "go.mongodb.org/mongo-driver/mongo/readpref"


type store struct {
	client *mongo.Client
	DB *mongo.Database
}


func New(mongoURI string)(*store , error){
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel() // defer runs at the end of the function no matter what
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI)) 
	if err != nil {
		return nil, err
	}

	// Ping 
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database("envv") // TODO : make the db name configurable flag or env var

	return &Store{
		Client: client,
		DB:     db,
	}, nil // return this instance of store for CRUD 
}

func (s *store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

