package guilds

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// notesMarkdown renders CommonMark with bare-URL linking. Raw HTML is not
// passed through (no html.WithUnsafe), so it is escaped as text.
var notesMarkdown = goldmark.New(goldmark.WithExtensions(extension.Linkify))

// notesHTMLEscaper makes raw HTML in the source render as visible text instead
// of being silently omitted by goldmark.
var notesHTMLEscaper = strings.NewReplacer("<", "&lt;")

// notesPolicy is the strict allow-list applied after markdown rendering.
var notesPolicy = newNotesPolicy()

func newNotesPolicy() *bluemonday.Policy {
	policy := bluemonday.NewPolicy()
	policy.AllowElements("p", "strong", "em", "ul", "ol", "li")
	policy.AllowAttrs("href").OnElements("a")
	policy.AllowURLSchemes("http", "https")
	policy.RequireNoFollowOnLinks(true)
	policy.RequireParseableURLs(true)
	policy.AddTargetBlankToFullyQualifiedLinks(true)
	return policy
}

// RenderNotes turns officer-written raid notes (a markdown subset) into HTML
// that is safe to show to members. Unsupported constructs are flattened to text
// or dropped; an empty input renders to an empty string. The caller logs the error.
func RenderNotes(markdown string) (string, error) {
	if markdown == "" {
		return "", nil
	}
	var rendered bytes.Buffer
	if err := notesMarkdown.Convert([]byte(notesHTMLEscaper.Replace(markdown)), &rendered); err != nil {
		return "", fmt.Errorf("guilds: render raid notes: %w", err)
	}
	return notesPolicy.Sanitize(rendered.String()), nil
}
