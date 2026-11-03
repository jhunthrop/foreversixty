package main

import (
	"fmt"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// legalRequest is buildRequest held to the game's talent rules: a
// -talents string (or a published entry) the game would refuse, such as
// a row taken before its gate or more points than the band's level
// gives, stops the run with the talent and rule named, rather than
// measuring a build no player can hold.
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
	if err := layouts.Check(c.Class, c.Talents, c.Level); err != nil {
		return api.SimRequest{}, fmt.Errorf("-talents %q: %w", c.Talents, err)
	}
	return req, nil
}
