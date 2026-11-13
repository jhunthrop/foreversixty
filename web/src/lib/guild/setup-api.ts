// web/src/lib/guild/setup-api.ts
// The guild setup checklist's routes (design/specs/2026-11-09-guild-onboarding.md §4): the
// checklist read, a step's mark and its undo, and the bot's recruiting channel. Same transport
// and error class as api.ts; every write drops the guild's cached reads.
import { API_BASE_URL } from '../planner/config';
import { call, invalidateGuild, normalizeGuildDiscord, type GuildSettingsDiscord } from './api';

export type SetupStepKey =
  | 'officer_ranks'
  | 'discord_bot'
  | 'recruiting'
  | 'raid_schedule'
  | 'invite_link'
  | 'addon_sync'
  | 'report_visibility'
  | 'linked_role';

export type SetupState = 'done' | 'skipped' | 'todo';
export type SetupPlan = 'none' | 'guild' | 'raid_guild';

export interface SetupStep {
  key: SetupStepKey;
  state: SetupState;
  /** The state came from an officer's mark, so it can be taken back. */
  marked: boolean;
}

export interface SetupRankLine {
  rank_index: number;
  members: number;
  /** How many of them the site knows are level 60; absent from an API older than 2026-11-13. */
  at_sixty?: number;
  names: string[];
}

/** `GET /v1/guilds/{id}/setup`. */
export interface GuildSetup {
  steps: SetupStep[];
  done: number;
  total: number;
  finished: boolean;
  plan: SetupPlan;
  ranks: { officer_max_rank_index: number; officers: number; ranks: SetupRankLine[] };
  discord: {
    bot_available: boolean;
    installed: boolean;
    server_name: string | null;
    channel_name: string | null;
    recruit_channel_name: string | null;
    webhook_connected: boolean;
    linked_members: number;
    members: number;
  };
  recruiting: { open: boolean; pitch: boolean; nights: string; wanted: string[]; role_set: boolean };
  schedule: {
    timezone: string | null;
    planned: number;
    next: { id: number; title: string; starts_at: string; posted_at: string | null } | null;
  };
  invite: { made: boolean; joined: number };
  sync: { synced: number; accounts: number; window_days: number; missing: string[] };
  visibility: string;
  /** The Raid guild readiness board's counts for the next raid; null with Raid guild or no raid. */
  readiness: { raid_title: string; signed: number; failing: number; names: string[] } | null;
}

function setupPath(guildId: number): string {
  return `/v1/guilds/${guildId}/setup`;
}

/** The checklist; never cached, since every step reads live state. */
export function fetchGuildSetup(guildId: number, apiBase: string = API_BASE_URL): Promise<GuildSetup> {
  return call<GuildSetup>(setupPath(guildId), apiBase);
}

/** Marks a step done (a confirmable step) or skipped; answers the checklist. */
export async function markSetupStep(
  guildId: number,
  step: SetupStepKey,
  state: 'done' | 'skipped',
  apiBase: string = API_BASE_URL,
): Promise<GuildSetup> {
  const setup = await call<GuildSetup>(`${setupPath(guildId)}/${step}`, apiBase, {
    method: 'PUT',
    body: { state },
  });
  invalidateGuild(guildId, apiBase);
  return setup;
}

/** Takes a step's mark back; answers the checklist. */
export async function clearSetupStep(
  guildId: number,
  step: SetupStepKey,
  apiBase: string = API_BASE_URL,
): Promise<GuildSetup> {
  const setup = await call<GuildSetup>(`${setupPath(guildId)}/${step}`, apiBase, { method: 'DELETE' });
  invalidateGuild(guildId, apiBase);
  return setup;
}

/** Saves (or with '' clears) the bot's recruiting channel; answers the settings `discord`. */
export async function saveRecruitChannel(
  guildId: number,
  channelId: string,
  apiBase: string = API_BASE_URL,
): Promise<GuildSettingsDiscord> {
  const data = await call<GuildSettingsDiscord>(
    `/v1/guilds/${guildId}/discord/bot/recruit-channel`,
    apiBase,
    { method: 'PUT', body: { channel_id: channelId } },
  );
  invalidateGuild(guildId, apiBase);
  return normalizeGuildDiscord(data);
}
