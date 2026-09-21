// logs/cmd/forever-logs/mechanics_draft.go
// mechanics-draft reads a log and proposes a mechanics.Table per encounter:
// a starting point for the curator who edits logs/engine/mechanics/tables,
// not a replacement for their judgment. Every drafted mechanic carries a
// "draft:" note naming the evidence it was drafted from.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

func runMechanicsDraft(args []string, out, errOut io.Writer) error {
	fs := flag.NewFlagSet("mechanics-draft", flag.ContinueOnError)
	fs.SetOutput(errOut)
	encounterID := fs.Int64("encounter", 0, "only draft this encounter id")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return errors.New("usage: forever-logs mechanics-draft [-encounter id] <file>")
	}
	path := fs.Arg(0)
	// KeepEvents is on here, unlike the other read paths in this file: the
	// phase candidates need the boss's own health readings across the pull,
	// and those live only in the raw events, not in the summary.
	o, err := sessionOptions("local", baseTime(path), "", true)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	d := newDrafter(*encounterID)
	u := newUtilityDrafter()
	s := session.New(o)
	if err := stream(s, file, func(c session.Closed) error {
		d.addFight(c.Fight, c.Summary, s.Units(), c.Events)
		u.addFight(c.Summary, s.Units())
		return nil
	}); err != nil {
		return err
	}
	if err := d.write(out, errOut); err != nil {
		return err
	}
	// Utility evidence (spec §3.5) is spec-keyed, not encounter-keyed, so it
	// is reported once at the end rather than per encounter table.
	return u.write(errOut)
}

// drafter collects evidence for one or more encounters across every fight
// (pull) of that encounter the log contains, so a boss killed on the third
// try still gets credit for what it did on the first two.
type drafter struct {
	only  int64
	order []int64
	tabs  map[int64]*encounterDraft
}

func newDrafter(only int64) *drafter {
	return &drafter{only: only, tabs: map[int64]*encounterDraft{}}
}

// encounterDraft accumulates one encounter's spell evidence, and one
// draftPull per pull for the phase candidates, which need each pull's
// evidence kept apart rather than folded into one running total.
type encounterDraft struct {
	name             string
	totalDamageTaken int64
	spells           map[int64]*spellDraft
	pulls            []draftPull
}

// spellDraft accumulates one enemy spell's evidence: who it hit, how hard,
// and whether it ended anyone.
type spellDraft struct {
	name    string
	total   int64
	killed  int64
	players map[string]*playerHits
}

// playerHits is one player's history with a spell, split by whether the
// roster had them tanking at the moment of each hit. A player can be the
// tank on one pull and not another, so tankness is judged per hit, not
// once for the player across the whole encounter.
type playerHits struct {
	tankHits    int64
	nonTankHits int64
}

// addFight folds one closed fight's evidence into its encounter's draft.
// Trash fights and pulls that never reached an encounter carry no
// EncounterID and are skipped; only players (per the registry) count as
// evidence, since a mechanic that only ever hits pets or other enemies
// says nothing about what a raid should have avoided.
func (d *drafter) addFight(f fight.Fight, sum summary.Summary, reg *units.Registry, events []event.Event) {
	if f.Kind != fight.Encounter || f.EncounterID <= 0 {
		return
	}
	if d.only != 0 && f.EncounterID != d.only {
		return
	}
	ed := d.tabs[f.EncounterID]
	if ed == nil {
		ed = &encounterDraft{name: f.Name, spells: map[int64]*spellDraft{}}
		d.tabs[f.EncounterID] = ed
		d.order = append(d.order, f.EncounterID)
	}
	ed.addFight(sum, reg)
	ed.pulls = append(ed.pulls, buildDraftPull(f, sum, reg, events))
}

func (ed *encounterDraft) addFight(sum summary.Summary, reg *units.Registry) {
	tank := map[string]bool{}
	for _, r := range sum.Roster {
		tank[r.GUID] = r.Role == "tank"
	}
	for _, actor := range sum.DamageTaken {
		if u, ok := reg.Get(actor.GUID); !ok || !u.IsPlayer() {
			continue
		}
		for _, ab := range actor.Abilities {
			ed.totalDamageTaken += ab.Effective
			if ab.SpellID == 0 {
				continue // melee carries no mechanic to draft
			}
			// Hits counts direct landings and Ticks the periodic ones: a DoT or a
			// ground effect lands entirely in Ticks, and counting only Hits would
			// draft it as unavoidable with "0 non-tanks".
			ed.spell(ab.SpellID, ab.Name).hit(actor.GUID, ab.Hits+ab.Ticks, ab.Effective, tank[actor.GUID])
		}
	}
	for _, death := range sum.Deaths {
		if death.KillingBlow == nil {
			continue
		}
		if u, ok := reg.Get(death.GUID); !ok || !u.IsPlayer() {
			continue
		}
		if sd, ok := ed.spells[death.KillingBlow.SpellID]; ok {
			sd.killed++
		}
	}
}

func (ed *encounterDraft) spell(spellID int64, name string) *spellDraft {
	sd := ed.spells[spellID]
	if sd == nil {
		sd = &spellDraft{name: name, players: map[string]*playerHits{}}
		ed.spells[spellID] = sd
	}
	if name != "" {
		sd.name = name
	}
	return sd
}

func (sd *spellDraft) hit(guid string, hits, effective int64, isTank bool) {
	p := sd.players[guid]
	if p == nil {
		p = &playerHits{}
		sd.players[guid] = p
	}
	if isTank {
		p.tankHits += hits
	} else {
		p.nonTankHits += hits
	}
	sd.total += effective
}

// table renders the encounter's accumulated evidence as a mechanics.Table.
func (ed *encounterDraft) table(encounterID int64) mechanics.Table {
	ids := make([]int64, 0, len(ed.spells))
	for id := range ed.spells {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	ms := make([]mechanics.Mechanic, 0, len(ids))
	for _, id := range ids {
		ms = append(ms, ed.spells[id].draft(id, ed.totalDamageTaken))
	}
	return mechanics.Table{EncounterID: encounterID, Name: ed.name, Mechanics: ms}
}

// draft classifies one spell from its evidence:
//   - avoidable: it hit two or more non-tank players, or hit one non-tank
//     player more than once;
//   - unavoidable: every hit landed on a tank;
//   - otherwise: the evidence is too thin to call, so it drafts avoidable
//     with a note asking the curator to check it by hand.
func (sd *spellDraft) draft(spellID int64, totalDamageTaken int64) mechanics.Mechanic {
	players, nonTanks, repeatNonTank := 0, 0, false
	for _, p := range sd.players {
		players++
		if p.nonTankHits == 0 {
			continue // every hit on this player landed while they tanked
		}
		nonTanks++
		if p.nonTankHits >= 2 {
			repeatNonTank = true
		}
	}

	evidence := fmt.Sprintf("draft: hit %d players (%d non-tanks), %d%% of damage taken, killed %d",
		players, nonTanks, percent(sd.total, totalDamageTaken), sd.killed)

	var kind mechanics.Kind
	note := evidence
	switch {
	case nonTanks >= 2 || repeatNonTank:
		kind = mechanics.Avoidable
	case nonTanks == 0:
		kind = mechanics.Unavoidable
	default:
		// One non-tank, hit once: not enough evidence either way.
		kind, note = mechanics.Avoidable, "unclassified: check — "+evidence
	}
	return mechanics.Mechanic{SpellID: spellID, Name: sd.name, Kind: kind, Note: note}
}

func percent(part, whole int64) int64 {
	if whole <= 0 {
		return 0
	}
	return int64(math.Round(float64(part) / float64(whole) * 100))
}

func (d *drafter) write(out, errOut io.Writer) error {
	for _, id := range d.order {
		ed := d.tabs[id]
		t := ed.table(id)
		phases, evidence := draftPhases(ed.pulls)
		t.Phases = phases
		for _, line := range evidence {
			if _, err := fmt.Fprintln(errOut, line); err != nil {
				return err
			}
		}
		downtime, downtimeEvidence := draftDowntime(ed.pulls)
		t.Downtime = downtime
		for _, line := range downtimeEvidence {
			if _, err := fmt.Fprintln(errOut, line); err != nil {
				return err
			}
		}
		data, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return err
		}
		if _, err := out.Write(data); err != nil {
			return err
		}
		if _, err := fmt.Fprintln(out); err != nil {
			return err
		}
	}
	return nil
}

// draftRow is one enemy spell as seen once in a pull: its id, name, first
// offset into the pull, and how many times it happened. A cast row's count
// is how many times it started (mechanics.OnCastStart is what a channel
// logs, per the spec, so success is not what is counted here); an aura
// row's count is its track's own Applications.
type draftRow struct {
	SpellID   int64
	SpellName string
	firstMS   int64
	count     int64
}

// healthReading is the boss's health at one instant of a pull, read off an
// advanced block that named the encounter's own boss.
type healthReading struct {
	ms  int64
	pct float64
}

// draftPull holds one pull's evidence for phase candidates: the enemy cast
// rows and enemy aura tracks (read off the pull's summary.Casts and
// summary.Auras, keeping only rows whose GUID -- the caster for a cast, the
// first applier for an aura -- the registry does not know as a player) and
// the boss's health readings across the pull, so a candidate can be scored
// against the boss health it happened at.
type draftPull struct {
	enemyCasts []draftRow
	enemyAuras []draftRow
	health     []healthReading
}

// bossPctAt is the boss's health at or just before ms: the last reading at
// or before it, or the pull's first reading when nothing came before. Zero
// when the pull carries no boss health readings at all -- a log with no
// advanced logging enabled, say.
func (p draftPull) bossPctAt(ms int64) float64 {
	if len(p.health) == 0 {
		return 0
	}
	best := p.health[0]
	for _, r := range p.health {
		if r.ms > ms {
			break
		}
		best = r
	}
	return best.pct
}

// buildDraftPull turns one closed pull into the evidence draftPhases scores
// candidates against.
func buildDraftPull(f fight.Fight, sum summary.Summary, reg *units.Registry, events []event.Event) draftPull {
	return draftPull{
		enemyCasts: enemyCastRows(sum.Casts, reg),
		enemyAuras: enemyAuraRows(sum.Auras, reg),
		health:     bossHealthReadings(events, f.Name, reg, f.Start),
	}
}

// enemyCastRows is the pull's cast_start candidates. A cast that never
// succeeded (interrupted, or still going when the pull ended) carries no
// timestamp of its own in the summary -- CastRow.Sequence records only
// successes -- so it is timed at the pull's own start, which is the most a
// draft can say about when it happened.
func enemyCastRows(casts []summary.CastRow, reg *units.Registry) []draftRow {
	rows := make([]draftRow, 0, len(casts))
	for _, c := range casts {
		if u, ok := reg.Get(c.GUID); ok && u.IsPlayer() {
			continue
		}
		var firstMS int64
		if len(c.Sequence) > 0 {
			firstMS = c.Sequence[0]
		}
		rows = append(rows, draftRow{SpellID: c.SpellID, SpellName: c.SpellName, firstMS: firstMS, count: c.Started})
	}
	return rows
}

// enemyAuraRows is the pull's aura_applied candidates: one row per
// (target, spell) track whose first applier the registry does not know as
// a player.
func enemyAuraRows(auras []summary.AuraTrack, reg *units.Registry) []draftRow {
	rows := make([]draftRow, 0, len(auras))
	for _, tr := range auras {
		if len(tr.Appliers) == 0 {
			continue
		}
		if u, ok := reg.Get(tr.Appliers[0]); ok && u.IsPlayer() {
			continue
		}
		var firstMS int64
		if len(tr.Segments) > 0 {
			firstMS = tr.Segments[0].StartMS
		}
		rows = append(rows, draftRow{SpellID: tr.SpellID, SpellName: tr.Name, firstMS: firstMS, count: tr.Applications})
	}
	return rows
}

// bossHealthReadings collects the boss's health percentage across the pull,
// in log order: every advanced block that named a unit matching the
// encounter's own name, the same identity the fight list reads a wipe
// percentage from.
func bossHealthReadings(events []event.Event, bossName string, reg *units.Registry, start time.Time) []healthReading {
	var out []healthReading
	for _, e := range events {
		if !e.Adv.OK || e.Adv.MaxHP <= 0 || e.Adv.InfoGUID == "" {
			continue
		}
		if reg.Name(e.Adv.InfoGUID) != bossName {
			continue
		}
		out = append(out, healthReading{
			ms:  e.Time.Sub(start).Milliseconds(),
			pct: float64(e.Adv.CurrentHP) / float64(e.Adv.MaxHP) * 100,
		})
	}
	return out
}

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
