# Sim audit round 12: casters, 2026-11-03

Specs: mage-arcane, mage-fire, mage-frost, priest-shadow, druid-balance, shaman-elemental, every band. This lane also owned the haste question for priest-holy and priest-discipline.

## Method

- **Engine.** `lane/r12-casters` from the pin 2606bc05e. **Site.** `wt/r12-casters` from c6bdc855c, with `make simdb`. "Before" is `r12/sb-before` run from `r12/site-before/sim` and `r12/before.tsv`.
- **Client.** 1.60.1.70753 (round 11's layered copy `r11/c70753/builds/1.60.1.70753/raw`). The 70753 `hotfix-push.json` has no rows for the tables read here. Tooltips are from nether.wowhead.com/forever, read 2026-11-03.
- **Haste.** Weights were measured with the dps-per-point sweep plus spell haste (a local probe, not committed). Haste was also scanned with plain spec-breakdown runs at fixed haste (−5 to +20%, down to ±0.01%), paired by labelled RNG. One-iteration combat logs were diffed to find what changed.
- **Tests.** Every fix has a failing test first. Fork: `go test -tags=with_db` in `sim/mage`, `sim/core`, `sim/priest`, `sim/priest/healing`, `sim/priest/shadow`. Each of the 17 fork golden tests was also run alone; only TestP1Mage moves (regenerated in E 7ccb794fd). Site: `go test ./cmd/dps-per-point/`.
- **Runs.** spec-breakdown, 10,000 iterations, seeds 1 and 50001, both factions. That covers raid 60 for all 28 specs and bare 20 to 60 for the six specs plus Holy and Discipline.

## 1. Negative haste at 60 (Holy, Discipline, Fire, bare Elemental)

Real haste should never lower output once the rotation adapts. Three causes were found, and none is a true mana-bound result. All three come from one thing: the sim's events sit on exact grids (2 s regen ticks, the 1.5 s global, prepull times), so two events often happen at the same instant. At haste 0 those ties come out one way. Any haste, even ±0.01%, breaks them. The weights sweep's ±1% steps then measure the tie, not the haste.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Evocation lost one of its four regen ticks whenever it started on a tick.** Regen ticks every 2 s on a grid from the pull. Casts land on the global's 0.5 s grid, so Evocation (12051, an 8 s channel at +1500% regen) often starts exactly on a tick. The tick at the start belongs to the time before the channel. But the tick at the end ran after the channel's regen had already ended, because aura expiry runs before every pending action at the same moment. So that Evocation got 3 ticks, about 1,200 mana short at 60. Any haste moved the casts off the grid and won the tick back: bare Fire 60 read 337.8 at 0% haste, 341.7 at −0.01% and 338.7 at +0.01%. Now the regen ends just after that last tick, so every Evocation gets 4. With the regen grid's phase randomised per iteration (a probe), Fire came out the same as with the fix (342.0), and the haste curve went flat. | Seed-3 combat log: Evocation at 66.00 got ticks at 68, 70 and 72 only. `TestEvocationThatStartsOnARegenTickGetsFourTicks` failed (3 ticks), then passed. `sim/mage/evocation.go`. | CONFIRMED-BUG | old (Evocation's regen aura since upstream; round 10 R1 checked the amount, not the tick count) | E 7ccb794fd |
| H2 | **A cast that ended on the global finished a whole global late if another action spent the global at that moment.** Both healing priest rotations open with a 1.5 s Lightwell at −2.5 s, and the engine buffs Inner Fire at −1 s, the moment Lightwell ends. When a cast ends exactly as the global comes back, the engine finishes the cast inside the global's own event. The prepull Inner Fire ran first at that moment and spent the global. That pushed the global's event, and Lightwell's completion and mana with it, 1.5 s later (Lightwell landed at +0.5 s, not −1.0 s). Now a pending cast that ends before the new global gets its own completion event. | Log: `[-1.00] Casting 10952` (Inner Fire), then `[0.50] Completed cast 724`. `TestLightwellCompletesOnTimeWhenInnerFireIsCastAtItsEnd` failed (completed at 0s), then passed. `sim/core/gcd.go` SetGCDTimer. | CONFIRMED-BUG | old (core since upstream; the collision dates from b70e67d2b, 2026-10-08, the Lightwell opener) | E 1e7e01e4b |
| H3 | **The priest's precast Inner Fire was dropped for the whole fight whenever a prepull cast was still running at −1 s.** With any slowing of casts, Lightwell was still casting at −1 s, Inner Fire failed silently, and the priest saved its 283 mana. In a mana-bound fight that reads as +1.2% healing, so every −1% haste step looked like a big gain. This was the Holy bare −25.6 weight, and Discipline's the same way. Inner Fire now goes up as soon as the opener and the global are done. | `TestInnerFireIsCastWhenAPrepullCastRunsPastItsMoment` failed at −1% haste, then passed at −1, 0 and +1%. `sim/priest/healing/healing_priest.go`. | CONFIRMED-BUG | old (a6bd9de7e, 2026-10-07) | E 1e7e01e4b |
| H4 | **dps-per-point measured band 60 on the raid rotation's ranks.** A band's character wears the bare band's gear and the class kit, but its weights request named no preset. A level-60 request with no preset takes the raid preset's rank choices (`request.PublishedPreset`, added in round 9). So bare Elemental 60 cast full-rank Lightning Bolt on bare gear, ran dry, and measured 146 DPS where its published bare band does 198. Its spell-power figure came out 0.216 where bare measures 0.379, and its haste read −1.0 per 1%. The band requests now say bare, and the raid request says raid. Only band 60 was affected (bands 20 to 50 already default to bare). | `TestABandIsMeasuredUnderTheBarePresetAndTheRaidUnderTheRaids` failed (preset ""), then passed. Published `dps-per-point.json` band 60 for shaman-elemental is 0.2021. | CONFIRMED-BUG | R9 S 49356dbcf (the request gained a preset; dps-per-point did not send one) | S 2d9d55d90 |
| H5 | Fire's haste is close to zero on its own, and that is right. Spell haste does not shorten the global in Classic: the engine comment and `cast.go` leave the GCD unhasted. Scorch's 1.5 s cast already sits on the 1.5 s floor, and so does the 3-stack Pyroblast. Bare Fire is also mana-bound. So +20% haste is worth only about +0.5%. | Paired scans; `sim/core/cast.go:243` | Aligned (Era rule; Forever states no change) | - | - |
| H6 | A small tie at exactly 0% haste remains for Discipline and Elemental. Discipline reads 310.7 / 312.6 / 311.9 HPS at −0.01 / 0 / +0.01%, and bare Elemental 197.8 / 197.9 / 197.2 DPS, both within 0.6%. These are other same-instant ties (heal-profile pulses, the 2.0 s Lightning Bolt against the 2 s regen grid). Randomising the regen tick's phase per iteration would remove the regen kind everywhere. That probe moved raid and bare 60 by under 0.15% for every spec checked except bare Fire, but it would move every mana user's fork goldens, so it is left to the owner. | probe runs | OPEN (owner: random regen phase) | old | - |

Haste per 1% after the fixes, on the dps-per-point characters, seed 7. The "before" column is this lane's re-measurement at the pin; the published BiS weights come from an older engine.

| spec, band | before | after |
|---|---|---|
| priest-holy raid 60 (HPS) | −2.01 ± 0.16 | −0.10 ± 0.15 |
| priest-holy bare 60 (HPS) | −2.52 ± 0.04 | +0.71 ± 0.05 (bare preset) |
| priest-discipline raid 60 | −0.40 ± 0.08 | +0.83 ± 0.08 |
| priest-discipline bare 60 | −0.56 ± 0.04 | +0.73 ± 0.03 |
| mage-fire bare 60 (DPS) | −1.23 ± 0.18 | +0.80 ± 0.06 |
| shaman-elemental bare 60 | −1.04 ± 0.08 | +0.77 ± 0.04 |

## 2. Frost's leveling build (main's e2b3a9a53)

Old `230005/23023/253510130100301231`, new `2300052/5102/253510130101301221`, the same as the raid build. The change adds Missile Barrage 2, three more Incineration (5) and Ice Block. It drops Accelerant 2, Impact 3 and the third Winter's Chill point.

| # | finding | evidence | verdict |
|---|---|---|---|
| F1 | Bare 60: horde 346.8 / 346.9 to 367.2 / 367.2 (+5.9 / +5.9%), alliance 344.3 / 344.6 to 366.1 / 365.3 (+6.3 / +6.0%), seeds 1 / 50001. Bands 40 and 50 only trade the third Winter's Chill point for Ice Block, and their DPS is identical to the decimal. Bands 20 and 30 are the same build. | spec-breakdown, 10,000 iterations | gain ≥1% beyond error at both seeds |
| F2 | The spec tree holds: Frost keeps 31 points. Nothing the rotation casts is dropped. The APL casts Frostbolt, Ice Lance, Arcane Missiles (now with Missile Barrage), Ice Barrier, Evocation, Shoot, and Cone of Cold on five Winter's Chill stacks. That Cone of Cold line is refused at the ranker's 30 yards, and Winter's Chill stays at 2. Impact ("your Fire spells... stun") and Accelerant (Fireball, Frostfire Bolt, Arcane Wrath) touch only spells Frost never casts. | talents/mage.json texts; APL spell ids | passes; the dropped points did nothing for the spec (the Shadow-build precedent) |
| F3 | Client: Missile Barrage 400588 2/2 gives Frostbolt a 20% chance. Incineration 18459 5/5 gives Ice Lance +10% crit. | talents/mage.json; round 11 M6 engine check | Aligned |

**Kept.**

## 3. Elemental Tidal Focus at bare 50/60

| # | variant | bare 50 (a s1 / a s50001 / h s1 / h s50001) | bare 60 | verdict |
|---|---|---|---|---|
| T1 | Improved Healing Wave 5 → Tidal Focus 5 (`551033130000135/03/053352`) | +5.81 / +5.79 / +5.68 / +5.64% | +5.64 / +5.84 / +5.69 / +5.90% | Fails rule clause 1: drops Improved Healing Wave, an unmodeled guide talent (the leveling self-heal) |
| T2 | Totemic Focus 5 → Tidal Focus 5, keeping every unmodeled talent (`551033130000135/03/553302`) | −18.97 / −19.03 / −17.11 / −17.12% | +5.01 / +5.17 / +5.01 / +5.20% | Fails at 50: the build is walked top row first, so Tidal Focus is bought before Mindfulness |

**Held.** No variant passes the rule. The gain needs Improved Healing Wave given up, and that is the owner's leveling-build call (round 11 E7). Absolute horde seed 1: 50: 155.0 → 163.8; 60: 191.2 → 202.0.

## 4. Chain Lightning's cast time

| # | finding | evidence | confirmed? | origin | fix |
|---|---|---|---|---|---|
| C1 | Every rank (421, 930, 2860, 10605) is a **2.0 s** base cast: SpellMisc CastingTimeIndex 5, SpellCastTimes 5 Base 2000. It has a 6 s category cooldown and costs 315 mana at rank 4. Elemental Alacrity (16578: aura 107, misc 10, mask 0x3 / 0x1000, so Lightning Bolt, Chain Lightning and Lava Burst) takes 0.17 / 0.33 / 0.5 s off. So it is **1.5 s** with 3/3, which both builds take. The engine registers 2000 ms (`shaman/chain_lightning.go:56`) and applies Alacrity to it (`talents.go:551`). | client rows (one source plus the talent text) | Aligned | - | - |

## 5. Fresh re-audit

### Mage (Arcane, Fire, Frost)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Evocation's lost end tick (H1). It moves bare Fire most (it Evocates every fight), and the gnome mages at bare 30 to 60 (Frost alliance +1.5% at 30). | above | CONFIRMED-BUG | old | E 7ccb794fd |
| M2 | Landing openers. Fire's pre-pull Pyroblast at −6 s (6 s cast, lands on the pull plus its 24 yd/s travel). Frost's Ice Barrier at −4 s, then Frostbolt at −2.5 s (2.5 s with Improved Frostbolt). Arcane's pre-pull Frostbolt at −2.5 s and its Frostbolt under Presence of Mind. Removing Arcane's pre-pull costs 0.9 / 1.0% at raid; removing the PoM Frostbolt costs 0.4% at raid and is within error at bare. All are legal and timed as a player would play them. | spec-breakdown variants, both seeds | Aligned | - | - |
| M3 | Client rows for every spell the mage APLs cast at 60 match the engine: casts, costs, globals, categories, travel speeds (Frostbolt 28, Pyroblast 24, Fireball 24, Ice Lance 38, Scorch and Fire Blast none), Ice Barrier's 60 s buff and 30 s category, Presence of Mind's 180 s category, Arcane Power 15 s / 180 s. The fork's conformance golden (mage.golden.md) shows no mismatch. | SpellMisc, SpellCooldowns, SpellPower, SpellCastTimes | Aligned | - | - |
| M4 | Round 11's opens stand: Ice Lance's push-only category (M4), Fire Power on dot ticks (M10). | - | OPEN (unchanged) | - | - |

### Shadow

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Landing: Shadowform moved to −20 s, and Dark Sacrifice's paced line is round 11's. The core fix (H2) moves no Shadow number at any band. | before/after | Aligned | - | - |
| S2 | Conformance (priest.golden.md) shows no mismatch for Mind Flay, Mind Blast, Shadow Word: Pain, Devouring Plague, Shadow Word: Death or Shadowform. | fork golden | Aligned | - | - |

### Balance

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Starfire 9876 / 25298: 3.5 s, no travel time. Wrath 9912: 2.0 s, speed 20 (the engine has MissileSpeed 20, which Nature's Grace's no-missile check reads). | client rows | Aligned | - | - |
| B2 | Innervate (core `InnervateAura`, 20 s) has the same end-tick ordering as H1. Randomising the regen phase moved Balance by 0.00% bare and −0.05% raid, so it is below error here. Left to the auras lane (core/buffs.go). | probe | Noted (below error) | old | - |
| B3 | Round 11 B1 (Improved Moonfire's reach) is carried. | - | OPEN | - | - |

### Elemental

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | Landing lines are legal: Mana Spring pre-pull at −4 s; Lightning Bolt rank 9 pre-pull at −2.5 s (2.5 s with Alacrity: the client base is 2500 at ranks 9 and 10, which already carries Forever's cut); rank 10 while mana is ahead of the fight's share left; Fire Nova and Chain Lightning only on 2+ targets. Lightning Bolt 15207 / 15208: 178 / 196 base, +1.0 / +1.2 a level, 0.714, 190 / 220 mana, speed 20. Lava Burst 1326047: 2.5 s, 10 s category, 265 mana, the +20% dummy effect. Flame Shock 29228: 6 s shock category, 3 s ticks. Totems have the client's 1.0 s global (`shaman/totems.go:15`). | client rows | Aligned | - | - |
| E2 | Elemental Mastery 16166: SpellAuraOptions gives ProcCharges 1 (mask 87376, the row PoM and Clearcasting share), but the Forever tooltip only says "Lasts 15 sec" with no "next". Round 6 and 11's E6 stands. | SpellAuraOptions 128637; tooltip 2026-11-03 | OPEN (one source each way) | old | - |
| E3 | Searing Totem's bolts reach 20 yd (range 3) while the ranker stands casters at 30 yd. The totem stays where it is dropped, and the guide says to drop it within 20 yd of the boss, so the sim's always-in-range totem matches that play. | SpellRange 3 | Aligned (play assumption) | - | - |

## Cross-cutting

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | The core global fix (H2) applies to every class: a cast that ends on the global is no longer delayed by a prepull or other action spending the global at that moment. At raid 60 only warrior-arms moves, by −0.08 / −0.10% (alliance and horde, both seeds; confirmed by building without the change). Arms casts Slam, so a Slam that ends on its global now lands on time. No fork golden moves. | raid-60 grid, all 28 specs | (side effect of H2) | - | E 1e7e01e4b |
| X2 | Item spell haste counts twice. `character.go:269` multiplies CastSpeedMultiplier by an item's SpellHaste, and the same stat also reaches `stats[SpellHaste]` through EquipStats. MeleeHaste does the same. No published band carries haste (`haste_on_items` false everywhere), so nothing published moves. | engine reading | CONFIRMED-BUG (no published number), not fixed: sim/core stat plumbing, routed | old (upstream) | - |

## Before/after

Before: pin 2606bc05e on c6bdc855c. After: E 1e7e01e4b on S 2d9d55d90. 10,000 iterations. The dps-per-point fix moves no spec-breakdown number.

| Spec, band | horde s1 before → after | alliance s1 / s50001 | horde s1 / s50001 |
|---|---|---|---|
| mage-fire bare30 | 69.9 → 71.4 | +0.83 / +0.69% | +2.15 / +2.29% |
| mage-fire bare40 | 154.9 → 156.8 | +1.14 / +1.01% | +1.23 / +1.23% |
| mage-fire bare50 | 193.1 → 195.1 | +1.36 / +1.46% | +1.04 / +1.09% |
| mage-fire bare60 | 330.3 → 334.2 | +1.15 / +1.09% | +1.18 / +1.21% |
| mage-frost bare30 | 76.6 → 76.6 | +1.54 / +1.54% | 0 / 0 |
| mage-frost bare50 | 248.2 → 248.2 | +0.32 / +0.32% | 0 / 0 |
| mage-frost bare60 | 346.7 → 346.7 | +0.17 / +0.15% | 0 / 0 |
| mage-arcane bare30-60 | 214.1 → 214.4 (50) | +0.06 to +0.18% | +0.00 to +0.23% |
| priest-discipline bare60 (HPS) | 318.1 → 317.9 | −0.03 / −0.03% | −0.06 / −0.06% |
| warrior-arms raid60 | 731.7 → 731.1 | −0.08 / −0.10% | −0.08 / −0.10% |

Unchanged (under 0.05% at every seed and faction): every raid-60 spec except warrior-arms; druid-balance, priest-shadow, shaman-elemental and priest-holy at every bare band; every spec's bare 20; mage-frost bare 40 (+0.10% alliance only).

Causes: Fire and the gnome Frost and Arcane bands, H1; Discipline, H2's Lightwell timing; Arms, X1.

## Files to regenerate

- Fork: TestP1Mage.results was regenerated in E 7ccb794fd (only the forever_frost NoBuffs rows). Nothing else.
- Site: sim/request TestRotationLadder goldens for mage-fire, mage-frost, mage-arcane (H1) and warrior-arms (X1); the healer ladder rungs if they record Lightwell timing.
- `data/builds/<active>/dps-per-point.json`: every spec's band "60" (H4), then Data.lua.
- The BiS files' haste scale factors (published negatives came from older engines; see Measure at refresh) and bis/mage-frost's bare talents (the new leveling build).

## Measure at refresh

- Haste scale factors at bare and raid 60 for Holy, Discipline, Fire and Elemental, after H1 to H4. Expect zero to positive.
- dps-per-point band 60 for every spec whose bare and raid rank choices differ (Elemental, Shadow, the healers).
- Frost bare picks on the new leveling build (Missile Barrage values Intellect and crit differently).

## For other lanes

- **Healers lane:** H2 and H3 are in `sim/priest/healing` and core. Holy and Discipline now keep Inner Fire at every haste, and Lightwell lands at −1 s. No healer APL was edited. The heal profile's fixed pulse grid still makes small same-instant ties for Discipline (H6).
- **Guides lane:** Chain Lightning is 2.0 s base on every rank and 1.5 s with Elemental Alacrity 3/3, which both Elemental builds take. So "2.0 s" is the untalented figure, and the talented guide should say 1.5 s.
- **Physical lane:** warrior-arms raid 60 −0.08 to −0.10% from X1 (Slam now completes on its own global).
- **Auras lane:** Innervate's 20 s window has H1's end-tick ordering (`core/buffs.go InnervateAura`). It is below error for Balance; the helper pattern is in `mage/evocation.go`.
- **Itemscrit lane:** X2, item SpellHaste and MeleeHaste applied twice (`sim/core/character.go:258-275`).
- **Owner:** the random regen phase (H6) and Elemental's Tidal Focus (section 3).

## Still open

- H6 random regen phase; E2 Elemental Mastery's charge; Balance B3; mage M4.

## Sources

- 1.60.1.70753 client tables (round 11's layered copy): SpellMisc, SpellCastTimes, SpellCooldowns, SpellPower, SpellEffect (16578 row 694339, 16166), SpellAuraOptions 128637, SpellRange, SkillLineAbility 8724-8727; talents/mage.json, talents/shaman.json.
- nether.wowhead.com/forever/tooltip/spell/16166 (2026-11-03).
