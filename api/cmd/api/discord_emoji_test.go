package main

import (
	"bytes"
	"context"
	"image"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/HugoSmits86/nativewebp"

	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

const emojiCmdToken = "cmd-secret-bot-token"

// stubEmojiAPI holds uploads in memory and records the credentials it was built with.
type stubEmojiAPI struct {
	held    map[string]string
	created []string
}

func (s *stubEmojiAPI) ListAppEmoji(context.Context) ([]discordbot.AppEmoji, error) {
	var out []discordbot.AppEmoji
	for name, id := range s.held {
		out = append(out, discordbot.AppEmoji{ID: id, Name: name})
	}
	return out, nil
}

func (s *stubEmojiAPI) CreateAppEmoji(_ context.Context, name, _ string) (discordbot.AppEmoji, error) {
	s.held[name] = "id-" + name
	s.created = append(s.created, name)
	return discordbot.AppEmoji{ID: "id-" + name, Name: name}, nil
}

func (s *stubEmojiAPI) DeleteAppEmoji(_ context.Context, id string) error {
	for name, held := range s.held {
		if held == id {
			delete(s.held, name)
		}
	}
	return nil
}

func emojiEnv(key string) string {
	return map[string]string{"DISCORD_APPLICATION_ID": "app1", "DISCORD_BOT_TOKEN": emojiCmdToken}[key]
}

func TestDiscordEmojiUploadsOnlyTheMissingAndPrintsNoToken(t *testing.T) {
	icons := fakeIconDir(t)
	stub := &stubEmojiAPI{held: map[string]string{"fs_mage_frost": "77"}}
	var gotApp, gotToken string
	factory := func(app, token string) discordbot.EmojiAPI {
		gotApp, gotToken = app, token
		return stub
	}
	var out bytes.Buffer
	if err := runDiscordEmoji(context.Background(), []string{"--icons", icons}, emojiEnv, factory, &out); err != nil {
		t.Fatal(err)
	}
	if len(stub.created) != 27 || gotApp != "app1" || gotToken != emojiCmdToken {
		t.Fatalf("created %d app %q", len(stub.created), gotApp)
	}
	if err := runDiscordEmoji(context.Background(), []string{"--icons", icons}, emojiEnv, factory, &out); err != nil {
		t.Fatal(err)
	}
	if len(stub.created) != 27 {
		t.Fatalf("a second run uploaded %d", len(stub.created)-27)
	}
	if strings.Contains(out.String(), emojiCmdToken) || !strings.Contains(out.String(), "0 uploaded, 28 kept") {
		t.Fatalf("output:\n%s", out.String())
	}
}

func TestTheSetupCommandsAreRouted(t *testing.T) {
	routed := map[string]bool{}
	for _, job := range allCommands() {
		routed[job.name] = true
	}
	for _, name := range []string{"discord-emoji", "stripe-setup"} {
		if !routed[name] {
			t.Errorf("%s is not routed", name)
		}
	}
}

func TestDiscordEmojiNeedsBothVariables(t *testing.T) {
	err := runDiscordEmoji(context.Background(), nil, func(string) string { return "" }, nil, &bytes.Buffer{})
	if err == nil || strings.Contains(err.Error(), emojiCmdToken) {
		t.Fatalf("err %v", err)
	}
}

func TestNewestBuildSkipsPrebetaAndOrdersByNumber(t *testing.T) {
	root := t.TempDir()
	for _, d := range []string{"forever-prebeta", "1.60.1.9999", "1.60.1.70658", "1.15.9.69722"} {
		if err := os.Mkdir(filepath.Join(root, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if got, err := newestBuild(root); err != nil || got != "1.60.1.70658" {
		t.Fatalf("newest %q err %v", got, err)
	}
	if _, err := newestBuild(t.TempDir()); err == nil {
		t.Fatal("an empty root has a build")
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
