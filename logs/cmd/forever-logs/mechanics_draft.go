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
	o, err := sessionOptions("local", baseTime(path), "", false)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	d := newDrafter(*encounterID)
	s := session.New(o)
	if err := stream(s, file, func(c session.Closed) error {
		d.addFight(c.Fight, c.Summary, s.Units())
		return nil
	}); err != nil {
		return err
	}
	return d.write(out)
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

// encounterDraft accumulates one encounter's spell evidence.
type encounterDraft struct {
	name             string
	totalDamageTaken int64
	spells           map[int64]*spellDraft
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
func (d *drafter) addFight(f fight.Fight, sum summary.Summary, reg *units.Registry) {
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
			ed.spell(ab.SpellID, ab.Name).hit(actor.GUID, ab.Hits, ab.Effective, tank[actor.GUID])
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
		kind, note = mechanics.Avoidable, "unclassified: check"
	}
	return mechanics.Mechanic{SpellID: spellID, Name: sd.name, Kind: kind, Note: note}
}

func percent(part, whole int64) int64 {
	if whole <= 0 {
		return 0
	}
	return int64(math.Round(float64(part) / float64(whole) * 100))
}

func (d *drafter) write(out io.Writer) error {
	for _, id := range d.order {
		data, err := json.MarshalIndent(d.tabs[id].table(id), "", "  ")
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
