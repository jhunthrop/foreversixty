# Sim audit, round 2: physical group (2026-11-01)

Warrior (Arms, Fury, Protection), rogue (Assassination, Combat, Subtlety) and hunter (Beast Mastery, Marksmanship, Survival; Survival is melee). Client 1.60.1.70707. The baseline is the pinned engine 17d31ecfa. Fixes are on fork branch `lane/r2-physical` (E) and site branch `wt/r2-physical` (S). Unless a row says otherwise, measurements are `sim/cmd/spec-breakdown` at the level-60 raid preset, alliance, 10000 iterations at seeds 1 and 50001, as relative %. Every error is under ±0.1%.

Raw client tables (SpellEffect, SpellMisc, SpellCooldowns, SpellCategories, SkillLineAbility, ItemSparse, ItemEffect) were pulled from wago.tools for 70658 and 70707 and diffed. The scratch copies are in `scratchpad/r2/physical/research/dl/`. Forever tooltips come from `https://nether.wowhead.com/forever/tooltip/spell/<id>`. Those tooltips render the client's own text, so a tooltip plus the client table counts as one source, not two.

Status values: CONFIRMED-BUG means our sim was wrong. FOREVER-CHANGE means we were right and the web source is Era. OPEN means the item is not settled yet.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Pre-pull damage counted for nothing.** The engine records no damage before 0 (`sim/core/spell_result.go:425`). A pre-pull Aimed Shot at -3.5 s with a 2 s cast landed at about -1.5 s. Its damage was lost, but it still started the shared 6 s cooldown. Feral's Ravage at -1 s had the same problem. **Model:** a hit on the boss is what starts the fight, so a damaging pre-pull cast is timed to land at 0. The engine now holds any pre-pull cast that targets an enemy and deals direct or spell damage until it lands at 0, using the cast time at that moment plus missile travel. Buffs, marks, summons and heals keep their time. | `sim/core/apl.go` `prepullLandingHold`; test `sim/hunter/prepull_landing_test.go` failed first (0 damage over 43 hits). The fork `go test ./sim/...` goldens did not move. | CONFIRMED-BUG | E c880120ea |
| X2 | Effect of X1 on every DPS spec (alliance raid preset). | Feral +0.59/+0.50%, MM +0.57/+0.57%, BM +0.37/+0.35%. Unchanged: Balance, Bear, SV, Arcane, Fire, Frost, Ret, Prot paladin, Shadow, the rogues, Elemental (its Lightning Bolt at -2.5 s already lands at 0), Enhancement, the warlocks (Destruction's Immolate at -1.5 s already lands at 0) and the warriors. | measured | - |
| X3 | **Owner crit rule side effects.** No new bug. Details: Deadly Poison ticks crit on the spell table with the spell multiplier (1.5x). Their crit rate equals Instant Poison's landed crit rate: about 42% vs 42.4% for Combat, because Instant Poison's 12.7% miss is in the breakdown's denominator. Malice reaches both through the shared crit stat (`sim/core/spell_result.go:224-232`). Rend 35%, Deep Wounds (row "spell (12867)") 35%, Rupture 31%, Lacerate 29%, Lacerating Strikes 29%, Immolation Trap 34% and Serpent Sting 27% all crit. Volley is a channel with no PERIODIC_CAN_CRIT flag (SpellMisc 70707), so it ticks flat. No proc reads periodic crits: the only `OnPeriodicDamageDealt` handlers in warrior, rogue and hunter are Thousand Cuts and a TAQ set bonus, and no `CallbackOnPeriodicDamageDealt` proc exists anywhere. So tick crits fire no Seal Fate, Flurry or Deep Wounds. Real-log check: in the owner's 17 Sep beta sample (`web/public/sample-report/sample/`), Rend ticks crit 5 of 81 and Serpent Sting 13 of 296. | spec-breakdown rows; `sim/core/dot.go:395-401`, `sim/core/spell_outcome.go:38-70` | aligned | - |
| X4 | **Lacerating Strikes double-dipped target modifiers** (raised by the integrator after the casters' Ignite fix). The bleed is 40% of the damage Mongoose Bite did, and that hit already took the target's damage-taken modifiers. The ticks applied them a second time. Deep Wounds is clean: it snapshots average weapon damage (`sim/warrior/deep_wounds.go:73-84`). | `sim/hunter/lacerating_strikes.go:76`; test failed first | CONFIRMED-BUG (0% at the raid preset, which carries no physical damage-taken debuff) | E 83814b032 |
| X5 | Blood Fury: 2 min cooldown since 70707 (SpellCooldowns 55612: 180000 to 120000 ms). The engine already had 2 min. StartRecoveryTime 0 means no global cooldown, but the engine charged the default. | client rows; `sim/core/racial_blood_fury.go:72` | CONFIRMED-BUG | the warlock lane's E f44ba933d carries it. Mine was dropped to avoid a duplicate. My measurement before dropping it, horde orc: Arcane +1.9%, Fire +1.8%, Elemental +1.1%, Arms +0.9%, Enhancement +0.4%. |
| X6 | Windfury Totem on hunter pets (warlock lane's question). Hunter pets call `RefusePlayerBuffs` (`sim/hunter/pet.go:110`), and `applyPetBuffEffects` returns before any buff (`sim/core/buffs.go:439`). | code; BM breakdown shows no Windfury aura | aligned (warlock pets are the warlock lane's) | - |
| X7 | Real logs. Warcraft Logs has no Forever partition. foreverlogs.gg exposes only per-spec DPS spreads from beta dungeons at levels 15 to 30. Its report and API routes need a key. Our public API has no ranked encounters, and `/v1/reports` needs a session. The owner's level 1-3 sample report is the one source with ability detail. Survival leads the hunter specs on foreverlogs' median (130 against 71 and 69), which fits the melee ruling. No level-60 band can be compared. | https://foreverlogs.gg/statistics?phase=1&location=All+Dungeons ; `web/public/sample-report/sample/` | OPEN (no level-60 data until launch) | - |

## Warrior: Fury

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Sunder Armor threat now scales with AP (see P2). Fury's Sunder line adds TPS but no DPS. | | see P2 | E 939e4eedc |
| F2 | Raging Blows (round 1 held it at +0.7%). The current search winner is Raging Blows 1 for Enrage 1: +0.52/+0.61% (±0.08). | `talent-search -preset raid -keep` | HOLD (under 1%) | - |
| F3 | Cooldown timing against the execute phase. Holding Death Wish and Recklessness for `isExecutePhase E20` measured +0.65/+0.69% at 180 s, and +7% on the bare 60 band, where Recklessness otherwise goes out at the pull. The same gate, written as "execute or the last 30/15 s", made the fork's 60 s `forever_fury` suite rows **lose 13-30%**. I could not explain that in this pass, so the change was reverted. | `/scratchpad/r2/physical/rot/fury-F1.json`; fork `TestP1DPSWarrior` ShortSingleTarget | OPEN (settle why the gated cooldowns lose in a 60 s fight before adopting) | - |
| F4 | Normalized rage. The beta sample's per-swing rage tracks weapon speed, not damage: about 6.5-8 at 1.5-2.0 s swings and about 13.5 at 3.0-3.8 s. A crit gave the same 13.5 as a hit in one sample. The crit-rage factor (2.0 in the engine) is untested. | sample report events | aligned on shape; crit factor OPEN | - |
| F5 | Round-1 fixes on 70707. Engine tests pass for the Heroic Strike queue off-hand hit, Ruthless, Artisan's Precision and baseline Tactical Mastery. | `go test ./sim/warrior/...` | no regression | - |

## Warrior: Arms

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Round-1 Slam, Improved Slam and execute Recklessness checked on 70707. Slam casts 7.8 times and Recklessness 1.0 times a fight. Deep Wounds and Rend tick and crit. | breakdown | no regression | - |
| A2 | Talent search: the guide is the best build found. Deep Fury loses 3.0%. | talent-search | aligned | - |
| A3 | Hamstring is cast about 40 times a fight as a cheap filler, worth 1.7% of damage. That is a rotation choice already searched in round 1. | breakdown | aligned | - |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Revenge ignores AP.** Client 70707 rewrote every rank's description to `$m1+($<APRatio>*$apm)` with APRatio 0.25. The Forever tooltips agree: 25288 is 137.7-168.3 + 0.25 x AP and 11601 is 109.07-132.93 + 0.25 x AP. wowfdb logged it on 30 Oct. SpellEffect still carries an AP coefficient of 0, so the engine adds it by hand. | `sim/warrior/revenge.go` `revengeBaseDamage`; test `attack_power_scaling_test.go` failed first | CONFIRMED-BUG | E 939e4eedc |
| P2 | **Sunder Armor threat ignores AP.** Blizzard's 24 Sep notes say they "corrected the threat values on all ranks, including a small increase to threat generated from Attack Power" (via kami-labs); wowfdb has the same 24 Sep entry. sixtysim measured the size: 206 + 5% of AP. The client still shows flat 206 (11597 effect 63, AP coefficient 0): this is a server change. Two sources confirm the term exists. Only sixtysim gives the 5%. | `sim/warrior/sunder_armor.go` `sunderArmorFlatThreat`; same test file | CONFIRMED-BUG (coefficient single-source) | E 939e4eedc |
| P3 | P1 and P2 together, tank harness: own damage +24.1% (both seeds), threat/s 771.5 to 1037.0 (+34.4%), TMI unchanged at 39.1. Fork goldens: TPS rows of the DPS and tank warrior suites, plus Revenge damage. | `spec-breakdown -tank` | measured | - |
| P4 | Revenge rank 6 (25288) has a SkillLineAbility row (13218, AcquireMethod 0). Its book status comes from `data/curated/book-ranks.json`, not from the client. | client SkillLineAbility 70707 | OPEN (same as round-1 F11: the launch trainer list settles it) | - |

## Rogue: Assassination

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **The published 70707 raid set sims 253 DPS from white swings and poisons alone.** The engine refuses Mutilate without a dagger in each hand (round-1 E 56f97823e). The ranker's dagger table (`sim/cmd/leveling-bis/weapon_requirements.go`) held Backstab and Ambush but not Mutilate, so it kept Greenhammer and Ironwood Blade, a mace and a sword. Every pass then optimised a rotation that never fired. Mutilate's four ranks (1310707, 399956, 1241582, 1241584) now restrict both hands to daggers. The same band on Heartseeker and Scarlet Kris sims 576.2 (+128%). **A bis.yml run is needed** to re-rank the set, because the rest of it was chosen on white-only weights. | test `TestAPLRotationRequiresDaggerTrueForMutilate` failed first | CONFIRMED-BUG | S a76dfad65 |
| R2 | Stones beside poisons (see Combat C1). On daggers: 576.2 to 591.6 (+2.7%). | | CONFIRMED-BUG | E aded25fc4, S d2107e734 |
| R3 | Murder: creature-type mask 80, which is Humanoid plus Giant; the tooltip gives 4%. The engine applies it to those two types only. The raid target defaults to Humanoid. | client 14158; `sim/rogue/talents.go:140-157` | aligned (model); target-type policy OPEN | - |
| R4 | Poison AP scaling: the AP coefficient is 0 on 11354, 25349, 11340 and 11337, no Forever tooltip mentions AP, and wowfdb: "A test found none". | client SpellEffect 70707 | aligned | - |
| R5 | Venom (1310703, two 30% modifiers) stacks with Vile Poisons (16513, two 20% modifiers). The engine now adds them (`sim/rogue/venom.go:34-37`). The client tables cannot show whether the server adds or multiplies. | client rows | OPEN (needs a log with both up) | - |
| R6 | Talent search after R1. The search reads the band's gear, which is still mace and sword. | | OPEN (re-run after the bis.yml re-rank) | - |

## Rogue: Combat

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **A stone and a poison could not share a weapon.** The Deep Dive (7 Oct) lets them, but `Consumes` held one imbue enum per hand, and the kit's poison took the slot of the preset's stone. The fork adds `Consumes.main_hand_stone` and `off_hand_stone`, which hold a stone, weightstone or oil on a hand whose imbue is a poison or empty. Anything else is refused at build time. The site's `ResolvedPreset.Layer` moves a preset stone on a poisoned hand into that field. Measured: Combat +1.32%, Subtlety +1.86%, Assassination +2.45% (on its dead set). | tests `TestDenseStoneAppliesBesideAPoison` and `TestStoneFieldRefusesAPoison` (fork), `TestLayerPutsThePresetStoneBesideARoguePoison` (site), all failed first; `IDS.md` regenerated | CONFIRMED-BUG | E aded25fc4, S d2107e734 |
| C2 | Slice and Dice is still cast only while `currentTime < 15s` (8.8% uptime). Dropping the gate gives 77% uptime at +0.12/+0.13%. SnD-first variants at 83-95% uptime measured -0.05 to +0.07%. The model values SnD at about zero, because Forever's Eviscerate takes what white damage gains. The 15 s gate protects the low-level ladder rungs (APL notes). | `scratchpad/r2/physical/rot/combat-{A,B,C,D}.json` | aligned in DPS (within error); rotation reads unlike public TC. HOLD | - |
| C3 | Talent search: the guide is the best build found. Deep Assassination loses 3.5%. | talent-search | aligned | - |
| C4 | Second Blade Flurry (round 1): 2.00 casts a fight. | breakdown | no regression | - |

## Rogue: Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **Raid build adoption.** The tool's swap generator never makes the two-talent Dirty Deeds to Quietus move, so the build was assembled by hand: Dirty Deeds 2, Quietus 5 and Improved Ambush 1, for Cutthroat 3 (procs only off Backstab, which the Hemorrhage rotation never casts), Dirty Tricks 2 (crowd control) and Setup 3 (defensive). It measured +1.28% at both seeds (537.5 to 544.4, 538.0 to 544.9, ±0.05). It keeps Camouflage, Master of Deception, Thousand Cuts and 31 Subtlety points, and drops nothing the APL casts, so it passes every clause of the adoption rule. Round 1's legal swap was +0.70%; the tool's own winner (Cutthroat into Improved Ambush) is +0.32%. | `ts-out/rogue-subtlety`; spec-breakdown `-talents 00532310101400000-00000000000000000-5320201310013211501` | CONFIRMED-BUG (build) | S 08a782025 |
| S2 | Stones beside poisons: +1.86%. | C1 | CONFIRMED-BUG | E aded25fc4, S d2107e734 |
| S3 | The pull Ambush is counted (avg 925, 0.96 hits). | breakdown | no regression | - |

## Hunter: Beast Mastery

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | The pre-pull Aimed Shot counted for nothing (X1). It now lands at 0. | X1 | CONFIRMED-BUG | E c880120ea |
| B2 | Aspect of the Hawk and Trueshot Aura in 70707. The lower ranks were cut to sit under the top ranks: Hawk 20/35/50/70/90 became 15/22/30/38/47, with 14322 = 55 and 25296 = 120 unchanged. Trueshot 30/40/50/75 became 15/22/30/40, with 20906 = 50 unchanged. The engine reads 70707 (`sim/hunter/constants_auto_gen.go:257`, `sim/core/buff_ranks.go:127-133`), and Trueshot gives ranged AP only. The Forever tooltips match. | SpellEffect diff 70658 to 70707 | aligned | - |
| B3 | Hydra Shot 1293020: no SkillLineAbility row in 70658 or 70707, and no spell teaches it. trainables lists it as `class_spell`, which is the pipeline's fallback for class spells. No web report puts it in a spellbook. | client tables | aligned (not learnable; stays out) | - |
| B4 | Pet family | | OPEN (no Forever raid-pet consensus found) | - |
| B6 | **Raid build adoption** on the lane engine. Raid search with -keep: Careful Aim 5 and Mortal Shots 2 in, Improved Stings 1, Efficiency 1 and Rapid Killing 1 out (`5420001505001251/00524510012/0`, 31/20/0). Measured +1.93/+1.90% (627.7 to 639.8, 627.4 to 639.3, ±0.05). Every clause holds: at least 1% and beyond error; every unmodeled guide talent (Bestial Swiftness, Intimidation) kept; Beast Mastery holds 31; the APL casts none of the removed talents' spells; Bestial Wrath kept. | `ts-out2/hunter-beast-mastery` | CONFIRMED-BUG (build) | S a81d2997e |
| B5 | Item 15464 is Brute Hammer and did not change. Spell 15464 ("Increased Hit Chance 01", the Black Dragon Mail 2-piece) gained a spell-hit effect in 70707. The engine and simdb already count it as one 1% hit (`sim/core/client_sets_gen.go:57,254`; `data/pipeline/simdb/equip.py:352-366`). | ItemSparse, SpellEffect diff | aligned | - |

## Hunter: Marksmanship

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| M1 | Pre-pull Aimed Shot (X1): +0.57%. | X1 | CONFIRMED-BUG | E c880120ea |
| M2 | Round-1 shared Aimed and Multi-Shot timer, Careful Aim and Trueshot Aura: engine tests pass. | `go test ./sim/hunter` | no regression | - |
| M3 | Talent search on the lane engine. Winner 7/31/13: Rapid Killing 2, Surefooted 3 and Deflection 5 (unmodeled filler) for Endurance Training 3, Unleashed Fury 5, Ferocity 1 and Improved Stings 1. Measured +0.74/+0.81% (±0.06). Also, spec-breakdown sims the guide build at 619.5 where the search's own baseline reads 647.0: the two tools may not wear the same band. | `ts-out2/hunter-marksmanship` | HOLD (under 1%); tool baseline gap OPEN | - |

## Hunter: Survival

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| V1 | Lacerating Strikes double dip (X4). | | CONFIRMED-BUG | E 83814b032 |
| V2 | Hunter DoT types. Lacerate (24118-24120, 1299332) and Lacerating Strikes (1310536) are physical bleeds (mechanic 15). Serpent Sting, Immolation Trap and Explosive Trap are non-channelled DoTs, and all crit. Wyvern Sting's DoT (24336) is not in the engine and the rotation does not use it. | SpellMisc and SpellCategories 70707 | aligned | - |
| V3 | Talent search. The winner swaps Improved Concussive Shot 2 for Rapid Killing 2: +0.58/+0.59%. It drops an unmodeled guide talent and is under 1%. | talent-search | HOLD | - |
| V4 | Round-1 Aspect of the Falcon (99.8% uptime) and Rapid Fire: present. | breakdown | no regression | - |

## Before/after (lane engine 83814b032 plus site branch, against the pinned 17d31ecfa)

| Spec | seed 1 | seed 50001 | what moved it |
|---|---|---|---|
| Arms | +0.00% | +0.00% | (Sunder TPS only) |
| Fury | +0.00% | +0.00% | (Sunder TPS only) |
| Protection (DPS) | +24.06% | +24.05% | Revenge AP |
| Protection (threat/s, tank harness) | +34.4% | +34.4% | Revenge AP, Sunder AP |
| Assassination, published set | +2.45% | +2.49% | stones (Mutilate still dead until bis.yml) |
| Assassination, on daggers vs published set | +128% | +128% | S a76dfad65 (ranker) |
| Assassination, both on daggers | +2.67% | +2.67% | stones |
| Combat | +1.32% | +1.32% | stones |
| Subtlety | +1.86% | +1.86% | stones |
| Subtlety with the new build | +3.16% | +3.16% | stones and the build |
| Beast Mastery | +0.37% | +0.35% | pre-pull |
| Beast Mastery with the new build | +2.3% | +2.3% | pre-pull and the build |
| Marksmanship | +0.57% | +0.57% | pre-pull |
| Survival | +0.00% | +0.00% | |
| Feral (outside group) | +0.59% | +0.50% | pre-pull Ravage |

## Still open

- **Fury cooldowns in the execute phase (F3).** +0.7% at 180 s, but a sharp loss in the fork's 60 s suite that needs explaining before adoption.
- **Assassination** needs a bis.yml re-rank (S a76dfad65), then a talent search on dagger gear. Venom against Vile Poisons is additive in the engine and unverified in-game.
- **Marksmanship:** talent-search's baseline (647.0) differs from spec-breakdown's (619.5) for the same guide build. Find out why before the next MM build change.
- **Sunder's 5% coefficient** is single-source (sixtysim). **Revenge rank 6** book status needs the launch trainer list.
- **Murder** target-type policy for non-humanoid bosses, the **pet family**, and the **normalized-rage crit factor** remain open.
- **Real level-60 logs** do not exist before launch.
- **TestRotationLadder goldens** move for BM, MM, Feral and Elemental (pre-pull) and for the stone change. Left for the integrator, as the brief says.
- The site's `TestLayerPutsThePresetStoneBesideARoguePoison` needs the engine pin that carries `Consumes.main_hand_stone`.
