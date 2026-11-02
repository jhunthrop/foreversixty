package main

import (
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/wowsims/classic/sim/core"
)

// effectImplemented reports whether the engine this binary links registers
// an effect for the item: an item effect (core.NewItemEffect and every
// helper built on it, core.NewClientItemEffect included) or a weapon effect.
// It asks the engine rather than a list scraped from the fork's source,
// which went stale (round 9: seventeen picked items whose effects the
// engine runs read as unmodelled) and which a source scrape cannot keep
// complete when registrations go through helpers and loops.
func effectImplemented(itemID int) bool {
	inproc.Register()
	id := int32(itemID)
	return core.HasItemEffect(id) || core.HasWeaponEffect(id)
}
