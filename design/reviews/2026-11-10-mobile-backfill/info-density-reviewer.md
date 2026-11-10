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

