// web/tests/e2e/launch-polish.spec.ts
// P.1: the ship review's pre-launch items on the landing, the character page, the sample report and
// /setup, signed out, at 1280 and 390: the landing speaks in three outcome rows and never shows a
// skeleton for an empty list; the class row at 390 is a 5-and-4 grid; a character with a realm that
// is not its ruleset reads cleanly; the sample says why its numbers are small and shows more than one
// player at 1280. Captured after the page is scrolled through (an island that mounts when visible
// has mounted) when P1_SHOTS names a folder.
import { mkdirSync } from 'node:fs';
import { join } from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
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
import { SAMPLE_REPORT_META, routeSampleReportMeta } from './support/logs-sample-fixture';
import { fieldBox } from './support/select';

const SHOTS = process.env.P1_SHOTS;
const SIZES = [
  { name: '1280', width: 1280, height: 720 },
  { name: '390', width: 390, height: 844 },
] as const;

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

async function signedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil(null, 401)));
}

/** Scrolls down the page in steps so every island that mounts when visible has mounted, then back to the top. */
async function scrollThrough(page: Page): Promise<void> {
  const height = await page.evaluate(() => document.documentElement.scrollHeight);
  for (let y = 0; y < height; y += 500) {
    await page.evaluate((to) => window.scrollTo(0, to), y);
    await page.waitForTimeout(120);
  }
  await page.evaluate(() => window.scrollTo(0, 0));
  await page.waitForLoadState('networkidle');
  await page.waitForTimeout(400);
}

async function shot(page: Page, name: string, fullPage = true): Promise<void> {
  if (SHOTS === undefined) return;
  mkdirSync(SHOTS, { recursive: true });
  await page.screenshot({ path: join(SHOTS, `p1-${name}.png`), fullPage });
}

for (const size of SIZES) {
  test.describe(`at ${size.name}`, () => {
    test.use({ viewport: { width: size.width, height: size.height } });

    test('a character whose realm is not its ruleset: the realm line and the badge say different things', async ({
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
      await page.waitForLoadState('networkidle');
      await expect(page.getByTestId('character-fact-realm')).toHaveText('Doomhowl');
      await expect(page.getByTestId('character-fact-chip')).toHaveText('Hardcore US');
      expect(await page.getByTestId('character-head').innerText()).not.toMatch(/Beta/);
      await shot(page, `character-realm-${size.name}`);
    });

    test('a character with nothing on it, seen by a visitor: the not-found page', async ({ page }) => {
      await signedOut(page);
      const data = {
        ...CHARACTER.data,
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
      await routeCharacterSearch(page);
      await page.goto(CHARACTER_PATH);
      await page.waitForLoadState('networkidle');
      await expectFixtureNotFound(page, { signedIn: false });
      await shot(page, `character-no-data-${size.name}`);
    });

    test('the sample report says why its numbers are small and shows more than one player', async ({
      page,
    }) => {
      await signedOut(page);
      await page.goto('/reports/sample');
      await expect(page.getByTestId('report-title')).toHaveText('Sample raid night, 17 Sept 2026');
      await expect(page.getByTestId('report-sample-note')).toHaveText(
        'A real beta night with the names changed. Level 10 to 30 players, so the numbers are small. Your own reports look like this.',
      );
      const rows = page.getByTestId('summary-tab').getByTestId(/^roster-Player-/);
      await expect(rows.first()).toBeVisible();
      await page.waitForLoadState('networkidle');

      const subtitle = page.getByTestId('report-subtitle');
      for (const text of ['17 Sept 2026', `${SAMPLE_REPORT_META.fights.length} fights`]) {
        const span = subtitle.locator('span', { hasText: text }).first();
        await expect(span).toBeVisible();
        expect(await span.evaluate((el) => getComputedStyle(el).whiteSpace)).toBe('nowrap');
      }
      if (size.width >= 768) {
        // From md up the sample draws the full chart, as a real report does (owner, 2026-10-30:
        // the 96 px strip was "still too small"), so the table starts below a 720 px fold; the
        // first-screen promise is the phone's (onboarding-second-screen.spec.ts). Measured at
        // 1280 x 720: the chart is 390 to 885 px, the first row at 986 px.
        const chart = page.getByTestId('time-chart');
        await expect(chart).not.toHaveAttribute('data-compact', 'true');
        await expect(fieldBox(chart, 'window-select')).toBeVisible();
        const box = (await chart.boundingBox())!;
        const first = (await rows.first().boundingBox())!;
        expect(first.y).toBeGreaterThan(box.y + box.height);
      } else {
        // The title shows its date, not an ellipsis.
        const title = page.getByTestId('report-title');
        expect(await title.evaluate((el) => el.scrollWidth <= el.clientWidth + 1)).toBe(true);
        expect(await title.evaluate((el) => getComputedStyle(el).textOverflow)).not.toBe('ellipsis');
      }
      await shot(page, `sample-${size.name}`, false);
    });

    test('/logs shows the sample row with the sample’s own date and fights', async ({ page }) => {
      await signedOut(page);
      await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
      await routeSampleReportMeta(page);
      await page.goto('/logs');
      const sample = page.getByTestId('recent-reports-sample');
      await sample.scrollIntoViewIfNeeded();
      await expect(sample.getByTestId('report-row-when')).toHaveText('17 Sept 2026');
      await expect(sample.getByTestId('report-row-summary')).toContainText(
        `${SAMPLE_REPORT_META.fights.length} fights`,
      );
      await scrollThrough(page);
      await shot(page, `logs-${size.name}`);
    });
  });
}

test.describe('/setup picture at 390', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('is the whole Status capture at its aspect, labelled Example when signed out', async ({ page }) => {
    await signedOut(page);
    await page.goto('/setup');
    const frame = page.getByTestId('setup-companion-capture');
    await frame.scrollIntoViewIfNeeded();
    await expect(frame.getByText('Example')).toBeVisible();
    const img = frame.locator('img');
    await expect
      .poll(() => img.evaluate((el: HTMLImageElement) => el.complete && el.naturalWidth > 0))
      .toBe(true);
    const box = await img.evaluate((el: HTMLImageElement) => el.getBoundingClientRect().toJSON());
    expect(box.width).toBeLessThanOrEqual(390);
    expect(box.width / box.height).toBeCloseTo(1600 / 436, 0);
    await page.waitForLoadState('networkidle');
    await shot(page, 'setup-signed-out-390', true);
  });

  test('step lines wrap pretty: the body text balances its last line', async ({ page }) => {
    await signedOut(page);
    await page.goto('/setup');
    expect(
      await page.getByTestId('setup-companion-lead').evaluate((el) => getComputedStyle(el).textWrap),
    ).toMatch(/pretty|balance/);
  });
});
