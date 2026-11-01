# Sim-accuracy round 4: integration (2026-11-01)

Round 3's open items that the integrator owned, the integration, the pin and the review. The lane reports beside this file have the per-spec rows.

## Integrator findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **The /sim page fought a Humanoid by default.** Its default encounter had an empty creature type, which `sim/request` runs as a Humanoid (kept for requests stored before the field). The tier list, BiS, spec-breakdown and the searches fight `leveling.RankerEncounter()`, a target with no creature type. The control labelled the empty type "Any". Rogues read about 4% above their published number (Murder). The default is now "No creature type", and a stored empty type shows the Humanoid it actually fights. | `web/src/lib/sim/types.ts` DEFAULT_ENCOUNTER; `sim/request/request.go` `encounter`; settings.test.ts failed first (target type `''`) | CONFIRMED-BUG (page) | S c3cdc8d36 |
| I2 | **The /sim Raid-buffed preset was not the raid preset.** The list was a September copy of an adapter fixture: Blessing of Kings, Sanctity Aura (not learnable in Forever), Winter's Chill (no longer a crit debuff since 70532) and Shadow Weaving; no Heart of the Lion, Windfury Totem or Hunter's Mark. The buff and debuff half now comes from `data/curated/presets.json` at build time (`web/scripts/sync-raid-preset.mjs`), the preset every published level-60 number uses. This also makes the Kings with Heart of the Lion question moot for default page runs. | round 3 hybrids C3; settings.test.ts reads the curated file directly | CONFIRMED-BUG (page) | S c3cdc8d36 |
| I3 | **The coach's sim reference fought a Humanoid.** `simRefRequest` used `DefaultEncounter()` with no type. The reference is the band's BiS gear, preset and race, the published number's setup, so it now fights the type-neutral target. References are keyed by engine version, so the pin recomputes them. | `api/internal/coach/simref_job.go`; test failed first | CONFIRMED-BUG (tooling) | S c4d60af28 |
| I4 | **A rage potion on a class with no rage bar panicked the engine.** The /sim page filters it (`usableConsumables`), but the API and the tools pass any consumable, and a hunter or rogue run with Mighty Rage Potion dereferenced nil in `AddRage`. Found while measuring the page's consumables. | fork `sim/hunter/rage_potion_test.go` panicked first | CONFIRMED-BUG | E b41baf0a3 |
| I5 | **The /sim armor field named the level-63 boss's 3731 for every character.** The sheet resolved the preset without the character's level, and below 60 the web table used a linear fall floored at zero while the engine scales the level-60 figure by level (`api.armorBelowMaxLevel`). A level-33 character fights a level-36 target with 1980 armor. Found by the wow-player review. | settings.test.ts failed first (level 36: 0, want 1980) | CONFIRMED-BUG (page) | S f19f0a6e3 |
| I6 | **The /sim page's consumables are still its own list.** It carries one physical and one caster list, while the published numbers use the curated per-spec groups plus the class kit (poisons and self-buffs, `sim/leveling/kit.go`). Measured against the published setup (3000 iterations, horde): Fury +0.6%, Fire +1.8%, Destruction +0.5%, Combat -7.6% (no poisons), Marksmanship -1.8%, Survival +4.1%. Aligning it needs the page to resolve consumables per spec, as `request.Presets.Resolve` and the kit do. That is a design change, not a one-line fix. | `-consumes` runs against the published setup | OPEN (page design) | - |
| I7 | **Weapon damage source (round 3 open).** Settled by the physical lane (X1): simdb.bin's curve is the closer one. The item file was Era's `db.json` copied over every weapon, and the curve missed QualityModifier and the caster-weapon bit. Both now follow the client (1,027 of 1,036 melee weapons match Wowhead Forever). | physical.md X1 | CONFIRMED-BUG (data) | S 5043869bd |
| I8 | **The BiS ranker's weapon-DPS fix (S 425147a6f), checked against the bis.yml run that followed it** (41324083044 on f2c4446ea, published as 844619b6a at pin 640da215f). The band-60 raid weapons:<br>- Feral cat: the caster staff gives way to Hammer of the Titans (12796), both factions. Sane.<br>- Retribution horde: Blade of Hanna gives way to The Needler. That reflects the double-counted Seal of Righteousness (hybrids P1), so it re-picks at the next run.<br>- Survival's off hand: Heartseeker gives way to Teebu's Blazing Longsword. That pick is illegal (I10).<br>- Every other melee and hunter weapon held. | `git show 844619b6a:data/builds/1.60.1.70707/bis/<spec>.json` against 461b82caa | Aligned, except I10 | - |
| I10 | **The ranker put a main-hand-only weapon in the off hand.** The item file's slot reads "main_hand" for one-handers and "Main Hand" weapons alike. Once weapons were priced off the engine row, Teebu's Blazing Longsword (1728, Main Hand per Wowhead Classic and the engine's HandType) won Survival's band-60 off hand. The engine swings no off hand for it: 650.8 DPS against 742.6 with Heartseeker. The published set fell from 703.1 to 653.7 (horde) while the engine gained 5% for Improved Tracking. The ranker now carries the engine's hand type with its weapon row and keeps main-hand-only weapons out of the off-hand pool. A local re-rank gives Teebu's in the main hand and Heartseeker in the off hand: 768.6 horde, 757.1 alliance (+17.6% over the published set). The only other published case is Beast Mastery band 50 bare horde (Force of the Hippogryph). Found by the integrator's check of the bis.yml run. | `cmd/leveling-bis/engine_weapons_test.go` failed first (no hand type); scan of every published set | CONFIRMED-BUG (ranker) | S b11c281b7 |
| I9 | **Pet crit, hit and haste.** The warlock lane settled demon hit on a Blizzard developer statement (forever-bugs #394, single-source) and the hunter lane hunter pet hit on the client's channels. Crit and haste are blocked on level-60 logs: Blizzard's stated demon crit inheritance (forever-bugs #331) is not live. | warlock.md X2-X4; physical.md B1 | partly fixed; crit and haste BLOCKED on level-60 logs | lanes |

## Integration

- Fork `int/r4` merges `lane/r4-int`, `lane/r4-physical`, `lane/r4-hybrids`, `lane/r4-warlock` and `lane/r4-casters`. The only conflict was `sim/mage/TestP1Mage.results` (Eureka! and the mage cost fixes both moved it), regenerated at the merge. The full fork suite passes with `-tags=with_db`. The exception is `sim/web`, which already fails setup without the `binary_dist` package.
- The site merges `wt/r4-*`. The only conflict was `data/pipeline/normalize/gear.py`, where the physical lane mapped stat 83 (weapon damage) and the casters lane stats 84 to 89 (school spell damage); both mappings are kept.
- The pin is f8a5f94ce. The rotation ladder goldens are regenerated at the pin with `make simdb` and a scratch go.work.
- Gates, all exit 0:
  - `go vet` and `go test ./...` in sim
  - the coach and sims packages in api
  - `make apl-check` against the pin
  - `make styles-check`
  - the data tests for presets, APLs, genproto and stats
  - the full web vitest
  - astro check, eslint
  - the 33 sim e2e specs
- Five data tests already fail on origin/main before this round: `test_addondata` and `test_audit_addon`, Data.lua drift.

## Before/after

Pin 640da215f with origin/main, against pin f8a5f94ce with this branch. On today's published data: the BiS sets, builds and the simdb.bin committed today. Raid preset, band 60, 10,000 iterations. Healers are measured in effective HPS on the heal profile (now ×1.65, hybrids), tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 519.2 to 519.2 | +0.0 / +0.0 | 518.0 to 512.8 | -1.0 / -1.0 |
| druid-feral-bear | 674.2 to 667.8 | -0.9 / -1.0 | 694.0 to 686.4 | -1.1 / -1.0 |
| druid-feral | 582.8 to 582.8 | +0.0 / +0.0 | 585.8 to 582.3 | -0.6 / -0.6 |
| druid-restoration | 712.0 to 732.2 | +2.8 / +2.7 | 711.0 to 728.1 | +2.4 / +2.3 |
| hunter-beast-mastery | 640.5 to 644.8 | +0.7 / +0.8 | 656.9 to 661.1 | +0.6 / +0.6 |
| hunter-marksmanship | 646.2 to 649.2 | +0.5 / +0.5 | 660.9 to 665.7 | +0.7 / +0.8 |
| hunter-survival | 730.0 to 730.0 | +0.0 / +0.0 | 743.0 to 743.0 | +0.0 / +0.0 |
| mage-arcane | 663.6 to 601.8 | -9.3 / -9.4 | 665.7 to 605.5 | -9.0 / -9.2 |
| mage-fire | 714.7 to 625.8 | -12.4 / -12.5 | 710.9 to 618.9 | -12.9 / -13.0 |
| mage-frost | 559.3 to 547.2 | -2.2 / -2.0 | 557.3 to 543.3 | -2.5 / -2.4 |
| paladin-holy | 538.5 to 549.5 | +2.0 / +2.1 | 531.3 to 541.9 | +2.0 / +1.9 |
| paladin-protection | 693.5 to 693.5 | +0.0 / +0.0 | 721.2 to 689.1 | -4.5 / -4.4 |
| paladin-retribution | 589.5 to 632.7 | +7.3 / +7.0 | 649.6 to 583.3 | -10.2 / -10.1 |
| priest-discipline | 522.7 to 526.5 | +0.7 / +0.8 | 540.6 to 543.1 | +0.5 / +0.5 |
| priest-holy | 651.1 to 637.3 | -2.1 / -2.2 | 628.6 to 619.5 | -1.4 / -1.4 |
| priest-shadow | 668.1 to 665.4 | -0.4 / -0.4 | 669.2 to 668.4 | -0.1 / -0.2 |
| rogue-assassination | 601.0 to 601.0 | +0.0 / +0.0 | 614.5 to 614.5 | +0.0 / +0.0 |
| rogue-combat | 608.1 to 608.1 | +0.0 / +0.0 | 622.0 to 622.0 | +0.0 / +0.0 |
| rogue-subtlety | 545.3 to 545.3 | +0.0 / +0.0 | 559.7 to 559.7 | +0.0 / +0.0 |
| shaman-elemental | 464.7 to 464.7 | +0.0 / +0.0 | 462.6 to 466.3 | +0.8 / +0.8 |
| shaman-enhancement | 618.4 to 618.4 | +0.0 / +0.0 | 634.1 to 639.8 | +0.9 / +0.9 |
| shaman-restoration | 611.9 to 640.9 | +4.7 / +4.7 | 610.7 to 639.6 | +4.7 / +4.6 |
| warlock-affliction | 605.2 to 583.4 | -3.6 / -3.7 | 601.0 to 582.4 | -3.1 / -3.1 |
| warlock-demonology | 661.1 to 651.7 | -1.4 / -1.4 | 664.1 to 659.2 | -0.7 / -0.8 |
| warlock-destruction | 678.2 to 666.4 | -1.7 / -1.6 | 676.3 to 668.0 | -1.2 / -1.2 |
| warrior-arms | 726.5 to 726.5 | +0.0 / +0.0 | 745.2 to 751.2 | +0.8 / +0.8 |
| warrior-fury | 764.0 to 764.0 | +0.0 / +0.0 | 780.7 to 780.7 | +0.0 / +0.0 |
| warrior-protection | 1038.1 to 1038.1 | +0.0 / +0.0 | 1038.5 to 1026.4 | -1.2 / -1.3 |

What this table does not yet hold, because each needs a nightly regeneration:
- **The item-data fixes.** These are weapon curves, Molten Heart's weapon damage, ranged attack power, school spell damage, spell penetration, and the flat crit, hit and avoidance percentages. They reach published numbers once data.yml regenerates simdb.bin and bis.yml re-ranks. The lanes measured them locally:
  - On today's sets: Fire +1.3 to +1.5% (Kindling Stave's crit), warlocks +0.7 to +0.8%, rogues +0.5 to +0.7%, hunters +0.6%, Balance +0.4%.
  - After a re-rank: Shadow +1.5%, Elemental +1.2 to +1.9%, rogues up to +2.9%.
- **The new Fire raid build.** It adds +5.2% on top.
- **Horde Retribution's weapon re-pick.** On Arcanite Champion the horde number is +10.1% instead.

## wow-player review

First pass: SHIP Fire, Frost, Protection paladin, Restoration druid, Shadow, Balance, both ferals, the warlocks, the hunters, the rogues, the warriors, Enhancement and Elemental. The reviewer held four, and each was settled before the push:
- **Arcane, "never runs out of mana while Fire does".** A mana trace shows 3.1 s out of mana in 180. The list paces itself, Frostbolt when mana is behind, and Evocation costs 1.0 to 2.4%. Every cost matches the client. The -9% splits into rank 7 Missiles (rank 8 is an Ahn'Qiraj book) at -4.5%, one-charge Clearcasting at -5.3% and the cost fixes at -0.3%. SHIP.
- **Retribution, "twist on bare, faction gap, Unsealed Wrath's reading".** The Seal of Command lines are now gated on 30% mana and 1200 attack power (site 6660e6126), and bare is the best measured. The gap is the weapon: horde's Blade of Hanna against Arcanite Champion. Seal of Command without the twist is +7.2% under either reading. SHIP WITH FIXES: the horde number waits for bis.yml to re-pick the weapon.
- **Healers, "×1.65 fits a target band".** The lane justified it on gear separation: a three-slot downgrade costs the druid 53% more HPS and the shaman 25%, but the Holy priest 30% less. SHIP WITH FIXES: /tiers/healer ranks healers side by side, and the retune reorders Holy priest and Restoration shaman (0.8% apart). The healer tier page is held for the owner, see the ledger.
- **/sim, "target level 36 beside armor 3731".** The 36 was right (the fixture warrior is level 33). The 3731 was I5, now fixed. SHIP.
