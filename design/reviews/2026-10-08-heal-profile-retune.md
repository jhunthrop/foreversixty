# Healer profile retune (2026-10-08)

After the mana audit and the healer kit, the raid-ready healers covered 80 to 97% of the old profile's roughly 528
damage per second and the priests' mana lasted far past the fight, so healing met the damage and gear tied. This
retune makes the profile hard enough that gear separates, and states the rule so the next retune (Phase 2 gear)
repeats it. **Every figure below is measured under a stated incoming-damage profile and says nothing about which
healer class is stronger.** Band 60, alliance, 300 s fight, ranker output to a scratch directory (no nightly-owned file
touched).

## 1. The signal before (profile `onyxia-sized`, 528.5 damage per second)

Holy priest and Restoration druid, band 60, raid preset. Stat weights are the ranker's scale factors (Intellect = 1.00)
with their error; slot figures are the verified runner-up's effective HPS against the pick.

| | Holy priest | Restoration druid |
|---|---|---|
| Effective HPS / coverage / mana lasts | 514 / 97% / 2,290 s | 509 / 96% / 729 s |
| Healing power weight, error | 1.22, +/-0.40 (33%) | 0.74, +/-0.32 (44%) |
| Head: pick vs best alternative | Virtuous Crown vs Satin Hood: 0.0 (tie); vs the quest crown: 0.1 | Living Crown vs Feralheart Headdress: -1.2 |
| Chest: pick vs best alternative | Truefaith Vestments: -2.2 | Robes of the Exalted: -0.1 |

Across all 64 verified swaps of each spec (both factions, both presets, every slot), the median swap moved HPS by 0.7
(priest) and 0.7 (druid) against a median combined sim error of 0.4 to 0.5, and only 37 of 64 cleared the error.

**Verdict: no, gear could not be ranked on that profile.** Healing met the damage, the weights' error was a third to
a half of the healing power weight, and the top pick and its alternatives sat inside a point or two of effective HPS.

## 2. The rule

Scale the tank hit size and the raid-wide pulse damage together (ratio and cadence kept) by the **smallest
multiplier** at which, with the five healers at band 60:

1. the strongest raid-ready set covers about 80% of the incoming damage and the weakest about 60%;
2. every bare set (no raid consumables) covers less than half;
3. mana stays a live constraint: `mana_lasts_sec` is under the fight length for every healer bare and for at least
   one healer at raid.

The profile is a stand-in for Phase 1, scaled so sets separate. It is not a named boss, and its id and label say so.
The same rule is in `sim/cmd/leveling-bis/score_heal.go`'s doc comment and in `heal-profile.json`'s notes.

## 3. The sweep

Gear is re-chosen by the ranker at each multiplier. Columns per cell: effective HPS, coverage (HPS over incoming
damage), `mana_lasts_sec`, overheal.

**x1.25 (660.6 damage per second)**

| Spec | Raid | Bare |
|---|---|---|
| Holy priest | 642 / 97% / 406 s / 5.8% | 322 / 49% / 125 s / 0.7% |
| Discipline priest | 628 / 95% / 337 s / 7.4% | 383 / **58%** / 128 s / 1.7% |
| Holy paladin | 507 / 77% / 260 s / 0.6% | 179 / 27% / 82 s / 0.0% |
| Restoration shaman | 510 / 77% / 490 s / 1.3% | 216 / 33% / 102 s / 0.1% |
| Restoration druid | 599 / 91% / 279 s / 4.4% | 319 / 48% / 219 s / 4.4% |

**x1.5 (792.8 damage per second) - chosen**

| Spec | Raid | Bare |
|---|---|---|
| Holy priest | 669 / 84% / 236 s / 5.1% | 330 / 42% / 94 s / 0.2% |
| Discipline priest | 580 / 73% / 186 s / 2.4% | 373 / 47% / 101 s / 1.1% |
| Holy paladin | 496 / 63% / 236 s / 0.5% | 169 / 21% / 73 s / 0.0% |
| Restoration shaman | 565 / 71% / 280 s / 1.3% | 212 / 27% / 79 s / 0.0% |
| Restoration druid | 600 / 76% / 271 s / 3.8% | 318 / 40% / 218 s / 4.3% |

**x1.75 (924.9 damage per second)**

| Spec | Raid | Bare |
|---|---|---|
| Holy priest | 641 / 69% / 160 s / 2.3% | 307 / 33% / 82 s / 0.1% |
| Discipline priest | 598 / 65% / 138 s / 1.9% | 367 / 40% / 83 s / 0.8% |
| Holy paladin | 514 / 56% / 182 s / 0.3% | 178 / 19% / 60 s / 0.0% |
| Restoration shaman | 560 / 61% / 194 s / 0.9% | 208 / 22% / 63 s / 0.0% |
| Restoration druid | 604 / 65% / 272 s / 3.7% | 312 / 34% / 238 s / 4.4% |

**x2.0 (1,057.0 damage per second)**

| Spec | Raid | Bare |
|---|---|---|
| Holy priest | 622 / 59% / 126 s / 1.5% | 276 / 26% / 76 s / 0.1% |
| Discipline priest | 606 / 57% / 130 s / 1.3% | 366 / 35% / 74 s / 0.6% |
| Holy paladin | 530 / 50% / 128 s / 0.2% | 181 / 17% / 49 s / 0.0% |
| Restoration shaman | 531 / 50% / 126 s / 0.6% | 202 / 19% / 51 s / 0.0% |
| Restoration druid | 620 / 59% / 278 s / 3.5% | 319 / 30% / 222 s / 4.3% |

Reading against the rule:

- **x1.25** fails rule 2 (the bare Discipline priest covers 58%) and rule 1 (the strongest raid set covers 97%).
- **x1.5** meets all three: strongest raid set 84%, weakest 63%, every bare set under 50% (42, 47, 21, 27, 40), every
  bare mana under 300 s, and all five raid sets' mana under 300 s too (186 to 280 s).
- x1.75 and x2.0 also qualify on rule 2 and 3 but overshoot rule 1 (weakest 56 and 50%); the rule asks for the
  smallest multiplier, so they are not chosen.

The raid HPS columns are not monotonic in the multiplier because the ranker re-picks gear and talents' mana use at
each one; the coverage and mana columns are what the rule reads.

## 4. The chosen profile: `phase1-scaled`

`data/curated/heal-profile.json`: tank hit 1,150 -> 1,725 (every 3.3 s, about 523 per second), raid pulse 300 -> 450
(three members every 5 s, 270 per second), about 793 per second in all. Tank health, member health, cadence and the 25
percent spread are unchanged. The id changes from `onyxia-sized` to `phase1-scaled` and the label to "Phase 1 tank
hits and raid pulses, scaled so gear separates", because the figures are no longer an Onyxia-class boss's; the tank,
pulse, notes and summary reason fields say how they were chosen. Where the five guides named the profile ("the
Onyxia-sized tank hits and raid pulses") they now say "the scaled Phase 1 tank hits and raid pulses". The healer
panel caption and its profile block already render from the profile file (`healerFigureCaption(profileLabel)`,
`healerProfileTankLine(...)`); nothing in them or in the guide prose quotes a figure, so nothing else changes. Web
fixtures and their tests keep the old id and label as frozen sample data.

## 5. The result and the resolvability check

Re-run of all five healers, band 60, alliance, scratch out. This reproduces the x1.5 sweep row exactly (same seeds).

| Spec | Raid: HPS / coverage / mana lasts / overheal | Bare: HPS / coverage / mana lasts / overheal |
|---|---|---|
| Holy priest | 669 / 84% / 236 s / 5.1% | 330 / 42% / 94 s / 0.2% |
| Discipline priest | 580 / 73% / 186 s / 2.4% | 373 / 47% / 101 s / 1.1% |
| Holy paladin | 496 / 63% / 236 s / 0.5% | 169 / 21% / 73 s / 0.0% |
| Restoration shaman | 565 / 71% / 280 s / 1.3% | 212 / 27% / 79 s / 0.0% |
| Restoration druid | 600 / 76% / 271 s / 3.8% | 318 / 40% / 218 s / 4.3% |

Is the top pick now resolvable beyond the error? Yes for the slots, partly for the weights:

| | Holy priest before / after | Restoration druid before / after |
|---|---|---|
| Verified swaps beyond the combined sim error (of 64) | 37 / 58 | 37 / 51 |
| Median swap delta, HPS (median combined error) | 0.7 (0.4) / 2.5 (0.5) | 0.7 (0.5) / 5.1 (0.6) |
| Swaps larger than 3x the error | 29 / 53 | 24 / 44 |
| Healing power weight error, relative | 33% / 30% | 44% / 21% |

The healing power weight stays the noisiest row (still 20 to 45% across the five specs, with the Holy paladin's at
0.71 +/- 0.44), so a close call between two items that differ only in healing power is still not resolved by the
weights; the slot verifications, which sim the whole set, are what resolve it. Intellect, spirit, mp5 and crit sit at
4 to 14% error. Haste stays insignificant.

## 6. Tests

- `go test ./...` under `sim/`: passes after the five healer ladder goldens (`sim/request/testdata/ladder/*`) were
  adopted. Their healing-per-second and casts-per-level rows moved because the ladder fight is the profile fight; no
  rotation or engine change.
- `cd data && uv run pytest tests -q --no-cov -k 'heal or profile or preset'`: 11 passed.
- `cd web && npx vitest run src/lib/bis`: 13 files, 262 tests passed.
- No generated nightly-owned file was committed; the nightly `bis.yml` regenerates the healer files under the new
  profile.
