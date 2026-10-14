// companion/internal/addon/addon.go
// Package addon carries strings between the game and the site: the
// addon's character exports and its outbox out of SavedVariables and
// up to the API, and the builds and messages the site sends down into
// an inbox file the addon reads at load.
//
// The export strings are opaque here. FS1 is the addon's format and
// the site's; the companion moves it and never inspects it.
package addon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SavedVariablesName is the addon's saved-variables file.
const SavedVariablesName = "ForeverSixty"

// InboxName is the file the companion writes beside it.
const InboxName = "ForeverSixtyInbox"

// InboxGlobal is the global the inbox file assigns.
const InboxGlobal = "ForeverSixtyInbox"

// Export is one character's export string, ready to post. Ruleset is
// the second segment of the character key; whatever the addon wrote
// there is passed through untouched, because the beta log has not yet
// settled what Forever puts in it.
type Export struct {
	Name    string `json:"name"`
	Ruleset string `json:"ruleset"`
	Region  string `json:"region"`
	Export  string `json:"export"`
}

// Build is one build the site sent down for the addon to load.
type Build struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character,omitempty"`
	Code      string `json:"code"`
}

// InboxFormat is the inbox format this RenderInbox writes into every
// file as its top-level "format" field: contract 3 of
// docs/contracts/addon-v2.md (ForeverSixtyInbox format 2). The addon
// keys on it and ignores a file whose format is newer than it knows;
// within a format, forward compatibility comes from Message.Type (an
// addon build that does not know a Type skips that one message and
// keeps reading the rest). It replaces the informational "version"
// field the Wave C writer put in the same place, which nothing read.
const InboxFormat = 2

// The inbox message types contract 3 names. An addon build that
// predates a type skips any message carrying it.
const (
	MessageBuild      = "build"
	MessageUpgrade    = "upgrade"
	MessageWeights    = "weights"
	MessageSimResult  = "sim_result"
	MessageLastRaid   = "last_raid"
	MessageLootValues = "loot_values"
	MessageGuild      = "guild"
	MessageAck        = "ack"
)

// knownMessage is the set of types decoded field by field. Any other
// type is carried through as-is only when it is flat (decodeMessage).
var knownMessage = map[string]bool{
	MessageBuild: true, MessageUpgrade: true, MessageWeights: true, MessageSimResult: true,
	MessageLastRaid: true, MessageLootValues: true, MessageGuild: true, MessageAck: true,
}

// WeightEntry is one stat's weight inside a "weights" message.
type WeightEntry struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
}

// SimSlot is one slot's DPS change inside a "sim_result" message.
type SimSlot struct {
	Slot   string  `json:"slot"`
	ItemID int     `json:"item_id"`
	Delta  float64 `json:"delta"`
}

// LastRaidBoss is one rated boss fight inside a "last_raid" message.
// BiggestLoss is absent on a perfect card and Adherence until the
// rotation task lands; ConsumablesMissed is a list even when empty.
type LastRaidBoss struct {
	Name              string   `json:"name"`
	Percentile        int      `json:"percentile"`
	BiggestLoss       string   `json:"biggest_loss,omitempty"`
	LossSeconds       int      `json:"loss_seconds"`
	Deaths            int      `json:"deaths"`
	ConsumablesMissed []string `json:"consumables_missed"`
	Adherence         *float64 `json:"adherence,omitempty"`
}

// LootItemValue is one item's worth inside a "loot_values" message.
type LootItemValue struct {
	Value   float64 `json:"value"`
	Unit    string  `json:"unit"`
	Planned bool    `json:"planned"`
}

// Message is one typed inbox message (contract 3). Its fields are the
// API's MessageEntry (api/internal/addon/messages.go) JSON tag for JSON
// tag and in the same order, so the shared fixture round-trips byte for
// byte (contracts_test.go); the two are separate Go modules, so the tags
// are what keeps them in step.
//
// Type selects which of the fields below the addon should read; the
// rest are the zero value and simply not rendered (renderMessage).
// Character is optional exactly like Build.Character: empty reaches
// every character, addressed reaches one — the addon's
// Follow.sameCharacter matching rule, shared with builds.
type Message struct {
	Type      string `json:"type"`
	Character string `json:"character,omitempty"`

	// MessageBuild: one queued build, the same name and code as its
	// entry in Inbox.Builds.
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`

	// MessageUpgrade: the best upgrade Top Gear found for one slot,
	// Delta in Unit ("dps" or "score"; absent means "dps").
	Slot     string  `json:"slot,omitempty"`
	ItemID   int     `json:"item_id,omitempty"`
	ItemName string  `json:"item_name,omitempty"`
	Source   string  `json:"source,omitempty"`
	Delta    float64 `json:"delta,omitempty"`
	Unit     string  `json:"unit,omitempty"`

	// MessageWeights: the character's saved stat weights, which of
	// those stats are already capped, and what one point of the
	// weight-1 stat is worth. Spec also names a sim_result's spec.
	Spec        string        `json:"spec,omitempty"`
	Weights     []WeightEntry `json:"weights,omitempty"`
	Caps        []string      `json:"caps,omitempty"`
	DPSPerPoint *float64      `json:"dps_per_point,omitempty"`

	// MessageSimResult: the character's newest sim. DPS and Error are
	// always present on a sim_result, 0 included; Slots is a list even
	// when empty.
	DPS         *float64   `json:"dps,omitempty"`
	Error       *float64   `json:"error,omitempty"`
	GeneratedAt string     `json:"generated_at,omitempty"`
	Slots       *[]SimSlot `json:"slots,omitempty"`

	// MessageLastRaid: the character's newest rated raid.
	ReportID string         `json:"report_id,omitempty"`
	Date     string         `json:"date,omitempty"`
	Bosses   []LastRaidBoss `json:"bosses,omitempty"`

	// MessageLootValues: what each item of Raid is worth to the
	// character, keyed by item id.
	Raid  string                `json:"raid,omitempty"`
	Items map[int]LootItemValue `json:"items,omitempty"`

	// MessageGuild: the character's guild state block (design §3, §4):
	// claim state, the officer's pending-approval count (zero unless
	// this character is an officer or leader), and this character's
	// own rank.
	GuildName        string `json:"guild_name,omitempty"`
	ClaimState       string `json:"claim_state,omitempty"`
	PendingApprovals int    `json:"pending_approvals,omitempty"`
	Rank             string `json:"rank,omitempty"`

	// MessageAck: when the site last saw the character's export (RFC
	// 3339, UTC), and what became of its newest in-game sim request.
	ExportSeenAt    string `json:"export_seen_at,omitempty"`
	SimRequestState string `json:"sim_request_state,omitempty"`

	// Extra is every field of a message whose Type this build does not
	// know, carried through to the file as-is: each value a string or
	// a float64. It never travels as JSON.
	Extra map[string]any `json:"-"`
}

// Inbox is the body of GET /v1/addon/inbox (contract 3).
type Inbox struct {
	Format      int       `json:"format"`
	GeneratedAt string    `json:"generated_at"`
	Builds      []Build   `json:"builds"`
	Messages    []Message `json:"messages,omitempty"`
	// Skipped lists the messages that could not be carried to the
	// addon, and why. The sync logs each once per process.
	Skipped []SkippedMessage `json:"-"`
}

// SkippedMessage is one inbox message the companion could not carry:
// its 1-based position in the API's list, its type ("" when it had
// none) and why.
type SkippedMessage struct {
	Position int
	Type     string
	Why      string
}

// UnmarshalJSON reads the inbox message by message: a message of a
// type contract 3 names is decoded into Message's fields, one this
// build does not know is carried through as-is when it is a flat
// object of strings and numbers (the addon skips a type it does not
// know, so a new type never needs a new companion), and anything else
// is skipped and noted in Skipped. One bad message never costs the rest.
func (in *Inbox) UnmarshalJSON(b []byte) error {
	var wire struct {
		Format      int               `json:"format"`
		GeneratedAt string            `json:"generated_at"`
		Builds      []Build           `json:"builds"`
		Messages    []json.RawMessage `json:"messages"`
	}
	if err := json.Unmarshal(b, &wire); err != nil {
		return err
	}
	out := Inbox{Format: wire.Format, GeneratedAt: wire.GeneratedAt, Builds: wire.Builds}
	for i, raw := range wire.Messages {
		m, why := decodeMessage(raw)
		if why != "" {
			var head struct {
				Type string `json:"type"`
			}
			_ = json.Unmarshal(raw, &head) // a message with no readable type keeps ""
			out.Skipped = append(out.Skipped, SkippedMessage{Position: i + 1, Type: head.Type, Why: why})
			continue
		}
		out.Messages = append(out.Messages, m)
	}
	*in = out
	return nil
}

// decodeMessage reads one message, or says why it cannot be carried.
func decodeMessage(raw json.RawMessage) (Message, string) {
	var head struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &head); err != nil || head.Type == "" {
		return Message{}, "a message with no type"
	}
	if knownMessage[head.Type] {
		var m Message
		if err := json.Unmarshal(raw, &m); err != nil {
			return Message{}, fmt.Sprintf("a %q message that cannot be read: %v", head.Type, err)
		}
		return m, ""
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var fields map[string]any
	if err := dec.Decode(&fields); err != nil {
		return Message{}, fmt.Sprintf("a %q message that cannot be read: %v", head.Type, err)
	}
	m := Message{Type: head.Type, Extra: map[string]any{}}
	for k, v := range fields {
		switch {
		case k == "type":
		case k == "character":
			s, ok := v.(string)
			if !ok {
				return Message{}, fmt.Sprintf("a %q message whose character is not text", head.Type)
			}
			m.Character = s
		default:
			switch x := v.(type) {
			case string:
				m.Extra[k] = x
			case json.Number:
				f, err := x.Float64()
				if err != nil {
					return Message{}, fmt.Sprintf("a %q message with a number (%s) out of range", head.Type, k)
				}
				m.Extra[k] = f
			default:
				return Message{}, fmt.Sprintf("a %q message, a type this build does not know, "+
					"whose %q field is not text or a number", head.Type, k)
			}
		}
	}
	return m, ""
}

// ScanExports reads one SavedVariables file and returns every
// character export in it. Any table holding a string "export" field
// counts, wherever it sits, so a change to the addon's table shape
// does not silently stop the sync.
func ScanExports(path string) ([]Export, error) {
	saved, err := scanSaved(path)
	return saved.Exports, err
}

// SavedVariables is what the companion reads out of one
// SavedVariables file: the character exports and the outbox.
type SavedVariables struct {
	Exports []Export
	// Outbox is nil when there is nothing to send (see ScanOutbox).
	Outbox *Outbox
}

// scanSaved reads one SavedVariables file once for both of its parts.
// It only ever reads the file: the game owns it.
func scanSaved(path string) (SavedVariables, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return SavedVariables{}, err
	}
	top, err := ParseLua(string(b))
	if err != nil {
		return SavedVariables{}, fmt.Errorf("%s: %w", path, err)
	}
	var out []Export
	for _, v := range top {
		collect(v, &out)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Name != out[j].Name {
			return out[i].Name < out[j].Name
		}
		return out[i].Ruleset < out[j].Ruleset
	})
	ob, err := outboxOf(top)
	if err != nil {
		return SavedVariables{}, fmt.Errorf("%s: %w", path, err)
	}
	return SavedVariables{Exports: out, Outbox: ob}, nil
}

// collect walks a decoded value looking for export tables.
func collect(v any, out *[]Export) {
	t, ok := v.(*Table)
	if !ok {
		return
	}
	if s := t.Get("export"); s != "" {
		// The Phase 2 addon may still write the segment as "realm";
		// either key feeds the contract's ruleset field.
		ruleset := t.Get("ruleset")
		if ruleset == "" {
			ruleset = t.Get("realm")
		}
		*out = append(*out, Export{
			Name:    t.Get("name"),
			Ruleset: ruleset,
			Region:  t.Get("region"),
			Export:  s,
		})
	}
	for _, child := range t.Fields {
		collect(child, out)
	}
	for _, child := range t.Items {
		collect(child, out)
	}
}

// RenderInbox writes the Lua the addon loads. It is deterministic:
// the same inbox at the same time renders byte for byte the same. at
// is when the content was generated, not when the file is written —
// the sync re-renders only when the builds or messages change — which
// is what lets WriteInbox skip a pass that would change nothing. The
// API's own generated_at is not written for the same reason: it moves
// on every fetch.
//
// Builds are rendered exactly as the API sent them, and so are the
// messages: the API sends a "build" message beside each builds entry,
// and the companion adds none of its own.
func RenderInbox(in Inbox, at time.Time) []byte {
	var b strings.Builder
	b.WriteString("-- Written by the Forever Sixty companion. Do not edit:\n")
	b.WriteString("-- this file is replaced every ten minutes.\n")
	fmt.Fprintf(&b, "%s = {\n", InboxGlobal)
	fmt.Fprintf(&b, "\t[\"format\"] = %d,\n", InboxFormat)
	fmt.Fprintf(&b, "\t[\"generated_at\"] = %s,\n", quote(at.UTC().Format(time.RFC3339)))
	b.WriteString("\t[\"builds\"] = {\n")
	for _, bd := range in.Builds {
		b.WriteString("\t\t{\n")
		fmt.Fprintf(&b, "\t\t\t[\"id\"] = %s,\n", quote(bd.ID))
		fmt.Fprintf(&b, "\t\t\t[\"name\"] = %s,\n", quote(bd.Name))
		fmt.Fprintf(&b, "\t\t\t[\"character\"] = %s,\n", quote(bd.Character))
		fmt.Fprintf(&b, "\t\t\t[\"code\"] = %s,\n", quote(bd.Code))
		b.WriteString("\t\t},\n")
	}
	b.WriteString("\t},\n")
	b.WriteString("\t[\"messages\"] = {\n")
	for _, m := range in.Messages {
		renderMessage(&b, m)
	}
	b.WriteString("\t},\n}\n")
	return []byte(b.String())
}

// formatFloat renders a message's numeric fields the same way,
// however small or large: rounded to at most three decimals, trailing
// zeros dropped, never scientific notation (contract 3: the addon
// prints what it gets), so Codec.lua's number reader and a human
// reading the file see the same digits every render of the same value.
func formatFloat(f float64) string {
	if math.Abs(f) < 1e15 { // past this a float64 has no fraction to round
		f = math.Round(f*1000) / 1000
	}
	if f == 0 {
		f = 0 // no "-0"
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// field writes one `["key"] = value,` line at the message's depth.
func field(b *strings.Builder, indent, key, value string) {
	fmt.Fprintf(b, "%s[%s] = %s,\n", indent, quote(key), value)
}

// renderMessage writes one typed message, switching on Type to write
// only the fields that type actually uses: Message carries every
// type's fields at once, so a "guild" message read off the API has
// Slot/ItemID/etc at their zero value, and writing them anyway would
// put meaningless keys ("item_id" = 0) into a guild row. An optional
// field the API left out is left out here too. A type this build does
// not know writes its carried Extra fields, keys sorted.
func renderMessage(b *strings.Builder, m Message) {
	const in = "\t\t\t"
	b.WriteString("\t\t{\n")
	field(b, in, "type", quote(m.Type))
	if m.Character != "" {
		field(b, in, "character", quote(m.Character))
	}
	switch m.Type {
	case MessageBuild:
		field(b, in, "name", quote(m.Name))
		field(b, in, "code", quote(m.Code))
	case MessageUpgrade:
		field(b, in, "slot", quote(m.Slot))
		field(b, in, "item_id", strconv.Itoa(m.ItemID))
		field(b, in, "item_name", quote(m.ItemName))
		field(b, in, "source", quote(m.Source))
		field(b, in, "delta", formatFloat(m.Delta))
		if m.Unit != "" {
			field(b, in, "unit", quote(m.Unit))
		}
	case MessageWeights:
		field(b, in, "spec", quote(m.Spec))
		b.WriteString(in + "[\"weights\"] = {\n")
		for _, w := range m.Weights {
			fmt.Fprintf(b, "%s\t{ [\"stat\"] = %s, [\"weight\"] = %s },\n", in, quote(w.Stat), formatFloat(w.Weight))
		}
		b.WriteString(in + "},\n")
		b.WriteString(in + "[\"caps\"] = {\n")
		for _, c := range m.Caps {
			fmt.Fprintf(b, "%s\t%s,\n", in, quote(c))
		}
		b.WriteString(in + "},\n")
		if m.DPSPerPoint != nil {
			field(b, in, "dps_per_point", formatFloat(*m.DPSPerPoint))
		}
	case MessageSimResult:
		field(b, in, "spec", quote(m.Spec))
		if m.DPS != nil {
			field(b, in, "dps", formatFloat(*m.DPS))
		}
		if m.Error != nil {
			field(b, in, "error", formatFloat(*m.Error))
		}
		if m.GeneratedAt != "" {
			field(b, in, "generated_at", quote(m.GeneratedAt))
		}
		b.WriteString(in + "[\"slots\"] = {\n")
		if m.Slots != nil {
			for _, s := range *m.Slots {
				fmt.Fprintf(b, "%s\t{ [\"slot\"] = %s, [\"item_id\"] = %d, [\"delta\"] = %s },\n",
					in, quote(s.Slot), s.ItemID, formatFloat(s.Delta))
			}
		}
		b.WriteString(in + "},\n")
	case MessageLastRaid:
		field(b, in, "report_id", quote(m.ReportID))
		field(b, in, "date", quote(m.Date))
		b.WriteString(in + "[\"bosses\"] = {\n")
		for _, boss := range m.Bosses {
			renderBoss(b, in+"\t", boss)
		}
		b.WriteString(in + "},\n")
	case MessageLootValues:
		field(b, in, "raid", quote(m.Raid))
		ids := make([]int, 0, len(m.Items))
		for id := range m.Items {
			ids = append(ids, id)
		}
		sort.Ints(ids)
		b.WriteString(in + "[\"items\"] = {\n")
		for _, id := range ids {
			v := m.Items[id]
			fmt.Fprintf(b, "%s\t[%d] = { [\"value\"] = %s, [\"unit\"] = %s, [\"planned\"] = %t },\n",
				in, id, formatFloat(v.Value), quote(v.Unit), v.Planned)
		}
		b.WriteString(in + "},\n")
	case MessageGuild:
		field(b, in, "guild_name", quote(m.GuildName))
		field(b, in, "claim_state", quote(m.ClaimState))
		field(b, in, "pending_approvals", strconv.Itoa(m.PendingApprovals))
		field(b, in, "rank", quote(m.Rank))
	case MessageAck:
		field(b, in, "export_seen_at", quote(m.ExportSeenAt))
		if m.SimRequestState != "" {
			field(b, in, "sim_request_state", quote(m.SimRequestState))
		}
	default:
		keys := make([]string, 0, len(m.Extra))
		for k := range m.Extra {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			switch v := m.Extra[k].(type) {
			case string:
				field(b, in, k, quote(v))
			case float64:
				field(b, in, k, formatFloat(v))
			}
		}
	}
	b.WriteString("\t\t},\n")
}

// renderBoss writes one last_raid boss as a nested table.
func renderBoss(b *strings.Builder, indent string, boss LastRaidBoss) {
	in := indent + "\t"
	b.WriteString(indent + "{\n")
	field(b, in, "name", quote(boss.Name))
	field(b, in, "percentile", strconv.Itoa(boss.Percentile))
	if boss.BiggestLoss != "" {
		field(b, in, "biggest_loss", quote(boss.BiggestLoss))
	}
	field(b, in, "loss_seconds", strconv.Itoa(boss.LossSeconds))
	field(b, in, "deaths", strconv.Itoa(boss.Deaths))
	missed := make([]string, len(boss.ConsumablesMissed))
	for i, c := range boss.ConsumablesMissed {
		missed[i] = quote(c)
	}
	if len(missed) == 0 {
		field(b, in, "consumables_missed", "{}")
	} else {
		field(b, in, "consumables_missed", "{ "+strings.Join(missed, ", ")+" }")
	}
	if boss.Adherence != nil {
		field(b, in, "adherence", formatFloat(*boss.Adherence))
	}
	b.WriteString(indent + "},\n")
}

// InboxPath is where the inbox goes for one SavedVariables path.
func InboxPath(savedVariables string) string {
	return filepath.Join(filepath.Dir(savedVariables), InboxName+".lua")
}

// WriteInbox writes the inbox file unless the bytes already there are
// identical. It reports whether it wrote.
//
// It writes ForeverSixtyInbox.lua and nothing else: any other path is
// refused, so the game's own ForeverSixty.lua (ForeverSixtyDB, which
// the game owns and rewrites at logout) can never be overwritten by a
// path mistake. The write goes to a temporary file in the same folder
// and is renamed over the inbox, so the game never reads half a file.
//
// The game rewrites its own SavedVariables at logout, so an inbox
// written while the player is online may be replaced; the next
// ten-minute pass puts it back, and the addon reads it at the
// following load.
func WriteInbox(path string, body []byte) (bool, error) {
	if filepath.Base(path) != InboxName+".lua" {
		return false, fmt.Errorf("addon: refusing to write %s: the companion writes only %s.lua",
			path, InboxName)
	}
	if old, err := os.ReadFile(path); err == nil && string(old) == string(body) {
		return false, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".inbox-*.lua")
	if err != nil {
		return false, err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err := tmp.Write(body); err != nil {
		tmp.Close()
		return false, err
	}
	if err := tmp.Close(); err != nil {
		return false, err
	}
	if err := os.Chmod(name, 0o644); err != nil {
		return false, err
	}
	return true, os.Rename(name, path)
}
