// web/tests/e2e/item-tooltip.spec.ts
// Lane web-item-tooltips: hovering an equipped item's icon anywhere on the site shows its
// real tooltip -- icon, stats and source -- the same way TalentCell's own tooltip already
// works for talents (tenet 2, "an item is never just a name"). Runs against the default
// fixture suite: item 16963 (Helm of Wrath, gear.spec.ts's own equip target) carries stats,
// a set (Battlegear of Wrath) and armor, which is enough to exercise the panel without
// needing loot.json content the fixture build does not ship.
import { expect, test } from '@playwright/test';
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
test.describe('the BiS page shares one tooltip host across every row', () => {
  test('hovering a slot’s icon shows the item tooltip with icon, stats and level', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band10 = page.getByTestId('bis-band-alliance-10');
    await expect(band10).toBeVisible();

    const icon = band10.getByTestId('item-hover-19972'); // head slot, band 10, alliance
    await expect(icon).toBeVisible();
    await expect(page.getByTestId('item-tooltip')).toBeHidden();

    await icon.hover();
    const tooltip = page.getByTestId('item-tooltip');
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText('Lucky Fishing Hat');
    await expect(tooltip).toContainText('Head');
    await expect(tooltip).toContainText('43 Armor');
    await expect(tooltip).toContainText('+15 Stamina');
    const tooltipId = await tooltip.getAttribute('id');
    expect(tooltipId).not.toBeNull();
    await expect(icon).toHaveAttribute('aria-describedby', tooltipId!);

    await page.mouse.move(0, 0);
    await expect(tooltip).toBeHidden();
  });

  test('keyboard focus opens the tooltip and Escape closes it', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band10 = page.getByTestId('bis-band-alliance-10');
    const icon = band10.getByTestId('item-hover-19972');

    await icon.focus();
    const tooltip = page.getByTestId('item-tooltip');
    await expect(tooltip).toBeVisible();

    await page.keyboard.press('Escape');
    await expect(tooltip).toBeHidden();
  });

  test('only one tooltip is open at a time across the page’s delegated host', async ({ page }) => {
    await page.goto('/bis/hunter/marksmanship');
    const band10 = page.getByTestId('bis-band-alliance-10');
    const headIcon = band10.getByTestId('item-hover-19972');
    const chestIcon = band10.getByTestId('item-hover-3288');

    await headIcon.hover();
    await expect(page.getByTestId('item-tooltip')).toBeVisible();
    await expect(page.getByTestId('item-tooltip')).toContainText('Lucky Fishing Hat');

    await chestIcon.hover();
    await expect(page.getByTestId('item-tooltip')).toHaveCount(1);
    await expect(page.getByTestId('item-tooltip')).toContainText('Tribal Vest');
  });
});
