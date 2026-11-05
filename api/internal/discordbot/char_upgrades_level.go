package discordbot

import (
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// levelUpgrades is a healer's or tank's list (gear panel spec §A.5.3, amended 2026-11-05): the
// band /bis defaults to, each slot whose worn item is not the band's pick, ranked by the BiS
// list and item level, never DPS. Empty slots come first, then picks above what is worn by the
// most item levels, then the list's other picks in frame order. ok is false when the band is
// unknown or every slot wears its pick.
func (s *Service) levelUpgrades(card rankings.Character, spec specs.Spec, faction string) (upgradeList, bool) {
	level, ok := upgrades.LevelBand(card.Character.Level)
	if !ok {
		return upgradeList{}, false
	}
	band, found := s.loadBand(spec, level, faction, upgrades.PresetFor(level))
	if !found {
		return upgradeList{}, false
	}
	build, _ := s.d.Guilds.ActiveBuild()
	class := strings.ToLower(card.Character.Class)
	itemLevel := func(id int) int {
		if it, ok := s.d.Guilds.BuildItem(build, class, id); ok {
			return it.ItemLevel
		}
		return 0
	}
	worn := wornBySlot(card.Gear)
	picks := band.BySlot()
	var items []upgrade
	for _, c := range upgrades.Slots {
		sl, has := picks[c.Key]
		if !has || sl.ItemID <= 0 {
			continue
		}
		if u, ok := levelUpgrade(sl, worn[c.Key], itemLevel); ok {
			items = append(items, u)
		}
	}
	if len(items) == 0 {
		return upgradeList{}, false
	}
	return upgradeList{items: ordered(items), rank: rankLevel, band: bandLabels[phaseOfBand(level)]}, true
}

// levelUpgrade is one slot's pick over w (nil: the slot is empty) in item levels; ok is false
// when the pick is what is worn. A pick no higher than the worn item is still the list's pick.
func levelUpgrade(sl bis.Slot, w *rankings.GearWorn, itemLevel func(int) int) (upgrade, bool) {
	u := upgrade{slot: sl.Slot, id: sl.ItemID, name: sl.ItemName, kind: sl.SourceKind, source: sl.Source, worn: w}
	if w == nil {
		u.group, u.gain = groupEmpty, panelCopy.LevelEmpty
		return u, true
	}
	if w.ID == sl.ItemID {
		return upgrade{}, false
	}
	wornLevel := itemLevel(w.ID)
	if w.ItemLevel != nil {
		wornLevel = *w.ItemLevel
	}
	gap := itemLevel(sl.ItemID) - wornLevel
	if gap <= 0 || wornLevel == 0 {
		u.group, u.gain = groupUnmodelled, panelCopy.LevelOnList
		return u, true
	}
	u.group, u.value, u.gain = groupMeasured, float64(gap), levelGain(gap)
	return u, true
}

// levelGain is "+12 item levels", "+1 item level".
func levelGain(n int) string {
	if n == 1 {
		return panelCopy.LevelGainOne
	}
	return fill(panelCopy.LevelGain, "{n}", strconv.Itoa(n))
}
