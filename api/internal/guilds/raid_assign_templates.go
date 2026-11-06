// api/internal/guilds/raid_assign_templates.go
//
// The raid assignment templates (spec 2026-11-05-web-layout §7.2, owner decision: manual only,
// per encounter). Each boss of the client's raid table (raid_bosses_table.go) has one template:
// the slots a raid leader fills before the pull, in the order the call goes out. A boss the
// client adds later needs only its rows here; raid_assign_templates_test.go fails until it has
// them. A raid whose bosses the client does not carry yet reads "not in the game client yet".
package guilds

// Slot role filters: the roster group a slot is for, or any. The filter only orders the picker
// and flags an off-role pick; it never refuses one.
const (
	slotRoleTank   = "tank"
	slotRoleHealer = "healer"
	slotRoleMelee  = "melee"
	slotRoleRanged = "ranged"
	slotRoleAny    = "any"
)

// maxSlotKeyLen is raid_assignments.slot_key's check.
const maxSlotKeyLen = 40

// slotTemplate is one assignment slot. Count is how many names the slot takes at most; Warning
// is the short tag a raid warning line starts with ("MT"); Purpose is the one line that says
// what the slot does in the fight.
type slotTemplate struct {
	Key     string
	Label   string
	Warning string
	Purpose string
	Count   int
	Role    string
}

// encounterTemplates is each boss's slots by its DungeonEncounter id.
var encounterTemplates = map[int64][]slotTemplate{
	// Onyxia (Onyxia's Lair, 40): phase 1 on the ground, phase 2 in the air with whelps from
	// both caves and Deep Breath, phase 3 back down with Bellowing Roar and more whelps. Nothing
	// she casts is interrupted, and there is no curse, poison or magic to clear, so the template
	// has no interrupt or dispel slots.
	1084: {
		{Key: "main_tank", Label: "Main tank", Warning: "MT", Count: 1, Role: slotRoleTank,
			Purpose: "Picks her up and holds her side-on to the raid in phases 1 and 3."},
		{Key: "off_tank", Label: "Off tank (phase 3, fear)", Warning: "OT", Count: 2, Role: slotRoleTank,
			Purpose: "Builds threat from the landing and takes her when Bellowing Roar or Knock Away drops the main tank."},
		{Key: "main_tank_heals", Label: "Main tank healers", Warning: "MT heals", Count: 4, Role: slotRoleHealer,
			Purpose: "Heal the main tank only, through Flame Breath and Cleave, from behind her flank."},
		{Key: "off_tank_heals", Label: "Off tank healers", Warning: "OT heals", Count: 2, Role: slotRoleHealer,
			Purpose: "Keep the off tanks up, and swap to whoever has her after a fear."},
		{Key: "raid_heals", Label: "Raid healers", Warning: "Raid heals", Count: 6, Role: slotRoleHealer,
			Purpose: "Fireballs in phase 2, then whelp bites and lava eruptions in phase 3."},
		{Key: "fear_cover", Label: "Fear cover", Warning: "Fear Ward / Tremor", Count: 2, Role: slotRoleAny,
			Purpose: "Fear Ward on the main tank, or a Tremor Totem by the tank group, before each Bellowing Roar."},
		{Key: "whelps_west", Label: "Whelps, west cave", Warning: "Whelps W", Count: 3, Role: slotRoleAny,
			Purpose: "Pick up the west cave's whelps in phases 2 and 3 and kill them away from the raid."},
		{Key: "whelps_east", Label: "Whelps, east cave", Warning: "Whelps E", Count: 3, Role: slotRoleAny,
			Purpose: "Pick up the east cave's whelps in phases 2 and 3 and kill them away from the raid."},
		{Key: "breath_callers", Label: "Deep Breath callers", Warning: "Breath calls", Count: 2, Role: slotRoleAny,
			Purpose: "Call her flight line on voice so the raid clears it before Deep Breath lands."},
		{Key: "p2_ranged_west", Label: "Phase 2 ranged, west side", Warning: "P2 west", Count: 10, Role: slotRoleRanged,
			Purpose: "Spread along the west wall and keep casting while she is in the air."},
		{Key: "p2_ranged_east", Label: "Phase 2 ranged, east side", Warning: "P2 east", Count: 10, Role: slotRoleRanged,
			Purpose: "Spread along the east wall and keep casting while she is in the air."},
		{Key: "soulstones", Label: "Soulstones and rebirths", Warning: "SS / Rez", Count: 3, Role: slotRoleAny,
			Purpose: "Soulstone a healer before the pull; Rebirth goes to a tank first."},
	},
}

// templateOf is the boss's slots; false when the boss has none.
func templateOf(encounterID int64) ([]slotTemplate, bool) {
	slots, ok := encounterTemplates[encounterID]
	return slots, ok
}

// slotOf finds one slot of the boss's template.
func slotOf(encounterID int64, key string) (slotTemplate, bool) {
	slots, _ := templateOf(encounterID)
	for _, s := range slots {
		if s.Key == key {
			return s, true
		}
	}
	return slotTemplate{}, false
}

// tierRaidOf is the client's raid for a zone key; false when the client carries none for it.
func tierRaidOf(zoneKey string) (tierRaid, bool) {
	for _, r := range firstTierRaids {
		if r.Key == zoneKey {
			return r, true
		}
	}
	return tierRaid{}, false
}

// encounterOf is the boss of the zone's raid with that id, when the raid has it.
func encounterOf(zoneKey string, encounterID int64) (tierBoss, bool) {
	raid, ok := tierRaidOf(zoneKey)
	if !ok {
		return tierBoss{}, false
	}
	for _, b := range raid.Bosses {
		if b.EncounterID == encounterID {
			return b, true
		}
	}
	return tierBoss{}, false
}
