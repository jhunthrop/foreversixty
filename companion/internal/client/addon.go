// companion/internal/client/addon.go
// The two addon routes. Both carry strings the companion does not
// read: FS1 exports going up and addon codes coming down.
package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
)

// PostAddonExports uploads the character exports the addon saved.
func (c *Client) PostAddonExports(ctx context.Context, chars []addon.Export) error {
	if len(chars) == 0 {
		return nil
	}
	body, err := json.Marshal(struct {
		Characters []addon.Export `json:"characters"`
	}{chars})
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/addon/exports",
		Body: body, Type: "application/json",
	})
	return err
}

// AddonInbox fetches the builds the player chose on the site.
func (c *Client) AddonInbox(ctx context.Context) (addon.Inbox, error) {
	var out addon.Inbox
	if _, err := c.do(ctx, request{
		Method: http.MethodGet, Path: "/v1/addon/inbox", Out: &out,
	}); err != nil {
		return addon.Inbox{}, err
	}
	return out, nil
}
