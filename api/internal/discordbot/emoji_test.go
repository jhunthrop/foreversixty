package discordbot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/HugoSmits86/nativewebp"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

const emojiTestToken = "emoji-secret-bot-token"

// fakeEmojiDiscord is the application-emoji endpoints, holding what was uploaded.
type fakeEmojiDiscord struct {
	mu      sync.Mutex
	held    map[string]string // name → id
	next    int
	creates int
	deletes int
	auth    []string
	// limitOnce answers the next POST with a 429 (retry after 10 ms); goneIDs answer DELETE 10014.
	limitOnce bool
	goneIDs   map[string]bool
}

func (f *fakeEmojiDiscord) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.auth = append(f.auth, r.Header.Get("Authorization"))
	base := "/applications/app1/emojis"
	switch {
	case r.Method == http.MethodGet && r.URL.Path == base:
		items := []AppEmoji{{ID: "9", Name: "someone_elses"}}
		for name, id := range f.held {
			items = append(items, AppEmoji{ID: id, Name: name})
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"items": items})
	case r.Method == http.MethodPost && r.URL.Path == base && f.limitOnce:
		f.limitOnce = false
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"message":"You are being rate limited.","retry_after":0.01,"global":false}`))
	case r.Method == http.MethodDelete && f.goneIDs[strings.TrimPrefix(r.URL.Path, base+"/")]:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"Unknown Emoji","code":10014}`))
	case r.Method == http.MethodPost && r.URL.Path == base:
		var body struct{ Name, Image string }
		_ = json.NewDecoder(r.Body).Decode(&body)
		if !strings.HasPrefix(body.Image, "data:image/png;base64,") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		f.next++
		f.creates++
		id := fmt.Sprint(1300 + f.next)
		f.held[body.Name] = id
		_ = json.NewEncoder(w).Encode(AppEmoji{ID: id, Name: body.Name})
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, base+"/"):
		id := strings.TrimPrefix(r.URL.Path, base+"/")
		for name, held := range f.held {
			if held == id {
				delete(f.held, name)
			}
		}
		f.deletes++
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func newEmojiREST(t *testing.T) (*HTTPREST, *fakeEmojiDiscord) {
	t.Helper()
	fake := &fakeEmojiDiscord{held: map[string]string{}}
	srv := httptest.NewServer(fake)
	t.Cleanup(srv.Close)
	rest := NewHTTPREST("app1", emojiTestToken)
	rest.BaseURL = srv.URL
	return rest, fake
}

func fakeImages() []EmojiImage {
	var out []EmojiImage
	for _, sp := range raidspec.All() {
		out = append(out, EmojiImage{Name: sp.Emoji(), PNG: []byte{0x89, 'P', 'N', 'G'}})
	}
	return out
}

func TestSyncEmojiUploadsOnlyTheMissing(t *testing.T) {
	rest, fake := newEmojiREST(t)
	ctx := context.Background()
	fake.held["fs_druid_balance"] = "77"
	var out bytes.Buffer
	rep, err := SyncEmoji(ctx, rest, fakeImages(), false, &out)
	if err != nil || rep.Uploaded != 27 || rep.Kept != 1 || fake.held["fs_druid_balance"] != "77" {
		t.Fatalf("first run %+v err %v", rep, err)
	}
	rep, err = SyncEmoji(ctx, rest, fakeImages(), false, &out)
	if err != nil || rep.Uploaded != 0 || rep.Kept != 28 || fake.creates != 27 {
		t.Fatalf("second run %+v err %v creates %d", rep, err, fake.creates)
	}
	rep, err = SyncEmoji(ctx, rest, fakeImages(), true, &out)
	if err != nil || rep.Deleted != 28 || rep.Uploaded != 28 || fake.held["someone_elses"] != "" {
		t.Fatalf("replace %+v err %v", rep, err)
	}
	if strings.Contains(out.String(), emojiTestToken) || !strings.Contains(out.String(), "uploaded fs_hunter_beast_mastery") {
		t.Fatalf("output:\n%s", out.String())
	}
	for _, a := range fake.auth {
		if a != "Bot "+emojiTestToken {
			t.Fatalf("authorization %q", a)
		}
	}
}

// A 429 waits Discord's retry_after and tries once more; an emoji already gone deletes cleanly.
func TestSyncEmojiRetriesARateLimitOnce(t *testing.T) {
	rest, fake := newEmojiREST(t)
	fake.limitOnce = true
	rep, err := SyncEmoji(context.Background(), rest, fakeImages(), false, &bytes.Buffer{})
	if err != nil || rep.Uploaded != 28 || fake.creates != 28 {
		t.Fatalf("sync %+v err %v creates %d", rep, err, fake.creates)
	}
	fake.goneIDs = map[string]bool{"gone": true}
	if err := rest.DeleteAppEmoji(context.Background(), "gone"); err != nil {
		t.Fatalf("10014 = %v", err)
	}
}

func TestEmojiImagesAre128PixelPNGs(t *testing.T) {
	images, err := EmojiImages(fakeIconDir(t))
	if err != nil || len(images) != 28 {
		t.Fatalf("%d images, err %v", len(images), err)
	}
	for _, img := range images {
		cfg, err := png.DecodeConfig(bytes.NewReader(img.PNG))
		if err != nil || cfg.Width != emojiSide || cfg.Height != emojiSide || len(img.PNG) > emojiMaxBytes {
			t.Fatalf("%s: %+v %d bytes err %v", img.Name, cfg, len(img.PNG), err)
		}
	}
	if _, err := EmojiImages(t.TempDir()); err == nil {
		t.Fatal("an empty folder built images")
	}
}

// countingEmoji answers a fixed list and counts the reads.
type countingEmoji struct {
	items []AppEmoji
	reads int
	err   error
}

func (c *countingEmoji) ListAppEmoji(context.Context) ([]AppEmoji, error) {
	c.reads++
	return c.items, c.err
}
func (c *countingEmoji) CreateAppEmoji(context.Context, string, string) (AppEmoji, error) {
	return AppEmoji{}, nil
}
func (c *countingEmoji) DeleteAppEmoji(context.Context, string) error { return nil }

// The drainer draws the roster with the bot's spec emoji; with no emoji API, as text lines.
func TestTheDrainedPostDrawsSpecEmoji(t *testing.T) {
	for _, withEmoji := range []bool{true, false} {
		w := newPostWorld(t)
		if withEmoji {
			w.svc.emoji = &EmojiRegistry{API: &countingEmoji{items: []AppEmoji{{ID: "42", Name: "fs_paladin_holy"}}}}
		}
		w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) {
			f.Post.Roster = []notify.RosterLine{{Name: "Aurelith", Class: "paladin", Spec: "Holy", Role: "healer", State: "in"}}
		})
		w.enqueue(t, w.raidID)
		w.drain(t)
		creates := w.callsOf("CreateMessage")
		if len(creates) != 1 {
			t.Fatalf("creates %d", len(creates))
		}
		healers := ""
		for _, f := range creates[0].Message.Embeds[0].Fields {
			if f.Name == "Healers · 1" {
				healers = f.Value
			}
		}
		want := "Aurelith · Holy Paladin"
		if withEmoji {
			want = "<:fs_paladin_holy:42> Aurelith"
		}
		if healers != want {
			t.Fatalf("emoji %v: healers %q", withEmoji, healers)
		}
	}
}

func TestEmojiRegistryReadsOnceAndRereadsHourlyWhileMissing(t *testing.T) {
	now := time.Date(2026, 12, 10, 2, 0, 0, 0, time.UTC)
	api := &countingEmoji{items: []AppEmoji{{ID: "1", Name: "fs_druid_balance"}, {ID: "2", Name: "unrelated"}}}
	r := &EmojiRegistry{API: api, Now: func() time.Time { return now }}
	ctx := context.Background()
	if m := r.Marks(ctx); m["fs_druid_balance"] != "1" || m["unrelated"] != "" || api.reads != 1 {
		t.Fatalf("marks %v reads %d", m, api.reads)
	}
	r.Marks(ctx)
	now = now.Add(59 * time.Minute)
	r.Marks(ctx)
	if api.reads != 1 {
		t.Fatalf("reread inside the hour: %d", api.reads)
	}
	for _, sp := range raidspec.All() {
		api.items = append(api.items, AppEmoji{ID: sp.Key, Name: sp.Emoji()})
	}
	now = now.Add(time.Minute)
	if m := r.Marks(ctx); len(m) != 28 || api.reads != 2 {
		t.Fatalf("marks %d reads %d", len(m), api.reads)
	}
	now = now.Add(3 * time.Hour)
	r.Marks(ctx)
	if api.reads != 2 {
		t.Fatal("a complete set was read again")
	}
	if (*EmojiRegistry)(nil).Marks(ctx) != nil {
		t.Fatal("a nil registry has marks")
	}
}

// fakeIconDir is an icon folder holding a 64 px webp for every spec icon, as the active build's
// folder does, so the image path is tested without the generated web data.
func fakeIconDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	img := image.NewNRGBA(image.Rect(0, 0, 64, 64))
	for i := range img.Pix {
		img.Pix[i] = byte(i)
	}
	var buf bytes.Buffer
	if err := nativewebp.Encode(&buf, img, nil); err != nil {
		t.Fatal(err)
	}
	for _, sp := range raidspec.All() {
		if err := os.WriteFile(filepath.Join(dir, sp.Icon+".webp"), buf.Bytes(), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}
