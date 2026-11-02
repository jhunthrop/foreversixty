package main

import (
	"errors"
	"sync/atomic"
	"testing"
)

func TestOnlyTheRealEngineAndItsWrappersRankConcurrently(t *testing.T) {
	cases := []struct {
		name   string
		runner engineRunner
		want   bool
	}{
		{"real", realEngine{}, true},
		{"memo over tank over real", newMemoRunner(roleRunner(realEngine{}, specInfo{Role: roleTank})), true},
		{"healer on the real backend", healEngine{backend: realHealingBackend{}}, true},
		{"fake", &fakeEngine{}, false},
		{"memo over fake", newMemoRunner(&fakeEngine{}), false},
		{"tank over fake", roleRunner(&fakeEngine{}, specInfo{Role: roleTank}), false},
	}
	for _, c := range cases {
		if got := runsConcurrently(c.runner); got != c.want {
			t.Errorf("%s: runsConcurrently = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestForEachReportsTheFirstErrorByIndex(t *testing.T) {
	first, second := errors.New("first"), errors.New("second")
	for _, parallel := range []bool{false, true} {
		var ran atomic.Int64
		err := forEach(3, parallel, func(i int) error {
			ran.Add(1)
			switch i {
			case 1:
				return first
			case 2:
				return second
			}
			return nil
		})
		if !errors.Is(err, first) {
			t.Errorf("parallel %v: err = %v, want the first by index", parallel, err)
		}
		if !parallel && ran.Load() != 2 {
			t.Errorf("sequential: ran %d, want it to stop at the first error", ran.Load())
		}
	}
}
