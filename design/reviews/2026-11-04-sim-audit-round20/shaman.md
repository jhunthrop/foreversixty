# Sim-accuracy round 20: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70755 with the late hotfix re-layer. Site branch wt/r20-shaman (off int/r20 04d7da663), fork branch lane/r20-shaman (off int/r20 da54b7b06). Gear is the frozen BiS ($R/bis-frozen); raid 60 runs the guide raid build ($R/rt), bare bands the published band talents unless a row says otherwise. Ranks were judged level by level with `spec-breakdown -level`, forcing a rank by authoring its own id on the line it governs (the round-19 harness, copied to $R/shaman/sub.py and judge.sh).

**2 confirmed bugs, both fixed** (Elemental 1, Enhancement 1).
- Elemental: one of the four toll points under Elemental Fury was Elemental Devastation, melee crit for a caster that never swings. Origin `old` (round 17 E3 saw it and parked it with the owner; the owner's 2026-11-04 ruling hands such calls to the lane).
- Enhancement: the Maelstrom Weapon rank tag sat on the permanent trigger aura and printed a 100%-uptime "Maelstrom Weapon" line in the aura metrics. Origin `R19 73b7e850e`.
- Restoration: 0. Every round-19 row holds level by level.

Raid 60 is unchanged to the tenth for all three specs on both factions and seeds (table below).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Lightning Overload's own rank spells (408439 to 408477, 408479 to 408484), left at half the pre-push bolt by the 70755 push. Pending owner; not reopened. | lightning_bolt.go:41-62, chain_lightning.go:34-46 | pending owner | - | - |
| S2 | **Enhancement band 40 half-mana reserve (round 19 S6), re-judged.** Horde, bare 40, 3,000 iterations, seed 1: level 40 154.8 with the half reserve below Maelstrom 5/5 against 154.0 on a flat quarter (+0.5%); level 46 164.3 against 163.1 (+0.7%). The rule holds; the band's weights wait for the refresh. | `rows-recheck.txt` (lane dir), rotation variant `rot/enh-q25.json` | Aligned (Measure at refresh) | - | - |
| S3 | **The Maelstrom Weapon trigger showed as a 100%-uptime aura.**<br>What was wrong: round 19 put the talent rank on the permanent "Maelstrom Weapon Trigger" aura's ActionID tag so the Earth Shock line could read it with auraIsKnown; an aura with an ActionID is reported in the aura metrics, so every Enhancement breakdown listed "Maelstrom Weapon" at 100% uptime beside the real stacking buff.<br>Fix: the trigger has no ActionID again; a separate "Maelstrom Weapon Rank" aura, registered and never activated, carries the tag. auraIsKnown only asks whether an aura is registered. The rotation and the site's Earth Shock gate test are unchanged; no damage moves (raid 60 identical, below). | `TestMaelstromWeaponTriggerCarriesItsRank` extended (the tagged aura must never be active): failed, then passed; `TestEnhancementShocksDownToAQuarterOfItsManaOrHalfBelowFullMaelstrom` and the rotation smoke test pass on the new engine | CONFIRMED | R19 73b7e850e | E 511ec1ffe |
| S4 | **Round-19 rows re-judged level by level** (horde, 3,000 iterations, seed 1; healing in this profile is deterministic to 0.1 HPS):<br>- Chain Heal rank 1 (line 16) at 46, 49, 52, 55, 59: 164.4 / 169.8 / 267.1 / 263.4 / 270.8 against rank 2's 163.9 / 169.4 / 265.1 / 261.7 / 268.5 (+0.3% to +0.9%).<br>- Chain Heal rank 2 at bare 60: 326.8 against rank 3 323.5 (+1.0%, measured on a build with the row removed) and rank 1 326.4.<br>- Tank Healing Wave rank 7 (line 17) at 56, 57, 58: 262.8 / 264.9 / 270.5 against rank 6 261.5 / 260.5 / 263.5 and rank 8 259.7 / 261.7 / 269.5.<br>- Elemental Lightning Bolt rank 3 at 26 and 27: 33.6 / 34.3 against the top rank 32.3 / 33.0 (+4.0% / +3.9%) and rank 4 33.7 / 33.9 (a tie at 26, rank 3 +1.2% at 27).<br>Every row holds. | `rows-recheck.txt`; `$R/shaman/sb-norow` for the bare-60 Chain Heal check | Aligned | - | - |
| S5 | Enhancement's relic slot (round 18 S3): Totem of the Storm or Totem of Thunder once the ranker's relic gate lands. Nothing new to measure this round. | | Measure at refresh | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **A toll point that does nothing: Elemental Devastation in both Elemental builds.**<br>What was wrong: four points in tiers 1 to 3 are needed to open Elemental Fury's row, and both builds spent them on Elemental Devastation 1 and Elemental Warding 3. Elemental Devastation is "your Critical Strikes with direct damage spells increase your chance to get a Critical Strike with melee attacks by 9% for 10 sec": an Elemental shaman never swings a weapon, so the point is worth nothing in any fight. Round 17 E3 said no legal swap takes a point that does something; that is wrong on two or more targets. Trial by Fire (tier 3, legal once 15 points sit above it) is "Increases the damage done by your Fire Nova spell by 20% and reduces its cooldown by 4 sec, and increases the duration of your Magma Totem by 20 sec", and the rotation casts both on two or more targets (lines 4 and 6).<br>Decision: the tolls are Trial by Fire 2 and Elemental Warding 2 in both builds (5500331220001351). Web research (owner rule): Babyhoof's Mobalytics Forever Elemental guide takes Trial by Fire and calls it "fine to skip ... if you never plan to do dungeons" (https://mobalytics.gg/wow-forever/classes/elemental-shaman-guide-progress, read 2026-11-04); wowfdb's projected 60 build takes Trial by Fire 2 and no Elemental Warding (https://wowfdb.com/classes/shaman/elemental/talents/, read 2026-11-04). Two Warding points keep 7% less Fire, Frost and Nature damage taken, which a raider wants on fire bosses.<br>Cost: raid 60 identical to the tenth on both factions and seeds (667.3 against 667.3 at 3,000 iterations for the one-point swap; the before/after table at 10,000). Bare band 20 gains: the ladder walk's 11th point is now Elemental Alacrity 1 instead of Elemental Devastation 1 (horde 28.8 to 30.2, alliance 29.5 to 29.9, table below). Bands 30 to 60 are single-target identical.<br>Prose: elemental.md talents paragraph rewritten (Trial by Fire, Warding 2, why Devastation gets nothing); builds.test.ts, _prose.test.ts pass; wow-player on the rendered page below. | talents/shaman.json (Elemental Devastation 30160, Trial by Fire 16086); `TestGuideBuildsSpendNothingOnDeadTalents` (sim/leveling) now covers shaman-elemental's build and raidBuild: failed on Elemental Devastation 1, then passed | CONFIRMED | old (round 17 E3 parked it; the owner's "lanes decide" ruling of 2026-11-04 makes it the lane's) | S 376c2b210 |
| E2 | Every Elemental talent's max-rank client text against the engine: Convection 10%, Concussion 5% on Lightning Bolt / Chain Lightning / Earth Shock, Reverberation 1.0 s, Elemental Alacrity 0.5 s, Call of Flame 15% and +9 s Flame Shock, Elemental Focus 10% one charge, Elemental Warding 10%, Trial by Fire 20% / -4 s / +20 s Magma, Eye of the Storm 70%, Call of Thunder 5%, Lightning Overload 30%, Elemental Fury +100% at 5/5, Lightning Mastery -0.5 s and no Chain Lightning cooldown for 15 s on a 3 min cooldown (16166: cd 180000, dur 15000). All match talents.go, fire_nova.go, fire_totems.go. | talents/shaman.json rank text; spellconst 16166 | Aligned | - | - |
| E3 | Every spell the rotation casts against spellconst (70755 with the re-layer): Lightning Bolt 10 (15208) 2.5 s, 220 mana, 353 + 2.1 a level at 0.714; Chain Lightning 4 (10605) 2.5 s, 495, 510 + 3.3 at 0.714, 6 s; Flame Shock 6 (29228) 410, 292 + 3.8 at 0.214 up front and 80 a tick at 0.1, 6 s category; Earth Shock 7 (10414) 450, 531 + 3.4 at 0.386; Searing Totem 6 (10438) 170, 55 s, its Attack 47 at 0.017 on a 2.2 s cast; Magma Totem 4 (10587) 650, 20 s; Mana Spring 4 (10497) 100, 5 min; Fire Nova 5 (408345) 375, 10 s category. All match the engine's tables (spellconst_damage_test.go guards the rolls). | spellconst/shaman.json | Aligned | - | - |
| E4 | Chain Lightning's jumps: the client's SpellEffect rows for every rank (421, 930, 2860, 10605) carry EffectChainTargets 3 and EffectChainAmplitude 0.7; chain_lightning.go bounces to 3 targets at 0.70 a jump. Multi-target only; no round had checked the amplitude. | wago SpellEffect 1.60.1.70755 ($R/shaman/dl/SpellEffect.csv) | Aligned | - | - |
| E5 | Pre-pull (Mana Spring at -4 s, Lightning Bolt at -2.5 s), Lightning Mastery with trinkets and Blood Fury, Flame Shock on the 21 s Call of Flame timer, Searing Totem within 20 yards: as rounds 18 and 19 left them. The Earth Shock filler question stays closed (round 17 E1). | | Aligned | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | S3 above (the Maelstrom Weapon metrics line). | | CONFIRMED | R19 73b7e850e | E 511ec1ffe |
| N2 | **Maelstrom Weapon's proc rate (round 14 N1, OPEN) stays open, with the client now checked both ways.** 408498's SpellAuraOptions row (223898) is ProcChance 100, ProcCharges 0 and no SpellProcsPerMinuteID in 1.60.1.70755 (ProcTypeMask 65556) and identical in Season of Discovery 1.15.9.70576 (ProcTypeMask 20): the chance lives in a server script in both games, so the client cannot tell 50% a hit (the engine, matching the dummy effect's 50) from Season of Discovery's 10 a minute. No new log: beta report 2814 (2026-11-04 12:24 UTC, "Test") has no encounters and 2815 onward do not exist; the one beta shaman with the talent (report 2748) swung a 3.2 s two-hander, which both models fit. At stake +0.14% / +0.09% at raid 60 (round 14). Settle: a live log of an Enhancement shaman at 34 or above on a one-hander faster than 2.5 s. | $R/shaman/dl/SpellAuraOptions.csv, sod-SpellAuraOptions.csv (wago, 2026-11-04); foreverlogs reports 2812 to 2817 read 21:07 UTC | OPEN | - | - |
| N3 | **Improved Ghost Wolf 2 in the raid build, re-decided under the do-nothing rule.** Moving its two points to Ancestral Knowledge 5/5 is +0.43% at raid 60 (horde 862.8 / 862.8 against 859.1 / 859.1; alliance 860.3 / 860.4 against 856.6 / 856.8; 10,000 iterations, seeds 1 / 50001). Improved Stormstrike 2 instead is 0.00% (a raid shaman under Judgement of Wisdom and Mana Spring never casts on an empty bar).<br>Web research: every Forever raid build found takes Ancestral Knowledge 5/5 and Improved Stormstrike 2/2; two of four also take Improved Ghost Wolf (stillsixty https://www.stillsixty.com/classes/shaman/enhancement "Instant Ghost Wolf for raid movement and corpse runs"; Wowhead's Drakanthos guide, branch B https://www.wowhead.com/forever/guide/enhancement-shaman-leveling-guide-for-wow-forever-35070) and two do not (wowforeverbuilds https://wowforeverbuilds.com/guide/enhancement-shaman-pve-guide; NeverFF https://www.neverff.com/wow-forever/specs/shaman/enhancement), all read 2026-11-04.<br>Decision: keep Improved Ghost Wolf. It is a utility point the guide explains (instant, indoor Ghost Wolf when the boss moves), the same class of pick as the owner's "Elemental keeps Improved Reincarnation", and the community is split. The guide's Ancestral Knowledge line now says where the two points go for a raid that never asks you to run. Not a bug. | 10,000-iteration runs in the session log; enhancement.md:55 | Aligned (decided by research; cost recorded) | - | S 376c2b210 (prose) |
| N4 | S2 above: the Earth Shock gate holds at 40 and 46. | | Aligned | - | - |
| N5 | Proc rows against the client: Windfury Weapon's enchant proc 439431 is 20%, 1,500 ms ProcCategoryRecovery, mask 20 (white swings and melee specials), two weapon-damage hits 439440 / 439441 (effect 58); Flurry's buff 16257 has 3 charges spent on mask 4 (white swings) with a 500 ms recovery, its trigger 16256 on mask 20; Elemental Focus's Clearcasting 16246 one charge. windfury_weapon.go and talents.go match (rounds 10 and 13 said the same). | SpellAuraOptions rows 235799, 128111, 128910, 127152 | Aligned | - | - |
| N7 | **Guide-accuracy round 17 Sims requested 1 (integrator follow-up): the 25% Earth Shock gate against no gate at raid 60.** Guide raid build, frozen BiS, 10,000 iterations, seeds 1 / 50001: shipped gate horde 859.1 / 859.1, alliance 856.6 / 856.8; no gate (Earth Shock on cooldown whenever affordable) horde 859.1 / 859.1, alliance 856.6 / 856.8. Identical to the tenth; Earth Shock casts a fight 29.90 / 29.86 / 29.87 / 29.83 against 29.90 / 29.86 / 29.87 / 29.83 (the gate holds back about one shock in 500 fights). Time under 25% mana on the shipped rotation: a probe that casts Earth Shock only while mana is below 25% lands 0.00 casts a fight on both factions (10,000 iterations, seed 1), so the bar never sits under a quarter while a shock is ready; out of mana 0 s of 180.<br>Call: in a raid Earth Shock never waits. enhancement.md's Mana paragraph takes guide round 17's text (one rule, "stop shocking at a quarter", the raid as a closing sentence) with the clause round 17 allowed on this result: "so Earth Shock never needs to wait and you shock freely down to that quarter". The gate stays in the rotation (it is the rule without raid buffs). | `esgate-r20.txt`, `rot/enh-nogate.json`, `rot/enh-below25.json` (lane dir); guide-accuracy round17/shaman.md SHAMAN-R17-1 | Aligned (guides-handoff; prose) | - | S ada66984c |
| N6 | Every Enhancement talent's max-rank text against the engine: Thundering Strikes 5% all, Ancestral Knowledge 10% Intellect, Mental Dexterity 150%, Elemental Weapons 20% / 40% / 15%, Shamanistic Focus 45%, Flurry 30% for 3 swings, Stormstrike 150% and one 20% charge for 12 s, Mental Quickness 30%, Improved Stormstrike 100% / 50% regen 15 s, Maelstrom Weapon 20% a stack to 5 for 30 s, Rage of the Farseer 30% for 25 s (425336: 3 min, aura 342 melee and ranged haste only; its second effect is aura 61, the model scale). All match. | talents/shaman.json; spellconst 425336, 17364 | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | S4 above: the Chain Heal and Healing Wave rows hold at every level judged. Riptide rank 1 at 50 to 59 still loses (round 19 S3); no row. | | Aligned | - | - |
| R2 | Every Restoration talent's max-rank text against the engine: Improved Healing Wave -0.5 s, Tidal Focus -5% mana and +5% hit, Mindfulness 50%, Totemic Focus -25%, Tidal Mastery 5%, Restorative Totems +25% Mana Spring / +50% Healing Stream, Healing Way 25%, Purification 10%, Mana Tide 12 s at 3 s a tick (four ticks; the engine's 13 s aura lands the same four), Nature's Swiftness 3 min, Water Shield 3 globes of 2% on a "few seconds" recovery (3,500 ms in SpellAuraOptions), Riptide 445 over 15 s (89 x 5) and +25% Chain Heal. All match talents_restoration.go, mana_tide_totem.go, water_shield.go, riptide.go. | talents/shaman.json; spellconst 1325256, 1239243, 16188 | Aligned | - | - |
| R3 | Cast spells against spellconst: Healing Wave 9 (10396) 3.0 s, 560 mana, 1344 + 5.5 at 0.857; Lesser Healing Wave 6 (10468) 1.5 s, 380, 820 + 4.2 at 0.429; Chain Heal 3 (10623) 2.5 s, 405, 490 + 3.3 at 0.714, EffectChainTargets 3 at amplitude 0.5 on every rank (1064, 10622, 10623); Riptide 3 (1239243) 230 mana, 12 s category, 161 + 3.7 a tick at 0.1; Healing Stream 5 (10463) 80 mana, 5 min; Mana Tide 4 (1325256) 60 mana, 13 s, 5 min. All match. | spellconst/shaman.json; wago SpellEffect | Aligned | - | - |
| R4 | Raid 60 unchanged (757.1 horde, 759.4 alliance). | after.tsv | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | `TestGuideBuildsSpendNothingOnDeadTalents` (sim/leveling) now takes a class per case and reads both `build:` and `raidBuild:`; the druid cases are unchanged. Shared test file, no behaviour outside the test. | sim/leveling/dead_points_test.go | tool | - | S 376c2b210 |
| X2 | The fork's geared tests still panic "No item with id" in a fresh worktree without the item database (round 19 X1). Only the isolated Maelstrom test and the site's engine-backed tests were run. | | note | - | - |

## Before/after (10,000 iterations, frozen BiS, seeds 1 / 50001)

Before = $R/before.tsv (engine da54b7b06, site 04d7da663); after = engine 511ec1ffe on site 376c2b210 (7c76d2835 after it moves no number; the guide raid builds through $R/rt). Restoration is effective HPS.

| Spec, level | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Elemental raid 60 | horde | 666.1 to 666.1 | 0.00 / 0.00 |
| Elemental raid 60 | alliance | 669.4 to 669.4 | 0.00 / 0.00 |
| Enhancement raid 60 | horde | 859.1 to 859.1 | 0.00 / 0.00 |
| Enhancement raid 60 | alliance | 856.6 to 856.6 | 0.00 / 0.00 |
| Restoration raid 60 | horde | 757.1 to 757.1 | 0.00 / 0.00 |
| Restoration raid 60 | alliance | 759.4 to 759.4 | 0.00 / 0.00 |
| Elemental bare 20 (walked band talents, 5510000000000000 to 5500100000000000) | horde | 28.8 to 30.2 | +4.9 / +4.5 |
| Elemental bare 20 | alliance | 29.5 to 29.9 | +1.4 / +1.0 |
| Elemental bare 30 to 60 | both | unchanged | 0.00 |

## Files to regenerate

- sim/request/testdata/ladder/shaman-elemental.golden.md: the ladder's talent strings at 20 and 30 move (Elemental Alacrity, then Trial by Fire and Warding 2, in place of Elemental Devastation and Warding 3). `web/src/lib/guides/leveling-builds.test.ts` ("the guide leveling builds match the sim's ladder", shaman-elemental) reads that golden and fails until it is regenerated; builds.test.ts and _prose.test.ts pass now.
- data/builds/1.60.1.70755/addon-builds.json, then addon/ForeverSixty/Data.lua, and the planner's default-build vectors: the Elemental `build:` and `raidBuild:` changed.
- Coach starters, if they carry the Elemental builds.
- The Elemental BiS band entries' talent strings (bands 20 and 30) move with the next re-rank; not touched here.
- No apl, rank-choices or spellranks copy changed.

## Measure at refresh

- Elemental band 20 (the walk's 11th point is now Elemental Alacrity) and the band 20 / 30 talent strings in the BiS entries.
- Carried: Enhancement band 40 with the half-mana reserve; the Enhancement relic slot; Restoration bands 40 and 50 and bare 60 with the Chain Heal rows; Enhancement band 30 with Stormstrike.

## For other lanes / integrator

- Nothing here moves a spec outside shaman. The engine commit is shaman-only (sim/shaman/talents.go); raid 60 is identical for all three specs.
- The site commit touches the shared `sim/leveling/dead_points_test.go` (a shaman case added; the druid cases are as they were).
- Platform: the fork's UI and spec-breakdown no longer list a 100% "Maelstrom Weapon" aura; the dormant rank marker reports 0 uptime and 0 procs, which spec-breakdown's aura table already hides.

## For the guide lane

Shipped with this lane (captures in $R/shaman/cap):
- elemental.md builds and the talents paragraph (Trial by Fire 2, Elemental Warding 2, why Elemental Devastation gets nothing): HOLD on the first capture (the toll sentence miscounted "two of them do real work"; the Mindfulness and Improved Reincarnation sentences read as more toll points; a "drops Magma Totem" phrase in the brief to the reviewer looked like a contradiction with the priority list, which it was not), fixed to "two in Trial by Fire and two in Elemental Warding" with the Restoration sentences in their own paragraph (S 7c76d2835): SHIP on the re-capture (cap/ele-tolls-2.png). Optional SHOULD not taken: a lead-in for the Mindfulness paragraph.
- enhancement.md Ancestral Knowledge line (where the Improved Ghost Wolf points go for a raid that never moves): SHIP (cap/enh-talents-full.png).
- enhancement.md Mana paragraph (N7: round 17's text plus "Earth Shock never needs to wait"): SHIP (cap/enh-mana.png). Optional SHOULDs not taken: "five talent points", and naming the paladin behind Judgement of Wisdom. Optional SHOULD not taken: "a little more damage" as a number; prose keeps measured figures out.

## Still open

- Pending owner, not reopened: Lightning Overload's rank spells; Healing Wave 2 tooltip on a geared 60; level-60 book ranks (Strength of Earth 5, Grace of Air 3, Healing Wave 10).
- Maelstrom Weapon's proc rate (N2): a live log of a level 34+ Enhancement shaman on a fast one-hander settles it.
- The low-rank proc curve (priest P1) stays unmodelled; no shaman row depends on it (round 19 S9).
- Logs: last looked 21:07 UTC on 2026-11-04. foreverlogs beta reports end at 2814 (empty); `api.foreversixty.gg/v1/reports/recent` lists one beta report (2026-10-20). Live logs start after the 23:00 UTC launch. Warcraft Logs is blocked (gcloud login expired); not worked around.
