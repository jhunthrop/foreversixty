package upgrades

import "strings"

// UnmatchedPoints is how many of the band build's points the worn talents do not match
// (web/src/lib/home/talent-delta.ts): Σ max(band − worn, 0) over every talent cell. worn is
// each tree's rank string; band the three trees joined with "-". A missing or unreadable digit
// is 0.
func UnmatchedPoints(worn []string, band string) int {
	bandTrees := strings.Split(band, "-")
	unmatched := 0
	for tree := range max(len(worn), len(bandTrees)) {
		w, b := at(worn, tree), at(bandTrees, tree)
		for cell := range max(len(w), len(b)) {
			unmatched += max(digit(b, cell)-digit(w, cell), 0)
		}
	}
	return unmatched
}

func at(list []string, i int) string {
	if i < len(list) {
		return list[i]
	}
	return ""
}

func digit(s string, i int) int {
	if i >= len(s) || s[i] < '0' || s[i] > '9' {
		return 0
	}
	return int(s[i] - '0')
}
