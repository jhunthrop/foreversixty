// web/src/lib/guild/overview-api.ts
// The role-aware Overview's reads and writes (web layout spec 2026-11-05 section 7.1, 7.4): the
// member's private checklist, the officer's action counts and recent attendance, the comp target
// and the Discord ping. Same transport, cache and error class as api.ts; every write drops the
// guild's cached reads through `invalidateGuild`.
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import { GUILD_TTL_MS, call, invalidateGuild, type GuildKeyItem, type GuildRaidView } from './api';
import { raidPath, type RaidId } from './raids-api';

/** Seats per role (Go guilds.CompTarget). `default` is the zone's, `raid` an officer's edit. */
export interface GuildCompTarget {
  tanks: number;
  healers: number;
  melee: number;
  ranged: number;
  source: 'default' | 'raid';
}

export type GuildCompRole = 'tanks' | 'healers' | 'melee' | 'ranged';

/** The Late rows per role, beside the In-only role counts. */
export type GuildRoleCounts = Readonly<Record<GuildCompRole, number>>;

/** A raider with no sign-up for this raid, officers only (one row per account, its main). */
export interface GuildNotSignedRaider {
  character_key: string;
  name: string;
  class: string;
  account_key: string;
  discord_linked: boolean;
}

/** When the officer ping last went out and when the next one may (6 h throttle). */
export interface GuildRaidPingState {
  last_pinged_at: string | null;
  next_allowed_at: string | null;
}

export interface GuildPingResult {
  posted: boolean;
  mentioned: number;
  no_discord: number;
  no_discord_names: string[];
  next_allowed_at: string | null;
}

export type ChecklistState = 'ok' | 'fail' | 'not_tracked';
export type ChecklistKey =
  'level' | 'attunement' | 'gear_gap' | 'enchants' | 'flask_food' | 'fire_resistance';
export type ChecklistVerdict = 'ready' | 'almost' | 'getting_started';

export interface ChecklistItem {
  key: ChecklistKey;
  state: ChecklistState;
  reason?: string;
  level?: number;
  required_level?: number;
  key_item?: GuildKeyItem;
  /** What attunement asks for: the zone's key, or its quest chain (key_item is then the chain's item). */
  attunement_kind?: 'key' | 'quest';
  upgrades?: number;
  gain_dps?: number;
  unit?: string;
  slots?: string[];
  missing_slots?: string[];
  /** flask_food: the counted stacks in the bags, and what the minimum lacks. */
  supplies?: { id: number; name: string; kind: 'flask' | 'elixir' | 'food'; count: number }[];
  missing?: ('flask_or_elixir' | 'food')[];
  /** fire_resistance: the total. */
  resistance?: number;
}

export interface GuildChecklist {
  character: { character_key: string; name: string; class: string; spec: string; level: number | null };
  raid: { id: number; zone_key: string; zone_name: string } | null;
  verdict: ChecklistVerdict;
  items: ChecklistItem[];
}

export type AttendanceCell = 'present' | 'absent' | 'bench' | 'none';

export interface GuildRecentAttendance {
  raids: { id: number; starts_at: string; zone_name: string }[];
  rows: { name: string; class: string; account_key: string; cells: AttendanceCell[]; slipping: boolean }[];
}

export interface GuildOfficerActions {
  approvals: { count: number; names: string[] };
  /** Pending applications, the applicants' characters oldest first. Absent on an older API. */
  applications?: { count: number; names: string[] };
  not_signed: { raid_id: number | null; count: number; names: string[] };
  /** The loot council's newest night with drops (null with none): `count` is its drops still open.
   *  `tracked` is always true now; the row hides at a count of 0. */
  drops_to_award: {
    tracked: boolean;
    count: number | null;
    raid_id: number | null;
    label: string;
    awarded: number;
  } | null;
}

const guildPath = (guildId: number) => `/v1/guilds/${guildId}`;

function privateRead<T>(path: string, apiBase: string): Promise<T> {
  return query<T>(`${apiBase}${path}`, () => call<T>(path, apiBase), {
    scope: 'private',
    ttlMs: GUILD_TTL_MS,
    revalidateOnRead: true,
  });
}

/** The signed-in member's own checklist for a raid (the next one when `raidId` is omitted). */
export function fetchMyChecklist(
  guildId: number,
  raidId: RaidId | null = null,
  apiBase: string = API_BASE_URL,
): Promise<GuildChecklist> {
  const raid = raidId === null ? '' : `?raid=${encodeURIComponent(String(raidId))}`;
  return privateRead<GuildChecklist>(`${guildPath(guildId)}/me/checklist${raid}`, apiBase);
}

/** Officers: the slipping and benched raiders over the last `nights` raids. */
export function fetchRecentAttendance(
  guildId: number,
  nights = 4,
  apiBase: string = API_BASE_URL,
): Promise<GuildRecentAttendance> {
  return privateRead<GuildRecentAttendance>(`${guildPath(guildId)}/attendance/recent?n=${nights}`, apiBase);
}

/** Officers: what is waiting on them (approvals, non-signers, drops). */
export function fetchOfficerActions(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildOfficerActions> {
  return privateRead<GuildOfficerActions>(`${guildPath(guildId)}/officer/actions`, apiBase);
}

/** Officers: one Discord message under the raid's post naming the non-signers. */
export async function pingRaid(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<GuildPingResult> {
  const result = await call<GuildPingResult>(`${raidPath(guildId, raidId)}/ping`, apiBase, {
    method: 'POST',
  });
  invalidateGuild(guildId, apiBase);
  return result;
}

/** Officers: this raid's own seats per role; answers the raid view. */
export async function setCompTarget(
  guildId: number,
  raidId: RaidId,
  target: Readonly<Record<GuildCompRole, number>>,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(`${raidPath(guildId, raidId)}/comp-target`, apiBase, {
    method: 'PUT',
    body: target,
  });
  invalidateGuild(guildId, apiBase);
  return raid;
}

/** Officers: back to the zone's default seats; answers the raid view. */
export async function resetCompTarget(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(`${raidPath(guildId, raidId)}/comp-target`, apiBase, {
    method: 'DELETE',
  });
  invalidateGuild(guildId, apiBase);
  return raid;
}
