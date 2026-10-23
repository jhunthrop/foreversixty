// web/tests/e2e/guides-sections.spec.ts
// W.18: the spec guide reads like a class guide. This walks the Frost Mage page section by
// section at desktop (1280) and phone (390) widths: the single status note, the Overview
// blocks, the builds and the key talents, the rotation lists, the stat block, the
// best-in-slot table and its level-range switcher, the enchant and consumable tables, the
// race table and the leveling strip, and checks the page never scrolls sideways.
//
// The content comes from data/builds and the BiS ranker, not the planner fixture, so it is the
// same under either FOREVER_DATA; only the check that icon files actually load needs the real
// icon tree and so runs under FOREVER_DATA=real.
import { expect, test, type Page } from '@playwright/test';

const DESKTOP = { width: 1280, height: 900 };
const PHONE = { width: 390, height: 844 };
const PAGE = '/guides/mage/frost';

const SECTION_IDS = [
  'overview',
  'talents-and-builds',
  'rotation-and-priority',
  'stat-priority',
  'best-in-slot',
  'enchants-and-consumables',
  'races',
  'professions',
  'leveling',
];

async function noSidewaysScroll(page: Page): Promise<void> {
  const { scrollWidth, clientWidth } = await page.evaluate(() => ({
    scrollWidth: document.documentElement.scrollWidth,
    clientWidth: document.documentElement.clientWidth,
  }));
  expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
}

async function imagesLoaded(page: Page, selector: string): Promise<void> {
  if (process.env.FOREVER_DATA !== 'real') return;
  const broken = await page.evaluate(
    (css) =>
      [...document.querySelectorAll<HTMLImageElement>(css)]
        .filter((img) => img.complete && img.naturalWidth === 0)
        .map((img) => img.src),
    selector,
  );
  expect(broken).toEqual([]);
}

for (const [label, viewport] of [
  ['1280', DESKTOP],
  ['390', PHONE],
] as const) {
  test.describe(`Frost Mage guide at ${label}`, () => {
    test.use({ viewport });

    test('has the nine sections in order, the status note once, and the header facts', async ({ page }) => {
      await page.goto(PAGE);
      const ids = await page.locator('main h2[id]').evaluateAll((els) => els.map((el) => el.id));
      expect(ids.filter((id) => SECTION_IDS.includes(id))).toEqual(SECTION_IDS);

      await expect(page.getByTestId('guide-status-note')).toHaveCount(1);
      await expect(page.getByTestId('guide-status-note')).toContainText('The beta stops at level 30.');
      await expect(page.getByTestId('guide-header-meta')).toContainText(
        /DPS · Updated \d{1,2} [A-Z][a-z]{2} · built on client 1\.60\.1\.\d+ · beta-tested to level 30, level 60 simulated/,
      );
      await expect(page.locator('.class-header-summary')).toContainText('Frost is the Mage control');
      await noSidewaysScroll(page);
    });

    test('lays out the Overview as strengths, weaknesses and the raid role', async ({ page }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-overview-strengths').locator('li')).toHaveCount(4);
      await expect(page.getByTestId('guide-overview-weaknesses').locator('li')).toHaveCount(3);
      await expect(page.getByTestId('guide-overview-raid')).toContainText('Role:');
      await expect(page.getByTestId('guide-overview-raid')).toContainText('Brings:');
    });

    test('shows both builds, the raid-build diff line and the key talents with icons', async ({ page }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-build-leveling')).toBeVisible();
      await expect(page.getByTestId('guide-build-raid')).toBeVisible();
      await expect(page.getByTestId('guide-build-diff')).toContainText(
        /^The raid build moves \d+ points from /,
      );
      await expect(page.getByTestId('guide-load-build')).toBeVisible();
      await expect(page.getByTestId('guide-sim-raid-build')).toBeVisible();
      const key = page.getByTestId('guide-key-talents');
      await expect(key.locator('h3')).toHaveText('Key talents');
      await expect(key.locator('li').first()).toContainText('Improved Frostbolt');
      await expect(key.locator('li').first().locator('img')).toHaveCount(1);
      await imagesLoaded(page, '[data-testid="guide-key-talents"] img');
    });

    test('numbers an opener and a priority with spell icons', async ({ page }) => {
      await page.goto(PAGE);
      const priority = page.getByTestId('guide-rotation-priority');
      await expect(priority.locator('h3')).toHaveText('Priority');
      await expect(priority.locator('li').first()).toContainText('Ice Lance');
      await expect(priority.locator('li').first().locator('img')).toHaveCount(1);
      await expect(page.getByTestId('guide-rotation-opener').locator('h3')).toHaveText('Opener');
      await expect(page.getByTestId('guide-rotation-cooldowns').locator('h3')).toHaveText('Cooldowns');
      await imagesLoaded(page, '[data-testid="guide-rotation-lists"] img');
    });

    test('gives the stat block its pills, its table and the hit cap line', async ({ page }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('stat-priority-pills')).toBeVisible();
      await expect(page.getByTestId('guide-stat-table')).toBeVisible();
      await expect(page.getByTestId('guide-stat-table-hit-cap')).toHaveText(/^Hit cap: \d+(\.\d)?%/);
    });

    test('opens best in slot on 60 and switches level range without a script', async ({ page }) => {
      await page.goto(PAGE);
      const panel60 = page.getByTestId('guide-bis-panel-60');
      const panel30 = page.getByTestId('guide-bis-panel-30');
      await expect(panel60).toBeVisible();
      await expect(panel30).toBeHidden();
      await expect(page.getByTestId('guide-bis-radio-60')).toBeChecked();
      await expect(panel60.locator('tbody tr')).toHaveCount(17);
      await expect(panel60.getByTestId('guide-bis-row-60-head')).toContainText('Head');
      await expect(panel60.getByTestId('guide-bis-row-60-head').locator('img')).toHaveCount(1);
      await expect(panel60.getByRole('link', { name: 'Full list and alternatives' })).toHaveAttribute(
        'href',
        '/bis/mage/frost#band-alliance-60',
      );

      await page.getByTestId('guide-bis-pill-30').click();
      await expect(panel30).toBeVisible();
      await expect(panel60).toBeHidden();
      await expect(panel30.getByRole('link', { name: 'Full list and alternatives' })).toHaveAttribute(
        'href',
        '/bis/mage/frost#band-alliance-30',
      );
      await expect(panel30.locator('tbody tr')).toHaveCount(17);

      await page.getByTestId('guide-bis-pill-60').click();
      await expect(panel60).toBeVisible();
      await imagesLoaded(page, '[data-testid="guide-bis"] img');
      await noSidewaysScroll(page);
    });

    test('keeps the best-in-slot table inside its own scrolling box', async ({ page }) => {
      await page.goto(PAGE);
      const box = page.getByTestId('guide-bis-panel-60').getByRole('region');
      await expect(box).toHaveAttribute('tabindex', '0');
      const { scrollWidth, clientWidth } = await box.evaluate((el) => ({
        scrollWidth: el.scrollWidth,
        clientWidth: el.clientWidth,
      }));
      // The box may scroll (phone) or fit (desktop); it never makes the page scroll.
      expect(scrollWidth).toBeGreaterThanOrEqual(clientWidth);
      await noSidewaysScroll(page);
    });

    test('lists enchants and consumables from the data, only for the slots it can verify', async ({
      page,
    }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-enchants').locator('h3')).toHaveText('Enchants');
      await expect(page.getByTestId('guide-enchant-main_hand')).toContainText('+30 spell power');
      await expect(page.getByTestId('guide-consumables').locator('h3')).toHaveText('Consumables');
      await expect(page.getByTestId('guide-consumables')).toContainText('Flask of Supreme Power');
      await expect(page.getByTestId('guide-consumables')).toContainText('Elixir of Frost Power');
    });

    test('has a race table with the recommended races marked, under the pills', async ({ page }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-race-pills')).toBeVisible();
      const table = page.getByTestId('guide-race-table');
      await expect(table.locator('tbody tr')).toHaveCount(6);
      await expect(page.getByTestId('guide-race-row-gnome')).toContainText('Recommended');
      await expect(page.getByTestId('guide-race-row-gnome')).toContainText('Eureka!');
      await expect(page.getByTestId('guide-race-row-troll')).toContainText('Berserking');
      await expect(page.getByTestId('guide-race-row-human')).not.toContainText('Recommended');
    });

    test('lists each leveling band’s talents with icons, names and ranks, and a planner link', async ({
      page,
    }) => {
      await page.goto(PAGE);
      const strip = page.getByTestId('guide-leveling-strip');
      const rows = strip.locator('[data-testid^="guide-leveling-row-"]:not([data-testid$="-talents"])');
      await expect(rows).toHaveCount(5);
      for (const row of await rows.all()) {
        await expect(row.getByText('Load in planner')).toHaveAttribute(
          'href',
          /^\/planner\?spec=mage-frost&talents=/,
        );
        const chips = row.getByTestId('guide-leveling-row-talents').locator('li');
        expect(await chips.count()).toBeGreaterThan(3);
        await expect(chips.first().locator('img')).toHaveCount(1);
        await expect(chips.first()).toContainText(/\d\/\d/);
      }
      await imagesLoaded(page, '[data-testid="guide-leveling-strip"] img');
      await noSidewaysScroll(page);
    });
  });
}

test.describe('table of contents', () => {
  test('desktop: the sticky rail jumps to a section', async ({ page }) => {
    await page.setViewportSize(DESKTOP);
    await page.goto(PAGE);
    const rail = page.getByTestId('guide-toc-rail');
    await expect(rail).toBeVisible();
    await rail.getByRole('link', { name: 'Best in slot' }).click();
    await expect(page).toHaveURL(/#best-in-slot$/);
    await expect(page.locator('#best-in-slot')).toBeInViewport();
  });

  test('phone: the contents drop down from the top and jump to a section', async ({ page }) => {
    await page.setViewportSize(PHONE);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-toc-rail')).toBeHidden();
    const inline = page.getByTestId('guide-toc-inline');
    await expect(inline).toBeVisible();
    await inline.locator('summary').click();
    await inline.getByRole('link', { name: 'Races' }).click();
    await expect(page).toHaveURL(/#races$/);
    await expect(page.locator('#races')).toBeInViewport();
  });
});
