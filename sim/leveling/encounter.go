package leveling

import "github.com/jhunthrop/foreversixty/sim/api"

// RankerEncounter is the encounter every published number is simmed
// against: the default fight with no creature type. A raid boss is no
// slaying-gear target, so a "+45 Attack Power against Undead" bracer
// counts for nothing; the empty type would be the engine's Humanoid,
// which would credit it. The BiS ranker, the talent search,
// spec-breakdown and dps-per-point all fight this one target, so their
// baselines agree.
func RankerEncounter() api.EncounterSpec {
	e := api.DefaultEncounter()
	e.TargetType = api.TargetTypeUnknown
	return e
}
