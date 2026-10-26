// web/src/lib/guild/discord-settings.ts
// The guild settings' Discord section (raids spec §3.8): what a test, a remove or a return from
// Discord says and where focus goes, and the channel as it reads after each. The words are in
// discord-copy.ts. No API call is made in this module.
import { notificationsCopy } from '../account/notifications-copy';
import type { DiscordReturnFlag } from '../account/notifications';
import { GuildApiError, type GuildSettingsDiscord } from './api';
import { guildDiscordCopy } from './discord-copy';

export { guildDiscordCopy };

/** A status line's content. `stays` holds until the next press; `planEnded` turns the section
 *  read-only and takes focus (gate spec §4). */
export interface DiscordStatus {
  ok: boolean;
  message: string;
  stays?: boolean;
  planEnded?: boolean;
}

const STATUS_RATE_LIMITED = 429;
const STATUS_DISCORD_REFUSED = 502;
const PLAN_REQUIRED = 'guild_plan_required';

const isPlanRefusal = (error: unknown): boolean =>
  error instanceof GuildApiError && error.code === PLAN_REQUIRED;

const planEnded = (): DiscordStatus => ({ ok: false, message: guildDiscordCopy.planEnded, planEnded: true });

/** A test that Discord took: switched off, it says posts are on again (account §6.1). */
export function testSent(wasSwitchedOff: boolean): DiscordStatus {
  return {
    ok: true,
    message: wasSwitchedOff ? notificationsCopy.testSentSwitchedOn : notificationsCopy.testSent,
    stays: true,
  };
}

/** A refused test: Discord's own words on a 502, the wait line on a 429, the plan line on a 403. */
export function testRefused(error: unknown): DiscordStatus {
  if (isPlanRefusal(error)) return planEnded();
  if (!(error instanceof GuildApiError)) return { ok: false, message: notificationsCopy.testOurError };
  if (error.status === STATUS_DISCORD_REFUSED)
    return { ok: false, message: notificationsCopy.testFailed(error.message) };
  if (error.status === STATUS_RATE_LIMITED) return { ok: false, message: notificationsCopy.testRateLimited };
  return { ok: false, message: notificationsCopy.testOurError };
}

export const removed = (): DiscordStatus => ({ ok: true, message: notificationsCopy.removed });
export const removeRefused = (): DiscordStatus => ({ ok: false, message: notificationsCopy.removeFailed });
export const connectStuck = (): DiscordStatus => ({ ok: true, message: guildDiscordCopy.connectStuck });

/** A test that went through resets the failures and the switched-off reason (api task 10). */
export function clearedFailures(discord: GuildSettingsDiscord): GuildSettingsDiscord {
  return { ...discord, failures: 0, disabled_reason: null, last_error: null };
}

/** The channel once removed: nothing saved, and nothing to post by default. */
export function removedChannel(discord: GuildSettingsDiscord): GuildSettingsDiscord {
  return {
    ...clearedFailures(discord),
    connected: false,
    channel_id: null,
    discord_guild_id: null,
    post_default: false,
  };
}

/** The selectors focus moves to; scoped to the section so no other block's control matches. */
export const DISCORD_FOCUS = {
  connect: '#discord [data-testid="discord-connect"], #discord [data-testid="discord-replace"]',
  test: '#discord [data-testid="discord-test"]',
  removeCancel: '#discord [data-testid="discord-remove-cancel"]',
  status: '#guild-discord-status',
} as const;

/** What a return from Discord says, where focus goes, and whether the settings read is stale. */
export interface GuildReturnOutcome {
  status: DiscordStatus;
  focus: string;
  refresh: boolean;
}

const RETURN_OUTCOMES: Readonly<Record<DiscordReturnFlag, GuildReturnOutcome>> = {
  connected: {
    status: { ok: true, message: guildDiscordCopy.returnConnected, stays: true },
    focus: DISCORD_FOCUS.test,
    refresh: true,
  },
  denied: {
    status: { ok: true, message: guildDiscordCopy.returnDenied, stays: true },
    focus: DISCORD_FOCUS.connect,
    refresh: false,
  },
  failed: {
    status: { ok: false, message: guildDiscordCopy.returnFailed },
    focus: DISCORD_FOCUS.connect,
    refresh: false,
  },
};

export const guildReturnOutcome = (flag: DiscordReturnFlag): GuildReturnOutcome => RETURN_OUTCOMES[flag];
