// web/src/lib/guild/api.ts
// Every call the browser makes to the guild half of the API (spec section 2.6), reconciled
// field-exact against the real, landed api/internal/guilds handlers (2026-09-21 plan's
// reconciliation record). Follows rankings/api.ts's pattern: the shared transport
// (account/api.ts's requestEnvelope), a module-local error class, and typed reads/writes
// with no bespoke fetch() call anywhere else in this lane's own code.
import { AccountError, invalidateMe, requestEnvelope, type EnvelopeResult } from '../account/api';
import { invalidate, query, refetch } from '../data/query';
import { API_BASE_URL } from '../planner/config';

export const GUILD_API_FAILED = 'That did not work; try again';

export class GuildApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
    /** The API's own `error.code` ("guild_plan_required"); empty when the envelope carried none. */
    readonly code: string = '',
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
  /** `guilds.faction` (ruling 12.2.B, design/specs/2026-10-04-guild-page.md §12.2): a
   *  stored column, the majority race among the roster's own FS1-decoded race, refreshed
   *  on roster/export change -- never computed live per request. Optional and `null` are
   *  both "no faction": a pre-contract response that carries no field at all, and a tie or
   *  empty roster that resolves to none, read identically -- neutral band, no emblem, no
   *  watermark. */
  faction?: 'alliance' | 'horde' | null;
  /** Guild crest contract (docs/contracts/2026-10-05-guild-crest-api.md): the uploaded
   *  crest's served URL, carrying the object's own hash as `?v=` so a re-upload is a new
   *  URL (never a stale cached image). Optional and `null` both mean "no crest" -- read
   *  through `guild/mark.ts`'s `guildMarkSrc`, never compared directly, so every caller
   *  falls back to the faction logo the same way. */
  crest_url?: string | null;
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
  /**
   * True when this verified officer's guild has the officer tools (a plan or a scoped grant), false
   * otherwise; absent for members. The page treats a missing field as true: the server's 403 is the real gate.
   */
  officer_tools?: boolean;
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
  /** The account shows the Supporter mark; absent from an API that predates it. */
  supporter?: boolean;
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
  /** Task L6.14's addendum (openapi `GuildRosterAddon`): this character's line in the guild's
   *  newest addon roster snapshot, `null` when the snapshot does not carry it. The API sends
   *  it on every row; optional here so a home from before the addendum still reads. */
  addon?: GuildRosterAddon | null;
}

/** One member's line in the addon's roster snapshot (addon-v2 contract 4), matched to a
 *  roster row by name. `gear_gap` is the integer the member's addon announced in its hello:
 *  the hello carries no unit, and the Roster tab reads it as DPS. */
export interface GuildRosterAddon {
  /** `level`, `spec`, `gear_gap` and `build` are null when the member has not consented to
   *  share gear; `seen_at` and `snapshot_at` are always present when the object is. */
  level: number | null;
  /** A spec slug, e.g. `warrior-arms`. */
  spec: string | null;
  /** May be negative. */
  gear_gap: number | null;
  /** The build name the member's addon announced; null when it announced none or without consent. */
  build: string | null;
  /** The sender's local time, `YYYY-MM-DD HH:MM`, as stored. */
  seen_at: string;
  /** `addon_snapshot.received_at`, RFC 3339 UTC. */
  snapshot_at: string;
}

/** The home's `addon_snapshot` (openapi `GuildAddonSnapshot`): the guild's newest addon
 *  roster snapshot. */
export interface GuildAddonSnapshot {
  /** RFC 3339 UTC. */
  received_at: string;
  /** The character KEY whose export carried the snapshot (`us/pvp/thorgrim`), not a name. */
  by_character: string;
  /** Every member line the snapshot listed, those matching no roster row included. */
  members: number;
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
  /** Task L6.14: the guild's newest addon roster snapshot, `null` when it has none. Optional
   *  so a home from before the addendum still reads (the Roster tab treats absent as null). */
  addon_snapshot?: GuildAddonSnapshot | null;
  /** The roster as a person counts it (api `HomeView.members`): everyone the newest addon
   *  snapshot names plus every verified character, merged by name, each with a state. Optional
   *  so a home from before it still reads (the Roster tab then lists the account rows alone). */
  members?: GuildRosterEntry[];
  /** The numbers the count line above the roster reads (api `HomeView.counts`). */
  counts?: GuildRosterCounts;
}

/** One person on the roster (api `RosterEntry`). `approved` is a verified character,
 *  `waiting` an account's character no officer has approved, `no_account` a name the addon saw
 *  that no account has claimed, `not_seen` a member who joined by invite and is not in the
 *  snapshot. `character_key` is set when the entry has a `roster[]` or `pending[]` row. */
export type GuildEntryState = 'approved' | 'waiting' | 'no_account' | 'not_seen';

export interface GuildRosterEntry {
  name: string;
  character_key?: string;
  state: GuildEntryState;
  rank: GuildRank;
  class?: string;
  /** The snapshot's level; absent when the member withholds it. */
  level?: number;
  /** The snapshot's spec slug (`warrior-arms`); absent when withheld. */
  spec?: string;
  /** The sender's local `YYYY-MM-DD HH:MM`. */
  seen_at?: string;
}

export interface GuildRosterCounts {
  on_roster: number;
  in_game_no_account: number;
  waiting: number;
  not_seen: number;
}

/** The words `GET /v1/me` lists on a guild entry's `can` (api `guilds.Capabilities`). */
export type GuildCapability = 'approve' | 'remove' | 'settings' | 'award' | 'billing' | 'release';

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
/** GET /v1/guild-invites/{token}: which guild a valid token opens, answered without joining
 *  and without a session, so the invite page can name the guild to a signed-out visitor too. */
export interface InvitePeekResult {
  guild: GuildSummary;
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
/**
 * `allowEmpty` is for the crest DELETE route, the one endpoint in this module whose
 * contract success response is a bare 204 (docs/contracts/2026-10-05-guild-crest-api.md)
 * rather than an envelope carrying a `data` payload -- every other call here still throws
 * on a successful-but-empty response, since that would otherwise mean the API changed
 * shape under it silently.
 */
async function call<T>(
  path: string,
  apiBase: string,
  init: { method?: string; body?: unknown } = {},
  allowEmpty = false,
): Promise<T> {
  let result: EnvelopeResult<T>;
  try {
    result = await requestEnvelope<T>(path, apiBase, { ...init, failureMessage: GUILD_API_FAILED });
  } catch (error) {
    if (error instanceof AccountError) throw new GuildApiError(error.message, error.status, error.code ?? '');
    throw new GuildApiError(GUILD_API_FAILED, 0);
  }
  if (result.data === null && !allowEmpty) {
    throw new GuildApiError(result.message ?? GUILD_API_FAILED, result.status);
  }
  return result.data as T;
}

/** A guild-scoped private answer older than this is never painted: role, plan and claim
 *  are what a player acts on, and a join, a grant or an approval can land at any moment. */
const GUILD_TTL_MS = 60 * 1000;

/** The path `fetchGuildHome` itself builds; `invalidateGuild` below drops every cursor page. */
function guildHomeKey(guildId: number, cursor: string | undefined, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/home${cursor === undefined ? '' : `?cursor=${cursor}`}`;
}

function guildSettingsKey(guildId: number, apiBase: string): string {
  return `${apiBase}/v1/guilds/${guildId}/settings`;
}

/**
 * The one invalidation every guild mutation calls (claim, release, contest, confirm, approve,
 * approve all, remove, join, leave, consent, settings, rotate, crest, award, un-award): drops
 * every cached read under this guild (home and its cursor pages, settings, raids, progression,
 * readiness, loot) and `/v1/me`, which carries the player's membership, rank, plan and the
 * guild line on their character cards. The trailing slash keeps guild 1 from catching guild 12.
 * Never called ad hoc from a component.
 */
export function invalidateGuild(guildId: number, apiBase: string = API_BASE_URL): void {
  invalidate(`${apiBase}/v1/guilds/${guildId}/`);
  invalidateMe(apiBase);
}

/**
 * "Guild home/settings for a member" (spec section 3.2): private, 5 minutes, through the
 * shared client cache (web/src/lib/data/query.ts). Every mutation below calls
 * `invalidateGuild`. Revalidated on every read: `onFresh` gets the newer answer when the
 * background read differs from the cached one this call returned.
 */
export function fetchGuildHome(
  guildId: number,
  cursor?: string,
  apiBase: string = API_BASE_URL,
  onFresh?: (home: GuildHome) => void,
): Promise<GuildHome> {
  const path = cursor === undefined ? '' : `?cursor=${encodeURIComponent(cursor)}`;
  return query<GuildHome>(
    guildHomeKey(guildId, cursor, apiBase),
    () => call<GuildHome>(`/v1/guilds/${guildId}/home${path}`, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS, revalidateOnRead: true, onRevalidated: onFresh },
  );
}

export function fetchGuildSettings(
  guildId: number,
  apiBase: string = API_BASE_URL,
  onFresh?: (settings: GuildSettingsData) => void,
): Promise<GuildSettingsData> {
  return query<GuildSettingsData>(
    guildSettingsKey(guildId, apiBase),
    () => loadGuildSettings(guildId, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS, revalidateOnRead: true, onRevalidated: onFresh },
  );
}

function loadGuildSettings(guildId: number, apiBase: string): Promise<GuildSettingsData> {
  return call<GuildSettingsData>(`/v1/guilds/${guildId}/settings`, apiBase);
}

/**
 * The settings read asked again right now, whatever the cache holds: what the Settings page
 * calls when the window regains focus, since a CLI grant or a Stripe webhook cannot notify
 * the browser. The plan lives on `/v1/me`, so callers pair it with `refetchMe`.
 */
export function refetchGuildSettings(
  guildId: number,
  onData: (settings: GuildSettingsData) => void,
  apiBase: string = API_BASE_URL,
): void {
  refetch(
    guildSettingsKey(guildId, apiBase),
    () => loadGuildSettings(guildId, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
    onData,
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
  invalidateGuild(guildId, apiBase);
  return data;
}

export async function claimGuild(guildId: number, apiBase: string = API_BASE_URL): Promise<ClaimResult> {
  const data = await call<ClaimResult>(`/v1/guilds/${guildId}/claim`, apiBase, { method: 'POST' });
  invalidateGuild(guildId, apiBase);
  return data;
}

export async function confirmClaim(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<ClaimConfirmResult> {
  const data = await call<ClaimConfirmResult>(`/v1/guilds/${guildId}/claim/confirm`, apiBase, {
    method: 'POST',
  });
  invalidateGuild(guildId, apiBase);
  return data;
}

export async function releaseClaim(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<ClaimReleaseResult> {
  const data = await call<ClaimReleaseResult>(`/v1/guilds/${guildId}/claim/release`, apiBase, {
    method: 'POST',
  });
  invalidateGuild(guildId, apiBase);
  return data;
}

export async function contestClaim(guildId: number, apiBase: string = API_BASE_URL): Promise<ContestResult> {
  const data = await call<ContestResult>(`/v1/guilds/${guildId}/claim/contest`, apiBase, { method: 'POST' });
  invalidateGuild(guildId, apiBase);
  return data;
}

export async function rotateInvite(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<InviteRotateResult> {
  const data = await call<InviteRotateResult>(`/v1/guilds/${guildId}/invite/rotate`, apiBase, {
    method: 'POST',
  });
  invalidateGuild(guildId, apiBase);
  return data;
}

/** 404 (a `GuildApiError` with status 404) for an unknown or rotated-away token. */
export function peekInvite(token: string, apiBase: string = API_BASE_URL): Promise<InvitePeekResult> {
  return call<InvitePeekResult>(`/v1/guild-invites/${encodeURIComponent(token)}`, apiBase);
}

/**
 * No `guildId` is known until the API answers (the token alone does not name a guild), so
 * the invalidation runs after it: the guild the invite named and `/v1/me` are dropped
 * before the caller navigates, so the guild page it lands on renders the member view.
 */
export async function acceptInvite(
  token: string,
  apiBase: string = API_BASE_URL,
): Promise<InviteAcceptResult> {
  const data = await call<InviteAcceptResult>(`/v1/guilds/invite/${token}/accept`, apiBase, {
    method: 'POST',
  });
  invalidateGuild(data.guild.id, apiBase);
  return data;
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
  invalidateGuild(guildId, apiBase);
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
  invalidateGuild(guildId, apiBase);
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
  invalidateGuild(guildId, apiBase);
  return data;
}

export async function leaveGuild(guildId: number, apiBase: string = API_BASE_URL): Promise<LeftResult> {
  const data = await call<LeftResult>(`/v1/guilds/${guildId}/members/me`, apiBase, { method: 'DELETE' });
  invalidateGuild(guildId, apiBase);
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
  /** True when the export carried talents, so `talent_points_unspent: 0` is a real zero; false
   *  when it carried none (the 0 is then "not read yet"). Absent on a response from before it. */
  talent_points_read?: boolean;
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
  /** `null` for a tier-1 (fallback) candidate — contract's own loot section, 2026-10-04
   *  addition: no BiS file in this build names any raid-tier item yet, so every real
   *  candidate on a live roster is tier-1 today, and this is null for every one of them.
   *  A caller must branch on this before formatting a DPS figure, never call a number
   *  method on it directly (defect 5's own live-fix round: `candidate.gain_dps.toFixed(0)`
   *  threw for every real candidate and broke the whole Loot tab). */
  gain_dps: number | null;
  not_sim_checked: boolean;
  /** Tier-1 only (contract's loot section): the item's own item level minus what the
   *  candidate wears in that slot (or the item's own item level when the slot is empty).
   *  `null` for a tier-0 (BiS-matched) candidate. */
  ilvl_delta: number | null;
  attendance: GuildRosterAttendance;
  already_equivalent: boolean;
  /** The item they wear now in the drop's slot; absent or empty when the slot is empty. */
  replaces?: string;
  /** Awards this character has had in this guild this tier; absent reads as none. */
  items_won?: number;
}

export interface GuildLootAwardedTo {
  /** The row `DELETE .../loot/awards/{award_id}` takes: what Undo sends. */
  award_id: number;
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
  /** How many of the table's drops this row stands for (two items can share a name); absent
   *  on a response from before it, read as 1. */
  copies?: number;
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
  award_id: number;
  status: 'awarded';
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
  invalidateGuild(guildId, apiBase); // every loot page, not just the selected encounter
  return data;
}

export async function unawardLoot(
  guildId: number,
  awardId: number | string,
  apiBase: string = API_BASE_URL,
): Promise<{ status: 'removed' }> {
  const data = await call<{ status: 'removed' }>(`/v1/guilds/${guildId}/loot/awards/${awardId}`, apiBase, {
    method: 'DELETE',
  });
  invalidateGuild(guildId, apiBase); // every loot page, not just the selected encounter
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
  invalidateGuild(guildId, apiBase);
  return data;
}

// ---------------------------------------------------------------------------------------
// Guild crest (docs/contracts/2026-10-05-guild-crest-api.md). Officer or leader of a
// claimed guild; a moderator may remove but the web lane only ever offers the block to
// `role === 'officer'` (Guild.svelte/GuildSettingsTab.svelte).
// ---------------------------------------------------------------------------------------

export interface GuildCrestResult {
  crest_url: string;
}

/**
 * Multipart, field `image`. The body is a `FormData`, which `account/api.ts`'s
 * `requestEnvelope` now passes straight to `fetch` without JSON-encoding or setting its
 * own `content-type` -- the browser writes the `multipart/form-data; boundary=...` header
 * itself, which is why this never sets one by hand.
 */
export async function putGuildCrest(
  guildId: number,
  file: File,
  apiBase: string = API_BASE_URL,
): Promise<GuildCrestResult> {
  const body = new FormData();
  body.append('image', file);
  const data = await call<GuildCrestResult>(`/v1/guilds/${guildId}/crest`, apiBase, {
    method: 'PUT',
    body,
  });
  invalidateGuild(guildId, apiBase);
  return data;
}

/** 204 on success (the contract's own wording) -- `call`'s `allowEmpty` reads that as
 *  success rather than "the envelope carried no data," the only route in this module that
 *  needs it. */
export async function deleteGuildCrest(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<{ status: 'removed' }> {
  await call<null>(`/v1/guilds/${guildId}/crest`, apiBase, { method: 'DELETE' }, true);
  invalidateGuild(guildId, apiBase);
  return { status: 'removed' };
}
