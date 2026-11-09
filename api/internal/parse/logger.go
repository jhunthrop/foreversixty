package parse

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// creditLogger names the log's writer as the report's logging character
// when the upload named none, before completion attributes the report's
// guild from it. A report that already names one keeps it.
func (d Deps) creditLogger(ctx context.Context, rep *reports.Report, s *session.Session) error {
	if rep.LoggingCharacter != nil {
		return nil
	}
	key, ok := loggerKey(s.Units().All())
	if !ok {
		return nil
	}
	if err := d.Reports.CreditLogger(ctx, rep.ID, key); err != nil {
		return err
	}
	rep.LoggingCharacter = &key
	return nil
}

// loggerKey is the character key of the player who wrote the log: the one
// player unit the log flags as MINE. A browser upload names no logging
// character, so this is how its report is credited. ok is false when the
// log flags no player, or more than one name, as its own: nothing is
// guessed. The unit's own ruleset and region ("First-Last-PvP-US") key it;
// a unit that names neither takes the default realm ReportRealm uses.
func loggerKey(all []units.Unit) (key string, ok bool) {
	name := ""
	for _, u := range all {
		if !u.IsPlayer() || u.Flags&units.FlagAffiliationMine == 0 || u.Name == "" {
			continue
		}
		if name != "" && name != u.Name {
			return "", false
		}
		name = u.Name
	}
	if name == "" {
		return "", false
	}
	key = character.KeyFromUnit("us", character.RulesetNormal, name)
	return key, character.ValidKey(key)
}
