// web/tests/e2e/report-density-captures.spec.ts
// The logs lane's density-pass captures: /logs, /reports and every report tab on the bundled
// sample (pull 11, Saltspine; Deaths on pull 20) at 1440 and 390, full page, written to
// design/reviews/2026-11-09-p4-logs-density/<state>-<width>.png. Skipped unless
// DENSITY_CAPTURES=1: they are pictures for the reviewers, not assertions.
//   DENSITY_CAPTURES=1 npx playwright test report-density-captures.spec.ts --project=desktop --workers=1
import { join } from 'node:path';
import { expect, test } from './support/test';
import { openLogsState, type LogsParityState } from './support/logs-parity';
import { signedOut } from './support/selector';

const OUT = join(import.meta.dirname, '../../../design/reviews/2026-11-09-p4-logs-density');
const WIDTHS = [1440, 390] as const;
const PLAYER = 'Player-4184-0000000E';

const PAGES: readonly LogsParityState[] = ['logs', 'logs-you', 'reports-list'];

/** Each report capture: its name and the sample path it opens. */
const TABS: readonly (readonly [string, string])[] = [
  ['summary', '?tab=summary'],
  ['damage-done', '?tab=damage-done'],
  ['healing', '?tab=healing'],
  ['damage-taken', '?tab=damage-taken'],
  ['deaths', '?fight=20&tab=deaths'],
  ['mechanics', '?tab=mechanics'],
  ['timeline', '?tab=timeline'],
  ['threat', '?tab=threat'],
  ['casts', '?tab=casts'],
  ['buffs', '?tab=buffs'],
  ['debuffs', '?tab=debuffs'],
  ['resources', '?tab=resources'],
  ['compare', '?tab=compare'],
  ['rating', '?tab=rating'],
  ['coach', '?tab=coach'],
  ['player', `?tab=damage-done&source=${PLAYER}`],
];

test.skip(process.env.DENSITY_CAPTURES !== '1', 'review captures run on request');
test.skip(({ isMobile }) => isMobile, 'each width sets its own viewport');

for (const width of WIDTHS) {
  const phone = width < 768;
  test.describe(`${width}`, () => {
    test.use({ viewport: { width, height: phone ? 844 : 1000 }, timezoneId: 'UTC' });

    for (const state of PAGES) {
      test(`${state} ${width}`, async ({ page }) => {
        await openLogsState(page, state);
        await page.waitForTimeout(600);
        await page.screenshot({ path: join(OUT, `${state}-${width}.png`), fullPage: true });
      });
    }

    for (const [name, query] of TABS) {
      test(`report ${name} ${width}`, async ({ page }) => {
        await signedOut(page);
        await page.goto(`/reports/sample${query}`);
        await expect(page.getByTestId('report-page')).toBeVisible();
        await expect(page.getByTestId('report-skeleton')).toHaveCount(0);
        await page.waitForTimeout(1200);
        await page.screenshot({ path: join(OUT, `report-${name}-${width}.png`), fullPage: true });
      });
    }
  });
}
