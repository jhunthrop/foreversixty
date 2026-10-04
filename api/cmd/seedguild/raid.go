// api/cmd/seedguild/raid.go
//
// The four raid nights this seed writes: who fought, which boss (when one is published),
// how long, and whether it died.
//
// Forever's real first raid tier, opening 9 December 2026, is Barrow Deeps, Hyjal Summit
// and Onyxia's Lair - not Molten Core (data/curated/loot/forever-raid-phases.json's own
// notes field: "the first tier opens on 9 December and is Barrow Deeps, Hyjal Summit and
// Onyxia"; Molten Core is one of six later "Era" raids with no announced open date). Of
// the three, only Onyxia's Lair has a published encounter at all
// (logs/engine/mechanics/encounters.json: id 1084 "Onyxia", read from the Forever beta
// client itself). Barrow Deeps and Hyjal Summit are both named explicitly in that same
// file's own "no_client_rows_yet.raids" list - candidate maps exist (2817 Starfall Barrow
// Den, 2832 Nightmare Grove, 2995 Hyjal Crater per forever-raid-phases.json) but nothing
// sourced maps the announced zone names onto those ids, and no encounter or boss name is
// published for either. This seed writes those two nights' pulls with a null encounter_id
// and a generic "Pull N" name rather than inventing a boss - the design spec's own ruling
// (design/specs/2026-10-04-guild-page.md §12.1: "no named encounter published" rather than
// a fabricated one). Every pull is the 40-player raid DifficultyID (186, that same file's
// own note).
package main

import "time"

// raidDifficulty is the 40-player raid DifficultyID (logs/engine/mechanics/encounters.json).
const raidDifficulty = 186

// encOnyxia is Onyxia's Lair's one published encounter
// (logs/engine/mechanics/encounters.json).
const encOnyxia = 1084

const (
	zoneOnyxia      = "Onyxia's Lair"
	zoneBarrowDeeps = "Barrow Deeps"
	zoneHyjal       = "Hyjal Summit"
)

// fightPlan is one pull: a name, the encounter it belongs to (0 when none is published -
// Barrow Deeps and Hyjal Summit, or trash, neither of which ranks or counts toward
// progression), whether it ended in a kill, and how long it ran.
type fightPlan struct {
	Name        string
	EncounterID int64 // 0 for trash or an unpublished encounter
	// Trash marks filler that is never a real attempt (never wipes, never counts
	// toward a night's wipe total) - distinct from EncounterID == 0, which an
	// unpublished-but-real pull (Barrow Deeps/Hyjal Summit) also carries.
	Trash      bool
	Kill       bool
	DurationMS int
}

// reportPlan is one raid night: a title, the zone it was fought in, how many days before
// now it happened, and its fights in pull order.
type reportPlan struct {
	Tag       string // short, stable suffix for this report's id and seed_rows row_key
	Label     string // the zone-ish name the title is built from, e.g. "Onyxia"
	Zone      string
	DaysAgo   int
	StartHour int // the raid's local-feeling start hour, UTC, for a stable time of day
	Fights    []fightPlan
}

// raidPlans is the four raid nights this seed writes: one Barrow Deeps night and one
// Hyjal Summit night, both older than a week (so the "this week" filter also exercises
// the days-outside-the-window case), and two Onyxia's Lair nights within the last seven
// days (so GET /v1/guilds/{id}/home's "this week's reports" list has rows) - both ending
// in a kill, per the real first tier's one published boss. Every report carries one or
// two wipes among six to ten fights, per the seed's own contract.
func raidPlans() []reportPlan {
	trash := func(name string, ms int) fightPlan {
		return fightPlan{Name: name, DurationMS: ms, Kill: true, Trash: true}
	}
	pull := func(name string, kill bool, ms int) fightPlan {
		return fightPlan{Name: name, Kill: kill, DurationMS: ms}
	}
	boss := func(name string, enc int64, kill bool, ms int) fightPlan {
		return fightPlan{Name: name, EncounterID: enc, Kill: kill, DurationMS: ms}
	}
	return []reportPlan{
		{
			Tag: "barrow", Label: zoneBarrowDeeps, Zone: zoneBarrowDeeps, DaysAgo: 13, StartHour: 19,
			Fights: []fightPlan{
				pull("Pull 1", true, 180_000),
				pull("Pull 2", true, 200_000),
				pull("Pull 3", false, 240_000),
				pull("Pull 4", true, 220_000),
				pull("Pull 5", true, 190_000),
				pull("Pull 6", true, 210_000),
				pull("Pull 7", true, 230_000),
			},
		},
		{
			Tag: "hyjal", Label: zoneHyjal, Zone: zoneHyjal, DaysAgo: 11, StartHour: 19,
			Fights: []fightPlan{
				pull("Pull 1", true, 190_000),
				pull("Pull 2", true, 210_000),
				pull("Pull 3", true, 230_000),
				pull("Pull 4", false, 250_000),
				pull("Pull 5", false, 240_000),
				pull("Pull 6", true, 220_000),
				pull("Pull 7", true, 200_000),
				pull("Pull 8", true, 180_000),
			},
		},
		{
			Tag: "ony1", Label: "Onyxia", Zone: zoneOnyxia, DaysAgo: 6, StartHour: 19,
			Fights: []fightPlan{
				trash("Onyxian Whelp", 60_000),
				trash("Onyxian Whelp", 55_000),
				trash("Onyxian Whelp", 58_000),
				boss("Onyxia", encOnyxia, false, 200_000),
				boss("Onyxia", encOnyxia, false, 210_000),
				boss("Onyxia", encOnyxia, true, 260_000),
			},
		},
		{
			Tag: "ony2", Label: "Onyxia", Zone: zoneOnyxia, DaysAgo: 4, StartHour: 19,
			Fights: []fightPlan{
				trash("Onyxian Whelp", 50_000),
				trash("Onyxian Whelp", 52_000),
				trash("Onyxian Whelp", 54_000),
				trash("Onyxian Whelp", 56_000),
				boss("Onyxia", encOnyxia, false, 220_000),
				boss("Onyxia", encOnyxia, true, 250_000),
			},
		},
	}
}

// pullGap is the time between one pull ending and the next starting — invite/buff/
// run-back time, not simulated otherwise.
const pullGap = 4 * time.Minute

// schedule turns a reportPlan into its created_at, completed_at, and each fight's own
// start time, anchored on now.
func (p reportPlan) schedule(now time.Time) (createdAt, completedAt time.Time, starts []time.Time) {
	createdAt = time.Date(now.Year(), now.Month(), now.Day(), p.StartHour, 0, 0, 0, time.UTC).
		AddDate(0, 0, -p.DaysAgo)
	at := createdAt
	starts = make([]time.Time, len(p.Fights))
	for i, f := range p.Fights {
		starts[i] = at
		at = at.Add(time.Duration(f.DurationMS) * time.Millisecond).Add(pullGap)
	}
	completedAt = at
	return createdAt, completedAt, starts
}

// titleFor is this plan's report title, naming the actual weekday createdAt falls on
// rather than a hardcoded day name - this tool computes every date relative to the
// moment it runs (schedule, above), so a fixed day name would lie about the row's own
// created_at the moment the tool runs on a different weekday.
func (p reportPlan) titleFor(createdAt time.Time) string {
	return p.Label + " · " + createdAt.Weekday().String()
}

// kills and wipes count a plan's own fights (trash excluded from both), for the dry-run
// summary and tests.
func (p reportPlan) kills() int {
	n := 0
	for _, f := range p.Fights {
		if !f.Trash && f.Kill {
			n++
		}
	}
	return n
}

func (p reportPlan) wipes() int {
	n := 0
	for _, f := range p.Fights {
		if !f.Trash && !f.Kill {
			n++
		}
	}
	return n
}
