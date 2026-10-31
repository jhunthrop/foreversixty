package upgrades

import "sort"

// Group is where an upgrade sorts: empty slots, then measured gains, then picks with no
// figure or whose effect the sim leaves out.
type Group int

const (
	GroupEmpty Group = iota
	GroupMeasured
	GroupUnmodelled
)

// Order is items in the bot's order: empty slots in frame order, then measured gains
// largest first, then the rest in frame order. items arrive in frame order; key reads each
// item's group and gain. The input is not modified.
func Order[T any](items []T, key func(T) (Group, float64)) []T {
	out := append([]T(nil), items...)
	sort.SliceStable(out, func(i, j int) bool {
		gi, vi := key(out[i])
		gj, vj := key(out[j])
		if gi != gj {
			return gi < gj
		}
		return gi == GroupMeasured && vi > vj
	})
	return out
}
