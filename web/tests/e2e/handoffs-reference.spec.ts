// web/tests/e2e/handoffs-reference.spec.ts
import { expect, test } from './support/test';

// The class landing page lost its generic "Open the planner for this class" link when the
// guides were rebuilt (a bare class link opened an empty planner), and the prose's bare class
// link went in the 2026-11-09 copy sweep. The hand-off is the At a glance plate's and the
// leveling lane's spec-specific "Load this build" and "Load in planner" links
// (guides-redesign.spec.ts), each opening the planner on the guide's own spec.
test('a spec guide links to the planner for its spec', async ({ page }) => {
  await page.goto('/guides/warrior/fury');
  await expect(page.locator('main a[href^="/planner?"][href*="spec=warrior-fury"]').first()).toBeVisible();
});
