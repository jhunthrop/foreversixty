// web/src/lib/guild/attendance-contract.ts
// The attendance lane's API types (design/specs/2026-10-27-guild-attendance.md, "Contract this
// spec assumes", with the deltas of the API as built: `moves.previous` carries `id`, readiness
// rows carry `no_export`). The key names are the API's own; api.ts re-exports them and
// `GuildRaidView` carries them.
import type { GuildRaidRole, GuildSignupState } from './api';
import type { AwardRecord } from './loot-council-api';

export type RaidPreDoneStatus = 'planned' | 'locked';
export type RaidMatchHow = 'auto' | 'hand';
export type RaidAttendanceOutcome = 'present' | 'late' | 'absent' | 'no_show';
export type ConsumableBefore = 'stocked' | 'short' | 'unknown';
export type MoveMetric = 'dps' | 'hps';

/** A raid named by id, as the API's shared `RaidRefView` writes it. */
export interface RaidRef {
  id: number;
  title: string;
  starts_at: string;
}

/** The log a raid is tied to. */
export interface RaidMatch {
  report_id: string;
  zone_name: string;
  uploaded_at: string;
  uploaded_by_name: string;
  pulls: number;
  kills: number;
  duration_ms: number;
  how: RaidMatchHow;
}

/** The night's figures; they name nobody, so every viewer gets them. */
export interface RaidNight {
  signed: number;
  came: number;
  late: number;
  not_in_log: number;
  joined: number;
  excused: number;
  kills: number;
  wipes: number;
  deaths: number;
  duration_ms: number;
}

/** One character's night (officers: every row; a member: their own). */
export interface RaidAttendanceRow {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  signed: GuildSignupState | null;
  role: GuildRaidRole | null;
  outcome: RaidAttendanceOutcome;
  /** The name of the account's other character that was in the log (ruling R-A2). */
  came_as: string | null;
  first_seen_offset_s: number | null;
  pulls_present: number;
  pulls_total: number;
  /** The effective excused flag (an officer's fix included). */
  excused: boolean;
  /** The viewer's own character. */
  own: boolean;
  /** What the log and the sign-ups said, before any officer's fix; `outcome` is the effective one. */
  auto_outcome: RaidAttendanceOutcome;
  auto_excused: boolean;
  /** An officer's manual fix of this night, or null. */
  override: RaidAttendanceOverride | null;
}

/** The statuses an officer can set over the log (bench counts as attended). */
export type AttendanceOverrideStatus = 'present' | 'absent' | 'bench' | 'excused';

export interface RaidAttendanceOverride {
  status: AttendanceOverrideStatus;
  /** The officer's main character name, or `an officer`. */
  by_name: string;
  at: string;
}

export interface RaidKill {
  encounter_id: number;
  name: string;
  fight_index: number;
  killed: boolean;
  kill_ms: number | null;
  /** For a boss never killed: every pull. */
  wipes_before: number;
  deaths: number;
}

export interface RaidUnnamed {
  pulls: number;
  kills: number;
  wipes: number;
}

export interface RaidDeathRow {
  character_key: string;
  name: string;
  class: string;
  on_wipes: number;
  on_kills: number;
}

export interface RaidConsumableRow {
  character_key: string;
  name: string;
  class: string;
  before: ConsumableBefore;
  flask: boolean | null;
  food: boolean | null;
  potions: number | null;
}

/** Why a pull's consumables cannot be read: no character sheet in the log, or not read yet. */
export type PullConsumablesReason = 'no_character_sheet' | 'not_read';

/** One raider at one boss pull: true up (or applied within 10 s), false proven missing (flasks
 *  only today), null the log cannot say. */
export interface RaidPullConsumable {
  character_key: string;
  name: string;
  class: string;
  flask: boolean | null;
  food: boolean | null;
  weapon_enchant: boolean | null;
  /** Combat potions used in the pull. */
  potions: number;
  /** The names of what was up. */
  brought: string[];
}

export interface RaidPullConsumables {
  fight_index: number;
  name: string;
  attempt: number;
  reason: PullConsumablesReason | null;
  /** Null with a reason: nothing can be said (unknown, never missing). */
  rows: RaidPullConsumable[] | null;
}

export interface RaidConsumables {
  /** True when at least one boss pull carried a character sheet to read. */
  log_has_buffs: boolean;
  rows: RaidConsumableRow[];
  by_pull: RaidPullConsumables[];
}

/** The night's answer figures; they name nobody. */
export interface RaidReportHeadline {
  kills: number;
  wipes: number;
  /** Boss pulls: kills plus wipes. */
  attempts: number;
  /** Every pull's deaths, trash included. */
  deaths: number;
}

export type PullResult = 'kill' | 'wipe' | 'clear';

/** One damage event on a raider before their death. */
export interface RaidReportHit {
  /** Time into the pull. */
  at_ms: number;
  /** 1 is a melee swing. */
  spell_id: number;
  ability: string;
  source: string;
  amount: number;
  overkill: number;
}

export interface RaidReportDeath {
  /** Empty when the log name maps to no character of the night. */
  character_key: string;
  name: string;
  class: string;
  /** Time into the pull. */
  at_ms: number;
  killing_blow: RaidReportHit | null;
  /** Up to the last 3 hits before the death, oldest first. */
  last_hits: RaidReportHit[];
}

export interface RaidReportPull {
  fight_index: number;
  /** Null on trash. */
  encounter_id: number | null;
  name: string;
  /** The boss pull's number; null on trash. */
  attempt: number | null;
  result: PullResult;
  duration_ms: number;
  deaths: number;
  /** Null until the fight's summary has been read; a member reads their own only. */
  death_causes: RaidReportDeath[] | null;
}

/** One raider on one boss kill: DPS, HPS for a healer, and the fight's rating when computed. */
export interface RaidReportParse {
  character_key: string;
  name: string;
  class: string;
  role: string;
  encounter_id: number;
  encounter: string;
  fight_index: number;
  metric: MoveMetric;
  value: number;
  rating: number | null;
  /** Value's place among every ranked kill of the encounter for the spec and metric, 0 to 100; null with no digest yet. */
  percentile?: number | null;
}

/** One awarded item on the night report, with the council's decision record (api AwardRecord). */
export interface RaidLootAward extends AwardRecord {
  item_id: number;
  item_name: string;
  quality: number;
  /** The icon file name from the Loot tab's loot table; "" when the build has none. */
  icon: string;
  character_key: string;
  name: string;
  class: string;
  by_name: string;
  at: string;
}

export interface RaidLootGroup {
  encounter_id: number;
  boss: string;
  awards: RaidLootAward[];
}

export interface RaidMoveRow {
  character_key: string;
  name: string;
  class: string;
  metric: MoveMetric;
  encounter: string;
  before: number;
  after: number;
}

export interface RaidMoves {
  previous: RaidRef | null;
  rows: RaidMoveRow[];
}

/** `null` on the raid view until a log is matched. */
export interface RaidReport {
  headline: RaidReportHeadline;
  /** Every fight of the night in order. */
  pulls: RaidReportPull[];
  parses: RaidReportParse[];
  kill_order: RaidKill[];
  unnamed: RaidUnnamed | null;
  deaths: RaidDeathRow[];
  consumables: RaidConsumables;
  loot: RaidLootGroup[];
  moves: RaidMoves;
  /** RFC 3339: when the summary was last posted to Discord. */
  posted_at: string | null;
}

/** The raid view's attendance keys (`GuildRaidView` extends this). */
export interface RaidAttendanceFields {
  pre_done_status: RaidPreDoneStatus | null;
  match: RaidMatch | null;
  night: RaidNight | null;
  attendance: RaidAttendanceRow[] | null;
  report: RaidReport | null;
}

/** One log the match picker lists. */
export interface MatchCandidate {
  report_id: string;
  zone_name: string;
  uploaded_at: string;
  uploaded_by_name: string;
  pulls: number;
  kills: number;
  duration_ms: number;
  closest: boolean;
  /** Set when another raid already holds this log (ruling R-A7). */
  matched_raid: RaidRef | null;
}

export interface MatchCandidates {
  zone_known: boolean;
  rows: MatchCandidate[];
}

export interface PostReportResult {
  posted_at: string;
}

/** The filtered readiness read's raid (`?raid={id}`). */
export interface ReadinessRaid {
  id: number;
  title: string;
  starts_at: string;
  cutoff_at: string;
  /** The zone has an attunement (a key or a quest chain), which each row's `attuned` reads. */
  attunement?: boolean;
  /** The raid asks for fire resistance (Onyxia's Lair). */
  fire_resistance?: boolean;
}

/** The unfiltered readiness read's next raid. */
export interface ReadinessNextRaid {
  id: number;
  title: string;
  zone_name: string;
  starts_at: string;
}

/** The raid filter's readiness row keys (`GuildReadinessRow` extends this). */
export interface ReadinessRaidRowFields {
  /** A signed character with no export: key off this, never the zero `logged_at`. */
  no_export: boolean;
  /** `null` on a `no_export` row and on the unfiltered read. */
  synced_before_cutoff: boolean | null;
}

/** Roster `attendance_source`. */
export type AttendanceSource = 'raids' | 'reports';
