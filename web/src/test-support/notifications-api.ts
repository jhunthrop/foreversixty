// web/src/test-support/notifications-api.ts
// Fakes for the three Discord notification routes and a `/v1/me` block builder, in the idiom of
// history-api.ts: a fetch stub answering the section-10 envelope, with each request recorded.
// An unmatched request throws rather than reaching the network.
import { vi } from 'vitest';
import type { DiscordNotifications } from '../lib/account/api';
import { envelope, failure } from './sim-api';

export const DISCORD_ROUTE = /\/v1\/notifications\/discord(\/test)?$/;

/** The Discord server of a connected webhook; distinct from the channel and webhook ids. */
export const GUILD_ID = '1087654321098765432';

/** A saved, working pasted webhook (no guild known); override any field. */
export function discordBlock(overrides: Partial<DiscordNotifications> = {}): DiscordNotifications {
  return {
    channel_id: '1098765432101234567',
    source: 'pasted',
    guild_id: null,
    name: 'Raid pings',
    events: { sim_finished: true, upload_parsed: true, upload_private: false },
    disabled_reason: null,
    last_error: null,
    failures: 0,
    ...overrides,
  };
}

export interface RecordedRequest {
  method: string;
  path: string;
  body: unknown;
}

export interface NotificationsApiStub {
  install(): void;
  reset(): void;
  /** The block `PUT` answers with on 200. */
  setBlock(block: DiscordNotifications): void;
  /** Makes the next matching route answer a failure: `test` and `put` and `delete`. */
  failWith(route: 'put' | 'delete' | 'test', status: number, message: string): void;
  requests(): readonly RecordedRequest[];
}

export function createNotificationsApi(): NotificationsApiStub {
  let block = discordBlock();
  let failures: Partial<Record<'put' | 'delete' | 'test', { status: number; message: string }>> = {};
  let recorded: RecordedRequest[] = [];

  async function handle(request: Request): Promise<Response> {
    const url = new URL(request.url);
    if (!DISCORD_ROUTE.test(url.pathname))
      throw new Error(`notifications stub: no route for ${url.pathname}`);
    const text = await request.text();
    recorded = [
      ...recorded,
      { method: request.method, path: url.pathname, body: text === '' ? null : JSON.parse(text) },
    ];
    const route = url.pathname.endsWith('/test') ? 'test' : request.method === 'PUT' ? 'put' : 'delete';
    const failed = failures[route];
    if (failed !== undefined) return failure(failed.message, failed.status);
    if (route === 'put') return envelope(block);
    return new Response(null, { status: 204 });
  }

  return {
    install: () => {
      vi.stubGlobal('fetch', vi.fn(handle));
    },
    reset: () => {
      vi.unstubAllGlobals();
      block = discordBlock();
      failures = {};
      recorded = [];
    },
    setBlock: (value) => {
      block = value;
    },
    failWith: (route, status, message) => {
      failures = { ...failures, [route]: { status, message } };
    },
    requests: () => recorded,
  };
}
