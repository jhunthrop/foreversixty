// api/internal/guilds/loot_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
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

	if len(view.Items) != 22 {
		t.Fatalf("items = %d, want 22", len(view.Items))
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
}

// TestLootAwardRejectsAnUnknownItem checks CreateLootAward's own validation.
func TestLootAwardRejectsAnUnknownItem(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
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
	cands := s.lootCandidates(roster, helmOfWrath, map[string]int{}, 8)
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
