//go:build !js

package main

import "testing"

// sim/cmd/wasm is a js/wasm-only package: syscall/js does not build on a
// host platform, so there is nothing here to unit test. The four
// exports' behaviour is covered three ways instead:
//
//   - sim/combine's tests cover simSplit and simCombine, which are thin
//     wrappers over Split and Results;
//   - sim/cmd/forever-sim's tests cover the request -> engine -> adapter
//     pipeline that simRun runs, against the same code;
//   - the CI smoke test in .github/workflows/sim.yml instantiates the
//     built wasm under node and asserts all four exports exist and that
//     simRun returns a SimResult with a summary.
//
// This file exists so `go test ./...` does not report the package as
// untested without saying why.
func TestWasmIsCoveredElsewhere(t *testing.T) {
	t.Log("see the comment above: combine, forever-sim, and the CI smoke test")
}
