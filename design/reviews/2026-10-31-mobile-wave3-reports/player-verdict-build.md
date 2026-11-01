# F10 Your reports: wow-player build verdict (iOS and Android captures)

Reviewer: wow-player, 2026-10-31. Read on the captures in `build/` and checked against `docs/contracts/fixtures/app` (reports.list, reports.get, reports.me, reports.meters, reports.meters.trash, coach.fight, me.raid-night). I did not open every capture: the signed-out, empty, error, loading, offline, not-yours, flagged, failed, processing and Android twins I did not read; the `-large` set I read only for iOS kill, iOS list and Android damage meters.

The player's question: "How did I do last raid?" The answer is first on the report: Onyxia, 712 DPS, 91 percentile, New best. That works, and the meters look like Details!: class colours, spec icons, DPS in mono, your row marked.

## Figures checked, and they agree
- 712.4 DPS appears as 712 in Your parses, the meter row (6th of 40) and the raid-night best. 91 percentile agrees in both places.
- Kill 4:28 = 268 000 ms everywhere. 40 players. 3 deaths = Rokhan, Hexlash, Kazrik in the meters.
- Trash 6 pulls: 95 + 120 + 64 + 41 + 88 + 52 s = 460 s = 7:40. Matches the row, the sheet title and the trash meters.
- 3 boss pulls (2 wipes at 61% and 22%, 1 kill) and 1 boss / 1 kill on the list row agree.
- Dates: 14 Dec 2026 is a Monday, "Onyxia, Monday" is right, and the UTC stamps convert to the evening of the 14th in US time. 11 Dec and 29 Nov also convert correctly.
- Fury shows two warriors above you only in the meters (Grimtusk 745, and none above 712 apart from rogue/mage/lock), which a 91 percentile can live with.

## MUST
1. **Home `Logs` row names the wrong "last raid"** (`home-logs-ios.png`). On Velaren, the Orc Beast Mastery hunter, it reads "Your last raid: Blackrock Spire, 0 kills". That is the warrior's live, private, still-running dungeon run (fixture: 4 fights, 0 bosses), not a raid and not Velaren's. The list under it is all "logged by Obnoxious Yell". The player's last raid is Onyxia, 1 kill. Pick the last finished raid with a boss pull, or say "Live now" for the running one. Also make the Home character and the report character one story (one character, one number), or the row will be screenshotted as a bug.
2. **Live capture shows a finished raid as Live** (`report-live-and.png`). It is the Onyxia 14 Dec report with a kill, a New best and 3 boss pulls, under the banner "Live. Updated 20 s ago." The list says that report is complete; the live one is Blackrock Spire. Capture the live state on the Blackrock Spire report (4 fights, no kills, "No kills of yours in this report."), or it contradicts the list on the same fixture.
3. **Largest text on the meter sheet loses your row** (`meters-damage-large-and.png`). The sheet shows 4 players, then half a screen of empty space, then "Show all 40". Your row is 6th, so you cannot find yourself, and the spec promises 11 rows. Also "Damage taken" wraps into a second line and breaks out of the segment bar. Show at least 11 rows or scroll, and keep your own row visible. Re-capture `meters-damage-large-ios.png` after the fix; I did not read it.

## SHOULD
1. **Three rows all called "Onyxia"** (`report-kill-ios.png`, `report-wipes-ios.png`). Wipe 1:12, Wipe 2:41, Kill 4:28 with nothing to tell the pulls apart. Warcraft Logs and Details! number them. Print "Onyxia · Pull 1" and so on, or "Try 1".
2. **Kill is slower than the wipe before it.** The raid got Onyxia to 22% in 2:41 with 11 deaths, then the kill takes 4:28 with 3 deaths. A raider reads that as a worse pull. Make the kill roughly 3:10 to 3:40, and keep 712 DPS and the coach text consistent (the coach quotes gaps at 1:12 and 2:40 to 2:51, which still fit).
3. **Fixture years.** `reports.get` fight start and end stamps say 2025-12-15, the report is 2026-12-15. Fix in `f10_fixtures.py`; if the app ever prints a fight time, it is a year early. Not visible on screen today.
4. **Trash deaths.** `reports.get` says trash pull 2 had 1 death; `reports.meters.trash` sums to 0. Make them agree.
5. **Two coaches, two stories for the same night.** `coach.fight` says "No flask at the pull"; `me.raid-night` says "No Elixir of the Mongoose or Juju Power" and "Whirlwind sat unused for 40 seconds", while the fight coach's top finding is Bloodthirst. In Classic you do not stack flask and elixirs, so a veteran will argue. Align the raid-night notes with the fight coach.
6. **Coach footer.** "Based on: Measured against Bloodthirst's 6-second cooldown and your own rage bar. From the character sheet in your log... Measured against your own rage bar." runs three sentences together and repeats itself. Dedupe and separate them (`report-premium-ios.png`).
7. **"New best" with nothing to beat.** Show the old number ("was 668 DPS") or the tag is a claim. The API sends `ranked: 1840` and the card never says "of 1,840 Fury Warriors"; add it under 91 percentile. "91 percentile" reads better as "91st percentile".
8. **Largest text** (`report-kill-large-ios.png`, `reports-list-large-ios.png`). The crest shrinks to about 24 px, the chevron and visibility glyph shrink to specks, and the Live, Processing and Failed tags drop below the title. Keep the crest and glyphs at a readable size; the tag can stay below.
9. **Meters**: no active time or share of total (Details! has both). Optional on a phone. Name rows are not tappable by design; fine.
10. **Damage taken bars.** Kraggor at 1,210 DTPS makes every other bar a stub, so the bars say nothing below row 2. Scale the bar to the first non-tank or cap it.

## Data I do not believe (for the accuracy reviewer)
- Onyxia total raid damage: 40 players at 600 to 790 DPS for 268 s is roughly 6 to 7 million damage, which is far above her health plus whelps. Check the sum against the client's Onyxia health.
- A 40-man Onyxia kill at 712 DPS being 6th of 40 and 91st percentile: plausible for Fury, but check the percentile curve against the Fury ladder at 1,840 ranked.
- Sapphiron Drape (17078) as an Onyxia drop replacing Cape of the Black Baron at item level 63: confirm both from the client.

## Verdict: HOLD

Fix the three MUSTs (Home Logs row, live capture on the wrong report, largest-text meter sheet), then re-capture those, and it ships. The report and list answer the question first and the numbers agree everywhere I checked.

---

# Round 2

Reviewer: wow-player, 2026-10-31. Read after the fix round: `report-kill-ios`, `report-kill-large-ios`, `report-premium-ios`, `home-logs-ios`, `report-live-and`, `meters-damage-large-and`, and the fixtures `reports.get`, `reports.get.live`, `reports.list`, `reports.meters.trash`, `coach.fight`, `me.raid-night`. I did not re-read the Android twins of the kill and home captures or the iOS large meter capture.

## Earlier MUSTs
| MUST | Now |
|---|---|
| Home Logs row named a live dungeon as "last raid" | **Fixed.** "Your last raid: Onyxia's Lair, 1 kill". |
| Live capture drew a finished raid | **Fixed.** `report-live-and` is Blackrock Spire, Tuesday 15 December 2026 (a Tuesday), Private, "No kills of yours in this report.", "No boss pulls", Trash 4 pulls 3:44. 48 + 71 + 39 + 66 s = 224 s = 3:44. |
| Largest text lost your row | **Fixed on Android.** Rows 1 to 4, then your 6th row with the gold edge, then "Show all 40". The "Damage taken" tab still wraps to two lines inside the bar, but it stays in the bar. |

## Earlier SHOULDs
- Pull numbers: done ("Onyxia · pull 1 of 3"). Good.
- Second wipe is now 4:52 at 22%, longer than the 4:28 kill. Reads right.
- Fixture dates: 2026 throughout, trash starts 02:05Z after the report (01:58Z) and the raid start (01:00Z). The Onyxia list row and `reports.get` agree on upload and heartbeat times (02:55:40Z, 02:55:30Z), and the last fight ends 02:55:12Z.
- Trash deaths agree (0), the raid-night coach notes now match the fight coach (Bloodthirst, flask, rage). Largest-text crest is bigger and the chevron legible.
- The kill sheet head at largest text on Android wraps "Onyxia · pull 3 of 3 Kill · 4:28 · 40 players" onto two lines. Readable.

## Still open, none blocking
1. **Home Logs row is another character's raid.** Home is "Next for Velaren" (hunter); the raid named is the warrior's. Say whose ("Obnoxious Yell: Onyxia's Lair, 1 kill") or scope the row to the shown character. Account-wide is defensible, but unlabelled it reads as Velaren's.
2. **Coach footer still repeats itself** (`report-premium-ios`): "Measured against Bloodthirst's 6-second cooldown and your own rage bar; From the character sheet...; Measured against your own rage bar." The first and third labels are not the same string, so the join does not remove them. Also a capital letter after each "; ". Dedupe in the coach data or lower-case after the semicolon.
3. Deferred by you: "was N DPS" on New best, "of 1,840", damage-taken bar scaling. Fine to follow.
4. Accuracy items from round 1 (Onyxia raid damage total, percentile curve, Sapphiron Drape and Cape of the Black Baron item level) still go to the accuracy reviewer.

## Verdict: SHIP

No MUSTs remain.
