# F16 guild extras, build review (veteran Classic raider)

Verdict: HOLD. Seen: iOS roster-all, history-page1, calendar-on, f16.6-event, raids-tab; Android roster-all. The other captures were not opened, so nothing here clears them.

## MUST
1. ios/f16.1-roster-all.png: every row has an empty dark square where the class/spec icon goes. Android (android/f16.1-roster-all.png) shows real spec icons. A blank box in a roster is the first thing a guildmate screenshots.
2. ios/f16.4-calendar-on.png: the "Raids in my calendar" switch is a solid yellow pill with no knob. I cannot tell on from off. Check android/f16.4-calendar-on.png the same way.
3. ios/f16.7-raids-tab.png: the guild head says "24 on the roster", the Roster row under it says "38 on the roster". The two numbers disagree on one screen.
4. Both roster captures: "8 of 8" has no unit. 8 what? The fixture says nights. Write "8 of 8 raids" and say the window ("last 8 nights").
5. ios/f16.3-history-page1.png (fixture guild.raids.history): the Onyxia nights are not believable.
   - Onyxia is killed on Wed 9 Dec and Thu 10 Dec, then Sun 13, Thu 17, Sun 20, Thu 24, Sun 27 and Thu 31 Dec. Onyxia has a 5-day lockout, so the same 40 cannot clear it on back-to-back nights or every 3 days.
   - Seven nights show 40, 39 or 38 raiders from a roster of 38 (34 raiders, 4 with no account). Five names in each night's present list are not on the roster at all.
   - Make it one Onyxia clear per 5 days at most, with raiders no larger than the roster.
6. ios/f16.7-raids-tab.png: the raid night reads "Onyxia, Tuesday" beside "Mon 14 Dec". The history list has no Monday Onyxia; its Onyxia nights are Sun 13 and Thu 17. Titles and dates must agree.
7. Calendar data disagree: calendar.ics has Onyxia on 11 and 18 Dec with "Pull 7:45 pm CST", while the app shows Sat 19 Dec, invites 6:45 pm, "Pull 7:00 pm" (ios/f16.6-event.png). Pick one set of raids and times.

## SHOULD
1. ios/f16.6-event.png: the title "Onyxia's Lair · In" means nothing to a raider. Use "Onyxia's Lair (signed up)". The location repeats the zone ("OLYMPUS XXVII · Onyxia's Lair"). The notes are partly covered by Delete Event, so they read "7:00 pm. 24 of 40 signed up", and the signed-up count goes stale in a static event.
2. Roster rows wrap badly: "Protection Warrior · 60 · Guild / Master" (ios/f16.1-roster-all.png) and "Marksmanship Hunter · 60 / · Member" start a line with a stray dot. Put the rank on its own line or drop "· 60" for level-60 characters.
3. iOS and Android disagree on the Rank sort. iOS goes Kraggor, Draknar, Mendara, Thokgar. Android goes Kraggor, Mendara, Thokgar, Draknar.
4. Raid history lists five-player Blackrock Depths and Lower Blackrock Spire under "Raid history" with "5 players". Either label the page "Logged runs" or filter to 10 and 40 player raids. The per-fight data is right for what it shows: Onyxia 40, Upper Blackrock Spire 10, 6 bosses in each five-player run.
5. The zones list in the fixture offers "Hyjal Summit" and "Barrow Deeps". Neither is a Classic Era raid, and Hyjal is a TBC zone. Both are unreachable by 9 Dec, so remove them or label them clearly.
6. The roster fixture's rows say "ruleset": "hardcore" while the guild and character keys say pvp. That is a data mismatch for the accuracy reviewer.

## Passed
- Spec names are full ("Holy Priest", "Feral Bear Druid").
- All class and spec combos are legal for Horde.
- The Enhancement Shaman carries no off-hand weapon issue (not shown).
- Plain members read "Member".
- Android has four tabs.
- No raid is dated before 9 Dec 2026.
- Boss lists and counts are right for the Blackrock Spire (upper and lower) and Blackrock Depths nights.

Written to /Users/jh/code/forever/.claude/worktrees/mobile-guild-extras/design/reviews/2026-10-31-mobile-wave3-guild/player-verdict-build.md

---

# Round 2 (all 48 iOS and 52 Android captures opened; fixtures re-checked)

Verdict: HOLD. Round-1 MUSTs 1 to 4, 6 and 7 are fixed. Round-1 MUST 5 is fixed on dates, lockout and roster size (Onyxia Wed 9 and Mon 14 Dec, 37 and 38 raiders). Spec icons draw, the switch has a knob, 38 matches everywhere, "8 of 8 nights" has a unit, pvp ruleset, rank ties in name order. Both platforms now match.

## MUST
1. Attendance contradicts the history it is built from. The footer says "Attendance is the last 8 raid nights with a log" (ios/f16.1-roster-all.png, android/f16.1-roster-foot.png). In the fixtures, Kraggor and Mendara read "8 of 8" but appear in 1 of the last 8 logged nights. Draknar reads "7 of 8" and appears in 1. Voidsong and Moonveil (Waiting) read "0 of 8 nights" yet both are in the Mon 14 Dec Onyxia log. The history is mostly 5- and 10-player dungeon groups, so nobody could be at 8 of 8. Make attendance count the same nights the history shows, or count raids only.
2. The player's own character is not in the guild. The Raids tab (ios/f16.7-raids-tab.png, android/f16.7-raids-tab.png) shows "Obnoxious Yell, Fury · DPS", level 58, signed up In. Obnoxious Yell is not among the 38 on the roster. The roster's own viewer is Lightningbolltz. "YOU WERE THERE" on the history rows (android/f16.3-history-page1.png) cannot be him. A raider looks for himself first and finds nobody. Put him on the roster, or use Lightningbolltz everywhere.
3. The calendar does not match the raids list. calendar.ics has Onyxia on Sat 19 Dec and a second Onyxia on Thu 24 Dec, 0 of 40 signed up. The Raids tab "ALSO PLANNED" shows Barrow Deeps on Tue 22 Dec and Onyxia on Sat 26 Dec at 8/40, Locked. The event in ios/f16.6-event.png is right, but the feed behind "Add to Calendar" disagrees with the Raids tab. The feed has no Barrow Deeps at all, and the Sat 26 Dec Onyxia is missing from it.

## SHOULD
1. ios/f16.7-raids-tab-xl.png: the Roster row icon overlaps "38 on the roster · attendance". The guild name touches the left edge. ios/you-calendar-row-xl.png has the Privacy and Terms rows squeezed together.
2. ios/f16.7-raids-tab-app.png reads "Invites in 49 days 4 h" while the golden captures read "2 days 2 h". Take the live capture on the fixture clock.
3. "Sat 26 Dec ... 8/40 LOCKED" is a week out. Locked that early looks like a bug. Barrow Deeps reads "11 signed up" with no "of N".
4. Still ask for "Onyxia's Lair (In)" over "Onyxia's Lair · In", since "In" alone reads oddly in a calendar. This is a taste call you can overrule on the spec.
5. Raid history is 23 nights in 10 days, mostly five-player runs, and Wed 9 Dec has a Blackrock Depths run starting at 1:58 pm. No guild logs this many dungeon nights. Limit the history to Onyxia and Upper Blackrock Spire, or label the page "Logged runs".
6. Android shows no "Raid nights" section on the Raids tab at the first screen. I did not check below the fold.

Barrow Deeps and Hyjal are accepted as Forever's own raids. Boss lists, sizes and spec names are right.

---

# Round 3 (all 48 iOS and 53 Android captures opened as sheets; fixtures re-checked by script)

Verdict: HOLD on one small MUST. Everything else from round 2 is fixed.

Round-2 MUSTs, checked:
- Attendance now equals each person's presence in the newest 8 logged nights (script check, 0 mismatches).
- Obnoxious Yell is on the roster (Fury Warrior, Officer, 3 of 8) and is the "YOU WERE THERE" character.
- calendar.ics matches the Raids tab exactly: Onyxia Sat 19 Dec (In, 24 of 40), Barrow Deeps Tue 22 Dec (Not signed up, 11 signed up), Onyxia Sat 26 Dec 8 pm (In, 19 of 40). The pull times convert correctly to Chicago time.
- Onyxia dates and sizes are fine (Wed 9 Dec with 37 raiders, Mon 14 Dec with 38, five-day lockout kept).
- The 26 Dec row now reads "IN · DPS", not Locked.
- The iOS live captures show "in 2 days 2 h".

## MUST
1. Your own level disagrees between screens. The roster says "Obnoxious Yell, Fury Warrior · 60 · Officer" (ios/f16.1-roster-all.png, android/f16.1-roster-all.png). The character chip on the Raids tab says "Fury · 58" (ios/f16.7-raids-tab.png). The raid roster fixture's addon block says level 60, while me.json says 58. Pick one level.

## SHOULD
1. Attendance is honest but reads badly. Everyone shows "2 of 8 nights" or "1 of 8", and the best is "3 of 8" (ios/f16.1-roster-all.png). A raid leader would read that as a guild that never shows up. The cause is that the footer says "raid nights" but the count includes five-player runs. Count only 10 and 40-player raids, and use a history window that holds at least 8 of them.
2. The feed's start time is the pull (7:00 pm) but the app's own event starts at the invites (6:45 pm, ios/f16.6-event.png). This was disclosed. A raider with both a feed and an app calendar sees two different start times. Fix when the API gets an invite time.
3. XL on iOS (ios/f16.7-raids-tab-xl.png, ios/you-calendar-row-xl.png): rows stack to one word per line and the icons get large. It is readable, so just tighten it.
4. The event title "Onyxia's Lair · In" was kept by design. Barrow Deeps reads "11 signed up" with no size, which is accepted.
5. Raid history is still mostly dungeon runs (23 nights in 10 days). Accepted per the spec.

Passed: spec names are full, class and spec combos are legal for Horde, boss lists and sizes are right, both platforms match, the Android tab count is as known.
