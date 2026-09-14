//go:build race

package bench

// raceEnabled reports that this binary was built with the race detector, whose
// instrumentation slows the parser several-fold. The budget test measures the
// production build, which the workflow's bench job runs without -race.
const raceEnabled = true
