package guilds

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// Applying (here from the app) tells the guild's officers: guild.applicant, linking to the applicant sheet, with
// Decline as its one action (Review is the notification itself).
func TestApplyPublishesTheApplicantToOfficers(t *testing.T) {
	h := newHTTPHarness(t)
	pub := &recordingPublisher{}
	h.svc.Events = pub
	h.svc.PublicBaseURL = "https://foreversixty.gg"
	gid, officer := recruitGuild(t, h, "Heralds")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "herald-app@example.com", "mirelle")

	// The app's own token applies (spec 2026-11-08-app-flows: Apply from the Raid tab).
	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "device", DeviceID: "phone", DeviceKind: auth.DeviceKindApp}
	var a Application
	h.data(h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)), &a)

	e := pub.waitFor(t, events.GuildApplicant, 1)[0]
	if e.Audience.Kind != events.ToGuildOfficers || e.Audience.GuildID != gid {
		t.Fatalf("audience %+v", e.Audience)
	}
	if want := fmt.Sprintf("https://foreversixty.gg/app/guild/%d/applicants/%d", gid, a.ID); e.Link != want {
		t.Fatalf("link %q, want %q", e.Link, want)
	}
	if !strings.HasPrefix(e.Title, "New applicant: mirelle") {
		t.Fatalf("title %q", e.Title)
	}
	if e.Body == "" || e.Data["guild_id"] != strconv.FormatInt(gid, 10) || e.Data["application_id"] != strconv.FormatInt(a.ID, 10) ||
		e.Data["applicant_name"] != "mirelle" || e.Data["applicant"] != strings.TrimPrefix(e.Title, "New applicant: ") {
		t.Fatalf("event %+v", e)
	}
	if len(e.Actions) != 2 || e.Actions[1].ID != "application_decline" || e.Actions[1].Method != http.MethodPost ||
		e.Actions[1].Path != fmt.Sprintf("/v1/guilds/%d/applications/%d/decline", gid, a.ID) {
		t.Fatalf("actions %+v", e.Actions)
	}
	if e.DedupeKey != fmt.Sprintf("guild.applicant:%d", a.ID) {
		t.Fatalf("dedupe %q", e.DedupeKey)
	}
}

func TestApplicantActionsAcceptAndDecline(t *testing.T) {
	got := applicantActions(2, 31)
	want := []struct{ id, title, path string }{
		{"application_accept", "Accept", "/v1/guilds/2/applications/31/accept"},
		{"application_decline", "Decline", "/v1/guilds/2/applications/31/decline"},
	}
	if len(got) != len(want) {
		t.Fatalf("actions %+v", got)
	}
	for i, w := range want {
		if a := got[i]; a.ID != w.id || a.Title != w.title || a.Path != w.path || a.Method != http.MethodPost || a.Body != nil {
			t.Errorf("action %d %+v, want %+v", i, a, w)
		}
	}
}

func TestApplicantCopy(t *testing.T) {
	spec, class, level := "Holy", "Priest", 60
	c := ApplicantCharacter{Name: "Mirelle", SpecName: &spec, Class: &class, Level: &level}
	if got := applicantTitle(c); got != "New applicant: Mirelle, Holy Priest 60" {
		t.Fatalf("title %q", got)
	}
	if got := applicantTitle(ApplicantCharacter{Name: "Ashen"}); got != "New applicant: Ashen" {
		t.Fatalf("bare title %q", got)
	}
	if got := applicantBody("Heralds", Application{Message: "  "}); got != "Applied to Heralds." {
		t.Fatalf("empty body %q", got)
	}
	long := ""
	for range 50 {
		long += "abcde "
	}
	got := applicantBody("Heralds", Application{Message: long})
	if r := []rune(got); len(r) > applicantBodyMax+1 || r[len(r)-1] != '…' {
		t.Fatalf("long body %q", got)
	}
}
