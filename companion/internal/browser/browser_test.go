package browser

import (
	"errors"
	"testing"
)

func TestOnlyHTTPSLinksAndLoopbackHTTPAreOpened(t *testing.T) {
	for raw, want := range map[string]bool{
		"https://foreversixty.gg/connect?code=KQ4-7TM": true,
		"http://127.0.0.1:8080/connect":                true,
		"http://localhost:8080/":                       true,
		"http://foreversixty.gg/":                      false,
		"javascript:alert(1)":                          false,
		"file:///C:/Windows/System32/calc.exe":         false,
		"ms-msdt:/id":                                  false,
		"//foreversixty.gg/":                           false,
		"https://":                                     false,
		"https://user:pw@foreversixty.gg/":             false,
		"":                                             false,
		"not a url":                                    false,
	} {
		if got := Check(raw) == nil; got != want {
			t.Errorf("Check(%q) allowed = %v, want %v", raw, got, want)
		}
	}
}

func TestEachPlatformOpensTheLinkAsOneArgument(t *testing.T) {
	const link = "https://foreversixty.gg/logs?a=1&b=2"
	for goos, want := range map[string]string{
		"windows": "rundll32 url.dll,FileProtocolHandler " + link,
		"darwin":  "open " + link,
		"linux":   "xdg-open " + link,
	} {
		name, args := command(goos, link)
		got := name
		for _, a := range args {
			got += " " + a
		}
		if got != want {
			t.Errorf("%s: %q, want %q", goos, got, want)
		}
		if args[len(args)-1] != link {
			t.Errorf("%s: the link is not the last argument", goos)
		}
	}
}

func TestOpenRefusesBeforeStartingAnything(t *testing.T) {
	old := start
	defer func() { start = old }()
	var ran []string
	start = func(name string, args ...string) error { ran = append(ran, name); return nil }
	if err := Open("file:///etc/passwd"); !errors.Is(err, ErrRefused) || len(ran) != 0 {
		t.Fatalf("err = %v, started %v", err, ran)
	}
	if err := Open("https://foreversixty.gg/"); err != nil || len(ran) != 1 {
		t.Fatalf("err = %v, started %v", err, ran)
	}
	start = func(string, ...string) error { return errors.New("no browser") }
	if err := Open("https://foreversixty.gg/"); err == nil {
		t.Fatal("a failed start was not reported")
	}
}
