package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestTheAppConfigRoute(t *testing.T) {
	w := newHomeWorld(t)
	w.svc.Settings = Settings{WebCheckoutRegions: []string{"US", "GB"}, MinVersionIOS: "1.2.0", MinVersionAndroid: "1.1.0"}
	rec := w.get(t, "/v1/app/config")
	if rec.Code != http.StatusOK || rec.Header().Get("Cache-Control") != "public, max-age=300" {
		t.Fatalf("%d %q", rec.Code, rec.Header().Get("Cache-Control"))
	}
	var body struct {
		Data Config `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	c := body.Data
	if strings.Join(c.WebCheckoutRegions, ",") != "US,GB" || c.MinAppVersion.IOS != "1.2.0" || c.MinAppVersion.Android != "1.1.0" ||
		c.ActiveBuild != homeBuild {
		t.Fatalf("config = %+v", c)
	}
	// No region is [], never null: an app reads it as "link nowhere".
	w.svc.Settings.WebCheckoutRegions = nil
	if rec := w.get(t, "/v1/app/config"); !strings.Contains(rec.Body.String(), `"web_checkout_regions":[]`) {
		t.Fatalf("no regions = %s", rec.Body)
	}
	w.svc.Data.(*fakeData).active = ""
	if rec := w.get(t, "/v1/app/config"); rec.Code != http.StatusServiceUnavailable || rec.Header().Get("Cache-Control") != "" {
		t.Fatalf("no build = %d %q", rec.Code, rec.Header().Get("Cache-Control"))
	}
}
