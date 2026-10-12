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
//	                   best-in-slot set for the character's level band (to 60) in the
//	                   same data build (sims.GearScore over sims.FullGearScore)
//	experience  0-700  25 a distinct dungeon cleared + 40 a distinct raid boss killed, by
//	                   any character the account has linked, capped
//	parses      0-600  the mean of the character's best parse percentile on each boss
//	attendance  0-200  the share of its guild's nights in the last 90 days it was at
//	discipline  0-100  the share of its rated fights pulled with a flask, a food and
//	                   every enchantable slot enchanted
//
// score is their sum; tier names the band the score falls in. A character with no rated
// fight has no score at all (ErrNoLogs), never a zero.
//
// The rater also writes a snapshot of each player's score onto the fight it just rated
// (RollUpScores: rating_scores' score, gear_score, experience, attendance, tier). The
// character page's score_history reads it (rankings.Store.scoreHistory: the last snapshot of
// each UTC date), with one gap: the backfill job re-rates a fight by withdrawing and
// rewriting its rows and does not roll up, so a backfill-re-rated row has NULL snapshot
// columns and is left out of the history (a date keeps a point only while one of its fights
// still carries a snapshot). ScoreV2 itself never reads the snapshot; it always computes from what is stored now.

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
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

// specAliases maps a spec slug as the combat log names it (logs/engine/units.RetailSpecName
// gives retail names) onto the slug the build data keys specs by
// (data/curated/stat-weights.json, data/builds/<build>/bis/<spec>.json). With it every
// Classic class/spec pair RetailSpecName can produce resolves to one of the data's 28 spec
// slugs; TestEveryLoggedClassicSpecResolvesToADataSpec holds that.
var specAliases = map[string]string{
	"rogue-outlaw":   "rogue-combat",
	"druid-guardian": "druid-feral-bear",
}

// DataSpecSlug is the build data's spec slug ("warrior-arms") for a logged class and spec
// ("Warrior", "Arms"), the one the addon inbox prices loot under. Empty when either is.
func DataSpecSlug(class, spec string) string {
	if class == "" || spec == "" {
		return ""
	}
	return dataSpecSlug(class, spec)
}

// dataSpecSlug is the build data's slug for a logged class and spec.
func dataSpecSlug(class, spec string) string {
	slug := specSlug(class, spec)
	if alias, ok := specAliases[slug]; ok {
		return alias
	}
	return slug
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
	// GearScore is contract 5's gear_score: the gear component in points (0-500), which
	// compares across specs. GearScoreRaw is the raw sims.GearScore, rounded, in the spec's
	// own weight units. ItemLevel is the equipped item level (0 when unknown).
	GearScore    int
	GearScoreRaw int
	ItemLevel    int
	// Attendance is contract 5's attendance: the fraction (0-1) of the guild's nights in
	// AttendanceWindow the character was at.
	Attendance float64
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

	raw, full, err := s.gearScore(ctx, characterKey, fights.spec)
	if err != nil {
		return ScoreV2{}, err
	}
	out.GearScoreRaw = int(math.Round(raw))
	if full > 0 {
		out.Components.Gear = points(GearMax*raw/full, GearMax)
	}
	out.GearScore = out.Components.Gear
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
		out.Attendance = float64(attended) / float64(nights)
		out.Components.Attendance = points(AttendanceMax*out.Attendance, AttendanceMax)
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
		out.spec = dataSpecSlug(*class, *spec)
	}
	return out, nil
}

// gearScore is sims.GearScore of the character's stored addon export under spec, and
// the mark it is measured against, sims.FullGearScore of the same spec in the export's own
// data build, for the character's own level band (sims.GearBand, the addon's Gear.bandFor
// rule): until 60 a character is measured against the best set for its level, not for
// 60. An export with no level= section (an addon older than that section) is measured
// against the level-60 set, as every export was before the band rule. No export or no
// spec is no gear (0, 0). Gear the loaded data cannot score - an unknown spec, build or
// class, or a spec with no best-in-slot set - is also no gear, with one log line naming
// the spec, rather than an error: the other four components still stand.
func (s *Store) gearScore(ctx context.Context, key, spec string) (raw, full float64, err error) {
	if spec == "" {
		return 0, 0, nil
	}
	var export string
	err = s.Pool.QueryRow(ctx, `select export from addon_exports where character_key = $1`, key).Scan(&export)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, 0, nil
	}
	if err != nil {
		return 0, 0, fmt.Errorf("rating: export of %s: %w", key, err)
	}
	decoded, ok := fs1.Decode(export)
	if !ok {
		s.logger().Warn("rating", "op", "gear", "character", key, "spec", spec,
			"err", "the stored export is not an FS1 code")
		return 0, 0, nil
	}
	raw, err = sims.GearScore(ctx, export, spec)
	var band int
	var hasBand bool
	if err == nil {
		band, hasBand, err = sims.GearBand(ctx, decoded.DataBuild, spec, gearLevel(decoded))
	}
	if err == nil && !hasBand {
		// Below the first band (level 20) the spec has no best-in-slot set to measure
		// against, so there is no mark: the gear component is 0, not a share of a
		// level-60 set the character cannot wear yet.
		return raw, 0, nil
	}
	if err == nil {
		full, err = sims.FullGearScore(ctx, decoded.DataBuild, spec, band)
	}
	if errors.Is(err, sims.ErrGearUnscorable) || errors.Is(err, sims.ErrNoBestInSlot) {
		s.logger().Warn("rating", "op", "gear", "character", key, "spec", spec, "err", err)
		return 0, 0, nil
	}
	if err != nil {
		return 0, 0, fmt.Errorf("rating: gear score of %s: %w", key, err)
	}
	return raw, full, nil
}

// gearLevel is the level an export is measured at: its level= section, else 60 (an export
// from an addon older than that section, measured against the level-60 set as every export
// was before the band rule). The gear component and the gear panel (CharacterGear) share it,
// so both read the same band.
func gearLevel(decoded fs1.Decoded) int {
	if decoded.HasLevel {
		return decoded.Level
	}
	return sims.FullGearBand
}

// ItemLevel is the item level ScoreV2 carries (itemLevel), for a reader that needs it for a
// character with no score yet: the nightly data addon's no-logs record.
func (s *Store) ItemLevel(ctx context.Context, key string) (int, error) {
	return s.itemLevel(ctx, key)
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
// rankings character page places them (rankings.Store.CharacterFights' best[], the page's
// own bests without its anonymise check, identity or score history); 0 with none.
func (s *Store) meanBestParse(ctx context.Context, key string) (float64, error) {
	parts := strings.SplitN(key, "/", 3)
	if len(parts) != 3 {
		return 0, nil
	}
	page, _, err := (&rankings.Store{Pool: s.Pool}).CharacterFights(ctx, parts[0], parts[1], parts[2])
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

// guildNights counts the nights of the character's verified guild since since and how
// many of them the character was at. A night is a UTC date on which one of the guild's
// reports that is not private has a fight (fights.start_ms, the fight's own fought_at); the
// character was at it when it has a rated or ranked row in one of those reports fought on
// that date. No verified guild is no nights.
func (s *Store) guildNights(ctx context.Context, key string, since time.Time) (nights, attended int, err error) {
	err = s.Pool.QueryRow(ctx,
		`with guild_reports as (
		   select r.id from reports r
		   join guild_characters gc on gc.guild_id = r.guild_id
		   where gc.character_key = $1 and gc.verified_at is not null and r.visibility <> 'private'
		 ), nights as (
		   select distinct (to_timestamp(f.start_ms / 1000.0) at time zone 'UTC')::date as night
		   from fights f join guild_reports g on g.id = f.report_id
		   where to_timestamp(f.start_ms / 1000.0) >= $2
		 ), present as (
		   select (rs.fought_at at time zone 'UTC')::date as night
		   from rating_scores rs join guild_reports g on g.id = rs.report_id
		   where rs.player_key = $1
		   union
		   select (m.fought_at at time zone 'UTC')::date
		   from fight_metrics m join guild_reports g on g.id = m.report_id
		   where m.player_key = $1 and m.state <> 'removed'
		 )
		 select (select count(*) from nights),
		        (select count(*) from nights n where n.night in (select night from present))`,
		key, since).Scan(&nights, &attended)
	if err != nil {
		return 0, 0, fmt.Errorf("rating: guild nights of %s: %w", key, err)
	}
	return nights, attended, nil
}

// disciplineFlasks and disciplineFoods are every flask and every food in the consumables
// catalogue (logs/engine/mechanics/consumables/catalogue.json), across all roles: a
// category is met by any consumable of that category, not only one listed for the
// player's own role (a warrior on a tank's food is fed).
//
// The catalogue's weapon_enchant category (sharpening stones, wizard oils) is left out of
// the requirement - treated as unknown - until a real beta log shows how stones and oils
// appear in COMBATANT_INFO (as an aura, or as a temporary enchant on the weapon).
var disciplineFlasks, disciplineFoods = consumableCategories()

func consumableCategories() (flasks, foods map[int64]bool) {
	flasks, foods = map[int64]bool{}, map[int64]bool{}
	for _, rc := range consumables.Load().Roles {
		for _, e := range rc.Flask {
			flasks[e.SpellID] = true
		}
		for _, e := range rc.Food {
			foods[e.SpellID] = true
		}
	}
	return flasks, foods
}

// fightDiscipline is whether row's player pulled fully prepared: a flask and a food
// (disciplineFlasks, disciplineFoods), and every equipped slot of disciplineEnchantSlots
// enchanted. nil is "not known": no combatant info with gear for the player.
func fightDiscipline(fight summary.Summary, row summary.RosterRow) *bool {
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
	full := carries(*combatant, disciplineFlasks) && carries(*combatant, disciplineFoods) && enchanted(*combatant)
	return &full
}

// carries is true when the combatant has any consumable of the category.
func carries(c summary.CombatantRow, category map[int64]bool) bool {
	for _, a := range c.Consumables {
		if category[a.SpellID] {
			return true
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
// player's row of the fight f just rated (rating_scores' score, gear_score - the gear
// points - experience points, attendance - the fraction - and tier), the snapshot the
// character page's score_history reads; the file comment says where that history has gaps.
// A player with no score yet (ErrNoLogs: the fight was not public, say) is left NULL. Every
// player is tried; the errors are returned together.
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
			`update rating_scores set score = $4, gear_score = $5, experience = $6, attendance = $7, tier = $8
			 where report_id = $1 and fight_index = $2 and player_key = $3`,
			f.ReportID, f.FightIndex, key, sc.Score, sc.GearScore, sc.Components.Experience,
			sc.Attendance, sc.Tier); err != nil {
			errs = append(errs, fmt.Errorf("rating: roll up %s/%d/%s: %w", f.ReportID, f.FightIndex, key, err))
		}
	}
	return errors.Join(errs...)
}
