package coach

import (
	"regexp"
	"slices"
	"strings"
	"testing"
)

// capitalisedRun is two or more capitalised words in a row: what a person's name looks like.
var capitalisedRun = regexp.MustCompile(`\b[A-Z][a-z]+(?:[ \t]+[A-Z][a-z]+)+\b`)

// functionWords are the closed-class English words a sentence may open with ("In Forever",
// "The Succubus", "Each Succubus"): capitalised by grammar, never part of a person's name.
var functionWords = map[string]bool{
	"a": true, "an": true, "the": true, "this": true, "that": true, "these": true, "those": true,
	"each": true, "every": true, "any": true, "all": true, "both": true, "no": true, "your": true,
	"its": true, "their": true, "in": true, "on": true, "at": true, "of": true, "for": true,
	"from": true, "with": true, "without": true, "by": true, "to": true, "into": true, "after": true,
	"before": true, "until": true, "if": true, "when": true, "while": true, "once": true, "then": true,
	"and": true, "but": true, "or": true, "so": true, "it": true, "you": true, "she": true, "he": true,
	"they": true,
}

// gameWords is every lowercased word of every dictionary name (the client's spell, talent, item
// and boss names and the consumables catalogue): "Forever" from "Lordaeron Forever Tabard",
// "Succubus" from "Summon Succubus". A capitalised run made only of these is game prose.
func gameWords(d *NameDict) map[string]bool {
	words := map[string]bool{}
	for key := range d.names {
		for _, w := range strings.Fields(key) {
			words[w] = true
		}
	}
	return words
}

// personLikeRuns returns each capitalised run in s that could be a person's name: it carries
// a word that is not a class or spec label, not a word of the game's own names and not a
// function word.
func personLikeRuns(s string, labels, game map[string]bool) []string {
	var out []string
	for _, run := range capitalisedRun.FindAllString(s, -1) {
		if slices.ContainsFunc(strings.Fields(run), func(w string) bool {
			key := strings.ToLower(w)
			return !labels[key] && !game[key] && !functionWords[key]
		}) {
			out = append(out, run)
		}
	}
	return out
}

func TestPersonLikeRuns_GameWordsAreNotNames(t *testing.T) {
	labels := map[string]bool{"warlock": true, "affliction": true}
	game := map[string]bool{"forever": true, "succubus": true, "summon": true, "drain": true, "soul": true}
	for _, s := range []string{
		"In Forever, Drain Soul has a chance to give a shard.",
		"The Succubus costs a shard where the Imp is free.",
		"Each Succubus summon costs one.",
		"- In Forever the totem reaches the raid.",
		"This is the Affliction Warlock as the guide describes it.",
	} {
		if got := personLikeRuns(s, labels, game); len(got) != 0 {
			t.Errorf("%q: flagged %v", s, got)
		}
	}
}

func TestPersonLikeRuns_PeopleAreFlagged(t *testing.T) {
	game := map[string]bool{"forever": true, "succubus": true}
	for s, want := range map[string]string{
		"Ask Justin Hunthrop for the pull timer.":    "Justin Hunthrop",
		"Play it the way Bob Smith does.":            "Bob Smith",
		"Justin Hunthrop runs this pull.":            "Justin Hunthrop",
		"The healers follow Forever Bobson's calls.": "Forever Bobson",
		"In Bob Smith's guild the rule differs.":     "Bob Smith",
	} {
		got := personLikeRuns(s, map[string]bool{}, game)
		if !slices.ContainsFunc(got, func(run string) bool { return strings.Contains(run, want) }) {
			t.Errorf("%q: got %v, want %q flagged", s, got, want)
		}
	}
}

func TestPersonLikeRuns_ReadsTheRealDictionary(t *testing.T) {
	game := gameWords(buildNames(t))
	for _, w := range []string{"forever", "succubus"} {
		if !game[w] {
			t.Errorf("%q is not a word of any game name", w)
		}
	}
	if got := personLikeRuns("In Forever, the Succubus tanks. The Succubus hits.", nil, game); len(got) != 0 {
		t.Errorf("flagged %v", got)
	}
	if got := personLikeRuns("Ask Justin Hunthrop for the pull timer.", nil, game); len(got) != 1 {
		t.Errorf("a person's name against the real dictionary: got %v", got)
	}
}
