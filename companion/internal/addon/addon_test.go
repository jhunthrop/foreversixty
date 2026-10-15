package addon

import (
	"context"
	"encoding/json"
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

func TestTheInboxRendersTypedMessagesThatReadBack(t *testing.T) {
	body := RenderInbox(Inbox{Messages: []Message{
		{
			Type: MessageUpgrade, Character: "us/hardcore/morrowlyn",
			Slot: "chest", ItemID: 11726, ItemName: `Robe of the "Void"`,
			Source: "Molten Core", Delta: 41.5,
		},
		{
			Type: MessageWeights, Character: "us/hardcore/morrowlyn", Spec: "fury",
			Weights: []WeightEntry{{Stat: "strength", Weight: 1}, {Stat: "hit", Weight: 0}},
			Caps:    []string{"hit"},
		},
		{
			Type: MessageGuild, Character: "us/hardcore/morrowlyn",
			GuildName: "Sanguine", ClaimState: "claimed", PendingApprovals: 2, Rank: "officer",
		},
	}}, t0)
	top, err := ParseLua(string(body))
	if err != nil {
		t.Fatalf("the rendered inbox does not parse: %v\n%s", err, body)
	}
	root := top[InboxGlobal].(*Table)
	if keyOf(root.Fields["format"]) != "2" {
		t.Errorf("format = %q", keyOf(root.Fields["format"]))
	}
	if _, has := root.Fields["version"]; has {
		t.Errorf("the inbox still writes the old version field: %+v", root.Fields)
	}
	messages := root.Fields["messages"].(*Table)
	if len(messages.Items) != 3 {
		t.Fatalf("messages = %+v", messages)
	}

	upgrade := messages.Items[0].(*Table)
	if upgrade.Get("type") != MessageUpgrade || upgrade.Get("character") != "us/hardcore/morrowlyn" ||
		upgrade.Get("slot") != "chest" || keyOf(upgrade.Fields["item_id"]) != "11726" ||
		upgrade.Get("item_name") != `Robe of the "Void"` || upgrade.Get("source") != "Molten Core" ||
		keyOf(upgrade.Fields["delta"]) != "41.5" {
		t.Errorf("upgrade = %+v", upgrade.Fields)
	}

	weights := messages.Items[1].(*Table)
	if weights.Get("type") != MessageWeights || weights.Get("spec") != "fury" {
		t.Errorf("weights = %+v", weights.Fields)
	}
	weightRows := weights.Fields["weights"].(*Table)
	if len(weightRows.Items) != 2 {
		t.Fatalf("weight rows = %+v", weightRows)
	}
	first := weightRows.Items[0].(*Table)
	if first.Get("stat") != "strength" || keyOf(first.Fields["weight"]) != "1" {
		t.Errorf("first weight = %+v", first.Fields)
	}
	caps := weights.Fields["caps"].(*Table)
	if len(caps.Items) != 1 || keyOf(caps.Items[0]) != "hit" {
		t.Errorf("caps = %+v", caps)
	}

	guild := messages.Items[2].(*Table)
	if guild.Get("type") != MessageGuild || guild.Get("guild_name") != "Sanguine" ||
		guild.Get("claim_state") != "claimed" || keyOf(guild.Fields["pending_approvals"]) != "2" ||
		guild.Get("rank") != "officer" {
		t.Errorf("guild = %+v", guild.Fields)
	}

	if string(RenderInbox(Inbox{Messages: []Message{
		{Type: MessageGuild, Character: "us/hardcore/morrowlyn", GuildName: "Sanguine",
			ClaimState: "claimed", PendingApprovals: 2, Rank: "officer"},
	}}, t0)) == string(body) {
		t.Fatal("expected a different render for a different message set")
	}
}

// TestAnUnrecognisedMessageTypeStillRenders proves renderMessage never
// panics or drops the envelope for a Type this companion build does not
// special-case: a future message kind (or an unaddressed one, added by
// a site release ahead of the next companion build) still shows up as
// "type" and, when present, "character" — the fields Follow.inbox-style
// forward compatibility depends on — with none of the other kinds'
// fields spuriously present.
func TestAnUnrecognisedMessageTypeStillRenders(t *testing.T) {
	body := RenderInbox(Inbox{Messages: []Message{{Type: "future_kind", Character: "us/hardcore/morrowlyn"}}}, t0)
	top, err := ParseLua(string(body))
	if err != nil {
		t.Fatalf("did not parse: %v\n%s", err, body)
	}
	row := top[InboxGlobal].(*Table).Fields["messages"].(*Table).Items[0].(*Table)
	if row.Get("type") != "future_kind" || row.Get("character") != "us/hardcore/morrowlyn" {
		t.Errorf("row = %+v", row.Fields)
	}
	if _, has := row.Fields["slot"]; has {
		t.Errorf("an unrecognised type carried an upgrade field: %+v", row.Fields)
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

func (f *fakeAPI) PostAddonExports(_ context.Context, chars []Export, _ *Outbox) (ExportsResult, error) {
	if f.err != nil {
		return ExportsResult{}, f.err
	}
	f.posted = append(f.posted, chars...)
	return ExportsResult{Stored: len(chars)}, nil
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

// TestAChangedMessageRewritesTheInboxEvenWithTheSameBuilds proves the
// deep-equal half of the sync's change check (sync.go, reflect.DeepEqual
// on Messages) actually does something: the builds slice comparison
// above (slices.Equal) is unaware of Messages entirely, so a version
// of this check that only compared builds would pass this test's first
// half and then silently stop noticing a weights or guild message that
// changed underneath an unchanged build list.
func TestAChangedMessageRewritesTheInboxEvenWithTheSameBuilds(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "SavedVariables")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(path, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	same := []Build{{ID: "b1", Name: "Holy", Code: "FSB1:x"}}
	api := &fakeAPI{inbox: Inbox{Builds: same, Messages: []Message{
		{Type: MessageGuild, Character: "us/hardcore/morrowlyn", GuildName: "Sanguine", PendingApprovals: 1},
	}}}
	s := New(Options{Paths: func() []string { return []string{path} }, API: api})
	if err := s.Poll(t.Context(), t0); err != nil {
		t.Fatal(err)
	}
	inbox := InboxPath(path)
	if err := os.Chtimes(inbox, t0, t0); err != nil {
		t.Fatalf("the inbox was not written: %v", err)
	}

	// Same builds, a changed message: the inbox must still be rewritten.
	api.inbox = Inbox{Builds: same, Messages: []Message{
		{Type: MessageGuild, Character: "us/hardcore/morrowlyn", GuildName: "Sanguine", PendingApprovals: 2},
	}}
	if err := s.Poll(t.Context(), t0.Add(InboxEvery)); err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(inbox)
	if err != nil {
		t.Fatal(err)
	}
	if fi.ModTime().Equal(t0) {
		t.Error("a changed message with unchanged builds never reached the addon")
	}
}

// savedVariablesWithOutbox is savedVariables plus the outbox the addon
// leaves at logout (contract 4), written the way the game writes a
// saved-variables table: bracketed keys, positional list entries.
const savedVariablesWithOutbox = `
ForeverSixtyDB = {
	["characters"] = {
		["Obnoxious-Normal"] = {
			["name"] = "Obnoxious",
			["ruleset"] = "normal",
			["region"] = "us",
			["export"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
		},
	},
	["outbox"] = {
		["format"] = 1,
		["export"] = {
			["character"] = "Obnoxious-Yell",
			["code"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
			["saved_at"] = "2026-11-04 14:02",
		},
		["sim_request"] = {
			["character"] = "Obnoxious-Yell",
			["spec"] = "warrior-arms",
			["requested_at"] = "2026-11-04 14:10",
			["gear_code"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
		},
		["roster_snapshot"] = {
			["guild"] = "OLYMPUS XXVII",
			["members"] = {
				{
					["name"] = "Obnoxious-Yell",
					["level"] = 20,
					["spec"] = "warrior-arms",
					["gear_gap"] = 14,
					["build"] = "Arms leveling",
					["seen_at"] = "2026-11-04 14:01",
				}, -- [1]
			},
		},
	},
}
`

func writeSaved(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "ForeverSixty.lua")
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestScanOutboxReadsTheWholeOutbox(t *testing.T) {
	ob, err := ScanOutbox(writeSaved(t, savedVariablesWithOutbox))
	if err != nil {
		t.Fatal(err)
	}
	if ob == nil {
		t.Fatal("the outbox was not found")
	}
	if ob.SavedAt != "2026-11-04 14:02" || ob.RequestedAt != "2026-11-04 14:10" {
		t.Errorf("outbox stamps = %q, %q", ob.SavedAt, ob.RequestedAt)
	}
	var got map[string]any
	if err := json.Unmarshal(ob.Body, &got); err != nil {
		t.Fatalf("the outbox body is not JSON: %v\n%s", err, ob.Body)
	}
	if got["format"] != 1.0 {
		t.Errorf("format = %v", got["format"])
	}
	sim := got["sim_request"].(map[string]any)
	if sim["spec"] != "warrior-arms" || sim["requested_at"] != "2026-11-04 14:10" {
		t.Errorf("sim_request = %+v", sim)
	}
	members := got["roster_snapshot"].(map[string]any)["members"].([]any)
	if len(members) != 1 || members[0].(map[string]any)["level"] != 20.0 {
		t.Errorf("members = %+v", members)
	}
	// The character exports are read from the same file, as before.
	exports, err := ScanExports(writeSaved(t, savedVariablesWithOutbox))
	if err != nil || len(exports) != 1 || exports[0].Name != "Obnoxious" {
		t.Errorf("exports = %+v, %v", exports, err)
	}
}

func TestAnOutboxThatIsMissingOrOfAnotherFormatIsNothingToSend(t *testing.T) {
	for name, body := range map[string]string{
		"missing":   savedVariables,
		"format 2":  strings.Replace(savedVariablesWithOutbox, `["format"] = 1`, `["format"] = 2`, 1),
		"no format": strings.Replace(savedVariablesWithOutbox, `["format"] = 1,`, ``, 1),
		"no export code": strings.Replace(savedVariablesWithOutbox,
			`["code"] = "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627",
			["saved_at"]`, `["saved_at"]`, 1),
	} {
		t.Run(name, func(t *testing.T) {
			path := writeSaved(t, body)
			before, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			ob, err := ScanOutbox(path)
			if err != nil {
				t.Fatal(err)
			}
			if ob != nil {
				t.Errorf("an outbox to send was found: %s", ob.Body)
			}
			after, err := os.ReadFile(path)
			if err != nil || string(after) != string(before) {
				t.Error("reading the outbox changed the saved-variables file")
			}
		})
	}
}

func TestWriteInboxRefusesAnyFileButTheInbox(t *testing.T) {
	dir := t.TempDir()
	game := filepath.Join(dir, SavedVariablesName+".lua")
	if err := os.WriteFile(game, []byte(savedVariables), 0o600); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{game, filepath.Join(dir, "Other.lua"), filepath.Join(dir, "ForeverSixtyInbox.lua.bak")} {
		if wrote, err := WriteInbox(path, RenderInbox(Inbox{}, t0)); err == nil || wrote {
			t.Errorf("WriteInbox(%s) = %v, %v; want a refusal", filepath.Base(path), wrote, err)
		}
	}
	after, err := os.ReadFile(game)
	if err != nil || string(after) != savedVariables {
		t.Fatal("the game's own saved variables were touched")
	}
	if wrote, err := WriteInbox(filepath.Join(dir, InboxName+".lua"), RenderInbox(Inbox{}, t0)); err != nil || !wrote {
		t.Fatalf("the inbox itself was refused: %v, %v", wrote, err)
	}
}

func TestNumbersRenderWithNoExponentAndAtMostThreeDecimals(t *testing.T) {
	for in, want := range map[float64]string{
		4.2: "4.2", 1: "1", 0: "0", 0.1 + 0.2: "0.3", 412.6: "412.6",
		1e21: "1000000000000000000000", 0.00001: "0", 2.71828: "2.718", -3.14159: "-3.142",
	} {
		if got := formatFloat(in); got != want {
			t.Errorf("formatFloat(%v) = %q, want %q", in, got, want)
		}
	}
}

func TestTheInboxDecodesEveryV2TypeAndCarriesOnlyFlatUnknownOnes(t *testing.T) {
	var in Inbox
	err := json.Unmarshal([]byte(`{"format":2,"generated_at":"2026-11-04T21:03:00Z","builds":[],"messages":[
		{"type":"ack","character":"us/normal/obnoxious","export_seen_at":"2026-11-04T14:03:00Z","sim_request_state":"queued","requested_at":"2026-11-04T14:11:00Z"},
		{"type":"future_flat","character":"us/normal/obnoxious","label":"plain","count":3,"ratio":0.25},
		{"type":"future_nested","character":"us/normal/obnoxious","rows":[1,2]},
		{"type":"sim_result","character":"us/normal/obnoxious","spec":"warrior-arms","dps":"not a number"},
		{"type":"loot_values","character":"us/normal/obnoxious","raid":"Molten Core","items":{"17076":{"value":6.3,"unit":"dps","planned":true}}}
	]}`), &in)
	if err != nil {
		t.Fatal(err)
	}
	if in.Format != 2 || in.GeneratedAt != "2026-11-04T21:03:00Z" {
		t.Errorf("envelope = %d %q", in.Format, in.GeneratedAt)
	}
	if len(in.Messages) != 3 {
		t.Fatalf("messages = %+v", in.Messages)
	}
	if in.Messages[0].Type != MessageAck || in.Messages[0].ExportSeenAt != "2026-11-04T14:03:00Z" ||
		in.Messages[0].SimRequestState != "queued" ||
		in.Messages[0].RequestedAt != "2026-11-04T14:11:00Z" {
		t.Errorf("ack = %+v", in.Messages[0])
	}
	flat := in.Messages[1]
	if flat.Type != "future_flat" || flat.Character != "us/normal/obnoxious" ||
		flat.Extra["label"] != "plain" || flat.Extra["count"] != 3.0 || flat.Extra["ratio"] != 0.25 {
		t.Errorf("flat unknown = %+v", flat)
	}
	if loot := in.Messages[2]; loot.Items[17076].Value != 6.3 || !loot.Items[17076].Planned {
		t.Errorf("loot_values = %+v", loot)
	}
	if len(in.Skipped) != 2 || in.Skipped[0].Type != "future_nested" || in.Skipped[0].Position != 3 ||
		in.Skipped[1].Type != MessageSimResult || in.Skipped[1].Position != 4 || in.Skipped[1].Why == "" {
		t.Errorf("skipped = %+v", in.Skipped)
	}

	// The carried unknown renders as itself, keys in a fixed order.
	body := string(RenderInbox(Inbox{Messages: []Message{flat}}, t0))
	row := parseInbox(t, body).Fields["messages"].(*Table).Items[0].(*Table)
	if row.Get("label") != "plain" || keyOf(row.Fields["count"]) != "3" || keyOf(row.Fields["ratio"]) != "0.25" {
		t.Errorf("rendered unknown = %+v", row.Fields)
	}
	if strings.Index(body, `"count"`) > strings.Index(body, `"label"`) {
		t.Errorf("the unknown message's keys are not sorted:\n%s", body)
	}
}

func parseInbox(t *testing.T, body string) *Table {
	t.Helper()
	top, err := ParseLua(body)
	if err != nil {
		t.Fatalf("the rendered inbox does not parse: %v\n%s", err, body)
	}
	return top[InboxGlobal].(*Table)
}

// hostile is every byte sequence a string field could carry that would break a naive Lua
// writer: long-bracket closers, a newline, a quote, a backslash and a NUL.
const hostile = "a]]b]=]c\nd\"e\\f\x00g"

// TestHostileStringsInEveryFieldLoadUnderLua51 puts hostile through every string field the
// v2 messages carry (and the envelope's), renders, and has the real interpreter compare each
// field to the same string built in Lua from its pieces.
func TestHostileStringsInEveryFieldLoadUnderLua51(t *testing.T) {
	h := hostile
	one := 1.5
	in := Inbox{
		Builds: []Build{{ID: h, Name: h, Character: h, Code: h}},
		Messages: []Message{
			{Type: MessageBuild, Character: h, Name: h, Code: h},
			{Type: MessageUpgrade, Character: h, Slot: h, ItemID: 1, ItemName: h, Source: h, Delta: 1, Unit: h},
			{Type: MessageWeights, Character: h, Spec: h, Weights: []WeightEntry{{Stat: h, Weight: 1}},
				Caps: []string{h}, DPSPerPoint: &one},
			{Type: MessageSimResult, Character: h, Spec: h, DPS: &one, Error: &one, GeneratedAt: h,
				Slots: &[]SimSlot{{Slot: h, ItemID: 1, Delta: 1}}},
			{Type: MessageLastRaid, Character: h, ReportID: h, Date: h, Bosses: []LastRaidBoss{{
				Name: h, Percentile: 1, BiggestLoss: h, ConsumablesMissed: []string{h, h}, Adherence: &one}}},
			{Type: MessageLootValues, Character: h, Raid: h, Items: map[int]LootItemValue{17076: {Value: 1, Unit: h}}},
			{Type: MessageGuild, Character: h, GuildName: h, ClaimState: h, Rank: h},
			{Type: MessageAck, Character: h, ExportSeenAt: h, SimRequestState: h, RequestedAt: h},
			{Type: h, Character: h, Extra: map[string]any{h: h}},
		},
	}
	body := RenderInbox(in, t0)

	// The pure-Go reader gets the same strings back.
	root := parseInbox(t, string(body))
	msgs := root.Fields["messages"].(*Table).Items
	if got := msgs[1].(*Table).Get("item_name"); got != h {
		t.Errorf("item_name = %q", got)
	}
	if got := msgs[8].(*Table).Get(h); got != h {
		t.Errorf("an unknown type's hostile key and value = %q", got)
	}

	ran, ok := runLua51(t, body, `
local H = "a]]b]=]c" .. "\n" .. "d" .. string.char(34) .. "e" .. string.char(92) .. "f" .. string.char(0) .. "g"
assert(#H == 16, "the Lua-side string is wrong: " .. #H)
local function eq(v, what) assert(v == H, what .. " = " .. string.format("%q", tostring(v))) end
local I = ForeverSixtyInbox
local b = I.builds[1]
eq(b.id, "builds.id"); eq(b.name, "builds.name"); eq(b.character, "builds.character"); eq(b.code, "builds.code")
local m = I.messages
assert(#m == 9, "#messages = " .. #m)
for i = 1, 9 do eq(m[i].character, "character " .. i) end
eq(m[1].name, "build.name"); eq(m[1].code, "build.code")
eq(m[2].slot, "upgrade.slot"); eq(m[2].item_name, "upgrade.item_name"); eq(m[2].source, "upgrade.source"); eq(m[2].unit, "upgrade.unit")
eq(m[3].spec, "weights.spec"); eq(m[3].weights[1].stat, "weights.stat"); eq(m[3].caps[1], "caps")
eq(m[4].spec, "sim_result.spec"); eq(m[4].generated_at, "sim_result.generated_at"); eq(m[4].slots[1].slot, "slots.slot")
eq(m[5].report_id, "last_raid.report_id"); eq(m[5].date, "last_raid.date")
local boss = m[5].bosses[1]
eq(boss.name, "boss.name"); eq(boss.biggest_loss, "boss.biggest_loss")
eq(boss.consumables_missed[1], "consumables_missed[1]"); eq(boss.consumables_missed[2], "consumables_missed[2]")
eq(m[6].raid, "loot_values.raid"); eq(m[6].items[17076].unit, "items.unit")
eq(m[7].guild_name, "guild_name"); eq(m[7].claim_state, "claim_state"); eq(m[7].rank, "rank")
eq(m[8].export_seen_at, "export_seen_at"); eq(m[8].sim_request_state, "sim_request_state"); eq(m[8].requested_at, "requested_at")
eq(m[9].type, "unknown type"); eq(m[9][H], "unknown field")
`)
	if ok {
		t.Logf("Lua 5.1 check ran through %q", ran)
	}
}
