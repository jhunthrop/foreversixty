import { describe, expect, it } from 'vitest';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { discordBlock } from '../../test-support/notifications-api';
import type { DiscordNotifications, Me } from './api';
import {
  blockState,
  channelLabel,
  failureWarning,
  looksLikeWebhookURL,
  premiumRowLinked,
  premiumRowValue,
  webhookInputError,
} from './notifications';

function me(entitled: boolean, discord?: DiscordNotifications | null): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ notifications: entitled }),
    ...(discord === undefined ? {} : { notifications: { discord } }),
  };
}

describe('blockState', () => {
  it('is free when not entitled and nothing is saved', () => {
    expect(blockState(me(false, null)).state).toBe('free');
    expect(blockState(me(false)).state).toBe('free');
  });

  it('is empty for a payer with nothing saved, including an older /v1/me with no key', () => {
    expect(blockState(me(true, null)).state).toBe('empty');
    expect(blockState(me(true)).state).toBe('empty');
  });

  it('is saved for a working webhook and carries it', () => {
    const discord = discordBlock();
    expect(blockState(me(true, discord))).toEqual({ state: 'saved', discord });
  });

  it('is disabled for any non-null reason while entitled', () => {
    expect(blockState(me(true, discordBlock({ disabled_reason: 'failed five times' }))).state).toBe(
      'disabled',
    );
  });

  it('is lapsed when a webhook is kept without the entitlement, whatever the reason', () => {
    expect(blockState(me(false, discordBlock())).state).toBe('lapsed');
    expect(blockState(me(false, discordBlock({ disabled_reason: 'failed five times' }))).state).toBe(
      'lapsed',
    );
    expect(blockState(me(false, discordBlock({ disabled_reason: 'plan lapsed' }))).state).toBe('lapsed');
  });

  it('is free when an older /v1/me carries no entitlements at all', () => {
    const bare: Me = { ...me(true), entitlements: undefined };
    expect(blockState(bare).state).toBe('free');
  });
});

describe('premiumRowValue', () => {
  const both = { sim_finished: true, upload_parsed: true };
  it.each([
    ['free', me(false, null), 'Off'],
    ['payer, nothing saved', me(true, null), 'Off · set up'],
    ['two events', me(true, discordBlock({ events: both })), 'On · 2 events'],
    [
      'one event',
      me(true, discordBlock({ events: { sim_finished: false, upload_parsed: true } })),
      'On · 1 event',
    ],
    [
      'no event',
      me(true, discordBlock({ events: { sim_finished: false, upload_parsed: false } })),
      'Off · no events chosen',
    ],
    [
      'switched off',
      me(true, discordBlock({ disabled_reason: 'failed five times', failures: 5 })),
      'Off · failed five times',
    ],
    ['lapsed', me(false, discordBlock()), 'Paused · plan lapsed'],
  ])('reads %s', (_name, account, expected) => {
    expect(premiumRowValue(account)).toBe(expected);
  });

  it('links in every state but free', () => {
    expect(premiumRowLinked(me(false, null))).toBe(false);
    expect(premiumRowLinked(me(true, null))).toBe(true);
    expect(premiumRowLinked(me(false, discordBlock()))).toBe(true);
  });
});

describe('failureWarning', () => {
  it('is null with no failures and at five (the notice takes over)', () => {
    expect(failureWarning(discordBlock({ failures: 0 }))).toBeNull();
    expect(failureWarning(discordBlock({ failures: 5 }))).toBeNull();
  });

  it('uses the spec sentences', () => {
    expect(failureWarning(discordBlock({ failures: 3, last_error: 'Missing Permissions' }))).toBe(
      'The last 3 posts failed. At 5 in a row it switches off. Last error from Discord: Missing Permissions.',
    );
    expect(failureWarning(discordBlock({ failures: 1, last_error: 'Unknown Webhook' }))).toBe(
      'The last post failed. At 5 in a row it switches off. Last error from Discord: Unknown Webhook.',
    );
  });

  it('omits the error clause when there is none', () => {
    expect(failureWarning(discordBlock({ failures: 2, last_error: null }))).toBe(
      'The last 2 posts failed. At 5 in a row it switches off.',
    );
  });
});

describe('channelLabel', () => {
  it('names the channel by id and is null without one', () => {
    expect(channelLabel('1098765432101234567')).toBe('channel 1098765432101234567');
    expect(channelLabel(null)).toBeNull();
  });
});

describe('looksLikeWebhookURL', () => {
  const ok = 'https://discord.com/api/webhooks/123456789/abc-DEF_ghi';

  it.each([
    ok,
    `  ${ok}  `,
    'https://discordapp.com/api/webhooks/1/tok',
    'https://discord.com/api/v10/webhooks/1/tok',
    'https://discord.com/api/v6/webhooks/1/tok',
    `${ok}?wait=true#x`,
  ])('accepts %s', (input) => {
    expect(looksLikeWebhookURL(input)).toBe(true);
  });

  it.each([
    '',
    'not a url',
    'http://discord.com/api/webhooks/1/tok',
    'https://ptb.discord.com/api/webhooks/1/tok',
    'https://canary.discord.com/api/webhooks/1/tok',
    'https://discord.com.evil.test/api/webhooks/1/tok',
    'https://evil.test/api/webhooks/1/tok',
    'https://user:pw@discord.com/api/webhooks/1/tok',
    'https://discord.com:8443/api/webhooks/1/tok',
    'https://discord.com/api/v5/webhooks/1/tok',
    'https://discord.com/api/v11/webhooks/1/tok',
    'https://discord.com/api/webhooks/abc/tok',
    'https://discord.com/api/webhooks/1',
    'https://discord.com/api/webhooks/1/tok/extra',
  ])('rejects %s', (input) => {
    expect(looksLikeWebhookURL(input)).toBe(false);
  });
});

describe('webhookInputError', () => {
  it('asks for a paste when empty or blank', () => {
    expect(webhookInputError('')).toBe('Paste a Discord webhook URL.');
    expect(webhookInputError('   ')).toBe('Paste a Discord webhook URL.');
  });

  it('gives the one message for any bad value and never echoes it', () => {
    const secret = 'https://evil.test/api/webhooks/1/SECRETTOKEN';
    const message = webhookInputError(secret);
    expect(message).toBe(
      'That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.',
    );
    expect(message).not.toContain('SECRETTOKEN');
  });

  it('is null for a good value', () => {
    expect(webhookInputError('https://discord.com/api/webhooks/1/tok')).toBeNull();
  });
});
