# Leveling BiS: the best gear at every level band, from the simulator

**Date:** 2026-09-28. **Owner ask:** "use the simulator and top gear to show what is the BiS at each level range as the person levels ... at level 20 a BiS for all equippable level 20 gear, with faction isolation for quests, etc. ... as they progress in level, it shows the new items that are BiS."

## Yes, and with what exists

- **Candidates:** `data/builds/<active>/items.json` carries every item's required level, class mask, inventory type, quality and faction restriction (23,909 items incl. the wowhead supplement).
- **Sources:** the engine's item database (`assets/database/db.json`, wowsims-forever) names, per item, its quest (id and name; the item's faction restriction carries the quest's side: 336 Alliance, 332 Horde, 560 both among 1,228 quest items), its vendor, its dungeon or raid boss, its crafting profession, its reputation standing, and for 3,292 items the creature and zone it drops from. The site's `loot.json` flattens quests into one bucket today; it grows.
- **Scoring:** the simulator's stat-weights mode (`sim/request/weights.go`) run for the spec at the band's level on the accuracy program's ladder character gives the band's weights; an item's score is the weighted sum of its stats (the planner already computes item stats per level through curves). Instant to rank thousands of candidates.
- **Verification:** Top Gear (`sim/bulk`, mode gear) sims the ranked set at that level to confirm the DPS and to settle the close calls per slot.

## Built once per spec and band, never per character

The owner (2026-09-28): "the BiS lists should be built once per class and level. It shouldn't be character by character." The ranking and the Top Gear verification run once per written spec per band, nightly, and the result is published as data. Per spec rather than per class because the stat weights that rank the items differ by spec (Balance and Feral want opposite gear); a class with one damage spec on the site is one list. A signed-in player's "your character" view is a lookup against that published list -- which of the saved export's slots the list beats -- and never a simulation of that character.

## Definitions

- **Bands:** 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60. A band's list is "the best you can wear at level L".
- **Eligible at L:** required level <= L; class allowed; armor proficiency by level (mail for hunters and shamans from 40, plate for warriors and paladins from 40; before that the lower class); weapon skills the class can train; faction: the item's restriction and the quest's side must match the chosen faction; raid items excluded below 60 (a leveling list is about what a leveling character can get); PvP rank items shown only with their rank; crafted items shown with the profession as a note (a "with professions" toggle); reputation items with the standing.
- **Obtainable, not lucky:** a slot's BiS is the best item with a named source (quest, dungeon, vendor, crafted, rep). World drops with a known zone are listed separately as "if it drops" candidates, never as the answer, because a leveling player cannot farm a 0.1% green.
- **New at this band:** the diff against the previous band per slot: the item, its source, the DPS gain the verification sim measured for the whole set.

## The pipeline

1. **Data (`data/pipeline/loot`):** enrich `loot.json` from db.json: quest sources become `{quest_id, name, faction}` per item; add `vendor` and `zone-drop` kinds (npc, zone name, zone level range from `zones`). Keep the existing kinds. Tests pin the counts.
2. **Ranking (`sim/cmd/leveling-bis`, Go):** for each written spec and band: build the ladder character (talents truncated; the same rule the ladder test states), run the weights mode once, score every eligible candidate per slot, pick greedily with the two-slot and one-hand/two-hand rules `sim/bulk/expand.go` already encodes, then run one Top Gear pass of the chosen set plus the runner-up per slot to confirm. Emit `data/builds/<active>/leveling-bis/<spec>.json`: per band, per slot, the item, its source, the score, the verified set DPS, and the diff from the previous band.
3. **Schedule:** nightly beside the sim validation job (Cloud Run `sim-run`/`sim-validate` recipes in `api/README.md`); the output is data, committed by the pipeline like loot.json, and published under `/data/<build>/`.
4. **Site:** a "Leveling gear" section per spec guide and a page `/gear/<class>/<spec>?level=20&faction=horde`: a band slider, per-slot rows with source badges (quest name and zone, boss and dungeon, vendor, profession, rep), a "New at 25" strip, an Alliance/Horde toggle, and for a signed-in player a "your character" mode that marks which slots are upgrades over the saved export and links each to the simulator's Top Gear with the candidate preloaded.
5. **Addon (design pass wave C):** Data.lua carries, per class/spec/band, the BiS item ids; the tooltip says "BiS for level 20-24 Marksmanship"; the level-up toast lists the new BiS items with their source.

## What is honest about it

The list is only as good as the sources: an item db.json has no source for is invisible, so the page says "N items had no known source" per band. Faction isolation covers quest side and item restriction; neutral quests show for both. Weights-based ranking can misorder two close items; the Top Gear pass settles the top two per slot but not every pair, and the page marks the verified rows.

## First steps (now)

- Lane `bis-data`: the loot enrichment (quest id/name/faction, vendor, zone-drop) with tests.
- Lane `bis-proto`: the ranking command for ONE spec (hunter-marksmanship) at bands 20, 30 and 40, both factions, writing the JSON and a readable markdown so the output shape can be judged before the page is built.
