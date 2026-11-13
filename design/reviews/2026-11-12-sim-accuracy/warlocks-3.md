# Sim accuracy 2026-11-12: warlocks, third pass (Affliction and Destruction, Demonology where shared)

Lane brief: `design/reviews/sim-accuracy/brief.md`. Lane `sim-warlock3`. Site from origin/main 62e835b00, fork from master 4c595ee6b (the pin). Client on the site: 1.60.1.70790 with its push layer; live client 1.60.1.70888 with push 113766 (the 10 Nov class push), read from the wago tables the 2026-11-11 lane cached and from wago's push records fetched again today. A player called the warlock sim out; this is the third pass after `2026-11-11-sim-accuracy/warlock.md` (the hotfix model) and `2026-11-12-sim-accuracy/warlocks-logs.md` (the first post-hotfix log pass).

**Totals: 2 confirmed engine bugs, both client mismatches the fresh read found, both fixed in the fork, neither on a published number (no card casts either spell): Bane of Havoc had no global cooldown where push 113766 gives it the ordinary 1.5 s one, and the engine still registered Curse of Shadow, a spell the Forever client does not have. 1 build moved under the adoption rule: the Affliction leveling build re-spends Improved Drains (dead since the 70707 text took Wrack out of it) and the Demonology fillers on Improved Bane of Agony, full Malevolence, Bane, Nightfall and Aftermath from 42, +5.6% at bare 60 (B3). 1 stale coach card fixed (B4). 0 new data since the last pass: Blizzard's 11 and 12 Nov posts and every wago push after 113766 carry no warlock row; the two new foreverlogs reports (2848, 2854) hold the first logged Destruction warlock with Conflagrate and three more Succubi. The Succubus crit item is narrowed, not closed: the owner's armory sheet carries no crit gear, so the engine's "owner's gear crit" explanation is out; two Succubi at 24 read 0.8 and 1.3%, hers at 36 to 37 reads 3.9%, the Voidwalkers 0 of 1,238. Wrack is still unlogged: the highest warlock in all 26 public reports is 36 to 37. **

## Method

1. **Engine against the client, spell by spell.** Every spell the three cards and the four builds touch (`sim/request/apl/warlock-*.apl.json`, `data/curated/apl/warlock-*.json`, `data/builds/1.60.1.70790/addon-builds.json`) read in the fork at 4c595ee6b (`shadowbolt.go`, `immolate.go`, `conflagrate.go`, `incinerate.go`, `soul_fire.go`, `searing_pain.go`, `shadowburn.go`, `drain_life.go`, `drain_soul.go`, `death_coil.go`, `lifetap.go`, `corruption.go`, `curses.go`, `siphon_life.go`, `wrack.go`, `talents.go`, `imp.go`, `succubus.go`, `voidwalker.go`, `felhunter.go`, `pet.go`, `summon_demon.go`, `demonic_brand.go`, `items.go`, the four set files, `armors.go`, `core/dot.go`'s clock rate) against `spellconst/warlock.json` (amount, variance, growth a level, coefficient, cost, cast, cooldown, duration, period, level, max level, family mask) and `talents/warlock.json` (every rank text of all 48 talents), then every one of those ids diffed across the 70790 and 70888 SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellCategories and SpellClassOptions tables, then push 113766's 325 records (13 pages, every locale) and the 24 pushes wago lists after it (12 Nov), filtered to warlock ids.
2. **Blizzard since the last pass.** The hotfixes thread (t/2403040) to its 12 Nov posts: no warlock line after 10 Nov. forever-bugs: no warlock issue updated since 8 Nov.
3. **Logs.** foreverlogs.gg now lists 26 public reports (2815 to 2854; 2848 and 2854 are new since the last pass and were fetched with every stream for their warlocks and pets, `warlocks-3/fetchw.py` is the last lane's script). Every warlock with 20 casts or more in the ten full-stream reports (2840 to 2854) was re-read with `warlocks-3/analyze3.py`: cast mix, damage shares, DoT clocks and ticks per application, crit rates and multipliers, Life Tap counts and amounts, pet share and crit. The 11 Nov lane's damage cache of 2815 to 2839 served the pooled demon crit table. Our API's reports table, read through the Cloud SQL proxy once the owner renewed the gcloud login (the `foreversixty` DATABASE_URL secret, read-only queries, the deployed sha 9deecc23b): 7 reports in all; the five uploaded since 10 Nov (three private, two public, 12 Nov) parsed to 0 fights, and no `fight_metrics` row of any report carries a warlock, so there is nothing of ours to check Wrack or the Succubus against. The public pair the last pass read (z2opt2x2gs5u, gxrgvz4yzsi2) are unchanged.
4. **Sims** on the Windows box only (`tools/remote-test.sh --dir sim --engine-sha`, one job at a time, `nice -n 15`, GOMAXPROCS=5): the fork's warlock and core packages at the test and fix shas; `spec-breakdown` at 10,000 iterations (seeds 1 and 50001 at raid 60, seed 1 on the bare bands 20 to 60) at the old pin and the new (`warlocks-3/run.sh`, `jobs-raid.txt`, `jobs-bare.txt`, `results/`); the talent and rotation searches at raid 60 and bare 60 for both specs (`warlocks-3/searches.sh`, `results/searches.txt`).

## Engine against the client

Every number the cards use matched. The table lists the rows read; "70888" means the row is identical in both clients unless stated.

| spell or talent | client row | engine | result |
|---|---|---|---|
| Shadow Bolt 686 to 25307 | amounts 13 to 268, growth 0.2 to 1.9 a level, 0.486/0.629/0.8/0.857, costs 25 to 380, casts 1.7/2.2/2.8/3.0 s, missile 20 yd/s | `shadowbolt.go` | aligned |
| Immolate 348 to 25309 | hit 8 to 158 (+0.7 to 2.3 a level), tick 3 to 55, 0.2 and 0.13, 15 s at 3 s, costs 25 to 380 | `immolate.go` | aligned |
| Conflagrate 1293817 to 18932 | 95 to 282 at 0.2188 to 0.2418 wide, 0.429, category cooldown 10 s (SpellCooldowns 70888), costs 100 to 255, learned 25/32/40/48/54/60 | `conflagrate.go` | aligned |
| Incinerate 412758 to 1293813 | 97/145/217 at 0.15 wide, 0.714, 2.5 s, effect 1 dummy 25 (the Immolate bonus) | `incinerate.go` | aligned |
| Soul Fire 6353, 17924 | 377/431, 1.0, 6 s, category cooldown 60 s | `soul_fire.go` | aligned |
| Searing Pain 5676 to 17923 | 23 to 114, 0.429, 1.5 s, costs 45 to 168 | `searing_pain.go` | aligned |
| Shadowburn 17877 to 18871 | 66 to 266, 0.429, category cooldown 15 s, costs 105 to 365 | `shadowburn.go` | aligned |
| Drain Life, Drain Soul, Death Coil | 10 to 51 a second at 0.1; 17 to 84 every 3 s at 0.1; 272/359/454 (+2.2 to 3 a level) at 0.214, 2 min | the three files | aligned |
| Life Tap 1454 to 11689 | 20 to 420 growing one a level to MaxLevel 16 to 66, dummy effect (no coefficient) | `lifetap.go` | aligned (the logs' per-tap amounts agree again, L6 below) |
| Corruption 172 to 25311 | 10 to 73 a tick, 0.2, 12/15/18 s at 3 s, 2 s cast, costs 35 to 340 | `corruption.go` | aligned |
| Bane of Agony 980 to 11713 | 6 to 46 a tick, 0.133, 24 s at 2 s, costs 25 to 215 | `curses.go` | aligned |
| Siphon Life 18265 to 18881 | 11 to 41 a tick, 0.05, 30 s at 3 s | `siphon_life.go` | aligned |
| Wrack 1316697 | 70790: 36 +2.5 a level, 0.286, aura 271 +10%; push 113766 (SpellEffect 1349556, 1349557): 25.2 +1.75 a level to 60, 0.2002, effect 1 a dummy 100 (the speed-up); SpellMisc 862720 keeps the channel attributes | `wrack.go` hand row | aligned with the push |
| Bane of Doom 603 | 1742 at 4.0 after 60 s, 60 s cooldown, cost 300 | `curses.go` | aligned |
| Curse of the Elements 440892 to 1311680 | -30/45/60/75 resistance (misc 126, every magic school), +4/6/8/10%, 5 min, costs 50 to 200 | `curses.go`, `core.CurseOfElementsEffects` | aligned |
| Curse of Recklessness 704 to 11717 | -175/285/395/505 armor, 2 min, costs 35 to 115 | `core.CurseOfRecklessnessAura` ranks | aligned (the guide's 505 is right) |
| Curse of Shadow 17862, 17937 | **no row in any table of 70790 or 70888** (`spells.json` null; `web/src/lib/sim/buffs.ts` and `data/curated/raid-cover.json` already say so) | `curses.go` registered rank 1 at 50 and rank 2 at 60 | **bug, fixed (B2)** |
| Amplify Curse 18288 | +50% on misc 8, 3 min, 30 s, no global | `curses.go` | aligned |
| Bane of Havoc 1225228 | 15% for 5 min, 5% of base mana; 70790 has no cooldown row; **push 113766 adds SpellCooldowns 102518 (StartRecoveryTime 1500) and SpellCategories 174286 (StartRecoveryCategory 133)** | `talents.go` registered it with no global | **bug, fixed (B1)** |
| Summon Imp 688, Succubus 712 | 10 s, 80% and 100% of base mana | `summon_demon.go` | aligned |
| Firebolt 3110 to 11763 | ranks 5 to 7 25/35/44 (+0.4 to 0.6 a level), 0.571; ranks 1 to 4 the push's 6/11/15/18 (SpellEffect 681913, 685570 to 685572) | `imp.go` hand rows for 1 to 4 | aligned |
| Lash of Pain 7814 to 11780 | 16 to 50, 0.429, category cooldown 12 s, costs 65 to 160 | `succubus.go` | aligned |
| Demon Armor 706 to 11735 | 210 to 570 armor, 3 to 15 Shadow resistance | `armors.go` | aligned |
| Affliction talents | Improved Life Tap 10/20, Suppression 1% hit and 4% threat a point, Improved Corruption 0.4 s and 2% a point, Improved Drains 7/13/20 on Drain Life, Drain Soul and Drain Mana, Improved Bane of Agony 5/10, Amplify Curse, Pandemic 33/67/100 on Corruption, both banes, the drains, Siphon Life and Wrack, Malevolence 1% a point, Nightfall 2/4% on Corruption, Drain Soul, Drain Life and Wrack, Siphon Life, Soul Siphon 4/8/12 a counted effect to 3 (Corruption, Siphon Life, Bane of Agony: push 113766's text), Shadow Mastery 2% a point, Wrack | `talents.go` | aligned |
| Demonology talents | Improved Imp 10/20/30, Demonic Embrace 3% a point, Unholy Power 2% a point, Fel Vitality 5% a point, Demonic Energies 50/100 of Life Tap, Improved Sayaad 10% a point on Lash, Demonic Sacrifice 15/2%/15/3%, Master Summoner 2 s and 20% a point, Decimation 45/90 and 3/6 and 20/40, Demonic Brand 2/4/6 charges and 17/33/50 threat, Soul Link 30% and 3%, Demonic Knowledge 33/67/100% of level, Master Demonologist 2% a point, Demonic Pact | `talents.go`, `demonic_brand.go`, `soul_fire.go` | aligned |
| Destruction talents | Improved Shadow Bolt 4% a point for 12 s (aura 270 on the caster's own Shadow), Bane 0.1/0.4 s a point, Cataclysm 3/6/10, Aftermath 10% a point on the hit, Ruin 20% a point, Shadowburn, Intensity, Agonizing Flames 3/7/10 both halves, Conflagrate, Bane of Havoc, Fire and Brimstone 8/17/25, Shadow and Flame 4% and 20% a point, Incinerate; Devastation and Emberstorm are not in the tree | `talents.go` | aligned |
| Crit multipliers | Pandemic 3 and Ruin 5 both read 2.0x on a 1.5x base (`core.critMultiplierAs`) | | aligned with the logs' 2.0x ticks (L4) |
| Orc Command 20575 | effect 1 (the second 5% on Firebolt and Lash) 5 in 70790, 0 in 70888 (SpellEffect 697943, the only warlock row the client diff changes) | `core/racial_command.go` | aligned (the last pass's W6) |
| 70790 to 70888 diff, whole warlock id set | SpellEffect 697943 (Command) and SpellMisc 815375 (1262316, The Quick and the Dead, attributes); nothing else | | 0 new data beyond push 113766 |
| wago pushes after 113766 (24 pushes dated 12 Nov, build 70888) | no record carries a warlock id | | 0 new data |

The Dot clock rate (`core/dot.go` `SetClockRate`) rescales the next tick, the later ticks and the expiry on the tick grid and resets when Wrack's channel ends; the three hastened DoTs are the push's three and nothing else (W3 of the 11 Nov pass stands).

## Logs against the engine

Ten full-stream reports, 22 warlocks with 20 or more casts, owners at 10 to 37 by the ranks they cast (`warlocks-3/analysis.txt`).

| # | finding | evidence | engine | result |
|---|---|---|---|---|
| L1 | **The first logged Destruction warlock.** Ozzie (2854, 25 or 26): Conflagrate rank 1 (1293817) 25 casts, 17 hits at 91 to 125 non-crit (median 109) and 157 on a crit (1.44x), casts never closer than 9.9 s; the client's 95.9 at 24% wide plus 0.429 of a leveling spell power is 73 to 119 plus about 10. In 10 of 17 cases no Immolate tick follows the hit until a recast, as a Conflagrate without Shadow and Flame should leave it. Immolate's hit 62 to 70 and its tick 15 to 18 sit on the rank 3 row with Aftermath and his spell power. | 2854 | `conflagrate.go` (10 s category cooldown, 0.429, consumes Immolate), `immolate.go` | aligned |
| L2 | **DoT clocks and ticks per application.** Corruption every 3.0 s (median 3,000 to 3,001 ms on 2,629 ticks), Bane of Agony every 2.0 s (6,373), Immolate 3.0 s (1,407), Drain Life 1.0 s (282), Drain Soul 3.0 s (70); full runs are 6 Corruption ticks at rank 3 (206 of 561 runs, the rest cut by the kill), 12 Agony ticks (153 of 477 at rank 2), an Immolate hit plus 5 ticks (70 of 234). | all ten | the dot configs | aligned |
| L3 | **DoT ticks crit, at spell crit, 1.5x without Pandemic and 2.0x with it.** Corruption 6.8%, Agony 6.9%, Immolate 6.7%, Drain Life 6.4% of ticks; Valpuri's ticks 2.02 to 2.05x (Pandemic 3), Lacra's 1.81 to 1.85x (two points), the rest 1.5x. | pooled | `core.critMultiplierAs`, Pandemic | aligned |
| L4 | **Nightfall.** Marinko 33 Shadow Trances on 677 Corruption, Drain Life and Drain Soul ticks (4.9%), Lacra 9 on 289 (3.1%); pooled 4.3% against the talent's 4% at two points. | 2848, 2850 | `talents.go` (every tick rolls) | aligned |
| L5 | **Life Tap.** Every tap in a report is the same amount to the point (Sträyed 226 on 52 and 218 on 8 after a Spirit buff fell off; Valpuri 174 on 91; Det 127 on 4; Ozzie 137): the rank's amount plus one a level above its own plus Spirit, as `lifetap.go` has it. Rates 0.5 to 8.6% of casts on dungeon pulls with drinking between, against the bare cards' tap-below-10%-mana over 180 s with no drinking (explained in the last pass, L10). | all ten | `lifetap.go` | aligned |
| L6 | **Demon spells.** Firebolt crit 296 of 4,949 (6.0%; engine 5.9%), Lash of Pain 76 of 1,314 (5.8%; engine 6.4%), both 1.5x. Firebolt's rank 2 to 4 amounts sit on the push rows (last pass L1). | pooled | `pet.go`, `imp.go`, `succubus.go` | aligned |
| L7 | **Succubus melee crit (the last pass's O1).** Three Succubi now: Zahnia (Sträyed) 47 of 1,206 swings (3.9%) at owner level 36 to 37; Vilna (Nibbles) 4 of 519 (0.8%) and Jhornia (Marinko) 6 of 445 (1.3%) at owner level 24 to 27 (2848). Two Voidwalkers (Void, ChuckNorris) crit 0 of 1,238. The owner's foreverlogs armory snapshot (2845 and 2853, the one warlock snapshot the site returns) carries Intellect 117 to 118, Agility 50, Spirit 91 to 104, 21 to 22 Demonology points, and fifteen green and white items with no crit on any of them (Silk Headband, Sorcerer Collar, Greenweave Mantle, Filigreed Shadow Gown, Ivycloth Sash, Watcher's Leggings, Sage's Boots, Elder's Bracers, Ivycloth Gloves, Skull Ring, Seal of Sylvanas, Worgenbane Talisman, Pillager's Cloak, Crescent Staff, Charred Wand). | 2845, 2848, 2853 | `pet.go` `demonMeleeCritPercent` 0.7 for every demon plus the owner's gear crit, suppressed to 0 against a boss | **held with evidence (H1)**: her 3.9% is not his gear; the two level-24 Succubi sit near the engine's rate and the Voidwalkers below it, so the rate is a Succubus-and-level matter (her pet's skill against level 22 to 30 dungeon mobs, or an Agility term the engine does not carry), not a flat demon rate. At raid 60 the engine's Succubus already crits 2.9% from the gear crit she inherits after the boss's suppression; a Succubus rate three points higher would move her 13% share of Demonology's damage by about 0.1% of the total. Needs a level 60 Succubus log on a boss. |
| L8 | **Wrack, the doubled DoT clock and the raid-60 Affliction card.** Nobody casts 1316697 in any of the 26 public reports; the highest warlock is Sträyed at 36 to 37 (Lacra, the one Affliction player with Amplify Curse and Shadow Trance, is 30 to 36). | all 26 | W1 to W3 of the 11 Nov pass | **held, unlogged (H2)** |
| L9 | **Pet share and the leveling style.** Imp players 9 to 15% (Firehook 9.0, Punker 9.7, Valpuri 10.5, Lacra 11.6, Ozzie 15.0, Det 15.4), Succubus players 21 to 36% (Marinko 21.3, Nibbles 26.1, Sträyed 35.7 on Searing Pain and Demonic Brand). Firehook and Nibbles shoot their wands for 876 and 873 of about 1,100 casts. | 2847 to 2854 | the bare breakdowns (below) | the Succubus share at 20 and 30 in the bare cards sits inside the logged range once the wand-heavy players are set aside; the Demonic Brand style stays the last pass's O2 |

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| B1 | **Bane of Havoc had no global cooldown.** The 70790 tables have no cooldown row for 1225228 (gcd_ms 0 in spellconst), and the engine registered it GCD-less on that. Push 113766 adds SpellCooldowns 102518 (StartRecoveryTime 1500) and SpellCategories 174286 (StartRecoveryCategory 133): the ordinary global, as every other curse has. No single-target card casts it, so no published number moves; a player's own two-target sim and the leveling walk's Havoc line now pay the global. | wago push 113766 (10 Nov) | new data | E b80063d67 (`talents.go`); test E f53b2b7be `TestBaneOfHavocTakesTheGlobalCooldown`, red on the box (`results/box-fork-tests-red.txt`) |
| B4 | **The Affliction coach card was stale against the 11 Nov rotation** (Bane of Doom at sixty, Agony only without Doom, the old Amplify Curse and Nightfall levels). Prose only; the API's coach reads it. | the card against `data/curated/apl/warlock-affliction.json` | old (the 11 Nov lane moved the guide and the card, not the coach) | site, this landing |
| B2 | **Curse of Shadow was still registered** (rank 1 at 50, rank 2 at 60) though 17862 and 17937 have no row in any table of either client and the site already says Forever deleted it (`web/src/lib/sim/buffs.ts`, `raid-cover.json`, the class guide). No card or preset casts or applies it, so nothing published moves; a player's own rotation could cast a spell that does not exist. | the client tables | old | E b80063d67 (`curses.go`, `warlock.go`); test E f53b2b7be `TestCurseOfShadowIsNotRegistered`, red on the box |

## Builds and rotations (adoption rule)

The talent and rotation searches at raid 60 (raid builds) and bare 60 (leveling builds), alliance BiS gear, seed 7 (`results/searches.txt`; the `rotation-search` reads the raid build under the raid preset since the Arcane lane's fix):

| search | result |
|---|---|
| rotation-search, Affliction, raid 60 | no wrong and no missing line; the one accepted mutation recasts Corruption at 3 s or less instead of under 3 s, +6.3 +/- 1.2 (+0.9%); measured again below on every DoT at 10,000 iterations it clears the rule (R1, adopted) |
| rotation-search, Destruction, raid 60 | the curated rotation is the best found (806.5 +/- 1.3; best variant +0.0 +/- 1.8) |
| rotation-search, Destruction, bare 60 | the curated rotation is the best found (495.4 +/- 0.7; best variant +0.0 +/- 1.0) |
| rotation-search, Affliction, bare 60 | no wrong and no missing line; the one mutation (Siphon Life after Immolate) is +1.7 +/- 0.6 (+0.4%), under the rule's 1%; the Shadow Trance bolt line is dead on the old leveling build (no Nightfall) and live on the new one |
| talent-search, Affliction, raid 60 (keep Wrack, Siphon Life, Amplify Curse, Nightfall, Demonic Sacrifice, Suppression) | the guide's raid build is the best among the builds that keep every unmodeled guide talent (686.4 +/- 0.9); the overall winner re-spends the two Grim Reach points on Improved Life Tap (+0.3%) and is refused. Improved Shadow Bolt earns its five points (+11.0 +/- 2.2, the 11 Nov reviewer's doubt closed), Nightfall +16.2, Demonic Sacrifice +29.7 |
| talent-search, Destruction, raid 60 (keep Conflagrate, Shadowburn, Incinerate, Bane of Havoc, Suppression) | the guide's raid build is the best found (806.5 +/- 1.3) |
| talent-search, Destruction, bare 60 | the guide is the best among the builds that keep every unmodeled guide talent; the overall winner re-spends Molten Skin, Destructive Reach, Grim Reach and Soul Harvest on Pandemic, Malevolence, Cataclysm and Aftermath (+3.1%) and is refused: the guide keeps Molten Skin and Destructive Reach for the quests on purpose |
| talent-search, Affliction, bare 60 (keep Wrack, Siphon Life, Amplify Curse, Suppression) | **B3, adopted:** the best build keeping every unmodeled guide talent (Grim Reach, Soul Harvest) re-spends Improved Drains 3, three of Unholy Power's five and the three survival-only Demonology fillers (Improved Health Funnel, Demonic Aegis, Improved Voidwalker) on Improved Bane of Agony 2, Malevolence 3 more, Nightfall 2 and Aftermath 5: 432.1 +/- 0.4 against 409.0 +/- 0.4, +23.1 (+5.6%) |

**B3, the Affliction leveling build.** Improved Drains went into the order when its text named Wrack (70685); the 70707 text took Wrack out and the engine followed, so the three points at 42 to 44 have credited nothing since, and the fillers at 45 to 60 were spent ahead of damage talents the tree still had. The new order from 42: Improved Bane of Agony 42 and 43, Malevolence 44 to 46, Bane 47 to 51, Nightfall 52 and 53, Aftermath 54 to 58, Unholy Power 59 and 60; the level 60 build is `25522020135201351/0002/005005`. It keeps every unmodeled talent, holds the spec tree and drops nothing the card casts (no drain is in the rotation; Immolate stays in it, so Bane and Aftermath pay). The walk at the checkpoint levels, bare band gear, 10,000 iterations, seed 1, both factions (`results/walk-base.txt` against `results/walk-new.txt`):

| level | horde | alliance |
|---|---|---|
| 10 to 40 | unchanged | unchanged |
| 43 | +1.9% | +1.9% |
| 45 | +3.0% | +3.0% |
| 48 | +4.1% | +4.1% |
| 50 | +4.7% | +4.6% |
| 53 | +4.8% | +4.7% |
| 55 | +3.6% | +3.6% |
| 58 | +4.7% | +4.7% |
| 60 | +6.0% | +5.6% |

Every cell beyond its 0.03 to 0.11 DPS error; the order's own levels come out as written (addon-builds' ORDER line in `walk-new.txt`).

The guide's build line, its Talents bullets (Improved Bane of Agony, Malevolence, Nightfall, Bane and Aftermath), the raid-build paragraph and the leveling list name the new levels; the card's Shadow Trance note covers both builds; `data/curated/ladder-first.json` carries the order (and `sim/leveling`'s copy), the walk test pins 45 to 59, the dead-points test asks for Nightfall 2, Malevolence 5 and Improved Drains 0, the Affliction ladder golden is regenerated (its rungs, a bare-kit character in a weapon and a wand: 40 107.8 to 107.4 on R1's earlier recast, 50 163.0 to 167.7 and 60 296.8 to 310.5 on the build, with the Shadow Trance bolt now cast at 60), and `warlock-affliction` joins `knownStaleBISFiles` until the re-rank (MR93). The round 17 ruling that held Nightfall at 0 ("the leveling rotation never casts Shadow Bolt") predates the 11 Nov card's Shadow Trance line.

**R1, the 3 s or less recast.** The raid search's hint, measured with every DoT's recast at 3 s or less instead of under 3 s (`aff-le3.json`, `jobs-le3.txt`, 10,000 iterations, seeds 1 and 50001):

| fight | under 3 s (shipped) | 3 s or less | gain |
|---|---|---|---|
| raid 60 horde, seeds 1 / 50001 | 749.80 / 750.57 | 758.83 / 758.96 | +1.20% / +1.12% |
| raid 60 alliance, seeds 1 / 50001 | 759.60 / 760.07 | 768.68 / 768.90 | +1.20% / +1.16% |
| bare 60 horde, seeds 1 / 50001 | 403.12 / 403.12 | 403.11 / 403.13 | unchanged |
| bare 40 horde, seeds 1 / 50001 | 180.55 / 180.56 | 180.31 / 180.29 | -0.13% / -0.15% |
| bare 50 horde, seed 1 | 262.50 | 262.35 | -0.06% |
| bare 20 and 30 | the clause is gated on knowing Wrack | | unchanged |

Errors 0.05 to 0.28 DPS. **Adopted (R1):** the raid preset, the one a player called out, gains 1.1 to 1.2% beyond error on both seeds and factions, nothing the card casts is dropped, bare 60 is level, and the bare 40 and 50 cells give back 0.2 DPS. The difference is a DoT left with exactly one tick's worth: under the old line it lapsed a tick into the channel; under the new one it is refreshed before it. The card's three DoT lines, the Amplify Curse check, their notes and panel text, the guide's DoTs paragraph and the coach card now read 3 s or less; `apl-sync` carries the rotation to the fork (06822d834) and the site copy; the Affliction ladder golden is regenerated on it.

**The coach card (B4, prose).** `data/curated/coach/warlock-affliction.md` still played the pre-10 Nov card: Bane of Doom at sixty, Bane of Agony only while no Doom was up, Amplify Curse from twenty-nine or twenty-four, the Shadow Trance bolt from thirty-four. It now says what the 11 Nov card and this build do: Agony all fight at sixty, the three DoTs recast under 3 s before a Wrack from forty, Amplify Curse from 22 (20 in the raid build) on the next Agony refresh, the Shadow Trance bolt from 43 in the raid build and 52 while leveling, cast between two Wracks.

## Before and after

Raid 60, guide raid builds, published BiS, 10,000 iterations, seeds 1 / 50001, pin 4c595ee6b against b80063d67 (B1 and B2 alone; `results/before-*.txt`, `after-*.txt`; the breakdowns differ only in the order of their zero-damage rows), then the landing's pin 06822d834 with the R1 card (`final-*.txt`, seed 1):

| spec | horde | alliance |
|---|---|---|
| Affliction | 749.80 / 750.57, unchanged | 759.60 / 760.07, unchanged |
| Destruction | 878.41 / 878.91, unchanged | 888.03 / 887.86, unchanged |
| Demonology | 801.40 / 801.69, unchanged | not run (no shared change) |

Bare, horde, seed 1: Affliction 20 to 60 and Destruction 20 to 60 unchanged to the hundredth (60.60, 103.23, 180.55, 262.50, 403.12; 61.88, 106.06, 183.06, 296.10, 490.93). At the landing's pin with the R1 card (`final-raid-1.txt`, `final-bare-1.txt`): Affliction raid 60 758.83 horde and 768.68 alliance (+1.2%), bare 40 180.31 (-0.13%), bare 50 262.35 (-0.06%), bare 20, 30 and 60 unchanged; Destruction and Demonology identical to the before column. The B3 build reaches the bare bands only through the re-rank (MR91); the walk table above is its measure. The raid 60 fights the report compares with the logs: Affliction is Wrack 36.6%, Corruption 24.9%, Bane of Agony 23.4%, Siphon Life 9.7%, the Nightfall Shadow Bolt 5.4% (6 casts on 7.7 trances), 2.6 Life Taps; Destruction is Incinerate 33.0%, Immolate 14.3%, Corruption 12.4%, Conflagrate 10.8%, Bane of Doom 10.6%, the Succubus 9.9% (her swings crit 2.9% from the gear crit she inherits), Shadowburn 6.8%, Bane of Agony 2.2%, 14.3 Life Taps (8% of casts), Shadow and Flame up 96% and 92%. Bare 20 and 30: the Succubus is 16.7% and 15.0% of Affliction's damage, 16.0% and 14.2% of Destruction's; Life Tap is 18% of casts; Destruction's Shadowburn fires 9.8 times a fight, a shard each before Shadow and Flame.

## Held and open

- **H1.** The Succubus's melee crit (L7): not the owner's gear; a Succubus-and-level rate the engine does not carry. About 0.1% of Demonology's raid DPS either way. Needs a level 60 Succubus on a boss.
- **H2.** Wrack, its DoT speed-up and the under-3 s recast rule (L8): no level 40 or higher Affliction log exists in any public source.
- **Soul Shards.** The engine keeps no shard count. The Destruction card casts Shadowburn on cooldown from 20; before Shadow and Flame (35 to 39 in the leveling order) every cast spends a shard, 12 on a 180 s fight, and the card note says to bring them. At raid 60 Shadow and Flame 5 refunds every one. Stated, not modelled; no change.
- Our own uploads: none since 10 Nov holds a fight (five reports, 0 fights each); a 60 Affliction or Destruction upload with advanced logging would close H1 and H2 at once.
- The last passes' leftovers stand: a Searing Pain and Demonic Brand leveling style for Demonology (O2), the level-by-level Affliction talent list, the Execute idea.

## Measure at refresh (ledger MR92, MR93)

- MR92: Bane of Havoc's global at the client move to 70888 (issue #21): `talents.go` carries push 113766's SpellCooldowns row by hand beside `wrack.go`'s and `imp.go`'s (MR41); once the generated tables carry it, read them again (test unchanged). Nothing published moves.
- MR93: the Affliction leveling build (B3): warlock-affliction bare bands 40 to 60, talents from 42, picks and weights; then drop its `knownStaleBISFiles` entry.

## Sim changes in this landing

- Fork `lane/sim-warlock3` on master 4c595ee6b: f53b2b7be (the two tests, red on the box), b80063d67 (B1, B2), 06822d834 (the synced Affliction card, R1), 5515ae0ef (master bc54f7ba5 merged in: the dps warriors lane's Deep Wounds fix, which main already pinned). Site pin 5515ae0ef; the fork's master fast-forwards to it.
- Shared code: none (`core` untouched).
- Goldens: `sim/request/testdata/ladder/warlock-affliction.golden.md` regenerated on the box for B3 and R1; Demonology's and Destruction's unchanged.
- Site: the Affliction guide (build line, Talents bullets, the raid-build paragraph, the DoTs paragraph, the leveling list), `data/curated/apl/warlock-affliction.json` and its site copy, `data/curated/ladder-first.json` and `sim/leveling`'s copy, the coach card, the walk and dead-points tests, `knownStaleBISFiles`, this report, the scripts and outputs under `warlocks-3/`, the ledger row, MR92 and MR93. wow-player: NO-SHIP on the first read (three MUSTs, all taken), SHIP on the second (`warlocks-3-wow-player.md`).

## Gates

GATES GO HERE.

## Sources

- Client 1.60.1.70790 (`data/builds/1.60.1.70790/spellconst/warlock.json`, `talents/warlock.json`, `spells.json`); the 2026-11-11 lane's wago CSVs of 70790 and 70888 (SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellCategories, SpellClassOptions); wago.tools hotfix push 113766 (325 records) and the 12 Nov pushes.
- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts to 12 Nov).
- forever-bugs, https://github.com/ClassicWoWCommunity/forever-bugs (warlock issues updated since 8 Nov: none).
- foreverlogs.gg reports 2815 to 2854 (`/api/reports/<id>`, `/combatants-roster`, `/meta_actors`, `/meta_events`); the armory snapshot in the 2845 and 2853 rosters.
- foreversixty API `/v1/reports/recent`.
