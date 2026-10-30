# BiS content-phase gate

2026-10-30, lane `bis-phase-gate`. Follows `design/reviews/2026-10-30-bis-phase-audit.md`, which found 181 "Raid-ready, Phase 1" picks from content that is not open at launch.

## What changed

**One curated file.** `data/curated/content-phases.json` now says which content is open when. Each gate has:

- a phase key from `data/curated/phases.json` (`launch`, `raids-1`), or `later` with the date `unannounced`
- the original patch, a confidence (`inferred` or `uncertain`) and a reason
- the ids it covers, keyed by stable game ids:
  - `npcs`: a vendor's or a dungeon boss's npc id. World sources carry no npc id in loot.json, so they are matched by the npc's name.
  - `quests`: quest ids
  - `factions`: reputation faction ids, with an optional `from_standing`
  - `zones`: an instance's or zone source's zone id
  - `kinds`: a loot source kind
  - `items`: item ids, used only where nothing else works

A gate whose phase is at or before `launch` is open. When Blizzard confirms a date, the owner edits that gate's `phase` and `date` and nothing else.

**Both languages read it.**

- `data/pipeline/loot/content_phases.py` writes the gates into `loot.json`. This runs in both `loot` and `loot-merge`, after the raid overlay and before the quest turn-in gate.
- `sim/cmd/leveling-bis/contentphases.go` applies the same gates to whatever `loot.json` the ranker loads, through `loadGatedLootIndex`. A gate therefore holds even on a `loot.json` generated before the gate existed.

**The hand lists are gone.**

- Moved from Go into the curated file: `raidLockedQuestOpens` (the AQ War Effort quests 8756, 8789 and 8790), `repFactionRaidPhaseOpens` (factions 609 and 910) and `worldBossSources`.
- Moved from Python into the curated file: `WORLD_BOSS_NPC_NAMES` and `REP_FACTION_RAID_PHASE_OPENS`.
- The rule "a quest reward inherits the reputation of its item" (the Earthstrike case) moved from `apply_quest_opens_gate` into the content-phase step.

**A missing id keeps its gate.** If a gated id is not in tonight's sources, the pipeline logs a warning and the ranker writes one to stderr. The gate stays in the file. This is the fix for commit `3592e40a5`, which deleted the Anachronos gate when the vendor dropped out of one fetch.

**Duplicate quest entries are merged before gating.** Each item's quest entries are collapsed by `quest_id`. The merged entry keeps the reputation fields and the strictest `opens`. In Go, every copy of a duplicated quest takes the strictest phase of any copy. A bare copy can no longer defeat the gate its twin carries. This covered:

- Treasure of the Timeless One
- The Only Prescription
- every Zandalar quest
- the Alterac Valley insignia quests

On build 70658, `loot-merge` went from 4,481 quest entries to 3,133.

**Bosses inside open dungeons get their own source.** A gated boss is split into a sibling source `<instance id>:<gate key>`. The sibling keeps the instance's kind, name and zone, carries `opens`, and leaves the rest of the instance open. For example, Balzaphon now sits in `dungeon:stratholme:scourge-invasion`. Splits on 70658:

- The Scourge Invasion bosses in Stratholme, Scholomance, Dire Maul, Scarlet Monastery, Shadowfang Keep and Razorfen Downs
- The Tier 0.5 summons in Blackrock Spire, Scholomance, Stratholme and Dire Maul

**`vendor_cost` skips gated vendors.** A whole-vendor gate such as Anachronos is not split a second time into `vendor:<id>:later`.

## The gates

| Gate | Phase | Original patch | Confidence | Keys |
|---|---|---|---|---|
| `world-bosses` | raids-1 (2026-12-09) | 1.2 to 1.8 | inferred | 6 npcs (moved from the hand lists, phase unchanged) |
| `aq-war-effort` | later | 1.9 | inferred | Anachronos; Cenarion Circle, Brood of Nozdormu; 10 quests, including the three that were hand-listed in Go |
| `silithus-abyssal-council` | later | 1.8 | inferred | 15 npcs (Abyssal Lords, Dukes, Templars, Everun, Mayna, Twilight Prophet); Silithus zone drops; 27 Cenarion Hold quests |
| `emerald-nightmare-lead-in` | later | 1.8 | uncertain | Twilight Corrupter |
| `zul-gurub-zandalar` | later | 1.7 | inferred | Zandalar Tribe; 41 quests; the six Bloodsoul and Darksoul smithing items |
| `scourge-invasion` | later | 1.11 | inferred | Balzaphon, Lord Blackwood, Revanchion, Scorn, Sever, Lady Falther'ess, Shadow of Doom; 6 quests |
| `argent-dawn-armaments` | later | 1.11 | inferred | quests 9221 to 9228 (Argent Dawn reputation itself stays open) |
| `tier-0-5` | later | 1.10 | uncertain | Mokvar, Deliana and their two guard copies; Valthalak, Mor Grayhoof, Isalien, Kormok, Jarien, Sothos, the Heirlooms chest, Magma Lord Bokk; 65 quests |
| `honor-system` | **launch (2026-11-04), open** | 1.4 | confirmed | every `pvp` source; the six rank quartermasters |
| `battleground-reputation` | later | 1.5 to 1.7 | uncertain | the six battleground factions; 8 quartermasters; Ivus, Lokholar, Korrak; the Alterac Valley zone; 12 quests |
| `darkspear-islands` | later | Forever | uncertain | Creeg Bothunk |
| `darkmoon-faire` | later | 1.6 | uncertain | Lhara; 13 deck and ticket quests |
| `dire-maul` | later | 1.3 | inferred | the Dire Maul zone; the Gordok Ogre Suit quests |
| `season-of-discovery-vendor-stock` | later | Season of Discovery | uncertain | Meilosh, Hasana, Lightspark, Miranda Breechlock |
| `forever-new-vendors` | later | Forever | uncertain | Pix Xizzix and the four other Undermine traders (265574 to 265577) |
| `ahn-qiraj-crafting` | later | 1.9 | inferred | the six Obsidian crafted items (items) |
| `late-recipes` | later | 1.9 to 1.11 | uncertain | Stormshroud Gloves, Titanic Leggings (items) |

On build 70658, every listed id matches a source, so tonight there are no missing-id warnings. Three additions go beyond the audit's table:

- `ahn-qiraj-crafting`: Light Obsidian Belt became Marksmanship's waist pick, and Obsidian pieces became Fury's, once the earlier picks were gated. Every Obsidian recipe needs Obsidian Shards, which drop only in Ahn'Qiraj and have no source in `loot.json`, so the recipe-and-reagent gate could not see them.
- The Bloodsoul and Darksoul smithing items, added to `zul-gurub-zandalar` for the same reason. They need Souldarite, which drops only in Zul'Gurub. They surfaced for the Protection specs, Retribution and Arms.
- `emerald-nightmare-lead-in`: Twilight Corrupter, kept apart from the Silithus gate.

## Uncertain sources (gated and listed)

These are gated under the house rule because their original phase, or their status in Forever, is unsure:

- **Battleground reputations, quartermasters and the Alterac Valley summons.** Blizzard gives no launch list.
- **Darkspear Islands (Creeg Bothunk).** No timing.
- **Darkmoon Faire.** No schedule.
- **Tier 0.5.** Forever is silent. The 226xxx ids are Season of Discovery re-itemisations.
- **Season of Discovery-era stock** at Meilosh and the three Argent quartermasters. They list only 2278xx items, with no cost and no standing.
- **The Forever-new Undermine traders.** Pix Xizzix, Winklespark, Gezzy Gunkgear, Rettrick and Zippie Fizzbolt are labelled "New in Forever", with unknown currency and timing. Only Pix Xizzix was in the audit. The other four sell the same kind of 274xxx stock and two of them appeared as band-60 alternatives, so they are gated with it.
- **Stormshroud Gloves and Titanic Leggings.** Flagged by id range only.
- **Twilight Corrupter.**

## Before and after

How this was run:

- `loot-merge` regenerated loot.json from the committed inputs. This is the nightly's filter stage, `uv run python -m pipeline loot-merge --build 1.60.1.70658 --engine <fork at 1f542def3>`.
- The ranker then ran at band 60 with `-weights-iterations 200`, the setting `bis.yml` uses.
- "Before" is the committed nightly output for 1.60.1.70658 at engine 1f542def3.

Neither the regenerated data nor the ranked files are committed.

### The four reviewed items and the ring

| Item | Before | After |
|---|---|---|
| Signet Ring of the Bronze Dragonflight (234032 to 234035, 234202) | finger1 pick in all 28 specs | in no band-60 pick or alternative |
| Chains of the Lich (23125) | Fire Mage neck (and 6 other casters) | gone; Fire Mage neck is Archlight Talisman (quest "Hidden Treasures") |
| Elemental Focus Band (20682) | Fire Mage finger2 | gone; Band of the Unicorn (world drop) |
| Talisman of Ascendance (22678) | Fire Mage trinket1 | gone; Rune of the Guard Captain (Horde) / Frozen Heart of the Mountain (Alliance) |
| Dryad's Wrist Bindings (19595) | Fire Mage wrist | gone; Netherflame Cuffs (Tailoring) |

### Replacement picks, "Raid-ready, Phase 1", Horde

**Fire Mage** (set DPS 846.5 to 724.1):

| Slot | Before | After |
|---|---|---|
| head | Sorcerer's Crown (Saving the Best for Last) | Red Mageweave Headband (Tailoring) |
| neck | Chains of the Lich (Balzaphon) | Archlight Talisman (Hidden Treasures) |
| shoulder | Rugged Mantle of the Timbermaw (Meilosh) | Mantle of the Timbermaw (Tailoring) |
| back | Arcanoweave Cloak (Pix Xizzix) | Mageflame Cloak (world drop) |
| wrist | Dryad's Wrist Bindings (Warsong Outriders) | Netherflame Cuffs (Tailoring) |
| hands | Inferno Gloves (Tailoring, Dire Maul pattern) | Gloves of Spell Mastery (Tailoring) |
| waist | Knowledge of the Timbermaw (Meilosh) | Fiery Waistcord (Tailoring) |
| legs | Sentinel's Silk Leggings (Illiyana Moonblaze) | Cindercloth Pants (Tailoring) |
| feet | Sorcerer's Boots (Anthion's Parting Words) | Omnicast Boots (BRD, Golem Lord Argelmach) |
| finger1 | Signet Ring (Anachronos) | Maiden's Circle (world drop) |
| finger2 | Elemental Focus Band (Prince Skaldrenox) | Band of the Unicorn (world drop) |
| trinket1 | Talisman of Ascendance (Epic Armaments) | Rune of the Guard Captain (quest) |

**Fury Warrior** (set DPS 894.1 to 796.8):

| Slot | Before | After |
|---|---|---|
| neck | Beads of Ogre Might (Falrin's Vendetta) | Skibi's Pendant (world drop) |
| shoulder | Darkspear Pauldrons (Creeg Bothunk) | Lightforge Spaulders (UBRS, The Beast) |
| back | Howler's Furs (Pix Xizzix) | Cape of the Black Baron (Baron Rivendare) |
| wrist | Forest Stalker's Bracers (Warsong Outriders) | Deepfury Bracers (world drop) |
| hands | Raider Gauntlets (Creeg Bothunk) | Devilsaur Gauntlets (Leatherworking) |
| waist | Belt of Preserved Heads (A Collection of Heads) | Might of the Timbermaw (Leatherworking) |
| legs | Titanic Leggings (Blacksmithing) | Devilsaur Leggings (Leatherworking) |
| feet | Boots of Heroism (Anthion's Parting Words) | Shadowcraft Boots (Rattlegore) |
| finger1 | Signet Ring (Anachronos) | Masons Fraternity Ring (Divino-matic Rod) |
| finger2 | Don Julio's Band (Frostwolf Clan) | Falcon's Hook (Uldaman) |
| trinket2 | Darkmoon Card: Maelstrom | Frozen Heart of the Mountain (Enchanting) |
| main_hand | Ravencrest's Legacy (Treasure of the Timeless One) | Ironwood Blade (Blacksmithing) |
| off_hand | Shadowsong's Sorrow (Treasure of the Timeless One) | Greenhammer (Blacksmithing) |
| ranged | The Purifier (Epic Armaments) | Precisely Calibrated Boomstick (world drop) |

**Holy Priest** (set DPS 737.4 to 657.3):

| Slot | Before | After |
|---|---|---|
| head | Virtuous Crown (Saving the Best for Last) | Devout Crown (Darkmaster Gandling) |
| neck | Wavefront Necklace (Lord Skwol) | Glowing Eye of Mordresh (Razorfen Downs) |
| shoulder | Argent Elite Shoulders (Hasana) | Devout Mantle (Solakar Flamewreath) |
| back | Hide of the Wild (Leatherworking) | Wingveil Cloak (Sunken Temple) |
| wrist | Virtuous Bracers (An Earnest Proposition) | Devout Bracers (Crimson Priest) |
| hands | Virtuous Mitts (Mokvar) | Devout Gloves (Archivist Galford) |
| waist | Virtuous Belt (Just Compensation) | Wisdom of the Timbermaw (Tailoring) |
| feet | Incandescent Mooncloth Boots (Meilosh) | Mooncloth Boots (Tailoring) |
| finger1 | Signet Ring (Anachronos) | Band of the Hierophant (world drop) |
| finger2 | Blessed Band of Light (Pix Xizzix) | Eye of Adaegus (world drop) |
| trinket1 | Darkspear Voodoo Seal (Creeg Bothunk) | Shard of the Splithooves (Heroes of Old) |
| trinket2 | Serenity Field (Pix Xizzix) | Uther's Strength (world drop) |

**Marksmanship Hunter** (set DPS 777.8 to 724.1):

| Slot | Before | After |
|---|---|---|
| head | Beastmaster's Cap (Saving the Best for Last) | Black Dragonscale Helm (Leatherworking) |
| neck | Amulet of the Darkmoon (1200 Tickets) | Skibi's Pendant (world drop) |
| shoulder | Darkspear Pauldrons (Creeg Bothunk) | Shadowcraft Spaulders (Cannon Master Willey) |
| wrist | Forest Stalker's Bracers (Warsong Outriders) | Beaststalker's Bindings (Fleshflayer Ghoul) |
| hands | Raider Gloves (Creeg Bothunk) | Mourning Handguards (Leatherworking) |
| waist | Belt of Preserved Heads (A Collection of Heads) | Sagebrush Girdle (Vyletongue Corruption). The first run picked Light Obsidian Belt, which led to the `ahn-qiraj-crafting` gate. |
| legs | Sentinel's Chain Leggings (Illiyana Moonblaze) | Windrunner Legguards (world drop) |
| feet | Beastmaster's Treads (Mokvar) | Shadowcraft Boots (Rattlegore) |
| finger1 | Signet Ring (Anachronos) | Masons Fraternity Ring (Divino-matic Rod) |
| finger2 | Cutthroat's Signet (Pix Xizzix) | Ring of the Underwood (world drop) |

The picks now look like the Classic Phase 1 pre-raid lists: the Devout, Beaststalker and Magister's dungeon sets, Devilsaur, Cape of the Black Baron and the Timbermaw crafts. Set DPS falls by 6 to 15 percent. That is the cost of dropping Tier 0.5, the battleground epics, the Darkmoon cards and the item-level 80 ring.

### Every spec at band 60

After the four specs above, all 28 specs were ranked at band 60, each against the final gates. The ranked files were then dropped into the build directory, temporarily, and the two published-file checks were run against them (`FOREVER_BIS_SANITY=1`):

- `TestNoBannedItemIsAPublishedBand60Pick` passes.
- `TestTheBand60PicksComeOnlyFromContentOpenAtLaunch` passes.
- Against the committed files, the same checks fail with 1,035 banned-item hits.

Across the 1,804 band-60 picks (28 specs, both factions, both presets):

- 1,450 picks changed.
- The four reviewed items and the Signet Rings went from 245 picks to 0. They are also in no alternative.
- Source kinds, before: quest 396, vendor 678, crafted 361, rep 168, dungeon 69, world 66, world drop 66.
- Source kinds, after: crafted 707, dungeon 463, world drop 361, quest 251, world 20, with no vendor or reputation picks.

Raid-ready preset, Horde, set score before and after. The tank specs show their tank score, not damage.

| Spec | Before | After | Slots changed |
|---|---|---|---|
| druid-balance | 493.9 | 416.1 | 12 |
| druid-feral-bear | 336.8 | 324.8 | 14 |
| druid-feral | 664.7 | 604.0 | 12 |
| druid-restoration | 678.4 | 611.7 | 13 |
| hunter-beast-mastery | 781.9 | 734.0 | 10 |
| hunter-marksmanship | 777.8 | 724.1 | 10 |
| hunter-survival | 776.4 | 702.3 | 16 |
| mage-arcane | 651.9 | 565.9 | 13 |
| mage-fire | 846.5 | 724.1 | 12 |
| mage-frost | 522.6 | 445.4 | 12 |
| paladin-holy | 569.4 | 498.7 | 12 |
| paladin-protection | 480.1 | 447.3 | 12 |
| paladin-retribution | 626.0 | 594.6 | 15 |
| priest-discipline | 637.6 | 559.4 | 13 |
| priest-holy | 737.4 | 657.3 | 12 |
| priest-shadow | 564.5 | 518.2 | 12 |
| rogue-assassination | 737.5 | 655.6 | 13 |
| rogue-combat | 704.7 | 618.4 | 12 |
| rogue-subtlety | 612.9 | 532.9 | 14 |
| shaman-elemental | 453.6 | 390.3 | 13 |
| shaman-enhancement | 677.9 | 619.5 | 14 |
| shaman-restoration | 630.3 | 581.7 | 12 |
| warlock-affliction | 833.2 | 737.3 | 11 |
| warlock-demonology | 779.1 | 688.1 | 11 |
| warlock-destruction | 881.0 | 774.2 | 11 |
| warrior-arms | 794.6 | 713.7 | 14 |
| warrior-fury | 894.1 | 796.8 | 14 |
| warrior-protection | 316.3 | 296.9 | 15 |

Crafted items with 19000 to 23999 ids that are still picked, not audited here. Their recipes are trainer-taught or have no source in `loot.json`, so no gate can see them yet:

- Timbermaw crafts: Might, Wisdom and Mantle of the Timbermaw, Timbermaw Brawlers
- Argent Dawn crafts: Dawn Treaders, Golden Mantle of the Dawn, Argent Shoulders
- Blue Dragonscale Leggings (one pick)

The Timbermaw and Argent Dawn recipes were Phase 1 in Classic 2019. Blue Dragonscale Leggings may be a 1.8 Cenarion recipe.

Ranking was not a bottleneck. Each spec took about 55 seconds at band 60.

## Honor opens at launch, and the mislabelled PvP sources it exposed

Opening `honor-system` broke both launch tests on Dryad's Wrist Bindings (19595, gate `battleground-reputation`). The item had a source "PvP rank 12 · Legionnaire · Horde", so it looked obtainable through honor.

**Where the row came from.** The client's `ItemSparse.RequiredPVPRank` for 19595 is 0 (checked against a fresh wago.tools export of 1.60.1.70658), and classic-db's `required_honor_rank` is 0 as well. Wowhead was not involved. The row was made in our own pipeline, in two steps:

1. **The title fallback.** `HONOR_RANK_BY_TITLE` ranked classic items by name alone. "Legionnaire's Band", a Warsong Outriders reward, got rank 12. So did ordinary greens such as Legionnaire's Leggings and the Champion's armour.
2. **The vendor backfill.** `pvp_faction._backfill_vendor_items_into_pvp_buckets` then treated Kelm Hargunth, the Horde Warsong Gulch quartermaster, as a rank-12 quartermaster because he sells one "ranked" item. It filed his whole stock under `pvp:rank-12:horde`: Dryad's Wrist Bindings, Forest Stalker's Bracers and the Season of Discovery Warsong items. The same happened to ordinary armour vendors that sell Legionnaire's Leggings, which put Scalemail and Polished Scale gear under the same rank.

`loot-merge` then recovered these ranks from the committed `loot.json` every night, so the mistake was carried forward from build 70009 onwards.

**The fix, in the pipeline:**

- **classic-db is authoritative for every classic id it names, zero included.** `classicdb_items.honor_ranks` drops classic-unranked ids from the client ranks, from the ranks recovered in merge mode, and from the title fallback. `pvp_faction.drop_non_rank_rewards` removes them from every pvp source after the split and backfill. On the rows wago served (19,257 of the client's items), the client and classic-db agree on all 441 ids both rank, and the client ranks no classic id that classic-db leaves unranked.
- **Merge mode keeps a recovered rank only with evidence on the item.** `pvp_faction.ranks_with_item_evidence` requires a rank title or a faction restriction. Every rank the client states meets this test.
- **The backfill adds only items that name their own rank title.** A reputation quartermaster's other stock is no longer filed under a rank.
- **The ranker applies the same classic-db rule** (`sim/cmd/leveling-bis/classicranks.go`, reading the build's committed `raw/classicdb/item_template.json`). The gate therefore also holds on the committed `loot.json` the nightly has not yet regenerated.

**Faction mismatches, checked across every pvp source.** After `loot-merge` with the fix:

- 664 pvp listings remain, down from 774. All 110 removed listings are non-rank items.
- No item is on both sides.
- No item's faction restriction or rank title disagrees with its bucket.
- Every remaining item carries a rank title or a faction restriction.

One general asymmetry turned up and is fixed. The title table named Horde's rank-11 to rank-13 titles (Blood Guard's, Legionnaire's) but not Alliance's (Knight-Lieutenant's, Knight-Captain's, Knight-Champion's). Captain Dirgehammer's re-itemised Knight-Lieutenant's sets therefore reached Alliance lists with no rank and no cap, while their Horde twins were capped. Every title from rank 10 up is now listed for both sides, except Alliance's rank-10 "Knight's", which also begins ordinary names such as Knight's Seal of the Dawn and Knight's Lance. A test holds the table to the two rank ladders.

"Offered to Alliance too" was a flaw in my launch test, not in the data. The test passed the faction itself as the item's restriction, which waives the side check. It now passes each item's real restriction from `items.json`, as the ranker does.

Dryad's Wrist Bindings stays on the banned list: it is gated by its battleground reputation, and it has no honor source. The honor-system examples on the list (Marshal's Silk Bracers, Knight-Lieutenant's Plate Helm) are no longer banned now that the gate is open, and no edit to the list was needed.

**What honor adds to band-60 lists.** Six specs were re-ranked at band 60 with honor open and the pvp fix applied: Fury Warrior, Combat Rogue, Retribution Paladin, Survival Hunter, Enhancement Shaman and Fire Mage.

- No default pick changed. The owner's existing `pvpRankCap` (rank 10, 2026-09-29: "a fresh level 60 can realistically get") keeps rank 11 and above out of the picks.
- Honor gear now appears as alternatives in every spec: 4 to 19 per list. These are not the 1.12 rank 10 to 13 epic sets; they are Forever's re-itemised honor gear:
  - Knight-Lieutenant's and Blood Guard's (rank 11)
  - Knight-Captain's and Legionnaire's (rank 12)
  - Lieutenant Commander's and Champion's (rank 14)
  - Marshal's, General's, Field Marshal's and Warlord's (ranks 16 and 17)
  - the Grand Marshal's and High Warlord's weapons (rank 18)
- For example, Fury Warrior's head alternatives are Champion's Plate Helm (rank 14, 15.0 score behind) and Blood Guard's Plate Helm (rank 11, 18.7 behind).
- One pick moved off honor gear. In the honor-open run before the title fix, Alliance Retribution's feet pick was Knight-Lieutenant's Lamellar Sabatons (rank 11, uncapped because the title was missing). With the fix it is capped like its Horde twin and listed as an alternative, so Alliance and Horde now match.
- Raising `pvpRankCap` is how honor gear would become default picks. That is an owner decision, separate from this gate.

## Every surface reads the same generated data

- **Site BiS pages and home upgrades** read `data/builds/<build>/bis/<spec>.json` (`web/src/lib/bis/load.ts`, `web/src/lib/home/upgrades.ts`).
- **Discord /bis panel** reads `bis/*.json` through `api/internal/bis/source_build.go`, the same files the API image carries.
- **Addon `Data.lua`** is rendered from `bis/*.json` by `pipeline addon-data`, which the `bis.yml` nightly runs right after ranking.
- **Site sim tools.** The Droptimizer and source picker read `loot.json`'s `opens`, which the pipeline now writes.

`bis.yml`'s input hash covers `data/curated` and `sim`, so the next nightly re-ranks every spec. Its `loot-merge` step writes the gates into `loot.json`, and the addon data is regenerated from the new picks. Nothing regenerated is committed here.

## Tests

- **`data/tests/test_content_phases.py`** (20 tests):
  - the curated file validates: known phase, a date that matches it, ids unique across gates
  - each key type reaches `loot.json`
  - a boss splits out of an open dungeon
  - `from_standing` works
  - an item gate splits the item out of its source
  - a gate dated at launch is open, and a stricter existing `opens` wins
  - duplicate quests merge and the gated copy wins
  - the Earthstrike inheritance holds
  - a source missing from the fetch keeps its gate and the curated file is untouched
  - a launch test: against the active build's committed `loot.json`, no banned item has a rankable open way
- **`sim/cmd/leveling-bis/contentphases_test.go`**:
  - the same rules on the Go side, including the missing-source test
  - `TestNoBannedItemIsObtainableAtLaunch`, which runs on every CI run. Over the committed `loot.json` with the gates applied, no banned item has a source a level-60 character of either faction can use. The ranker picks only from such sources, so no banned item can be a pick or alternative in any spec.
- **`sim/cmd/leveling-bis/testdata/launch-banned.json`** is the shared banned list. It holds the four reviewed items, every Signet Ring id, and at least one item for every gate. Both languages read it. It bans an item only while its gate is closed, so opening a gate needs no test edit.
- **`phase60_test.go`**:
  - `TestTheBand60PicksComeOnlyFromContentOpenAtLaunch` now requires an open source after the curated gates, for every source kind. Before, it only checked `opens == ""` on dungeon and world sources.
  - The new `TestNoBannedItemIsAPublishedBand60Pick` scans every published pick, alternative and runner-up.
  - Both read the committed `bis/*.json`, which this lane does not regenerate. So they run in `bis.yml`'s sanity step (`FOREVER_BIS_SANITY=1`), after the nightly has re-ranked. Against today's committed files they fail with 1,035 banned-item hits, which is the before state.
- **Replaced tests.** The old fallback tests that pinned the hand lists now run through `loadGatedLootIndex` with the real curated file.

## For the owner to confirm

1. **The house rule.** Forever launches "before Molten Core", which means the original Classic Phase 1 content set. Everything that opened after Phase 1 is gated until a date is announced.
2. **Honor rank rewards: decided.** Coordinator ruling: an announced date opens a gate, and Blizzard's PvP article dates honor ranks and rank rewards to launch. `honor-system` is `launch` / `2026-11-04`, `confirmed`. See "Honor opens at launch" below.
3. **Battlegrounds:** Warsong Gulch, Arathi Basin, Alterac Valley and Darkspear Islands at launch.
4. **Tier 0.5, the Darkmoon Faire and Dire Maul.**
5. **The Season of Discovery-era vendor stock and the five Forever-new Undermine traders.**
6. **World bosses stay at `raids-1`.** That is the pipeline's earlier choice. Blizzard has dated no world boss, so under the house rule they could equally be `later`. The launch lists are unaffected either way.
7. **The "Raid-ready, Phase 1" preset label** (audit item 7) is unchanged. With these gates it is now true of the gear.
8. **Blue Dragonscale Leggings** (one pick) is unverified. See the full-run list above.
9. **World-boss items in zone sources.** Lord Kazzak's and Azuregos' items are also listed under their zones' `zone:` sources, which the ranker never picks. The ranker is unaffected, but the site's sim tools still treat those items as open.
