package environ

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"

	"envv/src/model"
)

// hash the full repo to track changes
func HashRepo(envs []model.EnvSet) string {
	hashes := make([]string, 0, len(envs))

	for _, env := range envs {
		hashes = append(hashes, env.Hash)
	}

	sort.Strings(hashes)

	h := sha256.New()
	for _, hash := range hashes {
		h.Write([]byte(hash))
	}

	return hex.EncodeToString(h.Sum(nil))
}
