package envv

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"envv/src/config"
	"envv/src/environ"
	"envv/src/store"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [folder]",
	Short: "Import .env from MongoDB for a specific folder",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		folderName := args[0]

		// get uri
		if err := config.Load(); err != nil {
			return err
		}
		cfg := config.Get()

		// Find repo root and ID
		repoRoot, err := environ.FindGitRoot(".")
		if err != nil {
			return err
		}
		repoID := environ.HashRepoID(repoRoot)

		
		envPath := folderName
		if envPath == "root" || envPath == "/" || envPath == "" {
			envPath = "."
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		st, err := store.New(cfg.MongoURI)
		if err != nil {
			return err
		}
		defer st.Close(context.Background())

		envStore := store.NewEnvSetStore(st.DB)

		// Get from DB
		storedEnv, err := envStore.Get(ctx, repoID, envPath)
		if err != nil {
			return fmt.Errorf("failed to find .env for folder '%s': %w", folderName, err)
		}

		// Target file path
		targetPath := filepath.Join(repoRoot, envPath, ".env")
		fmt.Printf("Importing .env to %s\n", targetPath)

		// Write to file
		if err := environ.WriteEnvFile(targetPath, storedEnv.Vars); err != nil {
			return err
		}

		fmt.Println("Import completed successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
