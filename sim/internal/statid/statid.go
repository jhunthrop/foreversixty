// Package statid holds the one mapping between the stat weights
// vocabulary and the engine's Stat enum.
//
// It exists because two lanes need it: sim/request resolves a
// WeightsSpec's stat ids into the engine's StatWeightsRequest, and
// sim/adapter resolves them again to read the engine's
// StatWeightsResult back into the envelope's named rows. Putting the
// mapping in a leaf package under sim/internal, rather than in
// sim/request, is what lets sim/adapter depend on it without depending
// on sim/request - sim/request's own test suite links sim/adapter (the
// rotation smoke test runs a built request through it), and a
// sim/adapter that imported sim/request back would be a cycle.
package statid

import (
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
)

// spellings are the enum names strcase.Snake gets wrong.
//
// There is one. Snake inserts a separator before every capital, so
// MP5 comes out "m_p5"; the contract pins it as "mp5" (10.8) and so
// does every addon that reads a Pawn string. Listing the exception
// beats hand-writing all forty-one ids, which would stop matching the
// enum the first time the pin moved.
var spellings = map[string]string{"MP5": "mp5"}

// statID is one enum value's id: the value name with the Stat prefix
// stripped, in lower snake case, with the spellings above applied.
func statID(name string) string {
	bare := strings.TrimPrefix(name, "Stat")
	if fixed, ok := spellings[bare]; ok {
		return fixed
	}
	return strcase.Snake(bare)
}

// byID is the engine's Stat enum, keyed by the id the weights panel
// sends, so StatAttackPower is "attack_power". It is built from the
// enum's own name table rather than written down, so a stat the engine
// gains is offerable the day the pin moves.
var byID = func() map[string]proto.Stat {
	out := make(map[string]proto.Stat, len(proto.Stat_name))
	for value, name := range proto.Stat_name {
		out[statID(name)] = proto.Stat(value)
	}
	return out
}()

// Parse maps a stat id onto the engine's enum.
func Parse(id string) (proto.Stat, bool) {
	s, ok := byID[id]
	return s, ok
}

// Known lists every stat id a weights request may name, sorted.
func Known() []string {
	out := make([]string, 0, len(byID))
	for id := range byID {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}
