# Sim audit, round 3: physical group (2026-11-01)

Warrior (Arms, Fury, Protection), rogue (Assassination, Combat, Subtlety) and hunter (Beast Mastery, Marksmanship, Survival; Survival is melee). Client 1.60.1.70707. Baseline: engine pin c8d29ea33. Fixes are on fork branch `lane/r3-physical` (E) and site branch `wt/r3-physical` (S). Unless a row says otherwise, measurements are `sim/cmd/spec-breakdown`, level-60 raid preset, 10000 iterations at seeds 1 and 50001, relative %. Errors are under ±0.1% unless given.

Status values: CONFIRMED-BUG means our sim was wrong. FOREVER-CHANGE means we are right and the web source is Era. OPEN says what would settle it.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **The two searches fought a different target from everything they are compared with.** talent-search and rotation-search simmed `api.DefaultEncounter()`. Its empty target type becomes the engine's Humanoid (`sim/request/request.go:530-536`). The BiS ranker, spec-breakdown and dps-per-point fight a target with no creature type. This is round 2's Marksmanship gap (M3, 647.0 against 619.5): Improved Tracking added 4.2% on the Humanoid target, and Murder added 4% to every rogue search. The one encounter is now `leveling.RankerEncounter()`, and all five tools use it. | tests `TestSearchFightsTheRankersTypeNeutralTarget` (talent-search) and `TestRotationSearchFightsTheRankersTypeNeutralTarget` (rotation-search) failed first with target type `""` | CONFIRMED-BUG (tooling) | S b6db6f55e |
| X2 | **Creature-type policy (assigned item 3).** Phase 1 bosses, from Wowhead classic NPC tooltips read 2026-10-31 (`nether.wowhead.com/classic/tooltip/npc/<id>`): Humanoid ×5 (Lucifron 12118, Gehennas 12259, Shazzrah 12264, Sulfuron 12098, Majordomo 12018), Elemental ×3 (Garr 12057, Geddon 12056, Ragnaros 11502), Giant (Golemagg 11988), Beast (Magmadar 11982), Dragonkin (Onyxia 10184). **Decision:** published numbers stay on the type-neutral target, so slaying gear counts for nothing. Two talents had to follow from that. Improved Tracking applies on all 11 bosses, so it now counts on the type-neutral boss (X3). Murder applies on 6 of 11, so it is kept out of the search's reach instead (R1). In the measured specs, only the rogues (Murder) and Marksmanship/Survival (Improved Tracking) move with target type; the other 14 DPS specs do not move at all (scan at 1000 iterations, humanoid against unknown). | `spec-breakdown -target-type` (new flag, S 02c29aed1) | decided | - |
| X3 | **Improved Tracking was worth nothing in every published hunter number.** It matched only the seven listed types, and the ranker's target has no type. A hunter tracks whatever the boss is, and every Phase 1 boss is a trackable type. An Unknown-type target now counts unless the encounter is a training dummy. Mechanical still gets nothing. Survival +5.00%, Marksmanship +4.18% (Beast Mastery takes no points). The fork's suites fight a Demon and do not move. | client tooltip 24293 ("...increased by 5%"); `sim/hunter/talents.go` `improvedTrackingApplies`; test `TestImprovedTrackingCountsOnATypeNeutralBossButNotOnADummy` failed first (ratio 1.000 against 1.05) | CONFIRMED-BUG | E 701a61a40 |
| X4 | The fork's Settings suite stands every player 30 yd away (`sim/core/test_generators.go:184`). A melee warrior there never swings: its Heroic Strike queue sits for 59.5 s, and the 60 s row is about three Bloodthirsts. That is why round 2's gated cooldowns "lost 13-30%" in `TestP1DPSWarrior` ShortSingleTarget (F1). It is a harness artifact, not a rotation result, and it affects every melee suite in the fork. | debug run of the suite request: no auto-attack action at all | OPEN (harness; fork test-only; needs a melee distance per class in the generator and a regeneration of every melee golden) | - |
| X5 | Round-2 side effects re-checked: the pre-pull landing hold (the hunters' -3.5 s Aimed Shot lands with 0.95 hits; Subtlety's Ambush at 0 counts; the warriors' Battle Shout at -1 s is a buff and keeps its time), Blood Fury off the global cooldown (Arms orc: 2 casts in 180 s), Heart of the Lion (10% stats plus 280 AP, `sim/core/buff_ranks.go:135-142`, tooltip 409583), stones beside poisons (Combat and Subtlety carry both on each hand). No new bug. | breakdowns `scratchpad/r3/physical/bd/` | aligned | - |
| X6 | Tools: spec-breakdown gains `-duration` for the DPS fight (the site's control allows 20-600 s) and `-target-type`. | tests `TestBuildRequestTakesTheFightLength` (failed first: 60 s fought 180), `TestBuildRequestTakesTheTargetType` | tooling | S c78603e7e, S 02c29aed1 |
| X7 | Real logs: still no level-60 Forever logs (launch is 4 Nov). | ledger | OPEN | - |

## Warrior: Fury

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **Round 2's F3, settled.** Gating Death Wish and Recklessness on the execute phase alone loses at most fight lengths: -23% at 30 s, -5% at 60 s, -1% at 90 s and -1.3% at 300 s, where it gives up the second Death Wish. It wins only at 120-180 s. Death Wish lasts 30 s with a 3 min cooldown, and Recklessness 15 s with a 30 min cooldown (SpellCooldowns 54733 and 54205; Forever tooltips 12328 and 1719). A 20% execute window in a short fight is shorter than either buff. **New rule:** Death Wish goes now if a second use still fits (at least 210 s left), otherwise in the execute phase or the last 30 s. Recklessness goes in the execute phase or the last 15 s. It beats the old rotation at every length measured (horde, 5000 iterations, seed 1): 30 s +0.3%, 60 s +1.2%, 90 s +1.5%, 120 s +1.0%, 180 s +0.4%, 300 s +0.6%. At 10000 iterations and two seeds: 60 s +1.44/+1.43%, 180 s horde +0.34/+0.44%, alliance +0.80/+0.89%, bare 60 band +6.6%. The site's default fight is 180 s with 20% variation (`sim/api/envelope.go:391`), and the coach's log references use real fight lengths, so short fights count. The rotation's notes and coach prose follow. | `scratchpad/r3/physical/rot/fury-G.json`; `data/tests/test_apl.py` name sidecar updated (Recklessness 1719) | CONFIRMED-BUG (rotation) | S 3cf977743, E 2dbb91aee (fork copy and the regenerated `TestP1DPSWarrior` goldens) |
| F2 | Normalized-rage crit factor 2.0 (round 2 F4). Blizzard's 1 October beta notes say crits give "100% increased Rage (was 75%)", which is the 2.0 (`sim/core/rage.go:63-66`). The owner's 17 Sep sample predates that note and has one ambiguous crit swing. Its per-swing rage fits the speed constants: 2.29 s to 7.5-7.7, 3.01 s to 13.5 (two-hand 4.5/s). | `scratchpad/r2/physical/logs/sample_rage.txt` | aligned on one official source; OPEN for a log check | - |
| F3 | Bloodthirst is 45% of AP plus 48 (tooltip 23894, `sim/warrior/bloodthirst.go:30`). Death Wish in Forever is +20% physical damage done and +5% damage taken (tooltip 12328). | Forever tooltips | aligned | - |

## Warrior: Arms

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | The Fury cooldown rule tried on Recklessness: no change at 60, 180 or 300 s. Arms only reaches Berserker Stance in the execute phase, so Recklessness already lands there. | `rot/arms-R.json` | aligned | - |
| A2 | Mortal Strike is weapon damage plus 160 (tooltip 21553). Blood Fury is 10% AP (tooltip 20572; `sim/core/racial_blood_fury.go:66`) and is cast twice in 180 s off the global cooldown. | breakdown | aligned | - |
| A3 | Zero new findings beyond A1-A2 (rotation, talent masks of the cast spells, Slam and Hamstring as in round 2). | | aligned | - |

## Warrior: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Revenge rank 6 (25288), settled.** It is an Ahn'Qiraj book rank. It shares its SkillLineAbility batch with the other warrior AQ books: rows 13217 (Heroic Strike 9, 25286), 13218 (Revenge 6) and 13219 (Battle Shout 7). All three carry AcquireMethod 0, exactly like trainer ranks (Revenge 5 is row 6108), so the table cannot tell a book from a trainer rank. The trainables rule therefore gives no answer, and the curated book list stands. Ahn'Qiraj is not in Phase 1, so the rank cannot be learned at launch. The sim casts rank 5 (11601). | SkillLineAbility 70707; `data/curated/book-ranks.json:257`; breakdown | aligned (not learnable in Phase 1) | - |
| P2 | Sunder's +5% of AP threat: still single-source (sixtysim). The Forever tooltip (11597) says only "a high amount of threat". | tooltip 11597 | OPEN (needs a threat log) | - |
| P3 | Revenge tooltip at 70707: 137.7-168.3 + 0.25 × AP, as modelled. | tooltip 25288 | aligned | - |

## Rogue: Assassination

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **The raid search could trade Murder away.** On the type-neutral target, Murder is worth 0 (probe +0.00), but it gives +4% on 6 of the 11 Phase 1 bosses. Every rogue raid build takes it. It is now in the raid keeps for all three rogue specs. | `design/reviews/talent-search/raid-keeps.json`; tooltip 14158 | CONFIRMED-BUG (search input) | S 38bc96fb9 |
| R2 | **Venom with Vile Poisons (assigned item 2).** Both are aura 108 (add-percent spell modifier), with the same misc values (0 damage, 22 periodic) and the same class masks (8192 on Instant Poison, 65536 on Deadly Poison). The 1.12-based server sums percent modifiers of one kind (cmangos-classic `Player::ApplySpellMod`). So additive, as the engine has it (`sim/rogue/venom.go:34-37`). Tooltips: Venom 30%, Vile Poisons 4% a rank. That is one client source plus an emulator convention, not a Forever measurement. Side finding: neither talent's damage mask covers Wound Poison (mask 268435456), yet the engine scales Wound Poison with both. No rotation or preset uses Wound Poison, so the effect is 0%. | SpellEffect 693770-693771, 1340417-1340418; SpellClassOptions | aligned (additive); Wound Poison scaling OPEN (0% today) | - |
| R3 | **Talent search on daggers (round 2 R6).** The published Assassination set is still the mace and sword one: origin/main's file is the nightly at 17d31ecfa, and the bis.yml run in progress had not landed. With Heartseeker (12783) and Scarlet Kris (5267) swapped into the raid band locally (not committed), the search keeps the guide build: 591.8 ± 0.9, nothing beyond error. | `scratchpad/r3/physical/ts-out/rogue-assassination-raid.md` | aligned | - |
| R4 | The published set still sims about 260 DPS, because Mutilate is dead on a mace and sword. This is not a sim bug: it waits on the bis.yml re-rank (round 2 S a76dfad65). | breakdown 259.5 | OPEN (bis.yml) | - |

## Rogue: Combat

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | Poisons by hand: Deadly on the main hand and Instant on the off hand, against the swap: -0.42% for the swap. | `-consumes` swap | aligned | - |
| C2 | Zero new findings in the rotation (Backstab builder on a dagger, Eviscerate, two Blade Flurry casts, Adrenaline Rush) and the build (round 2 C3 still holds). Slice and Dice's 15 s gate is round 2's C2 HOLD. | breakdown | aligned | - |

## Rogue: Subtlety

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | "Ghostly Strike on cooldown" fires 1.5 times a fight, not about 9. Hemorrhage sits lower in the list but takes the energy first: the APL casts it as soon as energy covers it, so energy never pools to Ghostly Strike's 40. The cooldown itself is right (20 s; SpellCooldowns 54818; tooltip 14278; engine test). Pooling energy for Ghostly Strike (three gate variants) measured -0.04% to -0.15%, because the two strikes do about the same damage per energy. | `rot/sub-{A,B,C}.json`; debug log | aligned in DPS; the APL note overstates it | - |
| S2 | Swapping poison hands (Instant on the main hand): +0.47%, under 1%. | `-consumes` swap | HOLD | - |

## Hunter: Beast Mastery

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | The pre-pull Aimed Shot lands at 0 (0.95 hits). The cat is out at full uptime. Bestial Wrath is cast twice. The build takes no Improved Tracking, so X3 does not move it. | breakdown | aligned | - |
| B2 | Pet family (round 2 B4). | | OPEN (no Forever raid-pet consensus) | - |

## Hunter: Marksmanship

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| M1 | **Baseline gap (assigned item 5): explained and fixed.** It was the target type (X1). | X1 | CONFIRMED-BUG (tooling) | S b6db6f55e |
| M2 | Improved Tracking 5/5 now counts: +4.18%. | X3 | CONFIRMED-BUG | E 701a61a40 |
| M3 | Aimed Shot is cast only before the pull. It shares Multi-Shot's 6 s cooldown, and the APL prefers the instant Multi-Shot on purpose (APL note). | breakdown | aligned | - |

## Hunter: Survival

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| V1 | Improved Tracking 5/5 now counts: +5.00%. | X3 | CONFIRMED-BUG | E 701a61a40 |
| V2 | Lone Wolf (Marksmanship tier talent, 415370) is 20% damage done with no pet (aura 79, all schools); the engine gives 1.2 and sims no pet. Lacerating Strikes and Lacerate tick and crit. | SpellEffect 1085262; `sim/hunter/talents.go:565-570` | aligned | - |

## Before/after (lane engine 701a61a40 plus site branch, against the pin c8d29ea33 with origin/main's Fury rotation; alliance raid preset)

| Spec | seed 1 | seed 50001 | what moved it |
|---|---|---|---|
| Arms | +0.00% | +0.00% | |
| Fury | +0.80% | +0.89% | F1 (horde 180 s +0.34/+0.44%, 60 s +1.44/+1.43%) |
| Protection | +0.00% | +0.00% | |
| Assassination | +0.00% | +0.00% | (dead set until bis.yml) |
| Combat | +0.00% | +0.00% | |
| Subtlety | +0.00% | +0.00% | |
| Beast Mastery | +0.00% | +0.00% | |
| Marksmanship | +4.18% | +4.18% | Improved Tracking |
| Survival | +5.00% | +5.00% | Improved Tracking |

No spec outside this group moves. The target-type scan over every DPS spec moved only the rogues and the two hunters, and the search-tool change does not touch spec-breakdown or the ranker.

## Still open

- **Melee in the fork's Settings suite stands at 30 yd** (X4). The suite cannot judge a melee rotation until the generator gives melee specs melee range.
- **Page against published:** the /sim page's default target is `""`, which means Humanoid, so the page shows rogues 4% above their published number. The fix is a page default, outside this lane.
- **Sunder's 5% AP threat** is single-source. **The rage crit factor** rests on one Blizzard note.
- **Assassination** waits on the bis.yml re-rank. Wound Poison takes Vile Poisons and Venom in the engine, against the client masks (0% today).
- **Pet family**; **real level-60 logs** after launch.
- **TestRotationLadder:** Fury's rows move (F1). Left for the integrator.
