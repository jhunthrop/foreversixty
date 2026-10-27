// web/src/lib/billing/seats-api.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { heldSeat, SEAT_CANDIDATES } from '../../test-support/billing-fixtures';
import { invalidateGuild } from '../guild/api';
import { BillingApiError } from './api';
import {
  acceptSeatOffer,
  declineSeatOffer,
  giveSeat,
  moveSeat,
  seatCandidates,
  withdrawSeat,
} from './seats-api';

vi.mock('../guild/api', () => ({ invalidateGuild: vi.fn() }));

const API = 'https://api.foreversixty.test';
type GlobalFetch = (...args: Parameters<typeof fetch>) => Promise<Response>;

interface Seen {
  method: string;
  path: string;
  body: unknown;
}

function envelope(
  data: unknown,
  status = 200,
  error: { code: string; message: string; fields?: Record<string, string> } | null = null,
): Response {
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

afterEach(() => {
  vi.unstubAllGlobals();
  vi.mocked(invalidateGuild).mockClear();
});

describe('seat reads', () => {
  it('GETs the candidates and unwraps the list', async () => {
    const seen = stubFetch(() => envelope({ candidates: SEAT_CANDIDATES }));
    expect(await seatCandidates(5, API)).toEqual(SEAT_CANDIDATES);
    expect(seen[0]).toMatchObject({ method: 'GET', path: '/v1/guilds/5/seats/candidates' });
  });
});

describe('seat writes', () => {
  it('gives, moves and withdraws on the slot path with the user id', async () => {
    const seen = stubFetch(() => envelope(heldSeat(2)));
    await giveSeat(5, 2, 31, API);
    await moveSeat(5, 2, 32, API);
    await withdrawSeat(5, 2, API);
    expect(seen.map(({ method, path }) => `${method} ${path}`)).toEqual([
      'POST /v1/guilds/5/seats/2/give',
      'POST /v1/guilds/5/seats/2/move',
      'POST /v1/guilds/5/seats/2/withdraw',
    ]);
    expect(seen[0]?.body).toEqual({ user_id: 31 });
    expect(seen[1]?.body).toEqual({ user_id: 32 });
  });

  it('returns the seat as it now stands', async () => {
    stubFetch(() => envelope(heldSeat(1)));
    expect((await giveSeat(5, 1, 31, API)).state).toBe('held');
  });

  it('accepts and declines an offer, resolving on an empty answer', async () => {
    const seen = stubFetch(() => new Response(null, { status: 204 }));
    await acceptSeatOffer(5, API);
    await declineSeatOffer(5, API);
    expect(seen.map(({ method, path }) => `${method} ${path}`)).toEqual([
      'POST /v1/guilds/5/seats/offer/accept',
      'POST /v1/guilds/5/seats/offer/decline',
    ]);
  });
});

describe('seat errors', () => {
  it('leaves the guild cache alone when an action fails', async () => {
    stubFetch(() => envelope(null, 409, { code: 'seat_locked', message: 'Locked.' }));
    await expect(moveSeat(5, 1, 31, API)).rejects.toBeInstanceOf(BillingApiError);
    await expect(acceptSeatOffer(5, API)).rejects.toBeInstanceOf(BillingApiError);
    expect(invalidateGuild).not.toHaveBeenCalled();
  });

  it('drops the guild cache after a successful give', async () => {
    stubFetch(() => envelope(heldSeat(1)));
    await giveSeat(5, 1, 31, API);
    expect(invalidateGuild).toHaveBeenCalledWith(5, API);
  });

  it('propagates the code and movable_at of a 409 seat_locked', async () => {
    stubFetch(() =>
      envelope(null, 409, {
        code: 'seat_locked',
        message: 'Locked.',
        fields: { movable_at: '2026-12-02T00:00:00Z' },
      }),
    );
    const failure = await moveSeat(5, 1, 31, API).catch((error: unknown) => error);
    expect(failure).toBeInstanceOf(BillingApiError);
    expect(failure).toMatchObject({
      status: 409,
      code: 'seat_locked',
      fields: { movable_at: '2026-12-02T00:00:00Z' },
    });
  });

  it('propagates not_member', async () => {
    stubFetch(() => envelope(null, 409, { code: 'not_member', message: 'No.' }));
    await expect(giveSeat(5, 1, 99, API)).rejects.toMatchObject({ code: 'not_member' });
  });

  it('reports a network failure with status 0 and no code', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn<GlobalFetch>(async () => {
        throw new TypeError('offline');
      }),
    );
    await expect(withdrawSeat(5, 1, API)).rejects.toMatchObject({ status: 0, code: null });
  });
});
