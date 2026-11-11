package wow

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// flavorInfo writes a .flavor.info the way the game does (owner's install,
// 2026-11-09: "Product Flavor!STRING:0" then the product, CRLF on Windows).
func flavorInfo(t *testing.T, dir, product string) {
	t.Helper()
	body := "Product Flavor!STRING:0\r\n" + product + "\r\n"
	if err := os.WriteFile(filepath.Join(dir, FlavorInfo), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

// The owner's folder: _anniversary_, _classic_beta_, _forever_ and _retail_,
// each with its .flavor.info. The live game alone is found: the ended beta
// is another game now (2026-11-10, beta characters reached the site).
func TestTheOwnersInstallIsReadByProduct(t *testing.T) {
	root := t.TempDir()
	for folder, product := range map[string]string{
		"_anniversary_": "wow_anniversary", "_classic_beta_": "wow_classic_beta", "_forever_": "wowf", "_retail_": "wow",
	} {
		flavorInfo(t, fakeInstall(t, root, folder, "ACCOUNT#1").Path, product)
	}
	got := Scan(root)
	if len(got) != 1 {
		t.Fatalf("Scan = %+v, want _forever_ alone", got)
	}
	if got[0].Flavor != LiveFlavor || got[0].Product != "wowf" || filepath.Base(got[0].Path) != "_forever_" {
		t.Errorf("found %+v, want the live game", got[0])
	}
}

// A renamed or custom install is the live game when its file says wowf: it
// is listed and named WoW Forever, and the beta beside it is not.
func TestAFolderNamedAnythingIsTheLiveGameWhenItsFileSaysWowf(t *testing.T) {
	root := t.TempDir()
	beta := fakeInstall(t, root, "_classic_beta_", "ACCOUNT#1")
	flavorInfo(t, beta.Path, "wow_classic_beta")
	custom := fakeInstall(t, root, "WoWF Live", "ACCOUNT#1")
	flavorInfo(t, custom.Path, "WOWF")
	got := Scan(root)
	if len(got) != 1 || got[0].Path != custom.Path {
		t.Fatalf("Scan = %+v, want the renamed live folder alone", got)
	}
	if got[0].Flavor != LiveFlavor || FlavorName(got[0].Flavor) != "WoW Forever" {
		t.Errorf("renamed live = %+v", got[0])
	}
	// Picked directly, it is accepted too.
	if picked, err := Pick(custom.Path); err != nil || len(picked) != 1 || picked[0].Path != custom.Path {
		t.Errorf("Pick(renamed live) = %+v, %v", picked, err)
	}
}

// The product wins over the folder name: a _forever_ folder that holds
// another game is refused, and a beta product in a folder named _forever_
// is the beta.
func TestTheProductWinsOverAMismatchedFolderName(t *testing.T) {
	root := t.TempDir()
	wrong := fakeInstall(t, root, "_forever_", "ACCOUNT#1")
	flavorInfo(t, wrong.Path, "wow_anniversary")
	if got := Scan(root); len(got) != 0 {
		t.Errorf("Scan = %+v, want nothing: the file says Anniversary", got)
	}
	_, err := Pick(wrong.Path)
	var r *Refusal
	if !errors.As(err, &r) || r.Words() != "That is Classic Anniversary. "+PickWords {
		t.Errorf("Pick(_forever_ holding Anniversary) = %v, want the refusal naming Classic Anniversary", err)
	}
	// An unknown product in a Forever-named folder is refused without a name.
	odd := t.TempDir()
	flavorInfo(t, fakeInstall(t, odd, "_forever_", "ACCOUNT#1").Path, "wowxptr")
	if _, err := Pick(filepath.Join(odd, "_forever_")); !errors.As(err, &r) || r.Words() != PickWords {
		t.Errorf("Pick(_forever_ holding an unknown product) = %v", err)
	}

	// A beta product in a folder named _forever_ is the ended beta: refused,
	// and named.
	other := t.TempDir()
	beta := fakeInstall(t, other, "_forever_", "ACCOUNT#1")
	flavorInfo(t, beta.Path, "wow_classic_beta")
	if got := Scan(other); len(got) != 0 {
		t.Errorf("Scan = %+v, want nothing: the file says the beta", got)
	}
	if _, err := Pick(beta.Path); !errors.As(err, &r) || r.Words() != "That is Classic beta. "+PickWords {
		t.Errorf("Pick(_forever_ holding the beta) = %v, want the refusal naming Classic beta", err)
	}
}

// Without the file the folder name decides, as before; an empty or
// header-only file is no product at all.
func TestTheFolderNameDecidesWhenTheFileIsMissingOrEmpty(t *testing.T) {
	root := t.TempDir()
	live := fakeInstall(t, root, "_forever_", "ACCOUNT#1")
	retail := fakeInstall(t, root, "_retail_", "ACCOUNT#1")
	if err := os.WriteFile(filepath.Join(retail.Path, FlavorInfo), []byte("Product Flavor!STRING:0\r\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := productAt(retail.Path); got != "" {
		t.Errorf("header-only product = %q", got)
	}
	got := Scan(root)
	if len(got) != 1 || got[0].Path != live.Path || got[0].Product != "" {
		t.Errorf("Scan = %+v, want _forever_ by its name", got)
	}
}
