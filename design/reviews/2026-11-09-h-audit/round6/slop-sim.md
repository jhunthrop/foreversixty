# AI-slop review, sim pages, round 6 (2026-11-10)

Captures: round6-captures sim-landing-in, sim-result, sim-saved, sim-drops at 1440 and 390. Handed-off and fixture items ignored (Bloodrage curated prose, timeline toggle, gear strip, "Item level N", BiS target).

## Verdicts
- sim-landing-in: PASS WITH FIXES
- sim-result: PASS WITH FIXES
- sim-saved: PASS
- sim-drops: PASS WITH FIXES

## Fixed since round 5
Landing rotation no longer has prose under steps 1, 7, 8, 9, 10 or the "Other cooldowns" row. Your sims rows on result are date-only. Result at 390 is a native phone run (settings stack, DPS spread, folded Rage / Buffs / Rotation).

## Remaining
- sim-landing-in, 1440 + 390 (SimRotationSection.svelte): 21 Battle Shout (step 5) prose "Whenever it drops: 10 rage for the melee attack power of every party and raid member within 20 yards." restates the tooltip. Drop it; keep prose on Death Wish, Recklessness, Execute, Overpower, Heroic Strike.
- sim-result, 1440 + 390: 22 "Execute phase: Last 25% of the fight" with sub-line "0:45 of a 3:00 fight" (This run); give one figure. 22 "Upgrades for this character on Best in slot" under Gear; use "Upgrades on Best in slot" as on landing.
- sim-saved: none. Header meta "3:00 ± 20% fight, raid buffs, 8 consumables" has a stray "fight"; optional.
- sim-drops, 1440 + 390: 19/21 Crafted caption under the professions repeats itself: "No date yet: Blacksmithing, Blacksmithing, Blacksmithing, Blacksmithing, Engineering, Leatherworking, Leatherworking, Tailoring" plus a "First raids, 9 December 2026" clause per profession. One line per distinct date, names once (check for a duplicate bug in the date-line builder).
- sim-drops: Dungeons, Reputation, PvP captured collapsed for the fifth round; expand them in the capture.

## Authentic, keep
904.4 DPS hierarchy; ability table with miss/dodge/glance; execute-phase timeline; rage ledger with lost-at-cap; buff uptimes with counts; On you / On the target buff lists; raid date line; Droptimizer Crafted rows with icons, counts, ilvl; landing-in Your sims.
