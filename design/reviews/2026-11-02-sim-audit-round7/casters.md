# Sim audit round 7: casters (Mage and Priest), 2026-11-01

Specs: mage-arcane, mage-fire, mage-frost and priest-shadow, plus the damage side of priest-discipline and priest-holy. The Healers lane owns healing.

## Method

- **Engine.** Lane branch `lane/r7-casters` from the pin `b57765fb9`. The before column runs a second worktree detached at the pin.
- **Site.** Branch `wt/r7-casters` from origin/main `9cfb91c47`, with `make simdb` (the committed simdb.bin from data.yml's `c0bd0bd3d`).
- **Client.** The 70707 tables with the hotfix push layered in (the pipeline lane's layered `raw/`, fetched 2026-11-01 21:42 UTC).
- **Masks.** Every mage and priest talent effect was dumped with its aura, misc value, amount and class mask, and each mask was resolved against the spells' own SpellClassOptions rows. The scripts are in the lane scratch dir (`tools/talents.py`, `tools/maskhit.py`).
- **Runs.** `spec-breakdown` at 10,000 iterations, seeds 1 and 50001, both factions. Bands: raid 60, and bare 60, 50, 40 and 20.
- **Tests.** Every confirmed bug got a failing test first. I saw it fail, then saw it pass.
- **Not regenerated.** The item-data fix reaches published numbers at the next data.yml run. I did not commit a regenerated simdb. The "after, with new simdb" column comes from a scratch build of the simdb from the same layered tables.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1a | Ignite rolls crits into one burn (round 6 F1): its side effects. | `sim/mage/ignite.go`. A crit adds 40% of its finished damage to what the burn has left, over a fresh 4 s. The ticks keep a multiplier of 1 and ignore target modifiers (beta note "no longer counts damage bonuses twice"). The ticks crit (owner ruling). Master of Elements (`talents.go:391`) and Combustion (`talents.go:441`, `:473`) read only `OnSpellHitDealt` for mage spells, so Ignite's ticks neither refund mana nor spend Combustion charges. That matches Combustion's "non-periodic critical strikes". Ignite is a Fire effect for Fire Power (`ignite.go:73`). Raid Fire: Ignite is 14.3% of damage at 45 ticks a fight. | Aligned | - |
| I1b | Fire Power's bonus never reaches the Fireball and Pyroblast burns. | Fire Power 11124 is aura 4 (a server script, "per your Fire effects on the target"). The engine adds it only while the cast's `ApplyEffects` runs (`fire_power.go:50`). The burns are applied after travel, so they never had it, before round 6's dynamic DoTs or after. Now that the burns re-read the caster every tick, this choice decides whether Forever's dynamic ticks count the target's Fire effects. The client cannot say. | OPEN. Settled by a level-60 log of Pyroblast ticks with and without Fire effects on the target. Worth at most about 1% of Fire. | - |
| I1c | Improved Scorch is self-only and out of the raid preset. | Raid Fire still keeps its own stacks: Scorch is 33% of damage, and the APL's upkeep line stays (round 6 F5, held under the adoption rule). | Aligned (hold stands) | - |
| I1d | Ice Lance's Frozen x4 on spell power, and Arcane Blast on cast. | `ice_lance.go:104` multiplies the whole hit. 1342606 is now 300 in the layered table, which matches the 70707 text. Raid Frost: Ice Lance averages 3,504 at 72% crit (Shatter +50%). Arcane Blast has no missile speed. | Aligned | - |
| I1e | Robe of the Archmage. | One use a fight in all three raid mage sets (`item 14152`, 1.00 a fight, 499.5 mana). The gems and potions keep their cooldowns. | Aligned | - |
| I1f | Dynamic DoTs: Shadow Word: Pain, Devouring Plague, Mind Flay, and the Fireball, Pyroblast and Flamestrike burns. | `core/dot.go:267` re-takes `Snapshot` each tick, which also re-reads crit. Mind Flay's Improved Mind Flay ramp multiplies the per-tick snapshot and restores it (`mind_flay.go:91`), so the dynamic read keeps the ramp. Raid Shadow: the DoTs crit 20%, the same chance as Mind Blast before Inner Focus. | Aligned | - |
| I1g | Mental Agility and Spiritual Healing masks. | Re-checked against SpellClassOptions: Mental Agility covers Shadow Word: Pain, Devouring Plague and Shadow Word: Death, and not Mind Flay or Mind Blast. Raid Shadow costs: Mind Flay 102.5 (205 x 0.5), Mind Blast 175. But the cost check exposed how Shadowform stacks with Mental Agility; see P1. | Aligned (masks); P1 is a new bug | - |
| I2 | **Three wands in simdb were still Physical:** Pyric Caduceus 11748, Noxious Shooter 17745 and Peacekeeper Wand 274655. Pyric Caduceus is the band-50 wand of every mage, Shadow priest and warlock set. | Round 6 gave client and classic-db wands their school. Wands that come only from the wowhead supplement never called `_apply_wand_school` (`data/pipeline/simdb/items.py`, `build_wowhead_sim_items`). Wowhead Forever states `dmgtype1` for them: Pyric Caduceus 2 (Fire), Noxious Shooter 3 (Nature), Peacekeeper Wand 5 (Shadow). classic-db's item_template says Fire for 11748 too. A scratch simdb built from the same tables differs from the committed one in exactly these three items. Test `test_a_wowhead_wand_carries_the_school_its_dmgtype1_states` failed 3 of 4, then passed. | CONFIRMED-BUG (data; two sources for 11748) | S 953c25aa9 |
| I2b | Wand multipliers. | Wand Specialization (6057/14524) is aura 79 on wands, which the engine applies as Shoot's multiplier. Fire Power, Critical Mass, Combustion and Improved Scorch's crit are class-mask or mage-spell effects, and none touches Shoot (`talents.go`, `fire_power.go:47`). Shadowform and Darkness are aura 79 on the Shadow school, so they do raise a Shadow wand, as in Classic. Arcane Instability's crit did not reach the wand; see A1. | Aligned except A1 | - |
| I2c | Do wands use the spell hit table? | The engine rolls Shoot on the spell table (`core/wand.go`), so Precision (aura 55) and Shadow Focus (class mask) do not reach it. The emulators roll wands (DmgClass ranged) on the ranged table. Either way neither talent applies. Bare wand miss is 15 to 17% against a target 3 levels up. | OPEN (table choice; no talent moves) | - |
| I3 | Ledger rows: Ice Lance's coefficient, Ignite's 0.73, Missile Barrage on a miss, Arcane Concentration per missile, Master of Elements on Clearcasting, Shadowfiend. | No level-60 Forever logs yet (launch is 4 November). No new evidence. | OPEN (blocked on logs) | - |
| I4 | Mage and Shadow priest raid sets after bis run 41363895819. | The run was still in progress at the end of this lane (last update 23:48 UTC). The published sets are still the `8baffe5fc` ranking, which has Torch of Light in all four raid sets. | Not done: the integrator re-checks at the merge | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **Arcane Instability's crit reached only mage spells.** Wand shots and item procs got none of its 3%. | 15058 has three effects. Effect 0 is aura 108 misc 0 (direct damage) and effect 1 is aura 108 misc 22 (periodic), both class-masked: these are the damage half. Effect 2 is aura 290 (all crit chance) with no class mask. The engine had the crit as a class-spell modifier. It is now a character crit stat. Test `TestArcaneInstabilityCritReachesEveryCrit` failed (Shoot +0.0000), then passed. No golden moved. Only Arcane's build takes the talent. | CONFIRMED-BUG (fresh, client) | E 75aa6f778 |
| A2 | Arcane Power and Arcane Instability split direct damage (misc 0) from periodic damage (misc 22). The periodic mask names only Fireball, Flamestrike, Pyroblast and Blizzard, so Frostfire Bolt's burn should get neither. The engine's `DamageMultiplierAdditive` covers both halves. | No published rotation casts Frostfire Bolt, and Arcane's Missiles bolts are direct (in the misc 0 mask). | Noted (no published number moves) | - |
| A3 | Fresh, checked and aligned: Arcane Blast 1239700 and stack 400573, Missile Barrage 400588, Arcane Concentration 11213, Arcane Impact, Arcane Mind, Arcane Power 12042 (+30% damage, +30% cost, 3 min), Presence of Mind (category 180 s), Evocation 12051 (aura 110 +1500%, aura 134 100%, 8 min), Orc Blood Fury 20572 (10%, 2 min). | Mask checks against SpellClassOptions. Raid breakdown: Arcane Missiles 45%, Frostbolt 30%, Arcane Blast 25%. | Aligned | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Fire Power on the burns (I1b). | | OPEN | - |
| F2 | Fire wand and Ignite: a Fire wand crit does not feed Ignite (`ignite.go:40`, spell damage only). | Ignite 11119 is aura 4, a server script. Its text says "Fire damage spells". | OPEN (bare bands only; small) | - |
| F3 | Elixir of Firepower (+10) rather than Greater Firepower (+40) in the Fire set. | In 70707, item 21546 is "Elixir of Holy Power" (spell 1310077, +40 Holy). Forever turned Greater Firepower into a Holy elixir, so `presets.json`'s note "the client table has none" is right. | Aligned (Forever change) | - |
| F4 | Fresh, checked and aligned: Combustion 11129/28682 (aura 107 +10% a stack, 3 charges on 65536, 3 min category), Critical Mass, Incineration, Heating Up, Master of Elements (Fire and Frost crits, mage spells only), Improved Fire Blast, Scorch, Fire Blast and Pyroblast hit (Precision 3/3 plus gear: 8.5% miss). | | Aligned | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Fresh, checked and aligned: Ice Barrier 13033 and Frigid Power 1326891 (15%, plus 5% while shielded; the misc 22 rows cover Blizzard and Frostfire Bolt's periodic damage), Fingers of Frost (2 charges, 15%), Shatter (+50% at 3/3), Ice Shards, Frost Channeling, Winter's Chill, Cold Snap, Troll Berserking 20554 (10%, 15 s, 3 min). | Raid breakdown: Frostbolt 68%, Arcane Missiles 16%, Ice Lance 16%. | Aligned | - |
| R2 | Ice Lance's coefficient (ledger). | | OPEN (logs) | - |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Shadowform's 50% cost cut added to Mental Agility's 10% instead of multiplying with it.** A top-rank Shadow Word: Pain cost 188 (470 x 0.4) instead of 211.5 (470 x 0.9 x 0.5). Devouring Plague cost 394 instead of 443. | Shadowform 15473 effect 3 is aura 72 (MOD_POWER_COST_SCHOOL_PCT, Shadow at -50), a school cost multiplier. Mental Agility 14520 is aura 108 misc 14, a spell modifier. The server applies a spell's modifiers first and the school multiplier after them. That ordering is from cmangos and TrinityCore's power-cost code, the only source. The client's two aura types support it. The engine already has `SchoolCostMultiplier`, applied after the spell modifiers (`core/spell.go:715`), and Shadowform now uses it. Test `TestShadowformCostMultipliesWithMentalAgility` failed (188 against 211.5), then passed. Raid Shadow never runs dry, so only the bare bands move. | CONFIRMED-BUG (fresh; client aura types plus the emulators' cost order, one family) | E 475f421a4 |
| P2 | **Silent Resolve cut only Holy threat.** | 14523 effect 0 is aura 10 on school mask 34, Holy plus Shadow. The 70707 text reads "the threat generated by your spells". Test `TestSilentResolveCoversShadowThreat` failed (1.0), then passed (0.70). | CONFIRMED-BUG (fresh, client; threat only, unpublished) | E 05f285314 |
| P3 | Fresh, checked and aligned: Twin Disciplines 1225132 (direct mask: Shadow Word: Death, Penance, Holy Nova; periodic mask: Shadow Word: Pain, Devouring Plague, Mind Flay, Renew), Darkness 15259 (aura 79 Shadow, 10%), Shadowform's crit-damage mask (Mind Blast, Mind Flay, Shadow Word: Pain, Devouring Plague, Shadow Word: Death), Shadow Focus (class mask, 5%), Improved Mind Blast, Improved Shadow Word: Pain, Inner Focus (crit mask excludes periodic), Devouring Plague 985 mana with a 24 s category, Mind Flay 205 mana. | Raid breakdown: Mind Flay 41%, Shadow Word: Pain 19%, Mind Blast 19%, Devouring Plague 12%, Shadow Word: Death 8%. Mind Blast and Shadow Word: Death miss 10.4% (Shadow Focus 5% plus gear). | Aligned | - |

## Priest: Discipline and Holy (damage side)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Searing Light was a 5% additive share of Smite and Holy Fire.** | 14909 effect 0 is aura 79 (damage done) on school mask 2, Holy: every Holy damage the priest deals, Penance and Holy Nova's damage halves included, multiplied like Darkness. It now multiplies the Holy school. Test `TestSearingLightIsAllHolyDamageDone` failed, then passed. No published rotation casts Holy damage. | CONFIRMED-BUG (fresh, client; /sim custom rotations only) | E 05f285314 |
| H2 | Power in Light (Smite and Penance +10% on targets with your Holy Fire or Shadow Word: Pain) is named, not modelled (`talents.go:43`). | No published rotation casts Smite or Penance on an enemy. | Noted (/sim gap) | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | The Pyric Caduceus fix also lifts the band-50 bare warlock sets (Affliction, Demonology and Destruction all wear it). | `bis/warlock-*.json`, band 50 bare, both factions | Outside the group: re-measure | S 953c25aa9 |
| X2 | Engine percent cost modifiers (`SpellMod_PowerCost_Pct`) are summed. For example, Arcane Power's +30% and Frost Channeling's -15% give 115%, where the modern client may multiply its percent spell modifiers. | Not settled by the client. It moves nothing published: no raid mage runs dry, and Arcane does not take Frost Channeling. | OPEN | - |
| X3 | `go env GOWORK` inside the lane engine worktree finds the lane's own go.work one directory up, so `go test ./sim/...` there fails setup. Use `GOWORK=off` for the fork suite. | tooling | Noted | - |

## Before/after

Before: pin `b57765fb9` on origin/main `9cfb91c47`. After: lane engine `475f421a4` on site `953c25aa9`, with the committed simdb. "After, new simdb" adds a scratch simdb built from the same layered tables (what data.yml will produce). Horde DPS, seed 1. Changes are seed 1 / seed 50001. Alliance is within 0.3 points of horde on every row.

| Spec, band | before | after | change | after, new simdb | change | causes |
|---|---|---|---|---|---|---|
| Arcane raid 60 | 589.6 | 589.6 | 0.0 / 0.0% | 589.6 | 0.0% | A1 (no wand shots) |
| Arcane bare 60 | 318.2 | 318.4 | +0.1 / +0.1% | 318.4 | +0.1% | A1 |
| Arcane bare 50 | 194.6 | 194.6 | 0.0 / 0.0% | 196.6 | +1.0 / +1.1% | I2 |
| Fire raid 60 | 656.7 | 656.7 | 0.0% | 656.7 | 0.0% | - |
| Fire bare 50 | 181.8 | 181.8 | 0.0% | 190.1 | +4.6 / +4.6% | I2 |
| Frost raid 60 | 590.1 | 590.1 | 0.0% | 590.1 | 0.0% | - |
| Frost bare 50 | 193.0 | 193.0 | 0.0% | 198.8 | +3.0 / +3.3% | I2 |
| Shadow raid 60 | 713.9 | 714.3 | +0.1 / +0.0% | 714.3 | +0.1% | P1 (never short of mana) |
| Shadow bare 60 | 356.0 | 349.0 | -2.0 / -2.0% | 349.0 | -2.0% | P1 |
| Shadow bare 50 | 189.8 | 189.8 | 0.0 / 0.0% | 196.5 | +3.5 / +3.6% | I2, P1 |
| Shadow bare 40 | 128.3 | 128.3 | 0.0% | 128.3 | 0.0% | - |

Alliance Shadow bare 60 is 360.0 to 351.6 (-2.3 / -2.3%), and alliance Shadow bare 50 is 195.9 to 201.3 (+2.8%). Every spec at bare 40 and bare 20 is within 0.2%. The full grid (both factions, both seeds, five bands) is `meas1.txt` and `meas-shadow2.txt` in the lane scratch dir.

## Still open

- Fire Power on the Fireball and Pyroblast burns under dynamic DoTs (I1b, F1).
- A Fire wand crit and Ignite (F2).
- Which table a wand rolls on (I2c).
- Summed against multiplied percent cost modifiers (X2).
- The ledger's log-blocked rows (I3): Ice Lance's coefficient, Ignite's 0.73, Missile Barrage on a miss, Arcane Concentration per missile, Master of Elements on Clearcasting, Shadowfiend.
- The raid set re-check after bis run 41363895819 lands (I4).
- Power in Light is unmodelled (H2).

## Sources

- The 70707 client tables with the wago.tools hotfix push layered in (pipeline lane, fetched 2026-11-01 21:42 UTC): SpellEffect, SpellClassOptions, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, Spell, SpellName, ItemSparse, Item.
- Wowhead Forever item feed as carried in `data/builds/1.60.1.70707/raw/wowhead/items.json` (`dmgtype1`).
- classic-db item_template (`raw/classicdb/item_template.json`), Pyric Caduceus `damage_type` 2.
- cmangos and TrinityCore `CalculatePowerCost`: the school cost multiplier is applied after the spell modifiers.
