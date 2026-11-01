// web/tests/e2e/support/discord-notifications.ts
// Stubs the three Discord notification routes for the account-page specs and records what the
// page sent, so a spec can assert that the typed URL travelled once and never came back.
import type { Page, Route } from '@playwright/test';
import type { DiscordNotifications } from '../../../src/lib/account/api';
import { mockEntitlements } from '../../../src/test-support/billing-fixtures';
import { DEAD_API_BASE } from '../../hermetic';

/** The channel the saved webhook posts to; deliberately not the webhook id a pasted URL carries. */
export const CHANNEL_ID = '1198765432101234568';

/** The Discord server of a connected webhook; distinct from the channel id and any webhook id. */
export const GUILD_ID = '1087654321098765432';

/** The API origin the e2e build links to (playwright.config.ts's webServer env). */
export const E2E_API_BASE = DEAD_API_BASE;
export const CONNECT_URL_GLOB = `${E2E_API_BASE}/v1/notifications/discord/connect`;
export const DISCORD_URL_GLOB = '**/v1/notifications/discord';
export const DISCORD_TEST_URL_GLOB = '**/v1/notifications/discord/test';

export function discordBlock(overrides: Partial<DiscordNotifications> = {}): DiscordNotifications {
  return {
    channel_id: CHANNEL_ID,
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

/** The `/v1/me` `notifications` key for a payer's account. */
export function notificationsKey(
  discord: DiscordNotifications | null,
  connect = false,
): {
  notifications: { connect: boolean; discord: DiscordNotifications | null };
} {
  return { notifications: { connect, discord } };
}

function envelopeBody(data: unknown): string {
  return JSON.stringify({ ok: true, data, error: null, request_id: 'req-e2e' });
}

function failureBody(message: string): string {
  return JSON.stringify({ ok: false, data: null, error: { message }, request_id: 'req-e2e' });
}

export interface DiscordRoutes {
  /** Request bodies the page sent to PUT, oldest first. */
  puts: unknown[];
  tests: number;
  deletes: number;
}

export interface DiscordRouteOptions {
  /** The block PUT answers with. */
  block?: DiscordNotifications;
  /** PUT answers this status and message instead. */
  putFailure?: { status: number; message: string };
  /** The test route answers this instead of 204. */
  testFailure?: { status: number; message: string };
}

/** What the API stores and answers for a PUT: the saved block carrying the events it was sent,
 *  so the page's patch of `me` never puts a switch back to its old position. */
function savedBlock(put: { events?: DiscordNotifications['events'] }): DiscordNotifications {
  return put.events === undefined ? discordBlock() : discordBlock({ events: put.events });
}

function fulfillFailure(route: Route, failed: { status: number; message: string }): Promise<void> {
  return route.fulfill({
    status: failed.status,
    contentType: 'application/json',
    body: failureBody(failed.message),
  });
}

export async function routeDiscordNotifications(
  page: Page,
  options: DiscordRouteOptions = {},
): Promise<DiscordRoutes> {
  const seen: DiscordRoutes = { puts: [], tests: 0, deletes: 0 };
  await page.route(DISCORD_URL_GLOB, (route) => {
    const request = route.request();
    if (request.method() === 'DELETE') {
      seen.deletes += 1;
      return route.fulfill({ status: 204 });
    }
    const put = request.postDataJSON() as { events?: DiscordNotifications['events'] };
    seen.puts.push(put);
    if (options.putFailure !== undefined) return fulfillFailure(route, options.putFailure);
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: envelopeBody(options.block ?? savedBlock(put)),
    });
  });
  await page.route(DISCORD_TEST_URL_GLOB, (route) => {
    seen.tests += 1;
    if (options.testFailure !== undefined) return fulfillFailure(route, options.testFailure);
    return route.fulfill({ status: 204 });
  });
  return seen;
}

/** What the stubbed `/v1/me` answers, as the notification routes are exercised. */
export interface DiscordMeStages {
  entitled?: boolean;
  /** The API says the one-click connect is on. */
  connect?: boolean;
  /** The deployment has notifications off: /v1/me carries no `notifications` key. */
  featureOff?: boolean;
  initial: DiscordNotifications | null;
  /** What /v1/me answers once the PUT has been seen. */
  afterPut?: DiscordNotifications | null;
  /** What /v1/me answers once a test has been sent. */
  afterTest?: DiscordNotifications | null;
  /** Awaited before every `/v1/me` and `/v1/devices` answer: holds the page in its loading state. */
  hold?: () => Promise<void>;
}

export function envelope(data: unknown): { status: number; contentType: string; body: string } {
  return { status: 200, contentType: 'application/json', body: envelopeBody(data) };
}

function currentBlock(stages: DiscordMeStages, seen: DiscordRoutes): DiscordNotifications | null {
  if (seen.deletes > 0) return null;
  if (seen.tests > 0 && stages.afterTest !== undefined) return stages.afterTest;
  if (seen.puts.length > 0 && stages.afterPut !== undefined) return stages.afterPut;
  return stages.initial;
}

/** Signs a member in; /v1/me follows what the stubbed notification routes have been sent. */
export async function signInMember(page: Page, stages: DiscordMeStages, seen: DiscordRoutes): Promise<void> {
  await page.route('**/v1/me', async (route) => {
    await stages.hold?.();
    await route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ notifications: stages.entitled ?? true }),
        ...(stages.featureOff === true
          ? {}
          : notificationsKey(currentBlock(stages, seen), stages.connect ?? false)),
      }),
    );
  });
  await page.route('**/v1/reports**', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 20 })),
  );
  await page.route('**/v1/devices', async (route) => {
    await stages.hold?.();
    await route.fulfill(envelope([]));
  });
}
