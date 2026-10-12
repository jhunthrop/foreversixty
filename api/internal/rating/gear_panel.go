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

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
)

// CharacterGear is the gear panel of the character keyed key (Store.CharacterGear).
func (s *Service) CharacterGear(ctx context.Context, key string) (rankings.CharacterGear, error) {
	return s.Store.CharacterGear(ctx, key)
}

// CharacterGear is the gear panel of the character keyed key: one row per export slot, and
// rankings.GearGapTotal of them. No export (or one that does not decode) is no rows.
//
//   - worn: the export's item in the slot, named from the export's data build; null when the
//     slot is empty.
//   - pick: sims.BandPicks for the spec, band and the faction of the export's race; null when
//     the character has no spec, its level is below the first band, the faction or the
//     spec's bis file is unknown, or the band has no pick for the slot.
//   - gap: 0 when the worn item is the pick or scores at least as much; null when there is no
//     pick, no mark (sims.FullGearScore cannot be read) or either item is not in the data;
//     otherwise GearMax * (pick - worn) / full, one decimal. An empty slot scores 0.
func (s *Store) CharacterGear(ctx context.Context, key string) (rankings.CharacterGear, error) {
	out := rankings.CharacterGear{Slots: []rankings.GearSlot{}}
	var export string
	err := s.Pool.QueryRow(ctx, `select export from addon_exports where character_key = $1`, key).Scan(&export)
	if errors.Is(err, pgx.ErrNoRows) {
		return out, nil
	}
	if err != nil {
		return out, fmt.Errorf("rating: export of %s: %w", key, err)
	}
	decoded, ok := fs1.Decode(export)
	if !ok {
		s.logger().Warn("rating", "op", "gear_panel", "character", key, "err", "the stored export is not an FS1 code")
		return out, nil
	}
	fights, err := s.ratedFights(ctx, key)
	if err != nil {
		return out, err
	}
	spec := fights.spec
	picks, full, err := s.bandPicks(ctx, key, decoded, spec)
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
		if row.Gap, err = slotGap(ctx, decoded.DataBuild, spec, full, row.Worn, row.Pick); err != nil {
			return out, err
		}
		out.Slots = append(out.Slots, row)
	}
	out.GapTotal = rankings.GearGapTotal(out.Slots)
	return out, nil
}

// bandPicks is the band's picks for the export's character under spec, and the band's mark
// (sims.FullGearScore; 0 when it cannot be read). Data the loaded build cannot answer for is
// no picks or no mark, with one log line, never an error.
func (s *Store) bandPicks(ctx context.Context, key string, decoded fs1.Decoded, spec string) (map[string]sims.BandPick, float64, error) {
	if spec == "" {
		return nil, 0, nil
	}
	band, ok, err := sims.GearBand(ctx, decoded.DataBuild, spec, gearLevel(decoded))
	if err != nil || !ok {
		return nil, 0, s.unknownGear(key, spec, err)
	}
	var picks map[string]sims.BandPick
	if faction := sims.RaceFaction(decoded.DataBuild, decoded.RaceSlug); faction != "" {
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

// unknownGear logs err and drops it when it says the loaded data cannot answer
// (ErrGearUnscorable, ErrNoBestInSlot); any other error is returned.
func (s *Store) unknownGear(key, spec string, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, sims.ErrGearUnscorable) || errors.Is(err, sims.ErrNoBestInSlot) {
		s.logger().Warn("rating", "op", "gear_panel", "character", key, "spec", spec, "err", err)
		return nil
	}
	return fmt.Errorf("rating: gear panel of %s: %w", key, err)
}

// slotGap is one row's gap (see CharacterGear).
func slotGap(ctx context.Context, build, spec string, full float64, worn *rankings.GearWorn, pick *rankings.GearPick) (*float64, error) {
	if pick == nil {
		return nil, nil
	}
	if worn != nil && worn.ID == pick.ID {
		zero := 0.0
		return &zero, nil
	}
	if full <= 0 {
		return nil, nil
	}
	pickScore, ok, err := sims.ItemGearScore(ctx, build, spec, pick.ID)
	if err != nil || !ok {
		return nil, ignoreUnscorable(err)
	}
	wornScore := 0.0
	if worn != nil {
		if wornScore, ok, err = sims.ItemGearScore(ctx, build, spec, worn.ID); err != nil || !ok {
			return nil, ignoreUnscorable(err)
		}
	}
	gap := gearGapPoints(pickScore, wornScore, full)
	return &gap, nil
}

// gearGapPoints is the Gear component's scale on one slot: the points pick would add over
// worn against the mark full, one decimal; 0 when worn scores at least as much as pick.
func gearGapPoints(pick, worn, full float64) float64 {
	if worn >= pick {
		return 0
	}
	return math.Round(GearMax*(pick-worn)/full*10) / 10
}

// ignoreUnscorable drops ErrGearUnscorable (the slot's gap is then null); any other error
// is returned.
func ignoreUnscorable(err error) error {
	if err == nil || errors.Is(err, sims.ErrGearUnscorable) {
		return nil
	}
	return err
}
