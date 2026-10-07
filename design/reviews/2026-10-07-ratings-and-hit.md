# Ratings and hit: the ranker's units and the hit weight (2026-10-07)

Lane `ratings`. Site branch `ratings` (from `main`), fork branch `ratings` (from `forever`).
All figures are build 1.60.1.70009, band 60, the default level-63 target, weights at 100
iterations per direction, the scratch runs described at the end.

## Findings in one paragraph

The ranker did not over-score rating items. `main.go` has divided every candidate's hit, crit,
dodge, parry, block and defense by the build's `combatratings.txt` factor since the rating-units
lane, so Lionheart Helm already scored as 2% hit and 2% crit. The unit bug ran the other way: the
division was blind, so the 347 items whose hit, crit, dodge, parry, block or defense comes from an
on-equip aura (a literal percent, Fury Visor's 1 hit and 1 crit) were scored at a tenth to a
fourteenth of their value. That is fixed with a per-item `percent_stats` field. The small hit weight
was a separate defect in the engine's stat-weight sweep: a +-20 step across a floor and a cap
averaged a saturated response and published about a fifth of the real marginal value. That is fixed
in the fork, and the band output now carries `hit_to_cap`. A third finding, outside this lane's
files, is that the Fury weights character does not dual wield (see "The weights character").

## The units chain

| Stage | What `hit` / `crit` mean | Where |
| --- | --- | --- |
| Client `ItemSparse` | rating points (ITEM_MOD 31, 32, 12-15) | `normalize/gear.py` `STAT_BY_MODIFIER_ID` |
| Planner `items/<class>.json` `stats` | the same rating points, unchanged | `resolve_item_values` |
| On-equip spell (`equip.STAT_AURAS` 47, 49, 51, 52, 54) | a literal percent | `_merge_effect_stats`, `classicdb_items.equip_stats` |
| classic-db `raw_stats` | rating points | `classicdb_items.planner_stats` |
| wowhead planner payload | rating points | `wowhead_items.planner_stats` |
| `simdb.bin` | percent: ratings divided by the level-60 factor, aura percent untouched | `simdb/ratings.py`, `items.py`, `enchants.py` |
| Engine | one point is one percent (`HitRatingPerHitChance` and `CritRatingPerCritChance` are 1) | `spell_result.go` |
| Ranker `score()` | must be engine percent | `data.go` `convertRatingStats` |

The factors are 10 hit rating, 14 crit rating, 10 haste, 10 expertise per percent (verified in
`combatratings.txt` for every level), and 12 dodge, 15 parry, 5 block, 1 defense at level 60.
Only the simdb path and the ranker divide; the planner JSON keeps the client's number so a tooltip
matches the game.

Before this lane the planner JSON did not say which share of a rating-family stat was an aura. The
ranker therefore divided all of it. `GearItem.percent_stats` (`models.py`) now records the aura
share, set in `build_class_items` (from `EffectIndex.stats`) and `classicdb_items.to_gear_item`
(from `equip_stats`); a wowhead row has none. `convertRatingStats` computes
`(amount - percent) / factor + percent`, which is exact even for an item that carries both (none do
on this build: `_merge_effect_stats` refuses to sum the two units, and a data test pins that every
`percent_stats` entry is a positive share of `stats`). 347 distinct items carry one; every one of
them in the warrior file is a classic-db row.

`items/<class>.json` was patched in place rather than regenerated: the local raw export has drifted
from the committed build (a regenerate rewrote 700 to 1500 unrelated rows per class and the
talents files), and the nightly owns that drift. The patch adds `percent_stats` to every row from a
regenerate's values and changes nothing else; the nightly's next normalize emits the same field.

Pinned by tests: Lionheart Helm scores 2 hit and 2 crit, Fury Visor 1 and 1, Arcanoweave Cloak 1
hit (`ratings_units_test.go`, on the committed warrior data); a 1% aura hit and a 10-rating hit
score identically; the Python data tests assert the same two items' `percent_stats`.

Enchants and set bonuses: the ranker scores neither from stats. Enchants never reach `score()`,
and set bonuses are decided by an engine run (`sets.go`), where `simdb.bin` already holds percent.
`simdb/enchants.py` already divides enchant ratings. Nothing to change there.

## The hit curve

The 20000-iteration paired-seed DPS of the actual weights character (`weightsRequest` at band 60,
bare preset), with `n` points of hit added through `BonusStats`. Level 63 target, 300 skill: base
miss 8%, suppression 1%, so specials stop missing at 9% hit.

### Fury warrior, the character the sweep measures (one weapon, 0% hit)

| hit | -2 | -1 | 0 | +1 | +2 | +5 | +9 | +12 | +20 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| DPS | 126.53 | 126.53 | 126.53 | 126.53 | 128.22 | 133.33 | 140.15 | 140.15 | 140.15 |

Flat to +1 (the suppression dead zone), 1.69 DPS per point from +1 to +9, flat from +9: the yellow
cap sits exactly at 9%. With one weapon the white swings share that cap.

The old sweep took +-20 around 0: the low side is floored (0 DPS change), the high side runs past
the cap (140.15 - 126.53 = 13.62 over 20 points = 0.68 per point), and the two are averaged to 0.34
DPS per point. That is the published 4.59 attack power per percent (0.459 per rating point; 0.34 /
0.0757 DPS per attack power point). The marginal value at the live edge is 1.69, which is 22.4
attack power per percent: the new sweep's published figure. The curve and the sweep agree.

### Fury warrior with a real dual wield pair (the same character, two one-hand swords)

| hit | -2 | -1 | 0 | +1 | +2 | +5 | +9 | +12 | +20 | +28 | +30 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| DPS | 173.37 | 174.42 | 175.31 | 176.16 | 178.87 | 187.08 | 197.91 | 204.30 | 218.71 | 227.63 | 227.63 |

Slope 2.7 per point from +1 to +9 (specials and whites), about 1.8 per point from 9 to 28 (white
swings only, with the 19 point dual wield penalty), flat from 28: the white cap is 9 + 19 = 28.
Below +1 the off-hand's own table is still live (Human's sword skill moves its dead zone), which is
why the curve is not flat at -2 to 0.

### Rogue (Combat), 3% hit from talents

| abs. hit | 0 | 1 | 2 | 3 | 4 | 5 | 8 | 9 | 12 | 15 | 23 | 28 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| DPS | 216.02 | 216.18 | 218.47 | 220.78 | 223.13 | 225.43 | 232.35 | 234.62 | 240.43 | 246.17 | 260.51 | 269.30 |

Dead zone to 1, 2.3 per point to 9, about 1.9 to 1.8 per point after (whites), cap at 28.
`hit_to_cap` reads specials 6, white 25 from the 3% baseline.

### Mage (Fire)

| hit | -2 | -1 | 0 | +1 | +2 | +5 | +9 | +12 | +16 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| DPS | 182.25 | 185.18 | 188.23 | 191.17 | 194.27 | 203.47 | 216.23 | 222.86 | 224.32 |

Spell hit has no floor and a long linear run (3.0 per point to 12, flat by 16). A +-1 step is exact
there; the old +-20 reached across the cap and under-read it. No `hit_to_cap` is published for a
spec that does not swing a weapon: the spell cap depends on school and talent bonuses the engine
applies per spell.

## The fixes

Fork (`sim/core`, commit on branch `ratings`):

- `hitprofile.go`: `ComputeHitProfile` builds the request's environment without simulating and
  reads the player's total hit, the main-hand (or ranged) table's suppression and miss, and whether
  white swings carry the dual wield penalty. `hitStepWindow` picks the sweep window: +-1 where
  neither side is floored or capped, the live part of it otherwise (from the suppression edge up
  one point when the character sits in the dead zone, down one point when it sits on the cap), and
  none past the white cap, where hit is worth nothing and the weight is left at 0.
- `statweight.go`: hit takes that window instead of +-20. `computeStatWeights` reads any window
  that is not symmetric about the baseline as the slope between its two simulations, with a
  hard-cap check against the low simulation. Attack power, strength, agility, intellect, armor and
  mana keep their +-20 steps. A request with no player keeps +-1.

Site (`sim/cmd/leveling-bis`, branch `ratings`):

- `data.go`: `candidate.PercentStats`, `convertRatingStats(stats, percent, factors)`.
- `hitcap.go`, `simrun.go`, `report.go`, `main.go`: `engineRunner.HitProfileFor`; every band record
  carries `hit_to_cap` as `{baseline, specials, white?}` in engine percent. `baseline` is the weights
  character's hit, `specials` the distance to the 9% yellow cap, `white` the distance to the dual
  wield white cap, present only for a dual wielder. Casters publish none. The web types are not
  changed in this lane; `hit_to_cap` is an additive key.

## Before and after

Scratch runs, `-bands 60`, `-out` a scratch directory; "before" is the unchanged binary, "units" is
the site fix over the old sweep, "after" is both. Hit and crit are the published
`weight_per_percent` (attack power per percent for the melee specs; the mage's own reference stat).

| Spec / preset | Hit before | Hit after | Crit | Hit/crit before | Hit/crit after | `hit_to_cap` |
| --- | --- | --- | --- | --- | --- | --- |
| warrior-fury bare | 4.59 | 22.42 (+-2.68) | 32.08 | 0.14 | 0.70 | 0 baseline, specials 9 |
| warrior-fury raid | 6.83 | 33.40 (+-3.21) | 31.42 | 0.22 | 1.06 | 0 baseline, specials 9 |
| rogue-combat bare | 14.42 | 25.45 (+-1.90) | 33.60 | 0.43 | 0.76 | 3 baseline, specials 6, white 25 |
| rogue-combat raid | 14.88 | 30.20 (+-2.46) | 32.36 | 0.46 | 0.93 | 3 baseline, specials 6, white 25 |
| mage-fire bare | 8.73 | 12.52 (+-1.13) | 7.99 | 1.09 | 1.57 | none |
| mage-fire raid | 9.13 | 12.16 (+-1.02) | 11.88 | 0.77 | 1.02 | none |

Per rating point (what the site rail shows) the Fury bare figures are 0.459 before and 2.242
after against crit 2.292.

Picks changed, summed over the four records of each spec (bare and raid, alliance and horde; 17
slots each, 68 picks):

| Spec | Units fix alone | Sweep fix on top | Before to after |
| --- | --- | --- | --- |
| warrior-fury | 9 | 26 | 29 |
| rogue-combat | 4 | 12 | 14 |
| mage-fire | 12 | 18 | 22 |

Bare, alliance, warrior-fury, after both: neck Medallion of the Dawn becomes Mark of Fordring, back
Cape of the Black Baron becomes Howler's Furs, wrist Forest Stalker's Bracers becomes Battleborn
Armbraces, feet Drudge Boots becomes Boots of Heroism, ring Protector's Band becomes Signet Ring of
the Bronze Dragonflight. The units fix alone moved the neck (Mark of Fordring), shoulder and ring.

## The weights character: Fury is not dual wielding

`ladderMeleeWeapons` (`character.go`, not this lane's file) picks the two highest-DPS one-handers
whose planner slot allows `main_hand` or `off_hand`. For Fury at band 60 that is Andonisus
(22736) and Glaive of the Defender (23051). Both are `HandTypeMainHand` in the engine's database.
`Equipment.EquipItem` puts a main-hand-only weapon in the main hand, so the second overwrites the
first and the off hand stays empty: `AutoAttacks.IsDualWielding` is false and the character swings
one weapon (126.5 DPS at 0% hit against 175.3 with a real one-hand pair). So the white dual wield
cap (28%) is not in play in the Fury sweep or in `hit_to_cap` (which says so honestly: specials 9,
no white figure). The planner `slot` does not separate one-hand from main-hand-only weapons, so the
fix needs the item's hand type in the planner data or an engine-side check in the ladder; it is for
the lane that owns `character.go`. Once Fury dual wields, its hit weight rises (2.7 against 1.7
DPS per point) and `hit_to_cap` gains a white figure with no change here.

## Follow-ups

- Rogue-combat's profile reads dual wielding, so its weapon pair is valid.
- The sweep's hit window uses the main-hand table. The off-hand table can differ by a point of
  suppression through weapon skill (Human swords), visible as the sub-9% slope above; that is below
  the sweep's resolution and not modelled.
- `hit_to_cap` is the weights character's distance, not the BiS set's. The site can subtract the
  set's hit from the specials figure itself; the engine cap (baseline plus specials) is 9 for every
  melee spec against this boss.

## Reproduce

`make simdb && make sim/internal/spellranks/spellranks.json`, then
`go run ./sim/cmd/leveling-bis -spec warrior-fury -bands 60 -out <scratch>` under a `go.mod`
replace pointing at the fork's `ratings` branch. Scratch outputs for this review: before, units and
after directories under the session scratchpad; none of them is committed.
