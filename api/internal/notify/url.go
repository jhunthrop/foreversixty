package notify

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

// ErrNotAWebhook is returned for anything that is not a Discord incoming
// webhook URL. It names no part of the URL.
var ErrNotAWebhook = errors.New("notify: not a Discord webhook URL")

const webhookPathPrefix = "/api/webhooks/"

// webhookHosts are the only hosts accepted, exactly (no subdomains).
var webhookHosts = map[string]bool{"discord.com": true, "discordapp.com": true}

var (
	webhookIDPattern    = regexp.MustCompile(`^[0-9]+$`)
	webhookTokenPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)
)

// ParsedURL is the credential pair a webhook URL carries. The URL holds the
// webhook id, not the channel id.
type ParsedURL struct{ WebhookID, Token string }

// URL rebuilds the canonical webhook URL.
func (p ParsedURL) URL() string {
	return "https://discord.com" + webhookPathPrefix + p.WebhookID + "/" + p.Token
}

// ParseWebhookURL accepts exactly https://discord.com or discordapp.com
// /api/webhooks/<digits>/<token>: no userinfo, port, query, fragment or
// extra path.
func ParseWebhookURL(raw string) (ParsedURL, error) {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme != "https" || u.User != nil || u.Port() != "" ||
		!webhookHosts[u.Hostname()] || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" {
		return ParsedURL{}, ErrNotAWebhook
	}
	rest, ok := strings.CutPrefix(u.EscapedPath(), webhookPathPrefix)
	id, token, found := strings.Cut(rest, "/")
	if !ok || !found || !webhookIDPattern.MatchString(id) || !webhookTokenPattern.MatchString(token) {
		return ParsedURL{}, ErrNotAWebhook
	}
	return ParsedURL{WebhookID: id, Token: token}, nil
}
