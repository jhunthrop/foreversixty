# data

Pipeline that turns a WoW client build into the JSON and icons the site and the API read.

## Commands
```
uv sync
uv run python -m pipeline fetch --product wow_classic_era            # latest build
uv run python -m pipeline fetch --product <forever product> --build <build>
uv run python -m pipeline normalize --build <build>
uv run python -m pipeline icons --build <build>
uv run python -m pipeline diff --from <build> --to <build>
uv run ruff check . && uv run pytest
```

`fetch` and `icons` are the only commands that use the network. `normalize` and
`diff` are offline and fully unit-tested against the fixtures in `tests/fixtures/`.

## Layout
```
builds/<build>/raw/*.csv        downloaded DB2 exports (gitignored)
builds/<build>/*.json           flat entities: zones, dungeons, items, spells,
                                classes, races, talents, sets, combos
builds/<build>/talents/<class-slug>.json   per-class trees with rank descriptions
builds/<build>/items/<class-slug>.json     per-class equippable items
builds/<build>/icons/<name>.webp           64x64 icons for every talent and item above
builds/<build>/manifest.json    build, product, fetched_at, sha256 per emitted file
curated/{classes,races,combos}.json        hand-maintained Forever facts (committed)
diffs/<from>__<to>.json         added / removed / changed per entity
.icon-cache/<file data id>.blp  raw icon cache (gitignored), so reruns are cheap
```

The site reads `web/src/data/active-build.json` to decide which build directory to
publish; it is a one-line file, `{ "build": "<build>" }`, owned by `web/`. Point it at a
new build only after that build's JSON is committed here.

## Curated Forever facts
`curated/classes.json` and `curated/races.json` attach `forever_changes` to a row by
`slug`; `curated/combos.json` lists the legal race and class pairs. Every change entry
and every combo marked `new_in_forever` must carry at least one source:

```json
{ "text": "...", "sources": [{ "label": "...", "url": "https://...", "kind": "blizzard" }] }
```

`kind` is one of `blizzard`, `datamined`, `community`, `site` — the same vocabulary as the site's source pills in `web/src/lib/sources.ts`, so a curated source renders without mapping. An
unsourced claim stops the pipeline on purpose: the planner never states a Forever fact
it cannot point at.

Skyborne was carried as a placeholder race with the local sentinel id `900` and
`"placeholder": true` until the 1.60 client (Forever beta) exported real rows: it
turned out to need two, `high-order-skyborne` (id 95, Alliance) and
`windshaper-skyborne` (id 96, Horde) — the client's own per-faction split for a
neutral race, the same pattern it uses for Pandaren and Dracthyr. `curated/races.json`
now carries one entry per slug (only `slug` and `forever_changes`, matched onto the
client row the normal way) instead of the placeholder fields, and `curated/combos.json`
points its Skyborne pairs at 95 and 96 instead of 900. Because era's `ChrRaces` has no
Skyborne row at all, this makes the shared `curated/` directory validate only against
a build that has both real rows — 1.60 and any later one — not against the frozen
`builds/1.15.9.69722`; see `tests/test_curated.py`'s `BETA_BUILD`.

## Sept 17 checklist
1. Find the Forever product key on https://wago.tools/builds (it appears when the beta client is on the CDN).
2. Run the workflow_dispatch in GitHub Actions with that product, or run fetch, normalize and icons locally.
3. If a table 404s or a column is missing, fix `TABLES` in `pipeline/wago.py` or the
   normalizer, add a fixture row, keep the golden tests green.
4. If `normalize` logs "items not emitted" it also exits non-zero, which stops the
   workflow before it can commit a build with no `items/`. An item stat modifier id is
   missing from `STAT_BY_MODIFIER_ID` in `pipeline/normalize/gear.py`. Add it with the
   right planner stat key, or `None` if the planner does not track that stat, and run
   `normalize` again. Everything except `items/` is still written by the failed run.
5. Curate the Forever facts under `curated/` with their sources.
6. Commit `builds/<build>/` and bump `web/src/data/active-build.json`.

## Known gaps
- The Classic Era client has no `JournalInstance` table, so `dungeons.json` is empty for
  `wow_classic_era` builds. `JournalInstance` is the only entry in `OPTIONAL_TABLES` in
  `pipeline/wago.py`; any other missing table fails the fetch on purpose.
- Classic items grant crit, hit, spell power, mp5 and defence through equip spells rather
  than through `StatModifier_bonusStat_*`, so those keys do not appear on Era items. Only
  what the client states in a column is emitted.
- Armour and weapon proficiency is not in the client tables; `pipeline/proficiency.py`
  holds the Classic 1.x proficiencies and must be revisited for Forever.
- Some items have `IconFileDataID` 0 in `Item.csv`, meaning the client itself ships no art
  for them. Those items are emitted with `PLACEHOLDER_ICON` (`inv_misc_questionmark`, the
  client's own placeholder) rather than an empty name, which would resolve to
  `icons/.webp` and 404. On `wow_classic_era` 1.15.9.69722 this affects **33 items across
  174 per-class records** — mostly `Monster - Item` rows, but also real items such as
  Gorehowl (227688) and Flowing Scarf (209423). Each one logs a warning naming the item.
  An item whose `IconFileDataID` is nonzero but absent from `ManifestInterfaceData` takes
  the same fallback; no Era item currently does.
- A talent takes the same fallback when its first rank's spell has no icon the client can
  resolve, and is named `Unknown talent <id>` when that spell has no `SpellName` row.
  Both log a warning naming the talent. No Era talent currently hits either branch, but a
  Forever client with new talent spells is exactly where it would first happen, and an
  empty `icon` or `name` would reach the site as `icons/.webp` and a blank cell.
- `tests/test_build_conformance.py` checks the committed build directory against the
  Phase 1 interface contract (keys, slot and stat vocabularies, prereq and rank
  integrity, icon coverage, manifest coverage). Its constants are transcribed from the
  contract, not imported from `pipeline`; if it fails, the contract wins.
- The 1.60 client (Forever beta) is a modern client build: `ChrRaces` carries internal
  rows a player can never pick (`PlayableRaceBit` -1) that share a `Name_lang` with a
  real playable race — e.g. id 33 "Human", an unused "ThinHuman" body-type row with
  placeholder flavour text, alongside the real Human at id 1. `normalize_races` drops
  any row with `PlayableRaceBit` -1 before slugifying, so it can never collide with a
  real race in `curated.merge_curated`'s by-slug lookup. Classic Era's `ChrRaces` has no
  `PlayableRaceBit` column at all, and every one of its rows is kept, as before.
- The 1.60 client's `ItemSparse` has no `Resistances_*` or `StatModifier_bonusAmount_*`
  columns at all: armour and stat amounts are computed client-side from curve tables
  (`RandPropPoints`, `ItemArmorTotal`, `ItemArmorQuality`) that are not in `TABLES` and
  this pipeline does not resolve. Per "only what the client states in a column is
  emitted," a missing column reads as no data (armour 0, no stat contribution) rather
  than an error. `_has_gear_value` still drops an armour piece with nothing to compare
  it on, so **no non-weapon item survives from the 1.60 build's `items/` today** —
  weapons are exempt from that clause and are unaffected. Resolving the curve tables is
  future work, not attempted here: the formula is not stated in a column either, and
  guessing at it risks silently wrong stats rather than an honest gap.
