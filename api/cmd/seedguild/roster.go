// api/cmd/seedguild/roster.go
//
// The mock roster this tool seeds: 24 WoW-plausible characters (no real player or
// streamer names) covering all nine classes, a tank, six healers and seventeen DPS, two
// Skyborne, two officers and three characters left unverified so the guild home's approve
// button has rows to act on. Pure data plus pure helpers — no I/O, no randomness — so
// rosterTest can assert the shape without a database.
//
// Depth for the raid-control-centre panels (Roster/Raids/Progression/Readiness/Loot,
// 2026-10-04 scope note): attendance varies per raid night (Account, SkipReports),
// deaths concentrate on a few raiders (DeathProne), two accounts carry a main and an alt
// (Account, shared between two mockCharacter rows), and gear readiness varies
// (NonBisSlots, MissingEnchantSlot, UnspentTalents - see readiness.go for how each turns
// into the actual FS1 export).
package main

// role is the three roles fight_metrics.role and the home roster care about.
const (
	roleTank   = "tank"
	roleHealer = "healer"
	roleDPS    = "dps"
)

// mockCharacter is one roster entry: who they are, what they play, and how this seed
// treats their membership (rank, verification, consent) and raid presence.
type mockCharacter struct {
	Name      string
	Class     string // lowercase class slug, matches fs1's ClassSlug
	Spec      string // display spec name, used to pick a BiS file and to label talents
	Role      string
	RaceSlug  string
	BisFile   string // basename under data/builds/<build>/bis, "" when no BiS exists for this spec
	Officer   bool
	Unverifed bool // true for the 3 characters an officer still needs to approve
	Consent   string
	ItemLevel int

	// Account groups two characters under one account (a main and an alt) when set to
	// the same non-empty value; empty means "this character's own account". Two
	// characters sharing an Account must carry the same Consent (consent is a
	// guild_members column, one row per account) - mockRoster's own pairs do.
	Account string

	// SkipReports names the raidPlans tags (see raid.go) this character does not
	// attend - fight_metrics carries no row for them on those nights, the same as a
	// raider who did not show up. Nil/empty means full attendance.
	SkipReports []string

	// DeathProne weights this character into rollDeaths' victim pool more heavily - a
	// few raiders who reliably find the fire, the way most real guilds have one.
	DeathProne bool

	// NonBisSlots is how many of this character's BiS gear slots are swapped for a
	// real alternative item from that slot's own BiS alternatives list (gear.go),
	// rather than the top pick - a readiness gap, not a fabricated item.
	NonBisSlots int

	// MissingEnchantSlot is one gear slot (by name) that, for this character alone,
	// carries no enchant even though every other character's copy of that slot would -
	// a readiness gap the future readiness board (design spec §9) is meant to surface.
	MissingEnchantSlot string

	// UnspentTalents is true for a character whose talent string deliberately totals
	// fewer than the level-60 maximum.
	UnspentTalents bool
}

// mockRoster is the fixed 24-character roster, in the order the seed writes them. The
// order itself is not meaningful to the database; it is kept stable so a dry run's "first
// two rows" and a rerun's row keys line up exactly.
func mockRoster() []mockCharacter {
	return []mockCharacter{
		{Name: "Thornwicke", Class: "warrior", Spec: "Protection", Role: roleTank, RaceSlug: "human", BisFile: "",
			Consent: "gear", ItemLevel: 62, Account: "thornwicke-household"},
		{Name: "Ironbrand", Class: "paladin", Spec: "Holy", Role: roleHealer, RaceSlug: "human", BisFile: "",
			Officer: true, Consent: "gear", ItemLevel: 63},
		{Name: "Valomira", Class: "paladin", Spec: "Holy", Role: roleHealer, RaceSlug: "human", BisFile: "",
			Consent: "gear", ItemLevel: 60, SkipReports: []string{"mc2"}},
		{Name: "Seraphine", Class: "paladin", Spec: "Retribution", Role: roleDPS, RaceSlug: "human", BisFile: "paladin-retribution.json",
			Consent: "gear", ItemLevel: 65, Account: "thornwicke-household"},
		{Name: "Dunmaro", Class: "hunter", Spec: "BeastMastery", Role: roleDPS, RaceSlug: "dwarf", BisFile: "hunter-beast-mastery.json",
			Consent: "gear", ItemLevel: 61, SkipReports: []string{"mc1"}, NonBisSlots: 2},
		{Name: "Pinefeather", Class: "hunter", Spec: "Marksmanship", Role: roleDPS, RaceSlug: "dwarf", BisFile: "hunter-marksmanship.json",
			Consent: "gear", ItemLevel: 58, NonBisSlots: 3},
		{Name: "Skytalon", Class: "hunter", Spec: "Survival", Role: roleDPS, RaceSlug: "skyborne", BisFile: "hunter-survival.json",
			Consent: "gear", ItemLevel: 59, DeathProne: true},
		{Name: "Shiverknife", Class: "rogue", Spec: "Assassination", Role: roleDPS, RaceSlug: "gnome", BisFile: "rogue-assassination.json",
			Officer: true, Consent: "gear", ItemLevel: 66, Account: "shiverknife-household", MissingEnchantSlot: "wrist"},
		{Name: "Blacktide", Class: "rogue", Spec: "Combat", Role: roleDPS, RaceSlug: "gnome", BisFile: "rogue-combat.json",
			Consent: "gear", ItemLevel: 57, SkipReports: []string{"mc1", "mc2", "ony2"}},
		{Name: "Nightlatch", Class: "rogue", Spec: "Subtlety", Role: roleDPS, RaceSlug: "human", BisFile: "rogue-subtlety.json",
			Unverifed: true, Consent: "roster", ItemLevel: 55, DeathProne: true},
		{Name: "Hollowmend", Class: "priest", Spec: "Holy", Role: roleHealer, RaceSlug: "gnome", BisFile: "",
			Consent: "gear", ItemLevel: 64, SkipReports: []string{"ony1", "ony2"}},
		{Name: "Dawnweaver", Class: "priest", Spec: "Holy", Role: roleHealer, RaceSlug: "human", BisFile: "",
			Consent: "roster", ItemLevel: 60, SkipReports: []string{"ony2"}},
		{Name: "Voidcall", Class: "priest", Spec: "Shadow", Role: roleDPS, RaceSlug: "gnome", BisFile: "priest-shadow.json",
			Unverifed: true, Consent: "gear", ItemLevel: 56, UnspentTalents: true},
		{Name: "Stonetide", Class: "shaman", Spec: "Restoration", Role: roleHealer, RaceSlug: "dwarf", BisFile: "",
			Consent: "gear", ItemLevel: 62, SkipReports: []string{"ony1"}},
		{Name: "Emberfall", Class: "shaman", Spec: "Elemental", Role: roleDPS, RaceSlug: "dwarf", BisFile: "shaman-elemental.json",
			Consent: "gear", ItemLevel: 63, NonBisSlots: 4},
		{Name: "Groundshock", Class: "shaman", Spec: "Enhancement", Role: roleDPS, RaceSlug: "dwarf", BisFile: "shaman-enhancement.json",
			Consent: "gear_bags", ItemLevel: 67, DeathProne: true},
		{Name: "Frostquill", Class: "mage", Spec: "Frost", Role: roleDPS, RaceSlug: "gnome", BisFile: "mage-frost.json",
			Consent: "gear", ItemLevel: 68, SkipReports: []string{"mc1"}, NonBisSlots: 2},
		{Name: "Cinderglass", Class: "mage", Spec: "Fire", Role: roleDPS, RaceSlug: "gnome", BisFile: "mage-fire.json",
			Consent: "gear", ItemLevel: 59},
		{Name: "Skyarcana", Class: "mage", Spec: "Arcane", Role: roleDPS, RaceSlug: "skyborne", BisFile: "mage-arcane.json",
			Unverifed: true, Consent: "gear", ItemLevel: 58},
		{Name: "Doomwhisper", Class: "warlock", Spec: "Affliction", Role: roleDPS, RaceSlug: "gnome", BisFile: "warlock-affliction.json",
			Consent: "gear", ItemLevel: 69, SkipReports: []string{"mc1"}},
		{Name: "Grimsworn", Class: "warlock", Spec: "Demonology", Role: roleDPS, RaceSlug: "gnome", BisFile: "warlock-demonology.json",
			Consent: "gear", ItemLevel: 60, SkipReports: []string{"mc1", "mc2", "ony1"}},
		{Name: "Hexmarrow", Class: "warlock", Spec: "Destruction", Role: roleDPS, RaceSlug: "gnome", BisFile: "warlock-destruction.json",
			Consent: "gear", ItemLevel: 61, Account: "shiverknife-household", SkipReports: []string{"mc2"}, MissingEnchantSlot: "main_hand"},
		{Name: "Mossveil", Class: "druid", Spec: "Restoration", Role: roleHealer, RaceSlug: "night-elf", BisFile: "",
			Consent: "gear", ItemLevel: 64},
		{Name: "Starnettle", Class: "druid", Spec: "Balance", Role: roleDPS, RaceSlug: "night-elf", BisFile: "druid-balance.json",
			Consent: "gear", ItemLevel: 70, SkipReports: []string{"ony2"}, UnspentTalents: true},
	}
}

// classes returns the distinct class slugs mockRoster covers.
func classes(roster []mockCharacter) map[string]bool {
	out := map[string]bool{}
	for _, c := range roster {
		out[c.Class] = true
	}
	return out
}

// countBy tallies roster rows matching pred.
func countBy(roster []mockCharacter, pred func(mockCharacter) bool) int {
	n := 0
	for _, c := range roster {
		if pred(c) {
			n++
		}
	}
	return n
}

// accountKey is the account a character belongs to: its own Account group, or its own
// name when it does not share an account with anyone.
func (c mockCharacter) accountKey() string {
	if c.Account != "" {
		return c.Account
	}
	return c.Name
}

// attends reports whether this character is present at the raid night tagged reportTag.
func (c mockCharacter) attends(reportTag string) bool {
	for _, t := range c.SkipReports {
		if t == reportTag {
			return false
		}
	}
	return true
}
