package guilds

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Supporters answers whether an account shows the supporter mark; nil means
// nobody does (the fixture build, and tests that do not care).
type Supporters interface {
	IsSupporter(ctx context.Context, userID int64) (bool, error)
}

// markSupporters sets Supporter on every row of the given lists from its owning account,
// asking once per distinct account however many characters it has. A cosmetic flag never
// takes the home down: no service, or a failed ask, leaves that account's rows false (the
// failure logged at warn).
func (s *Service) markSupporters(ctx context.Context, lists ...[]RosterRow) {
	if s.Supporters == nil {
		return
	}
	answers := map[int64]bool{}
	for _, rows := range lists {
		for i := range rows {
			owner := rows[i].ownerUserID
			if owner == 0 {
				continue
			}
			supporter, asked := answers[owner]
			if !asked {
				ok, err := s.Supporters.IsSupporter(ctx, owner)
				if err != nil {
					s.logger().Warn("guilds", "id", httpx.RequestIDFrom(ctx), "op", "supporter", "user_id", owner, "err", err)
				}
				supporter = ok && err == nil
				answers[owner] = supporter
			}
			rows[i].Supporter = supporter
		}
	}
}

// Entitlements answers whether a user holds a per-guild feature; nil means nobody does.
type Entitlements interface {
	CanForGuild(ctx context.Context, guildID, userID int64, feature entitlements.Feature) (bool, entitlements.Reason, error)
}

// officerTools reports whether the officer views are unlocked for guildID. Like the
// supporter mark, a failed ask never takes the home down: it reads false and logs at warn.
func (s *Service) officerTools(ctx context.Context, guildID, userID int64) bool {
	if s.Entitlements == nil {
		return false
	}
	ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, userID, entitlements.FeatureOfficerViews)
	if err != nil {
		s.logger().Warn("guilds", "id", httpx.RequestIDFrom(ctx), "op", "officer_tools", "guild_id", guildID, "err", err)
		return false
	}
	return ok
}
