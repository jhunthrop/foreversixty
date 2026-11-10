# AI-slop review: guides, round 4 (2026-11-10)

Captures: round4-captures/guides, guides-you, guides-class, guides-spec at 1440 and 390 (all eight opened). Handed-off items (handoffs.md) and fixture icons do not fail a page.

## /guides: PASS
Nine class plates, spec rows, title band, no copy. Blank spec icons are fixture.

## /guides-you: PASS
- 17 "Your spec" pill beside the yellow rule on Beast Mastery (1440, 390, wraps under the name at 390). GuideSpecRow.astro, f-selected's, handed off.
- 22 "Your guide: Beast Mastery Hunter >" prefix in the band. lib/guides/copy.ts, f-content's, handed off.

## /guides/hunter: PASS
Spec plates with point splits, "What changed from 1.12" list, Races rail with "new" marks, plain Sources list. Round 3 Pets sentence trimmed to one clause. No definitional blurbs.

## /guides/hunter/survival: PASS WITH FIXES (no visual tells 1, 2, 5, 8, 12, 13, 15)
Non-handed-off:
- 22 Stat band footer "1 point of agility = 0.63 DPS at level 60": the band already says Level 60 and the weights table carries the DPS column. Drop the line, or print "0.63 DPS per agility". lib/guides/copy.ts (stat lane).
- 22 Weights table header "DPS per point or rating" (1440 Stat priority section): "DPS per point" is the unit; "or rating" qualifies nothing. Same lane.
- 23/21 Key talents sidebar, Lone Wolf: the tooltip line plus "Forever's pets copy only a tenth of your attack power and take no stat buffs from other players, so in a raid the talent beats the pet..." is a paragraph restating the Overview and Talents prose (1440 and 390). Keep "20% more damage with every attack while no pet is out." Rendered by GuideBuildsLane.
Handed off, still present, not failed on:
- 23 Lede "Survival is a melee spec in Forever: two weapons, Raptor Strike queued..." (survival.md, f-content). Same file: "You fire no shots in this spec. Your bow or gun is there for its stats." and "Strength is attack power directly. Plain attack power gives only itself."
- Sticky bar "On this page / Change" at 390 (Toc, f-selected); "DPS · Updated" band; broken spell icons; dead plate beside Talents prose; empty Strider Kick, Aspect of the Beast, Rapid Fire icons.

## Verdicts
- /guides: PASS
- /guides-you: PASS
- /guides/hunter: PASS
- /guides/hunter/survival: PASS WITH FIXES (three small text fixes, none visual)
