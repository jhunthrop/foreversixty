package guilds

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// ownerOnlyScores is a scorer whose one score is hidden: it answers it only to owner, the way
// rating.Service.ForeverScore reads the viewer off the request (score_visibility.go).
type ownerOnlyScores struct {
	owner int64
	asked []int64
}

func (o *ownerOnlyScores) ScoreOf(ctx context.Context, _ string) (*int, error) {
	viewer := auth.ActorFrom(ctx).UserID
	o.asked = append(o.asked, viewer)
	if viewer != o.owner {
		return nil, nil
	}
	score := 74
	return &score, nil
}

// An applicant who hid their Forever Score shows none on the officers' sheet: the sheet asks
// the scorer as the officer, never as the applicant.
func TestAHiddenScoreIsOffTheApplicantSheet(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Hushed")
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "hushed-app@example.com", "hushed")
	h.as(uid)
	if res := h.do(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)); res.StatusCode != http.StatusCreated {
		t.Fatalf("apply = %d: %s", res.StatusCode, readBody(res))
	}
	scores := &ownerOnlyScores{owner: uid}
	h.svc.Scores = scores
	h.as(officer)
	var page struct{ Rows []OfficerApplication }
	h.data(h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/applications", gid), ""), &page)
	if len(page.Rows) != 1 || page.Rows[0].Sheet == nil || page.Rows[0].Sheet.Score != nil {
		t.Fatalf("rows %+v", page.Rows)
	}
	if len(scores.asked) != 1 || scores.asked[0] != officer {
		t.Fatalf("the sheet asked as %v, want the officer %d", scores.asked, officer)
	}
}
