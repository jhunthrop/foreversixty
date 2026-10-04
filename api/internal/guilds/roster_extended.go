// api/internal/guilds/roster_extended.go
//
// The roster-row enrichment the contract's RosterRow adds beyond what HomeRoster's own
// join already carries (docs/contracts/2026-10-04-guild-centre-api.md): attendance, best
// parse, rating and last-report-at, each its own small query over this guild's reports
// rather than one unmanageable join. Also the Overview summary and the standing line,
// which both read the already-enriched roster rather than querying a second time.
package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// attendanceWindow is the contract's own "last 8 guild reports" (design spec §4.B).
const attendanceWindow = 8

// enrichRoster fills in Attendance, BestParse, Rating and LastReportAt on every row of out,
// in place - called once by HomeRoster after the base query, never per row.
func (s *Store) enrichRoster(ctx context.Context, guildID int64, out []RosterRow) error {
	byKey := make(map[string]*RosterRow, len(out))
	for i := range out {
		byKey[out[i].CharacterKey] = &out[i]
	}

	nights, err := s.attendanceWindowSize(ctx, guildID)
	if err != nil {
		return err
	}
	for i := range out {
		out[i].Attendance.Nights = nights
	}
	present, err := s.attendancePresent(ctx, guildID)
	if err != nil {
		return err
	}
	for key, n := range present {
		if row, ok := byKey[key]; ok {
			row.Attendance.Present = n
		}
	}

	bestParses, err := s.bestParses(ctx, guildID)
	if err != nil {
		return err
	}
	for key, bp := range bestParses {
		if row, ok := byKey[key]; ok {
			b := bp
			row.BestParse = &b
		}
	}

	ratings, err := s.latestRatings(ctx, guildID)
	if err != nil {
		return err
	}
	for key, rv := range ratings {
		if row, ok := byKey[key]; ok {
			r := rv
			row.Rating = &r
		}
	}

	lastReportAt, err := s.lastReportAtByOwner(ctx, guildID)
	if err != nil {
		return err
	}
	for i := range out {
		if t, ok := lastReportAt[out[i].ownerUserID]; ok {
			out[i].LastReportAt = &t
		}
	}
	return nil
}

// attendanceWindowSize is how many of the guild's own reports the attendance window
// actually spans - attendanceWindow, or fewer when the guild has not logged that many yet.
func (s *Store) attendanceWindowSize(ctx context.Context, guildID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from (select id from reports where guild_id = $1 order by created_at desc limit $2) x`,
		guildID, attendanceWindow).Scan(&n); err != nil {
		return 0, fmt.Errorf("guilds: attendance window: %w", err)
	}
	return n, nil
}

// attendancePresent is character_key -> how many of the attendance window's own reports
// that character appears in (fights.players, distinct by report so a character fought in
// more than one fight on the same night is still counted once).
func (s *Store) attendancePresent(ctx context.Context, guildID int64) (map[string]int, error) {
	rows, err := s.Pool.Query(ctx, `
		select character_key, count(distinct report_id) from (
		  select unnest(f.players) as character_key, f.report_id
		  from fights f
		  join (select id from reports where guild_id = $1 order by created_at desc limit $2) recent
		    on recent.id = f.report_id
		) present
		group by character_key`, guildID, attendanceWindow)
	if err != nil {
		return nil, fmt.Errorf("guilds: attendance present: %w", err)
	}
	defer rows.Close()
	out := map[string]int{}
	for rows.Next() {
		var key string
		var n int
		if err := rows.Scan(&key, &n); err != nil {
			return nil, fmt.Errorf("guilds: attendance present: %w", err)
		}
		out[key] = n
	}
	return out, rows.Err()
}

// bestParses is character_key -> that character's single best kill-fight parse across the
// guild's own reports - dps for every role but healer, hps for healer, matching
// rankings.Store's own metricOf convention without importing that package.
func (s *Store) bestParses(ctx context.Context, guildID int64) (map[string]BestParse, error) {
	rows, err := s.Pool.Query(ctx, `
		select distinct on (m.player_key) m.player_key, coalesce(m.role, ''),
		       m.metric_dps, m.metric_hps, coalesce(f.name, ''), m.report_id, m.fight_index
		from fight_metrics m
		join reports r on r.id = m.report_id
		join fights f on f.report_id = m.report_id and f.fight_index = m.fight_index
		where r.guild_id = $1 and m.kill and m.state <> 'removed'
		order by m.player_key,
		  (case when m.role = 'healer' then coalesce(m.metric_hps, 0) else coalesce(m.metric_dps, 0) end) desc`,
		guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: best parses: %w", err)
	}
	defer rows.Close()
	out := map[string]BestParse{}
	for rows.Next() {
		var (
			key, role, encounter string
			dps, hps             *float64
			bp                   BestParse
		)
		if err := rows.Scan(&key, &role, &dps, &hps, &encounter, &bp.ReportID, &bp.FightIndex); err != nil {
			return nil, fmt.Errorf("guilds: best parses: %w", err)
		}
		bp.Encounter = encounter
		if role == "healer" {
			bp.Metric = "hps"
			if hps != nil {
				bp.Value = *hps
			}
		} else {
			bp.Metric = "dps"
			if dps != nil {
				bp.Value = *dps
			}
		}
		out[key] = bp
	}
	return out, rows.Err()
}

// ratingComponentScore is the one field toComponentDTOs' own JSON shape this reader needs -
// api/internal/rating/cards.go's componentDTO, duplicated narrowly rather than imported, to
// keep this package from depending on api/internal/rating for a two-field read.
type ratingComponentScore struct {
	Name  string   `json:"name"`
	Score *float64 `json:"score"`
}

// latestRatings is character_key -> its latest sufficient rating_scores row's overall and
// six named components, for the guild's own reports.
func (s *Store) latestRatings(ctx context.Context, guildID int64) (map[string]RatingView, error) {
	rows, err := s.Pool.Query(ctx, `
		select distinct on (rs.player_key) rs.player_key, rs.overall, rs.components
		from rating_scores rs
		join reports r on r.id = rs.report_id
		where r.guild_id = $1 and not coalesce(rs.insufficient, false)
		order by rs.player_key, rs.fought_at desc`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: latest ratings: %w", err)
	}
	defer rows.Close()
	out := map[string]RatingView{}
	for rows.Next() {
		var (
			key string
			raw []byte
			rv  RatingView
		)
		if err := rows.Scan(&key, &rv.Overall, &raw); err != nil {
			return nil, fmt.Errorf("guilds: latest ratings: %w", err)
		}
		var components []ratingComponentScore
		if err := json.Unmarshal(raw, &components); err == nil {
			for _, c := range components {
				if c.Score == nil {
					continue
				}
				switch c.Name {
				case "output":
					rv.Output = *c.Score
				case "survival":
					rv.Survival = *c.Score
				case "mechanics":
					rv.Mechanics = *c.Score
				case "utility":
					rv.Utility = *c.Score
				case "preparation":
					rv.Preparation = *c.Score
				case "activity":
					rv.Activity = *c.Score
				}
			}
		}
		out[key] = rv
	}
	return out, rows.Err()
}

// lastReportAtByOwner is user_id -> the newest created_at of any report that account owns
// in this guild.
func (s *Store) lastReportAtByOwner(ctx context.Context, guildID int64) (map[int64]time.Time, error) {
	rows, err := s.Pool.Query(ctx,
		`select owner_id, max(created_at) from reports
		 where guild_id = $1 and owner_id is not null group by owner_id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: last report at: %w", err)
	}
	defer rows.Close()
	out := map[int64]time.Time{}
	for rows.Next() {
		var id int64
		var at time.Time
		if err := rows.Scan(&id, &at); err != nil {
			return nil, fmt.Errorf("guilds: last report at: %w", err)
		}
		out[id] = at
	}
	return out, rows.Err()
}

// ratingFloorPercentile is the Overview summary's own below_rating_floor cutoff (design
// spec §4.A: "a floor the guild or site defines, e.g. percentile 25 within role"). This
// reader applies it guild-wide rather than grouped by role - a deviation from the design
// spec's own "within role" phrasing, named in CONTROL_CENTRE.md, kept for this round since
// grouping a 20-30 raider guild by role would leave most role buckets too small for a
// percentile to mean anything.
const ratingFloorPercentile = 0.25

// belowRatingFloor counts roster rows whose overall rating sits at or below the guild's own
// 25th percentile, among every row that has a rating at all (design spec §4.A). Zero or one
// rated row never has a floor to sit below.
func belowRatingFloor(roster []RosterRow) int {
	var overalls []float64
	for _, row := range roster {
		if row.Rating != nil {
			overalls = append(overalls, row.Rating.Overall)
		}
	}
	if len(overalls) < 2 {
		return 0
	}
	sort.Float64s(overalls)
	idx := int(float64(len(overalls)-1) * ratingFloorPercentile)
	floor := overalls[idx]
	n := 0
	for _, row := range roster {
		if row.Rating != nil && row.Rating.Overall <= floor {
			n++
		}
	}
	return n
}

// homeSummary assembles the Overview tab's own glance row from the already-enriched roster
// plus two small guild-wide counts (named encounters down, pulls this tier) that do not fit
// naturally on any one roster row.
func (s *Store) homeSummary(ctx context.Context, guildID int64, roster []RosterRow) (SummaryView, error) {
	sv := SummaryView{UpdatedAt: time.Now().UTC()}
	for _, row := range roster {
		if row.Verified {
			sv.Raiders++
		} else {
			sv.WaitingForApproval++
		}
	}
	sv.BelowRatingFloor = belowRatingFloor(roster)

	if err := s.Pool.QueryRow(ctx, `
		select count(distinct f.encounter_id) from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.kill and f.encounter_id is not null`, guildID).
		Scan(&sv.NamedEncountersDown); err != nil {
		return SummaryView{}, fmt.Errorf("guilds: named encounters down: %w", err)
	}
	// pulls_this_tier: every fight ever logged for the guild. No raid-tier-start
	// timestamp exists anywhere queryable (the design doc's own 9 December 2026 tier
	// open is a future calendar date, not a stored marker) - see CONTROL_CENTRE.md.
	if err := s.Pool.QueryRow(ctx, `
		select count(*) from fights f join reports r on r.id = f.report_id where r.guild_id = $1`, guildID).
		Scan(&sv.PullsThisTier); err != nil {
		return SummaryView{}, fmt.Errorf("guilds: pulls this tier: %w", err)
	}
	return sv, nil
}

// standingFor is the member/officer-only "where do I stand" line (design spec §4.A.1):
// among this viewer's own roster rows, the most representative one (verified over
// unverified, known class/spec over unknown) ranked by item level against every other
// verified, gear-consented character of the same class and spec. nil when nothing in the
// roster can anchor it.
func standingFor(s *Store, roster []RosterRow, viewerUserID int64) *StandingView {
	var self *RosterRow
	for i := range roster {
		row := &roster[i]
		if row.ownerUserID != viewerUserID || !row.Verified || !hasGearConsent(row.Consent) ||
			row.className == "" || row.specStr == "" || row.ItemLevel == nil {
			continue
		}
		if self == nil {
			self = row
		}
	}
	if self == nil {
		return nil
	}

	type peer struct {
		key       string
		itemLevel int
	}
	var peers []peer
	for _, row := range roster {
		if row.Verified && hasGearConsent(row.Consent) && row.className == self.className &&
			row.specStr == self.specStr && row.ItemLevel != nil {
			peers = append(peers, peer{row.CharacterKey, *row.ItemLevel})
		}
	}
	sort.Slice(peers, func(i, j int) bool {
		if peers[i].itemLevel != peers[j].itemLevel {
			return peers[i].itemLevel > peers[j].itemLevel
		}
		return peers[i].key < peers[j].key
	})
	rank := 0
	for i, p := range peers {
		if p.key == self.CharacterKey {
			rank = i + 1
			break
		}
	}

	band, hasBand := s.loadBandFor(self.className, self.specStr, self.faction)
	cr := computeReadiness(self.Consent, band, hasBand, self.gear, self.enchants,
		self.bags, self.hasBagsSection, self.talentPointsSpent, self.hasTalents, self.ItemLevel, nil)

	return &StandingView{
		Spec: self.specStr, Class: self.className,
		SameSpecCount: len(peers), RankByItemLevel: rank, ItemLevel: *self.ItemLevel,
		NeedsBeforeNextRaid: cr.Needs(3),
	}
}
