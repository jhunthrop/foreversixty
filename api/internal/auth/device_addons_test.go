package auth

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func putAddons(t *testing.T, h *harness, token, body string) *http.Response {
	t.Helper()
	r, err := http.NewRequest(http.MethodPut, h.server.URL+"/v1/devices/self/addons", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")
	if token != "" {
		r.Header.Set("Authorization", "Bearer "+token)
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func listedDevice(t *testing.T, h *harness) Device {
	t.Helper()
	var list []Device
	h.decode(t, h.do(t, http.MethodGet, "/v1/devices", ""), &list)
	if len(list) != 1 {
		t.Fatalf("devices = %+v, want the one companion", list)
	}
	return list[0]
}

func TestACompanionReportsItsAddonsAndTheDeviceListCarriesThem(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	token := h.pairDeviceToken(t)

	if got := listedDevice(t, h).Addons; got != nil {
		t.Fatalf("addons before any report = %+v, want omitted", *got)
	}
	res := putAddons(t, h, token,
		`{"addons":[{"folder":"ForeverSixty","version":"2.0.37"},{"folder":"ForeverSixtyData","version":"2026.11.09"}]}`)
	wantStatus(t, res, http.StatusNoContent, "a companion's report")
	want := []DeviceAddon{{Folder: "ForeverSixty", Version: "2.0.37"}, {Folder: "ForeverSixtyData", Version: "2026.11.09"}}
	if got := listedDevice(t, h).Addons; got == nil || !reflect.DeepEqual(*got, want) {
		t.Fatalf("addons = %+v, want %+v", got, want)
	}

	// A later scan that finds none replaces the report with an empty list, not an absent one.
	wantStatus(t, putAddons(t, h, token, `{"addons":[]}`), http.StatusNoContent, "a report of none")
	if got := listedDevice(t, h).Addons; got == nil || len(*got) != 0 {
		t.Fatalf("addons after an empty report = %+v, want []", got)
	}
}

func TestADeviceAddonsReportIsTheCompanionsOwnAndWellFormed(t *testing.T) {
	h := newHarness(t)
	h.signIn(t, "raider@example.com")
	token := h.pairDeviceToken(t)

	res := h.do(t, http.MethodPut, "/v1/devices/self/addons", `{"addons":[]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("a browser session's report = %d, want 401", res.StatusCode)
	}
	for name, body := range map[string]string{
		"not json":       `{`,
		"no list":        `{}`,
		"unknown folder": `{"addons":[{"folder":"Details","version":"1"}]}`,
		"listed twice":   `{"addons":[{"folder":"ForeverSixty","version":"1"},{"folder":"ForeverSixty","version":"2"}]}`,
		"odd version":    `{"addons":[{"folder":"ForeverSixty","version":"<b>"}]}`,
		"long version":   `{"addons":[{"folder":"ForeverSixty","version":"` + strings.Repeat("1", maxAddonVersion+1) + `"}]}`,
	} {
		wantStatus(t, putAddons(t, h, token, body), http.StatusBadRequest, name)
	}
	if got := listedDevice(t, h).Addons; got != nil {
		t.Fatalf("addons after only refused reports = %+v, want omitted", *got)
	}
}
