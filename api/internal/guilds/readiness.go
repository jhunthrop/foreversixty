// api/internal/guilds/readiness.go
//
// GET /v1/guilds/{id}/readiness (contract step 5,
// docs/contracts/2026-10-04-guild-centre-api.md): member and officer only (design spec
// §4.0: "Not shown" to the public) - every member sees every row, an officer additionally
// gets nudge_text. Built on readiness_core.go's computeReadiness, the exact computation the
// roster standing line's own needs_before_next_raid already calls, so the two surfaces can
// never disagree about one character's own gaps.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// GearGapView/EnchantsView/ConsumablesView are the readiness row's own three check
// objects (contract's exact field names).
type GearGapView struct {
	Upgrades      int     `json:"upgrades"`
	GainDps       float64 `json:"gain_dps"`
	NotSimChecked int     `json:"not_sim_checked"`
	// Unit is what gain_dps counts: "dps" for a damage spec, null for a healer or tank, whose
	// BiS file measures no gain in a unit of its own, so no figure prints (week-one spec §4.2).
	Unit *string `json:"unit"`
}

type EnchantsView struct {
	MissingSlots []string `json:"missing_slots"`
	Checked      bool     `json:"checked"`
}

type ConsumablesView struct {
	State string `json:"state"`
}

// ReadinessRow is one character's own readiness board row.
type ReadinessRow struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
	Consent      string `json:"consent"`
	// Level is the character's level (export, else the addon roster), null when neither
	// carries one. Under 60 nothing is graded: gear_gap null, enchants unchecked, consumables
	// unknown, failing 0, no item_level_delta and no nudge_text (week-one spec §4.2).
	Level               *int            `json:"level"`
	GearGap             *GearGapView    `json:"gear_gap"`
	Enchants            EnchantsView    `json:"enchants"`
	Consumables         ConsumablesView `json:"consumables"`
	TalentPointsUnspent int             `json:"talent_points_unspent"`
	TalentPointsRead    bool            `json:"talent_points_read"`
	ItemLevel           *int            `json:"item_level"`
	ItemLevelDelta      *int            `json:"item_level_delta"`
	LoggedAt            time.Time       `json:"logged_at"`
	Failing             int             `json:"failing"`
	// SyncedBeforeCutoff is true when the export was synced more than READINESS_FRESH_DAYS ago,
	// judged against now on the filtered and unfiltered read alike (stale; null only on a
	// no_export row). NoExport is the raid-filtered read's signed character with no export.
	SyncedBeforeCutoff *bool   `json:"synced_before_cutoff"`
	NoExport           bool    `json:"no_export"`
	NudgeText          *string `json:"nudge_text,omitempty"`
	// Attuned is an officer's only: the attunement of the board's raid (the filtered raid's,
	// else the next raid's; raid_prep.go attunementOf), its key carried or its quest chain done:
	// true or false when read, null when the run needs none or the export cannot say.
	Attuned *bool `json:"attuned"`
	// Resistances is an officer's only (D14): the export's resistance totals at gear consent,
	// null without them.
	Resistances *fs1.Resistances `json:"resistances"`
	// FlaskFood is an officer's only (D14): whether the bags hold a flask or raid elixir and
	// buff food, null without the bags share or an export that names them.
	FlaskFood      *FlaskFoodView `json:"flask_food"`
	readinessScore float64
}

// FlaskFoodView is a Readiness row's flask and food on hand.
type FlaskFoodView struct {
	FlaskOrElixir bool `json:"flask_or_elixir"`
	Food          bool `json:"food"`
}

// flaskFoodViewOf is the row's flask and food, nil when it cannot be read.
func flaskFoodViewOf(row RosterRow) *FlaskFoodView {
	r, ok := suppliesOf(row)
	if !ok {
		return nil
	}
	return &FlaskFoodView{
		FlaskOrElixir: hasKind(r.Stacks, supplyFlask) || hasKind(r.Stacks, supplyElixir),
		Food:          hasKind(r.Stacks, supplyFood),
	}
}

// officerPrep fills a row's officer-only raid-prep fields for the board's zone (nil: no raid).
func officerPrep(rr *ReadinessRow, row RosterRow, zone *Zone) {
	if zone != nil {
		rr.Attuned = attunedFlag(attunementOf(*zone, row))
	}
	rr.Resistances = resistancesOf(row)
	rr.FlaskFood = flaskFoodViewOf(row)
}

// ReadinessView is the body of GET /v1/guilds/{id}/readiness.
type ReadinessView struct {
	MedianItemLevel *int           `json:"median_item_level"`
	GeneratedAt     time.Time      `json:"generated_at"`
	Rows            []ReadinessRow `json:"rows"`
	// Raid is set on the raid-filtered read only.
	Raid *ReadinessRaid `json:"raid,omitempty"`
}

// medianItemLevel is the middle value of the known item levels of the roster's raiders (a
// leveling character's item level is no raid yardstick, so it is left out), nil when none has
// one yet.
func medianItemLevel(roster []RosterRow) *int {
	var levels []int
	for _, row := range roster {
		if row.ItemLevel != nil && !leveling(row.Level) {
			levels = append(levels, *row.ItemLevel)
		}
	}
	if len(levels) == 0 {
		return nil
	}
	sort.Ints(levels)
	mid := levels[len(levels)/2]
	if len(levels)%2 == 0 {
		mid = (levels[len(levels)/2-1] + levels[len(levels)/2]) / 2
	}
	return &mid
}

// readinessRoster is the roster every readiness read grades: HomeRoster, with each row's
// level filled from the guild's addon roster snapshot where the export carries none.
func (s *Store) readinessRoster(ctx context.Context, guildID int64) ([]RosterRow, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return nil, err
	}
	// moderator/verifiedOfficer/actorID only affect MayRemove/MayApprove/standing, none
	// of which a readiness read surfaces - zero values are harmless here.
	roster, err := s.HomeRoster(ctx, guildID, g, 0, false, false)
	if err != nil {
		return nil, err
	}
	if _, _, err := s.homeAddonSnapshot(ctx, guildID, roster); err != nil {
		return nil, err
	}
	return roster, nil
}

// Readiness assembles the readiness board: one row per verified roster character
// (readiness is a pre-pull tool for raiders the guild has already vouched for - an
// unverified row has nothing to ready up for yet, the same reasoning the standing line's
// own viewer-consent gate uses one level up). officerView adds nudge_text to every graded row.
func (s *Store) Readiness(ctx context.Context, guildID int64, officerView bool, raid *Raid) (ReadinessView, error) {
	scope, err := s.newReadinessScope(ctx, raid)
	if err != nil {
		return ReadinessView{}, err
	}
	roster, err := s.readinessRoster(ctx, guildID)
	if err != nil {
		return ReadinessView{}, err
	}
	nudgeDay := ""
	var zone *Zone
	if officerView {
		if nudgeDay, err = s.nudgeDay(ctx, guildID, raid); err != nil {
			return ReadinessView{}, err
		}
		if zone, err = s.readinessZone(ctx, guildID, raid); err != nil {
			return ReadinessView{}, err
		}
	}

	now := time.Now()
	if scope != nil {
		now = scope.now
	}
	median := medianItemLevel(roster)
	view := ReadinessView{MedianItemLevel: median, GeneratedAt: now.UTC(), Rows: []ReadinessRow{}}
	seen := map[string]bool{}
	for _, row := range roster {
		if !row.Verified || !scope.includes(row.CharacterKey) {
			continue
		}
		seen[row.CharacterKey] = true
		cr := s.readinessOf(row, median)
		rr := readinessRowOf(row, cr)
		markFreshness(&rr, now)
		rr.readinessScore = float64(rr.Failing)*readinessFailWeight + cr.GearGap.GainDps
		if officerView && !cr.Leveling {
			text := nudgeText(row.Name, cr, nudgeDay, staleFailure(rr, now)...)
			rr.NudgeText = &text
		}
		if officerView {
			officerPrep(&rr, row, zone)
		}
		view.Rows = append(view.Rows, rr)
	}
	view.Raid = scope.view()
	view.Rows = append(view.Rows, scope.noExportRows(seen)...)
	sortReadinessRows(view.Rows)
	return view, nil
}

// readinessFailWeight puts every failing check ahead of any gear gain in the worst-first sort.
const readinessFailWeight = 100

// readinessRowOf is a roster row's board row for its computed readiness.
func readinessRowOf(row RosterRow, cr CharacterReadiness) ReadinessRow {
	rr := ReadinessRow{
		CharacterKey: row.CharacterKey, Name: row.Name, Consent: row.Consent, Level: cr.Level,
		Enchants:            EnchantsView{MissingSlots: cr.MissingEnchantSlots, Checked: cr.EnchantsChecked},
		Consumables:         ConsumablesView{State: cr.ConsumablesState},
		TalentPointsUnspent: cr.TalentPointsUnspent,
		TalentPointsRead:    cr.TalentsRead,
		ItemLevel:           cr.ItemLevel, ItemLevelDelta: cr.ItemLevelDelta,
		LoggedAt: row.LoggedAt, Failing: cr.FailingCount(),
	}
	if rr.Enchants.MissingSlots == nil {
		rr.Enchants.MissingSlots = []string{}
	}
	if row.Class != nil {
		rr.Class = *row.Class
	}
	if row.Spec != nil {
		rr.Spec = *row.Spec
	}
	if cr.GearChecked {
		rr.GearGap = &GearGapView{Upgrades: cr.GearGap.Upgrades, GainDps: cr.GearGap.GainDps, NotSimChecked: cr.GearGap.NotSimChecked}
		if cr.GearUnit != "" {
			unit := cr.GearUnit
			rr.GearGap.Unit = &unit
		}
	}
	return rr
}

// sortReadinessRows orders the board: graded rows worst-first (a no_export row after every
// row with an export), then leveling rows, highest level first, then by name (week-one spec
// §3.5).
func sortReadinessRows(rows []ReadinessRow) {
	sort.SliceStable(rows, func(i, j int) bool {
		a, b := rows[i], rows[j]
		aLeveling, bLeveling := leveling(a.Level), leveling(b.Level)
		switch {
		case aLeveling != bLeveling:
			return bLeveling
		case aLeveling:
			if *a.Level != *b.Level {
				return *a.Level > *b.Level
			}
			return a.Name < b.Name
		default:
			return a.readinessScore > b.readinessScore
		}
	})
}

// nudgeDay is the date the officer's nudge asks for a fix by ("Thu 10 Dec", the needs line's
// own date): the raid the board is cut to, else the guild's next planned raid, in the guild's
// raid timezone (UTC when it has none). Empty when no raid is planned. A bare weekday read
// "Check before Thursday." for a raid five weeks out (guild round 3).
func (s *Store) nudgeDay(ctx context.Context, guildID int64, raid *Raid) (string, error) {
	if raid == nil {
		next, ok, err := s.NextRaid(ctx, guildID, time.Now())
		if err != nil || !ok {
			return "", err
		}
		raid = &next
	}
	tz, err := s.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return "", err
	}
	return raidDay(raid.StartsAt, tz), nil
}

// raidDayLayout is a raid's date as the site's needs line names it (raid-time.ts dateLabel):
// "Thu 10 Dec".
const raidDayLayout = "Mon 2 Jan"

// raidDay is start's date ("Thu 10 Dec") in the IANA zone tz, UTC when tz is empty or unknown.
func raidDay(start time.Time, tz string) string {
	return start.In(zoneOrUTC(tz)).Format(raidDayLayout)
}

// zoneOrUTC loads the IANA zone tz, UTC when tz is empty or unknown.
func zoneOrUTC(tz string) *time.Location {
	if tz != "" {
		if l, err := time.LoadLocation(tz); err == nil {
			return l
		}
	}
	return time.UTC
}

// nudgeText is the officer-only clipboard message (design spec §4.E): the character's own
// name plus its own failText(), joined - the same wording the standing line's
// needs_before_next_raid already uses, never a second phrasing - then the day of the raid it
// is for (day, empty with no raid planned, when the message just lists what to fix). When the
// talent-points check assumed level 60 for want of a level on file (maxTalentPoints), the
// nudge says so explicitly rather than presenting that guess as a fact the officer might
// repeat to the raider. extraFails are failures the character's own readiness does not know
// (a stale sync before a raid); they read as part of the same list.
func nudgeText(name string, cr CharacterReadiness, day string, extraFails ...string) string {
	fails := append(cr.failText(), extraFails...)
	var out string
	switch {
	case len(fails) == 0:
		out = name + ": every readiness check passes."
	case day == "":
		out = name + ": " + strings.Join(fails, "; ") + "."
	default:
		out = name + ": " + strings.Join(fails, "; ") + ". Check before " + day + "."
	}
	if cr.TalentLevelAssumed {
		out += " (No level on file, so talent points assume level 60.)"
	}
	return out
}

// The readiness words a one-character answer uses: the board's own (web/src/lib/guild/
// readiness-view.ts consumablesLabel and NOT_SHARED).
const (
	readinessStocked   = "Stocked"
	readinessShort     = "Short: "
	readinessNotShared = "Not shared"
	// readinessReady and readinessNotChecked are the board's READINESS_READY (something read,
	// nothing failing: passesEveryCheck) and NOT_CHECKED_YET (shared, nothing read yet).
	readinessReady      = "Ready"
	readinessNotChecked = "Not checked yet"
	// readinessLeveling is a leveling character's word: its level, and that checks start at 60.
	readinessLeveling = "Level %d, checks start at %d"
	// readinessShortMax is how many failing checks follow "Short: ".
	readinessShortMax = 2
)

// ReadinessFacts is one roster character's readiness as the board computes it: the word,
// and when its export was last synced (nil off the roster or with no export).
type ReadinessFacts struct {
	Word     string
	SyncedAt *time.Time
	OnRoster bool
	// ItemLevel is the roster's and the board's item level (equipped average, else the latest
	// log's), so /char prints the same number; nil when neither is known.
	ItemLevel *int
	// GearUpgrades is the board's own gear upgrade count (the "N gear upgrades waiting" the word
	// can carry), so /char's "N more upgrades" counts the same list; GearChecked is false when
	// the board read no gear, and GearUpgrades is then 0.
	GearUpgrades int
	GearChecked  bool
}

// CharacterReadinessWord is the board's word for one character: Stocked, Short: and its
// first failing checks, or Not shared (off the verified roster, or nothing to read).
func (s *Service) CharacterReadinessWord(ctx context.Context, guildID int64, characterKey string) (string, error) {
	facts, err := s.CharacterReadinessFacts(ctx, guildID, characterKey)
	return facts.Word, err
}

// CharacterReadinessFacts is CharacterReadinessWord with the sync time beside it.
func (s *Service) CharacterReadinessFacts(ctx context.Context, guildID int64, characterKey string) (ReadinessFacts, error) {
	none := ReadinessFacts{Word: readinessNotShared}
	roster, err := s.Store.readinessRoster(ctx, guildID)
	if err != nil {
		return none, fmt.Errorf("guilds: readiness word: %w", err)
	}
	for _, row := range roster {
		if row.Verified && row.CharacterKey == characterKey {
			return s.Store.readinessFactsOf(row, medianItemLevel(roster)), nil
		}
	}
	return none, nil
}

func (s *Store) readinessFactsOf(row RosterRow, median *int) ReadinessFacts {
	cr := s.readinessOf(row, median)
	facts := ReadinessFacts{Word: readinessWord(cr, row.Consent), OnRoster: true, ItemLevel: row.ItemLevel,
		GearChecked: cr.GearChecked, GearUpgrades: cr.GearGap.Upgrades}
	if !row.LoggedAt.IsZero() {
		synced := row.LoggedAt
		facts.SyncedAt = &synced
	}
	return facts
}

// readinessWord is the board's word for one row, the site's rule (readiness-view.ts
// passesEveryCheck): leveling, else the failing checks, else Stocked or Ready when anything was
// read, else Not shared for a roster-only consent and Not checked yet for one that shares.
func readinessWord(cr CharacterReadiness, consent string) string {
	if cr.Leveling {
		return fmt.Sprintf(readinessLeveling, *cr.Level, raidLevel)
	}
	if fails := cr.Needs(readinessShortMax); len(fails) > 0 {
		return readinessShort + strings.Join(fails, ", ")
	}
	switch {
	case cr.ConsumablesState == consumablesStocked:
		return readinessStocked
	case cr.GearChecked || cr.EnchantsChecked || cr.TalentsRead || cr.ConsumablesState != consumablesUnknown:
		return readinessReady
	case hasGearConsent(consent):
		return readinessNotChecked
	}
	return readinessNotShared
}

func (s *Service) readiness(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	// Verified members only, as home and loot: a character still waiting on an officer reads
	// nobody's gear or readiness.
	member, err := s.Store.IsVerifiedMember(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "readiness", err, "could not load that guild's readiness just now")
		return
	}
	if !member {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you are not a member of that guild", nil)
		return
	}
	rank, verified, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "readiness", err, "could not load that guild's readiness just now")
		return
	}
	verifiedOfficer := verified && (rank == "officer" || rank == "leader")

	raid, filtered, ok := s.readinessRaid(w, r, guildID)
	if !ok {
		return
	}
	view, err := s.Store.Readiness(r.Context(), guildID, verifiedOfficer, raid)
	if err != nil {
		s.fail(w, r, "readiness", fmt.Errorf("readiness: %w", err), "could not load that guild's readiness just now")
		return
	}
	httpx.CachePrivate(w)
	if filtered {
		httpx.WriteOK(w, r, http.StatusOK, view)
		return
	}
	next, err := s.Store.readinessNextRaid(r.Context(), guildID, time.Now())
	if err != nil {
		s.fail(w, r, "readiness", fmt.Errorf("readiness: %w", err), "could not load that guild's readiness just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusOK, readinessWithNext{ReadinessView: view, NextRaid: next})
}
