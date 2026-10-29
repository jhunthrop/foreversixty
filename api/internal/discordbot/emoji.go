package discordbot

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// AppEmoji is one of the application's own emoji.
type AppEmoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// EmojiAPI is Discord's application-emoji endpoints (spec 2026-10-28 §P.9.1). An application's
// emoji work in every server the bot posts in, with no server permission.
type EmojiAPI interface {
	ListAppEmoji(ctx context.Context) ([]AppEmoji, error)
	CreateAppEmoji(ctx context.Context, name, imageDataURI string) (AppEmoji, error)
	DeleteAppEmoji(ctx context.Context, id string) error
}

func (r *HTTPREST) emojiPath() string {
	return "/applications/" + url.PathEscape(r.applicationID) + "/emojis"
}

// ListAppEmoji is GET /applications/{id}/emojis.
func (r *HTTPREST) ListAppEmoji(ctx context.Context) ([]AppEmoji, error) {
	var out struct {
		Items []AppEmoji `json:"items"`
	}
	if err := r.do(ctx, call{method: http.MethodGet, path: r.emojiPath(), botAuth: true, out: &out}); err != nil {
		return nil, err
	}
	return out.Items, nil
}

// CreateAppEmoji is POST /applications/{id}/emojis with a data:image/png;base64 image.
func (r *HTTPREST) CreateAppEmoji(ctx context.Context, name, imageDataURI string) (AppEmoji, error) {
	var out AppEmoji
	body := map[string]string{"name": name, "image": imageDataURI}
	if err := r.do(ctx, call{method: http.MethodPost, path: r.emojiPath(), botAuth: true, body: body, out: &out}); err != nil {
		return AppEmoji{}, err
	}
	if out.ID == "" {
		return AppEmoji{}, errors.New("discordbot: emoji answered without an id")
	}
	return out, nil
}

// codeUnknownEmoji is Discord's answer for an emoji that no longer exists.
const codeUnknownEmoji = 10014

// DeleteAppEmoji is DELETE /applications/{id}/emojis/{emoji}; one already gone is success.
func (r *HTTPREST) DeleteAppEmoji(ctx context.Context, id string) error {
	err := r.do(ctx, call{method: http.MethodDelete, path: r.emojiPath() + "/" + url.PathEscape(id), botAuth: true})
	var api *APIError
	if errors.As(err, &api) && api.Code == codeUnknownEmoji {
		return nil
	}
	return err
}

// The registry's timings: a list read is bounded, and while any spec's emoji is missing the
// list is read again at most once an hour (§P.9.1).
const (
	emojiReadTimeout = 5 * time.Second
	emojiRereadAfter = time.Hour
)

// EmojiRegistry holds the application's spec emoji ids by name, in memory only (§P.8: no table
// stores them; Discord's own list is the one source). It reads the list on first use, and again
// at most once an hour while any spec's emoji is missing, logging emoji_missing with the names.
// A spec whose emoji is missing draws the text line.
type EmojiRegistry struct {
	API EmojiAPI
	Now func() time.Time
	Log *slog.Logger

	mu     sync.Mutex
	marks  notify.Marks
	readAt time.Time
	read   bool
}

// Marks are the spec emoji ids by name; nil without an API (every line then draws as text).
func (r *EmojiRegistry) Marks(ctx context.Context) notify.Marks {
	if r == nil || r.API == nil {
		return nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	now := r.now()
	if !r.read || (len(missingEmoji(r.marks)) > 0 && now.Sub(r.readAt) >= emojiRereadAfter) {
		r.reload(ctx, now)
	}
	return r.marks
}

func (r *EmojiRegistry) now() time.Time {
	if r.Now == nil {
		return time.Now()
	}
	return r.Now()
}

func (r *EmojiRegistry) logger() *slog.Logger {
	if r.Log == nil {
		return slog.Default()
	}
	return r.Log
}

// reload reads the list; a failed read keeps what was held and is retried an hour on.
func (r *EmojiRegistry) reload(ctx context.Context, now time.Time) {
	r.read, r.readAt = true, now
	ctx, cancel := context.WithTimeout(ctx, emojiReadTimeout)
	defer cancel()
	items, err := r.API.ListAppEmoji(ctx)
	if err != nil {
		r.logger().Warn("discordbot", "op", "emoji_read", "outcome", "failed", "err", err)
	} else {
		r.marks = specMarks(items)
	}
	if missing := missingEmoji(r.marks); len(missing) > 0 {
		r.logger().Warn("discordbot", "op", "emoji_missing", "names", missing)
	}
}

// specMarks keeps the listed emoji that are spec marks.
func specMarks(items []AppEmoji) notify.Marks {
	wanted := map[string]bool{}
	for _, sp := range raidspec.All() {
		wanted[sp.Emoji()] = true
	}
	marks := notify.Marks{}
	for _, e := range items {
		if wanted[e.Name] && e.ID != "" {
			marks[e.Name] = e.ID
		}
	}
	return marks
}

// missingEmoji are the spec emoji names the marks lack, sorted.
func missingEmoji(marks notify.Marks) []string {
	var out []string
	for _, sp := range raidspec.All() {
		if marks[sp.Emoji()] == "" {
			out = append(out, sp.Emoji())
		}
	}
	slices.Sort(out)
	return out
}
