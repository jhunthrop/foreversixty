// api/internal/guilds/setup_test.go
package guilds

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

func (w raidWorld) setup(t *testing.T) SetupView {
	t.Helper()
	w.as(w.officer)
	res := w.h.do(http.MethodGet, w.path("/setup"), "")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("setup = %d: %s", res.StatusCode, readBody(res))
	}
	var v SetupView
	w.h.data(res, &v)
	return v
}

func (w raidWorld) mark(t *testing.T, step, state string, want int) SetupView {
	t.Helper()
	w.as(w.officer)
	method, body := http.MethodPut, `{"state":"`+state+`"}`
	if state == "" {
		method, body = http.MethodDelete, ""
	}
	res := w.h.do(method, w.path("/setup/%s", step), body)
	if res.StatusCode != want {
		t.Fatalf("%s %s = %d: %s", step, state, res.StatusCode, readBody(res))
	}
	var v SetupView
	if want == http.StatusOK {
		w.h.data(res, &v)
	}
	return v
}

func stateOf(v SetupView, step string) string {
	for _, s := range v.Steps {
		if s.Key == step {
			return s.State
		}
	}
	return ""
}

func TestSetupStartsAtToDoAndIsOfficerOnly(t *testing.T) {
	w := newRaidWorld(t, false)
	v := w.setup(t)
	if v.Total != len(SetupSteps) || len(v.Steps) != len(SetupSteps) || v.Finished || v.Plan != planNone {
		t.Fatalf("fresh = %+v", v)
	}
	for i, s := range v.Steps {
		if s.Key != SetupSteps[i] {
			t.Fatalf("step %d = %s, want %s", i, s.Key, SetupSteps[i])
		}
	}
	for _, step := range []string{StepOfficerRanks, StepRecruiting, StepRaidSchedule, StepInviteLink, StepReportVisibility, StepLinkedRole} {
		if got := stateOf(v, step); got != SetupTodo {
			t.Errorf("%s = %s, want todo", step, got)
		}
	}
	// Every seeded raider has a verified character and none has synced an export yet.
	if v.Sync.Accounts != 3 || v.Sync.Synced != 0 || len(v.Sync.Missing) != 3 || stateOf(v, StepAddonSync) != SetupTodo {
		t.Fatalf("sync = %+v", v.Sync)
	}
	w.as(w.member)
	if res := w.h.do(http.MethodGet, w.path("/setup"), ""); res.StatusCode != http.StatusForbidden {
		t.Fatalf("member = %d", res.StatusCode)
	}
}

func TestSetupDerivesFromTheGuildsData(t *testing.T) {
	w := newRaidWorld(t, true)
	w.exec(t, fmt.Sprintf(`update guilds set officer_max_rank_index = 2, default_visibility = 'guild',
		raid_timezone = 'America/Chicago', invite_token_rotated_at = now() where id = %d`, w.gid))
	w.exec(t, fmt.Sprintf(`update guild_characters set rank_index = 0 where guild_id = %d and rank = 'officer'`, w.gid))
	w.exec(t, fmt.Sprintf(`update guild_characters set rank_index = 3 where guild_id = %d and rank = 'member'`, w.gid))
	w.exec(t, fmt.Sprintf(`update guild_characters set source = 'invite' where guild_id = %d and user_id = %d`, w.gid, w.member2))
	w.createRaid(t)
	// Three of rank 3's characters are level 60 and one is 58: the rank line counts the three.
	w.exec(t, fmt.Sprintf(`insert into characters (key, region, ruleset, name, level)
		select gc.character_key, 'us', 'hardcore', gc.character_key,
		       case when row_number() over (order by gc.character_key) = 1 then 58 else 60 end
		from guild_characters gc where gc.guild_id = %d and gc.rank_index = 3
		on conflict (key) do update set level = excluded.level`, w.gid))
	v := w.setup(t)
	for _, step := range []string{StepOfficerRanks, StepRaidSchedule, StepInviteLink, StepReportVisibility} {
		if got := stateOf(v, step); got != SetupDone {
			t.Errorf("%s = %s, want done", step, got)
		}
	}
	if v.Plan != planRaidGuild || v.Schedule.Planned != 1 || v.Schedule.Next == nil || v.Schedule.Next.PostedAt != nil ||
		v.Invite.Joined != 1 {
		t.Fatalf("facts = plan %s, schedule %+v, invite %+v", v.Plan, v.Schedule, v.Invite)
	}
	if v.Readiness != nil {
		t.Fatalf("a Raid guild gets no preview: %+v", v.Readiness)
	}
	if len(v.Ranks.Ranks) != 2 || v.Ranks.Ranks[0].RankIndex != 0 || v.Ranks.Ranks[0].Members != 1 ||
		v.Ranks.Ranks[1].Members != 4 || v.Ranks.Ranks[1].AtSixty != 3 ||
		v.Ranks.Officers != 1 {
		t.Fatalf("ranks = %+v", v.Ranks)
	}
}

// Without the bot in the deployment, a connected webhook is the Discord step's done.
func TestSetupDiscordStepReadsTheWebhookWithoutTheBot(t *testing.T) {
	w := newRaidWorld(t, false)
	w.h.svc.Discord = fakeDiscord{enabled: true, hook: &notify.GuildWebhook{ChannelID: "1"}}
	v := w.setup(t)
	if !v.Discord.WebhookConnected || stateOf(v, StepDiscordBot) != SetupDone {
		t.Fatalf("discord = %+v, state %s", v.Discord, stateOf(v, StepDiscordBot))
	}
}

func TestSetupMarksSkipConfirmAndUndo(t *testing.T) {
	w := newRaidWorld(t, false)
	if v := w.mark(t, StepLinkedRole, SetupSkipped, http.StatusOK); stateOf(v, StepLinkedRole) != SetupSkipped {
		t.Fatalf("skip = %+v", v.Steps)
	}
	if v := w.mark(t, StepOfficerRanks, SetupDone, http.StatusOK); stateOf(v, StepOfficerRanks) != SetupDone || v.Done != 1 {
		t.Fatalf("confirm = %+v", v)
	}
	// A derived-only step takes no done mark; an unknown step none at all.
	w.mark(t, StepAddonSync, SetupDone, http.StatusBadRequest)
	w.mark(t, "dragons", SetupSkipped, http.StatusBadRequest)
	w.mark(t, StepRecruiting, "maybe", http.StatusBadRequest)
	if v := w.mark(t, StepOfficerRanks, "", http.StatusOK); stateOf(v, StepOfficerRanks) != SetupTodo {
		t.Fatalf("undo = %+v", v.Steps)
	}
	// Skipping everything finishes the checklist with nothing done.
	var v SetupView
	for _, step := range SetupSteps {
		v = w.mark(t, step, SetupSkipped, http.StatusOK)
	}
	if !v.Finished || v.Done != 0 {
		t.Fatalf("all skipped = %+v", v)
	}
	w.setContested(t, true)
	w.mark(t, StepLinkedRole, SetupDone, http.StatusConflict)
	w.as(w.member)
	if res := w.h.do(http.MethodPut, w.path("/setup/%s", StepLinkedRole), `{"state":"done"}`); res.StatusCode != http.StatusForbidden {
		t.Fatalf("member mark = %d", res.StatusCode)
	}
}

func TestSpecClassName(t *testing.T) {
	holy, resto := "Holy", "Restoration Druid"
	for _, c := range []struct {
		spec  *string
		class string
		want  string
	}{
		{&holy, "priest", "Holy Priest"},
		{&resto, "druid", "Restoration Druid"},
		{nil, "warrior", "Warrior"},
		{nil, "", ""},
	} {
		if got := specClassName(c.spec, c.class); got != c.want {
			t.Errorf("specClassName(%v, %q) = %q, want %q", c.spec, c.class, got, c.want)
		}
	}
}

// A guild without Raid guild sees what the readiness board would flag for its next raid.
func TestSetupPreviewsReadinessWithoutRaidGuild(t *testing.T) {
	w := newRaidWorld(t, false)
	raid := w.createRaid(t)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"in"}`)
	v := w.setup(t)
	if v.Readiness == nil || v.Readiness.Signed != 1 || v.Readiness.RaidTitle == "" || v.Readiness.Names == nil {
		t.Fatalf("preview = %+v", v.Readiness)
	}
}

// The visibility step reads the guild's upload default, never a report: a private default is
// done, the column default is to do, and the view carries the setting alone.
func TestSetupVisibilityStepReadsTheGuildDefaultNotAReport(t *testing.T) {
	for _, tc := range []struct {
		visibility string
		done       bool
	}{{"private", true}, {"guild", true}, {defaultVisibility, false}} {
		v := SetupView{DefaultVisibility: tc.visibility}
		if got := v.derivedDone()[StepReportVisibility]; got != tc.done {
			t.Errorf("default %q: done = %v, want %v", tc.visibility, got, tc.done)
		}
	}
}
