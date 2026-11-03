package sims

import (
	"fmt"

	simapi "github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// talentsLegal refuses a server run whose character, or any talent
// loadout it compares, holds a build the game refuses (sim/talents
// Check: a rank past a talent's last, a row taken before its gate, a
// prerequisite unmet, more points than the character's level gives).
// The envelope checks only a talent string's shape. A nil layout is no
// check; cmd/api logs that at startup.
func talentsLegal(layouts *talents.Layouts, req simapi.SimRequest) error {
	if layouts == nil {
		return nil
	}
	c := req.Character
	if err := layouts.Check(c.Class, c.Talents, c.Level); err != nil {
		return fmt.Errorf("the character's talents: %w", err)
	}
	if req.Bulk == nil {
		return nil
	}
	for _, l := range req.Bulk.Talents {
		if err := layouts.Check(c.Class, l.Talents, c.Level); err != nil {
			return fmt.Errorf("talent loadout %q: %w", l.Name, err)
		}
	}
	return nil
}
