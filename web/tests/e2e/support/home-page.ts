// web/tests/e2e/support/home-page.ts
// The home page's signed-in states from committed fixtures (tests/e2e/fixtures/home/, written by
// design/mocks/home_fixtures.py): the session cookie, /v1/me, the sims, reports and guild reads,
// and the trimmed BiS and item files. The parity gate and the home e2e both open home through it.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, type Page } from '@playwright/test';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures/home');
/** The fixtures' own "today": the export ages and the raid countdown read against it. */
export const HOME_NOW = new Date('2026-11-06T18:00:00Z');

export interface HomeState {
  /** A fixture file, or `pending` for a /v1/me that never answers (the loading state). */
  me?: string;
  sims?: string;
  reports?: string;
  guild?: string;
  /** The BiS file fails (the upgrades error state). */
  failBis?: boolean;
}

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

const envelope = (data: unknown) => ({ json: { ok: true, data, error: null, request_id: 'home' } });

/** Routes everything the signed-in home reads. Signed out (no `me`) it routes nothing. */
export async function routeHome(page: Page, state: HomeState): Promise<void> {
  if (state.me === undefined) return;
  await page.clock.setFixedTime(HOME_NOW);
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  if (state.me === 'pending') {
    await page.route('**/v1/me', () => new Promise<void>(() => {}));
    return;
  }
  const me = fixture(state.me);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me)));
  const sims = state.sims === undefined ? { rows: [], total: 0, page: 1, per_page: 20 } : fixture(state.sims);
  await page.route('**/v1/sims?**', (route) => route.fulfill(envelope(sims)));
  const reports =
    state.reports === undefined ? { rows: [], total: 0, page: 1, per_page: 100 } : fixture(state.reports);
  await page.route('**/v1/reports?mine=1**', (route) => route.fulfill(envelope(reports)));
  if (state.guild !== undefined) {
    const home = fixture(state.guild);
    await page.route('**/v1/guilds/*/home', (route) => route.fulfill(envelope(home)));
  }
  for (const [spec, cls] of [
    ['hunter-beast-mastery', 'hunter'],
    ['warrior-fury', 'warrior'],
  ]) {
    await page.route(`**/data/*/bis/${spec}.json`, (route) =>
      state.failBis
        ? route.fulfill({ status: 500, body: '' })
        : route.fulfill({ json: fixture(`bis-${spec}.json`) }),
    );
    await page.route(`**/data/*/items/${cls}.json`, (route) =>
      route.fulfill({ json: fixture(`items-${cls}.json`) }),
    );
  }
}

/** Opens home and waits for the state to settle: the island hydrated and its reads answered. */
export async function openHome(page: Page, state: HomeState): Promise<void> {
  await routeHome(page, state);
  await page.goto('/');
  await expect(page.locator('astro-island:has([data-testid="home-signed-in-root"]):not([ssr])')).toHaveCount(
    1,
  );
  if (state.me === undefined || state.me === 'pending') return;
  await expect(page.locator('[data-testid="home-upgrades"] [aria-busy="true"]')).toHaveCount(0);
  await page.waitForLoadState('networkidle');
}
