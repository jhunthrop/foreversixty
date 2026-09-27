# Wowhead item supplement: the gear the client files do not carry

**Date:** 2026-09-27. **Status:** approved for implementation (owner: "we're missing so many item ids").

## The problem, measured

The site's item data comes from wago.tools' export of the beta client's `ItemSparse`
table (`data/pipeline/wago.py`). On 2026-09-27, against the newest export
(1.60.1.70009, the client build players run):

| | |
|---|---|
| Items in wowhead's Forever gear planner (`nether.wowhead.com/forever/data/gear-planner?dv=100`, all `versionNum` 16001) | 11,239 |
| Of those, absent from wago's 70009 `ItemSparse` | **4,712** |
| Of those, equippable and uncommon or better | **4,687** |
| Their item levels | 10 to 100, most between 20 and 69 |
| Of one hunter's ten equipped items (Bow Jackzon) | 7 absent |

Blizzard is delivering much of Forever's itemization as server hotfixes. Hotfixed rows live
in the client's hotfix cache, not in the shipped DB2 files wago exports, so no future wago
build closes this gap on its own. Wowhead reads the live game and has every one of them.

## Calibration

On 6,527 items both sources carry, names agree on 400 of 400 sampled, and wowhead's stat
values equal the values `pipeline/normalize/item_curves.py` resolves from the client's
curves, one for one, for stamina, agility, strength, intellect, spirit, spell power,
attack power, mp5, crit, hit, defense, dodge, parry and the resistances. Wowhead also
states weapon damage and speed, which the 1.60 planner rows lack (`GearItem.damage_min`
is 0 on every 1.60 item today).

## Decision

Wowhead's Forever gear planner becomes a second item source, used for exactly the ids the
client's `ItemSparse` lacks ("supplement items"). The client stays the source of truth
for every id it has: effects, sets, sockets and sim resolution keep their current path.

One module owns the source: `data/pipeline/wowhead_items.py`.

- `fetch-wowhead --build <build>` downloads the payload to
  `builds/<build>/raw/wowhead-gear-planner.js` (raw is git-ignored) with a `.meta.json`
  beside it (url, fetched_at, sha256). The payload is undated upstream; the meta is our date.
- `load_items(path) -> list[WowheadItem]` parses `WH.setPageData("wow.gearPlanner.classicplus.item", {...})`
  (tolerant of the trailing commas the payload contains) into a typed record.
- `supplement(items, client_ids) -> list[WowheadItem]` is the ids the client lacks that
  pass the planner's own gates: `PLANNER_QUALITIES`, `MAX_PLAYER_LEVEL`,
  `SLOT_BY_INVENTORY_TYPE`, `is_junk_name` -- the same gates `build_class_items` applies.
- `to_gear_item(item)` and `to_item(item)` produce the pipeline's own `GearItem` and
  `Item` models. Stats map through `STAT_KEYS` (wowhead key -> planner key); a wowhead key
  the planner does not track maps to `None` and is counted, never raised.
- `class_allowed(item, class_id)` is wowhead's `classMask` bit `class_id - 1`, plus
  `proficiency.can_equip`, the same rule the client rows use.
- Wowhead's negative armour subclasses (-2 rings, -3 necks, -4 trinkets, -5 held,
  -6 cloaks, -7 tabards, -8 shirts) map to the client's: cloaks are cloth (1), the rest
  Miscellaneous (0).

## Where the supplement lands

1. `normalize` merges supplement items into `items.json` (`Item`) and every
   `items/<class>.json` (`GearItem`), logging the count. Sets: a supplement item's
   `itemset` id joins an existing set's `item_ids`; wowhead's set payload carries no set
   names, so a set the client lacks entirely is left out and counted (a known gap).
2. `simdb` builds a `SimItem` for each supplement item: type from inventory type, armour,
   stats through `simdb/statmap.py`, weapon type and hand type from class/subclass and
   inventory type, weapon damage and speed from wowhead, class allowlist from `classMask`,
   unique from `maxcount`, required level, set id. No on-equip effects (wowhead states them
   as prose only): counted in the log. `simitems.json` lists them.
3. `icons`: a supplement item names its icon (e.g. `inv_misc_cape_10`). The client's
   `ManifestInterfaceData` maps that name back to a file data id for the existing CASC
   download; a name the client lacks falls back to
   `https://wow.zamimg.com/images/wow/icons/large/<name>.jpg`, resized to the same 64x64 WebP.
4. The build moves to 1.60.1.70009 (the New build checklist), `fetch-wowhead` becomes a
   step of the data workflow between `fetch` and `normalize`, and
   `web/src/data/active-build.json` bumps.

## Out of scope for this round

Loot sources for supplement items (wowhead's `source`/`sourcemore` codes), set bonuses for
sets the client lacks, on-equip procs for supplement items, and using wowhead for weapon
damage on items the client does carry. Each is a counted gap in the log, not a silent one.

## Global constraints

- `uv run ruff check .` clean; `uv run pytest` with the 80% coverage gate.
- Every new gate or mapping has a fixture-backed unit test; nothing branches on a build string.
- Raw payloads are git-ignored; committed outputs are `builds/<build>/*.json`, `simdb.bin`, icons.
