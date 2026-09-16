// web/tests/e2e/report-mechanics.spec.ts
// Mechanics mode against the fixture report: fight 3 is encounter 9001 (Warden Kelthas),
// whose curated table lists one ability of every kind. Anima Lash is avoidable and killed
// Thalgrit (Player-4184-000000A4, the tank); Anima Surge is interruptible and went through;
// Wrack Soul is dispellable and ran its course; Anima Cascade never went out and Necrotic
// Wound is unavoidable, so both sit under "also in the table". Fight 4 is encounter 9002,
// which has no table at all.
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
  // By its own accessible name, not by position: every problem line has a link, and a
  // column of buttons that all read the same names none of them.
  await problems.getByRole('link', { name: 'Damage Taken for Thalgrit, Anima Lash' }).click();
  await expect(page).toHaveURL(/tab=damage-taken/);
  await expect(page).toHaveURL(/ability=334660/);
});

test('the death outranks the interrupt and the dispel, and each links to its own tab', async ({ page }) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  const lines = page.getByTestId('mechanics-problems').getByRole('listitem');
  await expect(lines).toHaveCount(3);
  // A death sorts above any count, whatever the counts are.
  await expect(lines.nth(0)).toContainText('died to it');
  await expect(lines.nth(1)).toContainText('Anima Surge went through 1 of 1 casts');
  await expect(lines.nth(2)).toContainText('Wrack Soul ran its course 1 of 1 times it landed');

  await expect(lines.nth(2).getByRole('link', { name: 'Dispels for Wrack Soul' })).toBeVisible();
  await lines.nth(1).getByRole('link', { name: 'Interrupts for Anima Surge' }).click();
  await expect(page).toHaveURL(/tab=interrupts/);
});

test('the mechanics nobody failed are listed rather than hidden', async ({ page }) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  const clean = page.getByTestId('mechanics-clean');
  await expect(clean).toContainText('Anima Cascade');
  await expect(clean).toContainText('Necrotic Wound');
  await expect(clean).toContainText('unavoidable, the fight’s own damage');
});

test('a player card says what to tell them: what hit them, and what did not', async ({ page }) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  const tank = page.getByTestId('mechanics-player-Player-4184-000000A4');
  await expect(tank).toContainText('Anima Lash');
  await expect(tank).toContainText(/avoidable/i);
  await expect(tank).toContainText('Never hit by: Anima Cascade.');

  // A player nothing landed on gets both halves too: the clean line, and the list of
  // what they stayed out of.
  const spared = page.getByTestId('mechanics-player-Player-4184-000000A1');
  await expect(spared).toContainText('Clean: nothing avoidable landed on them.');
  await expect(spared).toContainText('Never hit by: Anima Lash, Anima Cascade.');
});

test('a boss without a table says so and offers the draft', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=4&mode=mechanics');
  await expect(page.getByTestId('mechanics-no-table')).toContainText('No mechanics table');
});

test('the night rolls mechanics up per boss, out of that boss’s pulls', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=all&mode=mechanics');
  const night = page.getByTestId('mechanics-night');
  // Grouped under the boss whose table lists them, and counted out of that boss's own
  // pulls: the night folds Warden Kelthas and Skolex, and one denominator for both
  // would be the wrong number on every line.
  await expect(night).toContainText('Warden Kelthas');
  await expect(night).toContainText('Anima Lash');
  await expect(night).toContainText('hit someone on 1 of 1 pull');
  await expect(night).toContainText('most often Thalgrit');
  // The problems list says which boss each failure was on, since two bosses can list
  // one spell id and the night shows them side by side.
  await expect(page.getByTestId('mechanics-problems')).toContainText('on Warden Kelthas');
});

test('nothing is listed as unclassified when the table covers everything that hit a player', async ({
  page,
}) => {
  await page.goto(`${FIGHT}&mode=mechanics`);
  await expect(page.getByTestId('mechanics-mode')).toBeVisible();
  // The Warden's table lists the only ability that landed on a player, so the honesty
  // section has nothing to confess. It appears only when the table is silent about
  // something that hit someone.
  await expect(page.getByTestId('mechanics-unclassified')).toHaveCount(0);
});
