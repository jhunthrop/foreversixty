# Sim audit round 9: ranker (2026-11-02)

This lane owns `sim/cmd/leveling-bis`, the BiS pipeline code and `data/curated/content-phases.json`. It worked the four assigned items and two items the integrator routed here: X1 (the stale engine-effects list) and the ranks lane's per-preset rank choices. It then re-ranked all 28 specs at every band (20 to 60, raid and bare at 60, both factions) and audited every pick against the game and the owner's class rules. It made no engine change.

**Setup**
- Site: `wt/r9-ranker`, branched from origin/main 55a1212f7.
- Engine: the pin 40021374a for items 1 to 4. The final re-rank uses the items lane's `lane/r9-items` 13b4768d3, because that is the engine X1 is meant to read (Dormant Heart of the Mountain, `core.NewClientItemEffect`).
- Scratch: `scratchpad/r9/ranker/`. No output of any run is committed.
- Loot data. The committed loot.json carries the closed gates' `opens`, and the ranker keeps the stricter of a file's `opens` and the curated gate (`TestAContentGateKeepsAStricterComputedOpens`). So opening a gate in the curated file changes nothing until loot.json is rebuilt. I rebuilt loot.json offline with `python -m pipeline loot-merge` into two scratch copies of the build:
  - `loot-before`: origin/main's curated file;
  - `loot-after`: this branch's curated file.

  Both go through the same pipeline code and the same engine checkout, so their difference is the gate alone. (Neither copy equals the committed loot.json byte for byte: pipeline changes since the last data run move 1,052 lines on their own.)
- Item data. The `reputation` and `limit_category` fields this branch adds to `items/<class>.json` only reach the committed files at the next data run. For the final re-rank I wrote them into a scratch copy of the class files, from the client tables: ItemSparse 70707 from wago.tools, layered with the committed `raw/hotfixes/ItemSparse.csv`.
- Simdb. The committed `simdb.bin` still types a held-in-off-hand item as a main hand (round 8's note; the next data run fixes it). That artifact turned out to matter far more this round (K10). So the before/after comparison of picks runs on round 7's held-typed `simdb-held.bin` on both sides. The performance proof (item 4) runs on the committed simdb on both sides.

**In short: 9 CONFIRMED-BUGs, all fixed with tests first**
- K15: a set's cheapest pieces and its noisy screen;
- K16: a wand left behind by verify;
- K17: Forever zone reputations with no gate;
- K8: reputation standing and side;
- K9: Unique-Equipped families;
- K13: a quartermaster's side;
- K11: the stale effect list;
- K14: the bare-60 preset;
- K12: the round-8 tank probe. This one was a defect in the evidence: the ranker itself was right, and the open note is withdrawn.

The two owner rulings are done (K1, K2). Tank sweeps were already deterministic; the finding that they were not came from round 8's probe (K12). The ranker's work drops by 17% (sims run) and its wall time by 36% side by side (CPU 20%), with output identical to the byte (item 4).

## Assigned items

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K1 | **Owner ruling 1: the battleground gates open at launch.** `battleground-reputation` (WSG, AB and AV reputation, the quartermasters, the AV summons and quests) and `darkspear-islands` (Creeg Bothunk) move from `later` to `launch`, 2026-11-04, `confirmed`. Their `launch-banned.json` examples go, as that file's note asks. | Owner ruling. Blizzard sources:<br>- "How PvP Progression in World of Warcraft: Forever Works" (worldofwarcraft.blizzard.com/en-us/news/24303316, 7 Oct), on honor across every battleground map, Alterac Valley included;<br>- "Battle for the Darkspear Islands in World of Warcraft: Forever" (worldofwarcraft.blizzard.com/news/24307308, 7 Oct), levels 30 to 60, with Darkspear Raiders and Theramore Expeditionary Force reputations;<br>- "Forever goes live globally 4 November".<br>Fetched 2026-11-02. None of the three posts says "at launch" in so many words; the ruling does.<br>Tests first:<br>- `TestTheBattlegroundGatesOpenAtLaunch` (Go): both gates open; a Creeg Bothunk and a Warsong Outriders source come through open.<br>- `test_the_battlegrounds_open_at_launch` (Python).<br>Both failed, then passed. | Ruling (done) | S f94b40cf2 |
| K2 | **Owner ruling 2: the rank-10 cap keeps reading `title_rank`.** The comment at `pvpRankExceedsCap` (band.go) records the ruling: the cap should switch to honor cost once the client links honor costs to items. | ItemExtendedCost 1.60.1.70707 (wago.tools, 1,101 rows) has these columns: Money, MinFactionID, MinReputation, ItemID_0..4, ItemCount_0..4, CurrencyID_0..4 and CurrencyCount_0..4. No column names the item a cost row prices; that link lives in the server's vendor table. | Ruling (done) | S c8ab58049 |
| K12 | **Round 8's "tank sweeps disagree between seeds and some fail with 'stamina weighs nothing'" came from its probe, not the ranker.** The probe (`r8/ranker/site-exp/.../zz_probe4_test.go`) called `realEngine{}.RunWeights`. That is the damage sweep, and stamina deals no damage, so it refuses to normalise (`ErrNoWeights`), and the picks it did make followed damage. The ranker sweeps tanks on `tankEngine` (`roleRunner`, main.go). There every sweep point already shares one seed with labelled random streams, i.e. common random numbers across the stat deltas (engine `statweight.go`: `UseLabeledRands`, and one `RandomSeed` for the baseline and every low and high run). | Re-probe on the tank engine (`noise1.txt`): Protection warrior, Protection paladin and bear, bands 20 to 60, 100 and 400 iterations a direction.<br>- The same seed twice: **0 differing slots in all 30 cases**.<br>- Seeds 3 and 50003: 0 to 3 slots at 100 iterations and 0 to 2 at 400. Round 8 reported 3 to 15.<br>- No sweep failed; stamina is the positive reference everywhere.<br>Two seeds still differ by a slot or two where tank items sit within error of each other. The ranker publishes one fixed seed (now the named `weightsSeed`), so two runs give the same weights and the same picks.<br>Test: `TestTheTankSweepIsReproducibleAndValuesStamina`, on the real engine, band-20 Protection warrior. Two sweeps are bit-identical; stamina is positive beyond its error; the damage engine refuses the same request. It passes on origin/main's code, which is the point: nothing in the ranker was broken. | CONFIRMED-BUG in round 8's evidence, so the open note is withdrawn. Ranker: aligned. | S d47e460f5 (test and named seed) |
| P1 | **Run time: a sim memo.** Every ranker sim is deterministic: it names its own seed and iterations, and the engine draws from labelled streams. So a repeat request is an exact repeat. `memoRunner` (memo.go) answers repeated plain, weights and Shoot sims from memory, keyed by the request with the log-only character name cleared. An unseeded request is never reused. It wraps the role runner, and `unwrapRunner` keeps the tank and healer interfaces reachable. | Memo totals over the 28 specs: 41,195 sims run and 8,563 answered from memory (17.2%). The hits come from where the two weight sets reach the same sets, and from the guard, chooseRun, the final measurement and set completion re-measuring sets an earlier step already ran.<br>Tests: memo_test.go (one sim for a repeated request under three labels; a new seed or new gear runs again; unseeded always runs; weights reused; the tank runner stays reachable). | Done | S fd2fa8ecd |
| P2 | **Run time: the two weight sets and both factions rank concurrently.** In each band, the two weight sets, and each set's two factions before and after the trinket reconcile, share only read-only inputs. On a runner safe for concurrent use they now run at once. Safe runners are the real engine (the API already runs it concurrently) and the tank, healer and memo runners over it. The memo holds an in-flight entry, so two weight sets reaching one set at the same moment still simulate it once. The nightly's `-all` runs specs one after another, so this is what fills its cores. A test double that records its calls in order keeps the sequential path. | Race detector: `go build -race`, Survival, Protection warrior and Holy priest at band 20 → 0 data races. Unit test of eight concurrent callers under `-race`.<br>`forEach` returns the first error by index, so the error reported does not depend on scheduling. | Done | S cc3c4c627 |

### Item 4 proof: same picks, less time

- **Identity.** origin/main (`A`) and this branch's performance commits (`P`, cc3c4c627: memo plus concurrency, nothing else that moves picks on today's data) ran on the same scratch build (gates closed, committed simdb, published incumbents). All 336 band rows of the 28 specs are identical. Every band row compares equal: slots, alternatives, weights, set score, report fields. Only the two run-time fields differ.
- **Work.** 49,758 sim requests; the memo answers 8,563 of them (17.2%), so 41,195 run. Per spec the share runs from 13% (bear) to 29% (Feral cat).
- **Time.** Both binaries ran side by side on the same machine (load average about 120 from the other lanes), 4 specs, one process per spec, `/usr/bin/time`. Output is identical (0 rows differ):

| spec | origin/main wall s | this branch wall s | origin/main CPU s | this branch CPU s |
|---|---|---|---|---|
| druid-balance | 815 | 447 (−45%) | 587 | 440 (−25%) |
| hunter-survival | 917 | 572 (−38%) | 789 | 646 (−18%) |
| warrior-protection | 986 | 632 (−36%) | 792 | 667 (−16%) |
| priest-holy | 1275 | 901 (−29%) | 2063 | 1648 (−20%) |
| total | 3993 | 2552 (**−36%**) | 4231 | 3401 (**−20%**) |

The two full 28-spec runs above are not a fair wall-time pair, because the load changed under them: origin/main took 5,626 s at 4 specs at a time, this branch 4,280 s at 3 at a time. On a quiet 4-core runner (the nightly) the memo's 17 to 20% is the floor. The concurrency's gain grows with free cores, up to 2 to 4 times on a band's ranking step, which is about 90% of a spec's time (the weights sweeps are the rest).

Other cheap wins I looked at and did not take:
- **Skipping the second pass when both weight sets agree.** The memo already does this, and more: it also catches the agreements partway through a pass.
- **An early-out for dominated candidates.** Picks run on score and their cost is negligible. The time is in the trinket, effect, verify and set sims, and those compare sets whose stats differ, so no candidate is dominated in a way that saves a sim without risking a pick.

## Side effects of opening the gates

The offline rebuild says exactly what the next loot.json refresh changes. It is nightly-owned and held, so it is not committed:
- **34 sources lose `opens: later`.** These are the 20 WSG, AB and AV rep sources (friendly to exalted), the 9 quartermasters (Gaelden Hammersmith, Thanthaldis Snowgleam, Grunnda Wolfheart, Jekyll Flandring, Illiyana Moonblaze, Kelm Hargunth, Rutherford Twing, Samuel Hawke and Creeg Bothunk), the AV zone (2597), and the summoned Ivus, Lokholar, Korrak and Grunnda world sources.
- **52 quest entries lose `opens: later`.** These are Proving Grounds, the AV insignia chain (Rise and Be Recognized to The Eye of Command), Korrak, Hero of the Stormpike and Hero of the Frostwolf, the WSG and AB care packages, Take Five Bases, Control Five Bases, and the Past Victories and Past Efforts turn-ins.
- **347 equippable items gain an open source.** By band: 22 at 20, 58 at 30, 60 at 40, 62 at 50, 145 at 60. 75 of them are epic. These counts are before the standing, side and rank rules below.

The rules that bound them, checked on the re-rank:
- **Standing.** Revered and exalted only at 60 (`repStandingObtainable`). The rep sources carry their standing, and vendors inherit it. Creeg Bothunk's stock did not: see K8.
- **Side.** WSG, AB and AV are split by faction in loot.json. The Darkspear stock was not: see K8.
- **Rank cap.** Unchanged: no new PvP-rank item, and `title_rank` is still read (K2).
- **Unique-Equipped families.** See K9.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K8 | **Creeg Bothunk's stock would have been listed with no standing and for both factions.** No `rep` source names his 51 items, so his vendor row states neither standing nor side, and `vendorInheritsRepStandingGate` has nothing to copy. With the gate open, the ranker treated Exalted rewards as leveling gear and gave Horde-only rewards to the Alliance. Examples: Raider Gauntlets 272096 (level 50, Exalted) and Raider Handwraps 272098. | Client ItemSparse 70707: MinFactionID 2798 and MinReputation 4 to 7 on all 51. The hotfix layer has the rows the shipped table lacks: the cloaks, Reapers, Spellblades, Tomes, Bludgeons, Raider Gloves and Handguards, and the shoulders.<br>Faction 2798 is "Darkspear Raiders", ParentFactionID 892 (Horde Forces), in Faction 70707. The Alliance rewards are separate items (2799 Theramore Expeditionary Force, 2719xx) that no source in loot.json names.<br>Wowhead Forever tooltip, 272096: "Requires Darkspear Raiders - Exalted"; 272060: "Requires Darkspear Raiders - Friendly". Fetched 2026-11-02.<br>Fix: the pipeline writes `reputation` {faction_id, standing} on each class-file row (`reputation_for`). `applyItemReputations` (reputation.go) holds every source of the item to the stricter standing, and gives a Horde or Alliance Forces reputation its side (data.go's `repSide`: loot.json's 509, 510, 729, 730, 889 and 890, plus the client's 2798 and 2799 added here, under Faction.ParentFactionID 892 and 891; see For other lanes on loot.json's swapped 889/890). Neutral gates nothing.<br>The same gate also covers any other vendor-only reputation reward the client marks.<br>Tests first: `TestAnItemsClientReputationGatesEverySource`, `TestAnItemsReputationNeverLoosensASourcesStanding`, `TestTheClassFileCarriesTheItemsReputation`, and the Python `test_reputation_for_maps_the_client_standing` and `test_client_rows_carry_their_reputation_requirement`. | CONFIRMED-BUG (client table plus Wowhead; live once loot.json is rebuilt) | S 13967ecf3 |
| K9 | **Two trinkets of one Unique-Equipped family in one set.** The rule paired items by id and name only. With the gates open, the first re-rank put Rune of Perfection and Rune of Duty on bear bands 20 and 40 (both factions), and Darkspear Voodoo Seal with Relentless Raider's Seal on bear band 50 (Horde). The game allows one of each family. | Client ItemSparse 70707 `LimitCategory`, and ItemLimitCategory 70707 (each Quantity 1):<br>- 714 "Rune of Battle": the Runes;<br>- 715 "Seal of Battle": the Darkspear and Theramore seals;<br>- 716 "Talisman of Battle": Talisman of Arathor and Defiler's Talisman;<br>- 718 "Arena Master": Arena Master and Arena Grand Master;<br>- 708 "Artisan's Tier": 73 crafted tier pieces, one of which may be worn.<br>Every category that holds equippable items has Quantity 1.<br>Wowhead Forever 272061: "Unique-Equipped: Seal of Battle (1)".<br>Fix: the pipeline writes `limit_category`. A single rule, `pairMate` (pair.go), now decides id, name and category for every paired-slot check: pick, the trinket shortlist, effect ranking, verify swaps, confirmed stats, the faction reconcile and the alternatives.<br>Tests first: `TestOneLimitCategoryFillsOnlyOneOfAPair`, `TestTheClassFileCarriesTheLimitCategory`, and the Python `test_client_rows_carry_their_limit_category`. Today's published files have no such pair. | CONFIRMED-BUG (client tables plus Wowhead; live once the gates open) | S 6c239bffd |
| K13 | **An Alliance quartermaster's unsourced stock reached the Horde.** loot.json names no faction for Illiyana Moonblaze (Silverwing Sentinels) or Samuel Hawke (League of Arathor). Their stock with a rep source inherits that source's side. Their stock without one did not: 13 of Illiyana's 75 items and 9 of Samuel Hawke's 70, among them the level-25 Season of Discovery WSG set. So with the gate open, the re-rank before this fix (D2) did two things wrong:
- **Wrong-faction picks.** It put the Alliance Sentinel's leggings on Horde band-60 rows of 19 of the 26 specs it finished, in both bare and raid. These are Forever's item-level-78 WSG Exalted legs, 237814 to 237825, from Illiyana Moonblaze, with no faction restriction and no rep source.
- **Wrong-faction alternatives.** It listed Lorekeeper's Staff (212580) as a Horde band-30 main-hand alternative.

The Horde's own copies are the Outrider's pieces (237820, 237824) and Advisor's Gnarled Staff (212584), from Kelm Hargunth. | loot.json `vendor:14753` and `vendor:15127` have no `faction`, while `vendor:14754` (Kelm Hargunth) has `horde`. The Horde copy of the staff is Advisor's Gnarled Staff 212584 (Kelm Hargunth). Wowhead Forever tooltips for 212580 and 212584 give the same stats and no faction or reputation line (2026-11-02).<br>Fix: `vendorInheritsStockSide` (data.go) gives a vendor with no side the one side its rep-sourced stock names. A vendor naming both sides, or none, is left alone. The duplicate side table I had added in reputation.go is folded into data.go's `repSide`, which gains 2798 and 2799.<br>Test first: `TestAQuartermasterTakesTheSideOfItsRepStock`. The first fix read each item's last rep source, and the full re-rank still showed the Sentinel's legs on Horde rows of 21 specs. The cause: both sides' quartermasters sell the WSG bracers, so 16 of Illiyana's items carry both reps, and she counted as both sides. Now only single-side items count; the test gained a shared bracer and failed before the change (19f41188c). Samuel Hawke's mail pieces also carry MinFactionID 509 in the client, so K8 gates them too. | CONFIRMED-BUG (loot.json plus Wowhead; live once the gates open) | S 3e01ad51a, S 19f41188c |

## Routed by the integrator

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K11 | **The ranker's list of engine-modelled item effects was stale** (items lane X1). `effectids_generated.go` was scraped from fork 912a458f7 on 2026-10-07. Seventeen picked items whose effects the engine runs read as unmodelled, so they never got an engine-verified run and the page said "not simulated". Examples: Ankh of Life, Kindling Stave, Robe of the Archmage, Toxic Revenger and Black Menace. | The items lane's report, row X1. Fix: `effectImplemented` asks the linked engine after registration (`core.HasItemEffect` or `core.HasWeaponEffect`), and the scraped map and `sim/scripts/effectids.py` are deleted. `core.NewClientItemEffect` registrations (Dormant Heart of the Mountain on lane/r9-items) count with no ranker change.<br>Test first: `TestTheEngineDecidesWhichItemEffectsAreModelled` failed for those five, then passed. The relic test's unmodelled example moves off Kindling Stave, which the engine now models, to the first client-effect item the engine does not model. | CONFIRMED-BUG (routed) | S f99a222ec |
| K14 | **The bare-60 pass simmed the raid rank choices.** The ranks lane made rank choices per preset (wt/r9-ranks c0f19a3bb, `data/curated/rank-choices.json` rows with an optional `preset`) and gave `api.SimRequest` a `preset`. An empty preset takes `PublishedPreset(level)`, which is raid at 60. The ranker set no preset, so the bare-60 rank rows never reached the bare-60 sets. | The integrator's request. I merged wt/r9-ranks (S ba98afb47). Fix: `plainRequest` and `weightsRequest` set `Preset` from the pass (`specInfo.rankPreset`: raid for the raid pass, bare for every other), so every sim a band runs uses its own preset's ranks. `presetBare` is `request.BarePreset`.<br>Test first: `TestEveryPassRequestNamesItsPreset` failed (empty preset at 60), then passed. | CONFIRMED-BUG (routed) | S 65212f9be |
| K15 | **Set completion lost Fury band 30 Alliance's Defias Leather.** Each threshold tried one combination, the cheapest pieces by score, and the 100-iteration screen demanded a 0.5% lead. With the gates open, the cheapest four took the chest (+0.7%, under the 1% confirm margin), and the best four (no chest) screened only +0.3% at 100 iterations. The band then sat 0.5% below origin/main's set. | Re-sim, 10,000 iterations, seeds 1 and 50001, final engine: origin/main 62.3 / 62.3; this branch before the fix 62.0 / 62.0; the same set with the four Defias pieces 62.9 / 62.9. At the confirm harness (300 iterations, seed 7) those four measure 62.4 against 61.6.<br>Fix: `setTrialsFor` adds up to four close alternatives (one piece traded for one left out), and `runBestTrial` confirms the best screen whenever it leads at all. The confirm keeps swapMargin and the error bar. `setScreenMargin` is deleted.<br>Tests first: `TestTrySetCompletionTriesTheCloseAlternativesToTheCheapestPieces` and `TestTrySetCompletionConfirmsTheBestScreenThatLeadsAtAll` failed, then passed; the set tests' sim counts hold.<br>After: 62.9 / 62.9 (+1.0% on origin/main). | CONFIRMED-BUG (integrator's re-sim) | S 9ee115e7b, S 91164bc93 |
| K16 | **A wand the band's final set outgrew.** A wand's score rests on the measured wand credit, verify only swaps in the runner-up, and the wand's worth moves with the rest of the set (a mana-bound Shadow priest wands more with less Intellect). Shadow band 40 Alliance scored Jaina's Firestarter first and Goblin Igniter second, and verify swapped in the Igniter. Twisted Nether Wand and Jaina's were both better in the final set. | Re-sim, 10,000 iterations, seeds 1 and 50001: origin/main 151.7 / 151.8; this branch before the fix 147.8 / 147.9 (−2.6%); the same set with Twisted Nether Wand 151.4 / 151.4.<br>Fix: `rankWandSlot` measures the top three wands by score in the set, before verify. `finalWandCheck` runs again on the verified, set-completed picks and adopts a wand ahead by swapMargin. Both run only for specs that cast Shoot.<br>Tests first: `TestTheWandSlotKeepsTheBestMeasuredOfItsTopWands` and `TestTheFinalWandCheckReplacesAWandTheFinalSetOutgrew`.<br>After: Jaina's Firestarter, 151.9 / 151.9 (level with origin/main, +0.1%). | CONFIRMED-BUG (integrator's re-sim) | S 6bc63f449, S fd2b19a4f |
| K17 | **Two Forever reputations would have opened with no gate.** The items lane's client reputation sources (F3, wt/r9-items 3ef62621d) source The Watchers (2819: sixteen level-60 signets, Friendly to Exalted) and Powderfuse (2827: six level-40 rewards) at the next data run, and no gate named either. | Powderfuse is Faction 2827 under the Steamwheedle Cartel (169). Powderfuse Port is the goblin port on the Riverglades' east coast (warcraft.wiki.gg, patch 1.60.1), and Blizzard's "Forever goes live globally 4 November" names the Riverglades among the launch destinations. So Powderfuse opens at launch, confidence **inferred**: no Blizzard post names the faction.<br>The Watchers: Faction 2819 reads "an order of powerful jailers and marshals of night elf society". No Blizzard post, client source or Wowhead Forever page (faction, or signet 275968) gives a zone or how it is earned, so it is **gated later, uncertain**, per the house rule. Warden's Signet of Ferocity is its banned example.<br>Tests first: `TestTheForeverZoneReputationsFollowTheLaunchEvidence` (Go) and its Python twin. The active-build banned test lets the one id the committed loot.json cannot have yet (faction 2819) through `AWAITING_THE_NEXT_DATA_RUN`, to be emptied after the data run. | CONFIRMED-BUG (integrator's request; live at the next data run) | S 23811cb59 |
| K10 | **The held-item simdb artifact now inflates sets.** On the committed simdb, a held-in-off-hand item takes the main hand's place in every sim. With K11, Feral and bear one-handers such as Thornspike get engine-ranked, and the off-hand fill adds a held item (Tork Wrench, Satyr's Rod) beside them. In that sim the cat swings with no weapon and shows 0% miss and 0% dodge. | `sb` on the committed simdb, Feral band 30 Horde with Thornspike and Tork Wrench: 119.5 DPS. The same set without the held item: 72.3 DPS. The table shows miss 0.0% and dodge 0.0% for autos and Shred (8.0% and 6.5% without it), and Thornspike made no difference (119.5 either way).<br>On the held-typed simdb, the full set is 72.7, below origin/main's Manual Crowd Pummeler set (86.6).<br>The first re-rank on the committed simdb published +38% (cat) and +90% (bear tank score) at band 30 from this artifact alone. | CONFIRMED (data artifact; the next data run's simdb fixes it). Not a ranker bug, but see "For other lanes". | - |

## Re-rank: picks that change

`A2` is origin/main and `D5` is this branch at 19f41188c (with wt/r9-ranks merged) on lane/r9-items; the before/after below uses `D6` (91164bc93), which adds K15 and K16. Both use the held-typed simdb and the published incumbents (so the guard is live on both sides). The loot is before and after the gate, and the class files are without and with `reputation` and `limit_category`. Set score is at the verify harness (300 iterations, seed 7); healers are HPS and tanks are their own DPS (see Tanks below).

- **Where the battleground gear lands.** Of D5's 5,175 filled slots, 823 are battleground-sourced: 84 at band 20, 64 at 30, 110 at 40, 62 at 50 and 503 at 60 (bare and raid together).
  - Rings carry most of it: Protector's Band and Legionnaire's Band (14 and 13 specs), Advisor's and Lorekeeper's Rings (13 each), Don Julio's Band (Stormpike Guard or Frostwolf Clan, 20 specs).
  - Then come the WSG medallions, Battle Healer's Cloak and Caretaker's Cape (healers), Dryad's Wrist Bindings and Forest Stalker's Bracers, the AB girdles, and the WSG Sentinel's and Outrider's legs at 60.
  - Darkspear (Horde only): Darkspear Voodoo Seal (14 specs), Darkspear Signet, Insurgent's Band, and Darkspear Pauldrons (Exalted, band 60 only).
  - The AB Talismans of Battle go to tanks; the AV insignias show up at 60 raid.
- **The checks on the final run** (audit9 and audit.py) all pass: no client-revered or client-exalted reward below 60, no wrong-side reward or quartermaster, no two items of one Unique-Equipped family, and no PvP pick above rank 10 (the PvP picks are ranks 7 to 9). The only empty slots are honest ones: trinkets with no sourced item at bands 20 to 40, relics the engine does not model, and one Horde band-20 caster neck.
- **The engine's own effect list (K11) brings in, by measurement:** Gloves of Holy Might (42 slots), Ankh of Life (37), Thunderbrew's Boot Flask (31), Dormant Heart of the Mountain (14, healers at 50 and 60, on the items lane's engine), Spire of Hakkar (3), Kindling Stave (3) and Cloak of Flames (1).
- **Rank choices (K14, the ranks lane's per-preset rows) move set scores, not gear.** Comparing D3 (the same branch before the merge) with D5: Holy paladin +30.5% mean (band 50 Alliance 157.6 → 264.2), Elemental +19.7%, Restoration shaman +15.1%, Discipline +4.8%, Holy priest +2.3%, Balance +2.4%, Restoration druid +2.2%, Fire +2.4%. Those are the ranks lane's downranking gains, measured on the ranked sets. The before/after below shows all changes together.
- **Tanks.** `set_dps` on a tank row is the tank's own damage; the ranker picks on tank score. The five tank rows that fall in the table below are damage, on sets that won on tank score.

## Fresh audit of the ranker's rules

Round 8's rules R1 to R21, re-checked on the final re-ranks (D5 and D6: 28 specs, 336 band rows each) with round 8's `audit.py` (slot, level, faction, armor at 40, weapon skill, ranged type, raid source, two-hander and off hand, dual wield, main-hand-only, shield, dagger, pairs) and this round's `audit9.py` (client standing below 60, reputation side, limit-category pairs, a quartermaster's side).

| # | rule | verdict on the re-rank |
|---|---|---|
| R1 | Class may equip (paladins train axes, druids no polearms, rogues no axes) | Aligned in code (round 8's K1). The committed `items/paladin.json` still has no axes (subclasses 4 to 8 only); the next data run adds them. Retribution's picks are maces and polearms until then. |
| R2 | Mail from 40 (hunter, shaman), plate from 40 (warrior, paladin) | Aligned: no sub-40 mail or plate. |
| R3 | Required level, raised to a quest's floor | Aligned. Darkspear stock begins at level 28; the battleground's 30-to-60 bracket matches the bands. |
| R4 | Faction | Aligned after K8 and K13: no Horde-only reward on an Alliance row and no Alliance quartermaster stock on a Horde row (audit9: 0). |
| R5 | Phase gates | Aligned. WSG, AB, AV and Darkspear are open (K1). Raids, world bosses, Dire Maul, ZG, AQ, the Faire and the late recipes stay closed (`TestNoBannedItemIsObtainableAtLaunch` passes). |
| R6 | Revered and exalted only at 60 | Aligned after K8: no client-revered or client-exalted reward below 60 (audit9: 0). |
| R7 | Rare spawns give way | Aligned. |
| R8 | PvP rank above 10 is never a default | Aligned. The cap reads `title_rank` (K2), and the re-rank's PvP picks are all ranks 7 to 9. |
| R9 | No same id, name or Unique-Equipped family in a pair | Fixed (K9). audit9: 0 limit-category pairs. |
| R10 | Two-hander means an empty off hand; one-hander gets an off hand | Aligned on the held-typed simdb: no one-hander beside an empty off hand. (The committed simdb's artifact is K10.) |
| R11 | Dual wield only for rogues, Fury and hunters; shamans never | Aligned: Enhancement wears a two-hander in every band, with no shaman off-hand weapon. |
| R12 | Main-hand-only weapons never in the off hand | Aligned. |
| R13 | Daggers where the rotation casts Backstab, Ambush or Mutilate | Aligned for Assassination and Subtlety. Combat is still caught by its Backstab fallback (round 8, still open). |
| R14 | Ranged type by class | Aligned. |
| R15 | Shields for warrior, paladin and shaman only | Aligned. |
| R16 | Set bonuses | Aligned on the picks. |
| R17 | Weights | Tank sweeps are reproducible and value stamina (K12). The bare-60 pass now sweeps under its own rank choices (K14). |
| R18 | Weapon DPS term counts the off hand in full | Unchanged (round 8, still open). |
| R19 | Trinkets, with an engine tournament | The effect list is the engine's own answer now (K11). Thunderbrew's Boot Flask, Ankh of Life and Kindling Stave enter by measurement. |
| R20 | The incumbent guard | Works. It kept the published set in the one row where a data bug misleads the score (Balance band 30 Horde, Glimmering Staff; see For other lanes). |
| R21 | Survival is melee; one build at 60 | Aligned. Survival dual-wields melee weapons in every band (Teebu's Blazing Longsword with a dagger at 60) and wears a ranged slot for Shoot. |
| - | One job per page | The ranker publishes gear, the weights it ranked by, and set figures only. No talent or rotation advice comes out of `leveling-bis`. |

## Before/after

Origin/main (`A2`) against this branch (`D6`, head 91164bc93, the same as the final head for every pick: 23811cb59 only adds gates that no source names before the next data run). 28 specs and 336 band rows, both on the held-typed simdb and the published incumbents. The loot is closed against open, and the class files are without against with the two new fields. Set score is at the verify harness (300 iterations, seed 7). Healers are HPS; tanks are their own DPS.

| spec | rows with a changed pick | mean | largest gain | largest loss |
|---|---|---|---|---|
| druid-balance | 11 of 12 | +3.91% | +14.99% (30 A, 65.9 → 75.8) | - |
| druid-feral | 11 of 12 | +2.45% | +5.71% (60 H, 277.8 → 293.7) | -0.30% (30 A, 88.0 → 87.7) |
| druid-feral-bear | 12 of 12 | +7.63% | +22.52% (30 H, 28.4 → 34.8) | -0.06% (40 A, 64.2 → 64.2) |
| druid-restoration | 8 of 12 | +2.78% | +9.98% (30 H, 103.0 → 113.2) | - |
| hunter-beast-mastery | 11 of 12 | +1.87% | +5.05% (60 raid A, 672.0 → 705.9) | - |
| hunter-marksmanship | 9 of 12 | +1.94% | +5.60% (60 A, 302.8 → 319.8) | -0.34% (30 H, 138.8 → 138.4) |
| hunter-survival | 12 of 12 | +2.28% | +7.08% (60 A, 357.7 → 383.1) | - |
| mage-arcane | 12 of 12 | +1.48% | +2.89% (60 raid A, 609.4 → 627.0) | - |
| mage-fire | 11 of 12 | +4.05% | +10.13% (40 A, 149.1 → 164.2) | - |
| mage-frost | 11 of 12 | +1.29% | +4.51% (60 raid A, 599.9 → 627.0) | -0.05% (20 H, 36.7 → 36.7) |
| paladin-holy | 12 of 12 | +33.59% | +68.76% (50 A, 156.6 → 264.2) | - |
| paladin-protection | 12 of 12 | +2.89% | +12.15% (60 A, 158.0 → 177.2) | -3.32% (60 raid H, 451.4 → 436.4) |
| paladin-retribution | 12 of 12 | +1.99% | +2.97% (50 A, 177.8 → 183.0) | - |
| priest-discipline | 12 of 12 | +5.86% | +12.54% (40 H, 146.5 → 164.9) | - |
| priest-holy | 12 of 12 | +3.47% | +9.68% (50 H, 175.0 → 191.9) | - |
| priest-shadow | 11 of 12 | +3.99% | +7.93% (60 H, 360.9 → 389.5) | - |
| rogue-assassination | 12 of 12 | +2.73% | +4.06% (60 A, 215.9 → 224.6) | - |
| rogue-combat | 12 of 12 | +2.04% | +3.25% (40 H, 110.6 → 114.2) | - |
| rogue-subtlety | 12 of 12 | +2.38% | +4.78% (40 H, 94.8 → 99.4) | - |
| shaman-elemental | 10 of 12 | +21.83% | +51.32% (60 H, 152.4 → 230.6) | - |
| shaman-enhancement | 12 of 12 | +1.51% | +4.18% (40 H, 123.5 → 128.6) | - |
| shaman-restoration | 11 of 12 | +16.06% | +24.77% (40 H, 147.9 → 184.5) | - |
| warlock-affliction | 10 of 12 | +2.30% | +6.00% (20 A, 55.4 → 58.7) | -0.21% (40 A, 138.4 → 138.1) |
| warlock-demonology | 12 of 12 | +1.90% | +4.63% (60 A, 332.9 → 348.4) | -0.24% (40 A, 145.0 → 144.6) |
| warlock-destruction | 12 of 12 | +1.41% | +4.28% (60 raid A, 656.7 → 684.8) | - |
| warrior-arms | 12 of 12 | +1.76% | +5.10% (60 A, 247.4 → 260.0) | - |
| warrior-fury | 12 of 12 | +2.38% | +3.90% (60 A, 278.3 → 289.1) | - |
| warrior-protection | 12 of 12 | +1.94% | +10.63% (60 A, 183.1 → 202.6) | -2.86% (40 A, 96.8 → 94.0) |

rows > +0.5%: 274; rows < -0.5%: 5; total 336

**Five rows fall more than 0.5%, all tank rows on own damage.** The same five sets, re-measured on the tank score they rank on (3,000 iterations, seeds 1 and 50001, final engine):

| row | origin/main tank score | this branch tank score |
|---|---|---|
| Protection paladin 20 Alliance | 1465.8 / 1463.0 | 1520.8 / 1524.4 (+3.9%) |
| Protection paladin 40 Alliance | 6365.0 / 6368.3 | 6649.2 / 6645.0 (+4.4%) |
| Protection paladin 60 raid Horde | 41889.0 / 41832.8 | 43906.2 / 43963.5 (+4.9%) |
| Protection warrior 40 Alliance | 8964.2 / 8984.5 | 9314.5 / 9332.6 (+3.9%) |
| Protection warrior 40 Horde | 9443.9 / 9463.4 | 9857.6 / 9851.0 (+4.2%) |

**The two DPS rows from D5** that fell (Shadow 40 Alliance −1.0%, Fury 30 Alliance −0.6%) were really lower at 10,000 iterations on two seeds, and neither came from a legitimate restriction (no item of origin/main's sets is illegal under K8, K9 or K13). Both are fixed (K15, K16) and now hold or rise: Shadow 151.9 against origin/main's 151.7 and 151.8, Fury 62.9 against 62.3.

No sim DPS moves from this lane's code. The DPS moves in this table come from picks, plus the ranks lane's rank choices and the items lane's engine branch.

Run time of the final 28-spec run: 2,692 s at 4 specs at a time (D5, before K15 and K16, took 1,702 s under a lighter load). Round 8 measured 2,008 s on its branch. K15's extra screens and K16's wand runs add a few sims per band; the memo and concurrency cover them.

## For other lanes

- **Integrator, data run order.** Run the BiS refresh only on a data run that rebuilds `simdb.bin` with held typing (1bd3920b5). On the committed simdb, K11 makes Feral and bear pick one-handers with a held off hand, and the sim then wears the held item as the main hand (K10): +38% to +90% of artifact.
- **Integrator, cross-lane files.** S 13967ecf3 and S 6c239bffd add two optional fields to `data/pipeline/models.py` GearItem (`reputation`, `limit_category`) and write them in `data/pipeline/normalize/gear.py`. Every class-file row gains two keys at the next data run (null on most). The golden `data/tests/golden/items_warrior.json` was regenerated for them. The app contract (`docs/contracts`, `api/internal/contract`) reads class files without `additionalProperties: false`, so nothing else moves. The page could show the standing on a tooltip ("Requires Darkspear Raiders - Exalted").
- **Integrator, other specs.** All of this lane's commits move BiS picks (and published weights) for every group once the refresh runs. None moves sim DPS.
- **Pipeline (whoever owns `data/pipeline/loot`).** The Alliance Darkspear Islands rewards (Theramore Expeditionary Force, 2719xx: Theramore Gauntlets, Militant's Band, Seal of the Expedition, Theramore Signet and so on) have no source in loot.json. Wowhead lists only Creeg Bothunk, so Alliance lists cannot name them. The client has the items (MinFactionID 2799).
- **Items or engine lane.** In the engine, a Feral cat wearing a non-weapon item as its main hand (the committed simdb's held-item typing) swings with 0% miss and 0% dodge. It is harmless once held items are typed correctly, but the attack table should not depend on the weapon slot that way.
- **Items lane / pipeline.** `items/druid.json` (and the other class files) carry Glimmering Staff 249392, an Enchanting-made staff at level 25, with only +11 Intellect and +11 Stamina. Wowhead Forever (2026-11-02) gives it "Equip: Increases damage and healing done by magical spells and effects by up to 34". simdb has the spell power: the sim measures it 8.5 DPS above Bloodscalp Channeling Staff for Balance band 30. The ranker's score does not, so it ranks the staff below a 19-nature-power staff. On origin/main a verify run happens to catch it. With the gates open, the geared weights move, the run misses it, and the band keeps the published set (65.6 against 69.0). The equip-spell fold should be checked for Forever's 249xxx crafted items.
- **Items lane / integrator, for the record.** loot.json's WSG reputation ids are swapped against the client: in Faction 70707, 889 is Warsong Outriders (Horde Forces) and 890 is Silverwing Sentinels (Alliance Forces). The ranker's `repSide` follows loot.json's ids (889 alliance, 890 horde), so loot.json's sources come out right, as the items lane found. The item reputation gate (K8) reads client ids, but no client item names 889 or 890 in MinFactionID. If F3's client_rep ever sources an item under client 889 or 890, `repSide` would side it wrong. The fix would then be to key the ranker's table by name, or to correct loot.json's ids.
- **Integrator, after the data run.** Empty `AWAITING_THE_NEXT_DATA_RUN` in `data/tests/test_content_phases.py` once loot.json carries F3's sources (K17).
- **Page (`web/src/lib/items/source-labels.ts`).** Unchanged from round 8: `pvpSourceLabel` still prints "PvP rank 0" for an unranked reward.

## Still open

- **The two-seed tank spread.** Two weight seeds still differ in 0 to 3 armor slots (0 to 2 at 400 iterations) where tank items sit within error of each other. Picks are reproducible because the seed is fixed. A seed-independent pick would need a tie rule on the tank score's error; I left it alone.
- **Glimmering Staff's missing spell power** (items or pipeline lane): it costs Balance band 30 Horde 4.9% against origin/main's set once the gates open. The guard keeps the published set, so nothing published gets worse, but the better set is not found.
- **The committed simdb's held-item typing (K10).** Every pick here was measured on the held-typed simdb. The refresh needs a data run with the held-typed simdb.
- **Paladin axes** reach `items/paladin.json` only at the next data run (round 8's K1).
- **The Alliance Darkspear rewards** have no source in loot.json (pipeline).
- From round 8, unchanged: the dagger rule is broader than it needs to be (Combat); the off-hand weapon-DPS term counts full DPS; the ranker never recommends two copies of a non-unique ring; and the page's "PvP rank 0" label.
- **Run time.** The memo saves 17% of the sims. The rest is concurrency, which needs free cores. A next step would be to share the ladder and geared weight sets' common trinket tournaments across factions, which the reconcile already re-measures.
