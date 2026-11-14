# Forever Sixty: the store listing (CurseForge, Wago Addons, GitHub)

The one source for everything a player reads about the addon on a store. The dashboards
hold a copy of this file: when it changes, paste it again (neither site lets the release
workflow set a description). `md2html.mjs` renders the Description section to HTML for
CurseForge's editor; Wago and GitHub take the Markdown as it is. `gen_gallery.mjs` renders
the gallery plates (3200x1800 PNG) with the captions below. Each plate uses a lossless
capture from `captures/` when one exists, and otherwise the JPEG the /setup page uses;
`RECAPTURE.md` says how the owner takes the lossless ones. The description's inline pictures
are the /setup captures, served from foreversixty.gg/addon/, and only those that still show
the current surface.

| Field | Value |
|---|---|
| Project name | **Forever Sixty** (rename from "ForeverSixty"; the folder stays `ForeverSixty`) |
| One line (CurseForge summary, Wago tagline) | Play your class better. Your next upgrade and next talent, on the game's own frames. |
| Categories | Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Project icon and account avatar | `design/listing/icons/listing-icon-400.png` (the painted tile: the painted LX seal on the painted blue ground, the same image as the app icon and the Discord server icon) |
| Source | https://github.com/jhunthrop/foreversixty |
| Issues | https://github.com/jhunthrop/foreversixty/issues |
| Website | https://foreversixty.gg/setup |
| Required dependency (Relations) | Forever Sixty Data (project 1706707; Wago ANz7R564) |
| Gallery | the twelve plates in `gallery/`, in the order below, with these captions |

## Gallery captions

1. **The Forever Sixty line.** One line by the minimap with your next step, such as "Next talent at 25", or tonight's raid and its start time on raid day; a click opens the Forever Sixty window on that tab.
2. **The Forever Sixty window.** One window with tabs: Me, Builds, Gear, Rotation, Guild and Officer. Me has your Forever Score, your last raid boss by boss, what to do now and your sim.
3. **Item tooltips.** One line says what the item means for you, "+2.2 DPS", "Best in slot" or "Sidegrade", and Shift shows the rest.
4. **The talent window.** Numbers on the points you still have to take, the next one ringed in gold, and "Next at 25: Lethal Attacks · 12 of 15 · 3 off plan" under the trees.
5. **Bags and paperdoll.** A green arrow on the slot and on the bag item that beats it, a gold corner on a slot with an upgrade to chase.
6. **Quest rewards and rolls.** The green arrow on the reward or roll that is an upgrade, and "Sells for most" on the reward to vendor when none is.
7. **The player tooltip.** Just "Forever Score 1240" in the tier's colour, with the five parts on Shift.
8. **Chat badges.** "FS 1240 ·" before a sender's line in guild chat, Trade, LookingForGroup and whispers.
9. **The Guild tab.** Tonight's raid with your sign-up and job, your prep, the sign-ups, the next raids and where you stand in the guild.
10. **The Officer tab.** The roster with scores and sign-ups, tonight's readiness, a build to push to the guild, and the master loot session ranked by each raider's own gain.
11. **The Builds tab.** Your class's builds with one-click Follow, builds you pick on the site, a code from the planner, and your character to copy back to the site.
12. **Settings.** "Tooltips only" and one switch per surface, on the game's own Settings page.

## Uploading the gallery

The owner uploads by hand; nothing here logs in to either dashboard.

1. Run `node design/listing/gen_gallery.mjs` (Playwright comes from `web/node_modules`). Each
   line it prints must end in `ok`. A `RECAPTURE` line means that plate still uses a JPEG
   or a small crop, so see `RECAPTURE.md`; a plate with no capture at all stops the run. The
   plates are 3200x1800 PNG at 16:9, about 1 MB each. Neither store publishes a gallery size
   limit (checked 2026-11-03), and the generator fails any plate over 5 MB.
2. Delete the old plates on both dashboards first, so the old blurry copies do not stay
   beside the new ones.
3. **CurseForge** (project 1706689 > Gallery): upload `gallery/01-…png` to `12-…png` in
   order. For each image, set the title to the bold words of its caption and the description
   to the rest, then star plate 1 as featured. For **Forever Sixty Data** (1706707), upload
   plates 07, 08 and 09 with the same captions.
4. **Wago** (XKqEdoKy > Gallery): the gallery lists the newest upload first and shows 10 per
   page, so upload in reverse order, `12-…png` first and `01-…png` last. For ANz7R564, upload
   09, then 08, then 07.
5. Open each public page in a private window and check one plate at full size: text and
   frame edges should be sharp. If a store has shrunk the image, its lightbox will show it
   smaller than 3200 wide, but it should still not look blurry.

## Description

### The green arrow and one line on every item

![A green arrow on the character slot and the bag item that beats it](https://foreversixty.gg/addon/bag-arrow.jpg)

One line under any item says what it means for you, such as "+2.2 DPS", "Best in slot" or "Sidegrade". A green arrow marks the bag item, quest reward or loot roll that beats what you wear. The talent window rings your next point, and one panel shows your next upgrades, your rotation at your level and tonight's raid.

A new character follows the default leveling build for its spec, and the raid build from 60, so there is nothing to set up. The numbers are the ones on **[foreversixty.gg](https://foreversixty.gg)**, the planner, simulator and log site for World of Warcraft: Forever: the same stat weights, best in slot lists and Forever Score. Fan-made, not affiliated with Blizzard.

### Install

1. Install **Forever Sixty**.
2. The CurseForge and Wago apps install **Forever Sixty Data**, the nightly ratings pack, with it.
3. By hand, put the `ForeverSixty` and `ForeverSixtyData` folders in `Interface\AddOns`. The `ForeverSixtyInbox` folder in the same download is for the companion app.
4. Log in. A short tour shows what goes where.

For a build of your own, make it in the **[planner](https://foreversixty.gg/planner)**, press Copy for the game, then type `/fs` in game and paste it.

### What's new in 2.0.38

- A short tour on your first login: what the site puts in your game, what your fights send back, your guild's raids, and setting your character up. Skip it any time; `/fs help` opens it again. "Tooltips only" on its first page turns off everything but the tooltip lines.
- A new switch, "Show my Forever Score to others". Off, other players see no score for you; you still see yours.
- The talent window numbers only the points you still have to take, and its strip reads like "Next at 25: Lethal Attacks · 12 of 15 · 3 off plan".
- A talent's tooltip names its points in your build in one line, for example "Your build: point 12".
- The guild tab's count line is just members and online, for example "97 members · 3 online". Alone, the group panel says "Join a group to see its specs, gear and buffs."
- The minimap button wears the flat seal, sharp at its small size.

### The Forever Sixty line

One line beside the minimap with your next step. It always has words: your next talent ("Next talent at 25"), an upgrade in your bags to equip, with its gain, loot for you in the dungeon you are in, what a ready check found missing, or "Forever Sixty" when there is nothing to do. On raid day, from the evening before, it leads with your raid: tonight, the start time and the instance.

- **Click it** to open your panel.
- **Right-click it** to hide it until your next level; `/fs chip` brings it back.
- **Shift-drag it** to move it; Reset positions in Settings puts it back.
- It hides in combat, or stays at 25% if you turn "Hide the Forever Sixty line in combat" off.

### Your panel

Click the line, the Forever Sixty button on your character pane, or your keybind (Key Bindings, "Toggle Forever Sixty") and one panel answers what to do next on this character. Your class crest, level and Forever Score sit at the top; under them, only what has something to say:

- **Tonight**: the raid, its start, your role and your job. The day after, **Last raid**: each boss's percentile and your biggest loss.
- **Upgrades**: up to three, each with where it drops and the level it needs when that is above yours, and an Equip button for one already in your bags.
- **Talents**: points on plan, your next talent, and Change and Talents buttons. Under level 10 with no build, Pick your path loads a spec's leveling build.
- **Rotation at your level**: when to press each button, in plain words, never a spell you do not have yet. At 60 it folds away until you open it.

Opened from the character pane it docks to the pane. Drag its top to move it.

### Item tooltips

One line under the item says what it means for you: the green arrow and the gain ("+2.2 DPS"), "Best in slot", "Best in slot until level 30", "Sidegrade", or the level you can wear it, such as "Mail armor at level 40". Rings and trinkets name the one they replace. Healers read HPS.

At most one more line follows, the next upgrade for the slot or a better pick and where it drops. **Hold Shift** for the rest: the stat breakdown, "Also good:" alternatives, and where the weights came from (the spec's, or your own sim).

### The talent window

Open your talents and your build is drawn on them: a number on every point you still have to take, the next one ringed in gold, and a strip under the trees, "Next at 25: Lethal Attacks · 12 of 15 · 3 off plan". A talent's own tooltip names its points in your build. The client's talent buttons work exactly as before.

### Bags, paperdoll, quest rewards and rolls

- **A green arrow** on a character slot and on the bag item that beats it or fills it.
- **A gold corner** on a slot with an upgrade to chase; hover the slot and it names the pick.
- **Quest rewards**: the green arrow on the reward that is an upgrade, "Best in slot" under your pick, and "Sells for most" under the one to vendor when nothing helps.
- **Group loot rolls**: the green arrow on an upgrade and one short line, "+5.8 DPS · Best in slot", "Sidegrade" or "Not for you".

Another bag addon's bags are left alone, and `/fs diag` says so; the marks on your character pane still show.

### The Forever Score

![A guild chat line with the sender's FS badge in front of it](https://foreversixty.gg/addon/chat-badge-guild.jpg)

Hover a player and their tooltip carries one line, "Forever Score 1240" in the tier's colour, dated once the record is over a week old. Hold Shift for the five parts: Gear, Experience, Parses, Attendance and Preparation. In guild chat, Trade, LookingForGroup and whispers, a sender with a score gets "FS 1240 ·" before their line, muted once the record is three days old. A player with no rated fights shows nothing.

Scores come from the Forever Sixty Data pack, rated every night from public logs on foreversixty.gg. `/fs link` on your target copies their page on the site.

### The guild tab

![The guild frame's Forever Sixty roster](https://foreversixty.gg/addon/guild-tab.jpg)

A Forever Sixty button inside the guild window, beside Invite Member, opens your roster with Level, Spec, Gear, Score, Attendance and Gap. Click a column to sort; Online only filters. Guildmates who run the addon show their live build and gear gap; everyone else, the nightly record.

**Officers** see the approvals waiting on the site and can Push the raid build to guild behind a confirmation. Guildmates get it as a build they can choose; nothing changes for them until they do. Your own row goes live while "Share my build with my guild" is on.

### The group panel and the ready check

On the Raid tab of the Social window (O), or with `/fs group`, one row per member: Spec, Level, iLvl, Gap, Score, Enchants, and whether they run the addon, with their weakest slot under it. Members without the addon are read from inspect, paused in combat.

- After a ready check it lists who lacks flask, food or a weapon buff, and in Molten Core, Onyxia's Lair and Blackwing Lair which buffs, debuffs and resistances the bosses want and who brings them.
- **`/fs ready`** runs the check now, for your eyes only.
- **`/fs floor 1200`** tags the members scored under 1200.

### Loot council

When an epic drops and you are the master looter, the addon asks the group what it is worth to each member who can wear it and lists the answers as they come in, with a 10 second countdown. "Ask the group about rare drops too" adds rares. For yourself, `/fs loot` and a shift-clicked item gives its gain over what you wear and whether it is on your list.

### Builds and your character

`/fs` opens the build dialog on its Builds tab: the builds you picked on the site, and the box for a code from the planner. The preview lists the build's gear at its own level and what it replaces before you press Use this build. `/fs build raid`, `leveling` or `pvp` switches between your saved builds.

`/fs copy` opens Copy my character: your level, talents, gear, bags and professions in one code, selected for Ctrl+C. Paste it on the site's planner or simulator and they start from exactly what you have. With the companion app it travels by itself at logout, and the tab says "Site has your save from 1 h 35 min ago".

On raid day `/fs job` prints your jobs for the raid, and a line names them when you zone in.

### Settings

`/fs options` opens the addon's page on the game's Settings. The minimap button is off on a new install.

| Switch | What it does |
|---|---|
| Tooltips only | Turns off everything but the tooltip lines. |
| Show the Forever Sixty line | Beside the minimap. Shift-drag moves it. |
| Hide the Forever Sixty line in combat | On: the line hides when combat starts. Off: it stays, faded to 25%. |
| Marks on the talent window | Numbers your build's talents in order, rings the next one and adds a progress strip. |
| Marks on your bags and character pane | A green arrow on gear that beats yours, a gold corner on a slot with an upgrade to chase. |
| Show the minimap button | Left-click pastes or copies a build; right-click opens this page. Off: type /fs options or use Options > AddOns. |
| Gear tips on item tooltips | Hold Shift for the details. Off also hides player scores and the quest and roll marks. |
| Score badges in chat | Shows a player's Forever Score before their message in LFG, Trade, guild chat and whispers. |
| Show my Forever Score to others | Off: other players see no score for you. You still see yours. Applies after the companion's next upload. |
| Share my build with my guild | Guildmates with Forever Sixty see your level, spec, build and gear gap. Off: they see nothing. |
| Save my character for the companion app | Written at logout for the companion app on your PC, so the site can read your gear and talents. |
| Ask the group about rare drops too | As the master looter you ask the group about epics. On: rares too, whatever the loot threshold. Off by default. |

The page also lists the first commands, has Reset positions for the line, the minimap button and the build dialog, and says at the foot whether the addon data is current for this patch and the state of the player ratings.

### The companion app, optional

The companion app from **[foreversixty.gg/setup](https://foreversixty.gg/setup)**, for Windows and Mac, uploads your fights from the combat log for reports and the coach, carries your character to the site at logout, and brings back the builds you pick there, your newest sim and your last raid. Without it everything above still works, except Last raid and the sim in your panel.

### Commands

| Command | What it does |
|---|---|
| `/fs` | paste or copy a build |
| `/fs copy` | copy my character for the site |
| `/fs build` | switch between saved builds |
| `/fs chip` | show or hide the Forever Sixty line |
| `/fs inspect` | gear and talents of your target |
| `/fs link` | copy a player's site page |
| `/fs group` | your group's specs, gear and buffs |
| `/fs ready` | check flasks, food and buffs now |
| `/fs job` | your jobs for the next raid |
| `/fs loot` | what an item is worth to you |
| `/fs floor N` | tag members scored under N |
| `/fs options` | open the settings page |
| `/fs diag` | check what is wrong |
| `/fs help` | every command, and the tour again |

### If something looks wrong

- **`/fs diag`** lists what is working and what is not, one line each.
- **"Player ratings (on player tooltips): not installed"** on the Settings page means Forever Sixty Data is missing; install it beside the addon. **"too old"** means your addon manager has not updated it lately.
- **No line by the minimap**: a right-click hid it until your next level, or "Tooltips only" turned it off. `/fs chip` brings it back.
- **No marks in your bags**: another bag addon owns them. The marks on your character pane still show.
- Anything else: **[open an issue](https://github.com/jhunthrop/foreversixty/issues)** with the `/fs diag` output, or ask on **[Discord](https://discord.gg/NZhBtPwvHt)**.

No network (only the companion app talks to the site), no messages or kicks sent for you, and nothing opens in combat.

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt) · [Source and changelog](https://github.com/jhunthrop/foreversixty)
