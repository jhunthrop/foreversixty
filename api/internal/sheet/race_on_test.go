//go:build race

package sheet

// raceOn is set under -race, which slows rendering several times over; the timing budget is
// measured only without it.
const raceOn = true
