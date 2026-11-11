// web/tests/e2e/handoffs-setup.spec.ts
import { apiUnreachable, expect, test } from './support/test';
import { gotoSim, simCharacter } from './support/sim-page';
import { ACTIVE_BUILD } from './support/active-build';
import { addonCopy } from '../../src/lib/addon/copy';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/builds?*'));

test('the /setup paste box offers the planner and the simulator once an export decodes', async ({ page }) => {
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
    `/sim?source=addon&code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
});

test('a code from another format is refused by name, not silently dropped', async ({ page }) => {
  await page.goto('/setup');
  await page.getByTestId('addon-paste-code').fill('FS2:nope');
  await page.getByTestId('addon-paste-submit').click();
  await expect(page.getByTestId('addon-paste-error')).toHaveText('That code is FS2; this site reads FS1.');
  await expect(page.getByTestId('addon-paste-planner')).toHaveCount(0);
});

// The in-game window's own card retired (U9): the surfaces section is the evidence, with real
// client captures; How the addon works keeps its two flows.
test('How the addon works keeps two flows and never says the addon ships no screenshot', async ({ page }) => {
  await page.goto('/setup');
  const flows = page.getByTestId('setup-flows');
  await expect(flows.getByText(addonCopy.flowOutTitle)).toBeVisible();
  await expect(flows.getByText(addonCopy.flowInTitle)).toBeVisible();
  await expect(page.getByText('no screenshot', { exact: false })).toHaveCount(0);
});

test.describe('What it does in the game (L8.3)', () => {
  for (const width of [1280, 360]) {
    test(`at ${width}px: eleven surfaces in order, every capture lazy, sized, described and loaded once scrolled to`, async ({
      page,
    }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/setup');
      // Below 768 the grid folds behind one button (U5); open from 768.
      if (width < 768) await page.getByTestId('setup-surfaces-show').click();
      const surfaces = page.getByTestId('setup-surfaces').locator('[data-testid^="setup-surface-"]');
      await expect(surfaces).toHaveCount(11);
      expect(
        await surfaces.evaluateAll((els) =>
          els.map((el) => el.getAttribute('data-testid')?.replace('setup-surface-', '')),
        ),
      ).toEqual(['4.1', '4.2', '4.3', '4.4', '4.5', '4.6', '4.7', '4.8', '4.9', '4.10', '4.11']);
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

  test('the commands, as a key-binding list, and the companion handshake sit in the rail', async ({
    page,
  }) => {
    await page.goto('/setup');
    const commands = page.getByTestId('setup-commands');
    await expect(commands.locator('dt')).toHaveCount(addonCopy.commands.length);
    await expect(commands.locator('dt').first()).toHaveText('paste or copy a build');
    await expect(commands.locator('dd code').first()).toHaveText('/fs');
    await expect(commands.locator('dd code').last()).toHaveText('/fs help');
    await expect(page.getByTestId('setup-handshake')).toContainText('Site has your save from');
  });
});

// The point of the whole round: a character given to the site once is the site's current
// character everywhere, with no second paste.
test('an export pasted on /setup becomes the current character on the simulator and its tabs', async ({
  page,
}) => {
  await page.goto('/setup');
  await page
    .getByTestId('addon-paste-code')
    .fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:|who=Simfury:Realm`);
  await page.getByTestId('addon-paste-submit').click();

  // The header selector shows it at once, without a reload (the selector is the one character control).
  await expect(page.getByTestId('selector-name')).toHaveText('Simfury', { timeout: 15_000 });

  // A bare simulator URL (no ?code=) restores it and says so; so does a tool tab.
  await page.goto('/sim');
  await expect(simCharacter(page)).toBeVisible({ timeout: 15_000 });
  await expect(page.getByTestId('page-title')).toHaveText('Simfury');
  await page.goto('/sim/gear');
  await expect(page.getByTestId('selector-name')).toHaveText('Simfury', { timeout: 15_000 });

  // Forget means forgotten everywhere.
  await page.evaluate(() => {
    window.localStorage.removeItem('fs.currentCharacter');
    window.dispatchEvent(new Event('fs:current-character'));
  });
  await gotoSim(page);
  await expect(page.getByTestId('sim-landing')).toBeVisible();
  await expect(simCharacter(page)).toHaveCount(0);
});

// A visitor with no session hint is signed out from the first paint: a loaded export shows the
// sign-in hint at once, and the paste box never shrinks when the account read answers later.
test('a slow session check never moves the loaded paste box', async ({ page }) => {
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
  await page
    .getByTestId('addon-paste-code')
    .fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:|who=Simfury:Realm`);
  await page.getByTestId('addon-paste-submit').click();

  await expect(page.getByTestId('addon-paste-signin-hint')).toBeVisible();
  const box = page.getByTestId('addon-paste-box');
  const heightBefore = await box.boundingBox().then((b) => b?.height);
  const commandsTopBefore = await page
    .getByTestId('setup-commands')
    .boundingBox()
    .then((b) => b?.y);
  expect(heightBefore).toBeDefined();

  resolveMe?.();
  await page.waitForTimeout(300);
  await expect(page.getByTestId('addon-paste-signin-hint')).toBeVisible();
  expect(await box.boundingBox().then((b) => b?.height)).toBeGreaterThanOrEqual(heightBefore as number);
  expect(
    await page
      .getByTestId('setup-commands')
      .boundingBox()
      .then((b) => b?.y),
  ).toBe(commandsTopBefore);
});
