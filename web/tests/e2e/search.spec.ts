import { expect, test } from '@playwright/test';

test('search finds a dungeon page and navigates to it', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('/');
  const box = page.getByRole('searchbox');
  await expect(box).toBeFocused();
  await box.fill('Thanes');
  const result = page.getByRole('option', { name: /Hall of Thanes/ });
  await expect(result).toBeVisible();
  await result.click();
  await expect(page).toHaveURL(/\/dungeons\/hall-of-thanes$/);
});
