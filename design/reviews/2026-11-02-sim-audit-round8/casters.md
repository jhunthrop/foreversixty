# Sim audit round 8: casters, 2026-11-01

Specs: mage-arcane, mage-fire, mage-frost, priest-shadow, druid-balance, shaman-elemental.

## Method

- **Engine.** Lane branch `lane/r8-casters` from the pin `cd4dfdcc0`; the before column runs a second worktree detached at the pin.
- **Site.** Branch `wt/r8-casters` from origin/main `170b7b84c`, with `make simdb`. The before column runs a site worktree detached at origin/main.
- **Client.** The 70707 tables with the hotfix push layered in (pipeline lane, fetched 2026-11-01 21:42 UTC). A spellconst regenerated from them equals the committed `spellconst/<class>.json` for all four classes (0 differing spells), so the generated engine constants are current.
- **Coverage I leaned on, and checked.** The fork's conformance goldens (`sim/core/testdata/conformance/*.golden.md`) compare every registered spell's cost, cast time, cooldown, GCD, learn level, duration and damage roll to the client at levels 10 to 60. mage, priest, druid and shaman have 0 mismatch rows. The per-package `spellconst_damage_test.go` ladders hold every hand-written rank table (Fire Blast, Frostbolt, Arcane Blast, Starfire, Wrath, Moonfire, Insect Swarm, Lightning Bolt, Chain Lightning, the shocks, Lava Burst, Fire Nova, the fire totems). The sets golden covers the worn set bonuses.
- **New this round.**
  - Every rank each band casts, read from the breakdown, against SkillLineAbility and SpellLevels.
  - A scan of every spec at levels 10 to 50 for spells registered above their learn level.
  - Class regen formulas checked against two emulator families.
  - Out-of-mana time on every bare band.
- **Runs.** `spec-breakdown` at 10,000 iterations, seeds 1 and 50001, both factions, raid 60 and bare 60/50/40/30/20. The full grids are `before.txt`, `after.txt` and `delta.txt` in the lane scratch dir.
- **Tests.** Every confirmed bug got a failing test first. I saw it fail, then saw it pass.

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | Round 7 side effects, re-run: Shadowform x Mental Agility, Arcane Instability's all-crit on wands and procs, Silent Resolve and Searing Light schools, Moonkin Aura counted once, Insect Swarm crit. | Fork tests `TestShadowformCostMultipliesWithMentalAgility`, `TestArcaneInstabilityCritReachesEveryCrit`, `TestSilentResolveCoversShadowThreat`, `TestSearingLightIsAllHolyDamageDone` and `TestMoonkinCritIsGrantedOnce` pass on the lane. Measured: bare 60 Balance with `-buffs moonkin_aura` gives the same 265.4 DPS as without it, so the raid's aura does not stack with the druid's own form. Raid Insect Swarm crits 22.7%. | Aligned | - |
| I2 | The supplement wands' school (Pyric Caduceus and others) reaches published numbers only at the owner's single refresh. The committed simdb still has the old rows, so the band-50 numbers here do not show it yet. | round 7 I2 | Aligned (waits on the refresh) | - |
| I3 | Fire's Improved Scorch upkeep line (+0.4%) stays held. | No new evidence. Raid Fire: Scorch is 33% of damage, with the debuff at 5 stacks from the mage's own casts. | Held | - |
| I4 | Rank per band, 20/30/40/50/60, for every spell the six rotations cast. | Fireball, Frostbolt, Arcane Missiles, Scorch, Fire Blast, Pyroblast, Ice Lance, Ice Barrier, Mind Blast, Mind Flay, Shadow Word: Pain, Devouring Plague, Shadow Word: Death, Starfire, Wrath, Moonfire, Insect Swarm, Lightning Bolt, Flame Shock, Lava Burst, Fire Nova, Searing Totem and Mana Spring Totem are each the highest SkillLineAbility trainer rank at or below the band (SpellLevels). Forever's own levels are followed: Starfire r5 at 50 and r6 at 58, Ice Lance at 20/28/34/42/48/56, Shadow Word: Death at 32/40/48/56, Lava Burst at 30/40/50/60, Searing Totem at 10 to 60 in steps of 10. **One miss: Arcane Blast at 40 to 49 (A1).** | 1 CONFIRMED-BUG | S dbd2f6449 |
| I5 | Coefficients, cast times, costs and cooldowns against SpellEffect, SpellMisc, SpellPower and SpellCooldowns. | Conformance goldens: 0 mismatches. I spot-checked the hand-written tables against the layered rows: Fire Blast 10199 (438 base, 0.429, 8 s, 340 mana), Arcane Blast 1239697 (101, 1.0 a level, 15% base mana), Lava Burst 1326047 (220, 0.714, 10 s, 265 mana, +20% on Flame Shock), Flame Shock 29228 (166, 44 a 3 s tick, 410 mana, 6 s), Fire Nova 408345/408428 (375 mana, 10 s, 239 + 1.3 a level, 0.143), Lightning Overload's 408477 (98, 0.357: half of rank 10). | Aligned | - |
| I6 | Clearcasting, Elemental Focus, Elemental Mastery. | Elemental Focus 16164 is a 10% chance on a cast, and 16246 is -100% cost for one charge. The class mask covers Lightning Bolt, Chain Lightning, Lava Burst, the shocks and Fire Nova, but no totem (maskhit). Elemental Mastery 16166: -0.5 s on Lightning Bolt and Lava Burst (mask 1 and 4096), no Chain Lightning cooldown, 15 s, 180 s. All match `shaman/talents.go`. | Aligned | - |
| I7 | Master of Elements, Combustion, Ignite, Arcane Power, Presence of Mind. | Master of Elements 29074 refunds a share of base cost on Fire and Frost crits (`mage/talents.go:377`). The others are unchanged since round 7 and re-read against the 70707 text. Raid casts a fight: Arcane Power and Presence of Mind 1.49, Combustion 1.27. Those are 3 min cooldowns over a 180 s fight whose length varies. | Aligned | - |
| I8 | Shadow Weaving self-stacking. | 15257 procs from harmful spells (ProcTypeMask 65536). 15258 is aura 270, 2% Shadow damage taken from the caster, 5 stacks (CumulativeAura 5), 15 s (DurationIndex 8). The engine's self buff of 2% a stack is the same thing on one target. It adds a stack on every landed Mind Blast, Mind Flay, Shadow Word: Pain, Devouring Plague and Shadow Word: Death, once a cast, not once a tick. | Aligned | - |
| I9 | Icy Veins-equivalents. | Icy Veins 425121 is a Season of Discovery rune row (AcquireMethod 3), not learnable. Troll Berserking and Orc Blood Fury are modelled (round 7). | Aligned (none in Forever) | - |
| I10 | Mana regen: Spirit, mp5, Evocation, mana gems at their band ranks. | Mage and priest: 12.5 + Spirit/4 a tick. Druid: 15 + Spirit/5. Both match the emulators. **Shaman read 15 + Spirit/5 where the server gives 17 + Spirit/5 (E1).** Gems: Agate 400 ±12.5% (learned at 28), Jade 600 (38), Citrine 850 (48), Ruby 1100 (58). They match `ManaGemManaRestored`, and each band uses only the gems it has learned. Gems and Demonic Rune share category 1153, which matches the engine's `GetConjuredCD`. Evocation 12051 is +1500% regen and 100% while casting, for 8 s every 8 min. **Innervate was cast below its level 40 (B1).** | 2 CONFIRMED-BUG | E 13746e744, E 83e720c41 |
| I11 | BiS set bonuses, procs and use effects. | Magister's 2/3 piece (resist 8, +18 spell damage 9346), Ironfeather 2 piece (+20 spell damage 14799), Dreadmist and Elements single pieces (no bonus): all `declared, matches` in the sets golden. Frozen Heart of the Mountain is +29 Frost and Shadow only (school mask 48, `common/item_effects.go:3796`), so the Fire, Arcane, Nature and Elemental sets get only its 9 hit. Robe of the Archmage restores 500 ±50% every 5 min. Idol of the Moon gives Moonfire +33. Totem of Thunder gives Lightning Bolt +1% crit (relic mods generated). Mark of the Chosen and Uther's Strength proc only when struck, which never happens to a raid caster; Mark of the Chosen 590.0 against an empty slot 589.8 (noise). **Kindling Stave's Ignition proc (1300781) is not modelled** (Fire horde raid main hand), a known unmodelled-proc class from the 2026-10-07 relics review. | Aligned; Kindling Stave noted | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | **From level 40 to 49 an Arcane mage cast Arcane Blast rank 2, not rank 3.** | spellranks.json lists rank 3 twice. The NPC copy 42896 (level 76) comes first and the player's 1239697 (SpellLevels 148351: level 40, max 48; SkillLineAbility trainer row) second. The rank chain trusted the first row's level, so rank 3 counted as learned at 76. The rotation rewrite then cast 1239696 (78 base points against 101). A row above the level cap no longer sets a rank's learn level when a sibling within the cap exists (`sim/internal/spellranks/spellranks.go`, `buildRankChain`). Only this rank has the problem: a scan of every class's table finds no other above-cap duplicate. Test first: `TestARankAboveTheLevelCapDoesNotSetTheTiersLevel` (level 40 resolved to 1239696, then to 1239697). Band 40: +1.4%. | CONFIRMED-BUG (client SpellLevels plus SkillLineAbility) | S dbd2f6449 |
| A2 | Checked and aligned: Arcane Missiles ranks 5/6/7, Missile Barrage, Arcane Concentration, Arcane Impact, Arcane Mind, Arcane Instability, Arcane Meditation (50% at 3/3), Improved Channeling (pushback only), Wand Specialization, gnome and orc racials. | Raid breakdown: Arcane Missiles 46%, Frostbolt 29%, Arcane Blast 25%. | Aligned | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Checked and aligned: Scorch, Fire Blast, Pyroblast and Fireball ranks by band. Improved Fire Blast is -1 s a rank (11078 text "by 2 sec" at 2/2). Incineration covers Fire Blast, Ice Lance, Arcane Blast and Scorch. Accelerant covers Fireball, Frostfire Bolt and Arcane Wrath. Heating Up, Ignite, Critical Mass, Fire Power, Combustion. Scorch as the filler (APL notes, measured). | Raid breakdown: Scorch 33%, Pyroblast 30%, Fire Blast 22%, Ignite 14%. | Aligned | - |
| F2 | Kindling Stave's Ignition (1300781: harmful casts have a chance to deal 70 ±30% Fire) is unmodelled. The client states ProcChance 100 with a 10 ms recovery, but the text says "a chance", so the rate is server-side. | ItemEffect for 11750. design/reviews/2026-10-07-relics-procs.md lists it among the 59 unmodelled procs. | OPEN (rate unknown; undercounts horde raid Fire slightly) | - |

## Mage: Frost

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Checked and aligned: Frostbolt ranks 4/5/7/9/10, Ice Lance ranks (20/28/34/42/48/56), Ice Barrier ranks 11426/13031/13033 at 40/50/60, Fingers of Frost, Shatter, Winter's Chill, Frost Channeling, Ice Shards, troll Berserking. Ice Lance fires only on Fingers of Frost, so band 30 (no Fingers of Frost yet) never casts it, which is correct. | Raid breakdown: Frostbolt 68%, Arcane Missiles 16%, Ice Lance 16%. | Aligned | - |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Checked and aligned: Mind Flay, Mind Blast, Shadow Word: Pain, Devouring Plague and Shadow Word: Death ranks by band. Devouring Plague is baseline for every race (SkillLineAbility RaceMasks -1). Shadow Weaving (I8). Shadowform, Darkness, Twin Disciplines, Mental Agility, Inner Focus. Undead Touch of the Grave and gnome racials. | Raid breakdown: Mind Flay 41%, Shadow Word: Pain 19%, Mind Blast 19%, Devouring Plague 12%, Shadow Word: Death 8%. | Aligned | - |
| P2 | The band-20 build skips Mind Flay (Shadow Focus 5, Blackout 4, Spirit Tap 2). | Taking Mind Flay as the 11th point measures 36.1 against 38.6 DPS (horde): the band is out of mana for 80 s, and Mind Flay costs more than the wand. | Aligned | - |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Innervate was cast below its learn level.** Every bare band 20 and 30 Balance (and Restoration) druid got a free Innervate a fight. | SkillLineAbility teaches 29166 at SpellLevels 40 (`InnervateLevel`). `registerInnervateCD` registered it and added it as a mana cooldown at every level; its `RequiredLevel` is never enforced by core. The level scan found nothing else autocast below its level in this group. Test first: `TestInnervateIsLearnedAtForty` (registered at 30 and 39, then not). The druid conformance golden loses its Innervate rows at levels 10 to 38. The rotation rewrite and the ladder learn the same level (`singleTierLevelOverrides`, `idLearnLevelOverrides`). | CONFIRMED-BUG (client SkillLineAbility plus SpellLevels) | E 83e720c41, E d84ea3ead, S 0143f30ab |
| B2 | **The Starfire filler left bare-band druids out of mana for half the fight.** 82 s of 180 at band 20 and 94 s at band 30 once B1 lands; 39 s at 40 and 50 s at 50 before it. | Improved Wrath 5/5 (16814: -0.5 s and -50% mana) makes Wrath a 1.5 s cast at half the cost, and every bare build takes it. The rotation now casts Wrath while the mana pool is behind the fight's pace and Innervate is cooling down. Measured against pace-only and fixed thresholds: only this gate leaves raid 60 unchanged (567.3 against 567.3, 5,000 iterations). Test first: `TestBareBalanceDruidsDoNotSpendHalfTheFightOutOfMana` (bands 20 and 30 failed, then passed). The fork's copy is synced. | CONFIRMED-BUG (rotation; measured) | S 0143f30ab, E c6b747dd8 |
| B3 | Checked and aligned: Starfire, Wrath, Moonfire and Insect Swarm ranks by band (Forever's Starfire levels 50 and 58). Nature's Grace (16886: 10% cast speed, and the GCD 10% shorter for 3 s; the engine does not haste the GCD, so nothing is counted twice). Naturalist (+10% all damage). Eclipse, Moonglow, Improved Moonfire, Vengeance, Moonfury, Genesis, Moonkin Aura once (I1). | client rows, `druid/talents.go` | Aligned | - |
| B4 | Carried: Moonglow and Improved Wrath's cost cuts add (round 7 B2); Starfire rank 7 (25298) is a trainer row at 60 that IncludeAQ keeps out (round 7 B3). | | OPEN | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | **Shaman Spirit regen was 15 + Spirit/5 a 2 s tick; the Classic shaman rate is 17 + Spirit/5.** That is 1 mana a second more out of the five-second rule, and half that while casting with Mindfulness 3/3. | cmangos-classic `Unit::OCTRegenMPPerSpirit` ("case CLASS_SHAMAN: addvalue = (Spirit / 5 + 17)"; druid, hunter, paladin and warlock get 15). mangoszero `PlayerStats.cpp` and AscEmu `PlayerStats.cpp` say the same, and so does the Classic addon BetterCharacterStats. Forever's client carries no regen table to say otherwise. Test first: `TestShamanSpiritRegenIsSeventeenPlusAFifth` (29.89 against 30.89 at 224 Spirit), then pass. TestElemental and TestEnhancement goldens move with it. | CONFIRMED-BUG (emulators plus addon; client silent) | E 13746e744 |
| E2 | Checked and aligned: Lightning Bolt, Flame Shock, Lava Burst, Fire Nova, Searing Totem and Mana Spring ranks by band. Lightning Overload (3/7/10%, half damage and half coefficient). Elemental Alacrity, Convection, Concussion, Call of Flame, Elemental Fury, Elemental Mastery, Tidal Focus (+5% hit), Mindfulness. Mana Spring Totem lasts 300 s (DurationIndex 5), so one drop covers the fight. | client rows, `shaman/talents.go`, `lava_burst.go`, `fire_nova.go` | Aligned | - |
| E3 | Chain Lightning on cooldown at raid 60 is worse: 440.6 against 473.2 DPS (out of mana for 11 s). | rotation probe | Aligned (APL right) | - |
| E4 | Every bare band runs dry for 90 s or more, and a shaman has no wand. A rank-7 Lightning Bolt while behind the mana pace measured +14% at band 50 (108.4 to 123.4). The site always casts the highest learned rank, so no published rotation downranks. | probe `ele50-down.json` in the lane scratch dir | OPEN (owner policy: should bare bands model downranking? Mages and priests would gain too) | - |
| E5 | The bare-60 build is the leveling guide build: Improved Healing Wave, Totemic Focus and Natural Grace, and no Elemental Mastery. The raid build under the bare preset is +10.1%. The adoption rule holds it, because the winners drop unmodelled guide talents (design/reviews/talent-search/shaman-elemental.md). Bare 60 is not a published band. | measured, seeds 1 and 50001 | Noted (adoption rule) | - |
| E6 | The Stoneskin Totem aura is rank 6 (10408) at every band. | `core/buffs.go:561`. Defensive only; no DPS effect. | Noted | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | Core never enforces `Spell.RequiredLevel`, so a spell an ability file registers without its own level check exists at every level. The scan (every conformance preset at levels 10 to 50) found, besides Innervate: warrior 2458 (30), 1680 (36), 20569 (60), 871 (28), 2565 (16), 12323 (20), 1160, 6572, 19639 and 24407; rogue Vanish 1856 (22) and Feint 1966 (16); warlock Felhunter 691 (30) and Succubus 712 (20); shaman Healing Stream, Magma and Tremor ranks above level; mage Evocation (20) at 10. At bands 20 and 30 none of these fired in the published rotations of warrior, rogue, warlock, feral or enhancement (500-iteration check). | lane scan (test deleted after use); `levelscan.txt` in the lane scratch dir | Routed (no published number moves) | - |
| X2 | The elemental raid preset lists Major Mana Potion twice (the caster role and the shaman class row). The engine has one potion slot, so nothing is counted twice. | `data/curated/presets.json:161, 348` | Noted | - |

## Before/after

Before: pin `cd4dfdcc0` on origin/main `170b7b84c`. After: lane engine `d84ea3ead` on site `0143f30ab`, committed simdb, 10,000 iterations. Changes are seed 1 / seed 50001; absolute DPS is horde at seed 1. Every row not listed is within 0.15% on both seeds and both factions.

| Spec, band | before | after | change, horde | change, alliance | causes |
|---|---|---|---|---|---|
| Arcane bare 40 | 153.3 | 155.5 | +1.44 / +1.43% | +1.46 / +1.46% | A1 |
| Balance raid 60 | 567.3 | 567.0 | -0.05 / +0.0% | -0.11 / -0.18% | B2 (within error) |
| Balance bare 60 | 265.4 | 289.2 | +9.0 / +9.0% | +10.0 / +10.0% | B2 |
| Balance bare 50 | 156.7 | 196.9 | +25.7 / +25.7% | +26.8 / +26.8% | B2 |
| Balance bare 40 | 127.3 | 156.1 | +22.6 / +22.7% | +20.9 / +20.8% | B2 |
| Balance bare 30 | 61.3 | 65.5 | +6.9 / +6.9% | +6.5 / +6.5% | B1 (-24%), B2 |
| Balance bare 20 | 31.2 | 27.6 | -11.5 / -11.8% | -10.9 / -10.9% | B1 (-29%), B2 |
| Elemental raid 60 | 473.0 | 473.2 | +0.04% | +0.06% | E1 |
| Elemental bare 60 | 142.6 | 143.8 | +0.8 / +0.9% | +0.9 / +0.8% | E1 |
| Elemental bare 50 | 107.2 | 108.3 | +1.0 / +0.9% | +1.4 / +1.5% | E1 |
| Elemental bare 40 | 82.5 | 83.2 | +0.9 / +0.6% | +0.6 / +0.7% | E1 |
| Elemental bare 30 | 45.8 | 46.2 | +0.9 / +1.1% | +1.1 / +1.1% | E1 |
| Elemental bare 20 | 22.1 | 22.5 | +1.8 / +1.8% | +2.2 / +1.8% | E1 |
| Fire, Frost, Shadow, all bands | | | 0.0% | 0.0% | - |

Outside the group (horde, seed 1, 10,000 iterations):
- **Restoration druid:** bare 30 116.9 to 98.4 HPS (-15.8%); bare 20 51.8 to 40.8 (-21.2%). Cause: B1.
- **Restoration shaman:** raid 820.7 to 823.7 (+0.4%); bare 50 +0.6%; bare 30 +1.2%; bare 20 +1.5%. Cause: E1.
- **Enhancement:** within 0.6%.
- **Feral:** unchanged.

## For other lanes

- **Healers lane:** B1 removes Restoration druid's Innervate below 40, which costs 16 to 21% HPS at bare 20 and 30. E1 lifts Restoration shaman by 0.4 to 1.5%. Restoration's rotation already casts Innervate; the rewrite now drops that line below 40, the same as the engine.
- **Physical and warlock lanes:** X1's list of spells registered above their learn level. None fired in a published rotation at bands 20 and 30. A cooldown the APL autocasts would, so the list is worth a look before a rotation adds one.
- **Physical lane:** spellranks.json lists Mutilate rank 1 as 1329 (level 40) before 1310707 (level 30). The rank chain trusts the first row, so the rewrite treats Mutilate as learned at 40. Check which id the engine registers at 30.
- **Integrator:** sim/request TestRotationLadder goldens move for druid-balance (B2), druid-restoration and druid-feral (B1), mage-arcane (A1), and every shaman spec (E1). I did not regenerate them.

## Still open

- E4: downranking on bare bands (owner policy).
- F2: Kindling Stave's Ignition rate.
- B4: Wrath's cost stacking; Starfire rank 7 at trainers.
- Carried, blocked on level-60 logs: Fire Power on the burns, a Fire wand crit and Ignite, the wand hit table, Ice Lance's coefficient, Ignite's 0.73, Missile Barrage on a miss, Arcane Concentration per missile, Master of Elements on Clearcasting, Shadowfiend, Searing Totem's 2.2 s attack, Fire Nova's range.

## Sources

- The 70707 client tables with the hotfix push layered in (fetched 2026-11-01 21:42 UTC): SkillLineAbility, SpellLevels, SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellClassOptions, SpellDuration, ItemSparse, ItemEffect, ItemXItemEffect, ItemSetSpell.
- cmangos mangos-classic `src/game/Entities/Unit.cpp` `Unit::OCTRegenMPPerSpirit` (raw.githubusercontent.com, read 2026-11-01); GitHub code search (2026-11-01) for "Spirit / 5 + 17": mangoszero/server `PlayerStats.cpp`, AscEmu `PlayerStats.cpp`, MarcelineVQ/core `Unit.cpp`, BetterCharacterStats `helper.lua`.
