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
