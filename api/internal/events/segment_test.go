package events

import (
	"fmt"
	"testing"
)

func TestSegmentfKeepsTheCopyAsItsText(t *testing.T) {
	name, item := Name("Gorvash", "warrior"), Item("Deathbringer", 17068, 4)
	cases := []struct {
		format string
		args   []Segment
		want   []Segment
	}{
		{"%s was awarded %s by %s.", []Segment{name, item, Text("Olympus")},
			[]Segment{name, Text(" was awarded "), item, Text(" by Olympus.")}},
		{"%s", []Segment{Strong("1,184")}, []Segment{Strong("1,184")}},
		{"No verbs.", nil, []Segment{Text("No verbs.")}},
	}
	for _, c := range cases {
		t.Run(c.format, func(t *testing.T) {
			got := Segmentf(c.format, c.args...)
			if fmt.Sprint(got) != fmt.Sprint(c.want) {
				t.Fatalf("segments %v, want %v", got, c.want)
			}
			texts := make([]any, len(c.args))
			for i, a := range c.args {
				texts[i] = a.Text
			}
			if c.args != nil && SegmentsText(got) != fmt.Sprintf(c.format, texts...) {
				t.Fatalf("text %q, want the format's %q", SegmentsText(got), fmt.Sprintf(c.format, texts...))
			}
		})
	}
}
