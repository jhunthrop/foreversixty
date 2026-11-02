# Sim-accuracy round 7: pipeline lane

Branch: site `wt/r7-pipeline` (S) from origin/main 9cfb91c47. No engine change. Nothing pushed, pinned or dispatched. No nightly-owned data committed.

## 1. Wowhead feed coverage (owner ruling)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | fetch-wowhead refused a live feed listing under 90% of the seeded supplement, and only a manual `allow_shrink` dispatch carried the saved records. 70707's refetch listed 397 of 689 (58%) and needed it. | `data/pipeline/wowhead_items.py` `merge_live_feed` (before S febd69046) raised `WowheadCoverageError`; round 6 integration "After the push". | CONFIRMED (owner ruling) | S febd69046 |

What changed (S febd69046):

- **The carry is automatic.** `merge_live_feed` always carries the saved record of every seeded item the live feed no longer lists. Live records still win.
- **The floor is on the carried catalogue.** `build_floors.MIN_CARRIED_COVERAGE = 0.9` (renamed from `MIN_LIVE_COVERAGE`). The merge must keep 90% of the seeded supplement (the planner gear the client's ItemSparse lacks). A feed that is only missing items can no longer trip it.
- **A real drop still fails loudly.** The only way a record is lost is when its item has left the client's `Item` table. When that loses more than 10% of the seeded supplement, `WowheadCarryError` names the count, the share and the first 20 ids, and says to re-fetch. Nothing is written.
- **An item that left the client is dropped.** Decision: a seeded record whose item has no `Item` row (shipped or hotfix dump) is dropped, because the client tables are the truth. It is listed in the carry file's new `dropped` header and counted against the floor, so a truncated Item table cannot pass quietly. A record the client now covers with its own ItemSparse row is superseded, not lost, and does not count. On 70707 no record falls in either case: every one of the 10,556 records has an Item row (checked against the 70707 raw tables).
- **Every run logs its counts**, for example "the live feed lists 7 of 12 seeded supplement items (58%); 5 records carried from <seed>, 0 dropped (no Item row), 7 live, 12 written".
- **`allow_shrink` is gone from fetch-wowhead.** The `--allow-shrink` flag, its parameter on `fetch_wowhead` and `merge_live_feed`, and `WowheadCoverageError` are removed. data.yml's fetch-wowhead step no longer passes `$SHRINK`. That is the one workflow line this lane touched, plus its comment. The `allow_shrink` dispatch input itself stays, because it still gates normalize's two floors and the "No silent catalogue shrink" step. `loot`'s flag does nothing (`loot/__init__.py:144` says so); that and client.yml:324's hint belong to the ci lane.
- **tools/activate-build.sh** no longer tells the operator to add `-f allow_shrink=true` when the floors step stops a move.

Tests (written first, seen failing, then passing), in `data/tests/test_wowhead_carry.py`:
- a complete feed carries nothing;
- a feed at 58% carries the 5 saved records and passes;
- a carried catalogue that really shrank (5 of 12 items gone from the client's Item table) fails, and the file is untouched;
- a record whose item left the client is dropped, recorded and counted (11 of 12 passes);
- the existing fetch test now carries at 50% without a flag.

Touched suites: 103 passed across the 8 test files that touch wowhead items and floors; `tools/tests/test_activation.sh` 139 ok; ruff is clean.

A real-data check: the committed 70707 carry, with the 292 carried records taken out of the "live" side, keeps 289 of 289 seeded supplement items against the 70707 tables (100%).

## 2. The 14 loot sources dropped at c0bd0bd3d (for the owner; nothing changed)

**Why they dropped.** No upstream feed changed. `raw/classicdb/sources.json` differs only in `generated_at`. Item counts are unchanged (25,853). The cause is that two of our own generators disagree about battlegrounds:

- data.yml runs `loot` in full. It reads the client's `Map.csv`, where Alterac Valley (map 30) and Arathi Basin (map 529) are `InstanceType 3` (battleground). classic-db creatures and objects spawned on those maps resolve to the battleground zone. `instanced_zones` (`data/pipeline/loot/classicdb.py:184`) then gives a type-3 zone no loot source, so those drops vanish.
- bis.yml's nightly runs `loot-merge`. It has no Map.csv and recovers instance types from the committed loot.json (`types_from_committed_loot`), which only knows dungeons and raids. So the same spawns fall into flat `world:<npc>` buckets and into the level ranges of the world-drop pools.

So loot.json flips on every pair of runs: 3,198 at 2923e057f (data.yml), 3,212 at b6c17f743 (nightly), 3,199 at d40f4d6a4, 3,213 at 029cd1dff, **3,199 at c0bd0bd3d, and 3,213 again at 45875a020 (the nightly 4 minutes later; origin/main today)**. As of origin/main the 14 sources are back. The next data.yml run drops them again. `allow_shrink` had nothing to do with it: loot's flag does nothing.

**Net −14:** 15 named sources are gone, 3 world-drop level bands are renamed (the same items under new ranges), and 4 bands are new.

| # | source id (c0bd0bd3d^) | type | zone / map (classic-db spawn) | items it listed | still sourced after? |
|---|---|---|---|---|---|
| 1 | world:alterac-ram | world (creature) | Alterac Valley, map 30 | 17411 Steamsaw, 17642 Alterac Ram Hide | no |
| 2 | world:champion-guardsman | world (creature) | Alterac Valley (no spawn map; placed by the item's other records) | 17411 | no |
| 3 | world:commander-karl-philips | world (creature) | Alterac Valley, map 30 | 17411 | no |
| 4 | world:frostwolf-warrior | world (creature) | Alterac Valley (no spawn map; placed by the item's other records) | 17411 | no |
| 5 | world:stormpike-guardsman | world (creature) | Alterac Valley, map 30 | 17411 | no |
| 6 | world:frostwolf | world (creature) | Alterac Valley, map 30 | 17643 Frostwolf Hide | no |
| 7 | world:frostwolf-bloodhound | world (creature) | Alterac Valley, map 30 | 17643 | no |
| 8 | world:stabled-frostwolf | world (creature) | Alterac Valley, map 30 | 17643 | no |
| 9 | world:stabled-alterac-ram | world (creature) | Alterac Valley, map 30 | 17642 | no |
| 10 | world:coldtooth-supplies | world (object) | Alterac Valley, map 30 | 17542 Coldtooth Supplies | no |
| 11 | world:irondeep-supplies | world (object) | Alterac Valley, map 30 | 17522 Irondeep Supplies | no |
| 12 | world:frostwolf-banner | world (object) | Alterac Valley, map 30 | 17850 Frostwolf Banner | no |
| 13 | world:stormpike-banner | world (object) | Alterac Valley, map 30 | 17849 Stormpike Banner | no |
| 14 | world:hardpacked-snowdrift | world (object) | Alterac Valley, map 30 | 21038 Hardpacked Snowball | no |
| 15 | world:grunnda-wolfheart | world (creature; she is the Frostwolf supply vendor; classic-db also lists her stock as 0.02% "drops") | Alterac Valley, map 30; `opens: later` | 36 items, below | 23 of 36 keep the Frostwolf/Stormpike vendor and reputation sources |
| - | world:anvil | world (object), not dropped | Arathi Basin, map 529 | loses 3906 Balia'mah Trophy | no |

Grunnda Wolfheart's 13 items that lose their only source: 17348 Major Healing Draught, 17349 Superior Healing Draught, 17351 Major Mana Draught, 17352 Superior Mana Draught (consumables); 19029 Horn of the Frostwolf Howler (epic, mount); 19031 Frostwolf Battle Tabard (tabard); 19046 Frostwolf Battle Standard; 19301 Alterac Manna Biscuit; 19307 Alterac Heavy Runecloth Bandage; 19316 Ice Threaded Arrow and 19317 Ice Threaded Bullet (ammo); 19318 Bottled Alterac Spring Water; 19319 Harpy Hide Quiver and 19320 Gnoll Skin Bandolier (rare quiver and ammo pouch, ilvl 60, level 55). Her other 23 items keep `rep:frostwolf-clan:*` with vendors 13218 and 13219, plus `rep:stormpike-guard:exalted` and vendors 13216 and 13217 for the exalted ones: 19083, 19085, 19087-19090, 19095, 19096, 19099, 19101, 19103, 19308-19312, 19315, 19321, 19323-19325. 19483 keeps its quest source.

**Who references them:** nothing published. No item above is named in any `data/builds/1.60.1.70707/bis/*.json` (raid or leveling bands), in `addon-data.json`, in `data/curated/` (guides, presets, consumables) or in web/src. 19046 and 21038 appear only in `simconsumes.json`, the catalogue of every usable consumable, not any spec's published set. The spells.json hits on 19029, 19031 and 19319 are spell ids, not these items.

**The other differences in the same diff** (items keep a source; no gear):
- **PvP rank 5 Alliance and Horde, rank 10 Alliance.** These did not drop. They change shape: `faction_source` goes from `title` to `classic-db`, and `reitemised_from` gains 272694←15196, 272449←15197 and 272695←15198. The full run reads ranks from ItemSparse's RequiredPVPRank; the merge recovers them from the committed loot.json.
- **Two boss rows lose an item.** Blackrock Spire's Scorpok Stinger loses 9206 Elixir of Greater Strength (it moves to world_drop:44-51). Shadowfang Keep's npc 3864 loses 5120 Long Tail Hair.
- **36 world-drop pool items change band or chance.** The AV creatures widen or narrow the ranges: AV turn-ins 17306, 17326-17328, 17422, 17423, 17502-17504, 18143-18147, 18206, 18207 go from 50-62 to 55-55/55-61; the AV gray quest starters 18225-18234 drop to 0.02-0.04%. Also 5116, 5117, 5430, 5500, 5758, 9206, 11733, 11734, 11737, 12203, 12208 and 13467 (trade goods and librams).

For the owner to decide: should battleground spawns be sources (the merge's answer) or not (the full run's)? Either way, the two generators should agree. Today they flip every night a data.yml run happens.

## 3. Ledger pipeline follow-ups

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P2 | "Items still use the useHotfixes CSV without removals": does the server's push remove item rows that we keep? **No.** wago's push for 1.60.1.70707 has 868 ItemSparse and 868 Item entries: 678 valid (status 1), 190 invalid (status 3), **0 removed (status 2)**. 70685 and 70658 have no item entries at all. 70532 has 5,170 each: 5,129 valid, 41 invalid, 0 removed. | `hotfix_push.fetch_push("1.60.1.70707", ("ItemSparse","Item"))`, read 2026-11-01 (`r7/pipeline/itempush-<build>.json`). | Aligned | - |
| P3 | The 189 records whose newest 70707 entry is status 3 carry no values. Our layering keeps the row an earlier push gave (122 ItemSparse and 98 Item rows from 70532's push) or the hotfix dump's (9). That is the same rule the spell layer applies to status 3 (round 6: the previous row stands, counted as withheld). The 7 BiS-referenced items among them match Wowhead Forever's live tooltips: 1203 Aegis of Stormwind 1867 armor, +6 Str, +15 Sta; 273033 Gravespike Repeater 28-43 at 2.30, +8 AP. The others are 10750, 10782, 251963, 273031 and 286977. | `r7/pipeline/invalidated.tsv`; nether.wowhead.com/forever/tooltip/item/{1203,273033}, read 2026-11-01; committed items/warrior.json. | Aligned | - |
| P4 | ChrRaceRacialAbility (not fetched; racials hand-written). The 70707 push has 11 rows: Ruthless 1% crit, The Human Spirit 10%, Dwarf Artisan's Precision 1% hit (replaces Earth Fury), Orc Command +5% hunter and warlock pet damage (replaces Savage), Endurance 5% Health only, Plainsrunning no mounted speed, Wind Blessed 2% haste (both factions), Read Leyline 50% for 15 s, Skysight 15%. The engine (`sim/core/racials.go` at b57765fb9: lines 186-199, 217, 228, 246, 313, 429-437) and `data/curated/races.json` agree with every value. The remaining shipped rows (Berserking, Elune's Light, Beast Slaying, Big Game Hunter, Elemental Insight, Expansive Mind, Touch of the Grave) also agree. | `fetch_push` for ChrRaceRacialAbility, read 2026-11-01 (`r7/pipeline/otherpush-70707.json`); `https://wago.tools/db2/ChrRaceRacialAbility/csv?build=1.60.1.70707`. | Aligned | - |
| P5 | Skysight: the pushed racial summary says "15% for 15 sec". Our site text says 30 seconds. The client spell tables (Elemental Blessing 1259688, SpellDuration 30000) and Wowhead's server tooltip for 1259686 ("by 15% for 30 sec") both say 30 s, so the summary row is the stale one. Movement only; nothing simmed. | layered `SpellMisc`/`SpellDuration` 70707; nether.wowhead.com/forever/tooltip/spell/1259686, read 2026-11-01. | Aligned (two sources against the summary row) | - |
| P6 | Read Ley Line: the engine comment says "100% health and mana regeneration". The client and Wowhead say 50% for 15 s (`racials.go:524`). The racial is not modelled (no cooldown is registered), so no number moves. | as P4; nether.wowhead.com/forever/tooltip/spell/1259705. | Stale comment only, no change | - |
| P7 | SpellCategories, 69 pushed rows. The ones the sim reads: Ice Lance joins category 19 (0 → 19), the category Fire Blast uses, with an 8 s category recovery on both. The engine already shares the timer (`sim/mage/fire_blast.go:19-26`). The new Lacerate ranks 1328616 and 1328617 are DefenseType 2, Mechanic 15 (bleed); Ferocious Bite 1328618 is DefenseType 2, like the other ranks. Swift Judgement DefenseType 1 → 0 and Templar's Bulwark DefenseType 0 → 1 matter only to protection paladin (talent-gated). The rest are traps, NPC spells and utility. | `fetch_push` SpellCategories 70707; wago SpellCategories and layered SpellCooldowns CSVs. | Aligned (Swift Judgement and Templar's Bulwark: for the hybrids lane if either is ever cast) | - |
| P8 | SpellEquippedItems (21 pushed: 8 removals of rows the client does not ship, 10 invalid, 3 valid, none on a class spell) and SpellLabel (27, all new Forever spells or Dash/Track/Feline Grace labels) change nothing the engine reads. SpellProcsPerMinute has no push. Its only user in the client is Holy Smite 1326019 (2 PPM, 56 Holy). That is the on-equip proc of Lord Sakrasis' Scepter 5028 (held in off-hand, level 42), which the engine does not model and no BiS lists. | wago SpellProcsPerMinute and SpellAuraOptions (SpellProcsPerMinuteID 455) 70707; ItemEffect 236933 → item 5028; nether.wowhead.com/forever/tooltip/spell/1326019. | Aligned; 5028's proc OPEN (minor: leveling only, unmodelled) | - |

## Before/after

No engine change, and no site change that touches sim inputs. Item 1 changes only how fetch-wowhead treats a short feed: with today's feed it writes the same carry file the `allow_shrink` dispatch wrote. No DPS or HPS moves.

## Still open

- The loot.json battleground flip (item 2) waits for the owner's ruling. Once ruled, make `loot` and `loot-merge` agree: share one instance-type rule, or keep a battleground set in the merge.
- data.yml's `allow_shrink` dispatch input still gates normalize and the catalogue-shrink step. Whether it stays is for the ci lane. `loot --allow-shrink` is a no-op and could go with it, and client.yml:324's hint names it.
- Lord Sakrasis' Scepter's Holy Smite proc is unmodelled.

Sources: wago.tools hotfix push log (search per table, builds 70707, 70685, 70658 and 70532, read 2026-11-01); `https://wago.tools/db2/{ChrRaceRacialAbility,SpellCategories,SpellProcsPerMinute,SpellEquippedItems,ChrRaces}/csv?build=1.60.1.70707` (read 2026-11-01); Wowhead Forever tooltips: spell 1259686, 1259688, 1259705 and 1326019, item 1203 and 273033 (read 2026-11-01); git history of `data/builds/1.60.1.70707/loot.json` (2923e057f to 45875a020).
