// web/tests/captures/batch-h/capture.spec.ts
// The batch H integration captures (design/reviews/2026-11-10-batch-h): the merged build's sim
// result, a report's Player view, a guide spec page, the guild Raids and Applications tabs,
// /ratings and the shared build /b/{id}, at 1440 and 390, opened through the parity gate's own
// state openers. The /b page is the API's fixture HTML served as it is, with no chrome swap, so
// its header loads its scripts by the stable /_chrome/ paths (web/scripts/chrome-aliases.mjs).
// Not a test of behaviour.
//   CAPTURE_PORT=<port> npx playwright test -c tests/captures/batch-h/playwright.capture.config.ts
import { mkdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { prepareParityPage, REPO_ROOT, settleParityPage } from '../../e2e/support/parity';
import { openSimState } from '../../e2e/support/sim-parity';
import { openGuidesState } from '../../e2e/support/guides-page';
import { openGuildState } from '../../e2e/support/guild-parity';
import { routePlannerState } from '../../e2e/support/planner-page';

const OUT = process.env.CAPTURE_OUT ?? path.join(REPO_ROOT, 'design/reviews/2026-11-10-batch-h');
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const SHARED = path.join(REPO_ROOT, 'web/tests/e2e/fixtures/shared-build/survival.html');
const SHARED_RECORD = JSON.parse(
  readFileSync(path.join(REPO_ROOT, 'api/internal/site/testdata/shared-survival.json'), 'utf8'),
) as { id: string };

async function openReportPlayer(page: Page): Promise<void> {
  await page.goto('/reports/fixture2abcd?fight=3');
  await page.getByTestId('at-pull-links').first().getByRole('link', { name: 'Gear ›' }).click();
  await expect(page.getByTestId('report-gear')).toBeVisible();
}

async function openShared(page: Page): Promise<void> {
  await routePlannerState(page, 'out');
  const body = readFileSync(SHARED, 'utf8');
  await page.route(`**/b/${SHARED_RECORD.id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body }),
  );
  await page.goto(`/b/${SHARED_RECORD.id}`);
  // The header's selector island hydrates from /_chrome/ (its ssr attribute goes once it has).
  await expect(page.locator('astro-island[component-url^="/_chrome/"]').first()).not.toHaveAttribute(
    'ssr',
    /.*/,
  );
}

const PAGES: readonly [name: string, open: (page: Page) => Promise<void>][] = [
  ['sim-result', (page) => openSimState(page, 'result')],
  ['report-player', openReportPlayer],
  ['guide-mage-frost', (page) => openGuidesState(page, '/guides/mage/frost', 'out')],
  ['guild-raids', (page) => openGuildState(page, 'dense-raids')],
  ['guild-applications', (page) => openGuildState(page, 'dense-applications')],
  ['ratings', async (page) => void (await page.goto('/ratings'))],
  ['shared-build', openShared],
];

mkdirSync(OUT, { recursive: true });

for (const [name, open] of PAGES) {
  for (const width of WIDTHS) {
    test(`${name} at ${width}`, async ({ page }) => {
      test.slow();
      await page.setViewportSize({ width, height: HEIGHT });
      await prepareParityPage(page);
      await open(page);
      await settleParityPage(page);
      await page.screenshot({ path: path.join(OUT, `${name}-${width}.png`), fullPage: true });
    });
  }
}
