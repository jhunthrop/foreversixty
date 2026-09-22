package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// ErrNotFound is returned when a row that must exist does not: an unknown
// session, an expired or spent token, a revoked device.
var ErrNotFound = errors.New("auth: not found")

// ErrInvalidCharacter is returned when a Character's Key does not match
// its own Region, Ruleset and Name under the character package's key
// rules, or names a region or ruleset outside the contract.
var ErrInvalidCharacter = errors.New("auth: invalid character")

// ErrCharacterClaimed is returned when LinkCharacter is asked to move a
// character to an account that does not already own it, away from an
// account that does.
var ErrCharacterClaimed = errors.New("auth: character already claimed by another account")

// SessionTTL is how long a browser session lasts, as the contract sets it.
const SessionTTL = 30 * 24 * time.Hour

// LoginTokenTTL is how long an email magic link stays clickable, as the
// contract sets it.
const LoginTokenTTL = 20 * time.Minute

// PairingCodeTTL is how long a device pairing code stays claimable, as
// the contract sets it.
const PairingCodeTTL = 10 * time.Minute

// User is an account. Battletag and Email are pointers because the
// columns are nullable and the site distinguishes "no battletag" from
// an empty one: an account signed in by email has no battletag at all.
type User struct {
	ID        int64   `json:"id"`
	BnetSub   string  `json:"-"`
	Battletag *string `json:"battletag"`
	Email     *string `json:"email"`
	Role      string  `json:"role"`
	Anonymize bool    `json:"anonymize"`
	// BnetImportedAt is when the Battle.net import last ran for this
	// account, nil when it never has. Not serialized directly — Me's
	// own top-level bnet_imported_at field carries it (spec §6).
	BnetImportedAt *time.Time `json:"-"`
}

// PublicName is what strangers may be told an account is called: the
// battletag, or a stable pseudonym derived from the id.
//
// It is never the email address. A report's owner is published in the
// report body and in the Open Graph tags of an unauthenticated report
// page, and an account created by magic link has no battletag at all,
// so anything that falls back to the email publishes it. There is
// deliberately no method here that would return it.
//
// An account that has asked to be anonymous reads as the pseudonym
// even when it has a battletag: this is the one read path the
// anonymize flag governs. See the PATCH /v1/me description in
// openapi.yaml for the flag's exact scope.
func (u User) PublicName() string {
	if !u.Anonymize && u.Battletag != nil && *u.Battletag != "" {
		return *u.Battletag
	}
	return fmt.Sprintf("user-%d", u.ID)
}

// Session is a signed-in browser.
type Session struct {
	ID        string
	UserID    int64
	Method    string
	ExpiresAt time.Time
}

// Device is one paired companion install.
type Device struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Platform   string     `json:"platform"`
	CreatedAt  time.Time  `json:"created_at"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
	UserID     int64      `json:"-"`
}

// Character is a character linked to an account, or seen in a report.
type Character struct {
	Key     string `json:"key"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Name    string `json:"name"`
	Class   string `json:"class,omitempty"`
	// Realm, Level and Faction are what Blizzard's profile API adds
	// (spec §6); omitted when unknown (an export-sourced character that
	// has never been through the Battle.net import).
	Realm   string `json:"realm,omitempty"`
	Level   *int   `json:"level,omitempty"`
	Faction string `json:"faction,omitempty"`
	// Race and Gender come from the Battle.net account-profile capture
	// (spec §B/C); omitted when the character has never been through
	// the Battle.net import.
	Race   string `json:"race,omitempty"`
	Gender string `json:"gender,omitempty"`
	// ItemLevel is the character's equipped item level from the last
	// Battle.net profile capture; omitted when unknown.
	ItemLevel *int `json:"item_level,omitempty"`
	// AvatarURL and RenderURL are Blizzard's own render-CDN images for
	// this character (spec .superpowers/account-visual-brief.md §A4),
	// omitted when the Battle.net import has not captured them.
	AvatarURL *string `json:"avatar_url,omitempty"`
	RenderURL *string `json:"render_url,omitempty"`
	// Source is which path most recently wrote this row: "export" or
	// "bnet".
	Source string `json:"source"`
	// Guild is this character's current guild membership, omitted when
	// it has none.
	Guild *CharacterGuild `json:"guild,omitempty"`
}

// CharacterGuild is a character's guild_characters row, as GET /v1/me
// exposes it (spec §6).
type CharacterGuild struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Rank      string `json:"rank"`
	RankIndex *int   `json:"rank_index,omitempty"`
	Verified  bool   `json:"verified"`
}

// Guild is a guild the user belongs to.
type Guild struct {
	ID       int64  `json:"id"`
	Region   string `json:"region"`
	Ruleset  string `json:"ruleset"`
	Name     string `json:"name"`
	Rank     string `json:"rank,omitempty"`
	Consent  string `json:"consent,omitempty"`
	Verified bool   `json:"verified"`
	// Plan is non-nil only when this guild currently has an active guild
	// plan and the caller is a verified officer/leader of it (spec §1.4;
	// Ruling C — a non-officer member sees the features it unlocks, not
	// the billing detail).
	Plan *GuildBillingView `json:"plan,omitempty"`
}

// Store is every account read and write. One type rather than one per
// table: they are all small, they all belong to sign-in, and the handlers
// take them as one dependency.
type Store struct{ Pool *pgxpool.Pool }

const userColumns = `id, coalesce(bnet_sub, ''), battletag, email, role, anonymize, bnet_imported_at`

func scanUser(row pgx.Row) (User, error) {
	var u User
	if err := row.Scan(&u.ID, &u.BnetSub, &u.Battletag, &u.Email, &u.Role, &u.Anonymize, &u.BnetImportedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, ErrNotFound
		}
		return User{}, fmt.Errorf("auth: read user: %w", err)
	}
	return u, nil
}

// UpsertBnetUser finds or creates the account for a Battle.net subject,
// refreshing the battletag, which the player can change at any time.
func (s *Store) UpsertBnetUser(ctx context.Context, sub, battletag string) (User, error) {
	return scanUser(s.Pool.QueryRow(ctx,
		`insert into users (bnet_sub, battletag) values ($1, $2)
		 on conflict (bnet_sub) do update set battletag = excluded.battletag
		 returning `+userColumns, sub, battletag))
}

// UpsertEmailUser finds or creates the account for an email address.
func (s *Store) UpsertEmailUser(ctx context.Context, email string) (User, error) {
	return scanUser(s.Pool.QueryRow(ctx,
		`insert into users (email) values ($1)
		 on conflict (email) do update set email = excluded.email
		 returning `+userColumns, email))
}

// SetAnonymize turns the read-time pseudonym on or off for an account,
// and returns the account as it now stands.
func (s *Store) SetAnonymize(ctx context.Context, id int64, anonymize bool) (User, error) {
	return scanUser(s.Pool.QueryRow(ctx,
		`update users set anonymize = $2 where id = $1 returning `+userColumns, id, anonymize))
}

// User reads one account.
func (s *Store) User(ctx context.Context, id int64) (User, error) {
	return scanUser(s.Pool.QueryRow(ctx, `select `+userColumns+` from users where id = $1`, id))
}

// CreateSession stores a new browser session and returns it.
func (s *Store) CreateSession(ctx context.Context, userID int64, method string, ttl time.Duration) (Session, error) {
	sess := Session{ID: NewSessionID(), UserID: userID, Method: method, ExpiresAt: time.Now().UTC().Add(ttl)}
	if _, err := s.Pool.Exec(ctx,
		`insert into sessions (id, user_id, method, expires_at) values ($1, $2, $3, $4)`,
		sess.ID, sess.UserID, sess.Method, sess.ExpiresAt); err != nil {
		return Session{}, fmt.Errorf("auth: create session: %w", err)
	}
	return sess, nil
}

// SessionUser resolves a session id to its account, or ErrNotFound when
// the session is unknown or has expired.
func (s *Store) SessionUser(ctx context.Context, id string) (User, error) {
	return scanUser(s.Pool.QueryRow(ctx,
		`select `+userColumns+` from users
		 where id = (select user_id from sessions where id = $1 and expires_at > now())`, id))
}

// DeleteSession signs a browser out.
func (s *Store) DeleteSession(ctx context.Context, id string) error {
	if _, err := s.Pool.Exec(ctx, `delete from sessions where id = $1`, id); err != nil {
		return fmt.Errorf("auth: delete session: %w", err)
	}
	return nil
}

// CreateLoginToken stores a magic-link token's hash.
func (s *Store) CreateLoginToken(ctx context.Context, email string, hash []byte, ttl time.Duration) error {
	if _, err := s.Pool.Exec(ctx,
		`insert into login_tokens (token_hash, email, expires_at) values ($1, $2, now() + $3::interval)`,
		hash, email, fmt.Sprintf("%d seconds", int(ttl.Seconds()))); err != nil {
		return fmt.Errorf("auth: create login token: %w", err)
	}
	return nil
}

// RecentLoginTokens counts the links issued for an address within window,
// which is how the five-per-hour budget is enforced.
func (s *Store) RecentLoginTokens(ctx context.Context, email string, window time.Duration) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from login_tokens where email = $1 and created_at > now() - $2::interval`,
		email, fmt.Sprintf("%d seconds", int(window.Seconds()))).Scan(&n); err != nil {
		return 0, fmt.Errorf("auth: count login tokens: %w", err)
	}
	return n, nil
}

// ConsumeLoginToken spends a magic-link token and returns its address.
// The update is the check: a token already used, or expired, matches no
// row and comes back ErrNotFound, so two clicks cannot both sign in.
func (s *Store) ConsumeLoginToken(ctx context.Context, hash []byte) (string, error) {
	var email string
	err := s.Pool.QueryRow(ctx,
		`update login_tokens set used_at = now()
		 where token_hash = $1 and used_at is null and expires_at > now()
		 returning email`, hash).Scan(&email)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("auth: consume login token: %w", err)
	}
	return email, nil
}

// CreatePairingCode issues a device pairing code for a user.
func (s *Store) CreatePairingCode(ctx context.Context, userID int64, ttl time.Duration) (string, error) {
	code := NewPairingCode()
	if _, err := s.Pool.Exec(ctx,
		`insert into pairing_codes (code, user_id, expires_at) values ($1, $2, now() + $3::interval)`,
		code, userID, fmt.Sprintf("%d seconds", int(ttl.Seconds()))); err != nil {
		return "", fmt.Errorf("auth: create pairing code: %w", err)
	}
	return code, nil
}

// ClaimPairingCode spends a pairing code and returns whose it was.
func (s *Store) ClaimPairingCode(ctx context.Context, code string) (int64, error) {
	var userID int64
	err := s.Pool.QueryRow(ctx,
		`update pairing_codes set claimed_at = now()
		 where code = $1 and claimed_at is null and expires_at > now()
		 returning user_id`, code).Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, ErrNotFound
	}
	if err != nil {
		return 0, fmt.Errorf("auth: claim pairing code: %w", err)
	}
	return userID, nil
}

// CreateDevice stores a paired device and returns it.
func (s *Store) CreateDevice(ctx context.Context, userID int64, name, platform string, hash []byte) (Device, error) {
	d := Device{ID: NewDeviceID(), UserID: userID, Name: name, Platform: platform}
	if err := s.Pool.QueryRow(ctx,
		`insert into devices (id, user_id, name, platform, token_hash) values ($1, $2, $3, $4, $5)
		 returning created_at`, d.ID, userID, name, platform, hash).Scan(&d.CreatedAt); err != nil {
		return Device{}, fmt.Errorf("auth: create device: %w", err)
	}
	return d, nil
}

// DeviceByToken resolves a bearer token to its device and stamps
// last_seen_at, so the account page can show when a companion last
// uploaded. A revoked device matches nothing.
func (s *Store) DeviceByToken(ctx context.Context, hash []byte) (Device, error) {
	var d Device
	err := s.Pool.QueryRow(ctx,
		`update devices set last_seen_at = now()
		 where token_hash = $1 and revoked_at is null
		 returning id, user_id, name, platform, created_at, last_seen_at`, hash).
		Scan(&d.ID, &d.UserID, &d.Name, &d.Platform, &d.CreatedAt, &d.LastSeenAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Device{}, ErrNotFound
	}
	if err != nil {
		return Device{}, fmt.Errorf("auth: read device: %w", err)
	}
	return d, nil
}

// Devices lists a user's live devices, newest first.
func (s *Store) Devices(ctx context.Context, userID int64) ([]Device, error) {
	rows, err := s.Pool.Query(ctx,
		`select id, user_id, name, platform, created_at, last_seen_at from devices
		 where user_id = $1 and revoked_at is null order by created_at desc`, userID)
	if err != nil {
		return nil, fmt.Errorf("auth: list devices: %w", err)
	}
	defer rows.Close()
	out := []Device{}
	for rows.Next() {
		var d Device
		if err := rows.Scan(&d.ID, &d.UserID, &d.Name, &d.Platform, &d.CreatedAt, &d.LastSeenAt); err != nil {
			return nil, fmt.Errorf("auth: list devices: %w", err)
		}
		out = append(out, d)
	}
	return out, rows.Err()
}

// RevokeDevice revokes one of a user's own devices. It reports false when
// the id is not theirs or is revoked already.
func (s *Store) RevokeDevice(ctx context.Context, userID int64, id string) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`update devices set revoked_at = now() where id = $1 and user_id = $2 and revoked_at is null`, id, userID)
	if err != nil {
		return false, fmt.Errorf("auth: revoke device: %w", err)
	}
	return tag.RowsAffected() == 1, nil
}

// characterColumns is every column both Characters and CharactersByKeys
// select, in the order scanCharacterRows reads them.
const characterColumns = `c.key, c.region, c.ruleset, c.name, coalesce(c.class, ''),
	        coalesce(c.realm_name, ''), c.level, coalesce(c.faction, ''), c.source,
	        coalesce(c.race, ''), coalesce(c.gender, ''), c.equipped_item_level,
	        c.avatar_url, c.render_url,
	        g.id, g.name, gc.rank, gc.rank_index, gc.verified_at is not null`

// characterFrom is the join every character read shares: a character
// with its current guild membership, if any, attached (spec §6).
const characterFrom = `from characters c
	 left join guild_characters gc on gc.character_key = c.key
	 left join guilds g on g.id = gc.guild_id`

// scanCharacterRows reads every row of a characterColumns/characterFrom
// query into Characters, attaching a CharacterGuild wherever the guild
// join matched.
func scanCharacterRows(rows pgx.Rows) ([]Character, error) {
	defer rows.Close()
	out := []Character{}
	for rows.Next() {
		var c Character
		var guildID *int64
		var guildName, rank *string
		var rankIndex *int
		var verified bool
		if err := rows.Scan(&c.Key, &c.Region, &c.Ruleset, &c.Name, &c.Class,
			&c.Realm, &c.Level, &c.Faction, &c.Source,
			&c.Race, &c.Gender, &c.ItemLevel,
			&c.AvatarURL, &c.RenderURL,
			&guildID, &guildName, &rank, &rankIndex, &verified); err != nil {
			return nil, fmt.Errorf("auth: scan characters: %w", err)
		}
		if guildID != nil {
			c.Guild = &CharacterGuild{
				ID: *guildID, Name: *guildName, Rank: *rank, RankIndex: rankIndex, Verified: verified,
			}
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// Characters lists the characters linked to an account, each with its
// current guild membership (if any) attached (spec §6).
func (s *Store) Characters(ctx context.Context, userID int64) ([]Character, error) {
	rows, err := s.Pool.Query(ctx,
		`select `+characterColumns+` `+characterFrom+` where c.user_id = $1 order by c.key`, userID)
	if err != nil {
		return nil, fmt.Errorf("auth: list characters: %w", err)
	}
	return scanCharacterRows(rows)
}

// CharactersByKeys reads the same per-character shape Characters does,
// filtered to a specific set of keys rather than an account — what
// POST /v1/me/exports answers with for the keys it just wrote (spec §4.5).
func (s *Store) CharactersByKeys(ctx context.Context, keys []string) ([]Character, error) {
	if len(keys) == 0 {
		return []Character{}, nil
	}
	rows, err := s.Pool.Query(ctx,
		`select `+characterColumns+` `+characterFrom+` where c.key = any($1) order by c.key`, keys)
	if err != nil {
		return nil, fmt.Errorf("auth: characters by keys: %w", err)
	}
	return scanCharacterRows(rows)
}

// LinkCharacter records a character as belonging to an account. It
// validates that c.Key is exactly the key the character package would
// build from c.Region, c.Ruleset and c.Name — ErrInvalidCharacter
// otherwise, never a raw database error. If the key is new, or belongs
// to no one, or already belongs to userID, the row is created or its
// mutable fields (owner, class) are refreshed. If the key already
// belongs to a different account, it is not moved: LinkCharacter returns
// ErrCharacterClaimed rather than stealing it.
func (s *Store) LinkCharacter(ctx context.Context, userID int64, c Character) error {
	if !character.ValidRegion(c.Region) || !character.ValidRuleset(c.Ruleset) ||
		c.Key != character.Key(c.Region, c.Ruleset, c.Name) {
		return ErrInvalidCharacter
	}
	tag, err := s.Pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, class, user_id, refreshed_at)
		 values ($1, $2, $3, $4, nullif($5, ''), $6, now())
		 on conflict (key) do update set user_id = excluded.user_id,
		   class = coalesce(excluded.class, characters.class), refreshed_at = now()
		 where characters.user_id = excluded.user_id or characters.user_id is null`,
		c.Key, c.Region, c.Ruleset, c.Name, c.Class, userID)
	if err != nil {
		return fmt.Errorf("auth: link character: %w", err)
	}
	if tag.RowsAffected() == 0 {
		// The row already existed and its owner is neither userID nor
		// nobody: the WHERE guard refused the update.
		return ErrCharacterClaimed
	}
	return nil
}

// MemberKeys answers which of these character keys belong to an
// account. The ingest asks once per fight, so it is one statement
// rather than one per player.
func (s *Store) MemberKeys(ctx context.Context, keys []string) (map[string]bool, error) {
	out := make(map[string]bool, len(keys))
	if len(keys) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx,
		`select key from characters where key = any($1) and user_id is not null`, keys)
	if err != nil {
		return nil, fmt.Errorf("auth: member keys: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, fmt.Errorf("auth: member keys: %w", err)
		}
		out[key] = true
	}
	return out, rows.Err()
}

// Guilds lists the guilds an account is a member of, most recently
// active membership first — the "My guild" link in the header reads
// position [0].
func (s *Store) Guilds(ctx context.Context, userID int64) ([]Guild, error) {
	rows, err := s.Pool.Query(ctx,
		`select g.id, g.region, g.ruleset, g.name, m.rank, m.consent, m.verified_at is not null
		 from guilds g
		 join guild_members m on m.guild_id = g.id
		 where m.user_id = $1 order by m.refreshed_at desc`, userID)
	if err != nil {
		return nil, fmt.Errorf("auth: list guilds: %w", err)
	}
	defer rows.Close()
	out := []Guild{}
	for rows.Next() {
		var g Guild
		if err := rows.Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name, &g.Rank, &g.Consent, &g.Verified); err != nil {
			return nil, fmt.Errorf("auth: list guilds: %w", err)
		}
		out = append(out, g)
	}
	return out, rows.Err()
}

// GuildRank reports a user's rank in a guild, and whether they are a
// *verified* member at all. reports.mayView/mayEdit read this: an
// unverified guild_characters row (a bare export, possibly forged) must
// never grant view or edit rights on a guild-visible report.
func (s *Store) GuildRank(ctx context.Context, guildID, userID int64) (string, bool, error) {
	var rank string
	err := s.Pool.QueryRow(ctx,
		`select rank from guild_members where guild_id = $1 and user_id = $2 and verified_at is not null`,
		guildID, userID).Scan(&rank)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("auth: guild rank: %w", err)
	}
	return rank, true, nil
}
