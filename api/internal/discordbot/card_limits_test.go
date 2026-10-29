package discordbot

import (
	"encoding/json"
	"testing"
	"unicode/utf8"
)

// Discord's limits on a message with components.
const (
	maxCustomIDLen      = 100
	maxActionRows       = 5
	maxPlaceholderRunes = 150
)

// assertCardLimits fails the test when m breaks a Discord limit or a card rule: custom ids,
// row count, select option counts, label and placeholder lengths, button styles, the empty
// mention list and the ephemeral flag.
func assertCardLimits(t *testing.T, m Message) {
	t.Helper()
	if !m.Ephemeral {
		t.Error("card is not ephemeral")
	}
	if len(m.Components) > maxActionRows {
		t.Errorf("%d action rows", len(m.Components))
	}
	successes := 0
	for _, row := range m.Components {
		for _, c := range row {
			successes += checkComponent(t, c)
		}
	}
	if successes > 1 {
		t.Errorf("%d Success buttons", successes)
	}
	checkMentions(t, m)
}

func checkComponent(t *testing.T, c Component) (success int) {
	t.Helper()
	if utf8.RuneCountInString(c.CustomID) > maxCustomIDLen {
		t.Errorf("custom id %q over %d", c.CustomID, maxCustomIDLen)
	}
	if utf8.RuneCountInString(c.Label) > maxOptionRunes {
		t.Errorf("label %q too long", c.Label)
	}
	if utf8.RuneCountInString(c.Placeholder) > maxPlaceholderRunes {
		t.Errorf("placeholder %q too long", c.Placeholder)
	}
	if c.Kind == KindStringSelect {
		checkSelect(t, c)
		return 0
	}
	switch c.Style {
	case StyleSuccess:
		return 1
	case StyleSecondary, StyleLink:
	default:
		t.Errorf("button %q has style %d", c.Label, c.Style)
	}
	return 0
}

func checkSelect(t *testing.T, c Component) {
	t.Helper()
	if n := len(c.Options); n < 1 || n > maxSelectOptions {
		t.Errorf("select %q has %d options", c.CustomID, n)
	}
	for _, o := range c.Options {
		if utf8.RuneCountInString(o.Label) > maxOptionRunes || utf8.RuneCountInString(o.Description) > maxOptionRunes {
			t.Errorf("option %q too long", o.Label)
		}
	}
}

func checkMentions(t *testing.T, m Message) {
	t.Helper()
	raw, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var body struct {
		AllowedMentions struct {
			Parse []string `json:"parse"`
		} `json:"allowed_mentions"`
		Flags int `json:"flags"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatal(err)
	}
	if body.AllowedMentions.Parse == nil || len(body.AllowedMentions.Parse) != 0 {
		t.Errorf("allowed_mentions.parse = %v", body.AllowedMentions.Parse)
	}
	if body.Flags != ephemeralFlag {
		t.Errorf("flags = %d", body.Flags)
	}
}
