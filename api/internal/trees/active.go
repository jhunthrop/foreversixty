package trees

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
)

// ActiveBuildFile is the file that names the active client build. The site, the addon and the
// guides all follow web/src/data/active-build.json, which tools/activate-build.sh writes when it
// moves everything to a new build. A build that data.yml's fetch has landed under data/builds
// but no activation has moved to is on disk and loadable by version, and it must never change
// what the API serves, so the API reads the same file rather than picking the newest folder.
const ActiveBuildFile = "active-build.json"

// activeBuildFiles are where Load looks for ActiveBuildFile, first found wins:
//
//  1. beside the build directories: the image's /data/active-build.json, which api/Dockerfile
//     copies from web/src/data/active-build.json (the one file the activation writes);
//  2. the repository's web/src/data/active-build.json, for a dir that is the repository's own
//     data/builds (local runs, tools/dev-stack and the tests that read the real data).
func activeBuildFiles(dir string) []string {
	return []string{
		filepath.Join(dir, ActiveBuildFile),
		filepath.Join(dir, "..", "..", "web", "src", "data", ActiveBuildFile),
	}
}

// resolveActive is the active build among versions and the file that named it. With no file
// at all it falls back to the newest client build folder, logged, with from "": that is a
// fixture or temp directory, never the image. A file that cannot be read or decoded, names no
// build, or names a build that is not on disk fails the load: serving another build than the
// one the site shows is worse than not booting.
func resolveActive(dir string, versions []string) (version, from string, err error) {
	if len(versions) == 0 {
		return "", "", nil
	}
	for _, path := range activeBuildFiles(dir) {
		build, found, err := readActiveBuildFile(path)
		if err != nil {
			return "", "", err
		}
		if !found {
			continue
		}
		if !slices.Contains(versions, build) {
			return "", "", fmt.Errorf("trees: %s names build %q, which is not under %s", path, build, dir)
		}
		return build, path, nil
	}
	newest := newestVersion(versions)
	slog.Default().Warn("trees", "op", "active_build", "dir", dir, "fallback", "newest", "build", newest,
		"reason", "no "+ActiveBuildFile+" beside the builds or in the repository")
	return newest, "", nil
}

// readActiveBuildFile reads {"build": "..."} from path; found is false when the file is absent.
func readActiveBuildFile(path string) (build string, found bool, err error) {
	body, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("trees: read %s: %w", path, err)
	}
	var file struct {
		Build string `json:"build"`
	}
	if err := json.Unmarshal(body, &file); err != nil {
		return "", false, fmt.Errorf("trees: decode %s: %w", path, err)
	}
	if file.Build == "" {
		return "", false, fmt.Errorf("trees: %s names no build", path)
	}
	return file.Build, true, nil
}

// Active is the active build's data: the build ActiveBuildFile names (see resolveActive),
// which is what every "the current build" read in the API means: the rankings' spec
// inference, the Battle.net import's tables and FS1 stamp, the loot values' DPS per point,
// the BiS build. The second return is false when no build loaded at all.
func (d *Data) Active() (*Build, bool) {
	if d.active == "" {
		return nil, false
	}
	return d.Build(d.active)
}

// ActiveFrom is the file that named the active build, "" when Load fell back to the newest
// folder because no file was found. Startup logs it.
func (d *Data) ActiveFrom() string { return d.activeFrom }

// ReadActiveBuild is the build the ActiveBuildFile at path names, for a command that needs
// only the name and has no build directory to Load (the discord-emoji setup, which reads
// web/public/data). Unlike Load it has no fallback: an absent file is an error.
func ReadActiveBuild(path string) (string, error) {
	build, found, err := readActiveBuildFile(path)
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("trees: %s does not exist", path)
	}
	return build, nil
}
