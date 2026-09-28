// companion/internal/addon/addon.go
// Package addon carries strings between the game and the site: the
// addon's character exports out of SavedVariables and up to the API,
// and the builds the player chose on the site down into an inbox file
// the addon reads at load.
//
// The export strings are opaque here. FS1 is the addon's format and
// the site's; the companion moves it and never inspects it.
package addon

import (
	"fmt"
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

// InboxVersion is written into every rendered inbox as the "version"
// field: the schema this RenderInbox writes (2, for Wave C's typed
// messages — docs/superpowers/specs/2026-09-28-addon-character-aware-
// design.md §3, §4, §6). Nothing here READS it back; forward and
// backward compatibility both come from Message.Type instead (an
// addon build that does not know a Type ignores that one message and
// keeps reading the rest, the same way an addon that predates
// "messages" entirely never looks at the key and is unaffected by it
// appearing). It exists for support requests: naming the version a
// player's companion last wrote is one glance instead of a guess.
const InboxVersion = 2

// The typed inbox message kinds Wave C adds. An addon build that
// predates a kind ignores any message carrying it — see InboxVersion.
const (
	MessageUpgrade = "upgrade"
	MessageWeights = "weights"
	MessageGuild   = "guild"
)

// WeightEntry is one stat's weight inside a "weights" message.
type WeightEntry struct {
	Stat   string  `json:"stat"`
	Weight float64 `json:"weight"`
}

// Message is one typed inbox message beyond a build: the character's
// best upgrade from their last Top Gear run, their saved stat weights
// (with which stats are already capped), or their guild's state.
//
// Type selects which of the fields below the addon should read; the
// rest are the zero value and simply not rendered (renderMessage).
// Character is optional exactly like Build.Character: empty reaches
// every character, addressed reaches one — the addon's
// Follow.sameCharacter matching rule, shared with builds.
type Message struct {
	Type      string `json:"type"`
	Character string `json:"character,omitempty"`

	// MessageUpgrade: the best upgrade Top Gear found for one slot.
	Slot     string  `json:"slot,omitempty"`
	ItemID   int     `json:"item_id,omitempty"`
	ItemName string  `json:"item_name,omitempty"`
	Source   string  `json:"source,omitempty"`
	Delta    float64 `json:"delta,omitempty"`

	// MessageWeights: the character's saved stat weights, and which of
	// those stats are already capped — further points wasted.
	Spec    string        `json:"spec,omitempty"`
	Weights []WeightEntry `json:"weights,omitempty"`
	Caps    []string      `json:"caps,omitempty"`

	// MessageGuild: the character's guild state block (design §3, §4):
	// claim state, the officer's pending-approval count (zero unless
	// this character is an officer or leader), and this character's
	// own rank.
	GuildName        string `json:"guild_name,omitempty"`
	ClaimState       string `json:"claim_state,omitempty"`
	PendingApprovals int    `json:"pending_approvals,omitempty"`
	Rank             string `json:"rank,omitempty"`
}

// Inbox is the body of GET /v1/addon/inbox.
type Inbox struct {
	Builds   []Build   `json:"builds"`
	Messages []Message `json:"messages,omitempty"`
}

// ScanExports reads one SavedVariables file and returns every
// character export in it. Any table holding a string "export" field
// counts, wherever it sits, so a change to the addon's table shape
// does not silently stop the sync.
func ScanExports(path string) ([]Export, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	top, err := ParseLua(string(b))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
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
	return out, nil
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
// is when the builds were generated, not when the file is written —
// the sync re-renders only when the build set changes — which is what
// lets WriteInbox skip a pass that would change nothing.
func RenderInbox(in Inbox, at time.Time) []byte {
	var b strings.Builder
	b.WriteString("-- Written by the Forever Sixty companion. Do not edit:\n")
	b.WriteString("-- this file is replaced every ten minutes.\n")
	fmt.Fprintf(&b, "%s = {\n", InboxGlobal)
	fmt.Fprintf(&b, "\t[\"version\"] = %d,\n", InboxVersion)
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
// however small or large: six decimal places, never scientific
// notation, so Codec.lua's Lua number reader (and a human reading the
// file) sees the same digits every render of the same value.
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

// renderMessage writes one typed message, switching on Type to write
// only the fields that kind actually uses: Message carries every
// kind's fields at once, so a "guild" message read off the API has
// Slot/ItemID/etc at their zero value, and writing them anyway would
// put meaningless keys ("item_id" = 0) into a guild row. A Type this
// companion build does not recognise (a future kind the site started
// sending before this build was upgraded) falls through the switch
// with no case, so only "type" and "character" are written for it —
// forward-compatible by omission rather than by guessing its shape.
func renderMessage(b *strings.Builder, m Message) {
	b.WriteString("\t\t{\n")
	fmt.Fprintf(b, "\t\t\t[\"type\"] = %s,\n", quote(m.Type))
	if m.Character != "" {
		fmt.Fprintf(b, "\t\t\t[\"character\"] = %s,\n", quote(m.Character))
	}
	switch m.Type {
	case MessageUpgrade:
		fmt.Fprintf(b, "\t\t\t[\"slot\"] = %s,\n", quote(m.Slot))
		fmt.Fprintf(b, "\t\t\t[\"item_id\"] = %d,\n", m.ItemID)
		fmt.Fprintf(b, "\t\t\t[\"item_name\"] = %s,\n", quote(m.ItemName))
		fmt.Fprintf(b, "\t\t\t[\"source\"] = %s,\n", quote(m.Source))
		fmt.Fprintf(b, "\t\t\t[\"delta\"] = %s,\n", formatFloat(m.Delta))
	case MessageWeights:
		fmt.Fprintf(b, "\t\t\t[\"spec\"] = %s,\n", quote(m.Spec))
		b.WriteString("\t\t\t[\"weights\"] = {\n")
		for _, w := range m.Weights {
			fmt.Fprintf(b, "\t\t\t\t{ [\"stat\"] = %s, [\"weight\"] = %s },\n", quote(w.Stat), formatFloat(w.Weight))
		}
		b.WriteString("\t\t\t},\n")
		b.WriteString("\t\t\t[\"caps\"] = {\n")
		for _, c := range m.Caps {
			fmt.Fprintf(b, "\t\t\t\t%s,\n", quote(c))
		}
		b.WriteString("\t\t\t},\n")
	case MessageGuild:
		fmt.Fprintf(b, "\t\t\t[\"guild_name\"] = %s,\n", quote(m.GuildName))
		fmt.Fprintf(b, "\t\t\t[\"claim_state\"] = %s,\n", quote(m.ClaimState))
		fmt.Fprintf(b, "\t\t\t[\"pending_approvals\"] = %d,\n", m.PendingApprovals)
		fmt.Fprintf(b, "\t\t\t[\"rank\"] = %s,\n", quote(m.Rank))
	}
	b.WriteString("\t\t},\n")
}

// InboxPath is where the inbox goes for one SavedVariables path.
func InboxPath(savedVariables string) string {
	return filepath.Join(filepath.Dir(savedVariables), InboxName+".lua")
}

// WriteInbox writes the inbox file unless the bytes already there are
// identical. It reports whether it wrote.
//
// The game rewrites its own SavedVariables at logout, so an inbox
// written while the player is online may be replaced; the next
// ten-minute pass puts it back, and the addon reads it at the
// following load.
func WriteInbox(path string, body []byte) (bool, error) {
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
