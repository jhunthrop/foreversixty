package request

// Cooldown timings.
//
// The engine already carries them: proto.Cooldowns is a list of
// (ActionID, timings), where an empty timing list means "on cooldown,
// whenever the rotation allows" and a value is one usage at a fixed
// second. All this file does is name an action the way the settings bar
// names things - "spell:1719", "item:13452", or a consumable id from
// IDS.md - and refuse one it cannot resolve, because a major cooldown
// quietly dropped is a DPS number that is wrong and says nothing.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// ErrUnknownCooldown is returned for a cooldown id this module cannot
// turn into an engine action.
var ErrUnknownCooldown = errors.New("request: unknown cooldown")

// spellPrefix marks a cooldown id that names an engine spell rather
// than a client item or consumable: "spell:1719". itemPrefix is
// consumables.go's "item".
const spellPrefix = "spell"

// cooldownAction turns one cooldown id into the engine's ActionID.
func cooldownAction(id string, table *Consumables) (*proto.ActionID, error) {
	head, tail, qualified := strings.Cut(id, ":")
	switch {
	case qualified && head == spellPrefix:
		n, err := strconv.ParseInt(tail, 10, 32)
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("%w: %q names no spell", ErrUnknownCooldown, id)
		}
		return &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: int32(n)}}, nil
	case qualified && head == itemPrefix:
		n, err := strconv.ParseInt(tail, 10, 32)
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("%w: %q names no item", ErrUnknownCooldown, id)
		}
		return &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: int32(n)}}, nil
	}
	// A bare id is a consumable's, which only the build's table can
	// turn into an item. Saying the table is missing beats resolving
	// the id to nothing.
	if table == nil {
		return nil, fmt.Errorf("%w: %q is a consumable and no consumable table is loaded; pass one in Options, from data/builds/<build>/simconsumes.json", ErrUnknownCooldown, id)
	}
	itemID, ok := table.ItemID(id)
	if !ok {
		return nil, fmt.Errorf("%w: %q is not a spell, an item, or a consumable this build's table names exactly once", ErrUnknownCooldown, id)
	}
	return &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: int32(itemID)}}, nil
}

// cooldownsFor builds the engine's Cooldowns message, or nil when the
// character pinned none - which is the engine's own default and must
// stay distinguishable from "a message with no entries".
func cooldownsFor(specs []api.CooldownSpec, table *Consumables) (*proto.Cooldowns, error) {
	if len(specs) == 0 {
		return nil, nil
	}
	out := &proto.Cooldowns{Cooldowns: make([]*proto.Cooldown, 0, len(specs))}
	for _, cd := range specs {
		action, err := cooldownAction(cd.ID, table)
		if err != nil {
			return nil, err
		}
		out.Cooldowns = append(out.Cooldowns, &proto.Cooldown{Id: action, Timings: cd.AtSec})
	}
	return out, nil
}
