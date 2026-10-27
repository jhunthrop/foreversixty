import { describe, expect, it } from 'vitest';
import type { GuildBot, GuildBotChannel } from './api';
import { guildBotCopy } from './discord-bot-copy';
import {
  BOT_FOCUS,
  botReturnOutcome,
  groupChannels,
  guildBotState,
  preselectChannel,
} from './discord-bot-state';

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
};
const on = (bot: GuildBot | null, connected = false) => ({ bot_available: true, bot, connected });

describe('guildBotState', () => {
  it('follows the §3.1 precedence', () => {
    expect(guildBotState(on({ ...BOT, channel_id: null }), false)).toBe('picker');
    expect(guildBotState(on(BOT), false)).toBe('installed');
    expect(guildBotState(on(BOT), true)).toBe('picker');
    expect(guildBotState(on({ ...BOT, removed_in_discord: true }), false)).toBe('removed');
    expect(guildBotState(on(null, true), false)).toBe('webhook-only');
    expect(guildBotState(on(null, false), false)).toBe('nothing');
  });

  it('never answers a bot state when the deployment has the bot off', () => {
    expect(guildBotState({ bot_available: false, bot: BOT, connected: true }, false)).toBe('webhook-only');
    expect(guildBotState({ bot_available: false, bot: BOT, connected: false }, true)).toBe('nothing');
    expect(guildBotState({ bot: BOT, connected: false }, false)).toBe('nothing');
  });
});

const channel = (id: string, name: string, category: string | null = null): GuildBotChannel => ({
  id,
  name,
  category,
});

describe('preselectChannel', () => {
  const list = [channel('1', 'general'), channel('2', 'Raid-Signups', 'Guild'), channel('3', 'raids')];
  it('prefers the current channel, then a raid channel, then none', () => {
    expect(preselectChannel(list, '3')).toBe('3');
    expect(preselectChannel(list, null)).toBe('2');
    expect(preselectChannel(list, '99')).toBe('2');
    expect(preselectChannel([channel('1', 'general')], null)).toBeNull();
  });
});

describe('groupChannels', () => {
  it('puts the uncategorised first and keeps Discord order', () => {
    const groups = groupChannels([
      channel('1', 'a', 'Raids'),
      channel('2', 'b'),
      channel('3', 'c', 'Chat'),
      channel('4', 'd', 'Raids'),
    ]);
    expect(groups.map((g) => [g.category, g.channels.map((c) => c.id)])).toEqual([
      [null, ['2']],
      ['Raids', ['1', '4']],
      ['Chat', ['3']],
    ]);
  });
});

describe('botReturnOutcome', () => {
  it('maps each return flag', () => {
    expect(botReturnOutcome('installed', null)).toEqual({
      message: guildBotCopy.returns.installed,
      role: 'status',
      focus: BOT_FOCUS.select,
      holdsPicker: true,
    });
    expect(botReturnOutcome('denied', null)?.role).toBe('status');
    expect(botReturnOutcome('failed', null)).toMatchObject({ role: 'alert', focus: BOT_FOCUS.add });
    expect(botReturnOutcome('taken', { guild_name: 'Mograine Raiders' })?.message).toContain(
      'linked to Mograine Raiders.',
    );
    expect(botReturnOutcome('taken', null)?.message).toBe(guildBotCopy.returns.takenExpired);
    expect(botReturnOutcome('nonsense', null)).toBeNull();
  });
});
