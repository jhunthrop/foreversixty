// web/src/lib/guild/attendance-contract.ts
// The attendance lane's API types (design/specs/2026-10-27-guild-attendance.md, "Contract this
// spec assumes", with the deltas of the API as built: `moves.previous` carries `id`, readiness
// rows carry `no_export`). The key names are the API's own; api.ts re-exports them and
// `GuildRaidView` carries them.
import type { GuildRaidRole, GuildSignupState } from './api';

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
  excused: boolean;
  /** The viewer's own character. */
  own: boolean;
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

export interface RaidConsumables {
  /** Always false today: the log carries no buff events. */
  log_has_buffs: boolean;
  rows: RaidConsumableRow[];
}

export interface RaidLootAward {
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
  kill_order: RaidKill[];
  unnamed: RaidUnnamed | null;
  deaths: RaidDeathRow[];
  consumables: RaidConsumables;
  loot: RaidLootGroup[];
  moves: RaidMoves;
  /** RFC 3339 once an officer posted it to Discord; never posts twice. */
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
