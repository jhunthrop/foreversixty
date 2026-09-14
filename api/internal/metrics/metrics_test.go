package metrics

import (
	"errors"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
)

// fixtureRows is what the ingest derives from the fixture bundle's own
// Parquet: the rows a correct companion would have posted.
func fixtureRows(t *testing.T) []Row {
	t.Helper()
	f, err := engine.NewFixture("rpt")
	if err != nil {
		t.Fatal(err)
	}
	events, err := parquet.Unmarshal(f.Parquet)
	if err != nil {
		t.Fatal(err)
	}
	fi, sum, err := engine.Rebuild(1, engine.Header{
		EncounterID: 9001, Name: "Warden Kelthas", Difficulty: 8, Size: 5, Kill: true,
	}, events)
	if err != nil {
		t.Fatal(err)
	}
	return Derive(fi, sum)
}

func TestDeriveOneRowPerPlayerSortedByGUID(t *testing.T) {
	rows := fixtureRows(t)
	if len(rows) != 3 {
		t.Fatalf("rows = %d, want one per player", len(rows))
	}
	for i := 1; i < len(rows); i++ {
		if rows[i-1].PlayerGUID >= rows[i].PlayerGUID {
			t.Fatalf("rows are not sorted by player: %v", rows)
		}
	}
	for _, r := range rows {
		if r.EncounterID != 9001 || r.Difficulty != 8 || r.Size != 5 || !r.Kill {
			t.Fatalf("row %+v does not carry the fight header", r)
		}
		if r.DurationMS != 30000 {
			t.Fatalf("duration = %d, want the fixture's 30000", r.DurationMS)
		}
	}
}

func TestDeriveSplitsTheRolesTheEngineAssigned(t *testing.T) {
	roles := map[string]string{}
	for _, r := range fixtureRows(t) {
		roles[r.Role] = r.PlayerGUID
	}
	for _, want := range []string{"dps", "healer", "tank"} {
		if roles[want] == "" {
			t.Fatalf("roles = %v, want one of each of dps, healer, tank", roles)
		}
	}
}

func TestCompareAcceptsRowsThatMatch(t *testing.T) {
	rows := fixtureRows(t)
	if err := Compare(rows, rows); err != nil {
		t.Fatalf("identical rows should verify: %v", err)
	}
}

func TestCompareAcceptsFloatsWithinHalfAPercent(t *testing.T) {
	rows := fixtureRows(t)
	posted := append([]Row{}, rows...)
	for i := range posted {
		posted[i].MetricDPS *= 1.004
		posted[i].MetricHPS *= 1.004
	}
	if err := Compare(posted, rows); err != nil {
		t.Fatalf("a 0.4%% drift should verify: %v", err)
	}
}

// damageDealer is the index of a fixture row with a non-zero DPS, so a
// test that inflates a metric actually changes it.
func damageDealer(t *testing.T, rows []Row) int {
	t.Helper()
	for i, r := range rows {
		if r.MetricDPS > 0 {
			return i
		}
	}
	t.Fatal("the fixture has no row with damage")
	return 0
}

func TestCompareRejectsAnInflatedMetric(t *testing.T) {
	rows := fixtureRows(t)
	posted := append([]Row{}, rows...)
	posted[damageDealer(t, rows)].MetricDPS *= 2
	err := Compare(posted, rows)
	if err == nil {
		t.Fatal("a doubled metric must be rejected")
	}
	if got := err.Error(); got == "" {
		t.Fatal("the mismatch must name the field")
	}
}

func TestCompareRejectsMissingAndExtraRows(t *testing.T) {
	rows := fixtureRows(t)
	if err := Compare(rows[:2], rows); err == nil {
		t.Fatal("a missing row must be rejected")
	}
	swapped := append([]Row{}, rows...)
	swapped[0].PlayerGUID = "Player-4184-0000FFFF"
	if err := Compare(swapped, rows); err == nil {
		t.Fatal("a row for a player who was not there must be rejected")
	}
}

func TestCompareIgnoresClassSpecAndIlvlTheRebuildCannotSee(t *testing.T) {
	rows := fixtureRows(t)
	posted := append([]Row{}, rows...)
	for i := range posted {
		posted[i].Class, posted[i].Spec, posted[i].Ilvl = "Warrior", "Fury", 183
	}
	if err := Compare(posted, rows); err != nil {
		t.Fatalf("a rebuild that derived no class must not reject one: %v", err)
	}
	derived := append([]Row{}, rows...)
	derived[0].Class = "Mage"
	if err := Compare(posted, derived); err == nil {
		t.Fatal("a class the rebuild did derive must be enforced")
	}
}

func TestCompareRejectsChangedIntegers(t *testing.T) {
	rows := fixtureRows(t)
	for _, mutate := range []func(*Row){
		func(r *Row) { r.DamageTaken += 1 },
		func(r *Row) { r.ActiveMS += 1 },
		func(r *Row) { r.Deaths += 1 },
		func(r *Row) { r.DurationMS += 1 },
		func(r *Row) { r.Kill = !r.Kill },
		func(r *Row) { r.Role = "tank-ish" },
		func(r *Row) { r.Name = "Someone Else" },
	} {
		posted := append([]Row{}, rows...)
		mutate(&posted[0])
		if err := Compare(posted, rows); err == nil {
			t.Fatalf("a changed integer field must be rejected: %+v", posted[0])
		}
	}
}

func TestCompareReportsAMismatchThatNamesTheFieldApartFromTheValue(t *testing.T) {
	rows := fixtureRows(t)
	posted := append([]Row{}, rows...)
	i := damageDealer(t, rows)
	posted[i].MetricDPS *= 2

	var m *Mismatch
	if err := Compare(posted, rows); !errors.As(err, &m) {
		t.Fatalf("err = %v, want a *Mismatch", err)
	}
	if m.Field != "metric_dps" {
		t.Fatalf("field = %q, want metric_dps and nothing else", m.Field)
	}
	if m.PlayerGUID != rows[i].PlayerGUID {
		t.Fatalf("guid = %q, want %q", m.PlayerGUID, rows[i].PlayerGUID)
	}
	// The field is safe to hand back to the sender; the message, which
	// carries the engine's own number, is for the server's log.
	if strings.Contains(m.Field, "the events give") {
		t.Fatalf("the field must not carry the engine's value: %q", m.Field)
	}
	if !strings.HasPrefix(m.Error(), m.PlayerGUID+": ") || !strings.Contains(m.Error(), "the events give") {
		t.Fatalf("message = %q, want the player and the detail", m.Error())
	}

	// A row-count mismatch belongs to no player, so it stands alone.
	if err := Compare(rows[:2], rows); !errors.As(err, &m) || m.Field != "rows" {
		t.Fatalf("err = %v, want a rows mismatch", err)
	}
	if m.Error() != m.Detail {
		t.Fatalf("message = %q, want the detail alone", m.Error())
	}
}
