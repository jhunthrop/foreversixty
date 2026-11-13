# L5 Guild tab: info-density review (2026-11-13)

Verdict: FIX (4 must-fix items, then ranked additions)

Page job: are we raiding tonight, am I in, what is my job, am I ready, and where do I stand in the guild.
Reviewed: l5-raid-day, l5-queued, l5-old-companion, l5-no-companion, l5-no-pack against hub-guild.png; GuildHubTab.lua, GuildRaid.lua, contract 3 (raid_signup, raid_assignment, raid_roster), contract 5, Data.lua encounterNeeds/consumables.

What the tab can reach and does not render
- raid_roster per raid (counts, roles, jobs) is read for the top raid only; the Next raids rows show only your answer.
- raid_assignment has `warning` (the short tag, "OT") and `slot`; Tonight shows only label, purpose, partner.
- Ratings pack characters carry `item_level`, `fights`, `best_parses`, `attendance`, `tier`; guild carries `standing`, `nights`, `attendance`, `roster`. The top five shows Score and Att. only.
- Data.lua `encounterNeeds[npc].buffs/debuffs` (raid buff spell ids) exist; prep reads only resist and consumables.
- Follow.pushedBuilds stores { code, name, from } with no date.

## MUST FIX

1. Built Tonight lost what the approved board promised. hub-guild.png shows `Raid leader Brannoc` and `invites 19:45`; the built render has neither, and no spec change says so. The contract has no leader or invite field (raid_signup/raid_assignment/raid_roster carry raid, title, starts_at, size only). Either add `leader` and `invite_at` to raid_signup (contract 3, API) or delete both from the board and spec 4.5. Do not leave the board promising data that is not sent. Same for the guild line: the board shows `Molten Core 10/10 · Onyxia 1/1 · Thu, Sun, Tue`; the pack has no progress field (only standing, nights, attendance, roster), so built correctly shows standing, nights, attendance. Update the board, and list "guild progress" under needs data.
2. Next raids rows are a bare answer. raid_roster already arrives for every listed raid (one message each). Add a right-aligned coming figure per row (`12/40`, from signups.in+late over size; muted when size is absent) so a raider sees which night is filling. Place: GuildHubTab.drawNextRaids, model GuildRaid.nextRaids. Width: the row has 112+98 date/raid, so on the 480 window shorten the date to `Thu 20:00` (the weekday and clock; the day number is in Tonight) to fit the figure.
3. Role shortfall is invisible. `Healers 4 of 10 · Melee 2 of 12` is all muted one-colour text. Colour a group in the warning colour when coming < wanted and in success at or above it, so the gap reads at a glance; add the player's own role in the line by class colour or underline (your group, from raid_signup.role). Source: raid_roster.roles[*].coming/wanted, raid_signup.role.
4. Your prep stops at the first need. Onyxia's encounterNeeds also names buffs/debuffs. Show a `Buffs` row of the raid buffs you are missing right now (`encounterNeeds[...].buffs` against UnitBuff) only if the row can be filled honestly; otherwise leave it and list under needs data. Also `Missing` should show the count in bags (`Fire protection 0 of 3`) where Data.encounterNeeds/consumables give the kind: today it names only the item, so a player with 1 potion and one with 0 read the same. If Comms/RaidPrep already count by kind (Comms.lua:352), use that count.

## ADDITIONS, ranked by value to a raider

5. Your row in the standings: add the gap to the next place (`193 to Torvald`) in the guild heading's note or muted under your row. Source: member records via GuildRaid.standings.
6. Top-five columns: add an iLvl column from `characters[*].item_level` (muted, 24 wide), the Officer tab's width. Header click sorts Score, Att., iLvl. On 390-equivalent narrow windows the iLvl column folds first.
7. Row tooltips on the standings: score components (gear, experience, parses, attendance, discipline), `fights`, and `as_of` from the pack record; today the hover is a whisper click only. Source: characters[*].components.
8. Tonight's job: show `warning` as a gold tag beside the label (`OT`) as the web does, and when raid_assignment carries more than one encounter list each (drawJobs takes a list; confirm the render shows one row per encounter).
9. Officers sent: show how long ago (`Tue`). Follow.addPushed stores no time, so add `at` to the pushed entry (addon-owned, nothing invented) and render `from Brannoc, Tue`. The board shows it and the build omits it.
10. Sign-ups: add a one-row share bar (In / Late / Bench / Out / not signed, class-neutral quality colours already used) under the counts, 4 px, full width. Fills the dead strip under the counts in the left column (about 60 px of empty ground below the roles line in the raid-day render).
11. Officers (gate by the officer rank, names only to officers per contract): under Sign-ups, a muted `Not signed up: Ysolde, Torvald, ...` list of the raid_roster.members with answer none, class-coloured, click whispers, capped at six with `+10` that opens Officer. This is the same data Officer's roster owns, so keep to six names and link, do not restate the table.

## States
- No companion: Tonight is one muted line and the left column below it is empty; the whole left half is dead ground. Add under the line the Your prep block (it reads bags and Data.lua, no companion needed) with `for <next raid>` only if a raid is known; with none, show the attunement rows for the zone the client says you are saved to. If neither exists, leave the column as is and note it. Today the empty half is honest but unused.
- No pack: the right column drops to Officers sent and the left to one line; top five and standing are rightly left out, but the guild name and `you are Nth` could still come from the client roster (Theme.guildRoster) and the player's own place cannot. Show the guild name and the sign-up counts from raid_roster so the column is not one card.
- Old companion: `Sign up on the site` is fine. Stale `Ratings 8 days old` in the warning colour is right.
- Queued: footer says `Sends when the companion runs`, no dot; honest.

## Phones
The addon window has one width. At narrower windows fold iLvl, then Att., then the Next raids coming figure; Prep and Sign-ups stay.

## Needs data (does not exist; do not fake)
- Raid leader name and invite time (raid_signup or raid_roster field).
- Guild progression line (`Molten Core 10/10`): not in the ratings pack guild record.
- When a build was pushed (add to the addon's own pushed entry).
- Loot won across the guild, parses against spec median per boss, score history in game (spec 9 already names these).

## Round 2 (2026-11-13)

Verdict: PASS

- 1: leader, invite time and guild progression have no field in raid_signup, raid_roster or the pack; left out and listed as spec 9 follow-ups. The build no longer promises them. Accepted (board hub-guild.png should drop them when next touched).
- 2: the coming figure moved into the row hover (`12 of 40 coming`) because the 288 px column cannot take a fourth column. Accepted; the hover is the only place a night's fill shows, so a visible figure is a follow-up if the column widens.
- 3: roles are coloured by shortfall (Tanks 2 of 2 success, Healers/Melee/Ranged short in warning). Done.
- 4: `In your bags` row (`Flask 2 · Food 20`) added; the raid-buff row stays out (no honest per-player read of encounterNeeds buffs), listed under needs data. Accepted.
- Also confirmed in the redraw: Officers sent `from Brannoc, Tue`, own row 4 px apart, and the job line naming the boss.

Follow-ups still open (ranked additions 5-11, not blocking): gap to the next place, iLvl column and sortable headers, standings tooltip components, `warning` tag on the job, sign-up share bar, officer not-signed list, no-companion/no-pack dead halves.
Needs data: raid leader and invite time, guild progression, raid-buff check for prep.
