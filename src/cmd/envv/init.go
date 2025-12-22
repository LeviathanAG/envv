package envv

import (
	"context"
	"fmt"
	"time"

	"envv/src/config"
	"envv/src/environ"
	"envv/src/model"
	"envv/src/store"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize envv by scanning .env files and storing them in MongoDB",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Load developer env to load mongo uri
		if err := config.Load(); err != nil {
			return err
		}
		cfg := config.Get()

		//  Find root
		repoRoot, err := environ.FindGitRoot(".")
		if err != nil {
			return err
		}

		//  hash  repo ID
		repoID := environ.HashRepoID(repoRoot)

		// Scan
		envs, err := environ.ScanRepo(repoRoot, repoID)
		if err != nil {
			return err
		}

		if len(envs) == 0 {
			fmt.Println("No .env files found. Nothing to initialize.")
			return nil
		}

		//  Compute repowide hash 
		repoHash := environ.HashRepo(envs)
		fmt.Printf("Repository env hash: %s\n", repoHash)

		
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		st, err := store.New(cfg.MongoURI)
		if err != nil {
			return err
		}
		defer st.Close(context.Background())

		envStore := store.NewEnvSetStore(st.DB)
		repoStore := store.NewRepoStore(st.DB)

		// Check if repo already exists 
		storedRepo, err := repoStore.Get(ctx, repoID)
		if err == nil && storedRepo != nil {
			// Repo exists so compare hashes
			if storedRepo.EnvHash == repoHash {
				fmt.Println(" No changes detected. Repository is up to date.")
				return nil
			}
			fmt.Printf("Changes detected (stored: %s → new: %s)\n", storedRepo.EnvHash[:8], repoHash[:8])
		}

		//  Store env files into mongo
		for _, env := range envs {
			if err := envStore.Upsert(ctx, env); err != nil {
				return err
			}
			fmt.Printf("Stored %s/.env\n", env.EnvPath)
		}

		// Update repo metadata with new hash
		repo := model.Repo{
			RepoID:   repoID,
			RepoRoot: repoRoot,
			EnvHash:  repoHash,
		}
		if err := repoStore.Upsert(ctx, repo); err != nil {
			return err
		}

		fmt.Println("\nenvv init completed successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
