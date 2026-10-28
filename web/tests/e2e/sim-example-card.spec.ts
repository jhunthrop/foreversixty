// web/tests/e2e/sim-example-card.spec.ts
// The /sim landing's example card shows the real server run (design/specs/2026-10-27-launch-pricing.md
// B3.4, B8 check 13): the B5 label, the caption read from the data file, a reserved picture, and
// none of the invented figure the card used to carry.
import { expect, test } from '@playwright/test';
import { EXAMPLE_IMAGE_HEIGHT, EXAMPLE_IMAGE_WIDTH } from '../../src/lib/billing/example-image';
import { serverRunCaption } from '../../src/lib/billing/premium-example-copy';
import { landingCopy } from '../../src/lib/sim/landing-copy';

const INVENTED_FIGURE = '1,245 DPS';

test('the landing card shows the label, the data-file caption and a reserved picture', async ({ page }) => {
  await page.goto('/sim');
  const card = page.getByTestId('sim-example-result').first();
  await expect(card).toContainText(landingCopy.exampleCaption);
  await expect(card.getByTestId('sim-example-caption')).toHaveText(serverRunCaption());
  const image = card.locator('img');
  await expect(image).toHaveAttribute('width', String(EXAMPLE_IMAGE_WIDTH));
  await expect(image).toHaveAttribute('height', String(EXAMPLE_IMAGE_HEIGHT));
  await expect(image).toHaveAttribute('alt', serverRunCaption());
});

test('the invented figure appears nowhere on /sim', async ({ page }) => {
  await page.goto('/sim');
  expect(await page.content()).not.toContain(INVENTED_FIGURE);
});
