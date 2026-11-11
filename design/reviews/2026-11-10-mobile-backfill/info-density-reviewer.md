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

