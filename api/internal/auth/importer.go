// api/internal/auth/importer.go
package auth

import (
	"context"
	"time"
)

// ImportSummary is what one Battle.net import run produced (spec §4.1),
// logged by bnetCallback at INFO (success) or WARN (error). Defined here
// rather than in package bnetimport so this package need not import it:
// bnetimport already depends on auth transitively (through guilds, which
// reads auth.User), so the dependency runs one way only — auth defines
// the shape, bnetimport (and any other importer) returns it.
type ImportSummary struct {
	// Regions is which regions answered 200 for /profile/user/wow.
	Regions []string
	// Characters is how many characters rows were written.
	Characters int
	// Guilds is how many guild memberships were written.
	Guilds int
	// Skipped is collisions and level-too-low skips, combined.
	Skipped int
	// Unavailable is how many characters answered 404 for their public
	// profile (spec A2 — e.g. a Season of Discovery character under the
	// classic1x namespace). The characters row is still written; only
	// its guild membership and profile capture are unavailable.
	Unavailable int
}

// Importer runs the Battle.net character/guild import after a successful
// sign-in (spec §4.1). A Service with a nil Importer behaves exactly as
// login did before this feature existed — every existing auth test sees
// this. bnetimport.Service satisfies this interface.
type Importer interface {
	ImportAccount(ctx context.Context, userID int64, userToken string) (ImportSummary, error)
}

// importBudget bounds how long bnetCallback waits for the import before
// redirecting anyway (spec §4.1: "with a 5 second budget").
const importBudget = 5 * time.Second
