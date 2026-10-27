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

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
)

// An Origin-less, cookie-less signed POST reaches the endpoint through the
// full middleware chain: Discord sends neither.
func TestDiscordInteractionsPingThroughTheRouter(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	h := NewRouter(Deps{Version: "test", Auth: &auth.Authenticator{}, DiscordBot: discordbot.New(discordbot.Deps{PublicKey: pub})})
	body := `{"type":1}`
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	sig := hex.EncodeToString(ed25519.Sign(priv, []byte(ts+body)))
	// One past the router's per-IP budget: Discord's egress IPs are few.
	for i := 0; i <= routerRequestsPerMinute; i++ {
		req := httptest.NewRequest(http.MethodPost, discordbot.InteractionsPath, strings.NewReader(body))
		req.Header.Set("X-Signature-Ed25519", sig)
		req.Header.Set("X-Signature-Timestamp", ts)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK || strings.TrimSpace(rec.Body.String()) != `{"type":1}` {
			t.Fatalf("ping %d = %d %s", i+1, rec.Code, rec.Body)
		}
	}
}
