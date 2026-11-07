package sims

import (
	"context"
	"testing"
)

func importExport(t *testing.T, h *harness, key, export string) {
	t.Helper()
	ctx := context.Background()
	if _, err := h.store.Pool.Exec(ctx, `insert into characters (key, region, ruleset, name, class, level, user_id)
		values ($1, 'us', 'normal', 'Go Hogs', 'warrior', 20, $2) on conflict (key) do nothing`, key, h.owner); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(ctx, `insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		values ($1, $2, 'us', 'normal', 'Go Hogs', $3, 'blizzard', now(), now())
		on conflict (character_key) do update set export = excluded.export`, key, h.owner, export); err != nil {
		t.Fatal(err)
	}
}

// The auto-sim on import: one bare run per gear change, for any account, capped per day.
func TestImportSimRunsOncePerGearChange(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	const key = "us/normal/go-hogs"
	importExport(t, h, key, "FS1:1.60.1.70790:warrior:orc:0/0505/0:head=16866,chest=11726|level=20")
	for i, want := range []string{ImportQueued, ImportUnchanged} {
		if got, err := h.service.RequestFromImport(ctx, h.owner, key); err != nil || got != want {
			t.Fatalf("request %d: %q %v, want %q", i, got, err, want)
		}
	}
	importExport(t, h, key, "FS1:1.60.1.70790:warrior:orc:0/0505/0:head=16866,chest=11726,legs=10410|level=20")
	if got, _ := h.service.RequestFromImport(ctx, h.owner, key); got != ImportQueued {
		t.Fatalf("a gear change: %q", got)
	}
	if len(h.jobs.Calls) != 2 {
		t.Fatalf("job starts: %v", h.jobs.Calls)
	}
	var kind, preset, spec string
	var band int
	if err := h.store.Pool.QueryRow(ctx, `select kind, preset, spec, band from character_sims where character_key = $1 limit 1`, key).
		Scan(&kind, &preset, &spec, &band); err != nil || kind != "import" || preset != ImportPreset || spec != "warrior-fury" || band != 20 {
		t.Fatalf("character sim: %s %s %s %d %v", kind, preset, spec, band, err)
	}
	if got, _ := h.service.RequestFromImport(ctx, h.owner, "us/normal/nobody"); got != ImportNoGear {
		t.Fatalf("no export: %q", got)
	}
	if _, err := h.store.Pool.Exec(ctx, `insert into sims (id, user_id, spec, engine_version, lane, dps_mean, dps_error, iterations, result, state)
		select 'cap' || n, $1, 'warrior-fury', 'e', 'server', 0, 0, 1, '{}', 'done' from generate_series(1, 10) n`, h.owner); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(ctx, `insert into character_sims (sim_id, character_key, user_id, kind, spec, role, band, preset, gear_hash)
		select 'cap' || n, 'us/normal/other', $1, 'import', 'warrior-fury', 'dps', 20, 'bare', n::text from generate_series(1, 10) n`, h.owner); err != nil {
		t.Fatal(err)
	}
	importExport(t, h, key, "FS1:1.60.1.70790:warrior:orc:0/0505/0:head=16866|level=20")
	if got, _ := h.service.RequestFromImport(ctx, h.owner, key); got != ImportCapped {
		t.Fatalf("the 13th run of the day: %q", got)
	}
}
