package guilds

import (
	"errors"
	"testing"
	"time"
)

func TestRaidReadinessNamesARaid(t *testing.T) {
	w, first := readinessRaidWorld(t)
	ctx := t.Context()
	later := time.Now().Add(14 * 24 * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
	second := w.createRaids(t, raidBody(later, 0))[0]

	got, err := w.h.svc.RaidReadiness(ctx, w.gid, w.officer, second.ID)
	if err != nil || got.Raid == nil || got.Raid.ID != second.ID || got.Signed != 0 {
		t.Fatalf("second raid = %+v %v", got, err)
	}
	got, err = w.h.svc.RaidReadiness(ctx, w.gid, w.officer, first.ID)
	if err != nil || got.Raid == nil || got.Raid.ID != first.ID || got.Signed != 3 {
		t.Fatalf("first raid = %+v %v", got, err)
	}

	other := seedGuild(t, w.h.pool, "Elsewhere")
	if _, err := w.h.pool.Exec(ctx, `update guild_raids set guild_id = $1 where id = $2`, other, second.ID); err != nil {
		t.Fatal(err)
	}
	if got, err := w.h.svc.RaidReadiness(ctx, w.gid, w.officer, second.ID); err != nil || got.Raid != nil {
		t.Fatalf("another guild's raid = %+v %v", got, err)
	}
	if _, err := w.h.pool.Exec(ctx, `update guild_raids set status = 'cancelled' where id = $1`, first.ID); err != nil {
		t.Fatal(err)
	}
	if got, err := w.h.svc.RaidReadiness(ctx, w.gid, w.officer, first.ID); err != nil || got.Raid != nil {
		t.Fatalf("a cancelled raid = %+v %v", got, err)
	}
	if _, err := w.h.svc.RaidReadiness(ctx, w.gid, w.member, first.ID); !errors.Is(err, ErrNotOfficer) {
		t.Fatalf("member = %v", err)
	}
}
