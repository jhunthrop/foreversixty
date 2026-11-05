package discordbot

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/db"
)

// A guild event has no roles (events spec §3.5): In signs up at once, with no spec question and
// no Change spec.
func TestAnEventSignUpAsksNoSpec(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	w.exec(t, `update guild_raids set zone = 'event:guild', title = 'Guild meeting' where id = $1`, w.raidID)
	rep := w.press(t, "in")
	if len(rep.Message.Embeds) != 0 || !strings.HasPrefix(rep.Message.Content, "Zulmara is In") {
		t.Fatalf("event press = %+v", rep.Message)
	}
	for _, row := range rep.Message.Components {
		for _, c := range row {
			if c.Label == guideCopy.SpecChange {
				t.Fatal("an event outcome offers Change spec")
			}
		}
	}
}

// A spec the logs already name is known: the press signs up as it, with no question.
func TestALoggedSpecIsNotAskedAgain(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	ctx, at := context.Background(), time.Now().Add(-24*time.Hour)
	if err := db.EnsureMetricsPartition(ctx, w.st.Pool, at); err != nil {
		t.Fatal(err)
	}
	w.exec(t, `insert into reports (id, guild_id, visibility, status, created_at) values ('spec-log', $1, 'guild', 'complete', now())`, w.gid)
	if _, err := w.st.Pool.Exec(ctx, `insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec, role,
		metric_dps, deaths, encounter_id, duration_ms, kill, fought_at)
		values ('spec-log', 0, $1, 'Zulmara', 'paladin', 'Holy', 'healer', 0, 0, 1084, 60000, true, $2)`, keyZulmara, at); err != nil {
		t.Fatal(err)
	}
	rep := w.press(t, "in")
	if len(rep.Message.Embeds) != 0 || !strings.HasPrefix(rep.Message.Content, "Zulmara is In as Healer") {
		t.Fatalf("press = %+v", rep.Message)
	}
}

// Out with two characters, only one of them signed up: it is that one, with no select.
func TestOutTakesTheOneSignedUpCharacter(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
	if rep := w.pick(t, "in", keyZulmara); !strings.HasPrefix(rep.Message.Content, "Zulmara is In") {
		t.Fatalf("in = %+v", rep.Message)
	}
	rep := w.press(t, "out")
	if !strings.HasPrefix(rep.Message.Content, "Zulmara is Out") {
		t.Fatalf("out = %+v, want Zulmara out with no select", rep.Message)
	}
	// With neither signed up, Out still asks which.
	w2 := newPressWorld(t)
	w2.addChar(t, w2.gid, w2.uid, keyZulmara, "Zulmara", true)
	w2.addChar(t, w2.gid, w2.uid, keyDusk, "Duskstrider", true)
	if rep := w2.press(t, "out"); len(rep.Message.Components) == 0 || rep.Message.Components[0][0].Kind != KindStringSelect {
		t.Fatalf("out with nobody signed = %+v, want the select", rep.Message)
	}
}
