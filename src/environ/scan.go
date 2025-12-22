package environ

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"envv/src/model"

	"golang.org/x/sync/errgroup"
)

// ScanRepo walks the repo tree starting at repoRoot,
func ScanRepo(repoRoot string, repoID string) ([]model.EnvSet, error) {
	var results []model.EnvSet
	var mu sync.Mutex

	g, _ := errgroup.WithContext(context.Background())

	err := filepath.WalkDir(repoRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip dir
		if d.IsDir() {
			return nil
		}

		// shld only process .env. first ever performance optimization nice!
		if d.Name() != ".env" {
			return nil
		}

		// Launch a goroutine for each file
		g.Go(func() error {
			// Parse env vars from file
			vars, err := parseEnvFile(path)
			if err != nil {
				return err
			}

			// Determine env path relative to repo root
			relDir, err := filepath.Rel(repoRoot, filepath.Dir(path))
			if err != nil {
				return err
			}

			// Normalize root env path
			if relDir == "." {
				relDir = "."
			}

			// Compute hash for env doc
			hash := HashEnvSet(relDir, vars)

			env := model.EnvSet{
				RepoID:   repoID,
				RepoPath: repoRoot,
				EnvName:  envNameFromPath(relDir),
				EnvPath:  relDir,
				Vars:     vars,
				Hash:     hash,
			}

			mu.Lock()
			results = append(results, env)
			mu.Unlock()

			return nil
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Wait for all goroutines to complete hopefully it works
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

// parse the env file at given path into a hashset of key value pairs
func parseEnvFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	vars := make(map[string]string)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		vars[key] = val
	}

	return vars, scanner.Err()
}

// 
func envNameFromPath(relPath string) string {
	if relPath == "." {
		return "root"
	}
	return filepath.Base(relPath)
}
