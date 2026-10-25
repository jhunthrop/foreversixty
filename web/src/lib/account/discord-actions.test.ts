// @vitest-environment jsdom
// web/src/lib/account/discord-actions.test.ts
// The Discord block's flows (spec 2026-10-25-discord-notifications section 4.3) against the fetch
// stub. The block's click handlers only call these and draw the result.
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createNotificationsApi, discordBlock } from '../../test-support/notifications-api';
import { removeWebhook, saveEvents, saveWebhook, sendTest } from './discord-actions';
import { notificationsCopy } from './notifications-copy';

const URL_OK = 'https://discord.com/api/webhooks/123456789/abcDEF_token-1';
const BOTH = { sim_finished: true, upload_parsed: true };
const api = createNotificationsApi();

beforeEach(() => {
  api.install();
  document.cookie = 'fs_csrf=t';
});
afterEach(() => api.reset());

describe('saveWebhook', () => {
  it('refuses an empty field without a call', async () => {
    const result = await saveWebhook('  ', BOTH);
    expect(result).toEqual({ ok: false, target: 'field', message: notificationsCopy.emptyField });
    expect(api.requests()).toHaveLength(0);
  });

  it('refuses a non-Discord URL without a call and never echoes it', async () => {
    const result = await saveWebhook('https://example.com/hook/secret', BOTH);
    expect(result).toEqual({ ok: false, target: 'field', message: notificationsCopy.invalidUrl });
    expect(JSON.stringify(result)).not.toContain('secret');
    expect(api.requests()).toHaveLength(0);
  });

  it('sends the trimmed URL and the events in one PUT', async () => {
    const result = await saveWebhook(`  ${URL_OK} `, BOTH);
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
    const result = await saveWebhook(URL_OK, BOTH);
    expect(result).toEqual({ ok: false, target, message });
  });
});

describe('saveEvents', () => {
  it('sends events only, no webhook_url key', async () => {
    api.setBlock(discordBlock({ events: { sim_finished: false, upload_parsed: true } }));
    const result = await saveEvents({ sim_finished: false, upload_parsed: true });
    expect(result).toEqual({ ok: true, message: notificationsCopy.savedEvents });
    expect(api.requests()[0]?.body).toEqual({ events: { sim_finished: false, upload_parsed: true } });
  });

  it('says it could not save on a failure', async () => {
    api.failWith('put', 500, 'x');
    expect(await saveEvents(BOTH)).toEqual({
      ok: false,
      target: 'status',
      message: notificationsCopy.saveFailed,
    });
  });
});

describe('sendTest', () => {
  it('reads Sent on a 204', async () => {
    expect(await sendTest(false)).toEqual({ ok: true, message: notificationsCopy.testSent });
  });

  it('reads Posts are on again when the webhook was switched off', async () => {
    expect(await sendTest(true)).toEqual({ ok: true, message: notificationsCopy.testSentSwitchedOn });
  });

  it('shows Discord’s words on a 502', async () => {
    api.failWith('test', 502, 'Unknown Webhook');
    expect(await sendTest(false)).toEqual({
      ok: false,
      target: 'status',
      message: notificationsCopy.testFailed('Unknown Webhook'),
    });
  });

  it('has its own sentence for the rate limit and for our errors', async () => {
    api.failWith('test', 429, 'slow');
    expect((await sendTest(false)).message).toBe(notificationsCopy.testRateLimited);
    api.failWith('test', 500, 'boom');
    expect((await sendTest(false)).message).toBe(notificationsCopy.testOurError);
  });
});

describe('removeWebhook', () => {
  it('sends one DELETE and reads Removed', async () => {
    expect(await removeWebhook()).toEqual({ ok: true, message: notificationsCopy.removed });
    expect(api.requests().map((r) => r.method)).toEqual(['DELETE']);
  });

  it('says it could not remove on a failure', async () => {
    api.failWith('delete', 500, 'x');
    expect((await removeWebhook()).message).toBe(notificationsCopy.removeFailed);
  });
});
