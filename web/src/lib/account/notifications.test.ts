import { describe, expect, it } from 'vitest';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { discordBlock } from '../../test-support/notifications-api';
import type { DiscordNotifications, Me } from './api';
import {
  blockState,
  connectAvailable,
  discordReturnFlag,
  openInDiscordUrl,
  removeQuestionFor,
  sourceOf,
  failureWarning,
  looksLikeWebhookURL,
  premiumRowLinked,
  premiumRowValue,
  premiumRowVisible,
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
  });

  it('is off when /v1/me has no notifications key: the deployment has the feature off', () => {
    expect(blockState(me(false)).state).toBe('off');
    expect(blockState(me(true)).state).toBe('off');
  });

  it('is empty for a payer whose /v1/me carries the key with no webhook', () => {
    expect(blockState(me(true, null)).state).toBe('empty');
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
    const bare: Me = { ...me(true, null), entitlements: undefined };
    expect(blockState(bare).state).toBe('free');
  });
});

describe('premiumRowValue', () => {
  const both = { sim_finished: true, upload_parsed: true, upload_private: false };
  const none = { sim_finished: false, upload_parsed: false, upload_private: false };
  it.each([
    ['free', me(false, null), 'Off\u00A0·\u00A0Premium'],
    ['payer, nothing saved', me(true, null), 'Off\u00A0·\u00A0set\u00A0up'],
    ['two events', me(true, discordBlock({ events: both })), 'On\u00A0·\u00A02\u00A0events'],
    [
      'one event',
      me(true, discordBlock({ events: { ...none, upload_parsed: true } })),
      'On\u00A0·\u00A01\u00A0event',
    ],
    ['no event', me(true, discordBlock({ events: none })), 'Off\u00A0· no events chosen'],
    [
      'switched off',
      me(true, discordBlock({ disabled_reason: 'failed five times', failures: 5 })),
      'Off\u00A0· needs attention',
    ],
    ['lapsed', me(false, discordBlock()), 'Paused\u00A0· plan lapsed'],
    ['feature off, entitled', me(true), 'Off\u00A0·\u00A0Premium'],
    ['feature off, lapsed', me(false), 'Off\u00A0·\u00A0Premium'],
  ])('reads %s', (_name, account, expected) => {
    expect(premiumRowValue(account)).toBe(expected);
  });

  it('counts sim_finished plus upload_parsed only: upload_private is not an event', () => {
    const allThree = { ...both, upload_private: true };
    expect(premiumRowValue(me(true, discordBlock({ events: allThree })))).toBe(
      'On\u00A0·\u00A02\u00A0events',
    );
    expect(premiumRowValue(me(true, discordBlock({ events: { ...none, upload_private: true } })))).toBe(
      'Off\u00A0· no events chosen',
    );
  });

  it('shows no row when the deployment has notifications off, and a row in every other state', () => {
    expect(premiumRowVisible(me(true))).toBe(false);
    expect(premiumRowVisible(me(false))).toBe(false);
    expect(premiumRowVisible(me(false, null))).toBe(true);
    expect(premiumRowVisible(me(true, null))).toBe(true);
    expect(premiumRowVisible(me(true, discordBlock()))).toBe(true);
    expect(premiumRowVisible(me(false, discordBlock()))).toBe(true);
  });

  it('links in every state but free', () => {
    expect(premiumRowLinked(me(false, null))).toBe(false);
    expect(premiumRowLinked(me(true, null))).toBe(true);
    expect(premiumRowLinked(me(false, discordBlock()))).toBe(true);
  });

  it('does not link when the deployment has notifications off, whatever the entitlement', () => {
    expect(premiumRowLinked(me(true))).toBe(false);
    expect(premiumRowLinked(me(false))).toBe(false);
  });
});

describe('failureWarning', () => {
  it('is null with no failures and at five (the notice takes over)', () => {
    expect(failureWarning(discordBlock({ failures: 0 }))).toBeNull();
    expect(failureWarning(discordBlock({ failures: 5 }))).toBeNull();
  });

  it('uses the spec sentences', () => {
    expect(failureWarning(discordBlock({ failures: 3, last_error: 'Missing Permissions' }))).toEqual({
      failures: 'The last 3 posts failed. At 5 in a row it switches off. Send a test message to check.',
      lastError: 'Last error from Discord: Missing Permissions.',
    });
    expect(failureWarning(discordBlock({ failures: 1, last_error: 'Unknown Webhook' }))).toEqual({
      failures: 'The last post failed. At 5 in a row it switches off. Send a test message to check.',
      lastError: 'Last error from Discord: Unknown Webhook.',
    });
  });

  it('omits the error clause when there is none', () => {
    expect(failureWarning(discordBlock({ failures: 2, last_error: null }))).toEqual({
      failures: 'The last 2 posts failed. At 5 in a row it switches off. Send a test message to check.',
      lastError: null,
    });
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
    `${ok}?wait=true#x`,
    `${ok}?wait=true`,
    `${ok}#x`,
    `${ok}?`,
    `${ok}#`,
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

describe('connectAvailable', () => {
  it('reads the flag; a missing key or no notifications reads false', () => {
    const withFlag = (connect: boolean | undefined): Me => ({
      ...me(true, null),
      notifications: { ...(connect === undefined ? {} : { connect }), discord: null },
    });
    expect(connectAvailable(withFlag(true))).toBe(true);
    expect(connectAvailable(withFlag(false))).toBe(false);
    expect(connectAvailable(withFlag(undefined))).toBe(false);
    expect(connectAvailable(me(true))).toBe(false);
  });
});

describe('discordReturnFlag', () => {
  it.each([
    ['?discord=connected', 'connected'],
    ['?discord=denied', 'denied'],
    ['?discord=failed', 'failed'],
    ['?a=1&discord=failed', 'failed'],
    ['?discord=other', null],
    ['?discord=', null],
    ['?discord=Connected', null],
    ['?refreshed=1', null],
    ['', null],
  ])('%s -> %s', (search, expected) => {
    expect(discordReturnFlag(search)).toBe(expected);
  });
});

describe('connected destination helpers', () => {
  it('treats a missing source as pasted', () => {
    const { source: _source, ...older } = discordBlock();
    expect(sourceOf(older)).toBe('pasted');
    expect(sourceOf(discordBlock({ source: 'connected' }))).toBe('connected');
  });

  it('asks the remove question that matches the source', () => {
    expect(removeQuestionFor(discordBlock())).toContain('It stays in Discord');
    expect(removeQuestionFor(discordBlock({ source: 'connected' }))).toContain(
      'delete its webhook from Discord',
    );
  });

  it('builds the Open in Discord link from numeric ids only', () => {
    const known = discordBlock({ channel_id: '222', guild_id: '111' });
    expect(openInDiscordUrl(known)).toBe('https://discord.com/channels/111/222');
    expect(openInDiscordUrl(discordBlock({ guild_id: null }))).toBeNull();
    expect(openInDiscordUrl(discordBlock({ guild_id: '1/../x' }))).toBeNull();
  });
});
