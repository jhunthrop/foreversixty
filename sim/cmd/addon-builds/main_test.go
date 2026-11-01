package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// TestCommittedAddonBuildsFollowTheGuides: the addon's default talent
// builds (Data.lua's defaultBuilds) come from the committed
// addon-builds.json, which only the client activation used to write. A
// guide build changed since (Fire's raid build took Arcane Meditation in
// round 4 and Improved Fire Blast 2/2 in round 5) left the in-game addon
// on the old one. Regenerate with `go run ./sim/cmd/addon-builds -repo-root .`.
func TestCommittedAddonBuildsFollowTheGuides(t *testing.T) {
	root := filepath.Join("..", "..", "..")
	active, want, err := generate(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(root, "data", "builds", active, "addon-builds.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Error("data/builds/" + active + "/addon-builds.json is stale against the guides' builds; regenerate it with go run ./sim/cmd/addon-builds -repo-root .")
	}
}
