# Forever Sixty: the store listing (CurseForge, Wago Addons, GitHub)

The one source for everything a player reads about the addon on a store. The dashboards
hold a copy of this file: when it changes, paste it again (neither site lets the release
workflow set a description). `md2html.mjs` renders the Description section to HTML for
CurseForge's editor; Wago and GitHub take the Markdown as it is. `gen_gallery.mjs` renders
the gallery plates (3200x1800 PNG) with the captions below. Every plate's capture comes from
the addon's own mocks: `gen_captures.py` renders the 2.0.48 review boards again into
`captures/` (`RECAPTURE.md` says which, and which shots still wait on the game). The description's inline pictures
are the /setup captures, served from foreversixty.gg/addon/, and only those that still show
the current surface.

| Field | Value |
|---|---|
| Project name | **Forever Sixty** (rename from "ForeverSixty"; the folder stays `ForeverSixty`) |
| One line (CurseForge summary, Wago tagline) | Your next upgrade and next talent, raid sign-ups and Forever Scores, on the game's own frames. |
| Categories | Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Project icon and account avatar | `design/listing/icons/listing-icon-400.png` (the painted tile: the painted LX seal on the painted blue ground, the same image as the app icon and the Discord server icon) |
| Source | https://github.com/jhunthrop/foreversixty |
| Issues | https://github.com/jhunthrop/foreversixty/issues |
| Website | https://foreversixty.gg/setup |
| Required dependency (Relations) | Forever Sixty Data (project 1706707; Wago ANz7R564) |
| Gallery | the ten plates in `gallery/`, in the order below, with these captions |

## Gallery captions

1. **The Forever Sixty line.** One line by the minimap with the thing that matters now, such as "Onyxia's Lair 20:00 · In"; hover it for everything waiting, click it to open that tab.
2. **Me.** Your Forever Score and its five parts, your last raid boss by boss, what waits on the other tabs and your sim.
3. **Builds.** Your class's builds with the one you follow in gold and Follow on the rest, a box for a planner code and your character code for the site.
4. **Gear.** Your gear against the list for your level, the gain on every slot, missing enchants, the top upgrades and where they drop, the stat weights and the sources you farm.
5. **Rotation.** The priority for your level, the stat weights it is played for, and how closely you followed it on each boss of your last raid.
6. **Guild.** Tonight's raid with In, Late, Bench and Out, your job, your prep, the sign-ups, the next raids and your place in the guild.
7. **Officer.** The roster with scores, sign-ups, gear gap and enchants, tonight's readiness, a reminder for chat and a build to push to the guild.
8. **Bags and paperdoll.** A green arrow on the slot and on the bag item that beats it, and the gain on the slot's tooltip.
9. **Quest rewards and rolls.** The green arrow on the reward or roll that is an upgrade, with its gain and the list's own words.
10. **The player tooltip.** "Forever Score 1240" in the tier's colour, dated and greyed when the record is old, the five parts on Shift.

## Uploading the gallery

The owner uploads by hand; nothing here logs in to either dashboard.

1. Run `python3 design/listing/gen_captures.py`, then `node design/listing/gen_gallery.mjs`
   (Playwright comes from `web/node_modules`; if the Python step finds no browser, point
   `PLAYWRIGHT_CHROMIUM` at an installed Chromium). Each gallery line must end in `ok`. The
   plates are 3200x1800 PNG at 16:9, under 2 MB each. Neither store publishes a gallery size
   limit (checked 2026-11-03), and the generator fails any plate over 5 MB.
2. Delete the old plates on both dashboards first, so the old panel and flyout shots do not
   stay beside the new ones.
3. **CurseForge** (project 1706689 > Gallery): upload `gallery/01-…png` to `10-…png` in
   order. For each image, set the title to the bold words of its caption and the description
   to the rest, then star plate 1 as featured. For **Forever Sixty Data** (1706707), upload
   plates 10, 06 and 07 with the same captions.
4. **Wago** (XKqEdoKy > Gallery): the gallery lists the newest upload first and shows 10 per
   page, so upload in reverse order, `10-…png` first and `01-…png` last. For ANz7R564, upload
   07, then 06, then 10.
5. Open each public page in a private window and check one plate at full size: text and
   frame edges should be sharp.

## Description

### One window for your character, marks on the game's own frames

One line under any item says what it means for you, such as "+2.2 DPS", "Best in slot" or "Sidegrade". A green arrow marks the bag item, quest reward or loot roll that beats what you wear, and the talent window numbers your next points. Everything else lives in one window with six tabs: Me, Builds, Gear, Rotation, Guild and Officer. The line by your minimap opens it on whatever needs you.

A new character follows the default leveling build for its spec, and the raid build from 60, so there is nothing to set up. The numbers are the ones on **[foreversixty.gg](https://foreversixty.gg)**, the planner, simulator and log site for World of Warcraft: Forever: the same stat weights, best in slot lists and Forever Score. Fan-made, not affiliated with Blizzard.

### Install

1. Install **Forever Sixty**.
2. The CurseForge and Wago apps install **Forever Sixty Data**, the nightly ratings pack, with it.
3. By hand, put the `ForeverSixty` and `ForeverSixtyData` folders in `Interface\AddOns`. The `ForeverSixtyInbox` folder in the same download is for the companion app.
4. Log in. The window opens once on Me with three setup steps: pick your spec, hover an item you wear, get the companion app. `/fs tour` shows them again.

For a build of your own, make it in the **[planner](https://foreversixty.gg/planner)**, press Copy for the game, then paste it on the Builds tab.

### The Forever Sixty line

Your class crest and one line beside the minimap, the most pressing thing first: "Onyxia's Lair 20:00 · Sign up" on raid day, "Equip Beaststalker's Gloves +2.6 DPS", "Spend 1 point: Intimidation", "2 builds from the site", or how your last raid went. "+1" after it means one more thing waits.

- **Click it** to open the window on that tab. With a talent point to spend it opens the talent window too.
- **Hover it** for everything waiting, each with the tab it opens, and how long until the pull.
- **Right-click it** for Settings. **Shift-drag it** to move it.
- It hides in combat. Turn that off in Settings and it stays, faded, reading "After combat" until the fight ends.

The same window opens from `/fs`, your key binding, the Forever Sixty button on your character pane (on Gear) and the one on the guild frame (on Guild, or Officer for an officer). The minimap button, off on a new install, does exactly what the line does.

### Me

How you are doing. Your Forever Score with its tier, how far to the next, and the five parts, Gear, Experience, Parses, Attendance and Preparation, with the reason under the weakest one ("2 flasks missed"). Your last raid boss by boss: parse, biggest loss, deaths and consumables, then the median, best and worst. Right now lists what waits on the other tabs, each a link to its tab. Your sim shows your DPS, the three items that would gain you the most, and Sim again.

### Builds

Your class's builds, leveling and raid for every spec, with the one you follow in gold and one-click Follow on the rest. Up next lists your next points in order, with the level each one comes at. Paste a planner code to load a build, or copy your character for the site: talents, gear, bags and professions in one code. With the companion app the site gets it at logout, and the tab says when it last did.

### Gear

From level 10, what you wear against the list for your level. Pick a level band, and at 60 Raid-ready or Bare. Every slot shows "Best in slot" or the gain of the list's pick, "in your bags" when you already hold it, and "No enchant" where the pick carries one and yours does not. Upgrades lists the top four with the slot and where each drops, and More upgrades the next few. Under it the stat weights in DPS per point, from your own sim when you have one, and the sources the picks may use: untick PvP or Reputation and every pick, arrow and tooltip changes at once.

### Rotation

What to press. The priority for your level in plain words, with a level stepper to see what changes later, and the stat weights it is played for. With the companion app, How you played it lists every boss of your last raid, least followed first, with the biggest loss on each, and what you lost across the night to downtime and dying early.

### Guild

Raid night in one place. Tonight's raid, its start and how long until it, your sign-up with In, Late, Bench and Out, and your job for the fight. Your prep checks your bags for the attunement, the resistance the fight asks for and the consumables you are missing. Below it the sign-ups by role, the next raids with your answer for each, the guild's top five Forever Scores with your own place, and the build the officers sent.

Raids, sign-ups and jobs come from the guild's page on the site through the companion app; your answers go back the same way.

### Officer

For ranks that can edit officer notes. The roster sorted by sign-up and then score, with attendance, gear gap in DPS, missing enchants, item level, and a dot for who runs the addon right now. Tonight's readiness counts who runs the addon, who is off the raid build, enchants, attunements, flasks and the fight's protection potions, open jobs and the last ready check. Remind sends one line to guild chat about everyone not signed up, and Send to guild pushes the raid build; guildmates choose whether to follow it. Click a name to whisper, Shift-click to invite.

As master looter, looting a boss opens the session here: everyone whose addon answered, ranked by what the item gains them, what they won this week, attendance, whether it is their list's pick, and Give, which hands it out on a second click.

### Item tooltips

One line under the item says what it means for you: the green arrow and the gain ("+2.2 DPS"), "Best in slot", "Best for 20 to 29", "Sidegrade", or the level you can wear it, such as "Mail armor at level 40". Rings and trinkets name the one they replace. Healers read HPS.

At most one more line follows, the next upgrade for the slot or a better pick and where it drops. **Hold Shift** for the rest: the stat breakdown, "Also good:" alternatives, and where the weights came from (the spec's, or your own sim).

### Bags, paperdoll, quest rewards, rolls and talents

![A green arrow on the character slot and the bag item that beats it](https://foreversixty.gg/addon/bag-arrow.jpg)

- **A green arrow** on a character slot and on the bag item that beats it or fills it. **A gold corner** on a slot with an upgrade to chase. Alt-click a marked slot or bag item to open Gear on that slot.
- **Quest rewards**: the green arrow on the reward that is an upgrade, the list's own words under your pick, and "Sells for most" under the one to vendor when nothing helps.
- **Group loot rolls**: the green arrow on an upgrade and one short line, "+5.8 DPS · Best at level 60", "Sidegrade" or "Not for you".
- **The talent window**: a number on every point you still have to take, the next one ringed in gold, and a strip under the trees, "Next at 25: Lethal Attacks · 12 of 15 · 3 off plan". Click the strip to open Builds. The client's talent buttons work exactly as before.

Another bag addon's bags are left alone, and `/fs diag` says so; the marks on your character pane still show.

### The Forever Score

![A guild chat line with the sender's FS badge in front of it](https://foreversixty.gg/addon/chat-badge-guild.jpg)

Hover a player and their tooltip carries one line, "Forever Score 1240" in the tier's colour, dated and greyed once the record is over three days old. Hold Shift for the five parts. In guild chat, Trade, LookingForGroup and whispers, a sender with a score gets "FS 1240 ·" before their line. A player with no rated fights shows nothing.

Scores come from the Forever Sixty Data pack, rated every night from public logs on foreversixty.gg. `/fs link` on your target copies their page on the site.

### Settings

`/fs options` opens the addon's page on the game's Settings.

| Switch | What it does |
|---|---|
| Tooltips only | Turns off the line and every mark, leaving the tooltip lines. |
| Gear tips on item tooltips | Hold Shift for the details. Off also hides player scores and the quest and roll marks. |
| Marks on your bags and character pane | Green arrow: beats what you wear. Gold corner: an upgrade to chase. |
| Marks on the talent window | Numbers your build in order and rings the next talent. |
| Score badges in chat | In LFG, Trade, guild chat and whispers. |
| Show the Forever Sixty line | Beside the minimap. Shift-drag moves it. |
| Hide the Forever Sixty line in combat | Off: it fades to 25% instead. |
| Show the minimap button | Does what the line does. Off on a new install. |
| Share my build with my guild | Guildmates with Forever Sixty see your level, spec, build and gear gap. Off: they see nothing. |
| Show my Forever Score to others | Off: other players see no score for you. You still see yours. Applies after the companion's next upload. |
| Save my character for the companion app | Written at logout, so the site can read your gear and talents. |
| Ask the group about rare drops too | As master looter, rares as well as epics. Off by default. |

The page also lists the commands, has Reset positions for the line, the minimap button and the window, and says at the foot whether the addon data is current for this patch and the state of the Forever Scores.

### What's new in 2.0.48

- Forever Sixty is one window with tabs. The panel, the guild and group pop-outs and the welcome tour are gone; what they showed now sits on the tab it belongs to.
- A gold dot on a tab means something is waiting there: a point to spend, an upgrade in your bags, tonight's raid to sign up for.
- The line by your minimap opens the tab it names; right-click it for Settings.
- One key binding, "Toggle the Forever Sixty window". A key bound to the old panel opens it; a key bound to 2.0.47's window binding needs binding again.
- Gone: `/fs panel`, `/fs group`, `/fs floor`, `/fs job`, `/fs jobs`, `/fs chip` and `/fs welcome`.

### The companion app, optional

The companion app from **[foreversixty.gg/setup](https://foreversixty.gg/setup)**, for Windows and Mac, uploads your fights from the combat log for reports, carries your character to the site at logout, and brings back the builds you pick there, your newest sim, your last raid and your guild's raids, sign-ups and jobs. Without it the tooltips, marks, Builds and Gear still work; Me's last raid and sim, Rotation's last raid and the raid parts of Guild and Officer need it.

### Commands

| Command | What it does |
|---|---|
| `/fs` | open the Forever Sixty window |
| `/fs me`, `builds`, `gear`, `rotation`, `guild`, `officer` | open that tab |
| `/fs copy` | copy my character for the site |
| `/fs link` | copy a player's site page |
| `/fs tour` | the Set up steps on Me again |
| `/fs ready` | check flasks, food and buffs now |
| `/fs loot` | what an item is worth to you |
| `/fs options` | open the settings page |
| `/fs diag` | check what is wrong |
| `/fs help` | every command |

### If something looks wrong

- **`/fs diag`** lists what is working and what is not, one line each.
- **"Forever Scores: not installed"** on the Settings page means Forever Sixty Data is missing; install it beside the addon. **"too old"** means your addon manager has not updated it lately.
- **No line by the minimap**: "Show the Forever Sixty line" or "Tooltips only" in Settings turned it off.
- **No marks in your bags**: another bag addon owns them. The marks on your character pane still show.
- Anything else: **[open an issue](https://github.com/jhunthrop/foreversixty/issues)** with the `/fs diag` output, or ask on **[Discord](https://discord.gg/NZhBtPwvHt)**.

No network (only the companion app talks to the site), and nothing goes to chat or your guild unless you press Remind or Send to guild.

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt) · [Source and changelog](https://github.com/jhunthrop/foreversixty)
