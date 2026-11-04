# Sim-accuracy round 18: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70755 with the late hotfix re-layer. Site branch wt/r18-shaman (off int/r18 515fdff27), fork branch lane/r18-shaman (off int/r18 8c24d6202). Gear is the frozen BiS ($R/bis-frozen); talents are the guide builds unless a row says "band talents" (the published band entry).

**1 confirmed bug, fixed.**
- Enhancement: Earth Shock held back below half mana. Origin old. Guide text and card note ship with it.
- Elemental: 0. Restoration: 0.

Client note for later lanes: wago's `SpellCooldowns.csv?build=1.60.1.70755` is the base table without the hotfix re-layer (Stormstrike reads 8000 ms there, 6000 ms in `spellconst/shaman.json`). Read spellconst first.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Regression: tank Healing Wave rows at 38 to 46.** Every row rank wins or ties its neighbours at every level 36 to 48 (37: rank 4 106.2 vs rank 3 105.9; 41: rank 4 161.5 vs rank 5 160.6; 42: rank 5 159.9 vs rank 4 159.7; 46: rank 5 163.9 vs rank 6 162.2; 47: rank 6 163.4 vs 162.7 / 162.8). The rows in rank-choices.json are the round-17 rows and the sim copy is in sync. | $R/shaman/resto-probe-3k.txt (horde, bare, 3000 iterations, effective HPS) | Aligned | - | - |
| S2 | **Regression: Enhancement hit weight past the weapon cap.** `hitStepWindow` still measures up to the spell cap for a physical character with damage spells (sim/core/hitprofile.go:79); its tests pass on int/r18. The window's low side reaches one point below the weapon cap when hit sits within a point of it, the same "value of the last point you hold" choice the plain melee path makes at the cap. | TestHitStepWindow* (sim/core) | Aligned | - | - |
| S3 | **Enhancement's empty relic slot.** Measured each modelled totem on the frozen raid 60 set (10,000 iterations, horde; paired seeds):<br>- Totem of Thunder (crafted, Enchanting): +0.23 / +0.25 DPS at seeds 1 / 50001, the same on four more seeds. Real and tiny (0.03%).<br>- Totem of the Storm (272432, world drop 56 to 62): -0.04, +0.61, +1.76, +1.34, +0.77, +0.54 over six seeds, mean about +0.8 DPS (0.1%). Noisier because its Maelstrom roll shifts the random stream.<br>- Totem of Urgency and Burning Totem: 0.00 (Enhancement casts neither Lesser Healing Wave nor Flame Shock).<br>**How it should count:** every Enhancement shaman wears a totem, and a relic has no stats to lose, so an empty slot is the wrong answer. The ranker drops it because the relic gate (`trinketGainSignificant`, 2 x the combined unpaired standard error of two 100-iteration runs, about 1.4 DPS) can never resolve a sub-1-DPS gain. Two fixes, either enough: (a) gate a relic on the paired difference (same seed, per-iteration delta) instead of two independent errors; (b) for a relic slot only, publish the best measured modelled relic whose effect reaches a spell the rotation casts, marked as a near tie, instead of emptying the slot. Today's published alternative is Totem of Urgency, a healing relic, which is the worst suggestion for this spec. The right pick is Totem of the Storm where you have one, Totem of Thunder otherwise. Ranker code is shared; not changed here (no ranker runs this round). | sim/cmd/leveling-bis/report.go:1884-1891, weights.go:204-214, trinkets.go:58; $R/bis-frozen/shaman-enhancement.json band 60 rows | Ranker policy (For other lanes, Measure at refresh) | - | - |
| S4 | Tidal Totem (272431, Healing Wave cost -5%) is the one modelable shaman relic not registered. Its only source is the vendor Pix Xizzix, `opens: later` in loot.json, and the fork's item database lacks it (999f1085b said the same). Not a launch item. | loot.json vendor:227853; sim/shaman/relic_mods_auto_gen.go | Aligned (later phase) | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Earth Shock stopped at half mana; a quarter keeps more damage.**<br>What was wrong: line 7 cast Earth Shock only above 50% mana, "so you do not run dry without raid buffs". In a raid the shaman never runs dry (out of mana 0 s of 180 with or without the gate: Judgement of Wisdom and Mana Spring), so the reserve only dropped shocks. Without raid buffs the half reserve also cost damage at most levels.<br>Swept 0, 10, 20, 25, 30, 35, 40, 45, 50, 65% and pacing (mana above time left) and absolute mana floors (600 to 1200) at bands 20 to 60, then 10,000 iterations at seeds 1 and 50001 for the finalists. 25% wins at every band but 40.<br>Band 40 (levels 40 to 49) still prefers 50% by 0.3 to 0.85%: Maelstrom Weapon is 3/5 there, so the five-stack bolt still costs 40% of its mana. No single gate, absolute floor or Stormstrike-cooldown guard keeps that band level, and the guide gives one rule.<br>Fix: the gate is 25%. Card note 7, loop item 6 and the Mana paragraph follow (enhancement.md:77, :87). | Before/after table below; $R/shaman/es-gate-10k.txt; 3000-iteration sweeps in this report's session log | CONFIRMED | old (the line dates from the 2026-10-31 hybrids audit; no round re-measured the threshold) | S b2a0fd5b2, E d12738c79 (fork APL copy) |
| N2 | Band 30 with Stormstrike (ladder-first since round 15) casts it (29.8 a fight at 35). The published band 30 entry still has the old Flurry 4 talents; the refresh carries it. N1 measured on both: with Stormstrike, 25% is +2.3 to +2.8% at 30 and 35, +1.3% at 39. | `-talents 0000000000000000-055030310031000000-0000000000000000`, bare, 10,000 iterations | Aligned (Measure at refresh) | - | - |
| N3 | Talent rows against 70755: Mental Dexterity 150%, Mental Quickness 30%, Shamanistic Focus 45%, Flurry 30% (500 ms charge recovery, 15 s, 3 charges), Elemental Weapons 20/40/15%, Elemental Devastation 9% / 10 s, Maelstrom Weapon 20% a stack at 5/5 with 50% a melee hit (408498 effect 1), Stormstrike 150%, 6 s (hotfixed, spellconst), 8% base mana, 20% one charge on Lightning Bolt / Chain Lightning / Earth Shock, Rage of the Farseer 30% / 25 s / 3 min. All match the engine and its tests. | talents/shaman.json; spellconst/shaman.json; SpellAuraOptions 128111, 129538, 223898; sim/shaman/talents.go, stormstrike.go | Aligned | - | - |
| N4 | Published BiS effects at raid 60: Frozen and Molten Heart of the Mountain, Thunderbrew's Boot Flask registered; Rune of the Guard Captain is +14 AP (the forest tripling never applies to a raid boss); The Unstoppable Force's knockback stun is correctly unmodelled (no damage, bosses immune). Item crit on Black Dragonscale Helm, Timbermaw Tunic, Devilsaur pieces is every-table crit per law. | Wowhead Forever tooltips (fetched 2026-11-04); sim/common/item_effects.go, client_procs.go | Aligned | - | - |
| N5 | The rotation ladder golden's header prose (sim/request/ladder.go:1284, 1294) still says Enhancement dual-wields; the code (ladder.go:493-498) gives it a shield. Text only; no number moves. | sim/request/ladder.go | note (For other lanes: platform) | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | Raid 60 breakdown: Lightning Bolt 84.2 casts, 18.2 overloads (30% of casts, landed), Clearcasting 9.4 procs on 93 damage casts (10%), Lightning Mastery 1.48 casts (3 min), Flame Shock refreshed at 21 s (12 + Call of Flame 9). All match the client rows. | spec-breakdown raid 60 horde | Aligned | - | - |
| E2 | Pre-pull Lightning Bolt at -2.5 s is a 2.0 s cast with Elemental Alacrity, so it lands 0.5 s early; moving it to -2.0 s ties (666.1 vs 665.9). No change. | 10,000 iterations, seed 1 | Aligned | - | - |
| E3 | Published BiS effects: Totem of Thunder, Frozen Heart, Thunderbrew's Boot Flask registered; Hand of Edward the Odd's proc is on melee hits (CreateWeaponProcAura), which a caster never makes, so it is valued on its stats, as in game. | sim/common/item_effects.go:1410 | Aligned | - | - |
| E4 | Raid build tolls (Elemental Devastation 1, Elemental Warding 3) stay pending owner as in round 17. | | pending owner | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Healing Stream Totem lasts 5 minutes in 70755, so 2 casts a 300 s fight (Mana Tide takes the Water slot once) with 716 ticks is right. Riptide ticks crit, Healing Stream uses every-table crit only, Totem of Urgency +4% on Lesser Healing Wave: all as tested. | spellconst 10463; restoration breakdown raid 60 | Aligned | - | - |
| R2 | Seed S1 covers the rank rows; the round-17 Nature's Swiftness gate and rank hold. | | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The guide's stat table shows "Hit cap: reached" for Enhancement, and its hover says hit past the cap is worth nothing, while 22% of Enhancement's damage (Earth Shock, Maelstrom bolts, Searing Totem, Fire Nova) still misses 7%. Round 17's weight fix makes the weight non-zero, but the published `hit_to_cap` (melee shape: specials 0) and its wording do not mention the spell side. enhancement.md:93 "past the cap it is wasted" reads the same way. | web/src/lib/bis/hit-cap.ts, web/src/lib/guides/hit-cap-line.ts; sim/cmd/leveling-bis/hitcap.go:36 | note (For other lanes, For the guide lane) | - | - |

## Before/after (10,000 iterations, frozen BiS, seeds 1 / 50001)

Raid 60 is the guide build; bare bands use the published band talents. Before = the same binary with the round-start rotation (identical to $R/before.tsv at raid 60).

| Spec, band | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Enhancement raid 60 | horde | 855.3 to 859.1 | +0.45 / +0.41 |
| Enhancement raid 60 | alliance | 855.3 to 856.6 | +0.16 / +0.13 |
| Enhancement bare 60 | horde | 312.7 to 322.6 | +3.16 / +3.25 |
| Enhancement bare 60 | alliance | 311.3 to 318.2 | +2.20 / +2.17 |
| Enhancement bare 50 | horde | 217.0 to 219.0 | +0.92 / +0.97 |
| Enhancement bare 50 | alliance | 216.9 to 219.3 | +1.12 / +1.19 |
| Enhancement bare 40 | horde | 154.5 to 153.9 | -0.41 / -0.58 |
| Enhancement bare 40 | alliance | 154.2 to 153.3 | -0.62 / -0.32 |
| Enhancement level 46 (band 40) | horde | 164.4 to 163.0 | -0.85 / -0.80 |
| Enhancement level 53 (band 50) | horde | 222.3 to 225.0 | +1.21 / +1.35 |
| Enhancement level 35, band 30 with Stormstrike | horde | 100.3 to 102.6 | +2.30 / +2.29 |
| Elemental, Restoration | both | unchanged | 0.00 |

## Files to regenerate

- data/builds/1.60.1.70755/addon-data.json and addon/ForeverSixty/Data.lua: Enhancement card note 7 and the rotation-list clause, which still reads "On cooldown while you are above half mana" on the rendered page until the regen.
- sim/request/testdata/ladder/shaman-enhancement.golden.md (Earth Shock now fires to 25% mana at every rung).
- Coach starters, if they carry the Enhancement rotation.
- Hand-synced, not regenerated: sim/request/apl/shaman-enhancement.apl.json (site), ui/enhancement_shaman/apls/forever_enhancement.apl.json (fork, E d12738c79).

## Measure at refresh

- Enhancement bands 20 to 60 and raid with the 25% gate (weights move a little toward Intellect and mana at bare bands).
- Enhancement band 30 with Stormstrike (carried).
- Enhancement's relic slot (S3): Totem of the Storm or Totem of Thunder, if the ranker change lands.

## For other lanes

- **Ranker / platform (S3):** a relic slot publishes empty whenever the best relic's gain is under about 1.4 DPS, which every Enhancement totem is. Gate relics on the paired, same-seed delta, or publish the best modelled relic that touches a cast spell as a near tie. The alternative shown today (Totem of Urgency) is a healing relic.
- **Platform (X1):** Enhancement's "Hit cap: reached" line ignores the spell side (16% spell cap); the weight already counts it.
- **Platform (N5):** ladder.go golden header says Enhancement dual-wields; the code gives it a shield.

## For the guide lane

Shipped with this lane: enhancement.md:77 (loop item 6) and :87 (Mana), card note 7. wow-player: HOLD on the first capture (the Mana paragraph said the totems hit harder than the shock; Searing Totem does not, and the stat totems deal no damage), fixed to "pays for Stormstrike, which hits harder than the shock you skip, and for dropping your totems again"; then SHIP on the re-rendered page ($R/shaman/cap/enh-loop-2.png). The rotation-list clause on that page still reads "above half mana" until addon-data regenerates.

Not done here: enhancement.md:93 "it matters up to the cap; past the cap it is wasted" (X1). Hit past 9% still lands Earth Shocks and Maelstrom bolts up to 16%; suggest "it matters most up to the 9% cap for your swings, and keeps landing more of your shocks and bolts up to 16%".

## Still open

- Pending owner, not reopened: Healing Wave 2 tooltip on a geared 60; Lightning Overload rank spells; level-60 book ranks (Strength of Earth 5, Grace of Air 3, Healing Wave 10); Elemental Devastation 1 as the Elemental toll.
- Warcraft Logs is blocked (gcloud login expired). Not worked around.
