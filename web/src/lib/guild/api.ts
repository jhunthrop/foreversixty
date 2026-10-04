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
  /** Control-centre contract addition (2026-10-04): who holds a `claimed` guild, for the
   *  officer tools strip's "Claimed · by {name}" line. Absent on an unclaimed/pending/
   *  contested guild, and absent until every caller of this type is updated -- never
   *  rendered as a blank "by" when missing. */
  claimed_by_name?: string;
}

/** `GET .../home`'s `viewer` (control-centre contract, 2026-10-04): the server's own answer
 *  to "who is asking," read directly instead of re-derived from `/v1/me`'s guild list. A
 *  home response from before this contract landed carries no `viewer` at all -- every
 *  reader of this falls back to the pre-existing `/v1/me`-membership derivation in that
 *  case (`GuildViewerRole | undefined`). */
export type GuildViewerRole = 'public' | 'member' | 'officer' | 'moderator';

export interface GuildViewerView {
  role: GuildViewerRole;
  character_key: string | null;
  verified: boolean;
}

/** `GET .../home`'s `summary` -- the Overview tab's roster/progression glance, computed
 *  server-side so the page never re-derives a below-rating-floor count from a partial
 *  roster page. */
export interface GuildHomeSummary {
  raiders: number;
  waiting_for_approval: number;
  below_rating_floor: number;
  named_encounters_down: number;
  pulls_this_tier: number;
  updated_at: string;
}

/** `GET .../home`'s `standing` -- member and officer only; `null` when the viewer has no
 *  ranked peers to compare against (never a fabricated rank of 1 of 1). */
export interface GuildHomeStanding {
  spec: string;
  class: string;
  same_spec_count: number;
  rank_by_item_level: number;
  item_level: number;
  /** Up to 3 of the viewer's own failing readiness checks, server-phrased short clauses
   *  (e.g. "Enchant chest") -- the Overview tab's "before Thursday" sentence joins these
   *  directly, never re-deriving its own phrasing from the Readiness endpoint (spec §8: no
   *  new round trip from Overview). */
  needs_before_next_raid: string[];
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
/** `rating_scores.overall` and its six components (control-centre contract) -- `null`
 *  (never present-but-zero) when no `rating_scores` row exists yet for this character. */
export interface GuildRosterRating {
  overall: number;
  output: number;
  survival: number;
  mechanics: number;
  utility: number;
  preparation: number;
  activity: number;
}

export interface GuildRosterAttendance {
  present: number;
  nights: number;
}

export interface GuildRosterBestParse {
  metric: string;
  value: number;
  percentile: number | null;
  encounter: string;
  report_id: string;
  fight_index: number;
}

export interface GuildRosterRow {
  character_key: string;
  region: string;
  ruleset: string;
  name: string;
  class?: string;
  spec?: string;
  /** Control-centre contract addition: the roster/readiness filter bar's "Role" column,
   *  server-derived from `spec` -- never re-derived client-side from a spec-to-role table
   *  that could drift from the server's own (the web lane's earlier plan before this
   *  contract landed; superseded now that the field is sent directly). */
  role?: 'tank' | 'healer' | 'dps';
  rank: GuildRank;
  verified: boolean;
  /** addon_exports.updated_at within the last 24h (spec section 4.1, RULING 9). */
  logged_recently: boolean;
  /** `addon_exports.updated_at` itself (control-centre contract: "replaces logged_recently;
   *  keep logged_recently too"), for the Roster tab's "Last seen" sort column. */
  logged_at?: string;
  /** Present only at gear/gear_bags consent (spec section 3.2's fail-closed query rule). */
  item_level?: number;
  consent: GuildConsent;
  attendance?: GuildRosterAttendance;
  best_parse?: GuildRosterBestParse | null;
  rating?: GuildRosterRating | null;
  professions?: string[];
  /** `guild_characters.user_id`, same value for a main and its alts (control-centre
   *  contract) -- the Roster tab's "alt of {main}" tag groups on this, never on name
   *  similarity. */
  account_key?: string;
  /** Newest report this character's account uploaded; `null` when none. */
  last_report_at?: string | null;
  may_approve?: boolean;
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
  /** Control-centre contract addition: the server's own answer to "who is asking" --
   *  absent on a pre-contract home response, in which case callers fall back to the
   *  pre-existing `/v1/me`-membership derivation. */
  viewer?: GuildViewerView;
  claim: ClaimStateView;
  summary?: GuildHomeSummary;
  standing?: GuildHomeStanding | null;
  reports: GuildHomeReport[];
  next_cursor?: string;
  roster: GuildRosterRow[];
  /** Officer: every unverified member; member/public: `[]`. */
  pending?: GuildRosterRow[];
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

// ---------------------------------------------------------------------------------------
// Control-centre contract (docs/contracts/2026-10-04-guild-centre-api.md): Raids,
// Progression, Readiness and Loot. Every read goes through the same `call`/`query` pair
// above; every tab component that calls one of these catches a 404 itself (the api lane
// builds these in a separate worktree, so a tab must render its own honest empty state --
// never a crash -- until each endpoint exists) rather than this module swallowing it.
// ---------------------------------------------------------------------------------------

export interface GuildRaidFight {
  index: number;
  name: string | null;
  encounter_id: number | null;
  kill: boolean;
  duration_ms: number;
  deaths: number;
  players: number;
}

export interface GuildRaidPresentPlayer {
  character_key: string;
  name: string;
  class: string;
}

export interface GuildRaidTopParse {
  name: string;
  class: string;
  metric: string;
  value: number;
}

export interface GuildRaidRow {
  id: string;
  title: string;
  zone: string;
  created_at: string;
  duration_ms: number;
  fight_count: number;
  kill_count: number;
  wipe_count: number;
  raiders: number;
  deaths: number;
  top_parse: GuildRaidTopParse | null;
  fights: GuildRaidFight[];
  present: GuildRaidPresentPlayer[];
}

export interface GuildRaidsPage {
  rows: GuildRaidRow[];
  next_cursor: string | null;
}

function guildRaidsKey(guildId: number, cursor: string | undefined, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/raids${cursor === undefined ? '' : `?cursor=${cursor}`}`;
}

export function fetchGuildRaids(
  guildId: number,
  cursor?: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidsPage> {
  const path = cursor === undefined ? '' : `?cursor=${encodeURIComponent(cursor)}`;
  return query<GuildRaidsPage>(
    guildRaidsKey(guildId, cursor, apiBase),
    () => call<GuildRaidsPage>(`/v1/guilds/${guildId}/raids${path}`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
  );
}

export interface GuildParseRef {
  name: string;
  class: string;
  spec?: string;
  metric: string;
  value: number;
  report_id: string;
  fight_index: number;
}

export interface GuildProgressionPullsByNight {
  report_id: string;
  date: string;
  pulls: number;
  killed: boolean;
}

export interface GuildProgressionEncounter {
  encounter_id: number;
  name: string;
  zone: string;
  first_kill_at: string | null;
  pulls: number;
  kills: number;
  best_kill_ms: number | null;
  pulls_by_night: GuildProgressionPullsByNight[];
  deaths_per_pull: number;
  best_by_role: { tank?: GuildParseRef; healer?: GuildParseRef; dps?: GuildParseRef };
}

export interface GuildProgressionUnnamed {
  zone: string;
  pulls: number;
  nights: number;
}

export interface GuildProgressionTier {
  name: string;
  raids: string[];
  named_encounters: number;
  down: number;
}

export interface GuildProgressionPage {
  tier: GuildProgressionTier;
  encounters: GuildProgressionEncounter[];
  unnamed: GuildProgressionUnnamed[];
}

function guildProgressionKey(guildId: number, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/progression`;
}

export function fetchGuildProgression(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildProgressionPage> {
  return query<GuildProgressionPage>(
    guildProgressionKey(guildId, apiBase),
    () => call<GuildProgressionPage>(`/v1/guilds/${guildId}/progression`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
  );
}

export interface GuildReadinessGearGap {
  upgrades: number | null;
  gain_dps: number | null;
  not_sim_checked: number;
}

export interface GuildReadinessEnchants {
  missing_slots: string[];
  checked: boolean;
}

export interface GuildReadinessConsumables {
  state: 'stocked' | 'short' | 'unknown';
}

export interface GuildReadinessRow {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  consent: GuildConsent;
  gear_gap: GuildReadinessGearGap | null;
  enchants: GuildReadinessEnchants;
  consumables: GuildReadinessConsumables;
  talent_points_unspent: number;
  item_level: number | null;
  item_level_delta: number | null;
  logged_at: string;
  failing: number;
  /** Officer only. */
  nudge_text?: string;
}

export interface GuildReadinessPage {
  median_item_level: number;
  generated_at: string;
  rows: GuildReadinessRow[];
}

function guildReadinessKey(guildId: number, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/readiness`;
}

export function fetchGuildReadiness(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildReadinessPage> {
  return query<GuildReadinessPage>(
    guildReadinessKey(guildId, apiBase),
    () => call<GuildReadinessPage>(`/v1/guilds/${guildId}/readiness`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
  );
}

export interface GuildLootEncounter {
  encounter_id: number;
  name: string;
  zone: string;
  killed: boolean;
}

export interface GuildLootCandidate {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  gain_dps: number;
  not_sim_checked: boolean;
  attendance: GuildRosterAttendance;
  already_equivalent: boolean;
}

export interface GuildLootAwardedTo {
  character_key: string;
  name: string;
  at: string;
  by_name: string;
}

export interface GuildLootItem {
  item_id: number;
  name: string;
  icon: string;
  quality: number;
  slot: string;
  awarded_to: GuildLootAwardedTo | null;
  candidates: GuildLootCandidate[];
}

export interface GuildLootPage {
  encounters: GuildLootEncounter[];
  selected: number | null;
  items: GuildLootItem[];
}

function guildLootKey(guildId: number, encounterId: number | undefined, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/loot${encounterId === undefined ? '' : `?encounter_id=${encounterId}`}`;
}

export function fetchGuildLoot(
  guildId: number,
  encounterId?: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildLootPage> {
  const qs = encounterId === undefined ? '' : `?encounter_id=${encounterId}`;
  return query<GuildLootPage>(
    guildLootKey(guildId, encounterId, apiBase),
    () => call<GuildLootPage>(`/v1/guilds/${guildId}/loot${qs}`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
  );
}

export interface LootAwardResult {
  id: string;
  encounter_id: number;
  item_id: number;
  character_key: string;
  awarded_at: string;
}

/** Awards and un-awards invalidate every cached loot page for this guild, not just the
 *  currently-selected encounter -- a loot decision can be read back from any encounter
 *  query once the API supports filtering by item rather than encounter alone. */
function invalidateGuildLoot(guildId: number, apiBase: string): void {
  invalidate(`${apiBase}/v1/guilds/${guildId}/loot`);
}

export async function awardLoot(
  guildId: number,
  body: { encounter_id: number; item_id: number; character_key: string },
  apiBase: string = API_BASE_URL,
): Promise<LootAwardResult> {
  const data = await call<LootAwardResult>(`/v1/guilds/${guildId}/loot/awards`, apiBase, {
    method: 'POST',
    body,
  });
  invalidateGuildLoot(guildId, apiBase);
  return data;
}

export async function unawardLoot(
  guildId: number,
  awardId: string,
  apiBase: string = API_BASE_URL,
): Promise<{ status: 'removed' }> {
  const data = await call<{ status: 'removed' }>(`/v1/guilds/${guildId}/loot/awards/${awardId}`, apiBase, {
    method: 'DELETE',
  });
  invalidateGuildLoot(guildId, apiBase);
  return data;
}

export interface ApproveAllResult {
  approved: string[];
}

export async function approveAllRoster(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<ApproveAllResult> {
  const data = await call<ApproveAllResult>(`/v1/guilds/${guildId}/roster/approve-all`, apiBase, {
    method: 'POST',
  });
  invalidate(guildHomeKey(guildId, undefined, apiBase));
  invalidate(guildSettingsKey(guildId, apiBase));
  return data;
}
