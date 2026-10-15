# Forever Sixty Data: the store listing (CurseForge, Wago Addons)

The one source for the data pack's listing; paste it into both dashboards when it changes.
`md2html.mjs` renders the Description section to HTML for CurseForge's editor.

| Field | Value |
|---|---|
| Project name | **Forever Sixty Data** (rename from "ForeverSixtyData"; the folder stays `ForeverSixtyData`) |
| One line (CurseForge summary, Wago tagline) | The nightly Forever Scores and guild standings from foreversixty.gg, read in game by the Forever Sixty addon. Updated every night. |
| Categories | Data Export, Guild, Tooltip |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Logo | `design/logo/export/foreversixty-icon-1024.png` (the same mark; the name tells them apart) |
| Source | https://github.com/jhunthrop/foreversixty |
| Website | https://foreversixty.gg/ratings |
| Relations | Forever Sixty (project 1706689; Wago XKqEdoKy): this pack is its required dependency |
| Gallery | plates 6, 7 and 8 from the main listing's `gallery/`: the player tooltip, the chat badge, the guild tab |

## Description

## The Forever Score, in game

Forever Sixty Data is the ratings pack the **[Forever Sixty](https://foreversixty.gg/setup)** addon reads. It carries every public character's Forever Score and every verified guild's standing from **[foreversixty.gg](https://foreversixty.gg)** into the game, so a player tooltip, a chat line, the guild roster and the raid frame can show how someone performs without anyone leaving the game.

It holds no code. Install it beside Forever Sixty and the addon reads it; on its own it does nothing.

### What is in it

- **Forever Scores.** One number for a character's whole record, with the five parts behind it: gear (how close to the spec's best set for the level), experience (dungeons cleared and raid bosses killed), parses (how their best on each boss ranks against their spec), attendance (their guild's raid nights in the last 90 days) and preparation (flask, food and enchants on every logged pull). A parse is not a performance; the score says so.
- **Guild standings.** Progression, raid nights and roster size for every guild that raids with logs on the site.

### Updated every night

A new version is published each night after the site rates the day's logs. Your addon manager updates it like any other addon, and the numbers in the game match the site's. The addon dates what it shows ("as of 14 Oct") and says when the pack is too old to trust.

### Where the numbers come from

Every score comes from combat logs uploaded to foreversixty.gg with public visibility, rated by the same engine the site uses. Private and guild-only logs are never included. The site's **[ratings page](https://foreversixty.gg/ratings)** explains each part and how small samples are handled.

### What it never does

It never connects to the internet, reads nothing from your account, and does nothing in combat. It is a file the game loads, updated by your addon manager.

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt)
