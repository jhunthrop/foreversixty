// web/tests/e2e/onboarding-round6.spec.ts
// W.10: onboarding round 6 -- /setup gold on the first undone step only (setup spec 2026-11-07: the
// steps are one Plate, no progress chips), the companion picture shown; the landing card's label and its
// live dot; one prompt on an empty character page. Captured at 1280 and 390 after networkidle plus
// 500 ms (W10_SHOTS names the folder).
import { mkdirSync } from 'node:fs';
import { join } from 'node:path';
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { addonCopy } from '../../src/lib/addon/copy';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  routeCharacterSearch,
  routeGearIcons,
} from './support/character-fixture';
import { expectFixtureNotFound } from './support/character-not-found';

const SHOTS = process.env.W10_SHOTS ?? 'test-results/w10';
const SIZES = [
  { name: '1280', width: 1280, height: 900 },
  { name: '390', width: 390, height: 844 },
] as const;

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});
const ago = (ms: number): string => new Date(Date.now() - ms).toISOString();
const HOUR = 3_600_000;

const ME_DONE = {
  user: { id: 1, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false },
  characters: [
    {
      key: 'us/pvp/bow-jackzon',
      region: 'us',
      ruleset: 'pvp',
      name: 'Bow Jackzon',
      class: 'hunter',
      level: 24,
      build: { source: 'addon', captured_at: ago(HOUR) },
    },
  ],
  guilds: [],
};
const ME_NEW = { ...ME_DONE, characters: [] };
const DEVICE = {
  id: 'd1',
  name: 'JH-DESKTOP',
  platform: 'windows',
  created_at: ago(48 * HOUR),
  last_seen_at: ago(3 * 60_000),
};

async function signedIn(page: Page, me: unknown, devices: unknown[]): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(me)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil(devices)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
}
async function signedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil(null, 401)));
}
async function settle(page: Page): Promise<void> {
  await page.waitForLoadState('networkidle');
  await page.waitForTimeout(500);
}
async function shot(page: Page, name: string, fullPage = true): Promise<void> {
  mkdirSync(SHOTS, { recursive: true });
  await page.screenshot({ path: join(SHOTS, `w10-${name}.png`), fullPage });
}

/** An API answer for a character with a record and nothing on it: no gear, no fights, no score,
 *  and no other character by that name. */
async function routeEmptyCharacter(page: Page): Promise<void> {
  const data = {
    ...CHARACTER.data,
    character: { ...CHARACTER.data.character, class: 'priest', realm: 'Doomhowl' },
    gear: [],
    best: [],
    history: [],
    builds_seen: [],
    score_history: [],
  };
  // The not-found page asks for the same name elsewhere; none here.
  await routeCharacterSearch(page);
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...CHARACTER, data }),
    }),
  );
}

/** The steps carrying the page's one primary (the gold step, W.10), by step number. */
async function goldSteps(page: Page): Promise<string[]> {
  const plate = page.getByRole('region', { name: 'Your setup' });
  const found: string[] = [];
  for (const n of ['1', '2', '3']) {
    const action = plate.getByTestId(`setup-steps-step-${n}`).locator('a, button');
    if ((await action.count()) === 0) continue;
    // The primary Button is the gold gradient; a secondary is a flat ground.
    const image = await action.first().evaluate((el) => getComputedStyle(el).backgroundImage);
    if (image.includes('linear-gradient')) found.push(n);
  }
  return found;
}

/** The companion's Status capture: loaded at its own aspect, never a blank box. */
async function pictureShown(page: Page): Promise<void> {
  const frame: Locator = page.getByTestId('setup-companion-capture');
  await frame.scrollIntoViewIfNeeded();
  const img = frame.locator('img');
  await expect
    .poll(() => img.evaluate((el: HTMLImageElement) => el.complete && el.naturalWidth > 0))
    .toBe(true);
  const box = await img.evaluate((el: HTMLImageElement) => el.getBoundingClientRect().toJSON());
  expect(box.height).toBeGreaterThan(60);
  expect(box.width / box.height).toBeCloseTo(1600 / 436, 0);
}

for (const size of SIZES) {
  test.describe(`at ${size.name}`, () => {
    test.use({ viewport: { width: size.width, height: size.height } });

    test('/setup signed out: gold on Sign in only, the picture shown with its Example tag', async ({
      page,
    }) => {
      await signedOut(page);
      await page.goto('/setup');
      await settle(page);
      expect(await goldSteps(page)).toEqual(['1']);
      await expect(page.getByTestId('setup-steps-step-3')).toContainText(addonCopy.setupCompanionNote);
      await expect(page.getByTestId('setup-companion')).toContainText(addonCopy.setupCompanionNote);
      await pictureShown(page);
      await expect(page.getByTestId('setup-companion-capture').getByText('Example')).toBeVisible();
      await settle(page);
      await shot(page, `setup-signed-out-${size.name}`);
    });

    test('/setup signed in with no character: gold on Refresh from Battle.net only (B2)', async ({
      page,
    }) => {
      await signedIn(page, ME_NEW, []);
      await page.goto('/setup');
      await settle(page);
      await expect(page.getByTestId('setup-steps-step-2')).toContainText(addonCopy.setupAddonNote);
      expect(await goldSteps(page)).toEqual(['1']);
      await pictureShown(page);
      await settle(page);
      await shot(page, `setup-signed-in-pending-${size.name}`);
    });

    test('/setup signed in, everything done: three ticks and no gold', async ({ page }) => {
      await signedIn(page, ME_DONE, [DEVICE]);
      await page.goto('/setup');
      await settle(page);
      await expect(page.getByRole('img', { name: /^Step \d: Done$/ })).toHaveCount(3);
      expect(await goldSteps(page)).toEqual([]);
      await pictureShown(page);
      await settle(page);
      await shot(page, `setup-signed-in-done-${size.name}`);
    });

    test('/setup signed in with the addon done and no companion: gold on the companion step', async ({
      page,
    }) => {
      await signedIn(page, ME_DONE, []);
      await page.goto('/setup');
      await settle(page);
      expect(await goldSteps(page)).toEqual(['3']);
    });

    test('a character the visitor does not hold, with no data: the not-found page', async ({ page }) => {
      await signedOut(page);
      await routeEmptyCharacter(page);
      await page.goto(CHARACTER_PATH);
      await settle(page);
      await expectFixtureNotFound(page, { signedIn: false });
      await expect(page.getByText('Forever Score: no logs yet')).toHaveCount(0);
      await shot(page, `character-no-data-visitor-${size.name}`);
    });

    test('a character another account holds gets the same not-found page, without the sign-in row', async ({
      page,
    }) => {
      await signedIn(page, ME_DONE, []);
      await routeEmptyCharacter(page);
      await page.goto(CHARACTER_PATH);
      await settle(page);
      await expectFixtureNotFound(page, { signedIn: true });
    });

    test('the signed-in holder keeps their page: the head and the owner’s no-logs line', async ({ page }) => {
      const holder = {
        ...ME_DONE,
        characters: [
          {
            key: 'us/hardcore/elyra-duskvale',
            region: 'us',
            ruleset: 'hardcore',
            name: 'Elyra Duskvale',
            class: 'priest',
            level: 60,
          },
        ],
      };
      await signedIn(page, holder, []);
      await routeEmptyCharacter(page);
      await page.goto(CHARACTER_PATH);
      await settle(page);
      await expect(page.getByTestId('character-not-found')).toHaveCount(0);
      await expect(page.getByTestId('character-head').getByTestId('page-title')).toHaveText('Elyra Duskvale');
      await expect(page.getByTestId('character-no-logs')).toHaveText(
        'No logs yet. Upload a log to get a rating.',
      );
      await expect(page.getByTestId('character-gear-section')).toHaveCount(0);
      await expect(page.getByTestId('character-fact-chip')).toHaveText(/hardcore/i);
      await shot(page, `character-no-data-holder-${size.name}`);
    });

    test('a character with gear and no logs: the gear, then "No logs yet." under Forever Score', async ({
      page,
    }) => {
      await signedOut(page);
      const data = { ...CHARACTER.data, ...GEAR, best: [], history: [], builds_seen: [], score_history: [] };
      await routeGearIcons(page);
      await page.route('**/data/*/items/priest.json', (route) =>
        route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) }),
      );
      await page.route(CHARACTER_URL, (route) =>
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ ...CHARACTER, data }),
        }),
      );
      await page.goto(CHARACTER_PATH);
      await settle(page);
      await expect(page.getByTestId('character-gear-section')).toBeVisible();
      await expect(page.getByTestId('character-no-logs')).toHaveText('No logs yet.');
      await expect(page.getByTestId('current-character-bar-signed-out')).toHaveCount(0);
      await shot(page, `character-gear-first-${size.name}`);
    });
  });
}
