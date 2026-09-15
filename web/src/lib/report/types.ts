// web/src/lib/report/types.ts
// TypeScript mirrors of the engine's report JSON. Every field name here is the Go
// `json:` tag on the matching struct in logs/engine/summary, logs/engine/store,
// logs/engine/session and logs/engine/units; anything renamed there has to be renamed
// here, and src/fixtures/report/fixture.test.ts fails when the two drift.
//
// Two Go conventions leak through and must not be "tidied":
//   * `omitempty` fields are optional here, because Go omits their zero value.
//   * event.Item carries no `json:` tags at all, so its keys are the Go field names --
//     ID, ItemLevel, Enchants, BonusIDs, Gems -- with the capitalisation Go gave them.

/** summary.Ability */
export interface Ability {
  spell_id: number;
  name: string;
  school?: number;
  total: number;
  effective: number;
  overheal?: number;
  overkill?: number;
  absorbed?: number;
  resisted?: number;
  blocked?: number;
  hits: number;
  crits: number;
  ticks: number;
  /** Miss type (ABSORB, PARRY, DODGE, …) to count. */
  misses?: Record<string, number>;
  min: number;
  max: number;
}

/** summary.Pair — one source's or target's share of a row. */
export interface Pair {
  guid: string;
  name: string;
  total: number;
}

/** Set by the whole-night fold: where each pull sits on the night's clock. */
export interface PullMark {
  label: string;
  start_ms: number;
  end_ms: number;
  kill: boolean;
}

/** summary.Actor — one row of Damage Done, Damage Taken, Healing or Healing Taken. */
export interface Actor {
  guid: string;
  name: string;
  class?: string;
  total: number;
  effective: number;
  overheal?: number;
  absorbed?: number;
  active_ms: number;
  abilities: Ability[];
  targets: Pair[];
  /** One bucket per second from the fight's start. */
  series: number[];
  /** Set by the whole-night fold: the combat time this actor was present for, which their per-second figure divides by. */
  time_ms?: number;
  /** Set on the client when the row's totals and targets were measured from the fight's events. */
  measured?: boolean;
}

/** summary.DamageRef — one damage event kept for the deaths view. */
export interface DamageRef {
  at_ms: number;
  source_guid: string;
  source_name: string;
  spell_id: number;
  spell_name: string;
  amount: number;
  overkill?: number;
  absorbed?: number;
  hp_after?: number;
  max_hp?: number;
}

/** summary.HealRef: one heal on a dying player, engine 0.2.0 and later. */
export interface HealRef {
  at_ms: number;
  source_guid: string;
  source_name: string;
  spell_id: number;
  spell_name: string;
  amount: number;
  overheal?: number;
  absorbed?: number;
}

/** summary.AuraRef */
export interface AuraRef {
  spell_id: number;
  name: string;
  source_guid?: string;
  type?: string;
  at_ms: number;
  stacks?: number;
}

/** summary.Death */
export interface Death {
  guid: string;
  name: string;
  class?: string;
  at_ms: number;
  killing_blow?: DamageRef;
  /** Set by the whole-night fold: which pull this death happened in. */
  label?: string;
  last: DamageRef[];
  /** The last heals landed on the player before the death. Absent from summaries written before engine 0.2.0. */
  heals?: HealRef[];
  auras_held: AuraRef[];
  auras_lost: AuraRef[];
  release_ms?: number;
}

/** summary.Segment — one application window of an aura. */
export interface Segment {
  start_ms: number;
  end_ms: number;
  stacks: number;
  source_guid?: string;
}

/** summary.AuraTrack — one aura on one target. `type` is BUFF or DEBUFF. */
export interface AuraTrack {
  target_guid: string;
  target_name: string;
  spell_id: number;
  name: string;
  type: string;
  applications: number;
  max_stacks: number;
  uptime_ms: number;
  segments: Segment[];
  appliers: string[];
  /** Set by the whole-night fold: the combat time this track's target existed for, which uptime divides by. */
  time_ms?: number;
}

/** summary.CastRow. `sequence` holds millisecond offsets from the fight's start. */
export interface CastRow {
  guid: string;
  name: string;
  spell_id: number;
  spell_name: string;
  started: number;
  succeeded: number;
  failed: number;
  fail_reasons?: Record<string, number>;
  cast_time_ms: number;
  sequence: number[];
}

/** summary.ExchangeRow — one interrupt or dispel relationship. `kind` is interrupt or dispel. */
export interface ExchangeRow {
  kind: string;
  source_guid: string;
  source_name: string;
  target_guid: string;
  target_name: string;
  spell_id: number;
  spell_name: string;
  extra_spell_id: number;
  extra_spell_name: string;
  count: number;
}

/** summary.ResourceTrack. `power_type` is the game's power index (0 mana, 1 rage, 3 energy). */
export interface ResourceTrack {
  guid: string;
  name: string;
  power_type: number;
  series: number[];
  gained: number;
  spent: number;
  zero_ms: number;
}

/** summary.ThreatRow. `complete` is false while the threat model admits gaps. */
export interface ThreatRow {
  guid: string;
  name: string;
  threat: number;
  model_version: string;
  complete: boolean;
}

/** event.Item, which has no json tags: these are the Go field names verbatim. */
export interface GearItem {
  ID: number;
  ItemLevel: number;
  Enchants: number[] | null;
  BonusIDs: number[] | null;
  Gems: number[] | null;
}

/** summary.CombatantRow — gear, talents and consumables at pull. */
export interface CombatantRow {
  guid: string;
  name: string;
  spec_id?: number;
  spec?: string;
  item_level?: number;
  gear: GearItem[];
  talents: number[];
  consumables: AuraRef[];
  raid_buffs: AuraRef[];
  missing_buffs: number[];
}

/** summary.RosterRow — the per-player line the Summary tab shows. */
export interface RosterRow {
  guid: string;
  name: string;
  class?: string;
  /** "combatant_info" or "inferred". */
  class_source?: string;
  spec_id?: number;
  spec?: string;
  role: string;
  item_level?: number;
  active_ms: number;
  activity_pct: number;
  deaths: number;
  damage_done: number;
  healing_done: number;
  damage_taken: number;
  dps: number;
  hps: number;
  dtps: number;
}

/** summary.Summary — reports/<id>/fights/<n>/summary.json and live.json. */
export interface Summary {
  /** Set by the whole-night fold: the pulls, in order, on the night's clock. */
  pulls?: PullMark[];
  engine_version: string;
  fight_index: number;
  duration_ms: number;
  damage_done: Actor[];
  damage_taken: Actor[];
  healing: Actor[];
  healing_taken: Actor[];
  deaths: Death[];
  auras: AuraTrack[];
  casts: CastRow[];
  interrupts: ExchangeRow[];
  dispels: ExchangeRow[];
  resources: ResourceTrack[];
  threat: ThreatRow[];
  combatants: CombatantRow[];
  roster: RosterRow[];
}

/** units.Unit */
export interface Unit {
  guid: string;
  name: string;
  /** player, creature, pet, npc, object. */
  kind: string;
  npc_id?: number;
  flags: number;
  owner_guid?: string;
  class?: string;
  class_source?: string;
  spec_id?: number;
  item_level?: number;
  first_seen: string;
  last_seen: string;
}

/** session.Health */
export interface Health {
  engine_version: string;
  layout: string;
  layout_verified: boolean;
  layout_inferred: boolean;
  advanced_logging: boolean;
  missing_header: boolean;
  lines: number;
  parse_errors: number;
  unknown_events: Record<string, number>;
  dropped_lines: number;
  clock_jumps: number;
  year_rollovers: number;
  header_restarts: number;
}

/** store.FightEntry — one fight's line in report.json. `kind` is encounter or trash. */
export interface FightEntry {
  index: number;
  kind: string;
  name: string;
  encounter_id?: number;
  difficulty?: number;
  size?: number;
  kill: boolean;
  in_progress: boolean;
  zone?: string;
  start: string;
  end: string;
  duration_ms: number;
  players: string[];
  deaths: number;
  npc_kills: number;
  /** On a wipe, the boss's health the last time the log showed it; -1 when it never did. Absent from reports written before engine 0.2.0. */
  boss_health_pct?: number;
}

/** store.Report — reports/<id>/report.json. */
export interface ReportFile {
  report_id: string;
  engine_version: string;
  health: Health;
  fights: FightEntry[];
  units: Unit[];
}

export type Visibility = 'public' | 'unlisted' | 'private' | 'guild';
export type ReportStatus = 'live' | 'complete' | 'processing' | 'failed';

/** The `data` payload of GET /v1/reports/{id}. */
export interface ReportMeta {
  id: string;
  title: string;
  visibility: Visibility;
  owner: { id: number; battletag: string } | null;
  guild?: { id: number; name: string; ruleset: string; region: string };
  zone: string;
  status: ReportStatus;
  engine_version: string;
  fights: FightEntry[];
  players: string[];
  created_at: string;
  /** Where report.json and the fight files are served from, with no trailing slash. */
  data_base_url: string;
  /** The content-phase name the API derives from `fought_at` (e.g. "launch", "raids-1"). */
  phase?: string;
}

/**
 * The engine normalises every empty slice to `[]`, so this is a guard rather than a
 * routine conversion: it keeps one malformed or hand-written file from throwing inside a
 * `.map()` deep in a component. Call it at the loader boundary, not in render code.
 */
export function asArray<T>(value: T[] | null | undefined): T[] {
  return Array.isArray(value) ? value : [];
}
