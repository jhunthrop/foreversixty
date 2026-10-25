// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { createNotificationsApi, discordBlock } from '../../test-support/notifications-api';
import * as accountApi from './api';
import {
  NotificationsApiError,
  deleteDiscordWebhook,
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
  it('PUTs the URL and events and returns the block, then drops the cached /v1/me', async () => {
    const invalidate = vi.spyOn(accountApi, 'invalidateMe');
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
    expect(invalidate).toHaveBeenCalledWith(BASE);
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

  it('types 402 as not-entitled and 500 as failed, and does not invalidate', async () => {
    const invalidate = vi.spyOn(accountApi, 'invalidateMe');
    stub.failWith('put', 402, 'premium');
    await expect(saveDiscordWebhook({ events: EVENTS }, BASE)).rejects.toMatchObject({
      kind: 'not-entitled',
    });
    stub.failWith('put', 500, 'boom');
    await expect(saveDiscordWebhook({ events: EVENTS }, BASE)).rejects.toMatchObject({ kind: 'failed' });
    expect(invalidate).not.toHaveBeenCalled();
  });
});

describe('testDiscordWebhook', () => {
  it('POSTs to the test route and drops the cached /v1/me on 204', async () => {
    const invalidate = vi.spyOn(accountApi, 'invalidateMe');
    await expect(testDiscordWebhook(BASE)).resolves.toBeUndefined();
    expect(stub.requests()[0]).toMatchObject({ method: 'POST', path: '/v1/notifications/discord/test' });
    expect(invalidate).toHaveBeenCalledWith(BASE);
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
  it('DELETEs and drops the cached /v1/me', async () => {
    const invalidate = vi.spyOn(accountApi, 'invalidateMe');
    await deleteDiscordWebhook(BASE);
    expect(stub.requests()[0]).toMatchObject({ method: 'DELETE', path: '/v1/notifications/discord' });
    expect(invalidate).toHaveBeenCalledWith(BASE);
  });

  it('types a failure', async () => {
    stub.failWith('delete', 500, 'boom');
    await expect(deleteDiscordWebhook(BASE)).rejects.toMatchObject({ kind: 'failed' });
  });
});
