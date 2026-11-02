# Sim audit round 10: ranker (2026-11-02)

This lane owns `sim/cmd/leveling-bis`, the BiS pipeline code and `data/curated/content-phases.json`. It settled owner ruling 3 (The Watchers), took two handoffs from the integrator (the Warsong Gulch ids and Holy paladin's libram), and re-audited the ranker after round 9 on a full local re-rank. It made no engine change.

**Setup**
- Site: `wt/r10-ranker` from origin/main bb619a188, with wt/r10-items merged (711a12d20) for its Warsong Gulch fix.
- Engine: the pin 4d3129cae (`$R/base-engine`).
- A local copy of the held data run in `scratchpad/r10/ranker/root`, never committed. I ran `fetch`, `normalize --allow-shrink`, `loot --allow-shrink` and `simdb` on build 1.60.1.70707 with this branch's pipeline. The talents were then put back to the committed ones, because I skipped `fetch-wowhead-talents`. That gives what the next data run will give the ranker: class files with `reputation` and `limit_category`, round 9's client reputation sources, the held-item simdb typing and paladin axes. The ranker binary embeds that simdb.
- Re-rank: all 28 specs at bands 20 to 60, bare and raid at 60, both factions (336 rows), plus two side runs (below). No output is committed.

**In short: 4 CONFIRMED-BUGs, all fixed with tests first, plus ruling 3 settled**
- K18: a client reputation source opened four gated vendors' stock. Live once the data run lands; 16 picks in 4 specs on the local run.
- K19: the ranker read the Warsong Gulch sides with the fork's reversed ids (handoff).
- K20: a healer's trinket and relic re-measure at the ceiling was refused for asking too many iterations (handoff).
- K21: the BiS report still called crit one unified stat.

## Owner ruling 3: The Watchers

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| W1 | **The Watchers (faction 2819) is not launch content on the evidence. The gate stays `later`, now `inferred`, with the evidence in its "why".** | **What ties it to the Barrow Deeps:**<br>- Faction 2819 (70707): "An order of powerful jailers and marshals of night elf society. They stand vigilant to ensure that no threat can ever escape their extensive underground network of prisons."<br>- Blizzard, "World of Warcraft: Forever Raid Overview" (news.blizzard.com/en-us/article/24303009, fetched 2026-11-02): the Barrow Deeps is "the prison" that "once held Illidan Stormrage". It is a level-60, 10-player raid with entrances from Mount Hyjal, Winterspring and Blackmaw Hold, and the post asks "how the wardens plan to contain the danger below". It unlocks on 9 December with Hyjal Summit and Onyxia.<br>- 70707 client: Brooch of the Wardens 277136, "Grants access to the Vord'anaar Prison within Barrow Deeps" (Wowhead Forever tooltip). Its neighbour 277153, Intricate Filigree, needs Watchers Honored. Map 2817 is Starfall Barrow Den, and Starfall is in Winterspring, one of the raid's three entrances.<br><br>**No launch way in:**<br>- The client has no quest or reputation row that earns 2819: QuestFactionReward 70707 has only the two generic rows, and no Criteria asset names 2819.<br>- What 2819 gates: 16 signets (275968 to 275984, level 60, Bind on Pickup, Friendly to Exalted, "Unique-Equipped: Watcher's Ring"), Nelanna's Enchanted Jewel 277519 and Intricate Filigree 277153.<br>- Wowhead Forever's faction page and the signets' pages name no zone, NPC or quartermaster (fetched 2026-11-02).<br>- The datamined faction list (kami-labs, citing Wowhead's datamine news) does not put The Watchers among Mount Hyjal's three launch factions (Guardians of Hyjal, Barkskin Burrow, Cenarion Scouts).<br>- No Blizzard Forever post names the faction.<br><br>Ruling: open if launch content, else stay `later` with the evidence. So it stays `later` and should be revisited when the Barrow Deeps opens or Blizzard says how the reputation is earned. I kept `later` rather than `raids-1`, as the ruling asks: nothing proves the rewards come only from the raid.<br>Tests: `test_the_forever_zone_reputations_follow_the_launch_evidence` now pins `later`/`unannounced`/`inferred` and the evidence in the "why" (it failed first). The Go test comment is updated. | Settled: not launch | S 3741c143f |
| W2 | **The awaiting pass.** `AWAITING_THE_NEXT_DATA_RUN` keeps faction 2819: the committed loot.json has no Watchers source yet. I ran the banned-item test with the set emptied against the local data run:<br>- on origin/main's pipeline it FAILED: Sea Giant's Toe Ring 274746 had an open way in (K18);<br>- with K18 it PASSES.<br>So the integrator can empty the set after the data run, provided K18 is in it. | Local run, `tests/test_content_phases.py::test_no_banned_item_has_an_open_way_in_the_active_build` with the set emptied. | Awaiting pass kept (held output) | - |

## Confirmed bugs

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K18 | **A client reputation source opened gated vendors' stock.** Round 9's `apply_client_reputations` gave every item with a client MinFactionID its own `rep:` source, even when a vendor already sold it. Gating the vendor leaves that rep source open. So the four gated Undermine traders (`forever-new-vendors`: Winklespark, Gezzy Gunkgear, Rettrick, Zippie Fizzbolt) had their 27 Ratchet, Booty Bay, Gadgetzan and Everlook rewards open at launch. Examples: Sunburn Spaulders, Jhordy's Misplaced Screwdriver, Substandard Belt Chain, Sea Giant's Toe Ring. On the local run they were picked in 16 rows: Survival and Combat at 40 and 50, Arms at 40, Holy priest at 30 and 40. | ItemSparse 70707: 274746 MinFactionID 21 (Booty Bay), MinReputation 4. loot.json from the local run: `rep:booty-bay:friendly` with no `opens` beside `vendor:265575` with `opens: later`.<br>Fix: an item that a source or quest already names gets its side and no rep source. Its seller is the way in, and the ranker already holds every source to the item's own client standing (round 9's K8). Items nothing else names keep their rep source: the Theramore set, The Watchers.<br>Tests first: `test_a_gated_vendors_reputation_stock_stays_gated` and `test_a_vendor_item_keeps_its_vendor_as_its_way_in_and_gets_its_side` failed, then passed. The items lane's own new test still passes. | CONFIRMED-BUG (live at the next data run) | S 0158de259 |
| K19 | **The ranker read the Warsong Gulch sides with the fork's reversed ids** (items lane handoff). In client Faction 70707, 889 is Warsong Outriders (Horde Forces 892) and 890 is Silverwing Sentinels (Alliance Forces 891). The items lane's S 34e994332 makes loot.json keep those ids. `repSide` and `repFactionSwap` still had 890 as Horde. On the regenerated loot.json, every Horde WSG reward would have been sided Alliance and relabelled. The content-phases labels for 889 and 890 were swapped too. | Fix: `repSide` follows the client; the labels and the Go and Python fixtures use the client ids.<br>Tests first: `TestRepSideFollowsTheClientFactionIDs` and the rewritten `TestCorrectedRepSourceSwaps...` failed, then passed. On a loot.json rebuilt with the items lane's pipeline, all 21 named battleground rep sources match the new `repSide`. | CONFIRMED-BUG (handoff; live at the next data run) | S cdb749068 (merge 711a12d20) |
| K20 | **A healer's trinket and relic re-measure at the ceiling was refused** (healers lane handoff, Holy paladin's libram). The healer engine runs seven times the iterations it is asked for. The ceiling re-measure asks for 2000, so a healer sent 14000; the request layer accepts at most 10000 ("a split part's iterations must be between 1 and 10000"). On the local run that happened 38 times in four healer specs (Holy paladin 14, Restoration druid 16, Restoration shaman 6, Holy priest 2). Each such slot lost its second look at 2000 iterations. | Fix: the scaled count stops at `api.MaxIterations`.<br>Test first: `TestHealEngineNeverAsksForMoreThanTheRequestLayerAllows` failed (14000), then passed.<br>After: 0 failures. No pick changes on today's data (all five healer specs re-ranked: 0 of 60 rows differ), because nothing that was cut off gains beyond its error at the ceiling. | CONFIRMED-BUG (no pick moves today) | S 9c1199c19 |
| K21 | **The BiS report said crit is one stat on every table.** Since round 9 only item crit rating and the all-table auras reach both tables; Agility is melee and ranged, Intellect spell. | Test first: `TestWriteMarkdownStatesTheCritSplit`. | CONFIRMED (prose in the generated report, no number) | S 2a8937c85 |

Holy paladin's libram has a second cause, in the engine. Libram of Infusion (spell 1306429, aura 107, +6% crit, client families {2097152, 65536}) never reached the Holy Shock heal 1311605. The engine mapped family 1<<21 to the cast only. With the libram on, the heal showed a zero delta: Holy paladin 60 raid, 10,000 iterations, seeds 1 and 50001, 670.8 and 670.6 HPS with or without it. The healers lane fixed it (E ac11976ed). Until that is pinned, the slot stays empty (`no_dps_value`).

## Fresh audit of the ranker after round 9

Checked on the 336 rows of the local re-rank with `audit10.py`. It runs round 9's standing, side and limit-category checks, plus gated factions, shaman dual wield, off-hand weapons for specs that do not dual wield, Survival's melee main hand, wands on melee specs, and one-table crit items on the wrong spec. Result: **0 violations**.

| # | area | what I checked | verdict |
|---|---|---|---|
| A1 | Crit split and the weight vocabulary | One `crit` weight id. `statid.TableOnly` keeps MeleeCrit and SpellCrit out of every weights request. Every spec's `weight_stats` has `crit`. The sweep adds every-table crit, which is what an item's crit rating (stat 32) gives. So item scoring is right for melee, casters and healers alike. | Aligned |
| A2 | One-table crit on items | Only 4 equip links in the client carry aura 52, 57 or 290 (ItemEffect, ItemXItemEffect and SpellEffect 70707). Two of them are Blackhand's Breadth 13965 (1318954, 1% melee and ranged) and Eye of the Beast 13968 (1318956, 1% spell). Our class files give both 2% every-table crit from classic-db, because neither has an ItemSparse row, and no use effect (see For other lanes). Neither is picked or listed as an alternative on the run (UBRS quest rewards, floor 55). | Aligned for the ranker; item data open |
| A3 | Stones and wizard oil | The ranker never ranks enchants or imbues. Each pass sends its preset's consumables (bare = `leveling.KitConsumes`, the same source as the /sim bare sets in `published-presets.json`; raid at 60). | Aligned |
| A4 | Wands | The engine's Shoot is DefenseTypeMagic (`sim/core/wand_test.go`), so it reads Crit + SpellCrit and Intellect still raises wand crit. Every mage, warlock and priest row wears a wand. Healers score wands at 0 by design. | Aligned |
| A5 | Per-band rank choices and preset | Each pass names its preset (K14), and rank-choice rows apply by preset (`forPreset`). The bare pass sends the same kit /sim's bare default sends at every band. | Aligned |
| A6 | Class rules | Shamans: `offHandPool` gives specs that do not dual wield no weapons; Enhancement wears a two-hander in every row. Survival dual wields melee weapons in every row. Paladin axes are in the rebuilt class file (255 axes). Druids have no polearms. | Aligned |
| A7 | Phases and gates | Battleground and Darkspear open: Creeg Bothunk's stock comes through his vendor row, held to its client standing (K8); the Theramore set through its rep source. The Watchers, the Undermine traders (after K18), raids, Dire Maul and the rest stay closed. 0 gated picks. | Aligned after K18 |
| A8 | Unique-Equipped, standing, side | Round 9's checks: 0 limit-category pairs, 0 revered or exalted rewards below 60, 0 wrong-side rewards. | Aligned |
| A9 | Set completion and the incumbent guard | The set-completion screens run and adopt (example: Feral 30, Defias Leather 4-piece). For each spec's largest drop against the published file, I re-simmed published and new sets on the same engine at 2,000 to 3,000 iterations, seeds 1 and 50001: the new set is level or ahead in every DPS row (Feral 30 H 84.0 → 84.7, BM 20 A 63.4 → 65.0, MM 20 A 62.0 → 63.8, Survival 20 A 46.4 → 47.3, Enhancement 30 H 63.1 → 63.5, Arms 30 H 56.2 → 56.7, Retribution 30 A 68.8 → 68.9). The large drops against the published file are the engine and data changes since 1 November, not picks. | Aligned |
| A10 | Healer rows | Restoration druid 20 A: the new gear beats the published gear under the same talents (41.4 against 39.9 HPS). The ladder's band-20 talents are another matter (Still open). | Aligned (gear) |

## Before/after

Only BiS picks move. No sim DPS moves from this lane. Set score is at the verify harness (300 iterations, seed 7); healers are HPS.

| change | specs re-ranked | rows with a changed pick | set score | example, 10,000 iterations, seeds 1 / 50001 |
|---|---|---|---|---|
| K18 (unfixed against fixed loot) | Survival, Combat, Arms, Holy priest, Retribution, Enhancement | 16 of 72 (0 in Retribution and Enhancement) | −0.02% to −0.16% mean; worst −1.2% (Survival 40 A, 158.7 → 156.8) | Combat 40 H 115.2 / 115.3 → 114.1 / 114.1 (−1.0%); Holy priest 40 H 116.9 / 116.8 → 117.1 / 116.9 (level) |
| K20 | the five healers | 0 of 60 | 0 | - |
| K19 | - | 0 on the committed loot.json; moves Horde and Alliance WSG rewards once loot.json is rebuilt | - | - |

The drops are honest: those items are not on sale at launch.

## For other lanes

- **Integrator, data run.** Run the held data run with K18 in it, then empty `AWAITING_THE_NEXT_DATA_RUN` (W2). `TestRepSideMatchesTheBuildsFactionIDs` skips with "awaiting the next data run" until loot.json carries the client WSG ids. It is the one held-output test I left waiting.
- **Integrator, data run.** `normalize` refuses to write the build without `--allow-shrink`: items.json goes from 25,853 to 25,688 ids. Of the 165 that go, 105 are gear rows at item level 30 and required level 1 that copy Molten Core items (288908 Aurastone Hammer, 288919 Band of Accuria, the 2888xx Cenarion set); the rest are recipes and consumables. None of the 165 is in wago's 70707 ItemSparse export or the committed hotfix layer today. Single source; the data lane should confirm they are gone before re-baselining.
- **Items lane.** Blackhand's Breadth 13965 and Eye of the Beast 13968, as Forever links them:
  - Breadth: 1318954, 1% melee and ranged crit, plus a use effect (1318944, aura 306, 5-minute cooldown) and an aura-4 spell 1318945.
  - Eye: 1318956, 1% spell crit, plus use 1287840 (aura 55) and aura-4 spell 1318846.

  Our class files have 2% every-table crit and no use for both. They are not picked today.
- **Healers lane.** Libram of Infusion now reaches the heal with E ac11976ed. With K20, the next re-rank measures it.
- **Owner of `sim/leveling`'s talent ladder.** `LadderTalentString` fills the guide's targets top row first, so the bands below 60 take whichever early talent the guide happens to list first. Example: Restoration druid band 20 spends 5 points in Furor (5510). On the same new gear, 5051 heals 43.9 against 41.4 HPS at 3,000 iterations, seeds 1 and 50001 (+5.9%). The ladder says it "approximates a leveling build, it does not design one". The band talents are a talent-build question, not the ranker's.

## Still open

- The Watchers' gate: revisit when the Barrow Deeps opens on 9 December or Blizzard says how the reputation is earned.
- Holy paladin's libram slot stays empty until the healers lane's engine fix is pinned.
- From round 9, unchanged: the two-seed tank spread; Glimmering Staff's spell power; the dagger rule for Combat; the off-hand weapon-DPS term; two copies of a non-unique ring; "PvP rank 0" on the page.
