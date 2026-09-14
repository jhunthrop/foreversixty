// web/tests/e2e/report-tables.spec.ts
import { expect, test } from '@playwright/test';

const FIGHT = '/reports/fixture2abcd?fight=3';

test('the summary tab lists the roster with specs and deaths', async ({ page }) => {
  await page.goto(FIGHT);
  await expect(page.getByTestId('summary-tab')).toContainText('5 players');
  await expect(page.getByTestId('summary-tab')).toContainText('1 deaths');
  await expect(page.getByTestId('roster-Player-4184-000000A1')).toContainText('Protection');
  await expect(page.getByTestId('combatants')).toContainText('item level 183');
});

test('damage done ranks rows, colours names by class and expands to abilities', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=damage-done`);

  // The default source scope is "all friendlies" (spec section 4 and url.ts's own
  // defaultState), so Warden Kelthas -- a creature, not a player -- does not have a row
  // here at all: this is the raid's own damage ranking, the same table Warcraft Logs shows
  // by default. Baelgrim is the raid's top damage dealer in the fixture (4,400 vs.
  // Morrowlyn's 3,110), so the top row is theirs.
  const rows = page.getByTestId('actor-table').locator('li');
  await expect(rows.first()).toContainText('Baelgrim');
  const name = page.getByTestId('actor-Player-4184-000000A1').getByTestId('row-name');
  await expect(name).toHaveCSS('color', 'rgb(198, 155, 109)');

  await page.getByTestId('actor-Player-4184-000000A1').getByRole('button').first().click();
  await expect(page.getByTestId('row-detail')).toContainText('Slam');
  await expect(page.getByTestId('row-detail')).toContainText('Warden Kelthas');
});

test('the ability filter narrows every row to that ability', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=damage-done`);
  await page.getByTestId('filter-ability').selectOption({ label: 'Slam' });
  const rows = page.getByTestId('actor-table').locator('li');
  await expect(rows).toHaveCount(1);
  await expect(rows.first()).toContainText('Baelgrim');
});

test('boss damage only drops the boss’s own row and keeps the raid’s', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=damage-done`);
  await page.getByTestId('filter-boss').check();
  await expect(page.getByTestId('actor-table')).not.toContainText('Warden Kelthas · ');
  await expect(page.getByTestId('actor-table')).toContainText('Baelgrim');
});

test('a brushed window marks the split as approximate and says why', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=damage-done&start=0&end=10000`);
  await expect(page.getByTestId('approximate-note')).toContainText('Totals and per-second figures are exact');
  await expect(page.getByTestId('actor-table').getByTestId('row-amount').first()).toContainText('~');
});

test('rows carry a parse percentile from the API on a whole-fight encounter view', async ({ page }) => {
  await page.route('**/v1/rankings/percentile**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { percentile: 96.2 }, error: null, request_id: 'r' }),
    }),
  );
  await page.goto(FIGHT);
  await expect(page.getByTestId('roster-Player-4184-000000A1')).toContainText('96');
});
