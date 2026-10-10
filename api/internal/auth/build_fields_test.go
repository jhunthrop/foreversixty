package auth

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// writeHunterTreesFixture builds a minimal, three-tree trees.Data: one class (Hunter,
// slug "hunter") with Beast Mastery/Marksmanship/Survival trees in position order, and
// one race (Troll, slug "troll", horde) -- just enough for buildFieldsFromExport's race
// fill-in and primarySpec to resolve against a real *trees.Data rather than a nil one.
func writeHunterTreesFixture(t *testing.T) *trees.Data {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, "1.60.1.70009")
	if err := os.MkdirAll(filepath.Join(dir, "talents"), 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"classes.json": `[{"id":1,"name":"Hunter","slug":"hunter","color":"#abd473"}]`,
		"races.json":   `[{"id":1,"name":"Troll","slug":"troll","faction":"horde"}]`,
		"combos.json":  `[{"race_id":1,"class_id":1,"new_in_forever":false}]`,
		"talents/hunter.json": `{"build":"1.60.1.70009","class_id":1,"class_slug":"hunter","trees":[
			{"id":1,"name":"Beast Mastery","position":0,"talents":[]},
			{"id":2,"name":"Marksmanship","position":1,"talents":[]},
			{"id":3,"name":"Survival","position":2,"talents":[]}]}`,
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	data, err := trees.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Versions()) != 1 {
		t.Fatalf("versions = %v, want [1.60.1.70009]", data.Versions())
	}
	return data
}

// linkWithExport is this file's shared setup: an account with one character whose
// addon_exports row carries export, source and captured_at exactly as PutExports
// (the addon sync) or the Battle.net import would have written them.
func linkWithExport(t *testing.T, s *Store, userID int64, key, region, ruleset, name, class, source, export string) {
	t.Helper()
	ctx := context.Background()
	if err := s.LinkCharacter(ctx, userID, Character{
		Key: key, Region: region, Ruleset: ruleset, Name: name, Class: class,
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, $3, $4, $5, $6, $7, now(), now())`,
		key, userID, region, ruleset, name, export, source); err != nil {
		t.Fatal(err)
	}
}

func TestBuildFieldsDecodesGearTalentsPointsLevelAndRealm(t *testing.T) {
	s := testStore(t)
	s.Trees = writeHunterTreesFixture(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "bow@example.com")
	if err != nil {
		t.Fatal(err)
	}
	export := "FS1:1.60.1.70009:hunter:troll:503200000/0/0:head=12640,chest=11726|level=24|who=Bow:Skyborne"
	linkWithExport(t, s, u.ID, "us/pvp/bow", "us", "pvp", "Bow", "hunter", "addon", export)

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	c := chars[0]
	if c.Build == nil {
		t.Fatal("build = nil")
	}
	wantGear := map[string]int{"head": 12640, "chest": 11726}
	if len(c.Build.Gear) != 2 || c.Build.Gear["head"] != 12640 || c.Build.Gear["chest"] != 11726 {
		t.Fatalf("gear = %v, want %v", c.Build.Gear, wantGear)
	}
	if c.Build.Talents == nil {
		t.Fatal("talents = nil")
	}
	wantTrees := []string{"503200000", "0", "0"}
	if len(c.Build.Talents.Trees) != 3 || c.Build.Talents.Trees[0] != wantTrees[0] {
		t.Fatalf("talents.trees = %v, want %v", c.Build.Talents.Trees, wantTrees)
	}
	wantPoints := []int{10, 0, 0}
	if len(c.Build.Talents.Points) != 3 || c.Build.Talents.Points[0] != wantPoints[0] {
		t.Fatalf("talents.points = %v, want %v", c.Build.Talents.Points, wantPoints)
	}
	if c.Build.Level == nil || *c.Build.Level != 24 {
		t.Fatalf("build.level = %v, want 24", c.Build.Level)
	}
	if c.Build.DataBuild != "1.60.1.70009" {
		t.Fatalf("data_build = %q", c.Build.DataBuild)
	}
	// The character row itself gains level, race, faction and realm since Battle.net
	// never set any of them for this export-only character.
	if c.Level == nil || *c.Level != 24 {
		t.Fatalf("character level = %v, want 24", c.Level)
	}
	if c.Race != "Troll" || c.Faction != "horde" {
		t.Fatalf("race/faction = %q/%q, want Troll/horde", c.Race, c.Faction)
	}
	if c.Realm != "Skyborne" {
		t.Fatalf("realm = %q, want Skyborne", c.Realm)
	}
	if c.Spec != "Beast Mastery" {
		t.Fatalf("spec = %q, want Beast Mastery (the only tree with points)", c.Spec)
	}
}

func TestBuildFieldsBattleNetWinsOverTheExport(t *testing.T) {
	s := testStore(t)
	s.Trees = writeHunterTreesFixture(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "bnet@example.com")
	if err != nil {
		t.Fatal(err)
	}
	key := "us/pvp/bnetchar"
	if err := s.LinkCharacter(ctx, u.ID, Character{
		Key: key, Region: "us", Ruleset: "pvp", Name: "Bnetchar", Class: "hunter",
	}); err != nil {
		t.Fatal(err)
	}
	// Battle.net already set level, race, faction and realm directly on characters;
	// the export's own who=/level= must not override any of them.
	level := 60
	if _, err := s.Pool.Exec(ctx,
		`update characters set level = $2, race = 'Orc', faction = 'horde', realm_name = 'Faerlina'
		 where key = $1`, key, level); err != nil {
		t.Fatal(err)
	}
	export := "FS1:1.60.1.70009:hunter:troll:503200000/0/0:|level=24|who=Bow:Skyborne"
	if _, err := s.Pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, 'us', 'pvp', 'Bnetchar', $3, 'addon', now(), now())`,
		key, u.ID, export); err != nil {
		t.Fatal(err)
	}

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	c := chars[0]
	if c.Level == nil || *c.Level != 60 {
		t.Fatalf("level = %v, want Battle.net's 60, not the export's 24", c.Level)
	}
	if c.Race != "Orc" || c.Faction != "horde" {
		t.Fatalf("race/faction = %q/%q, want Battle.net's Orc/horde", c.Race, c.Faction)
	}
	if c.Realm != "Faerlina" {
		t.Fatalf("realm = %q, want Battle.net's Faerlina, not the export's Skyborne", c.Realm)
	}
	// Build's own fields still come from the export -- the precedence rule is about the
	// character row's gap-filling fields, not about Build at all.
	if c.Build == nil || c.Build.Level == nil || *c.Build.Level != 24 {
		t.Fatalf("build.level = %v, want the export's own 24", c.Build.Level)
	}
}

func TestBuildFieldsOmitsSpecOnATie(t *testing.T) {
	s := testStore(t)
	s.Trees = writeHunterTreesFixture(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "tie@example.com")
	if err != nil {
		t.Fatal(err)
	}
	// "5" and "5" base-36 are each a single point; Beast Mastery and Marksmanship tie.
	export := "FS1:1.60.1.70009:hunter:troll:5/5/0:"
	linkWithExport(t, s, u.ID, "us/pvp/tied", "us", "pvp", "Tied", "hunter", "addon", export)

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	if chars[0].Spec != "" {
		t.Fatalf("spec = %q, want omitted on a tie", chars[0].Spec)
	}
}

func TestBuildFieldsOmitsSpecWhenNoPointsAreSpent(t *testing.T) {
	s := testStore(t)
	s.Trees = writeHunterTreesFixture(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "zero@example.com")
	if err != nil {
		t.Fatal(err)
	}
	linkWithExport(t, s, u.ID, "us/pvp/fresh", "us", "pvp", "Fresh", "hunter", "addon",
		"FS1:1.60.1.70009:hunter:troll:0/0/0:")

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	if chars[0].Spec != "" {
		t.Fatalf("spec = %q, want omitted when no points are spent", chars[0].Spec)
	}
}

func TestBuildFieldsOfAMalformedExportLeavesTheCharacterIntact(t *testing.T) {
	s := testStore(t)
	s.Trees = writeHunterTreesFixture(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "malformed@example.com")
	if err != nil {
		t.Fatal(err)
	}
	capturedAt := time.Now().Truncate(time.Second)
	key := "us/pvp/oldaddon"
	if err := s.LinkCharacter(ctx, u.ID, Character{
		Key: key, Region: "us", Ruleset: "pvp", Name: "Oldaddon", Class: "hunter",
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, 'us', 'pvp', 'Oldaddon', 'not-an-fs1-string', 'addon', $3, now())`,
		key, u.ID, capturedAt); err != nil {
		t.Fatal(err)
	}

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 {
		t.Fatalf("characters = %v, err = %v (GET /v1/me must never fail over a bad export)", chars, err)
	}
	c := chars[0]
	if c.Build == nil || c.Build.Source != "addon" || !c.Build.CapturedAt.Equal(capturedAt) {
		t.Fatalf("build = %+v, want source/captured_at intact", c.Build)
	}
	if c.Build.Gear != nil || c.Build.Talents != nil || c.Build.Level != nil || c.Build.DataBuild != "" {
		t.Fatalf("build = %+v, want every decoded field omitted", c.Build)
	}
	if c.Level != nil || c.Race != "" || c.Faction != "" || c.Realm != "" || c.Spec != "" {
		t.Fatalf("character = %+v, want every gap-filling field left empty", c)
	}
}

// recordSyncAt writes one character_syncs row at a fixed time, the way synclog.Record
// would have when that sync happened.
func recordSyncAt(t *testing.T, s *Store, key, source, outcome string, when time.Time) {
	t.Helper()
	if _, err := s.Pool.Exec(context.Background(),
		`insert into character_syncs (character_key, source, outcome, created_at) values ($1, $2, $3, $4)`,
		key, source, outcome, when); err != nil {
		t.Fatal(err)
	}
}

func TestBuildCarriesMedianSyncGapAndSyncErrorFromHistory(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "cadence@example.com")
	if err != nil {
		t.Fatal(err)
	}
	linkWithExport(t, s, u.ID, "us/pvp/nightly", "us", "pvp", "Nightly", "hunter", "blizzard", "FS1:1.60.1.70009:hunter:troll:0/0/0:")
	base := time.Now().Add(-72 * time.Hour)
	for _, h := range []int{0, 24, 48} {
		recordSyncAt(t, s, "us/pvp/nightly", "blizzard", "ok", base.Add(time.Duration(h)*time.Hour))
	}
	recordSyncAt(t, s, "us/pvp/nightly", "blizzard", "bnet_refresh_failed", base.Add(60*time.Hour))

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 || chars[0].Build == nil {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	b := chars[0].Build
	if b.MedianSyncGapSec == nil || *b.MedianSyncGapSec != 24*3600 {
		t.Fatalf("median_sync_gap_sec = %v, want 86400", b.MedianSyncGapSec)
	}
	if b.SyncError == nil || *b.SyncError != "bnet_refresh_failed" {
		t.Fatalf("sync_error = %v, want bnet_refresh_failed", b.SyncError)
	}
}

func TestBuildSyncFieldsAreNullWithoutHistoryAndAlwaysInTheJSON(t *testing.T) {
	s := testStore(t)
	ctx := context.Background()
	u, err := s.UpsertEmailUser(ctx, "fresh@example.com")
	if err != nil {
		t.Fatal(err)
	}
	linkWithExport(t, s, u.ID, "us/pvp/fresh", "us", "pvp", "Fresh", "hunter", "addon", "not an export")
	recordSyncAt(t, s, "us/pvp/fresh", "addon", "ok", time.Now())
	recordSyncAt(t, s, "us/pvp/fresh", "addon", "ok", time.Now().Add(time.Hour))

	chars, err := s.Characters(ctx, u.ID)
	if err != nil || len(chars) != 1 || chars[0].Build == nil {
		t.Fatalf("characters = %v, err = %v", chars, err)
	}
	if chars[0].Build.MedianSyncGapSec != nil || chars[0].Build.SyncError != nil {
		t.Fatalf("build = %+v, want both nil (two syncs is not enough, and nothing failed)", chars[0].Build)
	}
	raw, err := json.Marshal(chars[0].Build)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"median_sync_gap_sec":null`, `"sync_error":null`} {
		if !strings.Contains(string(raw), want) {
			t.Fatalf("build JSON = %s, want it to contain %s", raw, want)
		}
	}
}
