package model

import "time"

// Repo represents  metadata stored in MongoDB for the repository.
type Repo struct {
	ID        string    `bson:"_id,omitempty"`
	RepoID    string    `bson:"repo_id"`
	RepoRoot  string    `bson:"repo_root"`
	EnvHash   string    `bson:"env_hash"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
