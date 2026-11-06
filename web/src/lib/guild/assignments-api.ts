// web/src/lib/guild/assignments-api.ts
// The raid assignments' contract and writes (web layout spec 2026-11-05 §7.2; Go
// guilds.RaidAssignments). Officers set a slot's names in order, clear one, and post them all to
// Discord as a reply under the raid's post. Manual only: nothing here picks a name. Same
// transport, cache and error class as api.ts; every write drops the guild's cached reads.
import { API_BASE_URL } from '../planner/config';
import { call, invalidateGuild, type GuildRaidGroup, type GuildRaidView } from './api';
import { raidPath, type RaidId } from './raids-api';

/** The roster group a slot is for, or any: orders the picker and flags an off-role pick. */
export type AssignmentRole = GuildRaidGroup | 'any';

/** A character signed In, Late or Bench, as the assignments name it. */
export interface AssignmentRaider {
  character_key: string;
  name: string;
  class: string;
  group: GuildRaidGroup;
  state: 'in' | 'late' | 'bench';
  /** The viewer's own character. */
  own: boolean;
}

/** A slot's pick; `off_role` is a warning, never a refusal. */
export interface AssignedName extends AssignmentRaider {
  off_role: boolean;
}

export interface AssignmentSlot {
  key: string;
  label: string;
  /** What the slot does in the fight, one line. */
  purpose: string;
  /** The most names the slot takes. */
  count: number;
  role: AssignmentRole;
  names: AssignedName[];
}

export interface EncounterAssignments {
  encounter_id: number;
  name: string;
  raid_zone: string;
  slots: AssignmentSlot[];
  /** One /rw line per filled slot. */
  raid_warning: string[];
}

export interface GuildRaidAssignments {
  encounters: EncounterAssignments[];
  /** The raid's name while the game client carries none of its bosses. */
  not_in_client: string[];
  /** Every character signed In, Late or Bench: the pickers' list. */
  candidates: AssignmentRaider[];
  /** RFC 3339; `null` before the first post. */
  posted_at: string | null;
}

export interface AssignmentPostResult {
  posted: boolean;
  posted_at: string | null;
  next_allowed_at: string | null;
}

const slotPath = (guildId: number, raidId: RaidId, encounterId: number, slot: string) =>
  `${raidPath(guildId, raidId)}/assignments/${encounterId}/${encodeURIComponent(slot)}`;

/** Officers: the slot's names, in order; answers the raid view. */
export async function setAssignment(
  guildId: number,
  raidId: RaidId,
  encounterId: number,
  slot: string,
  characterKeys: readonly string[],
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(slotPath(guildId, raidId, encounterId, slot), apiBase, {
    method: 'PUT',
    body: { character_keys: characterKeys },
  });
  invalidateGuild(guildId, apiBase);
  return raid;
}

/** Officers: empties the slot; answers the raid view. */
export async function clearAssignment(
  guildId: number,
  raidId: RaidId,
  encounterId: number,
  slot: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(slotPath(guildId, raidId, encounterId, slot), apiBase, {
    method: 'DELETE',
  });
  invalidateGuild(guildId, apiBase);
  return raid;
}

/** Officers: one Discord message under the raid's post listing every slot. */
export async function postAssignments(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<AssignmentPostResult> {
  const result = await call<AssignmentPostResult>(`${raidPath(guildId, raidId)}/assignments/post`, apiBase, {
    method: 'POST',
  });
  invalidateGuild(guildId, apiBase);
  return result;
}
