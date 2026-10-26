package parse

import (
	"context"
	"errors"
	"sync"
	"testing"
)

// fakeCompleted records every report the completion hook was told of.
type fakeCompleted struct {
	mu  sync.Mutex
	ids []string
}

func (f *fakeCompleted) ReportCompleted(_ context.Context, reportID string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.ids = append(f.ids, reportID)
}

func (f *fakeCompleted) all() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string{}, f.ids...)
}

func TestACompletedUploadLearnsItsZoneAndGuildThenCallsTheHook(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	ctx := context.Background()
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		t.Fatal(err)
	}
	var guildID int64
	if err := d.Reports.Pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Olympus') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := d.Reports.Pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, verified_at) values ($1, $2, $3, now())`,
		guildID, *rep.LoggingCharacter, *rep.OwnerID); err != nil {
		t.Fatal(err)
	}
	hook := &fakeCompleted{}
	d.Completed = hook

	if err := Report(ctx, d, reportID); err != nil {
		t.Fatal(err)
	}
	var zone *string
	var gid *int64
	if err := d.Reports.Pool.QueryRow(ctx,
		`select log_zone, matched_guild_id from reports where guild_id is null and id = $1`, reportID).Scan(&zone, &gid); err != nil {
		t.Fatal(err)
	}
	if zone == nil || *zone != "Blackrock Spire" {
		t.Fatalf("log_zone = %v, want the fixture log's Blackrock Spire", zone)
	}
	if gid == nil || *gid != guildID {
		t.Fatalf("matched_guild_id = %v, want %d (guild_id null)", gid, guildID)
	}
	if got := hook.all(); len(got) != 1 || got[0] != reportID {
		t.Fatalf("hook calls = %v, want the report once", got)
	}
}

func TestAFailedUploadDoesNotCallTheHook(t *testing.T) {
	d, reportID, objects, _ := fixtureUpload(t, false)
	objects.getErr = errors.New("no such object")
	hook := &fakeCompleted{}
	d.Completed = hook
	if err := Report(context.Background(), d, reportID); err == nil {
		t.Fatal("a missing upload must fail")
	}
	if got := hook.all(); len(got) != 0 {
		t.Fatalf("a failed parse called the hook: %v", got)
	}
}
