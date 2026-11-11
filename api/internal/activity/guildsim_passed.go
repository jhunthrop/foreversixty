package activity

// guildsim.passed (spec 2026-11-08-app-flows F4): when a member's finished sim moves them above
// others on their guild's ladder (the group the sim is in: role, band and preset), each member
// passed hears it. A move needs a place to move from: the member's previous finished sim in the
// group within LadderMaxAge. Passed is a member whose latest ladder figure sat at or above that
// previous figure and is now below the new one; the member's own other characters never count,
// and a ladder below ladderMin members (hidden, ruling R2) tells nobody.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/events"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

// passedCopy is the words of guildsim.passed.
// Body is the push's: the passer's figure, then Lead; the Activity row shows the figure itself
// and keeps Lead alone.
var passedCopy = struct{ Title, Body, Lead string }{
	Title: "%s passed you in the guild ladder",
	Body:  "%s %s, %s",
	Lead:  "+%s over your %s (%s, %s).",
}

// guildLinkFormat is the app link the event lands on: the Raid tab for that guild, where the
// guild sims ladder lives.
const guildLinkFormat = "%s/app/guild/%d"

// previousSimSQL is the character's latest finished sim in the group before simID, within the
// ladder's age.
const previousSimSQL = `select sm.dps_mean from character_sims cs
	join sims sm on sm.id = cs.sim_id and sm.state = 'done'
	where cs.character_key = $1 and cs.sim_id <> $2 and cs.role = $3 and cs.band = $4 and cs.preset = $5
	  and coalesce(sm.finished_at, sm.created_at) > $6 and coalesce(sm.finished_at, sm.created_at) <= $7
	order by coalesce(sm.finished_at, sm.created_at) desc limit 1`

// publishPassed publishes guildsim.passed to each member f's sim passed on ladder.
func (w *GuildSimWriter) publishPassed(ctx context.Context, f simFacts, simID string, ladder []LadderRow) error {
	if w.Events == nil || len(ladder) < ladderMin {
		return nil
	}
	var prev float64
	err := w.Store.Pool.QueryRow(ctx, previousSimSQL, f.key, simID, f.role, f.band, f.preset,
		w.now().Add(-LadderMaxAge), f.at).Scan(&prev)
	if isNoRows(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("previous sim: %w", err)
	}
	for _, r := range passedOn(ladder, f.key, prev, f.value) {
		events.Quietly(ctx, w.Events, w.logger(), w.passedEvent(f, simID, r))
	}
	return nil
}

// passedOn is the ladder rows key's move from prev to now passed: another account's figure at
// or above prev and below now.
func passedOn(ladder []LadderRow, key string, prev, now float64) []LadderRow {
	var self int64
	for _, r := range ladder {
		if r.Key == key {
			self = r.userID
		}
	}
	var out []LadderRow
	for _, r := range ladder {
		if r.Key != key && r.userID != self && r.Value >= prev && r.Value < now {
			out = append(out, r)
		}
	}
	return out
}

// passedEvent is guildsim.passed to r's account.
func (w *GuildSimWriter) passedEvent(f simFacts, simID string, r LadderRow) events.Event {
	_, unit := RoleWords(f.role)
	words := BandWords(f.band)
	if f.preset != barePreset {
		words += " " + f.preset
	}
	segments := Segmentf(passedCopy.Lead, Strong(textfmt.Figure(f.value-r.Value, f.value)),
		Text(textfmt.Figure(r.Value, r.Value)), Text(SpecName(f.spec)), Text(words))
	lead := events.SegmentsText(segments)
	value := f.value
	return events.Event{
		Kind: events.GuildSimPassed, Audience: events.Users(r.userID),
		DedupeKey: "guildsim.passed:" + simID + ":" + r.Key,
		Title:     fmt.Sprintf(passedCopy.Title, f.name),
		Body:      fmt.Sprintf(passedCopy.Body, textfmt.Figure(f.value, f.value), unit, lead), ActivityBody: &lead,
		Segments: segments, Crest: f.class, Figure: &value,
		Link: fmt.Sprintf(guildLinkFormat, w.PublicBaseURL, f.guildID), Icon: specs.ByKey[f.spec].Icon,
		Data: map[string]string{
			"guild_id": strconv.FormatInt(f.guildID, 10), "sim_id": simID, "character_key": f.key,
			"role": f.role, "band": strconv.Itoa(f.band), "preset": f.preset,
		},
	}
}
