package coach

// Drafts is every pack string the spec does not show, by key. They are drafts: Task 33's
// experience spec fixes them before any surface ships, and its spec test pins them then.
// A key starting "title_" is a finding title and keeps §5.2's shape (under eight words,
// at most one number); "night_" keys are night-line stems (§5.5) for codes other than
// core casts, whose stem is NightCoreCasts.
var Drafts = map[string]string{
	// F1: the detail when there is one counted gap.
	"detail_core_one_gap": "{actual} of a possible {possible} with {resource} in your bar. Longest gap {gap1}.",

	// F3 dead time.
	"title_idle":           "Nothing cast for {n} seconds",
	"detail_idle":          "{n} s idle of your {counted} counted, longest at {spans}.",
	"basis_idle_threshold": "From your log: dead time counts from 10 s and 5% of your counted time.",

	// F4: a cap with no downtime to leave out, and mana that ran low to the end.
	"detail_capped_no_downtime": "Your {resource} sat at its cap for {n} s, longest at {span}.",
	"detail_out_of_mana_to_end": "Your mana stayed under 5% from {t} to the end, {n} s in all.",

	// F5: no level in the log.
	"cannot_level": "Spell ranks: your log doesn't record your level.",

	// F6 preparation.
	"title_no_food":           "No food at the pull",
	"title_no_weapon_enchant": "No weapon oil or stone",
	"title_no_enchant":        "No enchant on {Item}",
	"title_potions":           "Combat potions: {n} used",
	"title_raid_buff":         "No {Buff} at the pull",
	"cannot_weapon_enchant":   "Weapon oils and stones: the coach can't read them from the log yet.",

	// F7: a melee swing's killing blow.
	"melee": "{Boss}'s melee",

	// F8 adherence.
	"title_adherence":  "{pct}% of casts from the rotation",
	"detail_adherence": "{pct}% of your casts were the written rotation's. Most cast outside it: {spells}.",

	// Strengths the §5.3 examples do not show.
	"strength_cpm":         "{Spell} at {you} a minute, at or above the comparison.",
	"strength_adherence":   "{pct}% of casts from the rotation, at or above the comparison.",
	"strength_overheal":    "{Spell} (Rank {r}) at {pct}% overheal, at or under the comparison.",
	"strength_consumables": "Every consumable the coach can check was up at the pull.",

	// Links the §5.3 examples do not show.
	"link_death_in_gap":      "You died in the {Spell} gap at {gap}.",
	"link_consumable_covers": "{Item} would have covered this.",

	// Night-line stems for codes other than core casts.
	"night_uptime":      "{Aura} uptime",
	"night_idle":        "Dead time",
	"night_resource":    "{Resource} wasted",
	"night_ranks":       "Spell ranks",
	"night_preparation": "Preparation at the pull",
	"night_death":       "Deaths",
	"night_adherence":   "Rotation adherence",
}
