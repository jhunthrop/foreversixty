# Shaman builds from the 2026-10-07 talent search

Source: the reports regenerated today on engine d1af3529f (read from the main checkout, `design/reviews/talent-search/shaman-{elemental,enhancement}.md`). Rule: the winner keeps every unmodeled guide talent, holds the spec's own tree as the majority, gains at least 1% of baseline beyond the combined error, and drops nothing the rotation casts.

## Elemental

- Old: `FS1:1.60.1.70009:shaman:dwarf:4532310300103051/0/553322:` (31/0/20)
- New: `FS1:1.60.1.70009:shaman:dwarf:553231130010305/01/553302:` (32/1/18)
- Variant: "guide, modeled non-damage points re-spent". Taken: +1 Convection, +1 Elemental Focus, +1 Thundering Strikes. Dropped: Lava Burst, 2 Tidal Focus.
- Search gain: +12.6 DPS on 127.1, about +9.9%, combined error 0.7.
- Rule check:
  - Unmodeled talents kept: the finalist's "Drops unmodeled guide talents" is "-", and Elemental Warding, Improved Healing Wave, Improved Reincarnation, Mindfulness, Natural Grace are all unchanged (Tidal Focus is modeled, no damage).
  - Own tree majority: 32 of 51 points in Elemental.
  - Gain: 12.6 against a threshold of 1% x 127.1 + 0.7 = 2.0. Clears it.
  - Rotation casts: ladder level-60 distinct casts 5 before and after, same top casts, no violations, no zero-cast or unresolved entries. Lava Burst is not in the rotation, so dropping its talent costs no cast. The rotation uses no talent-granted ability that moved (Elemental Mastery is not in this tree and no variant touches it).
- Ladder level 60: 71.8 -> 76.6, about +6.7%. Lower rungs move within about -2% to +4%, no regression violations.
- Rejected: "deep Enhancement" (overall winner, +23.9 DPS): 18/28/5 fails the spec-tree majority and drops Mindfulness, Natural Grace, Improved Reincarnation, Improved Healing Wave and Elemental Warding. "guide, all non-damage points re-spent" (+20.3) and "deep Elemental" (+20.3) both drop unmodeled Natural Grace, Improved Reincarnation, Improved Healing Wave. Variants 7 to 9 (+16 to +17) each drop an unmodeled talent too.

## Enhancement

- Old: `FS1:1.60.1.70009:shaman:dwarf:553322/253130030005102051/0:` (20/31/0)
- New: `FS1:1.60.1.70009:shaman:dwarf:32303/255130030005102051/052:` (11/33/7)
- Variant: "guide, modeled non-damage points re-spent". Taken: +1 Call of Flame, +2 Ancestral Knowledge, +5 Totemic Focus, +2 Mindfulness. Dropped: 2 Convection, 3 Concussion, 3 Reverberation, 2 Elemental Devastation.
- Search gain: +3.9 DPS on 203.6, about +1.9%, combined error 1.0.
- Rule check:
  - Unmodeled talents kept: the finalist's "Drops unmodeled guide talents" is "-". Earth's Grasp, Elemental Warding stay; the dropped Elemental talents are classed as modeled (measured at near zero for this spec).
  - Own tree majority: 33 of 51 in Enhancement.
  - Gain: 3.9 against a threshold of 1% x 203.6 + 1.0 = 3.04. Clears it by 0.9.
  - Rotation casts: Stormstrike (17364), Maelstrom Weapon and Flurry are untouched. Level-60 ladder row has 7 distinct casts before and after with Stormstrike still firing (22.5% -> 23.0% of casts), no violations.
- Ladder level 60: 124.6 -> 126.2, about +1.3%. Level 40 distinct casts fall from 8 to 7 and DPS from 65.2 to 63.8, because the ladder's 9-points-per-level truncation of the new string reaches the same talents in a different order (Ancestral Knowledge before Flurry); at level 60 nothing is lost, and no regression violation is raised.
- Rejected: "deep Enhancement" variants (+34.7 and +35.2) drop Maelstrom Weapon (-5), Reverberation, Concussion; "Guide, all non-damage points re-spent" (same DPS, +3.9) drops Elemental Warding, an unmodeled talent. Totemic Focus single swaps (+2.4 to +2.9) are subsumed by the adopted variant.

## Changes

- Guides: `web/src/content/guides/shaman/{elemental,enhancement}.md` (build string and Talents section).
- Goldens: `sim/request/testdata/ladder/shaman-{elemental,enhancement}.golden.md` regenerated. No smoke-test edit was needed; the smoke test passes unchanged.
- Verified: `go test ./...` under sim/, `vitest run src/content src/lib/guides` (341 passed).
- Not touched: shaman-enhancement still equips an off-hand weapon on the ladder gear rule (`off_hand:22819`), contrary to the no-dual-wield rule. That is ladder gear selection in the engine/test harness, outside this lane.
