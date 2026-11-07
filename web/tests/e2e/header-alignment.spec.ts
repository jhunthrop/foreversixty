// W.12 ruling 9: the planner and tier-list headers start on the body's left edge (48 px from 768 px,
// 18 px on a phone), not at the viewport's. A scoped `padding` that named the phone inset alone had
// zeroed the site column's gutter under both headers.
import { expect, test } from './support/test';

const WIDTHS = [1280, 1024, 768] as const;

async function left(locator: ReturnType<import('@playwright/test').Page['locator']>): Promise<number> {
  const box = await locator.first().boundingBox();
  if (box === null) throw new Error('not rendered');
  return Math.round(box.x);
}

test.describe('headers align with the body', () => {
  test.skip(({ isMobile }) => isMobile, 'Widths are set by the test');

  for (const width of WIDTHS) {
    // Spec 2026-11-08 4.1: the head's answer (the old facts) sits at the right from 900, so the
    // crest is the head's left edge; the body's is the trees' Plate (the tree-source line it was
    // measured by now sits inside the Plate's foot).
    test(`/planner at ${width}: the crest starts where the trees do`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/planner');
      const crest = page.getByTestId('planner-header-crest');
      await expect(crest).toBeVisible();
      await expect(page.getByTestId('planner-trees')).toBeVisible();
      const body = await left(page.getByTestId('planner-trees'));
      expect(body).toBe(48);
      expect(await left(crest)).toBe(body);
    });

    test(`/tiers at ${width}: the title and the plate start where the body does`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await page.goto('/tiers');
      const body = await left(page.locator('[data-testid="tier-faction-panel-alliance"]'));
      expect(body).toBe(48);
      expect(await left(page.getByTestId('tier-page').locator('h1'))).toBe(body);
      expect(await left(page.getByTestId('tier-page').getByTestId('tier-plate'))).toBe(body);
    });
  }

  test('on a phone both headers keep the 18 px inset', async ({ page }) => {
    await page.setViewportSize({ width: 360, height: 800 });
    await page.goto('/planner');
    await expect(page.getByTestId('planner-header-crest')).toBeVisible();
    expect(await left(page.getByTestId('planner-header-crest'))).toBe(18);
    await page.goto('/tiers');
    expect(await left(page.getByTestId('tier-page').locator('h1'))).toBe(18);
  });
});
