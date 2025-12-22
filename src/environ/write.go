package environ

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// WriteEnvFile writes a map of environment variables back to a .env file.
func WriteEnvFile(path string, vars map[string]string) error {
	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// Sort keys for deterministic output
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if _, err := f.WriteString(fmt.Sprintf("%s=%s\n", k, vars[k])); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	return nil
}
