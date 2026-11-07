// web/tests/e2e/first-impression.spec.ts
// W.8: the pages a visitor meets before they are a user (/, /setup, /logs), signed out, at 1280
// and 390. Each reads right and is captured for the review (W8_SHOTS names the folder; by default
// test-results/w8).
import { mkdirSync } from 'node:fs';
import { join } from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { CHARACTER_PATH, CHARACTER_URL, routeCharacterSearch } from './support/character-fixture';
import { expectFixtureNotFound } from './support/character-not-found';
import { routeSampleReportMeta } from './support/logs-sample-fixture';

const SHOTS = process.env.W8_SHOTS ?? 'test-results/w8';
const SIZES = [
  { name: '1280', width: 1280, height: 900 },
  { name: '390', width: 390, height: 844 },
] as const;

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

async function signedOutWithFeeds(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil(null, 401)));
  await page.route('**/v1/rankings/guilds?**', (route) => route.fulfill(fulfil({ rows: [] })));
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: 'bare00000001',
            title: 'Report',
            created_at: new Date().toISOString(),
            fight_count: 0,
            kill_count: 0,
          },
          {
            id: 'named0000001',
            title: 'Blackrock Depths',
            created_at: new Date().toISOString(),
            fight_count: 14,
            kill_count: 3,
            zone: 'Blackrock Depths',
            logging_character: { key: 'us/pvp/bow-jackzon', name: 'Bow Jackzon' },
          },
        ],
      }),
    ),
  );
  await routeSampleReportMeta(page);
}

for (const size of SIZES) {
  test.describe(`signed out at ${size.name}`, () => {
    test.use({ viewport: { width: size.width, height: size.height } });

    test('/setup is three steps in one plate, the companion section after it', async ({ page }) => {
      mkdirSync(SHOTS, { recursive: true });
      await signedOutWithFeeds(page);
      await page.goto('/setup');
      const plate = page.getByRole('region', { name: 'Your setup' });
      await expect(plate.getByRole('img', { name: /^Step \d: / })).toHaveCount(3);
      await expect(plate.getByRole('link', { name: 'Sign in with Battle.net' })).toBeVisible();
      await expect(page.getByTestId('setup-commands')).toBeVisible();
      await expect(page.getByText('For testers')).toHaveCount(0);
      await expect(page.getByTestId('setup-companion-pair')).toContainText(
        'Sign in with Battle.net inside the companion; it connects itself.',
      );
      await expect(page.getByTestId('setup-companion-platforms')).toHaveText(
        'Windows only; Mac is not supported yet. On a Mac, paste an export instead.',
      );
      await page.getByTestId('setup-companion-capture').scrollIntoViewIfNeeded();
      await expect(page.getByTestId('setup-companion-capture').locator('img')).toBeVisible();
      await page.screenshot({ path: join(SHOTS, `w8-setup-${size.name}.png`), fullPage: true });
    });

    test('/logs lists only reports with a zone or a kill, in the "Your reports" row shape', async ({
      page,
    }) => {
      mkdirSync(SHOTS, { recursive: true });
      await signedOutWithFeeds(page);
      await page.goto('/logs');
      const feed = page.getByTestId('recent-reports');
      // The named report and the sample row; the bare "0 kills" report is not listed.
      await expect(feed.getByTestId('report-row')).toHaveCount(2);
      const named = feed.getByTestId('report-row').first();
      await expect(named).toContainText('Blackrock Depths');
      // The credited character is the title; the zone leads the muted line under it (W.9).
      await expect(named).toHaveAttribute('href', /^\/reports\//);
      await expect(named).toContainText('Bow Jackzon');
      await expect(named.getByTestId('report-row-summary')).toHaveText(
        'Blackrock Depths · 3 bosses · 14 fights',
      );
      await expect(page.getByTestId('reports-sign-in')).toBeVisible();
      await feed.scrollIntoViewIfNeeded();
      await page.screenshot({ path: join(SHOTS, `w8-logs-${size.name}.png`), fullPage: true });
    });
  });
}

test('an unknown character says so, with the lookup and the addon line', async ({ page }) => {
  mkdirSync(SHOTS, { recursive: true });
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.route(CHARACTER_URL, (route) => route.fulfill(fulfil(null, 404)));
  await routeCharacterSearch(page);
  // Only the fixture character's address is prerendered in the preview build; the line's own
  // name, region and ruleset are pinned in page-view's unit test.
  await page.goto(CHARACTER_PATH);
  await expectFixtureNotFound(page, { signedIn: false });
  await expect(page.getByTestId('character-elsewhere')).toHaveCount(0);
  await page.screenshot({ path: join(SHOTS, 'w8-character-not-found.png'), fullPage: true });
  // C10: Look up takes the name to the lookup page.
  await page.getByTestId('character-lookup').getByRole('button', { name: 'Look up' }).click();
  await expect(page).toHaveURL(/\/characters\?q=Elyra%20Duskvale#lookup$/);
});

test('the planner opens a visitor on the level-60 build, never a "range" line', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/planner');
  const card = page.getByTestId('band-compare');
  await expect(card).toContainText(/level 60 build/i, { timeout: 15_000 });
  await expect(card).not.toContainText('your level');
  await expect(card).not.toContainText('range');
});
