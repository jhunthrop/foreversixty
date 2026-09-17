# The planner on the client's own trees

Status: approved in conversation 2026-09-17 (the persona review of the planner and the owner's
three notes: tree order, connection links, the in-game look). Written from measurements on the
Forever beta client, build 1.60.1.69893.

## What it is for

The planner shows Forever's talent trees from a Wowhead pre-beta snapshot (`data/raw-forever/`,
470 talents, provisional spell ids 104724–113398) laid over Classic Era's tab data. The beta
client now carries the real trees, in the modern trait tables rather than the legacy `Talent`
table the pipeline reads: nine class trees of 50–54 nodes, 469 talents once two stale twin nodes are dropped, real spell ids (75 above the Era range, 43 of them in the 1.3M Forever range), each node's grid position, 96 prerequisite edges,
rank caps, and the tab order and background art names in the legacy `TalentTab` rows that the
client keeps for its UI. Reading those makes the planner true to the game in one change — the
order (Warrior: Arms, Fury, Protection), the positions, the links — and gives every consumer
keyed by spell id (the sim, the report's planner link) the ids the client will write.

The planner's look then follows the game's talent frame: each tree's own background art, the
frame, a per-tree point counter, and the three cell states players read at a glance, all
rendered in this site's palette rather than pasted in.

## Global constraints

- Both clients stay supported by the pipeline: Classic Era (legacy `Talent`/`TalentTab`) and the
  1.60 client (trait tables). The emitted per-class talent JSON keeps one schema, with fields
  added, never renamed; the Era build re-emits byte-identical except for added fields.
- The Wowhead snapshot stays committed as the record of what was believed before the beta; the
  pipeline stops reading it once the trait tables emit the same trees, and the diff between the
  two (names, ranks, positions, prerequisites) is written to `diffs/` and summarised in the
  report.
- Spell ids in the emitted trees are the client's. A `tree_version` moves with the data so
  shared builds made against the old trees keep validating against their own version; the API's
  talent data (`api/internal/trees`) is regenerated from the same emitted files.
- The site's active build switches to the beta build only when the planner e2e passes on it and
  the API validates a shared build against it; the switch is its own task at the end.
- Art: the game's own background textures, fetched from the client by file id the way icons
  are, composited per tree and processed in the pipeline (desaturated, darkened, tinted toward
  the site's palette; the exact treatment is a constant in one place) so the site ships one
  processed image per tree, never a raw texture. Frame, borders, counters and cell states are
  CSS in the game's proportions, not textures.
- Phone first: the tree art and links render at 390px with no horizontal scroll inside a tree;
  the links stay legible when a tree is the width of a phone.
- No third-party requests at runtime; everything the planner draws ships with the site.

## 1. Data: the trait tables

Measured on build 1.60.1.69893 (wago.tools serves each by name with `?build=`):

| Table | Rows | Carries |
|---|---|---|
| `TraitTree` | 17 | one tree per class (9 mapped through `SkillLineXTraitTree`), plus non-class trees (Field Guide-style systems of 9–16 nodes) that are not talents; a class tree's three tabs are its maximal `TraitNodeGroup` rows, ordered by their `PosX` band, which equals `TalentTab.OrderIndex` |
| `TraitNode` | 560 | `TraitTreeID`, `PosX`, `PosY` (grid position in a fixed unit), `Type`, `Flags` |
| `TraitNodeEntry` | 656 | `TraitDefinitionID`, `MaxRanks` |
| `TraitNodeXTraitNodeEntry` | 563 | node → entry |
| `TraitDefinition` | ~650 | one `SpellID` per talent (not per rank), `OverrideName`, `OverrideIcon`; per-rank values through `TraitDefinitionEffectPoints` → `CurvePoint` |
| `TraitEdge` | 96 | `LeftTraitNodeID` → `RightTraitNodeID`, `Type`, `VisualStyle`: the prerequisite links |
| `TraitCond` | 172 | the per-tier points-spent gates (5 per tier, 27 tabs) and three vestigial rows; a prerequisite's required rank is the prerequisite's own `MaxRanks` |
| `TraitNodeGroup`, `TraitNodeGroupXTraitNode` | 368 / 2,159 | groupings (tiers and/or specs; to be measured) |
| `SkillLineXTraitTree` | 9 | class skill line → tree |
| `TalentTab` (legacy) | 27 | `Name`, `OrderIndex`, `BackgroundFile`, `ClassMask`: the three tabs per class in the game's order |

Open questions the plan's first task measures rather than assumes: how a class's single trait
tree splits into its three tabs (by `PosX` column bands, by `TraitNodeGroup`, or by a tab id on
the node); how `PosX`/`PosY` quantise to the 4-column, 7-row grid; which `TraitCond` rows are the
per-tier points gates and which carry a prerequisite rank; and whether `TraitDefinition.SpellID`
is the talent's rank-1 spell with ranks resolved through `SpellName`/`Spell` the way the Wowhead
payload gave every rank's text.

The pipeline gains a trait reader used when the tables exist for the build (the 1.60 client) and
keeps the legacy reader for Era. The emitted per-class JSON gains, per tree: `order` (from
`OrderIndex`), `background` (the processed image's path); per talent: `spell_id` (the client's),
`row`/`column` from the node position, `prereq_talent_id`/`prereq_rank` from the edges and
conditions (the fields exist today), and ranks with their descriptions.

## 2. Planner: links and states

- Trees render in `order`. The Split readout and every place that lists the three trees follow it.
- A connector is drawn from each prerequisite cell to its dependent the way the game draws it: a
  straight vertical or horizontal line when they share a column or row, an elbow otherwise, as an
  SVG overlay on the tree grid; gold once the prerequisite's required rank is met, grey until
  then. The dependent stays locked (the existing refusal text, which already names the reason)
  until its prerequisite and tier gate are met; a prerequisite cannot drop below what its
  dependents need (the existing rules are audited for that case and a test added).
- Cell states: available (gold border), filled (green border, rank in the corner), maxed (green
  with the max rank), locked (grey, dimmed icon). The point-order strip uses the same states.
- The header banner stops claiming "Classic Era trees" once the client's trees are shown; it
  names the build the trees come from.
- Reset keeps its confirm step but the button reads "Reset…" so the confirm is expected.

## 3. Planner: the look

- Each tree draws its processed background behind the grid, the grid's cells sized so a tree
  reads at the game's proportions on desktop (three trees side by side) and one tree at a time on
  a phone (the existing tab strip).
- A per-tree point counter in each tree's header ("Arms 31"), and the class's remaining points
  beside the level, styled as the game's counters but in the site's type.
- The frame: a border and header treatment in the game's proportions from CSS tokens, not
  textures, so it holds in both themes.

## 4. API

`api/internal/trees` loads the emitted trees per `tree_version`; the new version's data is
generated from the same files, and a shared build is validated against the version it names.
The share flow's e2e covers: share on the new trees, open the link, open it on a phone.

## Testing

- Pipeline: unit tests on fixture rows for the trait reader (positions, edges, ranks, the tab
  split), the art processing (a fixture texture in, the processed size and palette out), and the
  Wowhead diff; the Era build's JSON byte-identical apart from added fields.
- Planner: unit tests for the rules (prerequisite lock, no dropping below a dependent's need,
  tier gates) and the connector geometry; e2e on the beta build's trees: order, a known build
  (31/20/0 warrior) entered and read back as 31/20/0, a link drawn and coloured, share and
  reopen on desktop and phone, the banner's wording.
- Persona: the single planner persona re-run once the whole plan has merged.
