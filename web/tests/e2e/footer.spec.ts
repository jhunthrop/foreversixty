import type { Locator } from '@playwright/test';
import { expect, test } from './support/test';

// Round 5's compact footer (design/specs/2026-11-09-directions.md, coordinator ruling 2026-11-09):
// on a phone every site link is a 44px row target and the legal links sit above the legal lines.
test.use({ viewport: { width: 390, height: 844 } });

const HIT_TARGET_PX = 44;

async function box(locator: Locator): Promise<{ x: number; y: number; width: number; height: number }> {
  const found = await locator.boundingBox();
  if (!found) throw new Error('element has no box');
  return found;
}

test('on a phone the compact footer carries its four site links as 44px targets', async ({ page }) => {
  await page.goto('/about');
  const nav = page.getByRole('navigation', { name: 'Footer' });
  await expect(nav.getByRole('link')).toHaveText([
    'Get set up',
    'Premium',
    'Join the Discord',
    'Report a bug',
  ]);
  for (const link of await nav.getByRole('link').all()) {
    expect((await box(link)).height).toBeGreaterThanOrEqual(HIT_TARGET_PX);
  }
  const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth);
  expect(overflow).toBeLessThanOrEqual(0);
});

test('the legal links and the buttons clear 44px and sit above the legal sentences', async ({ page }) => {
  await page.goto('/about');
  const footer = page.getByRole('contentinfo');
  for (const [label, href] of [
    ['Terms', '/terms'],
    ['Privacy', '/privacy'],
    ['Refunds', '/refunds'],
  ] as const) {
    const link = footer.getByRole('link', { name: label, exact: true });
    await expect(link).toHaveAttribute('href', href);
    const linkBox = await box(link);
    expect(linkBox.height).toBeGreaterThanOrEqual(HIT_TARGET_PX);
    expect(linkBox.width).toBeGreaterThanOrEqual(HIT_TARGET_PX);
  }
  const discord = footer.getByRole('link', { name: 'Join the Discord' });
  expect((await box(discord)).height).toBeGreaterThanOrEqual(HIT_TARGET_PX);
  const termsY = (await box(footer.getByRole('link', { name: 'Terms', exact: true }))).y;
  const trademarkY = (await box(footer.getByText('Not affiliated with'))).y;
  expect(termsY).toBeLessThan(trademarkY);
});
