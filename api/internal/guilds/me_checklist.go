// api/internal/guilds/me_checklist.go
//
// The member's private "Before the raid" checklist (spec 2026-11-05-web-layout §7.1, D12):
// one character's checks for one raid, read from the same computation as the Readiness board
// (computeReadiness), with a verdict in words: ready, almost or getting_started. Only the
// signed-in member's own characters are ever read; no route takes another member's.
package guilds

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

// Checklist states and verdicts.
const (
	CheckOK         = "ok"
	CheckFail       = "fail"
	CheckNotTracked = "not_tracked"

	VerdictReady          = "ready"
	VerdictAlmost         = "almost"
	VerdictGettingStarted = "getting_started"
)

// Checklist item keys, in the order the list reads.
const (
	CheckLevel          = "level"
	CheckAttunement     = "attunement"
	CheckGearGap        = "gear_gap"
	CheckEnchants       = "enchants"
	CheckFlaskFood      = "flask_food"
	CheckFireResistance = "fire_resistance"
)

// Why a check is not tracked, for the web's words.
const (
	reasonNoLevel          = "no_level"
	reasonKeysNotInExport  = "keys_not_in_export"
	reasonGearNotShared    = "gear_not_shared"
	reasonNoBisList        = "no_bis_list"
	reasonLeveling         = "leveling"
	reasonNeedsConsumables = "needs_named_consumables"
	reasonNeedsResistance  = "needs_resistance_totals"
)

// ChecklistItem is one check. The detail fields carry only what that key reads.
type ChecklistItem struct {
	Key   string `json:"key"`
	State string `json:"state"`
	// Reason says why a not_tracked check has no reading.
	Reason string `json:"reason,omitempty"`
	// level
	Level         *int `json:"level,omitempty"`
	RequiredLevel *int `json:"required_level,omitempty"`
	// attunement
	KeyItem *KeyItem `json:"key_item,omitempty"`
	// gear_gap, in the Readiness board's units
	Upgrades *int     `json:"upgrades,omitempty"`
	GainDps  *float64 `json:"gain_dps,omitempty"`
	Unit     *string  `json:"unit,omitempty"`
	Slots    []string `json:"slots,omitempty"`
	// enchants
	MissingSlots []string `json:"missing_slots,omitempty"`
}

// ChecklistCharacter is the character the checklist is for.
type ChecklistCharacter struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
	Level        *int   `json:"level"`
}

// ChecklistRaid is the raid the checklist is for.
type ChecklistRaid struct {
	ID       int64  `json:"id"`
	ZoneKey  string `json:"zone_key"`
	ZoneName string `json:"zone_name"`
}

// ChecklistView is GET /v1/guilds/{id}/me/checklist.
type ChecklistView struct {
	Character ChecklistCharacter `json:"character"`
	Raid      *ChecklistRaid     `json:"raid"`
	Verdict   string             `json:"verdict"`
	Items     []ChecklistItem    `json:"items"`
}

// checklistInput is everything one checklist reads, already loaded.
type checklistInput struct {
	zone      *Zone
	readiness CharacterReadiness
	// row is the character's roster row; onRoster is false when it has none (no export).
	row      RosterRow
	onRoster bool
	band     bis.Band
	hasBand  bool
}

// buildChecklist is the checklist's items and verdict; pure.
func buildChecklist(in checklistInput) ([]ChecklistItem, string) {
	items := []ChecklistItem{levelCheck(in)}
	if in.zone != nil && in.zone.KeyItem != nil {
		items = append(items, attunementCheck(*in.zone.KeyItem, in.row))
	}
	items = append(items, gearGapCheck(in), enchantsCheck(in),
		ChecklistItem{Key: CheckFlaskFood, State: CheckNotTracked, Reason: reasonNeedsConsumables},
		ChecklistItem{Key: CheckFireResistance, State: CheckNotTracked, Reason: reasonNeedsResistance})
	return items, checklistVerdict(items)
}

// checklistVerdict is getting_started when the level is short or unknown or the key is
// missing, ready when every tracked check passes, else almost. Gear the member does not share
// is almost too: nothing on file is not "ready" (a missing BiS list, every healer's and tank's
// today, does not count against them).
func checklistVerdict(items []ChecklistItem) string {
	verdict := VerdictReady
	for _, it := range items {
		gate := it.Key == CheckLevel || it.Key == CheckAttunement
		switch {
		case gate && (it.State == CheckFail || (it.Key == CheckLevel && it.State == CheckNotTracked)):
			return VerdictGettingStarted
		case it.State == CheckFail || it.Reason == reasonGearNotShared:
			verdict = VerdictAlmost
		}
	}
	return verdict
}

// requiredLevel is the zone's level, raidLevel without a raid or for a run that names none.
func requiredLevel(zone *Zone) int {
	if zone != nil {
		if n := zoneMinLevel(*zone); n > 0 {
			return n
		}
	}
	return raidLevel
}

func levelCheck(in checklistInput) ChecklistItem {
	need := requiredLevel(in.zone)
	it := ChecklistItem{Key: CheckLevel, RequiredLevel: &need}
	level := in.readiness.Level
	switch {
	case level == nil:
		it.State, it.Reason = CheckNotTracked, reasonNoLevel
	case *level >= need:
		it.State, it.Level = CheckOK, level
	default:
		it.State, it.Level = CheckFail, level
	}
	return it
}

// attunementCheck is ok when the key is in the export's bags or worn. The export lists only
// gear in bags (Export.lua isEquippable) and never the keyring, so a key that is not there is
// not tracked rather than missing: absence proves nothing.
func attunementCheck(key KeyItem, row RosterRow) ChecklistItem {
	it := ChecklistItem{Key: CheckAttunement, KeyItem: &key}
	if attuned := attunedOf(key, row); attuned != nil {
		it.State = CheckOK
		return it
	}
	it.State, it.Reason = CheckNotTracked, reasonKeysNotInExport
	return it
}

// attunedOf is true when the key is in the row's bags (at bags consent) or worn (at gear
// consent); nil otherwise, never false (see attunementCheck).
func attunedOf(key KeyItem, row RosterRow) *bool {
	yes := true
	if row.Consent == "gear_bags" {
		for _, id := range row.bags {
			if id == key.ID {
				return &yes
			}
		}
	}
	if hasGearConsent(row.Consent) {
		for _, id := range row.gear {
			if id == key.ID {
				return &yes
			}
		}
	}
	return nil
}

func gearGapCheck(in checklistInput) ChecklistItem {
	it := ChecklistItem{Key: CheckGearGap}
	cr := in.readiness
	if !cr.GearChecked {
		it.State, it.Reason = CheckNotTracked, gearNotTrackedReason(in)
		return it
	}
	upgrades, gain := cr.GearGap.Upgrades, cr.GearGap.GainDps
	it.Upgrades, it.Slots = &upgrades, gapSlots(in.band, in.row.gear)
	if cr.GearUnit != "" {
		unit := cr.GearUnit
		it.Unit, it.GainDps = &unit, &gain
	}
	it.State = CheckOK
	if upgrades > 0 {
		it.State = CheckFail
	}
	return it
}

// gearNotTrackedReason says why the gear gap has no reading.
func gearNotTrackedReason(in checklistInput) string {
	switch {
	case in.readiness.Leveling:
		return reasonLeveling
	case !in.onRoster || !gearOnFile(in.row.Consent, in.row.gear):
		return reasonGearNotShared
	}
	return reasonNoBisList
}

// gapSlots is every slot the band names a pick for that the worn gear is an upgrade away from,
// in the band's order (bis.GapFor's own rule).
func gapSlots(band bis.Band, gear map[string]int) []string {
	out := []string{}
	for _, slot := range band.Slots {
		if slot.ItemID <= 0 {
			continue
		}
		worn, has := gear[slot.Slot]
		if bis.VerdictFor(slot, worn, has).Upgrade {
			out = append(out, slot.Slot)
		}
	}
	return out
}

func enchantsCheck(in checklistInput) ChecklistItem {
	it := ChecklistItem{Key: CheckEnchants}
	cr := in.readiness
	switch {
	case !cr.EnchantsChecked && cr.Leveling:
		it.State, it.Reason = CheckNotTracked, reasonLeveling
	case !cr.EnchantsChecked:
		it.State, it.Reason = CheckNotTracked, reasonGearNotShared
	case len(cr.MissingEnchantSlots) > 0:
		it.State, it.MissingSlots = CheckFail, cr.MissingEnchantSlots
	default:
		it.State = CheckOK
	}
	return it
}

// readinessKeyItem is the key the board's raid asks for: the filtered raid's zone, else the
// guild's next raid's; nil when it needs none or none is planned.
func (s *Store) readinessKeyItem(ctx context.Context, guildID int64, raid *Raid) (*KeyItem, error) {
	if raid == nil {
		next, ok, err := s.NextRaid(ctx, guildID, time.Now())
		if err != nil || !ok {
			return nil, err
		}
		raid = &next
	}
	return zoneOf(raid.Zone).KeyItem, nil
}
