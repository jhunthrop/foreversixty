package discordbot

import (
	"context"
	"net/url"
	"sync"
)

// RESTCall is one recorded FakeREST call.
type RESTCall struct {
	Method, ServerID, ChannelID, MessageID, Token string
	Message                                       Message
	Commands                                      []Command
}

// FakeREST records calls and returns scripted answers.
type FakeREST struct {
	mu sync.Mutex
	// Err is returned by every call; ErrFor overrides it per method name.
	Err    error
	ErrFor map[string]error
	// MessageID is what CreateMessage returns (default "m1").
	MessageID string
	// ReplyID is what CreateMessage returns for a reply (default MessageID's answer).
	ReplyID string
	// Latest is what LatestMessageID returns.
	Latest string
	Layout ServerLayout
	// OnCall, when set, runs inside every call before it answers (a test acts while
	// Discord is being called).
	OnCall func(RESTCall)
	calls  []RESTCall
}

func (f *FakeREST) record(c RESTCall) error {
	if f.OnCall != nil {
		f.OnCall(c)
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, c)
	if err, ok := f.ErrFor[c.Method]; ok {
		return err
	}
	return f.Err
}

// Calls returns a copy of the recorded calls.
func (f *FakeREST) Calls() []RESTCall {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]RESTCall(nil), f.calls...)
}

func (f *FakeREST) CreateMessage(_ context.Context, channelID string, m Message) (string, error) {
	if err := f.record(RESTCall{Method: "CreateMessage", ChannelID: channelID, Message: m}); err != nil {
		return "", err
	}
	if m.ReplyTo != "" && f.ReplyID != "" {
		return f.ReplyID, nil
	}
	if f.MessageID == "" {
		return "m1", nil
	}
	return f.MessageID, nil
}

func (f *FakeREST) EditMessage(_ context.Context, channelID, messageID string, m Message) error {
	return f.record(RESTCall{Method: "EditMessage", ChannelID: channelID, MessageID: messageID, Message: m})
}

func (f *FakeREST) LatestMessageID(_ context.Context, channelID string) (string, error) {
	err := f.record(RESTCall{Method: "LatestMessageID", ChannelID: channelID})
	return f.Latest, err
}

// DeleteMessage keeps the REST contract: an already deleted message (10008) is success.
func (f *FakeREST) DeleteMessage(_ context.Context, channelID, messageID string) error {
	return alreadyGone(f.record(RESTCall{Method: "DeleteMessage", ChannelID: channelID, MessageID: messageID}))
}

func (f *FakeREST) ServerLayout(_ context.Context, serverID string) (ServerLayout, error) {
	err := f.record(RESTCall{Method: "ServerLayout", ServerID: serverID})
	return f.Layout, err
}

func (f *FakeREST) LeaveServer(_ context.Context, serverID string) error {
	return f.record(RESTCall{Method: "LeaveServer", ServerID: serverID})
}

func (f *FakeREST) PutCommands(_ context.Context, serverID string, cmds []Command) error {
	return f.record(RESTCall{Method: "PutCommands", ServerID: serverID, Commands: cmds})
}

func (f *FakeREST) EditOriginal(_ context.Context, token string, m Message) error {
	return f.record(RESTCall{Method: "EditOriginal", Token: token, Message: m})
}

func (f *FakeREST) FollowUp(_ context.Context, token string, m Message) error {
	return f.record(RESTCall{Method: "FollowUp", Token: token, Message: m})
}

// FakeOAuth records codes and returns scripted answers.
type FakeOAuth struct {
	mu          sync.Mutex
	Installed   InstalledServer
	Who         Identity
	InstallErr  error
	IdentityErr error
	installs    []string
	identities  []string
}

func (f *FakeOAuth) InstallURL(state string) string {
	return "https://discord.test/install?state=" + url.QueryEscape(state)
}

func (f *FakeOAuth) LinkURL(state string) string {
	return "https://discord.test/link?state=" + url.QueryEscape(state)
}

func (f *FakeOAuth) ExchangeInstall(_ context.Context, code string) (InstalledServer, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.installs = append(f.installs, code)
	return f.Installed, f.InstallErr
}

func (f *FakeOAuth) ExchangeIdentity(_ context.Context, code string) (Identity, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.identities = append(f.identities, code)
	return f.Who, f.IdentityErr
}

// InstallCodes returns the codes ExchangeInstall was called with.
func (f *FakeOAuth) InstallCodes() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.installs...)
}

// IdentityCodes returns the codes ExchangeIdentity was called with.
func (f *FakeOAuth) IdentityCodes() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.identities...)
}

var (
	_ REST  = (*FakeREST)(nil)
	_ OAuth = (*FakeOAuth)(nil)
)
