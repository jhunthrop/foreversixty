# Sim-accuracy round 25: platform lane

Site branch `wt/r25-platform` (from int/r25 ab2029586). Fork branch `lane/r25-platform` (from int/r25 bf727699f). Client 1.60.1.70790. BUILD MODE: touched tests only, no regens. Every sim went through `$R/sim.sh`. Every CONFIRMED-BUG has a test that was red before its fix.

**Confirmed and fixed: 11.** P1, P2 `R21 559e39867`; P3 `old`; R1 to R4 `R21 559e39867` plus `old` (blocked by P1 and P3); R5, R6 `R24 2ae0fe3b8` (judged on P1/P2 gear); H1, H2 `old`.
**Cross-class DPS movers:** fork a87011646 (sim/core) touches every class, but no published number moves (measured below). The site's below-band gear (S 2979f34dd) changes only `spec-breakdown -level 10` to `19`, which is a measuring tool.

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Still no new live Forever logs. `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct, no kills). Nothing in this lane needs a log. | last looked 17:21 UTC 6 Nov | OPEN (no data) | - | - |

## Seed 1: unarmed and bowless below 20 (cross-cutting)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Below 20 the band-20 gear lost its weapons.** `spec-breakdown -level 10` to `19` runs the band-20 entry and `wearableGear` dropped every slot whose item needed a higher level, with nothing in its place. Rogues fought unarmed below 19 (Blackfang 2236 needs 20, Assassin's Blade 1935 needs 19). Hunters had no bow below 20 (Ranger Bow 3021 needs 20). Survival lost its melee weapon, casters lost their wands, and so on. A combat rogue at 12 did 4.8 DPS. Now `request.LevelGear.Fit` refills a slot the level cannot hold. It takes the band's own runner-up first (the entry's `alternatives`, then `runners_up`, same handedness). If none fits, it takes the class item table's best weapon, shield or wand for the slot at that level, using the ladder's own pickers. A dagger stays a dagger (Backstab), a launcher stays a launcher, a two-hander stays a two-hander, and a two-hand main hand leaves no off hand. | `sim/cmd/spec-breakdown/main.go` (`levelGear`), `sim/request/level_gear.go`; test `TestNoSpecBelowTheLowestBandIsUnarmedOrBowless` (every spec, both factions, 10 to 19: nothing worn above its level, every weapon, shield or wand slot of band 20 filled, hunters hold a bow, gun or crossbow). Red: 1,555 violations. `TestACasterBelowTheBandKeepsAWand` was red too. | CONFIRMED-BUG (tool) | R21 559e39867 (`-level` below 20 arrived 4 Nov) | S 2979f34dd |
| P2 | **The below-band gate read the client's required level alone.** A quest reward the client lists at 0 was held at 10. Example: Forsaken Greataxe 251533, from The Wrath of Rath'mael (quest level 22, floor 19), on Retribution, Enhancement and Arms. The gate now uses `leveling.EffectiveRequiredLevel` with the loot.json quest floors, the level the ladder and the BiS ranker already use. Below 20 this also empties every armor slot whose item is a quest reward the level cannot have yet, unless a runner-up fits (armor is not refilled from the client table). | test `TestALevelBelowTheBandHoldsNoQuestRewardItCannotHaveYet` (red: level 12 held 251533) | CONFIRMED-BUG (tool) | R21 559e39867 | S 2979f34dd |
| P3 | **An APL condition on a spell the character does not have yet dropped out of its gate.** Round 22 (E 30bf8bdb0) fixed this for unknown auras, but spells still built to nil, and `and`/`or`/`not`/compare drop a nil operand. Beast Mastery and Marksmanship gate Serpent Sting and Arcane Shot on "Multi-Shot's time to ready >= 0.5 OR mana <= 25%". Below 18, with no Multi-Shot, the first term vanished, so both shots waited for a quarter mana that never came. A level 10 to 17 hunter fired neither. The engine now reads an unknown spell as one you cannot cast: canCast, isReady, inFlight and isChanneling read false; timeToReady reads 24 h (after any fight); castTime, travelTime, cpm, channeledTicks and currentCost read 0. spellIsKnown is unchanged. Beast Mastery bare at 12 goes from 32.0 to 43.3 DPS. | fork `sim/core/apl_values_spell.go` (`unknownSpellValue`); test `TestAnUnknownSpellReadsAsOneYouCannotCastInsteadOfDroppingTheGate` (red with the fix stashed); site comment `sim/request/rotation_ranks.go` castKeys updated | CONFIRMED-BUG (engine, shared) | old (round 22 fixed auras only) | E a87011646 |

## Seed 2: the rank rows the bare bands blocked (level by level, bare band 20, 10,000 iterations, seeds 1 and 50001)

Method: an exploration pass at 3,000 iterations over every ranked spell the card casts that gains a rank at 10 to 19, then the leads at 10,000, both seeds and both factions (Troll and Dwarf for hunters, Tauren and Night Elf for Feral), then joint runs where two rows touch the same fight. Every row's forced rank was confirmed cast (per-spell ids in `-json`).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **Beast Mastery Arcane Shot: rank 1 at 14 to 16 and 18 to 19.** Every rank adds 15% of ranged attack power, so rank 1 keeps most of the hit for less mana. Rank 1 alone gains +0.12 to +0.42% at 14 to 16 (all four runs). At 17 every lower rank loses (-2.2% with both shots at rank 1). At 18 and 19, rank 1 of both shots gains +0.98 to +1.28%. At 13 the gain is +0.05%, within error, so 13 stays on rank 2. | `$R/platform/confirm-h-*.txt`, `joint-out.txt`, `joint2-out.txt`; test `TestRound25LowLevelHunterRankRows` (red before the rows) | CONFIRMED (rank row) | R21 559e39867 + old (blocked by P1, P3) | S e0223103d |
| R2 | **Beast Mastery Serpent Sting: rank 1 at 18 and 19,** with R1 (+0.98 to +1.28% together; alone +0.24 to +0.55%). At 15 and 16 rank 1 gains under 0.15% alone and loses with rank 1 Arcane Shot (-0.1 to -0.3%), so rank 2 stays. | same | CONFIRMED (rank row) | same | S e0223103d |
| R3 | **Marksmanship Arcane Shot: rank 1 at 14 to 15, 18 to 19, and at 13 for a Troll.** 14 and 15: +0.16 to +0.33% (all runs). 13: Troll +0.19/+0.25, Dwarf -0.33/-0.36, so the row is keyed to `troll`. 16: +0.07/+0.03 alone and -0.9% with both, so rank 2 stays. 17: -3.0% with both. 18 and 19: rank 1 of both +0.71 to +1.05%. | same | CONFIRMED (rank row) | same | S e0223103d |
| R4 | **Marksmanship Serpent Sting: rank 1 at 18 and 19,** with R3. At 15 rank 1 ties (0.00/-0.04). | same | CONFIRMED (rank row) | same | S e0223103d |
| R5 | **Feral Wrath: round 24's rank 1 row ran to 13; it now ends at 11, and rank 2 holds through 15.** Round 24 judged the row unarmed below 18 and in quest gear the level cannot hold (P1, P2). On the gear the level holds: rank 1 beats rank 2 by 10 to 11% at 10 and 11. At 12 and 13 the rank 2 learned at 6 beats rank 1 by 7.5 to 10%. At 14 and 15, rank 2 beats the rank 3 learned at 14 by 3.2 to 4.4%. From 16, rank 3 wins by 22% or more. | `confirm-h-*.txt`, `joint-out.txt`; tests `TestRound25FeralCasterRankRows`, `TestFeralCasterRankRowsBelowTwenty` (round 24's, moved to the new judgement) | CONFIRMED (rank row) | R24 2ae0fe3b8 (judged on P1/P2 gear) | S e0223103d |
| R6 | **Feral Moonfire: rank 1 at 10 to 12 (a Night Elf takes rank 2 at 11), rank 2 from 13.** At 10, rank 2 loses 0.3%. At 12 it loses 9% with Wrath rank 2. At 11 a Tauren loses 0.6% with rank 2 and a Night Elf gains 1.0 to 1.2%, so the rank 1 row at 11 names Tauren and both Skyborne races. At 13, with Wrath rank 2, rank 2 gains another 7% over Moonfire rank 1. Round 24's rank 2 row at 16 to 19 holds (rank 3 -9 to -11%). | same | CONFIRMED (rank row) | R24 2ae0fe3b8 | S e0223103d |
| A1 | **Energy and rage ranks: the newest rank wins,** as it must when every rank costs the same. Combat and Assassination Backstab at 12 and 13: rank 1 -8.7 to -15%. Fury Heroic Strike at 16 and 17: rank 2 -24 to -26%. Arms Rend at 10 and 11: rank 1 -14 to -17%. Battle Shout at 12 and 13: rank 1 -4 to -5%. Assassination and Combat cast no Eviscerate at 14 to 17 (see O2). Subtlety casts no Backstab below 20 (its band-20 main hand is a sword). | `explore1.txt` | Aligned | - | - |
| A2 | **Survival Raptor Strike: the newest rank wins** at every level from 10 to 19 (lower ranks -1.3 to -2.7%). | `explore1.txt` | Aligned | - | - |
| A3 | **Enhancement Earth Shock: the newest rank wins** (lower ranks -4 to -10%). Lightning Bolt is not cast below 20 by design: its only line wants five Maelstrom Weapon stacks and all five points, and the talent arrives from 37. Searing Totem has one rank below 20. | `explore1.txt`; `data/curated/apl/shaman-enhancement.json` | Aligned | - | - |
| A4 | **Retribution: the newest ranks win** for Seal of Righteousness (-0.1 to -8%), Holy Strike, Judgement of the Crusader and Seal of the Crusader at 10 to 19, except at 16. At 16, Holy Strike rank 1 and Seal of the Crusader rank 1 each gain +0.50 to +0.56% at both seeds. Two unrelated spells gaining the same amount at one level is a mana threshold, not a rank effect: that level's gear leaves the paladin out of mana for 19.9 s of 180, against 0.6 s at 15 and 1.2 s at 17, and rank 1 Holy Strike loses about 1.5% at 15 and 17. Declined: a single-level row on a gear artifact would read wrong on the card. | `confirm-h-*.txt`; Ret 15 to 17 breakdowns above | Aligned (16 declined, reason given) | - | - |

## Seed 3: platform leftovers

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The hotfix-push alert could never open its issue.** client.yml's push-alert saw the unlayered 70790 push on 5 and 6 Nov and failed both nights. `tools/client-alert.sh` creates its `hotfix-push-unlayered` label with a 127-character description, and GitHub refuses one over 100 (`HTTP 422: description is too long (maximum is 100 characters)`). That is why the stale record went unnoticed. The description is now 78 characters (the other label's is 77), and the stub `gh` in `tools/tests/test_activation.sh` refuses a long description the way GitHub does. | run https://github.com/jhunthrop/foreversixty/actions/runs/42210369952 (6 Nov) and 42019257346 (5 Nov); `bash tools/tests/test_activation.sh` red (7 failed), then `ok 165 tests` | CONFIRMED-BUG (alerting) | old (26f2ccc6f, 3 Nov) | S 13aac3758 |
| H2 | **Nothing flagged a push record that layered nothing.** The 70790 activation (e65ee3982) committed `hotfix-push.json` with every search total at 0. `hotfix_push.empty_push_reason` now names such a record. `layer_hotfix_push` logs it as a warning (activation is not blocked: a brand-new build may not be on wago yet). A data test holds the active build's committed record to a real push, so CI goes red on a record like e65ee3982's. | `data/pipeline/hotfix_push.py`; tests `test_a_record_whose_search_found_nothing_is_reported`, `test_the_active_builds_push_record_layered_a_push` (red: no such function; e65ee3982's record reports as empty); `test_hotfix_push.py` 44 passed, ruff clean | CONFIRMED-BUG (data guard) | old | S 13aac3758 |
| F1 | **Natural flask names in simbuffs: done.** Round 24's finish (28fb9c1f6, in int/r25) named all four `flask_of_natural_*` ids through `loot-merge --buffs-only`. The merge night keeps a committed entry. The full `loot` step names them from the client Item table. `test_loot_build.py` holds the active simbuffs to IDS.md. No code gap remains. | `data/builds/1.60.1.70790/simbuffs.json`; `data/pipeline/loot/buffs.py` | Aligned | - | - |

## Before/after

**P3 (fork a87011646), cross-class.** Raid 60: 28 specs, both factions, seeds 1 and 50001, 10,000 iterations, frozen BiS, guide raid builds, `$R/before.tsv` against `$R/platform/raid60-after.tsv`. All 112 values are identical. Bare band floors 20/30/40/50: 28 specs, horde, seed 1, 3,000 iterations, `$R/sb-before` against `$R/platform/sb` (`bare-before.tsv`, `bare-after.tsv`). All 112 values are identical. P3 acts only below the level where a gated spell is learned.

**P1/P2 (the tool), bare band 20, horde, seed 1, 3,000 iterations:** Combat 12: 4.76 -> 15.45. Combat 16: 9.47 -> 26.42. Beast Mastery 12: 11.97 -> 32.05 (43.27 with P3). Beast Mastery 16: 15.57 -> 47.42.

**Rank rows** (gain over the rotation's default rank, seeds 1 / 50001, horde then alliance; absolute horde seed 1):

| spec | level | row | horde | alliance | horde s1 abs |
|---|---|---|---|---|---|
| Beast Mastery | 14 | Arcane 1 | +0.26 / +0.31 | +0.12 / +0.16 | 50.11 -> 50.24 |
| Beast Mastery | 15 | Arcane 1 | +0.42 / +0.42 | +0.35 / +0.31 | 56.06 -> 56.30 |
| Beast Mastery | 16 | Arcane 1 | +0.36 / +0.32 | +0.33 / +0.29 | 64.42 -> 64.65 |
| Beast Mastery | 18 | Arcane 1 + Serpent 1 | +1.21 / +1.22 | +0.98 / +1.06 | 79.59 -> 80.55 |
| Beast Mastery | 19 | Arcane 1 + Serpent 1 | +1.28 / +1.20 | +1.13 / +1.13 | 82.17 -> 83.22 |
| Marksmanship | 13 | Arcane 1 (Troll) | +0.19 / +0.25 | (Dwarf -0.33 / -0.36, no row) | 46.15 -> 46.24 |
| Marksmanship | 14 | Arcane 1 | +0.18 / +0.18 | +0.17 / +0.16 | 50.34 -> 50.43 |
| Marksmanship | 15 | Arcane 1 | +0.27 / +0.25 | +0.27 / +0.33 | 55.94 -> 56.09 |
| Marksmanship | 18 | Arcane 1 + Serpent 1 | +0.86 / +0.85 | +1.05 / +1.04 | 79.74 -> 80.42 |
| Marksmanship | 19 | Arcane 1 + Serpent 1 | +0.87 / +0.88 | +0.72 / +0.71 | 82.85 -> 83.57 |
| Feral | 12 | Wrath 2 (was 1) | +9.80 / +9.95 | +8.90 / +9.06 | 8.50 -> 9.33 |
| Feral | 13 | Wrath 2 + Moonfire 2 (were 1, 1) | +16.70 / +16.71 | +15.99 / +16.13 | 8.31 -> 9.70 |
| Feral | 14 | Wrath 2 | +4.27 / +4.35 | +3.29 / +3.19 | 10.26 -> 10.70 |
| Feral | 15 | Wrath 2 | +3.71 / +3.78 | +4.30 / +4.13 | 10.40 -> 10.78 |
| Feral | 11 | Moonfire 2 (Night Elf, the default) | (Tauren -0.60 / -0.59, keeps 1) | +1.04 / +1.15 | - |

Error at 10,000 iterations is about +/- 0.03% of DPS on each run.

## Files to regenerate (integrator, at the pin)

- `TestRotationLadder` goldens (sim/request): every hunter rung below 18 (P3), the new hunter and Feral rows (R1 to R6), and any other rung where a line's gate names a spell not yet learned (P3). No raid or band-floor number moves.
- The addon data step (`addon-data.json`, `Data.lua` rotation rows; not the BiS picks): `addonrotation` reads `rank-choices.json` for the card's ranks (R1 to R6).
- Coach rules and starters (`persona_starter.py --all`) if they read the new rank ids (`addonrotation.coach_persona_spell_ids` reads rank-choices).
- `sim/rankchoices/rank-choices.json` was synced by hand (a copy of the curated file, as `make rank-choices-sync` does).

## Measure at refresh

- Nothing new for the BiS: raid and band checkpoints do not move.

## For other lanes

- **Every lane with rank rows at 10 to 19** (balance 4, Restoration druid 8, Holy paladin 1, Protection paladin 2, Discipline 1, Holy priest 1, Elemental 2, Restoration shaman 1): these were judged on the old below-band gear. That gear had no weapon or wand where the band's item needed more than the level, and it held quest rewards the level cannot have yet (P1, P2). Re-judge them with `spec-breakdown -level N -rank id=rank` on S 2979f34dd or later. A mana-bound caster row can move, as Feral's did.
- **Any lane:** a condition on a spell the character has not learned now reads as a spell it cannot cast (P3). A gate that relied on the old drop now holds. For example, `spellTimeToReady(X) < 1.5s` with X unlearned used to vanish and now reads false. No published number moved.
- **Rogue (O2):** at 14 to 17, Combat and Assassination cast Slice and Dice at one or two combo points every time it drops, so with Backstab's one point Eviscerate never reaches five. Check whether the low-level line wants a longer Slice and Dice or Eviscerate at fewer points.
- **Paladin (A4):** the band-20 Retribution gear at 16 runs out of mana for 20 s; at 15 and 17 it does not. Worth a look at the level-16 item mix.

## For the guide lane

- The rank lines changed in this lane: `hunter/beast-mastery.md:161`, `hunter/marksmanship.md:167`, `druid/feral.md:111` (the prose test passes).

## Still open

- L1 (logs).
- Faction is not read by the gear pickers below 20, nor by the ladder's (`pickGearItem`): an Alliance character can be handed a Horde quest reward (Staff and Axe of Orgrimmar, Wingblade). Quest faction is in loot.json. Small DPS effect; it touches the ladder goldens, so it is a separate change.
- Below 20 an armor slot the level cannot hold, with no runner-up that fits, stays empty. Armor is not refilled from the client table (weapons, shields and wands are).
- Fork `go vet ./sim/core` fails on `sim/core/test_generators.go:535` (copies a proto lock value). Pre-existing, not this lane's.
