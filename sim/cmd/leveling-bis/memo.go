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

// memoEntry is one request's answer; ready closes once it is in, so a
// caller asking for a sim another goroutine is already running waits for
// that run instead of starting its own.
type memoEntry[V any] struct {
	ready chan struct{}
	value V
}

// memoRunner answers repeated plain, weights and Shoot sims from memory
// and delegates everything else to the runner it wraps. It is safe for
// concurrent use.
type memoRunner struct {
	inner engineRunner

	mu      sync.Mutex
	plain   map[string]*memoEntry[memoPlain]
	weights map[string]*memoEntry[memoWeights]
	shoot   map[string]*memoEntry[memoPlain]
	hits    int
	misses  int
}

func newMemoRunner(inner engineRunner) *memoRunner {
	return &memoRunner{
		inner:   inner,
		plain:   map[string]*memoEntry[memoPlain]{},
		weights: map[string]*memoEntry[memoWeights]{},
		shoot:   map[string]*memoEntry[memoPlain]{},
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

// memoize returns cache's answer for key, running run once to get it.
func memoize[V any](m *memoRunner, cache map[string]*memoEntry[V], key string, run func() V) V {
	m.mu.Lock()
	if e, ok := cache[key]; ok {
		m.hits++
		m.mu.Unlock()
		<-e.ready
		return e.value
	}
	e := &memoEntry[V]{ready: make(chan struct{})}
	cache[key] = e
	m.misses++
	m.mu.Unlock()
	e.value = run()
	close(e.ready)
	return e.value
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
	v := memoize(m, m.plain, key, func() memoPlain {
		mean, stdErr, err := m.inner.RunPlainDPSWithError(req)
		return memoPlain{mean: mean, stdErr: stdErr, err: err}
	})
	return v.mean, v.stdErr, v.err
}

func (m *memoRunner) RunWeights(req api.SimRequest) (map[string]api.StatWeight, float64, error) {
	key, err := memoKey(req)
	if err != nil {
		return nil, 0, err
	}
	if key == "" {
		return m.inner.RunWeights(req)
	}
	v := memoize(m, m.weights, key, func() memoWeights {
		result, reference, err := m.inner.RunWeights(req)
		return memoWeights{result: result, reference: reference, err: err}
	})
	return v.result, v.reference, v.err
}

func (m *memoRunner) RunShootDPS(req api.SimRequest) (float64, error) {
	key, err := memoKey(req)
	if err != nil {
		return 0, err
	}
	if key == "" {
		return m.inner.RunShootDPS(req)
	}
	v := memoize(m, m.shoot, key, func() memoPlain {
		dps, err := m.inner.RunShootDPS(req)
		return memoPlain{mean: dps, err: err}
	})
	return v.mean, v.err
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
