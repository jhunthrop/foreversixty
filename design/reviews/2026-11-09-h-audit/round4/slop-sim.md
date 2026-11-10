# AI-slop review, sim pages, round 4 (2026-11-09)

Method: all nine pages opened at 1440 and 390. Result and saved at 390 are 7-10k px and were read from a downscaled full-page view (layout and section order, not small type). Handed-off items and fixture artefacts ignored. Weights haste caption kept per owner rule.

## First impression
Reads as a player's tool. No glow, gradient, glass, em dash or marketing verb. Round 3 fixes landed: weights iteration caption and column-header explainer gone, "Open Top Gear" link gone, specs "Raid BiS DPS at 60" and row carets gone, drops "All professions" gone, gear "Equipped" repeats and duplicate Name labels gone, landing-in 390 talents collapsed to a one-line split, result rotation text kept only on non-obvious steps, BiS line cut to "1,004.0 . 9.9% under".

## Verdicts
- sim-landing-out: PASS
- sim-landing-in: PASS WITH FIXES
- sim-result: PASS WITH FIXES
- sim-saved: PASS WITH FIXES
- sim-gear: PASS
- sim-talents: PASS
- sim-drops: PASS
- sim-weights: PASS
- sim-specs: PASS

## Remaining findings (not in handoffs.md)
sim-landing-in, 1440+390 (SimRotationSection.svelte)
- 7/21 grey paragraph under all 12 rotation steps, including "Other cooldowns as they come up: Trinkets, racials and potions, as each comes off cooldown." and the Whirlwind/Bloodthirst lines. Result and saved already trim to the non-obvious steps; do the same here. Roughly 1,550 px of Rotation column.
sim-result / sim-saved, 1440 (+390 same components)
- 22 BiS gap still said twice: header "+/-1.0 DPS . -99.6 vs BiS" and the DPS spread card "Level 60 raid BiS 1,004.0 . 9.9% under". Keep the card, drop "vs BiS" from the header.
- 22 "Fury Warrior" repeated on each Your sims row on result (same spec as the page); date alone.

## Not verifiable
- sim-drops: Dungeons, Reputation, PvP still captured collapsed (third round running); only Crafted is open. Not failed, recapture expanded.

## Authentic, keep
904.4 DPS hierarchy; ability table with miss/dodge/glance; execute-phase timeline; rage ledger with lost at cap; buff uptimes; Your sims; weights table with bars and rating conversions ("10 Hit rating = 1% Hit"); talent diff with +5/-5 per talent; class-coloured spec list on a plate with rank and DPS bars; Crafted rows with icons, counts and ilvl; raid dates line ("Onyxia's Lair: First raids, 9 December 2026").

## Handed off (noted, not failed)
Timeline Off/On toggle unlabeled (f-content). Sim gear stat strip, BiS DPS target, "Item level N" on every gear row (h-doll).
