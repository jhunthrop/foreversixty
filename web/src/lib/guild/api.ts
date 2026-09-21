// web/src/lib/guild/api.ts
// Every call the browser makes to the guild half of the API (spec section 2.6). Built
// against the spec's request/response shapes exactly, except GET /v1/guilds/{id}/home,
// whose shape the spec leaves unstated — this module's GuildHome/GuildHomeReport/
// GuildRosterRow is this lane's ruling on it (plan's "Rulings" section, item 1). Follows
// rankings/api.ts's pattern: the shared transport (account/api.ts's requestEnvelope), a
// module-local error class, and typed reads/writes with no bespoke fetch() call anywhere
// else in this lane's own code.
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

/** One row of "this week's reports" — kill/wipe counts per spec section 4.1. */
export interface GuildHomeReport {
  id: string;
  title: string;
  zone: string;
  created_at: string;
  kill_count: number;
  wipe_count: number;
}

/** Keyset-paginated the same way GET /v1/reports/recent is (spec section 4.1). */
export interface GuildHomeReportsPage {
  rows: GuildHomeReport[];
  next_cursor?: string;
}

/**
 * One roster row. Synthetic `account:`-prefixed characters (invite joins with no real
 * character, spec section 2.5) are never listed here — the API omits them entirely, so
 * this type carries no synthetic-row flag to check.
 *
 * `user_id` extends this lane's own Ruling 1 shape (the real API for this endpoint hasn't
 * landed, so this is cheap to add now): a roster row otherwise carries no account/owner
 * identifier, and without one there is no way to tell "one account with two verified
 * characters" from "two accounts with one character each" -- the distinction the
 * empty-roster heuristic in Guild.svelte needs. A numeric id is consistent with this
 * codebase's other numeric id fields (see `GuildSummary.id`, `MeGuild.id`).
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
  updated_at?: string;
  /** Present only at gear/gear_bags consent (spec section 3.2's fail-closed query rule). */
  item_level?: number;
  consent: GuildConsent;
  /** The owning account's id -- this lane's own ruling, see the type doc comment above. */
  user_id: number;
}

export interface GuildHome {
  guild: GuildSummary;
  viewer: { rank: GuildRank; verified: boolean; can_manage: boolean };
  reports: GuildHomeReportsPage;
  roster: GuildRosterRow[];
}

export interface GuildSettingsData {
  default_visibility: GuildVisibility;
  officer_max_rank_index: number;
  claimed_by: { battletag: string } | null;
  claim_pending: boolean;
  invite: { rotated_at: string | null };
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

export function fetchGuildHome(guildId: number, apiBase: string = API_BASE_URL): Promise<GuildHome> {
  return call<GuildHome>(`/v1/guilds/${guildId}/home`, apiBase);
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

export function rotateInvite(guildId: number, apiBase: string = API_BASE_URL): Promise<InviteRotateResult> {
  return call<InviteRotateResult>(`/v1/guilds/${guildId}/invite/rotate`, apiBase, { method: 'POST' });
}

export function acceptInvite(token: string, apiBase: string = API_BASE_URL): Promise<InviteAcceptResult> {
  return call<InviteAcceptResult>(`/v1/guilds/invite/${token}/accept`, apiBase, { method: 'POST' });
}

/**
 * A character_key is `<region>/<ruleset>/<name-slug>` (spec section 0) — its own literal
 * slashes — so it is percent-encoded whole before being interpolated into this one path
 * segment (plan ruling 3: the API's router must decode a literal %2F here for this call to
 * reach the right route).
 */
export function approveCharacter(
  guildId: number,
  characterKey: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildRosterRow> {
  return call<GuildRosterRow>(
    `/v1/guilds/${guildId}/characters/${encodeURIComponent(characterKey)}/approve`,
    apiBase,
    { method: 'POST' },
  );
}

export function removeCharacter(
  guildId: number,
  characterKey: string,
  apiBase: string = API_BASE_URL,
): Promise<RemovedResult> {
  return call<RemovedResult>(
    `/v1/guilds/${guildId}/characters/${encodeURIComponent(characterKey)}`,
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
