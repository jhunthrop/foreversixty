// web/tests/e2e/support/planner-page.ts
// /planner in the states its mocks draw (design/specs/2026-11-08-planner.md section 10, mocks
// design/mocks/web/web-planner*.html), from the files the mocks read: the trees, icons, tree art,
// band files and reference data of data/builds/<active build> are served to the fixture build in
// place of its trimmed ones (as support/sim-parity.ts serves the simulator's), and the signed-in
// account is fixtures/planner/me-34-talents.json (Felsnap, a level 34 Dwarf Beast Mastery hunter
// whose 25 points are real ranks: gap G5). planner_facts.mjs reads the same files, so both sides
// draw one build.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, type Page } from '@playwright/test';
import { plannerCopy } from '../../../src/lib/planner/copy';

const WEB_ROOT = path.join(import.meta.dirname, '..', '..', '..');
const BUILD = (
  JSON.parse(readFileSync(path.join(WEB_ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
    build: string;
  }
).build;
const REAL_BUILD_DIR = path.join(WEB_ROOT, '..', 'data', 'builds', BUILD);
const FIXTURES = path.join(WEB_ROOT, 'tests', 'e2e', 'fixtures');

/** The guide codes the guide mocks load (design/mocks/planner_facts.mjs `GUIDE`). */
export const PLANNER_CODES = {
  survival: `FS1:${BUILD}:hunter:night-elf:5/005005201/500230230050222151:`,
  holy: `FS1:${BUILD}:priest:human:00500503103/33505003030121431/0:`,
} as const;

/** The signed-in account: Felsnap with real ranks (gap G5). */
export const FELSNAP_ACCOUNT = 'planner/me-34-talents.json';

export type PlannerState = 'out' | 'felsnap';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'planner' }),
});
const failure = (status: number) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: false,
    data: null,
    error: { message: 'planner fixture' },
    request_id: 'planner',
  }),
});

export function plannerAccount(): { characters: { key: string; name: string }[] } {
  return JSON.parse(readFileSync(path.join(FIXTURES, FELSNAP_ACCOUNT), 'utf8')) as {
    characters: { key: string; name: string }[];
  };
}

/** Every /data/<build>/ file answered from the committed build when it is there. */
async function serveRealBuild(page: Page): Promise<void> {
  await page.route(`**/data/${BUILD}/**`, async (route) => {
    const rel = decodeURIComponent(new URL(route.request().url()).pathname).replace(`/data/${BUILD}/`, '');
    const file = path.join(REAL_BUILD_DIR, rel);
    if (!existsSync(file)) return route.continue();
    return route.fulfill({ path: file });
  });
}

/** Routes the account for `state` (signed out: 401) and the devices read the send consults. */
export async function routePlannerState(page: Page, state: PlannerState): Promise<void> {
  await serveRealBuild(page);
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  if (state === 'out') {
    await page.route('**/v1/me', (route) => route.fulfill(failure(401)));
    return;
  }
  const port = process.env.E2E_PORT ?? '4321';
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', url: `http://localhost:${port}` }]);
  const account = plannerAccount();
  await page.route('**/v1/me', (route) => route.fulfill(envelope(account)));
}

/**
 * Waits for the planner's own ready signals in order, each a thing that happened rather than a
 * longer clock: the island hydrated (Astro drops `ssr` from it; `client:after-paint` starts it
 * after the first paint, so a cold runner can spend most of a default wait before it begins),
 * then the trees drawn or the load failed. A failed load stops here with its message instead of
 * timing out on trees that are never coming (the 1440 flake on CI).
 */
export async function plannerReady(page: Page): Promise<void> {
  await expect(page.getByTestId('planner')).toBeAttached();
  await expect(page.locator('astro-island[ssr]:has([data-testid="planner"])')).toHaveCount(0);
  const trees = page.getByTestId('planner-trees');
  const failed = page.getByTestId('planner-load-error');
  await expect(trees.or(failed)).toBeVisible();
  if (await failed.isVisible())
    throw new Error(`the planner's talent data failed to load: ${await failed.innerText()}`);
}

/** Opens `url` in `state` and waits for the trees (or, signed in, the character's level) to settle. */
export async function openPlanner(page: Page, url: string, state: PlannerState = 'out'): Promise<void> {
  await routePlannerState(page, state);
  await page.goto(url);
  await plannerReady(page);
  // A build that matches its level build draws no level card; its guide row carries the Loaded tag.
  await expect(
    page
      .getByTestId('band-compare-differs')
      .or(page.getByTestId('band-compare-empty'))
      .or(page.getByTestId('planner-guide-builds').getByText(plannerCopy.guideBuildLoaded, { exact: true }))
      .first(),
  ).toBeVisible();
  await page.waitForLoadState('networkidle');
}

/** Clicks the talent named `name` (its cell's accessible name starts with it). */
export async function addTalent(page: Page, name: string): Promise<void> {
  await page
    .getByRole('button', { name: new RegExp(`^${name}, rank`) })
    .first()
    .click();
}
