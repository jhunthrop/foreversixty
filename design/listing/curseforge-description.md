# Forever Sixty: the store listing (CurseForge, Wago Addons, GitHub)

The one source for everything a player reads about the addon on a store. The dashboards
hold a copy of this file: when it changes, paste it again (neither site lets the release
workflow set a description). `md2html.mjs` renders the Description section to HTML for
CurseForge's editor; Wago and GitHub take the Markdown as it is. `gen_gallery.mjs` renders
the gallery plates (3200x1800 PNG) with the captions below. Every plate's capture comes from
the addon's own mocks: `gen_captures.py` renders the review boards again into
`captures/` (`RECAPTURE.md` says which, and which shots still wait on the game). The description's inline pictures
are the /setup captures, served from foreversixty.gg/addon/, and only those that still show
the current surface.

| Field | Value |
|---|---|
| Project name | **Forever Sixty** (rename from "ForeverSixty"; the folder stays `ForeverSixty`) |
| One line (CurseForge summary, Wago tagline) | A green arrow on the bag item, quest reward or loot roll that beats what you wear, with the gain in DPS. Simulated for Forever at your level and spec. |
| Categories | Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Project icon and account avatar | `design/listing/icons/listing-icon-400.png` (the painted tile: the painted LX seal on the painted blue ground, the same image as the app icon and the Discord server icon) |
| Source | https://github.com/jhunthrop/foreversixty |
| Issues | https://github.com/jhunthrop/foreversixty/issues |
| Website | https://foreversixty.gg/setup |
| Required dependency (Relations) | Forever Sixty Data (project 1706707; Wago ANz7R564) |
| Gallery | the eleven plates in `gallery/`, in the order below, with these captions |

## Gallery captions

1. **Quest rewards and rolls.** The green arrow and the gain on the reward or roll that beats what you wear. "Not for you" on a roll your spec can't use.
2. **Bags and paperdoll.** A green arrow on the slot and on the bag item that beats it, and the gain on the slot's tooltip.
3. **Talent window.** The points of your build still to take numbered on the trees, the next one ringed in gold. The line at the foot says what to spend now and counts points off your build in red.
4. **Builds.** Level 60 and raid builds for every spec. Follow one, or paste a planner code.
5. **Rotation.** What to press at your level, and how closely you followed it on each boss of your last raid.
6. **Gear.** Best in slot for your level against what you wear, with the gain on every slot. The top upgrades say where they drop.
7. **Guild.** Sign up for tonight's raid with In, Late, Bench or Out, see your job, and check your bags for what the fight asks.
8. **Officer.** The roster by sign-up and score, with gear gap and missing enchants. Remind posts one line to guild chat about everyone not signed up.
9. **Player tooltip.** "Forever Score 1240" in the tier's colour, dated and greyed when the record is old, the five parts on Shift.
10. **Your score and last raid.** Your Forever Score and its five parts, then your last raid boss by boss with your parse and biggest loss.
11. **By the minimap.** "Onyxia's Lair 20:00 · In". Hover for everything waiting; click to open that tab.

## Uploading the gallery

The owner uploads by hand; nothing here logs in to either dashboard.

1. Run `python3 design/listing/gen_captures.py`, then `node design/listing/gen_gallery.mjs`
   (Playwright comes from `web/node_modules`; if the Python step finds no browser, point
   `PLAYWRIGHT_CHROMIUM` at an installed Chromium). Each gallery line must end in `ok`. The
   plates are 3200x1800 PNG at 16:9, under 2 MB each. Neither store publishes a gallery size
   limit (checked 2026-11-03), and the generator fails any plate over 5 MB.
2. Delete the old plates on both dashboards first, so the old panel and flyout shots do not
   stay beside the new ones.
3. **CurseForge** (project 1706689 > Gallery): upload `gallery/01-…png` to `11-…png` in
   order. For each image, set the title to the bold words of its caption and the description
   to the rest, then star plate 1 (quest rewards and rolls) as featured. For **Forever Sixty
   Data** (1706707), upload plates 09, 07 and 08 with the same captions.
4. **Wago** (XKqEdoKy > Gallery): the gallery lists the newest upload first and shows 10 per
   page, so upload in reverse order, `11-…png` first and `01-…png` last; plate 11 lands alone
   on page 2. For ANz7R564, upload 08, then 07, then 09.
5. Open each public page in a private window and check one plate at full size: text and
   frame edges should be sharp.

## Description

### The arrow is right for Forever.

A green arrow marks the bag item, quest reward or loot roll that beats what you wear, with the gain: "+2.6 DPS". Every number is simulated for Forever's new talent trees, at your level and spec, not carried over from Classic.

The same simulation rings your next talent, grades your rotation boss by boss from your logs, and runs raid night for your guild: sign-ups, jobs, and who is missing flasks.

![A green arrow on the character slot and the bag item that beats it](https://foreversixty.gg/addon/bag-arrow.jpg)

Open source on **[GitHub](https://github.com/jhunthrop/foreversixty)**, with the changelog and an issue tracker; `/fs diag` lists what works and what does not. Fan-made, not affiliated with Blizzard.

### Install

The CurseForge and Wago apps add **Forever Sixty Data**, the nightly Forever Scores, with it. By hand, put the `ForeverSixty` and `ForeverSixtyData` folders in `Interface\AddOns`; the `ForeverSixtyInbox` folder in the same download is for the companion app. On first login the window opens once with three setup steps (`/fs tour` shows them again).

A new character follows its spec's Level 60 build, and the Raid build from 60, so there is nothing to set up. For a build of your own, make it in the **[planner](https://foreversixty.gg/planner)** on foreversixty.gg, press Copy for the game, then paste it on the Builds tab.

### Bags, paperdoll, quest rewards and rolls

- **A green arrow** on a character slot and on the bag item that beats it or fills it. **A gold corner** on a slot with an upgrade to chase. Alt-click a marked slot or bag item to open Gear on that slot.
- **Quest rewards**: the green arrow on the reward that is an upgrade, the same verdict as its tooltip under your pick ("Best for 20 to 29"), and "Sells for most" when nothing helps.
- **Group loot rolls**: the green arrow on an upgrade and one short line, "+5.8 DPS · Best at level 60", "Sidegrade" or "Not for you".

The arrow also shows in Baganator, Bagnon, AdiBags, ArkInventory and BetterBags. With any other bag addon, `/fs diag` says whether it shows there; the marks on your character pane always do.

### Item tooltips

One line under the item: the green arrow and the gain ("+2.2 DPS"), "Best in slot", "Best for 20 to 29", "Sidegrade", or the level you can wear it, such as "Mail armor at level 40". Rings and trinkets name the one they replace. Healers read HPS.

At most one more line follows, the next upgrade for the slot or a better pick and where it drops. **Hold Shift** for the rest: the stat breakdown, "Also good:" alternatives, and where the weights came from (the spec's, or your own sim).

### The talent window

A number on every point of your build you still have to take, the next talent ringed in gold, and a line at the foot of the window: "Spend 1 now: Bestial Swiftness · 22 of 24 · 1 off plan", or "Next at 25: Lethal Attacks" between levels, with points spent off your build counted in red. Click that line to open Builds. The talent buttons work exactly as before.

### Builds

Every spec's level 60 and raid builds, with the one you follow marked and one-click Follow on the rest. Up next lists your next points in order, with the level each one comes at. Paste a planner code to load a build, or copy your character for the site: talents, gear, bags and professions in one code. With the companion app the site gets it at logout, and the tab says when it last did.

### Rotation

What to press at your level, with a stepper for later levels and the stat weights it is played for. With the companion app, How you played it lists the bosses of your last raid, least followed first, with the biggest loss on each and downtime and early deaths totalled for the night.

### Gear

From level 10, what you wear against the best in slot for your level. Pick a level band, and at 60 Raid-ready or Bare (no raid buffs or consumables). Every slot shows "Best in slot" or the gain of the best pick, "in your bags" when you already hold it, and "No enchant" where the pick carries one and yours does not. Upgrades lists the top four with the slot and where each drops, and More upgrades the next few. Under it the stat weights (the same table as Rotation) and the sources the picks may use: untick PvP or Reputation and every pick, arrow and tooltip changes at once.

### Guild

Tonight's raid, its start and how long until it, your sign-up with In, Late, Bench and Out, and your job for the fight. Under it, what you are missing for the fight: the attunement, the resistance it asks for, flasks, food and potions. Below it the sign-ups by role, the next raids with your answer for each, the guild's top five Forever Scores with your own place, and the build the officers sent.

Raids, sign-ups and jobs come from the guild's page on the site through the companion app; your answers go back the same way.

### Officer

For ranks that can edit officer notes. The roster sorted by sign-up and then score, with attendance, gear gap in DPS, missing enchants, item level, and a dot for who runs the addon right now. Tonight's readiness counts who runs the addon, who is off the raid build, enchants, attunements, flasks and the fight's protection potions, open jobs and the last ready check. Remind sends one line to guild chat about everyone not signed up, and Send to guild pushes the raid build; guildmates choose whether to follow it. Click a name to whisper, Shift-click to invite.

As master looter, looting a boss opens the session here: everyone whose addon answered, ranked by what the item gains them, what they won this week, attendance, whether it is their best pick, and Give, which hands it out on a second click.

### The Forever Score

![A guild chat line with the sender's FS badge in front of it](https://foreversixty.gg/addon/chat-badge-guild.jpg)

Rated every night from public logs on foreversixty.gg and carried in by the Forever Sixty Data pack. Hover a player and their tooltip carries one line, "Forever Score 1240" in the tier's colour, dated and greyed once the record is over three days old. Hold Shift for the five parts: Gear, Experience, Parses, Attendance and Preparation. In guild chat, Trade, LookingForGroup and whispers, a sender with a score gets "FS 1240 ·" before their line. A player with no rated fights shows nothing. `/fs link` on your target copies their page on the site.

### The window

Six tabs: Me, Builds, Gear, Rotation, Guild and Officer. A gold dot on a tab means something waits there. Me shows your Forever Score and its five parts with the reason under the weakest one ("2 flasks missed"), your last raid boss by boss (parse, biggest loss, deaths and consumables), and your sim's DPS with the three items that would gain you the most.

The window opens from `/fs`, your key binding, the button on your character pane (on Gear), the one on the guild frame (on Guild, or Officer for an officer), and the line by your minimap.

### By the minimap

Your class crest and one line beside the minimap, the most pressing thing first: "Onyxia's Lair 20:00 · Sign up" on raid day, "Equip Beaststalker's Gloves +2.6 DPS", "Spend 1 point: Intimidation", or how your last raid went. "+1" after it means one more thing waits.

- **Click it** to open the window on that tab. With a talent point to spend it opens the talent window too.
- **Hover it** for everything waiting, each with the tab it opens, and how long until the pull.
- **Right-click it** for Settings. **Shift-drag it** to move it.
- It hides in combat. Turn that off in Settings and it stays, faded, reading "After combat" until the fight ends.

The minimap button, off on a new install, does the same.

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

### What's new in 2.0.49

- Gear and Rotation show the same stat weights table. Gear's hit line says Ranged hit or Melee hit for your spec, and ranged hunters no longer see a spell hit line.

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
- **No arrow in your bags**: `/fs diag` names your bag addon and says whether the arrow shows in it. The marks on your character pane still show.
- Anything else: **[open an issue](https://github.com/jhunthrop/foreversixty/issues)** with the `/fs diag` output, or ask on **[Discord](https://discord.gg/NZhBtPwvHt)**.

The addon makes no network calls (the companion app talks to the site) and writes nothing to chat unless you press Remind. Guildmates running Forever Sixty see your level, spec, build and gear gap until you turn off "Share my build with my guild".

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt) · [Source and changelog](https://github.com/jhunthrop/foreversixty)
