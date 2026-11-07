// web/tests/e2e/handoffs-report.spec.ts
import { expect, test } from './support/test';

test('each combatant at the pull has its own Sim link beside its gear', async ({ page }) => {
  // Reuses the same fixture report every other report-*.spec.ts file runs against. The Summary
  // rail's At pull (logs spec 2026-11-07 6.2): Gear › and Sim › per player.
  await page.goto('/reports/fixture2abcd?fight=3');
  const first = page.getByTestId('at-pull-card').getByTestId('at-pull-links').first();
  await expect(first).toBeVisible();
  const simLink = first.getByTestId('combatant-sim-link');
  await expect(simLink).toBeVisible();
  const href = await simLink.getAttribute('href');
  expect(href).toMatch(/^\/sim\?source=fight&ref=fixture2abcd%3A3%3A/);
});
