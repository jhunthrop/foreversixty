// web/tests/e2e/sample-report.spec.ts
// The sample report is bundled with the site (lib/report/sample.ts): /logs offers it, the
// click lands on fights, and no API is involved at any step.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
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
    `${SAMPLE_REPORT_META.zone} · ${SAMPLE_REPORT_META.fights.filter((fight) => fight.kind === 'encounter').length} bosses · ${SAMPLE_REPORT_META.fights.length} fights`,
  );
  await expect(sample.getByTestId('report-row-when')).toHaveText('8 Nov 2026');
  await sample.getByRole('link').first().click();

  await expect(page).toHaveURL(/\/reports\/sample$/);
  await expect(page.getByTestId('report-title')).toHaveText('Sample dungeon night, 8 Nov 2026');
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
  await expect(page.getByTestId('report-sample-note')).toContainText('A real dungeon night, names changed.');
  await expect(page).not.toHaveURL(/fight=/);
  // The sample opens on its boss pull: pull 11, Saltspine, a kill (a real report opens on its last kill).
  await expect(page.locator('select[data-testid="report-fight"]')).toHaveValue('11');
  // Class colours: no name keeps the plain text colour.
  const colours = await roster
    .getByTestId(/^roster-Player-/)
    .evaluateAll((rows) =>
      rows.slice(0, 40).map((row) => getComputedStyle(row.querySelector('a') as Element).color),
    );
  expect(new Set(colours).size).toBeGreaterThan(2);
  // A trash row prints its length once.
  const row1 = await page.locator('select[data-testid="report-fight"] option[value="1"]').innerText();
  expect((row1.match(/1:08/g) ?? []).length).toBe(1);

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
      // The Section head with its title and figure, the legend and the brush; the Window choice
      // is the control bar's (logs spec 6.1).
      await expect(chart.getByText('Damage over the fight', { exact: true }).first()).toBeVisible();
      // The whole fight's figure is the band's answer: the head carries no meta of its own.
      await expect(chart.getByTestId('chart-head-meta')).toHaveCount(0);
      await expect(chart.getByTestId('chart-legend')).toBeVisible();
      await expect(fieldBox(page, 'window-select')).toBeVisible();
      await expect(chart.getByTestId('overview-brush')).toBeVisible();
      // A plot of 140 px plus its 22 px axis.
      const plot = (await chart.getByTestId('chart-plot').boundingBox())!;
      expect(plot.height).toBeGreaterThanOrEqual(140 + 22);
      // The full chart's frame reserves its loaded height: nothing moves as it arrives.
      expect(await reportShift(page)).toBe(0);
    }
  });

  test('on a phone folds into a Disclosure, opened in place', async ({ page, isMobile }) => {
    test.skip(!isMobile, 'the phone layout');
    await page.goto('/reports/sample');
    const fold = page.getByTestId('chart-collapsed');
    await expect(fold).not.toHaveAttribute('open', '');
    await fold.locator('summary').click();
    await expect(page.getByTestId('time-chart').getByTestId('chart-legend')).toBeVisible();
  });
});

test.describe('the brush on a fight that does not end on a whole second', () => {
  // 1 s, 2 s and 5 s points (lib/chart/buckets.ts); the 10 s case is OverviewBrush.test.ts's.
  const FIGHTS = [
    { fight: 2, durationMs: 18_598, label: '0:18.6' },
    { fight: 4, durationMs: 129_839, label: '2:10' },
    { fight: 57, durationMs: 432_761, label: '7:13' },
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

test.describe('the sample’s counts say what they count', () => {
  // An open-world log sees every player in range: pull 1 has 57 of them and the logger's party is
  // four of them (the pull's own flags, engine 0.6.0). The default scope is the group; Everyone in range is one pick away.
  test('the default is the logger’s group, and nothing claims a raid', async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 900 });
    await page.goto('/reports/sample');
    const source = page.getByTestId('source-scope');
    await expect(source.locator('option:checked')).toHaveText('Your group');
    await expect(page.getByTestId('report-this-pull')).not.toContainText('Deaths');
    const chart = page.getByTestId('time-chart');
    await expect(chart.locator('[data-state="ready"]').first()).toBeVisible();
    await expect(chart.getByTestId('chart-head-meta')).toHaveCount(0);
    await expect(chart.getByTestId('chart-legend')).toContainText('Group median');
    await expect(chart).not.toContainText('Raid');
    // Nobody died on the boss pull: the group's chart marks no death.
    await expect(chart.getByTestId('death-marker')).toHaveCount(0);
    // A trash row's counts are true and say whose they are: every mob that died in range.
    await expect(page.locator('select[data-testid="report-fight"] option[value="1"]')).toHaveAttribute(
      'data-description',
      /· 10 mobs died/,
    );
    // The band (directions b1): the pull in the meta line, the group's DPS as its figure; This pull
    // repeats neither.
    await expect(page.getByTestId('report-subtitle')).toHaveText(
      'Excavation Site: Wetlands · Pull 11, Saltspine #1 · Kill · 0:39 · 5 players',
    );
    await expect(page.getByTestId('report-answer')).toHaveText(/345\s*group DPS/);
  });

  test('Everyone in range counts the players seen and totals, not a raid', async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 900 });
    await page.goto('/reports/sample?fight=1&source=everyone');
    await expect(page.getByTestId('summary-tab').getByTestId(/^roster-Player-/)).toHaveCount(57);
    const chart = page.getByTestId('time-chart');
    await expect(chart.locator('[data-state="ready"]').first()).toBeVisible();
    await expect(chart.getByTestId('chart-head-meta')).toHaveCount(0);
    await expect(fieldBox(chart, 'players')).toContainText(/Top 4, bottom 4 of \d+ players seen/);
    // The longer field still shows its value whole (the dropdown's fitted-field rule).
    const label = fieldBox(chart, 'players').locator('.sel-label');
    expect(await label.evaluate((el) => el.scrollWidth <= el.clientWidth + 1)).toBe(true);
    await expect(chart.getByTestId('death-marker').first()).toBeVisible();
    await expect(chart.getByTestId('chart-legend')).toContainText('Median player');
  });

  test('a pull the group was not in is everyone in range, said so', async ({ page }) => {
    // The logger is in every pull of this log: pull 7 is served as one the group missed.
    await page.route('**/sample-report/sample/report.json', async (route) => {
      const response = await route.fetch();
      const file = (await response.json()) as {
        fights: { index: number; players: string[]; group?: string[] | null }[];
      };
      const group = new Set(file.fights.flatMap((fight) => fight.group ?? []));
      for (const fight of file.fights)
        if (fight.index === 7) {
          fight.players = fight.players.filter((guid) => !group.has(guid));
          fight.group = [];
        }
      await route.fulfill({ response, json: file });
    });
    await page.setViewportSize({ width: 1440, height: 900 });
    await page.goto('/reports/sample?fight=7');
    await expect(page.getByTestId('source-scope').locator('option:checked')).toHaveText('Everyone in range');
    await expect(page.getByTestId('report-group-missing')).toHaveText('Your group wasn’t in this pull.');
    await expect(page.getByTestId('time-chart').getByTestId('chart-legend')).toContainText('Median player');
  });
});
