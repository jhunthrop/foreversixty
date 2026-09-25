// web/tests/e2e/handoffs-reference.spec.ts
import { expect, test } from '@playwright/test';

test('a class guide links to the planner for that class', async ({ page }) => {
  await page.goto('/guides/warrior');
  const link = page.getByRole('link', { name: 'Open the planner for this class' });
  await expect(link).toHaveAttribute('href', '/planner?class=warrior');
});
