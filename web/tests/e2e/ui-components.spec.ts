// web/tests/e2e/ui-components.spec.ts
// The shared component set on /design-sheet (polish spec 4.2): axe on every state, the switch
// operated with Space and Enter and announcing its state, a disabled switch out of the tab order,
// 44 px hit areas, and the settings row's phone layout (the value under the label below 480 px).
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';

const SHEET = '[data-testid="ui-sheet"]';

async function openSheet(page: Page): Promise<void> {
  await gotoHydrated(page, '/design-sheet', 'sheet-switch-live');
}

test.describe('component sheet', () => {
  test.use({ reducedMotion: 'reduce' });

  test('axe finds nothing in any state', async ({ page }) => {
    await openSheet(page);
    await expectNoAxeViolations(page, SHEET);
  });

  test('a switch announces its state and toggles with Space', async ({ page }) => {
    await openSheet(page);
    const sw = page.getByTestId('sheet-switch-live');
    await expect(sw).toHaveAttribute('role', 'switch');
    await expect(sw).toHaveAttribute('aria-checked', 'true');
    await sw.focus();
    await page.keyboard.press('Space');
    await expect(sw).toHaveAttribute('aria-checked', 'false');
    await page.keyboard.press('Space');
    await expect(sw).toHaveAttribute('aria-checked', 'true');
  });

  test('a switch also toggles with Enter and a click', async ({ page }) => {
    await openSheet(page);
    const sw = page.getByTestId('sheet-switch-live');
    await sw.focus();
    await page.keyboard.press('Enter');
    await expect(sw).toHaveAttribute('aria-checked', 'false');
    await sw.click();
    await expect(sw).toHaveAttribute('aria-checked', 'true');
  });

  test('a switch takes its name from its row label', async ({ page }) => {
    await openSheet(page);
    await expect(page.getByRole('switch', { name: 'A server sim finishes' })).toBeVisible();
    await expect(page.getByRole('switch', { name: 'Include private logs' })).toBeVisible();
  });

  test('a disabled switch cannot take focus and a saving one ignores a press', async ({ page }) => {
    await openSheet(page);
    const disabled = page.getByRole('switch', { name: 'Disabled' });
    await expect(disabled).toBeDisabled();
    const saving = page.getByRole('switch', { name: 'Saving' });
    await expect(saving).toHaveAttribute('aria-busy', 'true');
    await saving.click();
    await expect(saving).toHaveAttribute('aria-checked', 'true');
  });

  test('every switch has a 44 px hit area', async ({ page }) => {
    await openSheet(page);
    for (const sw of await page.getByRole('switch').all()) {
      const box = await sw.boundingBox();
      expect(box?.width).toBeGreaterThanOrEqual(44);
      expect(box?.height).toBeGreaterThanOrEqual(44);
    }
  });

  test('every button is at least 44 px tall below 1024', async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    await openSheet(page);
    for (const control of await page.locator('[data-testid^="sheet-buttons-"] button').all()) {
      const box = await control.boundingBox();
      expect(box?.height).toBeGreaterThanOrEqual(44);
    }
  });

  test('a link row is one 56 px target with a chevron', async ({ page }) => {
    await openSheet(page);
    const row = page.locator('[data-testid="sheet-rows"] a').first();
    const box = await row.boundingBox();
    expect(box?.height).toBeGreaterThanOrEqual(56);
    await expect(row.locator('[data-glyph="chevron"]')).toBeVisible();
  });

  test('below 480 px the value sits under its label, never beside it', async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    await openSheet(page);
    const row = page.locator('[data-testid="sheet-rows"] a').nth(1);
    const label = await row.locator('[id$="-label"]').boundingBox();
    const value = await row.getByTestId('settings-row-value').boundingBox();
    expect(value?.y).toBeGreaterThanOrEqual((label?.y ?? 0) + (label?.height ?? 0) - 1);
    const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth);
    expect(overflow).toBeLessThanOrEqual(0);
  });

  test('keyboard focus draws a 2 px gold ring on a primary button and a switch', async ({ page }) => {
    await openSheet(page);
    const targets = [
      page.locator('[data-testid="sheet-buttons-primary"] button').first(),
      page.getByTestId('sheet-switch-on'),
    ];
    for (const target of targets) {
      await target.focus();
      await page.keyboard.press('Shift+Tab');
      await page.keyboard.press('Tab');
      await expect(target).toBeFocused();
      const ring = await target.evaluate((el) => {
        const style = getComputedStyle(el);
        return { width: style.outlineWidth, style: style.outlineStyle };
      });
      expect(ring).toEqual({ width: '2px', style: 'solid' });
    }
  });
});
