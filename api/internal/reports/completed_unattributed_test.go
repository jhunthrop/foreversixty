package reports

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// Review Focus 1 through the completion path: an owner with no guild row, and an owner whose
// only row is an unverified logging character, give the report no guild, match no raid and log
// no error.
func TestFinishCompletionLeavesAnUnattributableOwnerAlone(t *testing.T) {
	verified := true
	for name, verifiedRow := range map[string]*bool{
		"no guild row": nil, "unverified logger only": new(bool), "control: verified logger": &verified,
	} {
		t.Run(name, func(t *testing.T) {
			h := newHarness(t)
			g := h.seedGuild("Alpha", "hardcore")
			if verifiedRow != nil {
				h.seedGuildCharacter(g, h.owner, harnessLoggingKey, *verifiedRow)
			}
			id := h.createReport(Public)
			var raidID int64
			if err := h.store.Pool.QueryRow(t.Context(), `
				insert into guild_raids (guild_id, zone, title, starts_at, duration_minutes, status, created_by)
				values ($1, 'raid:onyxias-lair', 'Night', now() - interval '1 hour', 180, 'planned', $2) returning id`,
				g, h.owner).Scan(&raidID); err != nil {
				t.Fatal(err)
			}
			if err := h.store.SetStatus(t.Context(), id, StatusComplete, "0.1.0", nil); err != nil {
				t.Fatal(err)
			}
			var logs bytes.Buffer
			log := slog.New(slog.NewTextHandler(&logs, nil))
			hook := &guilds.Service{Store: h.guilds, Log: log}
			FinishCompletion(t.Context(), h.store, hook, log, id, "Onyxia's Lair")

			var held *string
			if err := h.store.Pool.QueryRow(t.Context(), `select report_id from guild_raids where id = $1`, raidID).
				Scan(&held); err != nil {
				t.Fatal(err)
			}
			if verifiedRow != nil && *verifiedRow {
				// The control: the same night with a verified logger is attributed and matched.
				if !sameGuild(h.reportMatchedGuild(id), g) || held == nil || *held != id {
					t.Fatalf("control: matched guild %v raid holds %v, want %d and %s", h.reportMatchedGuild(id), held, g, id)
				}
				return
			}
			if h.reportGuild(id) != nil || h.reportMatchedGuild(id) != nil || held != nil {
				t.Fatalf("guild %v matched guild %v raid holds %v, want none", h.reportGuild(id), h.reportMatchedGuild(id), held)
			}
			if strings.Contains(logs.String(), "level=ERROR") {
				t.Fatalf("completion logged an error: %s", logs.String())
			}
		})
	}
}
