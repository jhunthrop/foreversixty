// web/tests/e2e/report-duplicates.spec.ts
import { readFile } from 'node:fs/promises';
import path from 'node:path';

import { expect, test } from './support/test';

import { pickFight, openTab } from './support/report-frame';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');

// Real fights repeat timestamps: a multi-target spell lands two casts in one millisecond,
// a trinket and its proc apply on the same tick, a DoT tick and a cleave hit a dying player
// in the same instant. The timelines and deaths views keyed their rows on the timestamp
// alone, so the first real wipe uploaded threw a duplicate-key error and rendered nothing.
test('a fight whose events share timestamps still renders its timelines and deaths', async ({ page }) => {
  const summary = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const cast = summary.casts[0];
  cast.sequence = [...cast.sequence, cast.sequence[0], cast.sequence[0]];
  const aura = summary.auras[0];
  aura.segments = [...aura.segments, { ...aura.segments[0] }];
  const death = summary.deaths[0];
  death.last = [...death.last, { ...death.last[0] }];
  summary.deaths = [...summary.deaths, { ...death, guid: death.guid, at_ms: death.at_ms + 1 }];

  const errors: string[] = [];
  page.on('pageerror', (error) => errors.push(error.message));
  await page.route('**/fights/3/summary.json*', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary) }),
  );

  await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
  await expect(page.locator('[data-testid^="lane-"]').first()).toBeVisible();

  await page.goto('/reports/fixture2abcd?fight=3&tab=deaths');
  await expect(page.getByTestId('deaths-tab')).toBeVisible();

  // The aura and cast tables keyed their own rows the same way, and the first real wipe
  // left both on "Loading the report."
  await page.goto('/reports/fixture2abcd?fight=3&tab=buffs');
  await expect(page.getByTestId('aura-table')).toBeVisible();
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts');
  await expect(page.getByTestId('cast-table')).toBeVisible();

  expect(errors).toEqual([]);
});

test('a link to a fight the report does not have says so instead of swapping silently', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=99');
  await expect(page.getByTestId('report-missing-fight')).toContainText('no fight 99');
  await expect(page.locator('select[data-testid="report-fight"]')).toHaveValue('3');
  // Picking a fight clears the notice.
  await pickFight(page, 4);
  await expect(page.getByTestId('report-missing-fight')).toHaveCount(0);
});

test('a kill reads as a kill, with per-second figures beside the totals', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await expect(page.locator('select[data-testid="report-fight"] option[value="3"]')).toHaveText(/· Kill ·/);
  await expect(page.locator('select[data-testid="report-fight"] option[value="3"]')).toHaveAttribute(
    'data-description',
    /1 death/,
  );
  // MUST 1: each amount beside its rate.
  // Below 1024 the head row is hidden (ui/StatTable's phone layout), still in the tree.
  const heads = page.getByTestId('summary-tab').getByRole('columnheader', { includeHidden: true });
  await expect(heads.filter({ hasText: 'Damage' }).first()).toBeAttached();
  await expect(heads.filter({ hasText: 'DPS' })).toBeAttached();
});

test('the whole night folds every boss pull into one table, and each pull opens from it', async ({
  page,
}) => {
  await page.goto('/reports/fixture2abcd?fight=all');
  await expect(page.locator('select[data-testid="report-fight"]')).toHaveValue('0');
  await expect(page.getByTestId('night-view')).toBeVisible();
  await expect(page.getByTestId('night-summary')).toContainText('2 boss pulls');
  await expect(page.getByTestId('night-bosses').locator('[data-testid^="night-boss-"]')).toHaveCount(2);
  await expect(
    page.getByTestId('night-players').locator('[data-testid^="night-player-"]').first(),
  ).toBeVisible();
  // The whole night keeps the tables and drops the chart: nothing draws a chart over a night.
  await expect(page.getByTestId('time-chart')).toHaveCount(0);
  await openTab(page, 'damage-done');
  await expect(page.getByTestId('actor-table')).toBeVisible();
  await openTab(page, 'summary');
  await page.getByTestId('night-bosses').getByRole('button', { name: '#1' }).first().click();
  await expect(page.locator('select[data-testid="report-fight"]')).toHaveValue('3');
  await expect(page.getByTestId('summary-tab')).toBeVisible();
});

test('a bare report link opens on the first boss pull, and fight changes enter history', async ({ page }) => {
  await page.goto('/reports/fixture2abcd');
  await expect(page.locator('select[data-testid="report-fight"]')).toHaveValue('3');
  await pickFight(page, 4);
  await expect(page).toHaveURL(/fight=4/);
  await page.goBack();
  await expect(page.locator('select[data-testid="report-fight"]')).toHaveValue('3');
});

test('picking a roster name narrows the page to that player', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  const names = page.getByTestId('roster-name');
  const first = await names.first().innerText();
  await names.first().click();
  await expect(page).toHaveURL(/source=/);
  await expect(page.getByTestId('summary-tab').getByTestId('roster-name')).toHaveCount(1);
  await expect(page.getByTestId('roster-name')).toHaveText(first);
});
