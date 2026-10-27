package discordbot

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

const commandHeading = "### 6.2 In Discord"

// descriptionSegments pairs each command's description in the spec with its name: the cell
// reads `/raid` `Your guild's…`; `/signup` `Sign up…`.
func descriptionSegments(t *testing.T) map[string]string {
	t.Helper()
	seg := specCell(t, commandHeading, "Command descriptions")
	out := map[string]string{}
	for i := 0; i+1 < len(seg); i += 2 {
		out[strings.TrimPrefix(seg[i], "/")] = seg[i+1]
	}
	return out
}

func TestCommandDefinitions(t *testing.T) {
	cmds := Commands()
	order := []string{"raid", "signup", "char", "bis", "report", "attendance", "readiness", "link", "help"}
	if len(cmds) != len(order) {
		t.Fatalf("%d commands", len(cmds))
	}
	desc := descriptionSegments(t)
	released := map[string]bool{"raid": true, "signup": true, "char": true, "bis": true, "link": true, "help": true}
	for i, c := range cmds {
		if c.Name != order[i] || c.Description != desc[c.Name] || c.Released != released[c.Name] {
			t.Errorf("command %d = %+v, want %s %q", i, c, order[i], desc[c.Name])
		}
		if len([]rune(c.Description)) > 100 {
			t.Errorf("%s description too long", c.Name)
		}
	}
	checkOptionDescriptions(t, cmds)
	checkChoices(t, cmds)
}

func checkOptionDescriptions(t *testing.T, cmds []Command) {
	t.Helper()
	inSpec := map[string]bool{}
	for _, seg := range specCell(t, commandHeading, "Option descriptions") {
		inSpec[seg] = true
	}
	for _, c := range cmds {
		for _, o := range c.Options {
			if !inSpec[o.Description] || len([]rune(o.Description)) > 100 {
				t.Errorf("/%s %s = %q is not a spec string", c.Name, o.Name, o.Description)
			}
		}
	}
}

func checkChoices(t *testing.T, cmds []Command) {
	t.Helper()
	byName := map[string]CommandOption{}
	for _, c := range cmds {
		for _, o := range c.Options {
			if len(o.Choices) > 0 {
				byName[c.Name+"."+o.Name] = o
			}
		}
	}
	var statuses []string
	for _, ch := range byName["signup.status"].Choices {
		statuses = append(statuses, ch.Name)
	}
	if strings.Join(statuses, ",") != "In,Late,Bench,Out" {
		t.Errorf("status choices = %v", statuses)
	}
	if n := len(byName["bis.slot"].Choices); n != 17 {
		t.Errorf("slot choices = %d", n)
	}
	if n := len(byName["bis.class"].Choices); n != 9 {
		t.Errorf("class choices = %d", n)
	}
	if n := len(byName["bis.phase"].Choices); n != 6 {
		t.Errorf("phase choices = %d", n)
	}
	for _, o := range cmds[1].Options {
		if o.Type != 3 {
			t.Errorf("option %s type %d", o.Name, o.Type)
		}
	}
	if !cmds[1].Options[0].Required || !cmds[2].Options[0].Required || !cmds[2].Options[0].Autocomplete {
		t.Error("required and autocomplete flags")
	}
}

func TestCommandsAreGuildContextOnly(t *testing.T) {
	raw, err := cmdJSON(Commands()[0])
	if err != nil || !strings.Contains(raw, `"contexts":[0]`) {
		t.Fatalf("%s %v", raw, err)
	}
}

func cmdJSON(c Command) (string, error) {
	b, err := c.MarshalJSON()
	return string(b), err
}

func TestHelpListsEveryCommand(t *testing.T) {
	w := newPressWorld(t)
	rep, err := w.svc.cmdHelp(context.Background(), interactionOf(typeCommand, pressDiscord, map[string]any{"name": "help"}))
	if err != nil {
		t.Fatal(err)
	}
	m := rep.Message
	if len(m.Embeds) != 1 || len(m.Components) != 0 || !m.Ephemeral {
		t.Fatalf("message = %+v", m)
	}
	e := m.Embeds[0]
	if e.Title != specString(t, commandHeading, "`/help` title", 0, nil) {
		t.Fatalf("title = %q", e.Title)
	}
	if e.Footer != "Forever Sixty" || e.Color != notify.BrandEmbed("", "", time.Time{}).Color {
		t.Fatalf("footer %q color %d", e.Footer, e.Color)
	}
	lines := strings.Split(e.Description, "\n")
	if len(lines) != 9 {
		t.Fatalf("%d lines", len(lines))
	}
	desc := descriptionSegments(t)
	for i, c := range Commands() {
		if want := fmt.Sprintf("**/%s** %s", c.Name, desc[c.Name]); lines[i] != want {
			t.Errorf("line %d = %q, want %q", i, lines[i], want)
		}
	}
}
