# Wowhead carry-forward (2026-10-30)

Lane `wt/wowhead-carry`. It fixes the data problem found in `design/reviews/2026-10-30-70707-dps-drop.md`. Nothing was pushed or dispatched.

## Exposure: is the live site at risk tonight?

**Yes, but only through activation. Nothing will regenerate 70658's items.**

- **The active build, 1.60.1.70658, is not re-fetched by anything scheduled.**
  - `fetch-wowhead` runs only in `data.yml`'s `fetch` job.
  - That job runs on a hand dispatch, or when `client.yml` calls it for a newer build that has no `items.json` yet.
  - `bis.yml` (nightly, 06:23) crawls wowhead's item and quest pages, then runs `fetch` (wago tables) and `loot`. Neither reads the gear-planner feed.
  - `engine.yml` (nightly, 05:07) re-fetches the wago tables and reads no wowhead file.
  - So 70658's committed items stay as they are.
- **The risk is `client.yml`, nightly at 03:41 UTC (it actually runs at about 10:00 UTC).**
  - Every morning it tries to activate the newest build. For 70707, which is already committed with the stripped stats, it skips the fetch and runs `tools/activate-build.sh 1.60.1.70707`.
  - On 29 Oct (70685) and 30 Oct (70707) the move failed only because the Mage guide builds do not re-encode onto the new talent trees. The errors were "Arcane Focus / Improved Fireball / Elemental Precision is not in the tree".
  - Nothing about items stopped it. Once the Mage guides are fixed for 70707 (the `mage-70707` lane), the next morning's run would:
    - activate 70707 with 396 wowhead items instead of 688: Kindling Stave at 0 SP, Spire of Hakkar at 18, and 22240 missing;
    - dispatch `bis.yml` with force, which re-ranks on those stats.
- **A new wago build tonight** would be fetched from the shrunken feed. Before this fix, all the gates would have passed it.
- **The soonest single stop is the activation guard in commit `51eebb18c`.**
  - `tools/activate-build.sh` now runs `pipeline build-floors --build NEW --against ACTIVE` before it changes anything.
  - On 70707 today it refuses: "1.60.1.70707 takes 396 items' stats from wowhead, below 90% of 1.60.1.70658's 688".
  - It reads only committed files, so it holds even before 70707 is regenerated.

## What changed

1. **Carry-forward** (`829a47855`, `9e3416828`), in `data/pipeline/wowhead_items.py`:
   - `fetch_wowhead` merges the live feed into a committed `raw/wowhead/items.json`, using `carry_forward` and `merge_live_feed`.
   - The seed is the build's own file, or else the newest other build's file, by version.
   - The live feed wins wherever it has the item.
   - A seeded record the feed lacks is carried. It is dropped only when the client's `Item` table loses the item, or `ItemSparse` gains a row for it. Both tables are read merged with the hotfix dump, as `normalize` and `simdb` read them.
   - The fetch refuses, and writes nothing, when the live feed lists less than `MIN_LIVE_COVERAGE = 0.9` of the seeded supplement (the planner gear the client lacks). `--allow-shrink` accepts the feed and still carries the rest.
   - Records are trimmed to the fields the pipeline reads, one per line, about 3.2 MB per build.
   - Every consumer reads the merged file through `raw_path`: `normalize/wowhead.py`, `simdb`, `icons`, `icons_fix`, `audit`. `levels` keeps reading the downloaded payload's base stats through `payload_path`.
   - `.gitignore` un-ignores only `data/builds/*/raw/wowhead/items.json`.
2. **Seeds** (`80393b2e0`):
   - **70658** had the most wowhead items (688). Its live payload was never saved, so its file was reconstructed:
     - today's live feed (10,248 records);
     - plus the 292 lost items, all 292 taken from 1.60.1.70291's saved raw payload (8 Oct).
     - Each record was accepted only if it reproduces 70658's committed class rows and class placement exactly. Weapon damage was checked against 70658's `simdb.bin`, because `loot` rewrites damage in `items/`.
     - None needed to be rebuilt from generated data.
     - The generator is at `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/seed70658.py`. That scratchpad is wiped after a few days, so it may already be gone.
   - **70707** is the live feed merged over 70658's file: 396 items live, 292 carried. It was written with `--allow-shrink`; without it the fetch refused at 396 of 688 (58%).
   - **Check:** 70658 was regenerated (normalize, loot, simdb) from its seed. `items.json` and the `simdb.bin` items came out identical to the committed ones. In the class files, every stat matched; only 2 required levels differed (concern 1).
3. **Workflows** (`5ddb0ba4f`):
   - In `data.yml`, "Seed the fill-in caches" now runs before `fetch-wowhead`, and its list includes `wowhead/items.json`.
   - `fetch-wowhead` now takes `$SHRINK`.
   - No other workflow fetches the gear planner. `client.yml` reaches it only by calling `data.yml`.
4. **Gates against the previous build** (`51eebb18c`). The thresholds are in `data/pipeline/build_floors.py`, and all are floors:
   - `check_item_sparse_completeness`: a build's first run is compared with the newest other build at 98% (`data.yml`'s own 2% tolerance). A build's own last run still may not lose a single item.
   - `_check_class_items_not_shrunk`: a first run uses the previous build's per-class counts (`previous_class_item_counts`), at 90% as before.
   - New `check_wowhead_items_floor`: the wowhead-sourced count must be at least 90% of the reference run's (the build's own last run, or else the previous build's). 70685's 7 against 688 fails it.
   - New CLI `pipeline build-floors`, used by `activate-build.sh`.

## Proof (70707 regenerated locally in a scratch root: normalize, loot and simdb)

| | items.json | stats from wowhead | Kindling Stave (11750) | Spire of Hakkar (10844) | 22240 |
|---|---|---|---|---|---|
| 70658 committed | 25,678 | 688 | 60 SP, wowhead | 62 SP, wowhead | present |
| 70707 committed | 25,534 | 396 | 0 SP, classic-db | 18 SP, classic-db | absent |
| 70707 regenerated with the carry | 25,688 | **688** | **60 SP, wowhead** | **62 SP, wowhead** | **present** |

- The 10 extra items are 70707's new client rows.
- Against 70658's `simdb.bin`, the regenerated one has 7 new ids. The only other changes are 16 client rows in the 275968–275984 range, which the client itself changed.

Mage raid bands were simmed with the review's probe setup: 70658 trees, guide builds and bands, 800 iterations, seed 7, pinned engine `1f542def3`, and only `simdb.bin` swapped.

| band | 70658 simdb | 70707 committed | 70707 with the carry |
|---|---|---|---|
| mage-arcane raid | 663.0 | 628.6 | **663.0** |
| mage-fire raid | 869.0 | 823.2 | **869.0** |
| mage-frost raid | 525.9 | 501.2 | **525.9** |

Tests: the data suite passes (1,611 passed, 8 skipped), and `ruff check` is clean. The new tests are `tests/test_wowhead_carry.py` (12) and `tests/test_build_floors.py` (9).

## Concerns and follow-ups

1. **Client rows are not carried, by design.**
   - A seeded record whose item has a client `ItemSparse` row is dropped (the stated rule). Wowhead also supplies required levels for client rows that state 0, and the feed lost two of those: 11902 and 11904.
   - They now resolve to the item-level proxy: 47 becomes 51 on 15 class rows.
   - This is minor, but it is the one difference the 70658 regeneration showed.
2. **A separate, older bug: the wowhead stat map misses keys wowhead now sends.**
   - `STAT_KEYS` maps `critstrkrtng` and `hitrtng`. Wowhead's records carry `critstrkpct` and `hitpct`, plus `def` and school spell power (`shasplpwr`, `firsplpwr` and so on). These are dropped without being counted.
   - Of the 688 supplement items:
     - 15 carry crit percent;
     - 23 carry hit percent;
     - 30 carry defense;
     - 28 carry school spell power.
   - Kindling Stave's 1% crit is one of them.
   - This is not a regression: 70658 has it too. It needs its own lane.
3. **Regeneration is still to do:**
   - Re-dispatch `data.yml` for 1.60.1.70707 (and 70685 if it is kept) once this lands. The seed is already committed, but the feed is at 58% coverage, so the dispatch needs `allow_shrink=true`.
   - That flag also disarms the normalize and loot shrink gates for that run. They would pass anyway, because the carry restores the count. A separate flag for the feed would be cleaner.
   - Then re-run the BiS ranking.
4. **70685 has no carry file.** A re-fetch of it would seed from the newest other build (70707) and carry from there.
5. **Two seeding paths.** The workflow's seed step and the Python fallback in `fetch_wowhead` both seed. The workflow's step picks the newest other build even when that build has no file; the Python fallback then picks the newest build that has one.
6. **The client-truth check from the review is still open.** It needs a fresh `DBCache.bin` and `pipeline hotfixes`.
