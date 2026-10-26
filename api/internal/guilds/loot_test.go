// api/internal/guilds/loot_test.go
package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// seedHunterCharacter seeds a verified hunter-marksmanship character with full BiS gear
// (16963 Helm of Wrath, one of Onyxia's own real drops and this spec's real band-60 pick
// for head per data/builds/1.60.1.70009/bis/hunter-marksmanship.json) so loot candidate
// ranking has a real band to match against.
func seedHunterCharacter(t *testing.T, h *httpHarness, gid int64, email, key string, headItemID int) int64 {
	t.Helper()
	uid := seedUser(t, h.pool, email)
	seedCharacter(t, h.pool, gid, uid, key, "member", true)
	export := fmt.Sprintf(`FS1:1.60.1.70009:hunter:dwarf:555555555/33/0:head=%d|level=60`, headItemID)
	if _, err := h.pool.Exec(context.Background(), `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ($1, $2, 'us', 'hardcore', $3, $4, now(), now())`, key, uid, key, export); err != nil {
		t.Fatal(err)
	}
	recomputeMembership(t, h.pool, gid, uid)
	if _, err := h.pool.Exec(context.Background(), `
		insert into fight_metrics (report_id, fight_index, player_key, class, spec, role, ilvl, faction, fought_at)
		values ($1, 0, $2, 'hunter', 'Marksmanship', 'dps', 60, 'alliance', now())`,
		"loot-fixture-"+key, key); err != nil {
		t.Fatal(err)
	}
	return uid
}

// TestLootRequiresMembership checks the member/officer-only gate.
func TestLootRequiresMembership(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	stranger := seedUser(t, h.pool, "loot-stranger@example.com")
	h.actor = auth.Actor{UserID: stranger, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/loot", gid), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a non-member reading loot = %d, want 403", res.StatusCode)
	}
}

// TestLootListsOnyxiasRealTwentyTwoItems checks the item list carries all 22 real drops
// with their real name/slot/icon, and that a slotless quest item never carries candidates.
//
// It does not assert any roster character appears as a candidate for a head item: checked
// against every one of data/builds/1.60.1.70009/bis/*.json's own band-60-alliance entries,
// none of Onyxia's 22 item ids appears as any spec's pick or alternative anywhere - the
// leveling-BiS catalogue this endpoint reads (api/internal/bis) only ever covers leveling
// gear up to 60, not raid-tier drops, since no raid-tier sim data has been published into
// those files yet. Candidate ranking's own logic (namesItem/lootCandidateLess) is unit
// tested directly below instead, against a hand-built band - see CONTROL_CENTRE.md for
// this gap, named as a real data limitation rather than a bug.
func TestLootListsOnyxiasRealTwentyTwoItems(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedHunterCharacter(t, h, gid, "loot-member@example.com", "us/hardcore/lootmember", 21524)

	member := seedUser(t, h.pool, "loot-reader@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/lootreader", "member", true)
	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}

	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/loot", gid), "")
	var view LootView
	h.data(res, &view)

	// Twenty-two drops, listed as 21 rows: the two "Head of Onyxia" items are one row of two.
	if len(view.Items) != 21 {
		t.Fatalf("items = %d, want 21", len(view.Items))
	}
	heads := 0
	for _, it := range view.Items {
		if it.Name == "Head of Onyxia" {
			heads++
			if it.Copies != 2 {
				t.Errorf("Head of Onyxia copies = %d, want 2", it.Copies)
			}
		} else if it.Copies != 1 {
			t.Errorf("%s copies = %d, want 1", it.Name, it.Copies)
		}
	}
	if heads != 1 {
		t.Fatalf("Head of Onyxia rows = %d, want 1", heads)
	}
	if view.Selected != 1084 || len(view.Encounters) != 1 || view.Encounters[0].Name != "Onyxia" {
		t.Fatalf("encounters/selected = %+v/%d, want Onyxia/1084", view.Encounters, view.Selected)
	}

	var helmOfWrath, scaleOfOnyxia *LootItemView
	for i := range view.Items {
		switch view.Items[i].ItemID {
		case 16963:
			helmOfWrath = &view.Items[i]
		case 15410:
			scaleOfOnyxia = &view.Items[i]
		}
	}
	if helmOfWrath == nil || helmOfWrath.Slot != "head" || helmOfWrath.Icon == "" || helmOfWrath.Name != "Helm of Wrath" {
		t.Fatalf("Helm of Wrath = %+v, want a named head item with an icon", helmOfWrath)
	}
	if scaleOfOnyxia == nil || scaleOfOnyxia.Slot != "" || len(scaleOfOnyxia.Candidates) != 0 {
		t.Fatalf("Scale of Onyxia = %+v, want no slot and no candidates", scaleOfOnyxia)
	}
}

func f64(v float64) *float64 { return &v }
func i(v int) *int           { return &v }

// TestCandidateSortWithinTier0ByGainDps unit-tests tier-0 (BiS-matched) ranking directly,
// independent of whether any real BiS file happens to name a raid item - see
// TestLootListsOnyxiasRealTwentyTwoItems's own doc comment for why that gap exists today.
func TestCandidateSortWithinTier0ByGainDps(t *testing.T) {
	cands := []LootCandidate{
		{CharacterKey: "low-gain", GainDps: f64(5), Attendance: Attendance{Present: 8}},
		{CharacterKey: "high-gain", GainDps: f64(40), Attendance: Attendance{Present: 2}},
		{CharacterKey: "tie-low-attendance", GainDps: f64(5), Attendance: Attendance{Present: 3}},
		{CharacterKey: "not-sim-checked", GainDps: nil, Attendance: Attendance{Present: 20}},
	}
	sortLootCandidates(cands)
	want := []string{"high-gain", "low-gain", "tie-low-attendance", "not-sim-checked"}
	if got := candidateKeys(cands); !equalStrings(got, want) {
		t.Fatalf("sorted = %v, want %v", got, want)
	}
}

// TestCandidateSortTier0AlwaysAheadOfTier1AndTier1ByIlvlDelta checks the follow-up fix's
// own ranking rule: every BiS-matched (tier 0) candidate sorts ahead of every fallback
// (tier 1) one regardless of either's own numbers, and tier 1 among itself sorts by
// ilvl_delta desc, then attendance.
func TestCandidateSortTier0AlwaysAheadOfTier1AndTier1ByIlvlDelta(t *testing.T) {
	cands := []LootCandidate{
		{CharacterKey: "fallback-high-delta", tier: 1, IlvlDelta: i(20), Attendance: Attendance{Present: 1}},
		{CharacterKey: "bis-tiny-gain", tier: 0, GainDps: f64(0.1), Attendance: Attendance{Present: 1}},
		{CharacterKey: "fallback-low-delta", tier: 1, IlvlDelta: i(5), Attendance: Attendance{Present: 8}},
		{CharacterKey: "fallback-tie-attendance", tier: 1, IlvlDelta: i(5), Attendance: Attendance{Present: 2}},
	}
	sortLootCandidates(cands)
	want := []string{"bis-tiny-gain", "fallback-high-delta", "fallback-low-delta", "fallback-tie-attendance"}
	if got := candidateKeys(cands); !equalStrings(got, want) {
		t.Fatalf("sorted = %v, want %v", got, want)
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestFallbackVerdict unit-tests the fallback tier's own pure item-level rule: a higher
// item level than what is worn qualifies; an empty slot always qualifies at the drop's
// own full item level; an equal or lower item level never qualifies.
func TestFallbackVerdict(t *testing.T) {
	cases := []struct {
		name          string
		dropLevel     int
		wornLevel     *int
		wantDelta     int
		wantQualifies bool
	}{
		{"empty slot", 70, nil, 70, true},
		{"real upgrade", 70, i(60), 10, true},
		{"equal level", 70, i(70), 0, false},
		{"downgrade", 60, i(70), 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			delta, qualifies := fallbackVerdict(c.dropLevel, c.wornLevel)
			if qualifies != c.wantQualifies {
				t.Fatalf("qualifies = %v, want %v", qualifies, c.wantQualifies)
			}
			if qualifies && delta != c.wantDelta {
				t.Fatalf("delta = %d, want %d", delta, c.wantDelta)
			}
		})
	}
}

func candidateKeys(cands []LootCandidate) []string {
	out := make([]string, len(cands))
	for i, c := range cands {
		out[i] = c.CharacterKey
	}
	return out
}

// TestLootAwardCreateAndDeleteIsOfficerOnly checks a plain member is refused, an officer
// can award and a second award for the same drop replaces the first, and delete removes it.
func TestLootAwardCreateAndDeleteIsOfficerOnly(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	grantOfficerTools(t, h, gid)
	member := seedHunterCharacter(t, h, gid, "loot-award-member@example.com", "us/hardcore/lootawardmember", 21524)
	officerUID := seedUser(t, h.pool, "loot-award-officer@example.com")
	seedCharacter(t, h.pool, gid, officerUID, "us/hardcore/lootawardofficer", "officer", true)
	recomputeMembership(t, h.pool, gid, officerUID)

	body := `{"encounter_id":1084,"item_id":16963,"character_key":"us/hardcore/lootawardmember"}`

	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), body)
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a plain member awarding loot = %d, want 403", res.StatusCode)
	}

	h.actor = auth.Actor{UserID: officerUID, Role: "user", Method: "session"}
	res = h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), body)
	var created map[string]any
	h.data(res, &created)
	awardID, ok := created["award_id"].(float64)
	if !ok || awardID <= 0 {
		t.Fatalf("created = %+v, want a positive award_id", created)
	}

	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/loot", gid), "")
	var view LootView
	h.data(res, &view)
	var helm *LootItemView
	for i := range view.Items {
		if view.Items[i].ItemID == 16963 {
			helm = &view.Items[i]
		}
	}
	if helm == nil || helm.AwardedTo == nil || helm.AwardedTo.CharacterKey != "us/hardcore/lootawardmember" {
		t.Fatalf("awarded_to = %+v, want the awarded character", helm)
	}
	// The page undoes an award by the id it reads on the item, so the view carries the id the
	// create call returned.
	if helm.AwardedTo.AwardID != int64(awardID) {
		t.Fatalf("awarded_to.award_id = %d, want %d (the id the award was created with)", helm.AwardedTo.AwardID, int64(awardID))
	}

	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/loot/awards/%d", gid, int64(awardID)), "")
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a plain member deleting an award = %d, want 403", res.StatusCode)
	}

	h.actor = auth.Actor{UserID: officerUID, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/loot/awards/%d", gid, int64(awardID)), "")
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("officer deleting an award = %d, want 200", res.StatusCode)
	}

	res = h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/loot", gid), "")
	h.data(res, &view)
	for _, item := range view.Items {
		if item.ItemID == 16963 && item.AwardedTo != nil {
			t.Fatalf("awarded_to after delete = %+v, want nil", item.AwardedTo)
		}
	}

	// Undoing it twice finds nothing the second time.
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/loot/awards/%d", gid, int64(awardID)), "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("undoing an award that is already gone = %d, want 404", res.StatusCode)
	}
}

// TestLootAwardUndoIsScopedToItsGuild checks an officer of another guild cannot undo this
// guild's award by guessing its id.
func TestLootAwardUndoIsScopedToItsGuild(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	grantOfficerTools(t, h, gid)
	seedHunterCharacter(t, h, gid, "undo-scope-member@example.com", "us/hardcore/undoscopemember", 21524)
	officerUID := seedUser(t, h.pool, "undo-scope-officer@example.com")
	seedCharacter(t, h.pool, gid, officerUID, "us/hardcore/undoscopeofficer", "officer", true)
	recomputeMembership(t, h.pool, gid, officerUID)
	h.actor = auth.Actor{UserID: officerUID, Role: "user", Method: "session"}
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid),
		`{"encounter_id":1084,"item_id":16963,"character_key":"us/hardcore/undoscopemember"}`)
	var created map[string]any
	h.data(res, &created)
	awardID := int64(created["award_id"].(float64))

	otherGuild := seedGuild(t, h.pool, "Other Forever")
	grantOfficerTools(t, h, otherGuild)
	otherOfficer := seedUser(t, h.pool, "undo-scope-other@example.com")
	seedCharacter(t, h.pool, otherGuild, otherOfficer, "us/hardcore/undoscopeother", "officer", true)
	recomputeMembership(t, h.pool, otherGuild, otherOfficer)
	h.actor = auth.Actor{UserID: otherOfficer, Role: "user", Method: "session"}
	res = h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/loot/awards/%d", otherGuild, awardID), "")
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("another guild's officer undoing this award = %d, want 404", res.StatusCode)
	}
	var n int
	if err := h.pool.QueryRow(context.Background(), `select count(*) from loot_awards where id = $1`, awardID).Scan(&n); err != nil || n != 1 {
		t.Fatalf("the award should still exist, count = %d err = %v", n, err)
	}
}

// TestLootAwardRejectsAnUnknownItem checks CreateLootAward's own validation.
func TestLootAwardRejectsAnUnknownItem(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	grantOfficerTools(t, h, gid)
	officerUID := seedUser(t, h.pool, "loot-invalid-officer@example.com")
	seedCharacter(t, h.pool, gid, officerUID, "us/hardcore/lootinvalidofficer", "officer", true)
	recomputeMembership(t, h.pool, gid, officerUID)

	h.actor = auth.Actor{UserID: officerUID, Role: "user", Method: "session"}
	body := `{"encounter_id":1084,"item_id":999999,"character_key":"us/hardcore/lootinvalidofficer"}`
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), body)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("awarding an unknown item = %d, want 400", res.StatusCode)
	}
}

func TestMain_lootStrings(t *testing.T) {
	// Sanity check: every onyxiaLoot name is non-empty, and the table has no duplicate ids.
	seen := map[int]bool{}
	for _, it := range onyxiaLoot {
		if it.Name == "" || strings.TrimSpace(it.Name) == "" {
			t.Fatalf("item %d has no name", it.ItemID)
		}
		if seen[it.ItemID] {
			t.Fatalf("duplicate item id %d", it.ItemID)
		}
		seen[it.ItemID] = true
	}
	if len(onyxiaLoot) != 22 {
		t.Fatalf("onyxiaLoot has %d entries, want 22", len(onyxiaLoot))
	}
}

// repoDataDirForTrees resolves ../../../data/builds from this test file's own location -
// the same pattern api/internal/bis/gain_test.go's repoDataDir uses - so these tests need
// no environment variable or working-directory assumption.
func repoDataDirForTrees(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not resolve this test file's own path")
	}
	return filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "data", "builds")
}

func realTreesStore(t *testing.T) *Store {
	t.Helper()
	dataDir := repoDataDirForTrees(t)
	treeData, err := trees.Load(dataDir)
	if err != nil {
		t.Fatalf("trees.Load: %v", err)
	}
	return &Store{DataDir: dataDir, Trees: treeData}
}

// TestBuildItemForResolvesARealOnyxiaDropByClass checks buildItemFor against the real
// data: Helm of Wrath (16963) lives in warrior.json at item level 76, and a class whose
// item table never carries it (e.g. a made-up slug) is refused.
func TestBuildItemForResolvesARealOnyxiaDropByClass(t *testing.T) {
	s := realTreesStore(t)
	it, ok := s.buildItemFor("warrior", 16963)
	if !ok || it.ItemLevel != 76 || it.Slot != "head" {
		t.Fatalf("buildItemFor(warrior, 16963) = %+v, ok=%v; want item_level 76, slot head", it, ok)
	}
	if _, ok := s.buildItemFor("warrior", 999999999); ok {
		t.Fatal("buildItemFor found an item id that does not exist, want false")
	}
}

// TestFallbackMatchQualifiesARealUpgradeAndRefusesADowngrade exercises fallbackMatch
// end to end against real per-class item data: a warrior with nothing equipped in head
// qualifies at the drop's own full item level; a warrior already wearing something with a
// higher item level than the drop does not.
func TestFallbackMatchQualifiesARealUpgradeAndRefusesADowngrade(t *testing.T) {
	s := realTreesStore(t)
	helmOfWrath := lootItem{ItemID: 16963, Slot: "head"}

	empty := RosterRow{className: "warrior", gear: map[string]int{}}
	delta, ok := s.fallbackMatch(empty, helmOfWrath)
	if !ok || delta != 76 {
		t.Fatalf("fallbackMatch (empty slot) = %d/%v, want 76/true", delta, ok)
	}

	// Red Winter Hat (21524, hunter.json's own leveling alt) is item level 1 - any real
	// head item outranks it, so a warrior "wearing" it (id reused purely for its known
	// low item level; buildItemFor is called with classSlug "warrior" below, and the id
	// alone determines item_level regardless of which class file first listed it, since
	// client item levels do not vary by class) still qualifies for the real upgrade.
	// To test a refusal, wear the drop item itself: no item outranks itself.
	wearingTheDrop := RosterRow{className: "warrior", gear: map[string]int{"head": 16963}}
	if _, ok := s.fallbackMatch(wearingTheDrop, helmOfWrath); ok {
		t.Fatal("fallbackMatch qualified a character already wearing the exact drop, want false")
	}

	noClassAccess := RosterRow{className: "not-a-real-class", gear: map[string]int{}}
	if _, ok := s.fallbackMatch(noClassAccess, helmOfWrath); ok {
		t.Fatal("fallbackMatch qualified an unresolvable class, want false")
	}
}

// TestLootCandidatesFallsBackToTier1ForARealOnyxiaDrop is the end-to-end check the
// follow-up fix asked for: lootCandidates, given a verified roster character with no BiS
// band match for Helm of Wrath (true for every character, since no BiS file in this repo
// names a raid-tier item - see CONTROL_CENTRE.md), still finds them as a tier-1 fallback
// candidate, with gain_dps nil, not_sim_checked true, and a real ilvl_delta.
func TestLootCandidatesFallsBackToTier1ForARealOnyxiaDrop(t *testing.T) {
	s := realTreesStore(t)
	roster := []RosterRow{
		{
			CharacterKey: "us/hardcore/fallbackwarrior", Name: "Fallbackwarrior", Verified: true,
			className: "warrior", specStr: "Fury", faction: "alliance", gear: map[string]int{},
		},
	}
	helmOfWrath := lootItem{ItemID: 16963, Slot: "head"}
	cands := s.lootCandidates(roster, helmOfWrath)
	if len(cands) != 1 {
		t.Fatalf("candidates = %+v, want exactly 1", cands)
	}
	c := cands[0]
	if c.GainDps != nil {
		t.Errorf("gainDps = %v, want nil", *c.GainDps)
	}
	if !c.NotSimChecked {
		t.Error("notSimChecked = false, want true")
	}
	if c.IlvlDelta == nil || *c.IlvlDelta != 76 {
		t.Errorf("ilvlDelta = %v, want 76", c.IlvlDelta)
	}
	if c.tier != 1 {
		t.Errorf("tier = %d, want 1", c.tier)
	}
}

// TestWornNameNamesTheItemInTheDropsSlot checks the "Replaces ..." name: the item the row wears
// in the slot, empty for an empty slot or one the class table cannot name.
func TestWornNameNamesTheItemInTheDropsSlot(t *testing.T) {
	s := realTreesStore(t)
	wearing := RosterRow{className: "warrior", gear: map[string]int{"head": 21524}}
	if got := s.wornName(wearing, "head"); got != "Red Winter Hat" {
		t.Errorf("wornName(head) = %q, want Red Winter Hat", got)
	}
	if got := s.wornName(wearing, "back"); got != "" {
		t.Errorf("wornName(empty slot) = %q, want empty", got)
	}
	if got := s.wornName(RosterRow{className: "warrior", gear: map[string]int{"head": 1}}, "head"); got != "" {
		t.Errorf("wornName(unknown item) = %q, want empty", got)
	}
}

// TestAwardCountsAndTheCandidateLine checks the candidate's "Replaces ..." name and its tier count:
// the count is the number of the guild's award rows naming that character, and the name is the item
// worn in the drop's slot.
func TestAwardCountsAndTheCandidateLine(t *testing.T) {
	counts := awardCounts(map[int]*AwardedTo{
		16963: {CharacterKey: "us/hardcore/a"},
		16939: {CharacterKey: "us/hardcore/a"},
		16947: {CharacterKey: "us/hardcore/b"},
	})
	if counts["us/hardcore/a"] != 2 || counts["us/hardcore/b"] != 1 || counts["us/hardcore/c"] != 0 {
		t.Fatalf("awardCounts = %v, want a:2 b:1 c:0", counts)
	}
	if got := awardCounts(nil); len(got) != 0 {
		t.Fatalf("awardCounts(nil) = %v, want empty", got)
	}

	s := realTreesStore(t)
	roster := []RosterRow{{
		CharacterKey: "us/hardcore/wearer", Name: "Wearer", Verified: true,
		className: "warrior", specStr: "Fury", faction: "alliance", gear: map[string]int{"head": 21524},
	}}
	cands := s.lootCandidates(roster, lootItem{ItemID: 16963, Slot: "head"})
	if len(cands) != 1 || cands[0].Replaces != "Red Winter Hat" {
		t.Fatalf("candidates = %+v, want one replacing Red Winter Hat", cands)
	}
}

// TestAwardLineNamesTheOfficersCharacter checks the officer on an award line is the account's
// main character in this guild (the highest item level of its characters here), "an officer" when
// the account has none, and "an officer" - never a pseudonym - once the account asks to be
// anonymous. A BattleTag never reaches the line.
func TestAwardLineNamesTheOfficersCharacter(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	ctx := context.Background()
	store := &Store{Pool: h.pool, Accounts: &auth.Store{Pool: h.pool}}

	addChar := func(uid int64, key, name string, ilvl int) {
		t.Helper()
		seedCharacter(t, h.pool, gid, uid, key, "officer", true)
		if _, err := h.pool.Exec(ctx, `
			insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
			values ($1, $2, 'us', 'hardcore', $3, 'FS1:x', now(), now())`, key, uid, name); err != nil {
			t.Fatal(err)
		}
		if _, err := h.pool.Exec(ctx, `
			insert into fight_metrics (report_id, fight_index, player_key, class, spec, role, ilvl, faction, fought_at)
			values ($1, 0, $2, 'hunter', 'Marksmanship', 'dps', $3, 'alliance', now())`,
			"award-line-"+key, key, ilvl); err != nil {
			t.Fatal(err)
		}
	}
	award := func(uid int64, itemID int, to string) {
		t.Helper()
		if _, err := h.pool.Exec(ctx, `
			insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by)
			values ($1, 1084, $2, $3, $4)`, gid, itemID, to, uid); err != nil {
			t.Fatal(err)
		}
	}

	// An officer with two characters here: the higher item level is the main.
	main := seedUser(t, h.pool, "award-line-main@example.com")
	addChar(main, "us/hardcore/awardalt", "Altwing", 52)
	addChar(main, "us/hardcore/awardmain", "Bowmain", 66)
	if _, err := h.pool.Exec(ctx, `update users set battletag = 'Officer#1234' where id = $1`, main); err != nil {
		t.Fatal(err)
	}
	// An officer whose account has no character in this guild at all.
	stranger := seedUser(t, h.pool, "award-line-none@example.com")
	// An officer who asked to be shown by pseudonym.
	hidden := seedUser(t, h.pool, "award-line-anon@example.com")
	addChar(hidden, "us/hardcore/awardhidden", "Shadowed", 60)
	if _, err := h.pool.Exec(ctx, `update users set anonymize = true, battletag = 'Hidden#9999' where id = $1`, hidden); err != nil {
		t.Fatal(err)
	}
	winner := seedUser(t, h.pool, "award-line-winner@example.com")
	addChar(winner, "us/hardcore/awardwinner", "Rootbound", 58)

	award(main, 16963, "us/hardcore/awardwinner")
	award(stranger, 16939, "us/hardcore/awardwinner")
	award(hidden, 16947, "us/hardcore/awardwinner")

	got, err := store.lootAwards(ctx, gid, 1084)
	if err != nil {
		t.Fatal(err)
	}
	for item, want := range map[int]string{16963: "Bowmain", 16939: "an officer", 16947: "an officer"} {
		if got[item] == nil || got[item].ByName != want {
			t.Fatalf("award of %d: by_name = %+v, want %q", item, got[item], want)
		}
	}
	for item, a := range got {
		if strings.Contains(a.ByName, "#") || strings.HasPrefix(a.ByName, "user-") {
			t.Fatalf("award of %d names an account (%q), want a character or \"an officer\"", item, a.ByName)
		}
	}
}

// Plan rows the gate tests seed: a scoped grant, and a whole Stripe guild plan.
const (
	scopedGrantSQL = `insert into entitlements (guild_id, plan, source, status, current_period_end, features)
		values ($1,'guild','grant','active', now()+interval '30 days','{officer_views}')`
	stripePlanSQL = `insert into entitlements (guild_id, plan, source, status, current_period_end, stripe_subscription_id)
		values ($1,'guild','stripe','active', now()+interval '30 days', $2)`
)

// grantOfficerTools wires the real entitlements store and gives gid a scoped grant.
func grantOfficerTools(t *testing.T, h *httpHarness, gid int64) {
	t.Helper()
	h.svc.Entitlements = &entitlements.Store{Pool: h.pool}
	if _, err := h.pool.Exec(t.Context(), scopedGrantSQL, gid); err != nil {
		t.Fatal(err)
	}
}

// TestLootAwardAndUndoNeedTheGuildPlan checks the plan gate: members learn nothing about the
// plan (forbidden first), an officer without it gets guild_plan_required, and either a scoped
// grant or a whole Stripe plan lets Award and Undo through.
func TestLootAwardAndUndoNeedTheGuildPlan(t *testing.T) {
	const planMessage = "Award and Undo are part of the guild plan."
	tests := []struct {
		name     string
		planSQL  string
		nilStore bool
		rank     string
		wantCode string // "" means allowed
	}{
		{"officer with a scoped grant", scopedGrantSQL, false, "officer", ""},
		{"officer with a Stripe plan", stripePlanSQL, false, "officer", ""},
		{"officer with no plan", "", false, "officer", "guild_plan_required"},
		{"officer with no entitlements wired", "", true, "officer", "guild_plan_required"},
		{"member with a scoped grant", scopedGrantSQL, false, "member", "forbidden"},
		{"member with no plan", "", false, "member", "forbidden"},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := newHTTPHarness(t)
			if !tc.nilStore {
				h.svc.Entitlements = &entitlements.Store{Pool: h.pool}
			}
			gid := seedGuild(t, h.pool, fmt.Sprintf("PlanGate%d", i))
			target := fmt.Sprintf("us/hardcore/plangatetarget%d", i)
			seedHunterCharacter(t, h, gid, fmt.Sprintf("plan-gate-target-%d@example.com", i), target, 21524)
			actorUID := seedUser(t, h.pool, fmt.Sprintf("plan-gate-actor-%d@example.com", i))
			seedCharacter(t, h.pool, gid, actorUID, fmt.Sprintf("us/hardcore/plangateactor%d", i), tc.rank, true)
			recomputeMembership(t, h.pool, gid, actorUID)
			if tc.planSQL != "" {
				args := []any{gid}
				if tc.planSQL == stripePlanSQL {
					args = append(args, fmt.Sprintf("sub_plan_gate_%d", gid))
				}
				if _, err := h.pool.Exec(t.Context(), tc.planSQL, args...); err != nil {
					t.Fatal(err)
				}
			}
			h.actor = auth.Actor{UserID: actorUID, Role: "user", Method: "session"}

			awardPath := fmt.Sprintf("/v1/guilds/%d/loot/awards", gid)
			body := fmt.Sprintf(`{"encounter_id":1084,"item_id":16963,"character_key":%q}`, target)
			res := h.do(http.MethodPost, awardPath, body)
			if tc.wantCode == "" {
				var created map[string]any
				h.data(res, &created)
				awardID := int64(created["award_id"].(float64))
				res = h.do(http.MethodDelete, fmt.Sprintf("%s/%d", awardPath, awardID), "")
				res.Body.Close()
				if res.StatusCode != http.StatusOK {
					t.Fatalf("undo = %d, want 200", res.StatusCode)
				}
				return
			}
			assertRefusal(t, res, tc.wantCode, planMessage)

			// Undo is gated the same way; the award row only has to exist for the path to resolve.
			var awardID int64
			if err := h.pool.QueryRow(t.Context(),
				`insert into loot_awards (guild_id, encounter_id, item_id, character_key, awarded_by)
				 values ($1,1084,16963,$2,$3) returning id`, gid, target, actorUID).Scan(&awardID); err != nil {
				t.Fatal(err)
			}
			res = h.do(http.MethodDelete, fmt.Sprintf("%s/%d", awardPath, awardID), "")
			assertRefusal(t, res, tc.wantCode, planMessage)
		})
	}
}

// assertRefusal checks a 403 with the given error code, and the plan message when it is the
// plan refusal.
func assertRefusal(t *testing.T, res *http.Response, wantCode, planMessage string) {
	t.Helper()
	defer res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("status = %d, want 403", res.StatusCode)
	}
	var env struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if env.Error.Code != wantCode {
		t.Fatalf("code = %q, want %q", env.Error.Code, wantCode)
	}
	if wantCode == "guild_plan_required" && env.Error.Message != planMessage {
		t.Fatalf("message = %q, want %q", env.Error.Message, planMessage)
	}
}
