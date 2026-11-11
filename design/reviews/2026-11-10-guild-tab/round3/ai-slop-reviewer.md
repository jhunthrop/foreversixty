# ai-slop-reviewer: Guild tab round 3 (iOS + Android), 2026-11-11

**Verdict: FAIL.** Only one of the three round-2 fixes landed, and only partly. One new row copy adds a tell 22. The round-2 note said "no re-review needed if 1 to 3 land as stated", and they did not land as stated.

I looked at `parity/{ios,android}-full-guild-tab-officer.png` and `parity/ios-full-guild-tab-leveler.png`, at `ios/dark/{officer-02,officer-06,leveler-01,leveler-04,link-01}` and `android/dark/officer-06`, and at the copy in `en.json` and `shared.json`.

The "full-page" renders are single 1179x2556 or 1081x2401 screens, not full pages. I could not see Gear gaps, Your attendance, your loot award in gold or the Assignments rows in any capture. My verdict on the round-3 additions rests on the copy keys and the code paths only.

First impression: still reads as written by a player. The raid head carries Onyxia's Lair, "As Thandrel, Damage", class-coloured names and "Ping on Discord". The panel system is inherited and the tab adds no new visual tell.

## Round-2 fixes

1. **Partly landed.** The key is now `Assigning is on the site.` ("for now" removed). It is still an apology line, and `GuildTabParts.kt:269` still renders it as a `FlowNote` under the Assignments rows, with `Open in planner ›` already in the header. Delete the line and the `guild.tab.assignOnSite` key. This is still tell 7.
2. **Not landed, or I could not verify.**
   - `6 open` still sits beside "Needs you" in both officer renders, and it restates the list under it.
   - `From your last log-out, 9 h ago` is not visible in any Guild tab capture; `raid.tab.stamp` and `raid.ready.stamp` still carry the long wording in `en.json`.
   - Check that the "Before you pull" block is gone from this tab. If it is still anywhere, shorten it to `9 h ago`.
3. **Not landed.** Both officer renders show `5 have not answered, 2 of them healers` with `Healers short by 3` under it. The Who is coming card carries the same fact (`Healers 5 of 8 · 3 short`). Drop the `Healers short by 3` line from Needs you (tell 19/22).

## New in round 3

4. **Tell 22, "Gear gap 41 DPS".** The row sits under a group label already named "Gear gaps", so the row should read `41 DPS`.
5. **Your attendance / "7 of 8 nights"** is fine.
6. **Your loot award in gold** reuses the "yours in gold" convention and adds no copy. Not seen, so not cleared visually.
7. **"Ping on Discord"** is fine.

## Still accepted

- **Tell 8:** accepted as the inherited panel system.
- **Tell 16:** the 31 / 4 / 36 spread on the guild head is still stat-tile-shaped. Optional polish.
- **Tell 14, minor:** the guild head shows `Olympus XXVII` in mixed case on iOS but `OLYMPUS XXVII` in the leveler capture. Pick one.
- **Text chevrons:** `Roster ›` under a heading named Roster is still redundant.

## Authentic, keep

The uneven hierarchy; faction-aware crest rings; `Late · 15 min` inline on the healer; class-coloured names and `Close to 60` with a level trail; `4 of 8 done`; `Plan a raid` with `Onyxia Wed, Barrow Deeps Sat`; `Recently on` / `just now`.

## Required to pass

1. Delete the `Assigning is on the site.` line and its key.
2. Drop `6 open` from the Needs you header, and confirm no log-out stamp wording is on this tab.
3. Drop `Healers short by 3` from Needs you.
4. Change `Gear gap {n} DPS` to `{n} DPS`.

After that I would expect PASS. Please include captures that show Gear gaps, Your attendance and the Assignments rows.
