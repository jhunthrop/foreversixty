package reports

import (
	"context"
	"fmt"
	"testing"
)

// seedGuildCharacter puts characterKey in guildID for userID, verified or not, inserting
// the user first if it is not there.
func (h *harness) seedGuildCharacter(guildID, userID int64, characterKey string, verified bool) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into users (id, email) values ($1, $2) on conflict (id) do nothing`,
		userID, fmt.Sprintf("u%d@example.com", userID)); err != nil {
		h.t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into guild_characters (guild_id, character_key, user_id, verified_at)
		 values ($1, $2, $3, case when $4 then now() end)`,
		guildID, characterKey, userID, verified); err != nil {
		h.t.Fatal(err)
	}
}

// reportMatchedGuild reads a report's matched_guild_id straight from the row.
func (h *harness) reportMatchedGuild(id string) *int64 {
	h.t.Helper()
	var g *int64
	if err := h.store.Pool.QueryRow(h.t.Context(), `select matched_guild_id from reports where id = $1`, id).Scan(&g); err != nil {
		h.t.Fatal(err)
	}
	return g
}

// reportGuild reads a report's guild_id straight from the row.
func (h *harness) reportGuild(id string) *int64 {
	h.t.Helper()
	rep, err := h.store.Get(context.Background(), id)
	if err != nil {
		h.t.Fatal(err)
	}
	return rep.GuildID
}

func sameGuild(got *int64, want int64) bool { return got != nil && *got == want }

// createReport logs as us/hardcore/baelgrim, the harness owner's character.
const harnessLoggingKey = "us/hardcore/baelgrim"

func TestAttributeGuildPicksTheLoggingCharactersVerifiedGuild(t *testing.T) {
	h := newHarness(t)
	a, b := h.seedGuild("Alpha", "hardcore"), h.seedGuild("Beta", "hardcore")
	h.seedGuildCharacter(a, h.owner, harnessLoggingKey, true)
	// The owner also has a verified alt elsewhere: the logging character still decides.
	h.seedGuildCharacter(b, h.owner, "us/hardcore/altbael", true)
	id := h.createReport(Public)

	got, err := h.store.AttributeGuild(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if !sameGuild(got, a) || !sameGuild(h.reportMatchedGuild(id), a) || h.reportGuild(id) != nil {
		t.Fatalf("attributed %v, stored %v, want guild %d", got, h.reportGuild(id), a)
	}
}

func TestAttributeGuildFallsBackToTheOwnersOnlyVerifiedGuild(t *testing.T) {
	h := newHarness(t)
	a, b := h.seedGuild("Alpha", "hardcore"), h.seedGuild("Beta", "hardcore")
	// The logging character is in Beta but unverified; the owner's one verified guild is Alpha.
	h.seedGuildCharacter(b, h.owner, harnessLoggingKey, false)
	h.seedGuildCharacter(a, h.owner, "us/hardcore/altbael", true)
	id := h.createReport(Public)

	got, err := h.store.AttributeGuild(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if !sameGuild(got, a) || !sameGuild(h.reportMatchedGuild(id), a) || h.reportGuild(id) != nil {
		t.Fatalf("attributed %v, want the owner's only verified guild %d", got, a)
	}
}

func TestAttributeGuildLeavesAnAmbiguousOwnerUnattributed(t *testing.T) {
	h := newHarness(t)
	a, b := h.seedGuild("Alpha", "hardcore"), h.seedGuild("Beta", "hardcore")
	h.seedGuildCharacter(a, h.owner, "us/hardcore/altone", true)
	h.seedGuildCharacter(b, h.owner, "us/hardcore/alttwo", true)
	h.seedGuildCharacter(b, h.owner, harnessLoggingKey, false)
	id := h.createReport(Public)

	got, err := h.store.AttributeGuild(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if got != nil || h.reportMatchedGuild(id) != nil || h.reportGuild(id) != nil {
		t.Fatalf("attributed %v, stored %v, want none for an owner in two guilds", got, h.reportGuild(id))
	}
}

func TestAttributeGuildNeverOverridesAGuildSetByHand(t *testing.T) {
	h := newHarness(t)
	a, b := h.seedGuild("Alpha", "hardcore"), h.seedGuild("Beta", "hardcore")
	h.seedGuildCharacter(a, h.owner, harnessLoggingKey, true)
	id := h.createReport(Public)
	if _, err := h.store.Pool.Exec(t.Context(), `update reports set guild_id = $2 where id = $1`, id, b); err != nil {
		t.Fatal(err)
	}

	got, err := h.store.AttributeGuild(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if !sameGuild(got, b) || !sameGuild(h.reportGuild(id), b) {
		t.Fatalf("attributed %v, stored %v, want the hand-set guild %d", got, h.reportGuild(id), b)
	}
}

func TestSetLogZoneWritesWhenUnsetOrDifferent(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	zone := func() *string {
		t.Helper()
		var z *string
		if err := h.store.Pool.QueryRow(t.Context(), `select log_zone from reports where id = $1`, id).Scan(&z); err != nil {
			t.Fatal(err)
		}
		return z
	}
	if zone() != nil {
		t.Fatal("a new report has a log zone")
	}
	for _, want := range []string{"Onyxia's Lair", "Onyxia's Lair", "Blackrock Spire"} {
		if err := h.store.SetLogZone(t.Context(), id, want); err != nil {
			t.Fatal(err)
		}
		if z := zone(); z == nil || *z != want {
			t.Fatalf("log_zone = %v, want %q", z, want)
		}
	}
	if err := h.store.SetLogZone(t.Context(), id, ""); err != nil {
		t.Fatal(err)
	}
	if z := zone(); z == nil || *z != "Blackrock Spire" {
		t.Fatalf("an empty zone overwrote log_zone: %v", z)
	}
}
