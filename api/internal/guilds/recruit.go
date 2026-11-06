// api/internal/guilds/recruit.go
//
// The public recruiting card (design/specs/2026-11-05-web-layout.md §7.2 Recruitment): the
// officer-edited fields on guilds (migration 0075) and the facts the site already holds about
// the guild (members, progress, the officers the guild page already names). Public, like the
// guild page itself.
package guilds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

const (
	// recruitPitchMax, recruitNightsMax and recruitNoteMax bound the officer's words (runes),
	// the same ceilings migration 0075's checks hold.
	recruitPitchMax  = 500
	recruitNightsMax = 120
	recruitNoteMax   = 80
	// recruitWantedMax is the most "what we look for" lines a card carries.
	recruitWantedMax = 12
	// recruitUpcomingDays and recruitUpcomingMax bound the planned raids the card derives its
	// nights from when the officers wrote none.
	recruitUpcomingDays = 14
	recruitUpcomingMax  = 8
)

// discordSnowflake is a Discord id: 17 to 20 digits, the 0075 check's own rule.
var discordSnowflake = regexp.MustCompile(`^[0-9]{17,20}$`)

// ErrInvalidRecruit is a recruiting PATCH whose fields break a rule; Field names which.
type ErrInvalidRecruit struct{ Field, Rule string }

func (e ErrInvalidRecruit) Error() string {
	return "guilds: invalid recruit " + e.Field + ": " + e.Rule
}

// WantedEntry is one "what we look for" line: a class, optionally one of its specs (a
// specs.json key), and a short note. SpecName is the catalogue's words for Spec, set on read.
type WantedEntry struct {
	Class    string  `json:"class"`
	Spec     *string `json:"spec"`
	SpecName *string `json:"spec_name"`
	Note     string  `json:"note"`
}

// RecruitNights is when the guild raids: the officers' own words, else nothing, and the
// guild's planned raids over the next two weeks (start times, the viewer's clock does the
// rest) while it recruits.
type RecruitNights struct {
	Text     *string     `json:"text"`
	Upcoming []time.Time `json:"upcoming"`
	Timezone *string     `json:"timezone"`
}

// RecruitOfficer is one guild master or officer as the guild page already names them.
type RecruitOfficer struct {
	Name  string  `json:"name"`
	Class *string `json:"class"`
	Spec  *string `json:"spec"`
	Rank  string  `json:"rank"`
}

// RecruitGuild is the card's guild line.
type RecruitGuild struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Region   string  `json:"region"`
	Ruleset  string  `json:"ruleset"`
	Faction  *string `json:"faction"`
	CrestURL *string `json:"crest_url"`
	Path     string  `json:"path"`
}

// RecruitCard is GET /v1/guilds/{id}/recruit.
type RecruitCard struct {
	Guild      RecruitGuild     `json:"guild"`
	Members    int              `json:"members"`
	Open       bool             `json:"open"`
	Pitch      string           `json:"pitch"`
	RaidNights RecruitNights    `json:"raid_nights"`
	Progress   TierView         `json:"progress"`
	Wanted     []WantedEntry    `json:"wanted"`
	Officers   []RecruitOfficer `json:"officers"`
	UpdatedAt  *time.Time       `json:"updated_at"`
}

// recruitFields is the guild's officer-edited recruiting columns.
type recruitFields struct {
	open          bool
	pitch, nights string
	wanted        []WantedEntry
	roleID        *string
	timezone      *string
	updatedAt     *time.Time
}

func (s *Store) recruitFields(ctx context.Context, guildID int64) (recruitFields, error) {
	var f recruitFields
	var wanted []byte
	err := s.Pool.QueryRow(ctx, `
		select recruit_open, recruit_pitch, recruit_nights, recruit_wanted, recruit_discord_role_id,
		       nullif(raid_timezone, ''), recruit_updated_at
		from guilds where id = $1`, guildID).
		Scan(&f.open, &f.pitch, &f.nights, &wanted, &f.roleID, &f.timezone, &f.updatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return f, ErrNotFound
	}
	if err != nil {
		return f, fmt.Errorf("guilds: recruit fields: %w", err)
	}
	if err := json.Unmarshal(wanted, &f.wanted); err != nil {
		return f, fmt.Errorf("guilds: recruit wanted: %w", err)
	}
	f.wanted = describeWanted(f.wanted)
	return f, nil
}

// describeWanted fills each line's SpecName from the catalogue; never nil.
func describeWanted(in []WantedEntry) []WantedEntry {
	out := make([]WantedEntry, 0, len(in))
	for _, w := range in {
		w.SpecName = nil
		if w.Spec != nil {
			if sp, ok := raidspec.ByKey(*w.Spec); ok {
				name := fullSpecName(sp)
				w.SpecName = &name
			}
		}
		out = append(out, w)
	}
	return out
}

// RecruitCard assembles the public card.
func (s *Store) RecruitCard(ctx context.Context, guildID int64) (RecruitCard, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return RecruitCard{}, err
	}
	f, err := s.recruitFields(ctx, guildID)
	if err != nil {
		return RecruitCard{}, err
	}
	card := RecruitCard{
		Guild: RecruitGuild{ID: g.ID, Name: g.Name, Region: g.Region, Ruleset: g.Ruleset, Faction: g.Faction,
			CrestURL: CrestURL(s.APIBaseURL, g.ID, g.CrestKey), Path: GuildPagePath(g)},
		Open: f.open, Pitch: f.pitch, Wanted: f.wanted, UpdatedAt: f.updatedAt,
	}
	if card.Members, err = s.verifiedAccounts(ctx, guildID); err != nil {
		return RecruitCard{}, err
	}
	if card.RaidNights, err = s.recruitNights(ctx, guildID, f); err != nil {
		return RecruitCard{}, err
	}
	prog, err := s.Progression(ctx, guildID, false)
	if err != nil {
		return RecruitCard{}, err
	}
	card.Progress = prog.Tier
	if card.Officers, err = s.publicOfficers(ctx, guildID); err != nil {
		return RecruitCard{}, err
	}
	return card, nil
}

// verifiedAccounts is the guild's member count: verified accounts, the billing and invite count.
func (s *Store) verifiedAccounts(ctx context.Context, guildID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `select count(*) from guild_members
		where guild_id = $1 and verified_at is not null`, guildID).Scan(&n); err != nil {
		return 0, fmt.Errorf("guilds: member count: %w", err)
	}
	return n, nil
}

// recruitNights is the officers' words; the planned raids are listed only while the guild
// recruits, so a closed card says nothing of when the guild is in the instance.
func (s *Store) recruitNights(ctx context.Context, guildID int64, f recruitFields) (RecruitNights, error) {
	out := RecruitNights{Upcoming: []time.Time{}, Timezone: f.timezone}
	if f.nights != "" {
		text := f.nights
		out.Text = &text
	}
	if !f.open {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `select starts_at from guild_raids
		where guild_id = $1 and status in ('planned', 'locked')
		  and starts_at > now() and starts_at < now() + make_interval(days => $2)
		order by starts_at limit $3`, guildID, recruitUpcomingDays, recruitUpcomingMax)
	if err != nil {
		return out, fmt.Errorf("guilds: recruit nights: %w", err)
	}
	starts, err := pgx.CollectRows(rows, pgx.RowTo[time.Time])
	if err != nil {
		return out, fmt.Errorf("guilds: recruit nights: %w", err)
	}
	for _, t := range starts {
		out.Upcoming = append(out.Upcoming, t.UTC())
	}
	return out, nil
}

// publicOfficers is the guild master and officers the public guild page already names: a
// verified leader or officer character with a boss kill in one of the guild's public, unflagged
// reports (the page's Roster bests), its account not anonymised; class and spec from that kill.
// One character per account, guild master first, then by name.
func (s *Store) publicOfficers(ctx context.Context, guildID int64) ([]RecruitOfficer, error) {
	rows, err := s.Pool.Query(ctx, `
		select name, class, spec, rank from (
		  select distinct on (gc.user_id) fm.player_name as name, nullif(fm.class, '') as class,
		         nullif(fm.spec, '') as spec, gc.rank
		  from guild_characters gc
		  join users u on u.id = gc.user_id
		  join lateral (
		    select m.player_name, m.class, m.spec from fight_metrics m join reports r on r.id = m.report_id
		    where m.player_key = gc.character_key and r.guild_id = $1 and m.kill and m.state <> 'removed'
		      and `+presence.EvidenceReport("r")+`
		    order by m.fought_at desc limit 1
		  ) fm on true
		  where gc.guild_id = $1 and gc.verified_at is not null and gc.rank in ('leader', 'officer')
		    and gc.character_key not like 'account:%' and not u.anonymize
		  order by gc.user_id, `+rankOrderSQL("gc")+`,
		           (gc.character_key = u.main_character_key) desc nulls last, fm.player_name
		) o
		order by case rank when 'leader' then 0 else 1 end, lower(name)`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: recruit officers: %w", err)
	}
	out, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (RecruitOfficer, error) {
		var o RecruitOfficer
		err := row.Scan(&o.Name, &o.Class, &o.Spec, &o.Rank)
		return o, err
	})
	if err != nil {
		return nil, fmt.Errorf("guilds: recruit officers: %w", err)
	}
	return out, nil
}

// RecruitPatch is an officer's edit: a nil field is left as it is. DiscordRoleID "" clears it.
type RecruitPatch struct {
	Open          *bool          `json:"open"`
	Pitch         *string        `json:"pitch"`
	RaidNights    *string        `json:"raid_nights"`
	Wanted        *[]WantedEntry `json:"wanted"`
	DiscordRoleID *string        `json:"discord_role_id"`
}

// validate trims the words and checks every rule, failing on the first field that breaks one.
func (p RecruitPatch) validate() (RecruitPatch, error) {
	out := p
	if p.Pitch != nil {
		v := strings.TrimSpace(*p.Pitch)
		if utf8.RuneCountInString(v) > recruitPitchMax {
			return p, ErrInvalidRecruit{"pitch", fmt.Sprintf("at most %d characters", recruitPitchMax)}
		}
		out.Pitch = &v
	}
	if p.RaidNights != nil {
		v := strings.TrimSpace(*p.RaidNights)
		if utf8.RuneCountInString(v) > recruitNightsMax {
			return p, ErrInvalidRecruit{"raid_nights", fmt.Sprintf("at most %d characters", recruitNightsMax)}
		}
		out.RaidNights = &v
	}
	if p.DiscordRoleID != nil {
		v := strings.TrimSpace(*p.DiscordRoleID)
		if v != "" && !discordSnowflake.MatchString(v) {
			return p, ErrInvalidRecruit{"discord_role_id", "a Discord role id, 17 to 20 digits"}
		}
		out.DiscordRoleID = &v
	}
	if p.Wanted != nil {
		wanted, err := validWanted(*p.Wanted)
		if err != nil {
			return p, err
		}
		out.Wanted = &wanted
	}
	return out, nil
}

// validWanted checks each line: a known class, a spec of that class when one is named, a short
// note; at most recruitWantedMax lines. SpecName is never stored.
func validWanted(in []WantedEntry) ([]WantedEntry, error) {
	if len(in) > recruitWantedMax {
		return nil, ErrInvalidRecruit{"wanted", fmt.Sprintf("at most %d lines", recruitWantedMax)}
	}
	out := make([]WantedEntry, 0, len(in))
	for _, w := range in {
		class := strings.ToLower(strings.TrimSpace(w.Class))
		if raidspec.ForClass(class) == nil {
			return nil, ErrInvalidRecruit{"wanted", "class must be one of " + strings.Join(raidspec.ClassOrder, ", ")}
		}
		entry := WantedEntry{Class: class, Note: strings.TrimSpace(w.Note)}
		if w.Spec != nil && strings.TrimSpace(*w.Spec) != "" {
			sp, ok := raidspec.ByKey(strings.TrimSpace(*w.Spec))
			if !ok || sp.Class != class {
				return nil, ErrInvalidRecruit{"wanted", "spec must be a spec of its class, like paladin-holy"}
			}
			entry.Spec = &sp.Key
		}
		if utf8.RuneCountInString(entry.Note) > recruitNoteMax {
			return nil, ErrInvalidRecruit{"wanted", fmt.Sprintf("a note is at most %d characters", recruitNoteMax)}
		}
		out = append(out, entry)
	}
	return out, nil
}

// UpdateRecruit applies a validated patch in one statement and stamps recruit_updated_at.
func (s *Store) UpdateRecruit(ctx context.Context, guildID int64, in RecruitPatch) error {
	p, err := in.validate()
	if err != nil {
		return err
	}
	var wanted []byte
	if p.Wanted != nil {
		if wanted, err = json.Marshal(*p.Wanted); err != nil {
			return fmt.Errorf("guilds: recruit wanted: %w", err)
		}
	}
	tag, err := s.Pool.Exec(ctx, `
		update guilds set
		  recruit_open = coalesce($2, recruit_open),
		  recruit_pitch = coalesce($3, recruit_pitch),
		  recruit_nights = coalesce($4, recruit_nights),
		  recruit_wanted = coalesce($5::jsonb, recruit_wanted),
		  recruit_discord_role_id = case when $6::text is null then recruit_discord_role_id
		                                 else nullif($6::text, '') end,
		  recruit_updated_at = now()
		where id = $1`, guildID, p.Open, p.Pitch, p.RaidNights, wanted, p.DiscordRoleID)
	if err != nil {
		return fmt.Errorf("guilds: update recruit: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// RecruitRoleID is the Discord role an accepted applicant is granted; nil when none is set.
func (s *Store) RecruitRoleID(ctx context.Context, guildID int64) (*string, error) {
	f, err := s.recruitFields(ctx, guildID)
	return f.roleID, err
}
