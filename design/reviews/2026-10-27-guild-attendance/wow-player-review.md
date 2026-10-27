# Guild attendance and the raid-night report: wow-player review (Lane G2)

Reviewer: wow-player persona (Classic raid officer since 2019, loot council, Details!, WCL, Raid-Helper and RCLootCouncil user).
Seen rendered: rows 1-11 and 13 in this folder (full pages and panels at 360, 390, 1440 and 2000 where captured).
Row 12 (the `report_ready` embed and the copied text in Discord clients) is not captured, so it is not approved here.

## Verdict: SHIP WITH FIXES

The bones are right. The night line `13 of 14 came · 1 kill · 3 wipes · 17 deaths · 2 h 41 min` is the
answer a raid leader wants and it is the first thing under the title. Kill order with kill time and
wipes, deaths per raider in Details! style, consumables against the readiness check, loot per boss,
and the five biggest changes are the things I would otherwise collect from WCL and RCLootCouncil by
hand. The words are guild words: `Not in the log`, `Late, first seen 22 min after the first pull`,
`Joined from the bench or without a sign-up`, `Excused`, and the member's line `If you were there under
another name, tell an officer.` "No-show" appears nowhere a player reads; it is only in capture file
names. `Felsnap as Mirestalk` counting as present is the thing that makes an officer trust the sheet.
What stops a clean SHIP: empty loot icon squares, figures that contradict each other on screen, a raider
who was not there shown as having used no consumables, and a finished raid that still offers `Cancel raid`.

## Findings

1. **must-fix. Every loot row has an empty icon square.** Seen in 01-done-officer-report (1440 and 390),
   02-member-present-1440 and 08-read-only-1440: `Head of Onyxia`, `Netherwind Crown`, `Ring of Binding`
   and `Vis'kag the Bloodletter` each sit beside a blank dark box. On Onyxia night the loot is what people
   screenshot, and four blank squares read as broken. Fix: load the item icons (ItemHover's 28 px icon) in
   the fixture build and in production. Do not ship a placeholder box; if an icon cannot load, show the
   question-mark icon the client uses, with the quality border.
2. **must-fix. The member "not in the log" view still shows a DPS figure.** 02-member-no-show-panel-1440:
   `Signed In as DPS. You are not in this log. If you were there under another name, tell an officer.`
   and under it `DPS 196 on Onyxia, up 16 since Thu 3 Dec`. A character who is not in the log has no DPS
   on that boss. Fix: when the outcome is not in the log, drop the deaths and the DPS/HPS change facts.
   Keep only `Excused` if set.
3. **must-fix. A raider who was not in the log is listed in Consumables as having used nothing.**
   01-done-officer-report-1440 and 08-read-only: `Hexbramble  SHORT  No flask  No food  No potions`.
   Hexbramble is `Not in the log` in Attendance. This is wrong, and it is the shaming row R-A3 rules out.
   Fix: Consumables lists only characters in the log. A came-as row uses the buffs of the character
   that was actually in the log.
4. **must-fix. The roster contradicts the raid.** 09-roster-raids-panel-1440: `Kraggor ... No raids
   yet`. Kraggor is the leader who uploaded the log and is `Present 4/4` on this raid. Zulmara reads
   `Came 6 of 8 · 24 of 30` on the roster but `came 7 of 8 raids` on her guild line (09-guild-line-390).
   An officer who sees either one stops trusting every attendance number. Fix: build both from the same
   attendance read, and make the fixture consistent so the captures prove it.
5. **must-fix. A done raid still shows the officer planning row.** 01-done-officer (1440 and 390) and
   08-read-only-1440: `Edit · Lock sign-ups · Copy the sign-up list · Cancel raid` stay in the band on a
   `DONE` raid, some greyed out and some live. "Cancel raid" on a raid we already killed gets mocked in
   guild chat. Fix: on `done`, remove the whole row; the only officer control in the band is `Not this log`.
   On an ended raid with no log (03-ended-officer captures), keep `Edit` (fixing the time helps the match)
   and `Match a log`. Remove `Lock sign-ups`, `Cancel raid` and `Copy the sign-up list`.
6. **must-fix. Notes and the sign-up roster sit open, with live dropdowns, below Attendance.**
   01-done-officer-1440, 08-read-only-1440 and 02-member-present-1440: `Notes` and `Sign-ups 14/40` are fully
   expanded, and every officer row still has an `In / Late / Bench / Out` select. Changing a sign-up after
   the raid quietly rewrites attendance. Spec R-A4 says these stay below Attendance, closed; the build does
   not do that. Fix: on `done`, show both as closed disclosures (`Notes`, `Sign-ups 14/40`), with the roster
   read-only when opened. After the raid, an officer edits in Attendance and `Excused`.
7. **must-fix. Death names are cut off on phone.** 01-done-officer-report-390: `Windtal...`, `Shadow...`,
   `Lightbr...`, `Mendw...` and `Arcane...`, and `3 on wipes · 1 on kills` wraps to three lines. The name is
   the one thing the Deaths list must show. Fix: below 768, shorten the bar to 64 px or drop it, keep the
   full name, and put `3 on wipes · 1 on kills` on a second line under it.
8. **should-fix. The pull count has no unit, and it looks like it contradicts Late.** 06-attendance-panel
   and 01: `Arcanemoor  Late, first seen 9 min after the first pull  4/4`. A raid leader reads "late, but
   there for every pull?" The 4 counts boss pulls, while late counts from the first trash pull. Fix: say
   `4 of 4 boss pulls` in full rows, or give the column a `Boss pulls` head. Compact rows can keep `4/4`
   under that head.
9. **should-fix. Signing Late and arriving when you said still counts as Late.** 06: Arcanemoor signed
   `Late · DPS` with the note `There by 8:30.` and is still in the `2 late` summary. Every attendance sheet
   I have kept treats someone who said they would be late, and was, as fine. Fix: a character signed Late
   is excused from Late by default. The row reads `Signed Late, first seen 9 min after the first pull`,
   and the summary does not count it as late.
10. **should-fix. The night line leaves out the people who joined.** Band, every done capture: `13 of 14
    came` while 15 raided (Soulgrave from the bench, Thornhide unsigned). A raid leader's first question
    is "how many did we have in?" Fix: `13 of 14 came · 2 joined · 1 kill · ...`, leaving the clause out
    at zero. That matches the Attendance summary.
11. **should-fix. The picker shows upload time, not raid time.** 04-picker-closest-panel-1440: rows read
    `Thu 10 Dec, 11:42 pm` and `Fri 11 Dec, 12:20 am`. Upload times drift, and a second logger's partial
    log goes up late. Officers match by when the pulls happened. Fix: lead with the log's first pull
    (`First pull 8:04 pm`) and show the upload time muted on the second line.
12. **should-fix. Biggest changes loses its heading when there is no previous raid.**
    07-bare-report-panel-1440: `First Onyxia's Lair raid logged. Changes show from the next one.` floats
    under the Loot line with no sub-head. Fix: keep the `Biggest changes` sub-head and put the line under
    it, the way Loot does.
13. **should-fix. The duration splits across lines on phone.** 01-done-officer-panel-390:
    `17 deaths · 2 h` on one line, `41 min` on the next. Fix: never break inside a figure and its unit
    (`2 h 41 min` stays on one line, and each clause between the dots stays whole).
14. **should-fix. Readiness shows gear for a raider with no export.** 10-readiness-on-panel-1440:
    `Shadowquill  4 upgrades · +28 DPS  BOOTS ... 59 (-2)  No export yet`. With no export, where did the
    gear gap come from? Fix: if the source is the site's saved gear, the cell says
    `Gear from the site, no addon export yet`; otherwise the row shows dashes.
15. **should-fix. Deaths for a came-as raider are under the signed name.** 01: Felsnap has one death in
    Deaths, but Felsnap played Mirestalk that night. Fix: show `Felsnap as Mirestalk` in Deaths,
    Consumables and Changes too, with Mirestalk's class crest if it differs.
16. **should-fix (a follow-up is acceptable). What a raid leader wants on Onyxia night is missing.**
    (a) For each wipe, the boss health and phase when it happened (`wipe at 38%, phase 2`), as in WCL's
    Fights list. Three bare wipes do not show whether we are close. (b) What killed each raider
    (`Deep Breath 2`, `Whelps 1`), as in Details!'s Deaths segment. (c) Fire resistance: the Notes say
    `Bring fire resistance`, but Consumables cannot show a `Greater Fire Protection Potion` and folds it
    into `Potions 3`. At least name the potions on hover.
17. **nit. `Bench and out` rows do not say which.** 06: `Frostnettle  0/4`, with 0/4 in the bright
    "missed pulls" colour. Fix: a muted `Out` or `Bench` in place of `0/4`.
18. **nit. A read-only officer sees the gate sentence twice.** 08-read-only-1440: `Planning and running
    raids are part of the guild plan. Subscribe the guild.` appears under the band and again under Report.
    Keep the one under the band.
19. **nit. In Your night the whole facts line is in the figure font.** 02-member-present-1440:
    `DPS 196 on Onyxia, up 16 since Thu 3 Dec` is all monospace; only the figures should be.
20. **nit. Missing space.** 10-readiness-on: `Sorted worst first.Nudge copies a message...`
21. **nit. A one-boss raid shows the loot count twice:** `LOOT 4`, then `ONYXIA 4`. That is fine for BWL;
    for a one-boss zone, drop the boss head.

## Data I do not believe (fixture; for the accuracy reviewer)

- `Onyxia 1:42` after 3 wipes, with 15 raiders. A fresh guild's Onyxia kill takes 3-6 minutes; 1:42 is
  speedrun territory, and Onyxia is a 40-player raid. Raiders will mock it if this fixture ends up in
  promotional captures. Use about 4:30 and a 38-40 roster.
- `Consumables 6` for 15 raiders: either the fixture is thin or rows are being dropped. Confirm every
  raider in the log gets a row.
- `Windtalon  1 upgrade · +0 DPS` on the readiness board: a +0 DPS item is not an upgrade.

## Answers to the two questions

- **Officer row on a done raid:** no. Remove `Edit`, `Lock sign-ups`, `Copy the sign-up list` and
  `Cancel raid`; `Not this log` in the match line is the only control in the band. An ended raid with no
  log keeps `Edit` and `Match a log` only (finding 5).
- **Notes and the sign-up roster below Attendance:** keep them, but closed and read-only. Attendance is
  the record of the night. The sign-up roster is the plan, and its selects must not stay live after the
  raid (finding 6).
