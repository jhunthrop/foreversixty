package addon

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// savedVariables is the shape the Phase 2 addon writes.
const savedVariables = `
ForeverSixtyDB = {
	["profileKeys"] = {
		["Morrowlyn - Sanguine"] = "Default",
	},
	["characters"] = {
		["Morrowlyn-Hardcore"] = {
			["name"] = "Morrowlyn",
			["realm"] = "Hardcore",
			["region"] = "us",
			["level"] = 60,
			["export"] = "FS1:1.15.9.69722:paladin:human:503200000/0/0:head=12640,chest=11726",
		},
		["Thalgrit-Normal"] = {
			["name"] = "Thalgrit",
			["ruleset"] = "normal",
			["region"] = "eu",
			["export"] = "FS1:1.15.9.69722:warrior:orc:0/310000000/0:head=12640",
		},
	},
	["dataBuild"] = "1.15.9.69722",
}
`

func TestScanExportsReadsEveryCharacter(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ForeverSixty.lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := ScanExports(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("exports = %+v", got)
	}
	if got[0].Name != "Morrowlyn" || got[0].Ruleset != "Hardcore" || got[0].Region != "us" ||
		!strings.HasPrefix(got[0].Export, "FS1:") {
		t.Errorf("first = %+v", got[0])
	}
	if got[1].Name != "Thalgrit" || got[1].Ruleset != "normal" {
		t.Errorf("second = %+v", got[1])
	}
}

func TestTheLuaReaderHandlesTheShapesSavedVariablesUses(t *testing.T) {
	src := `
-- a comment
--[[ a block
     comment ]]
Simple = "a \"quoted\" value\nwith a newline"
Numbers = { 1, -2.5, 1e3 }
Flags = { ["on"] = true, ["off"] = false, ["none"] = nil }
Named = { key = "value", ["bracket"] = "other" }
`
	top, err := ParseLua(src)
	if err != nil {
		t.Fatal(err)
	}
	if top["Simple"] != "a \"quoted\" value\nwith a newline" {
		t.Errorf("Simple = %q", top["Simple"])
	}
	nums, ok := top["Numbers"].(*Table)
	if !ok || len(nums.Items) != 3 || nums.Items[1] != -2.5 || nums.Items[2] != 1000.0 {
		t.Errorf("Numbers = %+v", top["Numbers"])
	}
	flags := top["Flags"].(*Table)
	if flags.Fields["on"] != true || flags.Fields["off"] != false {
		t.Errorf("Flags = %+v", flags.Fields)
	}
	named := top["Named"].(*Table)
	if named.Get("key") != "value" || named.Get("bracket") != "other" {
		t.Errorf("Named = %+v", named.Fields)
	}
}

func TestAMalformedFileIsAnErrorWithALineNumber(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ForeverSixty.lua")
	if err := os.WriteFile(path, []byte("A = {\nB = \"unterminated\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := ScanExports(path)
	if err == nil || !strings.Contains(err.Error(), "lua line") {
		t.Fatalf("err = %v", err)
	}
	if _, err := ScanExports(filepath.Join(t.TempDir(), "missing.lua")); err == nil {
		t.Fatal("a missing file was not an error")
	}
}

func TestTheInboxRendersLuaThatReadsBack(t *testing.T) {
	body := RenderInbox(Inbox{Builds: []Build{
		{ID: "b1", Name: `Holy "Bubble"`, Character: "Morrowlyn", Code: "FSB1:1.15.9.69722:paladin:001:head=12640"},
	}}, t0)
	top, err := ParseLua(string(body))
	if err != nil {
		t.Fatalf("the rendered inbox does not parse: %v\n%s", err, body)
	}
	root, ok := top[InboxGlobal].(*Table)
	if !ok {
		t.Fatalf("the inbox does not assign %s: %s", InboxGlobal, body)
	}
	if root.Get("generated_at") != "2026-12-09T20:00:00Z" {
		t.Errorf("generated_at = %q", root.Get("generated_at"))
	}
	builds := root.Fields["builds"].(*Table)
	if len(builds.Items) != 1 {
		t.Fatalf("builds = %+v", builds)
	}
	first := builds.Items[0].(*Table)
	if first.Get("name") != `Holy "Bubble"` || first.Get("code") != "FSB1:1.15.9.69722:paladin:001:head=12640" {
		t.Errorf("build = %+v", first.Fields)
	}
	if string(RenderInbox(Inbox{Builds: []Build{{ID: "b1", Name: `Holy "Bubble"`,
		Character: "Morrowlyn", Code: "FSB1:1.15.9.69722:paladin:001:head=12640"}}}, t0)) != string(body) {
		t.Error("two renders of the same inbox differ")
	}
}

func TestWriteInboxSkipsAnIdenticalFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "SavedVariables", InboxName+".lua")
	body := RenderInbox(Inbox{}, t0)
	wrote, err := WriteInbox(path, body)
	if err != nil || !wrote {
		t.Fatalf("first write = %v, %v", wrote, err)
	}
	wrote, err = WriteInbox(path, body)
	if err != nil || wrote {
		t.Fatalf("second write = %v, %v", wrote, err)
	}
	wrote, err = WriteInbox(path, RenderInbox(Inbox{}, t0.Add(time.Hour)))
	if err != nil || !wrote {
		t.Fatalf("changed write = %v, %v", wrote, err)
	}
}

func TestInboxPathSitsBesideTheSavedVariables(t *testing.T) {
	got := InboxPath("/w/WTF/Account/A#1/SavedVariables/ForeverSixty.lua")
	want := filepath.Join("/w/WTF/Account/A#1/SavedVariables", InboxName+".lua")
	if got != want {
		t.Fatalf("InboxPath = %q, want %q", got, want)
	}
}

// fakeAPI records what the sync sent and answers with a fixed inbox.
type fakeAPI struct {
	posted []Export
	inbox  Inbox
	err    error
}

func (f *fakeAPI) PostAddonExports(_ context.Context, chars []Export) error {
	if f.err != nil {
		return f.err
	}
	f.posted = append(f.posted, chars...)
	return nil
}

func (f *fakeAPI) AddonInbox(context.Context) (Inbox, error) { return f.inbox, nil }

func TestTheSyncUploadsOnAModificationAndWritesTheInboxOnItsTimer(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, t0, t0); err != nil {
		t.Fatal(err)
	}
	api := &fakeAPI{inbox: Inbox{Builds: []Build{{ID: "b1", Name: "Holy", Code: "FSB1:x"}}}}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api})

	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("posted %+v", api.posted)
	}
	if _, err := os.Stat(InboxPath(path)); err != nil {
		t.Fatalf("the inbox was not written: %v", err)
	}

	// Nothing changed: no second upload, and no inbox refresh until
	// the timer comes round.
	if err := s.Poll(t.Context(), t0.Add(time.Minute)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("a quiet pass uploaded again: %+v", api.posted)
	}

	// The player logs out: the file's modification time moves.
	if err := os.Chtimes(path, t0.Add(time.Hour), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if err := s.Poll(t.Context(), t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 4 {
		t.Fatalf("a logout did not upload: %+v", api.posted)
	}
}

func TestAFailedUploadIsRetriedOnTheNextPass(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	api := &fakeAPI{err: errors.New("offline")}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api})
	if err := s.Poll(t.Context(), t0); err == nil {
		t.Fatal("a failed upload was not reported")
	}
	api.err = nil
	if err := s.Poll(t.Context(), t0.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 2 {
		t.Fatalf("the retry did not happen: %+v", api.posted)
	}
}

func TestAMissingSavedVariablesFileIsNotAnError(t *testing.T) {
	api := &fakeAPI{}
	s := New(Options{
		Paths: func() []string { return []string{filepath.Join(t.TempDir(), "nope.lua")} },
		API:   api,
	})
	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	if len(api.posted) != 0 {
		t.Errorf("posted %+v", api.posted)
	}
}

func TestTheInboxIsNotRewrittenWhenTheBuildsHaveNotChanged(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	api := &fakeAPI{inbox: Inbox{Builds: []Build{{ID: "b1", Name: "Holy", Code: "FSB1:x"}}}}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api})
	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	inbox := InboxPath(path)
	if err := os.Chtimes(inbox, t0, t0); err != nil {
		t.Fatalf("the inbox was not written: %v", err)
	}

	// The ten-minute timer comes round and the site has the same
	// builds: the bytes are identical and the file is left alone.
	if err := s.Poll(t.Context(), t0.Add(InboxEvery)); err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(inbox)
	if err != nil {
		t.Fatal(err)
	}
	if !fi.ModTime().Equal(t0) {
		t.Error("the inbox was rewritten though the builds had not changed")
	}

	// A build the player added on the site does reach the addon.
	api.inbox = Inbox{Builds: append(api.inbox.Builds, Build{ID: "b2", Name: "Prot", Code: "FSB1:y"})}
	if err := s.Poll(t.Context(), t0.Add(2*InboxEvery)); err != nil {
		t.Fatal(err)
	}
	fi, err = os.Stat(inbox)
	if err != nil {
		t.Fatal(err)
	}
	if fi.ModTime().Equal(t0) {
		t.Error("a new build never reached the addon")
	}
	body, err := os.ReadFile(inbox)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "FSB1:y") {
		t.Errorf("the rewritten inbox does not carry the new build:\n%s", body)
	}
}
