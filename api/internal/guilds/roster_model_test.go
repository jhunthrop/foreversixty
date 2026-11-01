// api/internal/guilds/roster_model_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// TestMergeRosterTable pins the roster model: one entry per character the snapshot names and
// per account character, merged by name, each with a state, and the counts the page's count
// line reads. It is pure, so every shape of guild is one row of the table.
func TestMergeRosterTable(t *testing.T) {
	approved := func(name, key, rank string) RosterRow {
		return RosterRow{Name: name, CharacterKey: key, Rank: rank, Verified: true}
	}
	waiting := func(name, key, rank string) RosterRow {
		return RosterRow{Name: name, CharacterKey: key, Rank: rank}
	}
	line := func(name string) snapshotMember {
		return snapshotMember{Name: name, Level: 60, Spec: "Protection", SeenAt: "2026-10-20 18:00"}
	}

	cases := []struct {
		name     string
		rows     []RosterRow
		invited  []invitedMember
		snapshot []snapshotMember
		want     []RosterEntry
		counts   RosterCounts
	}{
		{
			name: "no snapshot lists the account rows alone",
			rows: []RosterRow{approved("Bow", "us/pvp/bow", "leader"), waiting("Ashenveil", "us/pvp/ashenveil", "officer")},
			want: []RosterEntry{
				{Name: "Bow", CharacterKey: "us/pvp/bow", State: EntryApproved, Rank: "leader"},
				{Name: "Ashenveil", CharacterKey: "us/pvp/ashenveil", State: EntryWaiting, Rank: "officer"},
			},
			counts: RosterCounts{OnRoster: 2, Waiting: 1},
		},
		{
			name:     "a snapshot name with no account is in game without one, by name, rank member",
			rows:     []RosterRow{approved("Bow", "us/pvp/bow", "leader")},
			snapshot: []snapshotMember{line("Bow"), line("Strayhound"), line("Oldmarrow-Realm")},
			want: []RosterEntry{
				{Name: "Bow", CharacterKey: "us/pvp/bow", State: EntryApproved, Rank: "leader"},
				{Name: "Oldmarrow", State: EntryNoAccount, Rank: "member", Level: i(60), Spec: strp("Protection"), SeenAt: "2026-10-20 18:00"},
				{Name: "Strayhound", State: EntryNoAccount, Rank: "member", Level: i(60), Spec: strp("Protection"), SeenAt: "2026-10-20 18:00"},
			},
			counts: RosterCounts{OnRoster: 3, InGameNoAccount: 2},
		},
		{
			name:     "names match in any case and a repeated snapshot line counts once",
			rows:     []RosterRow{approved("Bow Jackzon", "us/pvp/bow-jackzon", "leader")},
			snapshot: []snapshotMember{line("bow jackzon-Realm"), line("BOW JACKZON"), line("Strayhound"), line("strayhound")},
			want: []RosterEntry{
				{Name: "Bow Jackzon", CharacterKey: "us/pvp/bow-jackzon", State: EntryApproved, Rank: "leader"},
				{Name: "Strayhound", State: EntryNoAccount, Rank: "member", Level: i(60), Spec: strp("Protection"), SeenAt: "2026-10-20 18:00"},
			},
			counts: RosterCounts{OnRoster: 2, InGameNoAccount: 1},
		},
		{
			name:    "an invited member the snapshot does not name is not seen in game yet",
			rows:    []RosterRow{approved("Bow", "us/pvp/bow", "leader")},
			invited: []invitedMember{{UserID: 9, Name: "Driftwood", Class: strp("rogue"), Gear: true}},
			want: []RosterEntry{
				{Name: "Bow", CharacterKey: "us/pvp/bow", State: EntryApproved, Rank: "leader"},
				{Name: "Driftwood", State: EntryNotSeen, Rank: "member", Class: strp("rogue")},
			},
			counts: RosterCounts{OnRoster: 2, NotSeen: 1},
		},
		{
			name:     "an invited member the snapshot names is approved and seen, not counted twice",
			invited:  []invitedMember{{UserID: 9, Name: "Driftwood", Class: strp("rogue"), Gear: true}},
			snapshot: []snapshotMember{line("Driftwood")},
			want: []RosterEntry{
				{Name: "Driftwood", State: EntryApproved, Rank: "member", Class: strp("rogue"), Level: i(60), Spec: strp("Protection"), SeenAt: "2026-10-20 18:00"},
			},
			counts: RosterCounts{OnRoster: 1},
		},
		{
			name:     "an invited member who withholds gear shows no level or spec from the snapshot",
			invited:  []invitedMember{{UserID: 9, Name: "Driftwood"}},
			snapshot: []snapshotMember{line("Driftwood")},
			want: []RosterEntry{
				{Name: "Driftwood", State: EntryApproved, Rank: "member", SeenAt: "2026-10-20 18:00"},
			},
			counts: RosterCounts{OnRoster: 1},
		},
		{
			name:    "an invited member who is also an account character is listed once",
			rows:    []RosterRow{approved("Driftwood", "us/pvp/driftwood", "member")},
			invited: []invitedMember{{UserID: 9, Name: "driftwood", Gear: true}},
			want: []RosterEntry{
				{Name: "Driftwood", CharacterKey: "us/pvp/driftwood", State: EntryApproved, Rank: "member"},
			},
			counts: RosterCounts{OnRoster: 1},
		},
		{
			name: "a row the snapshot carries brings the snapshot's line along",
			rows: []RosterRow{{
				Name: "Bow", CharacterKey: "us/pvp/bow", Rank: "leader", Verified: true,
				Addon: &AddonMemberView{Level: i(60), Spec: strp("hunter-beast-mastery"), SeenAt: "2026-10-20 17:00"},
			}},
			snapshot: []snapshotMember{line("Bow")},
			want: []RosterEntry{
				{Name: "Bow", CharacterKey: "us/pvp/bow", State: EntryApproved, Rank: "leader", Level: i(60), Spec: strp("hunter-beast-mastery"), SeenAt: "2026-10-20 17:00"},
			},
			counts: RosterCounts{OnRoster: 1},
		},
		{
			name: "the guild from the walk: seven on the roster, two in game without an account, one waiting, one not seen",
			rows: []RosterRow{
				approved("Bow", "us/pvp/bow", "leader"),
				approved("Ashenveil", "us/pvp/ashenveil", "officer"),
				approved("Quillfeather", "us/pvp/quillfeather", "member"),
				waiting("Rootbound", "us/pvp/rootbound", "member"),
			},
			invited:  []invitedMember{{UserID: 4, Name: "Driftwood"}},
			snapshot: []snapshotMember{line("Bow"), line("Ashenveil"), line("Strayhound"), line("Oldmarrow")},
			counts:   RosterCounts{OnRoster: 7, InGameNoAccount: 2, Waiting: 1, NotSeen: 1},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, counts := mergeRoster(tc.rows, tc.invited, tc.snapshot)
			if counts != tc.counts {
				t.Errorf("counts = %+v, want %+v", counts, tc.counts)
			}
			if tc.want != nil && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("entries = %+v, want %+v", got, tc.want)
			}
			if len(got) != counts.OnRoster {
				t.Errorf("len(entries) = %d, on_roster = %d", len(got), counts.OnRoster)
			}
		})
	}
}

// TestCapabilitiesTable is the one table of who may do what in a guild: an officer is not a
// second guild master, the claimant is, and a moderator acts from outside.
func TestCapabilitiesTable(t *testing.T) {
	const me, other int64 = 7, 8
	claimedBy := func(id int64) Guild { return Guild{ClaimedBy: &id} }
	cases := []struct {
		name            string
		g               Guild
		moderator       bool
		verifiedOfficer bool
		want            []string
	}{
		{"a plain member", claimedBy(other), false, false, []string{}},
		{"an unverified viewer", Guild{}, false, false, []string{}},
		{"an officer under another's claim", claimedBy(other), false, true,
			[]string{CapApprove, CapRemove, CapSettings, CapAward}},
		{"an officer in an unclaimed guild", Guild{}, false, true,
			[]string{CapApprove, CapRemove, CapSettings, CapAward}},
		{"the guild master holding the claim", claimedBy(me), false, true,
			[]string{CapApprove, CapRemove, CapSettings, CapAward, CapBilling, CapRelease}},
		{"a claimant who lost their rank may still release", claimedBy(me), false, false, []string{CapRelease}},
		{"a moderator", claimedBy(other), true, false, []string{CapRemove, CapSettings, CapRelease}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Capabilities(tc.g, me, tc.moderator, tc.verifiedOfficer); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Capabilities = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestMayRemoveRowTable pins mayRemoveRow including the officer's own row.
func TestMayRemoveRowTable(t *testing.T) {
	const me, owner int64 = 7, 8
	claimedBy := func(id int64) Guild { return Guild{ClaimedBy: &id} }
	cases := []struct {
		name            string
		g               Guild
		ownerID         int64
		rank            string
		moderator       bool
		verifiedOfficer bool
		want            bool
	}{
		{"an officer removes a member", Guild{}, owner, "member", false, true, true},
		{"a member does not remove a member", Guild{}, owner, "member", false, false, false},
		{"an officer does not remove an officer", claimedBy(owner), owner, "officer", false, true, false},
		{"the claimant removes an officer", claimedBy(me), owner, "officer", false, true, true},
		{"nobody but a moderator removes a leader", claimedBy(me), owner, "leader", false, true, false},
		{"a moderator removes a leader", Guild{}, owner, "leader", true, false, true},
		{"a member removes their own row", Guild{}, me, "member", false, false, true},
		{"an officer has no Remove on her own row", claimedBy(owner), me, "officer", false, true, false},
		{"the claimant has no Remove on their own officer row", claimedBy(me), me, "officer", false, true, false},
		{"the claimant has no Remove on their own row, whatever its rank", claimedBy(me), me, "member", false, true, false},
		{"a leader has no Remove on their own row", Guild{}, me, "leader", false, true, false},
		{"a leader who holds the claim has none either", claimedBy(me), me, "leader", false, true, false},
		{"a moderator removes their own officer row", claimedBy(owner), me, "officer", true, false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := mayRemoveRow(tc.g, me, tc.ownerID, tc.rank, tc.moderator, tc.verifiedOfficer); got != tc.want {
				t.Errorf("mayRemoveRow = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestStoreCapabilitiesReadsTheViewersStanding runs Capabilities against real rows: the
// capability list /v1/me sends for an officer, the claimant, a member and a stranger.
func TestStoreCapabilitiesReadsTheViewersStanding(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	master := seedUser(t, h.pool, "cap-master@example.com")
	seedCharacter(t, h.pool, gid, master, "us/hardcore/capmaster", "leader", true)
	recomputeMembership(t, h.pool, gid, master)
	officer := seedUser(t, h.pool, "cap-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/capofficer", "officer", true)
	recomputeMembership(t, h.pool, gid, officer)
	member := seedUser(t, h.pool, "cap-member@example.com")
	seedCharacter(t, h.pool, gid, member, "us/hardcore/capmember", "member", true)
	recomputeMembership(t, h.pool, gid, member)
	waitingOfficer := seedUser(t, h.pool, "cap-waiting@example.com")
	seedCharacter(t, h.pool, gid, waitingOfficer, "us/hardcore/capwaiting", "officer", false)
	recomputeMembership(t, h.pool, gid, waitingOfficer)
	stranger := seedUser(t, h.pool, "cap-stranger@example.com")
	if _, err := h.pool.Exec(ctx, `update guilds set claimed_by = $1 where id = $2`, master, gid); err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name string
		uid  int64
		want []string
	}{
		{"the guild master", master, []string{CapApprove, CapRemove, CapSettings, CapAward, CapBilling, CapRelease}},
		{"an officer", officer, []string{CapApprove, CapRemove, CapSettings, CapAward}},
		{"a member", member, []string{}},
		{"an officer nobody has approved", waitingOfficer, []string{}},
		{"a stranger", stranger, []string{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := h.store.Capabilities(ctx, gid, tc.uid, false)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("can = %v, want %v", got, tc.want)
			}
		})
	}
	if _, err := h.store.Capabilities(ctx, gid+999, master, false); err == nil {
		t.Error("an unknown guild: want an error")
	}
}

// putSnapshot stores a roster snapshot for the guild as the addon's outbox would.
func putSnapshot(t *testing.T, h *httpHarness, gid int64, senderKey string, members string) {
	t.Helper()
	if _, err := h.pool.Exec(context.Background(),
		`insert into guild_roster_snapshots (guild_id, character_key, members, received_at)
		 values ($1, $2, $3::jsonb, now())`, gid, senderKey, members); err != nil {
		t.Fatal(err)
	}
}

// TestHomeMembersListsEveryoneTheAddonSaw walks the guild from the review: verified and
// waiting accounts, two names in game with no account, and a member who joined by invite and
// is not in the snapshot; then the same invited member once the snapshot names them.
func TestHomeMembersListsEveryoneTheAddonSaw(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	gid := seedGuild(t, h.pool, "Forever")

	master := seedUser(t, h.pool, "mem-master@example.com")
	seedCharacter(t, h.pool, gid, master, "us/hardcore/memmaster", "leader", true)
	if _, err := h.pool.Exec(ctx, `insert into characters (key, user_id, region, ruleset, name) values ('us/hardcore/memmaster', $1, 'us', 'hardcore', 'Memmaster') on conflict do nothing`, master); err != nil {
		t.Fatal(err)
	}
	seedExport(t, h.pool, master, "us/hardcore/memmaster", "us", "hardcore", "Memmaster")
	recomputeMembership(t, h.pool, gid, master)
	pending := seedUser(t, h.pool, "mem-pending@example.com")
	seedCharacter(t, h.pool, gid, pending, "us/hardcore/mempending", "officer", false)
	seedExport(t, h.pool, pending, "us/hardcore/mempending", "us", "hardcore", "Mempending")
	recomputeMembership(t, h.pool, gid, pending)

	// An invite-joined account: a verified synthetic row, and a character of its own that is
	// in no guild.
	invitee := seedUser(t, h.pool, "mem-invitee@example.com")
	seedCharacter(t, h.pool, gid, invitee, syntheticKey(invitee), "member", true)
	recomputeMembership(t, h.pool, gid, invitee)
	if _, err := h.pool.Exec(ctx, `insert into characters (key, user_id, region, ruleset, name, class)
		values ('us/hardcore/driftwood', $1, 'us', 'hardcore', 'Driftwood', 'rogue')`, invitee); err != nil {
		t.Fatal(err)
	}

	putSnapshot(t, h, gid, "us/hardcore/memmaster", `[
		{"name":"Memmaster","level":60,"spec":"hunter-beast-mastery","gear_gap":0,"seen_at":"2026-10-20 18:00"},
		{"name":"Mempending-Realm","level":60,"spec":"paladin-holy","gear_gap":3,"seen_at":"2026-10-20 18:05"},
		{"name":"Strayhound","level":60,"spec":"warrior-protection","gear_gap":0,"seen_at":"2026-10-20 17:00"},
		{"name":"Oldmarrow","level":58,"spec":"warrior-arms","gear_gap":0,"seen_at":"2026-10-19 21:00"}]`)

	read := func(uid int64) HomeView {
		h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
		var view HomeView
		h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""), &view)
		return view
	}
	states := func(v HomeView) map[string]string {
		out := map[string]string{}
		for _, e := range v.Members {
			out[e.Name] = e.State
		}
		return out
	}

	view := read(master)
	want := map[string]string{
		"Memmaster": EntryApproved, "Mempending": EntryWaiting, "Driftwood": EntryNotSeen,
		"Strayhound": EntryNoAccount, "Oldmarrow": EntryNoAccount,
	}
	if got := states(view); !reflect.DeepEqual(got, want) {
		t.Fatalf("member states = %v, want %v", got, want)
	}
	if view.Counts != (RosterCounts{OnRoster: 5, InGameNoAccount: 2, Waiting: 1, NotSeen: 1}) {
		t.Errorf("counts = %+v", view.Counts)
	}
	for _, e := range view.Members {
		if e.Name == "Mempending" && e.Rank != "officer" {
			t.Errorf("the waiting officer's rank = %q, want officer (the word the addon sent)", e.Rank)
		}
		if e.Name == "Driftwood" && (e.Class == nil || *e.Class != "rogue" || e.Rank != "member") {
			t.Errorf("the invited member = %+v, want a member rogue", e)
		}
	}
	// Roster and pending stay the account rows.
	if len(view.Roster) != 2 || len(view.Pending) != 1 {
		t.Errorf("roster/pending = %d/%d, want 2/1", len(view.Roster), len(view.Pending))
	}

	// The snapshot now names the invited member's character: seen, approved, no longer "not seen".
	if _, err := h.pool.Exec(ctx, `update guild_roster_snapshots set members = members ||
		'[{"name":"Driftwood","level":57,"spec":"rogue-combat","gear_gap":4,"seen_at":"2026-10-20 18:30"}]'::jsonb
		where guild_id = $1`, gid); err != nil {
		t.Fatal(err)
	}
	view = read(master)
	if got := states(view)["Driftwood"]; got != EntryApproved {
		t.Errorf("Driftwood once the snapshot names them = %q, want approved", got)
	}
	if view.Counts.NotSeen != 0 || view.Counts.OnRoster != 5 {
		t.Errorf("counts after the sighting = %+v, want 5 on the roster and none unseen", view.Counts)
	}

	// An invited account's own character appearing in the guild as a normal row (its player
	// installed the addon) replaces the invited line rather than doubling it.
	seedCharacter(t, h.pool, gid, invitee, "us/hardcore/driftwood", "member", true)
	seedExport(t, h.pool, invitee, "us/hardcore/driftwood", "us", "hardcore", "Driftwood")
	view = read(master)
	n := 0
	for _, e := range view.Members {
		if e.Name == "Driftwood" {
			n++
		}
	}
	if n != 1 {
		t.Errorf("Driftwood listed %d times, want once", n)
	}
}

// TestHomeMembersWithoutASnapshot is a guild the addon never sent a roster for: its account
// rows list, and nobody is "in game without an account".
func TestHomeMembersWithoutASnapshot(t *testing.T) {
	h := newHTTPHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "nosnap@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/nosnap", "member", true)
	seedExport(t, h.pool, uid, "us/hardcore/nosnap", "us", "hardcore", "Nosnap")
	recomputeMembership(t, h.pool, gid, uid)
	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	var view HomeView
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), ""), &view)
	if view.AddonSnapshot != nil || len(view.Members) != 1 || view.Counts != (RosterCounts{OnRoster: 1}) {
		t.Fatalf("members/counts/snapshot = %+v/%+v/%+v", view.Members, view.Counts, view.AddonSnapshot)
	}
}

// TestLootCountsAnInvitedMemberAsACandidate is the loot council's side of the new roster: an
// account that joined by invite has a character with gear on the site and no guild_characters
// row of the guild's, yet it is ranked, can be awarded (under the account's own row), and is
// named on the award.
func TestLootCountsAnInvitedMemberAsACandidate(t *testing.T) {
	h := newHTTPHarness(t)
	ctx := context.Background()
	// The class item tables the fallback tier reads.
	real := realTreesStore(t)
	h.store.Trees, h.store.DataDir = real.Trees, real.DataDir
	gid := seedGuild(t, h.pool, "Forever")
	grantOfficerTools(t, h, gid)
	seedHunterCharacter(t, h, gid, "inv-loot-member@example.com", "us/hardcore/invlootmember", 21524)

	invitee := seedUser(t, h.pool, "inv-loot-invitee@example.com")
	synthetic := syntheticKey(invitee)
	seedCharacter(t, h.pool, gid, invitee, synthetic, "member", true)
	recomputeMembership(t, h.pool, gid, invitee)
	if _, err := h.pool.Exec(ctx, `insert into characters (key, user_id, region, ruleset, name, class)
		values ('us/hardcore/invlootdrift', $1, 'us', 'hardcore', 'Invlootdrift', 'hunter')`, invitee); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(ctx, `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ('us/hardcore/invlootdrift', $1, 'us', 'hardcore', 'Invlootdrift',
		        'FS1:1.60.1.70009:hunter:dwarf:555555555/33/0:head=21524|level=60', now(), now())`, invitee); err != nil {
		t.Fatal(err)
	}

	officer := seedUser(t, h.pool, "inv-loot-officer@example.com")
	seedCharacter(t, h.pool, gid, officer, "us/hardcore/invlootofficer", "officer", true)
	recomputeMembership(t, h.pool, gid, officer)
	h.actor = auth.Actor{UserID: officer, Role: "user", Method: "session"}

	read := func() LootView {
		var view LootView
		h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/loot", gid), ""), &view)
		return view
	}
	candidateOf := func(view LootView, itemID int, name string) *LootCandidate {
		for _, it := range view.Items {
			if it.ItemID != itemID {
				continue
			}
			for i := range it.Candidates {
				if it.Candidates[i].Name == name {
					return &it.Candidates[i]
				}
			}
		}
		return nil
	}

	view := read()
	// Helm of Wrath is a head drop a hunter cannot use; Dragonstalker's Helm is theirs.
	cand := candidateOf(view, 16939, "Invlootdrift")
	if cand == nil {
		t.Fatalf("the invited member is not a candidate for Dragonstalker's Helm: %+v", view.Items)
	}
	if cand.CharacterKey != synthetic {
		t.Errorf("candidate key = %q, want the account row %q (the row an award names)", cand.CharacterKey, synthetic)
	}

	body := fmt.Sprintf(`{"encounter_id":1084,"item_id":16939,"character_key":%q}`, synthetic)
	res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/loot/awards", gid), body)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("awarding to an invited member = %d, want 200", res.StatusCode)
	}
	view = read()
	for _, it := range view.Items {
		if it.ItemID == 16939 {
			if it.AwardedTo == nil || it.AwardedTo.Name != "Invlootdrift" || it.AwardedTo.CharacterKey != synthetic {
				t.Errorf("awarded_to = %+v, want Invlootdrift under %s", it.AwardedTo, synthetic)
			}
		}
	}
}

// TestLootListsEachDropOnce: the table holds two items named Head of Onyxia (the Alliance and
// the Horde quest item, ids 18422 and 18423) - the double row was data, not a rendering fault -
// and the council lists the name once with Copies 2.
func TestLootListsEachDropOnce(t *testing.T) {
	seen := map[string]int{}
	for _, it := range onyxiaLoot {
		seen[it.Name]++
	}
	dupes := map[string]int{}
	for name, n := range seen {
		if n > 1 {
			dupes[name] = n
		}
	}
	if !reflect.DeepEqual(dupes, map[string]int{"Head of Onyxia": 2}) {
		t.Fatalf("names the table repeats = %v, want only Head of Onyxia twice", dupes)
	}
	items := []LootItemView{{Name: "A"}, {Name: "Head of Onyxia"}}
	if i := indexOfLootName(items, "Head of Onyxia"); i != 1 {
		t.Errorf("indexOfLootName = %d, want 1", i)
	}
	if i := indexOfLootName(items, "Nope"); i != -1 {
		t.Errorf("indexOfLootName(missing) = %d, want -1", i)
	}
}
