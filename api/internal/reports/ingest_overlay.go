package reports

import "github.com/jhunthrop/foreversixty/logs/engine/summary"

// overlayDerivedRoster fills the roster fields the companion left blank from
// the API's own rebuild of the fight: class, spec and level. It never
// overrides a value the companion did send.
func overlayDerivedRoster(posted *summary.Summary, rebuilt summary.Summary) {
	derivedAt := make(map[string]int, len(rebuilt.Roster))
	for j, row := range rebuilt.Roster {
		derivedAt[row.GUID] = j
	}
	for j := range posted.Roster {
		k, ok := derivedAt[posted.Roster[j].GUID]
		if !ok {
			continue
		}
		row, derived := &posted.Roster[j], rebuilt.Roster[k]
		if row.Class == "" {
			row.Class = derived.Class
		}
		if row.Spec == "" {
			row.Spec = derived.Spec
		}
		if row.Level == 0 {
			row.Level = derived.Level
		}
	}
}
