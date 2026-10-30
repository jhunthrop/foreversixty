package discordbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func TestSpecCopyMatchesTheSpec(t *testing.T) {
	g := guideCopy
	for key, want := range map[string]string{
		"spec.title": g.SpecTitle, "spec.lead": g.SpecLead, "spec.lead.change": g.SpecLeadChange,
		"spec.placeholder": g.SpecPlaceholder, "spec.change": g.SpecChange, "spec.stale": g.SpecStale,
	} {
		if got := guideString(t, "`"+key+"`", 0); got != want {
			t.Errorf("%s = %q, spec %q", key, want, got)
		}
	}
	if got := guideString(t, "`spec.kept`", 0); strings.TrimSpace(got) != strings.TrimSpace(g.SpecKept) || !strings.HasPrefix(g.SpecKept, " ") {
		t.Errorf("spec.kept = %q, spec %q", g.SpecKept, got)
	}
	if g.SpecPlaceholder != g.BisPlaceholderSpec {
		t.Error("spec.placeholder is shared with bis.placeholder.spec")
	}
}

// specIDs are the spec card's custom ids for the press world's raid and key.
func specPickID(w *pressWorld, state, key string) string {
	return fmt.Sprintf("gspc:1:pick:%d:%s:%s", w.raidID, state, key)
}

func (w *pressWorld) specPress(t *testing.T, id string, values ...string) (Reply, error) {
	t.Helper()
	data := map[string]any{"custom_id": id}
	if len(values) > 0 {
		data["values"] = values
	}
	return w.svc.components[prefixSpec](context.Background(), interactionOf(typeComponent, pressDiscord, data))
}

func (w *pressWorld) mustSpecPress(t *testing.T, id string, values ...string) Reply {
	t.Helper()
	rep, err := w.specPress(t, id, values...)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func (w *pressWorld) raidSpecOf(t *testing.T, key string) string {
	t.Helper()
	var spec *string
	if err := w.st.Pool.QueryRow(context.Background(), `select raid_spec from guild_characters where character_key = $1`, key).Scan(&spec); err != nil {
		t.Fatal(err)
	}
	if spec == nil {
		return ""
	}
	return *spec
}

func (w *pressWorld) signupRow(t *testing.T, key string) (state, role, spec string) {
	t.Helper()
	if err := w.st.Pool.QueryRow(context.Background(), `select state, role, spec from guild_raid_signups where raid_id = $1 and character_key = $2`,
		w.raidID, key).Scan(&state, &role, &spec); err != nil {
		t.Fatal(err)
	}
	return state, role, spec
}

// paladin makes Zulmara a paladin with no spec picked.
func paladin(t *testing.T, w *pressWorld) {
	t.Helper()
	w.addCharNoSpec(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.exec(t, `update characters set class = 'paladin' where key = $1`, keyZulmara)
}

func TestAnInWithNoSpecAsksWhichSpec(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	rep := w.press(t, "in")
	m := rep.Message
	if rep.Update || !m.Ephemeral || len(m.Embeds) != 1 || len(m.Components) != 1 || w.signups(t) != 0 {
		t.Fatalf("reply = %+v, signups %d", rep, w.signups(t))
	}
	e := m.Embeds[0]
	lead := fill(guideCopy.SpecLead, "{Name}", "Zulmara", "{State}", "In", "{Title}", "Onyxia's Lair", "{Date}", w.day())
	if e.Title != guideCopy.SpecTitle || e.Description != lead || e.Color != classColorOf("paladin") ||
		e.Thumbnail != pressPublic+"/icons/hd/crests/paladin.webp" || e.Footer != cardFooter || e.Timestamp.IsZero() || e.Author != "" {
		t.Fatalf("embed = %+v", e)
	}
	sel := m.Components[0][0]
	if sel.CustomID != specPickID(w, "in", keyZulmara) || sel.Placeholder != guideCopy.SpecPlaceholder || len(m.Components[0]) != 1 {
		t.Fatalf("select = %+v", sel)
	}
	var got []string
	for _, o := range sel.Options {
		got = append(got, o.Label+"/"+o.Description+"/"+o.Value)
		if o.Default {
			t.Fatalf("a default with no spec saved: %+v", o)
		}
	}
	if strings.Join(got, ",") != "Holy/Healer/paladin-holy,Protection/Tank/paladin-protection,Retribution/DPS/paladin-retribution" {
		t.Fatalf("options %v", got)
	}
	assertCardLimits(t, m)
}

func TestOutNeverAsksAndAKnownSpecSignsUpAtOnce(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	if rep := w.press(t, "out"); len(rep.Message.Embeds) != 0 || w.signups(t) != 1 {
		t.Fatalf("out = %+v", rep)
	}
	w.exec(t, `update guild_characters set raid_spec = 'paladin-protection' where character_key = $1`, keyZulmara)
	rep := w.press(t, "in")
	if !strings.HasPrefix(rep.Message.Content, "Zulmara is In as Tank") {
		t.Fatalf("in = %q", rep.Message.Content)
	}
}

// After the character select the card answers in place.
func TestTheSpecCardReplacesTheCharacterSelect(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
	rep := w.pick(t, "late", keyZulmara)
	if !rep.Update || len(rep.Message.Embeds) != 1 || rep.Message.Components[0][0].CustomID != specPickID(w, "late", keyZulmara) {
		t.Fatalf("pick = %+v", rep)
	}
}

func TestTheSpecPickSavesAndSignsUpWithItsRole(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	rep := w.mustSpecPress(t, specPickID(w, "in", keyZulmara), "paladin-holy")
	want := "Zulmara is In as Healer for Onyxia's Lair, " + w.day() + ". 1 of 40 signed up. Holy is kept for Zulmara."
	wantReply(t, rep, want, true, changeSpecButton(w, keyZulmara))
	if w.raidSpecOf(t, keyZulmara) != "paladin-holy" {
		t.Fatal("raid_spec not saved")
	}
	if state, role, spec := w.signupRow(t, keyZulmara); state != "in" || role != "healer" || spec != "Holy" {
		t.Fatalf("row %s %s %s", state, role, spec)
	}
	// The next press asks nothing.
	w.exec(t, `delete from guild_raid_signups`)
	if rep := w.press(t, "bench"); !strings.Contains(rep.Message.Content, "on the Bench") {
		t.Fatalf("next press = %+v", rep)
	}
}

func TestTheSpecPickRefusesForgeriesAndTheLockedRaid(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	other := seedUser(t, w.st.Pool)
	w.addChar(t, w.gid, other, keyDusk, "Duskstrider", true)
	cases := []struct {
		name, id, value, want string
	}{
		{"another account's key", specPickID(w, "in", keyDusk), "paladin-holy", pressCopy.Failure},
		{"another class's spec", specPickID(w, "in", keyZulmara), "mage-frost", pressCopy.Failure},
		{"out in the id", specPickID(w, "out", keyZulmara), "paladin-holy", pressCopy.Failure},
		{"garbled", "gspc:1:pick:x:in:" + keyZulmara, "paladin-holy", pressCopy.Failure},
	}
	for _, c := range cases {
		wantReply(t, w.mustSpecPress(t, c.id, c.value), c.want, true)
	}
	if w.signups(t) != 0 || w.raidSpecOf(t, keyZulmara) != "" {
		t.Fatal("a forged pick saved")
	}
	w.exec(t, `update guild_raids set status = 'locked' where id = $1`, w.raidID)
	wantReply(t, w.mustSpecPress(t, specPickID(w, "in", keyZulmara), "paladin-holy"), spec(t, "Press: refusals", 0, nil), true)
	if w.raidSpecOf(t, keyZulmara) != "" {
		t.Fatal("a locked raid saved the spec")
	}
}

func TestTheSpecPickWaitsForAnUnverifiedCharacter(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	w.exec(t, `update guild_characters set verified_at = null where character_key = $1`, keyZulmara)
	rep := w.mustSpecPress(t, specPickID(w, "in", keyZulmara), "paladin-holy")
	wantReply(t, rep, spec(t, "Press: refusals", 5, map[string]string{"Name": "Zulmara"}), true)
}

func TestChangeSpecNeverChangesTheState(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	w.mustSpecPress(t, specPickID(w, "bench", keyZulmara), "paladin-holy")
	chg := fmt.Sprintf("gspc:1:chg:%d:%s", w.raidID, keyZulmara)
	card := w.mustSpecPress(t, chg)
	lead := fill(guideCopy.SpecLeadChange, "{Name}", "Zulmara", "{Title}", "Onyxia's Lair", "{Date}", w.day())
	sel := card.Message.Components[0][0]
	if !card.Update || card.Message.Embeds[0].Description != lead || !sel.Options[0].Default || sel.Options[1].Default {
		t.Fatalf("change card = %+v", card)
	}
	setID := sel.CustomID
	rep := w.mustSpecPress(t, setID, "paladin-retribution")
	want := "Zulmara is already on the Bench for Onyxia's Lair, " + w.day() + ". Retribution is kept for Zulmara."
	wantReply(t, rep, want, true, changeSpecButton(w, keyZulmara))
	if state, role, spec := w.signupRow(t, keyZulmara); state != "bench" || role != "dps" || spec != "Retribution" {
		t.Fatalf("row %s %s %s", state, role, spec)
	}
	w.press(t, "out")
	rep = w.mustSpecPress(t, setID, "paladin-protection")
	wantReply(t, rep, "Protection is kept for Zulmara.", true)
	if state, _, _ := w.signupRow(t, keyZulmara); state != "out" || w.raidSpecOf(t, keyZulmara) != "paladin-protection" {
		t.Fatalf("out row %s", state)
	}
}

func TestAnOlderSpecCardIsStale(t *testing.T) {
	w := newPressWorld(t)
	rep := w.mustSpecPress(t, "gspc:0:pick:1:in:"+keyZulmara, "paladin-holy")
	wantReply(t, rep, guideCopy.SpecStale, true)
}

// A failed save keeps the card with its select and says pressCopy.Failure in the lead.
func TestASpecFailureKeepsTheCard(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	card := w.press(t, "in").Message
	embeds, _ := json.Marshal(card.Embeds)
	components, _ := json.Marshal(card.Components)
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": specPickID(w, "in", keyZulmara)})
	in.Source = sourceMessage{Embeds: embeds, Components: components}
	rep := cardFailureReply(in)
	var echoed []map[string]any
	if rep.Message.echo == nil || json.Unmarshal(rep.Message.echo.Embeds, &echoed) != nil || !rep.Update {
		t.Fatalf("reply = %+v", rep)
	}
	if echoed[0]["description"] != pressCopy.Failure || echoed[0]["title"] != guideCopy.SpecTitle ||
		string(rep.Message.echo.Components) != string(components) {
		t.Fatalf("echo = %v", echoed)
	}
	w.port.signUpErr = errors.New("boom")
	if _, err := w.specPress(t, specPickID(w, "in", keyZulmara), "paladin-holy"); err == nil {
		t.Fatal("a failed sign-up answered without an error")
	}
	if w.raidSpecOf(t, keyZulmara) != "" {
		t.Fatal("a failed sign-up saved the spec")
	}
	if _, refused := signupRefusal(guilds.ErrSignupRateLimited); !refused {
		t.Fatal("rate limit is a refusal")
	}
}

func (p *pressGuilds) SaveRaidSpec(ctx context.Context, cmd guilds.RaidSpecCommand) (guilds.RaidSpecResult, error) {
	if p.saveSpecErr != nil {
		return guilds.RaidSpecResult{}, p.saveSpecErr
	}
	return p.Service.SaveRaidSpec(ctx, cmd)
}

// A save that fails after the sign-up stood answers the outcome without spec.kept, and logs.
func TestAFailedSpecSaveStillAnswersTheSignUp(t *testing.T) {
	w := newPressWorld(t)
	paladin(t, w)
	w.port.saveSpecErr = errors.New("boom")
	rep := w.mustSpecPress(t, specPickID(w, "in", keyZulmara), "paladin-holy")
	want := "Zulmara is In as Healer for Onyxia's Lair, " + w.day() + ". 1 of 40 signed up."
	wantReply(t, rep, want, true, changeSpecButton(w, keyZulmara))
	if w.signups(t) != 1 || w.raidSpecOf(t, keyZulmara) != "" || !strings.Contains(w.log.String(), "save_failed") {
		t.Fatalf("signups %d spec %q log %s", w.signups(t), w.raidSpecOf(t, keyZulmara), w.log.String())
	}
}
