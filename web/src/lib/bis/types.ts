// web/src/lib/bis/types.ts
// The leveling BiS data contract, as design doc 2026-09-28-leveling-bis-design.md and this
// lane's own brief state it: one file per written spec, published once per class/spec/band
// rather than per character. Lane `bis-all` writes the real files (data/builds/<build>/bis/
// <spec>.json); this lane's fixture (src/data/fixtures/bis/) matches the same shape so the
// page can be built and tested before that lane lands on main.

export type Faction = 'alliance' | 'horde';

export interface BisStatWeight {
  stat: string;
  weight: number;
  /** The sweep's standard error on `weight` (leveling-bis writes it since 2026-09-28). */
  error?: number;
  /** `true` when the error is over a quarter of the weight: shown faint, never trusted. */
  insignificant?: boolean;
}

/**
 * `BisSlot` types every field but `swap_note` as required to match `bis/hover.ts` and
 * `BisSlotPopover.svelte` (both read straight through to `pick.item_id` once a pick is
 * found, with no `undefined` check). The real pipeline output does not honour that for a
 * slot the ranking found no source for: it writes `{ slot, verified: false }` alone, every
 * item field simply absent rather than a placeholder id. Code reading a `BisSlot` off a real
 * file checks for that with `hasKnownSource` (`source-cell.ts`) before trusting `item_id`
 * and friends, the same discipline `sourceBadgeLabel` already applies to `source_kind`.
 */
export interface BisSlot {
  slot: string;
  item_id: number;
  item_name: string;
  source: string;
  /** One of loot.ts's LootKind -- the picker's own kinds, so a BiS row's badge reads the
   *  same word as the Droptimizer's. */
  source_kind: string;
  /** The stat-weight estimate in the band's `score_unit`; absent on a row the ranker's
   *  own sim decided (trinkets, proc items, weapon pairs), which carries `sim_dps` instead
   *  (ranker-integrity-2, 2026-09-30: one number per row, never two units side by side). */
  score?: number;
  /** The measured set DPS with this item, for a sim-decided row. */
  sim_dps?: number;
  verified: boolean;
  /** A weapon slot where no sourced candidate carried a positive score and the ranker
   *  fell back to the best by item level rather than publish an empty weapon slot. */
  low_value?: boolean;
  /** Why an empty slot is empty: `two_hand_equipped`, `no_dps_value`, `no_sourced_item`,
   *  `effect_not_modelled`. */
  empty_reason?: string;
  /** The pick's own proc or use effect is not simulated (`hasImplementedEffect`,
   *  `sim/cmd/leveling-bis/report.go`) -- the row was ranked on its stats alone, which can
   *  let an unmodelled-effect item win by a margin smaller than its own blind spot (fourth
   *  wow-player sweep, day 3: "Serenity Field"). Never conflated with `empty_reason`'s own
   *  `effect_not_modelled` value, which describes an EMPTY slot with no sourced candidate at
   *  all, not a flag on a filled pick. */
  effect_unmodelled?: boolean;
  swap_note?: string;
  /** The next-best sourced candidates after this row's own pick (`sim/cmd/leveling-bis/
   *  report.go`'s own `slotRow.Alternatives`) -- up to three, ties (identical score to the
   *  pick) ranked first with `dps_delta` 0, then the next-best by score. Optional: a file
   *  published before this field existed carries no `alternatives` key at all;
   *  `normaliseBisFile` (`load.ts`) defaults that (and a literal JSON `null`) to `[]`, the
   *  same discipline it already applies to `coverage`/`new_at_band`. A separate lane renders
   *  this list -- this lane only carries the data through. */
  alternatives?: BisAlternative[];
}

/** One candidate `BisSlot.alternatives` names beyond the slot's own pick. */
export interface BisAlternative {
  item_id: number;
  item_name: string;
  sim_dps?: number;
  source_kind: string;
  source: string;
  /** The row's own DPS gap against the pick, always in real DPS -- the ranker
   *  (`sim/cmd/leveling-bis/report.go`'s `alternativeRow.DPSDelta`) never publishes a raw
   *  score-unit figure here (bis-ranker-integrity-3, 2026-09-29: that was this field's
   *  first cut and read as a positive DPS gap when it was really score points). Usually
   *  negative. For the one alternative `verified` is true on, this is instead the real,
   *  sim-measured delta -- see `verified`'s own doc. */
  dps_delta: number;
  /** True for at most one alternative per slot: the runner-up the ranker's own verify pass
   *  actually simmed against the pick, whose `dps_delta` above is that sim's real measured
   *  delta rather than a score estimate -- a runner-up that scored higher than the pick but
   *  lost the real sim must never publish a `dps_delta` that makes it look like the better
   *  fallback (owner review, tenet 8). Omitted (falsy) for every other row: a score estimate
   *  the ranker never simmed at all. */
  verified?: boolean;
  /** See `BisSlot.effect_unmodelled`'s own doc -- an alternative can carry the identical
   *  flag, for the same reason. */
  effect_unmodelled?: boolean;
}

/** One planner slot's coverage: how many items `eligible()` (sim/cmd/leveling-bis/eligible.go)
 *  passed for this band+faction, and how many of those `sourceFor()` could actually find a
 *  source for. A slot missing from `BisBand.coverage` had zero eligible candidates at all --
 *  not even an unsourced one. */
export interface BisCoverage {
  eligible: number;
  sourced: number;
}

export interface BisBand {
  spec: string;
  band: number;
  faction: Faction;
  race: string;
  talents: string;
  talent_points: number;
  weights: BisStatWeight[];
  slots: BisSlot[];
  set_dps: number;
  no_source_count: number;
  /** Lane `rank-guardrails`' guardrail A: planner slot -> `{ eligible, sourced }`
   *  (`sim/cmd/leveling-bis/report.go`'s own `Coverage` field). Required here the same way
   *  `new_at_band` is: a file the nightly published before this field existed carries no
   *  `coverage` key at all, and `normaliseBisFile` (`load.ts`) defaults that (and a literal
   *  JSON `null`) to `{}`, so every caller can read `band.coverage[slot]` with no existence
   *  check of its own. */
  coverage: Record<string, BisCoverage>;
  /** Precomputed by the pipeline: `"<slot>: <item name>"` for every slot whose BiS pick
   *  changed since the previous band, faction held constant. Empty at the lowest band (no
   *  previous band to diff against). */
  new_at_band: string[];
  weights_run_seconds: number;
  verify_run_seconds: number;
  /** The measured, un-normalised DPS this band's weights run found for one point of the
   *  spec's own reference stat (`sim/cmd/leveling-bis/report.go`'s own
   *  `bandReport.ReferenceDPSPerPoint`) -- the raw number every `weights[i].weight` ratio is
   *  normalised against, letting the page turn "Strength 1.99" into "= 1.99 *
   *  reference_dps_per_point DPS per point". Optional and defaults to `null`
   *  (`normaliseBisFile`, `load.ts`): a file published before this field existed carries no
   *  `reference_dps_per_point` key at all. */
  reference_dps_per_point?: number | null;
  /** The unit every stat-weight `score` on this band is in (`"reference_stat_points"`);
   *  a sim-decided row carries `sim_dps` instead of a score. */
  score_unit?: string;
}

export interface BisFile {
  spec: string;
  build: string;
  engine_version: string;
  generated_at: string;
  bands: BisBand[];
}

/** data/curated/specs.json's own shape -- the master list of written specs, one row per
 *  spec regardless of whether its BiS file exists yet. */
export interface SpecCatalogEntry {
  spec: string;
  class_slug: string;
  spec_slug: string;
  name: string;
  role: string;
  tree_index: number;
  reference_stat: string;
  weight_stats: string[];
}

/** `data/builds/<build>/items/<class>.json`'s own per-item shape, the fields this lane's
 *  source cell, item-level column and `ItemHover` stub read (`load.ts`'s `itemDetails`). */
export interface ItemDetail {
  name: string;
  quality: number;
  item_level: number;
  required_level: number;
  /** The client's icon name (e.g. `inv_jewelry_ring_26`), no extension, no path -- the same
   *  value `ItemHover`'s own item model carries. */
  icon: string;
  stats: Record<string, number>;
}

/** `loot.json`'s `quests` map: item id (as a string key) -> every quest that awards it,
 *  one entry per faction that has its own quest for the item. */
export interface LootQuestOption {
  quest_id: number;
  name: string;
  /** `'both'` for a neutral quest every faction can pick up -- the BiS row's own faction
   *  badge comes from the band it is rendered under, not this field (see
   *  `source-cell.ts`'s `resolveSourceCell`), so a `'both'` quest still shows the right
   *  faction word on each panel rather than a third label. */
  faction: Faction | 'both';
  min_level: number;
  level: number;
}

/** `loot.json`'s shape beyond `sim/loot.ts`'s own `LootFile` (which only reads `sources`):
 *  the quest-reward map a BiS quest row's name and level come from. Declared here rather
 *  than widening `sim/loot.ts`'s `LootFile` itself, which the Droptimizer picker also uses
 *  and has never needed the quest map for (every quest source there is one undifferentiated
 *  "Quests" pill, contract 6.1's own default-off kind). */
export interface LootQuestsFile {
  quests: Record<string, LootQuestOption[]>;
}

/**
 * The pre-resolved item `ItemHover`'s stub renders (its documented `model?` prop) -- every
 * `/bis` caller already has this from `itemDetails` (`load.ts`) at build time, so the stub
 * never reads or fetches on its own. The real component's own model may differ once lane
 * `web-item-tooltips` lands; this is this lane's own shape, used only by its stub and by the
 * page that builds it.
 */
export interface ItemHoverModel {
  name: string;
  quality: number;
  itemLevel: number;
  requiredLevel: number;
  icon?: string;
  stats?: Record<string, number>;
}

/** One slot's pick, before and after a band boundary, faction held constant -- the "what
 *  changed since <band>" panel's own row, and the source `load.ts`'s `changedSinceBand`
 *  computes it from for a band's "new" row markers too. Either side is undefined when the
 *  slot had, or still has, no known source. */
export interface ChangedSlot {
  slot: string;
  before?: BisSlot;
  after?: BisSlot;
}
