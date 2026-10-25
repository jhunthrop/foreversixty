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

const canonicalWebhookPrefix = "https://discord.com/api/webhooks/"

// webhookHosts are the only hosts accepted, exactly (no subdomains).
var webhookHosts = map[string]bool{"discord.com": true, "discordapp.com": true}

// webhookPathPattern accepts the unversioned path and the API versions 6-10
// that Discord has handed out in webhook URLs.
var webhookPathPattern = regexp.MustCompile(`^/api(?:/v(?:6|7|8|9|10))?/webhooks/([0-9]+)/([A-Za-z0-9_-]+)$`)

// ParsedURL is the credential pair a webhook URL carries. The URL holds the
// webhook id, not the channel id.
type ParsedURL struct{ WebhookID, Token string }

// URL rebuilds the canonical webhook URL.
func (p ParsedURL) URL() string {
	return canonicalWebhookPrefix + p.WebhookID + "/" + p.Token
}

// ParseWebhookURL accepts exactly https://discord.com or discordapp.com
// /api[/v6..v10]/webhooks/<digits>/<token>: no userinfo, port, query,
// fragment or extra path. The result is canonicalised by ParsedURL.URL.
func ParseWebhookURL(raw string) (ParsedURL, error) {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || u.Scheme != "https" || u.User != nil || u.Port() != "" ||
		!webhookHosts[strings.ToLower(u.Hostname())] || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" {
		return ParsedURL{}, ErrNotAWebhook
	}
	m := webhookPathPattern.FindStringSubmatch(u.EscapedPath())
	if m == nil {
		return ParsedURL{}, ErrNotAWebhook
	}
	return ParsedURL{WebhookID: m[1], Token: m[2]}, nil
}
