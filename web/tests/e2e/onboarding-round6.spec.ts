// web/tests/e2e/onboarding-round6.spec.ts
// W.10: onboarding round 6 -- /setup gold on the first undone step only, one label per progress chip
// that never wraps at 390, the companion picture's phone source; the landing card's label and its
// live dot; one prompt on an empty character page. Captured at 1280 and 390 after networkidle plus
// 500 ms (W10_SHOTS names the folder).
import { mkdirSync } from 'node:fs';
import { join } from 'node:path';
import { expect, test, type Locator, type Page } from '@playwright/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  routeGearIcons,
} from './support/character-fixture';
import { routeSampleReportMeta } from './support/logs-sample-fixture';

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

/** An API answer for a character with a record and nothing on it: no gear, no fights, no score. */
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
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...CHARACTER, data }),
    }),
  );
}

/** The visible step buttons drawn as the primary recipe (the gold gradient), by step number. */
async function goldSteps(page: Page): Promise<string[]> {
  const found: string[] = [];
  for (const button of await page.locator('.setup-btn').all()) {
    if (!(await button.isVisible())) continue;
    const image = await button.evaluate((el) => getComputedStyle(el).backgroundImage);
    if (image.includes('linear-gradient')) found.push((await button.getAttribute('data-step')) ?? '?');
  }
  return found;
}

async function chipRowIsOneLine(page: Page): Promise<void> {
  const tops = await page
    .locator('.setup-chip-link')
    .evaluateAll((links) => links.map((el) => Math.round(el.getBoundingClientRect().top)));
  expect(new Set(tops).size, `chip tops ${tops.join(',')}`).toBe(1);
  const labels = await page.locator('[data-testid="setup-chip-label"]:visible').allTextContents();
  expect(labels[1]).toBe('Addon');
  expect(labels[2]).toBe('Companion');
  expect(['Sign in', 'Signed in']).toContain(labels[0]);
  const height = await page
    .locator('.setup-chip-link')
    .first()
    .evaluate((el) => el.getBoundingClientRect().height);
  expect(height).toBeLessThanOrEqual(48);
}

async function pictureShowsPhoneCrop(page: Page, phone: boolean): Promise<void> {
  const frame: Locator = page.getByTestId('setup-companion-capture');
  await frame.scrollIntoViewIfNeeded();
  const img = frame.locator('img');
  await expect
    .poll(() => img.evaluate((el: HTMLImageElement) => el.complete && el.naturalWidth > 0))
    .toBe(true);
  const info = await img.evaluate((el: HTMLImageElement) => ({
    src: el.currentSrc,
    box: el.getBoundingClientRect().toJSON(),
  }));
  expect(info.src.includes('companion-status-phone')).toBe(phone);
  expect(info.box.height).toBeGreaterThan(60);
  if (phone) {
    // 100% of the panel's content width, about 314 px at 390.
    expect(info.box.width).toBeGreaterThan(300);
    expect(info.box.width / info.box.height).toBeCloseTo(734 / 172, 0);
    // The sub line is 26 px tall in the 734 px crop, so at this width it is at least 11 px.
    expect((26 * info.box.width) / 734).toBeGreaterThanOrEqual(11);
  }
}

for (const size of SIZES) {
  const phone = size.width < 640;
  test.describe(`at ${size.name}`, () => {
    test.use({ viewport: { width: size.width, height: size.height } });

    test('/setup signed out: gold on Sign in only, the chips on one row, the picture shown', async ({
      page,
    }) => {
      await signedOut(page);
      await page.goto('/setup');
      await settle(page);
      expect(await goldSteps(page)).toEqual(['1']);
      await chipRowIsOneLine(page);
      await expect(page.getByTestId('setup-how-it-arrives')).toHaveText(
        'The addon saves your character when you log out; the companion sends it to the site. No companion? Paste an export instead (see Try an export below).',
      );
      await page.getByTestId('setup-how-it-arrives').getByRole('link', { name: 'Try an export' }).click();
      await expect(page.getByTestId('setup-try-export')).toHaveJSProperty('open', true);
      await expect(page.getByTestId('setup-try-export')).toBeInViewport();
      await expect(page.getByTestId('setup-companion-note')).toHaveText(
        'Sends your character automatically and logs live. Without it, paste an export after each change.',
      );
      await pictureShowsPhoneCrop(page, phone);
      await expect(page.getByTestId('setup-companion-capture-label')).toHaveText('Example');
      await expect(page.getByTestId('setup-companion-capture-label')).toBeVisible();
      await settle(page);
      await shot(page, `setup-signed-out-${size.name}`);
    });

    test('/setup signed in, addon and companion pending: gold on the addon step only', async ({ page }) => {
      await signedIn(page, ME_NEW, []);
      await page.goto('/setup');
      await settle(page);
      await expect(page.getByTestId('setup-line-addon')).toHaveText(
        'No character yet. Log out once with the addon on and the companion running, or paste an export.',
      );
      expect(await goldSteps(page)).toEqual(['2']);
      await chipRowIsOneLine(page);
      await pictureShowsPhoneCrop(page, phone);
      await expect(page.getByTestId('setup-companion-capture-label')).toBeHidden();
      await settle(page);
      await shot(page, `setup-signed-in-pending-${size.name}`);
    });

    test('/setup signed in, everything done: no gold, the chips still one row with one label each', async ({
      page,
    }) => {
      await signedIn(page, ME_DONE, [DEVICE]);
      await page.goto('/setup');
      await settle(page);
      await expect(page.locator('.setup-chip[data-done="true"]')).toHaveCount(3);
      expect(await goldSteps(page)).toEqual([]);
      await chipRowIsOneLine(page);
      await pictureShowsPhoneCrop(page, phone);
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

    test('the landing after the beta: the example label, "levels", the next date in gold on Name reservation, full page', async ({
      page,
    }) => {
      await page.clock.install({ time: new Date('2026-10-23T12:00:00Z') });
      await signedOut(page);
      await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
      await page.route('**/v1/rankings/guilds?**', (route) => route.fulfill(fulfil({ rows: [] })));
      await routeSampleReportMeta(page);
      await page.goto('/');
      await page.clock.resume();
      await settle(page);
      await expect(page.getByTestId('home-hero-label')).toHaveCount(0);
      await expect(page.getByTestId('home-hero-link')).toContainText('Marksmanship Hunter best in slot');
      await expect(page.getByTestId('home-hero-bands')).toHaveText(
        'Leveling best in slot, levels 20, 30, 40, 50, 60',
      );
      await expect(page.getByTestId('home-dates-line').locator('[data-next]')).toHaveText(
        'Name reservation Oct 27',
      );
      await shot(page, `landing-after-beta-full-${size.name}`);
    });

    test('the landing while the beta runs: the dates line leads with the running beta', async ({ page }) => {
      await page.clock.install({ time: new Date('2026-10-20T12:00:00Z') });
      await signedOut(page);
      await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
      await page.route('**/v1/rankings/guilds?**', (route) => route.fulfill(fulfil({ rows: [] })));
      await routeSampleReportMeta(page);
      await page.goto('/');
      await page.clock.resume();
      await settle(page);
      await expect(page.getByTestId('home-hero-card')).toContainText('Beta is live');
      await expect(page.getByTestId('home-hero-bands')).toHaveText('Leveling best in slot, levels 10 to 30');
      await expect(page.getByTestId('home-dates-line').locator('[data-next]')).toHaveText(
        'Beta live through Oct 21',
      );
      await shot(page, `landing-beta-live-full-${size.name}`);
    });

    test('a character the visitor does not hold, with no data: one sentence, the lookup, one quiet link', async ({
      page,
    }) => {
      await signedOut(page);
      await routeEmptyCharacter(page);
      await page.goto(CHARACTER_PATH);
      await settle(page);
      await expect(page.getByTestId('character-not-found-line')).toHaveText(
        'Elyra Duskvale has no Forever Sixty data yet.',
      );
      await expect(page.getByTestId('character-not-found-form')).toBeVisible();
      await expect(page.getByTestId('character-not-found-next').getByRole('link')).toHaveText(
        'Get the addon',
      );
      await expect(page.getByTestId('character-not-found-next').getByRole('link')).toHaveCount(1);
      // The empty score bars and the tooltip lookalike are the holder's, not a visitor's.
      await expect(page.getByTestId('forever-score-empty')).toHaveCount(0);
      await expect(page.getByTestId('forever-score-empty-action')).toHaveCount(0);
      await expect(page.getByTestId('character-gear-none')).toHaveCount(0);
      await expect(page.getByTestId('character-in-game')).toHaveCount(0);
      await expect(page.getByText('Forever Score: no logs yet')).toHaveCount(0);
      await shot(page, `character-no-data-visitor-${size.name}`);
    });

    test('a character another account holds gets the same sentence, not the holder’s set', async ({
      page,
    }) => {
      await signedIn(page, ME_DONE, []);
      await routeEmptyCharacter(page);
      await page.goto(CHARACTER_PATH);
      await settle(page);
      await expect(page.getByTestId('character-not-found-line')).toHaveText(
        'Elyra Duskvale has no Forever Sixty data yet.',
      );
      await expect(page.getByTestId('forever-score-empty')).toHaveCount(0);
    });

    test('the signed-in holder keeps the full set: one primary button, one quiet line, no sign-in bar', async ({
      page,
    }) => {
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
      await expect(page.getByTestId('forever-score-empty-action')).toHaveText('Upload logs');
      await expect(page.getByTestId('forever-score-empty-next')).toHaveCount(0);
      await expect(page.getByTestId('character-gear-none')).toContainText('No gear on file yet.');
      await expect(page.getByTestId('character-gear-planner')).toHaveCount(0);
      await expect(page.getByTestId('character-ruleset-chip')).toHaveText(/hardcore/i);
      await shot(page, `character-no-data-holder-${size.name}`);
    });

    test('a character with gear and no logs: "No logs yet" in place of the score card, one button', async ({
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
      await expect(page.getByTestId('forever-score-empty')).toContainText('No logs yet');
      await expect(page.getByTestId('forever-score-empty-max')).toHaveText('Score runs to 2100.');
      await expect(page.getByTestId('forever-score-empty-gear')).toHaveCount(0);
      await expect(page.getByTestId('forever-score-empty-action')).toHaveText('Get the companion');
      await expect(page.getByTestId('current-character-bar-signed-out')).toHaveCount(0);
      await shot(page, `character-gear-first-${size.name}`);
    });
  });
}
