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
}

export interface BisSlot {
  slot: string;
  item_id: number;
  item_name: string;
  source: string;
  /** One of loot.ts's LootKind -- the picker's own kinds, so a BiS row's badge reads the
   *  same word as the Droptimizer's. */
  source_kind: string;
  score: number;
  verified: boolean;
  swap_note?: string;
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
  /** Precomputed by the pipeline: `"<slot>: <item name>"` for every slot whose BiS pick
   *  changed since the previous band, faction held constant. Empty at the lowest band (no
   *  previous band to diff against). */
  new_at_band: string[];
  weights_run_seconds: number;
  verify_run_seconds: number;
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
