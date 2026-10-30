// W.12 ruling 9: the planner and tier-list headers start on the body's left edge (48 px from 768 px,
// 18 px on a phone), not at the viewport's. A scoped `padding` that named the phone inset alone had
// zeroed the site column's gutter under both headers.
import { expect, test } from '@playwright/test';

const WIDTHS = [1280, 1024, 768] as const;

async function left(locator: ReturnType<import('@playwright/test').Page['locator']>): Promise<number> {
  const box = await locator.first().boundingBox();
  if (box === null) throw new Error('not rendered');
  return Math.round(box.x);
}

test.describe('headers align with the body', () => {
  test.skip(({ isMobile }) => isMobile, 'Widths are set by the test');

  for (const width of WIDTHS) {
    test(`/planner at ${width}: the crest and the facts start where the trees do`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/planner');
      const eyebrow = page.getByTestId('planner-header-crest');
      await expect(eyebrow).toBeVisible();
      await expect(page.getByTestId('planner-tree-source')).toBeVisible();
      const body = await left(page.getByTestId('planner-tree-source'));
      expect(body).toBe(48);
      expect(await left(eyebrow)).toBe(body);
      expect(await left(page.getByTestId('planner-facts'))).toBe(body);
    });

    test(`/tiers at ${width}: the eyebrow and the title start where the body does`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/tiers');
      const body = await left(page.locator('.tier-body'));
      expect(body).toBe(48);
      expect(await left(page.getByTestId('tier-page').locator('p.text-gold'))).toBe(body);
      expect(await left(page.getByTestId('tier-page').locator('h1'))).toBe(body);
    });
  }

  test('on a phone both headers keep the 18 px inset', async ({ page }) => {
    await page.setViewportSize({ width: 360, height: 800 });
    await page.goto('/planner');
    await expect(page.getByTestId('planner-header-crest')).toBeVisible();
    expect(await left(page.getByTestId('planner-header-crest'))).toBe(18);
    await page.goto('/tiers');
    expect(await left(page.getByTestId('tier-page').locator('p.text-gold'))).toBe(18);
  });
});
