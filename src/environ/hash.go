package environ

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

// HashEnvSet computes a deterministic hash for one env file
func HashEnvSet(envPath string, vars map[string]string) string {
	keys := make([]string, 0, len(vars))
	for k := range vars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	h := sha256.New()
	h.Write([]byte(envPath))

	for _, k := range keys {
		h.Write([]byte(k))
		h.Write([]byte("="))
		h.Write([]byte(vars[k]))
		h.Write([]byte("\n"))
	}

	return hex.EncodeToString(h.Sum(nil))
}
