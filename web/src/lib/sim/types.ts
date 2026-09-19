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
}

export interface EncounterSpec {
  duration_sec: number;
  variation: number;
  targets: number;
  execute_ratio: number;
  /** "" | "patchwerk" | "encounter:<encounter_id>". */
  profile: string;
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
  iterations: number;
  /** 0 means random; paired runs set it. */
  random_seed: number;
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
 * pool needs to pool an estimate. Not a new shape — confirm it with the engine lane.
 */
export type SimProgressUpdate = Pick<SimResult, 'iterations_run' | 'dps'>;

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
}

/** One row of GET /v1/sims?mine=1. */
export interface SimListRow {
  sim_id: string;
  spec: string;
  dps: number;
  engine_version: string;
  created_at: string;
  title: string;
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
}

export type SpecState = 'validated' | 'in_progress' | 'unsupported';

export interface SpecFidelity {
  spec: string;
  state: SpecState;
  /** Median DPS gap as a fraction; null until the validation job has run. */
  median_gap: number | null;
  parses: number;
  worst_actions: { name: string; sim_casts: number; actual_casts: number }[];
  engine_version: string | null;
  updated_at: string;
}

/** GET /v1/characters/{character_key}/sim-input. */
export interface SimInput {
  spec: string;
  /** Slot name to item id, the planner's own Gear shape. */
  gear: Record<string, number>;
  /** Talent ids in the order the points were spent. */
  talents: number[];
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
