// web/src/lib/guild/api.ts
// Every call the browser makes to the guild half of the API (spec section 2.6), reconciled
// field-exact against the real, landed api/internal/guilds handlers (2026-09-21 plan's
// reconciliation record). Follows rankings/api.ts's pattern: the shared transport
// (account/api.ts's requestEnvelope), a module-local error class, and typed reads/writes
// with no bespoke fetch() call anywhere else in this lane's own code.
import { AccountError, requestEnvelope, type EnvelopeResult } from '../account/api';
import { API_BASE_URL } from '../planner/config';

export const GUILD_API_FAILED = 'That did not work; try again';

export class GuildApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'GuildApiError';
  }
}

export type GuildRank = 'member' | 'officer' | 'leader';
export type GuildConsent = 'roster' | 'gear' | 'gear_bags';
export type GuildVisibility = 'public' | 'unlisted' | 'guild';

export interface GuildSummary {
  id: number;
  region: string;
  ruleset: string;
  name: string;
}

export type ClaimState = 'unclaimed' | 'pending' | 'claimed' | 'contested';

/**
 * GET .../home and GET .../settings both expose this (2026-09-21 security amendment).
 * `frozen` was added by a second, concurrent security review (same date): "contested" and
 * "frozen" are distinct -- a contest against an established (14+ day old) claim with
 * independently log-corroborated members is recorded as `state: 'contested'` but
 * `frozen: false`, so officer tools stay enabled and the claim is only queued for a
 * moderator. Only a young or uncorroborated contested claim is `frozen: true`. `frozen` is
 * meaningless (always `false`) at any state other than `'contested'`. Consumers must gate
 * "disable officer controls" on `frozen`, never on `state === 'contested'` alone.
 */
export interface ClaimStateView {
  state: ClaimState;
  since?: string;
  frozen: boolean;
}

export interface MemberRef {
  battletag: string;
}

export interface ClaimPendingView {
  by: MemberRef;
  expires_at: string;
}

/** One report row. No `zone`: HomeReport carries none (plan's reconciliation record,
 *  item 2) — an empty title falls back to guildHomeCopy.untitledReport, not a zone. */
export interface GuildHomeReport {
  id: string;
  title: string;
  created_at: string;
  fight_count: number;
  kill_count: number;
}

/**
 * One roster row. No `user_id`: the real API's RosterRow never sends one (plan's
 * reconciliation record, item 3) — "is this my own row" is derived from
 * `Me.characters[].key` instead, which already exact-matches `character_key`.
 */
export interface GuildRosterRow {
  character_key: string;
  region: string;
  ruleset: string;
  name: string;
  class?: string;
  spec?: string;
  rank: GuildRank;
  verified: boolean;
  /** addon_exports.updated_at within the last 24h (spec section 4.1, RULING 9). */
  logged_recently: boolean;
  /** Present only at gear/gear_bags consent (spec section 3.2's fail-closed query rule). */
  item_level?: number;
  consent: GuildConsent;
}

export interface GuildHome {
  guild: GuildSummary;
  claim: ClaimStateView;
  reports: GuildHomeReport[];
  next_cursor?: string;
  roster: GuildRosterRow[];
}

export interface GuildSettingsData {
  default_visibility: GuildVisibility;
  officer_max_rank_index: number;
  claimed_by: MemberRef | null;
  claim_pending: ClaimPendingView | null;
  claim: ClaimStateView;
  invite: { rotated_at: string | null };
}

export interface ContestResult {
  status: 'contested';
}

export interface ApproveResult {
  character_key: string;
  status: 'approved';
}

export interface ClaimResult {
  status: 'confirmed' | 'pending';
  expires_at?: string;
}
export interface ClaimConfirmResult {
  status: 'confirmed';
  claimed_by: { battletag: string };
}
export interface ClaimReleaseResult {
  status: 'released';
}
export interface InviteRotateResult {
  token: string;
  url: string;
  rotated_at: string;
}
export interface InviteAcceptResult {
  guild: GuildSummary;
  rank: 'member';
}
export interface RemovedResult {
  status: 'removed';
}
export interface LeftResult {
  status: 'left';
}
export interface UpdatedMember {
  consent: GuildConsent;
}

/**
 * Every read and write in this module goes through here, mirroring rankings/api.ts's
 * `get()`: failure is read off the HTTP status (requestEnvelope's own contract), and a
 * resolved-but-null `data` becomes a GuildApiError using the envelope's own message.
 */
async function call<T>(
  path: string,
  apiBase: string,
  init: { method?: string; body?: unknown } = {},
): Promise<T> {
  let result: EnvelopeResult<T>;
  try {
    result = await requestEnvelope<T>(path, apiBase, { ...init, failureMessage: GUILD_API_FAILED });
  } catch (error) {
    if (error instanceof AccountError) throw new GuildApiError(error.message, error.status);
    throw new GuildApiError(GUILD_API_FAILED, 0);
  }
  if (result.data === null) throw new GuildApiError(result.message ?? GUILD_API_FAILED, result.status);
  return result.data;
}

export function fetchGuildHome(
  guildId: number,
  cursor?: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildHome> {
  const query = cursor === undefined ? '' : `?cursor=${encodeURIComponent(cursor)}`;
  return call<GuildHome>(`/v1/guilds/${guildId}/home${query}`, apiBase);
}

export function fetchGuildSettings(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildSettingsData> {
  return call<GuildSettingsData>(`/v1/guilds/${guildId}/settings`, apiBase);
}

export function updateGuildSettings(
  guildId: number,
  patch: { default_visibility?: GuildVisibility; officer_max_rank_index?: number },
  apiBase: string = API_BASE_URL,
): Promise<GuildSettingsData> {
  return call<GuildSettingsData>(`/v1/guilds/${guildId}/settings`, apiBase, { method: 'PATCH', body: patch });
}

export function claimGuild(guildId: number, apiBase: string = API_BASE_URL): Promise<ClaimResult> {
  return call<ClaimResult>(`/v1/guilds/${guildId}/claim`, apiBase, { method: 'POST' });
}

export function confirmClaim(guildId: number, apiBase: string = API_BASE_URL): Promise<ClaimConfirmResult> {
  return call<ClaimConfirmResult>(`/v1/guilds/${guildId}/claim/confirm`, apiBase, { method: 'POST' });
}

export function releaseClaim(guildId: number, apiBase: string = API_BASE_URL): Promise<ClaimReleaseResult> {
  return call<ClaimReleaseResult>(`/v1/guilds/${guildId}/claim/release`, apiBase, { method: 'POST' });
}

export function contestClaim(guildId: number, apiBase: string = API_BASE_URL): Promise<ContestResult> {
  return call<ContestResult>(`/v1/guilds/${guildId}/claim/contest`, apiBase, { method: 'POST' });
}

export function rotateInvite(guildId: number, apiBase: string = API_BASE_URL): Promise<InviteRotateResult> {
  return call<InviteRotateResult>(`/v1/guilds/${guildId}/invite/rotate`, apiBase, { method: 'POST' });
}

export function acceptInvite(token: string, apiBase: string = API_BASE_URL): Promise<InviteAcceptResult> {
  return call<InviteAcceptResult>(`/v1/guilds/invite/${token}/accept`, apiBase, { method: 'POST' });
}

/**
 * The real router takes three literal path segments, never a combined, encoded
 * `character_key` (plan's reconciliation record, item 4). `region`/`ruleset` are always
 * one of a small fixed vocabulary already validated by `isRegion`/`isRuleset` upstream of
 * every caller — only `slug` needs encoding, matching how `fetchCharacter` in
 * `rankings/api.ts` already builds the sibling `/v1/characters/{region}/{ruleset}/{slug}`
 * path.
 */
export function approveCharacter(
  guildId: number,
  region: string,
  ruleset: string,
  slug: string,
  apiBase: string = API_BASE_URL,
): Promise<ApproveResult> {
  return call<ApproveResult>(
    `/v1/guilds/${guildId}/characters/${region}/${ruleset}/${encodeURIComponent(slug)}/approve`,
    apiBase,
    { method: 'POST' },
  );
}

export function removeCharacter(
  guildId: number,
  region: string,
  ruleset: string,
  slug: string,
  apiBase: string = API_BASE_URL,
): Promise<RemovedResult> {
  return call<RemovedResult>(
    `/v1/guilds/${guildId}/characters/${region}/${ruleset}/${encodeURIComponent(slug)}`,
    apiBase,
    { method: 'DELETE' },
  );
}

export function updateConsent(
  guildId: number,
  consent: GuildConsent,
  apiBase: string = API_BASE_URL,
): Promise<UpdatedMember> {
  return call<UpdatedMember>(`/v1/guilds/${guildId}/members/me`, apiBase, {
    method: 'PATCH',
    body: { consent },
  });
}

export function leaveGuild(guildId: number, apiBase: string = API_BASE_URL): Promise<LeftResult> {
  return call<LeftResult>(`/v1/guilds/${guildId}/members/me`, apiBase, { method: 'DELETE' });
}
