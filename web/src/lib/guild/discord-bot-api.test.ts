import { describe, expect, it } from 'vitest';
import { normalizeGuildDiscord, normalizeGuildSettings, type GuildSettingsData } from './api';
import { guildBotInstallUrl } from './discord-bot-api';
import { guildBotState } from './discord-bot-view';

const WEBHOOK = {
  available: true,
  connected: false,
  channel_id: null,
  discord_guild_id: null,
  disabled_reason: null,
  last_error: null,
  failures: 0,
  post_default: false,
};

const BOT_READ = {
  server_id: '1',
  server_name: 'Olympus',
  channel_id: '10',
  channel_name: 'raid-signups',
  linked_members: 0,
  members: 0,
  failures: 0,
  last_error: null,
  removed_in_discord: false,
  channel_unusable: null,
};

describe('guildBotInstallUrl', () => {
  it('builds the install URL on the API origin', () => {
    expect(guildBotInstallUrl(7, 'https://api.example')).toBe(
      'https://api.example/v1/guilds/7/discord/bot/install',
    );
  });
});

describe('normalizeGuildDiscord', () => {
  it('reads bot_available true with no bot key as unavailable, not as state 5', () => {
    const broken = normalizeGuildDiscord({ ...WEBHOOK, bot_available: true });
    expect(broken.bot_available).toBe(false);
    expect(guildBotState(broken, false)).toBe('nothing');
  });

  it('leaves a contract-true read alone', () => {
    const read = { ...WEBHOOK, bot_available: true, bot: null };
    expect(normalizeGuildDiscord(read)).toBe(read);
    const off = { ...WEBHOOK };
    expect(normalizeGuildDiscord(off)).toBe(off);
  });

  it('reads an absent channel_unusable as null', () => {
    const bot = { ...BOT_READ };
    delete (bot as { channel_unusable?: unknown }).channel_unusable;
    const read = normalizeGuildDiscord({ ...WEBHOOK, bot_available: true, bot });
    expect(read.bot?.channel_unusable).toBeNull();
  });

  it('normalizes the discord key of a settings read, and tolerates its absence', () => {
    const base = { default_visibility: 'guild' } as GuildSettingsData;
    expect(normalizeGuildSettings(base)).toBe(base);
    const settings = normalizeGuildSettings({ ...base, discord: { ...WEBHOOK, bot_available: true } });
    expect(settings.discord?.bot_available).toBe(false);
  });
});
