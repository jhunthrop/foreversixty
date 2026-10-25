// web/tests/e2e/support/discord-notifications.ts
// Stubs the three Discord notification routes for the account-page specs and records what the
// page sent, so a spec can assert that the typed URL travelled once and never came back.
import type { Page, Route } from '@playwright/test';
import type { DiscordNotifications } from '../../../src/lib/account/api';

/** The channel the saved webhook posts to; deliberately not the webhook id a pasted URL carries. */
export const CHANNEL_ID = '1198765432101234568';

export const DISCORD_URL_GLOB = '**/v1/notifications/discord';
export const DISCORD_TEST_URL_GLOB = '**/v1/notifications/discord/test';

export function discordBlock(overrides: Partial<DiscordNotifications> = {}): DiscordNotifications {
  return {
    channel_id: CHANNEL_ID,
    events: { sim_finished: true, upload_parsed: true },
    disabled_reason: null,
    last_error: null,
    failures: 0,
    ...overrides,
  };
}

/** The `/v1/me` `notifications` key for a payer's account. */
export function notificationsKey(discord: DiscordNotifications | null): {
  notifications: { discord: DiscordNotifications | null };
} {
  return { notifications: { discord } };
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
    seen.puts.push(request.postDataJSON());
    if (options.putFailure !== undefined) return fulfillFailure(route, options.putFailure);
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: envelopeBody(options.block ?? discordBlock()),
    });
  });
  await page.route(DISCORD_TEST_URL_GLOB, (route) => {
    seen.tests += 1;
    if (options.testFailure !== undefined) return fulfillFailure(route, options.testFailure);
    return route.fulfill({ status: 204 });
  });
  return seen;
}
