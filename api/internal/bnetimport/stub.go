package bnetimport

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// Stub is the importer wired while BNET_IMPORT_CHARACTERS is off (config.go): Blizzard
// serves no Forever namespace yet, only Era (classic1x) and progression realms, so an
// account import filled the site with Season of Discovery and Era characters that are not
// Forever characters at all. It writes nothing, never stamps users.bnet_imported_at, and
// reports every region unavailable; sign-in itself still works. The day the namespace
// appears (the nightly probe's namespace_appeared warning), the flag turns the real
// Service back on with no code change.
type Stub struct {
	Regions []string
}

// ImportAccount satisfies auth.Importer without touching Battle.net or the database.
func (s *Stub) ImportAccount(context.Context, int64, string) (auth.ImportSummary, error) {
	return auth.ImportSummary{Regions: s.Regions, Unavailable: len(s.Regions)}, nil
}
