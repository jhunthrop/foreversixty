// api/internal/guilds/raids_discord_budget_test.go
package guilds

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// hungPoster is Discord never answering: every Post blocks until its context ends.
type hungPoster struct{ notify.FakePoster }

func (h *hungPoster) Post(ctx context.Context, _ string, _ notify.Embed) error {
	<-ctx.Done()
	return ctx.Err()
}

// syncBuffer is a log sink safe to read while the server writes.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

// realChannel saves a sealed channel for w's guild and posts through the real deliverer.
func realChannel(t *testing.T, w raidWorld, poster notify.Poster) *notify.GuildStore {
	t.Helper()
	st := &notify.GuildStore{Pool: w.h.pool}
	sealer, err := notify.NewSealer(make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	box, err := sealer.Seal("https://discord.com/api/webhooks/1/guild-token")
	if err != nil {
		t.Fatal(err)
	}
	if err := st.Save(t.Context(), notify.GuildSaveInput{
		GuildID: w.gid, ChannelID: "1", TokenBox: box, ConnectedBy: w.officer,
	}); err != nil {
		t.Fatal(err)
	}
	w.h.svc.RaidEvents = &notify.GuildDeliverer{Store: st, Sealer: sealer, Poster: poster}
	return st
}

func TestRaidPostsStopAtTheBudgetAndLogTheSkipped(t *testing.T) {
	w, _, _ := postingWorld(t)
	realChannel(t, w, &hungPoster{})
	logs := &syncBuffer{}
	w.h.svc.Log = slog.New(slog.NewTextHandler(logs, nil))
	const budget = 300 * time.Millisecond
	w.h.svc.raidPostBudget = budget

	began := time.Now()
	views := w.createRaids(t, postingBody(futureStart(), 3, "true"))
	if took := time.Since(began); took > budget+2*time.Second {
		t.Fatalf("create took %v with a %v budget", took, budget)
	}
	if len(views) != 4 {
		t.Fatalf("created %d raids", len(views))
	}
	want := fmt.Sprintf("op=raid_post_skipped guild_id=%d raid_ids=\"[%d %d %d]\"", w.gid, views[1].ID, views[2].ID, views[3].ID)
	if !strings.Contains(logs.String(), want) {
		t.Fatalf("logs lack %q:\n%s", want, logs.String())
	}
}

func TestCreateRaidAnswersWhenThePostFails(t *testing.T) {
	w, _, _ := postingWorld(t)
	poster := &notify.FakePoster{PostErr: &notify.DiscordError{Status: 500, Body: "boom"}}
	st := realChannel(t, w, poster)
	views := w.createRaids(t, postingBody(futureStart(), 3, "true"))
	if len(views) != 4 || len(poster.Posts()) != 4 {
		t.Fatalf("views = %d, posts = %d", len(views), len(poster.Posts()))
	}
	hook, ok, err := st.Get(t.Context(), w.gid)
	if err != nil || !ok || hook.Failures != 4 || hook.LastError == nil || *hook.LastError != "boom" {
		t.Fatalf("channel after failed posts = %+v ok=%v err=%v", hook, ok, err)
	}
}
