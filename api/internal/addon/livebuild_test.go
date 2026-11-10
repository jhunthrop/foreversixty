package addon

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// The owner's six beta characters (2026-11-10) carried 70374 and 70450;
// 70532 to 70755 were pre-launch "wowf" builds. Live starts at MinLiveBuild.
func TestPreLiveProblemRefusesOnlyClientBuildsBeforeLaunch(t *testing.T) {
	for export, refused := range map[string]bool{
		"FS1:1.60.1.70374:hunter:orc:0/0/0:":         true,
		"FS1:1.60.1.70450:hunter:orc:0/0/0:|level=5": true,
		"FS1:1.60.1.70755:mage:human:0/0/0:":         true,
		"FS1:1.15.9.69722:mage:human:0/0/0:":         true,
		"FS1:" + MinLiveBuild + ":mage:human:0/0/0:": false,
		"FS1:1.60.1.70800:mage:human:0/0/0:":         false,
		"FS1:1.60.10.1:mage:human:0/0/0:":            false, // numeric, not text, order
		"FS1:aaa":                                    false, // no client build: opaque
		"FS1::mage:human:0/0/0:":                     false,
		"FS2:1.60.1.70374:hunter:orc:0/0/0:":         false, // not FS1
		"":                                           false,
	} {
		problem := preLiveProblem("Bow Jackzon", export)
		if (problem != "") != refused {
			t.Errorf("preLiveProblem(%q) = %q, want refused %v", export, problem, refused)
		}
	}
	got := preLiveProblem("Bow Jackzon", "FS1:1.60.1.70450:hunter:orc:0/0/0:")
	want := "Bow Jackzon's export is from the beta (build 1.60.1.70450), not live WoW Forever. Log in to WoW Forever and export again."
	if got != want {
		t.Errorf("message = %q, want %q", got, want)
	}
}

func TestParseFS1BuildReadsTheHead(t *testing.T) {
	for export, want := range map[string]string{
		"FS1:1.60.1.70790:warrior:orc:0/0/0:|level=60": "1.60.1.70790",
		"FS1:aaa":                      "aaa",
		"FS1:":                         "",
		"FS2:1.60.1.70790:warrior:orc": "",
		"nonsense":                     "",
	} {
		if got, ok := ParseFS1Build(export); got != want || ok != (want != "") {
			t.Errorf("ParseFS1Build(%q) = %q, %v; want %q", export, got, ok, want)
		}
	}
}

// apiError is the error half of a failed envelope.
type apiError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

func errorOf(t *testing.T, res *http.Response) apiError {
	t.Helper()
	defer res.Body.Close()
	var env struct {
		Error apiError `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	return env.Error
}

const betaExport = "FS1:1.60.1.70450:hunter:orc:0/0/0:|level=24"

// A companion that still reads the beta folder sends a beta export: the
// whole list is refused on "characters", which the companion logs and does
// not send again, and nothing is stored.
func TestTheCompanionSyncRefusesABetaExport(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodPost, "/v1/addon/exports",
		`{"characters":[{"name":"Bow Jackzon","ruleset":"normal","region":"us","export":"`+betaExport+`"}]}`)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	e := errorOf(t, res)
	if e.Code != CodePreLiveExport || e.Fields["characters"] == "" ||
		!strings.Contains(e.Message, "Bow Jackzon") || !strings.Contains(e.Message, "1.60.1.70450") {
		t.Fatalf("error = %+v, want pre_live_export naming the character and build", e)
	}
	exports, err := h.store.Exports(context.Background(), h.owner)
	if err != nil || len(exports) != 0 {
		t.Fatalf("exports = %+v, %v; want none stored", exports, err)
	}
}

// The paste box shows the API's message as it stands, so it must be the
// player's sentence.
func TestThePasteRefusesABetaExportInPlainWords(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, "/v1/me/exports",
		`{"exports":[{"name":"Bow Jackzon","ruleset":"normal","region":"us","export":"`+betaExport+`"}]}`)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	if e := errorOf(t, res); e.Message != preLiveProblem("Bow Jackzon", betaExport) {
		t.Fatalf("message = %q", e.Message)
	}
}

// The outbox's own export is held to the same floor.
func TestAnOutboxWithABetaExportIsRefused(t *testing.T) {
	h := newOutboxHarness(t, true)
	outbox := `{"format":1,"export":{"character":"Obnoxious-Yell","code":"` + betaExport + `","saved_at":"2026-11-10 12:00"}}`
	res := h.do(http.MethodPost, "/v1/addon/exports", withOutbox(outbox))
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	if e := errorOf(t, res); !strings.Contains(e.Fields["outbox"], "from the beta") {
		t.Fatalf("error = %+v, want the outbox refused as a beta export", e)
	}
}
