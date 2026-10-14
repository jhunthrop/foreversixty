package rating

// The character page's gear panel (Task L6.7): for each of the 17 slots an FS1 export names,
// in the export's order, what the character wears (its newest addon export, decoded by
// fs1.Decode as GET /v1/me does), the band's pick for the slot (sims.BandPicks: the data
// addon's pick for the spec, band and faction), and the gap: the Gear component's points
// the pick would add over the worn item.
//
// The spec, the level and the band are the Gear component's own (gearScore): the newest
// rated fight's spec, the export's level= section (else 60) and sims.GearBand. The gap uses
// the component's own scale, GearMax * (pick - worn) / full, where pick and worn are each
// item's sims.ItemGearScore (the per-item term GearScore and FullGearScore sum) and full is
// sims.FullGearScore for the band. The gaps do not add up to the component's distance to
// GearMax: the mark is the band's best set across presets and factions, which need not be
// the character's own faction's bare picks, and an item better than its pick offsets a weak
// slot in the component but has a gap of 0 here.
//
// Two slots' gaps are read together rather than one by one (fix round 1):
//
//   - a ring or trinket pair: the worn pair is matched to the picked pair whichever way round
//     gives the smaller total, so a pick worn in the other slot of its pair reads 0;
//   - the hands, when the picked or the worn main hand is two-handed: the weapons are one
//     unit, max(0, (picked main + picked off) - (worn main + worn off)) on the main hand,
//     and the off-hand 0 (null when the band has no off-hand pick).
//
// Which class's tables answer: the worn item's name, item level and two-handedness come from
// the export's own class (the class that wears it); every score, and a pick's
// two-handedness, from the spec's class (the tables FullGearScore scores the mark in). The
// two are the same class for any real character. An item the tables do not have is
// nameless, unscorable (its gap is null) and counts as one-handed.

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

// ScoreAndGear is the character page's Forever Score and gear panel, from one read of the
// character's rated fights and one of its export (characterInputs). An anonymised account has
// neither. A panel that cannot be read is logged and left empty; the error is the score's.
func (s *Service) ScoreAndGear(ctx context.Context, key string) (rankings.ForeverScore, rankings.CharacterGear, error) {
	empty := rankings.CharacterGear{Slots: []rankings.GearSlot{}}
	anon, err := s.Store.anonymized(ctx, key)
	if err != nil {
		return rankings.ForeverScore{}, empty, err
	}
	if anon {
		return rankings.ForeverScore{}, empty, nil
	}
	in, err := s.Store.characterInputs(ctx, key)
	if err != nil {
		return rankings.ForeverScore{}, empty, err
	}
	gear, err := s.Store.characterGear(ctx, key, in)
	if err != nil {
		s.logger().Error("rating", "op", "gear_panel", "character", key, "err", err)
		gear = empty
	}
	sc, err := s.Store.scoreV2(ctx, key, in)
	if errors.Is(err, ErrNoLogs) {
		return rankings.ForeverScore{}, gear, nil
	}
	if err != nil {
		return rankings.ForeverScore{}, gear, err
	}
	return toForeverScoreDTO(sc), gear, nil
}

// characterGear is the gear panel over inputs already read: one row per export slot, and
// rankings.GearGapTotal of them. No export (or one that does not decode) is no rows.
//
//   - worn: the export's item in the slot; null when the slot is empty.
//   - pick: sims.BandPicks for the spec, band and faction (the export's race's, else
//     characters.faction); null when the character has no spec, its level is below the first
//     band, its faction or the spec's bis file is unknown, or the band has no pick for the
//     slot.
//   - gap: 0 when the worn item is the pick or scores at least as much; null when there is no
//     pick, no mark (sims.FullGearScore cannot be read) or either item is not in the data;
//     otherwise GearMax * (pick - worn) / full, one decimal (an upgrade under 0.05 points
//     reads 0). An empty slot scores 0. Pairs and two-handers are read together (file
//     comment).
func (s *Store) characterGear(ctx context.Context, key string, in characterInputs) (rankings.CharacterGear, error) {
	out := rankings.CharacterGear{Slots: []rankings.GearSlot{}}
	if !in.hasExport {
		return out, nil
	}
	decoded, ok := fs1.Decode(in.export)
	if !ok {
		s.logger().Warn("rating", "op", "gear_panel", "character", key, "err", "the stored export is not an FS1 code")
		return out, nil
	}
	spec := in.fights.spec
	faction := sims.RaceFaction(decoded.DataBuild, decoded.RaceSlug)
	if faction == "" {
		faction = in.faction
	}
	picks, full, err := s.bandPicks(ctx, key, decoded, spec, faction)
	if err != nil {
		return out, err
	}
	for _, slot := range sims.GearSlots {
		row := rankings.GearSlot{Slot: slot}
		if id := decoded.Gear[slot]; id > 0 {
			row.Worn = &rankings.GearWorn{ID: id}
			if item, ok := sims.GearItem(decoded.DataBuild, decoded.ClassSlug, id); ok {
				name, ilvl := item.Name, item.ItemLevel
				row.Worn.Name, row.Worn.ItemLevel = &name, &ilvl
			}
		}
		if p, ok := picks[slot]; ok {
			row.Pick = &rankings.GearPick{ID: p.ItemID, Name: p.Name, Source: p.Source, Chance: p.Chance}
		}
		out.Slots = append(out.Slots, row)
	}
	g := gapReader{ctx: ctx, build: decoded.DataBuild, spec: spec, wornClass: decoded.ClassSlug, full: full,
		scores: map[int]itemScore{}}
	if err := g.fill(out.Slots); err != nil {
		return out, err
	}
	out.GapTotal = rankings.GearGapTotal(out.Slots)
	return out, nil
}

// bandPicks is the band's picks for the export's character under spec and faction, and the
// band's mark (sims.FullGearScore; 0 when it cannot be read). Data the loaded build cannot
// answer for is no picks or no mark, never an error (unknownGear).
func (s *Store) bandPicks(ctx context.Context, key string, decoded fs1.Decoded, spec, faction string) (map[string]sims.BandPick, float64, error) {
	if spec == "" {
		return nil, 0, nil
	}
	band, ok, err := sims.GearBand(ctx, decoded.DataBuild, spec, gearLevel(decoded))
	if err != nil || !ok {
		return nil, 0, s.unknownGear(key, spec, err)
	}
	var picks map[string]sims.BandPick
	if faction != "" {
		if picks, err = sims.BandPicks(ctx, decoded.DataBuild, spec, band, faction); err != nil {
			if err := s.unknownGear(key, spec, err); err != nil {
				return nil, 0, err
			}
			picks = nil
		}
	}
	full, err := sims.FullGearScore(ctx, decoded.DataBuild, spec, band)
	if err != nil {
		return picks, 0, s.unknownGear(key, spec, err)
	}
	return picks, full, nil
}

// unknownGear drops err when it says the loaded data cannot answer: a spec with no
// best-in-slot file or band (ErrNoBestInSlot: every tank and healer spec without one, so
// logged at debug level only) or data that cannot be scored (ErrGearUnscorable, a warning).
// Any other error is returned.
func (s *Store) unknownGear(key, spec string, err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, sims.ErrNoBestInSlot):
		s.logger().Debug("rating", "op", "gear_panel", "character", key, "spec", spec, "err", err)
		return nil
	case errors.Is(err, sims.ErrGearUnscorable):
		s.logger().Warn("rating", "op", "gear_panel", "character", key, "spec", spec, "err", err)
		return nil
	}
	return fmt.Errorf("rating: gear panel of %s: %w", key, err)
}

// itemScore is one item's sims.ItemGearScore; known is false when the tables lack it.
type itemScore struct {
	score float64
	known bool
}

// gapReader fills a panel's gaps (characterGear's rules).
type gapReader struct {
	ctx                    context.Context
	build, spec, wornClass string
	full                   float64
	scores                 map[int]itemScore
}

// score is id's gear score for the spec, read once per panel.
func (g *gapReader) score(id int) (float64, bool, error) {
	if v, ok := g.scores[id]; ok {
		return v.score, v.known, nil
	}
	v, known, err := sims.ItemGearScore(g.ctx, g.build, g.spec, id)
	if err != nil && !errors.Is(err, sims.ErrGearUnscorable) {
		return 0, false, err
	}
	known = known && err == nil
	g.scores[id] = itemScore{v, known}
	return v, known, nil
}

// sum is the gear score of ids taken together (0 for none); ok is false when one is unknown.
func (g *gapReader) sum(ids ...int) (float64, bool, error) {
	total := 0.0
	for _, id := range ids {
		v, ok, err := g.score(id)
		if err != nil || !ok {
			return 0, false, err
		}
		total += v
	}
	return total, true, nil
}

// gap is one worn item (nil: nothing worn) against one pick.
func (g *gapReader) gap(worn *rankings.GearWorn, pick *rankings.GearPick) (*float64, error) {
	if pick == nil {
		return nil, nil
	}
	if worn != nil && worn.ID == pick.ID {
		return ptrFloat(0), nil
	}
	return g.points([]int{pick.ID}, wornIDs(worn))
}

// points is the gap of the picked items over the worn ones taken together; null when there is
// no mark or an item is unknown.
func (g *gapReader) points(picked, worn []int) (*float64, error) {
	if g.full <= 0 {
		return nil, nil
	}
	pickScore, ok, err := g.sum(picked...)
	if err != nil || !ok {
		return nil, err
	}
	wornScore, ok, err := g.sum(worn...)
	if err != nil || !ok {
		return nil, err
	}
	return ptrFloat(gearGapPoints(pickScore, wornScore, g.full)), nil
}

// fill sets every row's gap: each slot on its own, then the ring and trinket pairs and the
// hands read together.
func (g *gapReader) fill(rows []rankings.GearSlot) error {
	at := map[string]*rankings.GearSlot{}
	for i := range rows {
		at[rows[i].Slot] = &rows[i]
		var err error
		if rows[i].Gap, err = g.gap(rows[i].Worn, rows[i].Pick); err != nil {
			return err
		}
	}
	for _, pair := range [][2]string{{"finger1", "finger2"}, {"trinket1", "trinket2"}} {
		a, b := at[pair[0]], at[pair[1]]
		crossA, err := g.gap(b.Worn, a.Pick)
		if err != nil {
			return err
		}
		crossB, err := g.gap(a.Worn, b.Pick)
		if err != nil {
			return err
		}
		if crossedIsSmaller(a.Gap, b.Gap, crossA, crossB) {
			a.Gap, b.Gap = crossA, crossB
		}
	}
	return g.hands(at["main_hand"], at["off_hand"])
}

// crossedIsSmaller reports whether the crossed matching of a pair (crossA, crossB) beats the
// straight one (a, b): fewer gaps that cannot be scored, else a smaller sum. A tie keeps the
// straight matching.
func crossedIsSmaller(a, b, crossA, crossB *float64) bool {
	nulls := func(x, y *float64) int {
		n := 0
		if x == nil {
			n++
		}
		if y == nil {
			n++
		}
		return n
	}
	sum := func(x, y *float64) float64 {
		t := 0.0
		if x != nil {
			t += *x
		}
		if y != nil {
			t += *y
		}
		return t
	}
	straightNulls, crossedNulls := nulls(a, b), nulls(crossA, crossB)
	if crossedNulls != straightNulls {
		return crossedNulls < straightNulls
	}
	return sum(crossA, crossB) < sum(a, b)
}

// hands reads the two hands as one unit when the picked or the worn main hand is two-handed.
// With no main-hand pick the slots keep their own gaps.
func (g *gapReader) hands(mh, oh *rankings.GearSlot) error {
	if mh.Pick == nil {
		return nil
	}
	pickTwoHand := g.twoHanded(g.specClass(), mh.Pick.ID)
	wornTwoHand := mh.Worn != nil && g.twoHanded(g.wornClass, mh.Worn.ID)
	if !pickTwoHand && !wornTwoHand {
		return nil
	}
	picked := []int{mh.Pick.ID}
	if oh.Pick != nil {
		picked = append(picked, oh.Pick.ID)
	}
	worn := append(wornIDs(mh.Worn), wornIDs(oh.Worn)...)
	var unit *float64
	if sameIDs(picked, worn) {
		unit = ptrFloat(0)
	} else {
		var err error
		if unit, err = g.points(picked, worn); err != nil {
			return err
		}
	}
	mh.Gap, oh.Gap = unit, nil
	if unit != nil && oh.Pick != nil {
		oh.Gap = ptrFloat(0)
	}
	return nil
}

// specClass is the class slug of the spec ("warrior" of "warrior-arms").
func (g *gapReader) specClass() string {
	class, _, _ := strings.Cut(g.spec, "-")
	return class
}

// twoHanded reports whether class's tables mark id two-handed; an item they lack is not.
func (g *gapReader) twoHanded(class string, id int) bool {
	it, ok := sims.GearItem(g.build, class, id)
	return ok && it.TwoHand
}

// gearGapPoints is the Gear component's scale on one slot: the points pick would add over
// worn against the mark full, one decimal; 0 when worn scores at least as much as pick.
func gearGapPoints(pick, worn, full float64) float64 {
	if worn >= pick {
		return 0
	}
	return math.Round(GearMax*(pick-worn)/full*10) / 10
}

func wornIDs(w *rankings.GearWorn) []int {
	if w == nil {
		return nil
	}
	return []int{w.ID}
}

// sameIDs reports whether a and b hold the same ids, in any order.
func sameIDs(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	seen := map[int]int{}
	for _, id := range a {
		seen[id]++
	}
	for _, id := range b {
		if seen[id] == 0 {
			return false
		}
		seen[id]--
	}
	return true
}

func ptrFloat(v float64) *float64 { return &v }
