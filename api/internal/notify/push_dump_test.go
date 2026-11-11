//go:build payloaddump

package notify

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestDumpPushExamples writes the push notifications built from notify facts (title, body,
// link) to $REVIEW_PAYLOADS_DIR/push.json for the review captures (go test -tags payloaddump).
func TestDumpPushExamples(t *testing.T) {
	dir := os.Getenv("REVIEW_PAYLOADS_DIR")
	if dir == "" {
		t.Skip("REVIEW_PAYLOADS_DIR is not set")
	}
	at := time.Date(2026, time.November, 10, 22, 0, 0, 0, time.UTC)
	site := "https://foreversixty.gg"
	type push struct{ ID, Title, Body, Link string }
	var events []push
	add := func(id, title, body, link string) { events = append(events, push{id, title, body, link}) }
	sim := SimFinishedEvent(site, SimFinished{SimID: "abc123", Spec: "warrior-fury", Character: "Kraggor",
		DPS: 1234.5, Title: "BiS check", CreatedAt: at.Add(-48 * time.Second), FinishedAt: at})
	add("sim", sim.Title, sim.Body, sim.Link)
	headline := SimFinishedEvent(site, SimFinished{SimID: "abc124", Spec: "mage-frost", DPS: 812,
		Headline: "+41 DPS from Mana Igniting Cord", FinishedAt: at})
	add("sim-headline", headline.Title, headline.Body, headline.Link)
	report := UploadParsedEvent(site, UploadParsed{ReportID: "rep1", Title: "Tuesday Molten Core",
		Zone: "Molten Core", Fights: 10, Kills: 9, FinishedAt: at})
	add("report", report.Title, report.Body, report.Link)
	untitled := UploadParsedEvent(site, UploadParsed{ReportID: "rep2", Fights: 1, Kills: 1, FinishedAt: at})
	add("report-untitled", untitled.Title, untitled.Body, untitled.Link)
	out, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "push.json"), append(out, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
}
