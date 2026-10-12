package rankings

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// CharacterGuild is a character's verified guild membership as the character page shows it:
// the guild's name and the character's rank in it (member, officer or leader). Never an id.
type CharacterGuild struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

// ScorePoint is one UTC date of a character's Forever Score history.
type ScorePoint struct {
	// Date is the UTC date, YYYY-MM-DD.
	Date  string `json:"date"`
	Score int    `json:"score"`
	Tier  string `json:"tier"`
}

// ScoreHistoryDates is how many dates the score history keeps: the newest ones.
const ScoreHistoryDates = 30

// anonymized reports whether key's owning account has asked to be anonymous: the join
// rating.Store.anonymized makes for GET .../rating, so the two endpoints hide the same
// characters. A key with no characters row, or no linked account, is never anonymized.
func (s *Store) anonymized(ctx context.Context, key string) (bool, error) {
	var anon bool
	err := s.Pool.QueryRow(ctx,
		`select u.anonymize from characters c join users u on u.id = c.user_id where c.key = $1`, key).Scan(&anon)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("rankings: anonymize check %s: %w", key, err)
	}
	return anon, nil
}

// identity fills who the character is in game, each field left null when no source knows it:
//
//   - guild: its verified guild membership (guild_characters.verified_at set), as { name,
//     rank }. An unverified membership is no guild.
//   - level and race: the newest of the two sources that carry them, field by field. The
//     sources are the stored addon export (fs1.Decode's level= section and race slug, dated
//     by addon_exports.captured_at) and the Battle.net import (characters.level and .race,
//     dated by characters.imported_at). The newer source's value wins; when it does not
//     carry the field (an export with no level= section, a race slug the build data does not
//     know) the older source's value stands. On a tie the Battle.net import wins, and so
//     does a characters row with no imported_at (its values are undated, never older). An
//     export's race slug is named by the build data (Store.Trees): "night-elf" is "Night
//     Elf", the form the Battle.net import stores.
//   - realm: character.Realm, the order the nightly data addon keys by.
func (s *Store) identity(ctx context.Context, ref *CharacterRef) error {
	var (
		bnetRealm, bnetRace, export, guildName, guildRank, unit *string
		bnetLevel                                               *int
		importedAt, capturedAt                                  *time.Time
	)
	err := s.Pool.QueryRow(ctx,
		`select c.realm_name, c.level, c.race, c.imported_at, ae.export, ae.captured_at, g.name, gc.rank,
		        (select rs.player_name from rating_scores rs where rs.player_key = k.key and rs.player_name <> ''
		          order by rs.fought_at desc limit 1)
		 from (select $1::text as key) k
		 left join characters c on c.key = k.key
		 left join addon_exports ae on ae.character_key = k.key
		 left join guild_characters gc on gc.character_key = k.key and gc.verified_at is not null
		 left join guilds g on g.id = gc.guild_id`, ref.Key).
		Scan(&bnetRealm, &bnetLevel, &bnetRace, &importedAt, &export, &capturedAt, &guildName, &guildRank, &unit)
	if err != nil {
		return fmt.Errorf("rankings: read identity %s: %w", ref.Key, err)
	}
	if guildName != nil && guildRank != nil {
		ref.Guild = &CharacterGuild{Name: *guildName, Rank: *guildRank}
	}

	var exportLevel *int
	var exportRace, whoRealm string
	if export != nil {
		if decoded, ok := fs1.Decode(*export); ok {
			if decoded.HasLevel {
				exportLevel = &decoded.Level
			}
			exportRace = s.raceName(decoded.DataBuild, decoded.RaceSlug)
			if decoded.HasWho {
				whoRealm = decoded.Realm
			}
		}
	}
	if bnetLevel != nil && *bnetLevel <= 0 {
		bnetLevel = nil
	}
	race := strings.TrimSpace(deref(bnetRace))
	level := bnetLevel
	// The export is newer only than a dated import. A characters row can carry a level with
	// no imported_at (cmd/seedguild writes one); its values are not known to be older, so
	// they stand and the export only fills what they lack.
	exportNewer := export != nil && capturedAt != nil && importedAt != nil && capturedAt.After(*importedAt)
	if exportNewer || level == nil {
		if exportLevel != nil {
			level = exportLevel
		}
	}
	if exportNewer || race == "" {
		if exportRace != "" {
			race = exportRace
		}
	}
	ref.Level = level
	ref.Race = nonEmpty(race)
	ref.Realm = nonEmpty(character.Realm(deref(bnetRealm), whoRealm, deref(unit)))
	return nil
}

// raceName is the display name of an export's race slug in the build data the export names,
// else the newest loaded build; "" when no build data is loaded or the slug is not a race
// there.
func (s *Store) raceName(dataBuild, slug string) string {
	if s.Trees == nil || slug == "" {
		return ""
	}
	build, ok := s.Trees.Build(dataBuild)
	if !ok {
		if build, ok = s.Trees.Latest(); !ok {
			return ""
		}
	}
	for _, r := range build.Races() {
		if r.Slug == slug {
			return r.Name
		}
	}
	return ""
}

// scoreHistory is the character's Forever Score over time, oldest first: one point per UTC
// date on which a rated fight landed, each the snapshot the rater rolled up onto the last of
// that date's fights (rating_scores.score and .tier, written by rating.Store.RollUpScores
// since migration 0035), for the newest ScoreHistoryDates dates. Only fights in public
// reports count, the rule ScoreV2's own fights follow. A fight with no snapshot (rated
// before 0035, or re-rated by the backfill, which does not roll up) is left out, so a date
// whose fights all lack one has no point. Empty, never nil.
func (s *Store) scoreHistory(ctx context.Context, key string) ([]ScorePoint, error) {
	rows, err := s.Pool.Query(ctx,
		`select to_char(d, 'YYYY-MM-DD'), score, tier from (
		   select distinct on ((rs.fought_at at time zone 'UTC')::date)
		          (rs.fought_at at time zone 'UTC')::date as d, rs.score, rs.tier
		     from rating_scores rs join reports r on r.id = rs.report_id
		    where rs.player_key = $1 and r.visibility = 'public'
		      and rs.score is not null and rs.tier is not null
		    order by (rs.fought_at at time zone 'UTC')::date desc, rs.fought_at desc,
		             rs.report_id desc, rs.fight_index desc
		    limit $2
		 ) last order by d`, key, ScoreHistoryDates)
	if err != nil {
		return nil, fmt.Errorf("rankings: read score history %s: %w", key, err)
	}
	defer rows.Close()
	out := []ScorePoint{}
	for rows.Next() {
		var p ScorePoint
		if err := rows.Scan(&p.Date, &p.Score, &p.Tier); err != nil {
			return nil, fmt.Errorf("rankings: scan score history %s: %w", key, err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func nonEmpty(s string) *string {
	if s = strings.TrimSpace(s); s == "" {
		return nil
	}
	return &s
}
