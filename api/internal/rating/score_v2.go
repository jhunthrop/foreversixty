package rating

// The Forever Score, version 2: one number for a character, built from five components
// (docs/superpowers/specs/2026-10-11-addon-v2-design.md §4.6; the nightly record that
// carries it is docs/contracts/addon-v2.md contract 5). This file is the formula's only
// home: the addon and the site display what it computes and never compute it themselves.
//
// Every constant the formula uses is named here, in one place, because the site explains
// each one in plain words and the thresholds are tuned from the live distribution. The
// ranges are the v1 definition:
//
//	gear        0-500  the character's gear score for its spec, against that spec's own
//	                   level-60 best-in-slot set (gearScoreAtFull)
//	experience  0-700  25 a distinct dungeon cleared + 40 a distinct raid boss killed, by
//	                   any character the account has linked, capped
//	parses      0-600  the mean of the character's best parse percentile on each boss
//	attendance  0-200  the share of its guild's nights in the last 90 days it was at
//	discipline  0-100  the share of its rated fights pulled with full consumables and
//	                   every enchantable slot enchanted
//
// score is their sum; tier names the band the score falls in. A character with no rated
// fight has no score at all (ErrNoLogs), never a zero.

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The component ceilings and rates (v1).
const (
	// GearMax is the gear component at a full best-in-slot set.
	GearMax = 500
	// ExperiencePerDungeon is earned once per distinct dungeon instance cleared.
	ExperiencePerDungeon = 25
	// ExperiencePerRaidBoss is earned once per distinct raid boss killed.
	ExperiencePerRaidBoss = 40
	// ExperienceMax caps the experience component.
	ExperienceMax = 700
	// ParsesMax is the parses component at a 100th-percentile mean.
	ParsesMax = 600
	// AttendanceMax is the attendance component at every guild night attended.
	AttendanceMax = 200
	// AttendanceWindow is how far back guild nights count.
	AttendanceWindow = 90 * 24 * time.Hour
	// DisciplineMax is the discipline component when every fight was fully prepared.
	DisciplineMax = 100
)

// The tiers, named for the client's own item qualities so the addon colours them with
// the client's quality colours. Provisional: set from the live distribution before week 2
// of the addon v2 programme closes.
const (
	TierCommon    = "common"
	TierUncommon  = "uncommon"
	TierRare      = "rare"
	TierEpic      = "epic"
	TierLegendary = "legendary"

	// UncommonFrom, RareFrom, EpicFrom and LegendaryFrom are the lowest score of each
	// tier; anything below UncommonFrom is common.
	UncommonFrom  = 500
	RareFrom      = 1000
	EpicFrom      = 1500
	LegendaryFrom = 2000
)

// Tier names the band score falls in.
func Tier(score int) string {
	switch {
	case score >= LegendaryFrom:
		return TierLegendary
	case score >= EpicFrom:
		return TierEpic
	case score >= RareFrom:
		return TierRare
	case score >= UncommonFrom:
		return TierUncommon
	default:
		return TierCommon
	}
}

// gearScoreAtFull is each spec's sims.GearScore of its own level-60 best-in-slot set: the
// highest of the band-60 sets in data/builds/1.60.1.70291/bis/<spec>.json, scored with that
// build's stat-weights.json, rounded. A spec's weights are in its own units (a healer's
// full set scores about 400, a rogue's about 5,700), so gear is measured against the
// spec's own best rather than one number for everyone. A spec missing here scores no gear
// component. Re-derive these when the weights or the best-in-slot lists change.
var gearScoreAtFull = map[string]float64{
	"druid-balance": 1150, "druid-feral": 3051, "druid-feral-bear": 1006, "druid-restoration": 428,
	"hunter-beast-mastery": 4476, "hunter-marksmanship": 4798, "hunter-survival": 5771,
	"mage-arcane": 1474, "mage-fire": 1766, "mage-frost": 1420,
	"paladin-holy": 393, "paladin-protection": 1347, "paladin-retribution": 4922,
	"priest-discipline": 514, "priest-holy": 518, "priest-shadow": 1209,
	"rogue-assassination": 5733, "rogue-combat": 5494, "rogue-subtlety": 5724,
	"shaman-elemental": 1309, "shaman-enhancement": 4872, "shaman-restoration": 490,
	"warlock-affliction": 1092, "warlock-demonology": 1256, "warlock-destruction": 1348,
	"warrior-arms": 4838, "warrior-fury": 5523, "warrior-protection": 1661,
}

// disciplineEnchantSlots are the gear slots a fully prepared player has enchanted, as
// indexes into COMBATANT_INFO's gear list (the client's inventory slot, less one): chest,
// feet, back and main hand - the four the guild readiness board checks
// (api/internal/guilds/readiness_core.go's enchantableSlots).
var disciplineEnchantSlots = []int{4, 7, 14, 15}

// ErrNoLogs is the answer for a character with no rated fight in a public report. It is
// never a zero score: the addon shows "no logs yet" and the site shows no score.
var ErrNoLogs = errors.New("rating: no logs")

// ScoreComponents are the score's five parts; they sum to the score.
type ScoreComponents struct {
	Gear, Experience, Parses, Attendance, Discipline int
}

// ScoreExperience is what the experience component counts. Dungeons and RaidKills are
// distinct instances and bosses across every linked character; Alts is how many linked
// characters there are besides this one.
type ScoreExperience struct {
	Dungeons, RaidKills, Alts int
}

// ScoreV2 is one character's Forever Score.
type ScoreV2 struct {
	Score int
	Tier  string
	// AsOf is the newest rated fight's time.
	AsOf       time.Time
	Components ScoreComponents
	// GearScore is the raw sims.GearScore of the character's gear, rounded; ItemLevel its
	// equipped item level (0 when unknown).
	GearScore int
	ItemLevel int
	// Spec is the newest rated fight's spec slug ("warrior-arms"), the one gear is scored
	// for; Fights is how many rated fights the score rests on.
	Spec       string
	Fights     int
	Experience ScoreExperience
}

// ScoreV2 computes characterKey's Forever Score from what is stored now. Only fights in
// public reports count toward the score's own fights (the rating endpoints' rule, spec
// docs/superpowers/specs/2026-09-21-performance-rating-design.md §5.1); experience and
// attendance count every report that is not private, since they publish only counts.
func (s *Store) ScoreV2(ctx context.Context, characterKey string) (ScoreV2, error) {
	fights, err := s.ratedFights(ctx, characterKey)
	if err != nil {
		return ScoreV2{}, err
	}
	if fights.count == 0 {
		return ScoreV2{}, ErrNoLogs
	}
	out := ScoreV2{AsOf: fights.newest, Spec: fights.spec, Fights: fights.count}

	gear, err := s.gearScore(ctx, characterKey, fights.spec)
	if err != nil {
		return ScoreV2{}, err
	}
	out.GearScore = int(math.Round(gear))
	if full := gearScoreAtFull[fights.spec]; full > 0 {
		out.Components.Gear = points(GearMax*gear/full, GearMax)
	}
	if out.ItemLevel, err = s.itemLevel(ctx, characterKey); err != nil {
		return ScoreV2{}, err
	}

	if out.Experience, err = s.experience(ctx, characterKey); err != nil {
		return ScoreV2{}, err
	}
	out.Components.Experience = points(float64(ExperiencePerDungeon*out.Experience.Dungeons+
		ExperiencePerRaidBoss*out.Experience.RaidKills), ExperienceMax)

	meanBest, err := s.meanBestParse(ctx, characterKey)
	if err != nil {
		return ScoreV2{}, err
	}
	out.Components.Parses = points(meanBest/100*ParsesMax, ParsesMax)

	nights, attended, err := s.guildNights(ctx, characterKey, time.Now().Add(-AttendanceWindow))
	if err != nil {
		return ScoreV2{}, err
	}
	if nights > 0 {
		out.Components.Attendance = points(AttendanceMax*float64(attended)/float64(nights), AttendanceMax)
	}

	if fights.disciplineKnown > 0 {
		out.Components.Discipline = points(DisciplineMax*float64(fights.disciplineFull)/float64(fights.disciplineKnown),
			DisciplineMax)
	}

	c := out.Components
	out.Score = c.Gear + c.Experience + c.Parses + c.Attendance + c.Discipline
	out.Tier = Tier(out.Score)
	return out, nil
}

// points rounds a component's raw value onto [0, max].
func points(v float64, max int) int {
	return int(math.Round(math.Max(0, math.Min(v, float64(max)))))
}

// ratedFightSummary is what the score reads off a character's rating_scores rows.
type ratedFightSummary struct {
	count  int
	newest time.Time
	spec   string
	// disciplineKnown counts fights whose discipline is recorded (rows written before
	// migration 0035 have none); disciplineFull counts those that were fully prepared.
	disciplineKnown, disciplineFull int
}

func (s *Store) ratedFights(ctx context.Context, key string) (ratedFightSummary, error) {
	var out ratedFightSummary
	var newest *time.Time
	var class, spec *string
	err := s.Pool.QueryRow(ctx,
		`select count(*), max(rs.fought_at), count(rs.discipline), count(*) filter (where rs.discipline),
		        (array_agg(rs.class order by rs.fought_at desc) filter (where coalesce(rs.spec, '') <> ''))[1],
		        (array_agg(rs.spec order by rs.fought_at desc) filter (where coalesce(rs.spec, '') <> ''))[1]
		 from rating_scores rs join reports r on r.id = rs.report_id
		 where rs.player_key = $1 and r.visibility = 'public'`, key).
		Scan(&out.count, &newest, &out.disciplineKnown, &out.disciplineFull, &class, &spec)
	if err != nil {
		return ratedFightSummary{}, fmt.Errorf("rating: rated fights of %s: %w", key, err)
	}
	if newest != nil {
		out.newest = *newest
	}
	if class != nil && spec != nil {
		out.spec = specSlug(*class, *spec)
	}
	return out, nil
}

// gearScore is sims.GearScore of the character's stored addon export under spec. No
// export, no spec, or gear the loaded data cannot score is no gear score (0) rather than
// an error: the other four components still stand.
func (s *Store) gearScore(ctx context.Context, key, spec string) (float64, error) {
	if spec == "" {
		return 0, nil
	}
	var export string
	err := s.Pool.QueryRow(ctx, `select export from addon_exports where character_key = $1`, key).Scan(&export)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("rating: export of %s: %w", key, err)
	}
	score, err := sims.GearScore(ctx, export, spec)
	if errors.Is(err, sims.ErrGearUnscorable) {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("rating: gear score of %s: %w", key, err)
	}
	return score, nil
}

// itemLevel is the Battle.net capture's equipped item level, else the newest ranked
// fight's, else 0.
func (s *Store) itemLevel(ctx context.Context, key string) (int, error) {
	var ilvl int
	err := s.Pool.QueryRow(ctx,
		`select coalesce(
		   (select equipped_item_level from characters where key = $1),
		   (select ilvl from fight_metrics where player_key = $1 and ilvl is not null
		     order by fought_at desc limit 1),
		   0)`, key).Scan(&ilvl)
	if err != nil {
		return 0, fmt.Errorf("rating: item level of %s: %w", key, err)
	}
	return ilvl, nil
}

// experience counts across the account's linked characters (character.LinkedCharacters,
// the characters GET /v1/me lists); a character no account owns counts alone.
//
// Dungeons is always 0 today: nothing stored says which logged instance is a dungeon (the
// fights and reports tables carry an encounter id and a free-text zone, and the build
// data's dungeons.json is still empty), and no dungeon logs exist yet. The fights table,
// classified by reports.zone against that dungeon list, is what will carry it. Until
// then every killed encounter counts as a raid boss.
func (s *Store) experience(ctx context.Context, key string) (ScoreExperience, error) {
	var out ScoreExperience
	var userID *int64
	err := s.Pool.QueryRow(ctx, `select user_id from characters where key = $1`, key).Scan(&userID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return out, fmt.Errorf("rating: owner of %s: %w", key, err)
	}
	keys := []string{key}
	if userID != nil {
		linked, err := character.LinkedCharacters(ctx, s.Pool, *userID)
		if err != nil {
			return out, err
		}
		for _, l := range linked {
			if l.Key != key {
				keys = append(keys, l.Key)
				out.Alts++
			}
		}
	}
	// Kills are read from the two tables keyed by character (fight_metrics, every ranked
	// fight; rating_scores, every rated one), not fights.players, which holds log GUIDs.
	err = s.Pool.QueryRow(ctx,
		`select count(distinct encounter_id) from (
		   select m.encounter_id from fight_metrics m join reports r on r.id = m.report_id
		    where m.player_key = any($1) and m.kill and m.state <> 'removed' and r.visibility <> 'private'
		   union
		   select rs.encounter_id from rating_scores rs join reports r on r.id = rs.report_id
		    where rs.player_key = any($1) and rs.kill and r.visibility <> 'private'
		 ) kills where coalesce(encounter_id, 0) <> 0`, keys).Scan(&out.RaidKills)
	if err != nil {
		return out, fmt.Errorf("rating: raid kills of %s: %w", key, err)
	}
	return out, nil
}

// meanBestParse is the mean of the character's best parse percentile per boss, as the
// rankings character page places them (rankings.Store.Character's best[]); 0 with none.
func (s *Store) meanBestParse(ctx context.Context, key string) (float64, error) {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 {
		return 0, nil
	}
	page, _, err := (&rankings.Store{Pool: s.Pool}).Character(ctx, parts[0], parts[1], parts[2])
	if err != nil {
		return 0, err
	}
	var sum float64
	var n int
	for _, b := range page.Best {
		if b.Percentile != nil {
			sum += *b.Percentile
			n++
		}
	}
	if n == 0 {
		return 0, nil
	}
	return sum / float64(n), nil
}

// guildNights counts the nights of the character's verified guild since since (a night
// is a UTC date on which the guild filed a report that is not private) and how many of
// them the character was ranked or rated in. No verified guild is no nights.
func (s *Store) guildNights(ctx context.Context, key string, since time.Time) (nights, attended int, err error) {
	err = s.Pool.QueryRow(ctx,
		`with nights as (
		   select r.id, (r.created_at at time zone 'UTC')::date as night
		   from reports r
		   join guild_characters gc on gc.guild_id = r.guild_id
		   where gc.character_key = $1 and gc.verified_at is not null
		     and r.visibility <> 'private' and r.created_at >= $2
		 )
		 select (select count(distinct night) from nights),
		        (select count(distinct n.night) from nights n
		          where exists (select 1 from rating_scores rs where rs.report_id = n.id and rs.player_key = $1)
		             or exists (select 1 from fight_metrics m
		                         where m.report_id = n.id and m.player_key = $1 and m.state <> 'removed'))`,
		key, since).Scan(&nights, &attended)
	if err != nil {
		return 0, 0, fmt.Errorf("rating: guild nights of %s: %w", key, err)
	}
	return nights, attended, nil
}

// fightDiscipline is whether row's player pulled fully prepared: every consumable
// category its role's catalogue lists for flask, food and weapon enchant present, and
// every equipped slot of disciplineEnchantSlots enchanted. nil is "not known" - no
// catalogue for the role, or no combatant info with gear for the player.
func fightDiscipline(fight summary.Summary, row summary.RosterRow, cat *consumables.RoleCatalogue) *bool {
	if cat == nil {
		return nil
	}
	var combatant *summary.CombatantRow
	for i := range fight.Combatants {
		if fight.Combatants[i].GUID == row.GUID {
			combatant = &fight.Combatants[i]
			break
		}
	}
	if combatant == nil || len(combatant.Gear) == 0 {
		return nil
	}
	full := carries(*combatant, cat.Flask) && carries(*combatant, cat.Food) &&
		carries(*combatant, cat.WeaponEnchant) && enchanted(*combatant)
	return &full
}

// carries is true when the combatant has any one of entries, or entries is empty (a
// category the role is not asked for).
func carries(c summary.CombatantRow, entries []consumables.Entry) bool {
	if len(entries) == 0 {
		return true
	}
	for _, e := range entries {
		for _, a := range c.Consumables {
			if a.SpellID == e.SpellID {
				return true
			}
		}
	}
	return false
}

// enchanted is true when every equipped slot of disciplineEnchantSlots carries a
// permanent enchant (the first of an item's enchant ids). An empty slot asks for none.
func enchanted(c summary.CombatantRow) bool {
	for _, slot := range disciplineEnchantSlots {
		if slot >= len(c.Gear) || c.Gear[slot].ID == 0 {
			continue
		}
		if e := c.Gear[slot].Enchants; len(e) == 0 || e[0] == 0 {
			return false
		}
	}
	return true
}

// RollUpScores writes each roster player's Forever Score, as it stands now, onto that
// player's row of the fight f just rated (rating_scores' gear_score, experience,
// attendance and tier). A player with no score yet (ErrNoLogs: the fight was not public,
// say) is left NULL. Every player is tried; the errors are returned together.
func (s *Store) RollUpScores(ctx context.Context, f RatedFight) error {
	var errs []error
	seen := map[string]bool{}
	for _, row := range f.Summary.Roster {
		key := character.KeyFromUnit(f.Region, f.Ruleset, row.Name)
		if seen[key] {
			continue
		}
		seen[key] = true
		sc, err := s.ScoreV2(ctx, key)
		if errors.Is(err, ErrNoLogs) {
			continue
		}
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if _, err := s.Pool.Exec(ctx,
			`update rating_scores set gear_score = $4, experience = $5, attendance = $6, tier = $7
			 where report_id = $1 and fight_index = $2 and player_key = $3`,
			f.ReportID, f.FightIndex, key, sc.GearScore, sc.Components.Experience,
			sc.Components.Attendance, sc.Tier); err != nil {
			errs = append(errs, fmt.Errorf("rating: roll up %s/%d/%s: %w", f.ReportID, f.FightIndex, key, err))
		}
	}
	return errors.Join(errs...)
}
