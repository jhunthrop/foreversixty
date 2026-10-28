package coach

import ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"

// percentOf scales a 0-1 share to percent; the F3, F7 and F8 checks share it.
const percentOf = 100.0

// countedBetween is the length of [from, to) outside the excluded windows.
func countedBetween(excluded []ratingengine.Window, from, to int64) int64 {
	if to <= from {
		return 0
	}
	total := to - from
	for _, w := range excluded {
		total -= max(min(w.ToMS, to)-max(w.FromMS, from), 0)
	}
	return total
}
