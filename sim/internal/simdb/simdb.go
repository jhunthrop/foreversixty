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

// Attach puts the database on every player of a built request, which is
// what makes an item id resolve. It is called once per request rather
// than once per process because the engine offers no other way in.
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
			}
		}
	}
	return nil
}
