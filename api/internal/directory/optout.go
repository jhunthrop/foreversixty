// api/internal/directory/optout.go
package directory

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// Opt-out states the page shows.
const (
	OptOutRemoved = "removed"
	OptOutWaiting = "waiting"
	OptOutNone    = "none"
)

const (
	// optOutCodeLife is how long a code waits for the character's guild note.
	optOutCodeLife = 7 * 24 * time.Hour
	// CodePrefix starts every opt-out code, so the addon can find one in a note.
	CodePrefix    = "FS-"
	codeAlphabet  = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	codeLength    = 6
	methodAccount = "battlenet"
	methodNote    = "note"
)

// codePattern is a whole opt-out code.
var codePattern = regexp.MustCompile(`^FS-[A-HJ-NP-Z2-9]{6}$`)

var (
	// ErrNoSuchCharacter is a name neither the directory nor Blizzard's public profile knows.
	ErrNoSuchCharacter = errors.New("directory: no such character")
	// ErrNotYours is an account asking to remove a character it does not hold.
	ErrNotYours = errors.New("directory: not a character of this account")
	// ErrBadCharacter is a region or name that cannot be a Forever character.
	ErrBadCharacter = errors.New("directory: not a character name")
)

// OptOuts removes characters from the directory at their players' request and keeps them out.
type OptOuts struct {
	Pool *pgxpool.Pool
	// Client looks up a name the directory does not hold, to keep it out from now on.
	Client *bnetapi.Client
}

// Code is an opt-out code handed to the asker.
type Code struct {
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}

// NoteCode is a code a guildmate's addon found in a member's public guild note.
type NoteCode struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// ValidCode reports whether s is a well-formed opt-out code.
func ValidCode(s string) bool { return codePattern.MatchString(s) }

// target is the character an opt-out names.
type target struct {
	region, slug string
	id           int64
}

// parseTarget reads a region and a character name, or ErrBadCharacter.
func parseTarget(region, name string) (target, error) {
	region = strings.ToLower(strings.TrimSpace(region))
	slug := character.Slug(name)
	if !directoryRegion(region) || !character.ValidSlug(slug) {
		return target{}, ErrBadCharacter
	}
	return target{region: region, slug: slug}, nil
}

// resolve is the Blizzard character id behind a region and name: the directory's, else the
// public profile's (one call). ErrNoSuchCharacter when neither knows it.
func (o *OptOuts) resolve(ctx context.Context, region, name string) (target, error) {
	t, err := parseTarget(region, name)
	if err != nil {
		return target{}, err
	}
	region, slug := t.region, t.slug
	err = o.Pool.QueryRow(ctx,
		`select bnet_character_id from character_directory where region = $1 and slug = $2
		 union all
		 select bnet_character_id from directory_optout_codes where region = $1 and slug = $2
		 limit 1`, region, slug).Scan(&t.id)
	if err == nil {
		return t, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return target{}, fmt.Errorf("directory: resolve %s/%s: %w", region, slug, err)
	}
	if o.Client == nil {
		return target{}, ErrNoSuchCharacter
	}
	profile, _, err := o.Client.Character(ctx, region, "", slug)
	if isGone(err) {
		return target{}, ErrNoSuchCharacter
	}
	if err != nil {
		return target{}, fmt.Errorf("directory: resolve %s/%s: %w", region, slug, err)
	}
	t.id = profile.ID
	return t, nil
}

// StartCode gives a code for the character, replacing any earlier one. The character is
// removed once its public guild note carries the code (VerifyNoteCodes).
func (o *OptOuts) StartCode(ctx context.Context, region, name string) (Code, error) {
	t, err := o.resolve(ctx, region, name)
	if err != nil {
		return Code{}, err
	}
	code, err := newCode()
	if err != nil {
		return Code{}, err
	}
	c := Code{Code: code, ExpiresAt: time.Now().Add(optOutCodeLife).UTC().Truncate(time.Second)}
	_, err = o.Pool.Exec(ctx,
		`insert into directory_optout_codes (region, bnet_character_id, slug, code, expires_at)
		 values ($1, $2, $3, $4, $5)
		 on conflict (region, bnet_character_id) do update set slug = excluded.slug, code = excluded.code,
		   created_at = now(), expires_at = excluded.expires_at, verified_at = null`,
		t.region, t.id, t.slug, c.Code, c.ExpiresAt)
	if err != nil {
		return Code{}, fmt.Errorf("directory: save code %s/%d: %w", t.region, t.id, err)
	}
	return c, nil
}

// Status is where the character's opt-out stands: removed, waiting on its code, or none.
func (o *OptOuts) Status(ctx context.Context, region, name string) (string, error) {
	t, err := o.resolve(ctx, region, name)
	if errors.Is(err, ErrNoSuchCharacter) {
		return OptOutNone, nil
	}
	if err != nil {
		return "", err
	}
	var removed, waiting bool
	err = o.Pool.QueryRow(ctx,
		`select exists (select 1 from directory_optouts where region = $1 and bnet_character_id = $2),
		        exists (select 1 from directory_optout_codes where region = $1 and bnet_character_id = $2
		                  and verified_at is null and expires_at > now())`, t.region, t.id).Scan(&removed, &waiting)
	if err != nil {
		return "", fmt.Errorf("directory: opt-out status: %w", err)
	}
	switch {
	case removed:
		return OptOutRemoved, nil
	case waiting:
		return OptOutWaiting, nil
	}
	return OptOutNone, nil
}

// RemoveOwned removes a character the signed-in account holds through Battle.net.
func (o *OptOuts) RemoveOwned(ctx context.Context, userID int64, region, name string) error {
	t, err := parseTarget(region, name)
	if err != nil {
		return err
	}
	var id int64
	err = o.Pool.QueryRow(ctx,
		`select bnet_character_id from characters
		  where user_id = $1 and region = $2 and split_part(key, '/', 3) = $3 and bnet_character_id is not null
		  limit 1`, userID, t.region, t.slug).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotYours
	}
	if err != nil {
		return fmt.Errorf("directory: owned character: %w", err)
	}
	return o.remove(ctx, t.region, id, methodAccount)
}

// VerifyNoteCodes removes every character whose public guild note carries its unexpired code,
// as a guildmate's addon read the notes of guild: only a character the directory has in that
// guild counts. Returns how many it removed.
func (o *OptOuts) VerifyNoteCodes(ctx context.Context, region, guild string, codes []NoteCode) (int, error) {
	removed := 0
	for _, nc := range codes {
		if !ValidCode(nc.Code) {
			continue
		}
		var id int64
		err := o.Pool.QueryRow(ctx,
			`select c.bnet_character_id from directory_optout_codes c
			   join character_directory d on d.region = c.region and d.bnet_character_id = c.bnet_character_id
			  where c.region = $1 and c.code = $2 and c.slug = $3 and c.expires_at > now() and c.verified_at is null
			    and lower(d.guild_name) = lower($4)`,
			region, nc.Code, character.Slug(nc.Name), guild).Scan(&id)
		if errors.Is(err, pgx.ErrNoRows) {
			continue
		}
		if err != nil {
			return removed, fmt.Errorf("directory: note code: %w", err)
		}
		if err := o.remove(ctx, region, id, methodNote); err != nil {
			return removed, err
		}
		removed++
	}
	return removed, nil
}

// remove records the opt-out and deletes the character from the directory, in one transaction.
func (o *OptOuts) remove(ctx context.Context, region string, id int64, method string) error {
	tx, err := o.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("directory: begin opt-out: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx,
		`insert into directory_optouts (region, bnet_character_id, method) values ($1, $2, $3) on conflict do nothing`,
		region, id, method); err != nil {
		return fmt.Errorf("directory: opt-out %s/%d: %w", region, id, err)
	}
	if _, err := tx.Exec(ctx,
		`delete from character_directory where region = $1 and bnet_character_id = $2`, region, id); err != nil {
		return fmt.Errorf("directory: opt-out delete %s/%d: %w", region, id, err)
	}
	if _, err := tx.Exec(ctx,
		`update directory_optout_codes set verified_at = coalesce(verified_at, now())
		  where region = $1 and bnet_character_id = $2`, region, id); err != nil {
		return fmt.Errorf("directory: opt-out code %s/%d: %w", region, id, err)
	}
	return tx.Commit(ctx)
}

// isOptedOut reports whether a character's player asked to be left out; a failed read is
// treated as opted out, so a database hiccup never stores one.
func (s *Service) isOptedOut(ctx context.Context, region string, id int64) bool {
	var out bool
	if err := s.Pool.QueryRow(ctx,
		`select exists (select 1 from directory_optouts where region = $1 and bnet_character_id = $2)`,
		region, id).Scan(&out); err != nil {
		return true
	}
	return out
}

// directoryRegion reports whether region is one the directory holds (Forever's two).
func directoryRegion(region string) bool { return region == "us" || region == "eu" }

func newCode() (string, error) {
	var b strings.Builder
	b.WriteString(CodePrefix)
	for range codeLength {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeAlphabet))))
		if err != nil {
			return "", fmt.Errorf("directory: code: %w", err)
		}
		b.WriteByte(codeAlphabet[n.Int64()])
	}
	return b.String(), nil
}
