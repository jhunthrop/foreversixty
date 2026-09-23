// api/internal/bnetbuild/slots.go
package bnetbuild

// SLOTS is web/src/lib/planner/types.ts's SLOTS constant, copied so an FS1 gear line this
// package writes uses exactly the slot names the planner's decoder (fs1.ts) accepts. A test
// (slots_test.go) reads types.ts and refuses to let the two drift.
var SLOTS = [...]string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
	"finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
}
