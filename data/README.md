# data

Pipeline that turns a WoW client build into the JSON and icons the site and the API read.

## Commands
```
uv sync
uv run python -m pipeline fetch --product wow_classic_era            # latest build
uv run python -m pipeline fetch --product <forever product> --build <build>
uv run python -m pipeline normalize --build <build>
uv run python -m pipeline icons --build <build>
uv run python -m pipeline tree-art --build <build>                    # then normalize again, so the manifest lists trees/
uv run python -m pipeline simdb --build <build>
uv run python -m pipeline simconst --build <build>
uv run python -m pipeline gametables --build <build>
uv run python -m pipeline specs
uv run python -m pipeline specs --check                               # CI's drift gate; writes nothing
uv run python -m pipeline diff --from <build> --to <build>
uv run python -m pipeline simproto --engine "$FOREVER_ENGINE_PATH"    # only when the engine's protos change
uv run ruff check . && uv run pytest
```

`fetch`, `icons`, `tree-art` and `gametables` are the only commands that use the network.
`normalize`, `diff`, `simdb`, `simconst` and `specs` are offline and fully unit-tested
against the fixtures in `tests/fixtures/`. `simproto` reads a local engine checkout and is
the only command that needs one.

## Layout
```
builds/<build>/raw/*.csv        downloaded DB2 exports (gitignored)
builds/<build>/*.json           flat entities: zones, dungeons, items, spells,
                                classes, races, talents, sets, combos
builds/<build>/talents/<class-slug>.json   per-class trees with rank descriptions
builds/<build>/items/<class-slug>.json     per-class equippable items
builds/<build>/icons/<name>.webp           64x64 icons for every talent and item above
builds/<build>/trees/<background>.webp     320x384 background for every talent tree above
builds/<build>/manifest.json    build, product, fetched_at, sha256 per emitted file
builds/<build>/simdb.bin        the engine's SimDatabase protobuf (items, enchants)
builds/<build>/simconsumes.json consumable items and the spells they cast
builds/<build>/spellconst/<class-slug>.json  per-spell constants keyed by spell id
builds/<build>/gametables/<name>.txt  the client's own base-mana, crit and rating curves
curated/{classes,races,combos}.json        hand-maintained Forever facts (committed)
curated/specs.json              the canonical 27-spec list
curated/apl/<spec_slug>.json    one default rotation per spec, with sources
proto/                          the engine's .proto files, vendored, with ENGINE_SHA
pipeline/simproto/              generated Python bindings (committed)
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

## Simulator outputs

`simdb` and `simconst` read the same `builds/<build>/raw/` CSVs `normalize` does, so run
`fetch` first. `simdb` also reads `sets.json`, so run `normalize` before it. `gametables`
reads nothing local but writes into the build directory, so it needs `normalize` to have
run too. All three refresh `manifest.json` when they finish, so the manifest still covers
the directory.

`gametables/` is the one output that does not come from DB2. Vanilla's per-class base
mana and its combat-rating conversions are in the client's `GameTables/*.txt`, which is
what the engine's own `tools/base_stats_parser.py` reads; `GameTables`, `CombatRatings`,
`CharBaseStats` and six other DB2 candidates all 404 on build 1.60.1.69893.

Every CASC read goes through `pipeline/casc.py`, which **pins the build**:
`GET /api/casc/<file data id>?version=<build>`. The version is not optional — the same
id is different content in different builds (`combatratings.txt` is 12,144 bytes on
1.60.1.69893 and 31,614 on Classic Era), an unknown version answers 400, and a bare
request serves whatever wago defaults to, which is retail. The cache is keyed by build
for the same reason. Tree art (`pipeline/art.py`) and icons (`pipeline/icons.py`) go
through the same helper: all three were pinned separately first, and one helper is what
stops a fourth caller forgetting. A 400, a redirect or an unreadable response fails the
run rather than writing a file; a 404 or an empty body is `CascMissing`, which icons
treat as "the client has no art for this one" and everything else treats as an error.

Pinned, six of the seven files the engine's parser names turn out not to exist in this
client at all — they answer 200 with an empty body on 1.60.1.69893 and on Classic Era
alike, because their file data ids were added to retail after Classic forked. So
`gametables/` carries the three the client does ship (`combatratings.txt`, `basemp.txt`
and `hppersta.txt`) and `pipeline/gametables.py` keeps the other six in
`ABSENT_FROM_THE_CLASSIC_LINEAGE` with a test behind it, so the engine lane reads the gap
off a constant. Per-race base Strength, Agility, Stamina, Intellect and Spirit are in
neither DB2 nor `GameTables` and are not emitted by anything here.

`pipeline/casc.py` imports `CACHE_DIR` and `_atomic_write` from `pipeline/icons.py`, so
`icons.py` imports the helper inside `download_icons` rather than at module scope — the
same local-import pattern `pipeline/normalize/__init__.py` already uses.

`simdb.bin` is exactly the engine's `proto.SimDatabase`. It has no item-set and no
consumable field: a set lives on each item as `set_id` + `set_name`, and the engine's
consumables are a hand-written enum table, so the raw material for regenerating that
table is written beside the protobuf as `simconsumes.json`. `random_suffixes` is empty
because `ItemRandomSuffix` 404s on build 1.60.1.69893.

Three things the sim needs that the planner's `items/` does not carry:

- **Weapon damage and speed.** `pipeline/simdb/weapons.py` resolves them from
  `ItemDamage{OneHand,TwoHand,Ranged,Wand,Thrown}` the way `item_curves.py` resolves
  armour: `dps = ItemDamage<kind>[ilvl][quality]`, `avg = dps × ItemDelay / 1000`,
  `min = int(avg × (1 − DmgVariance/2))`, `max = round(avg × (1 + DmgVariance/2))`.
  Verified against Classic Era's own literal `MinDamage_0`/`MaxDamage_0` on the 478
  weapons unchanged between the builds: 471 exact (98.5%).
- **Stats that live in a spell.** `pipeline/simdb/equip.py` resolves `ItemEffect` (linked
  through `ItemXItemEffect` on a modern client, `ParentItemID` on an older one) and
  `SpellItemEnchantment` equip spells through `SpellEffect`. On build 1.60.1.69893 that
  is 18 items and 1,329 of the 2,216 enchant rows (40 more name only a weapon skill,
  which `pb.SimEnchant` has no field for, and are dropped with a warning). Every aura
  type it meets must be classified in `STAT_AURAS` or `IGNORED_AURAS`; an unclassified
  one stops the run on purpose.
- **Armour and stats without a column.** These are not re-read here: `simdb` calls
  `normalize/gear.py`'s `resolve_item_values`, the same function the planner uses, so
  the two can never disagree about what an item is worth.

`python -m pipeline specs` regenerates `sim/specs/specs.go` and
`web/src/lib/sim/specs.ts` from `curated/specs.json`. Run it after editing that file and
commit all three. `python -m pipeline specs --check` writes nothing and exits non-zero if
either generated file has drifted from the curated list; CI's `test` job runs it, and
`tests/test_specs.py` asserts the same thing. It is the only gate on those two files, so
`data.yml` also triggers on `sim/specs/**` and `web/src/lib/sim/specs.ts`.

`python -m pipeline simproto --engine <path>` re-vendors the engine's `.proto` files into
`proto/` and regenerates `pipeline/simproto/`. Only run it when the engine's protobuf API
changes; `proto/ENGINE_SHA` records which engine commit the committed bindings came from,
and `tests/test_genproto.py` fails if a local engine checkout has drifted from them.

The simulator outputs are generated for **`1.60.1.69893` only**. `builds/1.15.9.69722` is
frozen: its full `normalize` no longer completes, so a simulator output for it could not
be regenerated or trusted. Nothing in `pipeline/` branches on the build string, so an
older-schema build works if one is ever fetched again.

## New build checklist
1. Find the Forever product key on https://wago.tools/builds (it appears when the beta client is on the CDN).
2. Run the workflow_dispatch in GitHub Actions with that product, or run fetch, normalize,
   icons and tree-art locally (re-run normalize once more after tree-art, so the manifest
   lists `trees/`).
3. If a table 404s or a column is missing, fix `TABLES` in `pipeline/wago.py` or the
   normalizer, add a fixture row, keep the golden tests green.
4. If `normalize` logs "items not emitted" it also exits non-zero, which stops the
   workflow before it can commit a build with no `items/`. An item stat modifier id is
   missing from `STAT_BY_MODIFIER_ID` in `pipeline/normalize/gear.py`. Add it with the
   right planner stat key, or `None` if the planner does not track that stat, and run
   `normalize` again. Everything except `items/` is still written by the failed run.
5. Run `simdb`, `simconst` and `gametables` for the new build, then commit `simdb.bin`,
   `simconsumes.json`, `spellconst/` and `gametables/`. If `simdb` raises on an
   unclassified aura, an unknown skill line or an unknown stat modifier id, classify it in
   `pipeline/simdb/equip.py` or `pipeline/normalize/gear.py` from the spells that use it
   and rerun -- do not widen a filter to make it pass. If `gametables` raises on an empty
   download, the new client does not ship that file: move it into
   `ABSENT_FROM_THE_CLASSIC_LINEAGE` and tell the engine lane, rather than dropping it
   quietly. If it raises a 400, the build string is wrong.
6. Re-check the APL ranks: `uv run pytest tests/test_apl.py -q --no-cov`. Forever may
   renumber spell ranks, and a rotation naming a rank the client does not have is a
   spell the engine cannot resolve.
7. Curate the Forever facts under `curated/` with their sources.
8. Commit `builds/<build>/` and bump `web/src/data/active-build.json`.

## Known gaps
- The Classic Era client has no `JournalInstance` table, so `dungeons.json` is empty for
  `wow_classic_era` builds. `JournalInstance` is one of several entries in
  `OPTIONAL_TABLES` in `pipeline/wago.py` (the curve tables and `ItemXItemEffect` are
  the others); any table not listed there fails the fetch on purpose.
- Classic items and enchants grant some of their stats through a spell rather than a
  column. `pipeline/simdb/equip.py` resolves those for `simdb.bin`, so the sim sees them;
  `items/<class-slug>.json`, which the planner reads, still shows only what the client
  states in a column or on the stat curve. On build 1.60.1.69893 the gap is 18 items --
  Forever states nearly everything on the curve -- but 1,329 of the 2,216 enchant rows.
- `SpellScaling` does not exist on the Classic lineage, so `spellconst` carries the
  coefficient columns exactly as the tables state them, zeros included. Vanilla
  coefficients are a convention the engine owns, not data.
- Per-race base Strength, Agility, Stamina, Intellect and Spirit are in neither DB2 nor
  the client's `GameTables` for build 1.60.1.69893, so nothing here emits them.
  `gametables/basemp.txt` covers base mana and `gametables/combatratings.txt` the rating
  conversions; base primary stats stay with the engine's own
  `tools/base_stats_parser.py`. Six of the seven files that parser names -- including
  `octbasempbyclass.txt` and both crit-chance pairs -- are not in this client at all: see
  `ABSENT_FROM_THE_CLASSIC_LINEAGE` in `pipeline/gametables.py`.
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
- **Closed:** the 1.60 client's `ItemSparse` has no `Resistances_*` or
  `StatModifier_bonusAmount_*` columns at all: armour and stat amounts are computed
  client-side from curve tables instead. `TABLES` now carries all five —
  `ItemArmorTotal`, `ItemArmorQuality`, `ItemArmorShield`, `ArmorLocation`,
  `RandPropPoints` — as `OPTIONAL_TABLES` (a future product that truly lacks them
  degrades to the pre-curve behaviour below rather than aborting the fetch).
  `pipeline/normalize/item_curves.py` resolves the same numbers the client itself
  would show:
  - `armour = ItemArmorTotal[ilvl][type] × ItemArmorQuality[ilvl][quality] ×
    ArmorLocation[slot][type]`, or `ItemArmorShield[ilvl][quality]` for a shield (no
    location term). A robe (`InventoryType` 20) has no row of its own in
    `ArmorLocation`; the client scores it the same as a chest (5).
  - `stat amount = RandPropPoints[ilvl][quality-column][slot group] ×
    StatPercentEditor_<n> / 10000`, where quality 2/3/4 map to RandPropPoints'
    `Good`/`Superior`/`Epic` columns and quality 5 (legendary, which has no column of
    its own) reuses `Epic`. `slot group` is `InventoryType` mapped to one of
    RandPropPoints' 5 budget buckets (`STAT_BUDGET_GROUP_BY_INVENTORY_TYPE`) — major
    armour pieces and 2H weapons in group 0, mid armour pieces and trinkets in 1,
    jewelry/cloak/held-offhand in 2, 1H weapons in 3, ranged in 4.
  - Both formulas were verified against build 1.60.1.69893's own tables and 2,830
    items whose id, name, item level and quality are unchanged from
    `builds/1.15.9.69722`: armour matched exactly on 2,197 of 2,200 real armour
    pieces (the 3 misses are synthetic QA-named test rows, e.g. "90 Green Warrior
    Gauntlets", not player gear), and stat amounts matched exactly on 1,357 of 1,364
    items whose stat *keys* are also unchanged (99.5%; some 1.60 items were
    re-itemized with different stats than their Era counterpart, which is expected
    and outside what this cross-check measures).
  - `_row_has_literal_amounts` (checking for `Resistances_0`) decides per row
    whether a build states amounts literally (Era) or needs the curve (1.60); an Era
    row's output is unaffected by curves being available.
  - Consequence: build 1.60.1.69893's `items/` went from 3,616 rows (all weapons,
    349–682 per class) to 23,669 rows (1,740–3,710 per class), 4,729 unique item
    ids, of which 15,297 (class-duplicated) rows carry nonzero armour. See the
    item-curves work's own report for the full before/after and a five-item
    hand-check.
  - What is still not curve-resolved: weapon damage, which `pipeline/simdb/weapons.py`
    now resolves for `simdb.bin` from the `ItemDamage*` curves but which
    `items/<class-slug>.json` still does not carry, so `_has_gear_value`'s weapon
    exemption is still needed and unchanged, and any stat whose
    `StatModifier_bonusStat_*` id is not in `STAT_BY_MODIFIER_ID` (unchanged behaviour:
    `ItemDataError`, not a guess).

## What the beta client changed about the talents

`diffs/wowhead-2026-09-14__1.60.1.69893.json` compares the pre-beta Wowhead
snapshot the planner was built on against the trees read out of the beta
client. Across all 27 trees: **1 talent added** (Hunter, Marksmanship:
Improved Serpent Sting), **2 removed** (Druid Balance: Balance of Nature;
Warrior Protection: Vitality), **2 renamed in place** (Rogue Combat: Restless
Blades is Flawless Execution; Warlock Affliction: Drain Hope is Wrack), **4
moved** (Shaman Restoration swaps Tidal Mastery and Totemic Focus; Warrior
Protection moves Bastion and Focused Rage), **0 rank caps changed**, and **1
prerequisite dropped** (Druid Balance: Nature's Majesty no longer requires
Nature's Splendor, which broke the pair's two-way link).

The snapshot stays committed as the record of what was believed before the
beta. Nothing in the pipeline reads it to build the active site's trees any
more; `python -m pipeline forever-talents` still regenerates the frozen
`forever-prebeta` build, which exists only so links shared against it keep
opening.

Regenerate the diff file with:

```bash
uv run python -m pipeline wowhead-diff --build 1.60.1.69893
```

## Re-emitting the Era build: four gaps this task found and fixed

Regenerating `builds/1.15.9.69722` for Task 6 (the first time this worktree
had ever fetched Era's raw tables, since they are git-ignored) surfaced four
pre-existing gaps unrelated to talent trees, all now fixed:

- `pipeline/normalize/item_curves.py`'s `load_item_curves` read only
  `RandPropPoints`' `GoodF_0`-style float columns; Era's own `RandPropPoints`
  has no `F` columns at all (only the integer `Good_0` style), so normalizing
  Era crashed with `KeyError: 'GoodF_0'`. Now it prefers the float column
  where the row has it and falls back to the integer column where it does
  not — both hold the same values on build 1.60.1.69893 where both are
  present.
- `pipeline/spelltext.py`'s `_base_points` preferred `EffectBasePointsF`
  whenever the column merely *existed*, which is true for both builds'
  `SpellEffect` tables — but only build 1.60.1.69893 ever populates it; Era's
  copy is uniformly `"0"` padding, so every Era description's numbers were
  silently zeroed (e.g. "Reduces the cost of your Heroic Strike ability by 1
  rage point" became "by 0 rage point"). Confirmed against
  every row of both builds' `SpellEffect` tables: Era's `EffectBasePoints`
  (int) is always populated and its `EffectBasePointsF` always `"0"`; build
  1.60.1.69893's `EffectBasePoints` is always empty and `EffectBasePointsF`
  always carries the real value. The function now prefers the integer column
  when present and falls back to the float column only when it is not.
- `pipeline/normalize/gear.py`'s level-60 armour sanity cap (2,000) rejected
  item 13375, "Crest of Retribution" — a real rare shield (`RequiredLevel`
  55) with 2,057 armour that was in the originally-committed Era build but
  had never been re-checked against the cap since it was added. Raised to
  2,100.
- `pipeline/forever.py`'s manifest digests for `forever-prebeta` were
  computed before `_repoint_to_placeholder` rewrote some of the very files
  being digested, so the manifest recorded stale digests for those files.
  The digest loop is now a shared `_file_digests` helper, called once after
  the repoint step.

These four are unrelated to each other and to the reader switch; each is
documented in its own right in the fix commit that made it and in
`.superpowers/sdd/2026-09-17-forever-trees/task-6-report.md`'s "Deviation 2".

`data/curated/races.json` and `combos.json` now name the beta client's two
real Skyborne rows (`high-order-skyborne` id 95, `windshaper-skyborne` id
96) rather than the retired placeholder; Classic Era's own `ChrRaces` has no
Skyborne row at all, so merging the curated files onto Era's races raises
`CuratedError` by design (see `tests/test_curated.py`'s `real_merged`
docstring — Era's committed `classes.json`/`races.json`/`combos.json` are a
frozen historical artifact now that Skyborne is real client data rather than
a placeholder, and are validated against the beta build instead). Task 6
does not touch curated data, so re-emitting Era for this task only
regenerated `talents.json`, `talents/*.json` and `manifest.json` — the three
outputs Task 6's reader switch actually changes — leaving
`classes.json`/`races.json`/`combos.json`/`items.json`/`items/`/`sets.json`/
`zones.json`/`dungeons.json`/`spells.json` untouched on disk. `manifest.json`
also picked up the correct hash for `classes.json` and `races.json`, which
the previously-committed manifest had recorded incorrectly (stale from
before those files were last regenerated) — `pipeline.manifest.verify` now
returns `[]` for the Era build, which it would not have before this task.
