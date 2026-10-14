// api/internal/dataaddon/job_test.go
package dataaddon

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

type recordingLogger struct {
	warnings []string
}

func (l *recordingLogger) Info(string, ...any) {}
func (l *recordingLogger) Warn(msg string, args ...any) {
	l.warnings = append(l.warnings, msg)
}
func (l *recordingLogger) Error(string, ...any) {}

func TestRunSkipsAnUnparseablePlayerKeyAndCountsIt(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)
	foughtAt := now.Add(-2 * 24 * time.Hour)

	if _, err := pool.Exec(ctx,
		`insert into reports (id, visibility, status, created_at) values ('dataaddon-job-bad', 'public', 'complete', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from reports where id = 'dataaddon-job-bad'`) })

	// rating_scores is partitioned by month on fought_at (migration 0022);
	// ensure the partition exists before inserting into it, the same way
	// rating.Store.RateFight does before its own insert.
	if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
		   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
		 values ('dataaddon-job-bad', 1, 'not-a-valid-key', 'x', 1, false, 50, 50, false, '[]', 'test', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { pool.Exec(ctx, `delete from rating_scores where report_id = 'dataaddon-job-bad'`) })

	log := &recordingLogger{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: now, Log: log,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.SkippedRows != 1 {
		t.Errorf("skipped = %d, want 1", result.SkippedRows)
	}
	if len(log.warnings) == 0 {
		t.Error("the skip should be logged")
	}
}

func TestRunSkipsTheUploadAndLogsWhyWhenNoBucketIsConfigured(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	log := &recordingLogger{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Build: "1.60.1.69893", Now: time.Now().UTC(), Log: log,
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Files["Data.lua"] == 0 {
		t.Error("Data.lua should still be rendered (and its size logged) even with no bucket")
	}
	found := false
	for _, w := range log.warnings {
		if w == "dataaddon" {
			found = true
		}
	}
	if !found {
		t.Error("the skipped upload should be logged")
	}
}

func TestRunUploadsEveryRenderedFileAndAManifest(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	up := &FakeUploader{}
	result, err := Run(ctx, Deps{
		Store: &Store{Pool: pool}, Upload: up, Bucket: "my-bucket",
		Build: "1.60.1.69893", Now: time.Now().UTC(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Files["Data.lua"] == 0 {
		t.Fatal("Data.lua should be rendered")
	}
	var sawData, sawManifest bool
	for _, c := range up.Calls {
		if c.Bucket != "my-bucket" {
			t.Errorf("bucket = %q, want my-bucket", c.Bucket)
		}
		if c.Object == "data-addon/Data.lua" {
			sawData = true
		}
		if c.Object == "data-addon/manifest.txt" {
			sawManifest = true
			if string(c.Data) != "Data.lua\n" {
				t.Errorf("manifest = %q, want \"Data.lua\\n\"", c.Data)
			}
		}
	}
	if !sawData || !sawManifest {
		t.Errorf("calls = %+v, want Data.lua and manifest.txt uploaded", up.Calls)
	}
}

// Name-Realm carries no region: two characters sharing one would each read as the other in
// game, so neither is written and both are counted.
func TestRunLeavesOutCharactersThatShareANameRealmKey(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	now := time.Date(2026, 9, 22, 4, 0, 0, 0, time.UTC)
	foughtAt := now.Add(-2 * 24 * time.Hour)
	if _, err := pool.Exec(ctx,
		`insert into reports (id, visibility, status, created_at) values ('dataaddon-job-twins', 'public', 'complete', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		pool.Exec(ctx, `delete from rating_scores where report_id = 'dataaddon-job-twins'`)
		pool.Exec(ctx, `delete from reports where id = 'dataaddon-job-twins'`)
	})
	if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	for i, key := range []string{"us/normal/dataaddon-twin", "eu/normal/dataaddon-twin"} {
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
			 values ('dataaddon-job-twins', $1, $2, 'Twin-ClassicBetaPvP', 1, false, 50, 50, false, '[]', 'test', $3)`,
			i+1, key, foughtAt); err != nil {
			t.Fatal(err)
		}
	}
	log := &recordingLogger{}
	result, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Now: now, Log: log})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(result.Rendered["Data.lua"]), `"Twin-ClassicBetaPvP"`) {
		t.Errorf("a shared Name-Realm key was written:\n%s", result.Rendered["Data.lua"])
	}
	if result.SkippedRows != 2 || len(log.warnings) == 0 {
		t.Errorf("skipped = %d with %d warnings, want both twins counted and logged", result.SkippedRows, len(log.warnings))
	}
}

// One character whose score cannot be read (here: its spec's best-in-slot file is
// unreadable, so ScoreV2 errors) is logged and counted, and the night's file is still
// written with everyone else; only past maxCharacterFailures does the run stop.
func TestRunWritesTheFileWhenOneScoreFails(t *testing.T) {
	pool := testPool(t)
	data := useRepoData(t)
	ctx := context.Background()
	now := time.Now().UTC()
	foughtAt := now.Add(-2 * 24 * time.Hour)

	// The repository's weights, and a warrior-arms bis file that does not decode.
	const build = "1.60.1.70291"
	dir := t.TempDir()
	weights, err := os.ReadFile(filepath.Join(repoBuilds, build, "stat-weights.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, build, "bis"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, build, "stat-weights.json"), weights, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, build, "bis", "warrior-arms.json"), []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	sims.SetGearData(data, dir)
	t.Cleanup(func() { sims.SetGearData(data, repoBuilds) })

	cleanup := func() {
		bg := context.Background()
		pool.Exec(bg, `delete from rating_scores where report_id = 'dav2fail-report'`)
		pool.Exec(bg, `delete from reports where id = 'dav2fail-report'`)
		pool.Exec(bg, `delete from addon_exports where character_key like '%/dav2fail-%'`)
		pool.Exec(bg, `delete from characters where key like '%/dav2fail-%'`)
		pool.Exec(bg, `delete from users where email = 'dav2fail@example.com'`)
	}
	cleanup()
	t.Cleanup(cleanup)
	var userID int64
	if err := pool.QueryRow(ctx, `insert into users (email) values ('dav2fail@example.com') returning id`).Scan(&userID); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ key, name string }{
		{"us/normal/dav2fail-good", "Goodscore"}, {"us/normal/dav2fail-bad", "Badscore"},
	} {
		if _, err := pool.Exec(ctx, `insert into characters (key, region, ruleset, name, realm_name, user_id)
			values ($1, 'us', 'normal', $2, 'Classic Beta PvE 2', $3)`, c.key, c.name, userID); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := pool.Exec(ctx, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at)
		values ('us/normal/dav2fail-bad', $1, 'us', 'normal', 'Badscore', $2, now())`, userID, goldenGear); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx,
		`insert into reports (id, visibility, status, created_at) values ('dav2fail-report', 'public', 'complete', $1)`,
		foughtAt); err != nil {
		t.Fatal(err)
	}
	if err := db.EnsureRatingsPartition(ctx, pool, foughtAt); err != nil {
		t.Fatal(err)
	}
	for i, c := range []struct{ key, class, spec string }{
		{"us/normal/dav2fail-good", "Mage", "Fire"}, {"us/normal/dav2fail-bad", "Warrior", "Arms"},
	} {
		if _, err := pool.Exec(ctx,
			`insert into rating_scores (report_id, fight_index, player_key, player_name, class, spec, encounter_id,
			   kill, overall, overall_uncapped, overall_capped, components, model_version, fought_at)
			 values ('dav2fail-report', $1, $2, $2, $3, $4, 1, false, 50, 50, false, '[]', 'test', $5)`,
			i+1, c.key, c.class, c.spec, foughtAt); err != nil {
			t.Fatal(err)
		}
	}

	log := &recordingLogger{}
	result, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Trees: data, Now: now, Log: log})
	if err != nil {
		t.Fatalf("one failing score stopped the run: %v", err)
	}
	file := string(result.Rendered["Data.lua"])
	if !strings.Contains(file, `["Goodscore-ClassicBetaPvE2"] = {`) {
		t.Errorf("the character whose score reads is missing:\n%s", file)
	}
	if strings.Contains(file, "Badscore") {
		t.Errorf("the character whose score failed was written:\n%s", file)
	}
	if result.SkippedRows != 1 {
		t.Errorf("skipped = %d, want 1 (the failed score)", result.SkippedRows)
	}

	saved := maxCharacterFailures
	maxCharacterFailures = 0
	t.Cleanup(func() { maxCharacterFailures = saved })
	if _, err := Run(ctx, Deps{Store: &Store{Pool: pool}, Trees: data, Now: now, Log: log}); err == nil {
		t.Error("past maxCharacterFailures the run should stop")
	}
	cancelled, cancel := context.WithCancel(ctx)
	cancel()
	maxCharacterFailures = saved
	if _, err := Run(cancelled, Deps{Store: &Store{Pool: pool}, Trees: data, Now: now, Log: log}); err == nil {
		t.Error("a cancelled run should stop")
	}
}
