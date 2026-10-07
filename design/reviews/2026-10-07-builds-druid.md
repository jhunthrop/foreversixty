# Druid builds from the 2026-10-07 talent search

Source reports: `design/reviews/talent-search/druid-balance.md` and `druid-feral.md` (engine d1af3529f, client 1.60.1.70009). Restoration is not a DPS spec and is untouched.

## Balance (adopted)

- Old: `FS1:1.60.1.70009:druid:night-elf:5222211015401051/0/55333:`
- New: `FS1:1.60.1.70009:druid:night-elf:5232221115400051/05/503301:` (34/5/12)
- Variant: finalist 9, "guide, modeled non-damage points re-spent". Search gain +38.7 DPS (+/-1.2 combined) on a 181.6 guide, about +21%.
- Moves: -5 Furor, -3 Natural Shapeshifter, -1 Nature's Grace; +1 Moonglow, +1 Nature's Reach, +1 Nature's Splendor, +5 Heart of the Wild, +1 Reflection.

Rule check:
1. Keeps every unmodeled talent: the "drops unmodeled guide talents" column is "-". Furor and Natural Shapeshifter are modeled and measure zero; Nature's Grace is now "modeled, no damage" (-0.45/point). Improved Entangling Roots (1), Nature's Focus (5) and Subtlety (3) are kept.
2. Own tree is the majority: 34 Balance of 51 points.
3. Gain: 38.7 is about 21%, far above 1% of the guide and the +/-1.2 error.
4. Rotation casts: level-60 ladder row has the same top spells (Starfire 9876, Insect Swarm 24977, Moonfire 9912, 1259799) with none newly idle; no `zero_casts` line for druid-balance before or after.

Ladder level 60: 108.3 to 121.1, about +11.8% (golden regenerated; levels 20 to 50 move similarly).

Rejected: finalists 1 to 8 (up to about +25%) all drop Subtlety, Nature's Focus, Improved Entangling Roots (all unmodeled), so they fail rule 1. Revisit when the engine gains those talents.

## Feral (adopted)

- Old: `FS1:1.60.1.70009:druid:night-elf:01/55232232121032012001/5053:`
- New: `FS1:1.60.1.70009:druid:night-elf:01/55232032121032012001/50532:` (1/35/15)
- Variant: finalist 9, "guide, modeled non-damage points re-spent". Search gain +5.2 DPS (+/-0.5) on 281.6, about +1.8%.
- Moves: -2 Thick Hide (3 to 1), +2 Natural Shapeshifter.

Rule check:
1. Drops column "-": Thick Hide is modeled and measures zero. All unmodeled talents (Brutal Impact, Feral Charge, Feral Instinct, Feral Swiftness, Nature's Focus, Subtlety) are kept.
2. 35 of 51 points in Feral.
3. Gain about 1.8% against a 0.2% combined error: clears the 1% margin beyond error.
4. Rotation: level-60 top spells unchanged. `zero_casts` for Ferocious Bite (22829) at levels 38, 40 and 60 is identical before and after; it predates this change. No new idle ability.

Ladder level 60: 188.1 to 192.6, about +2.4%.

Rejected: finalists 1 to 8 (up to +6.3%) drop Subtlety, Feral Instinct, Brutal Impact, Feral Swiftness or Feral Charge, all unmodeled.

## Verification

`go test ./...` under sim/ passes (no smoke expectation changes needed); `npx vitest run src/content src/lib/guides` passes (14 files).
