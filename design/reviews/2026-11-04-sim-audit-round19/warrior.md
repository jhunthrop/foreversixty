# Round 19, lane warrior (Arms, Fury, Protection)

Site `wt/r19-warrior` on int/r19 (f91e9d3f9). Engine `lane/r19-warrior` on fork int/r19 at 1789ffd5e. Client 1.60.1.70755 with the late hotfix re-layer (layered raw tables in `scratchpad/c755/fetch/.../raw/`, `talents/warrior.json`, `spelltips/warrior.json`, Wowhead Forever tooltips).

**Result: 3 confirmed warrior sim bugs, all `old`, all fixed in the engine.** The new source this round is the Forever bug tracker run by the Classic community, where Blizzard's liaison answers (github.com/ClassicWoWCommunity/forever-bugs), plus the Forever warrior research group's raw beta combat logs (github.com/ppach/marrow.github.io `eternal/data/logs`, github.com/magey/forever-warrior discussions). Rounds 16 to 18 never read either one. Seed S1 is settled by these logs.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Heroic Strike and Cleave counted as white swings for procs.** The engine gives both next-swing abilities a main-hand auto bit (`ProcMaskMeleeMHSpecial \| ProcMaskMeleeMHAuto`, from upstream SoD 543752637). So each Heroic Strike hit could proc Unbridled Wrath and spent a Flurry charge and a Reckless Swings charge. In the client, all three auras carry ProcTypeMask 4, "done melee auto attack": Unbridled Wrath 12322, the Flurry charge aura 12966 (3 charges) and Reckless Swings 1329555 (4 charges). The Forever logs say a Heroic Strike is not an auto attack for that mask. (1) Beta build 1.60.1, 2026-09-23 (`ryqenz_zomdaya_uw_dw`, `siax_uw_dw`; forever-bugs #105): I counted it myself. Zomd and Siax landed 228 Heroic Strikes and Cleaves and had 472 Unbridled Wrath procs at a 35-36% rate on white swings. Not one proc followed a Heroic Strike or Cleave on its own. (2) Build **70450**, 2026-10-20, after Unbridled Wrath's rate fix (magey/forever-warrior discussion #30, Tawl, Unbridled Wrath 5/5 and Flurry 1/5, 1H and shield). Over 1,064 landed white swings and 228 landed Heroic Strikes, 640 procs came out at 60.2% of the white swings, and 0 followed a Heroic Strike. Flurry charges: 300 doses and 120 removals, every one after a white swing or a missed white swing, and none after the 228 landed Heroic Strikes. Nine came right after a *missed* Heroic Strike; that tiny case is not modelled. Classic Era agrees for Flurry (Aurana's sim notes, round 17). Fix: one rule, `countsAsWhiteSwing` (a pure auto, with no special bit), for Unbridled Wrath, Flurry charge use and Reckless Swings. Heroic Strike keeps its auto bit for everything else (swing timer, other procs). Tests: `white_swing_procs_test.go` failed first ("Heroic Strike / Cleave (mask 20): countsAsWhiteSwing = true, want false") and passes now. | logs above, `tools/uwhs.py`, `tools/flurry.py` in `$R/warrior/`; SpellAuraOptions for 12322, 12966 and 1329555 | CONFIRMED | old (rounds 16-18 kept it OPEN without these logs) | E af3bc587b |
| P5 | **Revenge and Shield Slam threat constants.** Searched for a blue post and Forever sources: wow.gg Forever Protection (updated 2026-09-18), mobalytics and rankedboost Forever pages, kami-labs patch digests, and every forever-bugs issue that mentions threat (#63, #64, #66, #278). None gives a number for Revenge or Shield Slam. Combat logs carry no threat. #63 (Sunder) and #64 ("Many Buffs and Debuffs now cause 0 threat", Battle Shout) are already in the engine. Still OPEN: a threat-meter capture (the #63 method, `UNIT_THREAT_LIST_UPDATE` deltas) or a blue post settles it. Ranks below 60 have the same problem: Shield Slam uses 508 at ranks 1-3 too, against the Classic table's 160/190/220/250. Revenge's formula does scale with level. Not changed, because the base values themselves are the open question. | URLs above; `revenge.go:19`, `shield_slam.go:22` | OPEN (carried) | - | - |
| P5b | **Heroic Strike threat was rank 8's at every rank.** The client has no threat effect for Heroic Strike. The engine's 145 is rank 8's row in the Classic server spell_threat table (mangoszero/database `World/Setup/FullDB/spell_threat.sql`; the same rows in mangosArchives databaseZero_Rel19 and tortoise-wow). That table lists 20, 39, 59, 78, 98, 118, 137, 145 and 175 by rank, and Wowhead's Classic threat overview also gives rank 9 as 175. But every rank got 145, so a level-20 tank's rank 3 Heroic Strike had rank 8's threat. Now it is per rank. Level 60 is unchanged (rank 8; rank 9 is the AQ book). Caution: these are one lineage of emulator data, and Forever's real number is as open as P5. The fix only makes the engine agree with its own source. Test `TestHeroicStrikeThreatIsPerRank` failed first at ranks 1-7 and 9, and passes now. | sources above | CONFIRMED (single lineage, flagged) | old | E fae61cd3d |
| AC | **Arcanite Champion's 1 PPM** (Arms raid main hand, both factions). ItemEffect 99501 is TriggerType 2 (chance on hit) with no rate. The 70755 tables have no SpellProcsPerMinute row, and 16916 has no SpellAuraOptions. Web: Wowhead Classic 16916 lists no rate; the Blizzard Classic forum thread "Weaponsmith - champion or reaper?" says "not sure of the PPM", and the SoD threads cover only the Refined version. No source. | ItemEffect, SpellAuraOptions; URLs in "Sources" | OPEN (shared item code) | - | - |
| X | **Coefficient lead.** For every warrior effect the engine scales by a coefficient, I checked the client's EffectBonusCoefficient. Heroic Strike, Execute, Overpower, Mortal Strike, Slam, Bloodthirst, Revenge, Shield Slam, Hamstring, Cleave and Spearing Strike are all 1 in the client and `BonusCoefficient: 1` in the engine (flat physical bonus). Rend and Thunder Clap are 0 in both. Deep Wounds' dummy 12162 is 0, and its ticks never read the coefficient. Attack-power terms against the tooltips: Bloodthirst 0.45 ("45% of your Attack Power plus 48"), Revenge 0.25 (`[109.07 + (0.25 * 1) * AP]`), Shield Slam plus block value and no AP (640-670). Warriors cast no spell-power spell. | SpellEffect; spelltips; nether.wowhead.com/forever | Aligned | - | - |

## Rank rows 20-59, level by level

Warriors have no rows in `data/curated/rank-choices.json`. Every level casts the highest rank learned. I judged each ranked spell against the client at every level from 20 to 59:

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **No lower rank can win.** Every rank of Heroic Strike (15 rage), Rend (10), Execute (15), Overpower (5), Sunder Armor (15), Revenge (5), Shield Slam (20), Thunder Clap (20), Demoralizing Shout (10), Bloodthirst (30), Mortal Strike (30), Slam (15, 15 s), Hamstring and Cleave has the same cost, cooldown, cast and duration rules as the top rank. Each lower rank does strictly less: less damage, less armor reduction, less threat (Sunder's effect 63: 34/75/117/158/206) and a smaller attack-power cut (Demoralizing Shout, whose MaxLevel caps each rank's growth). A warrior is never mana-bound, so no lower rank can gain at any level. | SpellPower, SpellCooldowns, SpellEffect, SpellLevels per rank (`tools/ranks.py`) | Aligned | - | - |
| R2 | **Cast check (the round-19 hunter trap).** The engine registers one rank per spell, the rank learned at the character's level (`rankAtLevel`). So I simmed all 120 levels, 20-59 for each spec (bare, 300 iterations, `-json`). Every APL spell is cast under the id the site rewrites it to, and each changes at its trainer level: Execute 5308/20658/20660/20661/20662 at 24/32/40/48/56; Heroic Strike 285→1608→11564→11565→11566→11567; Revenge 6572→6574→7379→11600→11601; Sunder 7386→7405→8380→11596→11597; Thunder Clap 8198→8204→8205→11580→11581, the trainable family (SkillLineAbility); Demoralizing Shout 1160→6190→11554→11555→11556; Mortal Strike, Bloodthirst, Slam, Rend and Overpower likewise. Nothing silently drops out. Two notes: one rank per level means a future downranked row would need the engine to register every rank first, as round 19 hunter did for Serpent Sting. And Cleave registers its top rank at every level, which no warrior rotation casts. | `$R/warrior/sweep/` | Aligned | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Deep Wounds (all three specs take 3/3).** Blizzard on forever-bugs #234 (2026-10-05), describing Forever's Deep Wounds against vanilla's: it "Rolls over its damage when refreshed", "Doesn't reset its tick timer when it is refreshed" and "Doesn't scale with Attack Power ... it scales with your weapon damage and not with your Attack Power at all". The engine got all three wrong: (a) it read the average weapon damage with attack power in it; (b) a new crit replaced the running bleed, losing what it had left; (c) a refresh restarted the ticks and threw a stray extra tick. Now a crit adds 20% a point of the hand's average weapon damage (off hand at half) to what the bleed has left, spread over four ticks on the running timer, and the bleed ends on its last tick. At raid 60, Fury's Deep Wounds goes from 24 ticks a fight at 82 to 59 ticks at 141 (5.1% of damage, up from 1.3%), because a warrior who crits every second no longer throws the bleed away. Round 6 (warlock Y4) cited #234 for a different question and filed it "Aligned". Tests: `deep_wounds_formula_test.go` (attack power, rollover) and `dps_warrior/deep_wounds_refresh_test.go` (ticks at 3/6/9/12/15 s after a refresh at 4 s; the old code gave six ticks). Both failed first and pass now. | forever-bugs #234 (Blizzard reply); talent text 20/40/60% over 12 sec | CONFIRMED | old (round 6 looked and missed) | E c647736e6 |
| A2 | Rotation after the fixes: Arms queues Heroic Strike only above 90 rage, so S1 barely touches it (-0.25%, from Reckless Swings and Flurry charges). | measure | Aligned | - | - |
| A3 | 70755 durations and costs against the engine: Rend 9/12/15/18/21/21/21 s, Thunder Clap 10-30 s, Mortal Strike 10 s, Recklessness 15 s / 30 min, Death Wish 30 s / 3 min / 10 rage, Spearing Strike 15 rage / 20 s / 75% (+75% on giants and dragonkin), Reckless Swings 8 s / 4 charges / 15 s / 25% base health. All match. Every talent curve in TraitDefinitionEffectPoints 70755 matches the engine. Dual Wield Specialization has a curve for an effect index 2 that the spell has no row for, so it does nothing. | SpellMisc, SpellDuration, SpellPower, CurvePoint | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | After S1, Unbridled Wrath rage drops from 97 to 79 a fight, Flurry uptime goes from 83.6% to 86.5%, and Reckless Swings uptime from 23% to 29% (horde, seed 1). Net +0.05 to +0.26%. | `bd-fury-*` | (S1) | - | - |
| F2 | The Heroic Strike rage gate re-probed after S1 (horde, seed 1, 3000 iterations): 45 → 873.6, 50 → 874.3, 60 (current) → 874.7, 70 → 873.4. 60 stays. | `-rotation` | Aligned | - | - |
| F3 | Unbridled Wrath's rate, 60% at 5/5, is confirmed on 70450 (640 procs on 1,064 landed white swings), after the 1 October fix (#105). | log above | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Tank fight after the fixes: 1088.6 / 1079.1 threat/s, alliance / horde, seed 1 (-0.8%). The loss is Deep Wounds losing its attack power; Protection crits too rarely for the rollover to make it back. Bare bands lose 0.4-2.8% threat from Heroic Strike's per-rank threat (P5b). | measure | (A1, P5b) | - | - |
| P2 | Shield Block 2565: the comment in `shield_block.go:11` says the client "carries no charge". The 70755 SpellAuraOptions does carry ProcCharges 2 (and 7 s, 5 s cooldown), so the engine's two charges now have the client behind them as well. Comment only; not changed. | SpellAuraOptions | Aligned | - | - |
| P3 | Open tracker items with no Blizzard answer and no model change: #409 (Defensive Stance cuts rage from damage taken to 90%; the conversion uses spell level), #192 (parry against +2 mobs ~11%; sims stand behind the boss), #278 (rage-gain threat, partly fixed on 2026-10-08: 25 threat per 5 rage from a parry or dodge, which matches the engine's 5:1). | forever-bugs | OPEN / Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Eureka! with Execute** (Gnome only): forever-bugs #439 (open, 2026-10-27, a 16 October log) shows that Eureka!'s 10% rage cut gives Execute nothing, because Execute takes the whole bar anyway. The engine applies the cut, which leaves more rage to convert. It is one report on an open issue, for a race no guide picks, so it is not changed. | #439 | OPEN | - | - |
| C2 | Deep Wounds' attacker multiplier is applied at the proc and again on every tick (`deep_wounds.go`, "Double dips"). Blizzard removed Ignite's double dip, but no source covers Deep Wounds. | code | OPEN | - | - |
| C3 | Talent builds after the fixes: no change is warranted. Deep Wounds is worth more now, and every build already takes 3/3. | - | Aligned | - | - |

## Before/after (raid 60, 10,000 iterations, guide builds via `rt`, frozen BiS)

The lane base (fork 1789ffd5e) equals `$R/before.tsv` to the decimal for all eight Arms/Fury values. Seed 1 / seed 50001:

| Spec | before (alliance; horde) | after S1 only | after all | change |
|---|---|---|---|---|
| Arms | 820.4 / 820.4; 826.4 / 826.5 | 818.4 / 818.0; 824.2 / 824.4 | 853.0 / 852.0; 855.5 / 854.4 | +3.4 to +4.0% |
| Fury | 879.4 / 878.9; 873.3 / 873.2 | 879.8 / 879.9; 875.5 / 875.5 | 917.6 / 918.3; 909.8 / 910.5 | +4.2 to +4.5% |
| Protection (threat/s) | 1097.0 / 1096.3; 1088.0 / 1088.0 | unchanged | 1088.6 / 1087.9; 1079.1 / 1079.6 | -0.8% |

Bare bands, alliance, seeds 1 / 50001 (before → after): Protection threat 20: 73.6 → 72.3 (-1.8%), 30: 143.7 → 139.7 (-2.8%), 40: 371.8 → 367.7 (-1.1%), 50: 515.8 → 513.9 (-0.4%). Fury 20 -0.6 to -0.9%, 30 -0.2 to -0.3%, 40 +0.1 to +0.2%, 50 0.0%. Arms 20 0.0%, 30 -1.4%, 40 -0.3%, 50 +2.4%.

No commit touches anything outside `sim/warrior`. **No spec outside the class moves.**

## Where the raid-60 gain comes from (coordinator ask)

Measured cumulatively on temporary worktrees at each fork commit (10,000 iterations, frozen BiS, guide raid builds via `rt`). "Alone" is each fix applied to the base by itself. Each cell is seed 1 / seed 50001.

| Spec, faction | base 1789ffd5e | + Heroic Strike procs (af3bc587b) | + threat (fae61cd3d) | + Deep Wounds (c647736e6) = all | Deep Wounds alone |
|---|---|---|---|---|---|
| Arms alliance | 820.4 / 820.4 | 818.4 / 818.0 (-0.24 / -0.29%) | same (0.00%) | 853.0 / 852.0 (+3.97 / +3.85%) | 854.2 / 853.8 (+4.12 / +4.07%) |
| Arms horde | 826.4 / 826.5 | 824.2 / 824.4 (-0.27 / -0.25%) | same | 855.5 / 854.4 (+3.52 / +3.38%) | 855.6 / 854.8 (+3.53 / +3.42%) |
| Fury alliance | 879.4 / 878.9 | 879.8 / 879.9 (+0.05 / +0.11%) | same | 917.6 / 918.3 (+4.34 / +4.48%) | 916.5 / 916.1 (+4.22 / +4.23%) |
| Fury horde | 873.3 / 873.2 | 875.5 / 875.5 (+0.25 / +0.26%) | same | 909.8 / 910.5 (+4.18 / +4.27%) | 908.1 / 908.9 (+3.98 / +4.09%) |

The threat change does not move DPS, so the Heroic Strike change alone equals the "+ Heroic Strike procs" column. The two fixes are close to multiplicative: base × (Heroic Strike alone) × (Deep Wounds alone) lands within 0.1% of "all" in every cell (for example Fury horde seed 1: 1.0025 × 1.0398 = 1.0424, against 1.0418 measured). So the order effect is at noise level, and its sign is mixed.

Why each fix moves DPS (horde, seed 1, from `-json`):
- **Heroic Strike change, Fury +0.25%:** Heroic Strike no longer spends Flurry charges, so white swings keep them. Flurry uptime goes from 83.9% to 86.8%, and white swings from 60.7 to 61.4 (main hand) and 81.8 to 82.4 (off hand) a fight. That outweighs losing Unbridled Wrath on Heroic Strike: 97.6 rage down to 79.1, so 34.2 Heroic Strikes become 33.9.
- **Heroic Strike change, Arms -0.25%:** Arms has no Flurry, so the change only costs rage. Unbridled Wrath falls from 41.7 to 37.7, and Reckless Swings rage from 166 to 152, because with a slow two-hander some of the four charges now run out unused instead of going to a Heroic Strike. Heroic Strikes fall from 7.9 to 7.4.
- **Deep Wounds, +3.4 to +4.5%:** a crit used to replace the running bleed, so a warrior who crits about once a second threw most of each bleed away. Now every crit's 60% of weapon damage is kept. Fury's Deep Wounds goes from 24.3 ticks at 11.1 DPS to 58.8 ticks at 46.1 DPS, and Arms' from 34.3 ticks at 29.3 DPS to 58.2 ticks at 59.0 DPS. Losing attack power from the formula cuts each crit's bleed by more than half, but the rollover more than makes up for it.

## Race order on the final engine (guide lane's carried ask; frozen BiS, 10,000 iterations, seeds 1 / 50001)

- Arms Alliance: Skyborne 856.4/855.7, Night Elf 856.0/855.4, Dwarf 855.8/855.6, Human 853.0/852.0, Gnome 852.3/851.1. Horde: Skyborne 858.3/858.4, Orc 855.5/854.4, Undead 855.3/855.4, Troll 852.7/852.1, Tauren 847.1/846.3.
- Fury Alliance: Skyborne 918.9/920.2, Night Elf 918.6/919.8, Human 917.6/918.3, Gnome 916.1/917.5, Dwarf 914.2/913.3. Horde: Skyborne 914.2/914.1, Undead 913.7/913.1, Orc 911.9/911.6, Troll 909.8/910.5, Tauren 902.7/902.2.
- Protection (tank fight; damage taken /s, TMI, threat/s at seed 1): Night Elf 259.2, 22.9, 1084; Dwarf 265.2, 23.1, 1089; Human 264.2, 23.1, 1085; Gnome 263.9, 23.2, 1085; Skyborne 264.6, 23.2, 1081. Horde: Tauren 252.8, 23.2, 1079; Orc 252.4, 24.6, 1087; Undead 252.1, 24.7, 1101; Troll 251.4, 24.6, 1083; Skyborne 252.7, 24.6, 1084. Alliance and Horde run different raid presets, so their TMI figures do not compare across factions.
- Every pick holds under the 1% race rule. Before the Deep Wounds fix, Dwarf led the Alliance for Arms (824.1 against Night Elf's 820.2); after it, the two are level.

## Files to regenerate

- Site `sim/request` TestRotationLadder goldens: warrior-protection (confirmed differs); warrior-arms and warrior-fury probably as well.
- `data/builds/1.60.1.70755/dps-per-point.json` and the stat weights that feed addon-data (Deep Wounds no longer reads attack power; weapon damage and crit gain). Fork golden TestP1DPSWarrior shows hit 19.4 → 15.6 and crit 6.1 → 7.4 from S1 alone.
- BiS (warrior specs) at the CI re-rank.
- addon-data.json and Data.lua (Arms Reckless Swings card note).
- Fork goldens regenerated in this lane: `sim/warrior/dps_warrior/TestP1DPSWarrior.results` (E af3bc587b, E c647736e6) and `sim/warrior/tank_warrior/TestP1TankWarrior.results` (E c647736e6).

## Measure at refresh

- Warrior BiS after the re-rank: Deep Wounds now favours weapon damage and crit over attack power.
- Arms Alliance Boot Flask and Frozen Heart picks (carried).

## For other lanes

- Engine changes are confined to `sim/warrior`. Shared items left alone: Arcanite Champion's rate (AC); Eureka! with Execute (C1).
- The tracker plus the community warrior logs are a source no earlier round used: ClassicWoWCommunity/forever-bugs (Blizzard answers on each issue), ppach/marrow.github.io `eternal/data/logs` and magey/forever-warrior discussions (raw WoWCombatLog files). Other melee lanes may find Flurry or Windfury evidence there.

## For the guide lane

- The Unbridled Wrath and race lines are done in this lane (see "Guide text").
- Optional, arms.md:67 and fury.md:74: Deep Wounds could add that a crit on a running bleed adds to it rather than replacing it. The current text is not wrong.
- Fury Death Wish timing inconsistency (above).

## Guide text

- Done in this lane, in their own commits: arms.md and fury.md now say Unbridled Wrath procs on white swings; Arms Dwarf and Night Elf level; Arms Skyborne level with Night Elf on the Alliance and a hair ahead of Undead on the Horde; Fury Skyborne and Undead level on the Horde, with Tauren last; Reckless Swings "white swings" on the Arms page and the Arms card note (curated plus the sim/request copy). wow-player: SHIP on the rendered Arms, Fury and Protection pages after one round (the MUST was the Arms Skyborne line). Captures: `$R/warrior/shots2/`.
- Left for the guide lane: Fury's Death Wish timing ("last 30 seconds" in Cooldowns against "only if a second one fits" on the card), which predates this pass.
- The Arms card note change needs the addon-data and Data.lua regen to reach the addon.

## Still open

- P5: Revenge, Shield Slam and Thunder Clap threat (a threat-meter capture or a blue post).
- AC: Arcanite Champion's proc rate.
- C1: Eureka! with Execute; C2: Deep Wounds' attacker-multiplier double dip.
- Nine Flurry charges spent right after a *missed* Heroic Strike on 70450 (not modelled).
- WARRIOR-13 (Improved Shield Wall) stays with the owner.
- Logs: foreverlogs.gg last checked 2026-11-04 19:52 UTC. The newest report is 2814 ("Test"); 2815 and later return "Report not found". `api.foreversixty.gg/v1/reports/recent` shows one 2026-10-20 report. There are no live logs yet (launch is at 23:00 UTC). Warcraft Logs is blocked (gcloud login); I did not work around it.

## Sources

- forever-bugs: https://github.com/ClassicWoWCommunity/forever-bugs/issues/105, /234, /439, /63, /64, /278, /409, /192, /65 (read 2026-11-04).
- Logs: https://github.com/ppach/marrow.github.io/tree/master/eternal/data/logs (2026-09-23); https://github.com/magey/forever-warrior/discussions/30 (70450, 2026-10-20).
- Warrior compendium: https://ppach-warriorcompendium.share.connect.posit.cloud/known-bugs.html, rage-on-crit.html, flurry.html, abilities.html (v 2026-11-04).
- Threat tables: https://github.com/mangoszero/database (World/Setup/FullDB/spell_threat.sql), mangosArchives/databaseZero_Rel19, tortoise-wow/tortoise-wow sql/base/tw_world_spell_threat.sql.
- Forever guides with no threat numbers: https://wow.gg/guides/warrior-protection-forever-overview (2026-09-18), https://mobalytics.gg/wow-forever/guides/warrior-class-overview, https://rankedboost.com/world-of-warcraft/forever/stats/combat-rules/.
- Arcanite Champion: https://www.wowhead.com/classic/spell=16916, https://us.forums.blizzard.com/en/wow/t/weaponsmith-champion-or-reaper/166147.
