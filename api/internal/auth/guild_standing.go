// api/internal/auth/guild_standing.go
//
// The guild facts GET /v1/me adds to each guilds[] entry beyond membership itself (spec
// design/specs/2026-10-27-subscribe-and-join.md §8): identity, claim, verify, covered and
// subscribe. guilds.Store answers them in one read; this package only names the shapes and
// copies them onto the entries.
package auth

import "context"

// GuildClaimView is a guild's claim state as the viewer needs it. HolderName is the claim
// holder's or pending claimant's character in the guild, sent only to a verified officer.
type GuildClaimView struct {
	State        string  `json:"state"` // "unclaimed" | "pending" | "claimed" | "contested"
	Yours        bool    `json:"yours"`
	PendingYours bool    `json:"pending_yours"`
	HolderName   *string `json:"holder_name"`
}

// GuildVerifyView says whether anyone in the guild can approve a character.
type GuildVerifyView struct {
	HasApprover bool `json:"has_approver"` // at least one verified officer or leader
}

// GuildCoveredView is what covers the guild today, for any account with a character in it.
type GuildCoveredView struct {
	Tier        string  `json:"tier"`  // "guild" | "raid_guild"
	State       string  `json:"state"` // "trial" | "active" | "ending" | "past_due" | "lapsed"
	EndsAt      *string `json:"ends_at"`
	BilledBy    string  `json:"billed_by"` // "" for a grant and for a viewer with no verified character
	YouPay      bool    `json:"you_pay"`
	YouHoldSeat bool    `json:"you_hold_seat"`
	SeatsGiven  int     `json:"seats_given"`
	SeatsTotal  int     `json:"seats_total"`
}

// GuildSubscribeView is the one step the guild card offers the viewer.
type GuildSubscribeView struct {
	Step           string  `json:"step"`
	TrialUsed      *bool   `json:"trial_used"`
	AvailableAt    *string `json:"available_at"`
	Reason         *string `json:"reason"`
	OtherGuildName *string `json:"other_guild_name"`
}

// GuildStanding is everything GuildStandingSource knows about one guild for one viewer.
type GuildStanding struct {
	Realm       string
	CrestURL    *string
	Faction     *string // "alliance" | "horde" | nil
	RosterCount *int    // counts.on_roster + counts.in_game_no_account; nil without a snapshot
	Claim       *GuildClaimView
	Verify      *GuildVerifyView
	Covered     *GuildCoveredView
	Subscribe   *GuildSubscribeView
}

// GuildStandingSource answers each guild's standing for a viewer in one read. It is an
// interface because guilds imports this package.
type GuildStandingSource interface {
	GuildStanding(ctx context.Context, userID int64, guildIDs []int64) (map[int64]GuildStanding, error)
}

// attachStanding copies each guild's standing onto its entry. A nil source or a failed read
// logs and returns gs unchanged: the older shape, which the web draws as identity-only cards.
func (s *Service) attachStanding(ctx context.Context, userID int64, gs []Guild) []Guild {
	if s.Standing == nil || len(gs) == 0 {
		return gs
	}
	ids := make([]int64, len(gs))
	for i, g := range gs {
		ids[i] = g.ID
	}
	standing, err := s.Standing.GuildStanding(ctx, userID, ids)
	if err != nil {
		s.logger().Error("auth: me: guild standing", "user_id", userID, "err", err)
		return gs
	}
	out := make([]Guild, len(gs))
	for i, g := range gs {
		if st, ok := standing[g.ID]; ok {
			g.Realm, g.CrestURL, g.Faction, g.RosterCount = st.Realm, st.CrestURL, st.Faction, st.RosterCount
			g.Claim, g.Verify, g.Covered, g.Subscribe = st.Claim, st.Verify, st.Covered, st.Subscribe
		}
		out[i] = g
	}
	return out
}
