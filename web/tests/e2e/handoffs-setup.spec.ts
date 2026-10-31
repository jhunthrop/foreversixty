// web/tests/e2e/handoffs-setup.spec.ts
import { apiUnreachable, expect, test } from './support/test';
import { ACTIVE_BUILD } from './support/active-build';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/builds?*'));

test('the /setup paste box offers the planner and the simulator once an export decodes', async ({ page }) => {
  await page.goto('/setup');
  await page.getByTestId('setup-try-export').locator('summary').click();
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
    `/sim?source=addon&code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
});

test('a code from another format is refused by name, not silently dropped', async ({ page }) => {
  await page.goto('/setup');
  await page.getByTestId('setup-try-export').locator('summary').click();
  await page.getByTestId('addon-paste-code').fill('FS2:nope');
  await page.getByTestId('addon-paste-submit').click();
  await expect(page.getByTestId('addon-paste-error')).toHaveText('That code is FS2; this site reads FS1.');
  await expect(page.getByTestId('addon-paste-planner')).toHaveCount(0);
});

// Task L8.3 replaced "ships no screenshot of it": the page now shows real client captures,
// one section per surface, every one lazy, sized and described.
test('the page says the in-game UI is in beta testing and points at the surfaces section', async ({
  page,
}) => {
  await page.goto('/setup');
  // The how-it-works cards sit inside the one closed "Everything the addon does" disclosure.
  await page.getByText('Everything the addon does').click();
  await expect(page.getByText('in beta testing in game')).toBeVisible();
  await expect(page.getByText('no screenshot', { exact: false })).toHaveCount(0);
});

test.describe('What it does in the game (L8.3)', () => {
  for (const width of [1280, 360]) {
    test(`at ${width}px: eleven surfaces in order, every capture lazy, sized, described and loaded once scrolled to`, async ({
      page,
    }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/setup');
      await page.getByText('Everything the addon does').click();
      const surfaces = page.getByTestId('setup-surfaces').locator('[data-surface]');
      await expect(surfaces).toHaveCount(11);
      expect(await surfaces.evaluateAll((els) => els.map((el) => el.getAttribute('data-surface')))).toEqual([
        '4.1',
        '4.2',
        '4.3',
        '4.4',
        '4.5',
        '4.6',
        '4.7',
        '4.8',
        '4.9',
        '4.10',
        '4.11',
      ]);
      const images = page.getByTestId('setup-surfaces').locator('img');
      const count = await images.count();
      expect(count).toBe(12);
      for (let i = 0; i < count; i++) {
        const image = images.nth(i);
        await expect(image).toHaveAttribute('loading', 'lazy');
        expect(((await image.getAttribute('alt')) ?? '').length).toBeGreaterThan(20);
        // The box holds the image's own shape before it loads, so nothing below moves.
        const before = await image.boundingBox();
        await image.scrollIntoViewIfNeeded();
        await expect.poll(() => image.evaluate((el: HTMLImageElement) => el.naturalWidth)).toBeGreaterThan(0);
        const after = await image.boundingBox();
        expect(Math.abs((after?.height ?? 0) - (before?.height ?? -1))).toBeLessThan(1);
        expect(after?.width ?? 0).toBeLessThanOrEqual(width);
      }
      // Messaging (4.9) has no capture yet: no image, no placeholder.
      await expect(surfaces.nth(8).locator('img')).toHaveCount(0);
      const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth);
      expect(overflow).toBeLessThanOrEqual(0);
    });
  }

  test('the commands and the companion handshake follow the surfaces', async ({ page }) => {
    await page.goto('/setup');
    await page.getByText('Everything the addon does').click();
    const commands = page.getByTestId('setup-commands').locator('li');
    await expect(commands).toHaveCount(13);
    await expect(commands.first()).toHaveText('/fs: paste or copy a build');
    await expect(commands.last()).toHaveText('/fs help: every command');
    await expect(page.getByTestId('setup-handshake')).toContainText('Site has your save from');
  });
});

// The point of the whole round: a character given to the site once is the site's current
// character everywhere, with no second paste.
test('an export pasted on /setup becomes the current character on the simulator and its tabs', async ({
  page,
}) => {
  await page.goto('/setup');
  await page.getByTestId('setup-try-export').locator('summary').click();
  await page
    .getByTestId('addon-paste-code')
    .fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:|who=Simfury:Realm`);
  await page.getByTestId('addon-paste-submit').click();

  // The bar under the paste box shows it at once, without a reload.
  const chip = page.getByTestId('current-character-chip');
  await expect(chip).toContainText('Simfury');

  // A bare simulator URL (no ?code=) restores it and says so; so does a tool tab.
  await page.goto('/sim');
  await expect(page.getByTestId('sim-character')).toBeVisible({ timeout: 15_000 });
  await page.goto('/sim/gear');
  await expect(page.getByTestId('selector-name')).toHaveText('Simfury', { timeout: 15_000 });

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
// setup.astro immediately after AddonPasteBox), which has a fixed height of its own
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
  await page.getByTestId('setup-try-export').locator('summary').click();
  await page
    .getByTestId('addon-paste-code')
    .fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:|who=Simfury:Realm`);
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
