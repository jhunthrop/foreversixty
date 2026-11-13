# Builds window, wow-player review, round 2 (2026-11-13)

**Verdict: SHIP WITH FIXES**

Round 1 fixes confirmed on the re-rendered boards:
- The druid Up next now matches the data (Feral Charge at 25, Savage Fury 26 to 27, Predatory Strikes 28 to 30, Leader of the Pack 31, Blood Frenzy 32, Natural Instinct 33 to 34). "26 more points after these" is right (51 - 15 - 10).
- The hunter count checks out: 19 points remain (Now plus levels 43 to 60), six rows cover 18, "1 more point after these" is right.
- Off plan talents are named (Hawk Eye is Marksmanship, 3 ranks; Deflection is Survival, 5 ranks, both match the client data).
- The build name reads "level 60 build" everywhere, matching the row label. The raw `\194\183` is gone. "Levels 44 to 48" reads right.
- The wrong-code message now says what to copy and that it starts with FSB1.

## Remaining must-fix

1. The Off plan lines read badly: "Hawk Eye · Marksmanship · 3 off plan (Rank 3/3)". The player has 3 of 3 ranks, so "3 off plan" and "Rank 3/3" say the same thing twice. Use "Hawk Eye 3/3 · Marksmanship" and "Deflection 2/5 · Survival", and the red heading already says it is off plan.
2. Header "27 of 33 points · 5 points off plan · 1 to spend" still adds up only if the reader does the sum (27 + 5 + 1 = 33). Say "32 of 33 points spent" or keep the line and drop "of 33". Pick one reading.

## Should-fix, not blocking

- Spec "Level 60 build" against a player's word "leveling build" is now consistent, but "Level 60 and raid build" on a one-row spec is clumsy next to the two-row specs. "Level 60 / Raid build" would read cleaner.
- No respec cost line when Use this build replaces 15 points. A druid at 24 pays gold, and the red Replaces line should say so.
- Gear rows in the paste preview have no icons or sources (already on the follow-up list; fine for now).
