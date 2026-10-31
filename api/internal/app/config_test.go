package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// getConfig asks for the config as the app named by xApp ("" sends no header).
func (w *homeWorld) getConfig(t *testing.T, xApp string) (*httptest.ResponseRecorder, Config) {
	t.Helper()
	mux := http.NewServeMux()
	Mount(mux, w.svc)
	req := httptest.NewRequest(http.MethodGet, "/v1/app/config", nil)
	if xApp != "" {
		req.Header.Set(AppHeader, xApp)
	}
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	var body struct {
		Data Config `json:"data"`
	}
	if rec.Code == http.StatusOK {
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
			t.Fatal(err)
		}
	}
	return rec, body.Data
}

// decisionL is the production default: iOS links in the US, Android nowhere.
var decisionL = map[string][]string{"ios": {"US"}, "android": {}}

func TestTheAppConfigRoute(t *testing.T) {
	w := newHomeWorld(t)
	w.svc.Settings = Settings{WebCheckoutRegions: decisionL, MinVersionIOS: "1.2.0", MinVersionAndroid: "1.1.0",
		Features: []string{"raids"}}
	rec, c := w.getConfig(t, "ios/1.2.0 (42)")
	if rec.Code != http.StatusOK || rec.Header().Get("Cache-Control") != "private, max-age=300" {
		t.Fatalf("%d %q", rec.Code, rec.Header().Get("Cache-Control"))
	}
	if !strings.Contains(strings.Join(rec.Header().Values("Vary"), ","), AppHeader) {
		t.Fatalf("Vary = %q, want X-App", rec.Header().Values("Vary"))
	}
	if c.MinAppVersion.IOS != "1.2.0" || c.MinAppVersion.Android != "1.1.0" || c.ActiveBuild != homeBuild ||
		len(c.Features) != 1 || !c.Features["raids"] {
		t.Fatalf("config = %+v", c)
	}
	w.svc.Data.(*fakeData).active = ""
	if rec, _ := w.getConfig(t, "ios/1.2.0 (42)"); rec.Code != http.StatusServiceUnavailable || rec.Header().Get("Cache-Control") != "" {
		t.Fatalf("no build = %d %q", rec.Code, rec.Header().Get("Cache-Control"))
	}
}

// Decision L: the checkout regions follow the X-App platform, and anything else links nowhere.
func TestTheCheckoutRegionsFollowThePlatform(t *testing.T) {
	w := newHomeWorld(t)
	w.svc.Settings = Settings{WebCheckoutRegions: decisionL, MinVersionIOS: "1.0.0", MinVersionAndroid: "1.0.0"}
	for name, c := range map[string]struct{ xApp, want string }{
		"ios":              {"ios/1.0.0 (1)", `"web_checkout_regions":["US"]`},
		"android":          {"android/1.0.0 (1)", `"web_checkout_regions":[]`},
		"no header":        {"", `"web_checkout_regions":[]`},
		"unknown platform": {"windows/1.0.0 (1)", `"web_checkout_regions":[]`},
		"malformed":        {"ios", `"web_checkout_regions":[]`},
	} {
		rec, _ := w.getConfig(t, c.xApp)
		if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), c.want) {
			t.Errorf("%s: %d %s, want %s", name, rec.Code, rec.Body, c.want)
		}
	}
}

// No switch turned on is {}, never null or absent: the apps keep their bundled defaults.
func TestTheFeaturesMapIsAlwaysPresent(t *testing.T) {
	w := newHomeWorld(t)
	rec, _ := w.getConfig(t, "android/1.0.0 (1)")
	if !strings.Contains(rec.Body.String(), `"features":{}`) {
		t.Fatalf("no features = %s", rec.Body)
	}
}

func TestAppPlatform(t *testing.T) {
	for header, want := range map[string]string{
		"ios/1.2.0 (42)": "ios", "android/1.0.0 (7)": "android", "iOS/1.0.0 (1)": "ios",
		"": "", "ios": "", "ios/": "",
	} {
		if got := AppPlatform(header); got != want {
			t.Errorf("AppPlatform(%q) = %q, want %q", header, got, want)
		}
	}
}
