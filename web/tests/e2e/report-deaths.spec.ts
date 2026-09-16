// web/tests/e2e/report-deaths.spec.ts
import { expect, test } from '@playwright/test';

const DEATHS = '/reports/fixture2abcd?fight=3&tab=deaths';

test('the recap shows the killing blow, the last hits and the auras that were up', async ({ page }) => {
  await page.goto(DEATHS);

  const card = page.getByTestId('death-Player-4184-000000A4');
  await expect(card).toContainText('Thalgrit');
  await expect(card).toContainText('killed by Warden Kelthas');
  await expect(card).toContainText('Anima Lash');
  await expect(card).toContainText('100 overkill');
  // Two hits, not three: the boss's swing on the tank is logged only as a
  // SWING_DAMAGE_LANDED, which the engine reads for the target's health and never counts
  // as a hit of its own, so the recap is the two Anima Lash hits.
  await expect(card.locator('tbody tr')).toHaveCount(2);
  await expect(card).toContainText('Necrotic Wound');
});

test('a death sets the window to the twenty seconds before it', async ({ page }) => {
  await page.goto(DEATHS);
  await page.getByTestId('death-window').click();
  await expect(page).toHaveURL(/start=0&end=11000/);
});

test('the combatants list links a build into the planner and the planner opens it', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=summary');

  const link = page.getByTestId('combatant-build-link').first();
  await expect(link).toHaveText('Gear in the planner');
  const href = await link.getAttribute('href');
  expect(decodeURIComponent(href ?? '')).toContain('FS1:');
  expect(decodeURIComponent(href ?? '')).toContain('head=175850');

  await link.click();
  await page.waitForURL('**/planner?code=**');
  // The fixture's tank carries retail spell ids, not per-talent ranks, so the code this
  // link built is gear-only: Planner.svelte says so rather than claiming talents loaded
  // (see the gearOnly branch of its planner-code-note, added for this link).
  await expect(page.getByTestId('planner-code-note')).toContainText('Gear loaded');
});

test('deaths read top to bottom as cards at 360px with 44px controls', async ({ page }, testInfo) => {
  test.skip(testInfo.project.name !== 'mobile', 'phone layout');
  await page.goto(DEATHS);
  const button = (await page.getByTestId('death-window').boundingBox())!;
  expect(button.height).toBeGreaterThanOrEqual(44);
  expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
});
