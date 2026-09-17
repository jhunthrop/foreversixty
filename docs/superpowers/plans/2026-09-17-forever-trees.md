# Forever Trees Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** The planner draws Forever's real talent trees, read from the 1.60 beta client's trait tables — the game's tab order, grid positions, prerequisite links, rank caps and per-rank text — over the game's own background art, processed into this site's palette.

**Architecture:** The pipeline gains a trait reader used when a build's tables name class trees (`SkillLineXTraitTree`), and keeps the legacy `Talent` reader for Classic Era. Both emit the same per-class JSON, with two fields added. A second pipeline step fetches each tab's four background quadrants from CASC, composites them to one 320x384 image, processes it through one named constant and writes one WebP per tree. The planner draws prerequisite connectors as an SVG overlay whose geometry is a pure, unit-tested function, and the active build switches to the beta build only once the e2e passes on it.

**Tech Stack:** Python 3.12 + uv + pytest + pydantic + Pillow (`data/`), Go 1.25 (`api/`), Svelte 5 islands + Astro + Tailwind + vitest + Playwright (`web/`).

**Spec:** `docs/superpowers/specs/2026-09-17-forever-trees-design.md`

## Global Constraints

Copied from the spec's "Global constraints", verbatim:

- Both clients stay supported by the pipeline: Classic Era (legacy `Talent`/`TalentTab`) and the 1.60 client (trait tables). The emitted per-class talent JSON keeps one schema, with fields added, never renamed; the Era build re-emits byte-identical except for added fields.
- The Wowhead snapshot stays committed as the record of what was believed before the beta; the pipeline stops reading it once the trait tables emit the same trees, and the diff between the two (names, ranks, positions, prerequisites) is written to `diffs/` and summarised in the report.
- Spell ids in the emitted trees are the client's. A `tree_version` moves with the data so shared builds made against the old trees keep validating against their own version; the API's talent data (`api/internal/trees`) is regenerated from the same emitted files.
- The site's active build switches to the beta build only when the planner e2e passes on it and the API validates a shared build against it; the switch is its own task at the end.
- Art: the game's own background textures, fetched from the client by file id the way icons are, composited per tree and processed in the pipeline (desaturated, darkened, tinted toward the site's palette; the exact treatment is a constant in one place) so the site ships one processed image per tree, never a raw texture. Frame, borders, counters and cell states are CSS in the game's proportions, not textures.
- Phone first: the tree art and links render at 390px with no horizontal scroll inside a tree; the links stay legible when a tree is the width of a phone.
- No third-party requests at runtime; everything the planner draws ships with the site.

And from this repo's own working rules:

- Work on the `forever-trees` worktree branch, never on `main`.
- `git commit -F <file>` in its own Bash call, never chained with `git add`.
- Node comes from `~/.nvm/versions/node/v22.12.0/bin`: `export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH"` before any `npm`/`npx`.
- Python comes from uv: `export PATH="$HOME/.local/bin:$PATH"` and run everything as `uv run …` from `data/`.
- Playwright runs on port 4325 in this worktree: `E2E_PORT=4325 npx playwright test …`.
- Before any `npm run build`/`check`/`test` in `web/`: `FOREVER_DATA=fixture npm run sync:report-fixture` and `npm run sync:data` are run by the `pre*` scripts already; do not bypass them.
- Run only the tests for the files you touched (`uv run pytest tests/test_x.py`, `npx vitest run src/lib/planner/x.test.ts`, `go test ./internal/x/`); CI runs the rest.
- Type check `web/` with `NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"` and read the count.
- After `npm run build` in `web/`, check `ls dist/planner-island.js` exists: the island build runs as `postbuild` and a Svelte compile error there fails silently otherwise.

---

## Context: what was measured, and how

Everything below was measured on build `1.60.1.69893` before this plan was written. The raw tables are gitignored (`data/builds/*/raw/`), so reproduce them with:

```bash
cd data
export PATH="$HOME/.local/bin:$PATH"
uv sync
uv run python -m pipeline fetch --product wow_classic_beta --build 1.60.1.69893
cd builds/1.60.1.69893/raw
for T in TraitTree TraitNode TraitNodeEntry TraitNodeXTraitNodeEntry TraitDefinition \
         TraitEdge TraitCond TraitNodeGroup TraitNodeGroupXTraitNode TraitCurrency \
         SkillLineXTraitTree SkillLine TraitDefinitionEffectPoints CurvePoint; do
  curl -sS -H "User-Agent: foreversixty-pipeline/0.1 (+https://foreversixty.gg)" \
    -o "$T.csv" "https://wago.tools/db2/$T/csv?build=1.60.1.69893"
done
```

### 1. Which trees are class trees, and how each splits into three tabs

`SkillLineXTraitTree` has 9 rows, one per class. Each names a `SkillLine` whose `DisplayName_lang` and `SpellIconFileID` equal a `TalentTab` row's `Name_lang` and `SpellIconID`; that tab's `ClassMask` is `1 << (ChrClasses.ID - 1)`.

| tree | skill line | class |
|---|---|---|
| 1082 | 373 Enhancement | 7 Shaman |
| 1089 | 574 Balance | 11 Druid |
| 1091 | 50 Beast Mastery | 3 Hunter |
| 1100 | 184 Retribution | 2 Paladin |
| 1111 | 38 Combat | 4 Rogue |
| 1112 | 237 Arcane | 8 Mage |
| 1114 | 613 Discipline | 5 Priest |
| 1116 | 354 Demonology | 9 Warlock |
| 1117 | 26 Arms | 1 Warrior |

The other 8 `TraitTree` rows (1058, 1066, 1081, 1083, 1118, 1187, 1188, 1189) are not class trees and carry 1-36 nodes each.

The tab split is **both** `PosX` bands and `TraitNodeGroup`, and they agree. Every class tree has exactly three `TraitNodeGroup` rows whose node sets partition the tree (15-19 nodes each, spanning all seven rows); each group's nodes sit in one `PosX` band (1020-2820, 5020-6820, 9080-10880), and the band index equals the tab's `TalentTab.OrderIndex`. Only one node contradicts its band (Priest 105865, `PosX` 9280) and it is a stale duplicate — see (2). Checked against `data/raw-forever/wowhead-talents-2026-09-14.json`, which keys talents by `TalentTab` id: the split reproduces the snapshot's tree assignment for all 469 talents, with 4 cell differences and 0 rank differences.

Tab order and background name per class, left to right (talent counts after the two stale nodes in (2) are dropped):

| class | tab 0 | tab 1 | tab 2 |
|---|---|---|---|
| Warrior | 161 Arms `WarriorArms` 17 | 164 Fury `WarriorFury` 18 | 163 Protection `WarriorProtection` 18 |
| Paladin | 382 Holy `PaladinHoly` 18 | 383 Protection `PaladinProtection` 16 | 381 Retribution `PaladinCombat` 18 |
| Hunter | 361 Beast Mastery `HunterBeastMastery` 16 | 363 Marksmanship `HunterMarksmanship` 17 | 362 Survival `HunterSurvival` 18 |
| Rogue | 182 Assassination `RogueAssassination` 17 | 181 Combat `RogueCombat` 17 | 183 Subtlety `RogueSubtlety` 19 |
| Priest | 201 Discipline `PriestDiscipline` 18 | 202 Holy `PriestHoly` 17 | 203 Shadow `PriestShadow` 18 |
| Shaman | 261 Elemental `ShamanElementalCombat` 16 | 263 Enhancement `ShamanEnhancement` 18 | 262 Restoration `ShamanRestoration` 16 |
| Mage | 81 Arcane `MageArcane` 18 | 41 Fire `MageFire` 17 | 61 Frost `MageFrost` 19 |
| Warlock | 302 Affliction `WarlockCurses` 17 | 303 Demonology `WarlockSummoning` 19 | 301 Destruction `WarlockDestruction` 16 |
| Druid | 283 Balance `DruidBalance` 16 | 281 Feral Combat `DruidFeralCombat` 19 | 282 Restoration `DruidRestoration` 16 |

Total 469 talents from 471 class-tree nodes.

**The switch between readers**: Classic Era `1.15.9.69722` serves `TraitTree` (1 row) and `TraitNode` (9 nodes, all in non-class tree 1058) but **404s on `SkillLineXTraitTree`**, `TraitEdge`, `TraitCond`, `TraitDefinitionEffectPoints` and `TraitCurrency`. A non-empty `SkillLineXTraitTree` is therefore the exact test for "this build has class trait trees".

### 2. How PosX/PosY quantise

Step 600 in both axes. Column origins, one per tab index: **1020, 5020, 9080**. Row origin **2130**, rows 0-6, columns 0-3. So `column = round((PosX - origin[tab]) / 600)` and `row = round((PosY - 2130) / 600)`.

Six coordinates are wrong in the client and need repair first:

| node | class | talent | raw | repaired |
|---|---|---|---|---|
| 104982 | Hunter | Lightning Reflexes | `PosX` 102800, `PosY` 5740 | 10280, 5740 |
| 105003 | Hunter | Improved Serpent Sting | `PosY` 39300 | 3930 |
| 105865 | Priest | Holy Specialization | `PosX` 9280, `PosY` 21300 | 9280, 2130 |
| 110875 | Paladin | Improved Seal of Fury | `PosX` 5030 | 5030 |
| 110878 | Paladin | Swift Judgement | `PosX` 5030 | 5030 |
| 105916 | Warlock | Amplify Curse | `PosY` 3320 | 3320 |
| 105921 | Warlock | Improved Life Tap | `PosY` 2120 | 2120 |

Two kinds: a stray factor of ten (repair by dividing by 10 while the value exceeds the grid's last coordinate plus 10) and a +-10 wobble (`round()` absorbs it).

Two of the 471 nodes are **stale duplicates**: they carry the same `TraitDefinition.SpellID` as another node in the same tab, and are exactly the nodes whose raw coordinates need repair, while their live twins need none.

| stale | live twin | talent |
|---|---|---|
| 104982 (`102800`,`5740`) | 110859 (`10280`,`5130`) | Hunter Lightning Reflexes 19168 |
| 105865 (`9280`,`21300`) | 110855 (`6220`,`2130`) | Priest Holy Specialization 14889 |

Neither stale node is referenced by any `TraitEdge` or `TraitCond`, and neither is in its tab's cumulative gate groups. After dropping them and quantising the other 469, **every node lands on an integer cell in 0-6 x 0-3, and no two nodes in a tab share a cell.**

### 3. TraitEdge and where the required rank lives

`TraitEdge` has 96 rows, all `VisualStyle` 1. 71 are inside class trees (25 belong to the non-class trees). **Left is the prerequisite, right is the dependent**, confirmed against the snapshot's `requires` arrays.

Two of the 71 are the reverse leg of a two-way pair — Druid `Nature's Splendor`(row 2) -> `Nature's Majesty`(row 1) and Hunter `Bestial Wrath`(row 6) -> `Intimidation`(row 4). Dropping any edge whose left node sits in a **greater** row index than its right node leaves **69 prerequisites**, at most one per talent, no cycles, and exactly the direction the snapshot records. (The snapshot itself carries the Druid cycle; the client's edge direction resolves it.) 67 of the 69 join cells in the same column and 2 join cells in the same row (Priest `Mind Flay` -> `Improved Mind Flay`, Paladin `Holy Shock` -> `Divine Precision`); none needs an elbow on this build.

**The required rank is not in `TraitCond`.** All 172 `TraitCond` rows are `CondType` 0 and carry `SpentAmountRequired` against a `TraitNodeGroup`; `GrantedRanks`, `TraitNodeEntryID`, `RequiredLevel`, `Flags`, `SpecSetID` are zero in every row. The required rank equals the prerequisite's `TraitNodeEntry.MaxRanks` in **all 69** cases (zero mismatches against the snapshot's `qty`).

Classifying all 172 `TraitCond` rows:

- **162 are the per-tier points gates**: 6 per tab x 27 tabs, `TraitCurrencyID` 3820 (the class point currency, `TraitCurrency.SourcedMax` = 51 = `MAX_POINTS`), `SpentAmountRequired` in {5,10,15,20,25,30}, each against the cumulative group of rows 0..k of one tab. No class deviates from five points per tier.
- **6 belong to the non-class trees** 1187/1188/1189 (`TraitCurrencyID` 4225).
- **2 are all-zero rows** (id 43478 on tree 1114, id 50974 on tree 1089): no group, no currency, no amount.
- **1 is an oddity**: id 50984, tree 1111, group 11576 (Assassination row 4) and `TraitNodeID` 105709 (Mutilate), `SpentAmountRequired` 1.

Filtering to `TraitCurrencyID == 3820 and TraitNodeID == 0 and SpentAmountRequired > 0` yields exactly 18 rows per class tree — three copies each of 5, 10, 15, 20, 25, 30 — for all nine classes. That is the ladder check the reader asserts.

### 4. Ranks, spell ids and descriptions

`TraitNodeEntry.MaxRanks` matches the snapshot's rank caps for all 469 talents (**0 differences**).

`TraitDefinition.SpellID` is **one spell for the whole talent, not a rank-1 spell**: only 42 of 469 have `Spell.NameSubtext_lang == "Rank 1"`, and the legacy `Talent.SpellRank_*` chains contradict it (49 rank-count mismatches among the 346 that match at all; 76 multi-rank talents have no legacy row). There are no per-rank spell ids in this client.

Per-rank values come from **`TraitDefinitionEffectPoints`** (640 rows: `TraitDefinitionID`, 0-based `EffectIndex`, `OperationType` 0 in every row, `CurveID`) resolved through **`CurvePoint`** (`Pos_0` = rank 1..5, `Pos_1` = the effect's value at that rank; every `Pos_1` is a whole number and `PosPreSquish_*` is unused). 364 of the 469 talents carry curves, and **every multi-rank talent does**.

Yes, the descriptions need `data/pipeline/spelltext.py`, with two additions: the `$m<n>`/`$M<n>` token (the minimum of effect n — 724 occurrences across the 1320 rank descriptions, unhandled today) and a per-rank effect-value override. With both, **998 of the 1320 rank descriptions render byte-identical to the snapshot's**; the rest are Wowhead's own pre-beta wording and the `${…}` arithmetic the renderer deliberately leaves alone.

`TraitDefinition.OverrideName_lang`, `OverrideIcon` and `OverrideDescription_lang` are empty/zero for all 469, so names come from `SpellName` and icons from `SpellMisc.SpellIconFileDataID` — all 469 resolve to a row in `ManifestInterfaceData`.

### 5. Names: client versus snapshot

469 client talents against 470 in the snapshot. Every difference:

| kind | class / tab | cell | client | snapshot |
|---|---|---|---|---|
| moved | Shaman Restoration | (0,2) | Totemic Focus (16173, 5) | Tidal Mastery |
| moved | Shaman Restoration | (3,0) | Tidal Mastery (16194, 5) | Totemic Focus |
| removed | Druid Balance | (2,3) | — | Balance of Nature (5 ranks) |
| added | Hunter Marksmanship | (3,3) | Improved Serpent Sting (19464, 5) | — |
| renamed | Rogue Combat | (3,1) | Flawless Execution (1310711, 1) | Restless Blades |
| renamed | Warlock Affliction | (6,1) | Wrack (1316697, 1) | Drain Hope |
| moved | Warrior Protection | (4,3) | Bastion (16538, 5) | Vitality |
| removed | Warrior Protection | (5,0) | — | Focused Rage (3 ranks) |
| moved | Warrior Protection | (5,2) | Focused Rage (29787, 3) | Bastion |

Four names exist only in the snapshot (Balance of Nature, Restless Blades, Drain Hope, Vitality) and three only in the client (Improved Serpent Sting, Flawless Execution, Wrack). Of those three, two are renames in place (Flawless Execution for Restless Blades, Wrack for Drain Hope) and one is an addition. The client-only *nodes* beyond that are the two stale duplicates in (2), not real talents. Prerequisites: 69 client against 70 snapshot; the only differences are the Wrack rename and the snapshot's own reversed Druid edge.

Spell ids: 469 talents, range 5570 to 1317257, of which 75 are in the Forever range (>= 1,000,000) and 43 are >= 1,300,000.

### 6. Background art

Each `TalentTab.BackgroundFile` has four `ManifestInterfaceData` rows under `Interface\TALENTFRAME\`, and all 27 x 4 are present. Warrior Arms, for example: `WarriorArms-TopLeft` 136983, `-TopRight` 136984, `-BottomLeft` 136981, `-BottomRight` 136982.

Quadrant sizes are uniform across tabs (checked on `WarriorArms` and `PriestShadow`): TopLeft 256x256 RGB, TopRight 64x256 RGBA, BottomLeft 256x128 RGBA, BottomRight 64x128 RGBA. **The composite is 320 x 384.**

`data/pipeline/icons.py` already does the fetch-and-convert by file id: `httpx.Client(base_url="https://wago.tools", headers={"User-Agent": USER_AGENT})`, `client.get(f"/api/casc/{file_id}")`, cached at `.icon-cache/<file id>.blp` via `_atomic_write`, then `Image.open(io.BytesIO(blp))` (Pillow reads BLP natively) and `.convert("RGBA")`. Task 7 reuses `_atomic_write`, `CACHE_DIR`, `BASE_URL` and `USER_AGENT` and adds only the composite and the treatment.

Palette tokens for the tint come from `web/src/styles/tokens.css`: `--color-bg: #07090d`, `--color-raised: #0d111a`, `--color-card-top: #131824`, `--color-gold: #e5b955`, `--color-gold-deep: #a8762a`.

### 7. The web and API side as it stands

- `web/src/lib/planner/types.ts`: `Talent {id, name, icon, max_rank, tier, column, prereq_talent_id, prereq_rank, ranks[]}`, `TalentTree {id, name, position, talents[]}`, `TalentFile {build, class_id, class_slug, trees[]}`. `MAX_POINTS` 51, `POINTS_PER_TIER` 5.
- `web/src/lib/planner/rules.ts`: `indexTalents`, `validateOrder`, `canAddPoint`, `canRemovePoint`, `messages`. `canRemovePoint` already replays `validateOrder` over the shortened order, so removing a prerequisite a dependent still needs is **already refused** (`rules.test.ts:165` covers the rank-1 case); only the multi-rank case lacks a test.
- `web/src/lib/planner/grid.ts`: `gridCells`, `gridSize`, `moveFocus` — pure, unit-tested, no geometry.
- `web/src/components/planner/TreeGrid.svelte`: CSS grid, `gap-2` (8px), cells `h-11 w-11 md:h-12 md:w-12` (44px / 48px). `TalentCell.svelte` draws the border states (`border-gold` maxed, `border-gold-deep` filled, `border-line` available, `border-line-soft` + `opacity-50` locked) and the `rank/max` pill.
- `web/src/components/planner/Planner.svelte`: renders `store.talentIndex.trees` in `position` order with a per-tree header already showing `store.split[i]`; line 290 prints `ERA_DATA_NOTICE`; the Reset button (line ~538) reads `Reset` and opens a confirm.
- `web/src/lib/planner/config.ts`: `ERA_DATA_NOTICE = 'Classic Era trees shown until the beta client exports; Forever revamped talents replace them then.'`, asserted verbatim at `tests/e2e/planner.spec.ts:16`.
- `web/src/lib/planner/load.ts`: `dataUrl(build, file)` -> `/data/<build>/<file>`.
- `web/src/data/active-build.json` is `{"build": "forever-prebeta"}`. `web/scripts/sync-data.mjs` publishes every build under `data/builds/` whose manifest lists `talents/`, driven by `SYNC_ENTRIES`. `FOREVER_DATA=fixture` (the Playwright default) publishes `web/src/fixtures/planner` instead.
- `api/internal/trees/trees.go` loads every directory under `TREE_DATA_DIR` (`data/builds`) that has `combos.json` and `talents/`; `Build.Talent(classID, talentID)` indexes by talent id; `readJSON` ignores unknown fields, so added fields do not break it. `api/internal/builds/validate.go` rejects an unknown `tree_version` with `No talent data for tree version %s`. `api/internal/spec/spec.go` `Split` maps log talent ids through `Build.Talent` against `Data.Latest()`.

---

## File structure

**New**

| file | responsibility |
|---|---|
| `data/pipeline/normalize/traits.py` | Read the trait tables into class trees: tab split, grid cells, stale-node drop, prerequisites, tier-ladder check. No pydantic, no spell text. |
| `data/pipeline/curves.py` | `TraitDefinitionEffectPoints` + `CurvePoint` -> per-rank effect values. |
| `data/pipeline/normalize/trait_trees.py` | Assemble `ClassTalents` from the reader, the curves, `SpellText` and the icon table. The trait-table twin of `talent_trees.py`. |
| `data/pipeline/art.py` | Talent-frame quadrant ids, the composite, the one processing constant, the per-build fetch. |
| `data/pipeline/wowhead_diff.py` | Diff the emitted trees against the Wowhead snapshot into `diffs/`. |
| `data/tests/fixtures/traits/*.csv` | A synthetic one-class trait build exercising every rule above. |
| `web/src/lib/planner/connectors.ts` | Connector geometry: cell centres, path strings, which connectors a tree has and whether each is met. |

**Modified**

| file | change |
|---|---|
| `data/pipeline/wago.py` | 13 trait tables in `TABLES`, 5 of them in `OPTIONAL_TABLES`. |
| `data/pipeline/spelltext.py` | `$m`/`$M` token; `describe(spell_id, overrides)`. |
| `data/pipeline/models.py` | `TalentEntry.spell_id`, `TalentTree.background` — appended, never reordered. |
| `data/pipeline/normalize/talent_trees.py` | Fill the two new fields on the legacy path. |
| `data/pipeline/normalize/__init__.py` | Pick the reader; emit backgrounds' names. |
| `data/pipeline/__main__.py` | `tree-art` and `wowhead-diff` subcommands. |
| `api/internal/trees/trees.go` | Index talents by spell id; expose `TalentBySpellID`. |
| `api/internal/spec/spec.go` | Fall back to the spell-id index in `Split`. |
| `web/src/lib/planner/types.ts` | `Talent.spell_id`, `TalentTree.background`. |
| `web/src/components/planner/TreeGrid.svelte` | The SVG connector overlay and the background image. |
| `web/src/components/planner/TalentCell.svelte` | The four cell states, tightened. |
| `web/src/components/planner/OrderStrip.svelte` | The same states on the strip. |
| `web/src/components/planner/Planner.svelte` | Banner wording, per-tree counter, remaining points, `Reset…`. |
| `web/src/lib/planner/config.ts` | `ERA_DATA_NOTICE` -> `TREE_SOURCE_NOTICE`. |
| `web/scripts/sync-data.mjs` | Publish `trees/`. |
| `web/src/fixtures/planner/**` | The two new fields plus a fixture background. |
| `web/src/data/active-build.json` | The switch, in the last task. |

---

### A note on how the three tab groups are identified

The three groups that partition a class tree are exactly the tree's `TraitNodeGroup` rows that are **maximal by inclusion** — no other group of that tree is a proper superset of them. Verified on all nine class trees: 36-43 groups each, exactly 3 maximal, always disjoint and always covering every node of the tree, and their majority `PosX` bands are always `{0, 1, 2}`. Every other group is a row group or a cumulative rows-0..k gate group, and so a proper subset of one of the three. This is the rule the reader uses, because it is structural: it does not assume seven rows or any particular band.

---

### Task 1: The trait tables in the fetch, and the reader's fixtures

**Files:**
- Modify: `data/pipeline/wago.py:12-54`
- Modify: `data/tests/test_wago.py`
- Create: `data/tests/fixtures/traits/SkillLine.csv`
- Create: `data/tests/fixtures/traits/SkillLineXTraitTree.csv`
- Create: `data/tests/fixtures/traits/TalentTab.csv`
- Create: `data/tests/fixtures/traits/ChrClasses.csv`
- Create: `data/tests/fixtures/traits/TraitNode.csv`
- Create: `data/tests/fixtures/traits/TraitNodeEntry.csv`
- Create: `data/tests/fixtures/traits/TraitNodeXTraitNodeEntry.csv`
- Create: `data/tests/fixtures/traits/TraitDefinition.csv`
- Create: `data/tests/fixtures/traits/TraitEdge.csv`
- Create: `data/tests/fixtures/traits/TraitNodeGroup.csv`
- Create: `data/tests/fixtures/traits/TraitNodeGroupXTraitNode.csv`
- Create: `data/tests/fixtures/traits/TraitCond.csv`
- Create: `data/tests/fixtures/traits/TraitCurrency.csv`

**Interfaces:**
- Consumes: `pipeline.wago.TABLES`, `pipeline.wago.OPTIONAL_TABLES` (existing lists).
- Produces: the 13 trait tables in `TABLES`; `SkillLineXTraitTree`, `TraitEdge`, `TraitCond`, `TraitDefinitionEffectPoints` and `TraitCurrency` also in `OPTIONAL_TABLES`. The fixture directory `data/tests/fixtures/traits/`, a synthetic one-class trait build (tree 1117 Warrior, 3 tabs, 12 nodes of which 1 is stale, 5 edges of which 1 is reversed) that Tasks 2, 3, 4 and 5 all read.

- [ ] **Step 1: Write the failing test**

Append to `data/tests/test_wago.py`:

```python
from pipeline.wago import OPTIONAL_TABLES, TABLES

#: Every trait table the 1.60 reader needs. Classic Era 1.15.9.69722 serves
#: SkillLine, TraitNode, TraitNodeEntry, TraitNodeXTraitNodeEntry,
#: TraitDefinition, TraitNodeGroup, TraitNodeGroupXTraitNode and CurvePoint,
#: and 404s on the other five, so only those five are allowed to be missing.
TRAIT_TABLES = [
    "SkillLine",
    "SkillLineXTraitTree",
    "TraitNode",
    "TraitNodeEntry",
    "TraitNodeXTraitNodeEntry",
    "TraitDefinition",
    "TraitEdge",
    "TraitCond",
    "TraitNodeGroup",
    "TraitNodeGroupXTraitNode",
    "TraitCurrency",
    "TraitDefinitionEffectPoints",
    "CurvePoint",
]
ERA_MISSING = {
    "SkillLineXTraitTree",
    "TraitEdge",
    "TraitCond",
    "TraitCurrency",
    "TraitDefinitionEffectPoints",
}


def test_every_trait_table_is_fetched_once():
    missing = [t for t in TRAIT_TABLES if t not in TABLES]
    assert missing == [], f"TABLES is missing {missing}"
    assert len(TABLES) == len(set(TABLES)), "TABLES lists a table twice"


def test_only_the_tables_classic_era_lacks_are_optional():
    optional_traits = {t for t in TRAIT_TABLES if t in OPTIONAL_TABLES}
    assert optional_traits == ERA_MISSING
```

- [ ] **Step 2: Run the tests to verify they fail**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_wago.py -q
```
Expected: FAIL — `TABLES is missing ['SkillLine', 'SkillLineXTraitTree', …]`.

- [ ] **Step 3: Add the tables**

In `data/pipeline/wago.py`, after `"TalentTab",` in `TABLES`, insert:

```python
    # The 1.60 client's talent trees live in the modern trait tables rather than
    # the legacy Talent rows above, which it still ships for its own UI chrome.
    # pipeline/normalize/traits.py reads these; Classic Era has no class trait
    # trees and 404s on five of them (see OPTIONAL_TABLES).
    "SkillLine",
    "SkillLineXTraitTree",
    "TraitNode",
    "TraitNodeEntry",
    "TraitNodeXTraitNodeEntry",
    "TraitDefinition",
    "TraitEdge",
    "TraitCond",
    "TraitNodeGroup",
    "TraitNodeGroupXTraitNode",
    "TraitCurrency",
    "TraitDefinitionEffectPoints",
    "CurvePoint",
```

And extend `OPTIONAL_TABLES`, replacing its closing brace region so the set reads:

```python
OPTIONAL_TABLES = frozenset(
    {
        "JournalInstance",
        "ItemArmorTotal",
        "ItemArmorQuality",
        "ItemArmorShield",
        "ArmorLocation",
        "RandPropPoints",
        # Classic Era 1.15.9.69722 has no class trait trees: it serves the other
        # eight trait tables (all of whose rows belong to non-class trees) and
        # 404s on these five. A build without them falls back to the legacy
        # Talent reader, which is exactly what Era wants.
        "SkillLineXTraitTree",
        "TraitEdge",
        "TraitCond",
        "TraitCurrency",
        "TraitDefinitionEffectPoints",
    }
)
```

- [ ] **Step 4: Run the tests to verify they pass**

```bash
cd data && uv run pytest tests/test_wago.py -q
```
Expected: PASS.

- [ ] **Step 5: Write the fixture tables**

`data/tests/fixtures/traits/SkillLine.csv`:

```csv
ID,DisplayName_lang,SpellIconFileID
26,Arms,132292
38,Combat,132090
```

`data/tests/fixtures/traits/SkillLineXTraitTree.csv`:

```csv
ID,SkillLineID,TraitTreeID,Variant
1,26,1117,0
```

`data/tests/fixtures/traits/TalentTab.csv`:

```csv
ID,Name_lang,BackgroundFile,OrderIndex,ClassMask,SpellIconID
161,Arms,WarriorArms,0,1,132292
164,Fury,WarriorFury,1,1,132347
163,Protection,WarriorProtection,2,1,134952
181,Combat,RogueCombat,1,8,132090
```

`data/tests/fixtures/traits/ChrClasses.csv`:

```csv
ID,Name_lang,Filename
1,Warrior,WARRIOR
4,Rogue,ROGUE
```

`data/tests/fixtures/traits/TraitNode.csv` — tab 0 origin 1020, tab 1 origin 5020, tab 2 origin 9080; 900004 and 900013 carry the client's +-10 wobble, 900005 a stray factor of ten, and 900006 is the stale duplicate of 900002:

```csv
ID,TraitTreeID,PosX,PosY,Type,Flags,TraitSubTreeID
900001,1117,1020,2130,0,0,0
900002,1117,1620,2130,0,0,0
900003,1117,1620,2730,0,0,0
900004,1117,2220,2720,0,0,0
900005,1117,1620,33300,0,0,0
900006,1117,16200,2130,0,0,0
900011,1117,5620,2130,0,0,0
900012,1117,6220,2130,0,0,0
900013,1117,5030,2730,0,0,0
900021,1117,9080,2130,0,0,0
900022,1117,9080,3330,0,0,0
900023,1117,9680,2730,0,0,0
```

`data/tests/fixtures/traits/TraitNodeEntry.csv`:

```csv
ID,TraitDefinitionID,MaxRanks,NodeEntryType,TraitSubTreeID
800001,700001,3,0,0
800002,700002,5,0,0
800003,700003,5,0,0
800004,700004,3,0,0
800005,700005,2,0,0
800006,700006,5,0,0
800011,700011,5,0,0
800012,700012,5,0,0
800013,700013,1,0,0
800021,700021,2,0,0
800022,700022,1,0,0
800023,700023,2,0,0
```

`data/tests/fixtures/traits/TraitNodeXTraitNodeEntry.csv`:

```csv
ID,TraitNodeID,TraitNodeEntryID,_Index
1,900001,800001,0
2,900002,800002,0
3,900003,800003,0
4,900004,800004,0
5,900005,800005,0
6,900006,800006,0
7,900011,800011,0
8,900012,800012,0
9,900013,800013,0
10,900021,800021,0
11,900022,800022,0
12,900023,800023,0
```

`data/tests/fixtures/traits/TraitDefinition.csv` — 700006 repeats 700002's spell id, which is what makes 900006 a duplicate:

```csv
ID,SpellID,OverrideName_lang,OverrideIcon
700001,12282,,0
700002,16462,,0
700003,12295,,0
700004,12296,,0
700005,12163,,0
700006,16462,,0
700011,12321,,0
700012,12322,,0
700013,12323,,0
700021,12797,,0
700022,12809,,0
700023,12287,,0
```

`data/tests/fixtures/traits/TraitEdge.csv` — 1 vertical, 1 horizontal, 1 vertical over two rows, 1 elbow, and 1 reversed leg that must be dropped:

```csv
ID,VisualStyle,LeftTraitNodeID,RightTraitNodeID,Type
1,1,900002,900003,2
2,1,900011,900012,2
3,1,900021,900022,2
4,1,900021,900023,2
5,1,900023,900021,2
```

`data/tests/fixtures/traits/TraitNodeGroup.csv` — 600003 is a row group, a proper subset of 600000, so the maximal-by-inclusion rule has something to reject:

```csv
ID,TraitTreeID,Flags
600000,1117,0
600001,1117,0
600002,1117,0
600003,1117,0
```

`data/tests/fixtures/traits/TraitNodeGroupXTraitNode.csv`:

```csv
ID,TraitNodeGroupID,TraitNodeID,_Index
1,600000,900001,0
2,600000,900002,0
3,600000,900003,0
4,600000,900004,0
5,600000,900005,0
6,600000,900006,0
7,600001,900011,0
8,600001,900012,0
9,600001,900013,0
10,600002,900021,0
11,600002,900022,0
12,600002,900023,0
13,600003,900001,0
14,600003,900002,0
```

`data/tests/fixtures/traits/TraitCond.csv` — the ladder for three tabs, plus the two shapes the real build carries that must be filtered out (an all-zero row and one pinned to a node):

```csv
ID,CondType,TraitTreeID,TraitNodeGroupID,TraitNodeID,TraitCurrencyID,SpentAmountRequired
500001,0,1117,600000,0,3820,5
500002,0,1117,600000,0,3820,10
500003,0,1117,600000,0,3820,15
500004,0,1117,600000,0,3820,20
500005,0,1117,600000,0,3820,25
500006,0,1117,600000,0,3820,30
500011,0,1117,600001,0,3820,5
500012,0,1117,600001,0,3820,10
500013,0,1117,600001,0,3820,15
500014,0,1117,600001,0,3820,20
500015,0,1117,600001,0,3820,25
500016,0,1117,600001,0,3820,30
500021,0,1117,600002,0,3820,5
500022,0,1117,600002,0,3820,10
500023,0,1117,600002,0,3820,15
500024,0,1117,600002,0,3820,20
500025,0,1117,600002,0,3820,25
500026,0,1117,600002,0,3820,30
500031,0,1117,0,0,0,0
500032,0,1117,600003,900001,3820,1
```

`data/tests/fixtures/traits/TraitCurrency.csv`:

```csv
ID,Type,SourcedMax
3820,2,51
```

- [ ] **Step 6: Commit**

```bash
git add data/pipeline/wago.py data/tests/test_wago.py data/tests/fixtures/traits
```

```bash
git commit -F - <<'MSG'
feat(data): fetch the 1.60 client's trait tables

The 1.60 beta client keeps its talent trees in the modern trait tables, not
the legacy Talent rows. Add all thirteen to the fetch, and allow the five
Classic Era 404s on so an Era fetch still completes.

Also adds tests/fixtures/traits/: a synthetic one-class trait build that the
reader's tests are written against, with the two coordinate typos, the stale
duplicate node, the reversed edge and the tier ladder the real build carries.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 2: The trait reader — class trees, the tab split and the grid

**Files:**
- Create: `data/pipeline/normalize/traits.py`
- Create: `data/tests/traits_fixture.py`
- Create: `data/tests/test_traits.py`
- Modify: `data/pipeline/normalize/talents.py:6-18` (make `_class_id_from_mask` public so both readers share it)

**Interfaces:**
- Consumes: `pipeline.csvio.read_csv`; the fixture tables from Task 1.
- Produces:
  - `pipeline.normalize.talents.class_id_from_mask(tab_id: int, mask: int) -> int` (renamed from `_class_id_from_mask`; same body, same `TalentDataError`).
  - `pipeline.normalize.traits.TraitDataError(ValueError)`
  - `pipeline.normalize.traits.TraitRows` — a frozen dataclass of the thirteen row lists, fields `skill_line`, `skill_line_x_trait_tree`, `talent_tab`, `chr_classes`, `node`, `node_entry`, `node_x_entry`, `definition`, `edge`, `cond`, `currency`, `node_group`, `node_group_x_node`, each `list[dict[str, str]]`.
  - `pipeline.normalize.traits.TraitTalent` — frozen: `node_id: int`, `definition_id: int`, `spell_id: int`, `max_rank: int`, `row: int`, `column: int`, `prereq_node_id: int | None = None`, `prereq_rank: int | None = None`.
  - `pipeline.normalize.traits.TraitTab` — frozen: `tab_id: int`, `name: str`, `position: int`, `background: str`, `talents: tuple[TraitTalent, ...]`.
  - `pipeline.normalize.traits.TraitClassTree` — frozen: `class_id: int`, `tree_id: int`, `tabs: tuple[TraitTab, ...]`.
  - `pipeline.normalize.traits.has_trait_trees(skill_line_x_rows: list[dict[str, str]]) -> bool`
  - `pipeline.normalize.traits.grid_cell(pos_x: int, pos_y: int, tab_index: int) -> tuple[int, int]`
  - `pipeline.normalize.traits.read_trait_trees(rows: TraitRows) -> list[TraitClassTree]` (sorted by `class_id`; prerequisites stay `None` until Task 3).
  - Module constants `CLASS_CURRENCY_ID = 3820`, `CLASS_POINT_BUDGET = 51`, `GRID_COLUMNS = 4`, `GRID_ROWS = 7`, `CELL_STEP = 600`, `COLUMN_ORIGINS = (1020, 5020, 9080)`, `ROW_ORIGIN = 2130`, `POSITION_TOLERANCE = 10`, `TIER_GATES = (5, 10, 15, 20, 25, 30)`.

- [ ] **Step 1: Write the failing test**

```python
# data/tests/traits_fixture.py
"""Loading the synthetic trait build. A sibling module, not a test module, so
test_traits.py and test_normalize_trait_trees.py read the same fixture the same
way. pytest puts data/tests/ on sys.path for test files with no __init__.py, so
both import it as a top-level module."""

from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.normalize.traits import TraitRows

TRAITS = Path(__file__).parent / "fixtures/traits"

#: TraitRows field -> the CSV that fills it.
TABLES = {
    "skill_line": "SkillLine",
    "skill_line_x_trait_tree": "SkillLineXTraitTree",
    "talent_tab": "TalentTab",
    "chr_classes": "ChrClasses",
    "node": "TraitNode",
    "node_entry": "TraitNodeEntry",
    "node_x_entry": "TraitNodeXTraitNodeEntry",
    "definition": "TraitDefinition",
    "edge": "TraitEdge",
    "cond": "TraitCond",
    "currency": "TraitCurrency",
    "node_group": "TraitNodeGroup",
    "node_group_x_node": "TraitNodeGroupXTraitNode",
}


def trait_rows(**replace) -> TraitRows:
    """The fixture build, with any one table swapped out by keyword."""
    loaded = {field: read_csv(TRAITS / f"{name}.csv") for field, name in TABLES.items()}
    return TraitRows(**{**loaded, **replace})
```

```python
# data/tests/test_traits.py
import pytest

from pipeline.csvio import read_csv
from pipeline.normalize.traits import (
    COLUMN_ORIGINS,
    TraitDataError,
    grid_cell,
    has_trait_trees,
    read_trait_trees,
)
from traits_fixture import TRAITS, trait_rows


def test_a_build_without_the_class_link_table_has_no_trait_trees():
    assert has_trait_trees(read_csv(TRAITS / "SkillLineXTraitTree.csv")) is True
    assert has_trait_trees([]) is False


@pytest.mark.parametrize(
    ("pos_x", "pos_y", "tab", "cell"),
    [
        (1020, 2130, 0, (0, 0)),
        (2820, 5730, 0, (6, 3)),
        (6820, 3330, 1, (2, 3)),
        (10880, 5730, 2, (6, 3)),
        # The client's own +-10 wobble: Paladin PosX 5030, Warlock PosY 2120.
        (5030, 2730, 1, (1, 0)),
        (1020, 2120, 0, (0, 0)),
        # A stray factor of ten: Hunter PosX 102800, Priest PosY 21300.
        (102800, 5730, 2, (6, 2)),
        (9080, 21300, 2, (0, 0)),
    ],
)
def test_grid_cell_quantises_and_repairs_the_clients_typos(pos_x, pos_y, tab, cell):
    assert grid_cell(pos_x, pos_y, tab) == cell


def test_grid_cell_refuses_a_position_that_is_not_on_its_tabs_grid():
    # Priest node 105865 sits at PosX 9280 while its tab's columns are 5020..6820.
    with pytest.raises(TraitDataError, match="9280"):
        grid_cell(9280, 2130, 1)


def test_the_tree_is_split_into_the_three_tabs_in_the_games_order():
    (warrior,) = read_trait_trees(trait_rows())
    assert warrior.class_id == 1
    assert warrior.tree_id == 1117
    assert [(t.tab_id, t.name, t.position, t.background) for t in warrior.tabs] == [
        (161, "Arms", 0, "WarriorArms"),
        (164, "Fury", 1, "WarriorFury"),
        (163, "Protection", 2, "WarriorProtection"),
    ]


def test_every_talent_lands_on_its_own_cell_with_its_spell_and_rank_cap():
    (warrior,) = read_trait_trees(trait_rows())
    arms, fury, protection = warrior.tabs
    assert [(t.node_id, t.spell_id, t.max_rank, t.row, t.column) for t in arms.talents] == [
        (900001, 12282, 3, 0, 0),
        (900002, 16462, 5, 0, 1),
        (900003, 12295, 5, 1, 1),
        (900004, 12296, 3, 1, 2),
        (900005, 12163, 2, 2, 1),
    ]
    assert [(t.node_id, t.row, t.column) for t in fury.talents] == [
        (900011, 0, 1),
        (900012, 0, 2),
        (900013, 1, 0),
    ]
    assert [(t.node_id, t.row, t.column) for t in protection.talents] == [
        (900021, 0, 0),
        (900022, 2, 0),
        (900023, 1, 1),
    ]


def test_the_stale_twin_of_a_talent_is_dropped_and_the_live_one_kept():
    (warrior,) = read_trait_trees(trait_rows())
    arms = warrior.tabs[0]
    deflection = [t for t in arms.talents if t.spell_id == 16462]
    assert [t.node_id for t in deflection] == [900002]


def test_two_twins_that_both_need_repair_are_refused_rather_than_guessed():
    nodes = read_csv(TRAITS / "TraitNode.csv")
    for node in nodes:
        if node["ID"] == "900002":
            node["PosX"] = "16200"
    with pytest.raises(TraitDataError, match="16462"):
        read_trait_trees(trait_rows(node=nodes))


def test_two_talents_in_one_cell_are_refused():
    nodes = read_csv(TRAITS / "TraitNode.csv")
    for node in nodes:
        if node["ID"] == "900004":
            node["PosX"] = str(COLUMN_ORIGINS[0] + 600)  # onto 900003's cell
    with pytest.raises(TraitDataError, match="row 1 column 1"):
        read_trait_trees(trait_rows(node=nodes))


def test_a_tree_whose_groups_do_not_make_three_tabs_is_refused():
    groups = [g for g in read_csv(TRAITS / "TraitNodeGroup.csv") if g["ID"] != "600002"]
    members = [
        m for m in read_csv(TRAITS / "TraitNodeGroupXTraitNode.csv") if m["TraitNodeGroupID"] != "600002"
    ]
    with pytest.raises(TraitDataError, match="2 tabs"):
        read_trait_trees(trait_rows(node_group=groups, node_group_x_node=members))
```

- [ ] **Step 2: Run the tests to verify they fail**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_traits.py -q --no-cov
```
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.normalize.traits'`.

- [ ] **Step 3: Make the class-mask helper public**

In `data/pipeline/normalize/talents.py`, rename `_class_id_from_mask` to `class_id_from_mask` (declaration and the one call site inside `normalize_talents`), and extend its docstring's first line to:

```python
def class_id_from_mask(tab_id: int, mask: int) -> int:
    """The one class a `ClassMask` bit names, or TalentDataError if it names none or many.

    Shared by both talent readers: the legacy one below and
    `pipeline/normalize/traits.py`, which resolves the same masks off the same
    `TalentTab` rows.
```

- [ ] **Step 4: Write the reader**

```python
# data/pipeline/normalize/traits.py
"""The 1.60 client's trait tables, read as Classic's three-tab talent trees.

The beta client keeps Forever's talents in the modern trait tables. The legacy
`Talent` table it still ships holds Classic Era's talents, so reading that
would draw the wrong trees; `pipeline/normalize/talent_trees.py` stays for
Classic Era, which has no class trait trees at all (`SkillLineXTraitTree` 404s
there). `has_trait_trees` is the switch between the two.

Every rule here was measured on build 1.60.1.69893 and is asserted rather than
assumed, because the only thing that makes any of them true is that Blizzard's
data says so today:

* a class tree is one `SkillLineXTraitTree` names, and the skill line's display
  name and icon identify the `TalentTab` whose `ClassMask` gives the class;
* a tree's three tabs are its `TraitNodeGroup` rows that are maximal by
  inclusion -- every other group of the tree (a row group, or a cumulative
  rows-0..k points gate) is a proper subset of one of them;
* a tab's left-to-right position is the `PosX` band most of its nodes sit in,
  and that equals the tab's `TalentTab.OrderIndex`;
* positions quantise onto a 4x7 grid in steps of 600, and six coordinates in
  the build are wrong -- three by a stray factor of ten and four by +-10;
* two nodes are stale twins of a live node in the same tab: same spell id, and
  the coordinate that needs the stray digit removed.
"""

from __future__ import annotations

import logging
from collections import Counter
from dataclasses import dataclass

from pipeline.normalize.talents import class_id_from_mask

logger = logging.getLogger(__name__)

#: `TraitCurrency` row for a class's talent points. Its `SourcedMax` is 51,
#: the same budget `web/src/lib/planner/types.ts` calls MAX_POINTS.
CLASS_CURRENCY_ID = 3820
CLASS_POINT_BUDGET = 51

GRID_COLUMNS = 4
GRID_ROWS = 7
#: The distance between two neighbouring cells in `PosX`/`PosY` units.
CELL_STEP = 600
#: `PosX` of column 0, one per tab, left to right. The gaps between the three
#: are not equal (4000 then 4060), so a tab's origin cannot be derived from
#: tab index times a constant.
COLUMN_ORIGINS = (1020, 5020, 9080)
ROW_ORIGIN = 2130
#: How far off a cell's exact coordinate the client is allowed to be. Six rows
#: in build 1.60.1.69893 are off by exactly 10 (e.g. Paladin `PosX` 5030 for
#: column 5020); anything further out is a different kind of wrong.
POSITION_TOLERANCE = 10
#: Points that must already sit in a tab before its rows 1..6 open.
TIER_GATES = (5, 10, 15, 20, 25, 30)

_LAST_COLUMN_POS = COLUMN_ORIGINS[-1] + (GRID_COLUMNS - 1) * CELL_STEP
_LAST_ROW_POS = ROW_ORIGIN + (GRID_ROWS - 1) * CELL_STEP


class TraitDataError(ValueError):
    """The trait tables do not have the shape three-tab talent trees need."""


@dataclass(frozen=True)
class TraitTalent:
    node_id: int
    definition_id: int
    spell_id: int
    max_rank: int
    row: int
    column: int
    prereq_node_id: int | None = None
    prereq_rank: int | None = None


@dataclass(frozen=True)
class TraitTab:
    tab_id: int
    name: str
    position: int
    background: str
    talents: tuple[TraitTalent, ...]


@dataclass(frozen=True)
class TraitClassTree:
    class_id: int
    tree_id: int
    tabs: tuple[TraitTab, ...]


@dataclass(frozen=True)
class TraitRows:
    """The thirteen tables the reader needs, straight off `read_csv`."""

    skill_line: list[dict[str, str]]
    skill_line_x_trait_tree: list[dict[str, str]]
    talent_tab: list[dict[str, str]]
    chr_classes: list[dict[str, str]]
    node: list[dict[str, str]]
    node_entry: list[dict[str, str]]
    node_x_entry: list[dict[str, str]]
    definition: list[dict[str, str]]
    edge: list[dict[str, str]]
    cond: list[dict[str, str]]
    currency: list[dict[str, str]]
    node_group: list[dict[str, str]]
    node_group_x_node: list[dict[str, str]]


def has_trait_trees(skill_line_x_rows: list[dict[str, str]]) -> bool:
    """True when this build's tables name class trait trees.

    Classic Era serves eight of the thirteen trait tables -- their rows all
    belong to non-class trees -- and 404s on `SkillLineXTraitTree`, which the
    fetch writes as an empty table. One non-empty row there is the whole test.
    """
    return bool(skill_line_x_rows)


def _repair(value: int, last: int) -> int:
    """Strip a stray factor of ten from a coordinate.

    Three coordinates in build 1.60.1.69893 carry one: `PosX` 102800 for
    10280, `PosY` 39300 for 3930, `PosY` 21300 for 2130. Dividing while the
    value is past the grid's last coordinate (plus the +-10 the client is
    allowed) removes it and leaves every honest coordinate alone -- `PosY`
    5740 stays 5740 rather than becoming 574.
    """
    while value > last + POSITION_TOLERANCE:
        value //= 10
    return value


def _quantise(value: int, origin: int, count: int, axis: str, raw: int) -> int:
    offset = value - origin
    index = round(offset / CELL_STEP)
    if not 0 <= index < count or abs(offset - index * CELL_STEP) > POSITION_TOLERANCE:
        raise TraitDataError(
            f"{axis} {raw} is not on the grid: it resolves to {value}, which is "
            f"{offset} from origin {origin} and so no multiple of {CELL_STEP} "
            f"inside 0..{count - 1}"
        )
    return index


def grid_cell(pos_x: int, pos_y: int, tab_index: int) -> tuple[int, int]:
    """The (row, column) a node's raw `PosX`/`PosY` name in its tab's grid."""
    column = _quantise(
        _repair(pos_x, _LAST_COLUMN_POS), COLUMN_ORIGINS[tab_index], GRID_COLUMNS, "PosX", pos_x
    )
    row = _quantise(_repair(pos_y, _LAST_ROW_POS), ROW_ORIGIN, GRID_ROWS, "PosY", pos_y)
    return row, column


def _class_of_tree(rows: TraitRows) -> dict[int, int]:
    """Trait tree id -> class id, through the skill line's own talent tab."""
    lines = {int(r["ID"]): r for r in rows.skill_line}
    tabs = {(r["Name_lang"], int(r["SpellIconID"])): r for r in rows.talent_tab}
    out: dict[int, int] = {}
    for link in rows.skill_line_x_trait_tree:
        line = lines.get(int(link["SkillLineID"]))
        if line is None:
            raise TraitDataError(f"skill line {link['SkillLineID']} is not in SkillLine")
        key = (line["DisplayName_lang"], int(line["SpellIconFileID"]))
        tab = tabs.get(key)
        if tab is None:
            raise TraitDataError(f"no TalentTab matches skill line {key}")
        out[int(link["TraitTreeID"])] = class_id_from_mask(int(tab["ID"]), int(tab["ClassMask"]))
    return out


#: The `PosX` halfway between the end of one tab's columns and the start of
#: the next, used only to say which band a node is in.
_BAND_MIDPOINTS = tuple(
    (COLUMN_ORIGINS[i] + (GRID_COLUMNS - 1) * CELL_STEP + COLUMN_ORIGINS[i + 1]) // 2
    for i in range(len(COLUMN_ORIGINS) - 1)
)


def _band(pos_x: int) -> int:
    """Which of the three column bands a `PosX` falls in."""
    repaired = _repair(pos_x, _LAST_COLUMN_POS)
    return sum(repaired > midpoint for midpoint in _BAND_MIDPOINTS)


def _tab_node_sets(
    tree_id: int,
    tree_nodes: set[int],
    rows: TraitRows,
    node_rows: dict[int, dict[str, str]],
) -> list[set[int]]:
    """The tree's three tab node sets, ordered left to right.

    A tab group is a `TraitNodeGroup` of this tree that no other group of the
    tree is a proper superset of. Measured on all nine class trees: exactly
    three per tree, always disjoint, always covering every node, and their
    majority `PosX` bands are always 0, 1 and 2.
    """
    members: dict[int, set[int]] = {}
    group_ids = {int(g["ID"]) for g in rows.node_group if int(g["TraitTreeID"]) == tree_id}
    for link in rows.node_group_x_node:
        group = int(link["TraitNodeGroupID"])
        if group in group_ids:
            members.setdefault(group, set()).add(int(link["TraitNodeID"]))
    maximal = [s for s in members.values() if not any(s < other for other in members.values())]
    covered: set[int] = set()
    for group in maximal:
        covered |= group
    if len(maximal) != len(COLUMN_ORIGINS) or covered != tree_nodes:
        raise TraitDataError(
            f"tree {tree_id} splits into {len(maximal)} tabs covering "
            f"{len(covered)} of its {len(tree_nodes)} nodes; a class tree has "
            f"{len(COLUMN_ORIGINS)} tabs covering all of them"
        )
    if sum(len(group) for group in maximal) != len(covered):
        raise TraitDataError(f"tree {tree_id}'s tabs overlap")
    bands = {}
    for group in maximal:
        band, _ = Counter(_band(int(node_rows[node]["PosX"])) for node in group).most_common(1)[0]
        if band in bands:
            raise TraitDataError(f"tree {tree_id} has two tabs in column band {band}")
        bands[band] = group
    return [bands[index] for index in range(len(COLUMN_ORIGINS))]


def _talents_for_tab(
    node_ids: set[int],
    tab_index: int,
    tab_name: str,
    node_rows: dict[int, dict[str, str]],
    entry_of: dict[int, dict[str, str]],
) -> tuple[TraitTalent, ...]:
    kept = _drop_stale_twins(node_ids, tab_name, node_rows, entry_of)
    talents: list[TraitTalent] = []
    occupied: dict[tuple[int, int], int] = {}
    for node_id in sorted(kept):
        node = node_rows[node_id]
        entry = entry_of[node_id]
        row, column = grid_cell(int(node["PosX"]), int(node["PosY"]), tab_index)
        if (row, column) in occupied:
            raise TraitDataError(
                f"{tab_name} has nodes {occupied[(row, column)]} and {node_id} "
                f"both at row {row} column {column}"
            )
        occupied[(row, column)] = node_id
        talents.append(
            TraitTalent(
                node_id=node_id,
                definition_id=int(entry["TraitDefinitionID"]),
                spell_id=int(entry["SpellID"]),
                max_rank=int(entry["MaxRanks"]),
                row=row,
                column=column,
            )
        )
    return tuple(talents)


def _needs_repair(node: dict[str, str]) -> bool:
    return (
        int(node["PosX"]) > _LAST_COLUMN_POS + POSITION_TOLERANCE
        or int(node["PosY"]) > _LAST_ROW_POS + POSITION_TOLERANCE
    )


def _drop_stale_twins(
    node_ids: set[int],
    tab_name: str,
    node_rows: dict[int, dict[str, str]],
    entry_of: dict[int, dict[str, str]],
) -> set[int]:
    """Keep one node per spell id in a tab.

    Build 1.60.1.69893 leaves two superseded nodes behind -- Hunter 104982 and
    Priest 105865 -- each carrying the same spell as a live node in the same
    tab and each with a coordinate that needs a stray digit removed, which the
    live node never does. Nothing else separates them: both are in the tab's
    group, neither is named by an edge or a condition, and the flags do not
    distinguish them.
    """
    by_spell: dict[int, list[int]] = {}
    for node_id in node_ids:
        by_spell.setdefault(int(entry_of[node_id]["SpellID"]), []).append(node_id)
    kept: set[int] = set()
    for spell_id, ids in by_spell.items():
        if len(ids) == 1:
            kept.add(ids[0])
            continue
        live = [node_id for node_id in ids if not _needs_repair(node_rows[node_id])]
        if len(live) != 1:
            raise TraitDataError(
                f"{tab_name} has {len(ids)} nodes for spell {spell_id} "
                f"({sorted(ids)}) and {len(live)} of them are on the grid; "
                "which one the client draws cannot be decided from the tables"
            )
        logger.info(
            "%s: dropping stale node %s, superseded by %s (spell %s)",
            tab_name,
            sorted(set(ids) - set(live)),
            live[0],
            spell_id,
        )
        kept.add(live[0])
    return kept


def read_trait_trees(rows: TraitRows) -> list[TraitClassTree]:
    """One `TraitClassTree` per class, tabs left to right, sorted by class id."""
    class_of_tree = _class_of_tree(rows)
    node_rows = {int(n["ID"]): n for n in rows.node}
    entries = {int(e["ID"]): e for e in rows.node_entry}
    definitions = {int(d["ID"]): d for d in rows.definition}
    entry_of: dict[int, dict[str, str]] = {}
    for link in rows.node_x_entry:
        node_id = int(link["TraitNodeID"])
        if node_id in entry_of:
            raise TraitDataError(f"node {node_id} has more than one entry")
        entry = entries[int(link["TraitNodeEntryID"])]
        definition = definitions[int(entry["TraitDefinitionID"])]
        entry_of[node_id] = {
            "TraitDefinitionID": entry["TraitDefinitionID"],
            "MaxRanks": entry["MaxRanks"],
            "SpellID": definition["SpellID"],
        }

    tabs_by_class: dict[int, list[dict[str, str]]] = {}
    for tab in rows.talent_tab:
        class_id = class_id_from_mask(int(tab["ID"]), int(tab["ClassMask"]))
        tabs_by_class.setdefault(class_id, []).append(tab)

    trees: list[TraitClassTree] = []
    for tree_id, class_id in sorted(class_of_tree.items()):
        tree_nodes = {n for n, node in node_rows.items() if int(node["TraitTreeID"]) == tree_id}
        missing = sorted(tree_nodes - set(entry_of))
        if missing:
            raise TraitDataError(f"tree {tree_id} nodes {missing} have no entry")
        tab_rows = sorted(tabs_by_class.get(class_id, []), key=lambda t: int(t["OrderIndex"]))
        if len(tab_rows) != len(COLUMN_ORIGINS):
            raise TraitDataError(
                f"class {class_id} has {len(tab_rows)} TalentTab rows, "
                f"not {len(COLUMN_ORIGINS)}"
            )
        node_sets = _tab_node_sets(tree_id, tree_nodes, rows, node_rows)
        tabs = tuple(
            TraitTab(
                tab_id=int(tab_row["ID"]),
                name=tab_row["Name_lang"],
                position=index,
                background=tab_row["BackgroundFile"],
                talents=_talents_for_tab(
                    node_sets[index], index, tab_row["Name_lang"], node_rows, entry_of
                ),
            )
            for index, tab_row in enumerate(tab_rows)
        )
        trees.append(TraitClassTree(class_id=class_id, tree_id=tree_id, tabs=tabs))
    return sorted(trees, key=lambda t: t.class_id)
```

- [ ] **Step 5: Run the tests to verify they pass**

```bash
cd data && uv run pytest tests/test_traits.py tests/test_normalize_talents.py -q --no-cov
```
Expected: PASS, 11 passed.

- [ ] **Step 6: Lint**

```bash
cd data && uv run ruff check pipeline/normalize/traits.py pipeline/normalize/talents.py tests/test_traits.py
```
Expected: `All checks passed!`

- [ ] **Step 7: Commit**

```bash
git add data/pipeline/normalize/traits.py data/pipeline/normalize/talents.py data/tests/traits_fixture.py data/tests/test_traits.py
```

```bash
git commit -F - <<'MSG'
feat(data): read the client's trait tables as three-tab talent trees

A class tree is one SkillLineXTraitTree names; its three tabs are the
TraitNodeGroup rows maximal by inclusion, ordered by the PosX band most of
their nodes sit in, which is the tab's own OrderIndex. Positions quantise onto
the 4x7 grid in steps of 600 after a stray factor of ten is divided out, and
the two superseded nodes the build leaves behind are dropped by spell id.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 3: Prerequisites from the edges, and the tier ladder asserted

**Files:**
- Modify: `data/pipeline/normalize/traits.py`
- Modify: `data/tests/test_traits.py`

**Interfaces:**
- Consumes: everything Task 2 produced.
- Produces:
  - `pipeline.normalize.traits.check_tier_gates(rows: TraitRows, tree_ids: Iterable[int]) -> None` — raises `TraitDataError` when a tree's points gates are not three copies of `TIER_GATES`, or when `TraitCurrency` 3820's `SourcedMax` is not `CLASS_POINT_BUDGET`.
  - `read_trait_trees` now fills `TraitTalent.prereq_node_id` and `TraitTalent.prereq_rank`. A prerequisite's required rank is the prerequisite's own `max_rank`.

- [ ] **Step 1: Write the failing tests**

Append to `data/tests/test_traits.py`:

```python
from pipeline.normalize.traits import check_tier_gates


def test_an_edge_points_from_the_prerequisite_to_the_dependent():
    (warrior,) = read_trait_trees(trait_rows())
    links = {
        t.node_id: (t.prereq_node_id, t.prereq_rank)
        for tab in warrior.tabs
        for t in tab.talents
    }
    assert links == {
        900001: (None, None),
        900002: (None, None),
        # a vertical link, one row down
        900003: (900002, 5),
        900004: (None, None),
        900005: (None, None),
        # a horizontal link, same row
        900012: (900011, 5),
        900011: (None, None),
        900013: (None, None),
        # the reverse leg of the two-way pair is dropped, so 900021 keeps none
        900021: (None, None),
        # a vertical link spanning two rows, and an elbow
        900022: (900021, 2),
        900023: (900021, 2),
    }


def test_an_edge_that_leaves_its_tab_is_refused():
    edges = read_csv(TRAITS / "TraitEdge.csv")
    edges.append(
        {"ID": "6", "VisualStyle": "1", "LeftTraitNodeID": "900002", "RightTraitNodeID": "900012", "Type": "2"}
    )
    with pytest.raises(TraitDataError, match="different tabs"):
        read_trait_trees(trait_rows(edge=edges))


def test_a_second_prerequisite_for_one_talent_is_refused():
    edges = read_csv(TRAITS / "TraitEdge.csv")
    edges.append(
        {"ID": "6", "VisualStyle": "1", "LeftTraitNodeID": "900001", "RightTraitNodeID": "900003", "Type": "2"}
    )
    with pytest.raises(TraitDataError, match="900003"):
        read_trait_trees(trait_rows(edge=edges))


def test_the_tier_ladder_is_five_points_per_tier_in_every_tab():
    check_tier_gates(trait_rows(), [1117])


def test_a_tree_that_moves_a_tier_gate_is_refused():
    conds = [c for c in read_csv(TRAITS / "TraitCond.csv") if c["ID"] != "500013"]
    with pytest.raises(TraitDataError, match="points gates"):
        check_tier_gates(trait_rows(cond=conds), [1117])


def test_a_point_budget_other_than_fifty_one_is_refused():
    currency = [{"ID": "3820", "Type": "2", "SourcedMax": "60"}]
    with pytest.raises(TraitDataError, match="60"):
        check_tier_gates(trait_rows(currency=currency), [1117])
```

- [ ] **Step 2: Run the tests to verify they fail**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_traits.py -q --no-cov
```
Expected: FAIL — `ImportError: cannot import name 'check_tier_gates'`.

- [ ] **Step 3: Add the edge reader and the ladder check**

In `data/pipeline/normalize/traits.py`, extend the imports:

```python
from collections import Counter
from collections.abc import Iterable
from dataclasses import dataclass, replace
```

and add, after `_drop_stale_twins`:

```python
def check_tier_gates(rows: TraitRows, tree_ids: Iterable[int]) -> None:
    """Assert the client still gates a tab's rows at 5 points per tier.

    Every `TraitCond` row in build 1.60.1.69893 is `CondType` 0 and says only
    "this group needs N points spent in this tab's currency"; none of them
    carries a prerequisite rank. Filtering to the class currency, no pinned
    node and a non-zero amount leaves exactly eighteen rows per class tree --
    three tabs times 5/10/15/20/25/30 -- for all nine classes. The planner and
    the API both hardcode that ladder (`POINTS_PER_TIER`), so a build that
    changed it has to stop the pipeline rather than quietly disagree.
    """
    budgets = {int(c["ID"]): int(c["SourcedMax"]) for c in rows.currency}
    budget = budgets.get(CLASS_CURRENCY_ID)
    if budget != CLASS_POINT_BUDGET:
        raise TraitDataError(
            f"TraitCurrency {CLASS_CURRENCY_ID} allows {budget} points, "
            f"not the {CLASS_POINT_BUDGET} the planner spends"
        )
    want = sorted(TIER_GATES * len(COLUMN_ORIGINS))
    for tree_id in sorted(tree_ids):
        got = sorted(
            int(c["SpentAmountRequired"])
            for c in rows.cond
            if int(c["TraitTreeID"]) == tree_id
            and int(c["TraitCurrencyID"]) == CLASS_CURRENCY_ID
            and int(c["TraitNodeID"]) == 0
            and int(c["SpentAmountRequired"]) > 0
        )
        if got != want:
            raise TraitDataError(
                f"tree {tree_id}'s points gates are {got}, not {want}: the "
                "planner's five-points-per-tier rule would not match the client"
            )


def _with_prerequisites(tree: TraitClassTree, edge_rows: list[dict[str, str]]) -> TraitClassTree:
    """Fill in each talent's one prerequisite from `TraitEdge`.

    Left is the prerequisite and right the dependent, checked against the
    Wowhead snapshot's own `requires` arrays: 69 of the tree edges match it
    exactly. Two edges are the reverse leg of a two-way pair (Druid
    Nature's Splendor/Nature's Majesty, Hunter Bestial Wrath/Intimidation);
    dropping the leg whose prerequisite sits further down the tab leaves no
    cycles, at most one prerequisite per talent, and the direction the
    snapshot records. The required rank is not in the tables at all -- it is
    the prerequisite's own rank cap in all 69 cases.
    """
    by_node = {t.node_id: t for tab in tree.tabs for t in tab.talents}
    tab_of = {t.node_id: tab.tab_id for tab in tree.tabs for t in tab.talents}
    links: dict[int, tuple[int, int]] = {}
    for edge in edge_rows:
        left, right = int(edge["LeftTraitNodeID"]), int(edge["RightTraitNodeID"])
        # Edges of other trees, and the two edges into dropped stale nodes.
        if left not in by_node or right not in by_node:
            continue
        if tab_of[left] != tab_of[right]:
            raise TraitDataError(
                f"edge {edge['ID']} joins nodes {left} and {right} in different tabs"
            )
        prerequisite, dependent = by_node[left], by_node[right]
        if prerequisite.row > dependent.row:
            logger.info(
                "dropping edge %s: node %s is below its dependent %s",
                edge["ID"],
                left,
                right,
            )
            continue
        if right in links:
            raise TraitDataError(
                f"node {right} has two prerequisites, {links[right][0]} and {left}; "
                "the planner's rules carry one"
            )
        links[right] = (left, prerequisite.max_rank)
    return TraitClassTree(
        class_id=tree.class_id,
        tree_id=tree.tree_id,
        tabs=tuple(
            TraitTab(
                tab_id=tab.tab_id,
                name=tab.name,
                position=tab.position,
                background=tab.background,
                talents=tuple(
                    replace(
                        talent,
                        prereq_node_id=links[talent.node_id][0],
                        prereq_rank=links[talent.node_id][1],
                    )
                    if talent.node_id in links
                    else talent
                    for talent in tab.talents
                ),
            )
            for tab in tree.tabs
        ),
    )
```

Then in `read_trait_trees`, replace the final `return` with:

```python
    check_tier_gates(rows, class_of_tree)
    linked = [_with_prerequisites(tree, rows.edge) for tree in trees]
    return sorted(linked, key=lambda t: t.class_id)
```

- [ ] **Step 4: Run the tests to verify they pass**

```bash
cd data && uv run pytest tests/test_traits.py -q --no-cov
```
Expected: PASS, 17 passed.

- [ ] **Step 5: Lint**

```bash
cd data && uv run ruff check pipeline/normalize/traits.py tests/test_traits.py
```
Expected: `All checks passed!`

- [ ] **Step 6: Commit**

```bash
git add data/pipeline/normalize/traits.py data/tests/test_traits.py
```

```bash
git commit -F - <<'MSG'
feat(data): prerequisites from TraitEdge, and the tier ladder asserted

TraitEdge's left node is the prerequisite and its right the dependent; the two
reverse legs of the build's two-way pairs are dropped by comparing rows, which
leaves one prerequisite per talent and no cycles. The required rank is in no
table -- it is the prerequisite's own rank cap, which matches the Wowhead
snapshot in all 69 cases.

TraitCond carries only points-spent gates, and they are exactly 5/10/15/20/25/30
per tab for all nine classes; check_tier_gates stops the pipeline if a future
build disagrees with the planner's hardcoded rule.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 4: Per-rank values from the curves, and the spell text that renders them

**Files:**
- Create: `data/pipeline/curves.py`
- Create: `data/tests/test_curves.py`
- Create: `data/tests/fixtures/traits/TraitDefinitionEffectPoints.csv`
- Create: `data/tests/fixtures/traits/CurvePoint.csv`
- Modify: `data/pipeline/spelltext.py:9-36` (docstring, `TOKEN`), `:80-84` (`describe`), `:104-117` (`_substitute`)
- Modify: `data/tests/test_spelltext.py`
- Modify: `data/tests/fixtures/Spell.csv`, `data/tests/fixtures/SpellEffect.csv`, `data/tests/fixtures/SpellMisc.csv`

**Interfaces:**
- Consumes: `pipeline.spelltext.SpellText`, `Effect`, `SpellRow`.
- Produces:
  - `pipeline.curves.CurveDataError(ValueError)`
  - `pipeline.curves.RankPoints` — frozen, field `by_definition: dict[int, dict[int, dict[int, int]]]` (definition id -> rank -> 0-based effect index -> value); methods `for_rank(definition_id: int, rank: int) -> dict[int, int]` and `highest_rank(definition_id: int) -> int`.
  - `pipeline.curves.load_rank_points(effect_point_rows: list[dict[str, str]], curve_point_rows: list[dict[str, str]]) -> RankPoints`
  - `pipeline.spelltext.SpellText.describe(spell_id: int, overrides: Mapping[int, int] | None = None) -> str` — `overrides` maps a 0-based effect index to the value that effect displays. Existing single-argument callers are unaffected.

- [ ] **Step 1: Write the failing curve test**

```python
# data/tests/test_curves.py
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.curves import CurveDataError, load_rank_points

HERE = Path(__file__).parent
TRAITS = HERE / "fixtures/traits"


def fixture_points():
    return load_rank_points(
        read_csv(TRAITS / "TraitDefinitionEffectPoints.csv"),
        read_csv(TRAITS / "CurvePoint.csv"),
    )


def test_each_rank_gets_the_curve_value_for_every_effect_of_the_definition():
    points = fixture_points()
    # 700001 has one effect; 700002 has two, whose curves rise at different rates.
    assert points.for_rank(700001, 1) == {0: 1}
    assert points.for_rank(700001, 3) == {0: 3}
    assert points.for_rank(700002, 1) == {0: 2, 1: 20}
    assert points.for_rank(700002, 5) == {0: 10, 1: 100}


def test_highest_rank_is_the_last_point_on_the_curve():
    points = fixture_points()
    assert points.highest_rank(700001) == 3
    assert points.highest_rank(700002) == 5


def test_a_definition_with_no_curve_has_no_points_and_says_so_quietly():
    points = fixture_points()
    assert points.for_rank(700013, 1) == {}
    assert points.highest_rank(700013) == 0


def test_an_operation_other_than_setting_the_value_is_refused():
    rows = [
        {"ID": "1", "TraitDefinitionID": "700001", "EffectIndex": "0", "OperationType": "1", "CurveID": "90001"}
    ]
    with pytest.raises(CurveDataError, match="OperationType 1"):
        load_rank_points(rows, read_csv(TRAITS / "CurvePoint.csv"))


def test_a_curve_with_no_points_is_refused():
    rows = [
        {"ID": "1", "TraitDefinitionID": "700001", "EffectIndex": "0", "OperationType": "0", "CurveID": "99999"}
    ]
    with pytest.raises(CurveDataError, match="99999"):
        load_rank_points(rows, read_csv(TRAITS / "CurvePoint.csv"))
```

And the fixtures it reads.

`data/tests/fixtures/traits/TraitDefinitionEffectPoints.csv`:

```csv
ID,TraitDefinitionID,EffectIndex,OperationType,CurveID
23001,700001,0,0,90001
23002,700002,0,0,90002
23003,700002,1,0,90003
23004,700003,0,0,90004
23005,700005,0,0,90005
23010,700004,0,0,90004
23006,700011,0,0,90002
23007,700012,0,0,90002
23008,700021,0,0,90006
23009,700023,0,0,90006
```

`data/tests/fixtures/traits/CurvePoint.csv`:

```csv
Pos_0,Pos_1,PosPreSquish_0,PosPreSquish_1,ID,CurveID,OrderIndex
1,1,0,0,1,90001,0
2,2,0,0,2,90001,1
3,3,0,0,3,90001,2
1,2,0,0,4,90002,0
2,4,0,0,5,90002,1
3,6,0,0,6,90002,2
4,8,0,0,7,90002,3
5,10,0,0,8,90002,4
1,20,0,0,9,90003,0
2,40,0,0,10,90003,1
3,60,0,0,11,90003,2
4,80,0,0,12,90003,3
5,100,0,0,13,90003,4
1,3,0,0,14,90004,0
2,6,0,0,15,90004,1
3,9,0,0,16,90004,2
4,12,0,0,17,90004,3
5,15,0,0,18,90004,4
1,10,0,0,19,90005,0
2,20,0,0,20,90005,1
1,5,0,0,21,90006,0
2,10,0,0,22,90006,1
```

- [ ] **Step 2: Run the test to verify it fails**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_curves.py -q --no-cov
```
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.curves'`.

- [ ] **Step 3: Write the curve reader**

```python
# data/pipeline/curves.py
"""What a trait talent's effects are worth at each of its ranks.

The 1.60 client has no per-rank spells. A talent is one `TraitDefinition`
pointing at one spell, and `TraitDefinitionEffectPoints` says which curve
drives each of that spell's effects; the curve's points are the ranks, with
`Pos_0` the rank and `Pos_1` the value. Measured on build 1.60.1.69893: 640
rows, `OperationType` 0 in every one, `Pos_0` always 1..5, every `Pos_1` a
whole number, and every multi-rank talent covered.

`pipeline/normalize/trait_trees.py` hands the values for one rank to
`SpellText.describe` as effect overrides, which is how a rank gets its own
sentence out of the one description the client stores.
"""

from __future__ import annotations

from dataclasses import dataclass


class CurveDataError(ValueError):
    """The curve tables do not have the shape per-rank values need."""


@dataclass(frozen=True)
class RankPoints:
    """definition id -> rank -> 0-based effect index -> the effect's value."""

    by_definition: dict[int, dict[int, dict[int, int]]]

    def for_rank(self, definition_id: int, rank: int) -> dict[int, int]:
        """The effect values at `rank`, or {} when this definition has no curve."""
        return self.by_definition.get(definition_id, {}).get(rank, {})

    def highest_rank(self, definition_id: int) -> int:
        """The last rank any of this definition's curves names; 0 when it has none."""
        ranks = self.by_definition.get(definition_id)
        return max(ranks) if ranks else 0


def load_rank_points(
    effect_point_rows: list[dict[str, str]], curve_point_rows: list[dict[str, str]]
) -> RankPoints:
    curves: dict[int, dict[int, int]] = {}
    for row in curve_point_rows:
        rank = round(float(row["Pos_0"]))
        curves.setdefault(int(row["CurveID"]), {})[rank] = round(float(row["Pos_1"]))
    by_definition: dict[int, dict[int, dict[int, int]]] = {}
    for row in effect_point_rows:
        # 0 is "this curve gives the effect's value". Anything else would mean
        # the value is computed from the spell's own base points in a way this
        # module does not implement, and quietly emitting the base value would
        # put a wrong number in a tooltip.
        if row["OperationType"] != "0":
            raise CurveDataError(
                f"TraitDefinitionEffectPoints {row['ID']} has OperationType "
                f"{row['OperationType']}; only 0 (set the value) is understood"
            )
        curve_id = int(row["CurveID"])
        curve = curves.get(curve_id)
        if not curve:
            raise CurveDataError(
                f"TraitDefinitionEffectPoints {row['ID']} names curve {curve_id}, "
                "which has no points"
            )
        index = int(row["EffectIndex"])
        ranks = by_definition.setdefault(int(row["TraitDefinitionID"]), {})
        for rank, value in curve.items():
            ranks.setdefault(rank, {})[index] = value
    return RankPoints(by_definition)
```

- [ ] **Step 4: Run the curve test to verify it passes**

```bash
cd data && uv run pytest tests/test_curves.py -q --no-cov
```
Expected: PASS, 5 passed.

- [ ] **Step 5: Write the failing spell-text tests**

Append to `data/tests/test_spelltext.py`:

```python
def test_m_token_renders_the_effects_minimum():
    # 724 of the 1320 rank descriptions in build 1.60.1.69893 use $m.
    text = fixture_text()
    assert text.describe(12321) == "Increases the duration of your shouts by 10%."


def test_overrides_replace_an_effects_value_for_one_rank():
    text = fixture_text()
    assert text.describe(12321, {0: 30}) == "Increases the duration of your shouts by 30%."
    # The spell itself is untouched: the next call sees the client's own value.
    assert text.describe(12321) == "Increases the duration of your shouts by 10%."


def test_overrides_reach_an_effect_the_spell_has_no_row_for():
    text = fixture_text()
    assert text.describe(12777, {1: 7}) == "Deals 4 damage and stuns for 7 sec."
```

and add the two spells those tests need.

`data/tests/fixtures/Spell.csv` gains:

```csv
12321,,"Increases the duration of your shouts by $m1%.",
12777,,"Deals $m1 damage and stuns for $m2 sec.",
```

`data/tests/fixtures/SpellEffect.csv` gains:

```csv
12321,0,0,10,0,0
12777,0,0,4,0,0
```

`data/tests/fixtures/SpellMisc.csv` gains:

```csv
12321,0,0,132333
12777,0,0,132333
```

- [ ] **Step 6: Run the spell-text tests to verify they fail**

```bash
cd data && uv run pytest tests/test_spelltext.py -q --no-cov
```
Expected: FAIL — the `$m1` tokens come back verbatim, and `describe()` takes one argument.

- [ ] **Step 7: Extend the renderer**

In `data/pipeline/spelltext.py`:

Add to the module docstring's token list, after the `$o<n>` line:

```
  $m<n> / $M<n>        the minimum of effect <n>'s displayed value
```

Widen `TOKEN`'s kind class:

```python
    r"(?P<kind>[sSoOtTdDmM])"
```

Extend the imports:

```python
from collections.abc import Mapping
from dataclasses import dataclass, field, replace
```

Replace `describe`:

```python
    def describe(self, spell_id: int, overrides: Mapping[int, int] | None = None) -> str:
        """The spell's description with its `$`-tokens resolved.

        `overrides` maps a 0-based effect index to the value that effect
        displays, which is how one rank of a trait talent gets its own
        sentence: the 1.60 client stores one description per talent and puts
        the per-rank numbers on a curve (see `pipeline/curves.py`). Only the
        described spell's own effects are overridden -- a `$<id>s1` token
        still reads the referenced spell as the client stores it.
        """
        row = self._spells.get(spell_id)
        if row is None:
            return ""
        if overrides:
            row = replace(row, effects={**row.effects, **_overridden(row.effects, overrides)})
        return TOKEN.sub(lambda match: self._substitute(row, match), row.description)
```

Add beside it, at module level:

```python
def _overridden(effects: dict[int, Effect], overrides: Mapping[int, int]) -> dict[int, Effect]:
    """The overridden effects, keeping each one's tick period.

    Die sides go to zero: a curve names one value per rank, not a spread, so
    `$s` and `$m` must both render exactly that number.
    """
    out: dict[int, Effect] = {}
    for index, value in overrides.items():
        period = effects[index].period_ms if index in effects else 0
        out[index] = Effect(base_points=value, die_sides=0, period_ms=period)
    return out
```

And in `_substitute`, immediately after the `if kind == "o":` block and before the final `return`:

```python
        if kind == "m":
            high = low
```

- [ ] **Step 8: Run the spell-text tests to verify they pass**

```bash
cd data && uv run pytest tests/test_spelltext.py tests/test_normalize_talent_trees.py tests/test_normalize_gear.py -q --no-cov
```
Expected: PASS — 17 in `test_spelltext.py`, and the two suites that call `describe` with one argument unchanged.

- [ ] **Step 9: Lint**

```bash
cd data && uv run ruff check pipeline/curves.py pipeline/spelltext.py tests/test_curves.py tests/test_spelltext.py
```
Expected: `All checks passed!`

- [ ] **Step 10: Commit**

```bash
git add data/pipeline/curves.py data/pipeline/spelltext.py data/tests/test_curves.py data/tests/test_spelltext.py data/tests/fixtures
```

```bash
git commit -F - <<'MSG'
feat(data): per-rank talent values from the client's curves

The 1.60 client has no per-rank spells: a talent is one spell whose effects
are driven per rank by a curve named in TraitDefinitionEffectPoints. Read them
into RankPoints, and let SpellText.describe take those values as effect
overrides so each rank renders its own sentence.

Also teaches the renderer $m/$M, the minimum of an effect -- 724 of the beta
build's 1320 rank descriptions use it and it came out verbatim before. With
both, 998 of those 1320 render byte-identical to the Wowhead snapshot's text.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 5: The emitted schema, and the trait assembler

**Files:**
- Modify: `data/pipeline/models.py:85-101` (`TalentEntry`, `TalentTree`)
- Modify: `data/pipeline/normalize/talent_trees.py:22-40` (`_talent_name` -> `talent_name`), `:62-90` (fill the two new fields)
- Create: `data/pipeline/normalize/trait_trees.py`
- Create: `data/tests/test_normalize_trait_trees.py`
- Create: `data/tests/fixtures/traits/Spell.csv`, `SpellMisc.csv`, `SpellEffect.csv`, `SpellDuration.csv`, `SpellName.csv`, `ManifestInterfaceData.csv`
- Modify: `data/tests/fixtures/TalentTab.csv` (a `BackgroundFile` column)
- Modify: `data/tests/golden/warrior.json` (regenerated)

**Interfaces:**
- Consumes: `pipeline.normalize.traits.read_trait_trees`, `TraitRows`; `pipeline.curves.load_rank_points`, `RankPoints`; `pipeline.spelltext.SpellText`; `pipeline.icons.resolve_icon`; `pipeline.normalize.classes.slugify`.
- Produces:
  - `pipeline.models.TalentEntry.spell_id: int` — the client's spell for the talent, appended after `ranks`.
  - `pipeline.models.TalentTree.background: str` — the tab's `BackgroundFile`, lowercased, appended after `talents`. Names the image at `builds/<build>/trees/<background>.webp`.
  - `pipeline.normalize.talent_trees.talent_name(talent_id: int, spell_id: int, spell_names: dict[int, str]) -> str` (renamed from `_talent_name`).
  - `pipeline.normalize.trait_trees.build_trait_talent_trees(rows: TraitRows, class_rows: list[dict[str, str]], spell_names: dict[int, str], spell_text: SpellText, rank_points: RankPoints, icons: dict[int, str], build: str) -> list[ClassTalents]`
  - `pipeline.normalize.trait_trees.TraitRankError(ValueError)`

New fields go **last** in each model so the existing keys keep their place in `model_dump()`, which is what makes the Era build's re-emit a pure addition (Task 6).

The spec asks the emitted tree for an `order` "from `OrderIndex`". That field already exists and is called `position` — `build_talent_trees` has set it from `TalentTab.OrderIndex` since Phase 1, and `api/internal/trees` and `web/src/lib/planner/rules.ts` both sort by it. Do not add a second field; `background` and `spell_id` are the only two additions.

- [ ] **Step 1: Write the failing test**

```python
# data/tests/test_normalize_trait_trees.py
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.curves import load_rank_points
from pipeline.icons import PLACEHOLDER_ICON, icon_names
from pipeline.normalize.trait_trees import TraitRankError, build_trait_talent_trees
from pipeline.spelltext import load_spell_text
from traits_fixture import TRAITS, trait_rows


def fixture_records(**replace):
    return build_trait_talent_trees(
        trait_rows(**replace),
        read_csv(TRAITS / "ChrClasses.csv"),
        {int(r["ID"]): r["Name_lang"] for r in read_csv(TRAITS / "SpellName.csv")},
        load_spell_text(
            read_csv(TRAITS / "Spell.csv"),
            read_csv(TRAITS / "SpellMisc.csv"),
            read_csv(TRAITS / "SpellEffect.csv"),
            read_csv(TRAITS / "SpellDuration.csv"),
        ),
        load_rank_points(
            read_csv(TRAITS / "TraitDefinitionEffectPoints.csv"),
            read_csv(TRAITS / "CurvePoint.csv"),
        ),
        icon_names(read_csv(TRAITS / "ManifestInterfaceData.csv")),
        "1.60.1.69893",
    )


def test_only_the_class_with_a_trait_tree_is_emitted():
    records = fixture_records()
    assert [r.class_slug for r in records] == ["warrior"]
    assert records[0].build == "1.60.1.69893"
    assert records[0].class_id == 1


def test_the_three_trees_carry_the_games_order_and_their_background():
    (warrior,) = fixture_records()
    assert [(t.id, t.name, t.position, t.background) for t in warrior.trees] == [
        (161, "Arms", 0, "warriorarms"),
        (164, "Fury", 1, "warriorfury"),
        (163, "Protection", 2, "warriorprotection"),
    ]


def test_a_talent_is_its_node_with_the_clients_spell_name_and_icon():
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    heroic = arms[900001]
    assert heroic.name == "Improved Heroic Strike"
    assert heroic.icon == "ability_rogue_ambush"
    assert heroic.spell_id == 12282
    assert (heroic.tier, heroic.column, heroic.max_rank) == (0, 0, 3)


def test_every_rank_gets_its_own_text_from_the_curve():
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    assert [r.description for r in arms[900001].ranks] == [
        "Increases your Strength by 1%.",
        "Increases your Strength by 2%.",
        "Increases your Strength by 3%.",
    ]
    # Every rank carries the same spell id: the client has no per-rank spells.
    assert [r.spell_id for r in arms[900001].ranks] == [12282, 12282, 12282]


def test_a_talent_with_no_curve_repeats_the_clients_one_sentence():
    (warrior,) = fixture_records()
    fury = {t.id: t for t in warrior.trees[1].talents}
    assert fury[900013].max_rank == 1
    assert fury[900013].ranks[0].description == "Dazes the target for 6 sec."


def test_a_curve_with_more_points_than_the_talent_has_ranks_is_clamped():
    # 700004 is a 3-rank talent on a 5-point curve, the shape three of the
    # beta build's single-rank talents have.
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    assert len(arms[900004].ranks) == 3
    assert arms[900004].ranks[2].description == "Reduces the target's attack speed by 9%."


def test_a_multi_rank_talent_with_no_value_for_a_rank_is_refused():
    entries = read_csv(TRAITS / "TraitNodeEntry.csv")
    for entry in entries:
        if entry["ID"] == "800013":  # 1 rank and no curve; ask for 4
            entry["MaxRanks"] = "4"
    with pytest.raises(TraitRankError, match="700013"):
        fixture_records(node_entry=entries)


def test_prerequisites_are_carried_through_as_talent_ids():
    (warrior,) = fixture_records()
    arms = {t.id: t for t in warrior.trees[0].talents}
    assert (arms[900003].prereq_talent_id, arms[900003].prereq_rank) == (900002, 5)
    assert (arms[900001].prereq_talent_id, arms[900001].prereq_rank) == (None, None)


def test_a_talent_whose_spell_has_no_icon_falls_back_to_the_placeholder():
    (warrior,) = fixture_records()
    protection = {t.id: t for t in warrior.trees[2].talents}
    assert protection[900022].icon == PLACEHOLDER_ICON


def test_talents_are_sorted_by_cell():
    (warrior,) = fixture_records()
    for tree in warrior.trees:
        cells = [(t.tier, t.column) for t in tree.talents]
        assert cells == sorted(cells)
```

- [ ] **Step 2: Write the spell-side fixtures**

`data/tests/fixtures/traits/SpellName.csv`:

```csv
ID,Name_lang
12282,Improved Heroic Strike
16462,Deflection
12295,Tactical Mastery
12296,Improved Thunder Clap
12163,Impale
12321,Booming Voice
12322,Cruelty
12323,Piercing Howl
12797,Improved Bloodrage
12287,Improved Revenge
```

`data/tests/fixtures/traits/Spell.csv` — 12809 is deliberately absent, so `900022` has to fall back to the placeholder icon and an empty description:

```csv
ID,NameSubtext_lang,Description_lang,AuraDescription_lang
12282,,"Increases your Strength by $m1%.",
16462,,"Increases your Parry chance by $m1% and your Dodge chance by $m2%.",
12295,,"Keeps you in combat stance for $m1 sec.",
12296,,"Reduces the target's attack speed by $m1%.",
12163,,"Adds $m1% of your Strength to your critical strikes.",
12321,,"Increases the duration of your shouts by $m1%.",
12322,,"Increases your critical strike chance by $m1%.",
12323,,"Dazes the target for $d.",
12797,,"Generates $m1 extra rage.",
12287,,"Increases Revenge damage by $m1%.",
```

`data/tests/fixtures/traits/SpellEffect.csv`:

```csv
SpellID,DifficultyID,EffectIndex,EffectBasePoints,EffectDieSides,EffectAuraPeriod
12282,0,0,1,0,0
16462,0,0,2,0,0
16462,0,1,20,0,0
12295,0,0,3,0,0
12296,0,0,3,0,0
12163,0,0,10,0,0
12321,0,0,2,0,0
12322,0,0,2,0,0
12797,0,0,5,0,0
12287,0,0,5,0,0
```

`data/tests/fixtures/traits/SpellDuration.csv`:

```csv
ID,Duration
21,6000
```

`data/tests/fixtures/traits/SpellMisc.csv` — 12809 has no row at all, and 12287's icon file id names no manifest row, which is the second half of the placeholder path:

```csv
SpellID,DifficultyID,DurationIndex,SpellIconFileDataID
12282,0,0,132333
16462,0,0,132269
12295,0,0,132353
12296,0,0,132326
12163,0,0,132287
12321,0,0,132345
12322,0,0,132347
12323,0,21,132316
12797,0,0,132277
12287,0,0,999999
```

`data/tests/fixtures/traits/ManifestInterfaceData.csv`:

```csv
ID,FilePath,FileName
132333,Interface\ICONS\,Ability_Rogue_Ambush.blp
132269,Interface\ICONS\,Ability_Defend.blp
132353,Interface\ICONS\,Ability_Warrior_InnerRage.blp
132326,Interface\ICONS\,Ability_ThunderClap.blp
132287,Interface\ICONS\,Ability_SearingArrow.blp
132345,Interface\ICONS\,Ability_Warrior_WarCry.blp
132347,Interface\ICONS\,Ability_Warrior_Rampage.blp
132316,Interface\ICONS\,Ability_GolemStormBolt.blp
132277,Interface\ICONS\,Ability_Racial_BloodRage.blp
```

- [ ] **Step 3: Run the test to verify it fails**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_normalize_trait_trees.py -q --no-cov
```
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.normalize.trait_trees'`.

- [ ] **Step 4: Add the two fields to the models**

In `data/pipeline/models.py`, replace `TalentEntry` and `TalentTree` with:

```python
class TalentEntry(BaseModel):
    id: int
    name: str
    icon: str
    max_rank: int
    tier: int
    column: int
    prereq_talent_id: int | None
    prereq_rank: int | None
    ranks: list[TalentRank]
    #: The client's own spell for the talent, which is what every consumer
    #: keyed by spell id (the sim, the report's planner link) will see written
    #: by the game. Appended last, like `background` below: the emitted key
    #: order is the schema's only compatibility surface, so new fields go on
    #: the end and existing ones never move.
    spell_id: int


class TalentTree(BaseModel):
    id: int
    name: str
    position: int
    talents: list[TalentEntry]
    #: The tab's `BackgroundFile`, lowercased. The site fetches the processed
    #: image at `/data/<build>/trees/<background>.webp`.
    background: str
```

- [ ] **Step 5: Fill them on the legacy path**

In `data/pipeline/normalize/talent_trees.py`:

Rename `_talent_name` to `talent_name` (declaration and its one call site) and add to its docstring:

```python
def talent_name(talent_id: int, spell_id: int, spell_names: dict[int, str]) -> str:
    """The first rank's spell name, or a placeholder that names the missing row.

    Shared with `pipeline/normalize/trait_trees.py`, which names a talent the
    same way off the client's trait tables.
    """
```

In the `tabs` comprehension, carry the background:

```python
    tabs = [
        {
            "id": int(row["ID"]),
            "name": row["Name_lang"],
            "position": int(row["OrderIndex"]),
            "class_mask": int(row["ClassMask"]),
            "background": row["BackgroundFile"].lower(),
        }
        for row in tab_rows
    ]
```

In the `TalentEntry(...)` construction, add as the last argument:

```python
            # The legacy table's talent is a chain of per-rank spells, so the
            # talent's own spell is the first rank's -- the same id the client
            # writes when the talent is learned.
            spell_id=ranks[0],
```

And in the `TalentTree(...)` construction, add as the last argument:

```python
                background=tab["background"],
```

- [ ] **Step 6: Give the shared TalentTab fixture a BackgroundFile**

`data/tests/fixtures/TalentTab.csv` becomes:

```csv
ID,Name_lang,ClassMask,OrderIndex,BackgroundFile
161,Arms,1,0,WarriorArms
164,Fury,1,1,WarriorFury
```

- [ ] **Step 7: Write the assembler**

```python
# data/pipeline/normalize/trait_trees.py
"""The client's trait tables -> one talents/<class-slug>.json per class.

The trait-table twin of `talent_trees.py`: same output shape, same field
meanings, different source. Names, icons and descriptions still come from the
spell tables, so a talent looks the same to the site whichever reader emitted
it; what differs is that a trait talent is one spell whose per-rank numbers
come off a curve (`pipeline/curves.py`) rather than a chain of per-rank
spells.
"""

from __future__ import annotations

import logging

from pipeline.curves import RankPoints
from pipeline.icons import resolve_icon
from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.normalize.classes import slugify
from pipeline.normalize.talent_trees import talent_name
from pipeline.normalize.traits import TraitDataError, TraitRows, TraitTalent, read_trait_trees
from pipeline.spelltext import SpellText

logger = logging.getLogger(__name__)


class TraitRankError(ValueError):
    """A talent has ranks the client gives no values for."""


def build_trait_talent_trees(
    rows: TraitRows,
    class_rows: list[dict[str, str]],
    spell_names: dict[int, str],
    spell_text: SpellText,
    rank_points: RankPoints,
    icons: dict[int, str],
    build: str,
) -> list[ClassTalents]:
    slugs = {int(row["ID"]): slugify(row["Name_lang"]) for row in class_rows}
    records: list[ClassTalents] = []
    for tree in read_trait_trees(rows):
        slug = slugs.get(tree.class_id)
        if slug is None:
            raise TraitDataError(
                f"tree {tree.tree_id} names class {tree.class_id}, which has no ChrClasses row"
            )
        records.append(
            ClassTalents(
                build=build,
                class_id=tree.class_id,
                class_slug=slug,
                trees=[
                    TalentTree(
                        id=tab.tab_id,
                        name=tab.name,
                        position=tab.position,
                        background=tab.background.lower(),
                        talents=sorted(
                            (
                                _entry(talent, spell_names, spell_text, rank_points, icons)
                                for talent in tab.talents
                            ),
                            key=lambda t: (t.tier, t.column, t.id),
                        ),
                    )
                    for tab in tree.tabs
                ],
            )
        )
    return sorted(records, key=lambda r: r.class_id)


def _entry(
    talent: TraitTalent,
    spell_names: dict[int, str],
    spell_text: SpellText,
    rank_points: RankPoints,
    icons: dict[int, str],
) -> TalentEntry:
    highest = rank_points.highest_rank(talent.definition_id)
    if talent.max_rank > 1 and highest < talent.max_rank:
        raise TraitRankError(
            f"definition {talent.definition_id} (node {talent.node_id}, spell "
            f"{talent.spell_id}) has {talent.max_rank} ranks but its curves "
            f"reach rank {highest}"
        )
    if highest > talent.max_rank:
        # Three single-rank talents in build 1.60.1.69893 (Ice Block,
        # Thousand Cuts, Last Stand) sit on curves with points the node can
        # never reach. Clamping to MaxRanks is what the client draws.
        logger.debug(
            "definition %s has curve points up to rank %s but only %s ranks; clamping",
            talent.definition_id,
            highest,
            talent.max_rank,
        )
    return TalentEntry(
        id=talent.node_id,
        name=talent_name(talent.node_id, talent.spell_id, spell_names),
        icon=resolve_icon(
            spell_text.icon_file_id(talent.spell_id),
            icons,
            f"talent {talent.node_id} (spell {talent.spell_id})",
        ),
        max_rank=talent.max_rank,
        tier=talent.row,
        column=talent.column,
        prereq_talent_id=talent.prereq_node_id,
        prereq_rank=talent.prereq_rank,
        ranks=[
            TalentRank(
                spell_id=talent.spell_id,
                description=spell_text.describe(
                    talent.spell_id, rank_points.for_rank(talent.definition_id, rank)
                ),
            )
            for rank in range(1, talent.max_rank + 1)
        ],
        spell_id=talent.spell_id,
    )
```

- [ ] **Step 8: Run the new test to verify it passes**

```bash
cd data && uv run pytest tests/test_normalize_trait_trees.py -q --no-cov
```
Expected: PASS, 10 passed.

- [ ] **Step 9: Regenerate the legacy golden and run its suite**

```bash
cd data && uv run pytest tests/test_normalize_talent_trees.py -q --no-cov
```
Expected: FAIL on `test_warrior_trees_match_golden` — the golden has no `spell_id`/`background`. Regenerate it:

```bash
cd data && uv run python - <<'PY'
from pathlib import Path
import sys
sys.path.insert(0, ".")
from tests.test_normalize_talent_trees import build_warrior
from pipeline.normalize import write_model
write_model(build_warrior()[0], Path("tests/golden/warrior.json"))
PY
```

```bash
cd data && git diff --stat tests/golden/warrior.json && git diff tests/golden/warrior.json | grep -c '^-' 
```
Expected: the only removed lines are the diff's own `---` header line (count 1). Every other change is an added `"spell_id"` or `"background"` line.

```bash
cd data && uv run pytest tests/test_normalize_talent_trees.py -q --no-cov
```
Expected: PASS.

- [ ] **Step 10: Lint**

```bash
cd data && uv run ruff check pipeline tests
```
Expected: `All checks passed!`

- [ ] **Step 11: Commit**

```bash
git add data/pipeline/models.py data/pipeline/normalize/talent_trees.py data/pipeline/normalize/trait_trees.py data/tests
```

```bash
git commit -F - <<'MSG'
feat(data): emit talent trees from the client's trait tables

build_trait_talent_trees is the trait-table twin of build_talent_trees: same
output shape, names and icons still from the spell tables, but a talent is one
spell whose per-rank numbers come off a curve rather than a chain of per-rank
spells.

The emitted schema gains two fields, both appended so no existing key moves:
talent.spell_id, the id the client writes when the talent is learned, and
tree.background, the tab's own art name. The legacy reader fills both too.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 6: Wire the reader in, re-emit both builds, and diff against the snapshot

**Files:**
- Modify: `data/pipeline/normalize/talents.py` (add `flat_talents`)
- Modify: `data/pipeline/normalize/__init__.py:59-100` (pick the reader; optional tables)
- Modify: `data/pipeline/forever.py:78-92` (use `flat_talents`; carry the new fields)
- Modify: `data/pipeline/normalize/forever_talents.py:99-125` (set `spell_id` and `background`)
- Modify: `data/pipeline/__main__.py` (a `wowhead-diff` subcommand)
- Create: `data/pipeline/wowhead_diff.py`
- Create: `data/tests/test_wowhead_diff.py`
- Create: `data/tests/test_beta_build.py`
- Modify: `data/tests/test_build_conformance.py:27-37` (the two new keys)
- Modify: `data/tests/test_forever_talents.py` (the two new fields)
- Modify: `data/README.md` (the diff's summary)
- Regenerate: `data/builds/1.15.9.69722/**`, `data/builds/1.60.1.69893/**`, `data/builds/forever-prebeta/**`
- Create: `data/diffs/wowhead-2026-09-14__1.60.1.69893.json`

**Interfaces:**
- Consumes: everything Tasks 2-5 produced.
- Produces:
  - `pipeline.normalize.talents.flat_talents(records: Sequence[ClassTalents]) -> list[TalentNode]` — the flat `talents.json` rows for already-built per-class records.
  - `pipeline.wowhead_diff.diff_snapshot(snapshot: dict, records: Sequence[ClassTalents]) -> dict`
  - `pipeline.wowhead_diff.write_snapshot_diff(snapshot: str, build: str, root: Path = Path("builds"), out: Path = Path("diffs")) -> Path`
  - `python -m pipeline wowhead-diff --snapshot <path> --build <build>`

- [ ] **Step 1: Write the failing diff test**

```python
# data/tests/test_wowhead_diff.py
from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.wowhead_diff import diff_snapshot


def entry(talent_id, name, tier, column, max_rank=1, prereq=None, prereq_rank=None):
    return TalentEntry(
        id=talent_id,
        name=name,
        icon="icon",
        max_rank=max_rank,
        tier=tier,
        column=column,
        prereq_talent_id=prereq,
        prereq_rank=prereq_rank,
        ranks=[TalentRank(spell_id=1, description="x") for _ in range(max_rank)],
        spell_id=1,
    )


def records(talents):
    return [
        ClassTalents(
            build="b",
            class_id=1,
            class_slug="warrior",
            trees=[
                TalentTree(id=161, name="Arms", position=0, talents=talents, background="warriorarms")
            ],
        )
    ]


def snapshot(talents):
    return {
        "trees": {"161": {"id": 161, "description": "WarriorArms"}},
        "talents": {"161": {str(t["id"]): t for t in talents}},
    }


def snap_talent(talent_id, name, row, col, ranks=1, requires=()):
    return {
        "id": talent_id,
        "name": name,
        "row": row,
        "col": col,
        "ranks": [None] * ranks,
        "descriptions": {str(n): "x" for n in range(1, ranks + 1)},
        "requires": [{"id": rid, "qty": qty} for rid, qty in requires],
    }


def test_identical_trees_report_no_change():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Deflection", 0, 1)]),
        records([entry(900002, "Deflection", 0, 1)]),
    )
    assert result["totals"] == {"snapshot": 1, "build": 1}
    assert result["trees"] == []


def test_a_new_name_in_the_same_cell_is_a_rename():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Drain Hope", 6, 1)]),
        records([entry(905916, "Wrack", 6, 1)]),
    )
    (tree,) = result["trees"]
    assert tree["renamed"] == [{"cell": [6, 1], "snapshot": "Drain Hope", "build": "Wrack"}]
    assert tree["added"] == [] and tree["removed"] == []


def test_a_name_that_appears_in_another_cell_is_a_move():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Tidal Mastery", 0, 2), snap_talent(2, "Totemic Focus", 3, 0)]),
        records([entry(11, "Totemic Focus", 0, 2), entry(12, "Tidal Mastery", 3, 0)]),
    )
    (tree,) = result["trees"]
    assert tree["moved"] == [
        {"name": "Tidal Mastery", "snapshot": [0, 2], "build": [3, 0]},
        {"name": "Totemic Focus", "snapshot": [3, 0], "build": [0, 2]},
    ]
    assert tree["renamed"] == []


def test_an_empty_cell_on_one_side_is_an_addition_or_a_removal():
    result = diff_snapshot(
        snapshot([snap_talent(1, "Balance of Nature", 2, 3)]),
        records([entry(11, "Improved Serpent Sting", 3, 3, max_rank=5)]),
    )
    (tree,) = result["trees"]
    assert tree["removed"] == [{"cell": [2, 3], "name": "Balance of Nature", "ranks": 1}]
    assert tree["added"] == [{"cell": [3, 3], "name": "Improved Serpent Sting", "ranks": 5}]


def test_rank_and_prerequisite_changes_are_listed_by_name():
    result = diff_snapshot(
        snapshot(
            [
                snap_talent(1, "Nature's Majesty", 1, 2, ranks=2, requires=((2, 1),)),
                snap_talent(2, "Nature's Splendor", 2, 2, ranks=3),
            ]
        ),
        records(
            [
                entry(11, "Nature's Majesty", 1, 2, max_rank=2),
                entry(12, "Nature's Splendor", 2, 2, max_rank=1, prereq=11, prereq_rank=2),
            ]
        ),
    )
    (tree,) = result["trees"]
    assert tree["rank_changes"] == [{"name": "Nature's Splendor", "snapshot": 3, "build": 1}]
    assert tree["prerequisite_changes"] == [
        {"name": "Nature's Majesty", "snapshot": ["Nature's Splendor", 1], "build": None},
        {"name": "Nature's Splendor", "snapshot": None, "build": ["Nature's Majesty", 2]},
    ]
```

- [ ] **Step 2: Run it to verify it fails**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_wowhead_diff.py -q --no-cov
```
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.wowhead_diff'`.

- [ ] **Step 3: Write the diff**

```python
# data/pipeline/wowhead_diff.py
"""What changed between the Wowhead pre-beta snapshot and the client's trees.

`data/raw-forever/wowhead-talents-2026-09-14.json` is the record of what was
believed about Forever's talents before the beta client shipped, and the
planner was built on it. The client now says; this writes down every place the
two disagree, keyed by the cell each talent sits in, so the difference is a
committed artefact rather than something a reader has to take on trust.

The snapshot keys its talents by `TalentTab` id, which is the same id the
emitted trees use, so no name matching is needed to pair up trees.
"""

from __future__ import annotations

import json
from collections.abc import Sequence
from pathlib import Path

from pipeline.models import ClassTalents, TalentEntry, TalentTree


def _prereq(talent: TalentEntry, by_id: dict[int, TalentEntry]) -> list | None:
    if talent.prereq_talent_id is None:
        return None
    prereq = by_id.get(talent.prereq_talent_id)
    return [prereq.name if prereq else str(talent.prereq_talent_id), talent.prereq_rank]


def _snapshot_prereq(talent: dict, tree: dict) -> list | None:
    requires = talent.get("requires") or []
    if not requires:
        return None
    first = requires[0]
    return [tree[str(first["id"])]["name"], first["qty"]]


def _tree_diff(snapshot_tree: dict, tree: TalentTree) -> dict:
    """Every disagreement inside one tree.

    Names are the identity, not cells: a talent that keeps its name and moves
    is a move, and a cell whose occupant's name is on neither side's shared
    list is a rename. What is left over on one side only is an addition or a
    removal.
    """
    snap_cell = {t["name"]: (t["row"], t["col"]) for t in snapshot_tree.values()}
    build_cell = {t.name: (t.tier, t.column) for t in tree.talents}
    only_snapshot = set(snap_cell) - set(build_cell)
    only_build = set(build_cell) - set(snap_cell)

    snap_by_cell = {(t["row"], t["col"]): t for t in snapshot_tree.values()}
    build_by_cell = {(t.tier, t.column): t for t in tree.talents}
    renamed = [
        {
            "cell": list(cell),
            "snapshot": snap_by_cell[cell]["name"],
            "build": build_by_cell[cell].name,
        }
        for cell in sorted(set(snap_by_cell) & set(build_by_cell))
        if snap_by_cell[cell]["name"] in only_snapshot and build_by_cell[cell].name in only_build
    ]
    was_renamed = {change["snapshot"] for change in renamed}
    now_renamed = {change["build"] for change in renamed}

    snap_by_name = {t["name"]: t for t in snapshot_tree.values()}
    build_by_name = {t.name: t for t in tree.talents}
    added = [
        {"cell": list(build_cell[name]), "name": name, "ranks": build_by_name[name].max_rank}
        for name in sorted(only_build - now_renamed)
    ]
    removed = [
        {"cell": list(snap_cell[name]), "name": name, "ranks": len(snap_by_name[name]["ranks"])}
        for name in sorted(only_snapshot - was_renamed)
    ]
    shared = sorted(set(snap_cell) & set(build_cell))
    moved = [
        {"name": name, "snapshot": list(snap_cell[name]), "build": list(build_cell[name])}
        for name in shared
        if snap_cell[name] != build_cell[name]
    ]

    by_id = {t.id: t for t in tree.talents}
    rank_changes, prerequisite_changes = [], []
    for name in shared:
        before, after = snap_by_name[name], build_by_name[name]
        if len(before["ranks"]) != after.max_rank:
            rank_changes.append(
                {"name": name, "snapshot": len(before["ranks"]), "build": after.max_rank}
            )
        was, now = _snapshot_prereq(before, snapshot_tree), _prereq(after, by_id)
        if was != now:
            prerequisite_changes.append({"name": name, "snapshot": was, "build": now})
    return {
        "tree_id": tree.id,
        "tree": tree.name,
        "added": added,
        "removed": removed,
        "renamed": renamed,
        "moved": moved,
        "rank_changes": rank_changes,
        "prerequisite_changes": prerequisite_changes,
    }


_CHANGE_KEYS = ("added", "removed", "renamed", "moved", "rank_changes", "prerequisite_changes")


def diff_snapshot(snapshot: dict, records: Sequence[ClassTalents]) -> dict:
    talents_by_tree = snapshot.get("talents") or {}
    trees = []
    build_total = 0
    for record in sorted(records, key=lambda r: r.class_id):
        for tree in sorted(record.trees, key=lambda t: t.position):
            build_total += len(tree.talents)
            snapshot_tree = talents_by_tree.get(str(tree.id))
            if snapshot_tree is None:
                trees.append(
                    {
                        "tree_id": tree.id,
                        "tree": tree.name,
                        "class_slug": record.class_slug,
                        "missing_from_snapshot": True,
                    }
                )
                continue
            diff = _tree_diff(snapshot_tree, tree)
            if any(diff[key] for key in _CHANGE_KEYS):
                trees.append({**diff, "class_slug": record.class_slug})
    return {
        "totals": {
            "snapshot": sum(len(t) for t in talents_by_tree.values()),
            "build": build_total,
        },
        "trees": trees,
    }


def write_snapshot_diff(
    snapshot: str,
    build: str,
    root: Path = Path("builds"),
    out: Path = Path("diffs"),
) -> Path:
    repo = Path(__file__).resolve().parents[2]
    payload = json.loads((repo / snapshot).read_text(encoding="utf-8"))
    records = [
        ClassTalents.model_validate_json(path.read_text(encoding="utf-8"))
        for path in sorted((root / build / "talents").glob("*.json"))
    ]
    result = {"snapshot": snapshot, "build": build, **diff_snapshot(payload, records)}
    out.mkdir(parents=True, exist_ok=True)
    name = Path(snapshot).stem.replace("wowhead-talents-", "wowhead-")
    path = out / f"{name}__{build}.json"
    path.write_text(json.dumps(result, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    counts = {key: sum(len(t.get(key, [])) for t in result["trees"]) for key in _CHANGE_KEYS}
    print(
        f"{result['totals']['snapshot']} snapshot talents, {result['totals']['build']} in "
        f"{build}: " + ", ".join(f"{key} {value}" for key, value in counts.items())
    )
    return path
```

- [ ] **Step 4: Run the diff test to verify it passes**

```bash
cd data && uv run pytest tests/test_wowhead_diff.py -q --no-cov
```
Expected: PASS, 5 passed.

- [ ] **Step 5: Share the flat talent rows**

In `data/pipeline/normalize/talents.py`, add at the end:

```python
def flat_talents(records: Sequence[ClassTalents]) -> list[TalentNode]:
    """The flat `talents.json` rows for already-built per-class records.

    `normalize_talents` above reads the legacy `Talent` table directly; a build
    whose trees came from the trait tables has no such table to read (the one
    it ships holds Classic Era's talents), so the flat list is folded back out
    of the per-class records instead. Both produce the same `TalentNode`.
    """
    return [
        TalentNode(
            id=talent.id,
            tab_id=tree.id,
            tab_name=tree.name,
            class_id=record.class_id,
            tier=talent.tier,
            column=talent.column,
            spell_ids=[rank.spell_id for rank in talent.ranks],
            prereq_talent_id=talent.prereq_talent_id,
        )
        for record in records
        for tree in record.trees
        for talent in tree.talents
    ]
```

and extend its imports:

```python
from collections.abc import Sequence

from pipeline.models import ClassTalents, TalentNode
```

- [ ] **Step 6: Pick the reader in normalize_build**

In `data/pipeline/normalize/__init__.py`, add to the local imports inside `normalize_build`:

```python
    from pipeline.curves import load_rank_points
    from pipeline.normalize.talents import flat_talents, normalize_talents
    from pipeline.normalize.trait_trees import build_trait_talent_trees
    from pipeline.normalize.traits import TraitRows, has_trait_trees
```

(and drop the now-duplicated `from pipeline.normalize.talents import normalize_talents` line).

Below `t = lambda name: ...`, add:

```python
    # A table a build's client simply does not have (see wago.OPTIONAL_TABLES)
    # is written as a header-only CSV by the fetch, but a build fetched before
    # that table was added to TABLES has no file at all. Both mean "no rows".
    def optional(name: str) -> list[dict[str, str]]:
        path = raw / f"{name}.csv"
        return read_csv(path) if path.exists() else []
```

Replace the `write_json(normalize_talents(...))` line and the `build_talent_trees` loop with:

```python
    trait_rows = TraitRows(
        skill_line=optional("SkillLine"),
        skill_line_x_trait_tree=optional("SkillLineXTraitTree"),
        talent_tab=t("TalentTab"),
        chr_classes=class_rows,
        node=optional("TraitNode"),
        node_entry=optional("TraitNodeEntry"),
        node_x_entry=optional("TraitNodeXTraitNodeEntry"),
        definition=optional("TraitDefinition"),
        edge=optional("TraitEdge"),
        cond=optional("TraitCond"),
        currency=optional("TraitCurrency"),
        node_group=optional("TraitNodeGroup"),
        node_group_x_node=optional("TraitNodeGroupXTraitNode"),
    )
```

immediately before the Phase 1 block (it needs `class_rows`, which is already read above), then inside the Phase 1 block replace

```python
    shutil.rmtree(build_dir / "talents", ignore_errors=True)
    for record in build_talent_trees(
        t("Talent"), t("TalentTab"), class_rows, spell_names, spell_text, icons, build
    ):
        write_model(record, build_dir / "talents" / f"{record.class_slug}.json")
```

with

```python
    # The 1.60 client keeps Forever's talents in the trait tables; the legacy
    # Talent table it still ships holds Classic Era's, so reading that here
    # would draw the wrong trees. Classic Era has no class trait trees at all.
    if has_trait_trees(trait_rows.skill_line_x_trait_tree):
        talent_records = build_trait_talent_trees(
            trait_rows,
            class_rows,
            spell_names,
            spell_text,
            load_rank_points(optional("TraitDefinitionEffectPoints"), optional("CurvePoint")),
            icons,
            build,
        )
        flat = flat_talents(talent_records)
    else:
        talent_records = build_talent_trees(
            t("Talent"), t("TalentTab"), class_rows, spell_names, spell_text, icons, build
        )
        flat = normalize_talents(t("Talent"), t("TalentTab"))
    write_json(flat, build_dir / "talents.json")
    shutil.rmtree(build_dir / "talents", ignore_errors=True)
    for record in talent_records:
        write_model(record, build_dir / "talents" / f"{record.class_slug}.json")
```

and delete the old `write_json(normalize_talents(...), build_dir / "talents.json")` line from the Phase 0 block.

- [ ] **Step 7: Keep the pre-beta build buildable**

In `data/pipeline/normalize/forever_talents.py`, add `spell_id=int(talent["id"])` as the last argument of the `TalentEntry(...)` construction, and give `normalize_forever_talents` a `tree_backgrounds: dict[int, str]` parameter beside `tree_names`, used as `background=tree_backgrounds[tree_id]` in the `TalentTree(...)` construction (raising `ForeverTalentError(f"tree {tree_id} has no background in our tables")` when absent, the way the name does).

In `data/pipeline/forever.py`, collect it in the same loop that builds `tree_class`/`tree_names`:

```python
            tree_backgrounds[tree["id"]] = tree["background"]
```

pass it through, and replace the inline `flat = [...]` comprehension with

```python
    from pipeline.normalize.talents import flat_talents

    flat = [node.model_dump() for node in flat_talents(per_class)]
```

- [ ] **Step 8: Add the CLI subcommand**

In `data/pipeline/__main__.py`, after the `diff` parser:

```python
    wd = sub.add_parser(
        "wowhead-diff", help="diff a build's talent trees against the Wowhead snapshot"
    )
    wd.add_argument(
        "--snapshot",
        default="data/raw-forever/wowhead-talents-2026-09-14.json",
        help="the saved Wowhead payload; see data/raw-forever/README.md",
    )
    wd.add_argument("--build", required=True)
```

and in `main`:

```python
    elif args.command == "wowhead-diff":
        from pipeline.wowhead_diff import write_snapshot_diff

        write_snapshot_diff(args.snapshot, args.build)
```

- [ ] **Step 9: Update the conformance keys and add the beta-build test**

In `data/tests/test_build_conformance.py`, add `"background"` to `TREE_KEYS` and `"spell_id"` to `TALENT_KEYS`.

```python
# data/tests/test_beta_build.py
"""The committed 1.60 build must be the trees measured off the client.

test_build_conformance.py checks the Classic Era build against the Phase 1
interface contract. This one checks the beta build against the facts recorded
in docs/superpowers/plans/2026-09-17-forever-trees.md: the numbers are the
measurement, so a regeneration that moves one has to be looked at.
"""

import json
from functools import cache
from pathlib import Path

BUILD = "1.60.1.69893"
BUILD_DIR = Path("builds") / BUILD

#: class slug -> the three tree names left to right, and the talent count in each.
TABS = {
    "warrior": [("Arms", 17), ("Fury", 18), ("Protection", 18)],
    "paladin": [("Holy", 18), ("Protection", 16), ("Retribution", 18)],
    "hunter": [("Beast Mastery", 16), ("Marksmanship", 17), ("Survival", 18)],
    "rogue": [("Assassination", 17), ("Combat", 17), ("Subtlety", 19)],
    "priest": [("Discipline", 18), ("Holy", 17), ("Shadow", 18)],
    "shaman": [("Elemental", 16), ("Enhancement", 18), ("Restoration", 16)],
    "mage": [("Arcane", 18), ("Fire", 17), ("Frost", 19)],
    "warlock": [("Affliction", 17), ("Demonology", 19), ("Destruction", 16)],
    "druid": [("Balance", 16), ("Feral Combat", 19), ("Restoration", 16)],
}
TALENT_TOTAL = 469
PREREQUISITE_TOTAL = 69
GRID_ROWS = 7
GRID_COLUMNS = 4


@cache
def talents(slug: str) -> dict:
    return json.loads((BUILD_DIR / "talents" / f"{slug}.json").read_text(encoding="utf-8"))


def test_every_class_has_its_three_tabs_in_the_games_order():
    for slug, want in TABS.items():
        trees = talents(slug)["trees"]
        assert [t["position"] for t in trees] == [0, 1, 2], slug
        assert [(t["name"], len(t["talents"])) for t in trees] == want, slug


def test_the_build_carries_every_talent_measured_off_the_client():
    total = sum(len(t["talents"]) for slug in TABS for t in talents(slug)["trees"])
    assert total == TALENT_TOTAL


def test_every_tree_names_a_background_and_every_talent_a_spell():
    for slug in TABS:
        for tree in talents(slug)["trees"]:
            assert tree["background"] == tree["background"].lower() and tree["background"]
            for talent in tree["talents"]:
                assert talent["spell_id"] > 0, (slug, talent["id"])
                assert talent["ranks"][0]["spell_id"] == talent["spell_id"]


def test_every_talent_sits_alone_on_the_four_by_seven_grid():
    for slug in TABS:
        for tree in talents(slug)["trees"]:
            cells = [(t["tier"], t["column"]) for t in tree["talents"]]
            assert len(set(cells)) == len(cells), (slug, tree["name"])
            for tier, column in cells:
                assert 0 <= tier < GRID_ROWS and 0 <= column < GRID_COLUMNS, (slug, tier, column)


def test_every_prerequisite_is_above_or_beside_its_dependent_in_the_same_tree():
    total = 0
    for slug in TABS:
        for tree in talents(slug)["trees"]:
            by_id = {t["id"]: t for t in tree["talents"]}
            for talent in tree["talents"]:
                if talent["prereq_talent_id"] is None:
                    assert talent["prereq_rank"] is None, (slug, talent["id"])
                    continue
                total += 1
                prereq = by_id[talent["prereq_talent_id"]]
                assert prereq["tier"] <= talent["tier"], (slug, talent["name"])
                assert talent["prereq_rank"] == prereq["max_rank"], (slug, talent["name"])
    assert total == PREREQUISITE_TOTAL
```

- [ ] **Step 10: Re-emit both builds**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run python -m pipeline fetch --product wow_classic_era --build 1.15.9.69722
```
Expected: prints `1.15.9.69722`; the five tables Era lacks are logged as `table X not found for build 1.15.9.69722; writing empty table`.

```bash
cd data && uv run python -m pipeline fetch --product wow_classic_beta --build 1.60.1.69893
```
Expected: prints `1.60.1.69893`, no warnings.

```bash
cd data && uv run python -m pipeline normalize --build 1.15.9.69722 && uv run python -m pipeline normalize --build 1.60.1.69893
```
Expected: exit 0 both times.

```bash
cd data && uv run python -m pipeline icons --build 1.60.1.69893 && uv run python -m pipeline normalize --build 1.60.1.69893
```
Expected: the new talents' icons are downloaded, then the manifest is rewritten to list them.

```bash
cd data && uv run python -m pipeline forever-talents --from-build 1.15.9.69722 --build forever-prebeta
```
Expected: exit 0 — the pre-beta build still validates against the two new model fields.

- [ ] **Step 11: Verify the Era build changed only by addition**

```bash
git diff --numstat data/builds/1.15.9.69722/ | awk '$2 != "0"'
```
Expected: **no output**. Any line printed is a deleted line in the Era build, which breaks the spec's "byte-identical except for added fields"; stop and fix rather than committing.

```bash
git diff data/builds/1.15.9.69722/talents/warrior.json | grep '^+' | grep -v '^+++' | grep -cv '"spell_id"\|"background"'
```
Expected: `0` — every added line is one of the two new fields.

- [ ] **Step 12: Write the snapshot diff**

```bash
cd data && uv run python -m pipeline wowhead-diff --build 1.60.1.69893
```
Expected: one line reading
`470 snapshot talents, 469 in 1.60.1.69893: added 1, removed 2, renamed 2, moved 4, rank_changes 0, prerequisite_changes 1`
and the file `data/diffs/wowhead-2026-09-14__1.60.1.69893.json`.

- [ ] **Step 13: Summarise the diff in the data README**

The spec asks for the diff to be summarised in the report as well as written
to `diffs/`. Add to `data/README.md`, under a new heading at the end:

```markdown
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
```

If the counts the `wowhead-diff` run printed differ from the ones above, the
client data has moved since this plan was written: write down what the run
actually printed, and say so to the controller rather than editing the run to
match the plan.

- [ ] **Step 14: Run the suites this touched**

```bash
cd data && uv run pytest tests/test_wowhead_diff.py tests/test_beta_build.py tests/test_build_conformance.py tests/test_normalize_build.py tests/test_normalize_talents.py tests/test_forever_talents.py -q --no-cov
```
Expected: PASS.

```bash
cd data && uv run ruff check pipeline tests && uv run pytest -q
```
Expected: `All checks passed!` and the whole suite green at >= 80% coverage.

- [ ] **Step 15: Commit**

```bash
git add data/pipeline data/tests data/builds data/diffs
```

```bash
git commit -F - <<'MSG'
feat(data): emit the beta build's real trees, and diff them against the snapshot

normalize_build now picks the trait reader for any build whose tables name
class trait trees and keeps the legacy reader for Classic Era, so both clients
stay supported from one code path. The Era build re-emits with nothing
removed: the only changes are the added spell_id and background fields.

1.60.1.69893 now carries Forever's own 469 talents in the game's tab order,
with the client's spell ids, grid cells, 69 prerequisites and per-rank text.
diffs/wowhead-2026-09-14__1.60.1.69893.json records every place the pre-beta
Wowhead snapshot and the client disagree.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 7: The tree background art

**Files:**
- Create: `data/pipeline/art.py`
- Create: `data/tests/test_art.py`
- Modify: `data/pipeline/__main__.py` (a `tree-art` subcommand)
- Modify: `data/README.md` (the step's place in the pipeline order)
- Regenerate: `data/builds/1.60.1.69893/trees/*.webp`, `data/builds/1.15.9.69722/trees/*.webp`, `data/builds/forever-prebeta/trees/*.webp`

**Interfaces:**
- Consumes: `pipeline.icons.CACHE_DIR`, `pipeline.icons._atomic_write`; `pipeline.wago.BASE_URL`, `USER_AGENT`.
- Produces:
  - `pipeline.art.ArtDataError(ValueError)`
  - `pipeline.art.BackgroundTreatment` — frozen: `saturation: float`, `brightness: float`, `tint: tuple[int, int, int]`, `tint_strength: float`.
  - `pipeline.art.TREATMENT: BackgroundTreatment` — the one place the look is defined.
  - `pipeline.art.BACKGROUND_SIZE: tuple[int, int] = (320, 384)`; `QUADRANTS: dict[str, tuple[int, int]]`; `QUADRANT_SIZES: dict[str, tuple[int, int]]`.
  - `pipeline.art.talent_frame_ids(manifest_rows: list[dict[str, str]]) -> dict[str, dict[str, int]]`
  - `pipeline.art.process_pixel(rgb: tuple[int, int, int], treatment: BackgroundTreatment = TREATMENT) -> tuple[int, int, int]`
  - `pipeline.art.compose_background(quadrants: Mapping[str, bytes]) -> Image.Image`
  - `pipeline.art.background_webp(quadrants: Mapping[str, bytes], treatment: BackgroundTreatment = TREATMENT) -> bytes`
  - `pipeline.art.backgrounds_for_build(build: str, root: Path = Path("builds"), cache_dir: Path = CACHE_DIR, client: httpx.Client | None = None) -> int`
  - `python -m pipeline tree-art --build <build>`
- Reads: each emitted `builds/<build>/talents/<slug>.json`'s `trees[].background`; writes `builds/<build>/trees/<background>.webp`.

- [ ] **Step 1: Write the failing test**

```python
# data/tests/test_art.py
import io

import pytest
from PIL import Image

from pipeline.art import (
    BACKGROUND_SIZE,
    QUADRANT_SIZES,
    TREATMENT,
    ArtDataError,
    background_webp,
    compose_background,
    process_pixel,
    talent_frame_ids,
)

COLOURS = {
    "TopLeft": (200, 40, 40),
    "TopRight": (40, 200, 40),
    "BottomLeft": (40, 40, 200),
    "BottomRight": (200, 200, 40),
}


def quadrant_bytes(sizes=None) -> dict[str, bytes]:
    sizes = sizes or QUADRANT_SIZES
    out = {}
    for name, colour in COLOURS.items():
        buffer = io.BytesIO()
        Image.new("RGB", sizes[name], colour).save(buffer, "PNG")
        out[name] = buffer.getvalue()
    return out


def test_the_four_quadrants_compose_into_one_320x384_image():
    image = compose_background(quadrant_bytes())
    assert image.size == BACKGROUND_SIZE
    assert image.getpixel((10, 10)) == COLOURS["TopLeft"]
    assert image.getpixel((300, 10)) == COLOURS["TopRight"]
    assert image.getpixel((10, 300)) == COLOURS["BottomLeft"]
    assert image.getpixel((300, 300)) == COLOURS["BottomRight"]


def test_a_quadrant_of_the_wrong_size_is_refused():
    sizes = {**QUADRANT_SIZES, "TopRight": (32, 256)}
    with pytest.raises(ArtDataError, match="TopRight"):
        compose_background(quadrant_bytes(sizes))


def test_a_missing_quadrant_is_refused():
    quadrants = quadrant_bytes()
    del quadrants["BottomRight"]
    with pytest.raises(ArtDataError, match="BottomRight"):
        compose_background(quadrants)


def test_the_treatment_desaturates_darkens_and_tints_toward_the_palette():
    before = (200, 40, 40)
    after = process_pixel(before)
    # Less saturated: the spread between the channels shrinks.
    assert (max(after) - min(after)) < (max(before) - min(before)) / 2
    # Darker: nothing survives at its old brightness.
    assert sum(after) < sum(before) / 2
    # And pulled most of the way to --color-raised.
    for channel, tint in zip(after, TREATMENT.tint, strict=True):
        assert abs(channel - tint) <= 30


def test_the_processed_image_matches_the_treatment_pixel_for_pixel():
    data = background_webp(quadrant_bytes())
    with Image.open(io.BytesIO(data)) as image:
        assert image.format == "WEBP"
        assert image.size == BACKGROUND_SIZE
        rgb = image.convert("RGB")
        for point, name in (
            ((10, 10), "TopLeft"),
            ((300, 10), "TopRight"),
            ((10, 300), "BottomLeft"),
            ((300, 300), "BottomRight"),
        ):
            want = process_pixel(COLOURS[name])
            got = rgb.getpixel(point)
            # Pillow's own blend and enhance round in C; WebP at quality 90 is
            # lossy. Two levels per channel is well inside both and still
            # pins the treatment to process_pixel, which is the definition.
            assert all(abs(g - w) <= 2 for g, w in zip(got, want, strict=True)), (name, got, want)


def test_talent_frame_ids_pairs_each_background_with_its_four_quadrants():
    rows = [
        {"ID": "136983", "FilePath": "Interface\\TALENTFRAME\\", "FileName": "WarriorArms-TopLeft.blp"},
        {"ID": "136984", "FilePath": "Interface\\TALENTFRAME\\", "FileName": "WarriorArms-TopRight.blp"},
        {"ID": "136981", "FilePath": "Interface\\TALENTFRAME\\", "FileName": "WarriorArms-BottomLeft.blp"},
        {"ID": "136982", "FilePath": "Interface\\TALENTFRAME\\", "FileName": "WarriorArms-BottomRight.blp"},
        {"ID": "136900", "FilePath": "Interface\\TALENTFRAME\\", "FileName": "MageArcane-TopLeft.blp"},
        {"ID": "132154", "FilePath": "Interface\\ICONS\\", "FileName": "Ability_GolemThunderClap.blp"},
    ]
    assert talent_frame_ids(rows) == {
        "warriorarms": {
            "TopLeft": 136983,
            "TopRight": 136984,
            "BottomLeft": 136981,
            "BottomRight": 136982,
        },
        "magearcane": {"TopLeft": 136900},
    }
```

- [ ] **Step 2: Run it to verify it fails**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run pytest tests/test_art.py -q --no-cov
```
Expected: FAIL — `ModuleNotFoundError: No module named 'pipeline.art'`.

- [ ] **Step 3: Write the art module**

```python
# data/pipeline/art.py
"""Each talent tree's background, taken from the game and made ours.

The client draws a tree's background as four textures under
`Interface\\TALENTFRAME\\`: a 256x256 top-left, a 64x256 top-right, a 256x128
bottom-left and a 64x128 bottom-right, which tile into one 320x384 panel. All
27 tabs of build 1.60.1.69893 have all four. They arrive the same way icons
do -- by file data id from CASC, cached as BLP -- and leave as one processed
WebP per tree at `builds/<build>/trees/<background>.webp`, which is what the
site ships. A raw texture is never published: Blizzard's art at full saturation
under this site's type would read as a screenshot of the game rather than as
this site, and the treatment below is the one place that decision lives.
"""

from __future__ import annotations

import io
import json
import logging
from collections.abc import Mapping
from dataclasses import dataclass
from pathlib import Path

import httpx
from PIL import Image, ImageEnhance

from pipeline.icons import CACHE_DIR, _atomic_write
from pipeline.wago import BASE_URL, USER_AGENT

logger = logging.getLogger(__name__)


class ArtDataError(ValueError):
    """The client's talent-frame textures are not the shape a panel needs."""


#: Quadrant name -> where its top-left corner goes in the composite.
QUADRANTS: dict[str, tuple[int, int]] = {
    "TopLeft": (0, 0),
    "TopRight": (256, 0),
    "BottomLeft": (0, 256),
    "BottomRight": (256, 256),
}
#: Quadrant name -> the size the client ships it at.
QUADRANT_SIZES: dict[str, tuple[int, int]] = {
    "TopLeft": (256, 256),
    "TopRight": (64, 256),
    "BottomLeft": (256, 128),
    "BottomRight": (64, 128),
}
BACKGROUND_SIZE = (320, 384)
_TALENT_FRAME_DIR = "interface\\talentframe\\"
_BLP_SUFFIX = ".blp"


@dataclass(frozen=True)
class BackgroundTreatment:
    """How a talent-frame texture becomes a Forever Sixty tree background.

    Applied in this order: pull the colour out, take the light down, then
    blend toward one palette token so every tree sits on the same ground and
    the gold borders and white type stay the brightest things on the panel.
    """

    #: 0 leaves a greyscale image, 1 leaves the texture's own colour.
    saturation: float
    #: Multiplies the remaining light. 1 leaves it alone.
    brightness: float
    #: The colour everything is blended toward: --color-raised in
    #: web/src/styles/tokens.css, the panel the tree grid sits on.
    tint: tuple[int, int, int]
    #: How far toward `tint`. 0 leaves the darkened texture, 1 leaves a flat fill.
    tint_strength: float


TREATMENT = BackgroundTreatment(
    saturation=0.30,
    brightness=0.45,
    tint=(0x0D, 0x11, 0x1A),
    tint_strength=0.55,
)


def process_pixel(
    rgb: tuple[int, int, int], treatment: BackgroundTreatment = TREATMENT
) -> tuple[int, int, int]:
    """What `process_background` does to one pixel. The treatment's definition."""
    red, green, blue = rgb
    # ITU-R 601-2 luma, the transform Image.convert("L") uses.
    grey = red * 299 / 1000 + green * 587 / 1000 + blue * 114 / 1000
    out = []
    for channel, tint in zip(rgb, treatment.tint, strict=True):
        desaturated = grey + (channel - grey) * treatment.saturation
        darkened = desaturated * treatment.brightness
        out.append(round(darkened + (tint - darkened) * treatment.tint_strength))
    return (out[0], out[1], out[2])


def compose_background(quadrants: Mapping[str, bytes]) -> Image.Image:
    """The four textures tiled into one 320x384 RGB image.

    Alpha is dropped rather than composited: three of the four quadrants carry
    an alpha channel, the panel behind them is opaque in game, and pasting
    with a mask would leave the site's page colour showing through the edges
    of every tree.
    """
    panel = Image.new("RGB", BACKGROUND_SIZE, (0, 0, 0))
    for name, offset in QUADRANTS.items():
        data = quadrants.get(name)
        if data is None:
            raise ArtDataError(f"no {name} texture; a tree background needs all four")
        with Image.open(io.BytesIO(data)) as tile:
            if tile.size != QUADRANT_SIZES[name]:
                raise ArtDataError(
                    f"{name} is {tile.size}, not the {QUADRANT_SIZES[name]} it tiles at"
                )
            panel.paste(tile.convert("RGB"), offset)
    return panel


def process_background(
    image: Image.Image, treatment: BackgroundTreatment = TREATMENT
) -> Image.Image:
    """`process_pixel` over a whole image, through Pillow's own C loops."""
    rgb = image.convert("RGB")
    desaturated = ImageEnhance.Color(rgb).enhance(treatment.saturation)
    darkened = ImageEnhance.Brightness(desaturated).enhance(treatment.brightness)
    tint = Image.new("RGB", darkened.size, treatment.tint)
    return Image.blend(darkened, tint, treatment.tint_strength)


def background_webp(
    quadrants: Mapping[str, bytes], treatment: BackgroundTreatment = TREATMENT
) -> bytes:
    processed = process_background(compose_background(quadrants), treatment)
    buffer = io.BytesIO()
    processed.save(buffer, "WEBP", quality=90, method=6)
    return buffer.getvalue()


def talent_frame_ids(manifest_rows: list[dict[str, str]]) -> dict[str, dict[str, int]]:
    """Background name (lowercased) -> quadrant name -> file data id."""
    out: dict[str, dict[str, int]] = {}
    for row in manifest_rows:
        if row["FilePath"].lower() != _TALENT_FRAME_DIR:
            continue
        name = row["FileName"]
        if not name.lower().endswith(_BLP_SUFFIX):
            continue
        stem = name[: -len(_BLP_SUFFIX)]
        background, _, quadrant = stem.rpartition("-")
        if quadrant not in QUADRANTS:
            continue
        out.setdefault(background.lower(), {})[quadrant] = int(row["ID"])
    return out


def _wanted_backgrounds(build_dir: Path) -> set[str]:
    """Every `background` the build's emitted trees name."""
    names: set[str] = set()
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for tree in payload["trees"]:
            names.add(tree["background"])
    return names


def backgrounds_for_build(
    build: str,
    root: Path = Path("builds"),
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
) -> int:
    """Write one processed background per tree. Returns the number written."""
    from pipeline.csvio import read_csv

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    frames = talent_frame_ids(read_csv(raw / "ManifestInterfaceData.csv"))
    wanted = sorted(_wanted_backgrounds(build_dir))
    out_dir = build_dir / "trees"
    out_dir.mkdir(parents=True, exist_ok=True)
    cache_dir.mkdir(parents=True, exist_ok=True)
    own = client is None
    if client is None:
        client = httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    written = 0
    try:
        for background in wanted:
            target = out_dir / f"{background}.webp"
            if target.exists():
                continue
            ids = frames.get(background)
            if not ids or set(ids) != set(QUADRANTS):
                raise ArtDataError(
                    f"{background} has {sorted(ids or ())} in ManifestInterfaceData, "
                    f"not the four quadrants {sorted(QUADRANTS)}"
                )
            quadrants: dict[str, bytes] = {}
            for quadrant, file_id in ids.items():
                cached = cache_dir / f"{file_id}.blp"
                if not cached.exists():
                    response = client.get(f"/api/casc/{file_id}", timeout=60)
                    response.raise_for_status()
                    _atomic_write(cached, response.content)
                quadrants[quadrant] = cached.read_bytes()
            _atomic_write(target, background_webp(quadrants))
            written += 1
    finally:
        if own:
            client.close()
    print(f"{written} tree backgrounds written, {len(wanted)} referenced")
    return written
```

- [ ] **Step 4: Run the test to verify it passes**

```bash
cd data && uv run pytest tests/test_art.py -q --no-cov
```
Expected: PASS, 6 passed.

- [ ] **Step 5: Add the subcommand**

In `data/pipeline/__main__.py`, after the `icons` parser:

```python
    a = sub.add_parser("tree-art", help="download and process each talent tree's background")
    a.add_argument("--build", required=True)
```

and in `main`, after the `icons` branch:

```python
    elif args.command == "tree-art":
        from pipeline.art import backgrounds_for_build

        backgrounds_for_build(args.build)
```

- [ ] **Step 6: Generate the art for all three builds**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && for B in 1.15.9.69722 1.60.1.69893; do uv run python -m pipeline tree-art --build "$B"; uv run python -m pipeline normalize --build "$B"; done
```
Expected: `27 tree backgrounds written, 27 referenced` for each (the second build reuses the BLP cache, so it is fast), then two clean normalizes that fold `trees/` into each manifest.

The pre-beta build has no `raw/` of its own; copy the Era art into it the way `forever.py` already copies icons. In `data/pipeline/forever.py`, beside the existing `icons` copy:

```python
    if (src / "trees").is_dir() and not (dst / "trees").exists():
        shutil.copytree(src / "trees", dst / "trees")
```

```bash
cd data && uv run python -m pipeline forever-talents --from-build 1.15.9.69722 --build forever-prebeta
```
Expected: exit 0, and `ls builds/forever-prebeta/trees | wc -l` prints `27`.

- [ ] **Step 7: Check the weight**

```bash
cd data && du -sh builds/1.60.1.69893/trees && ls -S builds/1.60.1.69893/trees | head -1 | xargs -I{} du -h builds/1.60.1.69893/trees/{}
```
Expected: the directory under 1.5 MB and no single file over 80 KB. If either is over, raise `quality` to 80 in `background_webp` and regenerate rather than changing the treatment.

- [ ] **Step 8: Note the new step in the pipeline order**

In `data/README.md`, add `tree-art` to the run order beside `icons`, as:

```markdown
4. `uv run python -m pipeline tree-art --build <build>` — one processed background per talent tree, from the client's own `Interface\TALENTFRAME\` textures. Re-run `normalize` afterwards so the manifest lists them.
```

(renumbering the steps after it).

- [ ] **Step 9: Lint and run the suite**

```bash
cd data && uv run ruff check pipeline tests && uv run pytest -q
```
Expected: `All checks passed!` and green.

- [ ] **Step 10: Commit**

```bash
git add data/pipeline data/tests/test_art.py data/README.md data/builds
```

```bash
git commit -F - <<'MSG'
feat(data): one processed background per talent tree

Each tab's four Interface\TALENTFRAME textures tile into a 320x384 panel; the
panel is desaturated, darkened and blended toward --color-raised by one named
constant so the trees read as this site rather than as a screenshot of the
game. process_pixel is the treatment's definition and the test holds the image
to it pixel for pixel.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 8: The API reads the new fields, and resolves a talent by spell id

**Files:**
- Modify: `api/internal/trees/trees.go:22-40` (`Talent`, `Tree`), `:105-135` (`Build`), `:300-330` (`loadBuild`'s talent loop)
- Modify: `api/internal/trees/trees_test.go`
- Modify: `api/internal/trees/testdata/builds/test-1/talents/warrior.json`, `paladin.json`
- Modify: `api/internal/spec/spec.go:65-92` (`Split`)
- Modify: `api/internal/spec/spec_test.go`

**Interfaces:**
- Consumes: the emitted `talents/<slug>.json` from Task 6, which now carries `spell_id` per talent and `background` per tree.
- Produces:
  - `trees.Talent.SpellID int` with tag `json:"spell_id"`.
  - `trees.Tree.Background string` with tag `json:"background"`.
  - `func (b *Build) TalentBySpellID(classID, spellID int) (TalentRef, bool)`.
  - `spec.Inferrer.Split` falls back to `TalentBySpellID` for ids that are not talent ids.

Why the spell-id index: a combat log's `COMBATANT_INFO` talents field carries what the client writes, and on the 1.60 client that is the talent's spell id, not the trait node id the emitted `id` field uses. `Split` is the one place in the API that reads log talent ids, so it is the one place that needs both.

- [ ] **Step 1: Write the failing tests**

Append to `api/internal/trees/trees_test.go`:

```go
func TestTalentCarriesItsSpellIDAndTreeItsBackground(t *testing.T) {
	data, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, ok := data.Build("test-1")
	if !ok {
		t.Fatal("fixture build test-1 did not load")
	}
	arms := b.Trees(1)[0]
	if arms.Background != "warriorarms" {
		t.Fatalf("Arms background = %q, want %q", arms.Background, "warriorarms")
	}
	ref, ok := b.Talent(1, 101)
	if !ok {
		t.Fatal("talent 101 is missing")
	}
	if ref.SpellID != 12281 {
		t.Fatalf("talent 101 spell id = %d, want 12281", ref.SpellID)
	}
}

func TestTalentBySpellIDFindsTheTalentTheClientWrites(t *testing.T) {
	data, err := LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := data.Build("test-1")
	ref, ok := b.TalentBySpellID(1, 12281)
	if !ok || ref.ID != 101 {
		t.Fatalf("TalentBySpellID(1, 12281) = %+v, %v; want talent 101", ref, ok)
	}
	if _, ok := b.TalentBySpellID(1, 999999); ok {
		t.Fatal("an unknown spell id must report false")
	}
	if _, ok := b.TalentBySpellID(99, 12281); ok {
		t.Fatal("an unknown class must report false")
	}
}
```

Append to `api/internal/spec/spec_test.go`:

```go
func TestSplitReadsTalentsWrittenAsSpellIDs(t *testing.T) {
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := data.Build("test-1")
	i := New(b)
	// 12281 is talent 101's spell, in Arms; 20000 is talent 201's, in Fury.
	got := i.Split("Warrior", []int64{12281, 12281, 20000})
	want := []int{2, 1}
	if len(got) != len(want) {
		t.Fatalf("Split = %v, want %v", got, want)
	}
	for n := range want {
		if got[n] != want[n] {
			t.Fatalf("Split = %v, want %v", got, want)
		}
	}
}
```

- [ ] **Step 2: Run them to verify they fail**

```bash
cd api && go test ./internal/trees/ ./internal/spec/
```
Expected: FAIL — `ref.SpellID undefined`, `arms.Background undefined`, `b.TalentBySpellID undefined`.

- [ ] **Step 3: Give the fixture the new fields**

In `api/internal/trees/testdata/builds/test-1/talents/warrior.json`, add `"background": "warriorarms"` to the Arms tree object and `"background": "warriorfury"` to the Fury tree object, and add a `"spell_id"` to each talent equal to its first rank's `spell_id` (101 -> 12281, 102 -> 16462, 103 -> 12286, 104 -> 12834, and the same rule for every talent in the file, Fury's 201 -> 20000 included). Do the same in `paladin.json`, using that file's own tree names for the backgrounds (`"paladinholy"`, and so on for each tree present).

- [ ] **Step 4: Add the fields and the index**

In `api/internal/trees/trees.go`:

```go
type Talent struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	MaxRank        int    `json:"max_rank"`
	Tier           int    `json:"tier"`
	Column         int    `json:"column"`
	PrereqTalentID *int   `json:"prereq_talent_id"`
	PrereqRank     *int   `json:"prereq_rank"`
	Ranks          []Rank `json:"ranks"`
	// SpellID is the spell the client writes when the talent is learned,
	// which is what a combat log's COMBATANT_INFO carries. The ID above is
	// the client's trait node id, and the two are different numbers.
	SpellID int `json:"spell_id"`
}

type Tree struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position int      `json:"position"`
	Talents  []Talent `json:"talents"`
	// Background names the processed art at /data/<build>/trees/<background>.webp.
	Background string `json:"background"`
}
```

Add a field to `Build`, beside `talents`:

```go
	talentsBySpell map[int]map[int]TalentRef
```

initialise it in `loadBuild`'s `b := &Build{…}` literal:

```go
		talentsBySpell: map[int]map[int]TalentRef{},
```

fill it in the same loop that fills `byTalentID`, right after the `byTalentID[t.ID] = …` line:

```go
				bySpellID[t.SpellID] = byTalentID[t.ID]
```

with `bySpellID := map[int]TalentRef{}` declared beside `byTalentID := map[int]TalentRef{}`, and stored beside it:

```go
		b.talentsBySpell[c.ID] = bySpellID
```

And add the accessor beside `Talent`:

```go
// TalentBySpellID resolves the spell id the client writes to the talent it
// belongs to. A build emitted before talents carried a spell id indexes every
// talent under 0, so a lookup for 0 is refused rather than answering with an
// arbitrary talent.
func (b *Build) TalentBySpellID(classID, spellID int) (TalentRef, bool) {
	if spellID == 0 {
		return TalentRef{}, false
	}
	bySpell, ok := b.talentsBySpell[classID]
	if !ok {
		return TalentRef{}, false
	}
	t, ok := bySpell[spellID]
	return t, ok
}
```

- [ ] **Step 5: Let Split read either id**

In `api/internal/spec/spec.go`, replace the loop body inside `Split`:

```go
	for _, id := range talents {
		ref, ok := i.build.Talent(classID, int(id))
		if !ok {
			// The 1.60 client writes the talent's spell id, not the trait
			// node id the emitted data is keyed by. Both are tried, in that
			// order, so a log from either client reads correctly.
			ref, ok = i.build.TalentBySpellID(classID, int(id))
		}
		if !ok {
			continue
		}
		if n, ok := position[ref.TreeID]; ok {
			split[n]++
		}
	}
```

- [ ] **Step 6: Make Latest prefer a real client build**

`Data.Latest()` is what `spec.Split` reads a log against. `compareVersions`
compares segment by segment and falls back to string order for a segment that
is not a number, which ranks `forever-prebeta` above `1.60.1.69893` — so the
rankings would name specs from the pre-beta snapshot's trees for as long as
that build is published. Add the test first, to `api/internal/trees/trees_test.go`:

```go
func TestLatestPrefersAClientBuildOverANamedDataSet(t *testing.T) {
	for _, c := range []struct {
		versions []string
		want     string
	}{
		{[]string{"1.15.9.69722", "1.60.1.69893", "forever-prebeta"}, "1.60.1.69893"},
		{[]string{"1.15.9.69722", "1.9.1.1"}, "1.15.9.69722"},
		{[]string{"forever-prebeta"}, "forever-prebeta"},
	} {
		if got := newestVersion(c.versions); got != c.want {
			t.Fatalf("newestVersion(%v) = %q, want %q", c.versions, got, c.want)
		}
	}
}
```

Run it:

```bash
cd api && go test ./internal/trees/ -run TestLatestPrefers
```
Expected: FAIL — `newestVersion([1.15.9.69722 1.60.1.69893 forever-prebeta]) = "forever-prebeta", want "1.60.1.69893"`.

Then, in `api/internal/trees/trees.go`, add beside `compareVersions`:

```go
// isClientBuild reports whether a version looks like a client build string
// ("1.60.1.69893") rather than a named data set ("forever-prebeta"). Only
// client builds are comparable as version numbers, and Latest must not hand
// the rankings a data set that is not the newest client data.
func isClientBuild(version string) bool {
	segments := strings.Split(version, ".")
	if len(segments) < 2 {
		return false
	}
	for _, segment := range segments {
		if _, err := strconv.Atoi(segment); err != nil {
			return false
		}
	}
	return true
}
```

and make `newestVersion` prefer them:

```go
func newestVersion(versions []string) string {
	best := versions[0]
	for _, v := range versions[1:] {
		bestIsBuild, vIsBuild := isClientBuild(best), isClientBuild(v)
		if bestIsBuild != vIsBuild {
			if vIsBuild {
				best = v
			}
			continue
		}
		if compareVersions(v, best) > 0 {
			best = v
		}
	}
	return best
}
```

```bash
cd api && go test ./internal/trees/ -run TestLatest
```
Expected: PASS.

- [ ] **Step 7: Run the tests to verify they pass**

```bash
cd api && go test ./internal/trees/ ./internal/spec/ ./internal/builds/
```
Expected: PASS.

- [ ] **Step 8: Check the API loads the real beta build**

```bash
cd api && mkdir -p cmd/loadcheck && cat > cmd/loadcheck/main.go <<'GO'
package main

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func main() {
	d, err := trees.Load("../data/builds")
	if err != nil {
		panic(err)
	}
	fmt.Println("versions:", d.Versions())
	fmt.Println("skipped:", d.Skipped())
	b, ok := d.Build("1.60.1.69893")
	if !ok {
		panic("1.60.1.69893 did not load")
	}
	total := 0
	for _, c := range b.Classes() {
		for _, t := range b.Trees(c.ID) {
			total += len(t.Talents)
		}
	}
	arms := b.Trees(1)[0]
	first := arms.Talents[0]
	fmt.Println("talents:", total)
	fmt.Println("Arms:", arms.Name, arms.Background)
	fmt.Println("first Arms talent:", first.Name, first.SpellID)
	ref, found := b.TalentBySpellID(1, first.SpellID)
	fmt.Println("by spell id:", ref.Name, found)
}
GO
go run ./cmd/loadcheck
```
Expected:
```
versions: [1.15.9.69722 1.60.1.69893 forever-prebeta]
skipped: []
talents: 469
Arms: Arms warriorarms
first Arms talent: <a real talent name> <a non-zero spell id>
by spell id: <the same name> true
```

```bash
cd api && rm -rf cmd/loadcheck
```

- [ ] **Step 9: Vet and commit**

```bash
cd api && go vet ./... && go test ./...
```
Expected: no output from vet, all packages pass.

```bash
git add api
```

```bash
git commit -F - <<'MSG'
feat(api): carry the talent's spell id and the tree's background

The emitted trees gained two fields. Talent.SpellID is what a combat log's
COMBATANT_INFO carries on the 1.60 client, which is not the trait node id the
data is keyed by, so spec.Split now tries the talent id and then the spell id
and reads a log from either client correctly.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 9: Connector geometry, and the rules audited for a dependent's need

**Files:**
- Create: `web/src/lib/planner/connectors.ts`
- Create: `web/src/lib/planner/connectors.test.ts`
- Modify: `web/src/lib/planner/types.ts:22-42` (`Talent`, `TalentTree`)
- Modify: `web/src/lib/planner/rules.test.ts`
- Modify: `web/src/fixtures/planner/talents/warrior.json`
- Modify: `web/src/fixtures/planner/fixture.test.ts:51-70`

**Interfaces:**
- Consumes: `Talent`, `TalentTree` from `./types`.
- Produces:
  - `web/src/lib/planner/types.ts`: `Talent.spell_id: number`, `TalentTree.background: string`.
  - `connectors.ts`: `CELL_PX = 44`, `GAP_PX = 8`, `PITCH_PX = 52`, `HALF_PX = 22`; `interface Cell { tier: number; column: number }`; `interface Connector { id: string; from: number; to: number; d: string; met: boolean }`; `cellCentre(cell: Cell): { x: number; y: number }`; `connectorPath(from: Cell, to: Cell): string`; `gridViewBox(tiers: number, columns: number): string`; `connectorsFor(tree: TalentTree, ranks: Map<number, number>): Connector[]`.

- [ ] **Step 1: Write the failing geometry test**

```ts
// web/src/lib/planner/connectors.test.ts
import { describe, expect, it } from 'vitest';
import { cellCentre, connectorPath, connectorsFor, gridViewBox, PITCH_PX } from './connectors';
import type { Talent, TalentTree } from './types';

function talent(overrides: Partial<Talent> & Pick<Talent, 'id' | 'tier' | 'column'>): Talent {
  return {
    name: `Talent ${overrides.id}`,
    icon: 'icon',
    max_rank: 5,
    prereq_talent_id: null,
    prereq_rank: null,
    ranks: [],
    spell_id: 1,
    ...overrides,
  };
}

function tree(talents: Talent[]): TalentTree {
  return { id: 161, name: 'Arms', position: 0, background: 'warriorarms', talents };
}

describe('cellCentre', () => {
  it('puts row 0 column 0 half a cell in, and steps by the pitch', () => {
    expect(cellCentre({ tier: 0, column: 0 })).toEqual({ x: 22, y: 22 });
    expect(cellCentre({ tier: 1, column: 1 })).toEqual({ x: 22 + PITCH_PX, y: 22 + PITCH_PX });
    expect(cellCentre({ tier: 6, column: 3 })).toEqual({ x: 178, y: 334 });
  });
});

describe('connectorPath', () => {
  it('runs straight down the column, edge to edge', () => {
    // 67 of the beta build's 69 prerequisites are this shape.
    expect(connectorPath({ tier: 0, column: 1 }, { tier: 1, column: 1 })).toBe('M 74 44 V 52');
  });

  it('spans more than one row when the prerequisite is further up', () => {
    expect(connectorPath({ tier: 0, column: 0 }, { tier: 2, column: 0 })).toBe('M 22 44 V 104');
  });

  it('runs straight across the row to the right', () => {
    // Priest: Mind Flay (2,2) -> Improved Mind Flay (2,3).
    expect(connectorPath({ tier: 2, column: 2 }, { tier: 2, column: 3 })).toBe('M 148 126 H 156');
  });

  it('runs straight across the row to the left', () => {
    // Paladin: Holy Shock (4,1) -> Divine Precision (4,0).
    expect(connectorPath({ tier: 4, column: 1 }, { tier: 4, column: 0 })).toBe('M 52 230 H 44');
  });

  it('elbows down the prerequisite column and then across', () => {
    // No prerequisite in build 1.60.1.69893 needs this, but the shape is the
    // game's and a future build that adds one must not draw a wrong line.
    expect(connectorPath({ tier: 0, column: 0 }, { tier: 1, column: 1 })).toBe('M 22 44 V 74 H 52');
    expect(connectorPath({ tier: 0, column: 2 }, { tier: 1, column: 1 })).toBe('M 126 44 V 74 H 96');
  });
});

describe('gridViewBox', () => {
  it('is the grid with no trailing gap', () => {
    expect(gridViewBox(7, 4)).toBe('0 0 200 356');
    expect(gridViewBox(4, 2)).toBe('0 0 96 200');
  });
});

describe('connectorsFor', () => {
  const arms = tree([
    talent({ id: 1001, tier: 0, column: 1 }),
    talent({ id: 1002, tier: 1, column: 1, prereq_talent_id: 1001, prereq_rank: 5 }),
    talent({ id: 1003, tier: 2, column: 0 }),
  ]);

  it('draws one connector per prerequisite, unmet until the rank is reached', () => {
    expect(connectorsFor(arms, new Map())).toEqual([
      { id: '1001-1002', from: 1001, to: 1002, d: 'M 74 44 V 52', met: false },
    ]);
  });

  it('marks a connector met once the prerequisite holds the rank it needs', () => {
    expect(connectorsFor(arms, new Map([[1001, 4]]))[0].met).toBe(false);
    expect(connectorsFor(arms, new Map([[1001, 5]]))[0].met).toBe(true);
    expect(connectorsFor(arms, new Map([[1001, 5]]))[0].id).toBe('1001-1002');
  });

  it('ignores a prerequisite that is not in the tree', () => {
    const broken = tree([talent({ id: 1, tier: 1, column: 0, prereq_talent_id: 99, prereq_rank: 1 })]);
    expect(connectorsFor(broken, new Map())).toEqual([]);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

```bash
cd web && export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH" && npx vitest run src/lib/planner/connectors.test.ts
```
Expected: FAIL — cannot resolve `./connectors`.

- [ ] **Step 3: Add the two fields to the planner's types**

In `web/src/lib/planner/types.ts`:

```ts
export interface Talent {
  id: number;
  name: string;
  icon: string;
  max_rank: number;
  /** 0-based row. */
  tier: number;
  /** 0-based column. */
  column: number;
  prereq_talent_id: number | null;
  prereq_rank: number | null;
  /** Exactly `max_rank` entries. */
  ranks: TalentRank[];
  /** The spell the client writes when the talent is learned. */
  spell_id: number;
}

export interface TalentTree {
  id: number;
  name: string;
  /** Left-to-right order as in the client. */
  position: number;
  talents: Talent[];
  /** Names the tree's art at `/data/<build>/trees/<background>.webp`. */
  background: string;
}
```

- [ ] **Step 4: Write the geometry**

```ts
// web/src/lib/planner/connectors.ts
// Where the game draws a line from a prerequisite to the talent that needs it, as
// a path in the tree grid's own pixel space. Pure and unit-tested: TreeGrid puts
// these in an SVG overlay sized to the grid, so nothing here measures the DOM.
//
// Three shapes, which is what the client draws: straight down a column, straight
// across a row, and an elbow -- down the prerequisite's column to the dependent's
// row, then across into it. Build 1.60.1.69893 uses the first for 67 of its 69
// prerequisites and the second for the other two.
//
// Every line starts and ends at a cell's edge rather than its centre, so no line
// crosses an icon.
import type { TalentTree } from './types';

/** The talent button's side in CSS pixels; TalentCell is h-11 w-11. */
export const CELL_PX = 44;
/** The grid's gap; TreeGrid is gap-2. */
export const GAP_PX = 8;
/** Centre-to-centre distance between neighbouring cells. */
export const PITCH_PX = CELL_PX + GAP_PX;
export const HALF_PX = CELL_PX / 2;

export interface Cell {
  tier: number;
  column: number;
}

export interface Connector {
  /** Stable key for the `{#each}` that draws these. */
  id: string;
  /** The prerequisite's talent id. */
  from: number;
  /** The dependent's talent id. */
  to: number;
  /** An SVG path, in the coordinate space `gridViewBox` describes. */
  d: string;
  /** True once the prerequisite holds the rank the dependent needs. */
  met: boolean;
}

export function cellCentre({ tier, column }: Cell): { x: number; y: number } {
  return { x: column * PITCH_PX + HALF_PX, y: tier * PITCH_PX + HALF_PX };
}

export function connectorPath(from: Cell, to: Cell): string {
  const a = cellCentre(from);
  const b = cellCentre(to);
  if (from.column === to.column) {
    return `M ${a.x} ${a.y + HALF_PX} V ${b.y - HALF_PX}`;
  }
  const side = to.column > from.column ? 1 : -1;
  if (from.tier === to.tier) {
    return `M ${a.x + side * HALF_PX} ${a.y} H ${b.x - side * HALF_PX}`;
  }
  return `M ${a.x} ${a.y + HALF_PX} V ${b.y} H ${b.x - side * HALF_PX}`;
}

/** The viewBox for an overlay covering a grid of this many tiers and columns. */
export function gridViewBox(tiers: number, columns: number): string {
  return `0 0 ${columns * PITCH_PX - GAP_PX} ${tiers * PITCH_PX - GAP_PX}`;
}

export function connectorsFor(tree: TalentTree, ranks: Map<number, number>): Connector[] {
  const byId = new Map(tree.talents.map((talent) => [talent.id, talent]));
  const connectors: Connector[] = [];
  for (const talent of tree.talents) {
    if (talent.prereq_talent_id === null) continue;
    const prereq = byId.get(talent.prereq_talent_id);
    if (prereq === undefined) continue;
    connectors.push({
      id: `${prereq.id}-${talent.id}`,
      from: prereq.id,
      to: talent.id,
      d: connectorPath(prereq, talent),
      met: (ranks.get(prereq.id) ?? 0) >= (talent.prereq_rank ?? 0),
    });
  }
  return connectors.sort((a, b) => a.from - b.from || a.to - b.to);
}
```

- [ ] **Step 5: Run the geometry test to verify it passes**

```bash
cd web && npx vitest run src/lib/planner/connectors.test.ts
```
Expected: PASS, 11 tests.

- [ ] **Step 6: Give the fixture the two new fields**

In `web/src/fixtures/planner/talents/warrior.json`, add `"background": "fixture_arms"` to the Arms tree object and `"background": "fixture_fury"` to the Fury tree object, and add `"spell_id"` to every talent equal to its first rank's `spell_id` (1001 -> 10011, 1002 -> 10021, 1003 -> 10031, 1004 -> 10041, 1005 -> 10051, 1006 -> 10061, 1007 -> 10071, 2001 -> 20011, 2002 -> 20021, 2003 -> 20031, 2004 -> 20041, 2005 -> 20051, 2006 -> 20061, 2007 -> 20071).

In `web/src/fixtures/planner/fixture.test.ts`, replace the prerequisite assertion's tier check and add the new fields:

```ts
  it('points every prerequisite at an earlier or equal tier in the same tree', () => {
    for (const tree of talents.trees) {
      const byId = new Map(tree.talents.map((t) => [t.id, t]));
      for (const talent of tree.talents) {
        if (talent.prereq_talent_id === null) {
          expect(talent.prereq_rank).toBeNull();
          continue;
        }
        const prereq = byId.get(talent.prereq_talent_id);
        expect(prereq).toBeDefined();
        // The client has two same-row prerequisites (Priest Mind Flay ->
        // Improved Mind Flay, Paladin Holy Shock -> Divine Precision), so a
        // prerequisite is above or beside its dependent, never below.
        expect(prereq!.tier).toBeLessThanOrEqual(talent.tier);
        expect(talent.prereq_rank).toBeGreaterThan(0);
```

and add a test beside it:

```ts
  it('names a background per tree and the client spell per talent', () => {
    for (const tree of talents.trees) {
      expect(tree.background).toMatch(/^[a-z0-9_]+$/);
      for (const talent of tree.talents) {
        expect(talent.spell_id).toBe(talent.ranks[0].spell_id);
      }
    }
  });
```

- [ ] **Step 7: Audit the prerequisite-drop rule**

Append to `web/src/lib/planner/rules.test.ts`, inside the `describe` that already covers `canRemovePoint`:

```ts
  it('refuses dropping a prerequisite below the rank a dependent still needs', () => {
    const index = indexTalents(talents);
    // Deflection to rank 2 (what Tactical Mastery needs), then one point in it.
    const order = [1002, 1002, 1004];
    expect(canRemovePoint(index, order, 1002)).toEqual({
      ok: false,
      reason: messages.prereqMissing('Tactical Mastery', 2, 'Deflection'),
    });
    // With Deflection at 3 there is a point to spare, so it comes back out.
    expect(canRemovePoint(index, [1002, 1002, 1002, 1004], 1002)).toEqual({ ok: true });
  });
```

- [ ] **Step 8: Run the touched suites**

```bash
cd web && npx vitest run src/lib/planner/connectors.test.ts src/lib/planner/rules.test.ts src/fixtures/planner/fixture.test.ts src/lib/planner/load.test.ts src/lib/planner/store.test.ts src/lib/planner/derive.test.ts
```
Expected: PASS.

```bash
cd web && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: `- 0 errors`.

- [ ] **Step 9: Commit**

```bash
git add web/src/lib/planner web/src/fixtures/planner
```

```bash
git commit -F - <<'MSG'
feat(web): connector geometry for the talent prerequisite links

Three shapes, the ones the client draws: straight down a column, straight
across a row, and an elbow. Every line runs cell edge to cell edge so it never
crosses an icon, and the whole thing is a pure function over cells so the
overlay never measures the DOM.

Also carries the two new data fields into the planner's types and fixture, and
adds the missing test that a prerequisite cannot be dropped below the rank one
of its dependents still needs.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 10: The connector overlay, the cell states, and the wording

**Files:**
- Modify: `web/src/lib/planner/styles.ts`
- Create: `web/src/lib/planner/styles.test.ts`
- Modify: `web/src/lib/planner/config.ts:11-14`
- Create: `web/src/lib/planner/config.test.ts`
- Modify: `web/src/components/planner/TreeGrid.svelte`
- Modify: `web/src/components/planner/TalentCell.svelte:39-43`, `:100-119`
- Modify: `web/src/components/planner/OrderStrip.svelte:60-72`
- Modify: `web/src/components/planner/Planner.svelte:8`, `:290`, `:536-540`
- Modify: `web/tests/e2e/planner.spec.ts`

**Interfaces:**
- Consumes: `connectorsFor`, `gridViewBox`, `Connector` from `../../lib/planner/connectors`.
- Produces:
  - `styles.ts`: `type CellState = 'locked' | 'available' | 'filled' | 'maxed'`; `cellState(rank: number, maxRank: number, available: boolean): CellState`; `CELL_BORDER: Record<CellState, string>`; `CELL_PILL: Record<CellState, string>`.
  - `config.ts`: `PREBETA_BUILD = 'forever-prebeta'`; `treeSourceNotice(build: string): string`. `ERA_DATA_NOTICE` is deleted.
  - DOM contract for the e2e: each connector renders as `[data-testid="connector-<from>-<to>"]` carrying `data-met="true"|"false"`; each talent button keeps `data-testid="talent-<id>"` and `data-rank`, and gains `data-state` holding its `CellState`.

- [ ] **Step 1: Write the failing unit tests**

```ts
// web/src/lib/planner/styles.test.ts
import { describe, expect, it } from 'vitest';
import { CELL_BORDER, CELL_PILL, cellState } from './styles';

describe('cellState', () => {
  it('is locked when no point can go in yet', () => {
    expect(cellState(0, 5, false)).toBe('locked');
  });

  it('is available once a point could go in', () => {
    expect(cellState(0, 5, true)).toBe('available');
  });

  it('is filled while it holds points below its cap', () => {
    expect(cellState(1, 5, true)).toBe('filled');
    expect(cellState(4, 5, true)).toBe('filled');
  });

  it('is maxed at the cap, whether or not anything else could be added', () => {
    expect(cellState(5, 5, false)).toBe('maxed');
    expect(cellState(1, 1, false)).toBe('maxed');
  });

  it('has a border and a pill class for every state', () => {
    for (const state of ['locked', 'available', 'filled', 'maxed'] as const) {
      expect(CELL_BORDER[state]).toBeTruthy();
      expect(CELL_PILL[state]).toBeTruthy();
    }
  });
});
```

```ts
// web/src/lib/planner/config.test.ts
import { describe, expect, it } from 'vitest';
import { PREBETA_BUILD, treeSourceNotice } from './config';

describe('treeSourceNotice', () => {
  it('names the client build the trees were read from', () => {
    expect(treeSourceNotice('1.60.1.69893')).toBe(
      'Talent trees read from the game client, build 1.60.1.69893.',
    );
  });

  it('says so plainly while the trees are still the pre-beta snapshot', () => {
    expect(treeSourceNotice(PREBETA_BUILD)).toBe(
      'Talent trees from a pre-beta Wowhead snapshot; the client’s own trees replace them at the beta.',
    );
  });

  it('never claims the trees are Classic Era', () => {
    expect(treeSourceNotice('1.15.9.69722')).not.toContain('Classic Era');
  });
});
```

- [ ] **Step 2: Run them to verify they fail**

```bash
cd web && export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH" && npx vitest run src/lib/planner/styles.test.ts src/lib/planner/config.test.ts
```
Expected: FAIL — `cellState` and `treeSourceNotice` are not exported.

- [ ] **Step 3: Add the states**

Append to `web/src/lib/planner/styles.ts`:

```ts
/**
 * The four states a talent cell is read at a glance by, from design/DESIGN-SYSTEM.md
 * and the game's own frame: gold says "you can spend here", green says "you have",
 * and a dimmed grey says "not yet".
 */
export type CellState = 'locked' | 'available' | 'filled' | 'maxed';

export function cellState(rank: number, maxRank: number, available: boolean): CellState {
  if (rank >= maxRank) return 'maxed';
  if (rank > 0) return 'filled';
  return available ? 'available' : 'locked';
}

/** The cell's own border. `locked` dims the icon with it, which is one state, not two. */
export const CELL_BORDER: Record<CellState, string> = {
  maxed: 'border-kill',
  filled: 'border-kill',
  available: 'border-gold',
  locked: 'border-line-soft opacity-50',
};

/**
 * The rank pill in the cell's corner. Filled and maxed share the green border and
 * differ by the number they show, which is what tells them apart in game too.
 */
export const CELL_PILL: Record<CellState, string> = {
  maxed: 'border-kill text-kill',
  filled: 'border-kill text-text',
  available: 'border-gold text-gold',
  locked: 'border-line text-muted',
};
```

- [ ] **Step 4: Replace the banner constant**

In `web/src/lib/planner/config.ts`, delete `ERA_DATA_NOTICE` and add:

```ts
/** The build id of the pre-beta data set, whose trees came from Wowhead, not a client. */
export const PREBETA_BUILD = 'forever-prebeta';

/**
 * What the planner says above the trees about where they came from. It names the
 * build rather than the expansion: the site serves several builds at once (a shared
 * link renders against the build it was saved on), so "Classic Era trees" was both
 * wrong and unanswerable once the beta client's trees shipped.
 */
export function treeSourceNotice(build: string): string {
  return build === PREBETA_BUILD
    ? 'Talent trees from a pre-beta Wowhead snapshot; the client’s own trees replace them at the beta.'
    : `Talent trees read from the game client, build ${build}.`;
}
```

- [ ] **Step 5: Run the unit tests to verify they pass**

```bash
cd web && npx vitest run src/lib/planner/styles.test.ts src/lib/planner/config.test.ts
```
Expected: PASS, 8 tests.

- [ ] **Step 6: Draw the connectors**

In `web/src/components/planner/TreeGrid.svelte`, extend the script's imports and derivations:

```ts
  import { connectorsFor, gridViewBox } from '../../lib/planner/connectors';
```

```ts
  const connectors = $derived(connectorsFor(tree, store.ranks));
  const viewBox = $derived(gridViewBox(size.tiers, size.columns));
```

and wrap the existing `<div role="grid" …>` in a positioned box with the overlay in front of it in the DOM (so the cells, which are positioned too, paint over the lines):

```svelte
<!-- `w-fit` so the overlay is exactly the grid's box: the grid's columns are
     max-content, so a full-width wrapper would stretch every line sideways.
     preserveAspectRatio="none" lets one viewBox in base-breakpoint units cover
     the slightly larger md cells, and non-scaling-stroke keeps the line weight
     identical at both sizes. -->
<div class="relative w-fit">
  <svg
    class="pointer-events-none absolute inset-0 h-full w-full"
    {viewBox}
    preserveAspectRatio="none"
    aria-hidden="true"
  >
    {#each connectors as connector (connector.id)}
      <path
        d={connector.d}
        data-testid={`connector-${connector.id}`}
        data-met={connector.met}
        fill="none"
        stroke-width="2"
        stroke-linecap="round"
        style="vector-effect: non-scaling-stroke"
        class={connector.met ? 'stroke-gold' : 'stroke-line'}
      />
    {/each}
  </svg>
  <div
    bind:this={root}
    role="grid"
    tabindex={-1}
    aria-label={`${tree.name} talents`}
    data-testid={`tree-${tree.id}`}
    class="relative grid gap-2"
    style={`grid-template-columns: repeat(${size.columns}, minmax(0, max-content));`}
    onkeydown={onKeyDown}
  >
```

(closing the new wrapper `</div>` after the grid's own `</div>`).

- [ ] **Step 7: Use the states in the cell and the strip**

In `web/src/components/planner/TalentCell.svelte`, replace the `borderClass` derivation:

```ts
  import { CELL_BORDER, CELL_PILL, cellState } from '../../lib/planner/styles';
```

```ts
  const state = $derived(cellState(rank, talent.max_rank, available));
```

(and delete the `maxed` and `borderClass` derivations, keeping `rank` and `available`).

In the button, replace the `class={…}` and `data-rank` region with:

```svelte
    data-testid={`talent-${talent.id}`}
    data-rank={rank}
    data-state={state}
    class={`rounded-control bg-card-top relative flex h-11 w-11 items-center justify-center border md:h-12 md:w-12 ${CELL_BORDER[state]}`}
```

and the pill's class with:

```svelte
      class={`tabular rounded-pill bg-bg absolute -right-1 -bottom-1 border px-1 font-mono text-[11px] leading-[14px] ${CELL_PILL[state]}`}
```

In `web/src/components/planner/OrderStrip.svelte`, give each point the same border. Extend the script:

```ts
  import { CELL_BORDER, cellState } from '../../lib/planner/styles';
```

and in the `points` derivation add the state, so the strip reads the same as the grid:

```ts
  const points = $derived(
    store.order.map((id, i) => {
      const talent = store.talentIndex?.byId.get(id) ?? null;
      const rank = store.order.slice(0, i + 1).filter((other) => other === id).length;
      return {
        key: `${i}-${id}`,
        level: levelForIndex(i),
        talent,
        // A point in the strip is always spent, so it is filled or maxed; the
        // third argument only decides between available and locked.
        state: talent ? cellState(rank, talent.max_rank, true) : 'locked',
      };
    }),
  );
```

and on the `<img>`:

```svelte
                  class={`rounded-control h-9 w-9 border object-cover ${CELL_BORDER[point.state]}`}
```

(removing the `border-line` it carries today).

- [ ] **Step 8: Update the banner and the reset button**

In `web/src/components/planner/Planner.svelte`:

```ts
  import { DEFAULT_CLASS_SLUG, treeSourceNotice } from '../../lib/planner/config';
```

```svelte
  <p class="text-muted px-[18px] text-[13px] md:px-0">{treeSourceNotice(store.treeVersion)}</p>
```

and the reset button's label:

```svelte
              Reset…
```

- [ ] **Step 9: Update the browser suite**

In `web/tests/e2e/planner.spec.ts`, replace the banner assertion in the first test:

```ts
  await expect(page.getByText(`Talent trees read from the game client, build ${ACTIVE_BUILD}.`)).toBeVisible();
```

and add a test beside it:

```ts
test('a prerequisite link is drawn and turns gold when the rank is met', async ({ page }) => {
  await page.goto('/planner');
  // Tactical Mastery (1004) needs Deflection (1002) at rank 2.
  const link = page.getByTestId('connector-1002-1004');
  await expect(link).toHaveAttribute('data-met', 'false');
  await expect(page.getByTestId('talent-1004')).toHaveAttribute('data-state', 'locked');

  await page.getByTestId('talent-1002').click();
  await expect(link).toHaveAttribute('data-met', 'false');
  await page.getByTestId('talent-1002').click();
  await expect(link).toHaveAttribute('data-met', 'true');
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-state', 'filled');
  await expect(page.getByTestId('talent-1004')).toHaveAttribute('data-state', 'available');
});

test('reset says it will ask before it clears the build', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Reset…' }).click();
  await expect(page.getByText('Clear every point in this build?')).toBeVisible();
  await page.getByRole('button', { name: 'Keep the build' }).click();
  await expect(page.getByTestId('planner-spent')).toHaveText('1/51');
});
```

Search the rest of `web/tests/e2e/` for the old button name and the old notice and fix every hit:

```bash
cd web && grep -rn "name: 'Reset' \|'Reset'\|Classic Era trees" tests/ src/
```
Expected after the edits: no hits outside this plan's own text.

- [ ] **Step 10: Run the suites**

```bash
cd web && npx vitest run src/lib/planner src/fixtures/planner
```
Expected: PASS.

```bash
cd web && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: `- 0 errors`.

```bash
cd web && E2E_PORT=4325 npx playwright test tests/e2e/planner.spec.ts --project=desktop
```
Expected: PASS.

```bash
cd web && E2E_PORT=4325 npx playwright test tests/e2e/planner-phone.spec.ts tests/e2e/planner-share.spec.ts --project=mobile
```
Expected: PASS.

```bash
cd web && ls dist/planner-island.js
```
Expected: the file exists (the Playwright run built the site).

- [ ] **Step 11: Commit**

```bash
git add web/src/lib/planner web/src/components/planner web/tests/e2e/planner.spec.ts
```

```bash
git commit -F - <<'MSG'
feat(web): prerequisite links, the four cell states, and honest tree wording

Each tree grid now carries an SVG overlay drawing the game's prerequisite
links: grey until the prerequisite holds the rank its dependent needs, gold
after. Cells read in the four states a player scans for -- locked, available,
filled, maxed -- and the point-order strip uses the same ones.

The banner no longer claims Classic Era trees; it names the build the trees
were read from. Reset reads "Reset…" so the confirm step is expected.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 11: The tree art, the counters and the frame

**Files:**
- Modify: `web/scripts/sync-data.mjs:34-42` (`SYNC_ENTRIES`)
- Modify: `web/src/lib/planner/sync-data.test.ts`
- Create: `web/src/fixtures/planner/trees/fixture_arms.webp`, `web/src/fixtures/planner/trees/fixture_fury.webp`
- Modify: `web/src/fixtures/planner/manifest.json`
- Modify: `web/src/components/planner/TreeGrid.svelte`
- Modify: `web/src/components/planner/SummaryBar.svelte:55-70`
- Modify: `web/src/components/planner/Planner.svelte:466-482`
- Modify: `web/tests/e2e/planner.spec.ts`

**Interfaces:**
- Consumes: `dataUrl` from `../../lib/planner/load`; `TalentTree.background`; `MAX_POINTS` from `../../lib/planner/types`.
- Produces:
  - `SYNC_ENTRIES` gains `{ name: 'trees', kind: 'dir', required: false }`, so `builds/<build>/trees/` is published at `/data/<build>/trees/`.
  - DOM contract for the e2e: each tree panel carries `data-testid="tree-panel-<tree id>"` and its counter `data-testid="tree-points-<tree id>"`; the summary bar gains `data-testid="planner-remaining"`.

- [ ] **Step 1: Publish the directory**

In `web/scripts/sync-data.mjs`, add to `SYNC_ENTRIES`, after the `icons` entry:

```js
  { name: 'trees', kind: 'dir', required: false },
```

In `web/src/lib/planner/sync-data.test.ts`, add a case to whichever `describe` covers `SYNC_ENTRIES` (or, if none does, at the end of the file):

```ts
import { SYNC_ENTRIES } from '../../../scripts/sync-data.mjs';

describe('SYNC_ENTRIES', () => {
  it('publishes the per-tree background art, and does not require it', () => {
    const trees = SYNC_ENTRIES.find((entry) => entry.name === 'trees');
    expect(trees).toEqual({ name: 'trees', kind: 'dir', required: false });
  });
});
```

(if `SYNC_ENTRIES` is already imported in that file, reuse the existing import rather than adding a second).

- [ ] **Step 2: Run it to verify it fails**

```bash
cd web && export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH" && npx vitest run src/lib/planner/sync-data.test.ts
```
Expected: FAIL — `expected undefined to deeply equal { name: 'trees', … }`, then PASS once the entry is added. Re-run to confirm.

- [ ] **Step 3: Make the fixture art**

```bash
cd data && export PATH="$HOME/.local/bin:$PATH" && uv run python - <<'PY'
from pathlib import Path

from PIL import Image

from pipeline.art import BACKGROUND_SIZE, process_background

out = Path("../web/src/fixtures/planner/trees")
out.mkdir(parents=True, exist_ok=True)
for name, colour in (("fixture_arms", (150, 60, 40)), ("fixture_fury", (60, 80, 150))):
    process_background(Image.new("RGB", BACKGROUND_SIZE, colour)).save(
        out / f"{name}.webp", "WEBP", quality=90, method=6
    )
    print(name, (out / f"{name}.webp").stat().st_size, "bytes")
PY
```
Expected: two lines, each a few hundred bytes (a flat colour compresses to almost nothing).

```bash
cd web && shasum -a 256 src/fixtures/planner/trees/fixture_arms.webp src/fixtures/planner/trees/fixture_fury.webp
```

Add both to `web/src/fixtures/planner/manifest.json`'s `files` object, keyed `"trees/fixture_arms.webp"` and `"trees/fixture_fury.webp"` with the digests printed above, keeping the object's existing alphabetical order.

- [ ] **Step 4: Draw the background behind the grid**

In `web/src/components/planner/TreeGrid.svelte`, add to the script:

```ts
  import { dataUrl } from '../../lib/planner/load';
```

```ts
  const backgroundSrc = $derived(dataUrl(store.treeVersion, `trees/${tree.background}.webp`));
  // A build with no art still has to render a usable tree, so a missing image
  // leaves the panel's own background rather than a broken-image box.
  let artBroken = $state(false);
```

and inside the `relative w-fit` wrapper, before the `<svg>`:

```svelte
  {#if tree.background && !artBroken}
    <!-- The client's own 320x384 panel, processed in the pipeline. `object-fill`
         rather than `cover`: the art is drawn for exactly this grid, so stretching
         it to the grid's box is what lines its detail up with the cells. -->
    <img
      src={backgroundSrc}
      alt=""
      aria-hidden="true"
      width="320"
      height="384"
      loading="lazy"
      decoding="async"
      data-testid={`tree-art-${tree.id}`}
      class="rounded-control pointer-events-none absolute inset-0 h-full w-full object-fill"
      onerror={() => (artBroken = true)}
    />
  {/if}
```

- [ ] **Step 5: Add the counters and the frame**

In `web/src/components/planner/SummaryBar.svelte`, after the `Level` block and before `Split`:

```svelte
  <div class="flex flex-col gap-1">
    <span class="label text-muted">Left</span>
    <span
      class="tabular text-gold font-mono text-[20px] leading-11"
      data-testid="planner-remaining"
    >
      {MAX_POINTS - store.spent}
    </span>
  </div>
```

In `web/src/components/planner/Planner.svelte`, give each tree panel the game's title bar and name its counter:

```svelte
          <div
            id={`tree-panel-${tree.id}`}
            role="tabpanel"
            aria-labelledby={`tree-tab-${tree.id}`}
            data-testid={`tree-panel-${tree.id}`}
            class="border-line-warm bg-raised rounded-panel flex-col gap-3 border p-4 md:flex {activeTree === i
              ? 'flex'
              : 'hidden'}"
          >
            <!-- The game's tree header: name on the left, points in the tree on the
                 right, a rule under both. Warm border and gold number are the
                 design system's; the proportions are the client's. -->
            <header class="border-line-soft flex items-baseline justify-between border-b pb-2">
              <h2 class="section-title text-[15px]">{tree.name}</h2>
              <span
                class="tabular text-gold font-mono text-[15px]"
                data-testid={`tree-points-${tree.id}`}
              >
                {store.split[i] ?? 0}
              </span>
            </header>
            <TreeGrid {store} {tree} />
          </div>
```

- [ ] **Step 6: Cover it in the browser suite**

Append to `web/tests/e2e/planner.spec.ts`:

```ts
test('each tree draws its own art and counts its own points', async ({ page }) => {
  await page.goto('/planner');
  const art = page.getByTestId('tree-art-161');
  await expect(art).toHaveAttribute('src', `/data/${ACTIVE_BUILD}/trees/fixture_arms.webp`);
  // The art is a real image, not a broken one.
  await expect
    .poll(() => art.evaluate((node: HTMLImageElement) => node.naturalWidth))
    .toBeGreaterThan(0);

  await expect(page.getByTestId('tree-points-161')).toHaveText('0');
  await expect(page.getByTestId('planner-remaining')).toHaveText('51');
  await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('tree-points-161')).toHaveText('1');
  await expect(page.getByTestId('planner-remaining')).toHaveText('50');
});
```

And in `web/tests/e2e/planner-phone.spec.ts`, append:

```ts
test('a tree and its links fit a phone with no sideways scroll', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 800 });
  await page.goto('/planner');
  await expect(page.getByTestId('tree-panel-161')).toBeVisible();
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
  const grid = page.getByTestId('tree-161');
  const box = await grid.boundingBox();
  expect(box!.width).toBeLessThanOrEqual(390 - 36);
  await expect(page.getByTestId('connector-1002-1004')).toBeAttached();
});
```

- [ ] **Step 7: Run the suites**

```bash
cd web && npx vitest run src/lib/planner src/fixtures/planner
```
Expected: PASS.

```bash
cd web && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: `- 0 errors`.

```bash
cd web && E2E_PORT=4325 npx playwright test tests/e2e/planner.spec.ts --project=desktop
```
Expected: PASS.

```bash
cd web && E2E_PORT=4325 npx playwright test tests/e2e/planner-phone.spec.ts --project=mobile
```
Expected: PASS.

- [ ] **Step 8: Commit**

```bash
git add web/scripts/sync-data.mjs web/src/lib/planner web/src/fixtures/planner web/src/components/planner web/tests/e2e
```

```bash
git commit -F - <<'MSG'
feat(web): the game's tree art, per-tree counters and frame

Each tree draws the processed background the pipeline emits for it, behind the
connector overlay and the cells. The tree header becomes the game's title bar
-- name, points in that tree, a rule under both -- and the summary bar gains
the points the build has left to spend.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

---

### Task 12: Switch the site to the beta build

**Files:**
- Modify: `web/src/data/active-build.json`
- Modify: `web/tests/e2e/real-data.spec.ts`
- Modify: `web/package.json` (the `test:e2e:real` script)

**Interfaces:**
- Consumes: everything above. `ACTIVE_BUILD` (from `tests/e2e/support/active-build.ts`) is read from the file this task changes, so every fixture spec follows it without edits.
- Produces: `web/src/data/active-build.json` = `{"build": "1.60.1.69893"}`.

Note on the fixture: `FOREVER_DATA=fixture` publishes `web/src/fixtures/planner` **under the active build's directory name**, so the fixture suites keep working unchanged after the switch. The `build` field inside `src/fixtures/planner/talents/warrior.json` stays `1.15.9.69722` — it records where that hand-made data came from, and changing it to a beta build id would be false. `src/fixtures/planner/fixture.test.ts` asserts it for that reason; leave both alone.

- [ ] **Step 1: Write the failing real-data tests**

Append to `web/tests/e2e/real-data.spec.ts`:

```ts
/** The synced talent file for a class, trees in position order. */
interface SyncedTalent {
  id: number;
  name: string;
  tier: number;
  column: number;
  max_rank: number;
  prereq_talent_id: number | null;
  prereq_rank: number | null;
  spell_id: number;
}
interface SyncedTree {
  id: number;
  name: string;
  position: number;
  background: string;
  talents: SyncedTalent[];
}

function syncedTrees(slug: string): SyncedTree[] {
  const file = readSynced<{ trees: SyncedTree[] }>('talents', `${slug}.json`);
  return [...file.trees].sort((a, b) => a.position - b.position);
}

/** Spends `points` in a tree by always taking the leftmost cell that is open. */
async function spend(page: Page, treeId: number, points: number): Promise<void> {
  for (let i = 0; i < points; i += 1) {
    await page.locator(`[data-testid="tree-${treeId}"] [data-state="available"]`).first().click();
  }
}

test('the trees are laid out in the order the client draws them', async ({ page }) => {
  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  expect(syncedTrees('warrior').map((tree) => tree.name)).toEqual(['Arms', 'Fury', 'Protection']);
});

test('the banner names the build the trees were read from', async ({ page }) => {
  await openPlanner(page);
  await expect(
    page.getByText(`Talent trees read from the game client, build ${ACTIVE_BUILD}.`),
  ).toBeVisible();
});

test('a 31/20/0 warrior is spent and read back as 31/20/0', async ({ page }) => {
  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  const [arms, fury] = syncedTrees('warrior');

  await spend(page, arms.id, 31);
  await spend(page, fury.id, 20);

  await expect(page.getByTestId('planner-split')).toHaveText('31/20/0');
  await expect(page.getByTestId('planner-spent')).toHaveText('51/51');
  await expect(page.getByTestId('planner-remaining')).toHaveText('0');
  await expect(page.getByTestId('planner-level')).toHaveText('60');
  await expect(page.getByTestId(`tree-points-${arms.id}`)).toHaveText('31');
  await expect(page.getByTestId(`tree-points-${fury.id}`)).toHaveText('20');
});

test('a prerequisite link is drawn for every prerequisite the client has', async ({ page }) => {
  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  const pairs = syncedTrees('warrior').flatMap((tree) =>
    tree.talents
      .filter((talent) => talent.prereq_talent_id !== null)
      .map((talent) => ({ tree, talent })),
  );
  expect(pairs.length).toBeGreaterThan(0);
  for (const { talent } of pairs) {
    await expect(
      page.getByTestId(`connector-${talent.prereq_talent_id}-${talent.id}`),
    ).toBeAttached();
  }

  // The first one turns gold once its prerequisite is maxed.
  const { tree, talent } = pairs[0];
  const prereq = tree.talents.find((t) => t.id === talent.prereq_talent_id)!;
  const link = page.getByTestId(`connector-${prereq.id}-${talent.id}`);
  await expect(link).toHaveAttribute('data-met', 'false');
  await spend(page, tree.id, 0);
  for (let i = 0; i < prereq.max_rank; i += 1) {
    await page.getByTestId(`talent-${prereq.id}`).click();
  }
  await expect(link).toHaveAttribute('data-met', 'true');
  await expect(page.getByTestId(`talent-${prereq.id}`)).toHaveAttribute('data-state', 'maxed');
});

test('every tree ships its own processed art', async ({ page }) => {
  await openPlanner(page);
  for (const slug of readSynced<{ slug: string }[]>('classes.json').map((row) => row.slug)) {
    for (const tree of syncedTrees(slug)) {
      const response = await page.request.get(`/data/${ACTIVE_BUILD}/trees/${tree.background}.webp`);
      expect(response.status(), `${slug} ${tree.name}`).toBe(200);
      expect(response.headers()['content-type']).toContain('image/webp');
    }
  }
});

test('a build shared on the new trees reopens from its link, on desktop and on a phone', async ({
  page,
}) => {
  let body: { point_order: number[]; tree_version: string } | undefined;
  await page.route('**/v1/builds', async (route) => {
    body = route.request().postDataJSON();
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { id: 'real1234', url: 'https://foreversixty.gg/b/real1234' },
        error: null,
        request_id: 'req-1',
      }),
    });
  });

  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  const [arms, fury] = syncedTrees('warrior');
  await spend(page, arms.id, 31);
  await spend(page, fury.id, 20);
  await page.getByLabel('Title').fill('Arms PvE');
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/real1234');

  expect(body!.tree_version).toBe(ACTIVE_BUILD);
  expect(body!.point_order).toHaveLength(51);

  // /b/:id is rendered by the Go API in production; the static preview gets the
  // same markup the interface contract specifies, carrying the body just posted.
  const record = {
    id: 'real1234',
    class_id: 1,
    race_id: 1,
    tree_version: ACTIVE_BUILD,
    point_order: body!.point_order,
    gear: {},
    title: 'Arms PvE',
    created_at: '2026-09-17T09:00:00Z',
    views: 1,
  };
  const shared = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Arms PvE · 31/20/0 · Forever Sixty</title>
<link rel="stylesheet" href="/planner-island.css"></head>
<body><main id="main">
<div id="planner" data-build='${JSON.stringify(record)}' data-tree-version="${ACTIVE_BUILD}"></div>
<script type="module" src="/planner-island.js"></script>
</main></body></html>`;
  await page.route('**/b/real1234', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: shared }),
  );

  await page.goto('/b/real1234');
  await expect(page.getByTestId('planner-split')).toHaveText('31/20/0');
  await expect(page.getByLabel('Class')).toBeDisabled();

  await page.setViewportSize({ width: 390, height: 800 });
  await page.reload();
  await expect(page.getByTestId('planner-split')).toHaveText('31/20/0');
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
});
```

- [ ] **Step 2: Run them against the current active build to see them fail**

```bash
cd web && export PATH="$HOME/.nvm/versions/node/v22.12.0/bin:$PATH" && E2E_PORT=4325 FOREVER_DATA=real npx playwright test tests/e2e/real-data.spec.ts --project=desktop
```
Expected: FAIL — the active build is still `forever-prebeta`, whose trees are the Wowhead snapshot's, so the banner text, the tree art requests and the tab order assertions do not hold.

- [ ] **Step 3: Switch the active build**

`web/src/data/active-build.json`:

```json
{
  "build": "1.60.1.69893"
}
```

- [ ] **Step 4: Run the real-data suite**

```bash
cd web && E2E_PORT=4325 FOREVER_DATA=real npx playwright test tests/e2e/real-data.spec.ts --project=desktop
```
Expected: PASS, all tests.

- [ ] **Step 5: Run the fixture suites, which now publish under the new build id**

```bash
cd web && npx vitest run
```
Expected: PASS.

```bash
cd web && NO_COLOR=1 npx astro check 2>&1 | grep -E "^- [0-9]+ errors?"
```
Expected: `- 0 errors`.

```bash
cd web && E2E_PORT=4325 npx playwright test --project=desktop
```
Expected: PASS.

```bash
cd web && E2E_PORT=4325 npx playwright test --project=mobile
```
Expected: PASS.

```bash
cd web && ls dist/planner-island.js
```
Expected: the file exists.

- [ ] **Step 6: Prove the API validates a build against the new version**

```bash
cd api && go test ./...
```
Expected: PASS.

```bash
cd api && mkdir -p cmd/validatecheck && cat > cmd/validatecheck/main.go <<'GO'
package main

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func main() {
	d, err := trees.Load("../data/builds")
	if err != nil {
		panic(err)
	}
	b, ok := d.Build("1.60.1.69893")
	if !ok {
		panic("1.60.1.69893 did not load")
	}
	latest, _ := d.Latest()
	fmt.Println("latest:", latest.Version)

	arms := b.Trees(1)[0]
	var order []int
	for _, t := range arms.Talents {
		if t.Tier != 0 {
			continue
		}
		for r := 0; r < t.MaxRank; r++ {
			order = append(order, t.ID)
		}
	}
	in := builds.Input{ClassID: 1, RaceID: 1, TreeVersion: "1.60.1.69893", PointOrder: order}
	fmt.Println("points:", len(order), "errors:", builds.Validate(d, in))
}
GO
go run ./cmd/validatecheck
```
Expected: `latest: 1.60.1.69893` and `errors: map[]` for a legal row-0 order. If `builds.Input`/`builds.Validate` have different names, read `api/internal/builds/validate.go` and use the ones it exports; the point of the check is that a build made of real beta talent ids validates against `tree_version` `1.60.1.69893`.

```bash
cd api && rm -rf cmd/validatecheck
```

- [ ] **Step 7: Commit**

```bash
git add web/src/data/active-build.json web/tests/e2e/real-data.spec.ts web/package.json
```

```bash
git commit -F - <<'MSG'
feat(web): serve the planner on the beta client's own trees

active-build moves to 1.60.1.69893, so /planner draws Forever's real talents
in the game's tab order over the game's own art. forever-prebeta and
1.15.9.69722 stay published, so every build shared against them keeps opening
against the trees it was made on.

The real-data smoke now covers the order, a 31/20/0 warrior spent and read
back, every prerequisite link, every tree's art, and a share that reopens from
its link on desktop and on a phone.

Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropic.com>
Claude-Session: https://claude.ai/code/session_01Wq4TtMZmJRhEJwCNH2gj1h
MSG
```

- [ ] **Step 8: Hand the persona review back to the controller**

This is the last task. Report to the controller, in these words, that the plan is complete and that its acceptance is outside this worktree:

> Every group has merged. The planner persona review runs against a harness built from `main` and must be re-run there — this worktree never runs it. The persona to re-run is the single planner persona, and the things it should now see that it could not before are: the three trees in the game's order (Warrior: Arms, Fury, Protection), the prerequisite links drawn and coloured, the four cell states, each tree's own background art, the per-tree point counters, and a banner that names the build rather than claiming Classic Era.

Do not run the review, and do not merge to `main`: the controller commits and merges.
