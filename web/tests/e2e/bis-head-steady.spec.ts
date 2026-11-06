// web/tests/e2e/bis-head-steady.spec.ts
// Switching spec within a class must not move the page head: the answer block on the right is a
// different height per spec (one DPS figure, a tank's four, a healer's metrics row), so the head
// has a fixed floor and the title and control bar keep the same y.
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';

const WIDTHS = [1440, 2000] as const;
const CLASS_SPECS: readonly (readonly string[])[] = [
  ['paladin/retribution', 'paladin/protection', 'paladin/holy'],
  ['priest/holy', 'priest/shadow'],
];

async function controlBarTop(page: Page, spec: string): Promise<number> {
  await openBis(page, `/bis/${spec}?band=60`);
  const box = await page.getByTestId('bis-controls').boundingBox();
  expect(box).not.toBeNull();
  return box!.y;
}

for (const width of WIDTHS) {
  for (const specs of CLASS_SPECS) {
    test(`at ${width}px the control bar keeps its y across ${specs.join(', ')}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 1000 });
      const tops: number[] = [];
      for (const spec of specs) tops.push(await controlBarTop(page, spec));
      expect(new Set(tops).size).toBe(1);
    });
  }
}
