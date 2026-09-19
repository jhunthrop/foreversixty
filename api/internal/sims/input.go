package sims

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// inputMaxAge is how long an anonymous sim-input read may be cached,
// in seconds. It is short: a character's gear changes whenever they
// log out.
const inputMaxAge = 60

// Input is the character model the sim page starts from: the newest
// of the addon export and the character's last ranked fight.
//
// Armory is the contract's third source and is not here yet: nothing
// in this repository stores an Armory refresh. When it does, it joins
// the comparison below and this shape does not change.
type Input struct {
	Spec string `json:"spec"`
	// Gear is the source's own shape: the addon's export string for an
	// addon read, {"trinkets": […]} for a fight. web/src/lib/sim/character.ts
	// turns either into the envelope's CharacterSpec, which is where
	// that mapping belongs - doing it here would write it twice, in two
	// languages.
	Gear json.RawMessage `json:"gear"`
	// Talents is the fight's recorded talent split. The addon export is
	// an opaque string this repository never parses, so a character who
	// has never parsed has none.
	Talents string `json:"talents"`
	// Buffs are engine buff ids (sim/request/IDS.md), mapped from what
	// the last fight recorded.
	Buffs      []string  `json:"buffs"`
	CapturedAt time.Time `json:"captured_at"`
	Source     string    `json:"source"`
}

// FightRef is the fight an Input's spec and talents came from, so the
// handler can read that fight's summary for the buffs.
type FightRef struct {
	ReportID   string
	FightIndex int
	PlayerName string
}

// emptyGear is what Gear carries when the source recorded none, so
// the field is always valid JSON rather than a bare null.
var emptyGear = json.RawMessage(`{}`)

// SimInput assembles a character's model from the database. It reports
// false when the character is not known at all. The buffs come from the
// bucket and are the handler's to add.
func (s *Store) SimInput(ctx context.Context, key string) (Input, FightRef, bool, error) {
	out := Input{Gear: emptyGear, Buffs: []string{}}
	var (
		ref        FightRef
		export     *string
		exportAt   *time.Time
		fightSpec  *string
		split      *string
		fightAt    *time.Time
		trinkets   []int64
		reportID   *string
		fightIndex *int
		playerName *string
	)
	err := s.Pool.QueryRow(ctx,
		`select export, updated_at from addon_exports where character_key = $1`, key).
		Scan(&export, &exportAt)
	if err != nil && !isNoRows(err) {
		return Input{}, FightRef{}, false, fmt.Errorf("sims: read export %s: %w", key, err)
	}
	err = s.Pool.QueryRow(ctx,
		`select spec, talent_split, fought_at, trinkets, report_id, fight_index, player_name
		 from fight_metrics
		 where player_key = $1 and state <> 'removed'
		 order by fought_at desc limit 1`, key).
		Scan(&fightSpec, &split, &fightAt, &trinkets, &reportID, &fightIndex, &playerName)
	if err != nil && !isNoRows(err) {
		return Input{}, FightRef{}, false, fmt.Errorf("sims: read last fight %s: %w", key, err)
	}

	// Whatever the source of the gear, the spec, the talents and the
	// fight to read buffs from come from the newest ranked fight: it is
	// the only source in this repository that records any of them.
	if fightSpec != nil {
		out.Spec = *fightSpec
	}
	if split != nil {
		out.Talents = *split
	}
	if reportID != nil && fightIndex != nil {
		ref = FightRef{ReportID: *reportID, FightIndex: *fightIndex}
		if playerName != nil {
			ref.PlayerName = *playerName
		}
	}

	switch {
	case exportAt != nil && (fightAt == nil || exportAt.After(*fightAt)):
		// The addon export is the richer source of gear: it carries
		// what the character logged out in, slot by slot.
		out.Source, out.CapturedAt = "addon", *exportAt
		out.Gear = json.RawMessage(*export)
	case fightAt != nil:
		// The fallback, available to anyone who has ever parsed. A
		// ranked row keeps the trinkets, not the whole gear list, so
		// this is the thinner model of the two and the page says so
		// from the source pill.
		out.Source, out.CapturedAt = "fight", *fightAt
		if trinkets == nil {
			trinkets = []int64{}
		}
		gear, err := json.Marshal(map[string]any{"trinkets": trinkets})
		if err != nil {
			return Input{}, FightRef{}, false, fmt.Errorf("sims: encode gear %s: %w", key, err)
		}
		out.Gear = gear
	default:
		return Input{}, FightRef{}, false, nil
	}
	return out, ref, true, nil
}

func (s *Service) simInput(w http.ResponseWriter, r *http.Request) {
	region := strings.ToLower(r.PathValue("region"))
	ruleset := strings.ToLower(r.PathValue("ruleset"))
	if !character.ValidRegion(region) || !character.ValidRuleset(ruleset) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	key := region + "/" + ruleset + "/" + character.Slug(r.PathValue("name"))
	in, ref, ok, err := s.Store.SimInput(r.Context(), key)
	if err != nil {
		s.fail(w, r, "sim input", err, "could not read that character just now")
		return
	}
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	in.Buffs = s.recordedBuffs(r.Context(), ref)
	// A signed-in caller's answer is theirs; it must never reach a
	// shared cache another viewer is served from. The anonymous read
	// is the same public data the character page already serves and
	// takes the same short window.
	if auth.ActorFrom(r.Context()).Signed() {
		w.Header().Set("Cache-Control", "private, no-store")
	} else {
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(inputMaxAge))
	}
	httpx.WriteOK(w, r, http.StatusOK, in)
}

// recordedBuffs reads the buffs the character's last fight recorded and
// maps them onto the engine's ids. A deployment with no bucket, a fight
// whose summary has gone, and a summary with no row for this player all
// answer "no buffs" with a line saying which: the character model is
// still worth serving without them.
func (s *Service) recordedBuffs(ctx context.Context, ref FightRef) []string {
	if s.Summaries == nil || ref.ReportID == "" {
		return []string{}
	}
	sum, err := fightSummary(ctx, s.Summaries, ref.ReportID, ref.FightIndex)
	if err != nil {
		s.logger().Warn("sims", "op", "sim input", "report", ref.ReportID,
			"fight", ref.FightIndex, "err", err)
		return []string{}
	}
	c, ok := combatantNamed(sum, ref.PlayerName)
	if !ok {
		s.logger().Warn("sims", "op", "sim input", "report", ref.ReportID,
			"fight", ref.FightIndex, "err", "the summary has no combatant for "+ref.PlayerName)
		return []string{}
	}
	return BuffIDs(append(append([]summary.AuraRef{}, c.RaidBuffs...), c.Consumables...), s.logger())
}
