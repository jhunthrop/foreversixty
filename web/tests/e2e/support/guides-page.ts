// web/tests/e2e/support/guides-page.ts
// /guides, /guides/<class> and /guides/<class>/<spec> in their account states (guides spec
// 2026-11-08 sections 4.4 and 5.1) from the committed fixtures the guide mocks draw: Felsnap is
// fixtures/home/me-34.json (a level 34 Beast Mastery hunter), Thandrel fixtures/home/me-60.json (a
// level 60 Fury warrior), and the derived Survival and Fire accounts are fixtures/guides/
// (README.md there). guides_facts.mjs reads the same files, so both sides draw one account.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { expect } from './test';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../fixtures');

export type GuidesState = 'out' | 'hunter34' | 'survival34' | 'warrior60' | 'fire60' | 'loading' | 'error';

/** The account fixture each signed-in state serves from /v1/me. */
const ACCOUNTS: Readonly<Record<'hunter34' | 'survival34' | 'warrior60' | 'fire60', string>> = {
  hunter34: 'home/me-34.json',
  survival34: 'guides/me-survival-34.json',
  warrior60: 'home/me-60.json',
  fire60: 'guides/me-fire-60.json',
};

/** The answer's `data-state` values once the account read has settled (the page contract). */
const SETTLED_ANSWER = /^(ready|unknown|none|error)$/;
const ANSWER = '[data-testid="guide-answer"]';

const json = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});
const envelope = (data: unknown) => json({ ok: true, data, error: null, request_id: 'guides' });
const failure = (status: number) =>
  json({ ok: false, data: null, error: { message: 'guides fixture' }, request_id: 'guides' }, status);

function fixture(name: string): unknown {
  return JSON.parse(readFileSync(path.join(FIXTURES, name), 'utf8'));
}

/**
 * Routes /v1/me for `state`: 401 signed out; signed in, the session hint (the fs_csrf cookie)
 * and then the account fixture, a read that never answers (`loading`) or a 500 (`error`).
 */
export async function routeGuidesState(page: Page, state: GuidesState): Promise<void> {
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

/**
 * Opens `pathname` in `state` and waits for it to settle: signed in, until the header answer
 * (`[data-testid="guide-answer"]`) leaves `data-state="loading"` for a settled state; `loading`,
 * only until the reserved answer is visible (its read never answers, so no networkidle either);
 * signed out, the page alone. Then the network is idle.
 */
export async function openGuidesState(page: Page, pathname: string, state: GuidesState): Promise<void> {
  await routeGuidesState(page, state);
  await page.goto(pathname);
  const answer = page.locator(ANSWER);
  if (state === 'loading') {
    await expect(answer).toBeVisible();
    return;
  }
  if (state !== 'out') await expect(answer).toHaveAttribute('data-state', SETTLED_ANSWER);
  await page.waitForLoadState('networkidle');
}

/** The guide's leveling lane (GuideLevelingTree). */
export const levelingTree = (page: Page) => page.getByTestId('guide-leveling-tree');

/**
 * Opens a reading section's phone fold (ReadingSection `fold`): below 768 the page closes it
 * before the section scrolls into view (lib/ui/fold.ts), so its body has no box until a player
 * taps the fold's line. From 768 the line is hidden and the section is already open.
 */
export async function openFold(page: Page, anchor: string): Promise<void> {
  const line = page.locator(`#${anchor} details[data-fold] > summary`);
  if ((await line.count()) > 0 && (await line.isVisible())) await line.click();
}

/** Opens the Leveling fold, scrolls the lane into view and waits for its island (client:visible) to hydrate. */
export async function hydratedLevelingTree(page: Page): Promise<void> {
  await openFold(page, 'leveling');
  await levelingTree(page).scrollIntoViewIfNeeded();
  await expect(page.locator('astro-island:has([data-testid="guide-leveling-tree"])')).not.toHaveAttribute(
    'ssr',
    /.*/,
  );
}
