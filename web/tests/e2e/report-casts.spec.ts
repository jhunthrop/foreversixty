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

// Warden Kelthas starts Anima Surge at 14.0s and never lands it, so no window holds a
// success of it and the windowed rows alone would never show it. Brushed around the
// start, the measure gives it a row: one started, one cancelled.
test('a spell the window only saw started gets a row with its measured counts', async ({ page }) => {
  test.slow();
  await serveDuckdbRuntime(page);
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts&source=enemies&start=12000&end=20000');
  await expect(page.getByTestId('cast-measured-note')).toBeVisible({ timeout: 60_000 });
  const surge = page.getByTestId('cast-Creature-0-2085-2284-7855-169754-0000AA0002-334653');
  await expect(surge).toContainText('Anima Surge');
  await expect(surge.getByTestId('cast-cancelled')).toContainText('1');
});

// Ashfang bites once at 1.5s of fight 3. Under all friendlies the row is there, filed
// under its owner and saying which pet it was.
test('a pet’s cast row sits under its owner and says which pet cast it', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts');
  const bite = page.getByTestId('cast-Pet-0-2085-2284-7855-165189-01000000B1-17253');
  await expect(bite).toBeVisible();
  await expect(bite).toContainText('Thalgrit');
  await expect(bite).toContainText('via Ashfang');
});

test('picking the pet’s owner keeps the pet’s casts on screen', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts&source=Player-4184-000000A4');
  await expect(page.getByTestId('cast-Pet-0-2085-2284-7855-165189-01000000B1-17253')).toBeVisible();
});

test('picking another player leaves the pet’s casts out', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&tab=casts&source=Player-4184-000000A3');
  await expect(page.getByTestId('cast-Pet-0-2085-2284-7855-165189-01000000B1-17253')).toHaveCount(0);
});
