// web/tests/e2e/item-tooltip.spec.ts
// Lane web-item-tooltips: hovering an equipped item's icon anywhere on the site shows its
// real tooltip -- icon, stats and source -- the same way TalentCell's own tooltip already
// works for talents (tenet 2, "an item is never just a name"). Runs against the default
// fixture suite: item 16963 (Helm of Wrath, gear.spec.ts's own equip target) carries stats,
// a set (Battlegear of Wrath) and armor, which is enough to exercise the panel without
// needing loot.json content the fixture build does not ship.
import { expect, test, type Locator, type Page } from '@playwright/test';
import { openGear } from './support/planner';

test('hovering the head slot’s icon shows the item tooltip with icon, stats and level', async ({ page }) => {
  await page.goto('/planner');
  await openGear(page);

  await page.getByTestId('slot-head').click();
  await page.getByTestId('item-16963').click();
  await expect(page.getByTestId('slot-head')).toContainText('Helm of Wrath');

  const icon = page.getByTestId('item-hover-16963');
  await expect(icon).toBeVisible();
  await expect(page.getByTestId('item-tooltip')).toBeHidden();

  await icon.hover();
  const tooltip = page.getByTestId('item-tooltip');
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText('Helm of Wrath');
  await expect(tooltip).toContainText('Head');
  await expect(tooltip).toContainText('610 Armor');
  await expect(tooltip).toContainText('+25 Strength');
  await expect(tooltip).toContainText('+20 Stamina');
  await expect(tooltip).toContainText('Battlegear of Wrath');
  await expect(tooltip).toContainText('Requires Level 60');
  // The fixture's own icon file is a 1x1 placeholder (the data pipeline's stand-in for art
  // it has not produced yet), which ItemTooltip.svelte's own onload check turns into the
  // two-letter fallback mark instead of a blank image -- the same rule TalentCell's own
  // tooltip icon already follows. A broken/placeholder icon is the fixture reality here,
  // not a bug, so this asserts the fallback rather than an <img> the fixture cannot back.
  await expect(tooltip.getByText('He', { exact: true })).toBeVisible();
  const tooltipId = await tooltip.getAttribute('id');
  expect(tooltipId).not.toBeNull();
  await expect(icon).toHaveAttribute('aria-describedby', tooltipId!);

  // Moving off the icon closes it again -- the same hover/leave rule TalentCell's own
  // tooltip and the BiS popover both already follow.
  await page.mouse.move(0, 0);
  await expect(tooltip).toBeHidden();
});

test('keyboard focus opens the tooltip and Escape closes it', async ({ page }) => {
  await page.goto('/planner');
  await openGear(page);

  await page.getByTestId('slot-head').click();
  await page.getByTestId('item-16963').click();

  const icon = page.getByTestId('item-hover-16963');
  await icon.focus();
  const tooltip = page.getByTestId('item-tooltip');
  await expect(tooltip).toBeVisible();

  await page.keyboard.press('Escape');
  await expect(tooltip).toBeHidden();
});

test('only one item tooltip is open at a time', async ({ page }) => {
  await page.goto('/planner');
  await openGear(page);

  await page.getByTestId('slot-head').click();
  await page.getByTestId('item-16963').click();
  await page.getByTestId('slot-shoulder').click();
  await page.getByTestId('item-16966').click();

  const headIcon = page.getByTestId('item-hover-16963');
  const shoulderIcon = page.getByTestId('item-hover-16966');

  await headIcon.hover();
  await expect(page.getByTestId('item-tooltip')).toBeVisible();
  await expect(page.getByTestId('item-tooltip')).toContainText('Helm of Wrath');

  await shoulderIcon.hover();
  // Still exactly one tooltip on the page, now naming the shoulder piece.
  await expect(page.getByTestId('item-tooltip')).toHaveCount(1);
  await expect(page.getByTestId('item-tooltip')).toContainText('Spaulders of Wrath');
});

// Lane bis-tooltip-host (2026-09-29): /bis/<class>/<spec> no longer wraps each row in its
// own ItemHover island -- one shared BisTooltipHost island now delegates hover/focus/tap
// for the whole page (BisTooltipHost.svelte), so this contract needs its own coverage
// against real delegated listeners rather than one Svelte instance per row.
/** A gear row's item name alone -- the row's hover target wraps its slot label, "New"
 *  pill, stat line and source line too, none of which the tooltip repeats. Matches either
 *  a main pick's own name (`.gear-row-name`) or an alternative's (`.gear-row-alt-name`,
 *  fix round 1: real bands now publish alternatives, so the page's own `[data-testid^=
 *  "item-hover-"]` hosts are no longer main picks exclusively). */
async function rowItemName(row: Locator): Promise<string> {
  return (await row.locator('.gear-row-name, .gear-row-alt-name').first().innerText()).trim();
}

/** The band table the BiS page shows first (the nightly's lowest band, alliance). */
function firstVisibleBand(page: Page) {
  return page.locator('[data-testid^="bis-band-alliance-"]:visible').first();
}

test.describe('the BiS page shares one tooltip host across every row', () => {
  test('hovering a slot’s icon shows the item tooltip with icon, stats and level', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band = firstVisibleBand(page);
    await expect(band).toBeVisible();

    // The picks are the nightly's, so the test reads the first pill's item name off the
    // page and expects the tooltip to carry that name plus a real stat line.
    const icon = band.locator('[data-testid^="item-hover-"]').first();
    await expect(icon).toBeVisible();
    const itemName = await rowItemName(icon);
    expect(itemName.length).toBeGreaterThan(0);
    await expect(page.getByTestId('item-tooltip')).toBeHidden();

    // Hover the name text itself, not the host's own bounding-box centre (fix round 1: a
    // pick's host can carry alternatives below it, so its centre is not a safe hover point
    // to assume).
    await icon.locator('.gear-row-name, .gear-row-alt-name').first().hover();
    const tooltip = page.getByTestId('item-tooltip');
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText(itemName);
    await expect(tooltip).toContainText(/Requires Level \d+|Item Level \d+|\+\d+ /);
    const tooltipId = await tooltip.getAttribute('id');
    expect(tooltipId).not.toBeNull();
    await expect(icon).toHaveAttribute('aria-describedby', tooltipId!);

    await page.mouse.move(0, 0);
    await expect(tooltip).toBeHidden();
  });

  test('keyboard focus opens the tooltip and Escape closes it', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    // The shared host is a client:idle island; it opens for an already-focused row on
    // mount, but waiting for its ready marker keeps this test free of that race entirely.
    await page.locator('html[data-tooltip-host="ready"]').waitFor({ state: 'attached' });
    const icon = firstVisibleBand(page).locator('[data-testid^="item-hover-"]').first();

    await icon.focus();
    const tooltip = page.getByTestId('item-tooltip');
    await expect(tooltip).toBeVisible();

    await page.keyboard.press('Escape');
    await expect(tooltip).toBeHidden();
  });

  test('only one tooltip is open at a time across the page’s delegated host', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const pills = firstVisibleBand(page).locator('[data-testid^="item-hover-"]');
    const firstIcon = pills.nth(0);
    const secondIcon = pills.nth(1);
    const firstName = await rowItemName(firstIcon);
    const secondName = await rowItemName(secondIcon);
    expect(secondName).not.toBe(firstName);

    await firstIcon.locator('.gear-row-name, .gear-row-alt-name').first().hover();
    await expect(page.getByTestId('item-tooltip')).toBeVisible();
    await expect(page.getByTestId('item-tooltip')).toContainText(firstName);

    await secondIcon.locator('.gear-row-name, .gear-row-alt-name').first().hover();
    await expect(page.getByTestId('item-tooltip')).toHaveCount(1);
    await expect(page.getByTestId('item-tooltip')).toContainText(secondName);
  });

  // Tooltip-polish brief item 5 (owner screenshot 2026-09-29): the panel never extends below
  // the viewport -- lib/items/tooltip-position.ts flips it above the anchor, or clamps and
  // scrolls internally when neither side fits. A row anywhere in the middle of the page never
  // exercises that path (Playwright centers whatever it hovers), so this forces the anchor to
  // the very bottom edge of the viewport first.
  test('a row pinned to the bottom edge of the viewport opens the tooltip without spilling past it', async ({
    page,
  }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band = firstVisibleBand(page);
    await expect(band).toBeVisible();
    const icon = band.locator('[data-testid^="item-hover-"]').last();
    await expect(icon).toBeVisible();

    // `block: 'end'` aligns the anchor's own bottom edge with the viewport's bottom edge --
    // the same "nothing below it" case a long band list's last row hits on a short viewport,
    // without needing a fixture tall enough to push a real row there on its own.
    await icon.evaluate((el) => el.scrollIntoView({ block: 'end' }));
    await icon.hover();

    const tooltip = page.getByTestId('item-tooltip');
    await expect(tooltip).toBeVisible();
    const box = await tooltip.boundingBox();
    const viewport = page.viewportSize();
    expect(box).not.toBeNull();
    expect(viewport).not.toBeNull();
    // The +1 covers sub-pixel rounding between the two measurements, nothing more.
    expect(box!.y + box!.height).toBeLessThanOrEqual(viewport!.height + 1);
  });
});
