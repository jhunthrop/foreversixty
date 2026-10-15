# Forever Sixty Data: the store listing (CurseForge, Wago Addons)

The one source for the data pack's listing; paste it into both dashboards when it changes.
`md2html.mjs` renders the Description section to HTML for CurseForge's editor. Its inline
pictures are the main listing's captures, served from foreversixty.gg/addon/.

| Field | Value |
|---|---|
| Project name | **Forever Sixty Data** (rename from "ForeverSixtyData"; the folder stays `ForeverSixtyData`) |
| One line (CurseForge summary, Wago tagline) | The nightly Forever Scores and guild standings from foreversixty.gg, read in game by the Forever Sixty addon. Updated every night. |
| Categories | Data Export, Guild, Tooltip |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Project icon | `design/logo/export/foreversixty-icon-1024.png` (the LX seal; the name tells the two projects apart) |
| Source | https://github.com/jhunthrop/foreversixty |
| Website | https://foreversixty.gg/ratings |
| Relations | Forever Sixty (project 1706689; Wago XKqEdoKy): this pack is its required dependency |
| Gallery | plates 6, 7 and 8 from the main listing's `gallery/`: the player tooltip, the chat badge, the guild tab |

## Description

## The Forever Score, in game

Forever Sixty Data is the ratings pack the **[Forever Sixty](https://foreversixty.gg/setup)** addon reads. It carries every public character's Forever Score and every verified guild's standing from **[foreversixty.gg](https://foreversixty.gg)** into the game, so a player tooltip, a chat line, the guild roster and the raid frame can show how someone performs without anyone leaving the game.

It holds no code. Install it beside Forever Sixty and the addon reads it; on its own it does nothing.

### Install

1. Install it with Forever Sixty. It is the addon's required dependency, so the CurseForge and Wago apps add it for you; by hand, the `ForeverSixtyData` folder goes in `Interface\AddOns` next to `ForeverSixty`.
2. Keep it updating. A new version is published every night; your addon manager picks it up like any other addon.
3. Nothing to configure. The Forever Sixty addon reads it at login and says on its Settings page which night's pack it has: "Player ratings: as of 14 Oct".

### Where it shows up

![A player tooltip with the Forever Score and its five parts](https://foreversixty.gg/addon/player-tooltip-score.jpg)

**On every player tooltip.** "Forever Score 1240 · Rare · as of 13 Oct" in the tier colour; hold Shift for the five parts behind it. A player with no public logs reads "no logs yet", never a zero.

![A guild chat line with the sender's FS badge in front of it](https://foreversixty.gg/addon/chat-badge-guild.jpg)

**In chat.** "FS 1240 ·" before a sender's line in LFG, Trade, guild chat and whispers, muted once the record is more than three days old.

![The guild frame's Forever Sixty roster](https://foreversixty.gg/addon/guild-tab.jpg)

**On the guild frame and the raid frame.** The guild tab's score, attendance and gear-gap columns and the group panel's score column read this pack for every member who does not run the addon.

### What is in it

- **Forever Scores.** One number for a character's whole record, with the five parts behind it: gear (how close to the spec's best set for the level), experience (dungeons cleared and raid bosses killed), parses (how their best on each boss ranks against their spec), attendance (their guild's raid nights in the last 90 days) and preparation (flask, food and enchants on every logged pull). A parse is not a performance; the score says so.
- **Guild standings.** Progression, raid nights and roster size for every guild that raids with logs on the site.
- **The date it was made**, so the addon can say "as of 14 Oct" and tell you when the pack is too old to trust.

### Updated every night

A new version is published each night after the site rates the day's logs, and the numbers in the game match the site's. If the addon's Settings page reads "Player ratings: too old", update this pack; "not installed" means it is missing from `Interface\AddOns`.

### Where the numbers come from

Every score comes from combat logs uploaded to foreversixty.gg with public visibility, rated by the same engine the site uses. Private and guild-only logs are never included. The site's **[ratings page](https://foreversixty.gg/ratings)** explains each part and how small samples are handled.

### What it never does

It never connects to the internet, reads nothing from your account, and does nothing in combat. It is a file the game loads, updated by your addon manager.

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt)
