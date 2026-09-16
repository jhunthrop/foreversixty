// web/tests/e2e/report-tables.spec.ts
import { expect, test } from '@playwright/test';
import { serveDuckdbRuntime } from './support/duckdb-runtime';

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

test('a brushed window measures the table and each opened row from the fight’s events', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto(`${FIGHT}&tab=damage-done&start=0&end=10000`);
  await expect(page.getByTestId('approximate-note')).toBeVisible();

  // Amount is `actor.effective`: window.ts measures it directly from the one-second
  // series, so it is exact under a bare window and never carries the `~` mark.
  await expect(page.getByTestId('actor-Player-4184-000000A1').getByTestId('row-amount')).not.toContainText(
    '~',
  );

  // The per-ability split is prorated by the summary, so the page measures it from the
  // events by default: first the whole table, then each row as it is opened.
  await expect(page.getByTestId('table-measured')).toBeVisible({ timeout: 60_000 });
  await page.getByTestId('actor-Player-4184-000000A1').getByRole('button').first().click();
  await expect(page.getByTestId('row-measured')).toBeVisible({ timeout: 60_000 });
  await expect(page.getByTestId('row-detail')).not.toContainText('~');
});

test('boss damage only marks the split as approximate even at a whole-fight window', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto(`${FIGHT}&tab=damage-done`);
  await expect(page.getByTestId('approximate-note')).toHaveCount(0);

  await page.getByTestId('filter-boss').check();
  await expect(page.getByTestId('approximate-note')).toBeVisible();

  // A target/boss filter scales the per-ability and per-target splits the same way the
  // window does (filters.ts's `targetShare`), so the page measures the table from the
  // events -- and Amount, exact at a whole-fight window, never carries the mark.
  await expect(page.getByTestId('table-measured')).toBeVisible({ timeout: 60_000 });
  await expect(page.getByTestId('actor-Player-4184-000000A1').getByTestId('row-amount')).not.toContainText(
    '~',
  );
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

test('the Summary tab gives the healer a percentile from the healing metric, not DPS', async ({ page }) => {
  const requestUrls: string[] = [];
  await page.route('**/v1/rankings/percentile**', async (route) => {
    requestUrls.push(route.request().url());
    await route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { percentile: 88.1 }, error: null, request_id: 'r' }),
    });
  });
  // FIGHT carries no &tab=, so this is the Summary tab -- the product's default landing
  // view, where a healer's headline number must read as a healing parse.
  await page.goto(FIGHT);
  await expect(page.getByTestId('roster-Player-4184-000000A2')).toContainText('88');

  const healerRequest = requestUrls.find((url) => url.includes('spec=Holy'));
  expect(healerRequest).toBeDefined();
  expect(healerRequest).toContain('metric=hps');
  expect(healerRequest).not.toContain('metric=dps');
});

test('buffs and debuffs each show only their own kind, with a real uptime', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=buffs`);
  await expect(page.getByTestId('aura-table')).toContainText('Power Word: Fortitude');
  await expect(page.getByTestId('aura-table')).not.toContainText('Necrotic Wound');
  // 31 s of the fight's 60 s wall length (engine 0.2.6 measures a closed fight by it).
  await expect(page.getByTestId('aura-1243-Player-4184-000000A1').getByTestId('aura-uptime')).toHaveText(
    '51.7%',
  );

  await page.goto(`${FIGHT}&tab=debuffs`);
  await expect(page.getByTestId('aura-table')).toContainText('Necrotic Wound');
  await expect(page.getByTestId('aura-table')).not.toContainText('Power Word: Fortitude');
});

test('a brushed window draws the uptime bars from its own start, inside the track', async ({ page }) => {
  // A window starting at 20 s: a bar positioned on the fight's clock would sit past the
  // track's right edge and stretch the page sideways.
  await page.goto(`${FIGHT}&tab=buffs&start=20000&end=40000`);
  const row = page.getByTestId('aura-1243-Player-4184-000000A1');
  await expect(row).toBeVisible();
  const lefts = await row
    .locator('span[style*="left:"]')
    .evaluateAll((spans) => spans.map((span) => Number.parseFloat((span as HTMLElement).style.left)));
  expect(lefts.length).toBeGreaterThan(0);
  for (const left of lefts) expect(left).toBeGreaterThanOrEqual(0);
  for (const left of lefts) expect(left).toBeLessThan(100);
  const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth);
  expect(overflow).toBeLessThanOrEqual(0);
});

test('casts show the caster, the count and a sequence timeline', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=casts`);
  await expect(page.getByTestId('cast-table')).toContainText('Frostbolt');
  await expect(page.getByTestId('cast-Player-4184-000000A3-116')).toContainText('Morrowlyn');
});

test('interrupts, dispels, resources and threat each render or say they are empty', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=interrupts`);
  await expect(page.getByTestId('table-empty')).toContainText('Nothing was interrupted');

  await page.goto(`${FIGHT}&tab=resources`);
  await expect(page.getByTestId('resource-graphs')).toBeVisible();
  await expect(page.getByTestId('resource-graphs').getByTestId('resource-zero').first()).toBeVisible();

  await page.goto(`${FIGHT}&tab=threat`);
  await expect(page.getByTestId('threat-table')).toContainText('Baelgrim');
  await expect(page.getByTestId('threat-incomplete')).toContainText('does not yet carry every class');
});

test('the trash fight’s interrupts and dispels do have rows', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=1&tab=interrupts');
  await expect(page.getByTestId('exchange-table')).toContainText('Pummel');
  await page.goto('/reports/fixture2abcd?fight=1&tab=dispels');
  await expect(page.getByTestId('exchange-table')).toContainText('Purify');
});

test('interrupts and dispels state their counts are the whole fight’s, unaffected by the window', async ({
  page,
}) => {
  await page.goto('/reports/fixture2abcd?fight=1&tab=interrupts&start=0&end=1000');
  await expect(page.getByTestId('exchange-wholefight-note')).toContainText("whole fight's totals");
  await expect(page.getByTestId('exchange-table')).toContainText('†');
});

test('a brushed window marks cast counts and threat as approximate', async ({ page }) => {
  await page.goto(`${FIGHT}&tab=casts&start=0&end=10000`);
  await expect(page.getByTestId('cast-approximate-note')).toBeVisible();
  await expect(page.getByTestId('cast-time-note')).toContainText('whole fight');

  await page.goto(`${FIGHT}&tab=threat&start=0&end=10000`);
  await expect(page.getByTestId('threat-approximate-note')).toBeVisible();
});

// Sunwick's one Heal on Baelgrim in fight 3: 1,900 cast, 320 of it overheal. The pair
// carries the overheal from the engine, so the Over column is on the plain row, and the
// ability filter measures the healing table on a pull the way it does the damage tables.
test('a healing row shows what each target did not need, and the ability filter measures it', async ({
  page,
}) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto('/reports/fixture2abcd?fight=3&tab=healing');
  const row = page.getByTestId('actor-Player-4184-000000A2');
  await row.getByRole('button').first().click();
  await expect(row.getByTestId('target-overheal')).toContainText('320');

  await page.goto('/reports/fixture2abcd?fight=3&tab=healing&ability=2060');
  await expect(page.getByTestId('table-measured')).toBeVisible({ timeout: 60_000 });
  await expect(page.getByTestId('actor-Player-4184-000000A2')).toBeVisible();

  await page.goto('/reports/fixture2abcd?fight=3&tab=healing&ability=999999');
  await expect(page.getByTestId('table-empty')).toBeVisible({ timeout: 60_000 });
});

// The dropdown's own words pasted into a link: `ability=Frostbolt` resolves to the id of
// that ability in the tab, and a name the tab has no row for is said out loud instead of
// the whole table showing under a filter that silently did nothing.
test('an ability named in the url resolves to its id, or says it is unknown', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=damage-done&ability=Frostbolt');
  await expect(page).toHaveURL(/ability=116(&|$)/);
  await expect(page.getByTestId('filter-ability')).toHaveValue('116');

  await page.goto('/reports/fixture2abcd?fight=3&tab=damage-done&ability=Nothing%20Of%20The%20Sort');
  await expect(page.getByTestId('report-unknown-ability')).toContainText('Nothing Of The Sort');
  await expect(page).not.toHaveURL(/ability=/);
});
