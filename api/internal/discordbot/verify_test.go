package discordbot

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strconv"
	"testing"
	"time"
)

func signedAt(t *testing.T, priv ed25519.PrivateKey, at time.Time, body []byte) (sig, ts string) {
	t.Helper()
	ts = strconv.FormatInt(at.Unix(), 10)
	return hex.EncodeToString(ed25519.Sign(priv, append([]byte(ts), body...))), ts
}

func newKeys(t *testing.T) (ed25519.PublicKey, ed25519.PrivateKey) {
	t.Helper()
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	return pub, priv
}

func TestVerify(t *testing.T) {
	pub, priv := newKeys(t)
	now := time.Now()
	body := []byte(`{"type":1}`)
	good, ts := signedAt(t, priv, now, body)
	stale, staleTS := signedAt(t, priv, now.Add(-6*time.Minute), body)
	flipped := append([]byte(nil), body...)
	flipped[2] ^= 1

	cases := []struct {
		name    string
		sig, ts string
		body    []byte
		want    error
	}{
		{"good", good, ts, body, nil},
		{"flipped body", good, ts, flipped, ErrBadSignature},
		{"stale", stale, staleTS, body, ErrStale},
		{"empty signature", "", ts, body, ErrBadSignature},
		{"non-hex signature", "zz" + good[2:], ts, body, ErrBadSignature},
		{"bad timestamp", good, "soon", body, ErrBadSignature},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := Verify(pub, c.sig, c.ts, c.body, now); !errors.Is(err, c.want) {
				t.Fatalf("err = %v, want %v", err, c.want)
			}
		})
	}
	if err := Verify(nil, good, ts, body, now); !errors.Is(err, ErrBadSignature) {
		t.Fatalf("no key: %v", err)
	}
}
