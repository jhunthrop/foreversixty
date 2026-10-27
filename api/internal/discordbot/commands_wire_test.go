package discordbot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

const (
	maxChoicesPerOption = 25
	maxDiscordText      = 100
)

const repoBuildsDir = "../../../data/builds"

// bisSlotKeys is every band's slot keys, in order, from the BiS build the bot reads; every
// file and band must agree, and the slot choice values must be exactly these.
func bisSlotKeys(t *testing.T) []string {
	t.Helper()
	treeData, err := trees.Load(repoBuildsDir)
	if err != nil {
		t.Fatal(err)
	}
	build := (&guilds.Store{DataDir: repoBuildsDir, Trees: treeData}).BisBuild()
	files, err := filepath.Glob(filepath.Join(repoBuildsDir, build, "bis", "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no BiS files for build %q: %v", build, err)
	}
	var keys []string
	for _, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		var file struct {
			Bands []struct {
				Slots []struct {
					Slot string `json:"slot"`
				} `json:"slots"`
			} `json:"bands"`
		}
		if err := json.Unmarshal(raw, &file); err != nil {
			t.Fatalf("%s: %v", f, err)
		}
		for _, b := range file.Bands {
			var got []string
			for _, s := range b.Slots {
				got = append(got, s.Slot)
			}
			if keys == nil {
				keys = got
			} else if !slices.Equal(got, keys) {
				t.Fatalf("%s: slots %v differ from %v", f, got, keys)
			}
		}
	}
	return keys
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
	keys := bisSlotKeys(t)
	choices := bisOption(t, "slot").Choices
	if len(choices) != len(keys) {
		t.Fatalf("%d slot choices, the data has %d slots", len(choices), len(keys))
	}
	for i, ch := range choices {
		if ch.Value != keys[i] {
			t.Errorf("slot %s value = %q, want the data key %q", ch.Name, ch.Value, keys[i])
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
