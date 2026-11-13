package leveling

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// The addon's Builds pane shows one line under each build saying why the
// guide builds it that way (design/specs/2026-11-13-addon-builds-pane.md,
// section 3). The line is the guide's own sentence, read from the prose of
// its "## Talents and builds" section, never composed here.

// talentsSectionRe cuts the Talents section's body out of a guide: from its
// heading to the next level-two heading or the end of the file.
var talentsSectionRe = regexp.MustCompile(`(?ms)^## Talents and builds\s*\n(.*?)(?:^## |\z)`)

// whyLeads are the sentence openings that name one build's own why, in the
// order they are tried for each kind: the first lead any sentence of the
// section's prose opens with wins, and its first such sentence is the why.
var whyLeads = map[string][]string{
	"leveling": {
		"The leveling build",
		"The level 60 build",
		"You level",
		"While leveling",
		"Out in the world, keep the level 60 build",
		"Out of a raid",
	},
	"raid": {
		"The raid build",
		"In the raid build",
		"At 60, for raids",
		"In a raid",
	},
}

// shortWhy is the fewest characters a why may have: "There are two
// builds." or "The level 60 build is the one you level with." says little on
// its own, so the next sentence of its paragraph joins it.
const shortWhy = 60

// GuideBuildWhy reads a spec's guide and returns the one-sentence why for
// its build of `kind` ("leveling" or "raid"); "" when the guide has no
// Talents prose. An unknown kind is an error.
func GuideBuildWhy(repoRoot, class, specSlug, kind string) (string, error) {
	if _, ok := whyLeads[kind]; !ok {
		return "", fmt.Errorf("leveling: no why lead for build kind %q", kind)
	}
	path := filepath.Join(repoRoot, "web", "src", "content", "guides", class, specSlug+".md")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("leveling: reading %s: %w", path, err)
	}
	return BuildWhy(string(b), kind), nil
}

// BuildWhy is GuideBuildWhy over a guide's text: the first prose sentence of
// the Talents section that opens with one of `kind`'s leads, else the first
// sentence of the section's first prose paragraph; a sentence under shortWhy
// characters takes the next one of its paragraph with it. Markdown bold is
// stripped. "" when the section or its prose is missing.
func BuildWhy(guide, kind string) string {
	m := talentsSectionRe.FindStringSubmatch(guide)
	if m == nil {
		return ""
	}
	paragraphs := proseParagraphs(m[1])
	if len(paragraphs) == 0 {
		return ""
	}
	for _, lead := range whyLeads[kind] {
		for _, paragraph := range paragraphs {
			all := sentences(paragraph)
			for i, sentence := range all {
				if strings.HasPrefix(sentence, lead) {
					return withNext(all, i)
				}
			}
		}
	}
	return withNext(sentences(paragraphs[0]), 0)
}

// withNext is the i-th sentence, joined by the next one of its paragraph
// while it is under shortWhy characters, unless that next sentence opens
// another build's why.
func withNext(all []string, i int) string {
	if i >= len(all) {
		return ""
	}
	why := all[i]
	if len(why) < shortWhy && i+1 < len(all) && !opensAWhy(all[i+1]) {
		why += " " + all[i+1]
	}
	return why
}

// opensAWhy reports whether sentence opens with any kind's lead.
func opensAWhy(sentence string) bool {
	for _, leads := range whyLeads {
		for _, lead := range leads {
			if strings.HasPrefix(sentence, lead) {
				return true
			}
		}
	}
	return false
}

// proseParagraphs are the section's blank-line separated paragraphs that are
// not bullet lists, bold stripped and lines joined, in order.
func proseParagraphs(section string) []string {
	var out []string
	for _, block := range strings.Split(strings.ReplaceAll(section, "\r\n", "\n"), "\n\n") {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		if len(lines) == 0 || lines[0] == "" || strings.HasPrefix(lines[0], "- ") {
			continue
		}
		text := strings.Join(lines, " ")
		text = strings.ReplaceAll(text, "**", "")
		out = append(out, strings.Join(strings.Fields(text), " "))
	}
	return out
}

// sentenceEndRe finds a sentence end: a full stop, question or exclamation
// mark followed by a space and an uppercase letter or an opening quote.
var sentenceEndRe = regexp.MustCompile(`([.!?])\s+(?:[A-Z"“])`)

// sentences splits a paragraph into its sentences, each with its own final
// stop, in order.
func sentences(paragraph string) []string {
	var out []string
	rest := paragraph
	for {
		loc := sentenceEndRe.FindStringSubmatchIndex(rest)
		if loc == nil {
			break
		}
		end := loc[3]
		out = append(out, strings.TrimSpace(rest[:end]))
		rest = strings.TrimSpace(rest[end:])
	}
	if rest != "" {
		out = append(out, rest)
	}
	return out
}
