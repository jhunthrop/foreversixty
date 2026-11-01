// api/internal/guilds/roster_spec.go
//
// The spec words the guild home's roster carries for the apps (Wave 3 W3-C16): each roster[]
// row and members[] entry gains role, the full spec name ("Holy Paladin", never bare "Holy")
// and the spec's icon stem, all from the spec catalogue (raidspec), so the apps never derive a
// role or a name from a spec word themselves. They are read only from the spec and class the
// row already shows, so a row whose consent withholds its spec withholds these too.
package guilds

import (
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// specFacts is the catalogue spec a roster row resolves to.
type specFacts struct {
	Role, Name, Icon string
}

// fullSpecName is the spec as a player says it: the catalogue name, then the class ("Holy
// Paladin", "Feral Bear Druid").
func fullSpecName(s raidspec.Spec) string {
	class := s.Class
	if class != "" {
		class = strings.ToUpper(class[:1]) + class[1:]
	}
	return strings.TrimSpace(s.Name + " " + class)
}

// resolveSpec finds the catalogue spec of a row's spec word: a specs.json key ("paladin-holy",
// the addon snapshot's form) first, else a spec name within the class ("Holy", the logs' form),
// the role deciding a name two specs share. ok is false when nothing resolves.
func resolveSpec(class, spec *string, role string) (specFacts, bool) {
	if spec == nil || strings.TrimSpace(*spec) == "" {
		return specFacts{}, false
	}
	if s, ok := raidspec.ByKey(strings.TrimSpace(*spec)); ok {
		return specFacts{Role: s.Role, Name: fullSpecName(s), Icon: s.Icon}, true
	}
	if class == nil {
		return specFacts{}, false
	}
	s, ok := raidspec.Resolve(*class, *spec, role)
	if !ok {
		return specFacts{}, false
	}
	return specFacts{Role: s.Role, Name: fullSpecName(s), Icon: s.Icon}, true
}

// describeRosterRows fills each roster row's spec_name and spec_icon, and its role when the
// logs gave none. A row that does not resolve keeps what it had.
func describeRosterRows(rows []RosterRow) {
	for i := range rows {
		role := ""
		if rows[i].Role != nil {
			role = *rows[i].Role
		}
		facts, ok := resolveSpec(rows[i].Class, rows[i].Spec, role)
		if !ok {
			continue
		}
		rows[i].SpecName, rows[i].SpecIcon = strp(facts.Name), strp(facts.Icon)
		if rows[i].Role == nil {
			rows[i].Role = strp(facts.Role)
		}
	}
}

// describeMembers fills each members[] entry's role, spec_name and spec_icon: from the entry's
// own spec (the addon snapshot's), else from its roster row's (the logs'), matched by
// character key. An entry with neither keeps all three null.
func describeMembers(entries []RosterEntry, rows []RosterRow) {
	byKey := make(map[string]RosterRow, len(rows))
	for _, row := range rows {
		byKey[row.CharacterKey] = row
	}
	for i := range entries {
		e := &entries[i]
		row, hasRow := byKey[e.CharacterKey]
		role := ""
		if hasRow && row.Role != nil {
			role = *row.Role
		}
		facts, ok := resolveSpec(e.Class, e.Spec, role)
		if !ok && hasRow && e.CharacterKey != "" {
			facts, ok = resolveSpec(row.Class, row.Spec, role)
		}
		if !ok {
			continue
		}
		e.Role, e.SpecName, e.SpecIcon = strp(facts.Role), strp(facts.Name), strp(facts.Icon)
	}
}

func strp(s string) *string { return &s }
