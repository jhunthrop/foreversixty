// companion/internal/client/addon.go
// The two addon routes. Both carry strings the companion does not
// read: FS1 exports and the addon's outbox going up, addon codes and
// inbox messages coming down.
package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
)

// PostAddonExports uploads the character exports the addon saved and,
// when outbox is not nil, the addon's outbox in the same request
// (docs/contracts/addon-v2.md contract 4). The outbox is never sent
// alone: with no characters there is nothing to send.
func (c *Client) PostAddonExports(ctx context.Context, chars []addon.Export, outbox *addon.Outbox) (addon.ExportsResult, error) {
	if len(chars) == 0 {
		return addon.ExportsResult{}, nil
	}
	in := struct {
		Characters []addon.Export  `json:"characters"`
		Outbox     json.RawMessage `json:"outbox,omitempty"`
	}{Characters: chars}
	if outbox != nil {
		in.Outbox = outbox.Body
	}
	body, err := json.Marshal(in)
	if err != nil {
		return addon.ExportsResult{}, err
	}
	var out addon.ExportsResult
	if _, err := c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/addon/exports",
		Body: body, Type: "application/json", Out: &out,
	}); err != nil {
		return addon.ExportsResult{}, err
	}
	return out, nil
}

// AddonInbox fetches the inbox v2 body: the builds the player chose on
// the site and every typed message (addon.Inbox decodes it message by
// message).
func (c *Client) AddonInbox(ctx context.Context) (addon.Inbox, error) {
	var out addon.Inbox
	if _, err := c.do(ctx, request{
		Method: http.MethodGet, Path: "/v1/addon/inbox", Out: &out,
	}); err != nil {
		return addon.Inbox{}, err
	}
	return out, nil
}
