package server

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
)

// An Origin-less, cookie-less signed POST reaches the endpoint through the
// full middleware chain: Discord sends neither.
func TestDiscordInteractionsPingThroughTheRouter(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	h := NewRouter(Deps{Version: "test", DiscordBot: discordbot.New(discordbot.Deps{PublicKey: pub})})
	body := `{"type":1}`
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	req := httptest.NewRequest(http.MethodPost, discordbot.InteractionsPath, strings.NewReader(body))
	req.Header.Set("X-Signature-Ed25519", hex.EncodeToString(ed25519.Sign(priv, []byte(ts+body))))
	req.Header.Set("X-Signature-Timestamp", ts)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK || strings.TrimSpace(rec.Body.String()) != `{"type":1}` {
		t.Fatalf("ping = %d %s", rec.Code, rec.Body)
	}
}
