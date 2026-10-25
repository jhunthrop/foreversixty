// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { createNotificationsApi, discordBlock } from '../../test-support/notifications-api';
import * as query from '../data/query';
import type { Me } from './api';
import {
  NotificationsApiError,
  deleteDiscordWebhook,
  patchMeDiscord,
  saveDiscordWebhook,
  testDiscordWebhook,
} from './notifications-api';

const BASE = 'https://api.test';
const EVENTS = { sim_finished: true, upload_parsed: false };
const stub = createNotificationsApi();

beforeEach(() => stub.install());
afterEach(() => {
  stub.reset();
  vi.restoreAllMocks();
});

describe('saveDiscordWebhook', () => {
  it('PUTs the URL and events and returns the block, and leaves the cached /v1/me alone', async () => {
    stub.setBlock(discordBlock({ events: EVENTS }));
    const block = await saveDiscordWebhook(
      { webhook_url: 'https://discord.com/api/webhooks/1/tok', events: EVENTS },
      BASE,
    );
    expect(block.events).toEqual(EVENTS);
    expect(stub.requests()).toEqual([
      {
        method: 'PUT',
        path: '/v1/notifications/discord',
        body: { webhook_url: 'https://discord.com/api/webhooks/1/tok', events: EVENTS },
      },
    ]);
  });

  it('omits webhook_url for an events-only change', async () => {
    await saveDiscordWebhook({ events: EVENTS }, BASE);
    expect(stub.requests()[0]?.body).toEqual({ events: EVENTS });
  });

  it.each([400, 422])('types a %i as invalid with the API message', async (status) => {
    stub.failWith('put', status, 'bad webhook');
    await expect(saveDiscordWebhook({ webhook_url: 'x', events: EVENTS }, BASE)).rejects.toMatchObject({
      name: 'NotificationsApiError',
      kind: 'invalid',
      status,
      message: 'bad webhook',
    });
  });

  it('types 402 as not-entitled and 500 as failed', async () => {
    stub.failWith('put', 402, 'premium');
    await expect(saveDiscordWebhook({ events: EVENTS }, BASE)).rejects.toMatchObject({
      kind: 'not-entitled',
    });
    stub.failWith('put', 500, 'boom');
    await expect(saveDiscordWebhook({ events: EVENTS }, BASE)).rejects.toMatchObject({ kind: 'failed' });
  });
});

describe('testDiscordWebhook', () => {
  it('POSTs to the test route', async () => {
    await expect(testDiscordWebhook(BASE)).resolves.toBeUndefined();
    expect(stub.requests()[0]).toMatchObject({ method: 'POST', path: '/v1/notifications/discord/test' });
  });

  it('types 429 as rate-limited', async () => {
    stub.failWith('test', 429, 'slow down');
    await expect(testDiscordWebhook(BASE)).rejects.toMatchObject({ kind: 'rate-limited', status: 429 });
  });

  it('types 502 as discord and carries the Discord text, shortened to 140', async () => {
    stub.failWith('test', 502, 'Z'.repeat(300));
    const error = await testDiscordWebhook(BASE).catch((caught: unknown) => caught);
    expect(error).toBeInstanceOf(NotificationsApiError);
    expect(error).toMatchObject({ kind: 'discord', status: 502 });
    expect((error as NotificationsApiError).message).toHaveLength(140);
  });
});

describe('deleteDiscordWebhook', () => {
  it('DELETEs', async () => {
    await deleteDiscordWebhook(BASE);
    expect(stub.requests()[0]).toMatchObject({ method: 'DELETE', path: '/v1/notifications/discord' });
  });

  it('types a failure', async () => {
    stub.failWith('delete', 500, 'boom');
    await expect(deleteDiscordWebhook(BASE)).rejects.toMatchObject({ kind: 'failed' });
  });
});

describe('patchMeDiscord', () => {
  const me: Me = {
    user: { id: 1, battletag: null, email: null, role: 'member', anonymize: false },
    characters: [],
    guilds: [],
  };

  it('returns a new object with the block, leaves me untouched and writes the cache', () => {
    const write = vi.spyOn(query, 'setQueryData');
    const block = discordBlock();
    const patched = patchMeDiscord(me, block, BASE);
    expect(patched).not.toBe(me);
    expect(patched.notifications).toEqual({ discord: block });
    expect(me.notifications).toBeUndefined();
    expect(write).toHaveBeenCalledWith(`${BASE}/v1/me`, patched);
  });

  it('keeps the other channels and can clear discord', () => {
    const withOther = { ...me, notifications: { discord: discordBlock(), phone: 1 } } as unknown as Me;
    const patched = patchMeDiscord(withOther, null, BASE);
    expect(patched.notifications).toEqual({ discord: null, phone: 1 });
  });
});
