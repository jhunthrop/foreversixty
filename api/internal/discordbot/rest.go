package discordbot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	apiBase        = "https://discord.com/api/v10"
	requestTimeout = 5 * time.Second
	maxAnswerBytes = 1 << 20
	maxErrorRunes  = 200
)

// Discord JSON error codes the bot reacts to.
const (
	codeUnknownChannel = 10003
	codeUnknownGuild   = 10004
	codeUnknownMessage = 10008
	// Unknown Webhook and Invalid Webhook Token: an interaction token is
	// good for 15 minutes.
	codeUnknownWebhook = 10015
	codeInvalidWebhook = 50027
	codeMissingAccess  = 50001
	codeMissingPerms   = 50013
)

// Transport failures, reported without the request (it carries credentials).
var (
	ErrTimeout     = errors.New("discordbot: discord did not answer in time")
	ErrUnreachable = errors.New("discordbot: discord unreachable")
)

// REST is Discord's HTTP API as the bot uses it, behind an interface so
// features test without the network.
type REST interface {
	CreateMessage(ctx context.Context, channelID string, m Message) (string, error)
	EditMessage(ctx context.Context, channelID, messageID string, m Message) error
	ServerLayout(ctx context.Context, serverID string) (ServerLayout, error)
	LeaveServer(ctx context.Context, serverID string) error
	// PutCommands replaces the command set; serverID "" means global.
	PutCommands(ctx context.Context, serverID string, cmds []Command) error
	EditOriginal(ctx context.Context, interactionToken string, m Message) error
	FollowUp(ctx context.Context, interactionToken string, m Message) error
}

// APIError is a non-2xx answer from Discord. It never carries a credential.
type APIError struct {
	Status     int
	Code       int
	RetryAfter time.Duration
	// Global is a 429 from the bot's global limit rather than one route's bucket.
	Global  bool
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("discord: status %d code %d: %s", e.Status, e.Code, e.Message)
}

// RateLimited reports a 429.
func (e *APIError) RateLimited() bool { return e.Status == http.StatusTooManyRequests }

// MissingAccess reports that the bot may not see or use the target.
func (e *APIError) MissingAccess() bool {
	switch e.Code {
	case codeMissingAccess, codeMissingPerms, codeUnknownGuild:
		return true
	}
	return e.Status == http.StatusForbidden
}

// Gone reports that the channel or message no longer exists.
func (e *APIError) Gone() bool {
	switch e.Code {
	case codeUnknownChannel, codeUnknownMessage, codeUnknownWebhook, codeInvalidWebhook:
		return true
	}
	return false
}

// newClient is the 5 s client every Discord call uses. A redirect would carry
// a credential to another host; the 3xx is returned instead.
func newClient() *http.Client {
	return &http.Client{
		Timeout:       requestTimeout,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}
}

// HTTPREST talks to discord.com with the bot token.
type HTTPREST struct {
	// BaseURL overrides the API base; tests point it at httptest.
	BaseURL       string
	applicationID string
	botToken      string
	client        *http.Client
}

// NewHTTPREST returns the production client.
func NewHTTPREST(applicationID, botToken string) *HTTPREST {
	return &HTTPREST{BaseURL: apiBase, applicationID: applicationID, botToken: botToken, client: newClient()}
}

// call is one request: path under the base, whether it carries the bot token,
// the JSON body (nil for none) and where to decode the answer (nil to skip).
type call struct {
	method, path string
	botAuth      bool
	body, out    any
}

func (r *HTTPREST) do(ctx context.Context, c call) error {
	var reader io.Reader
	if c.body != nil {
		raw, err := json.Marshal(c.body)
		if err != nil {
			return fmt.Errorf("discordbot: encode request: %w", err)
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, c.method, r.BaseURL+c.path, reader)
	if err != nil {
		return errors.New("discordbot: bad request")
	}
	if c.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.botAuth {
		req.Header.Set("Authorization", "Bot "+r.botToken)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return transportError(ctx, err)
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return answerError(res)
	}
	if c.out == nil {
		return nil
	}
	if err := json.NewDecoder(io.LimitReader(res.Body, maxAnswerBytes)).Decode(c.out); err != nil {
		return errors.New("discordbot: unreadable answer")
	}
	return nil
}

// transportError drops the net/url error, which embeds the request URL.
func transportError(ctx context.Context, err error) error {
	var t interface{ Timeout() bool }
	if ctx.Err() != nil || (errors.As(err, &t) && t.Timeout()) {
		return fmt.Errorf("%w: %w", ErrTimeout, context.DeadlineExceeded)
	}
	return ErrUnreachable
}

func answerError(res *http.Response) error {
	raw, _ := io.ReadAll(io.LimitReader(res.Body, maxAnswerBytes))
	var body struct {
		Message    string  `json:"message"`
		Code       int     `json:"code"`
		RetryAfter float64 `json:"retry_after"`
		Global     bool    `json:"global"`
	}
	_ = json.Unmarshal(raw, &body)
	msg := body.Message
	if msg == "" {
		msg = http.StatusText(res.StatusCode)
	}
	if runes := []rune(msg); len(runes) > maxErrorRunes {
		msg = string(runes[:maxErrorRunes])
	}
	retry := body.RetryAfter
	if retry == 0 {
		retry, _ = strconv.ParseFloat(res.Header.Get("Retry-After"), 64)
	}
	return &APIError{
		Status: res.StatusCode, Code: body.Code, Message: msg,
		Global:     body.Global || res.Header.Get("X-RateLimit-Global") != "",
		RetryAfter: time.Duration(retry * float64(time.Second)),
	}
}

func (r *HTTPREST) CreateMessage(ctx context.Context, channelID string, m Message) (string, error) {
	var out struct {
		ID string `json:"id"`
	}
	err := r.do(ctx, call{method: http.MethodPost, path: "/channels/" + url.PathEscape(channelID) + "/messages", botAuth: true, body: m, out: &out})
	if err != nil {
		return "", err
	}
	if out.ID == "" {
		return "", errors.New("discordbot: message answered without an id")
	}
	return out.ID, nil
}

func (r *HTTPREST) EditMessage(ctx context.Context, channelID, messageID string, m Message) error {
	path := "/channels/" + url.PathEscape(channelID) + "/messages/" + url.PathEscape(messageID)
	return r.do(ctx, call{method: http.MethodPatch, path: path, botAuth: true, body: m.edit()})
}

func (r *HTTPREST) LeaveServer(ctx context.Context, serverID string) error {
	return r.do(ctx, call{method: http.MethodDelete, path: "/users/@me/guilds/" + url.PathEscape(serverID), botAuth: true})
}

func (r *HTTPREST) PutCommands(ctx context.Context, serverID string, cmds []Command) error {
	path := "/applications/" + url.PathEscape(r.applicationID)
	if serverID != "" {
		path += "/guilds/" + url.PathEscape(serverID)
	}
	if cmds == nil {
		cmds = []Command{}
	}
	return r.do(ctx, call{method: http.MethodPut, path: path + "/commands", botAuth: true, body: cmds})
}

// interactionPath is a webhook path; the interaction token is its credential,
// so these calls carry no bot token.
func (r *HTTPREST) interactionPath(token string) string {
	return "/webhooks/" + url.PathEscape(r.applicationID) + "/" + url.PathEscape(token)
}

func (r *HTTPREST) EditOriginal(ctx context.Context, interactionToken string, m Message) error {
	return r.do(ctx, call{method: http.MethodPatch, path: r.interactionPath(interactionToken) + "/messages/@original", body: m.edit()})
}

func (r *HTTPREST) FollowUp(ctx context.Context, interactionToken string, m Message) error {
	return r.do(ctx, call{method: http.MethodPost, path: r.interactionPath(interactionToken), body: m})
}

// ServerLayout reads the server's channels and roles and the bot's own roles.
// A bot's user id is its application id.
func (r *HTTPREST) ServerLayout(ctx context.Context, serverID string) (ServerLayout, error) {
	base := "/guilds/" + url.PathEscape(serverID)
	layout := ServerLayout{BotUserID: r.applicationID}
	var member struct {
		Roles []string `json:"roles"`
	}
	for _, c := range []call{
		{method: http.MethodGet, path: base + "/channels", botAuth: true, out: &layout.Channels},
		{method: http.MethodGet, path: base + "/roles", botAuth: true, out: &layout.Roles},
		{method: http.MethodGet, path: base + "/members/" + url.PathEscape(r.applicationID), botAuth: true, out: &member},
	} {
		if err := r.do(ctx, c); err != nil {
			return ServerLayout{}, err
		}
	}
	layout.BotRoleIDs = member.Roles
	return layout, nil
}
