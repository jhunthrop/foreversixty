# L5 Guild tab (raider view): guild-officer review

Page job: tell a raider, before the pull, whether we raid tonight, whether they are in, what their job is and whether they are ready.

Checked: l5-raid-day, l5-queued, l5-old-companion, l5-no-companion, l5-no-pack against hub-guild.png, spec 4.5/8/9, GuildHubTab.lua, GuildRaid.lua, contract 3 (raid_signup, raid_assignment, raid_roster, accepts).

What works: the four answers sit one click from the raid, the state shows at once (teal, "not sent yet", footer "Sends when the companion runs", the dot clears), the job reads with its purpose and partner, and Next raids carries your answer per row. Sections with no data drop out instead of showing blanks, and the no-companion, old-companion and no-pack states each say why. Value-test: a raider opens this on raid day without being told. An officer does not switch from the sheet on this tab alone, which is expected; that is L7.

## Verdict: FIX

### Numbered items

1. SHOULD. Tonight lacks the board's time-to-be-ready lines. The board has `Raid leader Brannoc`, `20:00 server` and `invites 19:45`. The build shows `Thu 10 Dec, 20:00 · in 3 h 12 min` only, and the time has no zone label. A raider needs to know whether 20:00 is server time or their own, and when invites go. Add "server" to the clock now (the contract's `starts_at` is UTC; say which clock you show). Raid leader and invite time are NOT in raid_signup (needs data: leader name and `invite_at` on raid_signup); until then drop them from the board rather than leave the gap unexplained.

2. SHOULD. The job line omits the boss. `raid_assignment` carries `encounter` and `warning` (the "MT" tag), but the render shows only `label`. "Phase 2 ranged, west side" does not say which fight when the raid is Molten Core. Show `encounter` as a muted prefix. Data exists, nothing new needed.

3. SHOULD. Your prep numbers have no target. `Fire resistance 42` is a number without context; the raider cannot tell if 42 is fine or short. Show it against the encounter's ask (for example `42 of 315`, coloured on shortfall). `Missing` names the consumable but not how many; show a count if the catalogue has one. If RaidPrep has no target value for a school, hide the row rather than show the bare figure.

4. SHOULD. Sign-ups roles do not mark shortfalls. `Tanks 2 of 4 · Healers 4 of 10 · Melee 2 of 12` is all muted. Colour the groups under target in the warning colour, and add the player's own group (their `role` from raid_signup) so a raider sees "your group is short". `1 of 2 jobs filled` means nothing to a raider; say `jobs filled` only for officers or add what is open to the member (needs data: unfilled labels are officer-only `assignments`).

5. SHOULD. Standing denominator is wrong while ranks are unset. With `raider_ranks` empty, GuildRaid.isRaider counts every level 60, so `you are 9th of 15` (old-companion render, 15) versus `9th of 31` on the board is a different population and swings with alts. Mark the fallback visibly (muted `all level 60s` in the note) until officers set ranks, or the figure reads as authoritative when it is not. Say "set raider ranks on the site" for officers in L7.

6. SHOULD. Standings attendance has no window. `Att.` 91% needs "90 days" in the column title or the guild line (the guild line has "41 raid nights in 90 days"; the column does not). Also the own-row gap (5 then 9) has no separator; add a thin gap or an ellipsis.

7. SHOULD. Officers sent drops the date the board shows (`from Brannoc, Tue`). A build sent three weeks ago reads as current. Add the day from `pushed`, and a raid-build label only when the build is the one for the shown raid.

8. NICE. Stale pack. The old-companion state shows `Ratings 8 days old` in warning colour, good; also grey the scores at that age.

9. NICE. Late has no minutes. Contract has `late_minutes` 15/30/60, but the four answers (unified format) are right; show `late_minutes` in the sign-up detail when the site recorded it.

### For L6 (chip lines)
- waitingOf only fires when the raid is tonight (isTonight). A raid tomorrow evening with no answer raises no dot or chip. Add a next-24-hour window, or officers will find raiders unsigned at pull time. Add the countdown to the chip line (`Onyxia 20:00, in 3 h 12 min, sign up`).
- The job item is keyed by raid and labels, good (dismissal does not reappear on the same job). Keep the sign-up item's fact keyed on raid id, as now.
- Queued answer clears the dot immediately, which is right; it must come back if the companion rejects or the queue drops (check Export.mergeRaidAnswers expiry).

### For L7 (Officer)
- Officers' names arrive only in the officer `raid_roster` (`members`, `assignments`); the roster-by-name, unanswered list, bench and open jobs come from there. Do not use the counts message for officer views.
- Actions are blocked in the addon (outbox accepts only `raid_answer`): bench, assign, award stay on the site; the Officer tab should deep-link, not pretend. Show who has not answered and a one-click `whisper` using nameClicks, which this tab already has.
- Push build needs rank index <= 1 or raid leader per contract; hide the control otherwise.
- Raider-rank setting on the site feeds this tab's standings (item 5); link to it.

VERDICT: FIX

## Round 2

Checked l5-raid-day and l5-old-companion against round-1-fixes.md. Items 1 (clock says "local", leader and invite time left out), 2 (boss on the job line when it differs), 4 (roles coloured short/met, jobs-filled dropped), 5 (fallback says "at level 60"), 7 (Officers sent has the day) are fixed. The standings gap has a 4 px separator and the attendance window sits on the guild line beside the table; accepted. Item 3: the bare `Fire resistance 42` is still shown. There is no target in the data, so I would hide the row, but it is a should, not a blocker; the other prep rows have context (icon, in your bags, named missing). Items 8 and 9 stay NICE.

Not must-fix, carry forward: hide the resistance row or add a target when RaidPrep has one (needs data); L6 next-24-hour dot window and the countdown on the chip line.

VERDICT: SHIP
