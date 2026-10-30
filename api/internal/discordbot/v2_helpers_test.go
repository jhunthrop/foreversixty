package discordbot

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// headOf reads a V2 card's container back as the embed it encodes (spec 2026-10-29 §A.7): the
// accent as Color, "-# " as Author, "### " as Title (its link as URL), the rest as Description,
// the accessory as Thumbnail. Footers and timestamps are gone in V2.
func headOf(t *testing.T, m Message) notify.Embed {
	t.Helper()
	e, ok := headFrom(m)
	if !ok {
		t.Fatalf("no container in %+v", m)
	}
	return e
}

// headFrom is headOf without a test; ok is false when m has no container.
func headFrom(m Message) (notify.Embed, bool) {
	for _, n := range m.Layout {
		if n.Kind != KindContainer || len(n.Children) == 0 {
			continue
		}
		var e notify.Embed
		if n.Accent != nil {
			e.Color = *n.Accent
		}
		first := n.Children[0]
		content := first.Content
		if first.Kind == KindSection {
			content = first.Children[0].Content
			e.Thumbnail = first.Accessory.Media
		}
		var rest []string
		for _, line := range strings.Split(content, "\n") {
			switch {
			case e.Author == "" && e.Title == "" && strings.HasPrefix(line, "-# "):
				e.Author = strings.TrimPrefix(line, "-# ")
			case e.Title == "" && strings.HasPrefix(line, "### "):
				e.Title = strings.TrimPrefix(line, "### ")
				if strings.HasPrefix(e.Title, "[") && strings.HasSuffix(e.Title, ")") {
					if at := strings.LastIndex(e.Title, "]("); at > 0 {
						e.Title, e.URL = e.Title[1:at], e.Title[at+2:len(e.Title)-1]
					}
				}
			default:
				rest = append(rest, line)
			}
		}
		e.Description = strings.Join(rest, "\n")
		return e, true
	}
	return notify.Embed{}, false
}

// textsOf is every Text Display of a V2 message, in order, nested ones included.
func textsOf(m Message) []string {
	var out []string
	var walk func([]Node)
	walk = func(nodes []Node) {
		for _, n := range nodes {
			if n.Kind == KindTextDisplay {
				out = append(out, n.Content)
			}
			walk(n.Children)
		}
	}
	walk(m.Layout)
	return out
}

// sectionsOf are a V2 message's Sections, in order.
func sectionsOf(m Message) []Node {
	var out []Node
	var walk func([]Node)
	walk = func(nodes []Node) {
		for _, n := range nodes {
			if n.Kind == KindSection {
				out = append(out, n)
			}
			walk(n.Children)
		}
	}
	walk(m.Layout)
	return out
}

// wireOf is v as Discord is sent it, decoded generically.
func wireOf(t *testing.T, v any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}

// contentOf is a message's line: its content, or on a V2 message its first top-level Text
// Display (the §A.7 encoding of a content line).
func contentOf(m Message) string {
	if !m.componentsV2() {
		return m.Content
	}
	if first := m.Layout[0]; first.Kind == KindTextDisplay {
		return first.Content
	}
	return ""
}

// hasCard reports a V2 message that carries a card's container (the encoding of an embed).
func hasCard(m Message) bool {
	for _, n := range m.Layout {
		if n.Kind == KindContainer {
			return true
		}
	}
	return false
}
