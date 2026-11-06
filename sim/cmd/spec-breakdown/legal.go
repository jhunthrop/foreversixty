package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// errIllegalRace marks a -race the class cannot be in the faction.
var errIllegalRace = errors.New("illegal race")

// legalRequest is buildRequest held to the game's race and talent rules: a
// -race the class cannot be in the faction, or a -talents string (or a
// published entry) the game would refuse, such as a row taken before its
// gate or more points than the band's level gives, stops the run with the
// race or the talent and rule named, rather than measuring a character no
// player can make.
func legalRequest(o options) (api.SimRequest, error) {
	req, err := buildRequest(o)
	if err != nil {
		return api.SimRequest{}, err
	}
	dir, err := buildDirectory(o.repoRoot)
	if err != nil {
		return api.SimRequest{}, err
	}
	layouts, err := talents.Load(filepath.Join(dir, "talents"))
	if err != nil {
		return api.SimRequest{}, err
	}
	c := req.Character
	if err := checkRace(dir, c.Class, o.faction, c.Race); err != nil {
		return api.SimRequest{}, err
	}
	if err := layouts.Check(c.Class, c.Talents, c.Level); err != nil {
		return api.SimRequest{}, fmt.Errorf("-talents %q: %w", c.Talents, err)
	}
	return req, nil
}

// checkRace refuses a race the class cannot be in the faction the run
// measures (the published entry's gear and buffs are that faction's), from
// the client's race-class combinations the BiS race sweep reads.
func checkRace(buildDir, class, faction, race string) error {
	legal, err := request.LegalRaces(buildDir, class, faction)
	if err != nil {
		return err
	}
	if slices.Contains(legal, race) {
		return nil
	}
	return fmt.Errorf("-race %q: %w: a %s %s is one of %s", race, errIllegalRace, faction, class, strings.Join(legal, ", "))
}
