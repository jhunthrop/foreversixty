// web/tests/e2e/bis-doll-keyboard.spec.ts
// The paperdoll is a listbox on the keyboard: one tab stop, arrows move between slots (empty
// ones skipped), Enter opens the slot in place below 1760 and closes it again; from 1760 the
// arrows move the side inspector with the focus.
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';

const URL = '/bis/hunter/marksmanship?band=60';
const focusedSlot = (page: import('@playwright/test').Page) =>
  page.evaluate(() => (document.activeElement as HTMLElement | null)?.dataset.dollSlot ?? null);

test('below 1760: one tab stop, arrows move, Enter opens and closes the slot in place', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  await openBis(page, URL);
  const doll = page.getByTestId('bis-doll');
  await expect(doll.locator('[data-doll-slot][tabindex="0"]')).toHaveCount(1);
  await doll.locator('[data-doll-slot][tabindex="0"]').focus();
  const first = await focusedSlot(page);
  await page.keyboard.press('ArrowDown');
  const second = await focusedSlot(page);
  expect(second).not.toBeNull();
  expect(second).not.toBe(first);
  await page.keyboard.press('Enter');
  await expect(page.getByTestId(`bis-doll-slot-${second}`)).toHaveAttribute('aria-selected', 'true');
  await expect(page.getByTestId(`bis-doll-detail-${second}`)).toBeVisible();
  await page.keyboard.press('Enter');
  await expect(page.getByTestId(`bis-doll-detail-${second}`)).toHaveCount(0);
  await page.keyboard.press('End');
  expect(await focusedSlot(page)).not.toBe(second);
});

test('from 1760: the side inspector follows the arrow keys', async ({ page }) => {
  await page.setViewportSize({ width: 1920, height: 1100 });
  await openBis(page, URL);
  const inspector = page.getByTestId('bis-doll-inspector');
  await expect(inspector).toBeVisible();
  await page.getByTestId('bis-doll').locator('[data-doll-slot][tabindex="0"]').focus();
  const before = await inspector.innerText();
  await page.keyboard.press('ArrowDown');
  const slot = await focusedSlot(page);
  await expect(page.getByTestId(`bis-doll-slot-${slot}`)).toHaveAttribute('aria-selected', 'true');
  await expect(inspector).not.toHaveText(before);
});
