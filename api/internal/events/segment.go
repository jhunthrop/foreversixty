package events

import "strings"

// The segment styles (a feed row's and an Activity row's segments[].style). Plain text has
// none. The feed (package activity) and the Activity inbox share this one builder.
const (
	StyleStrong = "strong"
	StyleName   = "name"
	StyleItem   = "item"
)

// Segment is one run of a row's sentence.
type Segment struct {
	Text    string `json:"text"`
	Style   string `json:"style,omitempty"`
	Class   string `json:"class,omitempty"`
	Quality *int   `json:"quality,omitempty"`
	ItemID  *int   `json:"item_id,omitempty"`
}

// Text is a plain segment.
func Text(s string) Segment { return Segment{Text: s} }

// Strong is a bold segment.
func Strong(s string) Segment { return Segment{Text: s, Style: StyleStrong} }

// Name is a character's name in its class colour.
func Name(s, class string) Segment { return Segment{Text: s, Style: StyleName, Class: class} }

// Item is an item's name in its quality colour.
func Item(name string, id, quality int) Segment {
	return Segment{Text: name, Style: StyleItem, ItemID: &id, Quality: &quality}
}

// Segmentf is format's sentence with each %s verb replaced by the next arg, the text between
// them plain: Segmentf("%s was awarded %s.", Name(..), Item(..)). Its SegmentsText is the
// format's fmt.Sprintf of the args' text, so a copy string stays the one source of both.
func Segmentf(format string, args ...Segment) []Segment {
	parts := strings.Split(format, "%s")
	out := make([]Segment, 0, len(parts)+len(args))
	for i, p := range parts {
		if p != "" {
			out = appendSegment(out, Text(p))
		}
		if i < len(args) && i < len(parts)-1 {
			out = appendSegment(out, args[i])
		}
	}
	return out
}

// appendSegment adds s, joining it to a plain segment before it when s is plain too.
func appendSegment(out []Segment, s Segment) []Segment {
	if n := len(out); n > 0 && isPlain(s) && isPlain(out[n-1]) {
		out[n-1].Text += s.Text
		return out
	}
	return append(out, s)
}

func isPlain(s Segment) bool { return s.Style == "" }

// SegmentsText is the segments' sentence as plain text: the body a push shows, so the body and
// the styled row never say different things.
func SegmentsText(segs []Segment) string {
	var b strings.Builder
	for _, s := range segs {
		b.WriteString(s.Text)
	}
	return b.String()
}
