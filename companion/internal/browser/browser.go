// companion/internal/browser/browser.go
// Package browser opens a link in the player's own browser. Every
// external link in the window goes through it (the webview would
// otherwise open a second, unbranded window of its own), and so does
// the Battle.net sign-in: the player logs in on a page they can see
// the address of, never inside this app.
package browser

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

// ErrRefused is a link that is not allowed to be opened.
var ErrRefused = errors.New("only https links are opened")

// Check allows an https URL with a host, and http only for the
// loopback addresses a development build of the site answers on. The
// local page's own JavaScript asks for these, so a file:, javascript:
// or custom-protocol URL must never reach the operating system.
func Check(raw string) error {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" || u.User != nil {
		return ErrRefused
	}
	switch u.Scheme {
	case "https":
		return nil
	case "http":
		switch u.Hostname() {
		case "localhost", "127.0.0.1", "::1":
			return nil
		}
	}
	return ErrRefused
}

// command is the process that opens a URL on a platform. The URL is
// one argument, never a shell string.
func command(goos, target string) (string, []string) {
	switch goos {
	case "windows":
		return "rundll32", []string{"url.dll,FileProtocolHandler", target}
	case "darwin":
		return "open", []string{target}
	default:
		return "xdg-open", []string{target}
	}
}

// start runs a command without waiting for it. Tests replace it.
var start = func(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	go func() { _ = cmd.Wait() }()
	return nil
}

// Open shows the URL in the system browser after Check.
func Open(raw string) error {
	if err := Check(raw); err != nil {
		return err
	}
	name, args := command(runtime.GOOS, raw)
	if err := start(name, args...); err != nil {
		return fmt.Errorf("open the browser: %w", err)
	}
	return nil
}
