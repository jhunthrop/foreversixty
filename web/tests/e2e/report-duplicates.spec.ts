// web/tests/e2e/report-duplicates.spec.ts
import { readFile } from 'node:fs/promises';
import path from 'node:path';

import { expect, test } from '@playwright/test';

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
  await page.route('**/fights/3/summary.json', (route) =>
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
  await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');
  // Picking a fight clears the notice.
  await page.getByTestId('fight-4').click();
  await expect(page.getByTestId('report-missing-fight')).toHaveCount(0);
});

test('a kill reads as a kill, with per-second figures beside the totals', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3');
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill');
  await expect(page.getByTestId('fight-3-outcome')).toHaveClass(/text-kill/);
  await expect(page.getByTestId('fight-3-deaths')).toContainText('1');
  await expect(page.getByTestId('roster-damage').first()).toContainText('/s');
});

test('the whole night folds every boss pull into one table, and each pull opens from it', async ({
  page,
}) => {
  await page.goto('/reports/fixture2abcd?fight=all');
  await expect(page.getByTestId('fight-all')).toHaveAttribute('aria-current', 'true');
  await expect(page.getByTestId('night-view')).toBeVisible();
  await expect(page.getByTestId('night-summary')).toContainText('2 boss pulls');
  await expect(page.getByTestId('night-bosses').locator('li')).toHaveCount(2);
  await expect(page.getByTestId('night-players').locator('li').first()).toBeVisible();
  // The whole night is not a fight: no tabs, no chart.
  await expect(page.getByTestId('mode-bar')).toHaveCount(0);
  await page.getByTestId('night-bosses').getByRole('button', { name: 'pull 1' }).first().click();
  await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');
  await expect(page.getByTestId('summary-tab')).toBeVisible();
});

test('a bare report link opens on the first boss pull, and fight changes enter history', async ({ page }) => {
  await page.goto('/reports/fixture2abcd');
  await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');
  await page.getByTestId('fight-4').click();
  await expect(page).toHaveURL(/fight=4/);
  await page.goBack();
  await expect(page.getByTestId('fight-3')).toHaveAttribute('aria-current', 'true');
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
