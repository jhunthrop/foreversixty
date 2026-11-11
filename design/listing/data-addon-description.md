# Forever Sixty Data: the store listing (CurseForge, Wago Addons)

The one source for the data pack's listing; paste it into both dashboards when it changes.
`md2html.mjs` renders the Description section to HTML for CurseForge's editor. Its inline
pictures are the main listing's captures, served from foreversixty.gg/addon/.

| Field | Value |
|---|---|
| Project name | **Forever Sixty Data** (rename from "ForeverSixtyData"; the folder stays `ForeverSixtyData`) |
| One line (CurseForge summary, Wago tagline) | Forever Scores on player tooltips, score badges in chat and the guild tab's scores for the Forever Sixty addon, updated nightly. |
| Categories | Data Export, Guild, Tooltip |
| Game version | 1.60.1 (World of Warcraft: Forever) |
| License | All Rights Reserved |
| Project icon | `design/listing/icons/listing-icon-400.png` (the painted tile: the painted LX seal on the painted blue ground, the same image as the main listing; the name tells the two projects apart) |
| Source | https://github.com/jhunthrop/foreversixty |
| Website | https://foreversixty.gg/ratings |
| Relations | Forever Sixty (project 1706689; Wago XKqEdoKy): this pack is its required dependency |
| Gallery | plates 7, 8 and 9 from the main listing's `gallery/`: the player tooltip, the chat badges, the guild tab |

## Description

## The Forever Score, in game

Forever Sixty Data is the ratings pack the **[Forever Sixty](https://foreversixty.gg/setup)** addon reads. It carries every public character's Forever Score and every verified guild's standing from **[foreversixty.gg](https://foreversixty.gg)** into the game, so a player tooltip, a chat line, the guild tab and the group panel can show how someone performs.

Data only. Install it beside Forever Sixty and the addon reads it; on its own it does nothing.

### Install

1. Install it with Forever Sixty. It is the addon's required dependency, so the CurseForge and Wago apps add it for you; by hand, the `ForeverSixtyData` folder goes in `Interface\AddOns` next to `ForeverSixty`.
2. Keep it updating. A new version is published every night; your addon manager picks it up like any other addon.
3. Nothing to configure. The Forever Sixty addon reads it at login, and the foot of its Settings page says whether the pack is current: "Player ratings (on player tooltips): current".

### Where it shows up

**On every player tooltip.** One line, "Forever Score 1240" in the tier's colour, dated once the record is over a week old; hold Shift for the five parts behind it. A player with no rated fights shows nothing, never a zero.

![A guild chat line with the sender's FS badge in front of it](https://foreversixty.gg/addon/chat-badge-guild.jpg)

**In chat.** "FS 1240 ·" before a sender's line in guild chat, Trade, LookingForGroup and whispers, muted once the record is more than three days old.

![The guild frame's Forever Sixty roster](https://foreversixty.gg/addon/guild-tab.jpg)

**On the guild tab and the group panel.** The guild tab's Gear, Score and Attendance columns and the group panel's Score column read this pack for every member who does not run the addon.

### What is in it

- **Forever Scores.** One number for a character's whole record, with the five parts behind it: gear (how close to the spec's best set for the level), experience (dungeons cleared and raid bosses killed), parses (how their best on each boss ranks against their spec), attendance (their guild's raid nights in the last 90 days) and preparation (flask, food and enchants on every logged pull).
- **Guild standings.** Progression, raid nights and roster size for every guild that raids with logs on the site.
- **The date it was made**, so the addon can tell you when the pack is too old to trust: "too old (made 9 Oct)".

### Updated every night

A new version is published each night after the site rates the day's logs, and the numbers in game match the site's. If the addon's Settings page reads "too old", update this pack; "not installed" means it is missing from `Interface\AddOns`.

### Where the numbers come from

Every score comes from combat logs uploaded to foreversixty.gg with public visibility, rated by the same engine the site uses. Private and guild-only logs are never included, and a player who turns off "Show my Forever Score to others" in the addon is left out. The site's **[ratings page](https://foreversixty.gg/ratings)** explains each part and how small samples are handled.

No network and nothing read from your account: it is a file the game loads, updated by your addon manager.

**[foreversixty.gg](https://foreversixty.gg)** · [Get set up](https://foreversixty.gg/setup) · [Discord](https://discord.gg/NZhBtPwvHt)
