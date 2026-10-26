// web/src/lib/guild/discord-settings.test.ts
import { describe, expect, it } from 'vitest';
import { notificationsCopy } from '../account/notifications-copy';
import { GuildApiError, type GuildSettingsDiscord } from './api';
import {
  DISCORD_FOCUS,
  clearedFailures,
  guildDiscordCopy,
  guildReturnOutcome,
  removedChannel,
  testRefused,
  testSent,
} from './discord-settings';
import { GUILD_DISCORD_RESERVE, guildDiscordReserve, guildDiscordState } from './discord-state';
import { guildDiscordConnectUrl } from './discord-api';

const CONNECTED: GuildSettingsDiscord = {
  available: true,
  connected: true,
  channel_id: '1198765432101234568',
  discord_guild_id: '1087654321098765432',
  disabled_reason: null,
  last_error: null,
  failures: 0,
  post_default: true,
};
const SWITCHED_OFF = {
  ...CONNECTED,
  failures: 5,
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
};
const NOT_CONNECTED = removedChannel(CONNECTED);

describe('guildDiscordState', () => {
  it('reads not connected, connected and switched off', () => {
    expect(guildDiscordState(NOT_CONNECTED)).toBe('not-connected');
    expect(guildDiscordState(CONNECTED)).toBe('connected');
    expect(guildDiscordState(SWITCHED_OFF)).toBe('switched-off');
  });

  it('reserves the state height, and the connected height on a connected return', () => {
    expect(guildDiscordReserve(NOT_CONNECTED, '')).toBe(GUILD_DISCORD_RESERVE['not-connected']);
    expect(guildDiscordReserve(NOT_CONNECTED, '?discord=connected')).toBe(GUILD_DISCORD_RESERVE.connected);
    expect(guildDiscordReserve(NOT_CONNECTED, '?discord=denied')).toBe(
      GUILD_DISCORD_RESERVE['not-connected'],
    );
    expect(guildDiscordReserve(SWITCHED_OFF, '')).toBe(GUILD_DISCORD_RESERVE['switched-off']);
  });
});

describe('the connect route', () => {
  it('is absolute to the API origin', () => {
    expect(guildDiscordConnectUrl(501, 'https://api.example.gg')).toBe(
      'https://api.example.gg/v1/guilds/501/notifications/discord/connect',
    );
  });
});

describe('test outcomes', () => {
  it('a sent test stays, and says posts are on again when it was switched off', () => {
    expect(testSent(false)).toEqual({ ok: true, message: notificationsCopy.testSent, stays: true });
    expect(testSent(true).message).toBe(notificationsCopy.testSentSwitchedOn);
  });

  it("shows Discord's words on a 502, the wait line on a 429, our line otherwise", () => {
    expect(testRefused(new GuildApiError('Unknown Webhook', 502)).message).toBe(
      'Discord said: Unknown Webhook.',
    );
    expect(testRefused(new GuildApiError('slow down', 429)).message).toBe(notificationsCopy.testRateLimited);
    expect(testRefused(new GuildApiError('boom', 500)).message).toBe(notificationsCopy.testOurError);
    expect(testRefused(new Error('offline')).message).toBe(notificationsCopy.testOurError);
  });

  it('a guild_plan_required refusal ends the officer tools', () => {
    const refused = testRefused(new GuildApiError('Planning…', 403, 'guild_plan_required'));
    expect(refused).toEqual({ ok: false, message: guildDiscordCopy.planEnded, planEnded: true });
    expect(guildDiscordCopy.planEnded).toBe("Not saved. This guild's officer tools have ended.");
  });
});

describe('the channel after a test or a remove', () => {
  it('a test clears the failures and the switched-off reason, keeping the channel', () => {
    expect(clearedFailures(SWITCHED_OFF)).toEqual({ ...CONNECTED, failures: 0 });
  });

  it('a remove forgets the channel and the remembered box', () => {
    expect(NOT_CONNECTED).toMatchObject({
      connected: false,
      channel_id: null,
      discord_guild_id: null,
      post_default: false,
    });
    expect(CONNECTED.connected).toBe(true);
  });
});

describe('returns from Discord', () => {
  it('connected: the guild line, focus on Send test message, a fresh settings read', () => {
    expect(guildReturnOutcome('connected')).toEqual({
      status: { ok: true, message: 'Connected. Send a test message to see it in the channel.', stays: true },
      focus: DISCORD_FOCUS.test,
      refresh: true,
    });
  });

  it('denied and failed: focus back on the connect control, no read', () => {
    expect(guildReturnOutcome('denied').status.message).toBe('Cancelled in Discord. Nothing was connected.');
    expect(guildReturnOutcome('failed')).toMatchObject({
      status: { ok: false },
      focus: DISCORD_FOCUS.connect,
      refresh: false,
    });
    expect(guildReturnOutcome('failed').status.message).toContain('be an officer of this guild');
  });
});
