package store

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


func (s *EnvSetStore) Upsert(ctx context.Context, env model.EnvSet) error {
	now := time.Now()

	env.UpdatedAt = now
	if env.CreatedAt.IsZero() {
		env.CreatedAt = now
	}

	filter := bson.M{
		"repo_id":  env.RepoID,
		"env_path": env.EnvPath,
	}

	update := bson.M{
		"$set": env,
	}

	_, err := s.coll.UpdateOne(
		ctx,
		filter,
		update,
		options.Update().SetUpsert(true),
	)

	return err
}

// fetch a single env file by repoID and envPath
func (s *EnvSetStore) Get(
	ctx context.Context,
	repoID string,
	envPath string,
) (*model.EnvSet, error) {

	var env model.EnvSet

	err := s.coll.FindOne(ctx, bson.M{
		"repo_id":  repoID,
		"env_path": envPath,
	}).Decode(&env)

	if err != nil {
		return nil, err
	}

	return &env, nil
}

// list all Env files for a repo
func (s *EnvSetStore) ListByRepo(
	ctx context.Context,
	repoID string,
) ([]model.EnvSet, error) {

	cur, err := s.coll.Find(ctx, bson.M{
		"repo_id": repoID,
	})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []model.EnvSet
	for cur.Next(ctx) {
		var env model.EnvSet
		if err := cur.Decode(&env); err != nil {
			return nil, err
		}
		results = append(results, env)
	}

	return results, cur.Err()
}

func (s *EnvSetStore) ExistsRepo(ctx context.Context, repoID string) (bool, error) {
	count, err := s.coll.CountDocuments(ctx, bson.M{
		"repo_id": repoID,
	})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ListEnvPaths returns existing env paths for a repo
func (s *EnvSetStore) ListEnvPaths(
	ctx context.Context,
	repoID string,
) ([]string, error) {

	cur, err := s.coll.Find(ctx, bson.M{
		"repo_id": repoID,
	}, options.Find().SetProjection(bson.M{
		"env_path": 1,
	}))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var paths []string
	for cur.Next(ctx) {
		var doc struct {
			EnvPath string `bson:"env_path"`
		}
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		paths = append(paths, doc.EnvPath)
	}

	return paths, cur.Err()
}
