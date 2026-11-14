package discordbot

import (
	"context"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/directory"
)

// pickWorld is the press world with the sign-up by character pick on: the guild's Blizzard id
// and a directory roster of two members, one of them held by the world's linked account.
func newPickWorld(t *testing.T) *pressWorld {
	t.Helper()
	w := newPressWorld(t)
	w.exec(t, `truncate character_directory`)
	w.exec(t, `update guilds set bnet_guild_id = 77 where id = $1`, w.gid)
	w.exec(t, `insert into character_directory (region, bnet_character_id, name, slug, class, level, guild_bnet_id, guild_name, guild_rank)
		values ('us', 101, 'Vela Thornbrook', 'vela-thornbrook', 'hunter', 41, 77, 'g', 0),
		       ('us', 102, 'Orrin Ashmantle', 'orrin-ashmantle', 'warrior', 12, 77, 'g', 3)`)
	w.svc.d.Claims = &directory.Claims{Pool: w.st.Pool}
	return w
}

func (w *pressWorld) signupAs(t *testing.T, discordUser string, options ...map[string]any) Reply {
	t.Helper()
	in := interactionOf(typeCommand, discordUser, map[string]any{"name": "signup", "options": options})
	in.DisplayName = "a raider"
	rep, err := w.svc.cmdSignup(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	return rep
}

// A Discord user with no account is offered the guild's directory roster, never the raids.
func TestPickAutocompleteOffersTheDirectoryRoster(t *testing.T) {
	w := newPickWorld(t)
	choices := w.suggest(t, "d-new", optCharacter, "")
	if len(choices) != 2 || choices[0].Name != "Vela Thornbrook · 41 Hunter" || choices[0].Value != directory.PickValue(101) {
		t.Fatalf("choices = %+v", choices)
	}
	if typed := w.suggest(t, "d-new", optCharacter, "orr"); len(typed) != 1 || typed[0].Value != directory.PickValue(102) {
		t.Fatalf("typed choices = %+v", typed)
	}
	if raids := w.suggest(t, "d-new", optRaid, ""); len(raids) != 0 {
		t.Fatalf("raid choices for a stranger = %+v", raids)
	}
	// The linked Battle.net account is offered its own characters only.
	if mine := w.suggest(t, pressDiscord, optCharacter, ""); len(mine) != 0 {
		t.Fatalf("a Battle.net account was offered picks: %+v", mine)
	}
}

// The first pick makes a Discord-only account, claims the character for officers to see, and
// signs it up In; a second Discord user is told it is taken; the account's next press works.
func TestSignupByPick(t *testing.T) {
	w := newPickWorld(t)
	// A new pick has no spec yet: the claim line leads the "Which spec tonight?" card.
	rep := w.signupAs(t, "d-new", opt(optCharacter, directory.PickValue(101)))
	lead := fill(pickCopy.Claimed, "{Name}", "Vela Thornbrook")
	if !strings.HasPrefix(rep.Message.Content, lead) {
		t.Fatalf("reply = %q", rep.Message.Content)
	}
	w.exec(t, `update guild_characters set raid_spec = $1 where character_key = 'us/normal/vela-thornbrook'`, pressSpec)
	in := w.signupAs(t, "d-new", opt(optStatus, "in"), opt(optCharacter, directory.PickValue(101)))
	if strings.HasPrefix(in.Message.Content, lead) || !strings.Contains(in.Message.Content, "Vela Thornbrook is In") {
		t.Fatalf("second pick reply = %q, want the In outcome alone", in.Message.Content)
	}
	var only bool
	if err := w.st.Pool.QueryRow(context.Background(),
		`select u.discord_only from users u join discord_links l on l.user_id = u.id where l.discord_user_id = 'd-new'`).Scan(&only); err != nil || !only {
		t.Fatalf("discord-only account = %v, %v", only, err)
	}
	if w.signups(t) != 1 {
		t.Fatalf("sign-ups = %d", w.signups(t))
	}
	taken := w.signupAs(t, "d-other", opt(optStatus, "in"), opt(optCharacter, directory.PickValue(101)))
	wantReply(t, taken, fill(pickCopy.Taken, "{Name}", "Vela Thornbrook"), false)
	// The picked character is the account's now: a press on the post signs it up with no pick.
	out := w.pressAs(t, "d-new", w.raidID, "out")
	if !strings.Contains(out.Message.Content, "Vela Thornbrook is Out") {
		t.Fatalf("press = %q", out.Message.Content)
	}
}

// A Battle.net account never picks; a Discord-only account with nothing picked is asked to pick.
func TestPickRefusals(t *testing.T) {
	w := newPickWorld(t)
	wantReply(t, w.signupAs(t, pressDiscord, opt(optCharacter, directory.PickValue(102))), pickCopy.OwnAccount, false)
	if _, _, err := w.svc.d.Claims.EnsureDiscordAccount(context.Background(), "d-empty", "a raider"); err != nil {
		t.Fatal(err)
	}
	wantReply(t, w.pressAs(t, "d-empty", w.raidID, "in"), pressCopy.PickFirst, false)
}

func TestPickCopyMatchesTheSpec(t *testing.T) {
	row := func(i int) string { return spec(t, "Pick", i, nil) }
	checks := map[string][2]string{
		"Choice":     {pickCopy.Choice, row(0)},
		"Claimed":    {pickCopy.Claimed, row(1)},
		"Taken":      {pickCopy.Taken, row(2)},
		"OwnAccount": {pickCopy.OwnAccount, row(3)},
		"PickFirst":  {pressCopy.PickFirst, row(4)},
		"TooMany":    {pickCopy.TooMany, row(5)},
	}
	for name, c := range checks {
		if c[0] != c[1] {
			t.Errorf("%s = %q, spec says %q", name, c[0], c[1])
		}
	}
}
