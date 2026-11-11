# Round 1 findings: what we change, and what a ruling keeps

Every BLOCKER and SHOULD from round 1 (info-density on the Android tabs, guild-officer on both Raid
tabs, brand-consistency on both apps), with the decision. ACCEPT means both platforms change, through
shared copy keys and tokens. KEEP cites the ruling that decides it; the reviewers get this file in
round 2 and may contest a KEEP.

## Rulings the round-1 verdicts did not have
- **Portraits** (brand B3): design/specs/2026-11-08-app-flows.md lines 142-146 are the owner's exception
  to the one crest language: the profile button and the menu's characters show the Battle.net portrait
  in the faction ring with the faction badge; without one, the button and the menu head show the person
  glyph (the class crest is Home's head and is never drawn twice); the Switch character rows show the
  crest (round 5). KEEP. Both platforms now draw the stand-in while loading and on failure, never an
  empty ring. The Gear medallion is the portrait, else the seal (redesign plan, Task M0
  `PaperdollHeader`): iOS draws the seal where Android draws the portrait, so iOS follows Android.
- **The answer control** (brand B1, B2): the round-5 app mocks (design/mocks/ws_app_flows.py line 90,
  `.r5raid .seg`) supersede spec 2026-11-07 row 143. The web's round-5 Segmented is a group with the
  current one white over the teal/cyan underline (web/src/components/ui/segmented.css), not gold.
  Android keeps the Material 3 segmented button (one 48 dp pill, the chosen one teal, rule 8). iOS
  ACCEPT in part: the four separate buttons, three of them blue, read as four actions; iOS gets a
  `Segmented` primitive in Sources/DesignSystem (one connected group, the chosen one teal, hairline
  dividers, 44 pt, the stacked fallback at accessibility sizes, the R14 haptic), used for the answers.
- **Home's head carries no figures** (info-density Home BLOCKER 1): app-flows line 165, owner ruling.
  The figures go in the Next upgrades section's meta instead: `13 upgrades · +157.6 DPS`, and the
  synced stamp under it. ACCEPT in that form.
- **One job per page** (memory, owner 2026-10-01): Gear owns gear, the planner owns talents. Talents
  rows on Home or Gear, and a guild block in the drawer (the Raid tab is the guild), are KEEP.
- **Android bar title at rest** (brand S2): the round-5 Android mocks put the tab name in the M3 top
  app bar beside the profile button (parity goldens android-raid, -gear, -activity match the mocks).
  KEEP (rule 8, M3 top app bar).
- **Activity times** (brand S3): the inbox buckets ages as the platform notification centres do
  (`8 h ago`, `6:02 pm`, `Tue`, `3 Nov`), the same on both apps. KEEP as a ruling (rule 8); Home's
  `New for you` rows use the same buckets so one row never reads two ways.

## Guild (guild-officer)
| # | Finding | Decision |
|---|---|---|
| G1 | Officers cannot see raid readiness | ACCEPT: an officer line under Who is coming, `4 signed In have not synced this week`, from the guild home roster (`logged_recently`) matched to the raid's In and Late sign-ups |
| G2 | No officer push | The server already sends `raid.short` 24 h out and `guild.applicant` to officers (api/internal/events/catalogue.go). ACCEPT the gap: `raid.short` names the short roles and the not-answered count (`21 of 40 signed, 24 hours out. Healers 5 of 8, 5 not answered.`); push fixtures for `raid.short` and `guild.applicant` join the contract so the apps' captures show them |
| G3 | Needs you below the fold | ACCEPT: for officers, Needs you sits directly under the raid head |
| G4 | Comp is roles only | ACCEPT: a `Covered` line from the raid's `cover[]` / `composition.buffs[]`: what is missing or short, by name, under Who is coming (`Missing Fortitude · Windfury 1 of 2`); nothing when all is covered |
| G5 | Late minutes not on Who is coming | ACCEPT: `Brightwen (Late · 15 min)` before the raid as during it |
| G6 | `raid.starting` push is generic | ACCEPT on the server: the body names what the player's ready check misses (`Starts in 60 minutes. No flask in your bags.`) when the ready check has a miss |
| G7 | `raid.posted` body lacks counts | ACCEPT on the server: `Sign-ups are open: Thu 5 Nov, 8:00 pm CST. 24 of 40 signed.` |
| G9 | Mention cooldown | The Not answered screen is another lane's; verified in code only, reported there |

## Density (info-density, Android; iOS follows for parity)
| Screen | Finding | Decision |
|---|---|---|
| Home | Hero figures | See rulings: section meta and stamp, ACCEPT in that form |
| Home | Raid card role split and countdown on the non-raid-day card | ACCEPT: the card's line gains the role tally (`2/2 · 5/8 · 11/30`, short role in wipe red), one shared tally component with the Raid tab |
| Home | Three rows, no count | ACCEPT: the meta above and `Gear ›` already goes to every upgrade |
| Home | Worn item on upgrade rows | ACCEPT: the source line ends `· replaces <worn>` when the API names the worn item |
| Home | Drop chance on the source line | ACCEPT: `· 4%` when `drop_chance` is set |
| Home | Best run as a row | ACCEPT: one row with the dungeon, its best drop's icon, gain and chance |
| Home | spec_week, movers, guild ladder, import_sim | These are web Home sections; the app's Home follows R0 (raid, upgrades, what changed). KEEP: they reach the player through `New for you` and Activity's Tiers and Guild ladder filters |
| Home | Feed item names in quality colour; the other player's crest | ACCEPT (both Home rows and Activity): `segments[]` styles and `crest` |
| Raid | Cover and buffs | ACCEPT (G4) |
| Raid | Invite time, size, bench and out | ACCEPT: head meta `Invites 7:45 pm · 40`; Who is coming ends with `Bench 1 · Out 4 · 14 not answered` |
| Raid | Officer notes | ACCEPT: the raid's `notes` under the head, two lines, the rest on the raid page |
| Raid | Class-coloured names | ACCEPT: names in the class colour in Who is coming |
| Raid | Melee and ranged split | ACCEPT on the Damage line: `11 of 30 · 5 melee, 6 ranged` |
| Raid | Before you pull icons and counts | ACCEPT: the item icon beside each check where the item is known |
| Raid | Footer rows' figures; last raid line | ACCEPT where the guild home has the figure (`Roster` members, `Loot history` last raid); `Last: Onyxia, 1 kill, 2 wipes` on Loot history's line |
| Gear | Header context | ACCEPT: `+157.6 DPS from 13 upgrades · synced 2 h ago` under the band line; the counter reads `0 of 17 slots match the list` |
| Gear | Sort, gain bars, preset segmented | KEEP: the paperdoll order and the ControlBar summary are the ruled anatomy (redesign plan M3; spec ControlBar row) |
| Gear | Source on the cells | KEEP: the cell is name, slot and gain (spec Paperdoll row); the sheet has the source |
| Gear | Not modelled row | ACCEPT: `2 gains not modelled yet` on the header line |
| Gear | Where to go: upgrades per place | ACCEPT: `2 upgrades · +20.4 DPS` on each place row |
| Activity | Chip counts | KEEP until the feed returns per-group unread counts (needs data, filed) |
| Activity | Per-kind trailing figure | ACCEPT where the row's segments carry the figure (sim DPS) |
| Drawer | Character rows | ACCEPT: `Fury Warrior · 58 · Officer` (guild rank) on the head; item level on the switch rows |
| Drawer | Sync error | ACCEPT: `sync_error` replaces the stamp in wipe red with Try again |
| Drawer | Plan and Devices values | ACCEPT: the plan's renewal or founder date; the newest device's name |

## Brand
| # | Finding | Decision |
|---|---|---|
| S1 | Pushes say `20:00 CST` | ACCEPT on the server: `raidWhenLayout` to `Mon 2 Jan, 3:04 pm MST` |
| S4 | One-off type sizes on Home, Raid, Gear, drawer | ACCEPT: named roles in design/tokens/tokens.json (hero name, card title, fact line, figure unit), regenerated for web, iOS and Android, used by both apps |
| S5 | Hardcoded colours (Paperdoll.kt, GameColors parsePink) | ACCEPT: into tokens |
| S6 | Tallies differ between the apps | ACCEPT: one tally model, `Healers 5 of 8 · 3 short` with tone when the raid has a target, `Healers 6` without one, on both |
| S7 | `Equipped` on Android cells only | ACCEPT: both draw `Equipped` on a worn pick (the information is real; the web paperdoll's worn mode is a different job) |
