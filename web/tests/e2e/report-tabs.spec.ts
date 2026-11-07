// web/tests/e2e/report-tabs.spec.ts
import { expect, test } from './support/test';
import { serveDuckdbRuntime } from './support/duckdb-runtime';
import { heldRoute } from './support/held-route';
import { openChartIfCollapsed } from './support/report-chart';
import { choose } from './support/select';
import { pickFight, openTab } from './support/report-frame';

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
// The rows are StatRows now (logs spec 2026-11-07 6.2), so the role picks them out.
function summaryRosterRows(page: import('@playwright/test').Page): import('@playwright/test').Locator {
  return page.getByTestId('summary-tab').locator('[role="row"][data-testid^="roster-"]');
}

const FIGHT_SELECT = 'select[data-testid="report-fight"]';

test('the report opens on its first fight with the chrome the spec sets', async ({ page }) => {
  await page.goto(REPORT);

  await expect(page.getByTestId('report-title')).toHaveText('Sanguine Depths, fixture night');
  // One fight select in the control bar (logs spec 6.1, decision R2), on the first boss pull.
  await expect(page.locator(FIGHT_SELECT)).toHaveValue('3');
  await expect(page.locator(`${FIGHT_SELECT} option[value="3"]`)).toHaveText(/Warden Kelthas #1 · Kill/);
  // One tab strip (decision R3): the modes and views of old are tabs, Rankings a link.
  const tabs = page.getByTestId('report-tabs');
  for (const label of [
    'Summary',
    'Damage done',
    'Healing',
    'Damage taken',
    'Deaths',
    'Mechanics',
    'Timeline',
  ])
    await expect(tabs.getByRole('link', { name: label, exact: true })).toBeVisible();
  await expect(tabs.getByRole('link', { name: 'Summary', exact: true })).toHaveAttribute(
    'aria-current',
    'page',
  );
  await expect(page.getByTestId('mode-bar')).toHaveCount(0);
  await expect(page.getByTestId('report-rankings-link')).toHaveText('Rankings for Warden Kelthas ›');
});

test('the fight list numbers boss pulls and hides trash that killed nothing', async ({ page }) => {
  await page.goto(REPORT);
  const options = page.locator(`${FIGHT_SELECT} option`);
  await expect(page.locator(`${FIGHT_SELECT} optgroup`).first()).toHaveAttribute('label', 'Boss pulls');
  await expect(options.first()).toHaveText('All boss pulls');
  await expect(page.locator(`${FIGHT_SELECT} option[value="3"]`)).toHaveText(/^Warden Kelthas #1/);
});

test('every control the spec names writes itself into the URL and reads back', async ({ page }) => {
  await page.goto(REPORT);
  await pickFight(page, 2);
  await openTab(page, 'healing');
  await expect(page).toHaveURL(/\?fight=2&tab=healing$/);
  await openTab(page, 'events');
  await expect(page).toHaveURL(/\?fight=2&tab=events$/);

  await page.goto(`${REPORT}?fight=3&tab=deaths&source=Player-4184-000000A1`);
  await expect(
    page.getByTestId('report-tabs').getByRole('link', { name: 'Deaths', exact: true }),
  ).toHaveAttribute('aria-current', 'page');
  await expect(page.getByTestId('source-scope')).toHaveValue('Player-4184-000000A1');
  await expect(page.locator(FIGHT_SELECT)).toHaveValue('3');
  // An old link's view lands on its tab.
  await page.goto(`${REPORT}?fight=3&view=queries`);
  await expect(page.getByTestId('report-tabs')).toContainText(/Queries/);
  await expect(page.getByTestId('queries-view')).toBeVisible();
});

test('the tabs past the first seven sit in More, which names the open one', async ({ page }) => {
  await page.goto(`${REPORT}?fight=3`);
  const wide = (page.viewportSize()?.width ?? 1280) >= 1760;
  test.skip(wide, 'from 1760 the strip shows twelve tabs');
  await openTab(page, 'casts');
  await expect(page).toHaveURL(/tab=casts/);
  await expect(page.locator('[data-testid="report-tabs-more"]:visible')).toHaveText(/More: Casts/);
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
    await expect(page.locator(FIGHT_SELECT)).toHaveValue('3');
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
  await pickFight(page, 2);

  await expect(page.locator(FIGHT_SELECT)).toHaveValue('2');
  await expect(page.getByTestId('report-fight-error')).toHaveText('Report data did not load');

  // And it goes away again: an alert that outlived the fight it described would be the
  // same mislabelling in reverse.
  await pickFight(page, 3);
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
async function heldFightRoute(
  page: import('@playwright/test').Page,
  fightIndex: number,
  settle: 'abort' | 'continue',
): Promise<{ started: Promise<void>; release: () => void }> {
  return heldRoute(page, `**/fights/${fightIndex}/summary.json*`, settle);
}

test('a stale fight failure does not fail the fight that is on screen', async ({ page }) => {
  await page.goto(`${REPORT}?fight=1`);
  await expect(summaryRosterRows(page)).toHaveCount(3);

  const slow = await heldFightRoute(page, 2, 'abort');
  // Fight 1 is on screen, so its trash entry is already open (and cannot be folded away).
  await pickFight(page, 2);
  await slow.started;

  // Fight 3 is picked and lands while fight 2 is still open.
  await pickFight(page, 3);
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

  const slow = await heldFightRoute(page, 2, 'continue');
  await page.route('**/fights/3/summary.json*', (route) => route.abort());

  await pickFight(page, 2);
  await slow.started;

  // Fight 3 is picked and fails while fight 2 is still open.
  await pickFight(page, 3);
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

test('a lazy mode shows a sized skeleton, not a blank panel, while its chunk loads', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  // Settled first, so the only chunk the held route below can catch is the one the click
  // asks for: everything the island needs to paint this fight is already fetched.
  await expect(page.getByTestId('report-title')).toBeVisible();
  await expect(page.getByTestId('summary-tab')).toBeVisible();

  // The chunk is held open rather than delayed by a fixed setTimeout: a timed delay races
  // the assertion below (miss the window and the skeleton is already gone), so the test
  // owns when the loading frame ends instead of hoping it outlasts a poll -- the shape
  // sim-saved.spec.ts holds its saved-sim fetch with, and handoffs-addon.spec.ts /v1/me.
  //
  // The pattern is the island's whole chunk family, not 'CompareMode*.js': the report
  // island is built by its own Vite pass (src/report-island.ts), which names every
  // dynamic chunk report-island-<hash>.js. The name-based pattern this replaces matched
  // nothing at all, so the delay it was supposed to impose never happened and the test
  // was only ever catching the real, unslowed loading frame by luck.
  const chunk = await heldRoute(page, '**/report-island-*.js', 'continue');
  await openTab(page, 'compare');
  await chunk.started;

  await expect(page.getByTestId('lazy-view-skeleton')).toBeVisible();

  chunk.release();
  await expect(page.getByTestId('lazy-view-skeleton')).toHaveCount(0);
  await expect(page.getByTestId('compare-mode')).toBeVisible();
});

test('compare puts two fights side by side with a per-player difference', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=compare');
  await expect(page.getByTestId('compare-mode')).toContainText('Pick a second fight');
  await choose(page.getByTestId('compare-with'), { index: 1 });
  // One table at every width: a phone folds the other figures under the name.
  await expect(page.getByTestId('compare-table')).toBeVisible();
  await expect(page.getByTestId('compare-delta').first()).toContainText(/[+\u2212]/);
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
  await page.goto('/reports/fixture2abcd?fight=3&tab=compare&cmetric=damage_done');

  const slow = await heldFightRoute(page, 1, 'continue');
  await choose(page.getByTestId('compare-with'), '1');
  await slow.started;

  // The visitor changes their mind before fight 1's summary has even arrived.
  await choose(page.getByTestId('compare-with'), '2');
  const delta = page.getByTestId('compare-Player-4184-000000A3').getByTestId('compare-delta');
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

test('the full event stream lists aura refreshes under the Auras chip', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto('/reports/fixture2abcd?fight=3&tab=events');
  await page.getByTestId('events-stream').click();
  await expect(page.getByTestId('event-list')).toContainText('refreshed Power Word: Fortitude', {
    timeout: 60_000,
  });
  // The Auras chip governs it: off, the refresh goes with the applications.
  await page.getByLabel('Auras', { exact: true }).uncheck();
  await expect(page.getByTestId('event-list')).not.toContainText('refreshed Power Word: Fortitude');
});

test('the chart bands the fight’s phases and names them at their left edge', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await openChartIfCollapsed(page);
  const bands = page.getByTestId('time-chart').getByTestId('phase-band');
  await expect(bands).toHaveCount(2);
  await expect(bands.first()).toHaveText('Phase 1');
  await expect(bands.nth(1)).toHaveText('Phase 2');
});

test('a phase is a window preset, so every table reads per phase in one click', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await openChartIfCollapsed(page);
  // The six-button preset row is a select now (design review 2026-09-26 finding 4); the
  // option labels are the old buttons' own names, unchanged.
  await choose(page.getByTestId('window-select'), { label: 'Phase 2 · 14.0s to 1:00' });
  await expect(page).toHaveURL(/start=14000&end=60000/);
});

test('a fight with no phases shows no bands and no phase presets', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4');
  await openChartIfCollapsed(page);
  await expect(page.getByTestId('phase-band')).toHaveCount(0);
  await expect(page.getByTestId('window-select')).not.toContainText('Phase');
});

test('the fight list says nothing about a phase on a kill', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await expect(page.locator(`${FIGHT_SELECT} option[value="3"]`)).toHaveText(/· Kill ·/);
});
