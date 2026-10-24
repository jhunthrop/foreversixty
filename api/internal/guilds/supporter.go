package guilds

import "context"

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
					s.logger().Warn("guilds", "op", "supporter", "user_id", owner, "err", err)
				}
				supporter = ok && err == nil
				answers[owner] = supporter
			}
			rows[i].Supporter = supporter
		}
	}
}
