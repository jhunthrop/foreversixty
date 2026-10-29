package discordbot

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// The boards' rosters (gen_guided_cards.mjs ROSTER_33, ROSTER_8, ROSTER_40), so the real
// drainer draws the same sign-ups the designer drew.
func rl(name, class, spec, role, state string) notify.RosterLine {
	return notify.RosterLine{Name: name, Class: class, Spec: spec, Role: role, State: state}
}

func boardRoster33() []notify.RosterLine {
	const t, h, d = "tank", "healer", "dps"
	return []notify.RosterLine{
		rl("Kraggor", "warrior", "Protection", t, "in"), rl("Ironvale", "warrior", "Protection", t, "in"),
		rl("Bramblehide", "druid", "Feral Bear", t, "in"), rl("Lightwell", "priest", "Holy", h, "in"),
		rl("Dawnmere", "priest", "Holy", h, "in"), rl("Embertide", "priest", "Discipline", h, "in"),
		rl("Aurelith", "paladin", "Holy", h, "in"), rl("Halcyra", "paladin", "Holy", h, "in"),
		rl("Mossgrove", "druid", "Restoration", h, "in"), rl("Vaelora", "druid", "Restoration", h, "in"),
		rl("Sablewind", "shaman", "Restoration", h, "in"), rl("Zulmara", "hunter", "Beast Mastery", d, "in"),
		rl("Windtalon", "hunter", "Marksmanship", d, "in"), rl("Brightshot", "hunter", "Marksmanship", d, "in"),
		rl("Kelthar", "hunter", "Survival", d, "in"), rl("Pyrewisp", "mage", "Fire", d, "in"),
		rl("Cindrel", "mage", "Fire", d, "in"), rl("Frostwhisk", "mage", "Frost", d, "in"),
		rl("Stormvale", "mage", "Frost", d, "in"), rl("Orivane", "mage", "Frost", d, "in"),
		rl("Sundermark", "paladin", "Retribution", d, "in"), rl("Nightbrand", "priest", "Shadow", d, "in"),
		rl("Corvane", "rogue", "Combat", d, "in"), rl("Duskmantle", "rogue", "Combat", d, "in"),
		rl("Ravenmoor", "rogue", "Assassination", d, "in"), rl("Ashenra", "shaman", "Enhancement", d, "in"),
		rl("Grimhollow", "warlock", "Destruction", d, "in"), rl("Talvorn", "warlock", "Destruction", d, "in"),
		rl("Vexlash", "warlock", "Affliction", d, "in"), rl("Wyrmbane", "warrior", "Fury", d, "in"),
		rl("Thornquill", "druid", "Balance", d, "in"), rl("Glimmerfen", "warlock", "Destruction", d, "late"),
		rl("Hollowmere", "rogue", "Combat", d, "late"), rl("Quillfeather", "hunter", "Marksmanship", d, "bench"),
		rl("Marrowind", "hunter", "Marksmanship", d, "bench"), rl("Selwyn", "mage", "Frost", d, "bench"),
		rl("Fennrick", "warlock", "Affliction", d, "out"), rl("Ysolde", "priest", "Shadow", d, "out"),
		rl("Garrowe", "warrior", "Arms", d, "out"), rl("Lunetta", "mage", "Fire", d, "out"),
	}
}

func boardRoster8() []notify.RosterLine {
	return []notify.RosterLine{
		rl("Kraggor", "warrior", "Protection", "tank", "in"), rl("Zulmara", "hunter", "Beast Mastery", "dps", "in"),
		rl("Pyrewisp", "mage", "Fire", "dps", "in"), rl("Corvane", "rogue", "Combat", "dps", "in"),
		rl("Frostwhisk", "mage", "Frost", "dps", "in"), rl("Vexlash", "warlock", "Affliction", "dps", "in"),
		rl("Wyrmbane", "warrior", "Fury", "dps", "in"), rl("Glimmerfen", "warlock", "Destruction", "dps", "late"),
		rl("Garrowe", "warrior", "Arms", "dps", "out"),
	}
}

func boardRoster40() []notify.RosterLine {
	var out []notify.RosterLine
	for _, r := range boardRoster33() {
		if r.State == "in" {
			out = append(out, r)
		}
	}
	return append(out,
		rl("Stonebrow", "warrior", "Protection", "tank", "in"), rl("Lumen", "priest", "Holy", "healer", "in"),
		rl("Tidecall", "shaman", "Restoration", "healer", "in"), rl("Hollowmere", "rogue", "Combat", "dps", "in"),
		rl("Selwyn", "mage", "Frost", "dps", "in"), rl("Ashgrave", "warlock", "Destruction", "dps", "in"),
		rl("Quillfeather", "hunter", "Marksmanship", "dps", "in"), rl("Briarthorn", "warrior", "Fury", "dps", "in"),
		rl("Glimmerfen", "warlock", "Destruction", "dps", "late"), rl("Marrowind", "hunter", "Marksmanship", "dps", "bench"),
		rl("Fennrick", "warlock", "Affliction", "dps", "bench"), rl("Embervale", "mage", "Fire", "dps", "bench"),
		rl("Ysolde", "priest", "Shadow", "dps", "out"), rl("Lunetta", "mage", "Fire", "dps", "out"),
	)
}

// countsOf are the counts a roster makes; the roles count In rows only.
func countsOf(roster []notify.RosterLine) notify.RaidCounts {
	var c notify.RaidCounts
	for _, r := range roster {
		switch r.State {
		case "in":
			c.In++
			switch r.Role {
			case "tank":
				c.Tanks++
			case "healer":
				c.Healers++
			default:
				c.DPS++
			}
		case "late":
			c.Late++
		case "bench":
			c.Bench++
		case "out":
			c.Out++
		}
	}
	return c
}

// boardMarks is the board's emoji ids: 1301000000000000001 on, in specs.json order.
func boardMarks() EmojiAPI {
	api := &countingEmoji{}
	for i, sp := range raidspec.All() {
		api.items = append(api.items, AppEmoji{ID: fmt.Sprint(1301000000000000001 + i), Name: sp.Emoji()})
	}
	return api
}

// rosterWorld is a post world on the boards' night with the board's emoji and roster.
func rosterWorld(t *testing.T, roster []notify.RosterLine) *postWorld {
	t.Helper()
	w := newPostWorld(t)
	w.svc.emoji = &EmojiRegistry{API: boardMarks()}
	setRoster(w, roster)
	return w
}

func setRoster(w *postWorld, roster []notify.RosterLine) {
	boardRaid(w, countsOf(roster))
	w.facts.edit(w.raidID, func(f *guilds.RaidPostFacts) { f.Post.Roster = roster })
}

func lastEdit(w *postWorld) Message {
	edits := w.callsOf("EditMessage")
	return edits[len(edits)-1].Message
}

// dumpPostFrames drives the real drainer through the boards' posts: scheduled, partly signed,
// the reminder as an edit and as a reply, full, and the webhook reminder.
func dumpPostFrames(d frameDumper) {
	t := d.t
	created := func(id, title string, roster []notify.RosterLine) {
		w := rosterWorld(t, roster)
		w.postedOnce(t)
		d.write(dumpFrame{ID: id, Title: title, Message: stamped(marshalled(t, w.callsOf("CreateMessage")[0].Message))})
	}
	created("post-scheduled", "posted on scheduling: nobody yet", nil)
	created("post-partial", "partly signed: 8 of 40, no healer yet", boardRoster8())
	created("post-full", "full: 40 of 40, bench and out", boardRoster40())

	w := rosterWorld(t, boardRoster33())
	w.postedOnce(t)
	w.rest.Latest = "m1"
	w.remindDue(t)
	d.write(dumpFrame{ID: "post-tomorrow", Title: "the reminder as an edit of the post", Message: stamped(marshalled(t, lastEdit(w)))})

	r := rosterWorld(t, boardRoster33())
	r.postedOnce(t)
	r.rest.Latest, r.rest.ReplyID = "other", "r1"
	r.remindDue(t)
	d.write(dumpFrame{ID: "post-reminder", Title: "the reminder as a reply under the post",
		Above: stamped(marshalled(t, lastEdit(r))), Message: marshalled(t, r.replies()[0].Message)})

	posted := r.facts.by[r.raidID].Post.RaidPosted
	hook := webhookOf(t, r.st.Pool, r.gid, r.uid, r.poster) // binding the bot dropped the webhook; testStore also resets w
	if _, sent, err := hook.RaidReminder(context.Background(), r.gid, posted); err != nil || !sent {
		t.Fatalf("webhook reminder = %v, %v", sent, err)
	}
	contents := r.poster.Contents()
	last := contents[len(contents)-1]
	d.write(dumpFrame{ID: "post-reminder-webhook", Title: "the reminder through a webhook",
		Message: stamped(marshalled(t, Message{Content: last.Content, Embeds: last.Embeds}))})
}

// dumpSpecFrames: Aurelith, a paladin with no spec, presses In on the post (33 of 40 already
// signed), picks Holy, and the outcome lands in place with Change spec.
func dumpSpecFrames(d frameDumper) {
	t := d.t
	w := newPressWorld(t)
	w.svc.now = func() time.Time { return boardClock }
	const aurelith = "us/normal/aurelith"
	w.addCharNoSpec(t, w.gid, w.uid, aurelith, "Aurelith", true)
	w.exec(t, `update characters set class = 'paladin' where key = $1`, aurelith)
	w.exec(t, `update guilds set raid_timezone = 'America/Chicago' where id = $1`, w.gid)
	w.exec(t, `update guild_raids set starts_at = $2 where id = $1`, w.raidID, boardStart)
	w.exec(t, `insert into guild_raid_signups (raid_id, character_key, name, class, spec, role, state, signed_by, acted_by)
		select $1, 'us/normal/other' || n, 'Other' || n, 'mage', 'Frost', 'dps', 'in', $2, $2 from generate_series(1, $3) n`,
		w.raidID, seedUser(t, w.st.Pool), signedUpOfForty)
	ask := w.pressAs(t, pressDiscord, w.raidID, "in")
	d.reply("spec-ask-card", "In pressed, the spec unknown: the one-select card", ask)
	pick := ask.Message.Components[0][0].CustomID
	d.reply("spec-ask-done", "the pick: the sign-up outcome in place, Change spec under it", w.mustSpecPress(t, pick, "paladin-holy"))
}
