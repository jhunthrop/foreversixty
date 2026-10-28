package coach

// The power ids the log's resource tracks and spellconst cost types use: the WoW client's
// power enum, as sim/adapter/adapter.go documents it (Focus, 2, is not a coach resource).
const (
	PowerMana   int64 = 0
	PowerRage   int64 = 1
	PowerEnergy int64 = 3
)

// PowerNames spells each coach resource the way player-facing copy does.
var PowerNames = map[int64]string{PowerMana: "mana", PowerRage: "rage", PowerEnergy: "energy"}
