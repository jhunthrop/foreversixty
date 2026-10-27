package discordbot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
)

const (
	maxChoicesPerOption = 25
	maxDiscordText      = 100
)

// The BiS data's slot keys, which the slot choice values must be exactly.
var bisSlotKeys = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
	"finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
}

func TestChoicesStayInsideDiscordLimits(t *testing.T) {
	for _, c := range Commands() {
		if len([]rune(c.Description)) > maxDiscordText {
			t.Errorf("/%s description too long", c.Name)
		}
		for _, o := range c.Options {
			if len(o.Choices) > maxChoicesPerOption || len([]rune(o.Description)) > maxDiscordText {
				t.Errorf("/%s %s over a limit", c.Name, o.Name)
			}
			for _, ch := range o.Choices {
				if ch.Name == "" || ch.Value == "" || len([]rune(ch.Name)) > maxDiscordText || len([]rune(ch.Value)) > maxDiscordText {
					t.Errorf("/%s %s choice %+v", c.Name, o.Name, ch)
				}
			}
		}
	}
}

func bisOption(t *testing.T, name string) CommandOption {
	t.Helper()
	for _, o := range Commands()[3].Options {
		if o.Name == name {
			return o
		}
	}
	t.Fatalf("/bis has no %s", name)
	return CommandOption{}
}

func TestSlotAndPhaseChoicesMatchTheSpecAndData(t *testing.T) {
	raw, err := readSpec()
	if err != nil {
		t.Fatal(err)
	}
	slotList := regexp.MustCompile(`choice of 17: ([^)]+)\)`).FindStringSubmatch(raw)
	phaseList := regexp.MustCompile("`phase` \\(choice: ([^)]+)\\)").FindStringSubmatch(raw)
	if slotList == nil || phaseList == nil {
		t.Fatal("spec lists moved")
	}
	for _, c := range []struct {
		option string
		list   string
	}{{"slot", slotList[1]}, {"phase", phaseList[1]}} {
		var names []string
		for _, ch := range bisOption(t, c.option).Choices {
			names = append(names, ch.Name)
		}
		if got := strings.Join(names, ", "); got != c.list {
			t.Errorf("%s names = %q, spec %q", c.option, got, c.list)
		}
	}
	for i, ch := range bisOption(t, "slot").Choices {
		if ch.Value != bisSlotKeys[i] {
			t.Errorf("slot %s value = %q, want the data key %q", ch.Name, ch.Value, bisSlotKeys[i])
		}
	}
}

func TestReleasedCommandsAreRouted(t *testing.T) {
	w := newPressWorld(t)
	for _, c := range Commands() {
		if _, ok := w.svc.commands[c.Name]; c.Released && !ok {
			t.Errorf("released /%s has no handler", c.Name)
		}
	}
}

func dispatchCommand(t *testing.T, w *pressWorld, server, name string) string {
	t.Helper()
	in := interactionOf(typeCommand, pressDiscord, map[string]any{"name": name})
	in.GuildID, in.ReceivedAt = server, w.svc.now()
	rec := httptest.NewRecorder()
	w.svc.dispatch(t.Context(), rec, in)
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
	var body struct {
		Data struct {
			Content string           `json:"content"`
			Embeds  []map[string]any `json:"embeds"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Data.Embeds) > 0 {
		return "embed:" + body.Data.Embeds[0]["title"].(string)
	}
	return body.Data.Content
}

func TestUnhandledCommandAnswersTheFailureSentence(t *testing.T) {
	w := newPressWorld(t)
	if got := dispatchCommand(t, w, pressServer, "char"); got != commonCopy.Failure {
		t.Fatalf("answer = %q", got)
	}
}

func TestHelpAnswersInAnUnboundServerAndADM(t *testing.T) {
	w := newPressWorld(t)
	for _, server := range []string{"srv-unbound", ""} {
		if got := dispatchCommand(t, w, server, "help"); got != "embed:"+commandCopy.HelpTitle {
			t.Errorf("server %q answer = %q", server, got)
		}
	}
	if got := dispatchCommand(t, w, "srv-unbound", "raid"); got != commonCopy.Unbound {
		t.Errorf("raid in an unbound server = %q", got)
	}
}

func readSpec() (string, error) {
	raw, err := os.ReadFile(specPath)
	return string(raw), err
}
