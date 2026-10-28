package reports

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// GuildRanker is all MayView needs of the accounts store.
type GuildRanker interface {
	GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error)
}

// MayView is the one rule for who may see a report (§19): public and unlisted to anyone;
// private to the owner and moderators; guild to the owner, moderators and verified members.
func MayView(ctx context.Context, actor auth.Actor, rep Report, ranker GuildRanker) bool {
	if rep.Visibility == Public || rep.Visibility == Unlisted {
		return true
	}
	if !actor.Signed() {
		return false
	}
	if actor.IsModerator() || (rep.OwnerID != nil && *rep.OwnerID == actor.UserID) {
		return true
	}
	if rep.Visibility == GuildTo && rep.GuildID != nil && ranker != nil {
		_, ok, err := ranker.GuildRank(ctx, *rep.GuildID, actor.UserID)
		return err == nil && ok
	}
	return false
}
