package rating

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// specSlug turns a roster row's display class/spec ("Warrior", "Protection") into the
// slug logs/engine/mechanics/utility and data/curated/specs.json use ("warrior-
// protection"): lowercase, spaces to hyphens. This is the same trivial transform
// logs/engine/rating's own unexported specSlug applies (rating.go) - duplicated here
// because that function is unexported and this lane may not export it (logs/ is
// read-only for this lane); it exists only to pick which curated file to load, never to
// influence scoring math itself, which the engine computes independently.
func specSlug(class, spec string) string {
	return toSlug(class) + "-" + toSlug(spec)
}

func toSlug(s string) string {
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == ' ':
			b = append(b, '-')
		case c >= 'A' && c <= 'Z':
			b = append(b, c+('a'-'A'))
		default:
			b = append(b, c)
		}
	}
	return string(b)
}

// curatedTablesFor selects the curated data one roster row's rating needs: the fight's
// own encounter mechanics table (nil if none curated yet - Score handles that by
// excluding Survival/Mechanics, §1.1), this player's spec's utility table, and this
// player's role's consumable catalogue.
func curatedTablesFor(row summary.RosterRow, encounterID int64) ratingengine.CuratedTables {
	var out ratingengine.CuratedTables
	if m, ok := mechanics.Load(encounterID); ok {
		out.Mechanics = &m
	}
	if u, ok := utility.Load(specSlug(row.Class, row.Spec)); ok {
		out.Utility = &u
	}
	if rc, ok := consumables.Load().For(row.Role); ok {
		out.Consumables = &rc
	}
	return out
}
