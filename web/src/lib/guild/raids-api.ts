// web/src/lib/guild/raids-api.ts
// The planned-raid calls (design/specs/2026-10-26-guild-raids.md section 9): one raid read and the
// officer and member writes. Same transport, cache and error class as api.ts, so a failure is a
// `GuildApiError` carrying the API's `code` (`raid_locked`, `raid_cancelled`, `guild_plan_required`,
// `raid_cap`, `forbidden`) and every write drops the guild's cached reads through `invalidateGuild`.
import { invalidate, query } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import {
  GUILD_TTL_MS,
  call,
  invalidateGuild,
  type GuildRaidRole,
  type GuildRaidView,
  type GuildSignupState,
} from './api';

export type RaidId = number | string;
export type RaidStatusAction = 'lock' | 'unlock' | 'cancel';

/** The create and edit body (`repeat_weeks` is create only: 0, or 3 for four raids in all). */
export interface RaidInput {
  zone: string;
  title: string;
  /** RFC 3339; the form converts the guild wall clock with `wallClockToInstant`. */
  starts_at: string;
  duration_minutes: number;
  notes: string;
  /** The guild's IANA zone; the API stores it when it changes. */
  timezone: string;
  post_to_discord: boolean;
  repeat_weeks?: 0 | 3;
}

export interface SignupInput {
  state: GuildSignupState;
  role: GuildRaidRole;
  note: string;
}

export interface SignupResult {
  raid: GuildRaidView;
  /** The name of the account's other character that moved to Bench (ruling R3), else `null`. */
  moved: string | null;
}

export function raidPath(guildId: number, raidId: RaidId): string {
  return `/v1/guilds/${guildId}/raids/${encodeURIComponent(String(raidId))}`;
}

/** Under `/v1/guilds/{id}/`, so `invalidateGuild` clears it with the guild's other reads. */
function raidKey(guildId: number, raidId: RaidId, apiBase: string): string {
  return `${apiBase}${raidPath(guildId, raidId)}`;
}

/** The raid page's read; revalidated on every read, `onFresh` gets a newer answer. */
export function fetchGuildRaid(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
  onFresh?: (raid: GuildRaidView) => void,
): Promise<GuildRaidView> {
  return query<GuildRaidView>(
    raidKey(guildId, raidId, apiBase),
    () => call<GuildRaidView>(raidPath(guildId, raidId), apiBase),
    { scope: 'private', ttlMs: GUILD_TTL_MS, revalidateOnRead: true, onRevalidated: onFresh },
  );
}

/** The raid page's 10 s poll: asks again right now, whatever the cache holds. Rejects on a
 *  failure, so the page can count failed polls (spec 4 "Poll"). */
export function refetchGuildRaid(
  guildId: number,
  raidId: RaidId,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  const key = raidKey(guildId, raidId, apiBase);
  invalidate(key);
  return query<GuildRaidView>(key, () => call<GuildRaidView>(raidPath(guildId, raidId), apiBase), {
    scope: 'private',
    ttlMs: GUILD_TTL_MS,
  });
}

/** Plans one raid, or four with `repeat_weeks: 3`; answers every raid made, soonest first. */
export async function createRaid(
  guildId: number,
  body: RaidInput,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView[]> {
  const data = await call<{ raids: GuildRaidView[] }>(`/v1/guilds/${guildId}/raids`, apiBase, {
    method: 'POST',
    body,
  });
  invalidateGuild(guildId, apiBase);
  return data.raids;
}

export async function updateRaid(
  guildId: number,
  raidId: RaidId,
  body: Omit<RaidInput, 'repeat_weeks'>,
  apiBase: string = API_BASE_URL,
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(raidPath(guildId, raidId), apiBase, { method: 'PATCH', body });
  invalidateGuild(guildId, apiBase);
  return raid;
}

/** A cancel may ask for the Discord post (spec 4 "Cancel confirm"); lane G1b reads it. */
export interface RaidStatusBody {
  post_to_discord: boolean;
}

export async function setRaidStatus(
  guildId: number,
  raidId: RaidId,
  action: RaidStatusAction,
  apiBase: string = API_BASE_URL,
  body?: RaidStatusBody,
): Promise<GuildRaidView> {
  const raid = await call<GuildRaidView>(`${raidPath(guildId, raidId)}/${action}`, apiBase, {
    method: 'POST',
    ...(body === undefined ? {} : { body }),
  });
  invalidateGuild(guildId, apiBase);
  return raid;
}

/** A character key is `region/ruleset/name`; the route takes its three segments, so each is
 *  encoded alone and the slashes stay (encoding the whole key would miss the route). */
export function characterKeyPath(characterKey: string): string {
  return characterKey.split('/').map(encodeURIComponent).join('/');
}

export async function putSignup(
  guildId: number,
  raidId: RaidId,
  characterKey: string,
  body: SignupInput,
  apiBase: string = API_BASE_URL,
): Promise<SignupResult> {
  const result = await call<SignupResult>(
    `${raidPath(guildId, raidId)}/signups/${characterKeyPath(characterKey)}`,
    apiBase,
    { method: 'PUT', body },
  );
  invalidateGuild(guildId, apiBase);
  return result;
}
