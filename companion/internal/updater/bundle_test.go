package updater

import (
	"archive/zip"
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const appName = "Forever Sixty.app"

// zipEntry is one file, folder or link of a test zip.
type zipEntry struct {
	name, body string
	mode       fs.FileMode
}

func makeZip(t *testing.T, entries ...zipEntry) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for _, e := range entries {
		h := &zip.FileHeader{Name: e.name, Method: zip.Deflate}
		h.SetMode(e.mode)
		w, err := zw.CreateHeader(h)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.Write([]byte(e.body)); err != nil {
			t.Fatal(err)
		}
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

// appZip is what ditto -c -k --keepParent makes of a release bundle.
func appZip(t *testing.T, binary string) []byte {
	return makeZip(t,
		zipEntry{name: appName + "/", mode: fs.ModeDir | 0o755},
		zipEntry{name: appName + "/Contents/", mode: fs.ModeDir | 0o755},
		zipEntry{name: appName + "/Contents/Info.plist", body: "<plist>new</plist>", mode: 0o644},
		zipEntry{name: appName + "/Contents/MacOS/foreversixty-companion", body: binary, mode: 0o755},
		zipEntry{name: appName + "/Contents/Resources/seal.icns", body: "icon", mode: 0o644},
		zipEntry{name: appName + "/Contents/_CodeSignature/CodeResources", body: "seal", mode: 0o644},
		zipEntry{name: appName + "/Contents/Resources/current", body: "seal.icns", mode: fs.ModeSymlink | 0o755},
		zipEntry{name: "__MACOSX/" + appName + "/._Contents", body: "fork", mode: 0o644},
	)
}

// newBundle lays out an installed bundle and returns its main executable.
func newBundle(t *testing.T) (bundle, exe string) {
	t.Helper()
	bundle = filepath.Join(t.TempDir(), "Applications", appName)
	exe = filepath.Join(bundle, "Contents", "MacOS", "foreversixty-companion")
	if err := os.MkdirAll(filepath.Dir(exe), 0o755); err != nil {
		t.Fatal(err)
	}
	for path, body := range map[string]string{exe: "old", filepath.Join(bundle, "Contents", "Info.plist"): "<plist>old</plist>"} {
		if err := os.WriteFile(path, []byte(body), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return bundle, exe
}

func TestBundleRootIsTheAppAroundAMainExecutable(t *testing.T) {
	for exe, want := range map[string]string{
		"/Applications/Forever Sixty.app/Contents/MacOS/foreversixty-companion":     "/Applications/Forever Sixty.app",
		"/Users/jh/Applications/X.app/Contents/MacOS/foreversixty-companion":        "/Users/jh/Applications/X.app",
		"/Users/jh/Downloads/foreversixty-companion":                                "",
		"/Applications/Forever Sixty.app/Contents/Resources/foreversixty-companion": "",
		"/opt/Forever Sixty/Contents/MacOS/foreversixty-companion":                  "",
	} {
		if got := BundleRoot(filepath.FromSlash(exe)); got != filepath.FromSlash(want) {
			t.Errorf("BundleRoot(%q) = %q, want %q", exe, got, want)
		}
	}
}

func TestApplyPendingSwapsTheWholeBundle(t *testing.T) {
	sk, pub := keypair(t)
	dir := t.TempDir()
	bundle, exe := newBundle(t)
	stagePending(t, dir, sk, appZip(t, "new"))

	swapped, err := ApplyPending(dir, pub, exe)
	if err != nil || !swapped {
		t.Fatalf("ApplyPending = %v, %v", swapped, err)
	}
	if b, _ := os.ReadFile(exe); string(b) != "new" {
		t.Errorf("exe is %q", b)
	}
	if fi, err := os.Stat(exe); err != nil || fi.Mode().Perm()&0o100 == 0 {
		t.Errorf("the new executable is not executable: %v %v", fi, err)
	}
	if b, _ := os.ReadFile(filepath.Join(bundle, "Contents", "Info.plist")); string(b) != "<plist>new</plist>" {
		t.Errorf("Info.plist is %q: the bundle was not replaced whole", b)
	}
	if target, err := os.Readlink(filepath.Join(bundle, "Contents", "Resources", "current")); err != nil || target != "seal.icns" {
		t.Errorf("link = %q, %v", target, err)
	}
	// Nothing is left beside the bundle: no previous copy, no staging folder,
	// no Finder shadows.
	entries, _ := os.ReadDir(filepath.Dir(bundle))
	if len(entries) != 1 || entries[0].Name() != appName {
		t.Errorf("beside the bundle: %v", entries)
	}
	if _, err := os.Stat(PendingPath(dir)); !errors.Is(err, fs.ErrNotExist) {
		t.Error("the stage was not cleared")
	}
}

func TestABundleUpdateNeverLandsOnABareBinaryOrTheReverse(t *testing.T) {
	sk, pub := keypair(t)

	dir, exe := t.TempDir(), newExe(t)
	stagePending(t, dir, sk, appZip(t, "new"))
	if _, err := ApplyPending(dir, pub, exe); err == nil || !strings.Contains(err.Error(), "not running from one") {
		t.Fatalf("a bundle over a bare binary: %v", err)
	}
	if b, _ := os.ReadFile(exe); string(b) != "old" {
		t.Errorf("exe is %q", b)
	}

	dir = t.TempDir()
	_, bexe := newBundle(t)
	stagePending(t, dir, sk, []byte("\xcf\xfa\xed\xfe a bare Mach-O"))
	if _, err := ApplyPending(dir, pub, bexe); err == nil || !strings.Contains(err.Error(), "not an app bundle") {
		t.Fatalf("a bare binary into a bundle: %v", err)
	}
	if b, _ := os.ReadFile(bexe); string(b) != "old" {
		t.Errorf("bundle exe is %q", b)
	}
	if _, err := os.Stat(PendingPath(dir)); !errors.Is(err, fs.ErrNotExist) {
		t.Error("a stage that can never apply was kept")
	}
}

func TestAMalformedBundleZipIsRefusedAndTheAppKept(t *testing.T) {
	for name, body := range map[string][]byte{
		"an entry outside the app": makeZip(t, zipEntry{name: appName + "/../evil", body: "x", mode: 0o644}),
		"two top-level folders": makeZip(t,
			zipEntry{name: appName + "/Contents/MacOS/foreversixty-companion", body: "x", mode: 0o755},
			zipEntry{name: "Other.app/x", body: "x", mode: 0o644}),
		"not an app":            makeZip(t, zipEntry{name: "folder/Contents/MacOS/foreversixty-companion", body: "x", mode: 0o755}),
		"a link out of the app": makeZip(t, zipEntry{name: appName + "/Contents/MacOS/foreversixty-companion", body: "x", mode: 0o755}, zipEntry{name: appName + "/l", body: "../../etc", mode: fs.ModeSymlink | 0o755}),
		"an absolute link":      makeZip(t, zipEntry{name: appName + "/Contents/MacOS/foreversixty-companion", body: "x", mode: 0o755}, zipEntry{name: appName + "/l", body: "/etc", mode: fs.ModeSymlink | 0o755}),
		"no main executable":    makeZip(t, zipEntry{name: appName + "/Contents/Info.plist", body: "x", mode: 0o644}),
		"only Finder shadows":   makeZip(t, zipEntry{name: "__MACOSX/._x", body: "x", mode: 0o644}),
		"a truncated zip":       append([]byte(nil), appZip(t, "new")[:40]...),
	} {
		t.Run(name, func(t *testing.T) {
			sk, pub := keypair(t)
			dir := t.TempDir()
			bundle, exe := newBundle(t)
			stagePending(t, dir, sk, body)
			if swapped, err := ApplyPending(dir, pub, exe); err == nil || swapped {
				t.Fatalf("ApplyPending = %v, %v", swapped, err)
			}
			if b, _ := os.ReadFile(exe); string(b) != "old" {
				t.Errorf("exe is %q", b)
			}
			entries, _ := os.ReadDir(filepath.Dir(bundle))
			if len(entries) != 1 {
				t.Errorf("beside the bundle: %v", entries)
			}
		})
	}
}

func TestAFailedBundleSwapRestoresTheRunningApp(t *testing.T) {
	sk, pub := keypair(t)
	dir := t.TempDir()
	bundle, exe := newBundle(t)
	stagePending(t, dir, sk, appZip(t, "new"))
	orig := rename
	defer func() { rename = orig }()
	rename = func(from, to string) error {
		if to == bundle {
			if strings.HasSuffix(from, "previous.app") {
				return orig(from, to) // the restore is allowed
			}
			return errors.New("denied")
		}
		return orig(from, to)
	}
	if _, err := ApplyPending(dir, pub, exe); err == nil {
		t.Fatal("no error")
	}
	if b, _ := os.ReadFile(exe); string(b) != "old" {
		t.Errorf("exe is %q", b)
	}
	// A swap that failed for a reason that may pass keeps the stage for the
	// next launch, as the binary swap does.
	if _, err := os.Stat(PendingPath(dir)); err != nil {
		t.Errorf("the stage was dropped: %v", err)
	}
}
