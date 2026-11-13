package leveling

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// The addon's Builds pane shows one line under each build saying why the
// guide builds it that way (design/specs/2026-11-13-addon-builds-pane.md,
// section 3). The guide owns that line: frontmatter `buildWhy:` for the
// leveling build and `raidBuildWhy:` for the raid build, each 64 characters
// or fewer (web/src/content.config.ts), written to stand alone in a row.

// whyKeys are the frontmatter keys of each build kind's why.
var whyKeys = map[string]string{"leveling": "buildWhy", "raid": "raidBuildWhy"}

// GuideBuildWhy reads a spec's guide and returns the why for its build of
// `kind` ("leveling" or "raid"). A guide with no raidBuild has one build, so
// its raid why is its buildWhy. An unknown kind, or a guide missing the why
// its builds need, is an error.
func GuideBuildWhy(repoRoot, class, specSlug, kind string) (string, error) {
	if _, ok := whyKeys[kind]; !ok {
		return "", fmt.Errorf("leveling: no why for build kind %q", kind)
	}
	path := filepath.Join(repoRoot, "web", "src", "content", "guides", class, specSlug+".md")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("leveling: reading %s: %w", path, err)
	}
	why, err := BuildWhy(string(b), kind)
	if err != nil {
		return "", fmt.Errorf("leveling: %s: %w", path, err)
	}
	return why, nil
}

// BuildWhy is GuideBuildWhy over a guide's text.
func BuildWhy(guide, kind string) (string, error) {
	key := whyKeys[kind]
	if kind == "raid" && frontmatterValue(guide, "raidBuild") == "" {
		key = whyKeys["leveling"]
	}
	why := frontmatterValue(guide, key)
	if why == "" {
		return "", fmt.Errorf("no %s line in the frontmatter", key)
	}
	return why, nil
}

// frontmatterValue is a quoted frontmatter value, "" when the key is absent:
// single-quoted (`key: '...'`, a doubled quote standing for one) or, as
// Prettier writes a value holding an apostrophe, double-quoted.
func frontmatterValue(guide, key string) string {
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:\s*(?:'((?:[^']|'')*)'|"((?:[^"\\]|\\.)*)")\s*$`)
	m := re.FindStringSubmatch(guide)
	switch {
	case m == nil:
		return ""
	case m[2] != "":
		return strings.ReplaceAll(strings.ReplaceAll(m[2], `\"`, `"`), `\\`, `\`)
	default:
		return strings.ReplaceAll(m[1], "''", "'")
	}
}
