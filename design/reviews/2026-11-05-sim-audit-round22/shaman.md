# Sim-accuracy round 22: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70790. Site branch wt/r22-shaman (off int/r22 b179f211c), fork branch lane/r22-shaman (off int/r22). Gear is the frozen BiS ($R/bis-frozen); raid 60 runs the guide raid build ($R/rt). "Before" is this branch's own baseline ($R/shaman/sb0, built before any change, so the shared 3% crit change is in both sides). Sources read: round 21 shaman report and integration "Open, for round 22", guide-accuracy rounds 22 to 24 (shaman, on origin/main), the platform round 22 report (P1), the MythicSim hybrids lane, Blizzard's Forever notes thread (https://us.forums.blizzard.com/en/wow/t/2360696, all 10 posts, read 5 Nov ~09:00 UTC: posts 7 and 8 of 26/27 Oct are racials, post 10 of 4 Nov is UI only; no new shaman line since round 21).

**No confirmed engine or rotation bug this round.** One confirmed card bug is handed to platform (X2: Stormstrike missing from the Enhancement card at 25 to 39). The fork's stat-weight expectation was stale (seed, fixed). Elemental's three card wordings shipped after a wow-player NO-SHIP and fix.

## Seed

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **TestElemental Phase4-6 StatWeights-Default: the Hit weight (5.72 / 5.71, Phase 6 6.88) is right; the expectation was stale.**<br>Not amd64-only: it fails the same on arm64 at 170a59c39. Cause: 7c8f710d0 (round 17, this lane) keeps measuring hit past the 9% weapon cap while spells still miss up to the spell cap. Elemental enables auto attacks, so its hit profile is "physical"; the Phase 4 and 5 sets sit past the weapon cap, where the old window returned 0 although Lightning Bolt still misses (raid 60 Lightning Bolt shows 6.5% misses). Phase 6 sat exactly at the cap (window -1..0 before, -1..+1 now). Against 170a59c39 only these three numbers differ in the results file. Sanity: the published raid weights put 1% hit at 11.9 spell power. | `sim/shaman/elemental/TestElemental.results`, `sim/core/hitprofile.go:79` | test stale (not a sim bug) | R17 7c8f710d0 | E 8a9f33f74 |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **Sub-20 spell penalty (platform P1): the rank rows do not depend on it being absent for heals; Lightning Bolt 3 does for damage.** The engine applies no sub-20 cut (platform: three Forever sources for heals, damage OPEN-low). Lightning Bolt rank 3 (spell level 14) is the row at 26-27 and 30-33. At 26, rank 3 beats rank 5 by 4% (33.6 against 32.3, bare, 3,000 iterations) on mana; a 22.5% cut of rank 3's coefficient would cost it about 5% of its damage, which would make the two about level. So those two rows stand only while damage has no sub-20 cut, the same OPEN platform holds. | `data/curated/rank-choices.json`; spec-breakdown `-preset bare -band 20 -level 26 -rank 15208=3/5` | OPEN (platform P1: a sub-20 damage-rank tooltip at two spell damage values) | - | - |
| E2 | **Round 21's three card SHOULDs, applied on main's round 23/24 wording, then a wow-player NO-SHIP and fix.** Flame Shock now says "from 25, with all three points of Call of Flame" (the talent is taken 23 to 25); Chain Lightning "for the same cast time as a Lightning Bolt, until Lightning Mastery makes the bolt faster"; the top-rank bolt's first sentence is a full rule. The filler first carried the level 60 ranks in its first sentence (the guide-r24 SHOULD); the wow-player found that wrong inside leveling rows (the band 30 row headed "Rank 5" then named ranks 10 and 8). It now reads "at the rank your mana can carry to the end of the fight", true at every band, and the 60 ranks follow in the rest of the note. **wow-player: SHIP** on the re-captured Play it panels (bands 20, 30, 60; captured with a local, uncommitted patch of addon-data's first sentences; the card reads addon-data, so the live card changes only after the regen). Remaining SHOULD: the band 60 panel's first sentence no longer names rank 10 / rank 8 (the guide page does). | `data/curated/apl/shaman-elemental.json` lines 2, 5, 7, 8; captures `$R/shaman/cap/` | wording (guide lane items) | guides-handoff | S 0a7c63a7a, S 6903cd50b |
| E3 | Fresh: Lightning Bolt ranks 1 to 10 (amount, per level, cast 1.5/2.0/2.5, cost, coefficient 0.429/0.571/0.714), Chain Lightning 4 (510, 2.5 s, 495, 6 s), Flame Shock 6 (292 + 80 a tick, 0.214/0.10), Searing Totem Attack (47, 0.017, 2.437 s logged interval), Magma Totem 4, Fire Nova 408428, Earth Shock 7 against spellconst 70790: aligned. Wowhead's Forever tooltips compute at the spell's MaxLevel (Lightning Bolt 10 344-383 = 353 + 2.1 x 5; Earth Shock 534-562 at 65), so their ranges read higher than a level 60 player's; the engine's level 60 values are right. | `lightning_bolt.go`, `fire_totems.go`, spellconst | Aligned | - | - |
| E4 | Raid build: Elemental Warding 2 and Trial by Fire 2 stay as the tolls to tier 4 (round 20 decision; Eye of the Storm, Reverberation and Elemental Devastation also do nothing at raid 60). Not reopened. | talents 70790 | Aligned | - | - |
| E5 | Chain Lightning on one target re-checked on reasoning (70790 made it 2.5 s and 495 mana): round 15 measured -1.9% while mana is ahead and -9.7% on cooldown before that change; it can only be worse now. | round 15 E4 | Aligned | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | **MythicSim lead (+1.25%) does not hold on the current engine.** Their talents (453003102-053030310051002251-02: Convection 4, Call of Flame 3, Trial by Fire 2, Improved Stormstrike 2 for Reverberation 3, Tidal Focus 3, Mindfulness 3, Improved Ghost Wolf 2) with their list (Magma Totem, Fire Nova while mana covers 20 a second left, Flame Shock upkeep) ported onto our current APL: horde raid 60 825.2 / 824.4 against the guide's 839.5 / 841.1 (-1.7% / -2.0%, 10,000 iterations). Their list on our talents: 802.1 (-4.3%, 3k); their talents on our list 815.7. It also drops Improved Ghost Wolf (unmodeled), so it fails the adoption rule twice. The lead was measured before Maelstrom Weapon doubled (round 21 N1): with stacks on every hit, Tidal Focus's spell hit and Mindfulness carry more. | `$R/shaman/mkenh.py`, `$R/shaman/rot/enh-*.json` | Not adopted | - | - |
| N2 | **Guide r24 sim requested: the quarter-mana Earth Shock gate against no gate** at raid 60: identical at both factions and both seeds (horde 839.5 / 841.1, alliance 833.7 / 833.9), 0.0 s out of mana. "Earth Shock never needs to wait" holds. | `$R/shaman/rot/enh-es-nogate.json` | Aligned (guide clause stands) | - | - |
| N3 | Near ties the 3% crit change could flip: Lightning Bolt at four stacks still loses (835.0 / 835.2 against 839.5 / 841.1). Mindfulness does something (without it 822.8, -1.9%, 3k), so no dead point. | spec-breakdown | Aligned | - | - |
| N4 | Fresh: Maelstrom procs equal landed melee hits (97.2 a fight: white, Windfury Totem and Stormstrike), Stormstrike 6 s / 8% base mana, Rage of the Farseer 180 s (1.49 casts in 180 s fights), Rockbiter 7 at 60 = 554 + 16.5 x 6 (Wowhead's 686 is its MaxLevel 62), Windfury Totem an aura beside Rockbiter (Forever), Flurry 78.8% uptime, glancing 40%, dodge 6.6%. | breakdown; client rows; nether.wowhead.com/forever/tooltip/spell/16316 | Aligned | - | - |
| N5 | **Stormstrike at 25 to 39: the engine casts it** (bare band 30, level 30, leveling talents with Stormstrike: 29 casts, 25.5% of damage). The guide-r20-to-24 carried "RequiredLevel 40" note is a card issue, not a sim one: X2. | spec-breakdown `-band 30 -level 30 -talents 0-055030310031` | Aligned (sim) | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Sub-20 penalty: Healing Wave 2 (20-25) and 3 (26-36) rows rely on no cut; heals are settled by three Forever sources (platform P1). | platform P1 | Aligned | - | - |
| R2 | Fresh look at the raid healer fight: Healing Stream (5 min totem, 2 s ticks, crit on the every-table crit only, logged), Riptide, Healing Way 8/17/25, Purification 2%/rank, Tidal Mastery 1%/rank all as earlier rounds settled; nothing in 70790 or this round touches them. | `-heal` breakdown; `talents_restoration.go`, `water_totems.go` | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Blizzard notes: no shaman line after 3 Nov. Racials in posts 7/8 (Dwarf Artisan's Precision +1% hit, Troll Berserking 15 s, Orc Blood Fury 2 min) are in `sim/core/racials.go`. | thread 2360696 | Aligned | - | - |
| X2 | **The Enhancement card leaves Stormstrike off every band from 25 to 39.** Both Enhancement orders take Stormstrike at 25, the engine casts it, and Forever players use it at 25 to 30 (forum thread https://us.forums.blizzard.com/en/wow/t/2375184, 7 Oct: a level 30 beta Enhancement's "full rotation ... flame shock > SS > fire nova > earth shock"; "stormstrike costs 125 mana for us, no matter if it's level 25 or level 60"; Icy Veins and Mobalytics Forever leveling guides put it at 25). The card's rows start at 40: since client 70753 `spellranks.json` tracks Stormstrike 17364 at its client spell level 40 (before: the SoD rune 410156 at 1), so `addonrotation._family_ranks`'s talent fallback (tier 3 = 25, `test_a_talents_own_spell_joins_the_rotation_from_its_tiers_level`) no longer applies, and `trainables/shaman.json` also lists it at 40. A narrow fix (talent spell keeps its tier's level) was written and reverted here: it trips `test_the_committed_rotations_never_show_an_ability_before_its_trainer_level`, which reads the committed trainables, so it needs trainables.py and a regen, and it moves hunter Counterattack (30 to 25), priest Vampiric Embrace (30 to 25), paladin Turn Evil and Improved Mend Pet too. Handed to platform. | `data/builds/1.60.1.70790/spellranks.json` (70707 vs 70753), `addon-data.json` ability_levels 17364: 40, `data/pipeline/addonrotation.py:153-161` | CONFIRMED (card) | new-data (70753) + guides-handoff (guide r20 to r24 carried it) | - (platform) |

## Before/after (10,000 iterations, frozen BiS, raid 60, guide raid builds)

No commit on this branch changes a rotation, a talent string or the engine's damage: the APL commits are notes only, the fork commit is a test expectation. Branch baseline (int/r22 with the shared 3% crit change), seed 1 / 50001:

| Spec | Alliance | Horde | vs $R/before.tsv |
|---|---|---|---|
| Elemental | 653.8 / 654.1 | 650.5 / 651.1 | -2.3% (crit change) |
| Enhancement | 833.7 / 833.9 | 839.5 / 841.1 | -2.8% / -2.9% (crit change) |
| Restoration (HPS) | 742.5 / 742.6 | 740.6 / 740.5 | -2.2% (crit change) |

## Files to regenerate

- data/builds/1.60.1.70790/addon-data.json and addon/ForeverSixty/Data.lua: the Elemental Flame Shock, Chain Lightning and both Lightning Bolt notes (the card and Play it panel read them from there). After the regen, re-capture the Elemental Play it panels at bands 20, 30, 60 and check they read as `$R/shaman/cap/playit.txt`.
- Fork: sim/shaman/elemental/TestElemental.results also moves with int/r22's 3% crit change (every DPS row and the Phase 6 weights fail on int/r22 for that reason, not this lane's); this lane's commit fixes only the three hit weights against 170a59c39.

## Measure at refresh

- Elemental raid 60 BiS wears Thunderbrew's Boot Flask (744: 8 Spirit and a fire-breath use) in a trinket slot, and Enhancement raid 60 has no relic: both point at the ranker's item pool (client_unconfirmed drops), not the sim.
- Carried: Enhancement bands 40 to 60 and the relic; Elemental band 20/30 talent strings; Restoration bands 40, 50, bare 60 Chain Heal rows; Enhancement band 30 with Stormstrike (its published band talents lack it).

## For other lanes

- **Platform:** X2 (talent spell levels in the card and trainables; a regen). The Elemental hit profile is "physical" because the spec enables auto attacks although its APL never swings; harmless now (the spell-cap branch measures it) but a caster flag would be cleaner.
- Nothing on this branch moves DPS for any spec, in or outside shaman.

## For the guide lane

- The Elemental note wording is done here (E2) with a wow-player SHIP; main's round 23/24 wording of the Searing Totem and Chain Lightning notes is carried into this branch, so the merge of the Chain Lightning note line will conflict: take this branch's line.
- Still SHOULD: the band 60 Play it panel's filler sentence could name rank 10 / rank 8 at 60 only (needs a band-aware note, not possible in one APL note). The wow-player also asked for a sanity check of "3 again from 30 to 33" (it is round 19's measured mana row).

## Still open

- Maelstrom Weapon flat 100% against 20 a minute (fast two-hander log); sub-20 damage penalty (E1); pending owner items as in the brief.
- Logs, last looked 09:20 UTC 5 Nov: api.foreversixty.gg lists only the 20 October report; foreverlogs.gg's API answers "API access not allowed"; the Warcraft Logs v1 key in Secret Manager now returns 401 "Invalid key specified" on every host (tell the owner).
