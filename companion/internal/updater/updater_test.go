package updater

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
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

// release serves a GitHub-shaped latest release with one asset.
func release(t *testing.T, tag string, binary, signature []byte) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	name := AssetName("linux", "amd64")
	mux.HandleFunc("/repos/"+Repo+"/releases/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"tag_name":%q,"assets":[{"name":%q,"browser_download_url":"%s/a"},`+
			`{"name":"%s.minisig","browser_download_url":"%s/a.minisig"}]}`,
			tag, name, srv.URL, name, srv.URL)
	})
	mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) { w.Write(binary) })
	mux.HandleFunc("/a.minisig", func(w http.ResponseWriter, r *http.Request) { w.Write(signature) })
	return srv
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

func TestAVerifiedReleaseIsStaged(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("#!/bin/sh\necho the new companion\n")
	srv := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	dir := t.TempDir()
	u := New(Options{Dir: dir, APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
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
	srv := release(t, "companion-v9.9.9", []byte("a trojan"), sign(t, sk, good))
	dir := t.TempDir()
	u := New(Options{Dir: dir, APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil {
		t.Fatal("a tampered download was accepted")
	}
	if _, err := os.Stat(PendingPath(dir)); err == nil {
		t.Fatal("the tampered binary was staged anyway")
	}
}

func TestAReleaseThatIsNotNewerIsIgnored(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("same version")
	srv := release(t, "companion-v1.0.0", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "linux", GOARCH: "amd64"})
	if staged, err := u.Poll(t.Context(), t0); err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
}

func TestABuildWithNoKeyNeverUpdates(t *testing.T) {
	u := New(Options{Dir: t.TempDir(), APIBase: "http://127.0.0.1:1", Version: "1.0.0"})
	staged, err := u.Poll(t.Context(), t0)
	if err != nil || staged {
		t.Fatalf("Poll = %v, %v", staged, err)
	}
	if _, err := u.Stage(t.Context(), Release{TagName: "companion-v2"}); !errors.Is(err, ErrNoKey) {
		t.Fatalf("Stage = %v, want ErrNoKey", err)
	}
}

func TestAReleaseMissingThisPlatformIsAnError(t *testing.T) {
	sk, pub := keypair(t)
	binary := []byte("linux only")
	srv := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	u := New(Options{Dir: t.TempDir(), APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
		GOOS: "windows", GOARCH: "amd64"})
	if _, err := u.Poll(t.Context(), t0); err == nil {
		t.Fatal("a release with no Windows asset was accepted")
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
	srv := release(t, "companion-v9.9.9", binary, sign(t, sk, binary))
	dir := t.TempDir()
	u := New(Options{Dir: dir, APIBase: srv.URL, PublicKey: pub, Version: "1.0.0",
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
