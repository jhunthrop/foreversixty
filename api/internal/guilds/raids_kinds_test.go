package guilds

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

// kindBody is a create body for zone at start, with a title.
func kindBody(zone, title, start string) string {
	return fmt.Sprintf(`{"zone":%q,"title":%q,"starts_at":%q,"duration_minutes":120,"notes":"",`+
		`"timezone":"America/Chicago","repeat_weeks":0,"post_to_discord":false}`, zone, title, start)
}

func kindStart(hours int) string {
	return time.Now().Add(time.Duration(hours) * time.Hour).UTC().Truncate(time.Minute).Format(time.RFC3339)
}

// The officer route's rate limit (raids_signup_handler.go) would stop a burst of real posts, so
// the cap and the runs both go through the store, which is where the cap lives.
func TestThirteenRunsAndEventsPostPastTheRaidCap(t *testing.T) {
	f := newRaidFixture(t)
	start := time.Now().Add(24 * time.Hour)
	for i := 0; i < RaidCap; i++ {
		f.createRaid(t, start.Add(time.Duration(i)*time.Hour))
	}
	if _, err := f.store.CreateRaid(t.Context(), f.input(start)); !errors.Is(err, ErrRaidCap) {
		t.Fatalf("13th raid: err = %v, want ErrRaidCap", err)
	}
	zones := []string{"dungeon:the-deadmines", "event:guild", "event:guild-10"}
	for i := 0; i < 13; i++ {
		in := f.input(start.Add(time.Duration(100+i) * time.Hour))
		in.Zone = zones[i%len(zones)]
		if _, err := f.store.CreateRaid(t.Context(), in); err != nil {
			t.Fatalf("run %d at the raid cap: %v", i, err)
		}
	}
	if _, err := f.store.CreateRaid(t.Context(), f.input(start)); !errors.Is(err, ErrRaidCap) {
		t.Fatalf("runs lifted the raid cap: err = %v", err)
	}
}

func TestARepeatOfRunsIsNotHeldByTheCap(t *testing.T) {
	f := newRaidFixture(t)
	start := time.Now().Add(24 * time.Hour)
	for i := 0; i < RaidCap; i++ {
		f.createRaid(t, start.Add(time.Duration(i)*time.Hour))
	}
	in := f.input(start)
	in.Zone = "dungeon:the-deadmines"
	if _, err := f.store.CreateRaids(t.Context(), f.guildID, "", []RaidInput{in, in, in, in}); err != nil {
		t.Fatalf("a repeat of dungeon runs at the raid cap: %v", err)
	}
}

func TestAnEventKeyAnswersItsSizeAndName(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaids(t, kindBody("event:guild-10", "", kindStart(48)))[0]
	if raid.Size == nil || *raid.Size != 10 {
		t.Fatalf("size = %v, want 10", raid.Size)
	}
	if raid.ZoneKind != ZoneEvent || raid.ZoneName != "Guild event" || raid.Title != "Guild event" {
		t.Fatalf("view = %q %q %q", raid.ZoneKind, raid.ZoneName, raid.Title)
	}
	open := w.createRaids(t, kindBody("event:guild", "PvP premade", kindStart(49)))[0]
	if open.Size != nil || open.Title != "PvP premade" {
		t.Fatalf("no-limit event = %v %q", open.Size, open.Title)
	}
}

func TestADungeonRunCarriesItsKindSizeAndKey(t *testing.T) {
	w := newRaidWorld(t, true)
	ubrs := w.createRaids(t, kindBody("dungeon:upper-blackrock-spire", "", kindStart(48)))[0]
	if ubrs.ZoneKind != ZoneDungeon || ubrs.Size == nil || *ubrs.Size != 10 {
		t.Fatalf("UBRS = %q %v", ubrs.ZoneKind, ubrs.Size)
	}
	if ubrs.KeyItem == nil || ubrs.KeyItem.Name != "Seal of Ascension" {
		t.Fatalf("UBRS key = %+v", ubrs.KeyItem)
	}
	dm := w.createRaids(t, kindBody("dungeon:the-deadmines", "", kindStart(49)))[0]
	if dm.Title != "The Deadmines" || dm.KeyItem != nil || dm.Size == nil || *dm.Size != 5 {
		t.Fatalf("Deadmines = %q %+v %v", dm.Title, dm.KeyItem, dm.Size)
	}
}

// validateRaid is the rule and takes its clock, so this holds after 9 December too.
func TestARaidCannotBePostedBeforeItOpens(t *testing.T) {
	withOpenDays(t)
	now := time.Date(2026, 11, 5, 12, 0, 0, 0, time.UTC)
	req := func(zone, start string) raidRequest {
		return raidRequest{
			Zone: zone, StartsAt: start, DurationMinutes: 180, Timezone: "America/Chicago",
		}
	}
	_, fields := validateRaid(req("raid:onyxias-lair", "2026-11-20T02:00:00Z"), now, nil, true)
	if got := fields["starts_at"]; got != "Onyxia's Lair opens Wed 9 Dec. Pick that day or later." {
		t.Fatalf("starts_at = %q", got)
	}
	_, fields = validateRaid(req("raid:barrow-deeps", "2026-12-09T02:00:00Z"), now, nil, true)
	if got := fields["starts_at"]; got != "Barrow Deeps opens Wed 9 Dec. Pick that day or later." {
		t.Fatalf("8 Dec evening in Chicago: starts_at = %q", got)
	}
	if _, fields := validateRaid(req("raid:onyxias-lair", "2026-12-10T02:00:00Z"), now, nil, true); fields != nil {
		t.Fatalf("opening night refused: %v", fields)
	}
	if _, fields := validateRaid(req("dungeon:the-deadmines", "2026-11-20T02:00:00Z"), now, nil, true); fields != nil {
		t.Fatalf("a dungeon run refused: %v", fields)
	}
}

func TestAnUnknownZoneAsksForADungeonEventOrRaid(t *testing.T) {
	now := time.Date(2026, 11, 5, 12, 0, 0, 0, time.UTC)
	req := raidRequest{Zone: "", StartsAt: "2026-11-20T02:00:00Z", DurationMinutes: 180, Timezone: "UTC"}
	_, fields := validateRaid(req, now, nil, true)
	if got := fields["zone"]; got != "Choose a dungeon, event or raid." {
		t.Fatalf("zone = %q", got)
	}
}

func TestZoneViewsCarryTheFormsFacts(t *testing.T) {
	withOpenDays(t)
	byKey := map[string]ZoneView{}
	for _, z := range zoneViews() {
		byKey[z.Key] = z
	}
	ony := byKey["raid:onyxias-lair"]
	if ony.Kind != ZoneRaid || ony.Opens == nil || !ony.Opens.Equal(time.Date(2026, 12, 9, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("Onyxia = %+v", ony)
	}
	raw, err := json.Marshal(byKey["dungeon:the-deadmines"])
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got["kind"] != ZoneDungeon || got["opens"] != nil || got["key_item"] != nil {
		t.Fatalf("Deadmines = %s", raw)
	}
	if got["min_level"].(float64) != 17 || got["max_level"].(float64) != 26 || got["size"].(float64) != 5 {
		t.Fatalf("Deadmines facts = %s", raw)
	}
}

func TestADungeonRunNeverMovesTheReadinessChip(t *testing.T) {
	f := newRaidFixture(t)
	ctx := t.Context()
	now := time.Now()
	run := f.input(now.Add(24 * time.Hour))
	run.Zone = "dungeon:the-deadmines"
	dungeon, err := f.store.CreateRaid(ctx, run)
	if err != nil {
		t.Fatal(err)
	}
	if chip, err := f.store.readinessNextRaid(ctx, f.guildID, now); err != nil || chip != nil {
		t.Fatalf("only a dungeon planned: chip = %+v err = %v", chip, err)
	}
	raid := f.createRaid(t, now.Add(48*time.Hour))
	chip, err := f.store.readinessNextRaid(ctx, f.guildID, now)
	if err != nil || chip == nil || chip.ID != raid.ID {
		t.Fatalf("chip = %+v err = %v, want the raid %d", chip, err, raid.ID)
	}
	band, err := f.store.NextRaidFor(ctx, f.guildID, f.userID, nil, now)
	if err != nil || band == nil || band.ID != dungeon.ID || band.ZoneKind != ZoneDungeon {
		t.Fatalf("band = %+v err = %v, want the dungeon run", band, err)
	}
}
