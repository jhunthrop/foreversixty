package guilds

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// The raid_roster rows read the raid page's own counts; an officer also gets every answer and
// every job, a member only the counts.
func TestAddonRaidRostersNamesOnlyForOfficers(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	w.as(w.member)
	if res := w.putSignup(raid.ID, memberKeyA, `{"state":"in","role":"tank"}`); res.StatusCode != http.StatusOK {
		t.Fatalf("sign-up = %d: %s", res.StatusCode, readBody(res))
	}
	now := time.Now().Add(time.Hour)
	ctx := context.Background()

	asMember, err := w.h.svc.AddonRaidRosters(ctx, w.member, []string{memberKeyA, memberKeyB}, now)
	if err != nil {
		t.Fatal(err)
	}
	if len(asMember) != 2 || asMember[0].CharacterKey != memberKeyA || asMember[1].CharacterKey != memberKeyB {
		t.Fatalf("one row per verified character: %+v", asMember)
	}
	r := asMember[0]
	if r.RaidID != raid.ID || r.Officer || r.Members != nil || r.JobList != nil {
		t.Fatalf("a member's row = %+v", r)
	}
	// The member's account answered; the officer and the second member have not.
	if r.Counts.In != 1 || r.Counts.NotSigned != 2 || r.Coming.Tanks != 1 || r.Wanted == nil || r.Wanted.Tanks != 2 {
		t.Fatalf("counts = %+v coming %+v wanted %+v", r.Counts, r.Coming, r.Wanted)
	}
	if r.Jobs == 0 || r.JobsFilled != 0 {
		t.Errorf("jobs %d of %d", r.JobsFilled, r.Jobs)
	}

	asOfficer, err := w.h.svc.AddonRaidRosters(ctx, w.officer, []string{officerKey}, now)
	if err != nil || len(asOfficer) != 1 || !asOfficer[0].Officer {
		t.Fatalf("officer rows = %+v %v", asOfficer, err)
	}
	var signed, waiting int
	for _, m := range asOfficer[0].Members {
		switch m.Answer {
		case SignupIn:
			signed++
			if m.Name != "Kraggor" || m.Group != "tank" {
				t.Errorf("signed row = %+v", m)
			}
		case AddonAnswerNone:
			waiting++
		}
	}
	if signed != 1 || waiting == 0 || len(asOfficer[0].JobList) == 0 {
		t.Errorf("officer members %+v jobs %+v", asOfficer[0].Members, asOfficer[0].JobList)
	}

	if past, err := w.h.svc.AddonRaidRosters(ctx, w.member, []string{memberKeyA}, now.Add(-30*24*time.Hour)); err != nil || len(past) != 0 {
		t.Errorf("a raid past the window is listed: %+v %v", past, err)
	}
	if none, err := w.h.svc.AddonRaidRosters(ctx, w.member, []string{memberKeyC}, now); err != nil || len(none) != 0 {
		t.Errorf("an unverified character gets rows: %+v %v", none, err)
	}
}
