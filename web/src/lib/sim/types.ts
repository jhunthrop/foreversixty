// web/src/lib/sim/types.ts
// TypeScript mirrors of sim/api/envelope.go. Every key here is that file's `json:` tag;
// the contract (docs/superpowers/specs/2026-09-14-simulator-interfaces.md) makes the JSON
// names authoritative and identical in both, so anything renamed here is renamed there.
//
// There is no `raw` and no protobuf here. The contract's rule is that no protobuf crosses a
// lane boundary: `CharacterSpec` is plain JSON, and Go turns it into the engine's
// RaidSimRequest inside our own sim.wasm. Nothing outside Go ever encodes or decodes one.
import type { Summary } from '../report/types';

/** Where the character came from. */
export type SourceKind = 'armory' | 'addon' | 'build' | 'fight' | 'manual';

/** The only iteration counts the contract allows. */
export const ITERATIONS = { live: 500, normal: 3000, precise: 10_000 } as const;
export type IterationCount = (typeof ITERATIONS)[keyof typeof ITERATIONS];

export interface CharacterSource {
  kind: SourceKind;
  /** character_key | "" | build id | "<report_id>:<fight_index>". */
  ref: string;
  /** RFC3339. */
  captured_at: string;
}

export interface GearSlot {
  /** The planner's slot names: head, neck, shoulder, … , ranged. */
  slot: string;
  item_id: number;
  enchant?: number;
  suffix?: number;
}

/**
 * Everything the engine needs about the player, in JSON. `sim/request` turns it into the
 * engine's RaidSimRequest inside our own wasm; nothing outside Go touches a protobuf.
 */
export interface CharacterSpec {
  name: string;
  /** race slug. */
  race: string;
  /** class slug. */
  class: string;
  level: number;
  /** The engine's talents string, e.g. "01102123133-12312312-". */
  talents: string;
  gear: GearSlot[];
  /** Buff ids from the settings bar. */
  buffs: string[];
  consumes: string[];
  professions?: string[];
  /** When to use each major cooldown and potion. Absent means "everything on cooldown". */
  cooldowns?: CooldownSpec[];
}

/** A movement window the APL's movement conditions honour (contract 1.5). */
export interface Movement {
  interval_sec: number;
  duration_sec: number;
  /** away: out of melee, no casting. casting: spells interrupted, melee continues. */
  kind: 'away' | 'casting';
}

/** A step in the target-count timeline. Overrides `targets` when the list is non-empty. */
export interface TargetCount {
  at_sec: number;
  count: number;
}

export const TARGET_TYPE_IDS = [
  'humanoid',
  'undead',
  'beast',
  'demon',
  'dragonkin',
  'elemental',
  'giant',
  'mechanical',
  'unknown',
] as const;
export type TargetType = (typeof TARGET_TYPE_IDS)[number];

export interface EncounterSpec {
  duration_sec: number;
  variation: number;
  targets: number;
  execute_ratio: number;
  /** "" | "patchwerk" | "encounter:<encounter_id>". */
  profile: string;
  /** The fight style's id: a label only. The fields above and below are what the engine reads. */
  style?: string;
  movement?: Movement;
  targets_over_time?: TargetCount[];
  /** 60..63; 63 is the default the engine assumes when this is absent. */
  target_level?: number;
  /** 0 means the level's preset, resolved inside the engine. */
  target_armor?: number;
  target_type?: TargetType | '';
  /** No debuffs, no execute, no armor reduction. */
  dummy?: boolean;
}

/** When to use a cooldown. Empty `at_sec` means "on cooldown" (contract 1.7). */
export interface CooldownSpec {
  /** "spell:<id>" or a consumable id from IDS.md. */
  id: string;
  at_sec: number[];
}

export const DEFAULT_ENCOUNTER: EncounterSpec = {
  duration_sec: 180,
  variation: 0.2,
  targets: 1,
  execute_ratio: 0.25,
  profile: '',
};

export interface SimRequest {
  engine_version: string;
  /** spec_slug. */
  spec: string;
  source: CharacterSource;
  character: CharacterSpec;
  encounter: EncounterSpec;
  /** With `target_error` set this is the ceiling, not the count. */
  iterations: number;
  /** 0 means random; paired runs set it. */
  random_seed: number;
  /**
   * When > 0, the run continues in `STEP_ITERATIONS_DEFAULT` steps until
   * `dps.error / dps.mean` is at or under this, or `iterations` is reached. 0 is a
   * fixed-count run. Contract 1.2.
   */
  target_error?: number;
  bulk?: BulkSpec;
  weights?: WeightsSpec;
}

/** Contract 1.2. The size of one step of a target-error run. */
export const STEP_ITERATIONS_DEFAULT = 1000;

export interface Candidate {
  /** IDS.md slot vocabulary; "" means "wherever it fits" (rings, trinkets, weapons). */
  slot: string;
  item_id: number;
  /** 0 inherits the equipped enchant for the slot where it fits. */
  enchant?: number;
  suffix?: number;
  /** equipped | bag | bank | search | drop:<source-id> | set:<name>. */
  origin: string;
  /**
   * Contract A6: the human name of where it came from ("Ragnaros"), which the page fills
   * from `loot.json` and the API's headline reads back off the substitution.
   */
  source_name?: string;
}

export interface TalentLoadout {
  name: string;
  talents: string;
}

export interface GearSet {
  name: string;
  gear: GearSlot[];
}

export interface BulkSpec {
  /**
   * gear | talents | drops. Contract A4: the mode decides the expansion and the design's
   * `combinations` boolean is gone -- `gear` takes the product of every candidate group,
   * `drops` and `talents` one substitution at a time.
   */
  mode: string;
  candidates: Candidate[];
  talents?: TalentLoadout[];
  sets?: GearSet[];
  /**
   * Contract A5: alternative consumable lists tried as candidates in `gear` mode. Each
   * inner list replaces `CharacterSpec.Consumes` for that combination.
   */
  consumables?: string[][];
  /** Slots never substituted. */
  locked?: string[];
  /** fast | normal | high. */
  precision: string;
  /** The lane's cap, echoed so a saved request says what bounded it. */
  cap: number;
}

export interface WeightsSpec {
  stats: string[];
  /** The stat normalised to 1.0. */
  reference: string;
}

export interface Estimate {
  mean: number;
  stddev: number;
  /** Standard error of the mean. */
  error: number;
  min: number;
  max: number;
}

/**
 * What `simRun` reports through its progress callback: the two fields of a `SimResult` the
 * pool needs to pool an estimate. The contract pins this to `Pick<SimResult,
 * 'iterations_run' | 'dps'>` and `sim/cmd/wasm/main.go` builds exactly that shape — settled,
 * not open.
 */
export type SimProgressUpdate = Pick<SimResult, 'iterations_run' | 'dps'>;

export interface Substitution {
  /**
   * item | talents | set | consumes. Contract 10.8 adds `consumes`: `sim/bulk` emits one
   * per combination that used an alternative consumable list (`BulkSpec.consumables`),
   * and `name` is that list's ids joined by ", ".
   */
  kind: string;
  slot?: string;
  item_id?: number;
  enchant?: number;
  suffix?: number;
  /**
   * The loadout or set name — and, per contract A6, an item's name too, filled from
   * simdb, so a combo row reads without a second lookup. For a `consumes` substitution
   * (10.8) it is the consumable ids joined by ", ".
   */
  name?: string;
  talents?: string;
  origin?: string;
  /** Contract A6: copied from the candidate. */
  source_name?: string;
}

export interface Combo {
  substitutions: Substitution[];
  dps: Estimate;
  /** Against `equipped`, paired at the same stage. */
  delta: Estimate;
  /** 0 for the leader's within-error group, then 1, 2, … */
  group: number;
}

export interface Stage {
  iterations: number;
  combos: number;
}

export interface StatWeight {
  stat: string;
  /** The reference stat is exactly 1. */
  weight: number;
  error: number;
}

/**
 * One cast of the median-DPS iteration. `at_ms` is negative during the pre-pull.
 *
 * Contract A12: the row carries the summary's own action-key form (`spell:23881`,
 * `item:13503`, `other:melee`) and nothing else -- no display name and no spell id. The
 * page resolves the name with `resolveActionName`, exactly as it already does for every
 * cast row, so the sample table can never disagree with the cast table about what an
 * action is called.
 */
export interface SampleCast {
  at_ms: number;
  action: string;
  target?: string;
  /** rage, energy, mana, combo_points … after the cast. */
  resources?: Record<string, number>;
}

export interface SimResult {
  sim_id?: string;
  engine_version: string;
  request: SimRequest;
  lane: 'browser' | 'server';
  dps: Estimate;
  iterations_run: number;
  /** Wall clock of the run. */
  duration_ms: number;
  summary: Summary;
  error?: string;
  /**
   * Stopped on request, not a failure; `summary`/`dps` are partial. Additive
   * (`aborted,omitempty` on `sim/api/envelope.go`'s `SimResult`) -- mirrored here because
   * Task 17's save flow refuses to save an aborted result rather than storing a partial
   * run under a player-chosen title.
   */
  aborted?: boolean;
  /** Ranked, best first. Bulk kinds only. */
  combos?: Combo[];
  /** The base character at the final stage. Bulk kinds only. */
  equipped?: Estimate;
  stages?: Stage[];
  weights?: StatWeight[];
  /** One iteration's casts, the median-DPS one. */
  sample?: SampleCast[];
}

/** One row of GET /v1/sims?mine=1. */
export interface SimListRow {
  sim_id: string;
  spec: string;
  dps: number;
  engine_version: string;
  created_at: string;
  title: string;
  /** run | gear | drops | talents | weights. Absent on a row saved before migration 0014. */
  kind?: string;
  /** The API's own one-line summary, e.g. "+41 DPS from Vis'kag". */
  headline?: string;
}

export interface SimListPage {
  rows: SimListRow[];
  total: number;
  page: number;
  per_page: number;
}

export interface SimProgress {
  state: 'queued' | 'running' | 'done' | 'error';
  iterations_done: number;
  dps?: number;
  stage?: number;
  combos_done?: number;
  combos_total?: number;
}

export type SpecState = 'validated' | 'in_progress' | 'unsupported';

export interface SpecFidelity {
  spec: string;
  state: SpecState;
  /** Median DPS gap as a fraction; null until the validation job has run. */
  median_gap: number | null;
  parses: number;
  /**
   * spell_id is the summary's row identity -- the same client id `compare.ts`'s two-tier
   * join keys on -- and is what a future "link this action" feature would join a sim cast
   * to a parsed one on. Unread today.
   */
  worst_actions: { spell_id: number; name: string; sim_casts: number; actual_casts: number }[];
  /** `specs.go` coalesces this into a plain Go string; it is never null. */
  engine_version: string;
  /** Null for a card nothing has measured yet (`specs.go`'s UpdatedAt *time.Time). */
  updated_at: string | null;
}

/** GET /v1/characters/{character_key}/sim-input. */
export interface SimInput {
  spec: string;
  /**
   * The source's own opaque shape, never the planner's slot-to-item map: `input.go`'s
   * `Gear json.RawMessage` is the addon's own export JSON for an addon-sourced read, or
   * `{"trinkets": […]}` for a fight-sourced one -- two different, source-dependent shapes
   * neither this file nor `openapi.yaml` (`gear: { type: object }`) pins down further.
   * `sources.ts` does not attempt to decode it (see H3 in the final whole-branch review).
   */
  gear: unknown;
  /**
   * The fight's recorded talent split -- points per tree, e.g. `"31/0/20"` -- never a
   * per-talent order (`fight_metrics.talent_split`; `input.go`: "the addon export is an
   * opaque string this repository never parses, so a character who has never parsed has
   * none"). Empty when nothing has recorded one.
   */
  talents: string;
  /**
   * Buff ids, in the engine's own vocabulary (`sim/request/IDS.md`) -- NOT spell ids. The
   * amended contract puts the mapping on the API side: "the web never maps spell ids
   * itself". So these travel into `CharacterSpec.buffs` unchanged, and an id the engine
   * refuses surfaces as the engine's own error rather than being silently dropped.
   */
  buffs: string[];
  /**
   * The character's race slug, when the API has one. The contract's `sim-input` row does
   * not list it and nothing records a race today: an addon export carries one in its FS1
   * string, a combat log does not, and there is no Armory. It is optional here rather than
   * absent because a race is not guessable -- Forever's racials are two actives and two
   * passives each -- so `fromStoredCharacter` refuses rather than substituting one, and
   * the day the companion records it this field is the whole change. See the Rulings
   * table, "the stored character has no race".
   */
  race?: string;
  captured_at: string;
  /**
   * Which source the API actually had. Today that is `"addon"` or `"fight"`, newest wins:
   * nothing in the repo stores an Armory refresh yet, so `"armory"` appears here only when
   * Blizzard's profile API for Forever exists. The shape does not change when it does.
   */
  source: SourceKind;
}
