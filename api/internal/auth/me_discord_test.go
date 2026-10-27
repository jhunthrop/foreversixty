package auth

import (
	"context"
	"net/http"
	"testing"
)

// fakeDiscordLinks answers DiscordLinkView from scripted values.
type fakeDiscordLinks struct {
	available bool
	link      *DiscordLinkBlock
	suggested *string
}

func (f *fakeDiscordLinks) LinkAvailable() bool { return f.available }

func (f *fakeDiscordLinks) DiscordLink(context.Context, int64) (*DiscordLinkBlock, *string, error) {
	return f.link, f.suggested, nil
}

type meDiscord struct {
	Available bool              `json:"discord_link_available"`
	Link      *DiscordLinkBlock `json:"discord_link"`
	Suggested *string           `json:"discord_link_suggested"`
}

func TestMeCarriesTheDiscordLink(t *testing.T) {
	h := newHarness(t)
	uid := h.seedUser(t, "me-discord@example.com")
	read := func() meDiscord {
		var body meDiscord
		h.decode(t, h.sessionRequest(t, http.MethodGet, "/v1/me", uid), &body)
		return body
	}
	olympus := "Olympus XXVII"

	if got := read(); got.Available || got.Link != nil || got.Suggested != nil {
		t.Fatalf("with no view wired = %+v, want unavailable", got)
	}
	fake := &fakeDiscordLinks{available: true, suggested: &olympus}
	h.svc.DiscordLinks = fake
	if got := read(); !got.Available || got.Link != nil || got.Suggested == nil || *got.Suggested != olympus {
		t.Fatalf("unlinked = %+v, want available, no link, suggested %q", got, olympus)
	}
	fake.link, fake.suggested = &DiscordLinkBlock{DisplayName: "zulmara"}, nil
	if got := read(); !got.Available || got.Link == nil || got.Link.DisplayName != "zulmara" || got.Suggested != nil {
		t.Fatalf("linked = %+v, want zulmara and no suggestion", got)
	}
	fake.available = false
	if got := read(); got.Available || got.Link != nil {
		t.Fatalf("link off = %+v, want unavailable", got)
	}
}
