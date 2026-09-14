// companion/internal/client/devices.go
// Pairing. The claim is the one request the companion sends without a
// token, because it is the request that fetches one.
package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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
