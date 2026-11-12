# Sim accuracy 2026-11-12: Arcane mage against the client and real fights

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-arcane`. Site from origin/main a38608c67, fork at the pin 15e2c6bae. Client on the site: 1.60.1.70790; live client 1.60.1.70888 with its hotfix push. Arcane had no dedicated engine-versus-client pass before this one: the 2026-11-11 casters lane did Fire, Frost and Balance, the 2026-11-11 hotfix lane changed Arcane Concentration's split, and round30-open moved Arcane's Frostbolt rank rows (MR49).

**Totals: 2 confirmed bugs, both old, neither new data; 1 build moved. One engine: a mage had its first mana gem at 23 (the Mana Agate item's required level) where Conjure Mana Agate trains at 28, so the leveling walk at 23 to 27 popped a gem the class cannot make yet (no band entry moves; the three mage rank rows at 24 to 27 were re-judged on the fixed engine and every one holds). One tool: `rotation-search` wore the guide's leveling build under the raid preset, so every raid line gated on a raid-build talent (Arcane's Pyroblast, Fire Blast and Arcane Wrath) was dead and the probe called removing them "no effect"; it reads `raidBuild` now, as `talent-search` does. Every other engine row for the Arcane rotation matched the client; the first three Arcane logs on foreverlogs agree with the engine on Missile Barrage (procs on the cast, a missed Blast included), the Missiles cadence (1 s, 0.5 s under Barrage, five missiles from rank 3, three at rank 1, the first one period after the channel starts) and Clearcasting's sources. One build moves under the adoption rule: the leveling build takes Arcane Wrath at 44 for the fourth rank of Incineration (+1.6 to +6.8% at 45 to 55, +0.4 to +1.0% at 58 and 60, both factions), with its guide, card and coach lines. Engine pin: 15e2c6bae to fc2ed8cd3, the fork master commit that merges this lane's 28524da82 with the warlocks lane's a63d30e26 (which landed on the site while this lane ran and does not carry the gem fix on its own).**

## Method

1. **Blizzard since round 29.** "WoW: Forever Hotfixes" (t/2403040, the 12 Nov copy the rogues lane cached): the mage lines are the 9 Nov Arcane Concentration split and Flamestrike ticks no longer consuming Clearcasting, and the 10 Nov Arcane Blast duration fix ("not properly 18 seconds") and the Research spell. The split landed on 2026-11-11 (M1); Flamestrike is AoE; the engine's Arcane Blast buff has been 18 s since 70707's text.
2. **Client and push.** The rogues lane's fresh wago.tools CSVs for 1.60.1.70888 with `useHotfixes=1` against the 70790 hotfixed set, over every SpellClassSet 3 spell, every mage talent rank and two hops of their trigger spells (756 spells; `arcane/client_diff_mage.py`, the rogue lane's script with the mage class ids). Fourteen tables.
3. **Engine against the 70790 tables**, every spell, talent, aura, proc, set bonus and trinket the Arcane card casts or the two guide builds take (`sim/mage/arcane_blast.go`, `arcane_missiles.go`, `arcane_wrath.go`, `missile_barrage.go`, `talents.go`, `evocation.go`, `mana_gem.go`, `armors.go`, `shoot.go`, `items.go`, `sim/core/mana.go`) against `data/builds/1.60.1.70790/spellconst/mage.json`, `talents/mage.json`, `trainables/mage.json`, `items.json`, `sets.json` and the raw SpellEffect, SpellAuraOptions, SpellMisc, SpellCooldowns and SpellCategories rows.
4. **Real fights.** Our API: report gxrgvz4yzsi2 (8 Nov) has 289 mage units, none in the logging party; the only mage casts it carries are passing Frostbolts (40 at rank 5) and Blinks, no Arcane mage. foreverlogs.gg: 25 public reports (2815 to 2853); every mage in the rosters of 2821, 2823, 2834, 2849, 2850 and 2853 and in the cached participants of 2815 to 2850. Arcane mages with a known spec: Magie (2823, BFD, 7 Nov, Arcane 14 / Frost 3 at level 26, 1,126 casts) and Caster (2834, ESW and City of Dalaran, 8 Nov, Arcane 23 at level 32, 637 casts); by cast mix, Rat (2821, BFD, 7 Nov: Arcane Blast, Arcane Wrath rank 1, Arcane Missiles rank 3, about level 27) and Bonely (2849, 11 Nov, a level 8 to 15 mage on rank 1 Missiles and a wand). Every event of each (damage, casts, avoidance, resources, buffs; `arcane/fetch.py`, `arcane/analyze.py`). All three Arcane logs predate the 9 Nov Arcane Concentration split. The 60 to 70 mages of unknown spec in 2849, 2850 and 2853 were not fetched (the report API rate-limits at a few hundred requests); none is at 56 or above by the dungeons logged.
5. **Sims** on the Windows box (`tools/remote-test.sh --dir sim --engine-sha`, one job at a time, `nice -n 15`, GOMAXPROCS=5): the fork's mage suite (`-tags with_db`, skipping the results-file suite the module cache cannot write), the three mage rank rows at 23 to 27 on the fixed engine (`spec-breakdown -level N -rank`, 10,000 iterations, seeds 1 and 50001, both factions; `arcane/rows.sh`, output `arcane/rows.out.txt`), the `rotation-search` suite, and the talent and rotation searches at raid 60 and bare 60 on the fixed tool (`arcane/searches.sh`, output `arcane/searches.out.txt`).

## Client: 70790 against 70888 (hotfixed)

Of the 756 mage spells, only two NPC spells are new in 70888 (Frost Armor 1283419, a 10 s cooldown aura with a 6136 Chilled trigger; Blast Wave 1291176, a 15-point AoE). No row of any spell or talent the Arcane card uses changes. The Arcane Concentration split is server side (11213 unchanged, as the hotfix lane found) and stays a hand value until a client or push row states it (MR46).

## Engine against the client (1.60.1.70790)

| area | engine | client | verdict |
|---|---|---|---|
| Arcane Blast | 32/78/101/168/236 at ranks 20/30/40/50/60, variance 0.151, 0.5/0.8/1.0/1.3/1.6 a level to 28/38/48/58/68, coefficient 0.429, 1.5 s, 15% of base mana; buff 400573: 12% of Intellect as spell damage a stack, +100% cost a stack, 4 stacks, 18 s from the first cast (a stack never refreshes) | 400574, 1239696, 1239697, 1239699, 1239700; 400573 aura 174 (12, misc 126), aura 108 (100, misc 14), CumulativeAura 4, 18,000 ms; SpellMisc speed 0 | aligned |
| Arcane Missiles | channel 3/4/5/5/5/5/5 s with a missile a second, 85/140/235/320/410/500/595 mana at 8/16/24/32/40/48/56; one missile 24/34/47/66/95/130/171 plus 0.2/0.3/0.4/0.6/0.7/0.8/0.9 a level, coefficient 0.286, speed 20; the first missile one tick after the channel starts | 5143 to 10212 (aura 23, period 1000, trigger 7268 to 10274); the missiles' SpellEffect rows; SpellMisc speed 20 | aligned (the logs confirm the cadence, below) |
| Arcane Missiles rank 8 | 25345 not registered (`core.IncludeAQ` off) | trainables list it at 56; `data/curated/book-ranks.json` marks it an Ahn'Qiraj book rank the site's rank tables skip, with Fireball 12 and Frostbolt 11 | standing ruling, no change (see Held) |
| Missile Barrage | 20/40% on Arcane Blast, 10/20% on Arcane Wrath, Fireball, Frostbolt, Frostfire Bolt, rolled on the completed cast; the next Missiles free (-100 added to the other cost mods), 0.5 s ticks, 15 s; Frost Blast not registered | 400588 aura 42 to 400589, ProcTypeMask 0x10000; 400589 aura 108 (-50, misc 1), aura 4 (-500), aura 108 (-100, misc 14), 15,000 ms | aligned (the logs settle cast against hit, below) |
| Arcane Concentration and Clearcasting | 2% a rank on a landed mage spell, 1 s proc cooldown, one roll shared by a Missiles channel's cast and its missiles (9 Nov), one charge of -1000% cost on the next damage spell, 15 s, not spent by the proccing spell nor by a Missile Barrage channel | 11213 ProcCategoryRecovery 1000, ProcTypeMask 0x15550; 12536 aura 108 (-1000, misc 14), ProcCharges 1, 15,000 ms | aligned |
| Arcane Wrath | 120/187/300/451 at 25/36/48/60, variance 0.201, 1.4/2.3/2.4/3.0 a level to 35/47/59/71, coefficient 1.0, 3.5 s, 6 s cooldown, 140/220/305/395 mana, speed 24, plus 5% of maximum mana as spell damage through the coefficient | 1326079 to 1326082; effect 1 dummy 5; SpellCooldowns 6000 and category 6000 | aligned |
| Arcane Power | +30% damage and +30% cost for 15 s, 3 min, exclusive with Power Infusion | 12042 three aura 108 rows (30, misc 0, 14 and 22), 15,000 ms, RecoveryTime 180,000 | aligned |
| Presence of Mind | the next cast with a cast time instant, 3 min from when the charge is spent | 12043 aura 108 (-100, misc 10), duration -1, ProcCharges 1, CategoryRecoveryTime 180,000 | aligned |
| Evocation | from 20; 8 s channel, 8 min, spirit regen x 16 and all of it while channelling | 12051 SpellLevel 20, aura 110 (1500), aura 134 (100), aura 226 every 1000 ms, 8,000 ms, RecoveryTime 480,000 | aligned |
| Mana gems | ranks at 23/38/48/58, 375 to 425, 550 to 650, 775 to 925, 1,000 to 1,200, one a fight, 2 min shared with runes | Conjure Mana Agate/Jade/Citrine/Ruby train at 28/38/48/58 (759, 3552, 10053, 10054); the Mana Agate item 5514 says level 23, the other three gems their conjure level | **G1: rank 1 at 28** |
| Mage Armor | 34/46/58, 50% of spirit regen while casting, 5/10/15 resistances | 6117, 22782, 22783: aura 134 (50), aura 22 (5/10/15, misc 126) | aligned |
| Arcane Meditation, Arcane Mind, Arcane Instability, Arcane Impact, Arcane Subtlety, Precision, Incineration, Accelerant, Improved Fire Blast, Heating Up | 17/33/50% regen while casting; 2% Intellect and 20% Arcane crit damage a rank; 1% damage and 1% crit a rank; 2% Arcane crit a rank; 15% Arcane threat a rank; 2% hit a rank; 2% crit a rank on Fire Blast, Ice Lance, Arcane Blast and Scorch; -0.1 s a rank on Fireball, Frostfire Bolt and Arcane Wrath; -1 s a rank on Fire Blast; -25% Pyroblast cast a stack, 3 stacks, 20 s | 18462; 11232 (aura 137 10, aura 108 100 misc 15); 15058 (aura 108 1, aura 108 1 misc 22, aura 290 1); 11242 (aura 107 6 misc 7); 11210 (aura 108 -30 misc 2); 29438 (aura 55 6); 18459; 11069 (aura 107 -500 misc 10); 400625 (aura 108 -25 misc 10, stacks 3, 20,000 ms); the rank descriptions in talents/mage.json | aligned |
| Wand Specialization, Shoot | 13/25% wand damage; a shot only when out of mana | 6057 descriptions | aligned |
| Mana regen and the five-second rule | 12.5 + Spirit/4 a 2 s tick (the mage and priest formula), halted 5 s after a cast that spends mana; a free cast (Clearcasting, Missile Barrage) does not halt it | unchanged since round 8's regen pass; the client's gametables carry no regen column | unchanged |
| Bare and raid sets' bonuses and trinkets | The Postmaster 3-piece (bare 60: Band, Treads, Seal) +23 spell damage from the client's set rows; Robe of the Archmage 375 to 625 mana on 5 min; Burst of Knowledge -150 mana a spell for 10 s; Eye of the Beast, Smolderweb's Eye, Frozen Heart of the Mountain, Thunderbrew's Boot Flask, Ankh of Life registered; no set among the Pearly or Dreamweave pieces (the client lists none) | sets.json 81; item_effects.go | aligned |

The guide's talent and rotation numbers were read against the same rows: Arcane Wrath at 3.3 s with Accelerant 2/5, Fire Blast every 6 s with Improved Fire Blast 2/2, Pyroblast at 1.5 s on three Heating Up stacks, a Barrage Missiles at 30% of its cost inside Arcane Power, the Mana Ruby at 1,200 down, the Robe at 625, the gem from 28, Mage Armor from 34. Every line holds.

## Real fights

| # | check | evidence | verdict |
|---|---|---|---|
| L1 | **Missile Barrage procs on the cast, not the hit.** The engine rolls it on the completed cast (400588's 0x10000 mask reads as a hit proc). Caster, Missile Barrage 2/2: of 10 missed Arcane Blasts, 5 gained Barrage within 150 ms (45 hits: 17, 38%); Rat: 1 miss, 1 Barrage (71 hits: 24, 34%); Magie, 1/2: 13 of 48 hits (27%), no miss. In every log the buff is logged after the cast event and before the damage event of the same Blast. | foreverlogs 2823, 2834, 2821 | Aligned (the engine's model) |
| L2 | **Missiles cadence and count.** Missiles fire 900 to 1,000 ms apart on a plain channel and 400 to 500 ms apart under Barrage; a full rank 3 or 4 channel is five missiles (Magie 27 of 47 channels, the rest clipped by a dying mob or a move; Rat 19 of 24; Caster 13 of 24), a rank 1 channel three (Bonely, 101 of 182). A Barrage channel started the instant a Blast completed fires its first missile 504 ms later: one tick period after the channel starts, as the engine's Dot does. | foreverlogs 2823, 2834, 2821, 2849 | Aligned |
| L3 | **Clearcasting's sources.** Every gain (Magie 52, Caster 24, Rat 61) follows a landed spell hit (Arcane Explosion, a Missiles missile, Arcane Blast, Fire Blast, Flamestrike, Frost Nova, Arcane Wrath); none follows a wand hit alone, though 11213's mask (0x15550) carries the ranged auto-attack bit. The engine's `OnSpellHitDealt` on mage spells only. | foreverlogs 2823, 2834, 2821 | Aligned |
| L4 | **Cast mix.** Dungeon packs: Arcane Explosion is 47 to 62% of the Arcane damage at 26 to 32 and Flamestrike 14 to 17%; single-target shares (Missiles 14 to 15%, Blast 3 to 9%, Wrath 5 to 8%, wand 7 to 11%) are the card's at those levels less the AoE. Nobody at 60, nobody with Arcane Power yet. | foreverlogs | Aligned where comparable; no boss-length single-target log |
| L5 | **Mana flow.** foreverlogs carries no power or regen events (only Mana Leak and Cannibalize gains); our API has no Arcane mage. Evocation's 16x and the five-second rule stay on round 8's regen pass. | both | Open (no usable log) |

## Confirmed and fixed

| # | finding | source | fix |
|---|---|---|---|
| G1 | **A mage had its first mana gem five levels before it can conjure one.** `ManaGemLevel[1]` was 23, the Mana Agate item's required level; Conjure Mana Agate trains at 28, and the other three gems already read their conjure level (38/48/58, which the client gives the items too). On the leveling walk at 23 to 27 the bare mage popped a 375 to 425 mana gem it cannot make; the band entries (20, 30 and up) never saw it. Now 28. | trainables 759 (level 28); items.json 5514 (23); old engine | fork 71476b429 (red: `TestManaGemsRegisterAtTheirConjureLevels`, levels 23 and 27 registered rank 1) and 28524da82 (`mana_gem.go`) |
| T1 | **`rotation-search` wore the leveling build under the raid preset.** `guideOrOverrideBuild` read `GuideBuildTalents` whatever the preset; `talent-search` reads `GuideRaidBuildTalents` under raid. On Arcane's raid search the Pyroblast, Fire Blast, Arcane Wrath and Fireball lines were dead (removing each: +0.0) and the leveling Frostbolt filler carried the fight (+140.8 when removed). Every past raid rotation-search report for a spec whose `raidBuild` differs from its leveling build was measured on the wrong talents; the curated rotations themselves were authored by hand and the talent searches were right. | site `sim/cmd/rotation-search/setup.go` | site 7721be107 (red: `TestPrepareWearsTheGuidesRaidBuildUnderTheRaidPreset`) |

**Before and after.** G1 moves no band entry: band 20 sims at 20 (no gem either way) and every band from 30 has the gem either way. The three mage rank rows inside 23 to 27 were re-judged on the fixed engine (10,000 iterations, seeds 1 and 50001, both factions; `arcane/rows.out.txt`):

| row | levels | on the fixed engine | verdict |
|---|---|---|---|
| Arcane Frostbolt rank 1 (24 to 28) | 24 to 27 | rank 1 beats rank 4 by +2.2%, +1.7%, +1.6%, +1.5% (alliance) and +2.0%, +1.3%, +1.7%, +1.6% (horde) at 24, 25, 26, 27; the mage sits 27 to 30 s of 180 out of mana on rank 1 against 42 to 47 s on rank 4 | holds |
| Frost Frostbolt rank 4 (26 to 28) | 26 and 27 | rank 4 beats rank 5 by +4.4% and +4.3% (alliance), +5.1% and +4.0% (horde) | holds |
| Fire Fireball rank 4 (24) | 24 | rank 4 over rank 5 by +0.09% and +0.19% (gnome, seeds 1 and 50001; the first inside the 0.08% error), +0.33% and +0.42% (undead) | holds |

The rows' notes in `data/curated/rank-choices.json` record the re-judge; the guide text is unchanged (it already names 28 for the gem and 24 to 28 for rank 1).

## Searches

On the fixed tool and the fixed engine (`arcane/searches.out.txt`; seed 7, 800 final iterations, the committed alliance BiS gear):

| search | result |
|---|---|
| talent-search, raid 60, raid build (keep: Arcane Blast, Missile Barrage, Presence of Mind, Arcane Power, Arcane Wrath, Pyroblast, Heating Up, Improved Fire Blast, Ignite, Arcane Subtlety, Precision) | the guide's raid build is the best found (727.7 +/- 1.5) |
| rotation-search, raid 60, raid build | the curated rotation is the best found (727.7 +/- 1.5; best variant +1.0 +/- 2.2). Before T1 the same search read 670.4 with Pyroblast, Fire Blast, Arcane Wrath and Fireball dead |
| rotation-search, bare 60, leveling build | the curated rotation is the best found (406.3 +/- 1.1) |
| talent-search, bare 60, leveling build (keep: Arcane Blast, Missile Barrage, Presence of Mind, Arcane Power, Precision) | **B1:** the best build keeping every unmodeled guide talent takes Arcane Wrath for the fourth rank of Incineration, +4.6 +/- 1.6 (+1.1%); the overall winner re-spends an unmodeled point (Arcane Resilience, Arcane Geometry or Improved Channeling) on Wrath, +5.7, and is refused by the rule |

**B1, adopted under the adoption rule.** The level-60 leveling build becomes `141005213100311531/3/05350003` (Incineration 3, Arcane Wrath 1); the ladder's own rule puts Wrath at 44, the first point after Precision, and leaves every other level of the order where it was (`walk-base.txt` against `walk-wrath.txt`, 10,000 iterations, seed 1, both factions, bare band gear at the checkpoint levels):

| level | alliance | horde |
|---|---|---|
| 10 to 43 | unchanged | unchanged |
| 45 | +6.8% | +6.8% |
| 48 | +5.3% | +5.4% |
| 50 | +2.9% | +2.7% |
| 53 | +2.0% | +1.7% |
| 55 | +1.9% | +1.6% |
| 58 | +0.6% | +0.4% |
| 60 | +1.0% | +0.5% |

At 60 on both seeds (`bare60-seeds.txt`): alliance +1.0% and +1.1%, horde +0.5% and +0.4%, every cell beyond its 0.08% error. The build keeps every unmodeled talent, holds the spec tree, drops nothing the card casts (Incineration's fourth rank is 2% crit on Blast and Fire Blast) and gains at every level it touches; the horde cell at 60 alone sits under the rule's 1% while the walk above it carries +1.6 to +6.8%. The guide's build line, its Arcane Wrath, Presence of Mind, Incineration, fillers and leveling-list lines, the coach card and the rotation notes name 44; `addon-builds.json`'s Arcane leveling order moves its level-44 point; the Arcane ladder golden is regenerated; `mage-arcane` joins `knownStaleBISFiles` until the re-rank (MR68). wow-player: HOLD on the first read (the opener and movement lines still sent Presence of Mind into Frostbolt after 44; the coach card told a raid-build mage to wait for 44 and dropped the raid build from its Wrath line; a sim-flavoured comparison in the Wrath entry), fixed, then SHIP on the second. The ladder golden's level-60 rung, a mage in a weapon and a wand with no Intellect gear, reads 269.7 against 271.2 (-0.6%, inside the ladder's own 1% noise); the bare BiS band at 60 is the figure that counts.

## Held

- **Arcane Precision before 40** stays HELD for the owner (round30-open A1): V3 (Precision at 31 to 33, Arcane Power at 43) gains at every level 31 to 42 but drops a cast talent for three levels. Not re-measured here; the gem fix touches 23 to 27 only, below the levels in question.
- **Arcane Missiles rank 8, Fireball rank 12, Frostbolt rank 11.** The 70790 trainables list them at 56 and 60 (SkillLineAbility carries the Ahn'Qiraj book ranks like any trainer rank); `data/curated/book-ranks.json` keeps them off the site's rank tables and the engine's spellbook until `core.IncludeAQ`. A standing ruling, not this lane's; the first level-56 Arcane log will show which rank of Missiles a Forever trainer sells.
- **Missile Barrage from a missed Arcane Wrath** is unmeasured (Rat's 32 Wraths all landed); the Blast evidence (L1) covers the proc's trigger.

## Measure at refresh (ledger MR67, MR68)

- **MR67, the mage walks at 23 to 27 on the conjure-level gem:** mage-arcane, mage-fire and mage-frost leveling walks at 23 to 27 (any published per-level figure there); the three rank rows are already re-judged above and hold.
- **MR68, Arcane on the 2026-11-12 level 60 leveling build:** mage-arcane bare bands 50 and 60: talents (from 44), picks, weights; then drop its `knownStaleBISFiles` entry.

## Sim changes in this landing

- **Fork `lane/sim-arcane`, fast-forward of master 15e2c6bae to 28524da82:**
  - 71476b429: the gem level test, red on the box.
  - 28524da82: `ManaGemLevel[1]` 23 to 28; the mage suite green on the box (`-tags with_db`, `-skip TestP1Mage`, whose results file the module cache cannot write).
- **Site:** `rotation-search` reads the raid build under the raid preset; its suite green on the box. Engine pin fc2ed8cd3 (`sim/enginever/version.go`, `web/src/lib/sim/version.ts`, the three sim fixtures). The rank-choices notes (both copies). This report, the scripts and outputs under `arcane/`, the ledger row and MR66.

## Gates

On the rebased branch (on origin/main dfa77f3d7, the warlocks lane's landing; pin fc2ed8cd3), on the box, one job at a time (`arcane/results/gates-sim.txt`, `gates-web.txt`):
- sim: `go vet ./...`; `go test` for request (the ladder goldens, mage-arcane regenerated), leveling, rankchoices, cmd/leveling-bis, cmd/rotation-search, cmd/addon-builds, cmd/talent-search: all ok. cmd/spec-breakdown ok on the lane's fork sha before the rebase. The fork's mage suite ok at 28524da82 (`-tags with_db`); fc2ed8cd3 adds only the warlocks lane's commits to it.
- web: `npm test` on the guide, guides, sim and guide-page suites (132 files, 2,503 tests): 2,502 pass; the one failure, `src/pages/guides/_spec.test.ts` "/guides/warrior/arms rotation rows for the raid build > tags the raid build's own talent rows and no other", fails the same way on origin/main dfa77f3d7 before this landing (the melee-hunters lane's Arms build move; for that lane). `astro check`: 0 errors, 0 warnings. eslint on the touched sim and guides modules: clean.
- No Playwright spec touches a changed surface (the guide page renders the same components with new text).
