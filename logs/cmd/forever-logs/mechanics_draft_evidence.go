// logs/cmd/forever-logs/mechanics_draft_evidence.go
// The three evidence passes mechanics-draft runs over the same drafted
// fights, alongside the avoidable/unavoidable classification
// mechanics_draft.go itself performs: phase candidates (an enemy cast or
// aura that fires once, consistently, on every pull), downtime candidates
// (the same evidence, read for a spell that instead recurs a consistent
// number of times per pull), and per-spec utility candidates (an aura a
// player was the source of). Split into its own file per spec §3.5,
// which extends the drafting tool with exactly these two extra passes.
package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// phaseCandidate is one spell that could open a phase: how it was seen, when, and on
// how many of the encounter's pulls.
type phaseCandidate struct {
	spellID   int64
	name      string
	on        string
	pulls     int
	firstMS   int64
	healthPct []float64
}

// count records a spell as a phase candidate for this pull only when it
// happened exactly once. A second row for the same spell this pull -- a
// debuff that landed on two different players, say, each its own row with
// its own count of one -- disqualifies it instead of silently keeping
// whichever row was seen first; so does a single row whose own count was
// never one to begin with.
func count(once map[int64]*phaseCandidate, spellID int64, name, on string, firstMS, cnt int64, healthPct float64) {
	existing, ok := once[spellID]
	switch {
	case !ok && cnt == 1:
		once[spellID] = &phaseCandidate{
			spellID: spellID, name: name, on: on, pulls: 1,
			firstMS: firstMS, healthPct: []float64{healthPct},
		}
	case !ok:
		once[spellID] = &phaseCandidate{spellID: spellID, name: name, on: on}
	default:
		existing.pulls = 0
	}
}

// spread is the distance between a float slice's lowest and highest value,
// zero for fewer than two entries -- one pull says nothing about
// consistency yet.
func spread(vals []float64) float64 {
	if len(vals) < 2 {
		return 0
	}
	min, max := vals[0], vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func minOf(vals []float64) float64 {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func maxOf(vals []float64) float64 {
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

// formatMS prints a millisecond offset as m:ss.
func formatMS(ms int64) string {
	total := ms / 1000
	if total < 0 {
		total = 0
	}
	return fmt.Sprintf("%d:%02d", total/60, total%60)
}

// draftPhases proposes phase triggers for one encounter. The rule is the spec's: an
// enemy cast or aura application that happened exactly once in every pull, at a boss
// health that did not move much between them. Anything that happened twice in a pull is
// a rotation, not a phase; anything that happened in one pull of five is a fluke.
func draftPhases(pulls []draftPull) ([]mechanics.Phase, []string) {
	seen := map[int64]*phaseCandidate{}
	for _, pull := range pulls {
		once := map[int64]*phaseCandidate{}
		for _, row := range pull.enemyCasts {
			count(once, row.SpellID, row.SpellName, mechanics.OnCastStart, row.firstMS, row.count, pull.bossPctAt(row.firstMS))
		}
		for _, row := range pull.enemyAuras {
			count(once, row.SpellID, row.SpellName, mechanics.OnAuraApplied, row.firstMS, row.count, pull.bossPctAt(row.firstMS))
		}
		for id, c := range once {
			if c.pulls == 0 {
				continue
			}
			found := seen[id]
			if found == nil {
				seen[id] = c
				continue
			}
			found.pulls++
			found.healthPct = append(found.healthPct, c.healthPct...)
			if c.firstMS < found.firstMS {
				found.firstMS = c.firstMS
			}
		}
	}
	kept := []*phaseCandidate{}
	for _, c := range seen {
		if c.pulls == len(pulls) && spread(c.healthPct) <= 10 {
			kept = append(kept, c)
		}
	}
	// Candidates come out of a map, so two at the same instant would otherwise
	// swap names between runs; the spell id breaks the tie the same way every time.
	sort.SliceStable(kept, func(i, j int) bool {
		if kept[i].firstMS != kept[j].firstMS {
			return kept[i].firstMS < kept[j].firstMS
		}
		return kept[i].spellID < kept[j].spellID
	})
	phases := make([]mechanics.Phase, 0, len(kept))
	evidence := make([]string, 0, len(kept))
	for i, c := range kept {
		name := fmt.Sprintf("Phase %d", i+2)
		phases = append(phases, mechanics.Phase{
			Name:   name,
			Starts: mechanics.PhaseStart{SpellID: c.spellID, On: c.on},
		})
		evidence = append(evidence, fmt.Sprintf(
			"%s: %s (%d) %s once on each of %d pulls, first at %s, boss at %.0f%%-%.0f%%",
			name, c.name, c.spellID, c.on, c.pulls, formatMS(c.firstMS), minOf(c.healthPct), maxOf(c.healthPct)))
	}
	return phases, evidence
}

// downtimePlaceholderMS is the duration a drafted downtime candidate
// carries until a curator corrects it: this tool sees that a trigger
// recurs on a steady cadence, never how long raid activity actually paused
// for that stretch in the log, so it proposes the same length spec §3.4's
// own worked example uses (a Firesworn detonation window) rather than
// guessing a different number with no more basis than that one.
const downtimePlaceholderMS = 3000

// downtimeCandidate is one enemy cast or aura that repeats a consistent
// number of times within every pull of an encounter -- the signature of a
// recurring forced-downtime mechanic (an add's death detonation, a
// transition that repeats every so often) rather than a one-off phase
// trigger. repeat is -1 once a spell's per-pull repeat count disagrees
// between pulls, which disqualifies it: an inconsistent cadence is not
// something a curator can hang one duration_ms number on.
type downtimeCandidate struct {
	name   string
	on     string
	repeat int64
	pulls  int
}

// draftDowntime proposes downtime-window candidates (spec §3.5): the same
// underlying evidence draftPhases reads (enemy casts and aura applications
// across every pull), read through the second lens the spec names --
// "repeats within a single phase rather than opening a new one" -- rather
// than requiring exactly one occurrence per pull the way a phase trigger
// does.
func draftDowntime(pulls []draftPull) ([]mechanics.Downtime, []string) {
	seen := map[int64]*downtimeCandidate{}
	for _, pull := range pulls {
		counts := map[int64]*downtimeCandidate{}
		for _, row := range pull.enemyCasts {
			tallyDowntimeRepeat(counts, row, mechanics.OnCastStart)
		}
		for _, row := range pull.enemyAuras {
			tallyDowntimeRepeat(counts, row, mechanics.OnAuraApplied)
		}
		for id, c := range counts {
			found := seen[id]
			switch {
			case found == nil:
				seen[id] = &downtimeCandidate{name: c.name, on: c.on, repeat: c.repeat, pulls: 1}
			case found.repeat != c.repeat:
				found.repeat = -1
			default:
				found.pulls++
			}
		}
	}
	var ids []int64
	for id, c := range seen {
		if c.repeat >= 2 && c.pulls == len(pulls) {
			ids = append(ids, id)
		}
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	downtime := make([]mechanics.Downtime, 0, len(ids))
	evidence := make([]string, 0, len(ids))
	for _, id := range ids {
		c := seen[id]
		downtime = append(downtime, mechanics.Downtime{
			Trigger:    mechanics.PhaseStart{SpellID: id, On: c.on},
			DurationMS: downtimePlaceholderMS,
			Note: fmt.Sprintf(
				"draft: %s (%d) recurred %dx per pull across all %d pulls -- verify the actual downtime length from the log before committing",
				c.name, id, c.repeat, c.pulls),
		})
		evidence = append(evidence, fmt.Sprintf(
			"downtime candidate: %s (%d) %s, %dx per pull on all %d pulls, placeholder duration %dms -- curator must verify",
			c.name, id, c.on, c.repeat, c.pulls, downtimePlaceholderMS))
	}
	return downtime, evidence
}

// tallyDowntimeRepeat records one pull's repeat count for a spell that
// happened more than once -- draftPhases' count() disqualifies exactly
// this case as "not a phase, a rotation"; draftDowntime is that
// disqualified case read as its own signal instead of discarded.
func tallyDowntimeRepeat(counts map[int64]*downtimeCandidate, row draftRow, on string) {
	if row.count < 2 {
		return
	}
	if existing, ok := counts[row.SpellID]; ok {
		if existing.repeat != row.count {
			existing.repeat = -1
		}
		return
	}
	counts[row.SpellID] = &downtimeCandidate{name: row.SpellName, on: on, repeat: row.count}
}

// utilityDrafter accumulates, per spec, evidence for the utility table's
// owned list (spec §3.2): every aura a player of that spec applied to
// something else -- a buff to an ally, a debuff to an enemy -- across
// every fight the log contains. Spec-keyed rather than encounter-keyed,
// since a spec's utility does not depend on which boss it was used
// against.
type utilityDrafter struct {
	bySpec map[string]map[int64]*utilityCandidate
}

// utilityCandidate is one spec's history with one spell it applied:
// whether it looks like a buff or a debuff, how many fights it was seen
// in, and its uptime share averaged across those fights.
type utilityCandidate struct {
	name          string
	target        string
	fights        int
	meanUptimePct float64
}

func newUtilityDrafter() *utilityDrafter {
	return &utilityDrafter{bySpec: map[string]map[int64]*utilityCandidate{}}
}

// addFight folds one fight's player-sourced auras into their casters'
// spec evidence.
func (d *utilityDrafter) addFight(sum summary.Summary, reg *units.Registry) {
	if sum.DurationMS <= 0 {
		return
	}
	specBySource := map[string]string{}
	for _, c := range sum.Combatants {
		if c.Spec != "" {
			specBySource[c.GUID] = c.Spec
		}
	}
	for _, tr := range sum.Auras {
		if len(tr.Appliers) == 0 {
			continue
		}
		pct := float64(tr.UptimeMS) / float64(sum.DurationMS) * 100
		for _, applier := range tr.Appliers {
			u, ok := reg.Get(applier)
			if !ok || !u.IsPlayer() {
				continue
			}
			spec := specBySource[applier]
			if spec == "" {
				continue
			}
			d.credit(spec, tr, pct)
		}
	}
}

func (d *utilityDrafter) credit(spec string, tr summary.AuraTrack, pct float64) {
	bySpell := d.bySpec[spec]
	if bySpell == nil {
		bySpell = map[int64]*utilityCandidate{}
		d.bySpec[spec] = bySpell
	}
	c := bySpell[tr.SpellID]
	if c == nil {
		target := "ally"
		if tr.Type == "DEBUFF" {
			target = "enemy"
		}
		c = &utilityCandidate{name: tr.Name, target: target}
		bySpell[tr.SpellID] = c
	}
	c.fights++
	// Running mean: c.meanUptimePct after n fights is the mean of the first
	// n readings, updated one fight at a time rather than summed and
	// divided at the end, so a fight processed out of memory order still
	// produces the same figure.
	c.meanUptimePct += (pct - c.meanUptimePct) / float64(c.fights)
}

// write prints one evidence line per spec/spell candidate, in the same
// "draft:"-prefixed convention this file's mechanic evidence already uses,
// for a curator to check the spell id against
// data/builds/1.60.1.69893/spells.json before adding a verified entry to
// logs/engine/mechanics/utility/tables.
func (d *utilityDrafter) write(errOut io.Writer) error {
	specs := make([]string, 0, len(d.bySpec))
	for spec := range d.bySpec {
		specs = append(specs, spec)
	}
	sort.Strings(specs)
	for _, spec := range specs {
		bySpell := d.bySpec[spec]
		ids := make([]int64, 0, len(bySpell))
		for id := range bySpell {
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
		for _, id := range ids {
			c := bySpell[id]
			if _, err := fmt.Fprintf(errOut,
				"utility draft: %s: %s (%d), target %s, seen on %d fights, mean uptime %.0f%% -- verify against spells.json before adding to utility/tables/%s.json\n",
				spec, c.name, id, c.target, c.fights, c.meanUptimePct, spec); err != nil {
				return err
			}
		}
	}
	return nil
}
