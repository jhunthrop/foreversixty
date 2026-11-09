package bnetimport

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The Battle.net import never defaults a ruleset. From 2026-11-05 to
// 2026-11-09 every sign-up fell through to "normal" when Blizzard stopped
// listing Forever's realms, which keyed PvP players and their guilds as
// normal. The only place the import may name normal is bnetapi.RulesetOf,
// for a realm Blizzard itself types NORMAL; everything else must resolve a
// ruleset from a source or skip the character.
func TestNoImportCodePathDefaultsTheRuleset(t *testing.T) {
	allowed := map[string]int{"../bnetapi/realms.go": 1} // case "NORMAL"
	for _, dir := range []string{".", "../bnetapi"} {
		files, err := filepath.Glob(filepath.Join(dir, "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range files {
			if strings.HasSuffix(path, "_test.go") {
				continue
			}
			src, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			got := strings.Count(string(src), "RulesetNormal") + strings.Count(string(src), `"normal"`)
			if got != allowed[filepath.ToSlash(path)] {
				t.Errorf("%s names the normal ruleset %d times, want %d: resolve it from a source or skip",
					path, got, allowed[filepath.ToSlash(path)])
			}
		}
	}
}
