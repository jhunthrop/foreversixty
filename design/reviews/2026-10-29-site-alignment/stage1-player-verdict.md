# Stage 1 site alignment + home reorder: player verdict

Reviewed on screenshots: BiS Fury Warrior 20-29 (1440, 390), guide Fury Warrior (1440), sim, planner, home signed-in and signed-out (1440).

## Verdict: SHIP WITH FIXES (no blockers found in the stage 1 changes)

Gear list, stat weights, "Play it" and "Where to get it" on BiS are all still there, in the same order and as dense as before. The card chrome only adds a header strip and padding. The 20-29 pick list (Veteran's Silvered Chain Helm, Defias set, Demon Band, Shadowfang, etc.) is unchanged and reads as before.

## Blockers
None introduced by stage 1.

## Shoulds
1. Guide, "Talents and builds": "Level 60 and raid build" shows three grey placeholder bars instead of the talent summary, and the build panel's talent icons render as broken-image boxes. Same in the before capture, so not a regression, but the talent block is the thing a player came for. Confirm this is a capture artifact (icons/data not loaded) and re-capture; if real on the live page it is a blocker for the guide.
2. Guide, overview note: "The beta stops at level 30... Class mechanics carry over from 1.12 where the demo showed no change." Launch is Nov 4; on launch day this is wrong for a level-60 guide and reads like a dev note. Retire or date-gate it.
3. Planner: shows only Arms and Fury; a warrior has three trees (Protection missing). Talent icons are broken boxes and tiles show two-letter stubs ("Im", "De", "Bo"). Planner is not yet converted (before and stage1 identical), so likely the capture, but re-check before it is called done.
4. Sim: tab strip is now two rows, QUICK SIM... SPEC SUPPORT then "PLAN SIM LOGS RANKINGS FOR WARRIOR FORGET". "FORGET" is a button-like word with no meaning to a player; and "Your server runs did not load" plus "No sim with that id." are error copy on the first screen. Both were in before; fix the copy while converting.
5. BiS: the "Stat weights" scale column still has no "1 Strength = X DPS" line; the per-point DPS column carries it, but the Sixty Upgrades/wowsims convention wants the sentence. Unchanged by stage 1.
6. BiS 390: list at 390 is a very long single column; fine for dense data, but add slot anchors or a jump list.

## Home
- Signed-in: "Your upgrades" now sits right under the talents/sim tiles. Slot, you wear, best in slot, gain in DPS: this is what a player wants first. Good. The one-line "Name reservation open through Nov 3 · Launch Nov 4 · First raids Dec 9" is correct and readable. No "Help build this": fine.
- Signed-out: "Play your class better." with class picker and character lookup above the fold. Good. Countdown "05 Days 22 Hours 05 Minutes" correct for Oct 29.
- Signed-in "Another class" row is still present; "Latest" still last. Nothing a player relies on got harder to find.
