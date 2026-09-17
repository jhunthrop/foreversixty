// web/tests/e2e/report-compare.spec.ts
// Compare's two questions. "What did the player above me do differently" is one player's
// abilities in two pulls, which is a row expanding. "What did I do differently from them"
// is two players in this pull, which is the second picker.
import { expect, test } from '@playwright/test';

const COMPARE = '/reports/fixture2abcd?fight=3&mode=compare';

/**
 * CompareMode.svelte renders a player's row twice at once -- the desktop `<tr>`
 * (`compare-<guid>`) and the phone `<li>` card (`compare-card-<guid>`) -- and hides
 * whichever the current width does not use with a `display` class rather than removing it
 * from the DOM, so both its own "Abilities" button and its own expansion panel
 * (`compare-abilities-<guid>` on the row, `compare-abilities-card-<guid>` on the card)
 * exist on the page regardless of project. This opens whichever one a viewer at the
 * current width can actually see and returns that panel, so the three ability-expansion
 * tests below exercise the phone card under `--project=mobile` the same way they exercise
 * the desktop row under `--project=desktop`, instead of only ever driving the row.
 */
async function openAbilities(
  page: import('@playwright/test').Page,
  guid: string,
): Promise<import('@playwright/test').Locator> {
  const button = page
    .locator(`[data-testid="compare-${guid}"], [data-testid="compare-card-${guid}"]`)
    .getByRole('button', { name: /abilities/i })
    .filter({ visible: true });
  await button.click();
  return page
    .locator(`[data-testid="compare-abilities-${guid}"], [data-testid="compare-abilities-card-${guid}"]`)
    .filter({ visible: true });
}

test('a player row expands to that player’s abilities in both pulls', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4`);
  const diff = await openAbilities(page, 'Player-4184-000000A1');
  await expect(diff).toBeVisible();
  // Baelgrim's Slam is 4,400 in fight 3 and 482,100 in fight 4.
  await expect(diff).toContainText('Slam');
  await expect(diff.getByTestId('ability-delta').first()).toContainText('−');
});

test('an ability only one pull has shows a dash on the other side', async ({ page }) => {
  await page.goto(`${COMPARE}&with=1`);
  const diff = await openAbilities(page, 'Player-4184-000000A1');
  await expect(diff).toContainText('—');
});

test('threat has no ability split, and the row says so rather than opening empty', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4&cmetric=threat`);
  const diff = await openAbilities(page, 'Player-4184-000000A1');
  await expect(diff).toContainText('no ability split');
});

test('picking a player compares the two inside this pull, and rides in the url', async ({ page }) => {
  await page.goto(COMPARE);
  await page.getByTestId('compare-vs').selectOption('Player-4184-000000A3');
  await expect(page).toHaveURL(/vs=Player-4184-000000A3/);
  // The player table collapses to the two of them.
  const rows = page.getByTestId('compare-cards').locator('li');
  await expect(rows).toHaveCount(2);
  // And the ability diff below is Slam against Frostbolt.
  const diff = page.getByTestId('compare-players');
  await expect(diff).toContainText('Slam');
  await expect(diff).toContainText('Frostbolt');
});

test('the link alone opens on two players', async ({ page }) => {
  await page.goto(`${COMPARE}&vs=Player-4184-000000A3`);
  await expect(page.getByTestId('compare-vs')).toHaveValue('Player-4184-000000A3');
  await expect(page.getByTestId('compare-players')).toContainText('Frostbolt');
});

test('whatever is on screen copies as a CSV', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4`);
  await expect(page.getByTestId('compare-mode').getByTestId('copy-csv').first()).toBeVisible();
});
