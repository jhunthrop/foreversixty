// web/src/lib/bis/types.ts
// The leveling BiS data contract, as design doc 2026-09-28-leveling-bis-design.md and this
// lane's own brief state it: one file per written spec, published once per class/spec/band
// rather than per character. Lane `bis-all` writes the real files (data/builds/<build>/bis/
// <spec>.json); this lane's fixture (src/data/fixtures/bis/) matches the same shape so the
// page can be built and tested before that lane lands on main.

export type Faction = 'alliance' | 'horde';

/** The per-second unit a band's figures are in: damage for the damage specs, effective
 *  healing for a healer (the ranker keeps the DPS key names for both). */
export type RateUnit = 'DPS' | 'HPS';

export interface BisStatWeight {
  stat: string;
  weight: number;
  /** The sweep's standard error on `weight` (leveling-bis writes it since 2026-09-28). */
  error?: number;
  /** `true` when the error is over a quarter of the weight: shown faint, never trusted. */
  insignificant?: boolean;
  /** `'rating'` for a rating-family stat (hit, crit, dodge, parry, block, defense) --
   *  `weight`/`error` above are per RATING POINT (the number on an item's own tooltip,
   *  e.g. "+14 Crit"), not per percent (`sim/cmd/leveling-bis/report.go`'s own
   *  `publishWeightRatingUnits`). Absent for every other stat (Agility, AP, RAP, SP, ...):
   *  today's flat, weight-per-point-of-the-stat-itself semantics, unchanged. */
  unit?: 'rating';
  /** This build's own gametables/combatratings.txt level-60 rating points per 1% for
   *  `stat` -- set only alongside `unit: 'rating'`, so a consumer can render "14 Crit
   *  rating = 1% Crit" without hardcoding the client's own conversion table a second time.
   *  `weight === weight_per_percent / rating_factor` holds exactly wherever this is set. */
  rating_factor?: number;
  /** The weight exactly as the weights sweep measured it, per SIM UNIT (one point of
   *  hit/crit/dodge/parry/block/defense percentage) -- typed and passed through but never
   *  rendered by this rail (reserved for the Stat Weights tool, tenet 8: never show a
   *  number this surface wasn't asked to explain). Equal to `weight` for a non-rating-
   *  family stat. */
  weight_per_percent?: number;
  /** `weight` (already per point) divided by this band's `scale_reference_stat`'s own
   *  `weight` -- the SimulationCraft-familiar convention this lane's brief asks for
   *  (`sim/cmd/leveling-bis/report.go`'s own `normalizeScaleFactors`): the single
   *  highest-weighted PER-POINT stat reads exactly `1`, every other stat a fraction of it.
   *  Absent on a file published before this lane -- `panel-view.ts`'s `scaleRowsFor`
   *  computes it client-side from `weight`/`error`/the band's own `reference_dps_per_point`
   *  when missing, so the rail never breaks on an older file. */
  scale_factor?: number;
  /** Absolute DPS per point of `stat`: `weight * reference_dps_per_point`. Independent of
   *  `scale_factor` -- needs no anchor, only this row's own `weight` and the band's own
   *  `reference_dps_per_point`. */
  dps_per_point?: number;
  /** `error`, divided by the same divisor `scale_factor` uses -- so a reader comparing two
   *  rows' "±" compares the same normalized units the rows' own `scale_factor` is in. */
  scale_error?: number;
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
/** The set bonus a pick was adopted for (`sim/cmd/leveling-bis/sets.go`'s `setBonusNote`):
 *  the ranker wears this piece to complete `pieces` of `set`, not for its own stats. */
export interface BisSetBonus {
  set: string;
  pieces: number;
  bonus: string;
}

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
  /** Present only on a slot the ranker adopted for the set bonus it completes. */
  set_bonus?: BisSetBonus;
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
  /** `"not_in_sim"` (`sim/cmd/leveling-bis/report.go`'s own `notInSimReason`) when this
   *  row's own pick carries an item id this build's embedded item database does not have
   *  -- the pick stays score-decided (never sim-verified), `verified` is forced `false`,
   *  and no `sim_dps`/`dps_delta`/`swap_note` is ever published for it (report.go's own
   *  `SimStatus` doc). Sibling of `effect_unmodelled`, same convention: absent (undefined)
   *  for every other row, never defaulted by `normaliseBisFile` (`load.ts`) -- a missing
   *  key simply means "not flagged," matching `effect_unmodelled`'s own default-free
   *  behaviour. */
  sim_status?: string;
  swap_note?: string;
  /** This row's own pick's measured DPS advantage over the one comparator a real swap sim
   *  actually measured it against, in that same run (`sim/cmd/leveling-bis/report.go`'s own
   *  `slotRow.DPSDelta`) -- always real and always non-negative, unlike the two absolute
   *  numbers `swap_note` may embed (each a full-set snapshot at the moment ITS OWN slot was
   *  decided, which can silently disagree with the band's own published `set_dps` and with
   *  each other row's snapshot -- see `report.go`'s own `dpsComparisonPhrase` doc). Distinct
   *  from `alternatives[].dps_delta`, which is per-alternative. Optional and defaults to
   *  `null` (`normaliseBisFile`, `load.ts`): a file published before this field existed, or a
   *  row never compared to anything a sim actually measured, carries no `dps_delta` key at
   *  all. */
  dps_delta?: number | null;
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

/** `bare`: a naked character (no raid buffs, debuffs or consumables). `raid`: the Phase 1
 *  raid-ready preset. */
export type BisPresetId = 'bare' | 'raid';

/** One buff, debuff or consumable a preset applies. */
export interface BisPresetEffect {
  id: number;
  label: string;
}

/** `BisFile.presets[<id>]`: what a non-bare preset puts on the simulated character. */
export interface BisPresetMeta {
  label: string;
  buffs: BisPresetEffect[];
  debuffs: BisPresetEffect[];
  consumes: BisPresetEffect[];
  notes?: string;
}

/** `BisBand.hit_to_cap` for a weapon user (`sim/cmd/leveling-bis/hitcap.go`): the weights
 *  character's distance to the hit caps, all in engine percent of hit (10 hit rating = 1%).
 *  Published for a band whose spec swings or shoots; a caster's is `BisSpellHitToCap`. It has
 *  no `kind` key, which is how the two shapes are told apart. */
export interface BisMeleeHitToCap {
  kind?: undefined;
  /** The percent of hit the band's character has. */
  baseline: number;
  /** Percent still worth its full weight for specials, before the 9% cap against a level-63
   *  target. */
  specials: number;
  /** Percent to the white-swing cap; present only for a dual-wield spec. */
  white?: number;
}

/** `BisBand.hit_to_cap` for a caster (`spellHitToCap` in `hitcap.go`): the distance to the
 *  16% spell hit cap against a level-63 target (17% base miss, 1% always left). */
export interface BisSpellHitToCap {
  kind: 'spell';
  /** The percent of hit the band's character has before any school talent. */
  baseline: number;
  /** Percent still worth its full weight for a spell with only that baseline. */
  spell: number;
  /** Present when a talent adds hit to some spells (Elemental Precision, Arcane Focus,
   *  Nature's Reach, Suppression): the distance for those spells, which is `spell` less the
   *  bonus. */
  school?: {
    names: string[];
    /** Those spells' total hit, in percent. */
    hit: number;
    /** Percent still worth its full weight for them. */
    to_cap: number;
  };
}

export type BisHitToCap = BisMeleeHitToCap | BisSpellHitToCap;

/** The role a band was ranked for. Absent on a file published before roles: `'dps'`
 *  (`normaliseBisFile` defaults it). */
export type BisRole = 'dps' | 'tank' | 'healer';

/** `BisBand.metrics` on a tank band: the boss-profile figures the tank headline shows.
 *  Within a tank band every slot's `sim_dps`/`dps_delta` and each alternative's `dps_delta`
 *  are in tank-score points (mitigation, risk and threat combined), never damage. */
export interface TankMetrics {
  /** Damage taken per second at the boss profile. */
  dtps: number;
  /** Theck-Meloree Index: lower is better. */
  tmi: number;
  /** Chance of death over the fight, 0..1. */
  chance_of_death: number;
  /** Threat per second. */
  tps: number;
  /** Hit points against the boss profile. */
  effective_health: number;
  /** The tank's own damage per second (what `set_dps` carries on a tank band). */
  dps: number;
}

/** A healer band's measured outcomes under its `profile` (contract: the healer ranker). */
export interface BisHealMetrics {
  /** Effective healing per second; equals the band's `set_dps`. */
  hps: number;
  raw_hps: number;
  /** 0..1 share of raw healing that landed on full health. */
  overheal_pct: number;
  /** Seconds until the first out-of-mana, capped at 3600; at or past the fight length = mana to spare. */
  mana_lasts_sec: number;
  /** Effective healing per point of mana spent. */
  hpm: number;
}

/** The incoming-damage profile a healer file ranks gear under
 *  (`data/curated/heal-profile.json`, republished as the file's top-level `heal_profile`). */
export interface BisHealProfile {
  id: string;
  label: string;
  summary: string;
  notes: string;
  duration_sec: number;
  damage_spread: number;
  tank: { health: number; hit_damage: number; swing_seconds: number; reason: string };
  members: { health: number; reason: string };
  pulse: { damage: number; interval_seconds: number; members: number; reason: string };
  sources: { label: string; url: string; kind: 'blessing' | 'site' | 'blizzard' }[];
}

export interface BisBand {
  /** The role the band was ranked for. Absent on a file published before roles:
   *  `normaliseBisFile` defaults it to `'dps'`. On a healer band every `*_dps*` key is
   *  effective healing per second. */
  role?: BisRole;
  /** The `heal_profile.id` a healer band was ranked under. */
  profile?: string;
  /** Present on a tank band (`TankMetrics`) or a healer band (`BisHealMetrics`) only;
   *  `normaliseBisFile` defaults a tank band's to `null`. */
  metrics?: TankMetrics | BisHealMetrics | null;
  /** Absent on a file published before presets: such an entry is bare. */
  preset?: BisPresetId;
  spec: string;
  band: number;
  faction: Faction;
  race: string;
  talents: string;
  talent_points: number;
  weights: BisStatWeight[];
  /** See `BisHitToCap`. Absent on a file published before the key, and on a spec with no hit table;
   *  `normaliseBisFile` (`load.ts`) defaults it to `null`. */
  hit_to_cap?: BisHitToCap | null;
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
  /** Set only when this band's whole weights sweep could not be trusted
   *  (`sim/cmd/leveling-bis/report.go`'s own `bandReport.WeightsReason`,
   *  `referenceMeasurementReason`) -- when present, `reference_dps_per_point` is absent and
   *  every entry in `weights` carries `insignificant: true`. Pipeline prose, never rendered
   *  verbatim (tenet 7): `panel-view.ts`'s `bandInfosFor` maps any non-empty value to the
   *  single `bisCopy.weightsUnmeasuredLine`, the same way `empty_reason` is mapped through
   *  `emptyReasonLabel` rather than printed as-is. Optional and defaults to `null`
   *  (`normaliseBisFile`, `load.ts`): a file published before this field existed carries no
   *  `weights_reason` key at all. */
  weights_reason?: string | null;
  /** The unit every stat-weight `score` on this band is in (`"reference_stat_points"`);
   *  a sim-decided row carries `sim_dps` instead of a score. */
  score_unit?: string;
  /** `true` when at least one published slot's own pick carries `sim_status ===
   *  "not_in_sim"` (`sim/cmd/leveling-bis/report.go`'s own `SetDPSPartial` doc) -- that
   *  item was silently unequipped before every sim this band's own `set_dps` was ever
   *  measured from, so `set_dps` stays a true, honestly-computed number for "this set,
   *  minus that one item," never the full set this band actually publishes. `set_dps`
   *  itself is unchanged either way -- this flag only names what it is not a measurement
   *  of. Optional and defaults to `false` (`normaliseBisFile`, `load.ts`), same pattern as
   *  `weights_reason`'s own default: a file published before this field existed carries no
   *  `set_dps_partial` key at all. */
  set_dps_partial?: boolean;
  /** The `weights` id every row's `scale_factor` was normalized against (the single
   *  highest-weighted PER-POINT stat -- never a haste stat, see `BisStatWeight.scale_factor`'s
   *  own doc) -- `sim/cmd/leveling-bis/report.go`'s own `normalizeScaleFactors`. `""`/absent
   *  when no row qualified (every row insignificant, or `weights_reason` is set); optional
   *  and defaults to `null` (`normaliseBisFile`, `load.ts`) the same way `weights_reason`
   *  does, for a file published before this lane. */
  scale_reference_stat?: string | null;
  /** A haste weight_stat's (`melee_haste`/`spell_haste`) own `scale_factor`, republished
   *  here at band level (owner correction, 2026-09-30, after player review: haste is not a
   *  per-point stat and the rail shows it as a one-line caption, never a table row) so the
   *  page never has to find and re-read the haste entry out of `weights` itself. `null`
   *  when this spec carries no haste weight_stat, or when `scale_reference_stat` itself is
   *  empty (no per-point anchor this band's own sweep trusted). Optional and defaults to
   *  `null` (`normaliseBisFile`, `load.ts`), same pattern as `weights_reason`. */
  haste_scale_factor?: number | null;
  /** Whether any candidate this band's own eligible() pass considered carries a nonzero
   *  haste stat -- `sim/cmd/leveling-bis/report.go`'s own `bandHasHasteCandidate`. Read
   *  directly (never inferred from a haste `weights` row's own `insignificant` flag, which
   *  answers a different, statistical question -- owner correction, 2026-09-30, after the
   *  caption's own doubled-suffix bug was found on screenshot review) to decide the rail's
   *  haste caption: "Haste: `<n>` per 1%", plus ", not in the table because no item at this
   *  band has it" only when this is `false`. Optional and defaults to `true`
   *  (`normaliseBisFile`, `load.ts`) for a file published before this lane -- an unknown
   *  band should not silently start showing a claim ("no item has it") an older file never
   *  made. */
  haste_on_items?: boolean;
}

export interface BisFile {
  spec: string;
  build: string;
  engine_version: string;
  generated_at: string;
  bands: BisBand[];
  /** Absent on a file published before presets. */
  presets?: Partial<Record<BisPresetId, BisPresetMeta>>;
  /** Healer files only: the incoming-damage profile every healer figure is measured under. */
  heal_profile?: BisHealProfile;
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
