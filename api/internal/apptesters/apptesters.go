// Package apptesters is the phone apps' beta list: the public signup behind
// foreversixty.gg/app (POST /v1/app-testers) and the owner's list of it
// (GET /v1/admin/app-testers). One row per address and platform; nothing is
// mailed from here.
package apptesters

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

// The two platforms a tester can ask for, as the table's check constraint spells them.
const (
	PlatformIOS     = "ios"
	PlatformAndroid = "android"
)

// RFC 5321's limits: the whole path, and the local part.
const (
	maxEmailLength = 254
	maxLocalLength = 64
)

// ErrInvalidEmail is NormalizeEmail's answer for anything but one plain address.
var ErrInvalidEmail = errors.New("apptesters: not one email address")

// Tester is one signup, as the owner's list shows it.
type Tester struct {
	Email     string    `json:"email"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"created_at"`
}

// ValidPlatform reports whether p is one of the two platforms.
func ValidPlatform(p string) bool {
	return p == PlatformIOS || p == PlatformAndroid
}

// NormalizeEmail trims and lowercases one plain address (no display name, no
// list), refusing one over RFC 5321's lengths or whose domain is not a dotted
// name.
func NormalizeEmail(raw string) (string, error) {
	in := strings.TrimSpace(raw)
	if in == "" || len(in) > maxEmailLength || strings.ContainsAny(in, " \t\r\n<>,") {
		return "", ErrInvalidEmail
	}
	addr, err := mail.ParseAddress(in)
	if err != nil || addr.Name != "" || addr.Address != in {
		return "", ErrInvalidEmail
	}
	local, domain, ok := strings.Cut(in, "@")
	if !ok || local == "" || len(local) > maxLocalLength || !validDomain(domain) {
		return "", ErrInvalidEmail
	}
	return strings.ToLower(in), nil
}

// validDomain is a dotted name with no empty label.
func validDomain(d string) bool {
	labels := strings.Split(d, ".")
	if len(labels) < 2 {
		return false
	}
	for _, l := range labels {
		if l == "" {
			return false
		}
	}
	return true
}
