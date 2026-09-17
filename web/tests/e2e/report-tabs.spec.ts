// web/tests/e2e/report-tabs.spec.ts
import { expect, test } from '@playwright/test';
import { serveDuckdbRuntime } from './support/duckdb-runtime';

const REPORT = '/reports/fixture2abcd';

// These tests care about which fight's data is on screen, and used to read that off a
// Task-10 placeholder's "N players in this fight." text. Task 11 replaced the placeholder
// with the summary tab's real roster list, and Task 14 removed the placeholder branch
// entirely -- so the signal these tests want is now the summary tab's own roster rows
// (SummaryTab.svelte's `<li data-testid={`roster-${row.guid}`}>`), one per player the
// selected fight's summary actually reports: 3 for fight 1, 1 for fight 2, 5 for fight 3
// (src/fixtures/report/fights/<n>/summary.json). Scoped to `li`, not just `[data-testid^=
// "roster-"]`: each row also carries a nested `data-testid="roster-active"` span for its
// Active column, which shares the "roster-" prefix and would otherwise double the count.
function summaryRosterRows(page: import('@playwright/test').Page): import('@playwright/test').Locator {
  return page.getByTestId('summary-tab').locator('li[data-testid^="roster-"]');
}

test('the report opens on its first fight with the chrome the spec sets', async ({ page }) => {
  await page.goto(REPORT);

  await expect(page.getByTestId('report-title')).toHaveText('Sanguine Depths, fixture night');
  await expect(page.getByTestId('fight-selector')).toBeVisible();
  await expect(page.getByTestId('fight-3')).toBeVisible();
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill');
  for (const mode of ['analyze', 'compare', 'rankings', 'mechanics']) {
    await expect(page.getByTestId(`mode-${mode}`)).toBeEnabled();
    await expect(page.getByTestId(`mode-${mode}`)).not.toContainText('later');
  }
  await expect(page.getByTestId('mode-replay')).toBeDisabled();
  await expect(page.getByTestId('mode-replay')).toContainText('later');
  for (const view of ['tables', 'timelines', 'events', 'queries']) {
    await expect(page.getByTestId(`view-${view}`)).toBeVisible();
  }
  await expect(page.getByRole('tab', { name: 'Damage Done' })).toBeVisible();
});

test('trash is folded away behind a count, and unfolds', async ({ page }) => {
  await page.goto(REPORT);
  await expect(page.getByTestId('fight-1')).toHaveCount(0);
  await page.getByTestId('toggle-trash').click();
  await expect(page.getByTestId('fight-1')).toBeVisible();
});

test('every control the spec names writes itself into the URL and reads back', async ({ page }) => {
  await page.goto(REPORT);

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await page.getByTestId('tab-healing').click();
  await page.getByTestId('view-events').click();

  await expect(page).toHaveURL(/\?fight=2&view=events&tab=healing$/);

  await page.goto(`${REPORT}?fight=3&view=tables&tab=deaths&source=Player-4184-000000A1`);
  await expect(page.getByTestId('tab-deaths')).toHaveAttribute('aria-selected', 'true');
  await expect(page.getByTestId('source-scope')).toHaveValue('Player-4184-000000A1');
  await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');
});

test('the source scope lists the fight’s players', async ({ page }) => {
  await page.goto(`${REPORT}?fight=3`);
  const options = page.getByTestId('source-scope').locator('option');
  await expect(options).toHaveCount(7);
  await expect(options.nth(0)).toHaveText('All friendlies');
  await expect(options.nth(1)).toHaveText('All enemies');
  await expect(options.filter({ hasText: 'Elyra Duskvale' })).toHaveCount(1);
});

// The defect this guards is in parseReportState, which is plan-mandated and total by
// design: it validates ?fight= as a run of digits and cannot know the report's fights,
// so ?fight=0 parses even though fight_index is 1-based. The island resolves it against
// report.json instead of asking the edge for fights/0/summary.json and failing the page.
test('a fight the report does not have falls back to the first boss pull', async ({ page }) => {
  for (const search of ['?fight=7', '?fight=999']) {
    await page.goto(`${REPORT}${search}`);
    await expect(page.getByTestId('report-error')).toHaveCount(0);
    await expect(page.getByTestId('fight-selector')).toBeVisible();
    await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');
    await expect(summaryRosterRows(page)).toHaveCount(5);
  }
  // Fight 0 is the whole night, which every report with a boss pull has.
  await page.goto(`${REPORT}?fight=0`);
  await expect(page.getByTestId('night-view')).toBeVisible();
});

// A live report's next fight appears in the selector before its summary is written, so a
// fight whose file is not there yet is the normal case rather than a corrupt one. The URL
// and the selector both move on the click, so without a signal the previous fight's numbers
// would sit under the new fight's label -- one fight's data read as another's.
test('a fight whose summary does not load says so instead of passing off the last one', async ({ page }) => {
  await page.goto(REPORT);
  await page.route('**/fights/2/summary.json*', (route) => route.abort());

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();

  await expect(page.getByTestId('fight-2')).toHaveAttribute('aria-current', 'true');
  await expect(page.getByTestId('report-fight-error')).toHaveText('Report data did not load');

  // And it goes away again: an alert that outlived the fight it described would be the
  // same mislabelling in reverse.
  await page.getByTestId('fight-3').click();
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
});

// Picking a fight does not cancel the one before it, so two summary requests can be in
// flight at once and settle in either order. Whichever settles last must not speak for a
// fight nobody is looking at: one fight's roster under another fight's name, or an alert
// that outlived or preceded the fight it describes, is the same mislabelling either way.
//
// The two tests below interleave rather than await in turn -- the sequential case is the
// one that cannot race -- by holding the slow fight's response open in the route handler
// until the fast one has already settled. The fixture's three fights have three, one and
// five players, so the rendered line names which fight is on screen.
async function heldRoute(
  page: import('@playwright/test').Page,
  fightIndex: number,
  settle: 'abort' | 'continue',
): Promise<{ started: Promise<void>; release: () => void }> {
  let markStarted = (): void => {};
  let release = (): void => {};
  const started = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));

  await page.route(`**/fights/${fightIndex}/summary.json*`, async (route) => {
    markStarted();
    await gate;
    await (settle === 'abort' ? route.abort() : route.continue());
  });

  return { started, release: () => release() };
}

test('a stale fight failure does not fail the fight that is on screen', async ({ page }) => {
  await page.goto(`${REPORT}?fight=1`);
  await expect(summaryRosterRows(page)).toHaveCount(3);

  const slow = await heldRoute(page, 2, 'abort');
  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await slow.started;

  // Fight 3 is picked and lands while fight 2 is still open.
  await page.getByTestId('fight-3').click();
  await expect(summaryRosterRows(page)).toHaveCount(5);
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);

  const failed = page.waitForEvent('requestfailed', (request) =>
    request.url().includes('/fights/2/summary.json'),
  );
  slow.release();
  await failed;
  await page.waitForTimeout(250);

  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
  await expect(summaryRosterRows(page)).toHaveCount(5);
});

test('a stale fight success neither clears the current error nor paints its roster', async ({ page }) => {
  await page.goto(`${REPORT}?fight=1`);
  await expect(summaryRosterRows(page)).toHaveCount(3);

  const slow = await heldRoute(page, 2, 'continue');
  await page.route('**/fights/3/summary.json*', (route) => route.abort());

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await slow.started;

  // Fight 3 is picked and fails while fight 2 is still open.
  await page.getByTestId('fight-3').click();
  await expect(page.getByTestId('report-fight-error')).toHaveText('Report data did not load');
  await expect(summaryRosterRows(page)).toHaveCount(3);

  const answered = page.waitForResponse((response) => response.url().includes('/fights/2/summary.json'));
  slow.release();
  await answered;
  await page.waitForTimeout(250);

  // Fight 2 has one player; seeing its roster or losing fight 3's alert would both be the
  // stale answer speaking for the selected fight.
  await expect(page.getByTestId('report-fight-error')).toHaveText('Report data did not load');
  await expect(summaryRosterRows(page)).toHaveCount(3);
});

test('compare puts two fights side by side with a per-player difference', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&mode=compare');
  await expect(page.getByTestId('compare-mode')).toContainText('Pick a second fight');
  await page.getByTestId('compare-with').selectOption({ index: 1 });
  // A phone gets one card per player instead of the table.
  const phone = (page.viewportSize()?.width ?? 1280) < 768;
  await expect(page.getByTestId(phone ? 'compare-cards' : 'compare-table')).toBeVisible();
  await expect(page.getByTestId(phone ? 'compare-card-delta' : 'compare-delta').first()).toContainText(
    /[+-]/,
  );
});

// Picking a second fight does not cancel the request for the one picked before it, so two
// of these can settle in either order -- the same hazard the fight-switch races above
// cover for the primary fight, now for CompareMode's own `right` fetch. Fight 3 (the
// fixture's current fight) has Morrowlyn at 3,110 damage; fight 1 has her at 1,484; fight
// 2 has her at 0. Picking fight 1, then fight 2 before fight 1's summary arrives, has to
// land on fight 2's zero (delta +3,110) and stay there once fight 1's late answer shows up
// -- not fall back to fight 1's 1,484 (delta +1,626), which is what an unguarded write
// would paint.
test('a stale compare answer does not overwrite the fight actually selected', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&mode=compare&cmetric=damage_done');

  const slow = await heldRoute(page, 1, 'continue');
  await page.getByTestId('compare-with').selectOption('1');
  await slow.started;

  // The visitor changes their mind before fight 1's summary has even arrived.
  await page.getByTestId('compare-with').selectOption('2');
  const phone = (page.viewportSize()?.width ?? 1280) < 768;
  const delta = phone
    ? page.getByTestId('compare-card-Player-4184-000000A3').getByTestId('compare-card-delta')
    : page.getByTestId('compare-Player-4184-000000A3').getByTestId('compare-delta');
  await expect(delta).toHaveText('+3,110');

  const answered = page.waitForResponse((response) =>
    /\/fights\/1\/summary\.json(\?|$)/.test(response.url()),
  );
  slow.release();
  await answered;
  await page.waitForTimeout(250);

  // Fight 1's late answer must not speak for the fight actually selected (fight 2).
  await expect(delta).toHaveText('+3,110');
});

test('rankings inside the report marks this report’s own row', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          rows: [
            {
              rank: 1,
              player: { key: 'us/normal/other-raider', name: 'Other Raider', class: 'Mage', spec: 'Fire' },
              guild: { name: 'Elsewhere', ruleset: 'normal', region: 'us' },
              value: 2200,
              size: 40,
              fought_at: '2026-12-09T22:10:00Z',
              duration_ms: 200000,
              talent_split: '0/31/20',
              trinkets: [],
              buff_count: 9,
              report_id: 'otherreport1',
              fight_index: 1,
              state: 'ok',
            },
            {
              rank: 2,
              player: { key: 'us/normal/baelgrim', name: 'Baelgrim', class: 'Warrior', spec: 'Protection' },
              guild: { name: 'The Last Watch', ruleset: 'normal', region: 'us' },
              value: 110,
              size: 5,
              fought_at: '2026-09-26T20:12:40Z',
              duration_ms: 40000,
              talent_split: '31/20/0',
              trinkets: [],
              buff_count: 4,
              report_id: 'fixture2abcd',
              fight_index: 3,
              state: 'at_risk',
            },
          ],
          total: 2,
          page: 1,
          per_page: 100,
          updated_at: '2026-12-09T22:15:00Z',
        },
        error: null,
        request_id: 'r',
      }),
    }),
  );

  await page.goto('/reports/fixture2abcd?fight=3&mode=rankings');
  await expect(page.getByTestId('rankings-rows')).toBeVisible();
  await expect(page.getByTestId('rankings-mine')).toContainText('Baelgrim');
  await expect(page.getByTestId('rankings-mine')).toContainText('at risk');
  await expect(page.getByRole('link', { name: /Full rankings/ })).toHaveAttribute(
    'href',
    '/rankings/warden-kelthas',
  );
});

test('rankings says so plainly on a trash fight', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=1&mode=rankings');
  await expect(page.getByTestId('rankings-mode')).toContainText('Trash is not ranked');
});

// Switching the metric does not cancel the request for the one picked before it, so two
// of these can be in flight and settle in either order too -- `wantedMetric`'s half of
// RankingsMode's guard. dps is held open; hps answers immediately and paints the screen;
// dps's late answer, once it does arrive, must not overwrite hps's numbers with its own.
test('a stale rankings answer for the metric that is no longer selected is dropped', async ({ page }) => {
  let markStarted = (): void => {};
  let release = (): void => {};
  const started = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));

  function pageBody(total: number): string {
    return JSON.stringify({
      ok: true,
      data: { rows: [], total, page: 1, per_page: 100, updated_at: '2026-12-09T22:15:00Z' },
      error: null,
      request_id: 'r',
    });
  }

  await page.route('**/v1/rankings?**', async (route) => {
    const metric = new URL(route.request().url()).searchParams.get('metric');
    if (metric === 'dps') {
      markStarted();
      await gate;
      await route.fulfill({ status: 200, contentType: 'application/json', body: pageBody(111) });
      return;
    }
    await route.fulfill({ status: 200, contentType: 'application/json', body: pageBody(222) });
  });

  await page.goto('/reports/fixture2abcd?fight=3&mode=rankings');
  await started;

  // The visitor picks Healing before Damage's request has even arrived.
  await page.getByTestId('rankings-metric').selectOption('hps');
  await expect(page.getByTestId('rankings-mode')).toContainText('222 ranked kills');

  const answered = page.waitForResponse((response) => response.url().includes('metric=dps'));
  release();
  await answered;
  await page.waitForTimeout(250);

  // Damage's late answer must not speak for the metric actually selected (Healing).
  await expect(page.getByTestId('rankings-mode')).toContainText('222 ranked kills');
});

// The fight selector stays visible in every mode, so a visitor can switch to a different
// encounter while a rankings request for the one they left is still in flight --
// `wantedFight`'s half of RankingsMode's guard. The fixture carries a second encounter
// fight (fight 4, Skolex the Insatiable) for exactly this: switching away from Warden
// Kelthas to Skolex is a real fight-to-fight race the way the metric test above is a
// metric-to-metric one, and both encounters are shown without unfolding trash. Warden
// Kelthas's request is held open; Skolex's answers immediately and paints the screen;
// Warden Kelthas's late answer, once it does arrive, must not overwrite Skolex's numbers.
test('a stale rankings answer for the fight that is no longer selected is dropped', async ({ page }) => {
  let markStarted = (): void => {};
  let release = (): void => {};
  const started = new Promise<void>((resolve) => (markStarted = resolve));
  const gate = new Promise<void>((resolve) => (release = resolve));

  function pageBody(total: number): string {
    return JSON.stringify({
      ok: true,
      data: { rows: [], total, page: 1, per_page: 100, updated_at: '2026-12-09T22:15:00Z' },
      error: null,
      request_id: 'r',
    });
  }

  await page.route('**/v1/rankings?**', async (route) => {
    const encounter = new URL(route.request().url()).searchParams.get('encounter');
    if (encounter === 'warden-kelthas') {
      markStarted();
      await gate;
      await route.fulfill({ status: 200, contentType: 'application/json', body: pageBody(111) });
      return;
    }
    await route.fulfill({ status: 200, contentType: 'application/json', body: pageBody(333) });
  });

  await page.goto('/reports/fixture2abcd?fight=3&mode=rankings');
  await started;

  // The visitor moves to the next kill before Warden Kelthas's request has even arrived.
  await page.getByTestId('fight-4').click();
  await expect(page.getByTestId('rankings-mode')).toContainText('333 ranked kills');
  await expect(page.getByRole('link', { name: /Full rankings/ })).toHaveAttribute(
    'href',
    '/rankings/skolex-the-insatiable',
  );

  const answered = page.waitForResponse((response) => response.url().includes('encounter=warden-kelthas'));
  release();
  await answered;
  await page.waitForTimeout(250);

  // Warden Kelthas's late answer must not speak for the fight actually selected (Skolex).
  await expect(page.getByTestId('rankings-mode')).toContainText('333 ranked kills');
});

// A death folded into the night moves with everything on its card: the hits before it
// keep their distance to the death, so "they took … over 3.0s" reads the same on the
// night as on the pull, not "over 7:29" for an eleven-second death.
test('a night death card keeps its own span', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=deaths');
  const onPull = (await page.getByTestId('deaths-tab').locator('li').first().innerText()).match(
    /over ([\d:.]+s?)/,
  )?.[1];
  await page.goto('/reports/fixture2abcd?fight=all&tab=deaths');
  const cards = page.getByTestId('deaths-tab').locator('li').filter({ hasText: 'Warden Kelthas' });
  await expect(cards.first()).toContainText(`over ${onPull}`);
});

test('the full event stream lists aura refreshes under Auras applied', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto('/reports/fixture2abcd?fight=3&view=events');
  await page.getByTestId('events-stream').click();
  await expect(page.getByTestId('event-list')).toContainText('refreshed Power Word: Fortitude', {
    timeout: 60_000,
  });
  // The "Auras applied" toggle governs it: off, the refresh goes with the applications.
  await page.getByLabel('Auras applied').uncheck();
  await expect(page.getByTestId('event-list')).not.toContainText('refreshed Power Word: Fortitude');
});

test('the chart bands the fight’s phases and names them at their left edge', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  const bands = page.getByTestId('time-chart').getByTestId('phase-band');
  await expect(bands).toHaveCount(2);
  await expect(bands.first()).toHaveText('Phase 1');
  await expect(bands.nth(1)).toHaveText('Phase 2');
});

test('a phase is a window preset, so every table reads per phase in one click', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await page
    .getByTestId('window-presets')
    .getByRole('button', { name: /Phase 2 · 14.0s to 1:00/ })
    .click();
  await expect(page).toHaveURL(/start=14000&end=60000/);
});

test('a fight with no phases shows no bands and no phase presets', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4');
  await expect(page.getByTestId('phase-band')).toHaveCount(0);
  await expect(page.getByTestId('window-presets')).not.toContainText('Phase');
});

test('the fight list says nothing about a phase on a kill', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill');
});
