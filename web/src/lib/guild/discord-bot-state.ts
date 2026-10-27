// web/src/lib/guild/discord-bot-state.ts
// Which of the bot spec §3.1 states the guild's Discord section is in, what each install return
// says and where focus goes, and the channel picker's pure helpers. No API call, no Svelte.
import type { GuildBotChannel, GuildSettingsDiscord } from './api';
import { guildBotCopy } from './discord-bot-copy';

export type GuildBotState = 'picker' | 'installed' | 'removed' | 'webhook-only' | 'nothing';

type BotFacts = Pick<GuildSettingsDiscord, 'bot_available' | 'bot' | 'connected'>;

/** §3.1's precedence. With the bot off for the deployment only the webhook states can occur. */
export function guildBotState(discord: BotFacts, changingChannel: boolean): GuildBotState {
  const bot = discord.bot_available === true ? (discord.bot ?? null) : null;
  if (bot === null) return discord.connected ? 'webhook-only' : 'nothing';
  if (bot.removed_in_discord) return 'removed';
  if (bot.channel_id === null || changingChannel) return 'picker';
  return 'installed';
}

/** Section heights per state, phone then desktop: the §3.1 caps for the picker (held on the
 *  install return until the read answers), measured rest heights for the others. */
export const GUILD_BOT_RESERVE: Readonly<Record<GuildBotState, string>> = {
  picker: 'min-h-[400px] md:min-h-[330px]',
  installed: 'min-h-[330px] md:min-h-[250px]',
  removed: 'min-h-[300px] md:min-h-[230px]',
  'webhook-only': 'min-h-[350px] md:min-h-[330px]',
  nothing: 'min-h-[340px] md:min-h-[250px]',
};

/** The query parameter the install callback leaves (`?discord_bot=`), its own name so it never
 *  collides with the webhook's `discord=`. */
export const BOT_RETURN_PARAM = 'discord_bot';
export const BOT_RETURN_FLAGS = ['installed', 'denied', 'failed', 'taken'] as const;
export type BotReturnFlag = (typeof BOT_RETURN_FLAGS)[number];

export function isBotReturnFlag(value: string): value is BotReturnFlag {
  return (BOT_RETURN_FLAGS as readonly string[]).includes(value);
}

/** Reads and drops `?discord_bot=` with `history.replaceState`, keeping every other query
 *  parameter and the hash. Null when the parameter is absent. */
export function takeBotReturnFlag(location: Location, history: History): string | null {
  const params = new URLSearchParams(location.search);
  const flag = params.get(BOT_RETURN_PARAM);
  if (flag === null) return null;
  params.delete(BOT_RETURN_PARAM);
  const query = params.toString();
  history.replaceState(
    history.state,
    '',
    `${location.pathname}${query === '' ? '' : `?${query}`}${location.hash}`,
  );
  return flag;
}

/** A test that went through resets the bot's failures (api: the same as the webhook's). */
export function clearedBotFailures(discord: GuildSettingsDiscord): GuildSettingsDiscord {
  if (discord.bot === null || discord.bot === undefined) return discord;
  return { ...discord, bot: { ...discord.bot, failures: 0, last_error: null } };
}

/** The selectors focus moves to, scoped to the section. */
export const BOT_FOCUS = {
  select: '#discord [data-testid="bot-channel"]',
  add: '#discord [data-testid="bot-add"]',
  change: '#discord [data-testid="bot-change"]',
  testButton: '#discord [data-testid="bot-test"]',
  removeCancel: '#discord [data-testid="bot-remove-cancel"]',
  status: '#guild-discord-status',
} as const;

export interface BotReturnOutcome {
  message: string;
  role: 'status' | 'alert';
  focus: string;
  /** The settings read is stale: the section holds the picker height until it answers. */
  holdsPicker: boolean;
}

/** §3.1's return table. The guild named by `taken` comes from the settings read, never the URL. */
export function botReturnOutcome(
  flag: string,
  refusal: { guild_name: string } | null,
): BotReturnOutcome | null {
  if (!isBotReturnFlag(flag)) return null;
  const { returns } = guildBotCopy;
  switch (flag) {
    case 'installed':
      return { message: returns.installed, role: 'status', focus: BOT_FOCUS.select, holdsPicker: true };
    case 'denied':
      return { message: returns.denied, role: 'status', focus: BOT_FOCUS.add, holdsPicker: false };
    case 'failed':
      return { message: returns.failed, role: 'alert', focus: BOT_FOCUS.add, holdsPicker: false };
    case 'taken':
      return {
        message: refusal === null ? returns.takenExpired : returns.taken(refusal.guild_name),
        role: 'alert',
        focus: BOT_FOCUS.add,
        holdsPicker: false,
      };
  }
}

const RAID_CHANNEL = /raid/i;

/** The current channel, else the first whose name holds `raid`, else none. */
export function preselectChannel(
  channels: readonly GuildBotChannel[],
  current: string | null,
): string | null {
  if (current !== null && channels.some((channel) => channel.id === current)) return current;
  return channels.find((channel) => RAID_CHANNEL.test(channel.name))?.id ?? null;
}

export interface ChannelGroup {
  category: string | null;
  channels: GuildBotChannel[];
}

/** Channels grouped by category in Discord's order, the uncategorised first. */
export function groupChannels(channels: readonly GuildBotChannel[]): ChannelGroup[] {
  const groups = new Map<string | null, GuildBotChannel[]>();
  for (const channel of channels) {
    groups.set(channel.category, [...(groups.get(channel.category) ?? []), channel]);
  }
  const uncategorised = groups.get(null);
  const named = [...groups].filter(([category]) => category !== null);
  return [
    ...(uncategorised === undefined ? [] : [{ category: null, channels: uncategorised }]),
    ...named.map(([category, list]) => ({ category, channels: list })),
  ];
}
