package golinks

import (
	"strings"
	"testing"
)

func TestValidSlug(t *testing.T) {
	for _, s := range []string{"reddit", "reddit-classicwow", "a", "0x", strings.Repeat("a", 40)} {
		if !ValidSlug(s) {
			t.Errorf("ValidSlug(%q) = false", s)
		}
	}
	for _, s := range []string{"", "-a", "Reddit", "a_b", "a/b", "a b", strings.Repeat("a", 41)} {
		if ValidSlug(s) {
			t.Errorf("ValidSlug(%q) = true", s)
		}
	}
}

func TestValidTargetKeepsEveryLinkOnTheSite(t *testing.T) {
	for _, s := range []string{"/", "/guides", "/sim", "/bis/hunter/marksmanship", "/premium?plan=guild", "/guides#rogue"} {
		if !ValidTarget(s) {
			t.Errorf("ValidTarget(%q) = false", s)
		}
	}
	for _, s := range []string{
		"", "guides", "//evil.example", "https://evil.example", "/\\evil.example", "\\\\evil",
		"/a b", "/a\tb", "javascript:alert(1)", "/" + strings.Repeat("a", MaxTargetLen),
	} {
		if ValidTarget(s) {
			t.Errorf("ValidTarget(%q) = true", s)
		}
	}
}

func TestValidLabel(t *testing.T) {
	if ValidLabel("") || ValidLabel("   ") || ValidLabel(strings.Repeat("a", MaxLabelLen+1)) {
		t.Fatal("an empty or long label was accepted")
	}
	if !ValidLabel("Reddit: r/classicwow") {
		t.Fatal("a plain label was refused")
	}
}
