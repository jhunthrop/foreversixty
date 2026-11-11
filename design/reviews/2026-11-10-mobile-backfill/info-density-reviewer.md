# info-density-reviewer: mobile backfill (Android tabs)

## Round 1 (2026-11-10, base ac2f37dec)

Android Home, Raid, Gear, Activity and the profile drawer are all THIN: each leaves API data it already fetches unrendered. I read the captures, the Kotlin models and the app fixtures, and checked the iOS and web captures only for platform differences. Nothing below repeats the BRIEF list.

## Home (HomeFlowScreen, HomeHero, HomeComposer, HomeRaidCard)

The page's job is "what do I do now": the next upgrade, tonight's raid, what changed. The hero is a banner holding a name and a level line, and nothing else.

**BLOCKER**
- **Hero figures.** The hero should carry the character's headline numbers in HomeHero.kt, as a trail line under "Level 60 Fury Warrior · Olympus XXVII". The data is `/v1/characters/{key}/home` `total_gain` ("+157.6 DPS across 13 upgrades" from `upgrade_count`, a Gear link) and `talents.state` with `points_differ` ("Talents differ by 1 point", linking to the planner). `synced_at` and `synced_from` give a quiet "Synced 2 h ago, Battle.net" stamp. `ApiModels.kt` parses `upgradeCount` and `totalGain`, but Home never renders them (`totalGain` is only used in `WidgetFigures.kt`).
- **Raid card.** "18 in, 3 late · you have not answered" has no role split and no countdown. Add `next_raid.counts.tanks/healers/dps` as a "2/2 · 5/8 · 11/30" trail, the pull countdown that Raid already computes, and a "Healers 5 of 8" warning in red. The data is `guild.home.next_raid.counts` and `size`. Raid already builds this, so Home should share the composable.
- **Three upgrade rows is too few.** The card shows three and cuts off, while `upgrade_count` can be 13. Put "+157.6 DPS, 13 upgrades" and a "See all 13" link in the Next upgrades section header (HomeFlowScreen.kt).

**SHOULD**
- **Upgrade rows lack the worn item.** `upgrades[].worn` (name, item_level, `known`) answers "replaces what". Add "Replaces Vest of Elements · ilvl 63" under the name in the upgrade row. A pick's own item_level (`pick.item_level`) belongs there too. The Gear item sheet has `worn` (36 hits in the app), so Home drops it where Gear shows it.
- **Drop chance.** `pick.drop_chance` (parsed; a hit in `ItemPageModel.kt`) goes on the source line, e.g. "Blackcrow · 8%". Today only the "Best run" footer carries it, as unlabelled "8%".
- **"Best run: Blackrock Spire, Blackcrow +1.9 DPS, 8%"** is a bold sentence with no structure. Make it a row: dungeon icon, name, best drop, gain, chance.
- **`/v1/home` `spec_week`.** Place A, 4 of 23, up from B last week (fields `grade`, `last_week_grade`, `place`, `of`). It is parsed in `ApiModels.kt` and drawn nowhere on Android. Add one dense row "Fury Warrior is A, was B · #4 of 23" linking to Tiers.
- **`/v1/home` `movers`.** The `rows` carry `rank`, `name` and `delta`, for a "Who moved" strip of three names with ± arrows.
- **Guild ladder.** The `guild_sims` row carries `you.rank` (7 of 14) and `rows[0].value` (the leader). Show "You are #7 of 14 · 114 DPS behind Thalrin". Today it appears only as feed prose.
- **`import_sim` (dps, simmed_at).** Show "Last sim 26.2 DPS, 4 days ago" with a re-run link.
- **New for you.** Rows have no section tag or timestamp format consistent with Activity, and the list stops at three. Reuse `ActivityRow`'s kind icon and add a "+N more" in the header (already "All activity (3)").

**NICE**
- Quality-colour the item names in the feed rows. `feed.rows[].segments[]` has `style:item` and `quality`, and the rows currently render plain white.
- iOS Home shows the "You are In" state line ("Late, 30 min" echo) and a source line "Crafted: Blacksmithing". The Android card is missing the "Crafted:" prefix from `source_line`, and prints "Blacksmithing" bare.

## Raid tab (RaidTabScreen, RaidTabParts, RaidTabModel)

The job is: get ready, be answered, and see who is coming. It is the thinnest tab for what `guild.raid` returns. The model reads `counts` roles and `signups`, and leaves everything else unrendered.

**BLOCKER**
- **Raid cover and buffs.** `guild.raid.cover[]` (windfury, grace of air and so on: `need`, `providers_in`, `providers_late`, `providers_bench`, `state`) and `composition.buffs[]` (`label`, `icon`, `present`, `count`) are a raid leader's first question: "do we have Fort, MotW, WF, soulstones". Add a "Covered" section between Who is coming and Before you pull. Show spell icons with a green or red state, and a short count for a gap ("WF 1 of 2, one late"). The data is in the fixture `guild.raid.json` and unused in raids/flows (only NotAnswered mentions bench or late). The web raid page shows it.
- **Raid facts are missing.** The card has the time and the countdown only. Add `invite_at` ("Invites 7:45 pm"), `duration_minutes` and the `size` of 40 to the head meta line. Then bench and out, from `counts.bench` and `counts.out`. At the moment "Damage 11 of 30" is the only place size shows, and bench and out are not shown anywhere on the tab.
- **Officer notes.** `guild.raid.notes` (or `notes_html`) is unrendered. "Bring fire resistance gear and flasks" is exactly what a raider reads before pulling. It belongs in the card header, trimmed to two lines, with a link to the raid page.

**SHOULD**
- **Who is coming.** "Thandrel, Kaelra, Thalrin and 8 more" is a list of names with no class colour and no spec icon. `signups[]` has `class`, `spec_icon` and `state`. Render class-coloured names with small spec icons, and the Late and Bench counts as chips (Late 3 · Bench 1). Show the unanswered count from `not_signed[]` ("14 not answered") when the viewer is an officer, or as a plain count otherwise.
- **Composition counts.** `counts.melee` and `counts.ranged` (and `composition.classes`) can follow the Damage line: "5 melee, 6 ranged".
- **Before you pull.** The checklist is good. Each row should link out: the missing flask to its item page, the enchant to its source (the "Where to get it" link already does this). Add the item icon beside the check (all four are items or spells with `icon` fields in `cdn.items.*` and `cdn.item-uses.json`), and a count ("3 in your bags") in a figure column.
- **Footer links list.** "Roster, Calendar, Loot history, Guild sims" are plain link rows. `summary` figures could show: `raiders`, `pulls_this_tier`, `named_encounters_down`. The "Last raid: 4 pieces" row needs only one figure, so give each row a number or drop it.
- **Last raid line.** `guild.raids.rows[]` (`kill_count`, `wipe_count`, `top_dps`, `deaths`) can sit in the footer: "Last: Onyxia, 1 kill, 2 wipes · top Morvane 697 DPS".

**NICE**
- Attendance: `me.characters[].guild.attendance` (6 of 8 nights) as "Attended 6 of 8" next to the answer.
- Raid dead state (`android-raid-no-raid`, `-none`): check that it lists the next raid date and the guild calendar link, not just a blank head.

## Gear (GearFlowScreen, GearComposer, GearItemSheet)

The job is gear and upgrades. It is the densest of the four, but the header number carries no context.

**BLOCKER**
- **Header context.** "1004.2 DPS" has no delta, and the line "4 of 17 slots on the list" sits next to 12 slots showing gains. Recheck the counter against `upgrade_count` and `already_best[]` (it must agree with the visible rows). Next to the DPS figure, add `total_gain` ("+157.6 if you take every upgrade") and `synced_at`, the same stamp as the hero. If the list shows what the row shows, the counter is wrong.

**SHOULD**
- **Sort.** Slots render in paperdoll order only. Offer sort by gain (the upgrade rows) and a "Best first" ranking. The `gain` field on each row is already there.
- **Gain bars.** Each gain is a bare "+14.5 DPS". Add a share bar or a rank chip (#1 to #8 upgrades) scaled to `total_gain`, the Gear equivalent of the web table's share bars.
- **Slot rows.** Add the source line and drop chance (`pick.source_line`, `drop_chance`) under the slot name, instead of only inside the sheet. Add the worn item and the ilvl change in the sheet header.
- **Talents.** `talents.state` and `points_differ` (e.g. "differs, 1 point") are not on the tab. Gear is not the planner, so show one status row linking out, not the tree.
- **`not_modelled_count` and `not_sim_checked_count`.** The "Gain not modelled yet" items (`modelled=false`, `gain_words`) need a row in the list, as the web BiS does (`BisScreen.kt` reads `notModelled`). Android Gear drops it.
- **Where to go tonight.** The rows are three lines of text. `cdn.loot.json` gives each boss its drop table and `drop_chance`: show the number of upgrades the dungeon offers, "2 upgrades · +20.4 DPS", with the best drop's icon.
- **Preset switch.** "60 raid · Alliance · Change" is a link. The preset and band (`preset`, `band_label`) are already sent: make Raid and Bare a segmented control on the tab, not a link to another screen.

**NICE**
- Item level on every row (`pick.item_level`), small, right-aligned.
- iOS Gear shows the same list with a Horde crest in the header; Android's header avatar is the character portrait. That is fine, but make sure both draw quality borders the same.

## Activity (ActivityScreen, ActivityModel)

The job is a log of what changed. It is readable, but short on levers.

**SHOULD**
- **Filter chips carry no counts.** Add per-group unread counts from the feed (`group`, `read`), e.g. "Raids 1". `activity.list.json` has `unread: 3` only in total, so counts need a per-group tally on the client.
- **Row body is the server string only.** `kind`, `group` and `link` exist per row. Show a kind-specific trailing figure: sim DPS and delta in tabular numerals on the right (as Home's upgrades do), a percentile chip on reports ("87th"), a +DPS chip on loot.
- **Item names.** "You got Deathbringer" and "Fury Warrior BiS changed" are plain. Use `feed.rows[].segments[]` (`style:item`, `quality`, `item_id`) for quality colour. This is the biggest single lift for "the game's own richness".
- **Class crest.** Rows for other players ("Kaelra passed you") should use `crest` from `feed.rows[]` (warrior), not the player's own crest. It is the same icon here as the hero, so it reads as "you".
- **Grouping.** "Today / Yesterday / This week" is fine. Add the guild name beside the raid rows and a day-of-week for "This week" rows, which only show "Tue".

**NICE**
- Swipe or long-press to mark one read; "Mark all read" is the only control.
- Empty state (`android-activity-empty`): list what appears here (raids, sims, logs, ladder, tiers) with each filter's icon, not blank copy.

## Profile drawer (ProfileMenu, MenuScreens, MenuModel)

The job is switching character and reaching the account.

**SHOULD**
- **Character rows.** Each row is crest, name, spec and level. Add from `/v1/me` `characters[]`: `item_level` ("ilvl 66"), guild rank (`guild.rank`: Officer) and attendance (`guild.attendance` 6 of 8). Add a +DPS figure from the character's `home.total_gain` if cached. The header should show "Officer · 6 of 8 nights" for the current character.
- **Guild block.** The drawer has no guild entry. `me.guilds[]` carries `name`, `rank`, `roster_count` and `verified_count` ("38 members, 32 verified"). Add a Guild section above Tools with one row linking to the guild page (shared with the Raid tab footer's Roster link).
- **Sync status.** `characters[].build.captured_at`, `median_sync_gap_sec` and `sync_error` back a Refresh status line. The head says "Updated 2 h ago", so add the error case in red and the gap ("syncs about every 90 min").
- **Account rows.** The rows show state but not figures: Plan should say the renewal or founder-price end date (`me.premium.json`), Devices the last device name (`devices.list.json`), Notifications the per-group on/off state.

**NICE**
- Tools list is nine bare words. Add one-line recent info: Guides "Fury Warrior updated", Tiers "Fury is A", Sims "948.6 DPS, 3 h ago" (all in `/v1/home`, `spec_week` and `import_sim`).
- Bottom of the drawer has dead space at 390 (the capture shows a third of the screen empty under Help and Privacy). Use it for the guild block above, not more links.

## Cross-platform

- iOS Home and Android Home are about equal, so no data is dropped on one side only except as listed. The iOS Home raid card shows "You are In", Android Home shows "you have not answered" until chosen (the answered state is already tracked in the BRIEF).
- iOS Gear row shows slot-name labels "Finger 1" and "Trinket 2". Android shows "Finger" and "Trinket" for two slots each, which is ambiguous when both rings are on the list. Use `slot_label` from the API ("Trinket 2").
- iOS Home source line uses "Crafted: Blacksmithing" (the API `source_line`), Android shows the trimmed text. Use the API string on both.

## Needs data

- Per-group unread counts for Activity chips (client tally, or a field on `/v1/feed`).
- A cached `total_gain` per character for the drawer (only the active character has it).
- A "last 3 pulls" or boss-progress figure for the Raid tab (only past-raid summaries exist in `guild.raids`).
- Reputation standing for "needs revered" rows (the text exists; the standing number does not).

## Verdict

**THIN** (fail). The MUST additions, in value order:

1. Raid tab: buff and cover coverage, invite time and duration, bench and out counts, officer notes (`guild.raid` `cover`, `composition.buffs`, `invite_at`, `duration_minutes`, `counts.bench/out`, `notes`).
2. Home hero: `total_gain`, `upgrade_count`, `talents.state`, `synced_at`, plus the raid card's role split (`next_raid.counts`).
3. Gear header: `total_gain`, a corrected "slots on the list" counter, and `synced_at`.

## Round 2 (2026-11-10, base 9bbf15e37)

# info-density-reviewer: round 2, Android tab roots and drawer

I judged the parity goldens and `.txt` goldens, `android-r2/raid-01-top` and `activity-01-top`, the drawer golden, and the strings and composer code for Gear. I did not open the flow-step PNGs or any iOS capture. The `fs2-home` squeeze is known and not re-reported.

## Round 1 findings

| # | Round-1 finding | Status |
|---|---|---|
| H1 | Home hero figures (BLOCKER) | FIXED in the ruled form. The Next upgrades meta reads `13 upgrades · +157.6 DPS`, and the stamp reads `Synced 2 h ago from the addon`. The `Gear ›` link and the stamp are now both on screen. I accept the ruling that the head stays figure-free (app-flows line 165). |
| H2 | Home raid card role split and countdown | FIXED. The card shows `Tanks 2/2 · Healers 5/8 · Damage 11/30` with short roles in red. On raid day the countdown and the `2 things before you pull` block show. |
| H3 | Three upgrade rows, no count | FIXED. The meta carries the count and total, and `Gear ›` goes to the full list. |
| H4 | Worn item, drop chance, Best run row | FIXED. The golden shows `· replaces Helm of Valor`, and Best run is a row with icon, `· 8%` and gain. |
| H5 | spec_week, movers, guild ladder, import_sim on Home | KEPT. The ruling is that these are web Home sections and the app reaches them through New for you and Activity. I accept it: `you are #7 of 14` and `Sim done: 948.6 DPS` are both already in the feed rows. |
| H6 | Item names in quality colour, other player's crest | FIXED. Emulator Activity shows Deathbringer in epic purple, and Home and Activity rows use the segment styles. |
| R1 | Raid cover and buffs (BLOCKER) | FIXED as a text line. See N1: it ships as plain text. |
| R2 | Invite time, size, bench and out (BLOCKER) | FIXED. `Invites 7:45 pm · 40` is in the head, and `Bench 1 · Out 3 · 14 not answered` ends Who is coming. |
| R3 | Officer notes (BLOCKER) | FIXED. Two lines, truncated with an ellipsis. |
| R4 | Class-coloured names, late minutes, melee and ranged | FIXED. Names are class-coloured (`raid-01-top`), `Brightwen (Late · 15 min)` is shown, and the Damage line reads `· 5 melee, 6 ranged`. |
| R5 | Before you pull icons and counts | FIXED. Item icons are in the emulator capture. The `3 in your bags` figure was already there. |
| R6 | Footer figures and last-raid line | FIXED. Roster shows `38 members, 31 at 60`, Calendar shows `Onyxia Wed, Barrow Deeps Sat`, Loot history shows `Last raid: 4 pieces`, and Guild sims shows `You are #7 of 14 DPS`. The no-raid state keeps the footer, so it is no longer blank. |
| R7 | Attendance (NICE) | STILL OPEN as a NICE. It is not on the tab and not in the drawer. |
| G1 | Gear header context (BLOCKER) | FIXED, with a new problem. See N2. |
| G2 | Sort, gain bars, source on cells, preset segmented | KEPT. The paperdoll order and the ControlBar ruling (redesign plan M3) stand. |
| G3 | Not-modelled row | FIXED as a header suffix: `2 gains not modelled yet`. The rows themselves are not shown, which I accept for the app. |
| G4 | Where to go: upgrades per place | FIXED. Rows read `1 upgrade · +1.9 DPS`, `2 upgrades · +20.4 DPS` and `1 upgrade · +6.8 DPS`. |
| G5 | Slot labels `Finger 1` and `Trinket 2` | STILL OPEN. The golden still shows `Finger` twice and `Trinket` twice. The `slot_label` field was not adopted. See N5. |
| A1 | Activity chip counts | KEPT until a feed field exists. This is on my needs-data list. |
| A2 | Trailing figure per kind | FIXED where the segment carries one. The emulator shows `1196.2` and `1184.3` on the right. |
| A3 | Item names in colour, crest for other players | FIXED. |
| A4 | Day-of-week and guild name on raid rows | STILL OPEN. This is a NICE: the buckets are `Tue` and `Sat` per the S3 ruling, and no guild name appears on raid rows. |
| D1 | Drawer character rows (ilvl, guild rank) | FIXED. The head reads `Fury Warrior · 60 · Raider` and the switch rows read `item level 61`. |
| D2 | Guild block in the drawer | KEPT under the one-job ruling, since the Raid tab is the guild. |
| D3 | Sync status and error | FIXED for the stamp, with `Refresh` kept on the same line. I did not see the red error state in a golden. |
| D4 | Plan and Devices values | PARTIAL. Devices reads `3 signed in · newest Pixel 9`, and Notifications shows `Raids, gear and sims, logs on`. Plan reads `Premium · Founder price` with no founder end date, which the ruling promised. See N3. |
| D5 | Dead space under Help at 390 | STILL OPEN as a NICE. The drawer golden has a third of the screen empty. With the guild block ruled out, fill it with the Tools sub-lines (below) or nothing. |

## New findings

**SHOULD**
- **N1: Covered line is wall-of-text with no icons.** The emulator `raid-01-top` shows three lines of red text: `Missing Blessing of Kings · Missing Blessing of Might · Missing Blessing of Wisdom · Missing Judgement of Wisdom · Missing Judgement of Light · Missing Blessings`. The same word "Missing" is repeated six times, and the last entry restates the first three. Drop the repeated prefix and put it once, as `Missing`, with the spell names after it. Put the spell icon beside each name from `composition.buffs[].icon` (fixture `guild.raid.json`). Fold it into the file that builds the line (`RaidTabParts.kt`). Do this because a leader scans for icons. At 390 the golden's one-line case (`Missing Fortitude · Windfury 1 of 2`) fits, but the full-gap case costs three lines.
- **N2: Gear counter and visible rows disagree.** Gear's counter reads `4 of 17 slots match the list`. The golden shows 5 cells labelled `Equipped` (Neck, Shoulder, Waist, Trinket, Off hand), 12 cells with a gain, and the header says `13 upgrades`. The counter comes from `slots.count { it.worn }` in `GearComposer.kt:79`. The header number in the same block says 13, while only 12 cells show a gain; Blackcrow (a ranged pick) may account for the 13th, which would be double-counted. I raised this as a BLOCKER in round 1 and the ruling says "the counter reads `0 of 17 slots match the list`". Reconcile all three numbers (counter 4, Equipped cells 5, header 13 against 12 gain cells) against `upgrade_count` and `already_best[]`, or explain the 13th. A raider will count the cells.
- **N3: Plan row has no date.** The decision says "the plan's renewal or founder date". The golden shows `Premium · Founder price`. Add the end date from `me.premium.json` (6 Jan 2027 per the pricing memory). If the fixture lacks the date, say so in the BRIEF.
- **N4: Home raid card tally has no "N short".** On Home the tally reads `Healers 5/8` in red, while Raid says `Healers 5 of 8 · 3 short`. The ruling S6 asked for one tally model on both. The compact form is acceptable, but the card should add the short count when a target exists, because red alone does not say how many.
- **N5: Slot labels.** The two rings and two trinkets are indistinguishable in the Gear golden (`Finger`, `Finger`, `Trinket`, `Trinket`) and in the sheet. Use `slot_label` from the API (`Finger 1`, `Trinket 2`). This is G5 above; it was unchanged.

**NICE**
- **N6: Gear cells show `Head ·` with a dangling separator** in the text golden (`Head ·` followed by the gain on the next line). The screenshot shows it as `Head · +14.5 DPS`, so only the text form has the stray separator; make the separator part of the same node so TalkBack does not read it alone.
- **N7: Activity empty state** (`android-activity-empty`) is still `Nothing yet. Loot, sims, reports and raids land here.` It names what fills it, which meets the standard, but it could carry the filter icons. NICE only.
- **N8: Drawer Tools are nine bare words.** Each could carry one figure that Home already fetches: `Tiers · Fury is A`, `Sims · 948.6 DPS, 3 h ago`. This also fills the empty third. Source: `/v1/home` `spec_week` and `import_sim`. NICE.
- **N9: Activity row `1196.2` repeats the body text `1196 DPS`** on the Brakka row (`activity-01-top`). Keep the trailing figure and cut the figure from the body, or the other way round, so the row says it once.

## Phones and font scale
Nothing new collapses at 390. At font scale 2 the only issue I saw is the one already known. The Gear grid reflows to two columns at 390 with 3-line names (`Forest Stalker's Bracers`), which is acceptable. Check N1 at font scale 2 once it is rebuilt.

## Needs data
- Per-group unread counts for the Activity chips (client tally, or a field on `/v1/feed`).
- The founder end date, if it is not in the fixture `me.premium.json`.
- A cached `total_gain` per character, so the switch rows can show a gain.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.

## Verdict

**DENSE WITH ADDITIONS**

## Round 3 (2026-11-10, base 7911bf87e)

# info-density-reviewer: round 3, Android tab roots and drawer

I judged these on the current build:
- the `android-r3/` step captures (raid-01, gear-01, home-01, activity-01, drawer-02);
- the parity goldens and their `.txt` files (home-60, raid-officer, gear, activity, menu, menu-low, activity-empty);
- `GearComposer.kt`.

I did not open the flow-step PNGs or the iOS captures. I did not re-check the fixture limits the BRIEF lists: the 12 d stamps, the tally without a target, and Late snapping back. Round-2 N3 (Plan founder end date) and A1 (Activity chip counts) are ruled data gaps. They stay under needs data and do not count against the verdict.

## Round-2 findings

| # | Round-2 finding | Status |
|---|---|---|
| N1 | Covered line: a wall of "Missing", no icons | FIXED in the golden. It reads `Missing Fort · WF 1 of 2` with the label once. The emulator capture is not clean; see F1. |
| N2 | Gear counter, Equipped cells, header count and gain cells disagree | FIXED in the golden. Counter 5 of 17, with 5 Equipped cells (Neck, Shoulder, Waist, Trinket 1, Off hand) and 12 gain cells. The header says 12 upgrades, and the 12 gains sum to the header's +47.5 DPS. The "Where to go" totals also reconcile (20.4 = 14.5 + 5.9). The emulator has one leftover; see F2. |
| N3 | Plan row has no founder end date | KEPT. The ruling accepts that no date exists in the API and the founder price holds while subscribed. The golden's `Premium · Founder price` is the honest figure. Needs data only if a `founder_until` field is ever added. |
| N4 | Home tally has no "N short" | FIXED. The golden reads `Tanks 2/2 · Healers 5/8 · 3 short · Damage 11/30 · 19 short`. |
| N5 | Slot labels `Finger 1` and `Trinket 2` | FIXED. The cells and the merged semantics nodes both carry them (`Finger 1, +6.8 DPS`). |
| N6 | `Head ·` dangling separator | FIXED for TalkBack. Each cell has a merged node (`Lionheart Helm, Head, +14.5 DPS`). The separator still appears as its own text line in the `.txt` golden, but it no longer reads alone. |
| N7 | Activity empty state could carry the filter icons | STILL OPEN as a NICE. It still reads `Nothing yet. Loot, sims, reports and raids land here.` It names what fills it, which meets the bar. |
| N8 | Drawer Tools are nine bare words | STILL OPEN as a NICE. The third-of-a-screen dead space from round 1 is gone, because the Account block now carries Notifications, Plan, Devices, Offline data and Settings with values. |
| N9 | Activity figure shown twice | FIXED. `activity-01-top` shows `1,196 DPS` once on the right. The body reads `+12 over your 1184 (Fury Warrior, 60)`, which is a different figure. |
| R7 | Attendance | STILL OPEN as a NICE. `me.characters[].guild.attendance` is not drawn anywhere. |
| A4 | Day-of-week and guild name on raid rows | KEPT. The S3 bucket ruling stands (`Sat`, `Tue`). |
| D3 | Sync error in the drawer | PARTIAL. The warning-tone stale stamp with `Refresh` shows in `drawer-02-open` (`Updated 12 d ago from Battle.net`). I saw no `sync_error` red state in any golden. |

Round-1 items I accepted in round 2 (Home hero, raid card split, upgrade meta, worn item, drop chance, Best run, officer cover and notes, Where to go figures) are unchanged and still hold in the goldens.

## New findings, ranked

**SHOULD**
- **F1: Covered line draws icons unevenly.** In `android-r3/raid-01-top`, `Blessing of Kings` has a spell icon. `Blessing of Might` and `Blessing of Wisdom` have none, only a gap where the icon would go, and the second row is indented to match. The fixture buffs also lack a short `label`, so full names show instead of `Kings`, `Might`, `Wisdom`. Either draw an icon for every name or none. Where `composition.buffs[].icon` is missing, collapse the gap. Fix this in the Covered renderer in `raids/flows/RaidTabParts.kt`. A raid leader scans this line by icon, and one icon in three reads as a bug.
- **F2: One Gear cell says nothing.** In `android-r3/gear-01-top`, Shoulder (`Officer's Pauldrons`) shows `Shoulder` with no gain and no `Equipped`. The counts add up only because it is the 17th cell: 3 Equipped, 13 gains, 1 silent. The header used to say `2 gains not modelled yet` and no longer shows it in either capture. `GearComposer.slotRow` sets `gain` only when `modelled` and `gain > 0`, so a not-modelled or zero-gain slot falls through blank. Give that case a one-word value (`Not modelled`, or `+0.0 DPS` when the gain is genuinely zero). The source is `home.notModelledCount` and `HomeUpgrade.modelled`. A raider who counts cells will find the hole.
- **F3: Home's tally never splits melee and ranged.** In `android-r3/home-01-top`, Home reads `Damage 15`, and the Raid tab reads `Damage 15 · 7 melee, 8 ranged`. Round-3 change 1 made `Damage N` equal melee plus ranged on both surfaces, and it does. But Home is a shorter form of the same fact, and the golden reads the same either way, so this stays minor. Add the split only if the card has room at 390.

**NICE**
- **F4: Drawer Notifications row has no value in the emulator** (`drawer-02-open`). Plan says `Free` and Devices has a value, but Notifications shows a bare label. The golden shows `Raids, gear and sims, logs on`. Confirm that the empty emulator row is the fixture's missing prefs and not a collapsed state.
- **F5: Activity rows lack a per-kind figure beyond sim DPS.** On Loot awarded (`Deathbringer`), the `+12.6 DPS` that Home's feed row carries is missing. The `figure` field is on the segment where the server sets it. Check whether the server sets it for `loot.awarded`.
- **F6: Gear still has no gear-talent status row.** This is KEPT under the one-job ruling, so it is not a finding. It is listed only to say I did not reopen it.

## Phones and font scale
Nothing new collapses at 390. The Gear grid stays two columns with 3-line names (`Precisely Calibrated Boomstick` takes four lines in the emulator and is still readable). F1's Covered line gets more cramped at `fs2-raid` if icons are added to every name, so check it at font scale 2 after the F1 fix.

## Needs data
- Per-group unread counts for the Activity chips (a field on `/v1/feed`, or a client tally).
- A founder end date, if the Plan row is ever to show one. Not in the API today.
- A cached `total_gain` per character, so the Switch rows can show a gain.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.

## Verdict

**DENSE WITH ADDITIONS**

Worth making soon, ranked:
1. F1: uniform icons on the Covered line.
2. F2: a value on the silent Gear cell, and restore the not-modelled count.
3. F5 and F4: the Loot row figure and the Notifications value in the emulator build.

## Round 4 (2026-11-11, base 693e7a952)

# info-density-reviewer: round 4, Android tab roots and drawer

I judged these on the current build:
- `android-r4/`: raid-01-top, raid-officer-02, home-01-top, light-gear, drawer-02-open.
- The parity goldens and `.txt` files: gear, activity, menu.
- `GearComposer.kt`, `GearCountsTest.kt`, `GearHeaderTest.kt`, `MenuScreens.kt` and `ProfileMenu.kt`.

I did not re-check the BRIEF's known items: the Might and Wisdom art, the 12 d stamps, the tally without a target, Late snapping back, and the three notification actions.

## Round-3 findings

| # | Round-3 finding | Status |
|---|---|---|
| F1 | Covered line draws icons unevenly | FIXED, with a fixture limit. In `raid-01-top` and `raid-officer-02` the line reads `Missing [icon] Kings · Might · Wisdom and 2 more`. The names have short labels, there is no gap where Might and Wisdom would have an icon, and the text closes up. One icon in three still looks uneven. That is the known fixture limit, not app behaviour. I raise N2 about the `and 2 more` tail. |
| F2 | Silent Gear cell, and the not-modelled count | FIXED. In `light-gear` the Shoulder cell reads `Shoulder · +0.0 DPS`, so no cell is blank. The counts reconcile: 3 Equipped, 13 gains, 1 at +0.0, which makes 17. The 13 gains sum to the header's +18.0. `GearCountsTest` pins it as "no cell is silent". The header keeps `· N gains not modelled yet`, and `GearHeaderTest` covers it. No golden or capture shows the `Not modelled` cell or the header suffix, because the fixture has `notModelledCount` 0, so I judged that path from the tests only. |
| F3 | Home's tally has no melee/ranged split | KEPT by ruling. Home reads `Tanks 3 · Healers 6 · Damage 15`, and the Raid tab carries `7 melee, 8 ranged`. I did not reopen it. |
| F4 | Drawer Notifications row has no value in the emulator | STILL OPEN, and it is a real bug. `drawer-02-open` shows a bare `Notifications`. The goldens (`android-menu.txt` line 28: `Raids, gear and sims, logs on`) only show it because `MenuParityTest.kt:122` hard-codes the string. `MenuScreens.kt:117` never sets `MenuLines.notifications`, so production always passes null. The string `menu_notifications_on` (`%1$s on`) is not used anywhere in `app/src/main`, so the real row is dead code. The golden shows a value the app never draws. See N1. |
| F5 | Loot row lacks a per-kind figure | FIXED. The golden reads `Onyxia · main hand, +12.6 DPS` on the Deathbringer row. The figure sits in the body, not in the trailing column the sim row uses. That is acceptable. |

## New findings, ranked

**SHOULD**
- **N1: Wire the Notifications row.** In `MenuScreens.kt:117` build `notifications` from the per-group on/off state, which comes from the `app.notifications.json` contract that the Notifications screen already loads. Render it with `menu_notifications_on` or `menu_notifications_off`. A raider checks `Raids, gear and sims on` before a raid night. Then make the golden's value come from the same builder, so the golden and the emulator cannot disagree again.
- **N2: Cap `and 2 more` on the Covered line.** The line shows three gaps and hides the rest. For an officer, the hidden gaps may be the ones that matter (Judgement, Windfury). Show every gap name, which costs one or two wrapped lines. If it must stay capped, make the `and N more` tail a tap target into the raid page's cover list. The source is `guild.raid.cover[]` (`state`), which the Raid tab already reads.

**NICE**
- **N3: Sim rows have no trailing figure.** `Kaelra passed you` carries `1,196 DPS` on the right. `Sim done: 948.6 DPS ±6` has its figure only in the title, and `+21 vs your last` is in the body. Use the same trailing column for sim rows (`feed.rows[].segments[].figure`).
- **N4: Raid footer.** The footer is below the fold, so I could not judge whether the Raid tab still has the per-row figures from round 2.

## Evidence defects (not the app's)
`android-r4/gear-01-top.png` and `activity-01-top.png` are captures of the system notification shade, not the app. The Gear and Activity emulator captures that carry weight are `light-gear` (the same state) and the parity goldens. Re-shoot those two before the next round.

## Phones and font scale
Nothing new collapses at 390. At fs2, N2 would cost one or two lines on Raid. Check `fs2-raid` when N2 lands.

## Needs data
- Per-group unread counts for the Activity chips, as a field on `/v1/feed` or a client tally.
- A founder end date, if the Plan row is ever to show one.
- A cached `total_gain` per character, so the Switch rows can show a gain.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.
- Spell art for Might and Wisdom on the fixture build, which is a fixture limit.

## Verdict

**DENSE WITH ADDITIONS**

Worth making soon, ranked:
1. N1: wire the Notifications row in production and drive the golden from the same builder.
2. N2: show every Covered gap, or make `and N more` a link.
3. N3: put the sim figure in the trailing column. Also re-shoot the two notification-shade captures.

## Round 5

Round 5 is DENSE WITH ADDITIONS: all three round-4 additions landed. I judged the Android r5 captures (home-01, raid-01 and raid-officer-02/03, gear-01, activity-01, drawer-02) and ios/light-raid, plus the fixtures. I did not open the iOS flow snapshots because they are still in CI.

## Round-4 additions

| Addition | Status |
|---|---|
| N1: drawer Notifications value | LANDED. `drawer-02-open` reads `Raids, gear and sims, logs, guild ladder, tiers, officer on`. It is no longer bare. |
| N2: `and N more` on Covered | LANDED. `raid-officer-02` shows `and 2 more` in the link colour, and the not-synced line has a chevron and a two-line row. Whether the tail opens the cover list I judged from the decisions file, not from a capture; `raid-officer-05-cover-list` exists, so I take it as shown. |
| N3: sim figure in the trailing column | LANDED. `activity-01-top` shows `190.4 DPS` right-aligned on the sim row, from `activity.list.json` `figure`. |
| Re-shoot of the notification-shade captures | LANDED. `gear-01-top` and `activity-01-top` are the app now. |

## Findings

**SHOULD**
1. **The Loot row has no figure, and the trailing column is uneven.** `activity.list.json` row 4163 (`loot.awarded`, Deathbringer) has no `figure`, so `activity-01-top` shows the trailing figure on the ladder and sim rows but not on loot. In round 4 I saw `+12.6 DPS` in the body on the parity golden. Either the golden carries data the fixture lacks, or the server sets it only sometimes. Make the fixture and the golden agree, and show the DPS gain in the trailing column when the sim has it. A raider wants to see what the drop is worth.
2. **The drawer Notifications value is a list of everything on.** `drawer-02-open` wraps `Raids, gear and sims, logs, guild ladder, tiers, officer on` onto two lines. When every group is on, collapse it to `All on`. Otherwise list only the groups that are off, so a raider sees at a glance what they will miss. The source is the per-group state in `app.notifications.json`. At fs2 (`fs2-drawer`) this line takes more space, so check it there.
3. **Gear's header is dense but flat.** `gear-01-top` reads `50 to 59 · Horde · 3 of 17 slots match the list · Change`, then the +18.0 DPS line. The cells show a per-slot gain, but none shows where the piece comes from, which Home's rows do (`Crafted: Blacksmithing`, `Lady Palanseer`). Add the source to the cell sheet if it is not there already. The sheet is `gear-03`, and I did not re-open it this round. It stays on the Gear tab, so the one-job rule is not touched.

**NIT**
4. The Activity captures `activity-01-top` and `activity-06` still read `203.6 DPS` and `+12.2 over your 191.4`. The decisions file says to judge by the house format (`204 DPS`), so this is capture age only. Re-shoot before the next round.
5. `raid-officer-03-roster-not-synced` is the roster screen, which is another lane's. The stat strip reads `4 of 38 · 4 in game, no account · 2 waiting`. It is not clear whether the last two are filtered too, so another lane should check that.
6. iOS parity holds. `ios/light-raid` matches Android row for row (Covered line, `40-man`, tallies with the melee/ranged split, tab badges 1 and 1). I have no iOS drift to add.

**BLOCKER**: none.

Rulings I did not reopen: Home's tally without the split (F3), the age on its own meta line, and the Might and Wisdom art.

## Needs data (unchanged from round 4)
- Per-group unread counts for the Activity chips.
- A founder end date for the Plan row.
- A cached `total_gain` per character for the Switch rows.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.

Files:
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/docs/contracts/fixtures/app/activity.list.json
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r5/

VERDICT: DENSE WITH ADDITIONS

## Round 6

Round 6 landed S2 on both platforms. S1 stays a KEEP I accept, and S3 got no ruling and is still open. There are no blockers.

I read the round-6 decisions, my round-5 verdicts, `android-r6/` (drawer, Activity, stale Gear, filtered roster) and `ios/r6-menu-01-open.png`. I did not open the iOS `r6-activity`, `r6-gear` or push-failure frames, or any fs2 captures.

### Round-5 additions

| Addition | Status |
|---|---|
| S2: drawer Notifications value | LANDED. Android `drawer-01-open` reads `All on`, and iOS `r6-menu-01-open` reads `Gear and sims, guild ladder, tiers off`. The value is one short line on both. |
| S1: loot figure | KEEP stands. `Loot awarded` has no trailing figure in `android-r6/activity-01-top.png` because `loot.awarded` carries no per-character gain on the server. The ruling answers my objection, so I list this as needs-data. |
| S3: Gear source | Still open. The round-6 decisions do not list it. |
| N4: house format | LANDED. Activity reads `204 DPS`, `+12 over your 191 (Fury Warrior, 60).` and `190 DPS`. Android `drawer-01-open` and iOS `r6-menu-01-open` both read `Updated 9 h ago from Battle.net`. |
| Roster stat plate | LANDED. `android-r6/raid-officer-02-roster-filtered.png` shows `4 of 38 on the roster` alone, with the filter chip `Not synced this week · 4`. |

### SHOULD
1. **Gear source is still missing.** `android-r5/gear-03-sheet.png` shows the item tooltip and the guildmates who need it. The runner-up row carries its source (`PvP rank 11 · Blood Guard · Horde`), but the equipped piece and the upgrade cells carry none. The data exists: `apps/android/.../gear/WhereToGo.kt` builds sources from `/data/{build}/loot.json`, with kind, name, boss, drop chance and profession. Add a `Where` line to the sheet for the selected item, and show a source column or sub-line on cells with a gain. Home's Best run line already shows this data, so no invented data is needed. Source text collapses to one wrapped line at 390.
2. **The two platforms show different Notifications states.** Android reads `All on` and iOS reads `Gear and sims, guild ladder, tiers off`. If both read the same `app.notifications.json`, one platform has a bug. If the fixtures differ, say so in the brief. I could not tell which from the captures.
3. **The roster row carries more figure than the Activity rows.** Roster rows show an 8-night pip bar and `2 of 8 nights`. Activity rows show no per-row context beyond the day. That is acceptable, but the `Show older` row sits above about 400 px of empty ground in `android-r6/activity-01-top.png`. With one page of three rows, the empty state is honest and fine.

### NIT
- The stale Gear frame (`android-r6/gear-01-stale.png`) reads `synced 9 d ago` in orange, which is good. The Main hand cell wraps at the phone width (`Main hand · Equipped`). This is the KEEP.
- `android-r6/raid-officer-02-roster-filtered.png` shows `Attendance is the last 8 raid nights with a log.` as grey explainer copy. It is on the roster screen, which is another lane's, so I only note it.

### Needs data
- A per-character gain on `loot.awarded`, which would put a trailing figure on loot rows. This is already in the round-6 KEEP.
- Per-group unread counts for the Activity chips.
- A cached `total_gain` per character for the Switch rows.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.
- A founder end date for the Plan row.

Files:
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r6/
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/ios/r6-menu-01-open.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/gear/WhereToGo.kt
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/docs/contracts/fixtures/app/app.notifications.json

VERDICT: DENSE WITH ADDITIONS

## Round 7

Round 7 landed both of my round-6 additions. I have no blocker. I read `android-r7/gear-01-top`, `gear-02-sheet-where` and `drawer-01-notifications-some-off`, and `ios/r7-gear-02-sheet-where-settled` and `r7-menu-notifications-some-off`. I did not open the All on and All off drawer frames, the push-failure frames or `r7-home-01-top`. I also did not open `r7-gear-03-sheet-where-scrolled`, so I did not check the iOS Where row against Android's on the scrolled sheet. The Where-line items below rest on the Android Helm sheet and the iOS Cape sheet.

**Round-6 additions**
- **S1, Gear source:** landed as a `Where` row in the sheet, not as a cell sub-line. Android reads `Crafted: Blacksmithing` and iOS reads `World drop` for the Dark Phantom Cape. The data is the one `WhereToGo.kt` already loads. The round-7 KEEP on cells without a source sub-line stands. The cells are already full at 390 with name, slot and gain, and one tap reaches the sheet.
- **S2, Notifications value:** landed. Android `drawer-01-notifications-some-off` and iOS `r7-menu-notifications-some-off` both read `Gear and sims, guild ladder, tiers off`. The fixture `app.notifications.json` is the same file on both.
- **Gear header:** reads `182 DPS` and `+18.0 DPS from 13 upgrades · synced 10 h ago` on both apps, in the house format.

**BLOCKER**
- None.

**SHOULD**
1. **The Where row carries no figure on the Android Helm sheet.** It reads `Crafted: Blacksmithing` with nothing after it.
   - Raiders ask what the profession skill is. For a boss drop they ask for the chance, which is `drop chance` in the `WhereToGo` source model.
   - Add the figure where the data has one, as a trailing value on the row, for example `Crafted: Blacksmithing · 225`. For a boss drop, `Onyxia · 18%`.
   - If `loot.json` has no figure for that item, as with the `World drop` cape, leave it off. That is the existing round-7 KEEP on loot rows without a figure.
   - I could not tell from the captures whether a boss-drop item exists in the fixture. Please confirm one is shown, so the instance, boss and chance path is actually exercised.
2. **The Android sheet is missing the gain line that iOS has.** The iOS Cape sheet ends the tooltip with `+0.8 DPS` in gold. The Android sheet shows the Helm, which is equipped, so it may be a case difference. Capture an Android sheet for an upgrade piece to confirm the two apps match. This is a capture gap, not a confirmed drift.
3. **The Android runner-up row has a source but no gain.** The Blood Guard's Plate Helm row reads `Runner-up · PvP rank 11 · Blood Guard · Horde`. iOS shows `Runner-up · World drop · −0.3 DPS` on its runner-up row. Add the gain to the Android row if the runner-up has one. I could not tell whether the Helm's runner-up has a gain in the data.

**NIT**
- Android's `Where to get it` sub-line under the source reads as an explainer label. If it is the tap target to the Where-to-go screen, as `FlowRows.kt:338` suggests, add a `›` so it reads as a link. The iOS row has no equivalent sub-line, so the two apps differ here. iOS has no drill-down from the row, or I missed it. Please align the two.
- The Android Helm cell and its sheet repeat `Equipped`, which is fine. The Main hand wrap at 390 is the old KEEP.
- The `3 guildmates have not shared their gear yet` footer is honest and useful.

**Needs data** (carried from round 6)
- A per-character gain on `loot.awarded`.
- Per-group unread counts for the Activity chips.
- A cached `total_gain` per character for the Switch rows.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.
- A founder end date for the Plan row.

Files:
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r7/gear-02-sheet-where.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/ios/r7-gear-02-sheet-where-settled.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/gear/flows/GearItemSheet.kt
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/gear/WhereToGo.kt

VERDICT: DENSE WITH ADDITIONS

## Round 8

Round 7's three additions (S1, S2, S3) have landed on both apps, and I have no blocker. I found no new gap that needs the owner's data. I did not open the Decline and Mention failure frames, the iOS All on and All off drawer frames, or the Home, Raid and Activity tabs. Those are unchanged by round 8's four rows, but I haven't re-checked them this round.

**Round-7 additions, checked**
- **S1, the Where row figure:** landed on both apps. `Crafted: Blacksmithing (275)` shows on the Dawn's Edge sheet (`android-r8/gear-03-sheet-crafted-upgrade.png`, `ios/r8-gear-sheet-crafted-upgrade-settled.png`) and `(245)` on the Helm. The boss-drop row reads `Onyxia, Onyxia's Lair, 8%` in `apps/android/app/src/test/screenshots/parity/android-item-sheet.png`. The world-drop cape keeps no figure because the data has none, which is the round-7 KEEP.
- **S2, the Android gain line:** landed. The gold `+2.8 DPS` and `+2.2 DPS` sit under the tooltip on the crafted-upgrade and vendor-upgrade sheets. The golden shows `+12.6 DPS`. The tooltip, the source row and the gain line now match iOS in structure.
- **S3, the runner-up gain:** landed. Android shows `Runner-up · Onyxia · 8% · −1.8 DPS`, `Crafted: Blacksmithing · −5.0 DPS` and `World drop · −1.3 DPS`, so the runner-up carries its source and its gain. The Blood Guard helm reads `Same DPS as the pick` because the data has no gain figure for it.
- **Gear `No gain`:** shown on the Officer's Pauldrons cell on both apps (`android-r8/gear-01-top.png`, iOS Gear behind the sheet).
- **Drawer Notifications value:** reads `Gear and sims, guild ladder, tiers off` (`ios/r8-menu-notifications-some-off.png`). Plan and Devices carry a value, `Free` and `3 signed in · newest iPhone`.

**BLOCKER**
- None.

**SHOULD**
1. **A runner-up's source repeats the row's own heading, and the two apps word it differently.**
   - On a crafted pick the heading already reads `Crafted: Blacksmithing (275)`, and the runner-up row repeats it.
   - Android reads `Runner-up · Crafted: Blacksmithing · −5.0 DPS`. iOS reads `Runner-up · Blacksmithing · −5.0 DPS`.
   - The Cape's runner-up shows `World drop` under a `World drop` heading on both apps.
   - Where the runner-up's source equals the pick's, show only `Runner-up · −0.3 DPS`. Where it differs, give both apps the same shared source words. That frees width on the Helm row, which currently wraps to a second line.
2. **No frame exercises a boss-drop runner-up on a device.** The brief says the fixture's first head runner-up is a PvP helm, so the boss-drop row (`Runner-up · Onyxia · 8% · −1.8 DPS`) is proven only in the Android golden. iOS has no boss-drop capture. Once the iOS golden run lands, confirm iOS draws the same row from `docs/contracts/fixtures/app/cdn.loot.json`. This is a capture gap, not a drift.

**NIT**
- `Same DPS as the pick` is a sentence in a row of figures. `±0.0 DPS` would not match the `No gain` rule either, so the app could drop the tail and show `Runner-up · PvP rank 11 · Blood Guard · Horde` alone. This is optional.
- The iOS Notifications value is clean at the top of the drawer. The value on Plan and Devices is tight but useful.
- The `3 guildmates have not shared their gear yet` footer is honest and still useful.
- Both apps' Gear cells stay without a source sub-line, which is the round-7 KEEP. A source row one tap away now carries the skill or chance, so I raise no new contest.

**Needs data** (carried over, none new)
- A per-character gain on `loot.awarded`.
- Per-group unread counts for the Activity chips.
- A cached `total_gain` per character for the Switch rows.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.
- A founder end date for the Plan row.

Files:
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r8/gear-02-sheet-crafted-skill.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r8/gear-03-sheet-crafted-upgrade.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r8/gear-04-sheet-upgrade.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/ios/r8-gear-sheet-crafted-upgrade-settled.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/ios/r8-gear-sheet-crafted-boss-runner-up-settled.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/test/screenshots/parity/android-item-sheet.png

VERDICT: DENSE WITH ADDITIONS

## Round 9

Round 9 rows 1 to 3 landed on both apps, and the round-8 S1 and S2 are closed or reduced to a capture gap the brief declares. I opened three Android sheets and two iOS sheets. I did not re-open Home, Raid, Activity, the drawer or the notification frames, and rounds 6 to 8 left them unchanged. The iOS crafted-upgrade frame (`ios/r9-gear-sheet-crafted-upgrade-settled.png`) and the `cdn.loot.json` fixture were not opened either.

**S1, runner-up repeating the pick's source:** closed on the frames I opened.
- The Android crafted helm sheet (`android-r9/gear-02-sheet-crafted-skill.png`) and the iOS one (`ios/r9-gear-sheet-crafted-pvp-runner-up-settled.png`) both read `Runner-up · PvP rank 11 · Blood Guard · Same DPS`. The PvP row has dropped the faction word.
- The iOS cape sheet (`ios/r9-gear-sheet-upgrade-settled.png`) reads `Runner-up · −0.3 DPS` under the `World drop` heading, so the repeated source is gone.
- On Android, `android-r9/gear-04-sheet-upgrade.png` is the vendor Leggings sheet, a different fixture item from the iOS cape. It reads `Runner-up · World drop · −1.3 DPS` under `Vendor: Lady Palanseer`, which differs from the pick's source, so naming it is correct.
- The two apps now share the same source words and the same same-source rule. Android's rule is `samePlace` in `GearItemSheet.kt`, and iOS's is `isPickPlace` in `ItemSheet.swift`.
- `Same DPS` is now a figure-like tail.

**S2, no boss-drop runner-up on a device:** this is the capture gap the brief declares. The Android golden `apps/android/app/src/test/screenshots/parity/android-item-sheet.txt` shows `Runner-up · −1.8 DPS` under `Onyxia, Onyxia's Lair, 8%`. The iOS boss-drop frame is the round-8 CI one, and the iOS golden run is still in flight per the brief.

**BLOCKER**
- None.

**SHOULD**
- None.

**NIT**
- `Replaces … / Equipped` keeps its icon, a KEEP.
- The Android Gear tab frame under the sheet is clipped by the sheet, so it only shows the header and the first row. That is fine for a sheet capture.
- iOS shows the gain line `+0.8 DPS` in the tooltip, matching Android.

**Needs data** (all KEEPs in `round-9-decisions.md`, none new)
- A per-character gain on `loot.awarded`.
- Per-group unread counts for the Activity chips.
- A cached `total_gain` per character for the Switch rows.
- A last-pulls or progress figure for the Raid tab.
- The reputation standing number for `needs revered` rows.
- A founder end date for the Plan row.

Files:
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r9/gear-02-sheet-crafted-skill.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/android-r9/gear-04-sheet-upgrade.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/ios/r9-gear-sheet-crafted-pvp-runner-up-settled.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/design/reviews/2026-11-10-mobile-backfill/ios/r9-gear-sheet-upgrade-settled.png
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/test/screenshots/parity/android-item-sheet.txt
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/android/app/src/main/kotlin/gg/foreversixty/app/gear/flows/GearItemSheet.kt
- /Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/ios/App/Items/ItemSheet.swift

VERDICT: DENSE

