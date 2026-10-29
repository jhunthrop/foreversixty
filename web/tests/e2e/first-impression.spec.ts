// web/tests/e2e/first-impression.spec.ts
// W.8: the pages a visitor meets before they are a user (/, /setup, /logs), signed out, at 1280
// and 390. Each reads right and is captured for the review (W8_SHOTS names the folder; by default
// test-results/w8).
import { mkdirSync } from 'node:fs';
import { join } from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';
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

    test('the landing page shows the next date first, no stale stamp, no "Loading"', async ({ page }) => {
      mkdirSync(SHOTS, { recursive: true });
      await signedOutWithFeeds(page);
      await page.goto('/');
      const timeline = page.getByTestId('home-timeline');
      await expect(timeline).toBeVisible();
      await expect(timeline).not.toContainText('Updated');
      const rows = timeline.getByTestId('home-timeline-row');
      await expect(rows.first()).not.toHaveAttribute('data-status', 'past');
      const statuses = await rows.evaluateAll((els) => els.map((el) => (el as HTMLElement).dataset.status));
      // Past dates come last: no past row is followed by a live, next or future one.
      expect(statuses.join(' ')).not.toMatch(/past (live|next|future)/);
      await expect(page.getByRole('heading', { name: 'Get started' })).toBeVisible();
      await expect(page.getByText('Five doors')).toHaveCount(0);
      await expect(page.getByTestId('home-get-set-up-signed-out')).toContainText(
        'Sign in, install the addon, and run the companion — one page.',
      );
      await page.getByTestId('home-around-the-site').scrollIntoViewIfNeeded();
      await expect(page.getByTestId('home-around-the-site')).toBeVisible();
      expect(await page.locator('main').innerText()).not.toContain('Loading');
      await page.screenshot({ path: join(SHOTS, `w8-landing-${size.name}.png`), fullPage: true });
    });

    test('/setup is three steps, the manual folded away, the companion after it', async ({ page }) => {
      mkdirSync(SHOTS, { recursive: true });
      await signedOutWithFeeds(page);
      await page.goto('/setup');
      await expect(page.getByTestId('setup-status')).toHaveAttribute('aria-label', 'Your progress');
      const manual = page.getByTestId('setup-everything');
      await expect(manual.locator('summary')).toContainText('Everything the addon does');
      await expect(manual).not.toHaveAttribute('open', '');
      await expect(page.getByTestId('setup-commands')).toBeHidden();
      await expect(page.getByText('For testers')).toHaveCount(0);
      await expect(page.getByTestId('setup-companion-pair')).toHaveText(
        'Sign in with Battle.net inside the companion; it connects itself.',
      );
      await expect(page.getByTestId('setup-companion-platforms')).toHaveText(
        'Windows only; Mac is not supported yet.',
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
      await expect(named.getByRole('link', { name: 'Bow Jackzon' })).toBeVisible();
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
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(fulfil(null, 404)));
  // Only the fixture character's address is prerendered in the preview build; the sentence's own
  // name, ruleset and region are pinned in CharacterNotFound.test.ts ("Nobody Here" on "PvP US").
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-not-found-line')).toHaveText(
    'No character called Elyra Duskvale on Hardcore US yet.',
  );
  await expect(page.getByTestId('character-not-found-form')).toBeVisible();
  await expect(page.getByTestId('character-not-found-next')).toContainText('Your character? Get the addon');
  await expect(page).toHaveTitle('Character not found');
  await page.screenshot({ path: join(SHOTS, 'w8-character-not-found.png'), fullPage: true });
});

test('the planner opens a visitor on the level-60 build, never "your level’s range"', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/planner');
  const card = page.getByTestId('band-compare');
  await expect(card).toContainText('60 build', { timeout: 15_000 });
  await expect(card).not.toContainText('your level');
  await expect(card).toContainText('the 60 level range');
});
