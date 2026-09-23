// api/internal/guilds/moderation.go
package guilds

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// ModerationClaimsPerPage is the moderation claim queue's page size,
// the same keyset-paginated shape GET /v1/reports/recent and
// GET /v1/guilds/{id}/home both already use.
const ModerationClaimsPerPage = 20

// ClaimEvidence is one side of a contested claim's plain, unweighted
// facts (item 4, fourth security review response): the site hands a
// moderator raw signal, never a verdict, since it is the moderator who
// decides which side is telling the truth.
type ClaimEvidence struct {
	RankIndex  *int    `json:"rank_index"`
	Verified   bool    `json:"verified"`
	VerifiedBy *string `json:"verified_by,omitempty"`
	// NightsInOthersReports (fifth security review response - renamed
	// from independent_nights) is how many distinct raid nights within
	// 30 days this character appears in a report of this guild's that
	// it did NOT itself upload. Read it as one input among several, not
	// a verdict: it says nothing about who the uploader actually is -
	// a squatter's own second account can upload a report naming a
	// sockpuppet's character (VerifyByLogs accepts exactly this, see
	// its own doc comment) - only that the appearance was not
	// self-reported.
	NightsInOthersReports int `json:"nights_in_others_reports"`
}

// ModerationParty is one side of a contested claim as the queue lists it.
type ModerationParty struct {
	Battletag string        `json:"battletag"`
	Evidence  ClaimEvidence `json:"evidence"`
}

// ModerationClaim is one open (contested, unresolved) claim.
type ModerationClaim struct {
	Guild       GuildIdentity   `json:"guild"`
	Claimant    ModerationParty `json:"claimant"`
	Contester   ModerationParty `json:"contester"`
	ClaimedAt   *time.Time      `json:"claimed_at"`
	ContestedAt time.Time       `json:"contested_at"`
}

const moderationCursorSep = "|"

type moderationCursor struct {
	ContestedAt time.Time
	GuildID     int64
}

func encodeModerationCursor(contestedAt time.Time, guildID int64) string {
	raw := contestedAt.UTC().Format(time.RFC3339Nano) + moderationCursorSep + strconv.FormatInt(guildID, 10)
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

func decodeModerationCursor(s string) (moderationCursor, bool) {
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return moderationCursor{}, false
	}
	contestedAt, idStr, ok := strings.Cut(string(raw), moderationCursorSep)
	if !ok || idStr == "" {
		return moderationCursor{}, false
	}
	t, err := time.Parse(time.RFC3339Nano, contestedAt)
	if err != nil {
		return moderationCursor{}, false
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return moderationCursor{}, false
	}
	return moderationCursor{ContestedAt: t, GuildID: id}, true
}

// evidenceFor reads the strongest single guild_characters row userID
// holds in guildID - verified over unverified, lowest rank_index (the
// guild master's own is 0) next, most recently refreshed last - and
// reports rank_index, verification state/source, and how many distinct
// raid nights within 30 days its character appears in this guild's
// reports that it does NOT own (the same independence VerifyByLogs
// itself now always applies, surfaced here as a plain count rather
// than folded into a pass/fail, since a moderator weighs the evidence
// directly rather than trusting a verdict the site already reached).
func (s *Store) evidenceFor(ctx context.Context, guildID, userID int64) (ClaimEvidence, error) {
	var ev ClaimEvidence
	err := s.Pool.QueryRow(ctx, `
		select gc.rank_index, gc.verified_at is not null, gc.verified_by,
		       (select count(distinct r.created_at::date) from fights f join reports r on r.id = f.report_id
		        where r.guild_id = gc.guild_id and gc.character_key = any(f.players)
		          and r.owner_id is distinct from gc.user_id and r.created_at >= now() - interval '30 days')
		from guild_characters gc
		where gc.guild_id = $1 and gc.user_id = $2
		order by (gc.verified_at is not null) desc, gc.rank_index nulls last, gc.refreshed_at desc
		limit 1
	`, guildID, userID).Scan(&ev.RankIndex, &ev.Verified, &ev.VerifiedBy, &ev.NightsInOthersReports)
	if errors.Is(err, pgx.ErrNoRows) {
		return ClaimEvidence{}, nil
	}
	if err != nil {
		return ClaimEvidence{}, fmt.Errorf("guilds: evidence for %d in guild %d: %w", userID, guildID, err)
	}
	return ev, nil
}

// openClaim is one contested guild row as the store reads it, before
// the service resolves battletags and evidence.
type openClaim struct {
	Guild       GuildIdentity
	ClaimantID  *int64
	ContesterID int64
	ClaimedAt   *time.Time
	ContestedAt time.Time
}

// OpenClaims lists every currently-contested claim, newest contest
// first, keyset-paginated - the moderation queue (item 4, fourth
// security review response). claimant is whichever of claimed_by/
// claim_pending_by is set (claimed takes priority), mirroring
// ResolveClaim's own resolution; nil only if a contested claim was
// somehow never actually claimed or pending, which contest state
// should never actually allow.
func (s *Store) OpenClaims(ctx context.Context, before *moderationCursor) ([]openClaim, error) {
	query := `select id, region, ruleset, name, claimed_by, claim_pending_by, claimed_at, claim_contested_at, claim_contested_by
		from guilds where claim_contested_at is not null`
	args := []any{}
	if before != nil {
		query += ` and (claim_contested_at, id) < ($1, $2)`
		args = append(args, before.ContestedAt, before.GuildID)
	}
	query += fmt.Sprintf(` order by claim_contested_at desc, id desc limit $%d`, len(args)+1)
	args = append(args, ModerationClaimsPerPage)

	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("guilds: open claims: %w", err)
	}
	defer rows.Close()
	out := []openClaim{}
	for rows.Next() {
		var c openClaim
		var claimedBy, pendingBy, contestedBy *int64
		if err := rows.Scan(&c.Guild.ID, &c.Guild.Region, &c.Guild.Ruleset, &c.Guild.Name,
			&claimedBy, &pendingBy, &c.ClaimedAt, &c.ContestedAt, &contestedBy); err != nil {
			return nil, fmt.Errorf("guilds: open claims: %w", err)
		}
		claimant := claimedBy
		if claimant == nil {
			claimant = pendingBy
		}
		c.ClaimantID = claimant
		if contestedBy != nil {
			c.ContesterID = *contestedBy
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// moderationClaims answers 404, not 403, for a non-moderator - the
// route's own existence is not advertised to anyone it does not
// concern (item 4, fourth security review response). No RequireSession
// wrap: an unauthenticated caller hits the same check and the same
// answer, rather than a 401 that would itself reveal the route needs
// signing in. Each side's evidence.nights_in_others_reports (fifth
// security review response) means only "not self-reported" - it says
// nothing about who the uploader actually is, so it is one fact among
// several for a moderator to weigh, never a verdict on its own.
func (s *Service) moderationClaims(w http.ResponseWriter, r *http.Request) {
	if !auth.ActorFrom(r.Context()).IsModerator() {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such route", nil)
		return
	}
	var before *moderationCursor
	if v := r.URL.Query().Get("cursor"); v != "" {
		c, ok := decodeModerationCursor(v)
		if !ok {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a cursor",
				map[string]string{"cursor": "the next_cursor a previous page returned"})
			return
		}
		before = &c
	}
	raw, err := s.Store.OpenClaims(r.Context(), before)
	if err != nil {
		s.fail(w, r, "moderation_claims", err, "could not read the claim queue just now")
		return
	}
	claims := make([]ModerationClaim, 0, len(raw))
	for _, c := range raw {
		item := ModerationClaim{Guild: c.Guild, ClaimedAt: c.ClaimedAt, ContestedAt: c.ContestedAt}
		if c.ClaimantID != nil {
			if u, uerr := s.Accounts.User(r.Context(), *c.ClaimantID); uerr == nil {
				item.Claimant.Battletag = u.PublicName()
			}
			if ev, everr := s.Store.evidenceFor(r.Context(), c.Guild.ID, *c.ClaimantID); everr == nil {
				item.Claimant.Evidence = ev
			}
		}
		if u, uerr := s.Accounts.User(r.Context(), c.ContesterID); uerr == nil {
			item.Contester.Battletag = u.PublicName()
		}
		if ev, everr := s.Store.evidenceFor(r.Context(), c.Guild.ID, c.ContesterID); everr == nil {
			item.Contester.Evidence = ev
		}
		claims = append(claims, item)
	}
	resp := map[string]any{"claims": claims}
	if len(raw) == ModerationClaimsPerPage {
		last := raw[len(raw)-1]
		resp["next_cursor"] = encodeModerationCursor(last.ContestedAt, last.Guild.ID)
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, resp)
}
