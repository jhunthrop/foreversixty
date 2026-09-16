import { expect, test } from '@playwright/test';
import { serveDuckdbRuntime } from './support/duckdb-runtime';

// Fight 3 brushed to its middle: Morrowlyn's one Frostbolt goes off at 5.0s, and the
// window's own cast lines say so. The summary keeps only whole-fight starts and
// failures, so before this the cell read a dash and the count a tilde.
test('a brushed window reads its casts from the fight’s events', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts&start=3000&end=10000');
  await expect(page.getByTestId('cast-measured-note')).toBeVisible({ timeout: 60_000 });
  const frostbolt = page.getByTestId('cast-Player-4184-000000A3-116');
  await expect(frostbolt).toContainText('1');
  await expect(frostbolt.getByTestId('cast-cancelled')).toContainText('0');
  await expect(page.getByTestId('cast-approximate-note')).toHaveCount(0);
});
