// web/tests/e2e/support/bis-page.ts
// Opening the BiS page and waiting for its island: the head script's pending mark is gone and
// Astro has hydrated the island (its `ssr` attribute drops), so the controls answer clicks.
import { expect, type Page } from '@playwright/test';

export async function openBis(page: Page, url: string): Promise<void> {
  await page.goto(url);
  await waitForBis(page);
}

export async function waitForBis(page: Page): Promise<void> {
  await expect(page.locator('html[data-bis-pending]')).toHaveCount(0);
  await expect(page.locator('astro-island:has([data-testid="bis-page"]):not([ssr])')).toHaveCount(1);
}

/** Chooses a control-bar option, opening the phone's collapsed control stack first when needed. */
export async function chooseControl(page: Page, param: string, value: string): Promise<void> {
  const option = page.getByTestId(`control-${param}-${value}`);
  if (!(await option.isVisible())) {
    await page.getByTestId('bis-controls').locator('[data-control-bar-toggle]').click();
  }
  await option.click();
}
