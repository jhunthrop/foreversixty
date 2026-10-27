// web/tests/e2e/support/guild-discord-bot.ts
// Stubs for the guild bot's routes (design/specs/2026-10-27-discord-bot.md §9) on top of the
// guild-discord stub: the channel list, the channel save, the bot test and the bot remove, each
// recorded. The settings read keeps answering the stub's `discord` key, which these routes change.
import type { Page, Route } from '@playwright/test';
import type { GuildBot, GuildBotChannel, GuildSettingsDiscord } from '../../../src/lib/guild/api';
import {
  E2E_API_BASE,
  NOT_CONNECTED,
  SETTINGS_GUILD_PATH,
  envelope,
  guildChannel,
  refused,
  stubGuildDiscord,
  type GuildDiscordStub,
  type GuildDiscordStubOptions,
} from './guild-discord';
import { SETTINGS_GUILD } from '../../../src/test-support/billing-fixtures';

export { E2E_API_BASE, NOT_CONNECTED, SETTINGS_GUILD_PATH, guildChannel };

export const BOT_INSTALL_HREF = `${E2E_API_BASE}/v1/guilds/${SETTINGS_GUILD.id}/discord/bot/install`;
export const SERVER_ID = '1087654321098765432';
export const RAID_CHANNEL_ID = '1198765432101234001';

export const BOT: GuildBot = {
  server_id: SERVER_ID,
  server_name: 'Olympus Discord',
  channel_id: RAID_CHANNEL_ID,
  channel_name: 'raid-signups',
  linked_members: 9,
  members: 31,
  failures: 0,
  last_error: null,
  removed_in_discord: false,
  channel_unusable: null,
};

export const CHANNELS: GuildBotChannel[] = [
  { id: '1198765432101234000', name: 'general', category: null },
  { id: RAID_CHANNEL_ID, name: 'raid-signups', category: 'Olympus Raids' },
  { id: '1198765432101234002', name: 'officers', category: 'Olympus Raids' },
];

/** The settings `discord` key with the bot on: a webhook state by default, plus overrides. */
export function botSettings(
  bot: Partial<GuildBot> | null,
  overrides: Partial<GuildSettingsDiscord> = {},
): GuildSettingsDiscord {
  return {
    ...NOT_CONNECTED,
    bot_available: true,
    bot: bot === null ? null : { ...BOT, ...bot },
    bot_refusal: null,
    ...overrides,
  };
}

export interface GuildBotStubOptions extends GuildDiscordStubOptions {
  /** The channel list; `'error'` answers 500. */
  channels?: GuildBotChannel[] | 'error';
  /** The first save answers 422 `channel_unusable` this many times. */
  unusableSaves?: number;
}

export interface GuildBotStub {
  /** The webhook stub's counters and the `discord` key. */
  base: GuildDiscordStub;
  channelReads: number;
  saves: string[];
  botTests: number;
  botDeletes: number;
}

const NO_CONTENT = { status: 204, body: '' };
const botRoute = (suffix: string): string => `**/v1/guilds/${SETTINGS_GUILD.id}/discord/bot${suffix}`;

/** Routes the settings page for The Last Watch plus the four bot routes. */
export async function stubGuildBot(page: Page, options: GuildBotStubOptions): Promise<GuildBotStub> {
  const base = await stubGuildDiscord(page, options);
  const seen: GuildBotStub = { base, channelReads: 0, saves: [], botTests: 0, botDeletes: 0 };
  let unusable = options.unusableSaves ?? 0;
  const withBot = (patch: (bot: GuildBot) => GuildBot | null): GuildSettingsDiscord | null => {
    const current = base.current();
    if (current === null || current.bot === null || current.bot === undefined) return current;
    const next = { ...current, bot: patch(current.bot) };
    base.set(next);
    return next;
  };

  await page.route(botRoute('/channels'), (route: Route) => {
    seen.channelReads += 1;
    if (options.channels === 'error')
      return route.fulfill(refused({ status: 500, code: 'internal', message: 'x' }));
    return route.fulfill(envelope({ channels: options.channels ?? CHANNELS }));
  });
  await page.route(botRoute('/channel'), (route: Route) => {
    const channelId = (route.request().postDataJSON() as { channel_id: string }).channel_id;
    seen.saves.push(channelId);
    if (unusable > 0) {
      unusable -= 1;
      return route.fulfill(refused({ status: 422, code: 'channel_unusable', message: 'unusable' }));
    }
    const name = (
      options.channels === undefined || options.channels === 'error' ? CHANNELS : options.channels
    ).find((channel) => channel.id === channelId)?.name;
    const next = withBot((bot) => ({
      ...bot,
      channel_id: channelId,
      channel_name: name ?? null,
      channel_unusable: null,
    }));
    return route.fulfill(envelope(next));
  });
  await page.route(botRoute('/test'), (route: Route) => {
    seen.botTests += 1;
    if (options.testRefusal !== undefined) return route.fulfill(refused(options.testRefusal));
    withBot((bot) => ({ ...bot, failures: 0, last_error: null }));
    return route.fulfill(NO_CONTENT);
  });
  await page.route(botRoute(''), (route: Route) => {
    if (route.request().method() !== 'DELETE') return route.fallback();
    seen.botDeletes += 1;
    withBot(() => null);
    return route.fulfill(NO_CONTENT);
  });
  return seen;
}
