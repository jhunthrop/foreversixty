// web/tests/e2e/handoffs-report.spec.ts
import { expect, test } from '@playwright/test';

test('each combatant in a report has its own Sim link beside the planner link', async ({ page }) => {
  // Reuses the same fixture report every other report-*.spec.ts file runs against.
  await page.goto('/reports/fixture2abcd?fight=3');
  const combatants = page.getByTestId('combatants').getByRole('listitem');
  await expect(combatants.first()).toBeVisible();
  const first = combatants.first();
  const simLink = first.getByTestId('combatant-sim-link');
  await expect(simLink).toBeVisible();
  const href = await simLink.getAttribute('href');
  expect(href).toMatch(/^\/sim\?source=fight&ref=fixture2abcd%3A3%3A/);
});
