// web/tests/e2e/planner-bis-hover.spec.ts
// Lane bis-hover-web: hovering a planner gear slot shows the leveling BiS pick for the
// character's own spec, faction and level band (lib/bis/hover.ts), fetched on demand from
// the per-spec static asset scripts/sync-data.mjs publishes rather than bundled into the
// planner island (design step 2 of this lane's brief).
//
// Runs against the default fixture suite. web/src/fixtures/planner/talents/hunter.json
// (this lane's own addition -- no hunter fixture existed before it) gives Marksmanship
// (talents 4001-4005) exactly 21 points of capacity across five tiers (5/5/5/5/1), so
// spending all of it reaches level 30 (BASE_LEVEL 9 + 21) in one tree, the level the leveling
// BiS band-10..60-step-5 rule rounds down to 30 too. src/data/fixtures/bis/hunter-
// marksmanship.json (lane bis-web's own fixture) is what the popover fetches for it; its
// alliance band 30 head slot is "Iron Cap of the Ladder", new since band 25's "Bone Cap of
// the Ladder" -- both asserted below.
import { expect, test, type Page } from '@playwright/test';
import { openGear } from './support/planner';

const MD_BREAKPOINT = 768;

/** Below md the planner shows one tree at a time behind a tab strip (TreeTabs.svelte); from
 *  md up every tree column is already visible. */
async function showTree(page: Page, name: string): Promise<void> {
  if ((page.viewportSize()?.width ?? MD_BREAKPOINT) < MD_BREAKPOINT) {
    await page.getByRole('tab', { name }).click();
  }
}

async function spendPoints(page: Page, talentId: number, times: number): Promise<void> {
  const cell = page.getByTestId(`talent-${talentId}`);
  for (let i = 0; i < times; i += 1) await cell.click();
}

test('hovering a gear slot shows the leveling BiS pick for the planner’s spec and band', async ({ page }) => {
  await page.goto('/planner?class=hunter&race=dwarf');
  await showTree(page, 'Marksmanship');
  await expect(page.getByTestId('talent-4001')).toBeVisible();

  await spendPoints(page, 4001, 5);
  await spendPoints(page, 4002, 5);
  await spendPoints(page, 4003, 5);
  await spendPoints(page, 4004, 5);
  await spendPoints(page, 4005, 1);
  // 21 points, all in Marksmanship (tree id 302): level 9 + 21 = 30. `planner-level` itself
  // only renders once a current-character pointer exists (SummaryBar.svelte's own
  // `hasCharacter` gate), which a fresh planner session never has, so the tree's own point
  // counter is what confirms the level here.
  await expect(page.getByTestId('tree-points-302')).toHaveText('21');

  await openGear(page);
  const headSlot = page.getByTestId('slot-head');
  await headSlot.hover();

  const popover = page.getByTestId('bis-hover-head');
  await expect(popover).toBeVisible();
  await expect(popover).toContainText('Best in slot at level 30');
  await expect(popover).toContainText('Iron Cap of the Ladder');
  await expect(popover).toContainText('New at 30');
  await expect(popover.getByTestId('bis-hover-was')).toContainText('Bone Cap of the Ladder');
  await expect(headSlot).toHaveAttribute('aria-describedby', 'bis-hover-head');

  // A link to the full band, not just the hover card's own summary.
  await expect(popover.getByRole('link', { name: /level 30 list/ })).toHaveAttribute(
    'href',
    '/bis/hunter/marksmanship?faction=alliance#band-alliance-30',
  );

  // Moving off the slot closes the popover again.
  await page.mouse.move(0, 0);
  await expect(popover).toBeHidden();
});

test('a spec with no leveling BiS list yet shows the honest empty state, not a blank card', async ({
  page,
}) => {
  await page.goto('/planner'); // default class has no committed BiS fixture
  await openGear(page);
  await page.getByTestId('slot-head').hover();
  const popover = page.getByTestId('bis-hover-head');
  await expect(popover).toBeVisible();
  await expect(popover.getByTestId('bis-hover-empty')).toBeVisible();
});
