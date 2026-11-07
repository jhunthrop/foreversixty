// web/tests/e2e/support/pages-page.ts
// The changelog, the content pages and the 404 in their account states (pages spec 2026-11-08
// section 5.1) from the committed fixture the mocks draw: Felsnap is fixtures/home/me-34.json (a
// level 34 Beast Mastery hunter), the account pages_facts.mjs reads. Signed out the account read
// answers 401; signed in the session hint (fs_csrf) and the fixture.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { expect } from './test';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures');

export type PagesState = 'out' | 'in' | 'loading' | 'error' | 'none';

const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const failure = (status: number) =>
  json({ ok: false, data: null, error: { message: 'pages fixture' }, request_id: 'pages' }, status);

/** Felsnap's account (fixtures/home/me-34.json). */
export function felsnapAccount(): { characters: unknown[] } & Record<string, unknown> {
  return JSON.parse(readFileSync(path.join(FIXTURES, 'home/me-34.json'), 'utf8')) as {
    characters: unknown[];
  };
}

/** Routes /v1/me for `state`: 401 signed out; else the session hint and the account read. */
export async function routePagesState(page: Page, state: PagesState): Promise<void> {
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
  const account = state === 'none' ? { ...felsnapAccount(), characters: [] } : felsnapAccount();
  await page.route('**/v1/me', (route) =>
    route.fulfill(json({ ok: true, data: account, error: null, request_id: 'pages' })),
  );
}

/** The 404 answer's settled states (CharacterAnswer's `data-state` once painted). */
const SETTLED_ANSWER = /^(ready|unknown|none|error|hidden)$/;

/**
 * Opens `pathname` in `state` and waits for it to settle: on the 404 signed in, until its answer
 * leaves `loading` (`loading` itself only until the reserve shows); then the network is idle.
 */
export async function openPagesState(page: Page, pathname: string, state: PagesState): Promise<void> {
  await routePagesState(page, state);
  await page.goto(pathname);
  const answer = page.getByTestId('nf-answer');
  if (state === 'loading') {
    await expect(answer).toBeVisible();
    return;
  }
  if (state !== 'out' && (await answer.count()) > 0) {
    await expect(answer).toHaveAttribute('data-state', SETTLED_ANSWER);
  }
  await page.waitForLoadState('networkidle');
}
