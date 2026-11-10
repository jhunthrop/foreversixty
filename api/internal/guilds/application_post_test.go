// api/internal/guilds/application_post_test.go
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

// postingAppBot is fakeAppBot with the recruiting post; each post arrives on posts.
type postingAppBot struct {
	fakeAppBot
	posts chan postedApplication
}

type postedApplication struct {
	guildID int64
	post    ApplicationPost
}

func (b *postingAppBot) PostApplication(_ context.Context, guildID int64, p ApplicationPost) (string, error) {
	b.posts <- postedApplication{guildID, p}
	return OutcomeSent, nil
}

// A filed application posts to the recruiting channel, off the request.
func TestApplyPostsToTheRecruitingChannel(t *testing.T) {
	h := newHTTPHarness(t)
	gid, officer := recruitGuild(t, h, "Posters")
	bot := &postingAppBot{posts: make(chan postedApplication, 1)}
	h.svc.ApplicationBot = bot
	h.svc.PublicBaseURL = testSiteBase + "/"
	openRecruiting(t, h, gid, officer)
	uid, key := applicant(t, h, "posters-app@example.com", "ashen")
	h.as(uid)
	if got := h.status(http.MethodPost, fmt.Sprintf("/v1/guilds/%d/applications", gid), applyBody(key, false)); got != http.StatusCreated {
		t.Fatalf("apply = %d", got)
	}
	select {
	case got := <-bot.posts:
		p := got.post
		if got.guildID != gid || p.Character != "ashen" || p.Message != "Two tiers of heals." || p.SpecClass != "Paladin" ||
			!strings.HasPrefix(p.ReviewURL, testSiteBase+"/guild/") || !strings.HasSuffix(p.ReviewURL, "/applications") {
			t.Fatalf("post = %+v", got)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("no recruiting post")
	}
}
