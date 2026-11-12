// web/tests/e2e/support/home-page.ts
// The home page's states from committed fixtures (tests/e2e/fixtures/home/, written by
// design/mocks/home_fixtures.py): the session cookie, /v1/me, /v1/feed, /v1/home, the guild read,
// and the trimmed BiS and item files. The parity gate and the home e2e both open home through it.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, type Page } from '@playwright/test';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures/home');
/** The fixtures' own "today": the feed's times and the countdowns read against it. */
export const HOME_NOW = new Date('2026-11-06T18:00:00Z');

export interface HomeState {
  /** A fixture file, or `pending` for a /v1/me that never answers (the loading state). */
  me?: string;
  /** A fixture file, `pending` (never answers) or `fail` (a 500). */
  feed?: string;
  /** The /v1/home answer. */
  home?: string;
  guild?: string;
}

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'home' } });

/** A fixture as the API's envelope, for a test that re-routes a read. */
export const homeEnvelope = (name: string) => envelope(fixture(name));

/** Every state routes the feed: the signed-out page's island reads it even while hidden. */
async function routeFeed(page: Page, feed = 'feed-empty.json'): Promise<void> {
  await page.route('**/v1/feed?**', (route) => {
    if (feed === 'pending') return new Promise<void>(() => {});
    if (feed === 'fail')
      return route.fulfill({
        status: 500,
        json: { ok: false, data: null, error: { code: 'internal', message: 'internal' } },
      });
    return route.fulfill(envelope(fixture(feed)));
  });
}

/** Routes everything home reads: signed out only the feed, signed in the rest too. */
export async function routeHome(page: Page, state: HomeState): Promise<void> {
  await page.clock.setFixedTime(HOME_NOW);
  await routeFeed(page, state.feed);
  if (state.me === undefined) return;
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  if (state.me === 'pending') {
    await page.route('**/v1/me', () => new Promise<void>(() => {}));
    return;
  }
  const me = fixture(state.me);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me)));
  if (state.home !== undefined) {
    const home = fixture(state.home);
    await page.route('**/v1/home?**', (route) => route.fulfill(envelope(home)));
  }
  if (state.guild !== undefined) {
    const home = fixture(state.guild);
    await page.route('**/v1/guilds/*/home', (route) => route.fulfill(envelope(home)));
  }
  for (const [spec, cls] of [
    ['hunter-beast-mastery', 'hunter'],
    ['warrior-fury', 'warrior'],
    ['warrior-protection', 'warrior'],
  ]) {
    await page.route(`**/data/*/bis/${spec}.json`, (route) =>
      route.fulfill({ json: fixture(`bis-${spec}.json`) }),
    );
    await page.route(`**/data/*/items/${cls}.json`, (route) =>
      route.fulfill({ json: fixture(`items-${cls}.json`) }),
    );
  }
}

/** Opens home and waits for the state to settle: the islands hydrated and their reads answered. */
export async function openHome(page: Page, state: HomeState): Promise<void> {
  await routeHome(page, state);
  await page.goto('/');
  await expect(page.locator('astro-island:has([data-testid="home-signed-in-root"]):not([ssr])')).toHaveCount(
    1,
  );
  if (state.me === 'pending') return;
  if (state.me === undefined) {
    await expect(page.locator('astro-island:has([data-testid="home-feed"]):not([ssr])')).toHaveCount(1);
  }
  if (state.home !== undefined) await expect(page.getByTestId('home-figure-sim')).toBeVisible();
  if (state.feed !== undefined && state.feed !== 'pending') {
    await expect(page.locator('[data-testid="home-feed"]:visible [aria-busy="true"]')).toHaveCount(0);
  }
  if (state.feed !== 'pending') await page.waitForLoadState('networkidle');
}
