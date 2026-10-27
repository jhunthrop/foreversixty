package discordbot

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

// MaxClockSkew is how far a request's timestamp may sit from our clock.
const MaxClockSkew = 5 * time.Minute

// MaxInteractionBytes caps an interaction body; the verifier reads no more.
const MaxInteractionBytes = 256 << 10

// Verification failures. Neither carries the signature or the body.
var (
	ErrBadSignature = errors.New("discordbot: bad interaction signature")
	ErrStale        = errors.New("discordbot: interaction timestamp outside the allowed skew")
)

// Verify checks Discord's Ed25519 signature: it signs the timestamp header
// (unix seconds) followed by the raw body. A signature that does not verify is
// ErrBadSignature; a good one whose timestamp is more than MaxClockSkew from
// now, either way, is ErrStale (a replay).
func Verify(pub ed25519.PublicKey, signatureHex, timestamp string, body []byte, now time.Time) error {
	if len(pub) != ed25519.PublicKeySize {
		return ErrBadSignature
	}
	sig, err := hex.DecodeString(signatureHex)
	if err != nil || len(sig) != ed25519.SignatureSize {
		return ErrBadSignature
	}
	seconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ErrBadSignature
	}
	message := make([]byte, 0, len(timestamp)+len(body))
	message = append(append(message, timestamp...), body...)
	if !ed25519.Verify(pub, message, sig) {
		return ErrBadSignature
	}
	skew := now.Sub(time.Unix(seconds, 0))
	if skew > MaxClockSkew || skew < -MaxClockSkew {
		return ErrStale
	}
	return nil
}
