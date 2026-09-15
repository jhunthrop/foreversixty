// Package auth is sign-in: browser sessions with CSRF, Battle.net OAuth,
// email magic links, and the per-device tokens the companion uploads
// with. Everything it stores is in migration 0005.
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base32"
	"encoding/base64"
	"strings"
)

// The identifier shapes the Phase 3 contract fixes.
const (
	// ReportIDChars is 12 lowercase base32 characters, 60 bits.
	ReportIDChars = 12
	// DeviceIDChars is 16 lowercase base32 characters.
	DeviceIDChars = 16
	// DeviceTokenChars is the random half of a device token.
	DeviceTokenChars = 32
	// DeviceTokenPrefix marks a device token wherever it appears.
	DeviceTokenPrefix = "fsd_"
	// PairingCodeChars is the code a person reads off the site and types
	// into the companion.
	PairingCodeChars = 8
)

var lowerBase32 = base32.StdEncoding.WithPadding(base32.NoPadding)

// Base32ID returns n random lowercase base32 characters. Each character
// carries five bits, so the caller's length is the entropy budget.
func Base32ID(n int) string {
	b := make([]byte, (n*5+7)/8)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand never fails on any platform this runs on; if it
		// ever did, continuing with a predictable id would be worse.
		panic("auth: crypto/rand: " + err.Error())
	}
	return strings.ToLower(lowerBase32.EncodeToString(b))[:n]
}

// NewReportID returns a fresh report id: never derived from content, so
// two identical raid nights are still two reports.
func NewReportID() string { return Base32ID(ReportIDChars) }

// NewDeviceID returns a fresh device id.
func NewDeviceID() string { return Base32ID(DeviceIDChars) }

// NewPairingCode returns a fresh pairing code.
func NewPairingCode() string { return Base32ID(PairingCodeChars) }

// NewSessionID returns an opaque 32-byte base64url session id.
func NewSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic("auth: crypto/rand: " + err.Error())
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

// NewDeviceToken returns a device token and the SHA-256 stored for it.
// The token is shown once; only the hash is ever written down.
func NewDeviceToken() (token string, hash []byte) {
	token = DeviceTokenPrefix + Base32ID(DeviceTokenChars)
	return token, TokenHash(token)
}

// TokenHash is the stored form of a secret: SHA-256 of its bytes.
func TokenHash(token string) []byte {
	sum := sha256.Sum256([]byte(token))
	return sum[:]
}

// SameToken compares two secrets in constant time.
func SameToken(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
