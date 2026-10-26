package guilds

import (
	"strings"
	"testing"
	"time"
)

const maxRenderDuration = 10 * time.Millisecond

func TestRenderNotes(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		exact       *string
		contains    []string
		notContains []string
	}{
		{name: "empty", input: "", exact: strPtr("")},
		{name: "bold", input: "**bold**", exact: strPtr("<p><strong>bold</strong></p>\n")},
		{name: "italic", input: "*soft*", contains: []string{"<em>soft</em>"}},
		{name: "unordered list", input: "- a\n- b", contains: []string{"<ul>", "<li>a</li>", "<li>b</li>"}},
		{name: "ordered list", input: "1. a\n2. b", contains: []string{"<ol>", "<li>a</li>"}},
		{
			name:        "raw html is escaped",
			input:       "<script>x</script>",
			contains:    []string{"x"},
			notContains: []string{"<script", "</script"},
		},
		{
			name:        "image dropped",
			input:       "![x](http://example.com/u.png)",
			notContains: []string{"<img", "src="},
		},
		{
			name:        "javascript link loses href",
			input:       "[x](javascript:alert(1))",
			contains:    []string{"x"},
			notContains: []string{"href", "javascript"},
		},
		{
			name:     "bare url is linked safely",
			input:    "https://foreversixty.gg",
			contains: []string{`<a href="https://foreversixty.gg"`, `rel="nofollow noopener"`, `target="_blank"`},
		},
		{
			name:        "heading flattened",
			input:       "# Title",
			contains:    []string{"Title"},
			notContains: []string{"<h1"},
		},
		{
			name:        "code block flattened",
			input:       "```\nfoo\n```",
			contains:    []string{"foo"},
			notContains: []string{"<pre", "<code"},
		},
		{
			name:        "table flattened",
			input:       "| a | b |\n|---|---|\n| 1 | 2 |",
			notContains: []string{"<table", "<tr", "<td"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RenderNotes(tc.input)
			if tc.exact != nil && got != *tc.exact {
				t.Fatalf("got %q, want %q", got, *tc.exact)
			}
			for _, want := range tc.contains {
				if !strings.Contains(got, want) {
					t.Errorf("output %q missing %q", got, want)
				}
			}
			for _, bad := range tc.notContains {
				if strings.Contains(got, bad) {
					t.Errorf("output %q must not contain %q", got, bad)
				}
			}
		})
	}
}

func TestRenderNotesTimingBudget(t *testing.T) {
	if testing.Short() {
		t.Skip("timing assertion skipped in short mode")
	}
	input := strings.Repeat("**bold** and a [link](https://example.com) - ", 50)[:2000]
	RenderNotes(input) // warm any lazy state before timing
	start := time.Now()
	RenderNotes(input)
	if elapsed := time.Since(start); elapsed > maxRenderDuration {
		t.Fatalf("rendering 2000 chars took %v, budget %v", elapsed, maxRenderDuration)
	}
}

func strPtr(s string) *string { return &s }
