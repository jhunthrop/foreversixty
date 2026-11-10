// api/internal/guilds/applications.go
//
// Guild applications (design/specs/2026-11-05-web-layout.md §7.2 Applications; migration 0075):
// a signed-in player applies with one of their own characters while the guild recruits, reads
// and withdraws their own, and officers list them with the character's facts and the votes.
// Accepting one is applications_decide.go.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

const (
	// applicationMessageMax and applicationAvailabilityMax bound the applicant's words (runes),
	// migration 0075's checks.
	applicationMessageMax      = 1000
	applicationAvailabilityMax = 200
	// applicationDecidedShown is how many decided applications the officer list carries after
	// every pending one.
	applicationDecidedShown = 20

	// The application states.
	ApplicationPending   = "pending"
	ApplicationAccepted  = "accepted"
	ApplicationDeclined  = "declined"
	ApplicationWithdrawn = "withdrawn"
)

// The refusals Apply, Withdraw and the decisions answer.
var (
	ErrNotRecruiting      = errors.New("guilds: guild is not taking applications")
	ErrAlreadyApplied     = errors.New("guilds: an application is already open")
	ErrAlreadyMember      = errors.New("guilds: already a member")
	ErrNotApplicantsChar  = errors.New("guilds: not the applicant's character")
	ErrWrongRealm         = errors.New("guilds: character is on another region or ruleset")
	ErrApplicationClosed  = errors.New("guilds: application is no longer pending")
	ErrInvalidApplication = errors.New("guilds: invalid application")
)

// ApplicationInput is what an applicant sends.
type ApplicationInput struct {
	CharacterKey string `json:"character_key"`
	Message      string `json:"message"`
	Availability string `json:"availability"`
	// DiscordDM is "Message me on Discord about this application", off by default.
	DiscordDM bool `json:"discord_dm"`
}

func (in ApplicationInput) normalized() (ApplicationInput, error) {
	out := ApplicationInput{CharacterKey: strings.TrimSpace(in.CharacterKey),
		Message: strings.TrimSpace(in.Message), Availability: strings.TrimSpace(in.Availability), DiscordDM: in.DiscordDM}
	if out.CharacterKey == "" || utf8.RuneCountInString(out.Message) > applicationMessageMax ||
		utf8.RuneCountInString(out.Availability) > applicationAvailabilityMax {
		return out, ErrInvalidApplication
	}
	return out, nil
}

// ApplicantCharacter is the character an application names, as the site knows it.
type ApplicantCharacter struct {
	CharacterKey string  `json:"character_key"`
	Name         string  `json:"name"`
	Class        *string `json:"class"`
	Spec         *string `json:"spec"`
	SpecName     *string `json:"spec_name"`
	// Group is the raid group the spec fights in (raidspec.Group: tank, healer, melee, ranged),
	// so the officer reads which open seat an applicant fills; null when the spec is unknown.
	Group     *string    `json:"group"`
	Level     *int       `json:"level"`
	ItemLevel *int       `json:"item_level"`
	BestParse *BestParse `json:"best_parse"`
	// Race and Faction are the character's from Battle.net, else its export's race; null when
	// neither says.
	Race    *string `json:"race"`
	Faction *string `json:"faction"`
}

// Application is one application as its applicant reads it.
type Application struct {
	ID           int64              `json:"id"`
	GuildID      int64              `json:"guild_id"`
	Status       string             `json:"status"`
	CreatedAt    time.Time          `json:"created_at"`
	DecidedAt    *time.Time         `json:"decided_at"`
	Message      string             `json:"message"`
	Availability string             `json:"availability"`
	DiscordDM    bool               `json:"discord_dm"`
	Character    ApplicantCharacter `json:"character"`
	userID       int64
	decidedBy    *int64
	roleOutcome  *string
	dmOutcome    *string
}

// Apply files userID's application to guildID.
func (s *Store) Apply(ctx context.Context, guildID, userID int64, raw ApplicationInput) (Application, error) {
	in, err := raw.normalized()
	if err != nil {
		return Application{}, err
	}
	if err := s.applicable(ctx, guildID, userID, in.CharacterKey); err != nil {
		return Application{}, err
	}
	var id int64
	err = s.Pool.QueryRow(ctx, `
		insert into guild_applications (guild_id, user_id, character_key, message, availability, discord_dm)
		values ($1, $2, $3, $4, $5, $6) returning id`,
		guildID, userID, in.CharacterKey, in.Message, in.Availability, in.DiscordDM).Scan(&id)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
		return Application{}, ErrAlreadyApplied
	}
	if err != nil {
		return Application{}, fmt.Errorf("guilds: apply: %w", err)
	}
	return s.application(ctx, guildID, id)
}

// applicable checks the guild recruits, the character is userID's own on the guild's region
// and ruleset, and the account is not already a verified member.
func (s *Store) applicable(ctx context.Context, guildID, userID int64, characterKey string) error {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return err
	}
	f, err := s.recruitFields(ctx, guildID)
	if err != nil {
		return err
	}
	if !f.open {
		return ErrNotRecruiting
	}
	var region, ruleset string
	err = s.Pool.QueryRow(ctx, `select region, ruleset from characters where key = $1 and user_id = $2`,
		characterKey, userID).Scan(&region, &ruleset)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotApplicantsChar
	}
	if err != nil {
		return fmt.Errorf("guilds: apply: read character: %w", err)
	}
	if region != g.Region || ruleset != g.Ruleset {
		return ErrWrongRealm
	}
	member, err := s.IsVerifiedMember(ctx, guildID, userID)
	if err != nil {
		return err
	}
	if member {
		return ErrAlreadyMember
	}
	return nil
}

const applicationColumns = `id, guild_id, user_id, status, created_at, decided_at, message, availability,
	discord_dm, character_key, decided_by, role_outcome, dm_outcome`

func scanApplication(row pgx.CollectableRow) (Application, error) {
	var a Application
	err := row.Scan(&a.ID, &a.GuildID, &a.userID, &a.Status, &a.CreatedAt, &a.DecidedAt, &a.Message,
		&a.Availability, &a.DiscordDM, &a.Character.CharacterKey, &a.decidedBy, &a.roleOutcome, &a.dmOutcome)
	return a, err
}

// applications runs query (selecting applicationColumns) and fills each row's character.
func (s *Store) applications(ctx context.Context, query string, args ...any) ([]Application, error) {
	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("guilds: applications: %w", err)
	}
	apps, err := pgx.CollectRows(rows, scanApplication)
	if err != nil {
		return nil, fmt.Errorf("guilds: applications: %w", err)
	}
	keys := make([]string, 0, len(apps))
	for _, a := range apps {
		keys = append(keys, a.Character.CharacterKey)
	}
	chars, err := s.applicantCharacters(ctx, keys)
	if err != nil {
		return nil, err
	}
	for i := range apps {
		if c, ok := chars[apps[i].Character.CharacterKey]; ok {
			apps[i].Character = c
		}
	}
	return apps, nil
}

// application is one application of guildID; ErrNotFound for another guild's or none.
func (s *Store) application(ctx context.Context, guildID, id int64) (Application, error) {
	apps, err := s.applications(ctx, `select `+applicationColumns+` from guild_applications
		where guild_id = $1 and id = $2`, guildID, id)
	if err != nil {
		return Application{}, err
	}
	if len(apps) == 0 {
		return Application{}, ErrNotFound
	}
	return apps[0], nil
}

// MyApplications is userID's applications to guildID, newest first; never nil.
func (s *Store) MyApplications(ctx context.Context, guildID, userID int64) ([]Application, error) {
	apps, err := s.applications(ctx, `select `+applicationColumns+` from guild_applications
		where guild_id = $1 and user_id = $2 order by created_at desc, id desc`, guildID, userID)
	if apps == nil && err == nil {
		apps = []Application{}
	}
	return apps, err
}

// Withdraw takes back userID's own pending application. ErrNotFound when it is not theirs,
// ErrApplicationClosed when it is decided or withdrawn already.
func (s *Store) Withdraw(ctx context.Context, guildID, userID, id int64) (Application, error) {
	a, err := s.application(ctx, guildID, id)
	if err != nil {
		return Application{}, err
	}
	if a.userID != userID {
		return Application{}, ErrNotFound
	}
	tag, err := s.Pool.Exec(ctx, `update guild_applications
		set status = 'withdrawn', decided_at = now(), decided_by = user_id
		where id = $1 and status = 'pending'`, id)
	if err != nil {
		return Application{}, fmt.Errorf("guilds: withdraw: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return Application{}, ErrApplicationClosed
	}
	return s.application(ctx, guildID, id)
}

// PendingApplicationsCount is how many applications to guildID wait for a decision: the
// officer Action needed row's count.
func (s *Store) PendingApplicationsCount(ctx context.Context, guildID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `select count(*) from guild_applications
		where guild_id = $1 and status = 'pending'`, guildID).Scan(&n); err != nil {
		return 0, fmt.Errorf("guilds: pending applications: %w", err)
	}
	return n, nil
}

// applicantCharacters is each key's class, spec, level, item level and best boss parse. Class
// and level from the character row, spec and the fallback item level from its newest ranked
// fight in a public, unflagged report, the item level the Battle.net equipped figure first. The
// best parse is the roster's own rule (bestParses) over every report, since an applicant has
// none of the guild's own yet.
func (s *Store) applicantCharacters(ctx context.Context, keys []string) (map[string]ApplicantCharacter, error) {
	out := map[string]ApplicantCharacter{}
	if len(keys) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `
		select c.key, c.name, coalesce(nullif(c.class, ''), fm.class), c.level,
		       coalesce(c.equipped_item_level, fm.ilvl), nullif(fm.spec, ''), coalesce(fm.role, ''),
		       nullif(c.race, ''), nullif(c.faction, ''), ae.export
		from characters c
		left join lateral (`+newestEvidenceFight("c.key")+`) fm on true
		left join addon_exports ae on ae.character_key = c.key
		where c.key = any($1)`, keys)
	if err != nil {
		return nil, fmt.Errorf("guilds: applicant characters: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var c ApplicantCharacter
		var role string
		var export *string
		if err := rows.Scan(&c.CharacterKey, &c.Name, &c.Class, &c.Level, &c.ItemLevel, &c.Spec, &role,
			&c.Race, &c.Faction, &export); err != nil {
			return nil, fmt.Errorf("guilds: applicant characters: %w", err)
		}
		s.raceFromExport(&c, export)
		if sp, ok := resolveSpec(c.Class, c.Spec, role); ok {
			c.SpecName, c.Group = &sp.Name, strp(sp.Group)
		}
		out[c.CharacterKey] = c
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return s.withBestParses(ctx, keys, out)
}

// raceFromExport fills a missing race and faction from the export's race.
func (s *Store) raceFromExport(c *ApplicantCharacter, export *string) {
	if export == nil || (c.Race != nil && c.Faction != nil) {
		return
	}
	d, ok := fs1.Decode(*export)
	if !ok || d.RaceSlug == "" {
		return
	}
	if c.Race == nil {
		race := raceWords(d.RaceSlug)
		c.Race = &race
	}
	if c.Faction == nil {
		if f, ok := s.factionForRace(d.RaceSlug); ok {
			c.Faction = &f
		}
	}
}

// raceWords is a race slug as a name: "night-elf" reads "Night Elf".
func raceWords(slug string) string {
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if p != "" {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

// newestEvidenceFight is the newest ranked fight of the character keyRef names in a public,
// unflagged report: its class, spec, role and item level.
func newestEvidenceFight(keyRef string) string {
	return `select m.class, m.spec, m.role, m.ilvl from fight_metrics m join reports r on r.id = m.report_id
		where m.player_key = ` + keyRef + ` and m.state <> 'removed' and ` + presence.EvidenceReport("r") + `
		order by m.fought_at desc limit 1`
}

// withBestParses sets each character's best parse.
func (s *Store) withBestParses(ctx context.Context, keys []string, chars map[string]ApplicantCharacter) (map[string]ApplicantCharacter, error) {
	best, err := s.bestParsesWhere(ctx, "m.player_key = any($1)", keys)
	if err != nil {
		return nil, err
	}
	for key, bp := range best {
		if c, ok := chars[key]; ok {
			b := bp
			c.BestParse = &b
			chars[key] = c
		}
	}
	return chars, nil
}
