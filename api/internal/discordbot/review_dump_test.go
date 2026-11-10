//go:build payloaddump

package discordbot

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// reviewDirEnv names the folder TestDumpReviewPayloads writes every Discord-visible message
// into, for the round-5 review boards (design/mocks/2026-11-09-f-discord/render.mjs): the guided
// cards, the gear panel, /help and the webhook posts. Needs TEST_DATABASE_URL, as the guided
// frames do.
const reviewDirEnv = "REVIEW_PAYLOADS_DIR"

// reviewSite is the public site the webhook posts link to and take the seal from.
const reviewSite = "https://foreversixty.gg"

func TestDumpReviewPayloads(t *testing.T) {
	dir := os.Getenv(reviewDirEnv)
	if dir == "" {
		t.Skip(reviewDirEnv + " is not set")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	d := frameDumper{t: t, dir: dir}
	dumpBisFrames(d)
	dumpCharFrames(d)
	dumpSignupFrames(d)
	dumpOfficerFrames(d)
	dumpPostFrames(d)
	dumpSpecFrames(d)
	dumpHelpFrame(d)
	dumpRaidAndLinkFrames(d)
	for id, m := range gearBoardFrames(t) {
		d.message(id, "gear panel: "+id, m)
	}
	for _, w := range webhookPosts() {
		d.write(dumpFrame{ID: w.id, Title: w.title, Message: webhookMessage(t, w.embed)})
	}
	writeEmojiMap(t, dir)
	writeTiles(t, dir)
}

// writeEmojiMap is emoji.json: each application emoji's name and the image it is drawn from
// (a spec's build icon by its file name, or a quality dot's PNG under emoji/), for the boards.
func writeEmojiMap(t *testing.T, dir string) {
	out := map[string]string{}
	for _, sp := range raidspec.All() {
		out[sp.Emoji()] = "icon:" + sp.Icon
	}
	dots, err := sheet.QualityDots()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "emoji"), 0o755); err != nil {
		t.Fatal(err)
	}
	for _, dot := range dots {
		if err := os.WriteFile(filepath.Join(dir, "emoji", dot.Name+".png"), dot.PNG, 0o600); err != nil {
			t.Fatal(err)
		}
		out[dot.Name] = "emoji/" + dot.Name + ".png"
	}
	raw, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "emoji.json"), append(raw, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}

// writeTiles draws every gear-panel item's thumbnail tile, as the tiles route serves it, into
// tiles/{id}.png.
func writeTiles(t *testing.T, dir string) {
	if err := os.MkdirAll(filepath.Join(dir, "tiles"), 0o755); err != nil {
		t.Fatal(err)
	}
	icons := sheet.NewFileIcons(filepath.Join(gearDataRoot(t), "icons"))
	for id, it := range gearItems(t) {
		icon, err := icons.Icon(it.Icon)
		if err != nil {
			t.Fatal(err)
		}
		png, err := sheet.Tile(icon, it.Quality)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "tiles", strconv.Itoa(id)+".png"), png, 0o600); err != nil {
			t.Fatal(err)
		}
	}
}

func dumpHelpFrame(d frameDumper) {
	w := newLookupWorld(d.t)
	w.svc.now = func() time.Time { return boardClock }
	d.reply("help", "/help", w.command(d.t, "d-stranger", "help", nil))
}

type webhookPost struct {
	id, title string
	embed     notify.Embed
}

// reviewRaid is the boards' raid as a webhook guild's post sees it.
func reviewRaid() notify.RaidPosted {
	return notify.RaidPosted{
		GuildName: "Olympus XXVII", Title: "Onyxia's Lair", Zone: "raid:onyxias-lair",
		StartsAt: boardStart, Timezone: "America/Chicago",
		RaidURL:  reviewSite + "/guild/us/normal/olympus-xxvii/raids/7",
		CrestURL: reviewSite + "/icons/hd/crests/warrior.webp",
		RaidID:   7, Counts: notify.RaidCounts{In: 23, Late: 2, Bench: 1, Out: 3, Tanks: 3, Healers: 6, DPS: 14},
		Size: 40, SizeKnown: true,
	}
}

func webhookPosts() []webhookPost {
	at := boardClock
	moved := boardStart.Add(-24 * time.Hour)
	return []webhookPost{
		{"webhook-test", "webhook: Send test message", notify.TestEmbed(reviewSite,
			map[notify.EventName]bool{notify.EventSimFinished: true, notify.EventUploadParsed: true}, at)},
		{"webhook-sim", "webhook: a server sim finished", notify.SimFinishedEmbed(reviewSite, notify.SimFinished{
			SimID: "abc123", Spec: "warrior-fury", Character: "Kraggor", DPS: 1234.5, Headline: "+41 DPS from Vis'kag",
			Title: "BiS check", CreatedAt: at.Add(-48 * time.Second), FinishedAt: at})},
		{"webhook-log", "webhook: a log finished parsing", notify.UploadParsedEmbed(reviewSite, notify.UploadParsed{
			ReportID: "rep1", Title: "Tuesday Molten Core", Zone: "Molten Core", Visibility: "public",
			Fights: 9, Kills: 7, FinishedAt: at})},
		{"webhook-guild-test", "webhook: the guild's test message", notify.GuildTestEmbed(reviewSite, "Olympus XXVII",
			"/guild/us/normal/olympus-xxvii", at)},
		{"webhook-raid-posted", "webhook: a raid posted", notify.RaidPostedEmbed(reviewRaid(), at)},
		{"webhook-raid-moved", "webhook: a raid moved", notify.RaidChangedEmbed(notify.RaidChanged{
			RaidPosted: reviewRaid(), Was: &moved}, at)},
		{"webhook-raid-cancelled", "webhook: a raid cancelled", notify.RaidChangedEmbed(notify.RaidChanged{
			RaidPosted: reviewRaid(), Cancelled: true}, at)},
		{"webhook-reminder", "webhook: the day-before reminder", notify.WebhookReminderEmbed(reviewRaid(), at)},
		{"webhook-report", "webhook: an officer posts the report", notify.ReportReadyEmbed(notify.ReportReady{
			RaidPosted: reviewRaid(),
			Kills:      []notify.ReportKill{{Boss: "Onyxia", Duration: 102 * time.Second}},
			Wipes:      1, Came: 22, Signed: 23, LootItems: 3, Time: 74 * time.Minute, Deaths: 6,
			Loot: []notify.ReportLoot{
				{Item: "Head of Onyxia", Winner: "Thalwyn"}, {Item: "Eskhandar's Collar", Winner: "Zulmara"},
				{Item: "Onyxia Tooth Pendant", Winner: "Kraggor"},
			},
			TopParses: []notify.ReportParse{
				{Name: "Zulmara", Metric: "DPS", Boss: "Onyxia", Value: 812},
				{Name: "Kraggor", Metric: "DPS", Boss: "Onyxia", Value: 776},
			},
			CoachReady: true,
		}, at)},
	}
}

// webhookMessage is the body a webhook post sends: the seal as the avatar and before the footer.
func webhookMessage(t *testing.T, e notify.Embed) json.RawMessage {
	seal := notify.AvatarURL(reviewSite)
	return marshalled(t, map[string]any{
		"username": notify.Username, "avatar_url": seal, "embeds": []notify.Embed{notify.WithSealFooter(e, seal)},
	})
}

// gearDataRoot is the repo's data folder for the gear panel's build, where its icons are.
func gearDataRoot(t *testing.T) string {
	root, err := filepath.Abs(filepath.Join("..", "..", "..", "data", "builds", gearBuild))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

// gearItems are the gear panel's items by id (testdata/gear-panel).
func gearItems(t *testing.T) map[int]trees.Item {
	raw, err := os.ReadFile(filepath.Join(gearData, gearBuild, "items.json"))
	if err != nil {
		t.Fatal(err)
	}
	var list []trees.Item
	if err := json.Unmarshal(raw, &list); err != nil {
		t.Fatal(err)
	}
	out := map[int]trees.Item{}
	for _, it := range list {
		out[it.ID] = it
	}
	return out
}

// dumpRaidAndLinkFrames are /raid's answer for a raider signed In, and /link before and after.
func dumpRaidAndLinkFrames(d frameDumper) {
	t := d.t
	w := newPressWorld(t)
	w.svc.now = func() time.Time { return boardClock }
	w.svc.d.PublicBaseURL = reviewSite
	w.addChar(t, w.gid, w.uid, keyZulmara, "Zulmara", true)
	w.exec(t, `update guilds set raid_timezone = 'America/Chicago' where id = $1`, w.gid)
	w.exec(t, `update guild_raids set title = 'Onyxia''s Lair', starts_at = $2 where id = $1`, w.raidID, boardStart)
	d.reply("raid-answer", "/raid: the next raid, before signing up", w.raidCmdAs(t, pressDiscord))

	h := newLinkHarness(t)
	ctx := context.Background()
	in := Interaction{DiscordUserID: "d-77"}
	ask, err := h.svc.cmdLink(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	d.reply("link-ask", "/link, not linked", ask)
	if err := h.svc.d.Store.SaveLink(ctx, h.uid, "d-77", "zulmara"); err != nil {
		t.Fatal(err)
	}
	done, err := h.svc.cmdLink(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	d.reply("link-done", "/link, linked", done)
}
