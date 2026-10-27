package seats

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/billingmail"
	"github.com/jhunthrop/foreversixty/api/internal/mail"
)

// clearSeatSQL empties a slot, keeping only what a returned seat shows.
const clearSeatSQL = `update premium_seats set state = 'empty', holder_user_id = null, given_by = null,
	given_at = null, offered_at = null, starts_at = null, movable_at = null,
	holder_gave_up_lookup_key = null, returned_name = $3, returned_at = $4
	where guild_id = $1 and slot = $2`

// release empties a held or accepted seat inside the guild lock. It records
// the last holder's name and the release time, and when the holder gave up
// their own Premium for the seat it opens one comeback and returns the
// seat-removed email for the caller to send after commit. An offered seat
// was never in use, so it only clears.
func (s *Service) release(ctx context.Context, tx pgx.Tx, guildID int64, row seatRow, now time.Time) (*mail.Message, error) {
	if row.holderID == nil {
		return nil, nil
	}
	if row.state == StateOffered {
		_, err := tx.Exec(ctx, clearSeatSQL, guildID, row.slot, nil, nil)
		return nil, wrap("clear offer", err)
	}
	ch, err := characterOf(ctx, tx, guildID, *row.holderID)
	if err != nil {
		return nil, err
	}
	if _, err := tx.Exec(ctx, clearSeatSQL, guildID, row.slot, ch.name, now); err != nil {
		return nil, wrap("release seat", err)
	}
	if row.gaveUpKey == nil {
		return nil, nil
	}
	return s.openComeback(ctx, tx, guildID, *row.holderID, *row.gaveUpKey, now)
}

// openComeback inserts the comeback for a released seat, once: ended_at is
// the seat's returned_at, so a repeat release hits the unique index and
// sends nothing.
func (s *Service) openComeback(ctx context.Context, tx pgx.Tx, guildID, userID int64, lookupKey string, endedAt time.Time) (*mail.Message, error) {
	price, ok := catalog.Lookup(lookupKey)
	if !ok {
		s.logger().Warn("seats", "op", "comeback", "guild", guildID, "user", userID, "price", lookupKey,
			"err", "price is not in the catalog, so there is nothing to come back to")
		return nil, nil
	}
	var guildName, email string
	if err := tx.QueryRow(ctx, `select g.name, u.email from guilds g, users u where g.id = $1 and u.id = $2`,
		guildID, userID).Scan(&guildName, &email); err != nil {
		return nil, wrap("comeback parties", err)
	}
	var id int64
	err := tx.QueryRow(ctx, `
		insert into seat_comebacks (user_id, guild_name, ended_at, price_lookup_key, expires_at)
		values ($1, $2, $3, $4, $5)
		on conflict (user_id, guild_name, ended_at) do nothing returning id`,
		userID, guildName, endedAt, lookupKey, endedAt.Add(catalog.ComebackValid)).Scan(&id)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, wrap("insert comeback", err)
	}
	msg, err := billingmail.SeatRemovedEmail(billingmail.SeatRemoved{
		To: email, Guild: guildName, EndedOn: endedAt, AmountCents: price.AmountCents,
		Interval: string(price.Interval), ComebackURL: s.comebackURL(price.Interval, id),
	})
	if err != nil {
		return nil, wrap("seat removed email", err)
	}
	return &msg, nil
}

func (s *Service) comebackURL(interval catalog.Interval, id int64) string {
	q := url.Values{"plan": {"premium"}, "interval": {string(interval)}, "comeback": {fmt.Sprint(id)}}
	return s.PublicBaseURL + "/premium/checkout?" + q.Encode()
}

// Sweep releases, in every guild with seats in use, the holders who are no
// longer verified members, and every seat of a guild whose plan has lapsed.
// It returns how many seats it released.
func (s *Service) Sweep(ctx context.Context, now time.Time) (int, error) {
	rows, err := s.Store.Pool.Query(ctx, `select distinct guild_id from premium_seats where state <> 'empty' order by guild_id`)
	if err != nil {
		return 0, wrap("sweep list", err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, wrap("sweep list", err)
	}
	total := 0
	for _, id := range ids {
		n, err := s.sweepGuild(ctx, id, now)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

// SweepGuild is Sweep for one guild.
func (s *Service) SweepGuild(ctx context.Context, guildID int64, now time.Time) error {
	_, err := s.sweepGuild(ctx, guildID, now)
	return err
}

func (s *Service) sweepGuild(ctx context.Context, guildID int64, now time.Time) (int, error) {
	tier, err := s.Store.activeTier(ctx, guildID)
	if err != nil {
		return 0, err
	}
	var outbox []mail.Message
	released := 0
	err = s.Store.withGuildLock(ctx, guildID, func(tx pgx.Tx) error {
		rows, err := s.Store.seatRows(ctx, tx, guildID)
		if err != nil {
			return err
		}
		for slot := 1; slot <= catalog.SeatsRaidGuild; slot++ {
			row, ok := rows[slot]
			if !ok || row.holderID == nil {
				continue
			}
			keep, err := s.keepSeat(ctx, tx, guildID, tier, row)
			if err != nil || keep {
				if err != nil {
					return err
				}
				continue
			}
			msg, err := s.release(ctx, tx, guildID, row, now)
			if err != nil {
				return err
			}
			released++
			if msg != nil {
				outbox = append(outbox, *msg)
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	s.send(ctx, outbox)
	return released, nil
}

// keepSeat is false for a seat whose plan has no access or whose holder is
// no longer a verified member.
func (s *Service) keepSeat(ctx context.Context, tx pgx.Tx, guildID int64, tier string, row seatRow) (bool, error) {
	if tier == "" {
		return false, nil
	}
	var member bool
	err := tx.QueryRow(ctx, `select exists (select 1 from guild_members
		where guild_id = $1 and user_id = $2 and verified_at is not null)`, guildID, *row.holderID).Scan(&member)
	return member, wrap("holder membership", err)
}
