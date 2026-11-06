// web/src/lib/guild/loot-council-api.ts
// The loot council's reads and writes (web layout spec 2026-11-05 §7.2, api/openapi.yaml
// loot/nights, loot/drops/{item_id}/candidates, loot/votes, loot/interest, loot/awards,
// export/loot.csv|json). A drop is keyed by night ("raid:<id>", "report:<id>", "live") and item.
// Same transport, cache and error class as api.ts; every write drops the guild's cached reads.
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import { GUILD_TTL_MS, call, invalidateGuild, type GuildRosterAttendance } from './api';

export type MsOs = 'ms' | 'os';
export type Vote = 'yes' | 'no';
export type DropKind = 'gear' | 'quest' | 'bag' | 'skinned';

/** An award's decision record. The vote counts are null for a member and on an old award. */
export interface LootDecision {
  award_id: number;
  item_id: number;
  item_name: string;
  icon: string;
  quality: number;
  character_key: string;
  name: string;
  class: string;
  ms_os: MsOs | null;
  reason: string | null;
  votes_yes: number | null;
  votes_no: number | null;
  votes_officers: number | null;
  /** The officer's character; "an officer" when unnamed, "" when none is recorded. */
  awarded_by: string;
  at: string;
}

export interface LootNightDrop {
  item_id: number;
  name: string;
  icon: string;
  quality: number;
  kind: DropKind;
  encounter_id: number;
  boss: string;
  state: 'open' | 'awarded';
  award: LootDecision | null;
}

export interface LootNight {
  night: string;
  raid_id: number | null;
  report_id: string | null;
  /** The night's day in the guild's raid zone, YYYY-MM-DD. */
  date: string;
  at: string;
  zone: string;
  /** Where an award made now is filed; the page offers the whole loot table on it. */
  current: boolean;
  drops: LootNightDrop[];
  decisions: LootDecision[];
}

export interface LootNights {
  officers: number;
  nights: LootNight[];
}

export interface LootCouncilVotes {
  yes: number;
  no: number;
  /** The officers who can vote. */
  voters: number;
  mine: Vote | null;
}

export interface LootCouncilCandidate {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  gain_dps: number | null;
  gain_unit: 'dps' | 'hps' | 'tps' | null;
  tier: 0 | 1;
  not_sim_checked: boolean;
  ilvl_delta: number | null;
  attendance: GuildRosterAttendance;
  already_equivalent: boolean;
  fills_empty_slot: boolean;
  replaces: string;
  replaces_item: { item_id: number; name: string; icon: string; quality: number } | null;
  items_won: number;
  bis_for_them: boolean;
  ms_os: MsOs | null;
  ms_os_source: 'declared' | 'derived' | null;
  votes: LootCouncilVotes;
}

export interface LootCouncilDrop {
  night: string;
  item_id: number;
  name: string;
  icon: string;
  quality: number;
  kind: DropKind;
  officers: number;
  /** Awarded on the night: votes and flags are closed. */
  closed: boolean;
  award: LootDecision | null;
  candidates: LootCouncilCandidate[];
}

export interface LootInterestFlag {
  item_id: number;
  character_key: string;
  ms_os: MsOs;
  set_by_officer: boolean;
}

export interface LootInterest {
  night: string;
  flags: LootInterestFlag[];
}

const lootPath = (guildId: number): string => `/v1/guilds/${guildId}/loot`;

function privateRead<T>(path: string, apiBase: string): Promise<T> {
  return query<T>(`${apiBase}${path}`, () => call<T>(path, apiBase), {
    scope: 'private',
    ttlMs: GUILD_TTL_MS,
    revalidateOnRead: true,
  });
}

/** Every verified member: the nights, newest first (a member's copy has no vote counts). */
export function fetchLootNights(guildId: number, apiBase: string = API_BASE_URL): Promise<LootNights> {
  return privateRead<LootNights>(`${lootPath(guildId)}/nights`, apiBase);
}

/** Officers: one drop of one night with its ranked candidates, votes and MS/OS. */
export function fetchDropCandidates(
  guildId: number,
  itemId: number,
  night: string,
  apiBase: string = API_BASE_URL,
): Promise<LootCouncilDrop> {
  const qs = `?night=${encodeURIComponent(night)}`;
  return privateRead<LootCouncilDrop>(`${lootPath(guildId)}/drops/${itemId}/candidates${qs}`, apiBase);
}

/** The viewer's own MS/OS flags on a night. */
export function fetchLootInterest(
  guildId: number,
  night: string,
  apiBase: string = API_BASE_URL,
): Promise<LootInterest> {
  return privateRead<LootInterest>(
    `${lootPath(guildId)}/interest?night=${encodeURIComponent(night)}`,
    apiBase,
  );
}

export interface DropTarget {
  night: string;
  item_id: number;
  character_key: string;
}

/** An officer's yes or no; null removes it. Answers the candidate's tally. */
export async function putLootVote(
  guildId: number,
  target: DropTarget,
  vote: Vote | null,
  apiBase: string = API_BASE_URL,
): Promise<LootCouncilVotes> {
  const data = await call<{ votes: LootCouncilVotes }>(`${lootPath(guildId)}/votes`, apiBase, {
    method: 'PUT',
    body: { ...target, vote },
  });
  invalidateGuild(guildId, apiBase);
  return data.votes;
}

/** A member's own MS or OS on a drop; null removes it. */
export async function putLootInterest(
  guildId: number,
  target: DropTarget,
  msOs: MsOs | null,
  apiBase: string = API_BASE_URL,
): Promise<MsOs | null> {
  const data = await call<{ ms_os: MsOs | null }>(`${lootPath(guildId)}/interest`, apiBase, {
    method: 'PUT',
    body: { ...target, ms_os: msOs },
  });
  invalidateGuild(guildId, apiBase);
  return data.ms_os;
}

export interface CouncilAward extends DropTarget {
  encounter_id: number;
  /** Null: the candidate's flag, else the derivation rule (a quest drop has none). */
  ms_os: MsOs | null;
  /** Trimmed; null when the officer gave none. */
  reason: string | null;
}

/** Officers: award a drop with its decision record; closes voting on it for the night. */
export async function awardDrop(
  guildId: number,
  award: CouncilAward,
  apiBase: string = API_BASE_URL,
): Promise<{ award_id: number }> {
  const data = await call<{ award_id: number }>(`${lootPath(guildId)}/awards`, apiBase, {
    method: 'POST',
    body: award,
  });
  invalidateGuild(guildId, apiBase);
  return data;
}

/** The history download, absolute to the API origin; from and to are YYYY-MM-DD, inclusive. */
export function lootHistoryHref(
  guildId: number,
  format: 'csv' | 'json',
  range: { from: string; to: string },
  apiBase: string = API_BASE_URL,
): string {
  const params = new URLSearchParams();
  if (range.from !== '') params.set('from', range.from);
  if (range.to !== '') params.set('to', range.to);
  const qs = params.toString();
  return `${apiBase}/v1/guilds/${guildId}/export/loot.${format}${qs === '' ? '' : `?${qs}`}`;
}
