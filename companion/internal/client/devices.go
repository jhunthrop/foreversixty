// companion/internal/client/devices.go
// Pairing. The claim is the one request the companion sends without a
// token, because it is the request that fetches one.
package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// Claim is the body of POST /v1/devices/claim.
type Claim struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// Device is the once-shown answer to a claim.
type Device struct {
	DeviceID string `json:"device_id"`
	Token    string `json:"token"`
	// AccountLabel is who the device now belongs to (battletag, else
	// email); empty from an older API.
	AccountLabel string `json:"account_label,omitempty"`
}

// TokenPrefix is what a device token starts with, per the contract.
const TokenPrefix = "fsd_"

// Claim exchanges a pairing code shown on the site for a device token.
// The token is returned once and never again, so the caller must store
// it before doing anything else.
func (c *Client) Claim(ctx context.Context, in Claim) (Device, error) {
	if in.Code == "" {
		return Device{}, errors.New("the pairing code is empty")
	}
	body, err := json.Marshal(in)
	if err != nil {
		return Device{}, err
	}
	var out Device
	if _, err := c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/devices/claim",
		Body: body, Type: "application/json", Anonymous: true, Out: &out,
	}); err != nil {
		return Device{}, err
	}
	if !strings.HasPrefix(out.Token, TokenPrefix) {
		return Device{}, errors.New("the API returned a token that is not a device token")
	}
	return out, nil
}

// ConnectRequest is the body of POST /v1/devices/connect.
type ConnectRequest struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// Connect is a started device authorization: what the companion keeps
// (ConnectID, which is the secret), what the player compares (UserCode)
// and where the browser goes (URL).
type Connect struct {
	ConnectID string `json:"connect_id"`
	UserCode  string `json:"user_code"`
	URL       string `json:"url"`
	// ExpiresIn is the lifetime in seconds, Interval the poll
	// period in seconds.
	ExpiresIn int `json:"expires_in"`
	Interval  int `json:"interval"`
}

// The answers a connect poll can carry.
const (
	ConnectPending  = "pending"
	ConnectApproved = "approved"
	ConnectDenied   = "denied"
	ConnectExpired  = "expired"
	// ConnectUsed is the 410 after the approved answer was collected
	// once; the token is gone.
	ConnectUsed = "used"
)

// ConnectState is one poll's answer. Token and DeviceID are set only
// with ConnectApproved, and only once.
type ConnectState struct {
	Status       string `json:"status"`
	DeviceID     string `json:"device_id,omitempty"`
	Token        string `json:"token,omitempty"`
	AccountLabel string `json:"account_label,omitempty"`
}

// Connect starts a sign-in through the site. Like Claim it carries no
// token: the device does not have one yet.
func (c *Client) Connect(ctx context.Context, in ConnectRequest) (Connect, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return Connect{}, err
	}
	var out Connect
	if _, err := c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/devices/connect",
		Body: body, Type: "application/json", Anonymous: true, Out: &out,
	}); err != nil {
		return Connect{}, err
	}
	if out.ConnectID == "" || out.UserCode == "" || out.URL == "" {
		return Connect{}, errors.New("the API answered a connect without an id, a code and a link")
	}
	if out.Interval <= 0 {
		out.Interval = 2
	}
	return out, nil
}

// ConnectStatus polls a started connect. The id is the credential, so
// the request is anonymous. A 404 is reported as expired (the row is
// gone) and a 410 as used, rather than as errors: both are answers.
// An approved answer without a device token is refused here, so a
// caller never stores a token it cannot use.
func (c *Client) ConnectStatus(ctx context.Context, connectID string) (ConnectState, error) {
	if connectID == "" {
		return ConnectState{}, errors.New("the connect id is empty")
	}
	var out ConnectState
	_, err := c.do(ctx, request{
		Method: http.MethodGet, Path: "/v1/devices/connect/" + url.PathEscape(connectID),
		Anonymous: true, Out: &out,
	})
	var ae *Error
	if errors.As(err, &ae) {
		switch ae.Status {
		case http.StatusNotFound:
			return ConnectState{Status: ConnectExpired}, nil
		case http.StatusGone:
			return ConnectState{Status: ConnectUsed}, nil
		}
	}
	if err != nil {
		return ConnectState{}, err
	}
	if out.Status == ConnectApproved && !strings.HasPrefix(out.Token, TokenPrefix) {
		return ConnectState{}, errors.New("the API returned a token that is not a device token")
	}
	return out, nil
}
