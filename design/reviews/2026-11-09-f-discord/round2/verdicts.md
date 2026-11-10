# Round 2 verdicts (lane f-discord, 2026-11-09)

Captures at cea40a96c.

## ux-designer: VERDICT: SHIP
- Legibility blocker fixed (phone readable). SHOULDS: "7 items:" colon; name sizes vary; row heights; picks not visible on the sheet; share head repeats the band; loot stack length; report card Killed line near the edge.

## wow-player: VERDICT: BLOCK
- BLOCKERS: "7 items:" dangling colon; card-report numbers contradict the report (1 h 14 min, 6 deaths, no wipes vs 2 h 41 min, 14 deaths, 3 wipes); bis-result source lines read kind words ("Alliance quest", "Vendors") instead of the sheet's sources.
- SHOULDS: sheet picks not shown; "Gain not modelled yet" a dead end; When and Guild time both shown; Buffs names long.

## ai-slop-reviewer: VERDICT: PASS
- Accepted the fan-made line (owner rule), "item level" and the lead panel corners. SHOULDS: Trinket 2 pick, twin big stats in the centre, Main hand meta wraps, "Killed" beside the count, the coach sentence.

## info-density-reviewer: VERDICT: DENSE WITH ADDITIONS
- Open: footer as band only, picks on gain rows, drop "Gain not modelled yet", Buffs line long at 8 of 40. LATER items accepted.

## guild-officer-reviewer: VERDICT: WORTH PAYING FOR
- SHOULDS: "Tanks · 3 of 2" should read as over target; the reminder says tomorrow twice; webhook raid posts could carry a role line. NICE: reminder fixture footer time.

## Round 3 fixes (this commit)
- Loot heading "7 items" (no colon); "10 items, the best 8 below".
- Report preview fixture is the same night as /report (3 wipes, 1:42 kill, 2 h 41 min, 14 deaths); a kill after wipes reads "Onyxia 1:42 after 3 wipes".
- Guided /bis result fixture carries the build's loot.json, so its source lines are the page's own.
- Sheet gain and upgrade lines shrink to 15 px to name their pick; a pick with no figure and no room leaves the line empty; the Upgrades figure is smaller than the item level.
- Over target reads "Tanks · 3 of 2, 1 over"; the reminder says tomorrow once ("is tomorrow, 8:00 PM.").
