package discordbot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// upgrade is one of the card's top three: a pick for a slot and its gain in the list's unit.
// kind is the pick's source kind when the BiS data names it ("" for a gear-panel pick the
// band does not list), for the BiS page's source line; source is the raw fallback.
type upgrade struct {
	slot, name, kind, source string
	id                       int
	gain                     string
}

// upgradeList is the card's upgrades, whether they read in DPS, and the band they are against.
type upgradeList struct {
	items []upgrade
	byDPS bool
	band  string
}

// upgrades are the card's top three. A non-tank spec reads in DPS against the band /bis
// defaults to for this character (phaseOfBand of its level band: the Phase 1 raid at 60),
// each the band's own measured gain of its pick over the worn item, largest first, so an
// item shows the same gain here as in /bis. A tank, or a character the band measured no
// gain for, reads in Forever Score points from the gear panel's own gaps against its band.
func (s *Service) upgrades(card rankings.Character, spec specs.Spec, guildFaction string) upgradeList {
	faction := s.characterFaction(card, guildFaction)
	if spec.Role != tankRole {
		if list, ok := s.dpsUpgrades(card, spec, faction); ok {
			return list
		}
	}
	return s.scoreUpgrades(card, spec, faction)
}

// characterFaction is the character's own faction (its race's), else the guild's.
func (s *Service) characterFaction(card rankings.Character, guildFaction string) string {
	if f := s.d.Guilds.RaceFaction(deref(card.Character.Race)); f != "" {
		return f
	}
	return guildFaction
}

// loadBand is the spec's band for faction; found is false when any part is unknown.
func (s *Service) loadBand(spec specs.Spec, band int, faction, preset string) (bis.Band, bool) {
	build := s.d.Guilds.BisBuild()
	if spec.Spec == "" || faction == "" || build == "" {
		return bis.Band{}, false
	}
	b, found, err := bis.LoadPresetBand(s.d.Guilds.BisDataDir(), build, spec.Spec, band, faction, preset)
	if err != nil {
		s.log.Warn("discordbot", "op", "char_band", "spec", spec.Spec, "err", err)
	}
	return b, found
}

func (s *Service) dpsUpgrades(card rankings.Character, spec specs.Spec, faction string) (upgradeList, bool) {
	level, ok := levelBand(card.Character.Level)
	if !ok {
		return upgradeList{}, false
	}
	phase := phaseOfBand(level)
	bandLevel, preset := phaseBand(phase)
	band, found := s.loadBand(spec, bandLevel, faction, preset)
	if !found {
		return upgradeList{}, false
	}
	worn := wornBySlot(card.Gear)
	type measured struct {
		slot bis.Slot
		gain float64
	}
	var gains []measured
	for _, sl := range band.Slots {
		if sl.ItemID <= 0 {
			continue
		}
		id, has := worn[sl.Slot]
		if v := bis.VerdictFor(sl, id, has); v.GainDps != nil && *v.GainDps > 0 {
			gains = append(gains, measured{sl, *v.GainDps})
		}
	}
	if len(gains) == 0 {
		return upgradeList{}, false
	}
	sort.SliceStable(gains, func(i, j int) bool { return gains[i].gain > gains[j].gain })
	list := upgradeList{byDPS: true, band: bandLabels[phase]}
	for _, g := range gains[:min(len(gains), maxUpgrades)] {
		list.items = append(list.items, upgrade{slot: g.slot.Slot, id: g.slot.ItemID, name: g.slot.ItemName,
			kind: g.slot.SourceKind, source: g.slot.Source, gain: signedDPS(g.gain)})
	}
	return list, true
}

// scoreUpgrades are the gear panel's three largest gaps, in points, against its bare band.
func (s *Service) scoreUpgrades(card rankings.Character, spec specs.Spec, faction string) upgradeList {
	list := upgradeList{}
	level, ok := levelBand(card.Character.Level)
	kinds := map[int]string{}
	if ok {
		list.band = bandLabels[fmt.Sprint(level)]
		if band, found := s.loadBand(spec, level, faction, bis.BarePreset); found {
			kinds = sourceKinds(band)
		}
	}
	for _, g := range upgradeSlots(card.Gear) {
		list.items = append(list.items, upgrade{slot: g.Slot, id: g.Pick.ID, name: g.Pick.Name,
			kind: kinds[g.Pick.ID], source: g.Pick.Source, gain: signedPoints(*g.Gap)})
	}
	return list
}

// sourceKinds is every item the band names, pick or alternative, by id, with its source kind.
func sourceKinds(band bis.Band) map[int]string {
	out := map[int]string{}
	for _, sl := range band.Slots {
		out[sl.ItemID] = sl.SourceKind
		for _, alt := range sl.Alternatives {
			if _, seen := out[alt.ItemID]; !seen {
				out[alt.ItemID] = alt.SourceKind
			}
		}
	}
	return out
}

func wornBySlot(gear []rankings.GearSlot) map[string]int {
	out := map[string]int{}
	for _, g := range gear {
		if g.Worn != nil {
			out[g.Slot] = g.Worn.ID
		}
	}
	return out
}

// upgradeSlots are the gear panel's slots with a pick that gains, largest gap first, at most three.
func upgradeSlots(gear []rankings.GearSlot) []rankings.GearSlot {
	var out []rankings.GearSlot
	for _, g := range gear {
		if g.Pick != nil && g.Gap != nil && *g.Gap > 0 && (g.Worn == nil || g.Worn.ID != g.Pick.ID) {
			out = append(out, g)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return *out[i].Gap > *out[j].Gap })
	return out[:min(len(out), maxUpgrades)]
}

// upgradesField says what the upgrades are ranked by and against which band, or that no gear
// is shared.
func upgradesField(card rankings.Character, list upgradeList) notify.EmbedField {
	if len(card.Gear) == 0 {
		return notify.EmbedField{Name: lookupCopy.TopUpgrades, Value: lookupCopy.NoGear}
	}
	name := lookupCopy.TopUpgrades
	if list.band != "" {
		name += fill(lookupCopy.UpgradesBand, "{band label}", list.band)
	}
	value := lookupCopy.UpgradesByScore
	if list.byDPS {
		value = lookupCopy.UpgradesByDPS
	}
	return notify.EmbedField{Name: name, Value: value}
}

// upgradeEmbeds are the upgrades as item embeds, each with its gain and the BiS page's source line.
func (s *Service) upgradeEmbeds(card rankings.Character, list upgradeList, faction string) []notify.Embed {
	build, _ := s.d.Guilds.ActiveBuild()
	class := strings.ToLower(card.Character.Class)
	out := make([]notify.Embed, 0, len(list.items))
	for _, u := range list.items {
		source := s.sourceLine(u.id, u.kind, faction, u.source)
		url := s.characterURL(card.Character.Key) + "#gear-slot-" + u.slot
		out = append(out, ItemEmbed(s.itemFacts(build, class, u.id, u.name, source, url), slotLabel(u.slot), u.gain, ""))
	}
	return out
}
