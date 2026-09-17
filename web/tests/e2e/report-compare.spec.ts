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

// Fight 3's DPS, rounded the way the cards round it: Baelgrim 73, Morrowlyn 52, Elyra 11.
test('the two cards read the two players’ own totals, against each other', async ({ page }) => {
  await page.goto(`${COMPARE}&vs=Player-4184-000000A3`);
  const cards = page.getByTestId('compare-cards').locator('li');
  await expect(cards).toHaveCount(2);
  // The base is the metric's leader when Source names nobody.
  const base = page.getByTestId('compare-card-Player-4184-000000A1');
  await expect(base).toContainText('73 in this pull');
  await expect(base).toContainText('52 for Morrowlyn');
  await expect(base.getByTestId('compare-card-delta')).toHaveText('+21');
  // And the other card is the same two numbers the other way round.
  const other = page.getByTestId('compare-card-Player-4184-000000A3');
  await expect(other).toContainText('52 in this pull');
  await expect(other).toContainText('73 for Baelgrim');
  await expect(other.getByTestId('compare-card-delta')).toHaveText('\u221221');
});

test('Source picks the base player, so a reader can compare themselves', async ({ page }) => {
  await page.goto(`${COMPARE}&vs=Player-4184-000000A3&source=Player-4184-000000A5`);
  await expect(page.getByTestId('compare-players-scope')).toContainText('Elyra Duskvale against Morrowlyn');
  const mine = page.getByTestId('compare-card-Player-4184-000000A5');
  await expect(mine).toContainText('11 in this pull');
  await expect(mine).toContainText('52 for Morrowlyn');
  await expect(mine.getByTestId('compare-card-delta')).toHaveText('\u221241');
  // The leader is not on the panel at all: the question was about these two.
  await expect(page.getByTestId('compare-card-Player-4184-000000A1')).toHaveCount(0);
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

// The fixture's two encounters are a boss with phases and a boss with none, so they share
// no phase: the picker says so rather than offering an alignment it cannot make.
test('the phase picker says when the two pulls share no phase', async ({ page }) => {
  await page.goto(`${COMPARE}&with=4`);
  await expect(page.getByTestId('compare-phase')).toHaveCount(0);
  await expect(page.getByTestId('compare-phase-note')).toContainText('no phase in common');
});

test('there is no phase picker before a second fight is picked', async ({ page }) => {
  await page.goto(COMPARE);
  await expect(page.getByTestId('compare-phase')).toHaveCount(0);
  await expect(page.getByTestId('compare-phase-note')).toHaveCount(0);
});
