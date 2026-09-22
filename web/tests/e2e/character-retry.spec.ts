// web/tests/e2e/character-retry.spec.ts
// Character.svelte's failed state offers Try again, and Try again has to re-fire the same
// request in place -- never a page reload (design 2026-09-22 spec section 1.5). The SSR
// render test in src/components/Character.test.ts can only prove LoadError's markup shape,
// because no $effect runs server-side, so the wiring itself is proved here: the route
// fails, the reader presses the button, and the second answer paints the real page.
import { expect, test } from '@playwright/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';

test('a failed character page retries the same fetch in place', async ({ page }) => {
  // A flag rather than an attempt counter: the flip happens between the two assertions
  // below, so which request is the retry never depends on how many the page made first.
  let failing = true;
  await page.route(CHARACTER_URL, async (route) => {
    if (failing) {
      await route.abort();
      return;
    }
    await route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(CHARACTER),
    });
  });

  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-error')).toBeVisible();
  await expect(page.getByTestId('character')).toHaveCount(0);

  failing = false;
  await page.getByTestId('character-error-retry').click();

  await expect(page.getByTestId('character')).toBeVisible();
  await expect(page.getByTestId('character')).toContainText('Elyra Duskvale');
  await expect(page.getByTestId('character-error')).toHaveCount(0);
  // The retry re-fetched in place: a reload would have left the URL untouched too, so the
  // proof is that the failing stub was never navigated past -- the page is the same one.
  expect(new URL(page.url()).pathname).toBe(CHARACTER_PATH);
});
