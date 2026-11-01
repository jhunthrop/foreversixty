# Sim audit round 6: healers

**Wild Growth:** it is a Restoration talent in Forever: the 31-point capstone, 1 point, behind Living Spirit 3/3. The talent gives rank 1 at 40, and the trainer teaches ranks 2 and 3 at 50 and 60. Rank 3 costs 1,050 mana and has a 3 s cast (2.5 s with Naturalist) and a 6 s cooldown. It heals the target and their party (five players) for 97 a tick plus 3.3% of healing power, over 7 ticks, front-loaded.

The engine has it right. The build skips it for one measured reason: our heal fight runs a druid out of mana before it runs out of time. In that fight Rejuvenation with Bloom heals more per point of mana than Wild Growth.

We found one real flaw on the way. The raid-wide damage never touched the healer, so a fifth of every Wild Growth landed on a druid at full health. That is fixed. With it fixed, the best of 36 Wild Growth lines comes out level with the guide rotation (-0.3%, inside error), not ahead of it. So the build stays: the adoption rule asks for +1%.

In a fight with heavy bursts of group damage, take it. The guide now says so.

Lane: wt/r6-healers (site) and lane/r6-healers (fork, from 8baffe5fc). Engine files touched:
- `proto/common.proto` and `sim/core/proto/common.pb.go`
- `sim/core/raid_damage.go`
- `sim/core/dot.go`
- `sim/druid/healing.go`
- `sim/priest/{renew,lightwell,binding_heal,heals,priest}.go`
- `sim/shaman/riptide.go`
- the conformance goldens
- `sim/paladin/holy/rotation_test.go`
- the four healer APL copies

The cherry-picked casters commit 9a5c25dff also touches `sim/mage/*` and `sim/priest/*` dots.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Raid-wide pulses never reached the healer.** `applyRaidDamageModel` picked pulse targets from the fake members only. The healer's own health never moved. So every group heal (Wild Growth, Tranquility, Prayer of Healing, Chain Heal's jumps, Lightwell, Healing Stream) put a fifth of its healing on a full-health healer, and Holy Binding's self-heal always overhealed. The profile itself calls the pulses "raid-wide damage (a breath, a wing buffet, a periodic shadow bolt volley)", and none of those skip the caster. The 2026-10-08 mana audit (section 9) had already named it: "the profile does not model the healer's own health". | `sim/core/raid_damage.go:52-75` at 8baffe5fc; `data/curated/heal-profile.json` pulse.reason. Fix: new `RaidDamageModel.pulses_reach_healer`, off by default so the fork's test profile does not move, and on in the curated profile (`pulse.reaches_healer`). Each pulse still hits three members, now drawn from five instead of four, so total incoming damage is unchanged. Tests first: `TestRaidDamagePulsesReachTheHealer` (fork) and `TestCuratedHealProfilePulsesReachTheHealer` (site). | CONFIRMED-BUG (code; the profile's own text) | E bc60a7231, S 924d8898e |
| X2 | **Heals over time snapshotted healing power.** Forever's periodic effects are dynamic. Blizzard on forever-bugs #137 (2026-10-01) writes of "DoT/HoT boosting effects that apply dynamically" under "dynamic DoTs", and #398 has Agony adjust "every tick". `DotConfig.DynamicSnapshot` re-runs a HoT's own OnSnapshot before each tick. It is set on Rejuvenation, Regrowth's HoT, Wild Growth, Tranquility, Renew, the Lightwell renew and Riptide. Healing Stream already computes each pulse as it lands. | Casters' E 9a5c25dff cherry-picked as E 2edc607c6. Test first: `TestDynamicHotReadsHealingPowerEachTick` (it would not compile before the field existed; its snapshot half holds the old behaviour). **Effect on the profile:** within 0.1%. No healer gains healing power in the middle of a HoT often enough to matter. Single Blizzard statement for HoTs (#137); #398 covers DoTs. | CONFIRMED-BUG (one Blizzard source, flagged) | E 5d9cdd67c |
| X3 | Coach downranks were judged against Ahn'Qiraj book ranks: Renew 25315 and Prayer of Healing 25316, which no Phase 1 priest can learn. `coachrules` now drops book ranks (curated/book-ranks.json) when it picks a downrank family's top rank. Book ranks stay in the id list, so their casts are still recognised. Auras and core lists are unchanged: tying their uptime to the stamped id broke the warrior web fixtures (Battle Shout 25289). | `data/pipeline/coachrules.py`; test first: `test_no_downrank_is_judged_against_an_ahnqiraj_book_rank` (it failed on Renew 25315). Only the two priest coach files move. | CONFIRMED-BUG | S c8e762ec1 |
| X4 | Priest APL notes stated ranks that were wrong. Discipline's said "Prayer of Healing at rank 1", but the list casts rank 4. Holy's said "every other spell keeps its top rank" without saying the tops are the Phase 1 trainer ranks. Both notes now name the Phase 1 tops: Prayer of Healing 4, Greater Heal 4, Renew 9. | Notes in `data/curated/apl/priest-*.json`; book-ranks.json. | CONFIRMED (prose) | S c8e762ec1 |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Wild Growth (owner report).** It is a talent (node 104910, tier 6, prereq Living Spirit 3) and trainable: SkillLineAbility ranks 408120 at 40, 1238214 at 50, 1238215 at 60. Rank 3: 97 a tick plus 1.6 a level, coefficient 0.033 a tick, 7 ticks over 7 s, 6 s cooldown, 1,050 mana, on the target's party. The hotfix push gives every rank cast-time index 14 (3 s); the shipped tables say instant. The engine matches all of this (`sim/druid/wild_growth.go`, `healing_tables.go:247`). **Why it is skipped:** not a trainables mistake, not a missing spell, not a zero value. Cast, it heals about 5 a mana with no overheal. But the profile is mana-bound: the druid's mana lasts 285 of 300 s, and overheal is 3%. Rejuvenation heals about 5.5 a mana before Bloom adds its Spirit-sized heal. Wild Growth's strength, healing a burst across five players in one cast, never comes up in a fight with steady pulses. Before X1, a fifth of each cast also overhealed on the druid itself. **Measured after X1, X2 and the R2 list:** 36 Wild Growth lines (2/3/4 of the five hurt, at 60/80%, three priority positions, paced or not), with the point taken from Swiftmend. Swiftmend eats a HoT before it can Bloom, so that point is nearly free (dropping it alone: +0.4%). The best line came out at 747.8 against 750.2 for the guide (seeds 1 and 50001, 10,000 iterations). | Client talents/druid.json; trainables; push diff rows 444, 445 and 1189; spec-breakdown runs (lane scratch `rot2/`, `wg_final.json`). | Aligned (settled; the adoption rule holds the build) | - |
| R2 | The list now heals the druid like a party member: a Rejuvenation line for player 0, and Tranquility counts two of the five. +0.8% over the old list on the X1 profile. | `data/curated/apl/druid-restoration.json`. | Follows X1 | S 8f4cddded |
| R3 | Every other signature spell was checked as trainable, implemented, valued and cast: Rejuvenation, Regrowth, Healing Touch, Tranquility, Swiftmend, Nature's Swiftness and Innervate. Lifebloom, Nourish and Tree of Life are SoD runes with no learn level, so they are not trainable. Nature's Focus (pushback) is in the build and does nothing in a sim without interrupts; the guide already says so. | trainables/druid.json. | Aligned | - |

## Priest: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | X1 is worth +6% here: Prayer of Healing, Holy Binding and Lightwell now heal a priest who is actually hurt. The list is unchanged. Widening Prayer of Healing's "three of four hurt" to "three of five" measured -1.0%, and per-member self lines -1.0%. | spec-breakdown runs. | Follows X1 | - |
| H2 | Greater Heal is never cast. Greater Heal rank 4 on the tank measured +0.1% (Heal rank 4 has the same 0.857 coefficient at less mana), so the list keeps Heal. Holy Nova (not taken) heals about 3 a mana across five. Binding Heal in front of Flash Heal measured +0.7%, because Holy Binding already heals the priest off Flash Heal. Neither change clears 1%. Circle of Healing is not trainable. Shadowfiend has no learn level (round 3). | trainables/priest.json; runs. | Aligned | - |

## Priest: Discipline

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| D1 | **The list never cast Greater Heal,** the priest's big heal. Greater Heal rank 4 as the tank's filler (Heal stays on the members), with the members' emergency Flash Heals paced at 70% of the fight clock: +3.9%. | runs (lane scratch `disc/`). | CONFIRMED-BUG (rotation) | S 8f4cddded |
| D2 | **Binding Heal was missing.** The 70707 push teaches it from the Holy skill line at 25, 32, 38, 44, 50 and 56 (pipeline P10). The 30 October hotfix had turned the old talent into Holy Binding, and the engine had neither. Each rank heals the target and the priest for the Flash Heal amount of its level. Client rows 401937 and 1240770-1240774; Wowhead Forever rank 6 (2026-11-01): "Heal yourself and a friendly target for 778 to 914", 380 mana, 1.5 sec cast. That is one source by the brief's rule, plus the push's SkillLineAbility rows. Now in the engine; conformance matches all six ranks. Discipline casts it in front of each emergency Flash Heal while the priest is under 90%: +7.7%. | Test first: `TestBindingHealHealsTheTargetAndThePriest`. unconfirmed: one crit roll or two (the engine rolls two). | CONFIRMED-BUG (missing trainable) | E bd67dab89, S c8e762ec1 |
| D3 | Penance, Power Word: Shield, Renew, Prayer of Healing, Lightwell and Dark Sacrifice: trainable, implemented and cast. Pacing Penance on members changed nothing. Power Infusion stays out (owner ruling). The mana still runs dry at about 127 s; nothing tried here moved that further. | runs. | Aligned | - |
| D4 | Benediction (integrator item 3): its Eye of Divinity drops from Majordomo in Molten Core, which `forever-raid-phases.json` lists as opening "later". So it is not obtainable in Phase 1, and its absence from loot.json's candidate pool is correct for now. | `data/curated/loot/forever-raid-phases.json`. | Aligned (note for the ranker lane) | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **The list precast Greater Blessing of Light,** which Forever removed (pipeline P5, the 70707 push). The line is gone: **-10.1% alliance / -10.5% horde**, the whole of this spec's move. The hybrids lane removes the engine and preset side. The fork rotation test stops expecting the cast. | `data/curated/apl/paladin-holy.json`; fork `sim/paladin/holy/rotation_test.go`. | CONFIRMED-BUG (removed spell) | E 671990588, S c8e762ec1 |
| P2 | The list heals the paladin like a party member (X1). Neutral (+0.2%), kept so the healer is not left at zero all fight. | | Follows X1 | S 8f4cddded |
| P3 | Light's Vigil (capstone) is in the build but never cast. Rank 3 costs 1,340 mana plus the Holy Shock, for a 704 + 14.3% party heal: about 2.5 healed a mana against Flash of Light's 5+. That stays true with the healer among the five. The Holy points left (Divine Strength, Improved Seals, Holy Conduit, Consecrated Ground) add nothing to a healer, so the point costs nothing. Holy Shock, Flash of Light, Holy Light, Divine Favor and Lay on Hands are trainable and implemented. Oath of Light and Beacon of Light have no trainer row. | engine `lights_vigil.go`; trainables. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **The list ended every fight with mana unspent while damage went unhealed:** mana lasted 337 s of 300, with 0.3% overheal. Chain Heal waited for a member under 70% and for half the mana bar; the list's own notes said a retune "was not done". Chain Heal now opens when any of the five (the shaman counted) is under 90%, on the fight-clock pace (mana at least 90% of the share of the fight left): **+6%**. This came from a grid of 72 variants on thresholds and gates; the plateau is flat within 0.5%. | runs (lane scratch `sham/`, `sham2/`). | CONFIRMED-BUG (rotation) | S 8f4cddded |
| S2 | Riptide, Chain Heal, Healing Wave, Lesser Healing Wave, Healing Stream, Mana Tide, Water Shield and Nature's Swiftness: trainable, implemented and cast. Earth Shield is a `class_spell` with no skill line (like Sacred Shield at 80), so it is not trainable. Healing Rain and Ancestral Guidance have no row. Chain Heal's jumps can reach the healer (its comment says otherwise; the code checks health bars). | trainables/shaman.json; `sim/shaman/chain_heal.go:62`. | Aligned | - |

## Before/after

Fork 8baffe5fc and origin/main 21d97955f, against fork 671990588 and site c8e762ec1. Raid preset, band 60, heal profile, 10,000 iterations, effective HPS.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| Restoration druid | 741.5 → 750.2 (+1.2%) | +1.2% | 739.2 → 748.8 (+1.3%) | +1.3% |
| Holy priest | 641.4 → 680.7 (+6.1%) | +6.0% | 623.2 → 662.6 (+6.3%) | +6.3% |
| Discipline priest | 538.4 → 611.9 (+13.7%) | +13.6% | 548.0 → 621.2 (+13.4%) | +13.3% |
| Holy paladin | 553.8 → 497.9 (-10.1%) | -10.2% | 544.2 → 487.0 (-10.5%) | -10.5% |
| Restoration shaman | 641.9 → 680.6 (+6.0%) | +5.9% | 640.2 → 677.1 (+5.8%) | +5.8% |

Each step's share (2,000 iterations, seed 1, alliance):
- X1 alone: druid +0.3, Holy +6.1, Discipline +1.5, paladin 0.0, shaman -0.3.
- X2: within 0.1 for every spec.
- Then R2, D1, D2, P1 and S1 as their rows give.

## Still open

- **Profile retune check.** Incoming damage is about 872 a second. After this round, the druid's raid set covers about 86% and the paladin's about 57%, both just outside the retune rule's 60 to 85%. The nightly ranker's sets and `hps_at_bounds` will show it. Retuning the profile, and /tiers/healer, belong to the integrator and the page lane.
- **Wild Growth under burst group damage.** The profile has steady pulses only. Launch raid logs (druid Wild Growth casts, overheal and mana at the end) would show whether Forever fights reward it.
- **Binding Heal's two halves:** one crit roll or two. One in-game cast with a crit settles it.
- **TestRotationLadder** fails for every spec on this branch, the healers among them. Per the brief it was not regenerated; the integrator does that at the pin.
- Fresh re-audit scope: this lane re-checked each healer's builds, rotations, signature spells, trainables and the heal profile. The spell amounts were left to the hybrids and casters lanes, as briefed.
