// web/tests/e2e/report-mechanics.spec.ts
// Mechanics mode against the fixture report: fight 3 is encounter 9001 (Warden Kelthas),
// whose curated table lists Anima Lash as avoidable, and Thalgrit (Player-4184-000000A4)
// is the tank it killed. Fight 4 is encounter 9002, which has no table.
import { expect, test } from '@playwright/test';

const FIGHT = '/reports/fixture2abcd?fight=3';

test('mechanics lists the avoidable hit and the death it caused, and links to the detail', async ({
  page,
}) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  await expect(page.getByTestId('mechanics-mode')).toBeVisible();
  const problems = page.getByTestId('mechanics-problems');
  await expect(problems).toContainText('Anima Lash');
  await expect(problems).toContainText('died to it');
  await problems
    .getByRole('button', { name: /Damage Taken/ })
    .first()
    .click();
  await expect(page).toHaveURL(/tab=damage-taken/);
  await expect(page).toHaveURL(/ability=334660/);
});

test('a player card says what to tell them', async ({ page }) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  const card = page.getByTestId('mechanics-player-Player-4184-000000A4');
  await expect(card).toContainText('Anima Lash');
  await expect(card).toContainText(/avoidable/i);
});

test('a boss without a table says so and offers the draft', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4&mode=mechanics');
  await expect(page.getByTestId('mechanics-no-table')).toContainText('No mechanics table');
});

test('the night rolls mechanics up per pull', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&mode=mechanics');
  await expect(page.getByTestId('mechanics-night')).toContainText('Anima Lash');
  await expect(page.getByTestId('mechanics-night')).toContainText(/pull/);
});
