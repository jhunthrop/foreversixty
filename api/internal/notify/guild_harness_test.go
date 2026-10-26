package notify

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"testing"
)

const guildPath = "/guild/us/normal/olympus-xxvii"

// fakeGuildAccess answers one scripted standing for every viewer.
type fakeGuildAccess struct {
	mu       sync.Mutex
	standing GuildStanding
	found    bool
	err      error
}

func (f *fakeGuildAccess) GuildStanding(context.Context, int64, int64) (GuildStanding, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.standing, f.found, f.err
}

func (f *fakeGuildAccess) set(s GuildStanding) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.standing = s
}

// officerWithPlan is the standing that may connect and test.
func officerWithPlan() GuildStanding {
	return GuildStanding{Officer: true, Plan: true, Name: "Olympus XXVII", Path: guildPath}
}

type guildHarness struct {
	*apiHarness
	access *fakeGuildAccess
	oauth  *FakeOAuthClient
	gid    int64
}

func newGuildAPI(t *testing.T, standing GuildStanding) *guildHarness {
	t.Helper()
	access := &fakeGuildAccess{standing: standing, found: true}
	oauth := connectFake()
	h := newAPIWith(t, oauth, func(s *Service) {
		s.Guilds = access
		s.GuildStore = &GuildStore{Pool: s.Store.Pool}
	})
	return &guildHarness{apiHarness: h, access: access, oauth: oauth, gid: seedGuild(t, h.svc.Store.Pool)}
}

func (g *guildHarness) route(suffix string) string {
	return fmt.Sprintf("/v1/guilds/%d/notifications/discord%s", g.gid, suffix)
}

// startGuildConnect runs the connect leg and returns the state Discord would echo back.
func (g *guildHarness) startGuildConnect() string {
	g.t.Helper()
	res, loc := g.get(g.route("/connect"))
	if res.StatusCode != http.StatusFound {
		g.t.Fatalf("connect status %d", res.StatusCode)
	}
	u, err := url.Parse(loc)
	if err != nil {
		g.t.Fatal(err)
	}
	return u.Query().Get("state")
}

func (g *guildHarness) assertGuildOutcome(loc, want string) {
	g.t.Helper()
	if loc != "https://foreversixty.gg"+guildPath+"/settings?discord="+want+"#discord" {
		g.t.Fatalf("Location = %q, want guild outcome %q", loc, want)
	}
}

func (g *guildHarness) guildStored() (GuildWebhook, bool) {
	g.t.Helper()
	w, ok, err := g.svc.GuildStore.Get(context.Background(), g.gid)
	if err != nil {
		g.t.Fatal(err)
	}
	return w, ok
}

// saveGuildHook stores rawURL as the guild's channel, sealed as the connect would.
func (g *guildHarness) saveGuildHook(rawURL string) {
	g.t.Helper()
	box, err := g.svc.Sealer.Seal(rawURL)
	if err != nil {
		g.t.Fatal(err)
	}
	if err := g.svc.GuildStore.Save(context.Background(), GuildSaveInput{
		GuildID: g.gid, ChannelID: testChan, TokenBox: box, ConnectedBy: g.uid,
	}); err != nil {
		g.t.Fatal(err)
	}
}
