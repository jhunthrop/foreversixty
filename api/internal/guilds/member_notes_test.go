package guilds

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

// notesGuild seeds a guild master, an officer and a member, each verified with an export.
func notesGuild(t *testing.T, h *httpHarness) (gid, leader, officer, member int64) {
	t.Helper()
	gid = seedGuild(t, h.pool, "Notes")
	leader = seedUser(t, h.pool, "notes-leader@example.com")
	seedVerifiedMember(t, h, gid, leader, "notesleader", "leader")
	officer = seedUser(t, h.pool, "notes-officer@example.com")
	seedVerifiedMember(t, h, gid, officer, "notesofficer", "officer")
	member = seedUser(t, h.pool, "notes-member@example.com")
	seedVerifiedMember(t, h, gid, member, "notesmember", "member")
	return gid, leader, officer, member
}

const secretNote = "Ask about the Thursday conflict."

func TestOfficerNotesAreOfficerOnly(t *testing.T) {
	h := newHTTPHarness(t)
	gid, leader, officer, member := notesGuild(t, h)
	notes := fmt.Sprintf("/v1/guilds/%d/members/%s/notes", gid, accountKeyFor(member))

	h.as(member)
	if got := h.status(http.MethodGet, notes, ""); got != http.StatusForbidden {
		t.Fatalf("member reads notes = %d, want 403", got)
	}
	if got := h.status(http.MethodPost, notes, `{"body":"x"}`); got != http.StatusForbidden {
		t.Fatalf("member writes a note = %d, want 403", got)
	}
	h.as(officer)
	if got := h.status(http.MethodPost, notes, `{"body":"   "}`); got != http.StatusBadRequest {
		t.Fatalf("empty note = %d, want 400", got)
	}
	stranger := seedUser(t, h.pool, "notes-stranger@example.com")
	if got := h.status(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/members/%s/notes", gid, accountKeyFor(stranger)), `{"body":"x"}`); got != http.StatusNotFound {
		t.Fatalf("note on a non-member = %d, want 404", got)
	}
	var first, second MemberNote
	h.data(h.do(http.MethodPost, notes, `{"body":"Always brings flasks."}`), &first)
	h.as(leader)
	h.data(h.do(http.MethodPost, notes, fmt.Sprintf(`{"body":%q}`, secretNote)), &second)
	if second.Author == nil || second.Author.Name != "notesleader" || !second.MayDelete {
		t.Fatalf("leader's note = %+v", second)
	}

	h.as(officer)
	var list struct{ Rows []MemberNote }
	h.data(h.do(http.MethodGet, notes, ""), &list)
	if len(list.Rows) != 2 || list.Rows[0].ID != second.ID || list.Rows[0].MayDelete || !list.Rows[1].MayDelete {
		t.Fatalf("officer's list = %+v (newest first, only their own deletable)", list.Rows)
	}
	if got := h.status(http.MethodDelete, fmt.Sprintf("%s/%d", notes, second.ID), ""); got != http.StatusForbidden {
		t.Fatalf("officer deletes the leader's note = %d, want 403", got)
	}
	if got := h.status(http.MethodDelete, fmt.Sprintf("%s/%d", notes, first.ID), ""); got != http.StatusOK {
		t.Fatalf("officer deletes own = %d", got)
	}
	h.as(leader)
	if got := h.status(http.MethodPost, notes, `{"body":"Officer's note."}`); got != http.StatusCreated {
		t.Fatalf("leader adds = %d", got)
	}
	h.as(officer)
	var officerNote MemberNote
	h.data(h.do(http.MethodPost, notes, `{"body":"Mine."}`), &officerNote)
	h.as(leader)
	if got := h.status(http.MethodDelete, fmt.Sprintf("%s/%d", notes, officerNote.ID), ""); got != http.StatusOK {
		t.Fatalf("guild master deletes any = %d", got)
	}
}

// body reads a response's raw body.
func body(t *testing.T, res *http.Response) string {
	t.Helper()
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

func TestNoMemberReadCarriesOfficerNotes(t *testing.T) {
	h := newHTTPHarness(t)
	gid, leader, _, member := notesGuild(t, h)
	if _, err := h.store.AddMemberNote(context.Background(), gid, member, leader, "leader", secretNote); err != nil {
		t.Fatal(err)
	}
	h.as(member)
	for _, path := range []string{
		fmt.Sprintf("/v1/guilds/%d/home", gid),
		fmt.Sprintf("/v1/guilds/%d/readiness", gid),
		fmt.Sprintf("/v1/guilds/%d/members/me/profile", gid),
		fmt.Sprintf("/v1/guilds/%d/members/%s/profile", gid, accountKeyFor(member)),
	} {
		res := h.do(http.MethodGet, path, "")
		if res.StatusCode != http.StatusOK {
			t.Fatalf("%s = %d", path, res.StatusCode)
		}
		if raw := body(t, res); strings.Contains(raw, secretNote) || strings.Contains(raw, `"notes"`) {
			t.Fatalf("%s carries officer notes: %s", path, raw)
		}
	}
}

func TestMemberProfileIsOfficerOrSelf(t *testing.T) {
	h := newHTTPHarness(t)
	gid, leader, officer, member := notesGuild(t, h)
	if _, err := h.store.AddMemberNote(context.Background(), gid, member, leader, "leader", secretNote); err != nil {
		t.Fatal(err)
	}
	if _, err := h.pool.Exec(context.Background(), `insert into loot_awards (guild_id, encounter_id, item_id, character_key,
		ms_os, reason, votes_yes, votes_no, votes_officers)
		values ($1, 1084, 18813, 'us/hardcore/notesmember', 'ms', 'Best in slot', 3, 1, 4)`, gid); err != nil {
		t.Fatal(err)
	}
	profile := fmt.Sprintf("/v1/guilds/%d/members/%s/profile", gid, accountKeyFor(member))

	h.as(0)
	if got := h.status(http.MethodGet, profile, ""); got != http.StatusUnauthorized {
		t.Fatalf("signed out = %d", got)
	}
	other := seedUser(t, h.pool, "notes-other@example.com")
	seedVerifiedMember(t, h, gid, other, "notesother", "member")
	h.as(other)
	if got := h.status(http.MethodGet, profile, ""); got != http.StatusForbidden {
		t.Fatalf("another member = %d, want 403", got)
	}
	h.as(officer)
	if got := h.status(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/members/u:999999/profile", gid), ""); got != http.StatusNotFound {
		t.Fatalf("unknown member = %d, want 404", got)
	}
	var p MemberProfile
	h.data(h.do(http.MethodGet, profile, ""), &p)
	if p.Member == nil || p.Member.CharacterKey != "us/hardcore/notesmember" || p.Attendance == nil || p.Readiness == nil {
		t.Fatalf("profile = %+v", p)
	}
	if p.Loot.Count != 1 || p.Loot.Items[0].Name == nil || *p.Loot.Items[0].Name != "Ring of Binding" {
		t.Fatalf("loot = %+v", p.Loot)
	}
	// The award's decision record rides with it; the guild median counts every verified raider,
	// the ones with nothing as zero (one award among this guild's raiders: the median is 0).
	if r := p.Loot.Items[0].AwardRecord; r.MsOs == nil || *r.MsOs != "ms" || r.Reason == nil || *r.Reason != "Best in slot" ||
		r.VotesYes == nil || *r.VotesYes != 3 || r.VotesNo == nil || *r.VotesNo != 1 || r.VotesOfficers == nil || *r.VotesOfficers != 4 {
		t.Fatalf("award record = %+v", r)
	}
	if p.Upgrades == nil {
		t.Fatal("upgrades is null; a profile with a readiness row sends a list, empty without a gap")
	}
	if p.Loot.GuildMedian != 0 {
		t.Fatalf("guild median = %d, want 0", p.Loot.GuildMedian)
	}
	if p.Attendance.Source == AttendanceSourceReports && (p.Loot.NightsSince != nil || p.Attendance.Streak != 0) {
		t.Fatalf("reports mode carries no nights: since %v, streak %d", p.Loot.NightsSince, p.Attendance.Streak)
	}
	if p.Notes == nil || len(*p.Notes) != 1 || (*p.Notes)[0].Body != secretNote {
		t.Fatalf("officer's notes = %v", p.Notes)
	}
	h.as(member)
	var self MemberProfile
	h.data(h.do(http.MethodGet, profile, ""), &self)
	if self.Member == nil || self.Notes != nil {
		t.Fatalf("self profile = %+v", self)
	}
}
