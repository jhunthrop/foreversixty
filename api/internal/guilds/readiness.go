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
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// GearGapView/EnchantsView/ConsumablesView are the readiness row's own three check
// objects (contract's exact field names).
type GearGapView struct {
	Upgrades      int     `json:"upgrades"`
	GainDps       float64 `json:"gain_dps"`
	NotSimChecked int     `json:"not_sim_checked"`
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
	CharacterKey        string          `json:"character_key"`
	Name                string          `json:"name"`
	Class               string          `json:"class"`
	Spec                string          `json:"spec"`
	Consent             string          `json:"consent"`
	GearGap             *GearGapView    `json:"gear_gap"`
	Enchants            EnchantsView    `json:"enchants"`
	Consumables         ConsumablesView `json:"consumables"`
	TalentPointsUnspent int             `json:"talent_points_unspent"`
	TalentPointsRead    bool            `json:"talent_points_read"`
	ItemLevel           *int            `json:"item_level"`
	ItemLevelDelta      *int            `json:"item_level_delta"`
	LoggedAt            time.Time       `json:"logged_at"`
	Failing             int             `json:"failing"`
	// SyncedBeforeCutoff and NoExport are the raid-filtered read's: whether the export was
	// synced within READINESS_FRESH_DAYS of the raid (null on the unfiltered read and on a
	// no_export row), and a signed character with no export at all.
	SyncedBeforeCutoff *bool   `json:"synced_before_cutoff"`
	NoExport           bool    `json:"no_export"`
	NudgeText          *string `json:"nudge_text,omitempty"`
	readinessScore     float64
}

// ReadinessView is the body of GET /v1/guilds/{id}/readiness.
type ReadinessView struct {
	MedianItemLevel *int           `json:"median_item_level"`
	GeneratedAt     time.Time      `json:"generated_at"`
	Rows            []ReadinessRow `json:"rows"`
	// Raid is set on the raid-filtered read only.
	Raid *ReadinessRaid `json:"raid,omitempty"`
}

// medianItemLevel is the middle value of every roster row's own known item level (nil
// values excluded), nil when nothing in the roster has one yet.
func medianItemLevel(roster []RosterRow) *int {
	var levels []int
	for _, row := range roster {
		if row.ItemLevel != nil {
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

// Readiness assembles the readiness board: one row per verified roster character
// (readiness is a pre-pull tool for raiders the guild has already vouched for - an
// unverified row has nothing to ready up for yet, the same reasoning the standing line's
// own viewer-consent gate uses one level up). officerView adds nudge_text to every row.
func (s *Store) Readiness(ctx context.Context, guildID int64, officerView bool, raid *Raid) (ReadinessView, error) {
	scope, err := s.newReadinessScope(ctx, raid)
	if err != nil {
		return ReadinessView{}, err
	}
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return ReadinessView{}, err
	}
	// moderator/verifiedOfficer/actorID only affect MayRemove/MayApprove/standing, none
	// of which this endpoint surfaces - zero values are harmless here.
	roster, err := s.HomeRoster(ctx, guildID, g, 0, false, false)
	if err != nil {
		return ReadinessView{}, err
	}

	median := medianItemLevel(roster)
	view := ReadinessView{MedianItemLevel: median, GeneratedAt: time.Now().UTC(), Rows: []ReadinessRow{}}
	seen := map[string]bool{}
	for _, row := range roster {
		if !row.Verified || !scope.includes(row.CharacterKey) {
			continue
		}
		seen[row.CharacterKey] = true
		band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
		cr := computeReadiness(row.Consent, band, hasBand, row.gear, row.enchants,
			row.bags, row.hasBagsSection, row.talentPointsSpent, row.hasTalents,
			row.level, row.hasLevel, row.ItemLevel, median)

		rr := ReadinessRow{
			CharacterKey: row.CharacterKey, Name: row.Name, Consent: row.Consent,
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
		}
		scope.apply(&rr)
		rr.readinessScore = float64(rr.Failing)*100 + cr.GearGap.GainDps
		if officerView {
			text := nudgeText(row.Name, cr, scope.staleFailure(rr)...)
			rr.NudgeText = &text
		}
		view.Rows = append(view.Rows, rr)
	}
	view.Raid = scope.view()
	view.Rows = append(view.Rows, scope.noExportRows(seen)...)

	sort.SliceStable(view.Rows, func(i, j int) bool {
		return view.Rows[i].readinessScore > view.Rows[j].readinessScore
	})
	return view, nil
}

// nudgeText is the officer-only clipboard message (design spec §4.E): the character's own
// name plus its own failText(), joined - the same wording the standing line's
// needs_before_next_raid already uses, never a second phrasing. When the talent-points
// check assumed level 60 for want of a level section on the export (maxTalentPoints), the
// nudge says so explicitly rather than presenting that guess as a fact the officer might
// repeat to the raider.
// extraFails are failures the character's own readiness does not know (a stale sync before a
// raid); they read as part of the same list.
func nudgeText(name string, cr CharacterReadiness, extraFails ...string) string {
	fails := append(cr.failText(), extraFails...)
	var out string
	if len(fails) == 0 {
		out = name + ": every readiness check passes."
	} else {
		out = name + ": "
		for i, f := range fails {
			if i > 0 {
				out += "; "
			}
			out += f
		}
		out += " - check before Thursday."
	}
	if cr.TalentLevelAssumed {
		out += " (no level on file - talent points assume level 60)"
	}
	return out
}

// The readiness words a one-character answer uses: the board's own (web/src/lib/guild/
// readiness-view.ts consumablesLabel and NOT_SHARED).
const (
	readinessStocked   = "Stocked"
	readinessShort     = "Short: "
	readinessNotShared = "Not shared"
	// readinessShortMax is how many failing checks follow "Short: ".
	readinessShortMax = 2
)

// ReadinessFacts is one roster character's readiness as the board computes it: the word,
// and when its export was last synced (nil off the roster or with no export).
type ReadinessFacts struct {
	Word     string
	SyncedAt *time.Time
	OnRoster bool
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
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return none, fmt.Errorf("guilds: readiness word: %w", err)
	}
	roster, err := s.Store.HomeRoster(ctx, guildID, g, 0, false, false)
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
	band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
	cr := computeReadiness(row.Consent, band, hasBand, row.gear, row.enchants,
		row.bags, row.hasBagsSection, row.talentPointsSpent, row.hasTalents,
		row.level, row.hasLevel, row.ItemLevel, median)
	facts := ReadinessFacts{Word: readinessWord(cr), OnRoster: true}
	if !row.LoggedAt.IsZero() {
		synced := row.LoggedAt
		facts.SyncedAt = &synced
	}
	return facts
}

func readinessWord(cr CharacterReadiness) string {
	if fails := cr.Needs(readinessShortMax); len(fails) > 0 {
		return readinessShort + strings.Join(fails, ", ")
	}
	if cr.ConsumablesState == "stocked" {
		return readinessStocked
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
	member, err := s.Store.IsMember(r.Context(), guildID, actor.UserID)
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
