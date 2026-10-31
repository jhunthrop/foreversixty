import type { Locator } from '@playwright/test';
import { expect, test } from './support/test';

// design/specs/2026-10-28-polish.md section 2.3: below 768 the four link columns are a 2 x 2
// grid (never one long column) and every link is a 44px row; the Discord button is reachable.
test.use({ viewport: { width: 390, height: 844 } });

const HIT_TARGET_PX = 44;

async function box(locator: Locator): Promise<{ x: number; y: number; width: number; height: number }> {
  const found = await locator.boundingBox();
  if (!found) throw new Error('element has no box');
  return found;
}

test('on a phone the four columns stack two-up with 44px link rows', async ({ page }) => {
  await page.goto('/about');
  const nav = page.getByRole('navigation', { name: 'Footer' });
  const labels = ['Play', 'Logs', 'Guilds', 'About'];
  const heads = await Promise.all(labels.map((label) => box(nav.getByText(label, { exact: true }).first())));
  expect(Math.abs(heads[0].y - heads[1].y)).toBeLessThan(2);
  expect(Math.abs(heads[2].y - heads[3].y)).toBeLessThan(2);
  expect(heads[1].x).toBeGreaterThan(heads[0].x + 100);
  expect(heads[2].y).toBeGreaterThan(heads[0].y + 100);

  const links = await nav.getByRole('link').all();
  expect(links.length).toBe(16);
  for (const link of links) {
    expect((await box(link)).height).toBeGreaterThanOrEqual(HIT_TARGET_PX);
  }
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
