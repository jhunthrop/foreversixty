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
	ItemLevel           *int            `json:"item_level"`
	ItemLevelDelta      *int            `json:"item_level_delta"`
	LoggedAt            time.Time       `json:"logged_at"`
	Failing             int             `json:"failing"`
	NudgeText           *string         `json:"nudge_text,omitempty"`
	readinessScore      float64
}

// ReadinessView is the body of GET /v1/guilds/{id}/readiness.
type ReadinessView struct {
	MedianItemLevel *int           `json:"median_item_level"`
	GeneratedAt     time.Time      `json:"generated_at"`
	Rows            []ReadinessRow `json:"rows"`
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
func (s *Store) Readiness(ctx context.Context, guildID int64, officerView bool) (ReadinessView, error) {
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
	for _, row := range roster {
		if !row.Verified {
			continue
		}
		band, hasBand := s.loadBandFor(row.className, row.specStr, row.faction)
		cr := computeReadiness(row.Consent, band, hasBand, row.gear, row.enchants,
			row.bags, row.hasBagsSection, row.talentPointsSpent, row.hasTalents,
			row.level, row.hasLevel, row.ItemLevel, median)

		rr := ReadinessRow{
			CharacterKey: row.CharacterKey, Name: row.Name, Consent: row.Consent,
			Enchants:            EnchantsView{MissingSlots: cr.MissingEnchantSlots, Checked: cr.EnchantsChecked},
			Consumables:         ConsumablesView{State: cr.ConsumablesState},
			TalentPointsUnspent: cr.TalentPointsUnspent,
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
		rr.readinessScore = float64(rr.Failing)*100 + cr.GearGap.GainDps
		if officerView {
			text := nudgeText(row.Name, cr)
			rr.NudgeText = &text
		}
		view.Rows = append(view.Rows, rr)
	}

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
func nudgeText(name string, cr CharacterReadiness) string {
	fails := cr.failText()
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

	view, err := s.Store.Readiness(r.Context(), guildID, verifiedOfficer)
	if err != nil {
		s.fail(w, r, "readiness", fmt.Errorf("readiness: %w", err), "could not load that guild's readiness just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}
