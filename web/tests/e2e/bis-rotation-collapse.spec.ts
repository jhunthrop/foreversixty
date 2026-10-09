// web/tests/e2e/bis-rotation-collapse.spec.ts
// The Play it list is one row per ability, not one per curated step: the healer lists have a
// step per target and threshold (Heal alone is ten), and a row that repeats the spell above it
// is the same ability again with its reason cut off.
import { expect, test } from '@playwright/test';

const SPEC_PAGES = ['/bis/priest/holy', '/bis/mage/fire', '/bis/warrior/arms'];

for (const path of SPEC_PAGES) {
  test(`${path}: no two consecutive Play it rows name the same spell and rank`, async ({ page }) => {
    await page.goto(path);
    const panels = page.getByTestId('bis-play-it');
    expect(await panels.count()).toBeGreaterThan(0);
    const lists = await panels.evaluateAll((nodes) =>
      nodes.map((node) =>
        [...node.querySelectorAll('.play-it-line')].map((row) =>
          [
            row.querySelector('.play-it-name')?.textContent?.trim(),
            row.querySelector('.play-it-rank')?.textContent?.trim() ?? '',
          ].join('|'),
        ),
      ),
    );
    for (const rows of lists) {
      rows.slice(1).forEach((row, index) => expect(row).not.toBe(rows[index]));
    }
  });
}

test('the holy priest Heal row says its conditions in words', async ({ page }) => {
  await page.goto('/bis/priest/holy');
  const heal = page
    .getByTestId('bis-play-it')
    .first()
    .locator('.play-it-line', { has: page.locator('.play-it-name', { hasText: /^Heal$/ }) });
  await expect(heal).toHaveCount(1);
  await expect(heal.locator('.play-it-condition')).toContainText(/Cast on the tank below \d+%/);
});
