// api/internal/guilds/raids_report_pulls.go
//
// The report's headline figures, its pull list with death causes and its consumables by pull
// (spec 2026-11-05-web-layout §7.2 "Raid report"). Pulls and figures name nobody and everyone
// reads them; a death and a consumables row name a raider and carry the own flag ownOnly
// narrows a member's report by.
package guilds

import (
	"context"
	"fmt"
	"strings"
)

// Pull results.
const (
	PullKill  = "kill"
	PullWipe  = "wipe"
	PullClear = "clear" // a trash pull the raid won
)

// ConsumablesNoSheet is a pull's consumables reason when the log carried no character sheet
// (COMBATANT_INFO) for anyone: what was up at the pull cannot be read. ConsumablesNotRead is a
// pull whose stored summary has not been read yet.
const (
	ConsumablesNoSheet = "no_character_sheet"
	ConsumablesNotRead = "not_read"
)

// ReportHeadline is the night's answer figures. Attempts are boss pulls (kills plus wipes);
// deaths count every pull, trash included, so they add up to the pull list.
type ReportHeadline struct {
	Kills    int `json:"kills"`
	Wipes    int `json:"wipes"`
	Attempts int `json:"attempts"`
	Deaths   int `json:"deaths"`
}

// PullDeathView is one death on a pull: who, how far into the pull, and what killed them.
// KillingBlow is null when the log showed no hit before the death.
type PullDeathView struct {
	CharacterKey string     `json:"character_key"`
	Name         string     `json:"name"`
	Class        string     `json:"class"`
	AtMS         int64      `json:"at_ms"`
	KillingBlow  *DeathHit  `json:"killing_blow"`
	LastHits     []DeathHit `json:"last_hits"`
	own          bool
}

// PullView is one fight of the night. Attempt is the boss pull's number (null on trash).
// DeathCauses is null until the fight's summary has been read (then Deaths still counts them).
type PullView struct {
	FightIndex  int             `json:"fight_index"`
	EncounterID *int            `json:"encounter_id"`
	Name        string          `json:"name"`
	Attempt     *int            `json:"attempt"`
	Result      string          `json:"result"`
	DurationMS  int64           `json:"duration_ms"`
	Deaths      int             `json:"deaths"`
	DeathCauses []PullDeathView `json:"death_causes"`
}

// PullConsumableRow is one raider's preparation at a boss pull: true up, false proven missing,
// null when the log cannot say. Brought names what was up.
type PullConsumableRow struct {
	CharacterKey  string   `json:"character_key"`
	Name          string   `json:"name"`
	Class         string   `json:"class"`
	Flask         *bool    `json:"flask"`
	Food          *bool    `json:"food"`
	WeaponEnchant *bool    `json:"weapon_enchant"`
	Potions       int64    `json:"potions"`
	Brought       []string `json:"brought"`
	own           bool
}

// PullConsumablesView is one boss pull's consumables. Rows is null with a Reason when the pull
// cannot be read.
type PullConsumablesView struct {
	FightIndex int                 `json:"fight_index"`
	Name       string              `json:"name"`
	Attempt    int                 `json:"attempt"`
	Reason     *string             `json:"reason"`
	Rows       []PullConsumableRow `json:"rows"`
}

// pullKeys maps a fight's log names to the night's character keys and classes.
type pullKeys map[int]map[string]keyClass

type keyClass struct{ key, class string }

func (s *Store) reportPullKeys(ctx context.Context, reportID string) (pullKeys, error) {
	rows, err := s.Pool.Query(ctx, `
		select fight_index, player_name, player_key, coalesce(lower(class), '')
		from fight_metrics where report_id = $1 and state <> 'removed'`, reportID)
	if err != nil {
		return nil, fmt.Errorf("guilds: report pull keys: %w", err)
	}
	defer rows.Close()
	out := pullKeys{}
	for rows.Next() {
		var n int
		var name string
		var kc keyClass
		if err := rows.Scan(&n, &name, &kc.key, &kc.class); err != nil {
			return nil, fmt.Errorf("guilds: report pull keys: %w", err)
		}
		if out[n] == nil {
			out[n] = map[string]keyClass{}
		}
		out[n][name] = kc
	}
	return out, rows.Err()
}

// fillPulls adds the headline, the pulls and the consumables by pull to rep.
func (s *Store) fillPulls(ctx context.Context, vc raidViewContext, rep *ReportView, reportID string, fights []reportFight) error {
	digests, err := s.pullDigests(ctx, reportID)
	if err != nil {
		return err
	}
	keys, err := s.reportPullKeys(ctx, reportID)
	if err != nil {
		return err
	}
	rep.Headline = headlineOf(fights)
	rep.Pulls = make([]PullView, 0, len(fights))
	rep.Consumables.ByPull = []PullConsumablesView{}
	attempts := map[int]int{}
	for _, f := range fights {
		p := pullOf(f, attempts)
		d := digests[f.index]
		if d.Read {
			p.DeathCauses = deathViews(d.Deaths, keys[f.index], vc)
		}
		rep.Pulls = append(rep.Pulls, p)
		if p.Attempt != nil {
			bp := consumablesOfPull(p, d, keys[f.index], vc)
			rep.Consumables.LogHasBuffs = rep.Consumables.LogHasBuffs || bp.Rows != nil
			rep.Consumables.ByPull = append(rep.Consumables.ByPull, bp)
		}
	}
	return nil
}

func headlineOf(fights []reportFight) ReportHeadline {
	var h ReportHeadline
	for _, f := range fights {
		h.Deaths += f.deaths
		if f.encounterID == nil {
			continue
		}
		h.Attempts++
		if f.kill {
			h.Kills++
		} else {
			h.Wipes++
		}
	}
	return h
}

// pullOf is one fight's row; attempts counts the boss pulls seen so far per encounter.
func pullOf(f reportFight, attempts map[int]int) PullView {
	p := PullView{FightIndex: f.index, EncounterID: f.encounterID, Name: f.name, DurationMS: f.durationMS, Deaths: f.deaths}
	switch {
	case f.encounterID == nil && f.kill:
		p.Result = PullClear
	case f.kill:
		p.Result = PullKill
	default:
		p.Result = PullWipe
	}
	if f.encounterID != nil {
		attempts[*f.encounterID]++
		n := attempts[*f.encounterID]
		p.Attempt = &n
	}
	return p
}

func deathViews(deaths []digestDeath, keys map[string]keyClass, vc raidViewContext) []PullDeathView {
	out := make([]PullDeathView, 0, len(deaths))
	for _, d := range deaths {
		kc := keys[d.Name]
		class := kc.class
		if class == "" {
			class = strings.ToLower(d.Class)
		}
		v := PullDeathView{CharacterKey: kc.key, Name: d.Name, Class: class, AtMS: d.AtMS, KillingBlow: d.KillingBlow,
			LastHits: d.Last, own: kc.key != "" && vc.ownsKey(kc.key)}
		if v.LastHits == nil {
			v.LastHits = []DeathHit{}
		}
		out = append(out, v)
	}
	return out
}

func consumablesOfPull(p PullView, d storedDigest, keys map[string]keyClass, vc raidViewContext) PullConsumablesView {
	out := PullConsumablesView{FightIndex: p.FightIndex, Name: p.Name, Attempt: *p.Attempt}
	switch {
	case !d.Read:
		out.Reason = textPtr(ConsumablesNotRead)
		return out
	case d.Consumables == nil:
		out.Reason = textPtr(ConsumablesNoSheet)
		return out
	}
	out.Rows = []PullConsumableRow{}
	for _, c := range d.Consumables {
		kc := keys[c.Name]
		row := PullConsumableRow{CharacterKey: kc.key, Name: c.Name, Class: kc.class, Flask: c.Flask, Food: c.Food,
			WeaponEnchant: c.WeaponEnchant, Potions: c.Potions, Brought: c.Brought, own: kc.key != "" && vc.ownsKey(kc.key)}
		if row.Brought == nil {
			row.Brought = []string{}
		}
		out.Rows = append(out.Rows, row)
	}
	return out
}
