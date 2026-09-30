// Package simdb carries the active build's item database into both
// artifacts.
//
// Forever re-itemises: the engine's own --tags=with_db table is
// vanilla's, and most of a vanilla item id resolves to nothing here. The
// data lane emits the real one as the engine's own SimDatabase protobuf
// at data/builds/<build>/simdb.bin, and `make simdb` copies the build
// named by web/src/data/active-build.json to simdb.bin beside this file,
// where it is embedded. The copy is generated and git-ignored; the
// source is committed.
//
// Neither artifact is built --tags=with_db. That tag would load vanilla's
// table and, worse, turn on core.WITH_DB, which makes the engine panic at
// init for any item effect whose id is not in the table - and the effects
// register before anything of ours could load a database.
//
// The database reaches the engine the way the engine's own web UI sends
// it: on proto.Player.Database, which core.NewCharacter folds into the
// global tables (core.addToDatabase is unexported, so this is the
// exported path, not a shortcut around one).
package simdb

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

//go:embed simdb.bin
var raw []byte

var load = sync.OnceValues(func() (*proto.SimDatabase, error) {
	db := &proto.SimDatabase{}
	if err := googleproto.Unmarshal(raw, db); err != nil {
		return nil, fmt.Errorf("simdb: the embedded item database is corrupt: %w", err)
	}
	if len(db.Items) == 0 {
		return nil, fmt.Errorf("simdb: the embedded item database has no items")
	}
	return db, nil
})

// itemIDs is the set of ids the embedded database carries, built once
// beside it: what UnequipUnknown checks a worn item against.
var itemIDs = sync.OnceValues(func() (map[int32]struct{}, error) {
	db, err := load()
	if err != nil {
		return nil, err
	}
	ids := make(map[int32]struct{}, len(db.Items))
	for _, item := range db.Items {
		ids[item.Id] = struct{}{}
	}
	return ids, nil
})

// Known reports whether id carries a row in the embedded item database
// -- the same table UnequipUnknown checks a worn item against. A
// ranker that equips id and trusts the engine to model its effect
// (effectids_generated.go's own claim) has to ask this first: an id
// this returns false for is silently stripped from every character's
// equipment by UnequipUnknown before the engine ever builds it, so its
// effect (registered in the engine by item id, sim/core/item_effects.go)
// never has a chance to apply -- a real sim of it still runs and still
// reports a real number, but that number is the set's DPS with the
// item simply not worn, not a measurement of its effect. A caller that
// ignores this and reads a resulting zero DPS delta as "this effect is
// real and worth nothing" publishes exactly the false tie this lane's
// brief calls the tenet-8 violation (Hand of Justice 11815, missing
// from this build's simdb.bin because its ItemSparse row is absent
// from the client export -- see data/pipeline/simdb/items.py's own
// "present in both tables" filter -- while its Item.csv row and
// items.json's own classic-db-fallback entry both still exist).
// Returns false (not an error) on a load failure: the caller already
// has its own path for a load failure (Attach/AttachWeights return an
// error first), so by the time Known is asked, id is either in the
// loaded table or it is not.
func Known(id int32) bool {
	ids, err := itemIDs()
	if err != nil {
		return false
	}
	_, ok := ids[id]
	return ok
}

// Attach puts the database on every player of a built request, which is
// what makes an item id resolve. It is called once per request rather
// than once per process because the engine offers no other way in.
//
// It also empties any worn slot whose item the database does not carry
// (UnequipUnknown): a character can wear a quality-1 keepsake ring or a
// totem the sim never scores, and the engine panics on an id it cannot
// find rather than skipping it. The site's strip already names such an
// item "not simmed", so the honest run is the one without it.
func Attach(req *proto.RaidSimRequest) error {
	db, err := load()
	if err != nil {
		return err
	}
	if req == nil || req.Raid == nil {
		return nil
	}
	for _, party := range req.Raid.Parties {
		for _, player := range party.Players {
			if player != nil {
				player.Database = db
				if _, err := UnequipUnknown(player); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// UnequipUnknown empties every equipment slot of the player whose item id
// the embedded database lacks and reports the ids it removed, in slot
// order. A player with no equipment, or every item known, is left alone.
func UnequipUnknown(player *proto.Player) ([]int32, error) {
	if player == nil || player.Equipment == nil {
		return nil, nil
	}
	known, err := itemIDs()
	if err != nil {
		return nil, err
	}
	var removed []int32
	for i, spec := range player.Equipment.Items {
		if spec == nil || spec.Id == 0 {
			continue
		}
		if _, ok := known[spec.Id]; !ok {
			removed = append(removed, spec.Id)
			player.Equipment.Items[i] = &proto.ItemSpec{}
		}
	}
	return removed, nil
}

// AttachWeights puts the database on a stat weights request's player,
// which is the same door as Attach's: proto.Player.Database, folded in
// by core.NewCharacter. A weights run equips the character the same
// way a DPS run does, so without it every item id resolves to nothing.
func AttachWeights(req *proto.StatWeightsRequest) error {
	db, err := load()
	if err != nil {
		return err
	}
	if req != nil && req.Player != nil {
		if _, err := UnequipUnknown(req.Player); err != nil {
			return err
		}
	}
	if req == nil || req.Player == nil {
		return nil
	}
	req.Player.Database = db
	return nil
}
