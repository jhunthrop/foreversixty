// web/tests/e2e/handoffs-reference.spec.ts
import { expect, test } from './support/test';

// The class landing page lost its generic "Open the planner for this class" link when the
// guides were rebuilt (a bare class link opened an empty planner). The hand-off now lives in
// each spec guide's own prose, and the At a glance plate and the leveling lane carry the
// spec-specific "Load this build" and "Load in planner" links (guides-redesign.spec.ts).
test('a spec guide links to the planner for its class', async ({ page }) => {
  await page.goto('/guides/warrior/fury');
  await expect(page.locator('main a[href="/planner?class=warrior"]').first()).toBeVisible();
});
