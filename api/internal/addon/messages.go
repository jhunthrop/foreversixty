// api/internal/addon/messages.go
// The Wave C typed inbox messages, added beside the existing queued
// builds (docs/superpowers/specs/2026-09-28-addon-character-aware-
// design.md §3, §4; wire contract in the Sept 28 amendment to
// docs/superpowers/specs/2026-09-14-phase-3-interfaces.md): a
// character's private guild state, its saved stat weights and which of
// them are already capped, and its best gear upgrade from its last Top
// Gear run.
package addon

import (
	"context"
	"sort"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

const (
	messageUpgrade = "upgrade"
	messageWeights = "weights"
	messageGuild   = "guild"
)

// WeightEntry is one stat's weight inside a "weights" message.
type WeightEntry struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
}

// MessageEntry is one typed inbox message beyond a build. Field-for-
// field, this is the wire contract the Sept 28 amendment to
// docs/superpowers/specs/2026-09-14-phase-3-interfaces.md names, and
// what companion/internal/addon.Message decodes into — the API and the
// companion are separate Go modules with no shared package, so the JSON
// tags are what keeps the two in step, not the Go type.
type MessageEntry struct {
	Type      string `json:"type"`
	Character string `json:"character,omitempty"`

	// MessageUpgrade
	Slot     string  `json:"slot,omitempty"`
	ItemID   int     `json:"item_id,omitempty"`
	ItemName string  `json:"item_name,omitempty"`
	Source   string  `json:"source,omitempty"`
	Delta    float64 `json:"delta,omitempty"`

	// MessageWeights
	Spec    string        `json:"spec,omitempty"`
	Weights []WeightEntry `json:"weights,omitempty"`
	Caps    []string      `json:"caps,omitempty"`

	// MessageGuild
	GuildName        string `json:"guild_name,omitempty"`
	ClaimState       string `json:"claim_state,omitempty"`
	PendingApprovals int    `json:"pending_approvals,omitempty"`
	Rank             string `json:"rank,omitempty"`
}

// GuildSource is enough of guilds.Store to build the "guild" message.
// guilds.Store satisfies it.
type GuildSource interface {
	CharacterGuildState(ctx context.Context, characterKey string) (guilds.CharacterGuildState, bool, error)
}

// SimSource is enough of sims.Store to build the "upgrade" and
// "weights" messages: the caller's own history, one page at a time, and
// one saved result at a time. sims.Store satisfies it.
type SimSource interface {
	Mine(ctx context.Context, userID int64, page int, kind string) (sims.Page, error)
	Get(ctx context.Context, id string) (simapi.SimResult, string, error)
}

// capEligibleStats are the rating stats a Classic sweep can hard-cap:
// past a threshold rating, more of the stat cannot matter (it cannot
// reduce a miss/dodge/parry/block chance below what the target's own
// avoidance already floors it at). This is the closed set weightCaps
// ever names — see its own comment for what signals a cap.
var capEligibleStats = map[string]bool{"hit": true, "expertise": true}

// weightCaps reads which of a saved weights run's stats the engine's
// own sweep already reports as capped: Insignificant with a Weight and
// Error both exactly zero is sim/api/weights.go's documented signal
// that the sweep skipped the stat entirely because the character was
// already past where more of it can matter — not a new server-computed
// rating breakpoint. Only capEligibleStats are ever named: a stat can
// read Insignificant/zero for the unrelated reason its own error swamps
// a genuinely small weight, and that is not a cap.
func weightCaps(weights []simapi.StatWeight) []string {
	var caps []string
	for _, w := range weights {
		if capEligibleStats[w.Stat] && w.Insignificant && w.Weight == 0 && w.Error == 0 {
			caps = append(caps, w.Stat)
		}
	}
	sort.Strings(caps)
	return caps
}

// characterKeysByNameSlug maps every one of this user's own characters'
// name slugs to its full character key, for resolveCharacter below.
func characterKeysByNameSlug(exports []Export) map[string]string {
	out := make(map[string]string, len(exports))
	for _, e := range exports {
		ruleset := character.RulesetFromRealm(e.Ruleset, "")
		out[character.Slug(e.Name)] = character.Key(e.Region, ruleset, e.Name)
	}
	return out
}

// resolveCharacter finds which of a user's own characters a saved sim's
// request was about, by name alone. SimRequest.Character carries no
// region or ruleset (sim/api/envelope.go's CharacterSpec) and a sim is
// never itself tied to a character_key (sims.Row has none), so this is
// a best effort: a name that matches none of the account's characters
// leaves the message unaddressed — reaching every character on the
// account, exactly like a build with no `character` — rather than
// guessing wrong.
func resolveCharacter(characterName string, byNameSlug map[string]string) string {
	if characterName == "" {
		return ""
	}
	return byNameSlug[character.Slug(characterName)]
}

// firstDone is the first row a page of a user's sim history offers that
// actually finished: a queued, running or errored row has no Weights or
// Combos worth reading, and Mine's page is already newest-first, so
// this is "the most recent usable run" without a second query.
func firstDone(rows []sims.Row) *sims.Row {
	for i := range rows {
		if rows[i].State == sims.StateDone {
			return &rows[i]
		}
	}
	return nil
}

// upgradeMessage is the account's single best upgrade from its most
// recent finished Top Gear run, best-effort addressed to one of the
// account's own characters (resolveCharacter).
//
// Scope note: the design speaks of "the five best upgrades"; this reads
// one, from the best-ranked combo whose substitution list changes
// exactly one slot — the case a Top Gear "best item for this slot"
// report is actually answering. A combo that regears several slots at
// once has no single per-slot delta to report (Combo carries only the
// whole combination's Delta), so it is skipped rather than misreported.
// Widening this to several messages is a producer-only change; the
// message shape does not move.
func (s *Service) upgradeMessage(ctx context.Context, userID int64, byNameSlug map[string]string) *MessageEntry {
	page, err := s.Sims.Mine(ctx, userID, 1, simapi.KindGear)
	if err != nil || len(page.Rows) == 0 {
		return nil
	}
	row := firstDone(page.Rows)
	if row == nil {
		return nil
	}
	res, _, err := s.Sims.Get(ctx, row.SimID)
	if err != nil {
		return nil
	}
	for _, combo := range res.Combos {
		var item *simapi.Substitution
		itemSubs := 0
		for i := range combo.Substitutions {
			if combo.Substitutions[i].Kind == simapi.SubstitutionItem {
				item = &combo.Substitutions[i]
				itemSubs++
			}
		}
		if itemSubs != 1 || item.Slot == "" || item.ItemID == 0 {
			continue
		}
		return &MessageEntry{
			Type: messageUpgrade, Character: resolveCharacter(res.Request.Character.Name, byNameSlug),
			Slot: item.Slot, ItemID: item.ItemID, ItemName: item.Name,
			Source: item.SourceName, Delta: combo.Delta.Mean,
		}
	}
	return nil
}

// weightsMessage is the account's saved stat weights from its most
// recent finished weights run, with weightCaps' capped-stat list,
// best-effort addressed the same way upgradeMessage is.
func (s *Service) weightsMessage(ctx context.Context, userID int64, byNameSlug map[string]string) *MessageEntry {
	page, err := s.Sims.Mine(ctx, userID, 1, simapi.KindWeights)
	if err != nil || len(page.Rows) == 0 {
		return nil
	}
	row := firstDone(page.Rows)
	if row == nil {
		return nil
	}
	res, _, err := s.Sims.Get(ctx, row.SimID)
	if err != nil || len(res.Weights) == 0 {
		return nil
	}
	weights := make([]WeightEntry, len(res.Weights))
	for i, w := range res.Weights {
		weights[i] = WeightEntry{Stat: w.Stat, Weight: w.Weight}
	}
	return &MessageEntry{
		Type: messageWeights, Character: resolveCharacter(res.Request.Character.Name, byNameSlug),
		Spec: res.Request.Spec, Weights: weights, Caps: weightCaps(res.Weights),
	}
}

// guildMessages is one "guild" message per character of this user's own
// that currently belongs to a guild.
func (s *Service) guildMessages(ctx context.Context, byNameSlug map[string]string) []MessageEntry {
	var out []MessageEntry
	for _, key := range byNameSlug {
		state, ok, err := s.Guilds.CharacterGuildState(ctx, key)
		if err != nil {
			s.logger().Warn("addon", "op", "messages_guild", "character_key", key, "err", err)
			continue
		}
		if !ok {
			continue
		}
		out = append(out, MessageEntry{
			Type: messageGuild, Character: key,
			GuildName: state.Name, ClaimState: state.Claim.State,
			PendingApprovals: state.PendingApprovals, Rank: state.Rank,
		})
	}
	return out
}

// messages builds every typed inbox message this user's account is
// entitled to: guild state per character, and one upgrade and one
// weights message from the account's saved sims. Guilds and Sims are
// each optional (Service's own doc comment); with either or both nil,
// or any one lookup's own error, the inbox still answers with whatever
// it could build — never a failed sync over an advanced feature a build
// or two never even reaches.
func (s *Service) messages(ctx context.Context, userID int64) []MessageEntry {
	exports, err := s.Store.Exports(ctx, userID)
	if err != nil {
		s.logger().Warn("addon", "op", "messages", "err", err)
		return nil
	}
	byNameSlug := characterKeysByNameSlug(exports)

	var out []MessageEntry
	if s.Guilds != nil {
		out = append(out, s.guildMessages(ctx, byNameSlug)...)
	}
	if s.Sims != nil {
		if m := s.upgradeMessage(ctx, userID, byNameSlug); m != nil {
			out = append(out, *m)
		}
		if m := s.weightsMessage(ctx, userID, byNameSlug); m != nil {
			out = append(out, *m)
		}
	}
	return out
}
