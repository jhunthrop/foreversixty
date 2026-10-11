// companion/internal/addon/lua51_test.go
// The real-interpreter check for what the writer renders: the inbox is loaded with Lua 5.1's
// own dofile, the interpreter the game runs, and a check chunk asserts on the table it built.
package addon

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// RunLua51 is runLua51 for the external addon_test package (the fakeapi-driven sync tests).
var RunLua51 = runLua51

// runLua51 writes body to a file and runs check under a real Lua 5.1 after dofile'ing it, so
// check reads the globals the file assigned (ForeverSixtyInbox) and fails with error() or
// assert(). It tries lua5.1 on PATH, then lua5.1 inside WSL for a Windows developer, and
// reports which ran. With neither it fails the test when CI is set (the workflow must install
// lua5.1) and otherwise reports false, leaving the caller to its pure-Go reader.
func runLua51(t testing.TB, body []byte, check string) (ran string, ok bool) {
	t.Helper()
	dir := t.TempDir()
	inbox := filepath.Join(dir, InboxName+".lua")
	script := filepath.Join(dir, "check.lua")
	if err := os.WriteFile(inbox, body, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(script, []byte("dofile(arg[1])\n"+check+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	type route struct {
		argv []string
		path func(string) (string, error)
	}
	native := func(p string) (string, error) { return p, nil }
	viaWSL := func(p string) (string, error) {
		out, err := exec.Command("wsl", "-e", "wslpath", "-a", filepath.ToSlash(p)).Output()
		return strings.TrimSpace(string(out)), err
	}
	for _, r := range []route{{[]string{"lua5.1"}, native}, {[]string{"wsl", "-e", "lua5.1"}, viaWSL}} {
		if _, err := exec.LookPath(r.argv[0]); err != nil {
			continue
		}
		probe := exec.Command(r.argv[0], append(append([]string{}, r.argv[1:]...), "-v")...)
		if err := probe.Run(); err != nil {
			continue // wsl without a distro or without lua5.1: try the next route
		}
		s, err1 := r.path(script)
		f, err2 := r.path(inbox)
		if err1 != nil || err2 != nil {
			continue
		}
		cmd := exec.Command(r.argv[0], append(append([]string{}, r.argv[1:]...), s, f)...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("the rendered inbox fails under Lua 5.1 (%s): %v\n%s\n--- inbox ---\n%s",
				strings.Join(r.argv, " "), err, out, body)
		}
		return strings.Join(r.argv, " "), true
	}
	if os.Getenv("CI") != "" {
		t.Fatal("CI is set and no Lua 5.1 interpreter was found on PATH or in WSL: " +
			"the workflow must install lua5.1 so the rendered inbox is checked against the client's Lua")
	}
	return "", false
}
