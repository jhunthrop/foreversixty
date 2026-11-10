package app_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/app"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// installAddon writes <game>/Interface/AddOns/<folder>/<folder>.toc naming version.
func installAddon(t *testing.T, game, folder, version string) {
	t.Helper()
	dir := filepath.Join(game, "Interface", "AddOns", folder)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, folder+".toc"), []byte("## Title: x\n## Version: "+version+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestTheScanListsEachShippedFolderWithTheNewestVersionAnyInstallHas(t *testing.T) {
	older, newer := t.TempDir(), t.TempDir()
	installAddon(t, older, "ForeverSixty", "2.0.30")
	installAddon(t, newer, "ForeverSixty", "2.0.37")
	installAddon(t, newer, "ForeverSixtyData", "data-v2026.11.09")
	installAddon(t, older, "Details", "1.0") // not ours: never reported

	got := app.ScanAddons([]string{older, newer})
	want := []client.InstalledAddon{{Folder: "ForeverSixty", Version: "2.0.37"}, {Folder: "ForeverSixtyData", Version: "2026.11.09"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("scan = %+v, want %+v", got, want)
	}
	if got := app.ScanAddons(nil); got == nil || len(got) != 0 {
		t.Fatalf("scan of no installs = %#v, want an empty list", got)
	}
}

func TestAPairedCompanionTellsTheSiteWhichAddonsAreInstalledOnlyWhenThatChanges(t *testing.T) {
	h := newHarness(t)
	installAddon(t, h.game, "ForeverSixty", "2.0.36")
	h.pairAndPrime()

	reports := h.srv.AddonReports()
	if len(reports) != 1 || !reflect.DeepEqual(reports[0], []client.InstalledAddon{{Folder: "ForeverSixty", Version: "2.0.36"}}) {
		t.Fatalf("reports after pairing = %+v, want the one install", reports)
	}
	// A minute on with nothing changed: no second report.
	h.app.Step(t.Context(), t0.Add(app.AddonScanEvery))
	if n := len(h.srv.AddonReports()); n != 1 {
		t.Fatalf("reports with nothing changed = %d, want 1", n)
	}
	// The addon updates: the next scan says so, not one before it.
	installAddon(t, h.game, "ForeverSixty", "2.0.37")
	h.app.Step(t.Context(), t0.Add(app.AddonScanEvery+time.Second))
	if n := len(h.srv.AddonReports()); n != 1 {
		t.Fatalf("reports inside the minute = %d, want 1", n)
	}
	h.app.Step(t.Context(), t0.Add(2*app.AddonScanEvery))
	reports = h.srv.AddonReports()
	if len(reports) != 2 || reports[1][0].Version != "2.0.37" {
		t.Fatalf("reports after the update = %+v, want 2.0.37 second", reports)
	}
}

func TestAReportTheSiteDidNotTakeIsSentAgainAtTheNextScan(t *testing.T) {
	h := newHarness(t)
	h.srv.FailNext("PUT /v1/devices/self/addons", 1)
	h.pairAndPrime()
	if n := len(h.srv.AddonReports()); n != 0 {
		t.Fatalf("reports after a refused one = %d, want 0", n)
	}
	h.app.Step(t.Context(), t0.Add(app.AddonScanEvery))
	reports := h.srv.AddonReports()
	if len(reports) != 1 || len(reports[0]) != 0 {
		t.Fatalf("reports after the retry = %+v, want one empty list", reports)
	}
}
