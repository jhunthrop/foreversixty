package leveling

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ReadActiveBuild is the client build the site currently serves
// (web/src/data/active-build.json's build).
func ReadActiveBuild(repoRoot string) (string, error) {
	type activeBuildFile struct {
		Build string `json:"build"`
	}
	path := filepath.Join(repoRoot, "web", "src", "data", "active-build.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", path, err)
	}
	var f activeBuildFile
	if err := json.Unmarshal(b, &f); err != nil {
		return "", fmt.Errorf("decoding %s: %w", path, err)
	}
	if f.Build == "" {
		return "", fmt.Errorf("%s carries no build", path)
	}
	return f.Build, nil
}
