// web/src/lib/billing/guild-billing-api.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { BillingApiError } from './api';
import {
  declineTransfer,
  exportHref,
  getGuildBilling,
  previewUpgrade,
  sendTransfer,
  upgrade,
  withdrawTransfer,
} from './guild-billing-api';
import { guildBillingRead, UPGRADE_PREVIEW } from '../../test-support/billing-fixtures';

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

interface Seen {
  method: string;
  path: string;
  body: unknown;
}

function envelope(data: unknown, status = 200, code?: string): Response {
  const error = status >= 400 ? { code: code ?? 'error', message: 'No.' } : null;
  return new Response(JSON.stringify({ ok: status < 400, data, error, request_id: 'r' }), {
    status,
    headers: { 'content-type': 'application/json' },
  });
}

function stubFetch(respond: () => Response): Seen[] {
  const seen: Seen[] = [];
  vi.stubGlobal(
    'fetch',
    vi.fn<GlobalFetch>(async (input) => {
      const request = input as Request;
      const text = await request.text();
      seen.push({
        method: request.method,
        path: new URL(request.url).pathname,
        body: text === '' ? undefined : JSON.parse(text),
      });
      return respond();
    }),
  );
  return seen;
}

afterEach(() => vi.unstubAllGlobals());

describe('guild billing reads', () => {
  it('GETs the billing read', async () => {
    const seen = stubFetch(() => envelope(guildBillingRead()));
    const read = await getGuildBilling(5, API);
    expect(read.verified_count).toBe(12);
    expect(seen[0]).toMatchObject({ method: 'GET', path: '/v1/guilds/5/billing' });
  });

  it('GETs the upgrade preview', async () => {
    const seen = stubFetch(() => envelope(UPGRADE_PREVIEW));
    expect(await previewUpgrade(5, API)).toEqual(UPGRADE_PREVIEW);
    expect(seen[0]).toMatchObject({ method: 'GET', path: '/v1/guilds/5/billing/upgrade-preview' });
  });

  it('answers a 404 as a BillingApiError for a member', async () => {
    stubFetch(() => envelope(null, 404, 'not_found'));
    await expect(getGuildBilling(5, API)).rejects.toMatchObject({ status: 404, code: 'not_found' });
  });
});

describe('guild billing writes', () => {
  it('POSTs the upgrade and returns the read', async () => {
    const seen = stubFetch(() => envelope(guildBillingRead()));
    await upgrade(5, API);
    expect(seen[0]).toMatchObject({ method: 'POST', path: '/v1/guilds/5/billing/upgrade' });
  });

  it('carries upgrade_failed on a declined card', async () => {
    stubFetch(() => envelope(null, 409, 'upgrade_failed'));
    const failure = await upgrade(5, API).catch((error: unknown) => error);
    expect(failure).toBeInstanceOf(BillingApiError);
    expect(failure).toMatchObject({ status: 409, code: 'upgrade_failed' });
  });

  it('posts transfer, withdraw and decline to their paths', async () => {
    const seen = stubFetch(() => envelope({}));
    await sendTransfer(5, 32, API);
    await withdrawTransfer(5, API);
    await declineTransfer(5, API);
    expect(seen.map(({ method, path }) => `${method} ${path}`)).toEqual([
      'POST /v1/guilds/5/billing/transfer',
      'POST /v1/guilds/5/billing/transfer/withdraw',
      'POST /v1/guilds/5/billing/transfer/decline',
    ]);
    expect(seen[0]?.body).toEqual({ to_user_id: 32 });
  });

  it('resolves a transfer answer with no data', async () => {
    stubFetch(() => new Response(null, { status: 204 }));
    await expect(withdrawTransfer(5, API)).resolves.toBeUndefined();
  });

  it('carries not_member for a target who is not an officer', async () => {
    stubFetch(() => envelope(null, 409, 'not_member'));
    await expect(sendTransfer(5, 9, API)).rejects.toMatchObject({ status: 409, code: 'not_member' });
  });
});

describe('exportHref', () => {
  it('points at the API origin for each kind', () => {
    expect(exportHref(5, 'attendance', API)).toBe(`${API}/v1/guilds/5/export/attendance.csv`);
    expect(exportHref(5, 'loot', API)).toBe(`${API}/v1/guilds/5/export/loot.csv`);
  });
});
