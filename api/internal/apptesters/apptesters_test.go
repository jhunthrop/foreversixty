package apptesters

import (
	"errors"
	"strings"
	"testing"
)

func TestNormalizeEmailTrimsAndLowercases(t *testing.T) {
	got, err := NormalizeEmail("  Thandrel@Example.COM ")
	if err != nil {
		t.Fatal(err)
	}
	if got != "thandrel@example.com" {
		t.Fatalf("got %q", got)
	}
}

func TestNormalizeEmailRefusesWhatIsNotOneAddress(t *testing.T) {
	longLocal := strings.Repeat("a", 65) + "@example.com"
	tooLong := "a@" + strings.Repeat("b", 250) + ".com"
	for _, in := range []string{
		"", " ", "not-an-email", "a@", "@example.com", "a@localhost", "Name <a@example.com>",
		"a@example.com, b@example.com", "a b@example.com", longLocal, tooLong, "a@example.com\nb",
		"a@.example.com", "a@example..com", "a@example.com.",
	} {
		if _, err := NormalizeEmail(in); !errors.Is(err, ErrInvalidEmail) {
			t.Errorf("NormalizeEmail(%q) = %v, want ErrInvalidEmail", in, err)
		}
	}
}

func TestValidPlatformIsIOSOrAndroid(t *testing.T) {
	for _, p := range []string{PlatformIOS, PlatformAndroid} {
		if !ValidPlatform(p) {
			t.Errorf("%q should be valid", p)
		}
	}
	for _, p := range []string{"", "iOS", "iphone", "windows"} {
		if ValidPlatform(p) {
			t.Errorf("%q should not be valid", p)
		}
	}
}
