# Guild tab: response to review round 4 (2026-11-11)

Verdicts: ux-designer SHIP, wow-player SHIP, guild-officer WORTH PAYING FOR, ai-slop FAIL on one line (now fixed, below).
Round 2's info-density DENSE and brand-consistency CONSISTENT stand.

## Fixed in round 5

| Item | Fix |
|---|---|
| ai-slop 1: the `Assigning is on the site.` line | The line and its key are gone; the section's link reads `Assign in planner ›`, which says where assigning happens in the one tap the brief asked for (spec 2.3a updated) |
| wow-player 1: `41 DPS` on a Restoration Druid | A gap reads in the addon's unit (Comms.gapUnit): a healer's in HPS, any other known spec's in DPS, `pts` with no spec; tests on both apps |
| ux should A: Android navigating rows had no chevron | `FlowRow(chevron = true)` on this lane's rows: Recent logs, Progression, Pages, the officer rows, Search guilds, Claim your guild |
| ux observation 2: Android renders two thirds empty | The full renders are cropped to their content |
| guild-officer should 2 / ai-slop: the gold award not seen | `parity/ios-crop-loot-yours.png`: Thandrel in gold on Ancient Cornerstone Grimoire, Kraggor in class colour |

## Follow-ups, named for their owners (not this lane's code)

| Owner | Item |
|---|---|
| Raid lanes (the shared at-60 raid block) | Needs you: drop the `6 open` count and the `Healers short by 3` line (Who is coming carries it); Before you pull's stamp reads `9 h ago` (ai-slop rounds 2 to 4) |
| Raid lanes | `Missing Fort · WF 1 of 2` needs its unit; Android Who is coming names lose class colour where iOS keeps it (wow-player 6) |
| Raid lanes | The fixture echo keeps the canned sign-up raid, so Late shows a short roster in fixture mode; the server keeps the rest (ux B). `Saved.` reserves its line (ux C) |
| Contract fixtures | Seed `not_signed_raiders` on `guild.raid.officer` so the device walk can tap Mention (guild-officer should 1); real Onyxia drops in `guild.loot.nights` and the guild-tab test data (wow-player 3) |
| Guild pass | `· you` beside your own award (wow-player 4); `ilvl` for `item level` (wow-player 5, ai-slop 2) to agree with the web roster first; officer order Missing raids and Gear gaps above Recently on (guild-officer round 3) |
| Owner ruling stands | A level-36 member signs up live (coordinator ruling 2026-11-10: allowed, with the quiet line); the guild head leads under 60 and follows the raid at 60 (spec 2.1/2.2) |
