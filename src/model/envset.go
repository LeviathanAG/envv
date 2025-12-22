package model

import "time"

type envset struct {
	ID 	  string    `bson:"_id,omitempty"`
	REPOID string   `bson:"repo_id"`
	
	Env_parent_folder_name string `bson:"env_parent_folder_name"`
	Envpath string `bson:"envpath"`


	vars map[string]string `bson:"vars"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Hash string `bson:"hash"`

}


// this is the document to be stored in the repo collection.

// for each directory/file contiaining env files, we create an envset document with its path, parent folder name, repo id and the key value pairs as a map in vars field.
// at the end we will also create a hash based on the vars to check for changes later on.
// also hash the overall documentset to check if new env files are created. this is if new env files are added to the repo.

