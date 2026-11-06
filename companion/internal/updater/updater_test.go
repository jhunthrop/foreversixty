package updater

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	minisign "github.com/jedisct1/go-minisign"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// keypair mints a minisign key for the test. The workflow's signing
// key is a real minisign secret key; this is the same algorithm with
// no passphrase.
func keypair(t *testing.T) (minisign.PrivateKey, string) {
	t.Helper()
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	var sk minisign.PrivateKey
	sk.SignatureAlgorithm = [2]byte{'E', 'd'}
	copy(sk.SecretKey[:], priv)
	pk := sk.PublicKey()
	raw := append(append(append([]byte{}, pk.SignatureAlgorithm[:]...), pk.KeyId[:]...),
		pk.PublicKey[:]...)
	return sk, base64.StdEncoding.EncodeToString(raw)
}

func sign(t *testing.T, sk minisign.PrivateKey, body []byte) []byte {
	t.Helper()
	sig, err := sk.Sign(body, minisign.SignOptions{Hashed: true})
	if err != nil {
		t.Fatal(err)
	}
	return sig.Encode()
}

// manifest builds the companion manifest a test feed serves, given the
// feed's base URL so the asset URLs can sit on it.
type manifest func(base string) Release

// signed is the manifest the workflow writes: one release with the
// linux/amd64 binary and its signature, each with the digest of the bytes
// the feed serves for it.
func signed(tag string, binary, signature []byte) manifest {
	return func(base string) Release {
		name := AssetName("linux", "amd64")
		return Release{Schema: ManifestSchema, Family: FeedFamily, TagName: tag, Assets: []Asset{
			{Name: name, URL: base + "/companion/a", SHA256: digest(binary)},
			{Name: name + ".minisig", URL: base + "/companion/a.minisig", SHA256: digest(signature)},
		}}
	}
}

func digest(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

// feed serves m at /dl/companion/latest.json (a nil m answers 404, as the
// feed does before the first release) and the two downloads behind it, and
// returns the feed base.
func feed(t *testing.T, m manifest, binary, signature []byte) string {
	t.Helper()
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	base := srv.URL + "/dl"
	mux.HandleFunc("/dl/companion/latest.json", func(w http.ResponseWriter, r *http.Request) {
		if m == nil {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(m(base))
	})
	mux.HandleFunc("/dl/companion/a", func(w http.ResponseWriter, r *http.Request) { w.Write(binary) })
	mux.HandleFunc("/dl/companion/a.minisig", func(w http.ResponseWriter, r *http.Request) { w.Write(signature) })
	return base
}

// release serves a feed whose companion manifest names tag, with one
// signed asset for linux/amd64.
func release(t *testing.T, tag string, binary, signature []byte) string {
	t.Helper()
	return feed(t, signed(tag, binary, signature), binary, signature)
}

func TestNewerComparesDottedVersions(t *testing.T) {
	for _, tc := range []struct {
		tag, current string
		want         bool
	}{
		{"companion-v1.0.1", "1.0.0", true},
		{"companion-v1.10.0", "1.9.3", true},
		{"companion-v1.0.0", "1.0.0", false},
		{"companion-v0.9.9", "1.0.0", false},
		{"companion-v2.0.0", "1.99.99", true},
	} {
		if got := Newer(tc.tag, tc.current); got != tc.want {
			t.Errorf("Newer(%q, %q) = %v", tc.tag, tc.current, got)
		}
	}
}

func TestAssetNamesCarryThePlatform(t *testing.T) {
	if got := AssetName("darwin", "arm64"); got != "foreversixty-companion_darwin_arm64" {
		t.Errorf("darwin = %q", got)
	}
	if got := AssetName("windows", "amd64"); got != "foreversixty-companion_windows_amd64.exe" {
		t.Errorf("windows = %q", got)
	}
}

func TestTheFeedDefaultsToOurOwnDomain(t *testing.T) {
	u := New(Options{})
	if u.o.FeedBase != "https://foreversixty.gg/dl" || u.feed.Base != u.o.FeedBase {
		t.Errorf("feed base = %q / %q", u.o.FeedBase, u.feed.Base)
	}
}

func TestAVerifiedReleaseIsStaged(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("#!/bin/sh\necho the new companion\n")
	base := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	dir := t.TempDir()
	u := New(Options{Dir: dir, FeedBase: base, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})

	staged, err := u.Poll(t.Context(), t0)
	if err != nil || !staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	got, err := os.ReadFile(PendingPath(dir))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(binary) {
		t.Fatalf("staged %q", got)
	}
	// The six-hour rule: a second poll straight away does nothing.
	if staged, err := u.Poll(t.Context(), t0.Add(time.Hour)); err != nil || staged {
		t.Fatalf("an early second poll = %v, %v", staged, err)
	}
}

func TestATamperedDownloadIsRefused(t *testing.T) {
	sk, pub := keypair(t)
	good := []byte("the real companion")
	// The manifest's digest matches the trojan, as it would if the feed
	// itself were rewritten: the signature is what refuses it.
	base := release(t, "companion-v9.9.9", []byte("a trojan"), sign(t, sk, good))
	dir := t.TempDir()
	u := New(Options{Dir: dir, FeedBase: base, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil || !strings.Contains(err.Error(), "does not verify") {
		t.Fatalf("a tampered download: %v", err)
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Fatal("the tampered binary was staged anyway")
	}
}

func TestADownloadThatDoesNotMatchItsDigestIsRefused(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("the real companion")
	sig := sign(t, sk, binary)
	m := func(base string) Release {
		rel := signed("companion-v9.9.9", binary, sig)(base)
		rel.Assets[0].SHA256 = digest([]byte("something else"))
		return rel
	}
	dir := t.TempDir()
	u := New(Options{Dir: dir, FeedBase: feed(t, m, binary, sig), PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil || !strings.Contains(err.Error(), "sha256") {
		t.Fatalf("a digest mismatch: %v", err)
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Fatal("a binary that missed its digest was staged")
	}
}

func TestADownloadOffTheFeedsOriginIsRefused(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("from somewhere else")
	sig := sign(t, sk, binary)
	hits := 0
	elsewhere := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		w.Write(binary)
	}))
	t.Cleanup(elsewhere.Close)
	m := func(base string) Release {
		rel := signed("companion-v9.9.9", binary, sig)(base)
		rel.Assets[0].URL = elsewhere.URL + "/a"
		return rel
	}
	dir := t.TempDir()
	u := New(Options{Dir: dir, FeedBase: feed(t, m, binary, sig), PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil || !strings.Contains(err.Error(), "off the feed's origin") {
		t.Fatalf("a cross-origin download: %v", err)
	}
	if hits != 0 {
		t.Errorf("the other host was asked %d times", hits)
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Fatal("a cross-origin binary was staged")
	}
}

func TestAFeedBaseThatIsNotHTTPSIsRefused(t *testing.T) {
	for _, base := range []string{"http://foreversixty.gg/dl", "ftp://127.0.0.1/dl", "not a url"} {
		f := Feed{Base: base, HTTP: http.DefaultClient}
		if _, err := f.Latest(t.Context(), FeedFamily); err == nil {
			t.Errorf("%q was read", base)
		}
	}
}

func TestAReleaseThatIsNotNewerIsIgnored(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("same version")
	base := release(t, "companion-v1.0.0", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), FeedBase: base, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	var buf bytes.Buffer
	u.o.Log = slog.New(slog.NewTextHandler(&buf, nil))
	if staged, err := u.Poll(t.Context(), t0); err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	if !strings.Contains(buf.String(), "newest companion release 1.0.0 is not newer than 1.0.0") {
		t.Errorf("log:\n%s", buf.String())
	}
}

func TestABuildWithNoKeyNeverUpdates(t *testing.T) {
	u := New(Options{Dir: t.TempDir(), FeedBase: "http://127.0.0.1:1", Version: "1.0.0"})
	staged, err := u.Poll(t.Context(), t0)
	if err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	if _, err := u.Stage(t.Context(), Release{TagName: "companion-v2"}); !errors.Is(err, ErrNoKey) {
		t.Fatalf("Stage = %v, want ErrNoKey", err)
	}
}

func TestAReleaseMissingThisPlatformIsNotAnUpdate(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("linux only")
	base := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), FeedBase: base, PublicKey: pub, Version: "1.0.0",
		GOOS: "windows", GOARCH: "amd64"})
	if staged, err := u.Poll(t.Context(), t0); err != nil || staged {
		t.Fatalf("Poll = %v, %v; a release with no Windows asset is skipped, not an error", staged, err)
	}
}

func TestAReleaseWithoutASignatureIsNotAnUpdate(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("unsigned")
	sig := sign(t, sk, binary)
	m := func(base string) Release {
		rel := signed("companion-v9.9.9", binary, sig)(base)
		rel.Assets = rel.Assets[:1]
		return rel
	}
	u := New(Options{Dir: t.TempDir(), FeedBase: feed(t, m, binary, sig), PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if staged, err := u.Poll(t.Context(), t0); err != nil || staged {
		t.Fatalf("Poll = %v, %v; a binary with no signature is no update, not an error", staged, err)
	}
}

func TestNoManifestYetIsNoUpdateAndNoError(t *testing.T) {
	_, pub := keypair(t)
	var buf bytes.Buffer
	u := New(Options{Dir: t.TempDir(), FeedBase: feed(t, nil, nil, nil), PublicKey: pub, Version: "0.2.5",
		GOOS: "linux", GOARCH: "amd64", Log: slog.New(slog.NewTextHandler(&buf, nil))})
	if staged, err := u.Poll(t.Context(), t0); err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	if !strings.Contains(buf.String(), "no companion release") {
		t.Errorf("log:\n%s", buf.String())
	}
}

func TestAManifestTagOfAnotherFamilyIsNoUpdate(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("x")
	base := release(t, "addon-v9.9.9", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), FeedBase: base, PublicKey: pub, Version: "0.2.5",
		GOOS: "linux", GOARCH: "amd64"})
	if rel, err := u.Latest(t.Context()); err != nil || rel.TagName != "" {
		t.Fatalf("Latest = %q, %v", rel.TagName, err)
	}
}

func TestAManifestOfAnotherSchemaOrFamilyIsAnError(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("x")
	sig := sign(t, sk, binary)
	for name, change := range map[string]func(*Release){
		"schema": func(r *Release) { r.Schema = 2 },
		"family": func(r *Release) { r.Family = "addon" },
	} {
		m := func(base string) Release {
			rel := signed("companion-v9.9.9", binary, sig)(base)
			change(&rel)
			return rel
		}
		u := New(Options{Dir: t.TempDir(), FeedBase: feed(t, m, binary, sig), PublicKey: pub, Version: "0.2.5",
			GOOS: "linux", GOARCH: "amd64"})
		if _, err := u.Poll(t.Context(), t0); err == nil {
			t.Errorf("a manifest with another %s was read", name)
		}
	}
}

func TestAnHTTPErrorFromTheFeedIsAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "rate limited", http.StatusForbidden)
	}))
	t.Cleanup(srv.Close)
	_, pub := keypair(t)
	u := New(Options{Dir: t.TempDir(), FeedBase: srv.URL, PublicKey: pub, Version: "0.2.5",
		GOOS: "linux", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil || !strings.Contains(err.Error(), "release feed answered 403") {
		t.Fatalf("a 403 from the feed: %v", err)
	}
}

func TestApplyPendingSwapsTheBinaryAndVerifiesAgain(t *testing.T) {
	sk, pub := keypair(t)
	dir := t.TempDir()
	exe := filepath.Join(t.TempDir(), "foreversixty-companion")
	if err := os.WriteFile(exe, []byte("the old companion"), 0o755); err != nil {
		t.Fatal(err)
	}
	if swapped, err := ApplyPending(dir, pub, exe); err != nil || swapped {
		t.Fatalf("with nothing pending = %v, %v", swapped, err)
	}

	body := []byte("the new companion")
	if err := os.WriteFile(PendingPath(dir), body, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "pending.minisig"), sign(t, sk, body), 0o600); err != nil {
		t.Fatal(err)
	}
	swapped, err := ApplyPending(dir, pub, exe)
	if err != nil || !swapped {
		t.Fatalf("ApplyPending = %v, %v", swapped, err)
	}
	got, err := os.ReadFile(exe)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(body) {
		t.Fatalf("the binary is %q", got)
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Error("the pending file was left behind")
	}
	// Both halves of the swap must be same-directory renames: the
	// install location and the application directory can sit on
	// different volumes, where a rename across them fails outright.
	if _, err := os.Stat(exe + ".previous"); err != nil {
		t.Errorf("the replaced binary is not beside the new one: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "previous")); err == nil {
		t.Error("the replaced binary was moved into the update directory")
	}
}

func TestApplyPendingRefusesAndClearsAnUnsignedStage(t *testing.T) {
	_, pub := keypair(t)
	dir := t.TempDir()
	exe := filepath.Join(t.TempDir(), "foreversixty-companion")
	if err := os.WriteFile(exe, []byte("old"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(PendingPath(dir), []byte("unsigned"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := ApplyPending(dir, pub, exe); err == nil {
		t.Fatal("an unsigned stage was applied")
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Error("the unsigned stage was not cleared")
	}
	if got, _ := os.ReadFile(exe); string(got) != "old" {
		t.Errorf("the running binary changed to %q", got)
	}
}

// TestApplyPendingRejectsAStageTamperedWithAfterStaging proves the
// apply-time re-verify actually does something: a pending file
// modified on disk after it was staged, with its original
// well-formed signature still sitting next to it, must be refused --
// this is the exact "it sat on disk in between" scenario the package
// exists for.
func TestApplyPendingRejectsAStageTamperedWithAfterStaging(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("the real companion")
	base := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	dir := t.TempDir()
	u := New(Options{Dir: dir, FeedBase: base, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if staged, err := u.Poll(t.Context(), t0); err != nil || !staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}

	// Tamper with the staged bytes after the fact, leaving the
	// original signature untouched.
	if err := os.WriteFile(PendingPath(dir), []byte("a trojan planted after staging"), 0o755); err != nil {
		t.Fatal(err)
	}

	exeDir := t.TempDir()
	exe := filepath.Join(exeDir, "foreversixty-companion")
	if err := os.WriteFile(exe, []byte("the running companion"), 0o755); err != nil {
		t.Fatal(err)
	}

	if _, err := ApplyPending(dir, pub, exe); err == nil {
		t.Fatal("a stage tampered with after staging was applied")
	}
	if got, _ := os.ReadFile(exe); string(got) != "the running companion" {
		t.Errorf("the running binary changed to %q", got)
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Error("the tampered stage was not cleared")
	}
	if entries, err := os.ReadDir(exeDir); err != nil {
		t.Fatal(err)
	} else if len(entries) != 1 {
		t.Errorf("stray files left next to the executable: %v", entries)
	}
}

func TestVerifyReportsWhichPieceIsWrong(t *testing.T) {
	sk, pub := keypair(t)
	body := []byte("payload")
	if err := Verify(pub, body, sign(t, sk, body)); err != nil {
		t.Fatal(err)
	}
	if err := Verify("not base64 at all", body, sign(t, sk, body)); err == nil {
		t.Error("a broken public key was accepted")
	}
	if err := Verify(pub, body, []byte("not a signature")); err == nil {
		t.Error("a broken signature was accepted")
	}
	if err := Verify(pub, []byte("other"), sign(t, sk, body)); err == nil {
		t.Error("a signature over other bytes was accepted")
	}
}

func stagePending(t *testing.T, dir string, sk minisign.PrivateKey, body []byte) {
	t.Helper()
	if err := os.WriteFile(PendingPath(dir), body, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "pending.minisig"), sign(t, sk, body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func newExe(t *testing.T) string {
	t.Helper()
	exe := filepath.Join(t.TempDir(), "foreversixty-companion")
	if err := os.WriteFile(exe, []byte("old"), 0o755); err != nil {
		t.Fatal(err)
	}
	return exe
}

func TestAnUpdateIsFollowedByTheNewBuildStarting(t *testing.T) {
	sk, pub := keypair(t)
	dir, exe := t.TempDir(), newExe(t)
	stagePending(t, dir, sk, []byte("new"))
	var buf bytes.Buffer
	var gotPath, gotDir string
	var gotArgs []string
	var order []string
	ok := ApplyAndRelaunch(LaunchOptions{
		Dir: dir, PublicKey: pub, Exe: exe, Args: []string{"-debug"}, WorkDir: "/work",
		Log:         slog.New(slog.NewTextHandler(&buf, nil)),
		BeforeStart: func() { order = append(order, "release") },
		Start: func(p string, a []string, d string) error {
			order = append(order, "start")
			gotPath, gotArgs, gotDir = p, a, d
			b, _ := os.ReadFile(p)
			if string(b) != "new" {
				t.Errorf("started %q before the new binary was in place", b)
			}
			return nil
		},
	})
	if !ok {
		t.Fatal("the caller was told to carry on")
	}
	if gotPath != exe || len(gotArgs) != 1 || gotArgs[0] != "-debug" || gotDir != "/work" {
		t.Errorf("started %q %v in %q", gotPath, gotArgs, gotDir)
	}
	if len(order) != 2 || order[0] != "release" {
		t.Errorf("order %v", order)
	}
	if !strings.Contains(buf.String(), "starting the new build") {
		t.Errorf("no log line:\n%s", buf.String())
	}
}

func TestAFailedStartKeepsRunningAndSaysTheUpdateStartsNextLaunch(t *testing.T) {
	sk, pub := keypair(t)
	dir, exe := t.TempDir(), newExe(t)
	stagePending(t, dir, sk, []byte("new"))
	var buf bytes.Buffer
	reacquired := false
	ok := ApplyAndRelaunch(LaunchOptions{
		Dir: dir, PublicKey: pub, Exe: exe,
		Log:              slog.New(slog.NewTextHandler(&buf, nil)),
		Start:            func(string, []string, string) error { return errors.New("access denied") },
		AfterFailedStart: func() { reacquired = true },
	})
	if ok {
		t.Fatal("told to exit with no new build running")
	}
	if !reacquired {
		t.Error("the claim was not retaken")
	}
	if !strings.Contains(buf.String(), "level=WARN") || !strings.Contains(buf.String(), "the update was installed; it starts on the next launch") {
		t.Errorf("log:\n%s", buf.String())
	}
	if b, _ := os.ReadFile(exe); string(b) != "new" {
		t.Errorf("exe is %q", b)
	}
}

func TestRefusedUpdatesLogAndLeaveARunnableBinary(t *testing.T) {
	sk, pub := keypair(t)
	_, otherPub := keypair(t)
	cases := map[string]func(dir string){
		"missing minisig": func(dir string) {
			os.WriteFile(PendingPath(dir), []byte("new"), 0o755)
		},
		"signature mismatch": func(dir string) { stagePending(t, dir, sk, []byte("new")) },
	}
	for name, setup := range cases {
		t.Run(name, func(t *testing.T) {
			dir, exe := t.TempDir(), newExe(t)
			setup(dir)
			key := pub
			if name == "signature mismatch" {
				key = otherPub
			}
			var buf bytes.Buffer
			started := false
			ok := ApplyAndRelaunch(LaunchOptions{Dir: dir, PublicKey: key, Exe: exe,
				Log:   slog.New(slog.NewTextHandler(&buf, nil)),
				Start: func(string, []string, string) error { started = true; return nil }})
			if ok || started {
				t.Fatal("started a build that was refused")
			}
			if !strings.Contains(buf.String(), "level=WARN") {
				t.Errorf("no WARN line:\n%s", buf.String())
			}
			if b, _ := os.ReadFile(exe); string(b) != "old" {
				t.Errorf("exe is %q", b)
			}
		})
	}
}

func TestASharingViolationOnTheSwapLeavesTheOldBinaryAndALogLine(t *testing.T) {
	sk, pub := keypair(t)
	dir, exe := t.TempDir(), newExe(t)
	stagePending(t, dir, sk, []byte("new"))
	orig := rename
	defer func() { rename = orig }()
	rename = func(from, to string) error {
		if from == exe {
			return errors.New("The process cannot access the file because it is being used by another process.")
		}
		return orig(from, to)
	}
	var buf bytes.Buffer
	ok := ApplyAndRelaunch(LaunchOptions{Dir: dir, PublicKey: pub, Exe: exe,
		Log:   slog.New(slog.NewTextHandler(&buf, nil)),
		Start: func(string, []string, string) error { t.Error("started"); return nil }})
	if ok || !strings.Contains(buf.String(), "level=WARN") {
		t.Fatalf("ok=%v log:\n%s", ok, buf.String())
	}
	if b, _ := os.ReadFile(exe); string(b) != "old" {
		t.Errorf("exe is %q", b)
	}
	ents, _ := os.ReadDir(filepath.Dir(exe))
	if len(ents) != 1 {
		t.Errorf("leftovers beside the exe: %v", ents)
	}
}

func TestAFailedSecondRenameRestoresThePreviousBinary(t *testing.T) {
	sk, pub := keypair(t)
	dir, exe := t.TempDir(), newExe(t)
	stagePending(t, dir, sk, []byte("new"))
	orig := rename
	defer func() { rename = orig }()
	rename = func(from, to string) error {
		if to == exe && strings.Contains(from, ".update-") {
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
}
