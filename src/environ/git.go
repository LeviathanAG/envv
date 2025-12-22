package environ

import (
	"errors"
	"os"
	"path/filepath"
)

// FindGitRoot walks upwards from startDir until it finds a .git directory.
// Returns the repo root path.
func FindGitRoot(startDir string) (string, error) {
	dir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	for {
		gitPath := filepath.Join(dir, ".git")
		info, err := os.Stat(gitPath)
		if err == nil && info.IsDir() {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached filesystem root
		}
		dir = parent
	}

	return "", errors.New("not inside a git repository")
}
