// Package metrics is the ranking row the companion sends, the ingest
// verifies, and the rankings store writes: one row per player per boss
// fight.
//
// The wire shape here is the Phase 3 contract's MetricsRow. The engine
// has summary.MetricRow, which is a different thing - one row per player
// per *metric*, carrying the report id and the engine version - so the
// contract's row lives here until the engine gains a summary.MetricsRow
// of this shape and both this package and the companion's copy can be
// deleted in its favour.
package metrics

import (
	"fmt"
	"math"
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Tolerance is how far a float may drift before verification calls it a
// mismatch: half a percent, as the contract sets it. Integers are exact.
const Tolerance = 0.005

// Row is one player's line for one fight. The JSON names are the
// contract's and the companion writes exactly these.
type Row struct {
	PlayerGUID  string  `json:"player_guid"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Spec        string  `json:"spec"`
	Role        string  `json:"role"`
	Ilvl        int64   `json:"ilvl"`
	MetricDPS   float64 `json:"metric_dps"`
	MetricHPS   float64 `json:"metric_hps"`
	DamageTaken int64   `json:"damage_taken"`
	ActiveMS    int64   `json:"active_ms"`
	Deaths      int     `json:"deaths"`
	EncounterID int64   `json:"encounter_id"`
	Difficulty  int64   `json:"difficulty"`
	Size        int64   `json:"size"`
	DurationMS  int64   `json:"duration_ms"`
	Kill        bool    `json:"kill"`
}

// Derive projects a rebuilt fight and summary onto the contract's rows,
// ordered by player GUID so two runs produce the same slice.
func Derive(f fight.Fight, s summary.Summary) []Row {
	out := make([]Row, 0, len(s.Roster))
	for _, r := range s.Roster {
		out = append(out, Row{
			PlayerGUID: r.GUID, Name: r.Name, Class: r.Class, Spec: r.Spec,
			Role: r.Role, Ilvl: r.ItemLevel,
			MetricDPS: r.DPS, MetricHPS: r.HPS, DamageTaken: r.DamageTaken,
			ActiveMS: r.ActiveMS, Deaths: r.Deaths,
			EncounterID: f.EncounterID, Difficulty: f.Difficulty, Size: f.Size,
			DurationMS: s.DurationMS, Kill: f.Kill,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].PlayerGUID < out[j].PlayerGUID })
	return out
}

// Compare reports the first way posted differs from derived, or nil when
// they agree. Integers must match exactly and floats must be within
// Tolerance, except for the three fields a rebuild cannot always see -
// class, spec, and item level - which are checked only when the rebuild
// derived one of its own.
func Compare(posted, derived []Row) error {
	if len(posted) != len(derived) {
		return fmt.Errorf("%d rows posted, the events produce %d", len(posted), len(derived))
	}
	byGUID := make(map[string]Row, len(posted))
	for _, r := range posted {
		byGUID[r.PlayerGUID] = r
	}
	for _, want := range derived {
		got, ok := byGUID[want.PlayerGUID]
		if !ok {
			return fmt.Errorf("no row posted for player %s", want.PlayerGUID)
		}
		if err := compareRow(want.PlayerGUID, got, want); err != nil {
			return err
		}
	}
	return nil
}

func compareRow(guid string, got, want Row) error {
	for _, f := range []struct {
		name      string
		got, want int64
	}{
		{"damage_taken", got.DamageTaken, want.DamageTaken},
		{"active_ms", got.ActiveMS, want.ActiveMS},
		{"deaths", int64(got.Deaths), int64(want.Deaths)},
		{"encounter_id", got.EncounterID, want.EncounterID},
		{"difficulty", got.Difficulty, want.Difficulty},
		{"size", got.Size, want.Size},
		{"duration_ms", got.DurationMS, want.DurationMS},
	} {
		if f.got != f.want {
			return fmt.Errorf("%s: %s is %d, the events give %d", guid, f.name, f.got, f.want)
		}
	}
	for _, f := range []struct {
		name      string
		got, want float64
	}{
		{"metric_dps", got.MetricDPS, want.MetricDPS},
		{"metric_hps", got.MetricHPS, want.MetricHPS},
	} {
		if !within(f.got, f.want) {
			return fmt.Errorf("%s: %s is %g, the events give %g", guid, f.name, f.got, f.want)
		}
	}
	for _, f := range []struct {
		name      string
		got, want string
	}{
		{"name", got.Name, want.Name},
		{"role", got.Role, want.Role},
	} {
		if f.got != f.want {
			return fmt.Errorf("%s: %s is %q, the events give %q", guid, f.name, f.got, f.want)
		}
	}
	// Class, spec, and item level come from COMBATANT_INFO, whose payload
	// the Parquet schema does not carry (see engine.Header), and the
	// companion's registry has seen the whole log where a rebuild sees
	// one fight. So they are checked only when the rebuild derived one of
	// its own; an empty or zero derivation is not evidence of a forgery.
	for _, f := range []struct {
		name      string
		got, want string
	}{
		{"class", got.Class, want.Class},
		{"spec", got.Spec, want.Spec},
	} {
		if f.want != "" && f.got != f.want {
			return fmt.Errorf("%s: %s is %q, the events give %q", guid, f.name, f.got, f.want)
		}
	}
	if want.Ilvl != 0 && got.Ilvl != want.Ilvl {
		return fmt.Errorf("%s: ilvl is %d, the events give %d", guid, got.Ilvl, want.Ilvl)
	}
	if got.Kill != want.Kill {
		return fmt.Errorf("%s: kill is %v, the events give %v", guid, got.Kill, want.Kill)
	}
	return nil
}

// within reports whether got is within Tolerance of want, relative to
// want, with an exact match required when want is zero.
func within(got, want float64) bool {
	if want == 0 {
		return got == 0
	}
	return math.Abs(got-want)/math.Abs(want) <= Tolerance
}
