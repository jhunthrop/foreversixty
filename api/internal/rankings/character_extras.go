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
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
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
	race := strings.TrimSpace(deref(bnetRace))
	level := newestLevel(bnetLevel, importedAt, export != nil, exportLevel, capturedAt)
	exportNewer := exportIsNewer(export != nil, capturedAt, importedAt)
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

// CharacterLevels is each of keys' level as the site knows it, every key present in the
// answer and nil when no source knows its level: the level the character page shows
// (identity's rule, newestLevel), from the Battle.net import or the stored addon export's
// level= section, whichever is newer. One query for all keys. The rater reads it once per
// fight for the rotation adherence band (api/internal/rating/adherence.go).
func CharacterLevels(ctx context.Context, q character.Querier, keys []string) (map[string]*int, error) {
	rows, err := q.Query(ctx,
		`select k.key, c.level, c.imported_at, ae.export, ae.captured_at
		 from unnest($1::text[]) as k(key)
		 left join characters c on c.key = k.key
		 left join addon_exports ae on ae.character_key = k.key`, keys)
	if err != nil {
		return nil, fmt.Errorf("rankings: read levels: %w", err)
	}
	defer rows.Close()
	out := make(map[string]*int, len(keys))
	for rows.Next() {
		var (
			key                    string
			bnetLevel              *int
			export                 *string
			importedAt, capturedAt *time.Time
		)
		if err := rows.Scan(&key, &bnetLevel, &importedAt, &export, &capturedAt); err != nil {
			return nil, fmt.Errorf("rankings: read levels: %w", err)
		}
		var exportLevel *int
		if export != nil {
			if decoded, ok := fs1.Decode(*export); ok && decoded.HasLevel {
				exportLevel = &decoded.Level
			}
		}
		out[key] = newestLevel(bnetLevel, importedAt, export != nil, exportLevel, capturedAt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rankings: read levels: %w", err)
	}
	return out, nil
}

// newestLevel is identity's level rule, shared with CharacterLevels: the Battle.net import's
// level (characters.level; a non-positive one is none) unless the export is newer
// (exportIsNewer) or the import carries none, and then the export's level= section when it
// has one.
func newestLevel(bnetLevel *int, importedAt *time.Time, hasExport bool, exportLevel *int, capturedAt *time.Time) *int {
	if bnetLevel != nil && *bnetLevel <= 0 {
		bnetLevel = nil
	}
	level := bnetLevel
	if exportIsNewer(hasExport, capturedAt, importedAt) || level == nil {
		if exportLevel != nil {
			level = exportLevel
		}
	}
	return level
}

// exportIsNewer reports whether the stored export is newer than the Battle.net import. The
// export is newer only than a dated import: a characters row can carry a level with no
// imported_at (cmd/seedguild writes one), and its values are not known to be older, so they
// stand and the export only fills what they lack.
func exportIsNewer(hasExport bool, capturedAt, importedAt *time.Time) bool {
	return hasExport && capturedAt != nil && importedAt != nil && capturedAt.After(*importedAt)
}

// raceName is the display name of an export's race slug in the build data the export names,
// else the active build; "" when no build data is loaded or the slug is not a race
// there.
func (s *Store) raceName(dataBuild, slug string) string {
	if s.Trees == nil || slug == "" {
		return ""
	}
	build, ok := s.Trees.Build(dataBuild)
	if !ok {
		if build, ok = s.Trees.Active(); !ok {
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
// since migration 0035), for the newest ScoreHistoryDates dates. Only fights in public,
// unflagged reports count (presence.EvidenceReport), the rule ScoreV2's own fights follow.
// A fight with no snapshot (rated before 0035, or re-rated by the backfill, which does not
// roll up) is left out, so a date whose fights all lack one has no point. Empty, never nil.
func (s *Store) scoreHistory(ctx context.Context, key string) ([]ScorePoint, error) {
	rows, err := s.Pool.Query(ctx,
		`select to_char(d, 'YYYY-MM-DD'), score, tier from (
		   select distinct on ((rs.fought_at at time zone 'UTC')::date)
		          (rs.fought_at at time zone 'UTC')::date as d, rs.score, rs.tier
		     from rating_scores rs join reports r on r.id = rs.report_id
		    where rs.player_key = $1 and r.visibility = 'public' and `+presence.EvidenceReport("r")+`
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
