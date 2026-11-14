// web/tests/e2e/addon-install-link.spec.ts
// The one "Install the addon" link (lib/addon/install-link.ts, owner 2026-11-15) on the gear, talent and
// guide pages: Where to go closes with it on the BiS page, the class guide's head carries it, and a signed-in
// planner's In the game ends with it (signed out, the plate's own addon link stands in).
import { expect, test } from './support/test';
import { openBis } from './support/bis-page';

test('the BiS page closes Where to go with Install the addon', async ({ page }) => {
  await openBis(page, '/bis/hunter/beast-mastery');
  const link = page.getByTestId('bis-where').getByTestId('bis-install-addon');
  await expect(link).toHaveText('Install the addon');
  await expect(link).toHaveAttribute('href', '/setup');
});

test('the class guide carries Install the addon in its head', async ({ page }) => {
  await page.goto('/guides/hunter/beast-mastery');
  const link = page.getByTestId('guide-install-addon');
  await expect(link).toHaveText('Install the addon');
  await expect(link).toHaveAttribute('href', '/setup');
});
