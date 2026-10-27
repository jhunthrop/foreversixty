import { describe, expect, it } from 'vitest';
import type { GuildBot, GuildSettingsDiscord } from './api';
import { GUILD_BOT_RESERVE, WEBHOOK_OFFER_RESERVE } from './discord-bot-view';
import { GUILD_DISCORD_RESERVE, guildDiscordReserve } from './discord-state';

const BOT: GuildBot = {
  server_id: '1',
  server_name: 'Olympus',
  channel_id: '10',
  channel_name: 'raid-signups',
  linked_members: 9,
  members: 31,
  failures: 0,
  last_error: null,
  removed_in_discord: false,
  channel_unusable: null,
};
type Webhook = Pick<GuildSettingsDiscord, 'connected' | 'disabled_reason'>;
const NONE: Webhook = { connected: false, disabled_reason: null };
const CONNECTED: Webhook = { connected: true, disabled_reason: null };
const SWITCHED_OFF: Webhook = { connected: true, disabled_reason: 'failed five times' };
const on = (bot: GuildBot | null, webhook: Webhook = NONE) => ({ ...webhook, bot_available: true, bot });

describe('guildDiscordReserve with the bot off', () => {
  it("is G1's, held at the connected height after a webhook return", () => {
    expect(guildDiscordReserve({ ...NONE, bot_available: false, bot: null }, '')).toBe(
      GUILD_DISCORD_RESERVE['not-connected'],
    );
    expect(guildDiscordReserve({ ...CONNECTED }, '')).toBe(GUILD_DISCORD_RESERVE.connected);
    expect(guildDiscordReserve({ ...NONE }, '?discord=connected')).toBe(GUILD_DISCORD_RESERVE.connected);
  });
});

describe('guildDiscordReserve with the bot on', () => {
  it("keeps G1's two webhook heights and adds only the offer", () => {
    expect(guildDiscordReserve(on(null, CONNECTED), '')).toBe(WEBHOOK_OFFER_RESERVE.connected);
    expect(guildDiscordReserve(on(null, SWITCHED_OFF), '')).toBe(WEBHOOK_OFFER_RESERVE['switched-off']);
  });

  it("reserves the bot's own states", () => {
    expect(guildDiscordReserve(on(null), '')).toBe(GUILD_BOT_RESERVE.nothing);
    expect(guildDiscordReserve(on(BOT), '')).toBe(GUILD_BOT_RESERVE.installed);
    expect(guildDiscordReserve(on({ ...BOT, channel_id: null }), '')).toBe(GUILD_BOT_RESERVE.picker);
    expect(guildDiscordReserve(on({ ...BOT, removed_in_discord: true }), '')).toBe(GUILD_BOT_RESERVE.removed);
  });

  it('holds the picker on an installed return, whatever the stale read says', () => {
    expect(guildDiscordReserve(on(null), '?discord_bot=installed')).toBe(GUILD_BOT_RESERVE.picker);
    expect(guildDiscordReserve(on(BOT), '?discord_bot=installed')).toBe(GUILD_BOT_RESERVE.picker);
  });

  it('holds the connected offer height on a webhook connected return', () => {
    expect(guildDiscordReserve(on(null), '?discord=connected')).toBe(WEBHOOK_OFFER_RESERVE.connected);
  });
});
