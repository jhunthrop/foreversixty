// web/tests/e2e/sample-report.spec.ts
// The sample report is bundled with the site (lib/report/sample.ts): /logs offers it, the
// click lands on fights, and no API is involved at any step.
import { expect, test, type Page } from '@playwright/test';
import { SAMPLE_REPORT_META } from './support/logs-sample-fixture';
import { fieldBox } from './support/select';

const SHOT_DIR = process.env.SAMPLE_SHOT_DIR;

/**
 * From the next navigation on, sums the layout shifts of anything inside the report. The site
 * header's own font-swap nudge (a 3 px text move in the nav, before the island mounts) is not
 * the report's and is left to the header's specs.
 */
async function recordReportShifts(page: Page): Promise<void> {
  await page.addInitScript(() => {
    const w = window as unknown as { __reportShift: number };
    w.__reportShift = 0;
    new PerformanceObserver((list) => {
      for (const entry of list.getEntries() as (PerformanceEntry & {
        value: number;
        hadRecentInput: boolean;
        sources?: { node?: Node | null }[];
      })[]) {
        const inReport = (entry.sources ?? []).some(
          (source) => source.node?.parentElement?.closest('#report') != null,
        );
        if (!entry.hadRecentInput && inReport) w.__reportShift += entry.value;
      }
    }).observe({ type: 'layout-shift', buffered: true });
  });
}

const reportShift = (page: Page): Promise<number> =>
  page.evaluate(() => (window as unknown as { __reportShift: number }).__reportShift);

test('the sample report opens from /logs and shows fights, with no API', async ({ page }, info) => {
  const api: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/v1/reports/sample')) api.push(request.url());
  });
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'no' }, request_id: 'r' }),
    }),
  );
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { rows: [] }, error: null, request_id: 'r' }),
    }),
  );

  await page.goto('/logs');
  const sample = page.getByTestId('recent-reports-sample');
  await sample.scrollIntoViewIfNeeded();
  await expect(sample.getByTestId('report-row-summary')).toHaveText(
    `${SAMPLE_REPORT_META.zone} · ${SAMPLE_REPORT_META.fights.length} fights`,
  );
  await expect(sample.getByTestId('report-row-when')).toHaveText('17 Sept 2026');
  await sample.getByRole('link').first().click();

  await expect(page).toHaveURL(/\/reports\/sample$/);
  await expect(page.getByTestId('report-title')).toHaveText('Beta night, 17 Sept 2026');
  await expect(page.getByText('17 Sept 2026').first()).toBeVisible();
  await expect(page.getByTestId('report-sample-badge')).toHaveText('Sample report');
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  await expect(page.getByTestId('report-owner')).toHaveCount(0);
  await expect(page.getByTestId('report-credit')).toHaveCount(0);
  await expect(page.getByTestId('report-error')).toHaveCount(0);
  await expect(page.getByText(SAMPLE_REPORT_META.zone).first()).toBeVisible();
  expect(api).toEqual([]);

  // L.8: a first look that convinces. Nothing the log cannot know is drawn; the class shows.
  const roster = page.getByTestId('summary-tab');
  await expect(roster).toBeVisible();
  await expect(roster.getByTestId(/^roster-Player-/).first()).toBeVisible();
  const body = (await page.locator('main').innerText()) ?? '';
  expect(body).not.toContain('REPLAY');
  expect(body).not.toContain('could not be reached');
  await expect(page.getByTestId('mode-replay')).toHaveCount(0);
  await expect(page.getByTestId('parse-unavailable')).toHaveCount(0);
  await expect(page.getByTestId('roster-percentile')).toHaveCount(0);
  await expect(page.getByTestId('roster-spec')).toHaveCount(0);
  await expect(page.getByTestId('report-sample-note')).toContainText(
    'A real beta night with the names changed. Level 10 to 30 players, so the numbers are small. Your own reports look like this.',
  );
  await expect(page).not.toHaveURL(/fight=/);
  // Opens on its first boss kill, else its longest boss attempt, else its longest pull. The
  // sample has no boss, so that is its longest pull (fight 6, 12:43), not the first.
  await expect(page.getByTestId('fight-6')).toHaveAttribute('aria-current', 'true');
  // Class colours: no name keeps the plain text colour.
  const colours = await roster
    .getByTestId(/^roster-Player-/)
    .evaluateAll((rows) =>
      rows.slice(0, 40).map((row) => getComputedStyle(row.querySelector('span') as Element).color),
    );
  expect(new Set(colours).size).toBeGreaterThan(2);
  // A trash row prints its length once.
  const row1 = await page.getByTestId('fight-1').innerText();
  expect((row1.match(/3:26/g) ?? []).length).toBe(1);

  if (SHOT_DIR !== undefined) {
    const width = info.project.name === 'mobile' ? 390 : 1280;
    await page.screenshot({ path: `${SHOT_DIR}/l8-sample-${width}.png`, fullPage: false });
  }
});

test.describe('the sample report’s chart', () => {
  test('on a desktop is the full chart, as a real report draws it', async ({ page, isMobile }) => {
    test.skip(isMobile, 'the desktop layout');
    await recordReportShifts(page);
    for (const width of [1280, 1440, 2000]) {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/reports/sample');
      const chart = page.getByTestId('time-chart');
      await expect(chart.locator('[data-state="ready"]').first()).toBeVisible();
      await expect(chart).not.toHaveAttribute('data-compact', 'true');
      // The head with its title and units, the legend, the Window presets and the brush.
      await expect(chart.getByText('Damage done', { exact: true }).first()).toBeVisible();
      await expect(chart.getByTestId('chart-head-meta')).toContainText('DPS');
      await expect(chart.getByTestId('chart-legend')).toBeVisible();
      await expect(fieldBox(chart, 'window-select')).toBeVisible();
      await expect(chart.getByTestId('overview-brush')).toBeVisible();
      // A plot of about 200 px (220 from 1440), plus its 22 px axis.
      const plot = (await chart.getByTestId('chart-plot').boundingBox())!;
      expect(plot.height).toBeGreaterThanOrEqual(200 + 22);
      // The full chart's frame reserves its loaded height: nothing moves as it arrives.
      expect(await reportShift(page)).toBe(0);
    }
  });

  test('on a phone stays the compact strip, once opened', async ({ page, isMobile }) => {
    test.skip(!isMobile, 'the phone layout');
    await page.goto('/reports/sample');
    await page.getByTestId('chart-toggle').click();
    const chart = page.getByTestId('time-chart');
    await expect(chart).toHaveAttribute('data-compact', 'true');
    await expect(chart.getByTestId('chart-legend')).toHaveCount(0);
    const box = (await chart.boundingBox())!;
    expect(box.height).toBeLessThanOrEqual(100);
  });
});

test.describe('the brush on a fight that does not end on a whole second', () => {
  // 1 s, 2 s and 5 s points (lib/chart/buckets.ts); the 10 s case is OverviewBrush.test.ts's.
  const FIGHTS = [
    { fight: 2, durationMs: 9_879, label: '0:09.9' },
    { fight: 9, durationMs: 80_261, label: '1:20' },
    { fight: 6, durationMs: 763_140, label: '12:43' },
  ];
  for (const { fight, durationMs, label } of FIGHTS) {
    test(`fight ${fight}: the whole-fight end handle sits on the track's end`, async ({ page }) => {
      // 1024 px: the full chart and its brush on both projects (a phone draws the strip, no brush).
      await page.setViewportSize({ width: 1024, height: 900 });
      await page.goto(`/reports/sample?fight=${fight}`);
      const chart = page.getByTestId('time-chart');
      await expect(chart.locator('[data-state="ready"]').first()).toBeVisible();
      const end = chart.getByTestId('window-end');
      await expect(end).toHaveAttribute('max', String(durationMs));
      // The browser's own reading of the input: a step that did not divide the fight snapped it down.
      expect(await end.inputValue()).toBe(String(durationMs));
      expect(await end.evaluate((input: HTMLInputElement) => input.valueAsNumber / Number(input.max))).toBe(
        1,
      );
      await expect(end).toHaveAttribute('aria-valuetext', label);
      await expect(chart.getByTestId('brush-domain-label')).toHaveText(label);
    });
  }
});
