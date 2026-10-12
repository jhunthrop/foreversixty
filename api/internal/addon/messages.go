// api/internal/addon/messages.go
// The Wave C typed inbox messages, added beside the existing queued
// builds (docs/superpowers/specs/2026-09-28-addon-character-aware-
// design.md §3, §4; wire contract in the Sept 28 amendment to
// docs/superpowers/specs/2026-09-14-phase-3-interfaces.md): a
// character's private guild state, its saved stat weights and which of
// them are already capped, and its best gear upgrade from its last Top
// Gear run.
//
// Inbox v2 (docs/contracts/addon-v2.md contract 3, the companion's
// ForeverSixtyInbox format 2) adds the build, sim_result, last_raid,
// loot_values and ack messages and two fields on the old ones: unit on
// upgrade, dps_per_point on weights. Every message is addressed by
// character key, region/ruleset/name-slug, the form the addon's
// sameCharacter matches.
package addon

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	simapi "github.com/jhunthrop/foreversixty/sim/api"
)

// InboxFormat is contract 3's format: the version GET /v1/addon/inbox answers
// and the companion writes into ForeverSixtyInbox.
const InboxFormat = 2

const (
	messageBuild      = "build"
	messageUpgrade    = "upgrade"
	messageWeights    = "weights"
	messageSimResult  = "sim_result"
	messageLastRaid   = "last_raid"
	messageLootValues = "loot_values"
	messageGuild      = "guild"
	messageAck        = "ack"
)

// InboxResponse is the body GET /v1/addon/inbox answers with (contract 3; the
// fixture docs/contracts/fixtures/inbox-v2.json): the format, when the API
// built it, the queued builds as they have always been listed (kept for one
// release beside the build messages that repeat them), and every message.
type InboxResponse struct {
	Format      int            `json:"format"`
	GeneratedAt string         `json:"generated_at"`
	Builds      []InboxEntry   `json:"builds"`
	Messages    []MessageEntry `json:"messages"`
}

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
//
// It now carries every field of contract 3 (docs/contracts/addon-v2.md), each
// omitted from the types that do not have it; contracts_test.go decodes the
// shared fixture into it strictly.
type MessageEntry struct {
	Type      string `json:"type"`
	Character string `json:"character,omitempty"`

	// messageBuild
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`

	// messageUpgrade
	Slot     string  `json:"slot,omitempty"`
	ItemID   int     `json:"item_id,omitempty"`
	ItemName string  `json:"item_name,omitempty"`
	Source   string  `json:"source,omitempty"`
	Delta    float64 `json:"delta,omitempty"`
	// Unit is what Delta is measured in, sims.UnitDPS or sims.UnitScore.
	Unit string `json:"unit,omitempty"`

	// messageWeights (Spec also names a sim_result's spec)
	Spec    string        `json:"spec,omitempty"`
	Weights []WeightEntry `json:"weights,omitempty"`
	Caps    []string      `json:"caps,omitempty"`
	// DPSPerPoint is the DPS one point of the weight-1 stat is worth for
	// Spec (sims.DPSPerPoint); nil, and absent, when the data has none.
	DPSPerPoint *float64 `json:"dps_per_point,omitempty"`

	// messageSimResult
	// DPS and Error are pointers so a sim_result always carries both, 0.0 included,
	// while every other type carries neither.
	DPS   *float64 `json:"dps,omitempty"`
	Error *float64 `json:"error,omitempty"`
	// GeneratedAt is when the run was queued, RFC 3339 UTC: a sim_result's, and a weights
	// message's (the run that made those weights).
	GeneratedAt string `json:"generated_at,omitempty"`
	// Slots is a pointer so a sim_result carries an empty list ("slots": [])
	// while every other type carries no slots at all.
	Slots *[]SimSlot `json:"slots,omitempty"`

	// messageLastRaid
	ReportID string         `json:"report_id,omitempty"`
	Date     string         `json:"date,omitempty"`
	Bosses   []LastRaidBoss `json:"bosses,omitempty"`

	// messageLootValues
	Raid  string                `json:"raid,omitempty"`
	Items map[int]LootItemValue `json:"items,omitempty"`

	// messageGuild
	GuildName        string `json:"guild_name,omitempty"`
	ClaimState       string `json:"claim_state,omitempty"`
	PendingApprovals int    `json:"pending_approvals,omitempty"`
	Rank             string `json:"rank,omitempty"`

	// messageAck: when the site stored the character's export. On a weights message from an
	// in-game run it is the export that request was made on (the ack's value back then);
	// absent on a website run and a request recorded before migration 0038.
	ExportSeenAt string `json:"export_seen_at,omitempty"`
	// SimRequestState is what became of the character's newest in-game sim
	// request (sims.SimRequestState*); absent when it never asked.
	SimRequestState string `json:"sim_request_state,omitempty"`
	// RequestedAt is when the site received that request (addon_sim_requests.created_at), RFC
	// 3339 UTC; absent when it never asked. The player's own wall time stays in the outbox.
	RequestedAt string `json:"requested_at,omitempty"`
}

// SimSlot is one slot's DPS change inside a sim_result.
type SimSlot struct {
	Slot   string  `json:"slot"`
	ItemID int     `json:"item_id"`
	Delta  float64 `json:"delta"`
}

// LastRaidBoss is one rated boss fight inside a last_raid message: the card's
// overall as the percentile, and the facts the rater stored with it
// (api/internal/rating/last_raid.go). Adherence is the rotation adherence the rater
// stored (api/internal/rating/adherence.go), absent when it could not be worked out and
// for a fight rated before migration 0037.
type LastRaidBoss struct {
	Name              string   `json:"name"`
	Percentile        int      `json:"percentile"`
	BiggestLoss       string   `json:"biggest_loss,omitempty"`
	LossSeconds       int      `json:"loss_seconds"`
	Deaths            int      `json:"deaths"`
	ConsumablesMissed []string `json:"consumables_missed"`
	Adherence         *float64 `json:"adherence,omitempty"`
}

// LootItemValue is one item's worth inside a loot_values message.
type LootItemValue struct {
	Value   float64 `json:"value"`
	Unit    string  `json:"unit"`
	Planned bool    `json:"planned"`
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
			Source: item.SourceName, Delta: combo.Delta.Mean, Unit: sims.UnitDPS,
		}
	}
	return nil
}

// weightsMessage is the fallback weights message: the account's saved stat weights from its
// most recent finished website weights run (an in-game request's run is not one: the
// character it was for gets its own message from inGameWeightsMessages), best-effort
// addressed the same way upgradeMessage is, with generated_at the run's queue time and no
// export_seen_at, since a website run names no export. It reads the newest page of the
// account's weights runs, so a page of nothing but in-game runs gives no fallback.
func (s *Service) weightsMessage(ctx context.Context, userID int64, byNameSlug map[string]string, exports map[string]Export) *MessageEntry {
	page, err := s.Sims.Mine(ctx, userID, 1, simapi.KindWeights)
	if err != nil {
		return nil
	}
	for _, row := range page.Rows {
		if row.State != sims.StateDone {
			continue
		}
		res, _, err := s.Sims.Get(ctx, row.SimID)
		if err != nil {
			return nil
		}
		if res.Request.Source.Kind == simapi.SourceAddon {
			continue
		}
		if len(res.Weights) == 0 {
			return nil
		}
		m := weightsEntry(resolveCharacter(res.Request.Character.Name, byNameSlug), res.Request.Spec, res.Weights, exports)
		if !row.CreatedAt.IsZero() {
			m.GeneratedAt = row.CreatedAt.UTC().Format(time.RFC3339)
		}
		return m
	}
	return nil
}

// inGameWeightsMessages is one weights message per character whose in-game request has a
// finished stat-weights run (the newest such request): that run's weights, generated_at its
// queue time, and export_seen_at the export the request was made on, when it was recorded
// (migration 0038).
func (s *Service) inGameWeightsMessages(ctx context.Context, userID int64, exports map[string]Export) []MessageEntry {
	rows, err := s.Store.inGameWeights(ctx, userID)
	if err != nil {
		s.logger().Warn("addon", "op", "messages_weights", "err", err)
		return nil
	}
	var out []MessageEntry
	for _, r := range rows {
		if len(r.Weights) == 0 {
			continue
		}
		m := weightsEntry(r.CharacterKey, r.Spec, r.Weights, exports)
		m.GeneratedAt = r.CreatedAt.UTC().Format(time.RFC3339)
		if r.ExportSeenAt != nil {
			m.ExportSeenAt = r.ExportSeenAt.UTC().Format(time.RFC3339)
		}
		out = append(out, *m)
	}
	return out
}

// weightsEntry is a weights message for key from a run's weights: weightCaps' capped-stat
// list, and the spec's DPS per point of its weight-1 stat (sims.DPSPerPoint) in the data build
// the addressed character's export names, or the newest loaded build when the run names none
// of the account's characters - present only when that weight-1 stat is the file's own
// reference.
func weightsEntry(key, spec string, ws []simapi.StatWeight, exports map[string]Export) *MessageEntry {
	weights := make([]WeightEntry, len(ws))
	for i, w := range ws {
		weights[i] = WeightEntry{Stat: w.Stat, Weight: w.Weight}
	}
	m := &MessageEntry{
		Type: messageWeights, Character: key, Spec: spec, Weights: weights, Caps: weightCaps(ws),
	}
	build := ""
	if e, ok := exports[m.Character]; ok {
		if d, ok := fs1.Decode(e.Export); ok {
			build = d.DataBuild
		}
	}
	// Only when the run is normalised to the file's own reference stat: a weight in some
	// other stat's points times this figure would be no number the addon can use.
	runWeights := make(map[string]float64, len(weights))
	for _, w := range weights {
		runWeights[w.Stat] = w.Weight
	}
	if p, ok := sims.DPSPerPoint(build, m.Spec); ok && p.Scales(runWeights) {
		dpp := p.DPS
		m.DPSPerPoint = &dpp
	}
	return m
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
// entitled to beyond the build messages the inbox makes from its queued
// builds (buildMessages): guild state per character; one upgrade message from the
// account's saved sims; one weights message per character (its newest finished in-game
// run, else the account's website run when that names it); and per character its
// newest finished in-game sim, its last raid, that raid's loot priced for
// it, and (with a companion paired) the ack. Guilds and Sims are each
// optional (Service's own doc comment); with either or both nil, or any
// one lookup's own error, the inbox still answers with whatever it could
// build — never a failed sync over an advanced feature a build or two
// never even reaches. now is the moment the last-raid window counts back
// from.
func (s *Service) messages(ctx context.Context, userID int64, now time.Time) []MessageEntry {
	out := []MessageEntry{}
	exports, err := s.Store.Exports(ctx, userID)
	if err != nil {
		s.logger().Warn("addon", "op", "messages", "err", err)
		return out
	}
	byNameSlug := characterKeysByNameSlug(exports)
	byKey := exportsByKey(exports)

	if s.Guilds != nil {
		out = append(out, s.guildMessages(ctx, byNameSlug)...)
	}
	// One weights message per character: its own in-game run first, else the account's
	// website run when that names a character no in-game run covers.
	weights := s.inGameWeightsMessages(ctx, userID, byKey)
	if s.Sims != nil {
		if m := s.upgradeMessage(ctx, userID, byNameSlug); m != nil {
			out = append(out, *m)
		}
		if m := s.weightsMessage(ctx, userID, byNameSlug, byKey); m != nil && weightsFor(weights, m.Character) == nil {
			weights = append(weights, *m)
		}
	}
	out = append(out, weights...)
	out = append(out, s.simResultMessages(ctx, userID)...)
	lastRaids, raids := s.lastRaidMessages(ctx, sortedExportKeys(byKey), now)
	out = append(out, lastRaids...)
	out = append(out, s.lootValuesMessages(ctx, raids, byKey, weights)...)
	out = append(out, s.ackMessages(ctx, userID)...)
	return out
}

// exportsByKey indexes a user's exports by character key.
func exportsByKey(exports []Export) map[string]Export {
	out := make(map[string]Export, len(exports))
	for _, e := range exports {
		out[character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name)] = e
	}
	return out
}

// sortedExportKeys is exports' character keys in order.
func sortedExportKeys(exports map[string]Export) []string {
	keys := make([]string, 0, len(exports))
	for k := range exports {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ---------------------------------------------------------------- inbox v2 builders --

// buildMessages is one "build" message per queued build the companion can write: the same
// character, name and code as its builds entry (contract 3 keeps both for one release).
func buildMessages(entries []InboxEntry) []MessageEntry {
	var out []MessageEntry
	for _, e := range entries {
		if e.Code == "" {
			continue
		}
		out = append(out, MessageEntry{Type: messageBuild, Character: e.Character, Name: e.Name, Code: e.Code})
	}
	return out
}

// simResultMessages is one "sim_result" per character with a finished in-game sim
// request: the newest one's spec, DPS, error and time.
//
// slots: RequestFromAddon (api/internal/sims/request_job.go) queues a stat-weights run, a
// SimRequest with no bulk section, so its stored result has no combinations and slots is
// an empty list; the addon prices a slot itself from the weights message (dps_per_point
// times the weight delta). Should the request ever queue a Top Gear run, slotDeltas reads
// its single-slot combinations here without a change to the message.
func (s *Service) simResultMessages(ctx context.Context, userID int64) []MessageEntry {
	rows, err := s.Store.simResults(ctx, userID)
	if err != nil {
		s.logger().Warn("addon", "op", "messages_sim_result", "err", err)
		return nil
	}
	out := make([]MessageEntry, 0, len(rows))
	for _, r := range rows {
		slots := slotDeltas(r.Combos)
		out = append(out, MessageEntry{
			Type: messageSimResult, Character: r.CharacterKey, Spec: r.Spec,
			DPS: ptr(round1(r.DPS)), Error: ptr(round1(r.Error)),
			GeneratedAt: r.CreatedAt.UTC().Format(time.RFC3339), Slots: &slots,
		})
	}
	return out
}

// slotDeltas is a stored bulk result's per-slot figures: every combination that changes
// exactly one item, best first, one per slot (upgradeMessage's own rule). Empty, not nil,
// for a plain run.
func slotDeltas(raw json.RawMessage) []SimSlot {
	slots := []SimSlot{}
	if len(raw) == 0 || string(raw) == "null" {
		return slots
	}
	var combos []simapi.Combo
	if err := json.Unmarshal(raw, &combos); err != nil {
		return slots
	}
	seen := map[string]bool{}
	for _, combo := range combos {
		var item *simapi.Substitution
		items := 0
		for i := range combo.Substitutions {
			if combo.Substitutions[i].Kind == simapi.SubstitutionItem {
				item = &combo.Substitutions[i]
				items++
			}
		}
		if items != 1 || item.Slot == "" || item.ItemID == 0 || seen[item.Slot] {
			continue
		}
		seen[item.Slot] = true
		slots = append(slots, SimSlot{Slot: item.Slot, ItemID: item.ItemID, Delta: round1(combo.Delta.Mean)})
	}
	return slots
}

// lastRaid is one character's last raid, as lastRaidMessages found it, for the loot
// values message that follows it.
type lastRaid struct {
	key, zone, spec string
}

// lastRaidMessages is one "last_raid" per character of the account with a rated fight in
// a report that is not private, fought within LastRaidWindow: the newest such report, every
// rated boss fight in it (kills and wipes), with the card's own overall as the
// percentile and the facts the rater stored with it (api/internal/rating/last_raid.go).
// A character with no such report gets no message. It also answers each character's raid
// zone and spec, for lootValuesMessages.
func (s *Service) lastRaidMessages(ctx context.Context, keys []string, now time.Time) ([]MessageEntry, []lastRaid) {
	if len(keys) == 0 {
		return nil, nil
	}
	fights, err := s.Store.lastRaidFights(ctx, keys, now.Add(-LastRaidWindow))
	if err != nil {
		s.logger().Warn("addon", "op", "messages_last_raid", "err", err)
		return nil, nil
	}
	var out []MessageEntry
	var raids []lastRaid
	for _, f := range fights {
		if len(out) == 0 || out[len(out)-1].Character != f.CharacterKey {
			out = append(out, MessageEntry{Type: messageLastRaid, Character: f.CharacterKey, ReportID: f.ReportID})
			raids = append(raids, lastRaid{key: f.CharacterKey, zone: f.Zone})
		}
		m := &out[len(out)-1]
		// Fights arrive oldest first, so the last one names the date and the spec.
		m.Date = f.FoughtAt.UTC().Format(time.DateOnly)
		raids[len(raids)-1].spec = rating.DataSpecSlug(f.Class, f.Spec)
		if f.Boss == "" {
			continue
		}
		missed := f.ConsumablesMissed
		if missed == nil {
			missed = []string{}
		}
		m.Bosses = append(m.Bosses, LastRaidBoss{
			Name: f.Boss, Percentile: int(math.Round(f.Overall)), BiggestLoss: f.BiggestLoss,
			LossSeconds: f.LossSeconds, Deaths: f.Deaths, ConsumablesMissed: missed, Adherence: f.Adherence,
		})
	}
	// A raid whose fights no log or fight row ever named has no boss to show.
	kept, keptRaids := out[:0], raids[:0]
	for i, m := range out {
		if len(m.Bosses) > 0 {
			kept, keptRaids = append(kept, m), append(keptRaids, raids[i])
		}
	}
	return kept, keptRaids
}

// lootValuesMessages is one "loot_values" per character with a last raid: the raid's loot
// priced for the character by sims.LootValues, under the character's own weights message
// when it has one (weights) and the spec's shipped weights otherwise. A raid whose zone names
// no raid in the build's loot table, a character with no export, or gear the loaded data
// cannot score gets no message.
func (s *Service) lootValuesMessages(ctx context.Context, raids []lastRaid, exports map[string]Export, all []MessageEntry) []MessageEntry {
	var out []MessageEntry
	for _, r := range raids {
		e, ok := exports[r.key]
		if !ok {
			continue
		}
		spec, runWeights := r.spec, map[string]float64(nil)
		if weights := weightsFor(all, r.key); weights != nil && weights.Spec != "" {
			spec = weights.Spec
			runWeights = make(map[string]float64, len(weights.Weights))
			for _, w := range weights.Weights {
				runWeights[w.Stat] = w.Weight
			}
		}
		if spec == "" {
			continue
		}
		set, ok, err := sims.LootValues(ctx, r.zone, e.Export, spec, runWeights)
		if err != nil {
			// Gear the data cannot score (a spec with no shipped weights, say) stays so until
			// a data deploy: said once per spec, not on every ten-minute poll.
			if !errors.Is(err, sims.ErrGearUnscorable) || warnOnce(spec) {
				s.logger().Warn("addon", "op", "messages_loot_values", "character_key", r.key, "spec", spec, "err", err)
			}
			continue
		}
		if !ok || len(set.Items) == 0 {
			continue
		}
		items := make(map[int]LootItemValue, len(set.Items))
		for id, v := range set.Items {
			items[id] = LootItemValue{Value: v.Value, Unit: v.Unit, Planned: v.Planned}
		}
		out = append(out, MessageEntry{Type: messageLootValues, Character: r.key, Raid: set.Raid, Items: items})
	}
	return out
}

// ackMessages is one "ack" per exported character when the account has a companion paired:
// when the site stored the export, what became of the character's newest in-game sim
// request (its run's done or failed once the run ended), and when the site received it.
func (s *Service) ackMessages(ctx context.Context, userID int64) []MessageEntry {
	rows, err := s.Store.acks(ctx, userID)
	if err != nil {
		s.logger().Warn("addon", "op", "messages_ack", "err", err)
		return nil
	}
	out := make([]MessageEntry, 0, len(rows))
	for _, r := range rows {
		m := MessageEntry{
			Type: messageAck, Character: r.CharacterKey,
			ExportSeenAt: r.SeenAt.UTC().Format(time.RFC3339), SimRequestState: r.SimRequestState,
		}
		if r.RequestedAt != nil {
			m.RequestedAt = r.RequestedAt.UTC().Format(time.RFC3339)
		}
		out = append(out, m)
	}
	return out
}

// weightsFor is the weights message among all addressed to key, or nil.
func weightsFor(all []MessageEntry, key string) *MessageEntry {
	for i := range all {
		if all[i].Type == messageWeights && all[i].Character == key {
			return &all[i]
		}
	}
	return nil
}

// unscorableWarned is the specs lootValuesMessages has already warned about.
var unscorableWarned sync.Map

// warnOnce reports whether this is the first warning for spec in this process.
func warnOnce(spec string) bool {
	_, seen := unscorableWarned.LoadOrStore(spec, true)
	return !seen
}

func ptr(v float64) *float64 { return &v }

// round1 rounds a figure the addon prints to a tenth.
func round1(v float64) float64 { return math.Round(v*10) / 10 }
