# wow-player review: addon 2.0.49 candidate (Gear + Rotation stat weights)

Seen rendered: 04-window-gear.png, 05-window-rotation.png, before-04-window-gear.png (BM hunter, level 60).

VERDICT: SHIP (no must-fixes), with 3 shoulds and 2 data checks for the accuracy reviewer.

## What reads right
- Gear and Rotation now show the same block: "Stat weights", Stat / Weight / DPS per point, top stat Hit rating 1.00. The unit is on screen, which is what the SimC convention wants. Rows are sorted high to low.
- "Hit rating" / "Crit rating" are the right words for rating stats. The old "Hit" / "Critical Strike" and the "per rating point" note are gone.
- Attack Power plus Ranged Attack Power as two identical 0.33 rows was a visible bug; one "Ranged attack power" row is correct for a ranged hunter.
- "Ranged hit: 2% to the 9% cap" reads like the client's own wording. The old "Spell hit: 9% to the 16% cap" on a BM hunter is gone from the screen, which is the right look for a ranged hunter.
- Aimed Shot "On cooldown, shared with Multi-Shot" matches what Classic hunters know (Aimed Shot and Multi-Shot share the 6 s lockout).
- Rotation arithmetic checks: eight bosses average 85%, downtime per boss sums to 92 s.

## Should
1. Priority rows 2 and 3 contradict each other. Bestial Wrath says "On cooldown, with Rapid Fire and trinkets", Rapid Fire says "As Auto Shot fires with Multi-Shot nearly ready". Pick one story: Rapid Fire is either lined up with Bestial Wrath or timed to Multi-Shot, not both.
2. The hit line ("Ranged hit: 2% to the 9% cap") sits in the same grey and size as the table rows right under Strength, so it reads as a seventh row with no value. Add a gap or label it as a cap line.
3. Rotation tab's "Low output" (Gehennas) has no unit or number while the others are seconds. Cosmetic, not from this change.

## Data I do not believe yet (for the accuracy reviewer)
- Strength 0.07 on a BM hunter. In Classic a hunter's Strength gives melee attack power only, no ranged attack power. A weight of 0.07 on a ranged spec means the sim values melee AP (Raptor Strike, melee auto in the dead zone, or the pet). Confirm that is modelled on purpose; if not, the row should not exist. It also uses one of the 6 row slots.
- Dropping the spell hit line for BM/MM. Classic players know Serpent Sting and Hunter's Mark roll the spell table (Nature, Arcane), and the old 16% spell cap is the number they remember for those. The claim that they now roll ranged hit in Forever needs the engine's own proof (hit table the sim uses for Serpent Sting), not only the absence of spells in the priority. If Serpent Sting and Hunter's Mark are in the priority and roll spell hit, the line was right and the "Spell hit" row should stay for BM.
- Agility 0.53 x 0.42 = 0.22 but the DPS column shows 0.23. Fine if the underlying weights are unrounded; check it is rounding, not a different source.
- Unchanged by this fix but a player will ask: "Black Dragonscale Helm" as head on a hunter. I know the Black Dragonscale set as boots, breastplate, leggings and shoulders; confirm that helm exists in the Forever client.
