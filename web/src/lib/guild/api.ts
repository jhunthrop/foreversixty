// web/src/lib/guild/api.ts
// Every call the browser makes to the guild half of the API (spec section 2.6), reconciled
// field-exact against the real, landed api/internal/guilds handlers (2026-09-21 plan's
// reconciliation record). Follows rankings/api.ts's pattern: the shared transport
// (account/api.ts's requestEnvelope), a module-local error class, and typed reads/writes
// with no bespoke fetch() call anywhere else in this lane's own code.
import type {
  GuildCompTarget,
  GuildNotSignedRaider,
  GuildRaidPingState,
  GuildRoleCounts,
} from './overview-api';
import type { GuildRaidAssignments } from './assignments-api';
import { AccountError, invalidateMe, requestEnvelope, type EnvelopeResult } from '../account/api';
import { invalidate, query, refetch, setQueryData } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import type {
  AttendanceSource,
  RaidAttendanceFields,
  ReadinessNextRaid,
  ReadinessRaid,
  ReadinessRaidRowFields,
} from './attendance-contract';

export type * from './attendance-contract';

export const GUILD_API_FAILED = 'That did not work; try again';

export class GuildApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
    /** The API's own `error.code` ("guild_plan_required"); empty when the envelope carried none. */
    readonly code: string = '',
    /** A 422's per-field lines, keyed by the API's field names. */
    readonly fields: Readonly<Record<string, string>> | null = null,
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

/** A claim holder or pending claimant, the way the game names them. */
export interface ClaimHolder {
  name: string;
  class: string | null;
  character_key: string;
}

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
  /** The signed-in account holds this claim. Filled by home and settings. */
  yours?: boolean;
  /** Who holds the claim, by their character on this guild's roster; never an account tag.
   *  Null when no claim is held or the holder has no character on it. */
  claimed_by?: ClaimHolder | null;
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
   * Present on every home response: true only for a verified officer or leader whose guild has the
   * officer tools (a plan or a scoped grant), false otherwise. The page treats a missing field
   * (an older API) as true: the server's 403 is the real gate.
   */
  officer_tools?: boolean;
  /** True only for a verified officer or leader of a Raid guild: the loot council and readiness board. */
  raid_guild_tools: boolean;
  /**
   * True when the guild has a verified officer or leader who can approve members. Absent on an
   * older API: unknown, read as "an approver exists" (today's note).
   */
  verify?: { has_approver: boolean };
}

/** `GET .../home`'s `summary` -- the Overview tab's roster/progression glance, computed
 *  server-side so the page never re-derives a below-rating-floor count from a partial
 *  roster page. */
export interface GuildHomeSummary {
  raiders: number;
  waiting_for_approval: number;
  below_rating_floor: number;
  /** The Progression tab's own tier counts: the band and the Overview print these. */
  tier: GuildProgressionTier;
  named_encounters_down: number;
  pulls_this_tier: number;
  updated_at: string;
}

/** The raid the needs line means (`standing.needs_raid`, guild loop round 3). */
export interface NeedsRaid {
  title: string;
  starts_at: string;
  /** The character the viewer's account signed up for that raid (In, then Late, then Bench);
   *  null with no such sign-up, when the line keeps the header character (guild round 3). */
  character_key?: string | null;
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
  /** The raid those needs are for, raid-first like Readiness and /raid (guild loop round 3): the
   *  needs line reads `before {Title}, {Thu 10 Dec}` with it. Absent on an older API. */
  needs_raid?: NeedsRaid | null;
}

/** One row of `GET /v1/guilds/search` (`data.rows`, at most 5): a stranger's "Did you mean". */
export interface GuildSearchRow {
  name: string;
  region: string;
  ruleset: string;
  /** The guild page's site path, `/guild/{region}/{ruleset}/{slug}`. */
  path: string;
}

export interface ClaimPendingView {
  by: ClaimHolder | null;
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
  /** The catalogue's full spec name ("Holy Paladin"), withheld with `spec`. */
  spec_name?: string;
  /** The raid group the row fights in (server-derived, as the raid page's sign-ups); absent
   *  with the spec, and from an API that predates it. */
  group?: GuildRaidGroup;
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
  /** The export's level, else the addon roster's; null with neither or below gear consent. */
  level?: number | null;
  /** The character the account marked as its main (`users.main_character_key`); absent on an API
   *  that does not send it yet, which reads as unmarked (guild round 3). */
  is_main?: boolean;
  consent: GuildConsent;
  attendance?: GuildRosterAttendance;
  /** Raid mode only: the same figure over the guild's last 30 logged raids. */
  attendance_30?: GuildRosterAttendance;
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
  /** The character's sim input is its FS1 export (api sims.ExportIsInputSQL; gear consent only),
   *  so the Open in simulator / planner slot has links to offer. Read at first paint to decide
   *  whether the row holds that slot at all; absent from an API that predates it. */
  sim_export?: boolean;
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
  /** The guild's next planned or locked raid, for verified members and officers only; absent
   *  for everyone else and when none is planned (design/specs/2026-10-26-guild-raids.md 3.4). */
  next_raid?: NextRaidView;
  /** `raids` once the guild has a logged raid: the roster's attendance then counts raids, not nights. */
  attendance_source?: AttendanceSource;
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
  /** The export's level, else the snapshot's; absent when the member withholds it. */
  level?: number;
  /** The snapshot's spec slug (`warrior-arms`); absent when withheld. */
  spec?: string;
  /** The sender's local `YYYY-MM-DD HH:MM`. */
  seen_at?: string;
  /** An invited member's account (`u:42`): they have no character row, so Remove names this. */
  account_key?: string;
  /** Set on an invited member's entry when the viewer's Remove would succeed (api `may_remove`). */
  may_remove?: boolean;
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
  /** The rank indexes the guild counts as raiders (the addon's Guild and Officer tabs), sorted;
   *  null until an officer says, and the addon then counts every level 60. */
  raider_rank_indexes: number[] | null;
  claimed_by: ClaimHolder | null;
  claim_pending: ClaimPendingView | null;
  claim: ClaimStateView;
  invite: GuildInviteView;
  /** The guild's Discord channel (raids spec §3.8): present only for a verified officer when
   *  the deployment has connect on. */
  discord?: GuildSettingsDiscord;
}

/** `GET /v1/guilds/{id}/settings`'s `discord` key. `disabled_reason` reads `failed five times`. */
export interface GuildSettingsDiscord {
  available: boolean;
  connected: boolean;
  channel_id: string | null;
  discord_guild_id: string | null;
  disabled_reason: string | null;
  last_error: string | null;
  failures: number;
  post_default: boolean;
  /** The deployment runs the Forever Sixty bot (discord bot spec §9); absent reads false. */
  bot_available?: boolean;
  /** The bot installed in this guild's Discord server; null or absent when it is not. */
  bot?: GuildBot | null;
  /** Why the officer's last install was refused: the guild that already holds the server. */
  bot_refusal?: { guild_name: string } | null;
}

/** The three permissions the bot needs in the raid channel, in the order the API lists them. */
export type GuildChannelPermission = 'View Channel' | 'Send Messages' | 'Embed Links';

/** Why the bot cannot post in its raid channel (spec amendment "channel can't be used"):
 *  `missing` is empty when the channel was deleted. */
export interface GuildChannelUnusable {
  reason: 'deleted' | 'permissions';
  missing: GuildChannelPermission[];
}

/** `settings.discord.bot`: `channel_id` and `channel_name` are null until the picker saves. */
export interface GuildBot {
  server_id: string;
  server_name: string;
  channel_id: string | null;
  channel_name: string | null;
  linked_members: number;
  members: number;
  failures: number;
  last_error: string | null;
  removed_in_discord: boolean;
  /** Null when the channel works; an older API's absent key is read as null. */
  channel_unusable: GuildChannelUnusable | null;
  /** Where new applications post (guild onboarding 2026-11-09); null or absent until picked. */
  recruit_channel_id?: string | null;
  recruit_channel_name?: string | null;
}

/** One entry of `GET …/discord/bot/channels`, in Discord's order. */
export interface GuildBotChannel {
  id: string;
  name: string;
  category: string | null;
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
  claimed_by: ClaimHolder | null;
}
export interface ClaimReleaseResult {
  status: 'released';
}
/** The guild's current invite link as settings reads it (api `guilds.InviteView`, spec 2026-11-05
 *  §4.1). `url` is the site path, set only for a verified officer and only while the API kept the
 *  link (it then shows again on a later visit); `made_by` is the maker's character on the roster. */
export interface GuildInviteView {
  rotated_at: string | null;
  url: string | null;
  made_by: ClaimHolder | null;
}
export interface InviteRotateResult {
  token: string;
  url: string;
  rotated_at: string;
  /** The settings shape of the link just made: its `url` is null when it cannot be shown again. */
  invite: GuildInviteView;
}
/** GET /v1/guild-invites/{token}: which guild a valid token opens, answered without joining
 *  and without a session, so the invite page can name the guild to a signed-out visitor too. */
export interface InvitePeekResult {
  guild: GuildSummary;
  /** Verified accounts in the guild. */
  members?: number;
  /** A signed-in viewer's character for this guild; null with none (and always, signed out). */
  joins_as?: InviteJoinsAs | null;
  /** The officer who made the link, by their character on the roster; null when unknown. */
  inviter?: ClaimHolder | null;
}
/** Which of the viewer's characters would join: on the guild's roster, or the one officers see
 *  until one of theirs shows on it. */
export interface InviteJoinsAs {
  name: string;
  class: string | null;
  race: string | null;
  level: number | null;
  on_roster: boolean;
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
export async function call<T>(
  path: string,
  apiBase: string,
  init: { method?: string; body?: unknown } = {},
  allowEmpty = false,
): Promise<T> {
  let result: EnvelopeResult<T>;
  try {
    result = await requestEnvelope<T>(path, apiBase, { ...init, failureMessage: GUILD_API_FAILED });
  } catch (error) {
    if (error instanceof AccountError)
      throw new GuildApiError(error.message, error.status, error.code ?? '', error.fields);
    throw new GuildApiError(GUILD_API_FAILED, 0);
  }
  if (result.data === null && !allowEmpty) {
    throw new GuildApiError(result.message ?? GUILD_API_FAILED, result.status);
  }
  return result.data as T;
}

/** A guild-scoped private answer older than this is never painted: role, plan and claim
 *  are what a player acts on, and a join, a grant or an approval can land at any moment. */
export const GUILD_TTL_MS = 60 * 1000;

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

/**
 * `bot_available: true` with no `bot` key breaks the contract (the key is an object or null).
 * It reads as the bot being unavailable, so the section draws the webhook states it can trust
 * rather than an add flow that cannot know whether a bot is installed.
 */
export function normalizeGuildDiscord(discord: GuildSettingsDiscord): GuildSettingsDiscord {
  if (discord.bot_available === true && discord.bot === undefined)
    return { ...discord, bot_available: false };
  if (discord.bot != null && discord.bot.channel_unusable === undefined)
    return { ...discord, bot: { ...discord.bot, channel_unusable: null } };
  return discord;
}

export function normalizeGuildSettings(settings: GuildSettingsData): GuildSettingsData {
  if (settings.discord === undefined) return settings;
  return { ...settings, discord: normalizeGuildDiscord(settings.discord) };
}

async function loadGuildSettings(guildId: number, apiBase: string): Promise<GuildSettingsData> {
  return normalizeGuildSettings(await call<GuildSettingsData>(`/v1/guilds/${guildId}/settings`, apiBase));
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
  onError?: () => void,
): void {
  refetch(
    guildSettingsKey(guildId, apiBase),
    () => loadGuildSettings(guildId, apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS },
    onData,
    onError,
  );
}

/**
 * `bot_refusal` is one-time (it names the refusal of the officer's last install): once the
 * section has shown it, the cached read drops it so a reload cannot draw it again.
 */
export async function clearCachedBotRefusal(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  const key = guildSettingsKey(guildId, apiBase);
  const cached = await query<GuildSettingsData>(key, () => loadGuildSettings(guildId, apiBase), {
    scope: 'private',
    ttlMs: GUILD_TTL_MS,
  });
  if (cached.discord?.bot_refusal == null) return;
  setQueryData(key, { ...cached, discord: { ...cached.discord, bot_refusal: null } });
}

export async function updateGuildSettings(
  guildId: number,
  patch: {
    default_visibility?: GuildVisibility;
    officer_max_rank_index?: number;
    /** Replaces the raider ranks; an empty list clears them to null. */
    raider_rank_indexes?: number[];
  },
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

/** An officer removes a member who joined by invite link, by their account (`account_key`). */
export async function removeMember(
  guildId: number,
  accountKey: string,
  apiBase: string = API_BASE_URL,
): Promise<RemovedResult> {
  const data = await call<RemovedResult>(
    `/v1/guilds/${guildId}/members/${encodeURIComponent(accountKey)}`,
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
  /** The night's best healer HPS and DPS-role DPS on a boss kill; absent when it has none. */
  top_hps?: GuildRaidTopParse | null;
  top_dps?: GuildRaidTopParse | null;
  fights: GuildRaidFight[];
  present: GuildRaidPresentPlayer[];
  /** The guild raid this night is matched to (a verified member's read only), else `null`. */
  raid_id?: number | null;
}

export interface GuildRaidsPage {
  rows: GuildRaidRow[];
  next_cursor: string | null;
  /** Verified members and officers only; the three keys are absent for every other viewer. */
  planned?: GuildRaidView[];
  zones?: GuildZoneView[];
  guild_timezone?: string | null;
  /** The guild channel's state (verified members only), for the create form with no raid planned. */
  discord?: GuildRaidDiscord;
}

// Planned raids and sign-ups (design/specs/2026-10-26-guild-raids.md section 9, "Contract this
// spec assumes"): the key names are the API's own.

export type GuildRaidStatus = 'planned' | 'locked' | 'done' | 'cancelled';
export type GuildSignupState = 'in' | 'late' | 'bench' | 'out';
export type GuildRaidRole = 'tank' | 'healer' | 'dps';
/** The roster group a sign-up fights in (the API's `raidspec.Group`, the Discord post's groups). */
export type GuildRaidGroup = 'tank' | 'healer' | 'melee' | 'ranged';

export interface GuildRaidCounts {
  in: number;
  late: number;
  bench: number;
  out: number;
  tanks: number;
  healers: number;
  /** Melee plus ranged; kept for the readiness page. */
  dps: number;
  /** In rows by group (absent from an API that predates them, which reads 0). */
  melee?: number;
  ranged?: number;
  /** In plus Late by class slug, one key per class the guild plays (zeros included). */
  classes?: Readonly<Record<string, number>>;
  /** Late rows by role (the role counts above are In only). */
  late_groups?: GuildRoleCounts;
}

/** A row's place in the roster and its catalogue spec (raid comp spec 4.1). Absent from an API
 *  that predates them: such a row reads its role's group, DPS as ranged, and its class alone. */
export interface GuildRaidRowComp {
  group?: GuildRaidGroup;
  /** The `specs.json` key of the row's spec, `null` when it does not resolve. */
  spec_key?: string | null;
}

export interface GuildRaidSignup extends GuildRaidRowComp {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  role: GuildRaidRole;
  state: GuildSignupState;
  note: string;
  /** The viewer's own character. */
  own: boolean;
  /** The account's other answered characters, best place first: the row is the account's best
   *  place, these ride its Also line and count nowhere (raids spec 3.7). Absent from an older API. */
  also?: GuildRaidAlsoSignup[];
}

/** Another answered character of a sign-up row's account. */
export interface GuildRaidAlsoSignup {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  spec_key: string | null;
  state: GuildSignupState;
}

export interface GuildRaidNotSigned extends GuildRaidRowComp {
  character_key: string;
  name: string;
  class: string;
  spec: string;
}

export interface GuildRaidViewerCharacter {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  verified: boolean;
  /** `null` while the character has not signed up. */
  state: GuildSignupState | null;
  role: GuildRaidRole;
  /** The character's saved raid spec in this guild (a slug from the spec catalogue, e.g.
   *  `paladin-holy`), `null` while unset. Absent from an API that predates it, which reads unset. */
  raid_spec?: string | null;
  /** The `specs.json` key of `spec`, `null` when it does not resolve. */
  spec_key?: string | null;
  /** The spec the picker starts on: `raid_spec`, else `spec_key`, else `null`. */
  default_spec?: string | null;
}

export interface GuildRaidView extends RaidAttendanceFields {
  id: number;
  zone: string;
  zone_name: string;
  /** What the zone is, which the page's words follow; absent from older fixtures, which read raid. */
  zone_kind?: GuildZoneKind;
  /** The key the run's door asks for; `null` or absent when none. */
  key_item?: GuildKeyItem | null;
  title: string;
  /** RFC 3339, UTC. */
  starts_at: string;
  duration_minutes: number;
  status: GuildRaidStatus;
  /** The curated file's raid size; `null` until sourced (ruling R6). */
  size: number | null;
  /** `null` until the guild sets a raid zone. */
  guild_timezone: string | null;
  /** Sanitised by the API: paragraphs, bold, italic, lists and links only. */
  notes_html: string;
  /** The raw markdown, for the edit form (absent from older fixtures). */
  notes?: string;
  counts: GuildRaidCounts;
  signups: GuildRaidSignup[];
  not_signed: GuildRaidNotSigned[];
  viewer: { characters: GuildRaidViewerCharacter[] };
  discord: GuildRaidDiscord;
  /** `null` while no post is scheduled for this raid. */
  discord_post: GuildRaidDiscordPost | null;
  /** The buff and debuff cover (D15), catalogue order; `null` for a dungeon or an event,
   *  absent from older fixtures. */
  cover?: GuildRaidCoverRow[] | null;
  /** Seats per role for this raid; `null` for a run with no size (dungeons, events). */
  comp_target?: GuildCompTarget | null;
  /** Officers only: raiders with no sign-up, one per account. */
  not_signed_raiders?: GuildNotSignedRaider[];
  /** Officers only: the Discord ping's throttle. */
  ping?: GuildRaidPingState;
  /** The raid leader's assignments per encounter; `null` for a dungeon or an event, absent from
   *  older fixtures. */
  assignments?: GuildRaidAssignments | null;
}

/** How a cover row's need is counted: per raid, or per five raiders in melee or at range. */
export type GuildRaidCoverPer = 'raid' | 'melee_group' | 'caster_group';

/** One buff or debuff of `GuildRaidView.cover` (Go guilds.CoverRow). The provider counts are the
 *  players signed In, Late and Bench who can provide it; `likely` is how many of the In and Late
 *  ones count on a talent their sign-up cannot show. `state` reads In plus Late against `need`. */
export interface GuildRaidCoverRow {
  key: string;
  name: string;
  kind: 'buff' | 'debuff';
  spell_id: number;
  /** The icon stem: `/data/{build}/icons/{stem}.webp`. */
  icon: string;
  need: number;
  per: GuildRaidCoverPer;
  providers_in: number;
  providers_late: number;
  providers_bench: number;
  likely: number;
  state: 'ok' | 'short' | 'missing';
}

/** The guild channel as the raid view carries it (raids spec §9): `available` is whether the
 *  deployment has Discord connect on; `post_default` the guild's remembered box. */
export interface GuildRaidDiscord {
  available: boolean;
  connected: boolean;
  post_default: boolean;
  /** The guild's remembered `Remind the day before` answer. */
  remind_default: boolean;
  /** The bot's channel name without the `#`; `null` for a webhook or when not connected. */
  channel_name: string | null;
}

/** A raid's scheduled Discord post (`GuildRaidView.discord_post`): each stamp is RFC 3339 or null. */
export interface GuildRaidDiscordPost {
  posted_at: string | null;
  post_at: string | null;
  reminded_at: string | null;
  remind: boolean;
}

/** What an officer posts (guild events spec 3.1). */
export type GuildZoneKind = 'raid' | 'dungeon' | 'event';

/** The item a run's door asks for: name and quality are the client's. */
export interface GuildKeyItem {
  id: number;
  name: string;
  quality: number;
  /** The one door the key opens when the run has a way in with no lock (`the service entrance`);
   *  absent when the run needs the key. */
  door?: string;
}

export interface GuildZoneView {
  key: string;
  name: string;
  kind: GuildZoneKind;
  size: number | null;
  /** RFC 3339; the day a raid opens, `null` for anything open now or without a date. */
  opens: string | null;
  min_level: number | null;
  max_level: number | null;
  key_item: GuildKeyItem | null;
}

/** The band line's fields: `GuildHome.next_raid`. */
export interface NextRaidView {
  id: number;
  zone: string;
  zone_name: string;
  /** Absent from older fixtures, which read raid. */
  zone_kind?: GuildZoneKind;
  title: string;
  starts_at: string;
  /** `planned` or `locked`; a locked raid's band segments are inert (absent from older fixtures). */
  status?: 'planned' | 'locked';
  size: number | null;
  guild_timezone: string | null;
  counts: Pick<GuildRaidCounts, 'in' | 'late' | 'tanks' | 'healers' | 'dps' | 'melee' | 'ranged'>;
  /** The current character's sign-up; `null` while unsigned. `note` travels with it. */
  own: { character_key: string; state: GuildSignupState; role: GuildRaidRole; note?: string } | null;
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
  /** The night's fastest kill; null on a night without one. */
  kill_ms: number | null;
  /** Every death across the night's pulls on the encounter. */
  deaths: number;
}

export interface GuildProgressionEncounter {
  encounter_id: number;
  name: string;
  zone: string;
  /** Omitted by the API with no kill; read with `typeof`, never `!== null`. */
  first_kill_at?: string | null;
  pulls: number;
  kills: number;
  /** Omitted by the API with no kill; read with `typeof`, never `!== null`. */
  best_kill_ms?: number | null;
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
  /** Every pull this tier, named and unnamed. Zero: nothing is logged in the tier yet. */
  pulls_total: number;
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
  /** What `gain_dps` counts: `dps` for a damage spec; null for a healer or tank, whose gap prints
   *  no figure. */
  unit: 'dps' | 'hps' | 'tps' | null;
}

export interface GuildReadinessEnchants {
  missing_slots: string[];
  checked: boolean;
}

export interface GuildReadinessConsumables {
  state: 'stocked' | 'short' | 'unknown';
}

export interface GuildReadinessRow extends ReadinessRaidRowFields {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  consent: GuildConsent;
  /** Export level, else the addon roster's; null with neither. Under 60 nothing is graded. */
  level: number | null;
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
  /** Officers only: the board raid's attunement done (true), not done (false), or unknown or
   *  not asked (null). Absent from an API that predates it. */
  attuned?: boolean | null;
  /** Officers only: the export's resistance totals; null without them. */
  resistances?: GuildReadinessResistances | null;
  /** Officers only: a flask or raid elixir, and buff food, in the bags; null when unknown. */
  flask_food?: { flask_or_elixir: boolean; food: boolean } | null;
}

export interface GuildReadinessResistances {
  fire: number;
  nature: number;
  frost: number;
  shadow: number;
  arcane: number;
}

export interface GuildReadinessPage {
  median_item_level: number;
  generated_at: string;
  rows: GuildReadinessRow[];
  /** The filtered read (`?raid={id}`) only. */
  raid?: ReadinessRaid;
  /** The unfiltered read only; `null` when no raid is planned or locked. */
  next_raid?: ReadinessNextRaid | null;
}

function readinessPath(guildId: number, raidId: number | undefined): string {
  const base = `/v1/guilds/${guildId}/readiness`;
  return raidId === undefined ? base : `${base}?raid=${raidId}`;
}

/** The readiness board; with `raidId`, only that raid's In and Late sign-ups (404 `not_found`
 *  for an unknown, foreign, done or cancelled raid). */
export function fetchGuildReadiness(
  guildId: number,
  raidId?: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildReadinessPage> {
  const path = readinessPath(guildId, raidId);
  return query<GuildReadinessPage>(`${apiBase}${path}`, () => call<GuildReadinessPage>(path, apiBase), {
    scope: 'private',
    ttlMs: GUILD_TTL_MS,
  });
}

export interface GuildLootEncounter {
  encounter_id: number;
  name: string;
  zone: string;
  killed: boolean;
  /** Separate nights with a kill (days in the guild's raid zone). Optional only for fixtures
   *  written before guild round 3; the API always sends it. */
  kill_nights?: number;
  /** The API's rule: kill_nights >= 3 reads "on farm"; fewer kill nights read "down". */
  on_farm?: boolean;
}

export interface GuildLootCandidate {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  /** `null` for a tier-1 (item level) candidate, and for a tier-0 one the sim never measured. A
   *  caller must branch on this before formatting a figure, never call a number method on it
   *  directly (defect 5's own live-fix round: `candidate.gain_dps.toFixed(0)` threw for every real
   *  candidate and broke the whole Loot tab). */
  gain_dps: number | null;
  /** What `gain_dps` counts; `null` when it is. */
  gain_unit: 'dps' | 'hps' | 'tps' | null;
  /** 0: the item is on the character's BiS list. 1: ranked by item level only. */
  tier: 0 | 1;
  not_sim_checked: boolean;
  /** Tier-1 only (contract's loot section): the item's own item level minus what the
   *  candidate wears in that slot (or the item's own item level when the slot is empty).
   *  `null` for a tier-0 (BiS-matched) candidate. */
  ilvl_delta: number | null;
  attendance: GuildRosterAttendance;
  already_equivalent: boolean;
  /** The drop fills an empty slot: say so, never the drop's item level as a gain (loot spec §3.2). */
  fills_empty_slot?: boolean;
  /** The item they wear now in the drop's slot; empty when the slot is empty. */
  replaces: string;
  /** The same item for the shared item link; `null` for an empty slot. */
  replaces_item: GuildLootWornItem | null;
  /** Awards this character has had in this guild this tier. */
  items_won: number;
}

export interface GuildLootWornItem {
  item_id: number;
  name: string;
  icon: string;
  quality: number;
}

/** Who an officer can hand a quest drop or the bag to, and the copy of the drop they would get. */
export interface GuildLootEligible {
  character_key: string;
  name: string;
  class: string;
  item_id: number;
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
  /** Logic only, never printed: the words come from `slotLine` over the three client fields. */
  slot: string;
  inventory_type: number;
  item_class: number;
  item_subclass: number;
  /** `gear`: a slotted drop the council ranks. `skinned`: Scale of Onyxia, never a pick. */
  kind: 'gear' | 'quest' | 'bag' | 'skinned';
  /** A quest drop's sourced line; empty otherwise. */
  quest_label: string;
  eligible: GuildLootEligible[];
  awarded_to: GuildLootAwardedTo | null;
  candidates: GuildLootCandidate[];
  /** How many of the table's drops this row stands for (two items can share a name). */
  copies: number;
  /** The boss the drop falls from, and that boss's raid. */
  encounter_id: number;
  boss: string;
  raid: string;
}

/** A first-tier raid of the drop picker; `in_client` false while the game client lacks it. */
export interface GuildLootRaid {
  key: string;
  name: string;
  in_client: boolean;
  /** Its bosses' encounter ids, in the table's order. */
  encounters: number[];
}

export interface GuildLootPage {
  encounters: GuildLootEncounter[];
  selected: number | null;
  /** Every boss's drops, boss by boss. */
  items: GuildLootItem[];
  raids: GuildLootRaid[];
  /** Verified roster characters whose level on file is 60. */
  at_level_cap: number;
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
