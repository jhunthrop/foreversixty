package seats

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
)

// maxFactRows bounds the seats read for one account: a seat held and a move
// offered in each guild is far more than anyone has.
const maxFactRows = 2 * catalog.SeatsRaidGuild

type factSeat struct {
	guildID   int64
	guildName string
	row       seatRow
}

// SeatFactsForUser reads the seat facts /v1/me carries for userID. A seat
// whose plan has lapsed, or whose slot the plan no longer has, shows nothing.
func (s *Store) SeatFactsForUser(ctx context.Context, userID int64, now time.Time) (auth.SeatFacts, error) {
	var facts auth.SeatFacts
	seats, err := s.factSeats(ctx, userID)
	if err != nil {
		return facts, err
	}
	tiers := map[int64]string{}
	for _, fs := range seats {
		standing, err := s.slotStands(ctx, tiers, fs.guildID, fs.row.slot)
		if err != nil {
			return facts, err
		}
		if !standing {
			continue
		}
		if err := s.addFact(ctx, &facts, userID, fs); err != nil {
			return facts, err
		}
	}
	facts.Ended, err = s.endedSeat(ctx, userID, now)
	return facts, err
}

func (s *Store) factSeats(ctx context.Context, userID int64) ([]factSeat, error) {
	rows, err := s.Pool.Query(ctx, `
		select s.guild_id, g.name, `+qualifiedSeatColumns+`
		from premium_seats s join guilds g on g.id = s.guild_id
		where s.holder_user_id = $1 or s.pending_user_id = $1
		order by s.guild_id, s.slot limit $2`, userID, maxFactRows)
	if err != nil {
		return nil, fmt.Errorf("seats: facts of %d: %w", userID, err)
	}
	out, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (factSeat, error) {
		var fs factSeat
		err := r.Scan(&fs.guildID, &fs.guildName, &fs.row.slot, &fs.row.state, &fs.row.holderID, &fs.row.givenBy,
			&fs.row.givenAt, &fs.row.offeredAt, &fs.row.startsAt, &fs.row.movableAt, &fs.row.gaveUpKey,
			&fs.row.returnedName, &fs.row.returnedAt, &fs.row.pendingID, &fs.row.pendingBy, &fs.row.pendingAt)
		return fs, err
	})
	if err != nil {
		return nil, fmt.Errorf("seats: facts of %d: %w", userID, err)
	}
	return out, nil
}

// slotStands is true when the guild's live whole plan has the slot; tiers
// caches each guild's tier for the one read.
func (s *Store) slotStands(ctx context.Context, tiers map[int64]string, guildID int64, slot int) (bool, error) {
	tier, cached := tiers[guildID]
	if !cached {
		var err error
		if tier, err = s.activeTier(ctx, guildID); err != nil {
			return false, err
		}
		tiers[guildID] = tier
	}
	return slot <= catalog.SeatsForTier(tier), nil
}

func (s *Store) addFact(ctx context.Context, facts *auth.SeatFacts, userID int64, fs factSeat) error {
	r := fs.row
	switch {
	case r.pendingID != nil && *r.pendingID == userID:
		from, err := s.giverName(ctx, fs.guildID, r.pendingBy)
		if err != nil || facts.Offer != nil || r.pendingAt == nil {
			return err
		}
		facts.Offer = &auth.SeatOffer{GuildID: fs.guildID, GuildName: fs.guildName, FromName: from, OfferedAt: *r.pendingAt}
	case r.state == StateOffered:
		from, err := s.giverName(ctx, fs.guildID, r.givenBy)
		if err != nil || facts.Offer != nil || r.offeredAt == nil {
			return err
		}
		facts.Offer = &auth.SeatOffer{GuildID: fs.guildID, GuildName: fs.guildName, FromName: from, OfferedAt: *r.offeredAt}
	case r.givenAt != nil:
		facts.HeldGuildIDs = append(facts.HeldGuildIDs, fs.guildID)
		if facts.Seat != nil {
			return nil
		}
		from, err := s.giverName(ctx, fs.guildID, r.givenBy)
		if err != nil {
			return err
		}
		facts.Seat = &auth.HeldSeat{GuildName: fs.guildName, GivenBy: from, Since: *r.givenAt, StartsAt: r.startsAt}
	}
	return nil
}

func (s *Store) giverName(ctx context.Context, guildID int64, giver *int64) (string, error) {
	if giver == nil {
		return "", nil
	}
	ch, err := characterOf(ctx, s.Pool, guildID, *giver)
	return ch.name, err
}

// endedSeat is the newest comeback still valid and unused.
func (s *Store) endedSeat(ctx context.Context, userID int64, now time.Time) (*auth.EndedSeat, error) {
	var e auth.EndedSeat
	err := s.Pool.QueryRow(ctx, `
		select guild_name, ended_at, id, price_lookup_key from seat_comebacks
		where user_id = $1 and used_at is null and expires_at > $2
		order by ended_at desc limit 1`, userID, now).
		Scan(&e.GuildName, &e.EndedAt, &e.ComebackID, &e.PriceLookupKey)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("seats: ended seat of %d: %w", userID, err)
	}
	return &e, nil
}
