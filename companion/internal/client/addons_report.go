package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// InstalledAddon is one Forever Sixty addon folder found under Interface/AddOns, with the version
// its TOC names (empty when the TOC has none).
type InstalledAddon struct {
	Folder  string `json:"folder"`
	Version string `json:"version"`
}

type addonsReport struct {
	Addons []InstalledAddon `json:"addons"`
}

// ReportAddons tells the site which Forever Sixty addons this device found installed, replacing
// its last report: /setup's addon step reads it. An empty list is a report of none, never null.
func (c *Client) ReportAddons(ctx context.Context, addons []InstalledAddon) error {
	if addons == nil {
		addons = []InstalledAddon{}
	}
	body, err := json.Marshal(addonsReport{Addons: addons})
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPut, Path: "/v1/devices/self/addons",
		Body: body, Type: "application/json",
	})
	return err
}
