# Forever Sixty: the store listing (CurseForge, Wago Addons, GitHub)

The one source for everything a player reads about the addon on a store. The dashboards
hold a copy of this file: when it changes, paste it again (neither site lets the release
workflow set a description). `md2html.mjs` renders the Description section to HTML for
CurseForge's editor; Wago and GitHub take the Markdown as it is. `gen_gallery.mjs` renders
the gallery plates from the client captures the /setup page already uses.

| Field | Value |
|---|---|
| Project name | **Forever Sixty** (rename from "ForeverSixty"; the folder stays `ForeverSixty`) |
| One line (CurseForge summary, Wago tagline) | Your next upgrade, your next talent and your Forever Score, on the game's own frames. No setup, no network, nothing in combat. |
| Categories | Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Logo | `design/logo/export/foreversixty-icon-1024.png` |
| Source | https://github.com/jhunthrop/foreversixty |
| Issues | https://github.com/jhunthrop/foreversixty/issues |
| Website | https://foreversixty.gg/setup |
| Required dependency (Relations) | Forever Sixty Data (project 1706707; Wago ANz7R564) |
| Gallery | the twelve plates in `gallery/`, in the order below, with these captions |

## Gallery captions

1. **The chip.** One line beside the minimap: the talent to take, the bag item to equip, loot here for you, or what the ready check found missing.
2. **Talent marks.** Your build's talents numbered in order on the talent window, the next one ringed, a strip counting the points your level allows.
3. **Item tooltips.** Whether an item beats what you wear and by how much, the next upgrade for the slot, where it drops, and two alternatives.
4. **Paper doll and bags.** A green arrow on the slot and the bag item that beats it; a gold corner on a slot worth chasing.
5. **The character-panel tab.** Your next five gear targets with their sources, your rotation at your level, and your build's progress.
6. **The player tooltip.** Every player's Forever Score, with the five parts behind it on Shift.
7. **Chat badges.** A sender's score in front of their line in guild, trade and whispers.
8. **The guild tab.** Your roster's level, spec, gear, score, attendance and gear gap, sortable, live for members who run the addon.
9. **The group panel.** Each member's spec, gear gap and score before the first pull; a ready check names who lacks flask, food, imbue or a class buff.
10. **Paste a build.** The one dialog of its own: paste a build from foreversixty.gg and follow it.
11. **Copy your character.** Ctrl+C sends your talents, gear and bags to the site's planner and simulator.
12. **Settings.** Eight switches on the game's own Settings page, each with one line saying what it does.

## Description

## Play your class better, in game

Forever Sixty puts the next thing that makes your character better where you would act on it: on the talent window, the item tooltip, the character panel, the raid frame. There is no window to open and nothing to configure. A new character follows the default build for its spec from the first level; every surface with nothing to say stays out of the way.

It is the in-game half of **[foreversixty.gg](https://foreversixty.gg)**, the planner, simulator and log site for World of Warcraft: Forever. The numbers in the game are the site's numbers: the same simulator, the same stat weights, the same Forever Score.

### What it shows you

- **The chip.** One small icon and line beside the minimap that names your next step: a talent point to spend, an upgrade waiting in your bags (click to equip), loot in the dungeon you just entered, what a ready check found missing, or last raid night's biggest loss. It fades in combat and hides with the UI.
- **Talent marks.** Your build's talents are numbered in order on the talent window, the next one is ringed, and a strip counts the points your level allows ("12 of 15 points · 3 off plan"). Hover a planned node for its place in the plan.
- **Item tooltips.** "Upgrade +4.2 DPS over Cracked Blacksmith Hammer" (healers read HPS), the next upgrade for that slot and the level it lasts until, where it drops with the boss and its chance, two alternatives, and a pair verdict for two one-handers against a two-hander.
- **Paper doll, bags, quest rewards and rolls.** A green arrow on the slot and the bag item that beats it, a gold corner on a slot whose next upgrade is worth chasing, "Best for Arms" on the quest reward that is one, and the tooltip's verdict beside the roll frame.
- **A Forever Sixty tab on the character panel.** Your next five gear targets with the source and drop chance of each, your rotation at your level with the rank to use, and your build's progress. At 60 it adds your simulation and your last raid: each boss's parse, biggest loss, deaths, consumables missed and rotation adherence.
- **The Forever Score.** Every player tooltip carries that player's score and tier, with the five parts behind it on Shift: gear, experience, parses, attendance and preparation. In guild chat, trade and whispers a sender with a score gets a badge before their line. Scores come from the Forever Sixty Data pack, updated nightly from public logs.
- **Guild tab.** On the guild frame: the roster's level, spec, gear, score, attendance and gear gap, sortable, live for members who run the addon. Officers can push a build to the guild.
- **Group panel and ready check.** On the raid frame: each member's spec, level, gear gap and score before the first pull. On a ready check, who lacks flask, food, weapon imbue or a class buff, named.
- **Loot council.** When an epic drops, the master looter sees what it is worth to each member who can wear it, by the simulator, before the roll.

### Builds, in and out

`/fs` opens the one dialog of the addon's own. Paste a build from the site's planner and the addon follows it, previewing what a paste would replace and never overwriting a raid build with a leveling one. `/fs copy` and Ctrl+C send your character, talents, gear and bags, to the site's planner and simulator. Builds you save on the site arrive in game through the desktop companion.

### Commands

| Command | What it does |
|---|---|
| `/fs` | paste or copy a build |
| `/fs copy` | copy my character for the site |
| `/fs build` | switch between saved builds |
| `/fs chip` | show or hide the chip |
| `/fs inspect` | gear and talents of your target |
| `/fs link` | copy a player's site page |
| `/fs group` | your group's specs, gear and buffs |
| `/fs ready` | check flasks, food and buffs now |
| `/fs loot` | what an item is worth to you |
| `/fs floor N` | tag members scored under N |
| `/fs options` | open the settings page |
| `/fs diag` | what is not working |
| `/fs help` | every command |

### What it never does

It never connects to the internet, never reads your account, never sends a message or a kick on your behalf, and does nothing in combat. It reads what the game client already shows you, and the data pack your addon manager updates. Every surface has a switch on the game's Settings page.

### Install both

Forever Sixty Data is a required dependency: it is the nightly ratings pack the scores and badges read. CurseForge and Wago install it with the addon. Installing by hand, take both from the same place.

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt) · [Source and changelog](https://github.com/jhunthrop/foreversixty)
