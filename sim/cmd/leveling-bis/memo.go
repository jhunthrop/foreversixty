package main

// The sim memo.
//
// Every sim the ranker runs is deterministic: each request names its own
// seed and iteration count, and the engine draws from seeded, labelled
// random streams. So two identical requests give the same answer, and the
// second one need not be simulated. Round 8 ranks a band under two weight
// sets start to finish (band_pipeline.go); wherever both sets reach the same
// candidate sets (the same first-pass picks, trinket tournaments, verify
// swaps and set measurements) the second pass's sims were exact repeats.
// The memo answers those from the first pass, and the guard, chooseRun and
// the final measurements reuse what the passes already measured.
//
// The key is the request with the character's name cleared: the name is a
// log label (bandCharacter's "verify", "trinket-rank", ...) the engine
// reads nowhere in a sim.

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core"
)

type memoPlain struct {
	mean, stdErr float64
	err          error
}

type memoWeights struct {
	result    map[string]api.StatWeight
	reference float64
	err       error
}

// memoRunner answers repeated plain and weights sims from memory and
// delegates everything else to the runner it wraps.
type memoRunner struct {
	inner engineRunner

	mu      sync.Mutex
	plain   map[string]memoPlain
	weights map[string]memoWeights
	shoot   map[string]memoPlain
	hits    int
	misses  int
}

func newMemoRunner(inner engineRunner) *memoRunner {
	return &memoRunner{
		inner:   inner,
		plain:   map[string]memoPlain{},
		weights: map[string]memoWeights{},
		shoot:   map[string]memoPlain{},
	}
}

// stats is the memo's hit and miss counts so far.
func (m *memoRunner) stats() (hits, misses int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.hits, m.misses
}

// memoKey is req's identity for the memo, or "" when req must always be
// simulated: a request with no seed draws a fresh one per run.
func memoKey(req api.SimRequest) (string, error) {
	if req.RandomSeed == 0 {
		return "", nil
	}
	req.Character.Name = ""
	b, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("keying a sim request: %w", err)
	}
	return string(b), nil
}

// lookup returns the cached value for key, counting a hit or a miss.
func lookup[V any](m *memoRunner, cache map[string]V, key string) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := cache[key]
	if ok {
		m.hits++
	} else {
		m.misses++
	}
	return v, ok
}

func store[V any](m *memoRunner, cache map[string]V, key string, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cache[key] = v
}

func (m *memoRunner) RunPlainDPS(req api.SimRequest) (float64, error) {
	mean, _, err := m.RunPlainDPSWithError(req)
	return mean, err
}

func (m *memoRunner) RunPlainDPSWithError(req api.SimRequest) (float64, float64, error) {
	key, err := memoKey(req)
	if err != nil {
		return 0, 0, err
	}
	if key == "" {
		return m.inner.RunPlainDPSWithError(req)
	}
	if v, ok := lookup(m, m.plain, key); ok {
		return v.mean, v.stdErr, v.err
	}
	mean, stdErr, err := m.inner.RunPlainDPSWithError(req)
	store(m, m.plain, key, memoPlain{mean: mean, stdErr: stdErr, err: err})
	return mean, stdErr, err
}

func (m *memoRunner) RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	key, err := memoKey(req)
	if err != nil {
		return nil, 0, err
	}
	if key == "" {
		return m.inner.RunWeights(req)
	}
	if v, ok := lookup(m, m.weights, key); ok {
		return v.result, v.reference, v.err
	}
	result, reference, err := m.inner.RunWeights(req)
	store(m, m.weights, key, memoWeights{result: result, reference: reference, err: err})
	return result, reference, err
}

func (m *memoRunner) RunShootDPS(req api.SimRequest) (float64, error) {
	key, err := memoKey(req)
	if err != nil {
		return 0, err
	}
	if key == "" {
		return m.inner.RunShootDPS(req)
	}
	if v, ok := lookup(m, m.shoot, key); ok {
		return v.mean, v.err
	}
	dps, err := m.inner.RunShootDPS(req)
	store(m, m.shoot, key, memoPlain{mean: dps, err: err})
	return dps, err
}

func (m *memoRunner) HitProfileFor(req api.SimRequest) (core.HitProfile, error) {
	return m.inner.HitProfileFor(req)
}

// unwrapRunner is the runner under any memo, where a role-only interface
// lives.
func unwrapRunner(runner engineRunner) engineRunner {
	if m, ok := runner.(*memoRunner); ok {
		return m.inner
	}
	return runner
}
