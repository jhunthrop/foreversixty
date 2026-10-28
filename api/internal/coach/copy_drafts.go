package coach

// Drafts is every pack sentence neither spec shows, by key: the experience spec's §8.1
// table (design/specs/2026-10-28-log-coach.md) fixed the rest, which copy.go now holds.
// A key starting "title_" is a finding title and keeps §5.2's shape (under eight words,
// at most one number).
var Drafts = map[string]string{
	// A pool whose median lacks the measurement (comparison): not a thin pool, so never
	// ThinPool's line, which would be false.
	"pool_lacks_measure": "Top {Spec} logs on {Boss} ({size}) don't record this yet.",

	// F6: weapon oils and stones (a temporary enchant; TitleNoWeaponEnchant is the
	// permanent one), and the strength when the sheet shows no flask but nothing missing.
	"title_no_weapon_oil":   "No weapon oil or stone",
	"detail_no_weapon_oil":  "You pulled with no weapon oil or stone.",
	"cannot_weapon_enchant": "Weapon oils and stones: the coach can't read them from the log yet.",
	"strength_consumables":  "Every consumable the coach can check was up at the pull.",

	// F7: the avoidable fact kept when a defensive was ready (killing blow, then avoidable,
	// then the defensive: §6.2's order); {since} is "the pull" or the time it came back.
	"detail_died_with_ready_avoidable": "{Killer} killed you at {t}. {Killer} is on this boss's list of damage you can avoid. {Spell} had been ready since {since}.",
	// F7 narrowed: what it could not check, said so (§20.4).
	"cannot_defensives":      "Defensive cooldowns: the coach's list for your spec isn't written yet.",
	"cannot_defensive_level": "Defensive cooldowns you never cast: your log doesn't record your level.",
	"cannot_death_killer":    "Your death at {t}: the log doesn't say what killed you.",
}
