// web/tests/e2e/planner-talent-tip.spec.ts
// Owner report 2026-10-31: hovering a talent in the right-hand column of a tree cut its
// tooltip off at the tree panel's right edge, so the rail beside it ("Improved T...") seemed
// to paint over the rest. The tooltip must be whole and on top wherever the cell sits: inside
// the viewport, and the topmost element at its own right and bottom edges.
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';

const WIDTHS = [1024, 1440, 2000];
const EDGE_INSET_PX = 6;

/** The test ids of the right-most cell in the top row and in the bottom row of the page's
 *  right-most visible tree: the two cells whose tooltips reach farthest right and down. */
async function rightColumnCells(page: Page): Promise<string[]> {
  return page.evaluate(() => {
    const cells = [...document.querySelectorAll<HTMLElement>('[data-testid^="talent-"][data-rank]')]
      .map((el) => ({ id: el.dataset.testid!, rect: el.getBoundingClientRect() }))
      .filter(({ rect }) => rect.width > 0);
    const right = Math.max(...cells.map(({ rect }) => rect.right));
    const column = cells.filter(({ rect }) => Math.abs(rect.right - right) < 2);
    column.sort((a, b) => a.rect.top - b.rect.top);
    return [...new Set([column[0].id, column[column.length - 1].id])];
  });
}

async function expectWholeAndOnTop(page: Page, tip: Locator): Promise<void> {
  const box = await tip.boundingBox();
  expect(box).not.toBeNull();
  const viewport = page.viewportSize()!;
  expect(box!.x).toBeGreaterThanOrEqual(0);
  expect(box!.y).toBeGreaterThanOrEqual(0);
  expect(box!.x + box!.width).toBeLessThanOrEqual(viewport.width);
  expect(box!.y + box!.height).toBeLessThanOrEqual(viewport.height);
  const probes = [
    { x: box!.x + box!.width - EDGE_INSET_PX, y: box!.y + box!.height / 2 },
    { x: box!.x + box!.width / 2, y: box!.y + box!.height - EDGE_INSET_PX },
    { x: box!.x + box!.width - EDGE_INSET_PX, y: box!.y + box!.height - EDGE_INSET_PX },
  ];
  for (const probe of probes) {
    const onTip = await tip.evaluate((el, { x, y }) => el.contains(document.elementFromPoint(x, y)), probe);
    expect(onTip, `the tooltip is the topmost element at (${probe.x}, ${probe.y})`).toBe(true);
  }
}

test.describe('a talent tooltip is never clipped by its tree panel', () => {
  for (const width of WIDTHS) {
    test(`right-column talents at ${width}px`, async ({ page }, testInfo) => {
      test.skip(testInfo.project.name !== 'desktop', 'the phone project has its own width');
      await page.setViewportSize({ width, height: 1000 });
      await page.goto('/planner');
      await page.getByTestId('band-compare').waitFor();
      for (const id of await rightColumnCells(page)) {
        const cell = page.getByTestId(id);
        await cell.evaluate((el) => el.scrollIntoView({ block: 'center' }));
        await cell.hover();
        const tip = page.getByRole('tooltip');
        await expect(tip).toBeVisible();
        await expectWholeAndOnTop(page, tip);
        await expect(cell).toHaveAttribute('aria-describedby', (await tip.getAttribute('id'))!);
        await page.mouse.move(0, 0);
        await expect(tip).toBeHidden();
      }
    });
  }

  test('a phone tooltip stays inside the screen and never scrolls the page sideways', async ({
    page,
  }, testInfo) => {
    test.skip(testInfo.project.name !== 'mobile', 'the desktop project covers the wide widths');
    await page.setViewportSize({ width: 360, height: 780 });
    await page.goto('/planner');
    await page.getByTestId('band-compare').waitFor();
    for (const id of await rightColumnCells(page)) {
      const cell = page.getByTestId(id);
      await cell.evaluate((el) => el.scrollIntoView({ block: 'center' }));
      await cell.focus();
      const tip = page.getByRole('tooltip');
      await expect(tip).toBeVisible();
      await expectWholeAndOnTop(page, tip);
      const { scrollWidth, clientWidth } = await page.evaluate(() => ({
        scrollWidth: document.documentElement.scrollWidth,
        clientWidth: document.documentElement.clientWidth,
      }));
      expect(scrollWidth).toBe(clientWidth);
      await page.keyboard.press('Escape');
      await expect(tip).toBeHidden();
      await expect(cell).toBeFocused();
    }
  });
});
