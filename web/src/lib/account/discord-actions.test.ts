// @vitest-environment jsdom
// web/src/lib/account/discord-actions.test.ts
// The Discord block's flows (spec 2026-10-25-discord-notifications section 4.3) against the fetch
// stub. The block's click handlers only call these and draw the result.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import * as accountApi from './api';
import type { Me } from './api';
import * as query from '../data/query';
import { createNotificationsApi, discordBlock } from '../../test-support/notifications-api';
import {
  refreshAfterConnect,
  removeWebhook,
  returnOutcome,
  saveEvents,
  saveWebhook,
  sendTest,
  takeReturnFlag,
} from './discord-actions';
import { notificationsCopy } from './notifications-copy';

const URL_OK = 'https://discord.com/api/webhooks/123456789/abcDEF_token-1';
const BOTH = { sim_finished: true, upload_parsed: true, upload_private: false };
const api = createNotificationsApi();
const ME_KEY = /\/v1\/me$/;

function meWith(discord: ReturnType<typeof discordBlock> | null = null): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'member', anonymize: false },
    characters: [],
    guilds: [],
    notifications: { discord },
  };
}

/** The `/v1/me` the cache was last given. */
function cachedDiscord(write: { mock: { calls: unknown[][] } }): unknown {
  const call = write.mock.calls.find(([key]) => ME_KEY.test(String(key)));
  return (call?.[1] as Me | undefined)?.notifications?.discord;
}

beforeEach(() => {
  vi.spyOn(accountApi, 'fetchMeFresh').mockResolvedValue(null);
  api.install();
  document.cookie = 'fs_csrf=t';
});
afterEach(() => {
  api.reset();
  vi.restoreAllMocks();
});

describe('saveWebhook', () => {
  it('refuses an empty field without a call', async () => {
    const result = await saveWebhook(meWith(), '  ', BOTH);
    expect(result).toEqual({ ok: false, target: 'field', message: notificationsCopy.emptyField });
    expect(api.requests()).toHaveLength(0);
  });

  it('refuses a non-Discord URL without a call and never echoes it', async () => {
    const result = await saveWebhook(meWith(), 'https://example.com/hook/secret', BOTH);
    expect(result).toEqual({ ok: false, target: 'field', message: notificationsCopy.invalidUrl });
    expect(JSON.stringify(result)).not.toContain('secret');
    expect(api.requests()).toHaveLength(0);
  });

  it('sends the trimmed URL and the events in one PUT', async () => {
    const result = await saveWebhook(meWith(), `  ${URL_OK} `, BOTH);
    expect(result).toEqual({ ok: true, message: notificationsCopy.savedFirst });
    expect(api.requests().map((r) => [r.method, r.body])).toEqual([
      ['PUT', { webhook_url: URL_OK, events: BOTH }],
    ]);
  });

  it.each([
    [422, 'invalid', 'field', notificationsCopy.invalidUrl],
    [502, 'discord', 'field', notificationsCopy.invalidUrl],
    [402, 'plans', 'plans', notificationsCopy.notEntitled],
    [500, 'failed', 'status', notificationsCopy.saveFailed],
  ])('maps a %i to the %s wording', async (status, _name, target, message) => {
    api.failWith('put', status, `raw ${URL_OK}`);
    const result = await saveWebhook(meWith(), URL_OK, BOTH);
    expect(result).toEqual({ ok: false, target, message });
  });
});

describe('saveEvents', () => {
  it('sends events only, no webhook_url key', async () => {
    api.setBlock(
      discordBlock({ events: { sim_finished: false, upload_parsed: true, upload_private: false } }),
    );
    const result = await saveEvents(meWith(), {
      sim_finished: false,
      upload_parsed: true,
      upload_private: false,
    });
    expect(result).toEqual({ ok: true, message: notificationsCopy.savedEvents });
    expect(api.requests()[0]?.body).toEqual({
      events: { sim_finished: false, upload_parsed: true, upload_private: false },
    });
  });

  it('sends all three keys, keeping the private value when the log box goes off', async () => {
    const events = { sim_finished: true, upload_parsed: false, upload_private: true };
    await saveEvents(meWith(), events);
    expect(api.requests()[0]?.body).toEqual({ events });
  });

  it('says it could not save on a failure', async () => {
    api.failWith('put', 500, 'x');
    expect(await saveEvents(meWith(), BOTH)).toEqual({
      ok: false,
      target: 'status',
      message: notificationsCopy.saveFailed,
    });
  });
});

describe('sendTest', () => {
  it('reads Sent on a 204', async () => {
    expect(await sendTest(meWith(), false)).toEqual({ ok: true, message: notificationsCopy.testSent });
  });

  it('reads Posts are on again when the webhook was switched off', async () => {
    expect(await sendTest(meWith(), true)).toEqual({
      ok: true,
      message: notificationsCopy.testSentSwitchedOn,
    });
  });

  it('shows Discord’s words on a 502', async () => {
    api.failWith('test', 502, 'Unknown Webhook');
    expect(await sendTest(meWith(), false)).toEqual({
      ok: false,
      target: 'status',
      message: notificationsCopy.testFailed('Unknown Webhook'),
    });
  });

  it('has its own sentence for the rate limit and for our errors', async () => {
    api.failWith('test', 429, 'slow');
    expect((await sendTest(meWith(), false)).message).toBe(notificationsCopy.testRateLimited);
    api.failWith('test', 500, 'boom');
    expect((await sendTest(meWith(), false)).message).toBe(notificationsCopy.testOurError);
  });
});

describe('removeWebhook', () => {
  it('sends one DELETE and reads Removed', async () => {
    expect(await removeWebhook(meWith(discordBlock()))).toEqual({
      ok: true,
      message: notificationsCopy.removed,
    });
    expect(api.requests().map((r) => r.method)).toEqual(['DELETE']);
  });

  it('says it could not remove on a failure', async () => {
    api.failWith('delete', 500, 'x');
    expect((await removeWebhook(meWith(discordBlock()))).message).toBe(notificationsCopy.removeFailed);
  });
});

describe('writing the answer into /v1/me', () => {
  it('a save patches the cache with the returned block and refetches once, silently', async () => {
    const write = vi.spyOn(query, 'setQueryData');
    const block = discordBlock({ channel_id: '42' });
    api.setBlock(block);
    await saveWebhook(meWith(), URL_OK, BOTH);
    expect(cachedDiscord(write)).toEqual(block);
    expect(accountApi.fetchMeFresh).toHaveBeenCalledTimes(1);
  });

  it('an events change patches the cache', async () => {
    const write = vi.spyOn(query, 'setQueryData');
    const block = discordBlock({
      events: { sim_finished: false, upload_parsed: true, upload_private: false },
    });
    api.setBlock(block);
    await saveEvents(meWith(discordBlock()), block.events);
    expect(cachedDiscord(write)).toEqual(block);
  });

  it('a remove patches discord to null', async () => {
    const write = vi.spyOn(query, 'setQueryData');
    await removeWebhook(meWith(discordBlock()));
    expect(cachedDiscord(write)).toBeNull();
    expect(accountApi.fetchMeFresh).toHaveBeenCalledTimes(1);
  });

  it('a test that works on a switched-off webhook clears the failures and the reason', async () => {
    const write = vi.spyOn(query, 'setQueryData');
    const off = discordBlock({ disabled_reason: 'failed five times', failures: 5, last_error: 'x' });
    await sendTest(meWith(off), true);
    expect(cachedDiscord(write)).toEqual({ ...off, failures: 0, disabled_reason: null, last_error: null });
  });

  it('a test on a healthy webhook writes nothing', async () => {
    const write = vi.spyOn(query, 'setQueryData');
    await sendTest(meWith(discordBlock()), false);
    expect(write).not.toHaveBeenCalled();
  });

  it('a failure writes nothing and does not refetch', async () => {
    const write = vi.spyOn(query, 'setQueryData');
    api.failWith('put', 500, 'x');
    await saveEvents(meWith(discordBlock()), BOTH);
    expect(write).not.toHaveBeenCalled();
    expect(accountApi.fetchMeFresh).not.toHaveBeenCalled();
  });
});

describe('takeReturnFlag', () => {
  afterEach(() => window.history.replaceState({}, '', '/'));

  it.each([
    ['/account?discord=connected#discord', 'connected', '/account#discord'],
    ['/account?discord=denied#discord', 'denied', '/account#discord'],
    ['/account?discord=failed#discord', 'failed', '/account#discord'],
    ['/account?discord=bogus#discord', null, '/account#discord'],
    ['/account?x=1&discord=failed#discord', 'failed', '/account?x=1#discord'],
  ])('%s reads %s and leaves %s', (start, flag, left) => {
    window.history.replaceState({}, '', start);
    expect(takeReturnFlag(window.location, window.history)).toBe(flag);
    expect(window.location.pathname + window.location.search + window.location.hash).toBe(left);
  });

  it('shows the outcome once: a second read finds nothing', () => {
    window.history.replaceState({}, '', '/account?discord=connected#discord');
    expect(takeReturnFlag(window.location, window.history)).toBe('connected');
    expect(takeReturnFlag(window.location, window.history)).toBeNull();
  });

  it('does not touch the address when there is no flag', () => {
    window.history.replaceState({}, '', '/account?refreshed=1#discord');
    const replace = vi.spyOn(window.history, 'replaceState');
    expect(takeReturnFlag(window.location, window.history)).toBeNull();
    expect(replace).not.toHaveBeenCalled();
  });
});

describe('returnOutcome', () => {
  it('connected: a polite status, focus on Send test message, a fresh read', () => {
    expect(returnOutcome('connected')).toEqual({
      status: { ok: true, message: notificationsCopy.returnConnected },
      focus: '[data-testid="discord-test"]',
      refresh: true,
    });
  });

  it.each([
    ['denied', true, notificationsCopy.returnDenied],
    ['failed', false, notificationsCopy.returnFailed],
  ] as const)('%s: focus on the connect control, no refresh', (flag, ok, message) => {
    const outcome = returnOutcome(flag);
    expect(outcome.status).toEqual({ ok, message });
    expect(outcome.focus).toContain('discord-connect');
    expect(outcome.focus).toContain('discord-replace');
    expect(outcome.refresh).toBe(false);
  });
});

describe('refreshAfterConnect', () => {
  it('re-reads /v1/me once and swallows a failure', async () => {
    const fresh = vi.spyOn(accountApi, 'fetchMeFresh').mockRejectedValue(new Error('offline'));
    refreshAfterConnect();
    await Promise.resolve();
    expect(fresh).toHaveBeenCalledTimes(1);
  });
});
