// web/src/lib/guild/api.ts
// Every call the browser makes to the guild half of the API (spec section 2.6), reconciled
// field-exact against the real, landed api/internal/guilds handlers (2026-09-21 plan's
// reconciliation record). Follows rankings/api.ts's pattern: the shared transport
// (account/api.ts's requestEnvelope), a module-local error class, and typed reads/writes
// with no bespoke fetch() call anywhere else in this lane's own code.
import { AccountError, requestEnvelope, type EnvelopeResult } from '../account/api';
import { invalidate, query } from '../data/query';
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
 * `frozen` is exactly `state === 'contested'`: two earlier, narrower freeze rules (a
 * young-claim-only test, then an independence-checked corroboration test) were each found
 * gameable by a squatter, so a later security-review response simplified this to "a
 * contest always freezes officer tools" -- there is no longer a `contested`-but-not-frozen
 * case. The field stays in the response shape (the API's own choice) so consumers may keep
 * gating on `frozen`, but it is always `true` whenever `state === 'contested'` and always
 * `false` otherwise.
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

/**
 * One report row. `zone` (added by a later security-review response, item 7) lets an
 * untitled report still show something identifying, matching every other report list in
 * this codebase (`ReportView.svelte`'s own `title === '' ? zone : title` pattern) -- an
 * empty title falls back to `zone`, and only when BOTH are empty does the row fall back to
 * `guildHomeCopy.untitledReport`.
 */
export interface GuildHomeReport {
  id: string;
  title: string;
  zone: string;
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
  /**
   * Computed server-side (a later security-review response, item 7) from the exact same
   * rank-protects-rank rule the DELETE .../characters/{key} route itself enforces, from
   * the viewpoint of whoever asked for this home. Replaces this lane's own earlier,
   * conservative client-side guess (hide Remove on any officer/leader row but self or
   * moderator) with an exact one: show Remove exactly when this is `true`, nothing more.
   * A stale or forged `true` still gets a 403 with a sentence server-side -- this field
   * only controls what renders, never what the API accepts.
   */
  may_remove: boolean;
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

const GUILD_TTL_MS = 5 * 60 * 1000;

/** Matches the path `fetchGuildHome` itself builds, so `invalidate(guildHomeKey(id,
 *  undefined, apiBase))` -- a prefix match -- catches every cursor page cached for this
 *  guild, not just the cursor-less first page. */
function guildHomeKey(guildId: number, cursor: string | undefined, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/home${cursor === undefined ? '' : `?cursor=${cursor}`}`;
}

function guildSettingsKey(guildId: number, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/settings`;
}

/**
 * "Guild home/settings for a member" (spec section 3.2): private, 5 minutes, through the
 * shared client cache (web/src/lib/data/query.ts). Every mutation below invalidates both
 * keys for the guild it acted on.
 */
export function fetchGuildHome(
  guildId: number,
  cursor?: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildHome> {
  const path = cursor === undefined ? '' : `?cursor=${encodeURIComponent(cursor)}`;
  return query<GuildHome>(
    guildHomeKey(guildId, cursor, apiBase),
    () => call<GuildHome>(`/v1/guilds/${guildId}/home${path}`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
  );
}

export function fetchGuildSettings(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildSettingsData> {
  return query<GuildSettingsData>(
    guildSettingsKey(guildId, apiBase),
    () => call<GuildSettingsData>(`/v1/guilds/${guildId}/settings`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
  );
}

export async function updateGuildSettings(
  guildId: number,
  patch: { default_visibility?: GuildVisibility; officer_max_rank_index?: number },
  apiBase: string = API_BASE_URL,
): Promise<GuildSettingsData> {
  const data = await call<GuildSettingsData>(`/v1/guilds/${guildId}/settings`, apiBase, {
    method: 'PATCH',
    body: patch,
  });
  invalidate(guildSettingsKey(guildId, apiBase));
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  return data;
}

export async function claimGuild(guildId: number, apiBase: string = API_BASE_URL): Promise<ClaimResult> {
  const data = await call<ClaimResult>(`/v1/guilds/${guildId}/claim`, apiBase, { method: 'POST' });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function confirmClaim(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<ClaimConfirmResult> {
  const data = await call<ClaimConfirmResult>(`/v1/guilds/${guildId}/claim/confirm`, apiBase, {
    method: 'POST',
  });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function releaseClaim(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<ClaimReleaseResult> {
  const data = await call<ClaimReleaseResult>(`/v1/guilds/${guildId}/claim/release`, apiBase, {
    method: 'POST',
  });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function contestClaim(guildId: number, apiBase: string = API_BASE_URL): Promise<ContestResult> {
  const data = await call<ContestResult>(`/v1/guilds/${guildId}/claim/contest`, apiBase, { method: 'POST' });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function rotateInvite(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<InviteRotateResult> {
  const data = await call<InviteRotateResult>(`/v1/guilds/${guildId}/invite/rotate`, apiBase, {
    method: 'POST',
  });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

/**
 * No `guildId` is known until the API answers (the token alone does not name a guild), so
 * this deliberately does not invalidate any guild's cached home/settings -- unlike every
 * other mutation in this file. The page that calls this navigates to the guild afterward,
 * which re-fetches regardless.
 */
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
export async function approveCharacter(
  guildId: number,
  region: string,
  ruleset: string,
  slug: string,
  apiBase: string = API_BASE_URL,
): Promise<ApproveResult> {
  const data = await call<ApproveResult>(
    `/v1/guilds/${guildId}/characters/${region}/${ruleset}/${encodeURIComponent(slug)}/approve`,
    apiBase,
    { method: 'POST' },
  );
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function removeCharacter(
  guildId: number,
  region: string,
  ruleset: string,
  slug: string,
  apiBase: string = API_BASE_URL,
): Promise<RemovedResult> {
  const data = await call<RemovedResult>(
    `/v1/guilds/${guildId}/characters/${region}/${ruleset}/${encodeURIComponent(slug)}`,
    apiBase,
    { method: 'DELETE' },
  );
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function updateConsent(
  guildId: number,
  consent: GuildConsent,
  apiBase: string = API_BASE_URL,
): Promise<UpdatedMember> {
  const data = await call<UpdatedMember>(`/v1/guilds/${guildId}/members/me`, apiBase, {
    method: 'PATCH',
    body: { consent },
  });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}

export async function leaveGuild(guildId: number, apiBase: string = API_BASE_URL): Promise<LeftResult> {
  const data = await call<LeftResult>(`/v1/guilds/${guildId}/members/me`, apiBase, { method: 'DELETE' });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}
