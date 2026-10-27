import { describe, expect, it } from 'vitest';
import type { GuildBot, GuildBotChannel, GuildChannelUnusable } from './api';
import { guildBotCopy, joinPermissions } from './discord-bot-copy';
import {
  BOT_FOCUS,
  botFailureWarning,
  botReturnOutcome,
  groupChannels,
  preselectAfterUnusable,
  preselectChannel,
} from './discord-bot-state';
import { guildBotState, guildBotView } from './discord-bot-view';

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
const on = (bot: GuildBot | null, connected = false) => ({ bot_available: true, bot, connected });

describe('guildBotView', () => {
  it('carries the bot exactly for the three bot states', () => {
    expect(guildBotView(on(BOT), false)).toEqual({ state: 'installed', bot: BOT });
    expect(guildBotView(on(null, true), false)).toEqual({ state: 'webhook-only', bot: null });
    expect(guildBotView({ bot_available: false, bot: BOT, connected: false }, false).bot).toBeNull();
  });
});

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

const DELETED: GuildChannelUnusable = { reason: 'deleted', missing: [] };
const NO_SEND: GuildChannelUnusable = { reason: 'permissions', missing: ['Send Messages'] };

describe('state 2b, the channel cannot be used', () => {
  it('sits below removed and the picker, above installed', () => {
    expect(guildBotState(on({ ...BOT, channel_unusable: DELETED }), false)).toBe('unusable');
    expect(guildBotState(on({ ...BOT, channel_unusable: DELETED, removed_in_discord: true }), false)).toBe(
      'removed',
    );
    expect(guildBotState(on({ ...BOT, channel_unusable: DELETED }), true)).toBe('picker');
    expect(guildBotState(on({ ...BOT, channel_unusable: null }), false)).toBe('installed');
  });

  it('reads the notice copy verbatim', () => {
    const copy = guildBotCopy.channelUnusable;
    expect(copy.title('deleted', 'raid-signups', 'Olympus')).toEqual({
      text: 'The raid channel #raid-signups was deleted in Olympus.',
      strong: null,
    });
    expect(copy.title('deleted', null, 'Olympus').text).toBe('The raid channel in Olympus was deleted.');
    expect(copy.title('permissions', 'raid-signups', 'Olympus')).toEqual({
      text: 'The bot cannot post in #raid-signups.',
      strong: '#raid-signups',
    });
    expect(copy.title('deleted', 'x'.repeat(70), 'Olympus').text).toContain(`#${'x'.repeat(59)}…`);
    expect(copy.detail('deleted', [], true)).toBe(
      'The bot cannot post there. Choose another channel. Raids post through your webhook channel until then.',
    );
    expect(copy.detail('deleted', [], false)).toBe(
      'The bot cannot post there. Choose another channel. Raids stop posting to Discord until then.',
    );
    expect(copy.detail('permissions', ['Send Messages', 'Embed Links'], false)).toBe(
      'In Discord, give the Forever Sixty role Send Messages and Embed Links in that channel, or choose another channel. Raids stop posting to Discord until then.',
    );
  });

  it('has no channel to open when channel_id is null (the picker draws instead)', () => {
    expect(guildBotState(on({ ...BOT, channel_id: null, channel_unusable: NO_SEND }), false)).toBe('picker');
  });

  it('joins permissions with commas and and', () => {
    expect(joinPermissions(['Send Messages'])).toBe('Send Messages');
    expect(joinPermissions(['Send Messages', 'Embed Links'])).toBe('Send Messages and Embed Links');
    expect(joinPermissions(['View Channel', 'Send Messages', 'Embed Links'])).toBe(
      'View Channel, Send Messages and Embed Links',
    );
  });

  it('preselects nothing when deleted, the current channel only when listed for permissions', () => {
    const list = [channel('1', 'general'), channel('2', 'raid-signups')];
    expect(preselectAfterUnusable(list, '2', DELETED.reason)).toBeNull();
    expect(preselectAfterUnusable(list, '2', NO_SEND.reason)).toBe('2');
    expect(preselectAfterUnusable(list, '9', NO_SEND.reason)).toBeNull();
    expect(preselectAfterUnusable(list, null, NO_SEND.reason)).toBeNull();
  });
});

describe('botFailureWarning', () => {
  const w = guildBotCopy.failureWarning;
  const failed = (failures: number, last_error: string | null) =>
    botFailureWarning({ ...BOT, failures, last_error });

  it('is absent until a post fails, and stays past five (the bot never switches itself off)', () => {
    expect(failed(0, null)).toBeNull();
    expect(failed(1, null)).toEqual({ failures: w.failures(1, 'raid-signups'), cause: null });
    expect(failed(7, 'discord status 503 code 0')?.failures).toBe(w.failures(7, 'raid-signups'));
  });

  it("names the permissions and the channel for Discord's permission refusals, never the raw error", () => {
    for (const error of [
      'discord status 403 code 50013',
      'discord status 403 code 50001',
      'discord status 404 code 50001',
    ])
      expect(failed(3, error)?.cause).toBe(w.permissions('raid-signups'));
  });

  it('says Discord did not answer for a 5xx or a transport failure, and nothing for any other refusal', () => {
    expect(failed(3, 'discord status 502 code 0')?.cause).toBe(w.noAnswer);
    expect(failed(3, 'context deadline exceeded')?.cause).toBe(w.noAnswer);
    expect(failed(3, 'discord status 400 code 50035')?.cause).toBeNull();
  });
});
