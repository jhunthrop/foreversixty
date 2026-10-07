# Warlock: Unstable Affliction and Curse of the Elements (2026-10-07)

Client build 1.60.1.70009. Fork branch `warlock-ua`, site branch `warlock-ua`.

## What the client says

### Unstable Affliction: present in the client, not learnable

The spellconst rows exist: 427717 (rank 1, tooltip level 40, 86 a tick), 1242971 (rank 2, level 60, 123) and 1242972 (rank 3, level 50, 174), each a 1.5 s cast, six 3 s ticks of shadow damage, sp_coefficient 0.2 a tick, family mask [0, 256, 0, 0]; the live text matches ("Only one Unstable Affliction or Immolate per Warlock can be active on any one target"). Rank 2 is learned at 60 and rank 3 at 50, which looks like a client slip but is what the table says. Helper rows 1219436, 427719 and 431748 are the dispel-silence and override-spell plumbing, not casts.

Evidence it cannot be learned in this build (coordinator's check of the SkillLineAbility table, wago, same build): no learn row for any of the three ids. `data/builds/1.60.1.70009/talents/warlock.json` names none of 427717, 1242971 or 1242972 and no talent text mentions the spell. So it was implemented and then withdrawn: the fork commit `ae4be8689` removes the registration, the Immolate exclusion and the table test. It is not in any rotation. The search tool still lists it as a candidate (spellranks.json carries the rows); it reads "no effect" because the engine does not register it.

### Curse of the Elements: reworked

Four learn rows, ids 440892, 1311676, 1311677, 1311680 at levels 20/30/40/50, costs 50/100/150/200 mana, instant, 5 min, effect 0 aura 22 (resistance -30/-45/-60/-75) and effect 1 aura 87 (damage taken +4/+6/+8/+10 percent), both over misc value 126, which is every magic school (arcane, fire, frost, holy, nature, shadow). "Only one Curse per Warlock can be active on any one target." The brief quoted the rank-3 text (8 percent); the level-60 curse is rank 4, 10 percent, 75 resistance (id 1311680).

The engine used to register vanilla ids 1490/11721/11722 at levels 40/50/60, shadow-school, with the raid aura (fire and frost only). Now `sim/warlock/curses.go` registers all four client ranks by level (ids, levels and costs read from `constants_auto_gen.go`), each with its own aura that multiplies damage taken for all six magic schools and removes the rank's resistance, and the cast replaces any other curse the warlock had on the target (same `ActiveCurseAura` rule as Agony, Doom, Recklessness). The warlock's own shadow and fire damage in a solo sim therefore gets the percentage.

Disagreement left in place: `core.CurseOfElementsAura` (the raid-debuff path, `core/debuffs.go`) is still the vanilla fire and frost, 10 percent, 75 resistance curse. Magnitudes agree with rank 4; the school set does not (client: all magic). It is outside the warlock lane and shared by every class's debuff setup.

## Tests and conformance

Fork tests first (`curse_of_the_elements_test.go`): four ranks registered at level 60 with the client ids, levels, costs and instant cast; each rank multiplies all six magic schools by 1.04/1.06/1.08/1.10 and leaves physical alone; each removes 30/45/60/75 resistance; casting Curse of Agony over it restores both. Conformance golden regenerated: the four Curse of the Elements rows read `match`, 300000 ms duration (damage columns n/a, the spell deals none). `go test --tags=with_db ./sim/warlock/... ./sim/conformance/ ./sim` pass; the `.results` goldens did not move.

## Rotation search (level 60, guide FS1 talents, committed alliance BiS band, seed 7, 800 iterations, paired)

The tool gained: damage-taken debuffs as probe candidates (`RaisesDamageTaken`, aura 87 over a finite duration) gated by `not auraIsActive` on the current target, and a replace mutator (a maintenance line becomes a maintenance candidate in place), because a one-per-target pair cannot be reached by insert then remove. Refresh mutators skip target-aura gates.

Curse variants, the Bane of Agony line replaced in place, DPS (mean +- error):

| Spec | Baseline (Agony) | Elements (rank 4) | Doom (603) | No curse |
|---|---|---|---|---|
| Affliction | 421.7 +- 0.4 | 402.9 +- 0.5 | 418.2 +- 0.6 | 375.1 +- 0.4 |
| Destruction | 397.4 +- 0.5 | 390.5 +- 0.6 | 397.5 +- 0.6 | 359.9 +- 0.6 |
| Demonology | 362.4 +- 0.4 | 340.7 +- 0.5 | 358.0 +- 0.5 | 316.1 +- 0.4 |

Elements loses 7 to 22 DPS against Agony in every spec, because the solo target has no magic resistance to strip and a 10 percent bonus on one caster's damage is worth less than Agony's own damage line. Doom ties Agony in Destruction (+0.1, error 0.8) and loses elsewhere. Unstable Affliction: nothing to measure (not learnable). Nothing beat the curated rotation by more than the combined error, so no curated APL, guide prose, ladder golden or smoke expectation changed, and `apl-sync` was not needed. Full reports regenerated at `design/reviews/rotation-search/warlock-*.md` (their baselines had gone stale against the committed rotations; the stale ones read 300.8 for Demonology).

The searches also found two ordering tweaks, unrelated to this lane's candidates and each +0.2 to +0.3 percent: swap Corruption and Agony in Affliction (+0.9, error 0.6) and swap Immolate and Conflagrate in Destruction (+1.2, error 0.8). Not adopted here; a candidate for the next rotation pass.

## Open questions, settled by the coordinator's learn-row check

- Corruption: the only level-60 learn row is 25311, so the engine's rank 7 is right (1223963 has none).
- Bane of Doom: the learn row is 603. The engine registers it under 603 but with the numbers of 449432 ("Curse of Doom", a different, non-learnable client row: 3200 damage after 60 s, coefficient 1.0). The 603 client row says 1742 damage at coefficient 4.0 with the same family mask. Conformance shows 603 as `not declared` (1742.00 vs none, 4.000 vs 1.000). Choosing id 603 is right; the damage numbers under it are not the 603 row's. Not changed here: it moves Doom, the curse that ties Agony in Destruction, and wants its own measured lane.
- AQ book ranks (Shadow Bolt 25307, Immolate 25309) are book-taught and phase-gated; left alone.
