// api/internal/account/store.go
package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/billing"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
)

// ErrNotFound is an account that no longer exists (deleted by a racing call).
var ErrNotFound = errors.New("account: no such account")

// Account is what the deletion reads about the account before it starts.
type Account struct {
	ID        int64
	Battletag *string
	Email     *string
	Role      string
}

// Staff reports whether the account is a moderator or admin. Staff accounts are closed by
// support: a moderator's claim resolutions keep their moderator (an on delete restrict key).
func (a Account) Staff() bool { return a.Role != "user" }

// Obligation is a guild the account answers for, which deleting it would leave in the lurch.
type Obligation struct {
	GuildID   int64
	GuildName string
	// Claim: the account holds the guild's claim (the guild master on the site).
	Claim bool
	// GuildPlan: the account pays for the guild's live Guild plan, not already set to end.
	GuildPlan bool
}

// Erased is what the transaction removed that needs work outside the database.
type Erased struct {
	// ObjectKeys are the storage objects of the deleted private reports and uploads.
	ObjectKeys []string
}

// Store is the deletion's database half.
type Store struct{ Pool *pgxpool.Pool }

// querier is the pool or a transaction.
type querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

const accountSQL = `select id, battletag, email, role from users where id = $1`

// Account reads the account, ErrNotFound when it is gone.
func (s *Store) Account(ctx context.Context, id int64) (Account, error) {
	return readAccount(ctx, s.Pool, accountSQL, id)
}

func readAccount(ctx context.Context, q querier, sql string, id int64) (Account, error) {
	var a Account
	err := q.QueryRow(ctx, sql, id).Scan(&a.ID, &a.Battletag, &a.Email, &a.Role)
	if errors.Is(err, pgx.ErrNoRows) {
		return Account{}, ErrNotFound
	}
	if err != nil {
		return Account{}, fmt.Errorf("account: read %d: %w", id, err)
	}
	return a, nil
}

// obligationsSQL is every guild whose claim the account holds or whose live Guild plan it
// pays for. A plan the payer already set to end at its period end is not an obligation: the
// guild was told it is ending, and the deletion only ends it sooner.
const obligationsSQL = `
	select g.id, g.name, coalesce(g.claimed_by = $1, false) as claim,
	  exists (select 1 from entitlements e
	          where e.guild_id = g.id and e.billing_user_id = $1 and not e.cancel_at_period_end
	            and ` + billing.LiveStripeSQL + `) as guild_plan
	from guilds g
	where g.claimed_by = $1
	   or exists (select 1 from entitlements e
	              where e.guild_id = g.id and e.billing_user_id = $1 and not e.cancel_at_period_end
	                and ` + billing.LiveStripeSQL + `)
	order by lower(g.name), g.id`

// Obligations lists the guilds that block the deletion; none means it may go ahead.
func (s *Store) Obligations(ctx context.Context, userID int64) ([]Obligation, error) {
	return readObligations(ctx, s.Pool, userID)
}

func readObligations(ctx context.Context, q querier, userID int64) ([]Obligation, error) {
	rows, err := q.Query(ctx, obligationsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("account: guild obligations of %d: %w", userID, err)
	}
	defer rows.Close()
	var out []Obligation
	for rows.Next() {
		var o Obligation
		if err := rows.Scan(&o.GuildID, &o.GuildName, &o.Claim, &o.GuildPlan); err != nil {
			return nil, fmt.Errorf("account: guild obligations of %d: %w", userID, err)
		}
		out = append(out, o)
	}
	return out, rows.Err()
}

// ObligationError is Erase's refusal when a guild obligation appeared after the check that
// preceded it (a claim confirmed meanwhile).
type ObligationError struct{ Guilds []Obligation }

func (e *ObligationError) Error() string {
	return fmt.Sprintf("account: %d guild obligation(s) block the deletion", len(e.Guilds))
}

// Erase deletes the account in one transaction. Postgres does most of it: every user-keyed
// table follows its own foreign key when the users row goes (cascade for what is the
// account's own, set null for shared records it merely touched). The statements before that
// handle what no key expresses, each with its reason below. The account row is locked first
// and the guild check repeated under the lock.
func (s *Store) Erase(ctx context.Context, userID int64) (Erased, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return Erased{}, fmt.Errorf("account: erase %d: begin: %w", userID, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	acct, err := readAccount(ctx, tx, accountSQL+` for update`, userID)
	if err != nil {
		return Erased{}, err
	}
	obligations, err := readObligations(ctx, tx, userID)
	if err != nil {
		return Erased{}, err
	}
	if len(obligations) > 0 {
		return Erased{}, &ObligationError{Guilds: obligations}
	}
	keys, err := erasePrivateReports(ctx, tx, userID)
	if err != nil {
		return Erased{}, err
	}
	if err := eraseActivity(ctx, tx, userID); err != nil {
		return Erased{}, err
	}
	if err := eraseCharacters(ctx, tx, userID); err != nil {
		return Erased{}, err
	}
	if err := eraseSignups(ctx, tx, userID); err != nil {
		return Erased{}, err
	}
	if err := eraseAccountRow(ctx, tx, acct); err != nil {
		return Erased{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return Erased{}, fmt.Errorf("account: erase %d: commit: %w", userID, err)
	}
	return Erased{ObjectKeys: keys}, nil
}

// step is one statement of the erase, named for its error.
type step struct {
	name string
	sql  string
}

// run executes steps in order with the same arguments.
func run(ctx context.Context, tx pgx.Tx, steps []step, args ...any) error {
	for _, st := range steps {
		if _, err := tx.Exec(ctx, st.sql, args...); err != nil {
			return fmt.Errorf("account: erase: %s: %w", st.name, err)
		}
	}
	return nil
}

// privateReportsSQL picks the reports that go with the account. Private reports were only
// ever the account's own: they go, with everything parsed from them. Public and unlisted
// reports were published (unlisted ones to anyone holding the link, and their files are
// served straight off the public bucket), and guild reports are the guild's attendance and
// loot record, shared with every raider in them: those stay, and the users row's own key
// (reports.owner_id on delete set null) unlinks the uploader, so no name is shown on them
// at all, the strictest form of the anonymise switch (auth.User.PublicName is never asked).
const privateReportsSQL = `select id from reports where owner_id = $1 and visibility = '` + reports.Private + `'`

// erasePrivateReports deletes the account's private reports, the rows parsed from them
// that carry no foreign key (fight metrics, ratings, coach profiles; every other report
// table cascades), and the account's uploads that belong to no kept report. It returns
// the storage keys of all of it, read before the rows go.
func erasePrivateReports(ctx context.Context, tx pgx.Tx, userID int64) ([]string, error) {
	ids, err := collectStrings(ctx, tx, privateReportsSQL, userID)
	if err != nil {
		return nil, fmt.Errorf("account: erase: private reports: %w", err)
	}
	keys, err := reportObjectKeys(ctx, tx, ids)
	if err != nil {
		return nil, err
	}
	uploads, err := collectStrings(ctx, tx,
		`delete from uploads where user_id = $1 and (report_id is null or report_id = any($2))
		 returning object_key`, userID, ids)
	if err != nil {
		return nil, fmt.Errorf("account: erase: uploads: %w", err)
	}
	if len(ids) > 0 {
		if err := run(ctx, tx, []step{
			{"fight metrics", `delete from fight_metrics where report_id = any($1)`},
			{"rating scores", `delete from rating_scores where report_id = any($1)`},
			{"coach profiles", `delete from coach_profiles where report_id = any($1)`},
			{"reports", `delete from reports where id = any($1)`},
		}, ids); err != nil {
			return nil, err
		}
	}
	return append(keys, uploads...), nil
}

// reportObjectKeys is every stored object of the reports: report.json, each fight's three
// files and each raw chunk, built with the reports package's own key layout.
func reportObjectKeys(ctx context.Context, tx pgx.Tx, ids []string) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var keys []string
	for _, id := range ids {
		keys = append(keys, reports.Keys(id).Report())
	}
	fights, err := tx.Query(ctx, `select report_id, fight_index from fights where report_id = any($1)`, ids)
	if err != nil {
		return nil, fmt.Errorf("account: erase: report fights: %w", err)
	}
	fightKeys, err := pgx.CollectRows(fights, func(row pgx.CollectableRow) ([]string, error) {
		var id string
		var n int
		err := row.Scan(&id, &n)
		k := reports.Keys(id)
		return []string{k.FightSummary(n), k.FightEvents(n), k.FightLive(n)}, err
	})
	if err != nil {
		return nil, fmt.Errorf("account: erase: report fights: %w", err)
	}
	chunks, err := tx.Query(ctx, `select report_id, start_offset from raw_chunks where report_id = any($1)`, ids)
	if err != nil {
		return nil, fmt.Errorf("account: erase: raw chunks: %w", err)
	}
	chunkKeys, err := pgx.CollectRows(chunks, func(row pgx.CollectableRow) (string, error) {
		var id string
		var offset int64
		err := row.Scan(&id, &offset)
		return reports.Keys(id).Raw(offset), err
	})
	if err != nil {
		return nil, fmt.Errorf("account: erase: raw chunks: %w", err)
	}
	for _, fk := range fightKeys {
		keys = append(keys, fk...)
	}
	return append(keys, chunkKeys...), nil
}

// eraseActivity deletes the feed rows about the account's characters (activity_events is
// keyed by character, not account) and its saved sims; character_sims and the addon's sim
// requests follow by key.
func eraseActivity(ctx context.Context, tx pgx.Tx, userID int64) error {
	return run(ctx, tx, []step{
		{"activity events", `delete from activity_events
			where character_key in (select key from characters where user_id = $1)`},
		{"sims", `delete from sims where user_id = $1`},
	}, userID)
}

// eraseCharacters handles the account's characters, which characters.user_id (on delete set
// null) makes unowned when the row goes. What the account supplied goes now: its addon
// exports follow the users row, and the addon half of each character's sync history goes
// here. A character only ever known from those exports (never imported from Battle.net,
// never seen in a public log) is the account's own data and is deleted outright. Every other
// character keeps what public sources say about it (the Battle.net public profile, which ages
// out within thirty days anyway, and the public logs behind the rankings).
func eraseCharacters(ctx context.Context, tx pgx.Tx, userID int64) error {
	return run(ctx, tx, []step{
		{"addon sync history", `delete from character_syncs
			where source = 'addon' and character_key in (select key from characters where user_id = $1)`},
		{"addon-only characters", `delete from characters c
			where c.user_id = $1 and c.source = 'export' and c.bnet_character_id is null
			  and not exists (select 1 from fight_metrics m where m.player_key = c.key)`},
	}, userID)
}

// eraseSignups answers the raid sign-ups, whose account keys are on delete restrict because
// a raid's roster groups by them. The account's own sign-ups go: they are its answers, and it
// leaves every guild with this deletion. A sign-up an officer wrote for someone else moves
// its acted_by to the account it answered for. (A raid the account created stays, without a
// creator: migration 0083.)
func eraseSignups(ctx context.Context, tx pgx.Tx, userID int64) error {
	return run(ctx, tx, []step{
		{"own sign-ups", `delete from guild_raid_signups where signed_by = $1`},
		{"sign-ups written for others", `update guild_raid_signups set acted_by = signed_by where acted_by = $1`},
	}, userID)
}

// eraseAccountRow deletes the email's leftover sign-in links (keyed by address, not account)
// and the users row, which takes the BattleTag, email, sealed Battle.net token and every
// cascading table with it: sessions, companion and app devices, push and live-activity
// tokens, Discord links, notification settings, guild memberships, entitlements and the
// Stripe customer link (the customer itself stays at Stripe), and the rest.
func eraseAccountRow(ctx context.Context, tx pgx.Tx, acct Account) error {
	if acct.Email != nil {
		if _, err := tx.Exec(ctx, `delete from login_tokens where email = $1`, *acct.Email); err != nil {
			return fmt.Errorf("account: erase: sign-in links: %w", err)
		}
	}
	tag, err := tx.Exec(ctx, `delete from users where id = $1`, acct.ID)
	if err != nil {
		return fmt.Errorf("account: erase: account row: %w", err)
	}
	if tag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}

// collectStrings runs a one-text-column query.
func collectStrings(ctx context.Context, q querier, sql string, args ...any) ([]string, error) {
	rows, err := q.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowTo[string])
}
