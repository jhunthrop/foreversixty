package main

import (
	"log/slog"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/sim/talents"
)

// activeVersion is the active client build's version for the startup log, "" when none loaded.
func activeVersion(data *trees.Data) string {
	b, ok := data.Active()
	if !ok {
		return ""
	}
	return b.Version
}

// inferrer names specs from the active client build's talent data.
func inferrer(data *trees.Data) *spec.Inferrer {
	b, ok := data.Active()
	if !ok {
		return spec.New(nil)
	}
	return spec.New(b)
}

// talentLayouts loads the active build's talent layout, which is what
// turns a fight's recorded talent ids into the engine's positional
// string. A build with no layout is not fatal: the scorer then refuses
// every character, which is what it already does for the missing race.
func talentLayouts(dir string, data *trees.Data, log *slog.Logger) *talents.Layouts {
	build, ok := data.Active()
	if !ok {
		log.Warn("sims", "state", "no client build", "effect", "no execution scores")
		return nil
	}
	layouts, err := talents.Load(filepath.Join(dir, build.Version, "talents"))
	if err != nil {
		log.Warn("sims", "state", "no talent layout", "err", err, "effect", "no execution scores")
		return nil
	}
	return layouts
}
