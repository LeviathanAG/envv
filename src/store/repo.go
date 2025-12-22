package store

import (
	"context"
	"time"

	"envv/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoStore struct {
	coll *mongo.Collection
}

func NewRepoStore(db *mongo.Database) *RepoStore {
	return &RepoStore{
		coll: db.Collection("repos"),
	}
}

// Get repo metadata
func (s *RepoStore) Get(ctx context.Context, repoID string) (*model.Repo, error) {
	var repo model.Repo
	err := s.coll.FindOne(ctx, bson.M{"repo_id": repoID}).Decode(&repo)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

// Upsert repo metadata
func (s *RepoStore) Upsert(ctx context.Context, repo model.Repo) error {
	now := time.Now()

	repo.UpdatedAt = now
	if repo.CreatedAt.IsZero() {
		repo.CreatedAt = now
	}

	_, err := s.coll.UpdateOne(
		ctx,
		bson.M{"repo_id": repo.RepoID},
		bson.M{"$set": repo},
		options.Update().SetUpsert(true),
	)

	return err
}
