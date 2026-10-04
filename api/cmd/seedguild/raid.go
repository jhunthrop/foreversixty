// api/cmd/seedguild/raid.go
//
// The four raid nights this seed writes: who fought, which boss, how long, and whether it
// died. Encounter ids and zone names come from logs/engine/mechanics/encounters.json
// (read from the Forever beta client itself, not invented) — Molten Core's ten bosses and
// Onyxia, both the 40-player difficulty (DifficultyID 186 per that file's own note).
package main

import "time"

// raidDifficulty is the 40-player raid DifficultyID (logs/engine/mechanics/encounters.json).
const raidDifficulty = 186

// Molten Core and Onyxia's Lair encounter ids, from
// logs/engine/mechanics/encounters.json's own read of DungeonEncounter.db2.
const (
	encLucifron          = 663
	encMagmadar          = 664
	encGehennas          = 665
	encGarr              = 666
	encShazzrah          = 667
	encBaronGeddon       = 668
	encSulfuronHarbinger = 669
	encGolemagg          = 670
	encMajordomoExecutus = 671
	encOnyxia            = 1084
)

const (
	zoneMoltenCore = "Molten Core"
	zoneOnyxia     = "Onyxia's Lair"
)

// fightPlan is one pull: a name, the encounter it belongs to (0 for trash, which never
// ranks and never counts toward progression), whether it ended in a kill, and how long it
// ran.
type fightPlan struct {
	Name        string
	EncounterID int64 // 0 for trash
	Kill        bool
	DurationMS  int
}

// reportPlan is one raid night: a title, the zone it was fought in, how many days before
// now it happened, and its fights in pull order.
type reportPlan struct {
	Tag       string // short, stable suffix for this report's id and seed_rows row_key
	Label     string // the zone-ish name the title is built from, e.g. "Molten Core"
	Zone      string
	DaysAgo   int
	StartHour int // the raid's local-feeling start hour, UTC, for a stable time of day
	Fights    []fightPlan
}

// raidPlans is the four raid nights this seed writes: two Molten Core nights older than a
// week (so the "this week" filter also exercises the days-outside-the-window case), two
// Onyxia nights within the last seven days (so GET /v1/guilds/{id}/home's "this week's
// reports" list has rows). Every report carries one or two wipes among six to ten fights,
// per the seed's own contract.
func raidPlans() []reportPlan {
	trash := func(name string, ms int) fightPlan { return fightPlan{Name: name, DurationMS: ms, Kill: true} }
	boss := func(name string, enc int64, kill bool, ms int) fightPlan {
		return fightPlan{Name: name, EncounterID: enc, Kill: kill, DurationMS: ms}
	}
	return []reportPlan{
		{
			Tag: "mc1", Label: zoneMoltenCore, Zone: zoneMoltenCore, DaysAgo: 13, StartHour: 19,
			Fights: []fightPlan{
				trash("Core Hound Pack", 85_000),
				trash("Flamewaker Guard", 60_000),
				boss("Lucifron", encLucifron, true, 240_000),
				boss("Magmadar", encMagmadar, true, 260_000),
				boss("Gehennas", encGehennas, true, 230_000),
				boss("Garr", encGarr, false, 300_000),
				boss("Garr", encGarr, true, 280_000),
				boss("Shazzrah", encShazzrah, true, 200_000),
			},
		},
		{
			Tag: "mc2", Label: zoneMoltenCore, Zone: zoneMoltenCore, DaysAgo: 11, StartHour: 19,
			Fights: []fightPlan{
				trash("Firelord Pack", 70_000),
				boss("Baron Geddon", encBaronGeddon, true, 250_000),
				boss("Sulfuron Harbinger", encSulfuronHarbinger, true, 260_000),
				trash("Flamewaker Elite", 50_000),
				boss("Golemagg the Incinerator", encGolemagg, false, 320_000),
				boss("Golemagg the Incinerator", encGolemagg, true, 300_000),
				boss("Majordomo Executus", encMajordomoExecutus, true, 280_000),
				trash("Core Rager", 40_000),
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
// moment it runs (schedule, above), so a fixed "Tuesday"/"Thursday" label would lie
// about the row's own created_at the moment the tool runs on a different weekday.
func (p reportPlan) titleFor(createdAt time.Time) string {
	return p.Label + " · " + createdAt.Weekday().String()
}

// kills and wipes count a plan's own fights, for the dry-run summary and tests.
func (p reportPlan) kills() int {
	n := 0
	for _, f := range p.Fights {
		if f.Kill {
			n++
		}
	}
	return n
}

func (p reportPlan) wipes() int {
	n := 0
	for _, f := range p.Fights {
		if f.EncounterID != 0 && !f.Kill {
			n++
		}
	}
	return n
}
