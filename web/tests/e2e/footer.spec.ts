import type { Locator } from '@playwright/test';
import { expect, test } from './support/test';

// Round 5's one-row footer (the mock, design/mocks/directions/png/g-home-1440.png): on a phone the
// row wraps, every link a 44px target, the fan-made line under the links, nothing overflowing.
test.use({ viewport: { width: 390, height: 844 } });

const HIT_TARGET_PX = 44;

async function box(locator: Locator): Promise<{ x: number; y: number; width: number; height: number }> {
  const found = await locator.boundingBox();
  if (!found) throw new Error('element has no box');
  return found;
}

test('on a phone the footer carries its seven links as 44px targets, the fan-made line under them', async ({
  page,
}) => {
  await page.goto('/about');
  const nav = page.getByRole('navigation', { name: 'Footer' });
  await expect(nav.getByRole('link')).toHaveText([
    'Get set up',
    'Addon',
    'Discord',
    'Premium',
    'Creators',
    'Report a bug',
    'Privacy',
    'Terms',
  ]);
  for (const link of await nav.getByRole('link').all()) {
    expect((await box(link)).height).toBeGreaterThanOrEqual(HIT_TARGET_PX);
  }
  await expect(nav.getByRole('link', { name: 'Terms', exact: true })).toHaveAttribute('href', '/terms');
  await expect(nav.getByRole('link', { name: 'Privacy', exact: true })).toHaveAttribute('href', '/privacy');
  const footer = page.getByRole('contentinfo');
  const line = footer.getByText('A fan-made player tool. Not affiliated with Blizzard.');
  const lastLink = await box(nav.getByRole('link').last());
  expect((await box(line)).y).toBeGreaterThan(lastLink.y);
  const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth);
  expect(overflow).toBeLessThanOrEqual(0);
});

test('/about and /terms carry the full trademark sentence and the operator line', async ({ page }) => {
  for (const path of ['/about', '/terms']) {
    await page.goto(path);
    const main = page.getByRole('main');
    await expect(main.getByText(/World of Warcraft and Warcraft are trademarks of Blizzard/)).toBeVisible();
    await expect(main.getByText(/registered trade name of Commish LLC/)).toBeVisible();
  }
});
