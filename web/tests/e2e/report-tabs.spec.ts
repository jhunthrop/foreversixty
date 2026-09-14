// web/tests/e2e/report-tabs.spec.ts
import { expect, test } from '@playwright/test';

const REPORT = '/reports/fixture2abcd';

test('the report opens on its first fight with the chrome the spec sets', async ({ page }) => {
  await page.goto(REPORT);

  await expect(page.getByTestId('report-title')).toHaveText('Sanguine Depths, fixture night');
  await expect(page.getByTestId('fight-selector')).toBeVisible();
  await expect(page.getByTestId('fight-3')).toBeVisible();
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill');
  for (const mode of ['analyze', 'compare', 'rankings']) {
    await expect(page.getByTestId(`mode-${mode}`)).toBeEnabled();
  }
  for (const mode of ['mechanics', 'replay']) {
    await expect(page.getByTestId(`mode-${mode}`)).toBeDisabled();
    await expect(page.getByTestId(`mode-${mode}`)).toContainText('later');
  }
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
test('a fight the report does not have falls back to the first one', async ({ page }) => {
  for (const search of ['?fight=0', '?fight=999']) {
    await page.goto(`${REPORT}${search}`);
    await expect(page.getByTestId('report-error')).toHaveCount(0);
    await expect(page.getByTestId('fight-selector')).toBeVisible();
    await expect(page.getByTestId('report-placeholder')).toHaveText('3 players in this fight.');
  }
});

// A live report's next fight appears in the selector before its summary is written, so a
// fight whose file is not there yet is the normal case rather than a corrupt one. The URL
// and the selector both move on the click, so without a signal the previous fight's numbers
// would sit under the new fight's label -- one fight's data read as another's.
test('a fight whose summary does not load says so instead of passing off the last one', async ({ page }) => {
  await page.goto(REPORT);
  await page.route('**/fights/2/summary.json', (route) => route.abort());

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

  await page.route(`**/fights/${fightIndex}/summary.json`, async (route) => {
    markStarted();
    await gate;
    await (settle === 'abort' ? route.abort() : route.continue());
  });

  return { started, release: () => release() };
}

test('a stale fight failure does not fail the fight that is on screen', async ({ page }) => {
  await page.goto(REPORT);
  await expect(page.getByTestId('report-placeholder')).toHaveText('3 players in this fight.');

  const slow = await heldRoute(page, 2, 'abort');
  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await slow.started;

  // Fight 3 is picked and lands while fight 2 is still open.
  await page.getByTestId('fight-3').click();
  await expect(page.getByTestId('report-placeholder')).toHaveText('5 players in this fight.');
  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);

  const failed = page.waitForEvent('requestfailed', (request) =>
    request.url().includes('/fights/2/summary.json'),
  );
  slow.release();
  await failed;
  await page.waitForTimeout(250);

  await expect(page.getByTestId('report-fight-error')).toHaveCount(0);
  await expect(page.getByTestId('report-placeholder')).toHaveText('5 players in this fight.');
});

test('a stale fight success neither clears the current error nor paints its roster', async ({ page }) => {
  await page.goto(REPORT);
  await expect(page.getByTestId('report-placeholder')).toHaveText('3 players in this fight.');

  const slow = await heldRoute(page, 2, 'continue');
  await page.route('**/fights/3/summary.json', (route) => route.abort());

  await page.getByTestId('toggle-trash').click();
  await page.getByTestId('fight-2').click();
  await slow.started;

  // Fight 3 is picked and fails while fight 2 is still open.
  await page.getByTestId('fight-3').click();
  await expect(page.getByTestId('report-fight-error')).toHaveText('Report data did not load');
  await expect(page.getByTestId('report-placeholder')).toHaveText('3 players in this fight.');

  const answered = page.waitForResponse((response) => response.url().includes('/fights/2/summary.json'));
  slow.release();
  await answered;
  await page.waitForTimeout(250);

  // Fight 2 has one player; seeing its roster or losing fight 3's alert would both be the
  // stale answer speaking for the selected fight.
  await expect(page.getByTestId('report-fight-error')).toHaveText('Report data did not load');
  await expect(page.getByTestId('report-placeholder')).toHaveText('3 players in this fight.');
});
