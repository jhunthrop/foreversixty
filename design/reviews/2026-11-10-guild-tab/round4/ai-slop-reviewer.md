# ai-slop-reviewer: Guild tab round 4 (iOS + Android), 2026-11-11

**Verdict: FAIL.** One Guild-tab-owned item remains, the same line raised in rounds 2 and 3.

Looked at `parity/ios-full-guild-tab-officer.png` and `parity/android-full-guild-tab-officer.png` (real full pages this time), `guild.tab.*` in `en.json` and `shared.assign.*` in `shared.json`.

First impression: made by someone who plays. The raid head, "As Thandrel, Damage", class-coloured names, "Late · 15 min", the loot list with Head of Onyxia "Not awarded yet", and "7 of 8 nights" are specific to this guild. Nothing new adds a visual tell.

## Guild tab's own items

1. **Tell 18/21, "Assigning is on the site."** A grey apology line at the bottom of the Assignments panel. "Open in planner ›" in the header is already the only action and the rows carry no chevrons, so they read as read-only. If the brief needs words, put them in the link: delete the line and the `guild.tab.assignOnSite` key, or rename the link to `Assign in planner ›`. I do not accept "the brief requires it verbatim in spirit"; the brief's intent (nobody thinks they can assign in the app) is met by the link label. This is the only item holding the FAIL.
2. **Minor, tell 22:** "item level 62" on every Roster row; `ilvl 63` or `63` would do. Optional.
3. **Minor:** "Roster" heading with a "Roster ›" link; `All 38 ›` or drop the link. Optional.
4. **Minor, case:** the guild head reads "Olympus XXVII" here; check the leveler capture matches.

## Round-3 fixes, checked

- "41 DPS": landed.
- Gear gaps, Your attendance (`Thandrel · 7 of 8 nights`), the loot section and the Assignments rows are now visible; no new tells.
- Loot in gold: point me at the capture where Thandrel wins an item; I did not see a gold line.

## Routed items (Needs you) and the routing

- "6 open" beside Needs you restates the two items under it (and 6 does not match what is shown).
- "Healers short by 3" repeats Who is coming's `Healers 5 of 8 · 3 short`.
- "From your last log-out, 9 h ago" under Before you pull should read `9 h ago`.

I accept the routing for the Guild tab's own verdict: the block is shared with Home, owned by the raid lanes, and this lane was told not to change it. But the lines are in the first screenful here. Name a tracked item against the raid lane in the response, not a note.

## Keep

The uneven hierarchy; the In / Late / Bench / Out control; "Ping on Discord" as the only filled button; class-coloured names and crests; Recently on / Missing raids / Gear gaps with real names and numbers; "1 of 1 down · 6 pulls"; Recent logs with real zones; "None planned" and "None waiting".

## Required to pass

1. Delete `Assigning is on the site.` and its key, or fold the words into the link as `Assign in planner ›`.
2. A tracked item against the raid lane for the Needs you lines and the log-out stamp, named in the response.

After 1, PASS WITH FIXES for the Guild tab, with item 2 open.
