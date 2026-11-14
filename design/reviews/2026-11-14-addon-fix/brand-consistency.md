# Brand consistency: addon Stat weights block (2026-11-14)

Surfaces: Gear and Rotation tabs. Web counterpart: BiS weight rail (web/src/lib/bis/copy.ts, hit-cap.ts) and the sim weights tables (web/src/lib/sim/copy.ts).

## Matches the web
- Title "Stat weights" equals bisCopy.statWeightsLabel.
- Stat names come from the web's statLabel words (Strength, Agility, Ranged attack power, Spell power, MP5); rating rows read "Hit rating", "Crit rating" as the web's rail does.
- Top stat 1.00, two decimals for weight and DPS per point, matching the SimC-normalized rail (weightsScaleNote). Columns Stat / Weight / DPS per point read as the rail's per-point table.
- Rating hover "10 hit rating is 1% hit" is the guides' ratingConversion wording.
- Same block, same rows and numbers on both tabs; colours from Theme tokens, sentence case, no em dashes.

## Items (none blocking)
1. SHOULD: hit cap line. Web hitCapLine words it "Hit to cap: 2% for specials" and "Spell hit to cap: N%". The addon says "Ranged hit: 2% to the 9% cap". It carries the cap figure the web omits, and the hub spec 11 leaves the wording to the client's own, but it is a second phrasing of one fact. Fix: either move the pattern into shared copy (L.gearHitCapTo from bisCopy.hitToCapLine words) or record the ruling in the spec.
2. SHOULD: haste line. Web weightsHasteCaption prints the scale factor ("Haste: 1.58 per 1%"). The addon prints "Haste: %.2f DPS per 1%% haste" (DPS, not the normalized weight). Different unit for the same label; not visible in these captures. Fix: print the scale factor as the web does, or label it DPS clearly as it does now and note the ruling.
3. NICE: the web's scale note ("normalized to Hit = 1.00, with DPS per point and the sim error") has no addon counterpart; the source note ("your sim, ...") is the addon's own. Acceptable for a window this size.
4. NICE: "Capped" in the Weight cell has no web twin; the web keeps the number. Keep only if the hover explains it.

Platform idiom (allowed): in-game hover tooltips for rating conversion, fixed-width columns.

VERDICT: PASS
