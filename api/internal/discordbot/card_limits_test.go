package discordbot

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

// Discord's limits on a message with components.
const (
	maxCustomIDLen      = 100
	maxActionRows       = 5
	maxPlaceholderRunes = 150
)

// assertCardLimits fails the test for every violation cardViolations finds in m.
func assertCardLimits(t *testing.T, m Message) {
	t.Helper()
	for _, v := range cardViolations(m) {
		t.Error(v)
	}
}

// cardViolations lists how m breaks a Discord limit or a card rule: custom ids, row count,
// select option counts, label and placeholder lengths, button styles, the empty mention list
// and the ephemeral flag.
func cardViolations(m Message) []string {
	var out []string
	if len(m.Components) > maxActionRows {
		out = append(out, fmt.Sprintf("%d action rows", len(m.Components)))
	}
	successes := 0
	for _, row := range m.Components {
		for _, c := range row {
			v, success := componentViolations(c)
			out = append(out, v...)
			successes += success
		}
	}
	if successes > 1 {
		out = append(out, fmt.Sprintf("%d Success buttons", successes))
	}
	return append(out, wireViolations(m)...)
}

func componentViolations(c Component) (out []string, success int) {
	if n := utf8.RuneCountInString(c.CustomID); n > maxCustomIDLen {
		out = append(out, fmt.Sprintf("custom id %q is %d long", c.CustomID, n))
	}
	if utf8.RuneCountInString(c.Label) > maxOptionRunes {
		out = append(out, fmt.Sprintf("label %q too long", c.Label))
	}
	if utf8.RuneCountInString(c.Placeholder) > maxPlaceholderRunes {
		out = append(out, fmt.Sprintf("placeholder %q too long", c.Placeholder))
	}
	if c.Kind == KindStringSelect {
		return append(out, selectViolations(c)...), 0
	}
	switch c.Style {
	case StyleSuccess:
		success = 1
	case StyleSecondary, StyleLink:
	default:
		out = append(out, fmt.Sprintf("button %q has style %d", c.Label, c.Style))
	}
	return out, success
}

func selectViolations(c Component) (out []string) {
	if n := len(c.Options); n < 1 || n > maxSelectOptions {
		out = append(out, fmt.Sprintf("select %q has %d options", c.CustomID, n))
	}
	for _, o := range c.Options {
		if utf8.RuneCountInString(o.Label) > maxOptionRunes || utf8.RuneCountInString(o.Description) > maxOptionRunes {
			out = append(out, fmt.Sprintf("option %q too long", o.Label))
		}
	}
	return out
}

// wireViolations checks what the message marshals to: the ephemeral flag and the mention list.
func wireViolations(m Message) []string {
	raw, err := json.Marshal(m)
	if err != nil {
		return []string{"marshal: " + err.Error()}
	}
	var body struct {
		AllowedMentions struct {
			Parse []string `json:"parse"`
		} `json:"allowed_mentions"`
		Flags int `json:"flags"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return []string{"unmarshal: " + err.Error()}
	}
	var out []string
	if body.AllowedMentions.Parse == nil || len(body.AllowedMentions.Parse) != 0 {
		out = append(out, fmt.Sprintf("allowed_mentions.parse = %v", body.AllowedMentions.Parse))
	}
	if body.Flags != ephemeralFlag {
		out = append(out, fmt.Sprintf("flags = %d, not ephemeral", body.Flags))
	}
	return out
}

func TestCardViolationsCatchEachBadInput(t *testing.T) {
	button := func(style ButtonStyle, id string) Component {
		return Component{Kind: KindButton, Style: style, Label: "Go", CustomID: id}
	}
	good := func() Message {
		return Message{Ephemeral: true, Components: []ActionRow{{button(StyleSecondary, "gbis:1:go:")}}}
	}
	if v := cardViolations(good()); len(v) != 0 {
		t.Fatalf("good card has violations %v", v)
	}
	rows := good()
	rows.Components = make([]ActionRow, maxActionRows+1)
	for i := range rows.Components {
		rows.Components[i] = ActionRow{button(StyleSecondary, "gbis:1:go:")}
	}
	twoSuccess := good()
	twoSuccess.Components = []ActionRow{{button(StyleSuccess, "a"), button(StyleSuccess, "b")}}
	emptySelect := good()
	emptySelect.Components = []ActionRow{{{Kind: KindStringSelect, CustomID: "s"}}}
	notEphemeral := good()
	notEphemeral.Ephemeral = false
	longID := good()
	longID.Components = []ActionRow{{button(StyleSecondary, strings.Repeat("x", maxCustomIDLen+1))}}
	for name, c := range map[string]struct {
		m    Message
		want string
	}{
		"101-character custom id": {longID, "custom id"},
		"six action rows":         {rows, "6 action rows"},
		"two Success buttons":     {twoSuccess, "2 Success buttons"},
		"select with no options":  {emptySelect, "0 options"},
		"missing ephemeral flag":  {notEphemeral, "not ephemeral"},
	} {
		v := strings.Join(cardViolations(c.m), "; ")
		if !strings.Contains(v, c.want) {
			t.Errorf("%s: violations %q lack %q", name, v, c.want)
		}
	}
}
