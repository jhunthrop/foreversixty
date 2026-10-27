package guilds

import (
	"context"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Supporters answers which accounts show the supporter mark, in one ask; nil means
// nobody does (the fixture build, and tests that do not care).
type Supporters interface {
	SupportersAmong(ctx context.Context, userIDs []int64) (map[int64]bool, error)
}

// markSupporters sets Supporter on every row of the given lists from its owning account,
// asking once per roster however many accounts and characters it has. A cosmetic flag never
// takes the home down: no service, or a failed ask, leaves every row false (the failure
// logged at warn).
func (s *Service) markSupporters(ctx context.Context, lists ...[]RosterRow) {
	if s.Supporters == nil {
		return
	}
	owners := distinctOwners(lists)
	if len(owners) == 0 {
		return
	}
	answers, err := s.Supporters.SupportersAmong(ctx, owners)
	if err != nil {
		s.logger().Warn("guilds", "id", httpx.RequestIDFrom(ctx), "op", "supporter", "err", err)
		return
	}
	for _, rows := range lists {
		for i := range rows {
			rows[i].Supporter = answers[rows[i].ownerUserID]
		}
	}
}

// distinctOwners lists each owning account of the rows once, skipping unowned rows.
func distinctOwners(lists [][]RosterRow) []int64 {
	seen := map[int64]bool{}
	var owners []int64
	for _, rows := range lists {
		for i := range rows {
			if owner := rows[i].ownerUserID; owner != 0 && !seen[owner] {
				seen[owner] = true
				owners = append(owners, owner)
			}
		}
	}
	return owners
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
