// api/internal/rating/viewer.go
package rating

import (
	"context"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// Accounts is the one method this package needs from auth.Store: whether userID is a
// verified member of guildID, and at what rank. Declared here, where it is consumed,
// rather than imported from reports.Accounts (which declares a second method this
// package never calls) - *auth.Store already satisfies both.
type Accounts interface {
	GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error)
}

// visible replicates reports.Service.mayView's exact rule (reports/handler.go:540-556)
// for a report rep: public and unlisted are visible to anyone; private is visible to the
// owner and moderators; guild is visible to the owner, moderators, and verified guild
// members. It cannot call mayView directly - that method is unexported on *reports.Service
// and this lane's file fence forbids editing reports/handler.go to export it (see this
// plan's Global Constraints and the ledger for the ruling this documents) - so it is
// rebuilt here from the same exported primitives mayView itself is built from
// (reports.Public/Unlisted/Private/GuildTo, auth.ActorFrom, Actor.Signed/IsModerator,
// Accounts.GuildRank). Any future change to mayView's rule must be mirrored here by hand;
// there is no way around that within this lane's constraints.
func visible(r *http.Request, rep reports.Report, accounts Accounts) bool {
	if rep.Visibility == reports.Public || rep.Visibility == reports.Unlisted {
		return true
	}
	a := auth.ActorFrom(r.Context())
	if !a.Signed() {
		return false
	}
	if a.IsModerator() || (rep.OwnerID != nil && *rep.OwnerID == a.UserID) {
		return true
	}
	if rep.Visibility == reports.GuildTo && rep.GuildID != nil && accounts != nil {
		if _, ok, err := accounts.GuildRank(r.Context(), *rep.GuildID, a.UserID); err == nil && ok {
			return true
		}
	}
	return false
}
