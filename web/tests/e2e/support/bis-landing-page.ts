// web/tests/e2e/support/bis-landing-page.ts
// /bis in its account states (design/specs/2026-11-07-bis-landing.md section 6) from the committed
// fixtures the landing mocks draw (fixtures/bis-landing/README.md): signed out, Felsnap at 24, 34
// and with no gear, Thandrel at 60, an account with no character, Felsnap with a finished server
// sim, the account read in flight and failed. The spec lists are served from this checkout's
// data/builds/<active build>/bis, the files design/mocks/bis_landing_facts.mjs counts against, so
// both sides count the same picks.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { ACTIVE_BUILD } from './active-build';
import { expect } from './test';
import { REPO_ROOT } from './parity';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures');

export type BisLandingState = 'out' | '24' | '34' | '60' | 'new' | 'nogear' | 'sim' | 'loading' | 'error';

const ACCOUNTS: Readonly<Record<Exclude<BisLandingState, 'out' | 'loading' | 'error'>, string>> = {
  '24': 'bis-landing/me-24.json',
  '34': 'home/me-34.json',
  '60': 'home/me-60.json',
  new: 'home/me-new.json',
  nogear: 'bis-landing/me-nogear.json',
  sim: 'bis-landing/me-sim.json',
};

/** The answer's `data-state` once the account read has settled (the page contract). */
const SETTLED_ANSWER = /^(ready|none|error|hidden)$/;

const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const envelope = (data: unknown) => json({ ok: true, data, error: null, request_id: 'bis-landing' });
const failure = (status: number) =>
  json(
    { ok: false, data: null, error: { message: 'bis landing fixture' }, request_id: 'bis-landing' },
    status,
  );

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

/** Serves every spec's list from data/builds, what the mock facts count against. */
async function routeBisFiles(page: Page): Promise<void> {
  await page.route('**/data/*/bis/*.json', (route) => {
    const file = path.basename(new URL(route.request().url()).pathname);
    const body = readFileSync(path.join(REPO_ROOT, 'data/builds', ACTIVE_BUILD, 'bis', file), 'utf8');
    return route.fulfill({ status: 200, contentType: 'application/json', body });
  });
}

/** Routes /v1/me for `state`: 401 signed out; signed in, the session hint and then the fixture,
 *  a read that never answers (`loading`) or a 500 (`error`). */
export async function routeBisLandingState(page: Page, state: BisLandingState): Promise<void> {
  await routeBisFiles(page);
  if (state === 'out') {
    await page.route('**/v1/me', (route) => route.fulfill(failure(401)));
    return;
  }
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  if (state === 'loading') {
    await page.route('**/v1/me', () => new Promise<void>(() => {}));
    return;
  }
  if (state === 'error') {
    await page.route('**/v1/me', (route) => route.fulfill(failure(500)));
    return;
  }
  const account = fixture(ACCOUNTS[state]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(account)));
}

/** Opens `pathname` in `state` and waits for the answer (and the finished line) to settle. */
export async function openBisLanding(page: Page, pathname: string, state: BisLandingState): Promise<void> {
  await routeBisLandingState(page, state);
  await page.goto(pathname);
  const answer = page.getByTestId('bis-answer');
  if (state === 'loading') {
    await expect(answer).toBeVisible();
    return;
  }
  if (state !== 'out') await expect(answer).toHaveAttribute('data-state', SETTLED_ANSWER);
  if (state === 'sim') await expect(page.getByTestId('finished-sim-line')).toBeVisible();
  await page.waitForLoadState('networkidle');
}
