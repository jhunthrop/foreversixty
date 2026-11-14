# Ranged hunter hit model (2026-11-14)

Scope: the finding from the addon 2.0.49 lane. The published BiS files gave Beast Mastery and Marksmanship a spell hit cap line, and the hunter stat weights might be skewed by it.

## Method

- Traced the hit types a spec "needs" from the BiS file back to the engine: `sim/cmd/leveling-bis/hitcap.go` (`hitToCapFor`, `hitToCapFromProfile`) publishes `spell` whenever `core.HitProfile.Spell` is set; the fork's `core.ComputeHitProfile` set it from `computeSpellHitProfile`, which walked the whole `unit.Spellbook`; the same flag opens the stat-weight sweep's hit step past the weapon cap (`core.hitStepWindow`).
- Hunter spellbook against the rotations: Explosive, Immolation and Freezing Trap and Volley carry `ProcMaskSpellDamage` and `SpellFlagAPL` (`sim/hunter/explosive_trap.go`, `immolation_trap.go`, `freezing_trap.go`, `volley.go`). `sim/request/apl/hunter-beast-mastery.apl.json` and `hunter-marksmanship.apl.json` cast none of them (Serpent Sting, Arcane Shot, Multi-Shot, Aimed Shot, Hunter's Mark, Rapid Fire, Bestial Wrath, Aspect of the Hawk). Survival's list casts Immolation Trap, so its spell line is real.
- Before/after on the Windows box (`tools/remote-test.sh --engine-sha`): the full `runSpec` pipeline for both specs at bands 50 and 60 (bare and raid, both factions), written to a temp directory and thrown away (nothing committed). A second probe read the weights character's hit profile for every other weapon-using and caster spec at 40 and 60.

## Findings

| # | finding | source | status |
|---|---|---|---|
| H1 | The hit profile counted every damaging spell the character knows, not the ones its rotation casts. BM and MM published `hit_to_cap.spell` (16% less their hit) for traps and Volley they never cast; the weights sweep would also treat hit past the ranged cap as still live for them. | fork `sim/core/hitprofile_spell.go` `computeSpellHitProfile` (whole spellbook) | Fixed, fork f9385b498: `rotationCastSpells` reads the APL's cast, channel, multidot and multishield actions at any depth, the prepull and the cooldowns an autocast action fires; a custom rotation or a spec's own action keeps the whole spellbook. Fork tests first (`hitprofile_rotation_test.go`). |
| H2 | Shadow priest's figure moves shape, not value: the spellbook's Holy spells (never cast) held the baseline at gear-only hit with Shadow Focus as a school figure. Every spell the rotation casts is Shadow, so the baseline now includes Shadow Focus and there is no school figure (raid 60: was "12%, Shadow 7%", now "7%"). | same | Fixed with H1; the guide caption counts a school talent in the baseline when no school figure is published (`web/src/lib/guides/hit-cap-line.ts`, test added). |
| O1 | Marksmanship's hit weight is zero at raid 60 while the published set sits at 7.7 to 8% hit: the geared weights sweep wears the first-pass set at 14.3% hit, past the 9% ranged cap. Same before and after this fix. | `geared_weights.go` (`gearedBandWeights`) | Open: a ranker question (weigh hit on a set nearer the cap, or keep hit's below-cap value), not a hit-model bug. |

## Before and after

Identical seeds, same commit, engine f0f90b5f9 against f9385b498:

- BM and MM bands 50 and 60 bare and raid 60, both factions: every weight, every pick and every `set_dps` identical (0.0% change). The weights characters sat at 0 to 7% hit (below the ranged cap, where the spell flag never touched the step) or at 13.3% (BM raid) and 14.3% (MM raid) where hit was already worth nothing either way.
- The only published change: `hit_to_cap.spell` is gone from every BM and MM band.
- Every other spec probed (Survival, the cat, the bear, Retribution, Protection paladin, Enhancement, the rogues, the warriors, Balance, Elemental, the mages, the warlocks): hit profile identical. Shadow priest: H2.
- DPS: the change touches only the hit profile, never a fight; no DPS moved and no golden moved.

## Shas

- Fork: f9385b498 (`lane/hunter-hit`, fast-forwarded onto master).
- Site pin: f9385b498.

## Refresh

MR113: the next bis run drops the spell line from the hunters and re-shapes Shadow's. No re-rank is needed for picks or weights.
