package main

import "testing"

// registerEngine and nextRunID are the two simrun.go helpers that do
// not themselves reach into the engine's sim loop (registerEngine only
// registers agent factories - sync.Once guarded, safe to call
// repeatedly; nextRunID is pure string formatting), so they are safe
// to exercise directly without triggering the real, slow engine path
// this lane's brief forbids in a test.

func TestRegisterEngineIsIdempotent(t *testing.T) {
	// Calling this more than once (as every other test in this package
	// that reaches runSpec/verifyBand/rankTrinketSlot already does,
	// indirectly, through the real registerEngineOnce guard those
	// functions used to call before this lane's engineRunner refactor)
	// must not panic or double-register anything - sync.Once's whole
	// contract.
	registerEngine()
	registerEngine()
}

func TestNextRunIDIsUniquePerCall(t *testing.T) {
	a := nextRunID()
	b := nextRunID()
	if a == b {
		t.Fatalf("nextRunID returned the same id twice: %q", a)
	}
	if a == "" || b == "" {
		t.Fatal("nextRunID returned an empty id")
	}
}

// A compile-time check that realEngine actually satisfies engineRunner
// (this lane's whole reason for existing: runSpec/verifyBand/
// rankTrinketSlot take engineRunner, and main's run() constructs a
// realEngine{} as the only production implementation).
var _ engineRunner = realEngine{}
