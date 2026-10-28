package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// simRefFactions is the order a ladder entry's faction is looked for: the ladder ranks
// Alliance first, as the spec guides list their recommended races.
var simRefFactions = []string{"alliance", "horde"}

// simRefName names the reference character in the engine's request.
const simRefName = "coach-reference"

// ladderCharacters is the coach-simref job's SimRefCharacters: the BiS ladder's own band
// character (sim/cmd/leveling-bis bandCharacter): the bare preset's race and talents from
// <dataDir>/<build>/bis/<spec>.json, at the band's level, in the band's picks (sims.BandPicks).
type ladderCharacters struct{ dataDir string }

var _ coach.SimRefCharacters = ladderCharacters{}

// ladderEntry is the part of one bis/<spec>.json band entry the reference needs.
type ladderEntry struct {
	Band    int    `json:"band"`
	Faction string `json:"faction"`
	Preset  string `json:"preset"`
	Race    string `json:"race"`
	Talents string `json:"talents"`
}

func (l ladderCharacters) Character(ctx context.Context, build, spec string, band int) (simapi.CharacterSpec, error) {
	entry, err := l.entry(build, spec, band)
	if err != nil {
		return simapi.CharacterSpec{}, err
	}
	picks, err := sims.BandPicks(ctx, build, spec, band, entry.Faction)
	if err != nil {
		return simapi.CharacterSpec{}, fmt.Errorf("coach sim ref: band picks: %w", err)
	}
	classSlug, _, _ := strings.Cut(spec, "-")
	ch := simapi.CharacterSpec{Name: simRefName, Race: entry.Race, Class: classSlug, Level: band,
		Talents: entry.Talents, Gear: gearOf(picks)}
	if leveling.NoMeleeAutoAttackSpecs[spec] {
		ch.DistanceFromTarget = leveling.CasterDistanceFromTarget
	}
	return ch, nil
}

// entry is the bare-preset band entry of spec in build, Alliance first.
func (l ladderCharacters) entry(build, spec string, band int) (ladderEntry, error) {
	path := filepath.Join(l.dataDir, build, "bis", spec+".json")
	raw, err := os.ReadFile(path)
	if err != nil {
		return ladderEntry{}, fmt.Errorf("coach sim ref: %w", err)
	}
	var file struct {
		Bands []ladderEntry `json:"bands"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		return ladderEntry{}, fmt.Errorf("coach sim ref: decode %s: %w", path, err)
	}
	for _, faction := range simRefFactions {
		for _, e := range file.Bands {
			bare := e.Preset == "" || e.Preset == sims.BarePreset
			if e.Band == band && e.Faction == faction && bare && e.Race != "" {
				return e, nil
			}
		}
	}
	return ladderEntry{}, fmt.Errorf("coach sim ref: no bare band %d entry for %s in %s", band, spec, build)
}

// gearOf is the picks as the engine's gear, in the export's slot order.
func gearOf(picks map[string]sims.BandPick) []simapi.GearSlot {
	gear := make([]simapi.GearSlot, 0, len(picks))
	for _, slot := range sims.GearSlots {
		if p, ok := picks[slot]; ok {
			gear = append(gear, simapi.GearSlot{Slot: slot, ItemID: p.ItemID})
		}
	}
	return gear
}
