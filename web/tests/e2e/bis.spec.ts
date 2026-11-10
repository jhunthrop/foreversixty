// web/tests/e2e/bis.spec.ts
// The Best in slot page on the in-game paperdoll (design/specs/2026-11-09-directions.md,
// Paperdoll): the index links to a spec page; one island owns faction and band, switched from the
// control bar with no reload; the frame draws a button per filled slot and its centre shows the
// selected slot's client tooltip, ending in what it is worth; the itinerary lists every pick by
// where it drops, each a hover target in the shared tooltip; What to stack carries the top weights
// and sends the rest to the sim; the controls and the sim link fit a 390 px phone. Picks come from
// the published nightly data, so the assertions read shapes and known copy, never a pinned item.
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { unrankedBisRoute } from './support/bis-file';
import { openBis, waitForBis, chooseControl } from './support/bis-page';

/** Set once per page load; a full navigation wipes it, so reading it back proves "no reload". */
async function markNoReload(page: Page): Promise<void> {
  await page.evaluate(() => {
    (window as unknown as { __noReloadMarker?: boolean }).__noReloadMarker = true;
  });
}

async function expectNoReloadSoFar(page: Page): Promise<void> {
  const marker = await page.evaluate(
    () => (window as unknown as { __noReloadMarker?: boolean }).__noReloadMarker,
  );
  expect(marker).toBe(true);
}

test('BiS: index links to a spec, faction and band controls switch the doll with no reload', async ({
  page,
}) => {
  await page.goto('/bis');
  await expect(page.getByRole('heading', { name: 'Best in slot', level: 1 })).toBeVisible();
  const marksmanshipLink = page.locator('a[href^="/bis/hunter/marksmanship"]').first();
  await expect(marksmanshipLink).toBeVisible();
  await marksmanshipLink.click();
  await expect(page).toHaveURL(/\/bis\/hunter\/marksmanship/);
  await waitForBis(page);

  const root = page.getByTestId('bis-page');
  await expect(root).toHaveAttribute('data-bis-faction', 'alliance');
  await expect(page.locator('[data-testid^="bis-band-alliance-"]')).toBeVisible();
  await expect(page.getByRole('heading', { level: 1 })).toHaveText(/^Marksmanship Hunter BiS/);
  // Every filled slot is a button on the frame, named by its slot and item.
  const slots = page.locator('[data-testid^="paperdoll-frame-slot-"]');
  expect(await slots.count()).toBeGreaterThan(8);
  await expect(slots.first()).toHaveAttribute('aria-label', /^[A-Z][a-z ]+: \S/);

  await markNoReload(page);
  await chooseControl(page, 'faction', 'horde');
  await expect(root).toHaveAttribute('data-bis-faction', 'horde');
  await expect(page).toHaveURL(/[?&]faction=horde/);
  await expectNoReloadSoFar(page);

  await chooseControl(page, 'faction', 'alliance');
  await chooseControl(page, 'band', '30');
  const band30 = page.getByTestId('bis-band-alliance-30');
  await expect(band30).toBeVisible();
  await expect(page).toHaveURL(/\/bis\/hunter\/marksmanship\/30-39(\?|$)/);
  await expectNoReloadSoFar(page);
  // Band 30's picks include a quest reward: the itinerary has its Quest stop.
  await expect(band30.getByTestId('bis-itinerary')).toContainText('Quest');
});

test('BiS: a spec with no ranked list yet shows the empty state, not a 404', async ({ page }) => {
  const route = unrankedBisRoute();
  test.skip(route === undefined, 'every written spec has a ranked list');
  const response = await page.goto(route!);
  expect(response?.status()).toBe(200);
  await expect(page.getByTestId('bis-empty-state')).toBeVisible();
  await expect(page.getByTestId('bis-empty-state')).toContainText('No leveling BiS list yet');
});

test('BiS: a slot shows its item in the centre, the itinerary opens the shared tooltip', async ({
  page,
  isMobile,
}) => {
  test.skip(isMobile, 'the frame shows from 423 px; hover is a mouse');
  await openBis(page, '/bis/hunter/beast-mastery');
  const centre = page.locator('[data-region="paperdoll-centre"]');
  const first = page.locator('[data-testid^="paperdoll-frame-slot-"][aria-pressed="true"]');
  const firstItem = (await first.getAttribute('aria-label'))!.split(': ')[1]!;
  await expect(centre.getByTestId('item-tooltip')).toContainText(firstItem);
  // It ends in what the pick is worth, never "+0.0".
  await expect(centre.getByTestId('bis-slot-gain')).toHaveText(/\S/);
  await expect(page.getByTestId('bis-doll')).not.toContainText('+0.0 DPS');

  // Hover another slot: the centre swaps, and no floating tooltip opens over the frame.
  const back = page.getByTestId('paperdoll-frame-slot-back');
  const backItem = (await back.getAttribute('aria-label'))!.split(': ')[1]!;
  await back.hover();
  await expect(back).toHaveAttribute('aria-pressed', 'true');
  await expect(centre.getByTestId('item-tooltip')).toContainText(backItem);
  await expect(page.locator('#bis-item-tooltip')).toHaveCount(0);

  // The itinerary's names open the page's one floating tooltip.
  const row = page.getByTestId('bis-itinerary').locator('[data-testid^="item-hover-"]').first();
  const rowName = (await row.innerText()).trim();
  await row.hover();
  const floating = page.locator('#bis-item-tooltip');
  await expect(floating).toBeVisible();
  await expect(floating).toContainText(rowName);
  await page.mouse.move(0, 0);
  await expect(floating).toBeHidden();
});

test('BiS: a covered off hand says the two-hander takes it', async ({ page, isMobile }) => {
  test.skip(isMobile, 'the frame shows from 423 px');
  await openBis(page, '/bis/hunter/beast-mastery');
  const offHand = page.getByTestId('paperdoll-frame-slot-off_hand');
  test.skip((await offHand.count()) === 0, 'no two-hander on this band with the current data');
  await offHand.click();
  await expect(page.locator('[data-region="paperdoll-centre"]')).toContainText(
    /is a two-hander: the off hand stays empty\./,
  );
});

test('BiS: What to stack lists the top weights, what 1.00 is worth, and sends the rest to the sim', async ({
  page,
}) => {
  await openBis(page, '/bis/hunter/marksmanship?band=60');
  const weights = page.getByTestId('bis-weights');
  await expect(weights).toBeVisible();
  await expect(weights.locator('dt')).toHaveCount(4);
  await expect(weights.locator('dd').first()).toHaveText('1.00');
  await expect(weights.getByTestId('bis-weights-unit')).toHaveText(/^1\.00 = \d\.\d\d\d DPS$/);
  await expect(weights.getByTestId('bis-weights-sim-link')).toHaveAttribute('href', /^\/sim\/weights\?code=/);
  // No other addon is ever named (owner rule, 2026-09-30).
  await expect(page.getByTestId('bis-page')).not.toContainText(/Pawn/);
});

test('BiS: Where to go links each stop of the itinerary', async ({ page }) => {
  await openBis(page, '/bis/hunter/beast-mastery');
  const links = page.getByTestId('bis-where').locator('a[href^="#stop-"]');
  const count = await links.count();
  expect(count).toBeGreaterThan(1);
  await expect(page.getByTestId('bis-itinerary').locator('h2')).toHaveCount(count);
  const target = (await links.first().getAttribute('href'))!.slice(1);
  await expect(page.locator(`#${target}`)).toHaveCount(1);
});

test.describe('the controls at 390', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('show all five bands, 44 px tall, with the selected one checked', async ({ page }) => {
    await page.goto('/bis/mage/fire?band=60');
    await waitForBis(page);
    await page.getByTestId('bis-controls').locator('[data-control-bar-toggle]').click();
    for (const band of [20, 30, 40, 50, 60]) {
      const option = page.getByTestId(`control-band-${band}`);
      await expect(option).toBeInViewport({ ratio: 1 });
      const box = await option.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
      expect((box?.x ?? 0) + (box?.width ?? 0)).toBeLessThanOrEqual(390);
    }
    await expect(page.getByTestId('control-band-60')).toHaveAttribute('aria-checked', 'true');
  });

  test('show the list form with the sim link inside the page', async ({ page }) => {
    await openBis(page, '/bis/warrior/fury');
    await expect(page.locator('.pdf-frame')).toBeHidden();
    await expect(page.locator('[data-testid^="paperdoll-list-slot-"]').first()).toBeVisible();
    const link = page.locator('[data-testid="bis-sim-link"]:visible');
    await expect(link).toHaveCount(1);
    const box = await link.boundingBox();
    expect((box?.x ?? 0) + (box?.width ?? 0)).toBeLessThanOrEqual(390);
    const width = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(width).toBeLessThanOrEqual(390);
  });
});

test('BiS: an item icon that fails to load draws the dark socket, never a broken image', async ({ page }) => {
  await page.route('**/icons/*.webp', (route) => route.fulfill({ status: 404, body: '' }));
  await openBis(page, '/bis/hunter/marksmanship?band=60');
  const doll = page.getByTestId('bis-doll');
  await expect(doll.locator('[data-icon-socket]:visible').first()).toBeVisible();
  await expect(doll.locator('img[data-icon]:visible')).toHaveCount(0);
});
