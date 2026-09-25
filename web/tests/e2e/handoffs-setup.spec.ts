// web/tests/e2e/handoffs-setup.spec.ts
import { expect, test } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';

test('the /addon paste box offers the planner and the simulator once an export decodes', async ({ page }) => {
  await page.goto('/setup');
  await page.getByTestId('addon-paste-code').fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`);
  await page.getByTestId('addon-paste-submit').click();

  await expect(page.getByTestId('addon-paste-error')).toHaveCount(0);
  const planner = page.getByTestId('addon-paste-planner');
  const sim = page.getByTestId('addon-paste-sim');
  await expect(planner).toHaveAttribute(
    'href',
    `/planner?code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
  await expect(sim).toHaveAttribute(
    'href',
    `/sim?code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
});

test('a code from another format is refused by name, not silently dropped', async ({ page }) => {
  await page.goto('/setup');
  await page.getByTestId('addon-paste-code').fill('FS2:nope');
  await page.getByTestId('addon-paste-submit').click();
  await expect(page.getByTestId('addon-paste-error')).toHaveText('That code is FS2; this site reads FS1.');
  await expect(page.getByTestId('addon-paste-planner')).toHaveCount(0);
});

test('the page says the in-game UI is in beta testing and ships no screenshot of it', async ({ page }) => {
  await page.goto('/setup');
  await expect(page.getByText('in beta testing in game')).toBeVisible();
  const images = await page.locator('main img').count();
  expect(images).toBe(0);
});

// The point of the whole round: a character given to the site once is the site's current
// character everywhere, with no second paste.
test('an export pasted on /addon becomes the current character on the simulator and its tabs', async ({
  page,
}) => {
  await page.goto('/setup');
  await page.getByTestId('addon-paste-code').fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`);
  await page.getByTestId('addon-paste-submit').click();

  // The bar under the paste box shows it at once, without a reload.
  const chip = page.getByTestId('current-character-chip');
  await expect(chip).toContainText('Warrior');

  // A bare simulator URL (no ?code=) restores it and says so; so does a tool tab.
  await page.goto('/sim');
  await expect(page.getByTestId('sim-character')).toBeVisible({ timeout: 15_000 });
  await page.goto('/sim/gear');
  await expect(page.getByTestId('current-character-chip')).toContainText('Warrior', { timeout: 15_000 });

  // Forget means forgotten everywhere.
  await page.goto('/setup');
  await page
    .getByTestId('current-character-chip')
    .getByRole('button', { name: /forget/i })
    .click();
  await page.goto('/sim');
  await expect(page.getByTestId('sim-character')).toHaveCount(0);
});

// The planner/sim links (data-testid="addon-paste-planner") sit ABOVE the reserved
// skeleton/hint block in the DOM (AddonPasteBox.svelte), so nothing below them moving can
// ever move their own boundingBox, which made an earlier version of this test
// tautological: it could not fail even when the reservation was sized wrong. This
// version instead measures two things that DO
// move when the reserved slot's height changes: the addon-paste-box section's own total
// height, and the top of current-character-chip (CurrentCharacterBar.svelte, rendered by
// addon.astro immediately after AddonPasteBox), which has a fixed height of its own
// (CHIP_HEIGHT) and so only shifts when something above it does.
test('a pending session check reserves the signed-out branch height, so the section and the chip below it never shrink once it resolves', async ({
  page,
}) => {
  let resolveMe: (() => void) | undefined;
  await page.route('**/v1/me', (route) => {
    void new Promise<void>((resolve) => {
      resolveMe = resolve;
    }).then(() =>
      route.fulfill({
        status: 401,
        contentType: 'application/json',
        body: '{"ok":false,"error":"unauthorized"}',
      }),
    );
  });
  await page.goto('/setup');
  await page.getByTestId('addon-paste-code').fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`);
  await page.getByTestId('addon-paste-submit').click();

  // The pasted export is already the current character (writeCurrent() runs synchronously
  // in submit(), independent of the /v1/me held pending below), so the chip is already on
  // the page while the skeleton is still showing.
  const chip = page.getByTestId('current-character-chip');
  await expect(chip).toBeVisible();
  await expect(page.getByTestId('addon-paste-status-skeleton')).toBeVisible();
  const sectionHeightBefore = await page
    .getByTestId('addon-paste-box')
    .boundingBox()
    .then((box) => box?.height);
  const chipTopBefore = await chip.boundingBox().then((box) => box?.y);
  expect(sectionHeightBefore).toBeDefined();
  expect(chipTopBefore).toBeDefined();

  resolveMe?.();
  await expect(page.getByTestId('addon-paste-signin-hint')).toBeVisible();
  const sectionHeightAfter = await page
    .getByTestId('addon-paste-box')
    .boundingBox()
    .then((box) => box?.height);
  const chipTopAfter = await chip.boundingBox().then((box) => box?.y);

  // Before the fix, ADDON_PASTE_STATUS_MIN_H reserved the taller signed-in form's height
  // (168px) for every pending check, so landing on the much shorter signed-out hint (this
  // test's route always answers 401) shrank the section by ~148.5px and pulled the chip up
  // by the same amount -- these two assertions catch exactly that regression.
  expect(sectionHeightAfter).toBeGreaterThanOrEqual(sectionHeightBefore as number);
  expect(chipTopAfter).toBeGreaterThanOrEqual(chipTopBefore as number);
});
