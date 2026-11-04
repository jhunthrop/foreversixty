# Round 15, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70755, site int/r15 fd320b60e, fork int/r15 d1d9f2775. Branches: site wt/r15-hunter, fork lane/r15-hunter.
Sims: spec-breakdown through the lock, 3000 iterations to explore, 10000 at seeds 1 and 50001 for every reported number.

Note on the baseline: the default raid-60 runs (before.tsv) use the published BiS talents, which are 70753 carry-overs. For Marksmanship the published raid-60 entry runs the **leveling** build (no Lone Wolf, 708 DPS), not the guide raidBuild (745.7). Every build measure below therefore passes the guide string with `-talents` on both sides. The refresh replaces the published talents (MR19/MR20).

## Seed items (guide-accuracy round 1)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Focused Fire before Endurance Training (Beast Mastery leveling).** The plain walk spent levels 15 to 18 on Endurance Training 1 to 4 (pet health and armor), then took Focused Fire (2% more damage for you and the pet) at 19 and 20. Focused Fire's row opens at 15. | addon-builds.json BM order `[1,1,3]x4` before `[1,2,1]x2`; talents/hunter.json 104975 (tier 1) and 104976. The published bands (10, 20, 30...) hold the same points either way, so no sim can separate them. Owner ruling 2026-11-04: on a tie, real leveling play decides, through a checked-in order override. | CONFIRMED (owner ruling) | old, guides-handoff | S 426825c1c |
| H2 | **Sniper Shot left out of the Marksmanship raid build.** The 2026-11-03 push doubled it (rank 3 590 over normalized weapon damage, 4 s cast, 15 s cooldown). The raid build had 33 Marksmanship points and stopped one short. Sniper Shot now takes the third Improved Stings point (Serpent Sting 20% to 13%), and the rotation casts it after Multi-Shot under Aimed Shot's own gate (next auto shot more than 1.2 s away). Gain: +2.5% at both seeds and both factions. Swapping out Surefooted instead gains only +0.6% (hit). Casting it on cooldown with no gate gains +0.3 to +1.0%. | spellconst 1310687/1310785/1310786 (effect 121, 320/450/590), engine constants_auto_gen.go:4473, sniper_shot.go (Mortal Shots, Lone Wolf and Ranged Weapon Specialization apply; Barrage does not, per its text). Warcraft Tavern's Forever MM guide ([link](https://www.warcrafttavern.com/forever/guides/marksmanship-hunter/), read 2026-11-04, no date shown) takes Sniper Shot and recasts it on cooldown. | CONFIRMED | old (rounds 13 and 14 took the pushed value and never re-asked the build), guides-handoff | S 1b923c4f0, E 7e25c12ab |
| H3 | **Rapid Killing and Rapid Recuperation missing from the Marksmanship leveling build.** The engine has no kills on a boss dummy, so it cannot price Rapid Killing's 20% next Shot after a kill, or Rapid Recuperation's 100% clause. It does model Rapid Recuperation's Serpent Sting half. The sim ties at every bare band: 20 +1.0%, 30 -0.6%, 40 +0.3%, 50 +0.5%, 60 +0.2 to +0.4%, and time out of mana drops at 60. Under the owner's tie rule the Forever guides decide, and they take both. New leveling build `532000050/005055200152105/5`: Rapid Killing 2 and Rapid Recuperation 2, paid for by Ferocity 1, Improved Stings 1 and two Barrage ranks. Other point sources tested (Improved Tracking 5 to 1 or 3; Efficiency 5 to 3) cost up to 1.5% at band 30. | talents/hunter.json 105005 and 104999 (Rapid Recuperation requires Rapid Killing 2); engine talents.go:554-640. Web, read 2026-11-04: Icy Veins Forever MM leveling ([link](https://www.icy-veins.com/wow-forever/marksmanship-hunter-leveling), as quoted in search: the pair "will solve almost all of your mana issues while leveling"; the page itself returned 403) and Robcat's Mobalytics guide, updated 2026-10-17 ([link](https://mobalytics.gg/wow-forever/profile/robcat94/classes/my-class), "Rapid Killing 2/2"). The one dissent is a PvP guide. The Aimed Shot rank-1 bare-60 rank row still wins on the new build (333.5 vs 319.8 without it). | CONFIRMED (tie, web consensus) | old, guides-handoff | S 1134376ac |
| H4 | **Deadly Aspects under Aspect of the Falcon.** The engine rolls the ranged proc only under Hawk. Falcon's tooltip says the talent "can now proc from and benefit melee attacks". | sim/hunter/talents.go:433-470; spelltips 469145 | PENDING OWNER (do not reopen) | - | - |

## Beast Mastery

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | Talents in the raid and leveling builds checked against 70755 rank text and the engine: Deadly Aspects 10%/30%/12 s (Hawk ranged; Beast melee), Focused Fire 2%, Unleashed Fury 15%, Ferocity 10%, Frenzy 100%/30%/8 s, Bestial Wrath 50%/18 s/2 min/12% base mana (2 casts, 412.8 mana), Bestial Discipline (pet focus, 50% casting regen), Intimidation and Bestial Swiftness as prerequisites. Endurance Training is tier-0 gate filler; every other tier-0/1 filler (Improved Aspect of the Monkey, Improved Revive Pet) does less. | talents/hunter.json; talents.go:115-125, 241, 433-470; breakdown at raid 60 | Aligned | - | - |
| B2 | Summon Hawk for the fourth Endurance Training point (guides S7). The engine models only the dive. The "continuing its assault for 18 sec" half (trigger 1312639, the hawk NPC's kit) has no client body, and the hawk shares Arcane Shot's timer. So no sim can price it, and the web is split: some guides call it core, one skips it for Arcane Shot. | summon_hawk.go:10-35; spellconst 1293241; web search 2026-11-04 (mobalytics BM guides, wow.gg overview) | OPEN: the hawk's own attack data (creature for 1312639, or a Forever log of hawk swings) would settle it | - | - |
| B3 | Horde race (guides S6): Orc 678.9, Undead 676.6, Troll 673.6 (3000 iterations, guide raid build). Orc is +0.8% over Troll, under 1%, so the race stays. | spec-breakdown -race | Aligned (race within 1% stays) | - | - |
| B4 | Pet: a cat at 1.2 s swing speed, full uptime (sim/request/request.go:685). Forever keeps Classic's per-family attack speeds (EU beta feedback thread, [link](https://eu.forums.blizzard.com/en/wow/t/634929/1)). Cat damage is normalized by speed, and a fast cat is the standard pick. Cat swings and Frenzy/Bite/Claw counts are consistent with the 1.2 s base, Frenzy and Leader of the Pack. Cat vs raptor (guides owner question 4) is a pet-family choice the request cannot express. | request.go:675-692; pet.go:40-67 | Aligned (choice); cat vs raptor OPEN for the owner's question | - | - |
| B5 | Rotation: Hunter's Mark, Bestial Wrath on cooldown, Rapid Fire as an auto leaves with Multi-Shot ready, Aimed Shot when the next auto is over 1.2 s away, Multi-Shot, Serpent Sting under 1.5 s, Arcane Shot above 25% mana. No mana loss at raid 60. | data/curated/apl/hunter-beast-mastery.json; breakdown | Aligned | - | - |

## Marksmanship

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | Sniper Shot and the leveling build: see H2 and H3. | | CONFIRMED (2) | | S 1b923c4f0, S 1134376ac, E 7e25c12ab |
| M2 | 70755 talent values in the engine: Barrage 5/10/15% (barrageDamageByRank), Sniper Shot 320/450/590, Trueshot Aura 20/35/50/75/100 by trainer rank, ranged attack power only (core/buff_ranks.go:127, client aura 124). Lone Wolf 20% with no pet, and the Lone Wolf build sims with no pet (request.go hunterTakesLoneWolf). Improved Stings 6/13/20%, Efficiency 3%/rank on Shot/Sting flags (Sniper Shot included), Ranged Weapon Specialization on every ranged spell but Serpent Sting and the hawk. Fork hunter spellconst/hotfix/talent tests pass on the lane head. | talents.go:31-60, 143-155, 393-403, 595-601; constants_auto_gen.go:4473, 5381-5392 | Aligned | - | - |
| M3 | The deep-Marksmanship raid build (5/46/0, Warcraft Tavern) without Surefooted and Improved Tracking: 717.8 against 764.3. The Survival dip (hit plus 5% damage) stays. Survivalist 5 is the gate to Surefooted, and every other tier-0/1 Survival filler (Deflection, Entrapment, Savage Strikes, Improved Wing Clip) does nothing for a ranged hunter either. | spec-breakdown 3000 iterations | Aligned | - | - |
| M4 | Pre-casting Sniper Shot at -3.5 s in place of the -2 s Aimed Shot: +0.5%, under 1%. Not adopted: the 4 s pre-cast also lands before the tank has threat (guides HUNTER-10). | spec-breakdown | Aligned (kept) | - | - |
| M5 | Horde race on the new raid build: Undead 767.9, Troll 764.3, Orc 762.1, all within 1%. | | Aligned | - | - |

## Survival (melee)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Melee kit and costs at raid 60: Raptor Strike 25.8 casts at 40 mana (100 less Resourcefulness 60%), Mongoose Bite 20.3 casts at 26 mana, opened by Expose Prey on the raid preset's Hunter's Mark (10%, 10 s). Strider Kick 19.1 of 22.5 possible (8 s), Immolation Trap 9.7 (18 s with Survivalist's Discipline 2/2), Lacerate 58 a tick. Predator's Edge: 30% melee-ability crit damage on Raptor Strike, Mongoose Bite, Strider Kick, Wing Clip, Counterattack, Lacerate and Lacerating Strikes, and off hand +50%. Savage Strikes +4% on the same set. Lightning Reflexes 15% Agility. No mana loss at raid or bare 60. | talents/hunter.json 110861, 104983, 104985, 110859; talents.go:326-352, 683-809; breakdown | Aligned | - | - |
| S2 | Counterattack (1 point) never fires in a raid (the boss swings at the tank). Survival has one build at 60 (owner ruling) that also levels solo, where mobs hit the hunter and parries open it, so the point does something for the spec. Deterrence is its prerequisite. | apl expected_idle 20910; survival_melee_test.go | Aligned | - | - |
| S3 | Bare-60 Intellect weight 2.60 (guides S2). The weights predate the push (engine 8baffe5fc). On the current tree bare 60 runs out of mana for 0 s, so the weight should fall at the refresh. | spec-breakdown -preset bare | Measure at refresh | - | - |
| S4 | Horde race: Undead 821.7, Orc 819.4, Troll 817.6, all within 1%. | | Aligned | - | - |
| S5 | Crusader on both weapons vs Mighty Intellect (guides S5): the enchant table scores stat enchants only, and spec-breakdown cannot swap enchants. | web/src/lib/guides/gear-extras.ts | Measure at refresh / for other lanes (ranker) | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Sniper Shot's trainer row (SkillLineAbility 59257, AcquireMethod 2 to 0 in the 70755 push, round-14 build X5) only makes rank 3 a trainer rank of the talent spell, like ranks 1 and 2 (both already 0). Sniper Shot is not baseline: trainables lists it as source talent, and the engine's talent gate is right. | wago SkillLineAbility 70755 (59255-59257); trainables/hunter.json | Aligned | - | - |
| X2 | Teebu's Blazing Longsword's Firebolt (1300753) crits at the hunter's spell crit (15.8% bare, separate from melee 10.6%), as the crit split says. | common/item_effects.go:2250 | Aligned | - | - |

## Before/after

Guide builds passed explicitly on both sides (before = $R/sb-before at site-before; after = the lane's engine and site). 10000 iterations, raid 60 unless noted.

| spec / band | faction | seed 1 before to after | seed 50001 before to after | % |
|---|---|---|---|---|
| Marksmanship raid 60 (H2) | alliance | 739.9 to 758.4 | 739.9 to 758.7 | +2.50 / +2.54 |
| Marksmanship raid 60 (H2) | horde | 745.5 to 764.3 | 745.7 to 764.7 | +2.52 / +2.55 |
| Marksmanship bare 60 (H3) | alliance | 330.7 to 331.8 | 330.5 to 331.7 | +0.33 / +0.36 |
| Marksmanship bare 60 (H3) | horde | 332.4 to 333.5 | 332.8 to 333.7 | +0.33 / +0.27 |

Bare bands 20 to 50 for H3 (3000 iterations, seed 1, horde): 66.8 to 67.5, 110.3 to 109.6, 168.2 to 168.7, 237.1 to 238.3.
H1 moves no band. The default raid-60 runs (published talents, before.tsv) are unchanged for all three specs at both seeds and factions, because the published files still carry the old talents.

## Files to regenerate

- data/builds/1.60.1.70755/addon-builds.json (BM leveling order H1; MM leveling and raid builds H2/H3). `TestCommittedAddonBuildsFollowTheGuides` fails until it is regenerated.
- Ladder goldens for hunter-marksmanship (new leveling walk) and, if they print the order, hunter-beast-mastery.
- addon-data (MM rotation gains the Sniper Shot line; builds); coach starters for hunter-marksmanship; dps-per-point for the hunters.
- No apl-sync is needed: the fork and sim/request copies of hunter-marksmanship were edited by hand to match.

## Measure at refresh

- Marksmanship raid 60 on the new raidBuild (Sniper Shot) and every bare band on the new leveling build, plus their weights.
- Published MM raid-60 talents are the leveling build, and Survival's raid-60 talents are not the guide build (70753 carry-over). The refresh must rank raid 60 on the guide raidBuild.
- Survival bare-60 Intellect weight (S3).

## For other lanes

- Ranker/page: Crusader vs Mighty Intellect for Survival (S5). The guide's enchant table cannot show proc enchants.
- Web: run builds.test.ts and rotation-lists.test.ts (no node_modules in this worktree, so not run here). The MM build strings decode to 51 legal points in the sim's decoder.

## For the guide lane

All in web/src/content/guides/hunter/marksmanship.md (text as on land/guides-r1):
- Line 63: "thirty points in Marksmanship take the talents that raise your big shots and your auto shot, sixteen in Beast Mastery take the pet talents and Deadly Aspects" -> "thirty-one points in Marksmanship take the talents that raise your big shots and your auto shot plus Rapid Killing and Rapid Recuperation, fifteen in Beast Mastery take the pet talents and Deadly Aspects". In the same paragraph: "thirty-three points in Marksmanship add Lone Wolf and the rest of Improved Stings" -> "thirty-three points in Marksmanship add Lone Wolf, a second point of Improved Stings and Sniper Shot".
- Line 66: "**Barrage** (3 points)" -> "**Barrage** (3 points in the raid build, 1 while leveling)".
- Line 74: "**Improved Stings** (3 points in the raid build): 20% more Serpent Sting damage" -> "**Improved Stings** (2 points in the raid build): 13% more Serpent Sting damage". Add bullets: "**Sniper Shot** (1 point, raid build): a 4 second shot for your ranged damage plus 590, every 15 seconds" and "**Rapid Killing** and **Rapid Recuperation** (2 points each, leveling build): Rapid Fire every 3 minutes, 20% on your next shot after a kill, and Serpent Sting keeps your mana ticking while you shoot".
- Line 79: "So the raid build drops the pet talents, takes Lone Wolf, and spends the points on Improved Stings and Surefooted's hit" -> "...spends the points on Improved Stings, Sniper Shot and Surefooted's hit". Replace the last sentence "Sniper Shot, the capstone, costs a point you would take from elsewhere in either build, and its 4 second cast holds your auto shot longer than any other shot you have, so neither build takes it." with "Sniper Shot, the capstone, costs the third point of Improved Stings and is worth more; it holds your auto shot for its whole cast, so start it right after an auto shot fires. The leveling build skips it for Rapid Killing and Rapid Recuperation, which keep you shooting between pulls."
- Rotation paragraph ("Then work the shared six second cooldown..."), after the Multi-Shot sentence, add: "Sniper Shot has its own 15 second cooldown: cast it whenever it is ready right after an auto shot fires, the same rule as Aimed Shot."
- data/curated/apl/hunter-marksmanship.json: the new Sniper Shot line (after Multi-Shot) has an empty `notes`. Proposed: "Sniper Shot whenever it is ready and an auto shot has just left. Its long cast holds the next auto, so it only pays right after one fires."

## Wave 2: Aspect of the Falcon (closes H4)

Lead: guide-accuracy round 1 (land/guides-r1, Q1/HUNTER-9/S1). Under Falcon, Deadly Aspects never procs from auto shots. The lead also flagged Falcon's own 5% for a sim check.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | **Under Falcon, no ranged proc.** Falcon's proc types are melee only. Deadly Aspects' Auto Shot mod reaches Aspect of the Hawk's mask only. The engine already rolls the ranged haste only under Hawk. A new test pins it. | wago 1.60.1.70755 SpellAuraOptions 239200 (469145): ProcTypeMask 20 = 0x4 melee auto + 0x10 melee-class spell, no ranged bit. SpellEffect 19552 effect 0: aura 107 misc 18, mask word0 1048576. SpellClassOptions: 1048576 is carried by Hawk 13165-14322 and 25296. | Aligned (closes H4; no DPS change) | guides-handoff | E b86fa7d39 (test) |
| W2 | **Falcon's own proc was modelled as Deadly Aspects.** The client gives Falcon a flat 5% (ProcChance 5) on landed melee autos and melee strikes, triggering Quick Strikes (469144, +30% melee speed, 12 s). It needs no talent, and Deadly Aspects cannot raise it: Falcon's class mask is 0, while the talent's two chance mods reach Hawk (word0 1048576) and Beast (word2 16, spells 13161 and 1299445-1299447). The engine instead rolled Deadly Aspects' 2% a rank on white hits only, and nothing without the talent. Beast (1299447: ProcChance 0, ProcTypeMask 4, melee autos only, plus the talent's 10%) was already right. | wago 70755 SpellAuraOptions 239200 and 246416, SpellEffect 469145 effect 2 (aura 42, trigger 469144), 19552 effects 0 and 1, SpellClassOptions 78076 (469145 mask 0) and 83855 (1299447 word2 16). Fix: aspect_of_the_falcon.go gets its own proc; talents.go drops the Falcon case. The round-1 test that encoded the old reading (TestDeadlyAspectsProcsMeleeHasteUnderAspectOfTheFalcon) is replaced by aspect_of_the_falcon_test.go (5% without the talent, 5% with Deadly Aspects 5, autos and strikes, 0 on ranged). | CONFIRMED | old, guides-handoff | E b86fa7d39 |

Measured (10000 iterations, horde, guide build; before = sb-before, after = lane head):

| Survival | seed 1 | seed 50001 |
|---|---|---|
| raid 60, Beast (published rotation) | 818.6 to 818.6 | 819.0 to 819.0 |
| raid 60, Falcon in place of Beast | 818.6 to 817.2 (-0.17%) | 819.0 to 816.1 (-0.35%) |
| bare 60, Beast | 321.4 to 321.4 | 321.4 to 321.4 |
| bare 60, Falcon in place of Beast | 321.4 to 317.1 (-1.34%) | 321.4 to 317.0 (-1.37%) |

Before the fix, Falcon and Beast simmed identical: the same 110 melee attack power and the same proc. After it, Beast leads, so Survival keeps Beast. No rotation casts Falcon: Beast Mastery and Marksmanship stand on Hawk and Survival on Beast. So no published number moves, for hunters or anyone else.

For the guide lane (wave 2):
- survival.md:80 (land/guides-r1): "and Deadly Aspects procs from your swings under either one, so Falcon is no upgrade for you" -> "and Deadly Aspects only works under Beast: Falcon has its own smaller chance to speed your swings instead, so Beast is the better aspect for you".
- beast-mastery.md:81 and marksmanship.md:83 (land/guides-r1): "under Falcon Deadly Aspects only procs from melee swings, never from your auto shots, so Hawk is the ranged aspect" -> "under Falcon Deadly Aspects does nothing, and Falcon's own proc speeds only melee swings, so Hawk is the ranged aspect".

Noted, not changed: Hawk's own proc types are 320 (0x40 ranged auto + 0x100 ranged-class spell), while the talent text says "Auto Shot". The engine follows the text (Auto Shot only), as wowsims classic does. A log of Quick Shots procs off Multi-Shot or Arcane Shot would settle it (OPEN).

## Still open

- H4: closed in wave 2.
- Hawk's proc types (320) against the talent's "Auto Shot" text (see wave 2).
- B2: Summon Hawk's 18 s assault (needs the hawk's attack data).
- B4: cat vs raptor (owner question 4; the request has no pet-family field).
- Carried from round 14: Ranged Weapon Specialization on Serpent Sting (H4), Lacerating Strikes on refresh (H7), the queued Raptor Strike lifting the off-hand miss penalty (H10).

## Wave 3: rank rows on the shared low-rank penalty

Merged int/r15 into both branches; both fast-forwarded (site 1e96d8df2, fork 9c511a2cd, which already contain this lane's waves 1 and 2).

**Does the penalty touch hunter shots? No, in either direction.** It scales only a spell's spell-power coefficients (core/spell.go:355-359), and it skips Physical-school spells (levelPenalty, spell.go:813).
- Aimed Shot, Multi-Shot, Sniper Shot and the melee kit are Physical, so they are skipped.
- Arcane Shot (arcane_shot.go:60), Serpent Sting (serpent_sting.go:59), Volley (volley.go:82) and Immolation Trap's DoT carry spell-power coefficient 0 (client EffectBonusCoefficient 0; E 05b72a3b5, MR14), so the penalty multiplies nothing.
- The cat's abilities are Physical.

The runs confirm it: every number below is identical to the decimal between the round's start engine and the merged one.

Rows and fixed-rank casts (10000 iterations, horde; before = $R/sb-before, after = merged lane head, no-row = merged head with the hunter rank rows removed):

| band | seed 1 before / after / no-row | seed 50001 before / after / no-row | verdict |
|---|---|---|---|
| BM bare 40 (Aimed Shot rank 1, 40-49) | 167.9 / 167.9 / 162.9 | 168.1 / 168.1 / 163.0 | keep (+3.1%) |
| BM bare 50 (no row) | 229.9 / 229.9 / 229.9 | 229.9 / 229.9 / 229.9 | unchanged |
| BM bare 60 (Aimed Shot rank 1) | 329.5 / 329.5 / 315.6 | 329.5 / 329.5 / 315.5 | keep (+4.4%) |
| MM bare 60 (Aimed Shot rank 1, wave-1 leveling build) | 333.5 / 333.5 / 319.8 | 333.7 / 333.7 / 320.2 | keep (+4.3%) |
| BM raid 60 | 673.7 / 673.7 / 673.7 | 673.3 / 673.3 / 673.3 | no row |
| MM raid 60 (wave-1 raid build and Sniper Shot line) | 764.3 / 764.3 / 764.3 | 764.7 / 764.7 / 764.7 | no row |

Against wave 1: the MM numbers equal wave 1's after-figures (bare 60 333.5 / 333.7, raid 764.3 / 764.7). The engine change cost the hunters nothing, so there is nothing to recover.

Fixed-rank casts in the APLs:
- Beast Mastery and Marksmanship: Hunter's Mark 4, Aimed Shot 6, Serpent Sting 8, Arcane Shot 8, Aspect of the Hawk 6, Sniper Shot 3.
- Survival: Raptor Strike 8, Counterattack 4, Mongoose Bite 4, Lacerate 4, Strider Kick 1, Immolation Trap 5, Wing Clip 3, Aspect of the Beast 4.

All are authored at max rank and rewritten to the learned rank per band. None takes spell power that the penalty could cut, so no row is added, dropped or re-picked. Survival has no rank rows.

Rows changed: none. Commits: this report section only.

## Wave 6: guide-accuracy round 3 seeds (Sniper Shot, Summon Hawk)

Merged int/r15 into both branches (site ef3195852, fork 6a4695cc0). Seeds come from `origin/land/guides-r3:design/reviews/guide-accuracy/round3/hunter.md` ("Sims requested" 1 and 2, D2, D3).

**Live push check (wago.tools, read 2026-11-04 ~05:00 UTC).** wago's 70755 search totals are about 3 times the committed record (Spell 7480 to 22952). A per-record query shows why: every hunter row is the same push 113534, listed three times, with the values we already layer. The rows checked were Aimed Shot 697764 (350), Sniper Shot 1340384/1340617/1340619 (320/450/590), Aspect of the Hawk 692138 (110), Arcane Shot 691069 and Serpent Sting 690655. No other push id touches them, nor the SpellEffect, SpellAuraOptions, SpellCooldowns or SpellClassOptions rows of Falcon, Beast, Deadly Aspects, Summon Hawk, Bestial Wrath, Rapid Fire, Multi-Shot or the Survival kit. Sniper Shot's trainer row 59257 is still push 113007 (AcquireMethod 0).

A caution for other lanes: the druid lane's `druid/hfx/diff.json` "live" tables appear to be the shipped files plus only the newer pushes. Every hunter row shows there as reverted to pre-push values (Aimed Shot 166, Sniper Shot 295), which the per-record query disproves.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W6-1 | **Sniper Shot in the Marksmanship raid build (round 3, sim 1).** Already adopted in wave 1 (H2), taking Improved Stings 3 to 2 rather than the Surefooted point. On round 3's race (Night Elf, alliance, 10000 iterations): published build 745.9 / 745.9, wave-1 build 765.3 / 765.2 (+2.6%), and round 3's variant (Surefooted 3 to 2) 760.0 / 760.2 (+1.9%). The Improved Stings variant stays. Round 3's rotation gate ("next auto shot at least 4 s away") can never be true: the auto timer is about 2 s. The wave-1 gate (over 1.2 s) and casting on cooldown were both measured in wave 1; the gate wins. | spec-breakdown, seeds 1 and 50001 | CONFIRMED (already fixed) | old, guides-handoff | S 1b923c4f0, E 7e25c12ab |
| W6-2 | **Summon Hawk for the fourth Endurance Training point (round 3, sim 2).** Night Elf, alliance, 10000 iterations: guide raid build 674.8 / 674.9. With Summon Hawk 1 for Endurance Training 4 and Summon Hawk rank 4 in Arcane Shot's place: 664.1 / 664.3 (-1.6%). Casting it ahead of Arcane Shot: 664.3 / 664.6. Only the dive (108 + 5% ranged attack power) can be modelled. The 18 s assault is trigger 1312639, whose only client effect is a dummy (effect 3, no amount, period or creature), so it is server script with no data. Not adopted; the guide line stands. | SpellEffect 1293527 effects 0-3, 1312639 effect 0 (wago 70755); summon_hawk.go | OPEN (model gap: the hawk's attack data, from a Forever log of hawk swings, would settle it) | - | - |
| W6-3 | **Round 3's D2 is wrong on the data, but its prose holds.** D2 says Falcon has no melee proc chance of its own. SpellAuraOptions 239200 gives 469145 ProcChance 5 (ProcTypeMask 20, melee autos and strikes). This lane's wave 2 models that (E b86fa7d39). The survival.md wording "not a clear upgrade" is still true: Falcon sims 0.2 to 1.4% under Beast. | wave 2 above | Aligned (prose); data note for the guide lane | - | - |

For the integrator: land/guides-r3 changes the race field of every hunter build string to `night-elf`. This lane changed Marksmanship's `build:` and `raidBuild:` digits (S 1b923c4f0, S 1134376ac). The merged lines must keep both: `FS1:1.60.1.70755:hunter:night-elf:532000050/005055200152105/5:` and `FS1:1.60.1.70755:hunter:night-elf:5/0052550011503051/50005003:`.

## Wave 7: guide-accuracy round 4 items

Merged int/r15 into both branches (site 9adb256d3, fork 43c137aec). Source: `origin/land/guides-r4:design/reviews/guide-accuracy/round4/hunter.md` (HUNTER-R4-4, Sims requested 1 and 2, the integrator's carried runs).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W7-1 | **Pet buffs.** The engine keeps every raid, party and individual buff off the hunter pet (sim/hunter/pet.go:103 `RefusePlayerBuffs`; core/buffs.go:423 returns before `applyBuffEffects`). Round 4 reads Blizzard's sentence as stat buffs only. The one source is Blizzard's Hunter and Druid deep dive ([link](https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid), no date on the page, read 2026-11-04): "Pets can no longer receive player buffs that increase their stats." It reads either way: "buffs that raise stats" or "player buffs, which raise stats". The client cannot settle it, since buff targeting is server side and no client row says which units a raid buff skips. A web search (2026-11-04) found no Forever source naming Battle Shout, Blessing of Might or Leader of the Pack on a pet: the Warcraft Tavern and Icy Veins Forever buff articles and the kami-labs deep-dive recap all restate the same sentence. The stake is large. On a temporary engine with the refusal removed (not committed), Beast Mastery raid 60 (Night Elf, seed 1, 3000 iterations) goes 674.9 to 852.6. That is the upper bound, with every buff including the stat buffs Blizzard clearly removed; the non-stat subset is some fraction of it. Not changed: the shared buff code is untouched, and a narrower list would be a guess. Marksmanship (Lone Wolf, no pet) and Survival (no pet) cannot move. | sim/hunter/pet.go:90-104, sim/core/buffs.go:418-445, sim/core/pet.go:138-141 | OPEN for the owner or a log. A Forever pet's tooltip or character frame with Battle Shout, Leader of the Pack or Blessing of Might up, or a beta log of pet melee AP or crit under them, would settle it. | - | - |
| W7-2 | **Summon Hawk not modelled.** Round 4 cites no new data. The integrator's carried run (674.81 / 674.91, "identical to the published build") put Summon Hawk in the build but never cast it. This lane's wave-6 run casts it in Arcane Shot's place: -1.6%. The 18 s assault (trigger 1312639) is still a client dummy with no data. | wave 6 W6-2 | OPEN (unchanged) | - | - |
| W7-3 | **"No Sniper Shot line in Marksmanship"** (round 4's integrator note). Round 4 read land/guides-r4 752c02e43, whose hunter-marksmanship.json predates this lane: 0 matches for 1310786 there. On int/r15 9adb256d3 the line is in data/curated/apl/hunter-marksmanship.json, sim/request/apl/hunter-marksmanship.apl.json and the fork's ui/hunter/apls/forever_marksmanship.apl.json (43c137aec). The guide strings carry both changes: `build: 'FS1:1.60.1.70755:hunter:night-elf:532000050/005055200152105/5:'`, `raidBuild: 'FS1:1.60.1.70755:hunter:night-elf:5/0052550011503051/50005003:'`. Round 4's 738.2 for the Surefooted variant measured lost hit with no Sniper Shot cast. With the line, the wave-1 build is +2.6% (wave 6). | grep on int/r15 | Aligned (already fixed: S 1b923c4f0, E 7e25c12ab) | - | - |

Raid 60 is unchanged by wave 7 (no fix). Values on the merged tree, from wave 6: Beast Mastery 674.8 / 674.9 and Marksmanship 765.3 / 765.2 (Night Elf, seeds 1 / 50001). Survival as in wave 2.

For the guide lane (wave 7): marksmanship.md:82 on land/guides-r4 keeps the reason for skipping Sniper Shot ("twice as long as Aimed Shot ... so neither build takes it"). The raid build now takes it. Use the wave-1 replacement in "For the guide lane" above.
