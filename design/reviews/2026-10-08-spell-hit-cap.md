# Spell hit to cap for casters (2026-10-08)

Lane `spell-hitcap`. Branches: site `spell-hitcap`, engine fork `spell-hitcap`.

## What ships

The BiS weight rail and the guide stat table say "Spell hit to cap: N% (school M%)" for a caster band, where the melee bands say "Hit to cap: N% for specials". The melee copy and shape are unchanged.

## Engine facts (verified in the fork)

- Base spell miss against a level-63 target for a level-60 caster: 17%. It lives in `sim/core/target.go`, `NewAttackTable`: `table.BaseSpellMissChance = UnitLevelFloat64(defender.Level-attacker.Level, 0.04, 0.05, 0.06, 0.17)`.
- The residual: `Spell.SpellChanceToMiss` floors a spell's miss at 1%. That literal is now the named constant `minSpellMissChance` (`spell_result.go`). Cap = 17 - 1 = 16 points, computed from the attack table, not hardcoded.
- Rating units: `HitRatingPerHitChance = 1` in the engine (the stat is already percent), and `data/builds/1.60.1.70009/gametables/combatratings.txt` row 60 has `Hit - Spell` = 10, which `data/pipeline/simdb/ratings.py` divides item hit by when it builds the simdb. So 10 rating on an item is 1 engine point, and the page's "10 hit rating is 1%" holds.
- A spell's hit is `stats.Hit + spell.BonusHitRating + school bonus + target bonus`, all via `Spell.SpellHitChance`. Talent hit is a per-spell mod (class mask), not a stat, so the profile reads each cast spell instead of modelling talents again.

## Engine change

`core.ComputeHitProfile` now also sets `Spell`, `SpellProfile` (`SpellHitProfile`: `Hit`, `SchoolHit`, `SchoolNames`, `Cap`, `ToSpellCap()`, `ToSchoolCap()`). It scans the unit's APL-castable, non-physical `ProcMaskSpellDamage` spells: `Hit` is the lowest total hit among them (what a spell with no school talent gets), `SchoolHit` the highest, `SchoolNames` the schools of the spells at that highest figure. Weapon specials, procs and pets are not counted. Tests: `sim/core/hitprofile_spell_test.go`.

## Ranker change

`hitToCapFor` returns the caster shape for `leveling.NoMeleeAutoAttackSpecs` and the unchanged melee shape otherwise:

```json
{"kind": "spell", "baseline": 0, "spell": 16, "school": {"names": ["Fire", "Frost"], "hit": 5, "to_cap": 11}}
```

`school` is absent when no talent adds hit. The melee shape has no `kind`. The report key is now a small wrapper type that decodes either shape (committed bis files still decode; tested).

## Published figures, band 60 raid (weights character, scratch run)

| Spec | spell | school | Where the school figure comes from |
| --- | --- | --- | --- |
| mage-fire | 16 | 11 (Fire, Frost; hit 5) | Elemental Precision, 5 ranks x 1% (live text: 1% a rank) |
| warlock-affliction | 11 | none (baseline 5) | Suppression, 1% a rank on every spell; see the ruling below |
| druid-balance | 16 | 12 (Arcane, Nature; hit 4) | Nature's Reach, 2 ranks x 2% (live text: 2% a rank) |
| shaman-elemental | 16 | none | no hit talent on the elemental tree (Tidal Focus is restoration) |

Sanity: the weights character is the bare ladder character (talents, buffs, one weapon, no armor), so its stat hit is 0 and the whole school figure is the talents. Every number equals the talent text times ranks, except warlock.

### Warlock before and after the Suppression fix (band 60 raid, set DPS of the ranked set)

Before is the engine at the first spell-hitcap commit (affliction-only, 2% a rank); after is the client's rows. All three builds take Suppression at 5 ranks.

| Spec | hit_to_cap before | hit_to_cap after | set DPS before | set DPS after |
| --- | --- | --- | --- | --- |
| warlock-affliction | spell 16, school 6 (Shadow, hit 10) | spell 11, baseline 5, no school | 894.5 | 903.9 (+1.1%) |
| warlock-demonology | spell 16, school 6 (Shadow, hit 10) | spell 11, baseline 5, no school | 793.0 | 826.0 (+4.2%) |
| warlock-destruction | spell 16, school 6 (Shadow, hit 10) | spell 11, baseline 5, no school | 795.5 | 827.7 (+4.0%) |

The set DPS also moves with the gear the changed hit weight picks, so it is the ranked set's figure, not the same gear before and after. Affliction gains little because it only lost the doubled bonus on its own spells while gaining Shadow Bolt and Searing Pain hit; the other two gain on every spell they cast.

## Open items for the owner

1. **Suppression (ruled: the client wins).** The client's rows for 18174 are aura 55 (spell hit) 5, aura 54 (hit) 5, both with spell class mask 0, and aura 10 (threat) -20 on misc 127 (all schools): 1% hit per rank on every spell, 4% less threat per rank on all schools. The engine had affliction spells only at 2% a rank. Fork `sim/warlock/talents.go` `applySuppression` now adds `stats.Hit` 1 per rank and multiplies `PseudoStats.ThreatMultiplier` by 1 - 0.04 per rank; tests `TestSuppressionAddsOnePercentSpellHitPerRankToEveryWarlockSpell` and `TestSuppressionReducesAllThreatFourPercentPerRank` were written first and failed against the old code. No `.results` or conformance golden moved: the committed DSRuin and SMRuin test builds do not take Suppression, and the conformance golden does not exercise it.
2. Spells flagged `SpellFlagBinary` use a different miss formula in `SpellChanceToMiss` (resists fold into the hit roll); the figure is the non-binary cap, 16, which the profile applies to every spell.
3. Debuffs that grant hit (applied as auras once a fight starts) are not in the profile, the same as for melee: it reads gear, talents and raid buffs at environment build.
4. The school label in the data is by school, not by class mask: for a warlock "Shadow" covers Corruption and Curse of Agony (Suppression) but Shadow Bolt does not get the talent. The page copy says only "(school M%)", so it does not claim more.

## Checks

`go test ./...` in `sim/` (site), `go test --tags=with_db ./sim/core/` (fork), `npx vitest run src/lib/bis`, `npm run check`, `npm run format:check` all pass.
