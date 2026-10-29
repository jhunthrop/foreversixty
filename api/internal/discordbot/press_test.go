package discordbot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// The install harness's fake answers nothing of the sign-up flow.
func (f *fakeGuilds) SignUp(context.Context, guilds.SignupCommand) (guilds.SignupResult, error) {
	return guilds.SignupResult{}, errors.New("fakeGuilds: no sign-ups")
}

func (f *fakeGuilds) SaveRaidSpec(context.Context, guilds.RaidSpecCommand) (guilds.RaidSpecResult, error) {
	return guilds.RaidSpecResult{}, errors.New("fakeGuilds: no specs")
}

func (f *fakeGuilds) SignupCandidates(context.Context, int64, int64) ([]guilds.SignupCandidate, error) {
	return nil, nil
}

func (f *fakeGuilds) PlannedRaids(context.Context, int64, time.Time) ([]guilds.Raid, error) {
	return nil, nil
}

func (f *fakeGuilds) GetRaid(context.Context, int64, int64) (guilds.Raid, bool, error) {
	return guilds.Raid{}, false, nil
}

func (f *fakeGuilds) NextRaidFor(context.Context, int64, int64, *string, time.Time) (*guilds.NextRaidView, error) {
	return nil, nil
}

func (f *fakeGuilds) MemberRaidView(context.Context, int64, int64, int64) (guilds.RaidView, bool, error) {
	return guilds.RaidView{}, false, nil
}

func (f *fakeGuilds) GuildName(context.Context, int64) (string, error) { return "", nil }

func (f *fakeGuilds) GuildRaidTimezone(context.Context, int64) (string, error) { return "", nil }

const (
	pressPublic  = "https://site.test"
	pressAPI     = "https://api.test"
	pressServer  = "srv1"
	pressDiscord = "d-1"
	pressHeading = "### 6.2 In Discord"
	keyZulmara   = "us/normal/zulmara"
	keyDusk      = "us/normal/duskstrider"
)

// pressGuilds is the real guilds service, with candidates or a sign-up error scripted when set.
type pressGuilds struct {
	*guilds.Service
	candidates []guilds.SignupCandidate
	signUpErr  error
	// candidateReads counts SignupCandidates calls.
	candidateReads int
}

func (p *pressGuilds) SignupCandidates(ctx context.Context, guildID, userID int64) ([]guilds.SignupCandidate, error) {
	p.candidateReads++
	if p.candidates != nil {
		return p.candidates, nil
	}
	return p.Service.SignupCandidates(ctx, guildID, userID)
}

func (p *pressGuilds) SignUp(ctx context.Context, cmd guilds.SignupCommand) (guilds.SignupResult, error) {
	if p.signUpErr != nil {
		return guilds.SignupResult{}, p.signUpErr
	}
	return p.Service.SignUp(ctx, cmd)
}

// pressWorld is Olympus XXVII bound to srv1 with a channel, a linked account with no
// characters yet, and one raid of Onyxia's Lair in two days.
type pressWorld struct {
	st     *Store
	gs     *guilds.Service
	port   *pressGuilds
	svc    *Service
	rest   *FakeREST
	log    *bytes.Buffer
	gid    int64
	uid    int64
	raidID int64
	// raidStart is the raid's start, for the outcomes' day.
	raidStart time.Time
}

func newPressWorld(t *testing.T) *pressWorld {
	t.Helper()
	st := testStore(t)
	ctx := context.Background()
	w := &pressWorld{st: st, rest: &FakeREST{}, log: &bytes.Buffer{}}
	w.uid = seedUser(t, st.Pool)
	w.gid = seedGuild(t, st.Pool, "Olympus XXVII")
	if _, err := st.BindServer(ctx, w.gid, pressServer, "Olympus", w.uid); err != nil {
		t.Fatal(err)
	}
	if err := st.SetChannel(ctx, w.gid, postChannel, "raid-signups"); err != nil {
		t.Fatal(err)
	}
	if err := st.SaveLink(ctx, w.uid, pressDiscord, "zulmara"); err != nil {
		t.Fatal(err)
	}
	w.gs = &guilds.Service{
		Store: &guilds.Store{Pool: st.Pool}, Accounts: &auth.Store{Pool: st.Pool}, PublicBaseURL: pressPublic,
		RaidPosts: st, Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	w.port = &pressGuilds{Service: w.gs}
	w.svc = New(Deps{Store: st, REST: w.rest, Guilds: w.port, PublicBaseURL: pressPublic, APIBaseURL: pressAPI, Log: slog_(w.log)})
	w.raidStart = time.Now().Add(48 * time.Hour).Truncate(time.Minute)
	w.raidID = w.addRaid(t, w.gid, w.raidStart)
	return w
}

// day is the raid's day as the press outcomes name it (the guild has no raid zone: UTC).
func (w *pressWorld) day() string { return w.raidStart.UTC().Format("Mon 2 Jan") }

func (w *pressWorld) addRaid(t *testing.T, guildID int64, starts time.Time) int64 {
	t.Helper()
	var id int64
	err := w.st.Pool.QueryRow(context.Background(), `insert into guild_raids
		(guild_id, zone, title, starts_at, duration_minutes, created_by)
		values ($1, 'raid:onyxias-lair', 'Onyxia''s Lair', $2, 180, $3) returning id`,
		guildID, starts.Truncate(time.Minute), w.uid).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

// pressSpec is the raid spec addChar's hunters bring: known, so a press signs up at once.
const pressSpec = "hunter-marksmanship"

// addChar puts a hunter of userID on guildID's roster, verified or pending, with its raid spec
// picked (addCharNoSpec leaves it unknown, so In, Late and Bench ask it first).
func (w *pressWorld) addChar(t *testing.T, guildID, userID int64, key, name string, verified bool) {
	t.Helper()
	w.addCharNoSpec(t, guildID, userID, key, name, verified)
	if _, err := w.st.Pool.Exec(context.Background(), `update guild_characters set raid_spec = $1 where character_key = $2`, pressSpec, key); err != nil {
		t.Fatal(err)
	}
}

// addCharNoSpec is addChar with no raid spec picked.
func (w *pressWorld) addCharNoSpec(t *testing.T, guildID, userID int64, key, name string, verified bool) {
	t.Helper()
	ctx := context.Background()
	if _, err := w.st.Pool.Exec(ctx, `insert into characters (key, region, ruleset, name, class, user_id)
		values ($1, 'us', 'normal', $2, 'hunter', $3) on conflict (key) do nothing`, key, name, userID); err != nil {
		t.Fatal(err)
	}
	if _, err := w.st.Pool.Exec(ctx, `insert into guild_characters (guild_id, character_key, user_id, rank, verified_at)
		values ($1, $2, $3, 'member', case when $4 then now() else null end)`, guildID, key, userID, verified); err != nil {
		t.Fatal(err)
	}
	w.recompute(t, guildID, userID)
}

// recompute re-derives userID's guild membership from its characters, as an approval does.
func (w *pressWorld) recompute(t *testing.T, guildID, userID int64) {
	t.Helper()
	ctx := context.Background()
	tx, err := w.st.Pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if err := guilds.RecomputeMembership(ctx, tx, guildID, &userID); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
}

func (w *pressWorld) setMain(t *testing.T, key string) {
	t.Helper()
	if _, err := w.st.Pool.Exec(context.Background(), `update users set main_character_key = $2 where id = $1`, w.uid, key); err != nil {
		t.Fatal(err)
	}
}

func (w *pressWorld) exec(t *testing.T, q string, args ...any) {
	t.Helper()
	if _, err := w.st.Pool.Exec(context.Background(), q, args...); err != nil {
		t.Fatal(err)
	}
}

func interactionOf(kind int, discordUser string, data any) Interaction {
	raw, _ := json.Marshal(data)
	return Interaction{ID: "i1", GuildID: pressServer, DiscordUserID: discordUser, Type: kind, Data: raw}
}

func (w *pressWorld) pressAs(t *testing.T, discordUser string, raidID int64, state string) Reply {
	t.Helper()
	rep, err := w.svc.press(context.Background(), interactionOf(typeComponent, discordUser,
		map[string]any{"custom_id": raidCustomID(raidID, state)}))
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func (w *pressWorld) press(t *testing.T, state string) Reply {
	t.Helper()
	return w.pressAs(t, pressDiscord, w.raidID, state)
}

func (w *pressWorld) pick(t *testing.T, state, key string) Reply {
	t.Helper()
	rep, err := w.svc.pick(context.Background(), interactionOf(typeComponent, pressDiscord,
		map[string]any{"custom_id": pickCustomID(w.raidID, state), "values": []string{key}}))
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

func (w *pressWorld) signups(t *testing.T) int {
	t.Helper()
	var n int
	if err := w.st.Pool.QueryRow(context.Background(), `select count(*) from guild_raid_signups`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func (w *pressWorld) raidURL() string {
	return fmt.Sprintf("%s/guild/us/normal/olympus-xxvii/raids/%d", pressPublic, w.raidID)
}

// spec is a "Press: …" segment of §6.2 with its markers filled.
func spec(t *testing.T, row string, index int, fill map[string]string) string {
	t.Helper()
	return specString(t, pressHeading, row, index, fill)
}

// wantButton is a link button (url), or with id a Secondary button such as Change spec.
type wantButton struct{ label, url, id string }

// changeSpecButton is the Change spec button under an outcome for key.
func changeSpecButton(w *pressWorld, key string) wantButton {
	return wantButton{label: guideCopy.SpecChange, id: fmt.Sprintf("gspc:1:chg:%d:%s", w.raidID, key)}
}

func wantReply(t *testing.T, rep Reply, text string, update bool, buttons ...wantButton) {
	t.Helper()
	m := rep.Message
	if m.Content != text || !m.Ephemeral || rep.Update != update {
		t.Fatalf("reply = %q ephemeral=%v update=%v, want %q update=%v", m.Content, m.Ephemeral, rep.Update, text, update)
	}
	var got []Component
	for _, row := range m.Components {
		got = append(got, row...)
	}
	if len(got) != len(buttons) {
		t.Fatalf("components = %+v, want %+v", got, buttons)
	}
	for i, b := range buttons {
		if b.id != "" {
			if got[i].Style != StyleSecondary || got[i].Label != b.label || got[i].CustomID != b.id {
				t.Fatalf("button %d = %+v, want %+v", i, got[i], b)
			}
			continue
		}
		if got[i].Style != StyleLink || got[i].Label != b.label || got[i].URL != b.url {
			t.Fatalf("button %d = %+v, want %+v", i, got[i], b)
		}
	}
	if _, err := json.Marshal(m); err != nil {
		t.Fatal(err)
	}
}

func TestPressCopyMatchesTheSpec(t *testing.T) {
	refusal := func(i int) string { return spec(t, "Press: refusals", i, nil) }
	outcome := func(i int) string { return spec(t, "Press: outcomes", i, nil) }
	pick := func(i int) string { return spec(t, "Press: select", i, nil) }
	checks := map[string][2]string{
		"NotLinked":    {pressCopy.NotLinked, spec(t, "Press: not linked", 0, nil)},
		"Link":         {linkCopy.LinkButton, spec(t, "Press: not linked", 1, nil)},
		"Locked":       {pressCopy.Locked, refusal(0)},
		"Cancelled":    {pressCopy.Cancelled, refusal(1)},
		"Ended":        {pressCopy.Ended, refusal(2)},
		"NotMember":    {pressCopy.NotMember, refusal(3)},
		"Setup":        {pressCopy.SetupButton, refusal(4)},
		"WaitingOne":   {pressCopy.WaitingOne, refusal(5)},
		"WaitingMany":  {pressCopy.WaitingMany, refusal(6)},
		"RateLimited":  {pressCopy.RateLimited, refusal(7)},
		"Failure":      {pressCopy.Failure, refusal(8)},
		"Prompt":       {pressCopy.SelectPrompt, pick(0)},
		"Placeholder":  {pressCopy.SelectPlaceholder, pick(1)},
		"Option":       {pressCopy.OptionDescription, pick(2)},
		"Main":         {pressCopy.OptionMain, pick(3)},
		"Showing":      {pressCopy.Showing, pick(4)},
		"In":           {pressCopy.InSaved + pressCopy.Count, outcome(0)},
		"Late":         {pressCopy.LateSaved, outcome(1)},
		"Note":         {pressCopy.NoteButton, outcome(2)},
		"Bench":        {pressCopy.BenchSaved, outcome(3)},
		"Out":          {pressCopy.OutSaved, outcome(4)},
		"Moved":        {pressCopy.Moved, outcome(5)},
		"Already":      {pressCopy.Already, outcome(6)},
		"AlreadyBench": {pressCopy.AlreadyBench, "{Name} " + outcome(7) + " for {Title}, {Date}."},
		"AlreadyOut":   {pressCopy.AlreadyOut, "{Name} " + outcome(8) + " for {Title}, {Date}."},
		"NoRaid":       {pressCopy.NoRaid, specString(t, pressHeading, "`/raid`", 5, nil)},
		"CountOpen":    {pressCopy.CountOpen, " " + specString(t, pressHeading, "Post fields", 12, nil) + "."},
	}
	for name, c := range checks {
		if c[0] != c[1] {
			t.Errorf("%s = %q, spec %q", name, c[0], c[1])
		}
	}
}

func TestPressFlowSentences(t *testing.T) {
	title := "Onyxia's Lair"
	zulmara := func(t *testing.T, w *pressWorld) { w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true) }
	inSentence := func(w *pressWorld, signed string) string {
		return spec(t, "Press: outcomes", 0, map[string]string{"Name": "Zulmara", "Role": "DPS", "Title": title, "Date": w.day(), "signed": signed, "size": "40"})
	}
	cases := []struct {
		name    string
		setup   func(*testing.T, *pressWorld)
		act     func(*testing.T, *pressWorld) Reply
		want    func(*pressWorld) string
		buttons func(*pressWorld) []wantButton
		update  bool
	}{
		{name: "unlinked",
			act:  func(t *testing.T, w *pressWorld) Reply { return w.pressAs(t, "d-stranger", w.raidID, "in") },
			want: func(*pressWorld) string { return spec(t, "Press: not linked", 0, nil) },
			buttons: func(w *pressWorld) []wantButton {
				return []wantButton{{label: linkCopy.LinkButton, url: w.svc.linkStartURL()}}
			}},
		{name: "cancelled", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.exec(t, `update guild_raids set status = 'cancelled' where id = $1`, w.raidID)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 1, nil) }},
		{name: "ended", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.exec(t, `update guild_raids set starts_at = now() - interval '5 hours' where id = $1`, w.raidID)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 2, nil) }},
		{name: "locked", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.exec(t, `update guild_raids set status = 'locked' where id = $1`, w.raidID)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 0, nil) }},
		{name: "not a member",
			want: func(*pressWorld) string {
				return spec(t, "Press: refusals", 3, map[string]string{"Guild": "Olympus XXVII"})
			},
			buttons: func(*pressWorld) []wantButton {
				return []wantButton{{label: pressCopy.SetupButton, url: pressPublic + "/setup"}}
			}},
		{name: "one pending", setup: func(t *testing.T, w *pressWorld) {
			w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", false)
		}, want: func(*pressWorld) string {
			return spec(t, "Press: refusals", 5, map[string]string{"Name": "Duskstrider"})
		}},
		{name: "several pending", setup: func(t *testing.T, w *pressWorld) {
			w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", false)
			w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", false)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 6, nil) }},
		{name: "one verified In", setup: zulmara,
			want:    func(w *pressWorld) string { return inSentence(w, "1") },
			buttons: func(w *pressWorld) []wantButton { return []wantButton{changeSpecButton(w, keyZulmara)} }},
		{name: "late", setup: zulmara,
			act: func(t *testing.T, w *pressWorld) Reply { return w.press(t, "late") },
			want: func(w *pressWorld) string {
				return spec(t, "Press: outcomes", 1, map[string]string{"Name": "Zulmara", "Role": "DPS", "Title": title, "Date": w.day()})
			},
			buttons: func(w *pressWorld) []wantButton {
				return []wantButton{changeSpecButton(w, keyZulmara), {label: pressCopy.NoteButton, url: w.raidURL() + "#signup"}}
			}},
		{name: "bench", setup: zulmara,
			act: func(t *testing.T, w *pressWorld) Reply { return w.press(t, "bench") },
			want: func(w *pressWorld) string {
				return spec(t, "Press: outcomes", 3, map[string]string{"Name": "Zulmara", "Title": title, "Date": w.day()})
			},
			buttons: func(w *pressWorld) []wantButton { return []wantButton{changeSpecButton(w, keyZulmara)} }},
		{name: "out", setup: zulmara,
			act: func(t *testing.T, w *pressWorld) Reply { return w.press(t, "out") },
			want: func(w *pressWorld) string {
				return spec(t, "Press: outcomes", 4, map[string]string{"Name": "Zulmara", "Title": title, "Date": w.day()})
			}},
		{name: "R3 move", update: true, setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
			if rep := w.pick(t, "in", keyDusk); !strings.Contains(rep.Message.Content, "Duskstrider is In") {
				t.Fatalf("first pick = %q", rep.Message.Content)
			}
		},
			act: func(t *testing.T, w *pressWorld) Reply { return w.pick(t, "in", keyZulmara) },
			want: func(w *pressWorld) string {
				return strings.TrimSuffix(inSentence(w, "1"), " 1 of 40 signed up.") +
					spec(t, "Press: outcomes", 5, map[string]string{"Other": "Duskstrider"})
			},
			buttons: func(w *pressWorld) []wantButton { return []wantButton{changeSpecButton(w, keyZulmara)} }},
		{name: "already", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.press(t, "in")
		}, want: func(w *pressWorld) string {
			return spec(t, "Press: outcomes", 6, map[string]string{"Name": "Zulmara", "State": "In", "Role": "DPS", "Title": title, "Date": w.day()})
		}, buttons: func(w *pressWorld) []wantButton { return []wantButton{changeSpecButton(w, keyZulmara)} }},
		{name: "rate limited", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.gs.SignupLimiter = httpx.NewKeyLimiter(1, time.Minute)
			w.press(t, "late")
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 7, nil) }},
		{name: "locked before only pending", setup: func(t *testing.T, w *pressWorld) {
			w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", false)
			w.exec(t, `update guild_raids set status = 'locked' where id = $1`, w.raidID)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 0, nil) }},
		{name: "locked before not a member", setup: func(t *testing.T, w *pressWorld) {
			w.exec(t, `update guild_raids set status = 'locked' where id = $1`, w.raidID)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 0, nil) }},
		{name: "service locked", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.port.signUpErr = fmt.Errorf("wrapped: %w", guilds.ErrRaidLocked)
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 0, nil) }},
		{name: "other error", setup: func(t *testing.T, w *pressWorld) {
			zulmara(t, w)
			w.port.signUpErr = errors.New("boom")
		}, want: func(*pressWorld) string { return spec(t, "Press: refusals", 8, nil) }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := newPressWorld(t)
			if c.setup != nil {
				c.setup(t, w)
			}
			act := c.act
			if act == nil {
				act = func(t *testing.T, w *pressWorld) Reply { return w.press(t, "in") }
			}
			var buttons []wantButton
			if c.buttons != nil {
				buttons = c.buttons(w)
			}
			wantReply(t, act(t, w), c.want(w), c.update, buttons...)
		})
	}
}

func TestPressLogsNoInteractionToken(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.port.signUpErr = errors.New("boom")
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": raidCustomID(w.raidID, "in")})
	in.Token = "secret-interaction-token"
	if _, err := w.svc.press(context.Background(), in); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(w.log.String(), in.Token) || !strings.Contains(w.log.String(), "discord_signup") {
		t.Fatalf("log = %s", w.log.String())
	}
}

// selectOf is the reply's one select.
func selectOf(t *testing.T, rep Reply) Component {
	t.Helper()
	if len(rep.Message.Components) != 1 || len(rep.Message.Components[0]) != 1 ||
		rep.Message.Components[0][0].Kind != KindStringSelect {
		t.Fatalf("components = %+v", rep.Message.Components)
	}
	return rep.Message.Components[0][0]
}

func TestSelectOrderAndCap(t *testing.T) {
	w := newPressWorld(t)
	for i := 1; i <= 30; i++ {
		w.addChar(t, w.gid, w.uid, fmt.Sprintf("us/normal/char%02d", i), fmt.Sprintf("Char%02d", i), true)
	}
	w.setMain(t, "us/normal/char17")
	rep := w.press(t, "in")
	sel := selectOf(t, rep)
	if len(sel.Options) != maxSelectOptions {
		t.Fatalf("options = %d", len(sel.Options))
	}
	first := sel.Options[0]
	if first.Label != "Char17" || first.Description != "Marksmanship Hunter · DPS"+pressCopy.OptionMain || first.Value != "us/normal/char17" {
		t.Fatalf("first option = %+v", first)
	}
	if sel.Options[1].Label != "Char01" || strings.HasSuffix(sel.Options[1].Description, pressCopy.OptionMain) {
		t.Fatalf("second option = %+v", sel.Options[1])
	}
	var starts time.Time
	if err := w.st.Pool.QueryRow(context.Background(), `select starts_at from guild_raids where id = $1`, w.raidID).Scan(&starts); err != nil {
		t.Fatal(err)
	}
	prompt := spec(t, "Press: select", 0, map[string]string{"State": "**In**", "Title": "Onyxia's Lair", "unix": fmt.Sprint(starts.Unix())})
	wantText := prompt + " " + spec(t, "Press: select", 4, map[string]string{"n": "30"})
	if rep.Message.Content != wantText || !rep.Message.Ephemeral || rep.Update {
		t.Fatalf("text = %q, want %q", rep.Message.Content, wantText)
	}
	if !strings.HasSuffix(rep.Message.Content, "Showing 25 of 30. Sign up the others on the raid page.") {
		t.Fatalf("text = %q", rep.Message.Content)
	}
	if sel.CustomID != pickCustomID(w.raidID, "in") || w.signups(t) != 0 {
		t.Fatalf("custom id %q, signups %d", sel.CustomID, w.signups(t))
	}
}

// forgetCandidates drops the bot's cached candidates, as 30 s passing would.
func (w *pressWorld) forgetCandidates() {
	w.svc.candidates.mu.Lock()
	defer w.svc.candidates.mu.Unlock()
	w.svc.candidates.by = nil
}

func TestAutocompleteCachesCandidatesPerAccount(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	clock := &testClock{t: time.Now().Truncate(time.Microsecond)} // Postgres keeps microseconds
	w.svc.now = clock.now
	for range 3 {
		w.suggest(t, pressDiscord, optCharacter, "z")
	}
	if w.port.candidateReads != 1 {
		t.Fatalf("candidate reads = %d, want 1 inside the window", w.port.candidateReads)
	}
	clock.advance(candidatesTTL)
	w.suggest(t, pressDiscord, optCharacter, "")
	if w.port.candidateReads != 2 {
		t.Fatalf("candidate reads = %d, want a fresh read after %s", w.port.candidateReads, candidatesTTL)
	}
}

func TestAPressReadsCandidatesFresh(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", false)
	w.suggest(t, pressDiscord, optCharacter, "") // caches the pending character
	w.exec(t, `update guild_characters set verified_at = now() where character_key = $1`, keyZulmara)
	w.recompute(t, w.gid, w.uid)
	if rep := w.press(t, "in"); !strings.HasPrefix(rep.Message.Content, "Zulmara is In as DPS") {
		t.Fatalf("press after approval = %q", rep.Message.Content)
	}
}

func TestSelectOptionsCarryTheLoggedSpec(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
	fought := time.Now()
	// fight_metrics is partitioned by month: make this month's, so a fresh database on any date works.
	if err := db.EnsureMetricsPartition(context.Background(), w.st.Pool, fought); err != nil {
		t.Fatal(err)
	}
	w.exec(t, `insert into reports (id, guild_id, visibility, status) values ('t9-spec', $1, 'public', 'complete')`, w.gid)
	w.exec(t, `insert into fight_metrics (report_id, fight_index, player_key, class, spec, role, ilvl, faction, fought_at)
		values ('t9-spec', 0, $1, 'hunter', 'Marksmanship', 'dps', 60, 'alliance', $2)`, keyZulmara, fought)
	sel := selectOf(t, w.press(t, "in"))
	var got string
	for _, o := range sel.Options {
		if o.Value == keyZulmara {
			got = o.Description
		}
	}
	if want := spec(t, "Press: select", 2, map[string]string{"Spec": "Marksmanship", "Class": "Hunter", "Role": "DPS"}); got != want {
		t.Fatalf("description = %q, want %q", got, want)
	}
}

func TestASelectWithNoUsableOptionLinksTheRaidPage(t *testing.T) {
	w := newPressWorld(t)
	long := strings.Repeat("k", maxOptionRunes+1)
	w.port.candidates = []guilds.SignupCandidate{
		{CharacterKey: long + "a", Name: "A", Class: "hunter", Role: "dps", Verified: true},
		{CharacterKey: long + "b", Name: "B", Class: "hunter", Role: "dps", Verified: true},
	}
	wantReply(t, w.press(t, "in"), pressCopy.Failure, false, wantButton{label: postCopy.RaidPage, url: w.raidURL()})
}

func TestAFailedPickStillReplacesTheSelect(t *testing.T) {
	w := newPressWorld(t)
	in := interactionOf(typeComponent, pressDiscord, map[string]any{"custom_id": pickCustomID(w.raidID, "in")})
	rep := w.svc.replyOf(in, handled{err: errors.New("boom")}, true)
	if !rep.Update || rep.Message.Content != commonCopy.Failure {
		t.Fatalf("reply = %+v", rep)
	}
	if rep := w.svc.replyOf(in, handled{err: errors.New("boom")}, false); rep.Update {
		t.Fatal("a press's failure must be a new message")
	}
}

func TestSelectOptionsStayInsideLimits(t *testing.T) {
	w := newPressWorld(t)
	long := "@everyone " + strings.Repeat("ẞ", 95)
	w.port.candidates = []guilds.SignupCandidate{
		{CharacterKey: "us/normal/a", Name: long, Class: "hunter", Spec: strings.Repeat("M", 100), Role: "dps", Verified: true, Main: true},
		{CharacterKey: "us/normal/b", Name: "<@123> **b**", Class: "warrior", Spec: "Protection", Role: "tank", Verified: true},
	}
	w.exec(t, `update guild_raids set title = $2 where id = $1`, w.raidID, "@everyone **"+strings.Repeat("x", 66))
	rep := w.press(t, "bench")
	sel := selectOf(t, rep)
	if sel.Placeholder != spec(t, "Press: select", 1, nil) || len(sel.CustomID) > 100 {
		t.Fatalf("placeholder %q, custom id %q", sel.Placeholder, sel.CustomID)
	}
	for _, o := range sel.Options {
		if utf8.RuneCountInString(o.Label) > maxOptionRunes || utf8.RuneCountInString(o.Description) > maxOptionRunes ||
			utf8.RuneCountInString(o.Value) > maxOptionRunes {
			t.Fatalf("option over the limits: %+v", o)
		}
	}
	if !strings.HasSuffix(sel.Options[0].Description, " Hunter · DPS · your main") {
		t.Fatalf("main description = %q", sel.Options[0].Description)
	}
	if sel.Options[1].Description != "Protection Warrior · Tank" {
		t.Fatalf("description = %q", sel.Options[1].Description)
	}
	if strings.Contains(rep.Message.Content, "**x") || utf8.RuneCountInString(rep.Message.Content) > 2000 {
		t.Fatalf("title not escaped: %q", rep.Message.Content)
	}
	body, err := json.Marshal(rep.Message)
	if err != nil || !strings.Contains(string(body), `"allowed_mentions":{"parse":[]}`) {
		t.Fatalf("body %s err %v", body, err)
	}
}

func TestPickReplacesTheEphemeral(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.addChar(t, w.gid, w.uid, keyDusk, "Duskstrider", true)
	other := seedUser(t, w.st.Pool)
	w.addChar(t, w.gid, other, "us/normal/vexlash", "Vexlash", true)

	wantReply(t, w.pick(t, "in", "us/normal/vexlash"), pressCopy.Failure, true)
	if n := w.signups(t); n != 0 {
		t.Fatalf("signups = %d after another account's character", n)
	}
	rep := w.pick(t, "late", keyZulmara)
	if !rep.Update || !strings.HasPrefix(rep.Message.Content, "Zulmara is Late as DPS") || w.signups(t) != 1 {
		t.Fatalf("pick = %+v, signups %d", rep, w.signups(t))
	}

	// The type 7 answer writes components explicitly, so the select goes.
	rec := httptest.NewRecorder()
	w.svc.writeReply(rec, Reply{Message: Message{Content: "done", Ephemeral: true}, Update: true}, respMessage)
	var got struct {
		Type int `json:"type"`
		Data struct {
			Content    string            `json:"content"`
			Components []json.RawMessage `json:"components"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Type != respUpdateMessage || got.Data.Content != "done" || got.Data.Components == nil ||
		!strings.Contains(rec.Body.String(), `"components":[]`) {
		t.Fatalf("update body = %s", rec.Body.String())
	}
	if !updateComponents()[pickIDPrefix] || updateComponents()[raidIDPrefix] {
		t.Fatal("pick must defer as an update, the post's buttons as a new message")
	}
}

func TestPressRefusesARaidOfAnotherGuild(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	other := seedGuild(t, w.st.Pool, "Mograine Raiders")
	w.addChar(t, other, w.uid, "us/normal/zulmaraalt", "Zulmaraalt", true)
	theirs := w.addRaid(t, other, time.Now().Add(24*time.Hour))

	wantReply(t, w.pressAs(t, pressDiscord, theirs, "in"), pressCopy.Failure, false)
	rep, err := w.svc.pick(context.Background(), interactionOf(typeComponent, pressDiscord,
		map[string]any{"custom_id": pickCustomID(theirs, "in"), "values": []string{"us/normal/zulmaraalt"}}))
	if err != nil {
		t.Fatal(err)
	}
	wantReply(t, rep, pressCopy.Failure, true)
	if n := w.signups(t); n != 0 {
		t.Fatalf("signups = %d", n)
	}
}

func TestPressMarksThePostAndNeverEditsIt(t *testing.T) {
	w := newPressWorld(t)
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	ctx := context.Background()
	if queued, err := w.st.EnqueuePost(ctx, w.raidID, w.gid); err != nil || !queued {
		t.Fatalf("enqueue = %v %v", queued, err)
	}
	w.exec(t, `update guild_raid_discord_posts set dirty_at = null, message_id = 'm1' where raid_id = $1`, w.raidID)
	w.press(t, "in")
	var dirty bool
	if err := w.st.Pool.QueryRow(ctx, `select dirty_at is not null from guild_raid_discord_posts where raid_id = $1`, w.raidID).Scan(&dirty); err != nil {
		t.Fatal(err)
	}
	if !dirty {
		t.Fatal("the press did not mark the post")
	}
	if calls := w.rest.Calls(); len(calls) != 0 {
		t.Fatalf("the press called Discord: %+v", calls)
	}
}
