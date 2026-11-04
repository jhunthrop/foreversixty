# Sim audit round 14: hybrids (2026-11-04)

Specs: Feral cat, Retribution, Enhancement. Bear and Protection are lane tanks'.

**Setup:**
- Client 1.60.1.70753 with the 2026-11-03 14:14 UTC push layered. Raw tables: `scratchpad/r11/c70753/builds/1.60.1.70753/raw`.
- Engine `lane/r14-hybrids` = fork int/r14 (5195a8889), unchanged. Site `wt/r14-hybrids` on int/r14 (72471fcb3).
- 10,000 iterations at seeds 1 and 50001. Before = `$R/before.tsv` (raid 60) and `$R/sb-before` (bands).
- Nothing nightly-owned was regenerated or committed.

**In short: 1 CONFIRMED-BUG, fixed (landing 3e3a14ff9).** The assigned item's premise (the Classic rule) does not hold in Forever: the engine is right and the landing's tool workaround was wrong.

## Assigned item: does weapon DPS reach Cat Form damage?

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Classic rule.** In Classic a feral form ignores the weapon's damage: the paw rolls `level x 0.85` to `level x 1.25` per second of swing (level capped at 60), and only the weapon's stats count. | cmangos mangos-classic `src/game/Entities/StatSystem.cpp:333-338` (`IsAttackSpeedOverridenShapeShift`; file's last commit 9c4fb5556, 2026-08-03) | Background | - | - |
| W2 | **Forever changed it.** Blizzard's druid deep dive: in Bear, Cat or Dire Bear Form the auto attack's DPS "is now the same as the DPS of the Druid's equipped weapon", the speed becomes 1.0 or 2.5 s, and "abilities that deal weapon damage likewise use these damage values". Weapon procs and weightstones work in forms. Round 1's three guides (Icy Veins, Mobalytics, SixtySim) agree. | us.forums.blizzard.com/t/2367361 (2026-09-30, re-read 2026-11-03); round 1 hybrids.md:68 | FOREVER-CHANGE (two sources) | - | - |
| W3 | **The client has no column for it.** Cat Form 768 carries the shapeshift (36), immunities (77), a 5 s periodic (23) and a new dummy (aura 4, row 1324553). That dummy is on every form, Travel and Aquatic too, so it is the consumables rule, not the weapon rule. Cat Form (Passive) 3025 gives 12 + 2 a level attack power from level 6 (120 at 60, as the engine's `2 x level`), Agility as attack power (aura 598, misc 1, 100%), threat -29%, and Faerie Fire free on a 6 s cooldown with a 1 s global. Wowhead's "12 plus Agility" is the base level's figure. | SpellEffect 679878-1324553, 680786-1359783; SpellLevels 98115; nether.wowhead.com/forever/tooltip/spell/768 (2026-11-03) | Aligned | - | - |
| W4 | **Where weapon damage reaches the cat (all intended):**<br>- The auto attack: `formPawWeapon` (`sim/druid/forms.go:50`) keeps the weapon's DPS and spread at a 1.0 s swing, normalized speed 1.0. The sweep's +1 weapon DPS reaches it through `BonusMHDps` (`core/attack.go:73`).<br>- Shred (155% of the paw hit plus 80), Claw (110% plus 115), Ravage and Windfury Totem's extra swing use the paw through `MHWeaponDamage`.<br>- Rake, Rip and Ferocious Bite are flat (client and Wowhead text).<br>- Nothing hidden: Omen's rate reads the paw's 1.0 s, so a slow weapon does not raise it.<br>The sweep's 2.3 DPS per weapon DPS at raid 60 matches the breakdown: white ~1.0, Windfury Totem ~0.3, Shred ~0.8 (0.39 a second x 1.55 x crit). Bands read 0.8 to 1.1 (no Windfury, fewer Shreds). | `forms.go`, `shred.go:76`, `claw.go:65`, `ravage.go:70`; Wowhead 9830/9850/9904; raid breakdown | Aligned (engine right) | - | - |
| W5 | **The landing's dps-per-point workaround zeroed a real term.**<br>**What was wrong:** 3e3a14ff9 added `formIgnoresWeapon` for druid-feral on the Classic rule, so the addon prices a weapon's DPS at nothing for a cat. That contradicts the engine, the guides ("Weapon damage per second matters in Forever", feral.md:92; druid index.md:27) and the BiS ranker, which picks cat weapons on DPS. The addon test that went red ("ignores a feral druid's weapon DPS", 9264bdd1a, 2026-10-12) was written before the engine fix (E 4cb26aa16, 2026-10-31).<br>**What changed:** `formIgnoresWeapon` is gone; the cat keeps its measured term. The real-data addon test becomes a fixture test that a cat prices weapon DPS on the main hand. The stale addon spec line (2026-10-11-addon-v2-design.md) is corrected.<br>**Should it stay as belt-and-braces?** No: it is wrong, not redundant. | Test first: `TestTheCatKeepsItsMeasuredWeaponTerm` (sim/cmd/dps-per-point) failed ("WeaponDPS:0, want the measured 1.1") and passes; addon `item_tooltip_spec.lua` 212/0 under busted; luacheck clean | CONFIRMED-BUG | landing 3e3a14ff9 (on an old addon test, 9264bdd1a) | S ffdb90ce4 |
| W6 | Unarmed paw: the engine falls back to 54.8 DPS (the fork's inherited figure), where Classic's formula gives 63 at 60. No published set is unarmed. | `forms.go:41`; W1 | OPEN (latent) | - | - |

## Feral cat (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Raid build read against the client text, talent by talent: Genesis, Nature's Majesty, Ferocity (Claw/Rake only, a path point at 60), Feral Swiftness and Feral Instinct (path points), Shredding Attacks 18, Savage Fury 10%, Sharpened Claws 6%, Predatory Strikes 150% of level, Predatory Discipline, Blood Frenzy 1/2 (50% extra combo point on builder crits), Leader of the Pack, Natural Instinct (mask resolves to Claw, Bite, Lacerate, Maul, Pounce, Primal Bite, Rake, Ravage, Rip, Shred, Swipe; no white swings), Rend and Tear (abilities only, never bleed ticks), Shifting Power, Naturalist 10%, Nature's Focus and Subtlety (no damage). | talents/druid.json; SpellEffect 1223242, 1223246, 16958; `feral_talents.go:84-115` | Aligned | - | - |
| C2 | Omen of Clarity: client row 16864 is ProcChance 100 with a 10 s recovery. The engine's 2 a minute on the paw plus 10 s is the logged rate (round 10: 346 gains against ~330 predicted). | SpellAuraOptions 129438; `talents.go:315-334` | Aligned | - | - |
| C3 | Faerie Fire 770 cast in Cat Form at band 20 is right in Forever: Cat Form (Passive) makes Faerie Fire free with a 6 s cooldown and a 1 s global (mask 1024), and the form stays up. | 3025 rows 1231426-8; band 20 breakdown | Aligned | - | - |
| C4 | Idol of the Dream (220606): Rip +2 s (446212), one extra tick. Shadowcraft 5/6 on the published set: +8 resist, +30 attack power, Crafted Shadows (utility), 20 Energy on hit (27787/27788; the client's 100% is not the server rate). | ItemEffect 184193; ItemSetSpell 2417-7912; `client_sets_gen.go:26,406` | Aligned (set rate is lane data's) | - | - |
| C5 | Tiger's Fury not learnable, Shifting Power the Energy source, no Rake in the rotation (61 + 102 over 9 s for 35 against Shred's ~700 for 42), Shred and Claw learn levels, Ravage opener from Prowl. | rounds 12/13 C5; Wowhead 9904 | Aligned | - | - |
| C6 | Bands 20 to 60 bare breakdowns: no unresolved cast, no idle form. | sb bare 20/40 | Aligned | - | - |

## Retribution (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Seal of Command: 70% weapon damage, 0.29 spell power, 1 s recovery, 7 a minute. In 70753 every rank triggers 20424 (Era had one proc spell a rank; 20947 is not in the client). The engine's 20947 label is cosmetic; the damage is identical. | SpellEffect 20375/20920/20424; Wowhead 20920 | Aligned (label only) | - | - |
| R2 | Unsealed Wrath matches its tooltip; the 33 SoC procs a fight come from stored copies under Righteousness. Its client mask (a 0% cost mod) holds the Seals and Judgement of Righteousness, which does nothing. | Wowhead 1310735; mask resolution | Aligned | - | - |
| R3 | Improved Seals 15%: client mask resolves to the Seal procs and Judgements of Command, Fury and Righteousness, exactly the engine's mask. | 20224; `paladin.go:83` | Aligned | - | - |
| R4 | Two-Handed Weapon Specialization: aura 79 school 1, Physical only (Holy Strike and the Seals do not get it), threat -20%. Vengeance: aura 79 school 3 (Physical and Holy), 3% a stack at 3/3, 3 stacks, 30 s. Champion of the Light: spell damage from Intellect, Hammer of Wrath -1.0 s (instant). Conviction, Precision, Benediction, Holy Conduit, Sanctified Judgement as their text. | SpellEffect 20111, 20049/20050, 1311084; `talents.go:85-260` | Aligned | - | - |
| R5 | Holy Strike: 50% of (normalized weapon + 93 at 60), 10 s, 20 mana; the flat is halved by effect order (121 before 31), which is how the server applies a percent effect to earlier flat bonuses. Wowhead's 97-121 is the rank's max-level roll (93 + 3.2 x 5). Consecration 312 over 8 s single target (96 + 216 first 4), paced. Hammer of Wrath 6 s. | SpellEffect 10333; Wowhead 10333/20924/24239 | Aligned | - | - |
| R6 | Libram of Law (272435): +4% Judgement damage, mask resolves to Judgement of Righteousness and Fury only (not Command). Worth +0.1% at raid 60 (700.9 to 701.6, 3,000 iterations), so the published ranged slot stays empty ("no_dps_value"). | relic_mods_auto_gen.go:36; SpellEffect 1291086 | Aligned | - | - |

## Enhancement (re-audit)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **Maelstrom Weapon's rate.** The engine reads effect 1's 50 as "50% a melee hit". 408498 is Season of Discovery's rune spell, reused by the Forever talent. In SoD its description binds that 50 to "chance increased by $s2% with Windfury Weapon and another $s2% with a two-hander"; the base rate is 10 a minute (wowsims/sod `sim/shaman/runes.go`, last commit d966c41c3, 2025-04-07). Forever's description drops the bonus sentence and keeps the dummy.<br>**One beta log** (report 2748, character 129323, Rage of the Storm 3.2 s two-hander, Windfury Weapon): 51 stack gains over 94 landed melee hits, 54% (95% range about 44 to 64%). That rules out SoD's bonuses (about 100%), but cannot tell 50% flat from 10 a minute (53% at 3.2 s).<br>**At stake:** 10 a minute on the published 3.3 s axe is +0.14 / +0.09% at raid 60 and +0.35% at bare 40 (horde, seed 1 / 50001). | Spell.csv 408498 for 1.15.9.70576 and 1.60.1.70753 (wago, 2026-11-03); Wowhead classic and forever 408498; `talents.go:782-797` | OPEN (settle: a log of a fast one-hander shaman with Maelstrom) | - | - |
| N2 | Totem of the Storm (272432) is the Forever relic: Lightning Bolt can grant Maelstrom at half the chance. Modelled (`items.go:134`). Its published delta of -5.8 is ranker noise. | ItemEffect 232245; spell 1291078 | Aligned | - | - |
| N3 | Taken talents against client text: Concussion 5%, Elemental Devastation 9%, Reverberation 3/5 (-0.6 s; Earth Shock 31.4 casts of a possible 33), Elemental Focus 10% (4.45 Clearcastings over about 45 casts), Thundering Strikes 5%, Ancestral Knowledge 6%, Mental Dexterity, Elemental Weapons, Shamanistic Focus 45%, Flurry, Stormstrike, Mental Quickness, Rage of the Farseer. Earth's Grasp, Tidal Focus and Mindfulness are path or utility points. | talents/shaman.json; raid breakdown | Aligned | - | - |
| N4 | No dual wield: Forest Defender's Axe, two-handed. The shaman's own Windfury Weapon is not doubled by the raid's Windfury Totem (no totem aura on the shaman). Fire Nova comes from Clearcasting; the beta logs show the same spell family (408342/408424). | bis json; breakdown; logs | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | wago lists a newer Forever build, **1.60.1.70755**, above the active 70753. The client watcher should pick it up; nothing here was measured on it. | wago.tools/api/builds (2026-11-03) | OPEN (new-data watch) | - | - |

## Before/after

No engine or sim change, so the numbers do not move. 10,000 iterations, raid 60:

| spec | horde s1 | horde s50001 | alliance s1 | alliance s50001 | horde s1 absolute |
|---|---|---|---|---|---|
| druid-feral | 0.00% | 0.00% | 0.00% | 0.00% | 660.4 -> 660.4 |
| paladin-retribution | 0.00% | 0.00% | 0.00% | 0.00% | 701.0 -> 701.0 |
| shaman-enhancement | 0.00% | 0.00% | 0.00% | 0.00% | 587.0 -> 587.0 |

Cat bare bands, horde: 20, 30, 40, 50 and 60 are identical at both seeds (45.0, 83.3, 115.5, 153.6, 266.8 at seed 1). W5 changes what the addon's tooltip prices, not what the engine deals.

## Files to regenerate

- `data/builds/1.60.1.70753/dps-per-point.json` for druid-feral (spec and bands), then `addon-data.json` and `addon/ForeverSixty/Data.lua`: the cat's `weaponDps` returns to its measured values (about 0.78 to 1.15 at bands, 2.33 at 60 on the last run).

## Measure at refresh

- None from this lane. The cat's BiS weapon picks already price weapon DPS (the ranker uses the engine), so the refresh does not move on W5.

## For other lanes

- **Platform (dps-per-point tool owner):** W5 edits `sim/cmd/dps-per-point/main.go` and its test, plus the addon spec, as this lane's assigned item. Please regenerate dps-per-point and Data.lua with the item list above.
- **Data (set bonuses):** Shadowcraft's 20 Energy proc (27787) has client ProcChance 100; the engine's rate is hand-written. Not re-checked here.
- **Coordinator:** X1, the 70755 build on wago.

## Still open

1. N1: Maelstrom Weapon's rate, 50% a hit or 10 a minute. Under 0.4% at any band; needs a fast-weapon log.
2. W6: unarmed paw figure (latent).
3. X1: build 70755.
4. Round 13's opens carry over unchanged: the band 20 Clearcasting Shred line, held under the 1% bar.
