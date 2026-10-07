package leveling

// NoMeleeAutoAttackSpecs is every spec whose real character never
// auto-attacks in melee at all (bis-ranker-integrity-5, item 7): a
// caster spec does not auto-attack in the leveling tournaments.
// wowsims-forever's engine gates a melee auto-swing purely on distance
// (sim/core/attack.go: a swing only starts when unit.DistanceFromTarget
// <= MaxMeleeAttackDistance, 5 yards), never on class or spec - every
// registered caster agent unconditionally enables melee auto-attacks,
// so the only lever that stops it is standing the character out of
// melee range (CasterDistanceFromTarget). Before this existed every
// caster stood in melee range, and Manual Crowd Pummeler won
// shaman-elemental and druid-balance's level-60 main_hand tournament
// off free melee DPS a real caster could never generate. Moved here
// from sim/cmd/leveling-bis so the talent search stands its casters
// exactly where the ranker does.
var NoMeleeAutoAttackSpecs = map[string]bool{
	"mage-arcane":         true,
	"mage-fire":           true,
	"mage-frost":          true,
	"warlock-affliction":  true,
	"warlock-demonology":  true,
	"warlock-destruction": true,
	"priest-shadow":       true,
	"shaman-elemental":    true,
	"druid-balance":       true,
	// A healer heals; it never swings in melee.
	"druid-restoration":  true,
	"paladin-holy":       true,
	"priest-discipline":  true,
	"priest-holy":        true,
	"shaman-restoration": true,
}

// CasterDistanceFromTarget clears MinRangedAttackDistance (12 yards,
// sim/core/constants.go) so a wand/Shoot-using character's own ranged
// auto-attack still fires, while also clearing MaxMeleeAttackDistance
// (5) so the melee auto-swing never starts - the same distance
// sim/core/test_generators.go's ranged-preset fixtures use.
const CasterDistanceFromTarget = 30
