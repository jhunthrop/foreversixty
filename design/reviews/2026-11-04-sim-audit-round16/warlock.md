# Sim-accuracy round 16: warlock (2026-11-04)

Specs: Affliction, Demonology, Destruction, bands 20 to 60, bare and raid. Client 1.60.1.70755 with the late hotfix re-layer. Site `wt/r16-warlock` from int/r16 (43efb9b14, then int/r16 c5915e4bb merged for wave 2); fork `lane/r16-warlock` from int/r16 (no engine change; int/r16 65c8e36cf merged). spec-breakdown at 10,000 iterations, seeds 1 and 50001, on `$R/bis-frozen` with each guide's raid build unless noted.

Totals: **2 CONFIRMED-BUGs**, both fixed on the site, no engine change.
- Demonology raid build: a point in Improved Voidwalker, which does nothing in a raid (origin `old`; held by rounds 13 and 14 as "within error" before the round-16 dead-point rule).
- Shadow Bolt rank rows measured only at band floors (Destruction 20-29 and 30-39, Demonology 20-29). Origin `R15 0bb0d948a` (the 20-29 rows) and `old` (Destruction's 30-39 row), both exposed by the per-level sweep on the MaxLevel + 12 penalty (priest E 18815822c).

## Logs (owner's standing order)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **No launch log exists yet.** `GET api.foreversixty.gg/v1/reports/recent` returns one report (z2opt2x2gs5u, 2026-10-20, 42 fights, no boss; already examined). foreverlogs.gg: `reports/public` newest is 2811 (2026-10-24, "Stocks as Beta Ends ... TF at lvl 30"); 2812 and 2813 (2026-10-29, "shield") are the last ids; 2814 to 2850, 2900, 3600, 4000, 10000 and 100000 are "Report not found". Forever launched today, so no character is near 40 (Wrack) or 60 yet. Warcraft Logs stays blocked (gcloud login), not tried. | foreverlogs.gg API, 2026-11-04 05:00 local | OPEN (re-check next round) | - | - |

## Seed items (round 15 regressions)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Wrack in the rotation and raid build: holds. Raid 60 horde 733.1 / 733.4 (= `before.tsv`). Wrack is 48.6% of the damage, 22.8 casts and 123 ticks a fight. The ticks a cast short of 6 are misses on the cast (spell miss about 9% on the published pre-raid gear; Immolate shows the same rate), not clipping: the APL sets no interrupt and nothing pushes the channel back in a DPS sim. `TestAfflictionCastsWrackWhenItsBuildTakesIt` passes (bare 50, bare 60, raid). | spec-breakdown; `sim/core/apl.go:255-290`, `cast.go:166` | Aligned | - | - |
| S2 | Bane of Doom over Agony: holds. Doom 2.41 casts (2.21 land: the misses), Agony 2.29 casts in the last minute, Amplify Curse 0.98 casts at 0.2% uptime (spent on the last Agony). `TestAfflictionRaidRotationKeepsBaneOfDoomAndAmplifiesTheLastAgony` passes. | spec-breakdown | Aligned | - | - |
| S3 | Curse of the Elements first: holds in all three curated APLs (Affliction line 2, Demonology 3, Destruction 0). `TestWarlockRotationsCurseOfTheElementsBeforeCorruption` passes. | APLs | Aligned | - | - |
| S4 | Soul Fire gated on Demonic Pact: holds. Raid 60 casts no Soul Fire; bare 60 does. `TestDemonologySoulFireOnlyWithoutDemonicPact` passes. | tests | Aligned | - | - |
| S5 | Shadow Bolt rows at 20-29: re-measured level by level in wave 2; see W1. | | (W1) | | |
| S6 | Eureka!'s client mask: unchanged and still OPEN. The live push has no row for 1259821. The engine follows the client (no Incinerate; a separate periodic +10% on Drain Life, Drain Soul and Wrack). No level-40+ Gnome Destruction log exists (L1). The guide follows the tooltip (guide round 10 D2: Wrack gets nothing on the page), so page and engine differ on Wrack for Gnome Affliction. | round 15 E1; live push 2026-11-04 | OPEN (a level 40+ Gnome log settles it) | - | - |
| S7 | **Wrack's 10% (coordinator item).** Decoded on the shipped 70755 tables with the live push layered (SpellEffect latest push 113557, SpellClassOptions 113539): 1316697 effect 1 (record 1349557, aura 271, base 10) has class mask word 0 = 0x402, words 1-3 = 0. Bit 0x2 is carried by Corruption (all twelve ids incl. 25311); bit 0x400 by Bane of Agony (980, 1014, 6217, 11711, 11712, 11713). The only other class-set-5 carriers are Invocation (426241, 426246, an SoD rune) and "Test Curse of Agony" 28608, neither castable. Bane of Doom 603 is word 0 0x80000000 + word 1 0x2, Siphon Life word 1 0x1, Wrack itself word 1 0x40000, Drain Life 0x8, Drain Soul 0x4000: all outside. The engine (`wrack.go`, `wrackVulnerableSpellCodes` = Corruption, Agony; `TestWrackVulnerabilityCoversOnlyCorruptionAndAgony`) matches. The tooltip ("your other Shadow damage over time effects") is broader; the client row decides. | wago SpellEffect / SpellClassOptions / SpellName 70755 + live push (`tools/layer.py`, `tools/mask.py` in $R/warlock) | Aligned | - | - |

## Affliction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Races on the Wrack raid build** (guide round 10's sim request). Raid 60, guide raid build, seeds 1 / 50001: Human 737.9 / 738.6, Gnome 736.7 / 737.0 (Human +0.2%); Orc 733.1 / 733.4, Undead 730.6 / 730.6 (Orc +0.4%), Troll 722.4 / 722.6 (-1.5% to Orc). The page's lines hold: Gnome "a step behind", Undead "as good as Orc" (within 1%), Troll "a step behind Orc". Gnome includes the client's periodic Eureka on Wrack (S6). | spec-breakdown -race | Aligned | guides-handoff (round 10) | - |
| A2 | Raid build `25520020135001351/00052/005004` re-decoded: every point acts on something the rotation casts or is a keep. Grim Reach 2 is range (utility; no talent left in reach does more: Fel Concentration only saves pushback, which a raid DPS sim never takes; Improved Drains and Soul Harvest touch drains and kills; Nightfall's instant Shadow Bolt measured -0.4% in round 15). | talents/warlock.json; `tools/dec.py` | Aligned | - | - |
| A3 | Talent masks re-read on 70755 + push: Malevolence 1310949 (crit +5) covers Corruption, both banes, Siphon Life, Wrack, the drains, Shadow Bolt, Shadowburn, Death Coil and the Succubus' Lash of Pain (Lash crit 10.3% Affliction / 6.0% Demonology / 8.6% Destruction with Malevolence 3, as the engine does); Pandemic 427712 and Shadow Mastery's periodic half 18271 effect 1 cover Corruption, Agony, Doom, Siphon Life, the drains and Wrack; Shadow Mastery's direct half covers Shadow Bolt, Shadowburn, Death Coil, Drain Mana; Improved Corruption 17810 (-2 s, +10% periodic) covers Corruption only. Soul Siphon 17804 is a dummy (aura 4, mask 0), so which effects it counts is server-side; the engine counts Corruption, the bane and Siphon Life. Whether the warlock's own Curse of the Elements also counts cannot be read from the client; it would not move a number, since the three DoTs already reach the 36% cap wherever Wrack is cast. | wago 70755 + push | Aligned (Soul Siphon count: note) | - | - |

## Demonology

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | **The raid build spent a point on Improved Voidwalker, which does nothing in a raid.** The raid build sacrifices the Imp before the pull (Demonic Pact keeps the sacrifice) and plays with the Succubus out, so the Voidwalker never comes out; Improved Voidwalker (18705) only touches Torment, Consume Shadows, Sacrifice and Suffering. Rounds 13 and 14 (D2) held it because the swap measured within error; the round-16 rule makes a provably dead point a bug by itself. The point moves to Improved Sayaad 2 (10% more Lash of Pain a point). Legal: the tier gates hold (t2 needs 10 below: 14; Decimation needs 15 below: 19; t4 needs 20: 21). Raid 60: horde 706.7 to 707.2 / 706.2 to 706.7, alliance 716.6 to 717.1 / 717.6 to 718.1 (+0.07% at both seeds, both factions, so it costs nothing). Demonic Aegis 2 in its place measured exactly the old number. New raid build `05/2035103021220001351/0550005`. Test first: `TestDemonologyRaidBuildSpendsNothingOnDeadDemons` (sim/leveling) failed (Improved Voidwalker 1, Improved Sayaad 1) and passes. Guide prose (demonology.md:60, :62) rewritten; wow-player review below. | talents/warlock.json 18705, 18754; runs | CONFIRMED-BUG | old (rounds 13 and 14 looked and held it) | S 0c9f72749, S cad97d206 |
| D2 | Other raid-build points: Improved Health Funnel 2 (Health Funnel is unmodeled; the guide's heavy-raid-damage Succubus funnel), Demonic Embrace 3 and Demonic Aegis 1 (survival keeps), Master Summoner 2 (the mid-fight resummon), Fel Vitality 3 (15% warlock mana, modeled). Improved Imp is already 0. Demonic Brand (1293695) needs Searing Pain casts and stays untaken. | talents.go:426, :1168 | Aligned | - | - |
| D3 | Phantom-talent check: the Soul Fire line reads `auraIsActive 440873` (Decimation's window) without `auraIsKnown`. It cannot misfire: Decimation comes at 29/31 in the leveling order (addon-builds index 19/21) and Soul Fire is learned at 48; the raid build has Decimation. | addon-builds.json; trainables 6353 | Aligned | - | - |

## Destruction

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Raid build `25500000033/0/0050355103101351` re-decoded: Bane of Havoc 1 is Incinerate's prerequisite; Cataclysm 3, Aftermath 5, Ruin 5, Agonizing Flames 3, Fire and Brimstone 3, Shadow and Flame 5 all act on cast spells. Conflagrate 15.75 casts (10 s cooldown), Shadowburn 10.7 (15 s), Incinerate 37.4; crit 39.8% on Conflagrate with Fire and Brimstone. | spec-breakdown; talents | Aligned | - | - |
| X2 | No talent-aura read in the APL; Conflagrate gates on Immolate, Shadowburn and Incinerate are casts of talent spells (dropped when untrained). | APL | Aligned | - | - |

## Cross-cutting (fresh pass)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **Live push on every warlock row.** Layered the live 70755 push over the shipped SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellCategories, SpellLevels, SpellAuraOptions, SpellClassOptions and SpellName (latest pushes 113557, 113539, 113548, 113595) and diffed every class-set-5 spell plus every talent spell: the only change is Wrack effect 0 (0.143 to 0.286, 0 to 2.5 a level), which the engine already reads. The Succubus, Imp, Demon Armor, Soul Link, Demonic Pact, Eureka! and Decimation rows are untouched. | wago 70755 + push, fetched 2026-11-04 | Aligned | - | - |
| C2 | Base points and coefficients, every rank (client): Corruption 0.2 a tick, Agony 0.133, Doom 1742 + 4.0, Siphon Life 0.05, Immolate 0.2 direct / 0.13 a tick, Shadow Bolt 0.486 / 0.629 / 0.8 / 0.857 from rank 4, Incinerate 0.714, Conflagrate and Shadowburn 0.429, Soul Fire 1.0, Drain Life 0.1: the engine's arrays match (`corruption.go:29`, `curses.go:40, :403`, `siphon_life.go:38`, `immolate.go:45-46`, `shadowbolt.go:40`, `incinerate.go:48`, `conflagrate.go:41`, `shadowburn.go:35`, `soul_fire.go:30`, `drain_life.go:32`). Wowhead Forever tooltips agree on Corruption 342/18 s, Agony 552/24 s, Doom 1742, Siphon Life 41/3 s, Wrack 86/s. | wago; nether.wowhead.com/forever | Aligned | - | - |
| C3 | Top ranks at 60: Corruption 25311 (r7), Shadow Bolt 25307 (r10) and Immolate 25309 (r8) have skill-line rows at 60; the APLs cast r6 / r9 / r7. This is the pending owner item "level-60 trainer book ranks"; not reopened. | trainables/warlock.json | pending owner | - | - |
| C4 | BiS gear on the frozen raid-60 files: no warlock set; the Highlander's / Defiler's Intent 3-piece (+1% spell crit) is one belt only. Frozen Heart of the Mountain (249469): Chthonic Power 1249110, 29 Frost and Shadow power for 15 s, 5 minute cooldown (Wowhead Forever), once a fight in the sim. Thunderbrew's Boot Flask and Uther's Strength do nothing for a ranged boss sim. | bis-frozen; items/warlock.json; wago 1249110 | Aligned | - | - |
| C5 | Demon melee crit 0.0% on a boss (0.7% base, less the level-63 suppression): the logged Forever rate (#331); demon crit at 60 stays open for a launch log. | `pet_crit_test.go` | OPEN (carried) | - | - |

## Wave 2: per-level rank rows and phantom talent lines

**W1. Shadow Bolt rows by level (CONFIRMED-BUG).** Rank 3 (MaxLevel 17) keeps its full bonus to 29 and rank 4 (MaxLevel 25) to 37 under the MaxLevel + 12 penalty, so the lower rank only wins while mana is short near a band floor. Swept every level 20 to 33 for Destruction and Demonology and 20 to 39 (and 40 to 59 for Demonology) at four points a band; the max learned rank against the two below it, Shadow Bolt pinned through `-rotation` probes, bare, `-level`, 3,000 iterations to screen and 10,000 at seeds 1 and 50001 for every level 20 to 33 (`$R/warlock/sweep10k.txt`, `sweep3k.txt`). I adopted a row where a rank wins by more than 0.2% at both seeds (about four times the run error); within 0.2% the default (highest learned rank) stands.

| level | Destruction (r3 / r4 / r5, seed 1; 50001 within 0.1) | Demonology |
|---|---|---|
| 20 | 59.8 / 59.3 | 55.0 / 54.6 |
| 21 | 60.5 / 60.1 | 55.6 / 55.3 |
| 22 | 61.2 / 60.8 | 56.2 / 56.1 |
| 23 | 61.9 / 61.6 | 56.8 / 56.9 |
| 24-25 | tie | r4 +0.2 |
| 26 | 68.6 / 68.9 | 62.9 / 63.8 |
| 28 | 71.4 / 71.7 / 71.6 | 65.7 / 66.7 / 66.0 |
| 29 | 71.9 / 72.3 / 72.6 | 66.3 / 67.2 / 66.9 |
| 30 | r4 100.2, r5 100.0 | r4 91.7, r5 91.3 |
| 31-32 | r4 = r5 within 0.2 | r4 = r5 within 0.2 |
| 33 | r5 103.4 > r4 103.2 | r5 94.1 > r4 93.9 |
| 36 / 39 | r6 ties r5 at 36; at 39 r6 115.5, r5 115.0, r4 112.3 | 39: r6 105.6, r5 105.1 |
| 40-59 | (Incinerate from 40) | highest rank wins or ties at every point |

New rows: Destruction rank 3 at 20-23 (was 20-29 rank 3 and 30-39 rank 4); Demonology rank 3 at 20-21 and rank 4 at 28-30 (was 20-29 rank 3). Gains against the old rows inside the bands: Destruction +1.0% at 29, +0.3% at 33, +2.8% at 39; Demonology +1.4% at 26-27, +0.9% at 29 (rank 4 against the old rank 3 there), +0.5% at 30 (rank 4 against the default rank 5). At the band floors the published numbers move only at 30: Destruction 100.2 to 100.0 (-0.2%, the tie), Demonology 91.3 to 91.7 (+0.4%). Both factions behave the same (no race key). Affliction: the highest rank wins at every level 20 to 39 (no row, none added). Test first: `TestWarlockShadowBoltRowsFollowTheLevel` (sim/request) failed at Destruction 26, 29, 33, 39 and Demonology 26, 29, 30, and passes. Both copies (data/curated and sim/rankchoices) changed; rank-choice tests in sim/request, sim/rankchoices, talent-search, rotation-search, spec-breakdown and the data pytest pass. No guide or card note names a Shadow Bolt rank, so no prose moves. Origin: `R15 0bb0d948a` (20-29 rows) and `old` (Destruction 30-39, round 15 W1 kept it on a floor measurement). Fix: S 71792db98.

**W2. Phantom talent lines: none.** Every warlock line checked:
- Affliction: no line reads a talent aura; Amplify Curse and Wrack are casts of talent spells (no talent, no spell, line dropped); the Doom gates read `dotIsActive 603`, which below 60 drops to the Agony-only gate the notes describe at every band.
- Demonology: the Soul Fire line reads Decimation's window (D3; Decimation always precedes Soul Fire); the prepull Imp, sacrifice and Soul Link lines gate on `auraIsKnown 425464` or are casts of the talent spells.
- Destruction: see X2.

## Before/after

Raid 60 (seed 1 absolute; %, seeds 1 / 50001):

| Spec | Faction | Seed 1 | % |
|---|---|---|---|
| Demonology (guide raid build, D1) | horde | 706.7 to 707.2 | +0.07 / +0.07 |
| Demonology | alliance | 716.6 to 717.1 | +0.07 / +0.07 |
| Affliction, Destruction | both | 733.1 / 776.4 unchanged | 0.00 |

Bare floors: Destruction 30 -0.2%, Demonology 30 +0.4%; every other floor unchanged. Mid-band gains in W1.

## Files to regenerate

- `data/builds/1.60.1.70755/addon-builds.json` (Demonology raid order: Improved Voidwalker out, second Improved Sayaad in); `TestCommittedAddonBuildsFollowTheGuides` fails until then.
- `addon-data.json` / Data.lua (warlock rank lines from the new rows) and the coach starters.
- Ladder goldens: warlock-demonology and warlock-destruction at levels 20 and 30 (rank rows).

## Measure at refresh

- Demonology raid 60 picks and weights on the Improved Sayaad build (Lash of Pain is pet spell damage; tiny).
- Destruction and Demonology band 20 and 30 weights and BiS (Shadow Bolt rank moved at 30).

## For other lanes

- None. No shared code changed; nothing outside the warlock specs moves.
- Platform: the addon resolves ranks only at its rungs (10/20/30/38/40/50/60), so Demonology's 28-30 rank 4 reaches the addon at 30 only (the shaman lane's note covers it).

## For the guide lane

- Done in this lane: demonology.md:60 and :62 (D1), wow-player SHIP.
- land/guides-r10 (not yet in int/r16) rewrites affliction.md and the warlock card notes; it does not touch the lines this lane changed. The Affliction APL's top-level `notes` and `engineering_notes` "5" still call Wrack the "#2 damage source behind Shadow Bolt"; they are not rendered. Left for the guide lane's merge so the files do not conflict.

## Still open

- Launch logs (L1): Wrack, Doom at 60, demon crit, Eureka on Incinerate (S6), Soul Siphon's count (A3).
- Demonology Horde race (owner question 1). Level-60 trainer book ranks (pending owner, C3).
- Siphon Life tick crit (held for the owner), Ignition's rate, shadow wands (round 14 list, unchanged).

## wow-player review (Demonology page, rendered at 1440)

**SHIP** on this lane's change, with no MUST. The reviewer checked the raid planner (Improved Sayaad 2/3, Improved Voidwalker 0/3, Improved Imp 0/3, Pact 1/1), the "moves 13 points" line and both prose passages against each other. The first pass said NO-SHIP for two reasons that came from the capture: the dev server could not load the Svelte islands or the icon files. A clean recapture fixed both.

Not this lane's, but a real contradiction a player would see (coordinator, please route): the Demonology Priority card shows "Soul Fire, every 6 seconds while Decimation is up" as step 1, while the prose says to skip it in a raid. The APL line is gated on no Demonic Pact, so the raid build never casts it, but this tree's card reader drops that gate. Two fixes have already landed on other branches and are not in int/r16 yet: main b53e82298 ("rotation cards read every gate the APL sets") and land/guides-r10's Soul Fire note ("In the raid build, skip it: ..."). Two SHOULDs, both minor:
- the planner's "2 more talents (2 points)" should name Improved Voidwalker and Demonic Energies (platform component);
- a clause on why Improved Sayaad stops at 2 of 3 (optional).
