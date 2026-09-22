// api/internal/bnetapi/media_test.go
package bnetapi

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"testing"
)

func TestCharacterMediaReadsAvatarAndMainRaw(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/character-media?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"assets": []map[string]string{
			{"key": "avatar", "value": "https://render.worldofwarcraft.com/us/character/whitemane/1-thoradin-avatar.jpg"},
			{"key": "inset", "value": "https://render.worldofwarcraft.com/us/character/whitemane/1-thoradin-inset.jpg"},
			{"key": "main-raw", "value": "https://render.worldofwarcraft.com/character/whitemane/1-thoradin-main-raw.png"},
			{"key": "main", "value": "https://render.worldofwarcraft.com/character/whitemane/1-thoradin-main.png"},
		}})

	c := newTestClient(fs)
	media, raw, err := c.CharacterMedia(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if media.AvatarURL != "https://render.worldofwarcraft.com/us/character/whitemane/1-thoradin-avatar.jpg" {
		t.Fatalf("AvatarURL = %q", media.AvatarURL)
	}
	if media.RenderURL != "https://render.worldofwarcraft.com/character/whitemane/1-thoradin-main-raw.png" {
		t.Fatalf("RenderURL = %q, want main-raw preferred over main", media.RenderURL)
	}
	if len(raw) == 0 || !strings.Contains(string(raw), "avatar") {
		t.Fatalf("raw = %s, want the verbatim assets body", raw)
	}
}

func TestCharacterMediaFallsBackToMainWhenMainRawIsAbsent(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/loneling/character-media?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"assets": []map[string]string{
			{"key": "avatar", "value": "https://render.worldofwarcraft.com/us/character/whitemane/2-loneling-avatar.jpg"},
			{"key": "main", "value": "https://render.worldofwarcraft.com/character/whitemane/2-loneling-main.png"},
		}})

	c := newTestClient(fs)
	media, _, err := c.CharacterMedia(context.Background(), "us", "whitemane", "Loneling")
	if err != nil {
		t.Fatal(err)
	}
	if media.RenderURL != "https://render.worldofwarcraft.com/character/whitemane/2-loneling-main.png" {
		t.Fatalf("RenderURL = %q, want the main asset as a fallback", media.RenderURL)
	}
}

func TestCharacterMediaDropsAnUntrustedHostAndLogsIt(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/thoradin/character-media?namespace=profile-classic1x-us",
		http.StatusOK, map[string]any{"assets": []map[string]string{
			{"key": "avatar", "value": "https://evil.example.com/avatar.jpg"},
			{"key": "main-raw", "value": "https://render.worldofwarcraft.com/character/whitemane/1-thoradin-main-raw.png"},
		}})

	c := newTestClient(fs)
	var buf bytes.Buffer
	c.Log = slog.New(slog.NewTextHandler(&buf, nil))
	media, _, err := c.CharacterMedia(context.Background(), "us", "whitemane", "Thoradin")
	if err != nil {
		t.Fatal(err)
	}
	if media.AvatarURL != "" {
		t.Fatalf("AvatarURL = %q, want dropped (untrusted host)", media.AvatarURL)
	}
	if media.RenderURL == "" {
		t.Fatal("RenderURL should still be set — only the untrusted asset is dropped")
	}
	if !strings.Contains(buf.String(), "media_host_rejected") || !strings.Contains(buf.String(), "evil.example.com") {
		t.Fatalf("log = %s, want op=media_host_rejected naming the rejected url", buf.String())
	}
}

func TestCharacterMediaPropagatesNotFound(t *testing.T) {
	fs := newFixtureServer(t)
	fs.json(http.MethodPost, "/token", http.StatusOK, map[string]any{"access_token": "tok", "expires_in": 3600})
	fs.json(http.MethodGet, "/profile/wow/character/whitemane/sodpop/character-media?namespace=profile-classic1x-us",
		http.StatusNotFound, nil)

	c := newTestClient(fs)
	if _, _, err := c.CharacterMedia(context.Background(), "us", "whitemane", "Sodpop"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}
