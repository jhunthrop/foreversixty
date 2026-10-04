// web/tests/e2e/logs-widths.spec.ts
// Logs landing spec (2026-10-04) §4.A's header band breaks out full-bleed the same way
// PlannerHeaderBand.svelte already does (`width: 100vw; margin-left: calc(50% - 50vw)`),
// which planner-widths.spec.ts's own history shows is exactly the kind of breakout that can
// push a page wider than its own viewport at a laptop width. Asserts the one invariant that
// matters: /logs never scrolls sideways, at every width this pass was asked to check, signed
// out and signed in.
import { expect, test } from '@playwright/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { routeSampleReportMeta } from './support/logs-hero-fixture';

const WIDTHS = [1024, 1100, 1280, 1440, 1920];

function fulfil(body: unknown, status = 200): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'e2e' }),
  };
}

async function noHorizontalScroll(page: import('@playwright/test').Page): Promise<void> {
  const { scrollWidth, clientWidth } = await page.evaluate(() => ({
    scrollWidth: document.documentElement.scrollWidth,
    clientWidth: document.documentElement.clientWidth,
  }));
  expect(scrollWidth, `scrollWidth (${scrollWidth}) should equal clientWidth (${clientWidth})`).toBe(
    clientWidth,
  );
}

test.describe('the logs landing never scrolls sideways', () => {
  for (const width of WIDTHS) {
    test(`signed out at ${width}px`, async ({ page }) => {
      await page.setViewportSize({ width, height: 1000 });
      await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
      await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
      await routeSampleReportMeta(page);
      await page.goto('/logs');
      await page.getByTestId('logs-hero-description').waitFor();
      await noHorizontalScroll(page);
    });

    test(`signed in at ${width}px`, async ({ page }) => {
      await page.setViewportSize({ width, height: 1000 });
      await page.route('**/v1/me', (route) => route.fulfill(fulfil(meAddonFixture)));
      await page.route('**/v1/devices', (route) => route.fulfill(fulfil([])));
      await page.route('**/v1/reports?mine=1**', (route) =>
        route.fulfill(
          fulfil({
            rows: [
              {
                id: 'ownreport01x',
                title: 'Molten Core, week 3',
                zone: 'Molten Core',
                status: 'complete',
                visibility: 'public',
                created_at: '2026-10-03T22:10:00Z',
                fight_count: 9,
                kill_count: 4,
              },
            ],
            total: 1,
            page: 1,
            per_page: 100,
          }),
        ),
      );
      await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
      // The hero's own per-player hook (§4.A.1) is a deferred third fetch; left unmocked so
      // it fails quietly (its own spec'd behaviour) rather than widening the fixture set this
      // width check does not need.
      await page.route('**/v1/reports/ownreport01x', (route) => route.abort());
      await page.goto('/logs');
      await page.getByTestId('logs-hero-title').waitFor();
      await noHorizontalScroll(page);
    });
  }
});
