// api/internal/addon/sync_events.go
//
// The addon sync's notification events (docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md):
// character.linked the first time a character's export arrives, and band.entered when a sync
// carries a character across into a new level range.
package addon

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// levelBands are the level ranges a character enters, lowest first.
var levelBands = []int{20, 30, 40, 50, 60}

// bandOf is the highest band at or below level; 0 below the first.
func bandOf(level int) int {
	band := 0
	for _, b := range levelBands {
		if level >= b {
			band = b
		}
	}
	return band
}

// exportLevel is the level= of an export; 0 when it carries none.
func exportLevel(export string) int {
	d, ok := fs1.Decode(export)
	if !ok || !d.HasLevel {
		return 0
	}
	return d.Level
}

// syncCopy is the words of the sync notifications.
var syncCopy = struct{ LinkedTitle, LinkedBody, BandTitle, BandBody string }{
	LinkedTitle: "%s is linked",
	LinkedBody:  "Your character synced for the first time. See its upgrades.",
	BandTitle:   "%s reached level %d",
	BandBody:    "A new level range: see the BiS list for it.",
}

// SyncEvents is what one sync tells the account: prior is each named character's level before
// the sync (absent: it had no export). Exports the sync did not carry tell nothing.
func SyncEvents(userID int64, prior map[string]int, exports []Export, publicBaseURL string) []events.Event {
	var out []events.Event
	for _, e := range exports {
		key := character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name)
		base := events.Event{
			Audience: events.Users(userID), Link: publicBaseURL + "/character/" + key,
			Data: map[string]string{"character_key": key},
		}
		level := exportLevel(e.Export)
		was, had := prior[key]
		switch {
		case !had:
			ev := base
			ev.Kind, ev.DedupeKey = events.CharacterLinked, "character.linked:"+key
			ev.Title, ev.Body = fmt.Sprintf(syncCopy.LinkedTitle, e.Name), syncCopy.LinkedBody
			out = append(out, ev)
		case was > 0 && bandOf(level) > bandOf(was):
			band := bandOf(level)
			ev := base
			ev.Kind, ev.DedupeKey = events.BandEntered, fmt.Sprintf("band.entered:%s:%d", key, band)
			ev.Title, ev.Body = fmt.Sprintf(syncCopy.BandTitle, e.Name, band), syncCopy.BandBody
			out = append(out, ev)
		}
	}
	return out
}

// priorLevels is each of keys' level in the account's stored export (0 when it has no
// level=); a key with no stored export of the account's is absent.
func (s *Store) priorLevels(ctx context.Context, userID int64, keys []string) (map[string]int, error) {
	rows, err := s.Pool.Query(ctx,
		`select character_key, export from addon_exports where user_id = $1 and character_key = any($2)`, userID, keys)
	if err != nil {
		return nil, fmt.Errorf("addon: prior levels: %w", err)
	}
	defer rows.Close()
	out := map[string]int{}
	for rows.Next() {
		var key, export string
		if err := rows.Scan(&key, &export); err != nil {
			return nil, fmt.Errorf("addon: prior levels: %w", err)
		}
		out[key] = exportLevel(export)
	}
	return out, rows.Err()
}

// exportKeys is each export's character key.
func exportKeys(exports []Export) []string {
	keys := make([]string, len(exports))
	for i, e := range exports {
		keys[i] = character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name)
	}
	return keys
}

// syncNotifier is a sync's notifications, read before the write and told after it.
type syncNotifier struct {
	userID int64
	prior  map[string]int
	ok     bool
}

// beforeSync reads the prior levels; a failed read is logged and tells nothing.
func (s *Service) beforeSync(ctx context.Context, userID int64, exports []Export) syncNotifier {
	if s.Events == nil || len(exports) == 0 {
		return syncNotifier{}
	}
	prior, err := s.Store.priorLevels(ctx, userID, exportKeys(exports))
	if err != nil {
		s.logger().Error("addon", "op", "sync_events", "user", userID, "err", err)
		return syncNotifier{}
	}
	return syncNotifier{userID: userID, prior: prior, ok: true}
}

// afterSync publishes the sync's events once its write stood.
func (s *Service) afterSync(ctx context.Context, n syncNotifier, exports []Export) {
	if !n.ok {
		return
	}
	for _, e := range SyncEvents(n.userID, n.prior, exports, s.PublicBaseURL) {
		events.Quietly(ctx, s.Events, s.logger(), e)
	}
}
