// api/internal/bnetapi/probe.go
package bnetapi

import (
	"context"
	"io"
	"net/http"
)

// ProbeResult is one game's answer to the namespace probe (spec §5).
type ProbeResult struct {
	Game   string
	Status int
}

// Probe requests GET /data/wow/realm/index?namespace=dynamic-{game}-{region}
// for each of games, in order, and reports each one's HTTP status. A
// request that fails outright (no response at all — a network error)
// reports status 0. Nothing here decides what a status means; the
// refresh job logs one line per game and compares against the
// configured game (spec §5).
func (c *Client) Probe(ctx context.Context, region string, games []string) []ProbeResult {
	out := make([]ProbeResult, 0, len(games))
	for _, game := range games {
		u := c.APIHost(region) + "/data/wow/realm/index?namespace=dynamic-" + game + "-" + region
		out = append(out, ProbeResult{Game: game, Status: c.probeStatus(ctx, u)})
	}
	return out
}

func (c *Client) probeStatus(ctx context.Context, rawURL string) int {
	token, err := c.AppToken(ctx)
	if err != nil {
		return 0
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return 0
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := c.do(req)
	if err != nil {
		return 0
	}
	defer res.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(res.Body, maxResponseBytes))
	return res.StatusCode
}
