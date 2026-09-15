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

// Inbox is the body of GET /v1/addon/inbox.
type Inbox struct {
	Builds []Build `json:"builds"`
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
	b.WriteString("\t},\n}\n")
	return []byte(b.String())
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
