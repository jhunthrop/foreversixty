package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// as makes userID the harness's actor (0 signs out).
func (h *httpHarness) as(userID int64) {
	h.actor = auth.Actor{UserID: userID, Role: "user", Method: "session"}
	if userID == 0 {
		h.actor = auth.Actor{}
	}
}

// status sends the request and answers its status, closing the body.
func (h *httpHarness) status(method, path, body string) int {
	h.t.Helper()
	res := h.do(method, path, body)
	_, _ = io.Copy(io.Discard, res.Body)
	res.Body.Close()
	return res.StatusCode
}

// seedSiteCharacter inserts a characters row owned by userID.
func seedSiteCharacter(t *testing.T, h *httpHarness, userID int64, key, region, ruleset, name, class string) {
	t.Helper()
	if _, err := h.pool.Exec(context.Background(), `
		insert into characters (key, region, ruleset, name, class, user_id, level, equipped_item_level)
		values ($1, $2, $3, $4, $5, $6, 60, 63)`, key, region, ruleset, name, class, userID); err != nil {
		t.Fatal(err)
	}
}

// recruitGuild seeds a guild with a verified officer and returns both.
func recruitGuild(t *testing.T, h *httpHarness, name string) (int64, int64) {
	t.Helper()
	gid := seedGuild(t, h.pool, name)
	officer := seedUser(t, h.pool, name+"-officer@example.com")
	seedVerifiedMember(t, h, gid, officer, name+"officer", "officer")
	return gid, officer
}

func openRecruiting(t *testing.T, h *httpHarness, gid, officer int64) {
	t.Helper()
	h.as(officer)
	if got := h.status(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/recruit", gid), `{"open":true}`); got != http.StatusOK {
		t.Fatalf("open recruiting = %d", got)
	}
}

func TestRecruitCardIsPublicAndOfficersEditIt(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Olympus")
	path := fmt.Sprintf("/v1/guilds/%d/recruit", gid)

	h.as(0)
	var card RecruitCard
	h.data(h.do(http.MethodGet, path, ""), &card)
	if card.Open || card.Guild.Name != "Olympus" || card.Members != 1 || len(card.Wanted) != 0 || len(card.Officers) != 0 {
		t.Fatalf("fresh card = %+v", card)
	}
	if got := h.status(http.MethodPatch, path, `{"open":true}`); got != http.StatusUnauthorized {
		t.Fatalf("signed-out patch = %d, want 401", got)
	}
	member := seedUser(t, h.pool, "recruit-member@example.com")
	seedVerifiedMember(t, h, gid, member, "recruitmember", "member")
	h.as(member)
	if got := h.status(http.MethodPatch, path, `{"open":true}`); got != http.StatusForbidden {
		t.Fatalf("member patch = %d, want 403", got)
	}

	h.as(officer)
	for _, bad := range []string{
		`{"wanted":[{"class":"deathknight"}]}`,
		`{"wanted":[{"class":"priest","spec":"paladin-holy"}]}`,
		`{"discord_role_id":"abc"}`,
		fmt.Sprintf(`{"pitch":%q}`, string(make([]byte, 501))),
	} {
		if got := h.status(http.MethodPatch, path, bad); got != http.StatusBadRequest {
			t.Fatalf("patch %s = %d, want 400", bad, got)
		}
	}
	body := `{"open":true,"pitch":"Two nights a week.","raid_nights":"Tue and Thu, 8 pm server",
		"wanted":[{"class":"Paladin","spec":"paladin-holy","note":"main heals"},{"class":"warrior"}],
		"discord_role_id":"123456789012345678"}`
	var settings struct {
		Card          RecruitCard `json:"card"`
		DiscordRoleID *string     `json:"discord_role_id"`
	}
	h.data(h.do(http.MethodPatch, path, body), &settings)
	if settings.DiscordRoleID == nil || *settings.DiscordRoleID != "123456789012345678" {
		t.Fatalf("role = %v", settings.DiscordRoleID)
	}
	w := settings.Card.Wanted
	if len(w) != 2 || w[0].Class != "paladin" || w[0].SpecName == nil || *w[0].SpecName != "Holy Paladin" || w[1].Spec != nil {
		t.Fatalf("wanted = %+v", w)
	}

	h.as(0)
	res := h.do(http.MethodGet, path, "")
	raw, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if !json.Valid(raw) || strings.Contains(string(raw), "123456789012345678") {
		t.Fatal("the public card carries the recruit role id")
	}
}

func TestRecruitCardNamesOnlyOfficersThePageAlreadyNames(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Named")
	hidden := seedUser(t, h.pool, "named-hidden@example.com")
	seedVerifiedMember(t, h, gid, hidden, "namedhidden", "leader")
	raider := seedUser(t, h.pool, "named-raider@example.com")
	seedVerifiedMember(t, h, gid, raider, "namedraider", "member")
	seedRankedReport(t, h.pool, gid, "rpt-named")
	if err := db.EnsureMetricsPartitions(context.Background(), h.pool, time.Now(), 0); err != nil {
		t.Fatal(err)
	}
	for i, key := range []string{"us/hardcore/Namedofficer", "us/hardcore/namedraider"} {
		if _, err := h.pool.Exec(context.Background(), `
			insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role, kill, encounter_id, fought_at, metric_dps)
			values ('rpt-named', $1, $2, $3, 'Paladin', 'Holy', 'healer', true, 1084, now(), 100)`, i, key, fmt.Sprint("Name", i)); err != nil {
			t.Fatal(err)
		}
	}
	openRecruiting(t, h, gid, officer)
	if _, err := h.pool.Exec(context.Background(), `insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Ony', now() + interval '2 days', 180, $2)`, gid, officer); err != nil {
		t.Fatal(err)
	}
	h.as(0)
	var card RecruitCard
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/recruit", gid), ""), &card)
	if len(card.Officers) != 1 || card.Officers[0].Name != "Name0" || card.Officers[0].Rank != "officer" {
		t.Fatalf("officers = %+v (the leader has no public kill, the raider is no officer)", card.Officers)
	}
	if len(card.RaidNights.Upcoming) != 1 || card.RaidNights.Text != nil {
		t.Fatalf("nights = %+v", card.RaidNights)
	}
}

// fakeAppBot scripts the bot's answers and records its calls.
type fakeAppBot struct {
	grant, dm string
	grants    []string
	dms       []string
}

func (f *fakeAppBot) GrantGuildRole(_ context.Context, _, _ int64, roleID string) (string, error) {
	f.grants = append(f.grants, roleID)
	return f.grant, nil
}

func (f *fakeAppBot) DirectMessage(_ context.Context, _ int64, text string) (string, error) {
	f.dms = append(f.dms, text)
	return f.dm, nil
}

// applicant seeds an account with a character on the guild's realm and answers both.
func applicant(t *testing.T, h *httpHarness, email, name string) (int64, string) {
	t.Helper()
	uid := seedUser(t, h.pool, email)
	key := "us/hardcore/" + name
	seedSiteCharacter(t, h, uid, key, "us", "hardcore", name, "paladin")
	return uid, key
}

func applyBody(key string, dm bool) string {
	return fmt.Sprintf(`{"character_key":%q,"message":"Two tiers of heals.","discord_dm":%v}`, key, dm)
}

func TestApplyRefusals(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Gates")
	path := fmt.Sprintf("/v1/guilds/%d/applications", gid)
	uid, key := applicant(t, h, "gates-app@example.com", "ashen")

	h.as(0)
	if got := h.status(http.MethodPost, path, applyBody(key, false)); got != http.StatusUnauthorized {
		t.Fatalf("signed out = %d", got)
	}
	h.as(uid)
	if got := h.status(http.MethodPost, path, applyBody(key, false)); got != http.StatusConflict {
		t.Fatalf("not recruiting = %d, want 409", got)
	}
	openRecruiting(t, h, gid, officer)
	h.as(uid)
	other, otherKey := applicant(t, h, "gates-other@example.com", "otherchar")
	_ = other
	if got := h.status(http.MethodPost, path, applyBody(otherKey, false)); got != http.StatusBadRequest {
		t.Fatalf("someone else's character = %d, want 400", got)
	}
	seedSiteCharacter(t, h, uid, "eu/hardcore/ashen", "eu", "hardcore", "ashen", "paladin")
	if got := h.status(http.MethodPost, path, applyBody("eu/hardcore/ashen", false)); got != http.StatusBadRequest {
		t.Fatalf("another region = %d, want 400", got)
	}
	if got := h.status(http.MethodPost, path, applyBody(key, false)); got != http.StatusCreated {
		t.Fatalf("apply = %d, want 201", got)
	}
	if got := h.status(http.MethodPost, path, applyBody(key, false)); got != http.StatusConflict {
		t.Fatalf("second open application = %d, want 409", got)
	}
	seedSiteCharacter(t, h, officer, "us/hardcore/gatesalt", "us", "hardcore", "gatesalt", "mage")
	h.as(officer)
	if got := h.status(http.MethodPost, path, applyBody("us/hardcore/gatesalt", false)); got != http.StatusConflict {
		t.Fatalf("a member applying to their own guild = %d, want 409", got)
	}
}

func TestApplicantReadsAndWithdrawsOnlyTheirOwn(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Mine")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "mine-app@example.com", "minechar")
	h.as(uid)
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, true)), &a)
	if a.Status != ApplicationPending || !a.DiscordDM || a.Character.Name != "minechar" || a.Character.ItemLevel == nil {
		t.Fatalf("application = %+v", a)
	}
	stranger := seedUser(t, h.pool, "mine-stranger@example.com")
	h.as(stranger)
	var mine struct{ Rows []Application }
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/applications/mine", gid), ""), &mine)
	if len(mine.Rows) != 0 {
		t.Fatalf("a stranger read %d applications", len(mine.Rows))
	}
	withdraw := fmt.Sprintf("/v1/guilds/%d/applications/%d/withdraw", gid, a.ID)
	if got := h.status(http.MethodPost, withdraw, ""); got != http.StatusNotFound {
		t.Fatalf("stranger withdraw = %d, want 404", got)
	}
	h.as(uid)
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/applications/mine", gid), ""), &mine)
	if len(mine.Rows) != 1 {
		t.Fatalf("own applications = %d", len(mine.Rows))
	}
	if got := h.status(http.MethodPost, withdraw, ""); got != http.StatusOK {
		t.Fatalf("withdraw = %d", got)
	}
	if got := h.status(http.MethodPost, withdraw, ""); got != http.StatusConflict {
		t.Fatalf("second withdraw = %d, want 409", got)
	}
	if n, err := h.store.PendingApplicationsCount(context.Background(), gid); err != nil || n != 0 {
		t.Fatalf("pending = %d, %v", n, err)
	}
}

func TestOfficersListAndVote(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Votes")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "votes-app@example.com", "votechar")
	h.as(uid)
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)), &a)

	list := fmt.Sprintf("/v1/guilds/%d/applications", gid)
	vote := fmt.Sprintf("/v1/guilds/%d/applications/%d/vote", gid, a.ID)
	if got := h.status(http.MethodGet, list, ""); got != http.StatusForbidden {
		t.Fatalf("applicant reads the officer list = %d, want 403", got)
	}
	if got := h.status(http.MethodPut, vote, `{"vote":"yes"}`); got != http.StatusForbidden {
		t.Fatalf("applicant votes = %d, want 403", got)
	}
	h.as(officer)
	if got := h.status(http.MethodPut, vote, `{"vote":"maybe"}`); got != http.StatusBadRequest {
		t.Fatalf("bad vote = %d", got)
	}
	var v Votes
	h.data(h.do(http.MethodPut, vote, `{"vote":"yes"}`), &v)
	if v.Yes != 1 || v.No != 0 || v.Voters != 1 || v.Mine == nil || *v.Mine != VoteYes {
		t.Fatalf("after yes = %+v", v)
	}
	h.data(h.do(http.MethodPut, vote, `{"vote":"no"}`), &v)
	if v.Yes != 0 || v.No != 1 || *v.Mine != VoteNo {
		t.Fatalf("after change = %+v", v)
	}
	var page struct {
		Rows    []OfficerApplication
		Pending int
	}
	h.data(h.do(http.MethodGet, list, ""), &page)
	if page.Pending != 1 || len(page.Rows) != 1 || page.Rows[0].Votes.No != 1 || page.Rows[0].Character.Name != "votechar" {
		t.Fatalf("list = %+v", page)
	}
	h.data(h.do(http.MethodDelete, vote, ""), &v)
	if v.Yes+v.No != 0 || v.Mine != nil {
		t.Fatalf("after remove = %+v", v)
	}
}

func TestAcceptJoinsTheApplicantAndAsksTheBot(t *testing.T) {
	h := newHTTPHarness(t)
	bot := &fakeAppBot{grant: OutcomeGranted, dm: OutcomeSent}
	h.svc.ApplicationBot = bot
	gid, officer := recruitGuild(t, h, "Accept")
	openRecruiting(t, h, gid, officer)
	h.as(officer)
	if got := h.status(http.MethodPatch, fmt.Sprintf("/v1/guilds/%d/recruit", gid), `{"discord_role_id":"123456789012345678"}`); got != http.StatusOK {
		t.Fatalf("set role = %d", got)
	}
	uid, key := applicant(t, h, "accept-app@example.com", "acceptchar")
	h.as(uid)
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, true)), &a)
	accept := fmt.Sprintf("/v1/guilds/%d/applications/%d/accept", gid, a.ID)
	if got := h.status(http.MethodPost, accept, ""); got != http.StatusForbidden {
		t.Fatalf("applicant accepts = %d, want 403", got)
	}

	h.as(officer)
	var d Decision
	h.data(h.do(http.MethodPost, accept, ""), &d)
	if d.Application.Status != ApplicationAccepted || !d.DiscordRole.Done || !d.DiscordDM.Done {
		t.Fatalf("decision = %+v", d)
	}
	if len(bot.grants) != 1 || bot.grants[0] != "123456789012345678" || len(bot.dms) != 1 {
		t.Fatalf("bot calls: grants %v dms %v", bot.grants, bot.dms)
	}
	if d.Application.DecidedBy == nil || d.Application.RoleOutcome == nil || *d.Application.RoleOutcome != OutcomeGranted {
		t.Fatalf("recorded = %+v", d.Application)
	}
	member, err := h.store.IsVerifiedMember(context.Background(), gid, uid)
	if err != nil || !member {
		t.Fatalf("accepted applicant is a verified member = %v, %v", member, err)
	}
	if got := h.status(http.MethodPost, accept, ""); got != http.StatusConflict {
		t.Fatalf("accept twice = %d, want 409", got)
	}
}

func TestAcceptWithoutARoleOrOptInSaysWhy(t *testing.T) {
	h := newHTTPHarness(t)
	bot := &fakeAppBot{grant: OutcomeGranted, dm: OutcomeSent}
	h.svc.ApplicationBot = bot
	gid, officer := recruitGuild(t, h, "NoRole")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "norole-app@example.com", "norolechar")
	h.as(uid)
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)), &a)
	h.as(officer)
	var d Decision
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications/%d/accept", gid, a.ID), ""), &d)
	if d.DiscordRole.Done || d.DiscordRole.Reason == nil || *d.DiscordRole.Reason != OutcomeNoRoleSet {
		t.Fatalf("role = %+v", d.DiscordRole)
	}
	if d.DiscordDM.Done || d.DiscordDM.Reason == nil || *d.DiscordDM.Reason != OutcomeNotOptedIn {
		t.Fatalf("dm = %+v", d.DiscordDM)
	}
	if len(bot.grants)+len(bot.dms) != 0 {
		t.Fatalf("the bot was asked: %v %v", bot.grants, bot.dms)
	}
}

func TestDeclineDMsOnlyWithTheOptInAndNeverGrants(t *testing.T) {
	h := newHTTPHarness(t)
	bot := &fakeAppBot{grant: OutcomeGranted, dm: OutcomeNotLinked}
	h.svc.ApplicationBot = bot
	gid, officer := recruitGuild(t, h, "Decline")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "decline-app@example.com", "declinechar")
	h.as(uid)
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, true)), &a)
	h.as(officer)
	var d Decision
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications/%d/decline", gid, a.ID), ""), &d)
	if d.Application.Status != ApplicationDeclined || len(bot.grants) != 0 || len(bot.dms) != 1 {
		t.Fatalf("decline = %+v, bot %v %v", d, bot.grants, bot.dms)
	}
	if d.DiscordDM.Reason == nil || *d.DiscordDM.Reason != OutcomeNotLinked {
		t.Fatalf("dm = %+v", d.DiscordDM)
	}
	if member, _ := h.store.IsVerifiedMember(context.Background(), gid, uid); member {
		t.Fatal("a declined applicant joined")
	}
}

func TestDecisionTextHasNoDash(t *testing.T) {
	for _, accept := range []bool{true, false} {
		if text := decisionText("Olympus", "Ash", accept); strings.Contains(text, "—") || !strings.HasPrefix(text, "Olympus") {
			t.Fatalf("text = %q", text)
		}
	}
}
