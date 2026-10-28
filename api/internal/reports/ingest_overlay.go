package reports

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// overlayDerivedRoster fills the roster fields the companion left blank: class,
// spec and item level from the API's own rebuild of the fight, and the
// character level from the site's character records (levels, by player GUID;
// the log never carries it). It never overrides a value the companion sent.
func overlayDerivedRoster(posted *summary.Summary, rebuilt summary.Summary, levels map[string]int64) {
	derivedAt := make(map[string]int, len(rebuilt.Roster))
	for j, row := range rebuilt.Roster {
		derivedAt[row.GUID] = j
	}
	for j := range posted.Roster {
		row := &posted.Roster[j]
		if row.Level == 0 {
			row.Level = levels[row.GUID]
		}
		k, ok := derivedAt[row.GUID]
		if !ok {
			continue
		}
		derived := rebuilt.Roster[k]
		if row.Class == "" {
			row.Class = derived.Class
		}
		if row.Spec == "" {
			row.Spec = derived.Spec
		}
		if row.ItemLevel == 0 {
			row.ItemLevel = derived.ItemLevel
		}
	}
}

// characterLevels looks up each roster player's character level, by GUID. A
// failed or absent lookup gives no levels: 0 reads "couldn't check" downstream,
// and the fight's storage never depends on it.
func (i *Ingest) characterLevels(ctx context.Context, rep Report, posted summary.Summary) map[string]int64 {
	if i.Levels == nil || len(posted.Roster) == 0 {
		return nil
	}
	region, ruleset := ReportRealm(rep)
	keys := make([]string, 0, len(posted.Roster))
	keyOf := make(map[string]string, len(posted.Roster))
	for _, row := range posted.Roster {
		key := character.KeyFromUnit(region, ruleset, row.Name)
		keys = append(keys, key)
		keyOf[row.GUID] = key
	}
	byKey, err := i.Levels.CharacterLevels(ctx, keys)
	if err != nil {
		i.logger().Error("ingest", "op", "levels", "report", rep.ID, "err", err)
		return nil
	}
	out := make(map[string]int64, len(byKey))
	for guid, key := range keyOf {
		if lvl := byKey[key]; lvl > 0 {
			out[guid] = lvl
		}
	}
	return out
}
