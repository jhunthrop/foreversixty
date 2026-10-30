package discordbot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// maxUpgrades is how many upgrade rows the /char panel lists (spec 2026-10-29 §A.5.3).
const maxUpgrades = 5

// tankRole is the spec role whose upgrades are ranked by score, not DPS.
const tankRole = "tank"

// weaponSlots are the slots whose empty gain no stat-weight score measures (the site's own
// upgrades.ts rule: a weapon's score bakes in its damage).
var weaponSlots = map[string]bool{"main_hand": true, "off_hand": true, "ranged": true}

// upgradeGroup orders the list: empty slots, then measured gains, then unmodelled picks.
type upgradeGroup int

const (
	groupEmpty upgradeGroup = iota
	groupMeasured
	groupUnmodelled
)

// upgrade is one slot whose pick beats what is worn: the pick, its gain in the list's unit
// (already worded: a figure or Gain not modelled yet), and what it replaces. kind is the
// pick's source kind when the BiS data names it, for the BiS page's source line; source the
// raw fallback.
type upgrade struct {
	slot, name, kind, source string
	id                       int
	gain                     string
	value                    float64
	group                    upgradeGroup
	worn                     *rankings.GearWorn
}

// upgradeList is the character's upgrades in the panel's order, whether they read in DPS,
// and the band they are against.
type upgradeList struct {
	items []upgrade
	byDPS bool
	band  string
}

// bySlot is the upgrades by slot key.
func (l upgradeList) bySlot() map[string]upgrade {
	out := make(map[string]upgrade, len(l.items))
	for _, u := range l.items {
		out[u.slot] = u
	}
	return out
}

// upgrades are the character's upgrades. A non-tank spec reads in DPS against the band /bis
// defaults to for this character (phaseOfBand of its level band: the Phase 1 raid at 60),
// each the band's own measured gain of its pick over the worn item, so an item shows the same
// gain here as in /bis. A tank, or a character the band measured no gain for, reads in Forever
// Score points from the gear panel's own gaps against its band.
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
	picks := band.BySlot()
	var items []upgrade
	for _, c := range slotChoices {
		sl, has := picks[c.Value]
		if !has || sl.ItemID <= 0 {
			continue
		}
		if u, ok := dpsUpgrade(band, sl, worn[c.Value]); ok {
			items = append(items, u)
		}
	}
	if len(items) == 0 {
		return upgradeList{}, false
	}
	return upgradeList{items: ordered(items), byDPS: true, band: bandLabels[phase]}, true
}

// dpsUpgrade is one slot's upgrade over w (nil: the slot is empty); ok is false when the pick
// is worn or measures no gain over it.
func dpsUpgrade(band bis.Band, sl bis.Slot, w *rankings.GearWorn) (upgrade, bool) {
	u := upgrade{slot: sl.Slot, id: sl.ItemID, name: sl.ItemName, kind: sl.SourceKind, source: sl.Source, worn: w}
	if w == nil {
		u.group, u.gain = groupEmpty, panelCopy.NotModelled
		if v, measured := emptyGain(band, sl); measured {
			u.gain, u.value = signedDPS(v), v
		}
		return u, true
	}
	v := bis.VerdictFor(sl, w.ID, true)
	if v.GainDps == nil || *v.GainDps < minShownGain {
		return upgrade{}, false
	}
	u.value = *v.GainDps
	if sl.EffectUnmodelled {
		u.group, u.gain = groupUnmodelled, panelCopy.NotModelled
		return u, true
	}
	u.group, u.gain = groupMeasured, signedDPS(u.value)
	return u, true
}

// emptyGain is the pick's gain over an empty slot: its whole stat-weight score in DPS (the
// site's upgrades.ts rule with a worn score of 0). measured is false for a weapon, a
// sim-decided pick (an on-use trinket carries no score), an unmodelled effect, or a band that
// states no DPS per point: no trustworthy figure exists then.
func emptyGain(band bis.Band, sl bis.Slot) (float64, bool) {
	if weaponSlots[sl.Slot] || sl.Score == nil || sl.EffectUnmodelled || band.ReferenceDpsPerPoint == nil {
		return 0, false
	}
	v := *sl.Score * *band.ReferenceDpsPerPoint
	return v, v >= minShownGain
}

// ordered is the panel's order: empty slots in frame order, then measured gains largest
// first, then unmodelled picks in frame order. items arrive in frame order.
func ordered(items []upgrade) []upgrade {
	out := append([]upgrade(nil), items...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].group != out[j].group {
			return out[i].group < out[j].group
		}
		return out[i].group == groupMeasured && out[i].value > out[j].value
	})
	return out
}

// scoreUpgrades are the gear panel's gaps, in points, against its bare band: empty slots
// first in frame order, then the largest gap.
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
	var items []upgrade
	for _, g := range upgradeSlots(card.Gear) {
		u := upgrade{slot: g.Slot, id: g.Pick.ID, name: g.Pick.Name, kind: kinds[g.Pick.ID], source: g.Pick.Source,
			gain: signedPoints(*g.Gap), value: *g.Gap, group: groupMeasured, worn: g.Worn}
		if g.Worn == nil {
			u.group = groupEmpty
		}
		items = append(items, u)
	}
	list.items = ordered(items)
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

// wornBySlot is the worn item of each slot that has one.
func wornBySlot(gear []rankings.GearSlot) map[string]*rankings.GearWorn {
	out := map[string]*rankings.GearWorn{}
	for _, g := range gear {
		if g.Worn != nil {
			out[g.Slot] = g.Worn
		}
	}
	return out
}

// upgradeSlots are the gear panel's slots, in frame order, with a pick that gains.
func upgradeSlots(gear []rankings.GearSlot) []rankings.GearSlot {
	at := map[string]int{}
	for i, c := range slotChoices {
		at[c.Value] = i
	}
	var out []rankings.GearSlot
	for _, g := range gear {
		if g.Pick != nil && g.Gap != nil && *g.Gap >= minShownGain && (g.Worn == nil || g.Worn.ID != g.Pick.ID) {
			out = append(out, g)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return at[out[i].Slot] < at[out[j].Slot] })
	return out
}

// upgradesText is the body's lead: Top upgrades for the band, then what they are ranked by,
// that no gear is shared, or that every slot wears the pick.
func upgradesText(card rankings.Character, list upgradeList) string {
	title := panelCopy.UpgradesTitle
	if list.band == "" {
		title, _, _ = strings.Cut(title, lineJoin)
	} else {
		title = fill(title, "{band label}", list.band)
	}
	switch {
	case len(card.Gear) == 0:
		return title + "\n" + lookupCopy.NoGear
	case len(list.items) == 0 && list.band != "":
		return title + "\n" + fill(panelCopy.UpgradesNone, "{band label}", list.band)
	case !list.byDPS:
		return title + "\n" + panelCopy.UpgradesByScore
	}
	return title + "\n" + panelCopy.UpgradesByDPS
}

// moreText is "-# {n} more upgrades on the character page.", "" when none are left.
func moreText(n int) string {
	switch {
	case n <= 0:
		return ""
	case n == 1:
		return panelCopy.UpgradesMoreOne
	}
	return fill(panelCopy.UpgradesMore, "{n}", fmt.Sprint(n))
}
