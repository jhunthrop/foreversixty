package main

// Concurrent ranking.
//
// A band ranks under each of its weight sets (band_pipeline.go), and each
// set ranks both factions up to the trinket reconcile and again after it.
// Those runs share only read-only inputs (items, loot index, weights), so
// they can run at once on a runner that is safe for concurrent use: the
// real engine, which the site's API already runs concurrently, and the
// tank, healer and memo runners built on it. A test double that records
// its calls in order is not, so it keeps the sequential order.

import "sync"

// concurrentSafe is what a runner safe for concurrent use implements.
type concurrentSafe interface{ concurrentSafe() bool }

// runsConcurrently reports whether r may be called from several goroutines.
func runsConcurrently(r engineRunner) bool {
	c, ok := r.(concurrentSafe)
	return ok && c.concurrentSafe()
}

func (realEngine) concurrentSafe() bool    { return true }
func (t tankEngine) concurrentSafe() bool  { return runsConcurrently(t.engineRunner) }
func (m *memoRunner) concurrentSafe() bool { return runsConcurrently(m.inner) }
func (h healEngine) concurrentSafe() bool {
	_, real := h.backend.(realHealingBackend)
	return real
}

// forEach runs do(i) for i in [0, n): at once when parallel, else in order.
// It returns the first error by index, so the error a run reports does not
// depend on scheduling.
func forEach(n int, parallel bool, do func(i int) error) error {
	errs := make([]error, n)
	if parallel {
		var wg sync.WaitGroup
		for i := range n {
			wg.Add(1)
			go func() {
				defer wg.Done()
				errs[i] = do(i)
			}()
		}
		wg.Wait()
	} else {
		for i := range n {
			if errs[i] = do(i); errs[i] != nil {
				break
			}
		}
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
