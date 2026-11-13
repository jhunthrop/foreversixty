# Fractional base values and the box-only weapon test (2026-11-13)

Data and test fix lane. Scope: the two items the 70938 engine follow left open (engine-70938.md, "Follow-up"): spell effect amounts were rounded in the data, and `TestWeaponsSwingAtTheirForeverSpeed` failed on the Windows box but not on CI.

**Result: the rounding is gone from the data, three engine tables that existed only because of it now read the client rows, and nothing published moves. The weapon-speed failure is not limited to the box. The committed 70938 item data is wrong, and CI never runs the test.**

## 1. Fractional base values

### Where the rounding happened

There was one place: `data/pipeline/spelltext.py` `_base_points`. It returned `round(float(EffectBasePointsF))`, which is Python's banker's rounding (36.5 became 36, 107.5 became 108). Every reader inherited the rounded value through `effect_amount`:

- `pipeline/simconst.py` (spellconst/*.json, which the engine vendors and generates its constants from)
- `pipeline/simdb/equip.py` (item equip stats)
- the tooltip renderer (`_bounds`, `_amount`, `work_out`)

Coefficients (`sp_coefficient`) and per-level growth (`points_per_level`) were already floats and never rounded. No other pipeline module rounds or truncates SpellEffect base points. `classicdb_items.py` and `audit/dumpdb.py` read Era's integer column, which has no fraction to lose.

### The fix (9052b36f1)

- `_base_points` and `effect_amount` return the client's exact float.
- `SpellEffectConstant.amount` is a float, and a whole amount is still written as an integer, so the generated files only change where a fraction exists.
- Display is the one place that rounds. A sentence rounds half up in `_bounds`/`_amount`, and a `${..}` formula works on the exact value and rounds half up in `work_out`. Rounding only at display is what the game does.
- New tests:
  - `test_effect_amount_keeps_the_clients_fraction` checks 25.2, 36.5 and 107.5.
  - `test_a_fractional_amount_is_written_exactly_and_a_whole_one_as_an_integer` covers spellconst.
  - `test_a_fractional_amount_is_kept_and_rounded_only_when_printed` checks Wrack: 60.2 prints 60, and a formula over six ticks gives 361 where the old rounding gave 360.

### Method

I fetched 70938's ten spell tables from wago.tools into a scratch build directory and layered the 70938 push over them with `layer_hotfix_push` against the committed store. Built with the old rounding, those tables reproduce the committed spellconst exactly: 0 mismatches across every field of every spell. Built with the fix, only the 117 amounts below change.

### (a) Every 70938 class-file row with a fraction

There are 117 effect rows on 108 spells: mage 43, hunter 28, priest 17, warlock 12, shaman 5, paladin 4, rogue 4, druid 3, warrior 1. Most are rank-0 NPC copies (Shoot, Throw, Fire Blast, Blast Wave, Holy Fire, Flash Heal 1232758, Baron Rivendare's Soul Drain and similar), which no engine rank reads. The player rows:

| Spell | Row | Client | Was | Reader |
|---|---|---|---|---|
| Wrack | 1316697 eff 0 | 25.2 | 25 | engine (was a hand row, MR41), conformance golden |
| Lightning Bolt overload r4, r5, r9, r10 | 408442, 408443, 408475, 408477 | 36.5, 54.5, 147.5, 176.5 | 36, 54, 148, 176 | engine (were hand rows, MR46) |
| Chain Lightning overload r1 | 408479 | 107.5 | 108 | engine (was a hand row, MR46) |
| Judgement of Fury r2 | 1311655 eff 0 | 36.9 | 37 | engine (`JudgementOfFuryDamage`), conformance golden |
| Judgement of Fury r1, r2 eff 2 | 1311650, 1311655 | 194.4, 316.8 | 194, 317 | no reader (effect 3, a dummy) |
| Seal of Fury r3 eff 0 | 20163 | 486.9 | 487 | generated `SealOfFuryBaseDamage` only; the engine reads its own proc table |
| Health Funnel r1, r2, r4, r7 eff 3 | 755, 3698, 3700, 11695 | 5.45, 10.49, 25.4, 61.2 | 5, 10, 25, 61 | no reader (not the generated primary effect, and the engine does not model Health Funnel) |
| Rupture rank 0 | 1271514 | 8.70 | 9 | `RuptureBaseDamage[0]`, unused (rank 0 is an internal copy) |
| Binding Heal r5 | 1240773 | 644.0001 | 644 | float noise only |

Tooltips: the five rendered texts that change belong to the overload proc spells (408443, 408475, 408477, 408479) and NPC Tranquility 25817. The site publishes those ids only as names in spells.json, never as text, so no visible tooltip, guide or card number changes. No guide or card quotes these values.

### (b) Hand rows retired (engine 5861a96ab, ff2172004, 45c009ce3; site pin f4da53474)

- **Wrack (MR41):** `WrackTickDamage` is `clientdamage.FromTable(WrackBaseDamage...)[1]` and `WrackTickCoefficient` is `WrackSpellCoeff[1]`. `TestWrackTickMatchesTheClientRow` checks the table against the vendored client row at every level, and `TestWrackTickReadsThe70888Push` keeps pinning 25.2, 60.2 and 0.2002.
- **Judgement of Fury:** `JudgementOfFuryDamage` was a hand table, and rank 2 stated 37. It is now `FromTable` over the generated rows, so rank 2 reads 36.9.
- **Lightning and Chain Lightning overloads (MR46):** the separately pinned hand rows and their test (`TestLightningOverloadRowsAreThe70888Client`) are gone. `TestLightningOverloadDamageMatchesClient` checks both tables against the client rows by id with the shared 0.05 roll tolerance, the same way every other shaman rank table is checked. The overload arrays stay in `lightning_bolt.go`. The generator cannot emit them, because they share "Lightning Bolt" and "Chain Lightning" with the hand-declared bolt arrays.
- Vendored client files and every class's `constants_auto_gen.go` were regenerated from the same tables.
- The conformance goldens moved only in the client columns (Wrack 25.00 to 25.20, Judgement of Fury rank 2's roll).
- The protection paladin suite results moved in the fifth decimal.
- The fork's `./sim/...` suite is green on the box. The one exception is `sim/web` [setup failed], which also fails on master because it needs binary_dist.

### (c) Before and after

Measured on the box: 10,000 iterations, seed 1, one sim at a time, alliance.

| Entry | 10cbe9f5e | 45c009ce3 |
|---|---|---|
| paladin-retribution bare 20 | 49.45 +/- 0.03 | 49.45 +/- 0.03 |
| paladin-protection bare 20 | 26.37 +/- 0.02 | 26.36 +/- 0.02 |
| paladin-retribution raid 60 | 833.35 +/- 0.51 | 833.35 +/- 0.51 |
| warlock-affliction raid 60 | 768.68 +/- 0.28 | 768.68 +/- 0.28 |
| shaman-elemental raid 60 | 747.98 +/- 0.31 | 747.98 +/- 0.31 |

Wrack and the overloads were already exact in the engine, so they cannot move. Judgement of Fury rank 2 only applies at levels 18 to 25, and the change is inside the error there. The site's ladder goldens do not move either. No number moves, so no guide prose, no wow-player SHIP and no re-rank row are needed for this part.

The committed `data/builds/1.60.1.70938/spellconst` and spell text regenerate through data.yml. I verified the scratch build byte for byte against the committed files apart from the 117 amounts, and the engine's vendored copies are that output.

## 2. `TestWeaponsSwingAtTheirForeverSpeed`: the failure is not limited to the box

- **The box runs the same data and engine as CI.** Its `simdb.bin` sha1 (9e97a2018f) is the committed 70938 file. The engine is the same pin, through the same `go mod edit -replace` that CI's pin-engine uses.
- **The committed data is wrong.** `data/builds/1.60.1.70938/items` has Fiery War Axe (870) at 2.9 s and 40.34 DPS, where 70790 had 3.6 s. Blade of Hanna (2801) went back from 2.1 s to 2.0 s. Those are exactly the round-5 rows `pipeline.wago_hotfixes` was written to fix.
- **Cause:** wago.tools' `db2/ItemSparse/csv?useHotfixes=1` export no longer applies push 16778086 on any build: 70532, 70707, 70790, 70888 and 70938 all return ItemDelay 2900 for 870. wago's hotfix page still records that push for 70938 with ItemDelay 3600 (the newest enUS entries carry `data: None`). The 70790 run on 2026-11-12 13:05 still got 3600 from the same export, so wago's export changed before the 70938 run (2026-11-13 06:53).
- **Why CI is green:** the repo variable `LAUNCH_MODE` is `true`, so sim.yml runs `go build` and `go vet` but not `go test`. Nothing on CI runs this test. tools/remote-test.sh and the box are correct, so they are unchanged. The test fails wherever it runs, on purpose.
- **Not fixed here:** the real fix is a data-source change, either crawling the item pushes from the hotfix page (about 80k ItemSparse entries for 70938) or cross-checking the CSV against the page. That needs its own data lane: **issue #27**.

### Box measurements affected

Since 2026-11-13 07:04 UTC (52cd11091, the 70938 data commit), every run on the active data, on the box or in CI, uses 870 at 2.9 s and 2801 at 2.0 s.

- The published sets that wear them are:
  - paladin-retribution bare 40 (870) and bare 60 (2801)
  - shaman-enhancement bare 40 (870)
  - warrior-arms bare 40 (870)
- No raid 60 entry wears either, so the raid 60 measurements of the engine-70938, casters-60, feral-60, rogues-60 and the other 2026-11-13 lanes are not affected.
- A lane that measured those bare 40 or bare 60 Retribution entries since 07:04 measured the wrong weapon. The 64 other 70938 item stat differences against 70790 are counted in issue #27 and still need checking.
- MR105 re-ranks those bands once #27 lands.

## Shas

- Site: 9052b36f1 (data fix), f4da53474 (engine pin 45c009ce3)
- Engine (wowsims-forever master): 5861a96ab, ff2172004, 45c009ce3
- Issue: #27
