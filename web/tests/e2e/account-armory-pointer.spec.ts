// web/tests/e2e/account-armory-pointer.spec.ts
// A current-character pointer written by a character page ('armory' source) used to send
// CurrentCharacterBar's effect into an infinite loop on /account (it read the state it had
// just written), which starved every other island: Your reports never left its skeleton.
import { expect, test } from '@playwright/test';
import { meBnetFixture } from '../../src/fixtures/me-bnet';

test('an armory current-character pointer does not stall the account page', async ({ page }) => {
  const errors: string[] = [];
  page.on('pageerror', (error) => errors.push(error.message));
  await page.route('**/v1/me', (route) =>
    route.fulfill({ json: { ok: true, data: meBnetFixture, error: null } }),
  );
  await page.route('**/v1/devices', (route) => route.fulfill({ json: { ok: true, data: [], error: null } }));
  await page.route('**/v1/reports**', (route) =>
    route.fulfill({ json: { ok: true, data: { rows: [], total: 0, page: 1, per_page: 20 }, error: null } }),
  );
  await page.addInitScript((key) => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: key,
        label: 'Thoradin · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  }, meBnetFixture.characters[0].key);
  await page.goto('/account');
  await expect(page.getByTestId('account-characters')).toBeVisible();
  await expect(page.getByTestId('my-reports-empty')).toBeVisible();
  await expect(page.getByTestId('my-reports-skeleton')).toHaveCount(0);
  expect(errors.filter((m) => m.includes('effect_update_depth_exceeded'))).toEqual([]);
});
