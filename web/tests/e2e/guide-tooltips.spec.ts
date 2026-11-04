// web/tests/e2e/guide-tooltips.spec.ts
// A spec guide names talents, spells, racials, enchants and items everywhere: the key-talent
// list, the build trees, the rotation and cooldown rows, the leveling strip, the gear and
// consumable tables, the races table and the prose itself. Each one opens the same game-style
// tooltip the rest of the site uses, on hover, keyboard focus or a tap, and Escape closes it.
// The tooltip data is fetched only once a reader asks for one.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';

const GUIDE = '/guides/warrior/fury';

async function openGuide(page: Page): Promise<string[]> {
  const bundleRequests: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/guide-tips/')) bundleRequests.push(request.url());
  });
  await page.goto(GUIDE);
  await page.locator('html[data-tooltip-host="ready"]').waitFor({ state: 'attached' });
  return bundleRequests;
}

test.describe('a spec guide opens the game tooltip for every name it shows', () => {
  test('a key talent opens on hover with its rank, and the data loads only then', async ({
    page,
    isMobile,
  }) => {
    test.skip(isMobile, 'hover is the desktop path; the tap test covers touch');
    const requests = await openGuide(page);
    expect(requests).toEqual([]);
    const talent = page.locator('[data-testid="guide-key-talents"] [data-game-tip]').first();
    await talent.scrollIntoViewIfNeeded();
    const name = (await talent.locator('.key-talent-name').innerText()).trim();
    await talent.hover();
    const tooltip = page.getByTestId('game-tooltip');
    await expect(tooltip).toBeVisible();
    await expect(tooltip.getByTestId('game-tooltip-name')).toHaveText(name);
    await expect(tooltip).toContainText(/Rank \d+\/\d+/);
    await expect(tooltip.getByTestId('game-tooltip-description')).not.toBeEmpty();
    expect(requests).toHaveLength(1);
    await page.mouse.move(0, 0);
    await expect(tooltip).toBeHidden();
  });

  test('a rotation row opens on keyboard focus and Escape closes it', async ({ page, isMobile }) => {
    test.skip(isMobile, 'keyboard path');
    await openGuide(page);
    const row = page.locator('[data-testid="guide-rotation-priority"] li[data-game-tip]').first();
    await row.scrollIntoViewIfNeeded();
    await row.focus();
    const tooltip = page.getByTestId('game-tooltip');
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText(/Instant|sec cast|Channeled|Passive/);
    await expect(row).toHaveAttribute('aria-describedby', (await tooltip.getAttribute('id'))!);
    await page.keyboard.press('Escape');
    await expect(tooltip).toBeHidden();
  });

  test('a spell named in the prose opens its tooltip on hover', async ({ page, isMobile }) => {
    test.skip(isMobile, 'hover is the desktop path');
    await openGuide(page);
    const mention = page.locator('main p span.gtip[data-game-tip^="spell:"]').first();
    await mention.scrollIntoViewIfNeeded();
    const name = (await mention.innerText()).trim();
    await mention.hover();
    await expect(page.getByTestId('game-tooltip-name')).toHaveText(name);
  });

  test('a gear pick and a consumable open the shared item tooltip', async ({ page, isMobile }) => {
    test.skip(isMobile, 'hover is the desktop path');
    await openGuide(page);
    const pick = page.locator('[data-testid="guide-bis"] .gbis-panel:visible [data-bis-item]').first();
    await pick.scrollIntoViewIfNeeded();
    await pick.locator('.gbis-name').hover();
    const itemTip = page.getByTestId('item-tooltip');
    await expect(itemTip).toBeVisible();
    await expect(itemTip).toContainText((await pick.locator('.gbis-name').innerText()).trim());
    await page.mouse.move(0, 0);
    await expect(itemTip).toBeHidden();

    const consumable = page.locator('[data-testid="guide-consumables"] [data-bis-item]').first();
    await consumable.scrollIntoViewIfNeeded();
    await consumable.hover();
    await expect(itemTip).toContainText(/^[\s\S]*Use: /);
  });

  test('a talent in the build tree opens the talent tooltip on hover and arrows move between cells', async ({
    page,
    isMobile,
  }) => {
    test.skip(isMobile, 'hover and keyboard path');
    await openGuide(page);
    // The tree hydrates when its section scrolls into view.
    await page.getByTestId('guide-tree-skeleton').first().scrollIntoViewIfNeeded();
    const columns = page.getByTestId('guide-tree-columns').first();
    await expect(columns).toBeVisible();
    const cell = columns.locator('[data-testid^="talent-"][role="group"]').first();
    await cell.hover();
    await expect(columns.getByRole('tooltip')).toBeVisible();
    await cell.focus();
    await page.keyboard.press('ArrowRight');
    await expect(columns.getByRole('tooltip')).toHaveCount(1);
  });

  test('a tap opens a talent on touch and a second tap closes it', async ({ page, isMobile }) => {
    test.skip(!isMobile, 'touch path');
    await openGuide(page);
    const chip = page.locator('[data-testid="guide-leveling-strip"] li[data-game-tip]').first();
    await chip.scrollIntoViewIfNeeded();
    await chip.tap();
    const tooltip = page.getByTestId('game-tooltip');
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText(/Rank \d+\/\d+/);
    const box = await tooltip.boundingBox();
    const viewport = page.viewportSize()!;
    expect(box!.x).toBeGreaterThanOrEqual(0);
    expect(box!.x + box!.width).toBeLessThanOrEqual(viewport.width);
    await chip.tap();
    await expect(tooltip).toBeHidden();
  });
});
