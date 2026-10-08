package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const raidFixtureTalents = `{"trees":[
 {"position":0,"talents":[{"id":1,"tier":0,"column":0,"max_rank":5}]},
 {"position":1,"talents":[{"id":11,"tier":0,"column":0,"max_rank":5},{"id":10,"tier":0,"column":1,"max_rank":5}]}
]}`

func writeRaidFixture(t *testing.T, frontmatter string) string {
	t.Helper()
	root := t.TempDir()
	for path, body := range map[string]string{
		"data/builds/testbuild/talents/hunter.json":     raidFixtureTalents,
		"web/src/content/guides/hunter/marksmanship.md": "---\n" + frontmatter + "---\n",
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

var raidFixtureSpec = specInfo{ClassSlug: "hunter", SpecSlug: "marksmanship"}

func TestPassTargetsRaidPassSpendsFromTheRaidBuild(t *testing.T) {
	root := writeRaidFixture(t, "build: 'FS1:testbuild:hunter:dwarf:2/50/0:'\nraidBuild: 'FS1:testbuild:hunter:dwarf:4/05/0:'\n")
	got, err := loadPassTargets(root, "testbuild", raidFixtureSpec)
	if err != nil {
		t.Fatal(err)
	}
	wantLeveling := map[int]int{1: 2, 11: 5, 10: 0}
	wantRaid := map[int]int{1: 4, 11: 0, 10: 5}
	if g := got.forPass(presetBare); !reflect.DeepEqual(g, wantLeveling) {
		t.Errorf("bare pass targets = %v, want the leveling build %v", g, wantLeveling)
	}
	if g := got.forPass(presetRaid); !reflect.DeepEqual(g, wantRaid) {
		t.Errorf("raid pass targets = %v, want the raid build %v", g, wantRaid)
	}
}

func TestPassTargetsWithoutRaidBuildReuseTheLevelingBuild(t *testing.T) {
	root := writeRaidFixture(t, "build: 'FS1:testbuild:hunter:dwarf:2/50/0:'\n")
	got, err := loadPassTargets(root, "testbuild", raidFixtureSpec)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.forPass(presetRaid), got.forPass(presetBare)) {
		t.Errorf("raid targets %v differ from leveling targets %v with no raidBuild", got.forPass(presetRaid), got.forPass(presetBare))
	}
}

func TestPassTargetsRejectAStaleRaidBuildStamp(t *testing.T) {
	root := writeRaidFixture(t, "build: 'FS1:testbuild:hunter:dwarf:2/50/0:'\nraidBuild: 'FS1:oldbuild:hunter:dwarf:4/05/0:'\n")
	if _, err := loadPassTargets(root, "testbuild", raidFixtureSpec); err == nil {
		t.Fatal("a raid build stamped for another client build: want an error, got nil")
	}
}
