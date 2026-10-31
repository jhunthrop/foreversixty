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
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';

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
        /DPS · Updated \d{1,2} [A-Z][a-z]{2}/,
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

    test('numbers a priority with spell icons and lists real cooldowns, with no opener', async ({ page }) => {
      await page.goto(PAGE);
      const priority = page.getByTestId('guide-rotation-priority');
      await expect(priority.locator('h3')).toHaveText('Priority');
      await expect(priority.locator('li').first()).toContainText('Ice Lance');
      await expect(priority.locator('li').first().locator('img')).toHaveCount(1);
      await expect(page.getByTestId('guide-rotation-opener')).toHaveCount(0);
      const cooldowns = page.getByTestId('guide-rotation-cooldowns');
      await expect(cooldowns.locator('h3')).toHaveText('Cooldowns');
      for (const name of ['Evocation', 'Cold Snap', 'Ice Block']) await expect(cooldowns).toContainText(name);
      await imagesLoaded(page, '[data-testid="guide-rotation-lists"] img');
    });

    test('gives the stat block its pills, the hit headline and the weights table', async ({ page }) => {
      await page.goto(PAGE);
      await expect(page.getByTestId('stat-priority-pills')).toBeVisible();
      await expect(page.getByTestId('guide-stat-table')).toBeVisible();
      // The talent is named from the active build's trees (the unit test in
      // src/pages/guides/_spec.test.ts pins which one), so a rename never fails this walk.
      await expect(page.getByTestId('guide-stat-table-hit-cap')).toHaveText(
        /^Hit cap: \d+(\.\d)?% · \d+(\.\d)?% from [^,]+, \d+(\.\d)?% from gear$/,
      );
      await expect(page.getByTestId('guide-stat-table')).toContainText('DPS per point');
      const pills = await page.locator('[data-testid^="stat-pill-"] span:last-child').allTextContents();
      const rows = await page.locator('[data-testid^="guide-stat-table-row-"] .stat-name').allTextContents();
      expect(pills).toEqual(rows);
    });

    test('opens best in slot on 60 and switches faction and level range without a script', async ({
      page,
    }) => {
      await page.goto(PAGE);
      const alliance60 = page.getByTestId('guide-bis-panel-alliance-60');
      await expect(alliance60).toBeVisible();
      await expect(page.getByTestId('guide-bis-panel-alliance-30')).toBeHidden();
      await expect(page.getByTestId('guide-bis-radio-60')).toBeChecked();
      await expect(alliance60.locator('tbody tr')).toHaveCount(17);
      await expect(alliance60.getByTestId('guide-bis-row-alliance-60-head')).toContainText('Head');
      await expect(alliance60.getByTestId('guide-bis-row-alliance-60-head').locator('img')).toHaveCount(1);
      await expect(alliance60.getByRole('link', { name: 'Full list and alternatives' })).toHaveAttribute(
        'href',
        '/bis/mage/frost?faction=alliance#band-alliance-60',
      );
      await expect(page.getByTestId('guide-bis-pill-60')).toHaveText('Level 60');

      await page.getByTestId('guide-bis-pill-30').click();
      await expect(page.getByTestId('guide-bis-panel-alliance-30')).toBeVisible();
      await expect(alliance60).toBeHidden();

      await page.getByTestId('guide-bis-faction-horde').click();
      const horde30 = page.getByTestId('guide-bis-panel-horde-30');
      await expect(horde30).toBeVisible();
      await expect(page.getByTestId('guide-bis-caption-horde-30')).toContainText('Troll (Horde)');
      await expect(horde30.getByRole('link', { name: 'Full list and alternatives' })).toHaveAttribute(
        'href',
        '/bis/mage/frost?faction=horde#band-horde-30',
      );

      await page.getByTestId('guide-bis-pill-60').click();
      await expect(page.getByTestId('guide-bis-panel-horde-60')).toBeVisible();
      await imagesLoaded(page, '[data-testid="guide-bis"] img');
      await noSidewaysScroll(page);
    });

    test('opens best in slot on the stored character own level range', async ({ page }) => {
      await page.addInitScript(() => {
        // 24 talent points spent: level 33, the 30 to 39 range.
        localStorage.setItem(
          'fs.currentCharacter',
          JSON.stringify({
            source: 'code',
            ref: 'FS1:1.60.1.70532:mage:gnome:55/55/4:',
            label: 'Test',
            classSlug: 'mage',
            savedAt: new Date().toISOString(),
          }),
        );
      });
      await page.goto(PAGE);
      await expect(page.getByTestId('guide-bis-radio-30')).toBeChecked({ timeout: 5000 });
      await expect(page.getByTestId('guide-bis-panel-alliance-30')).toBeVisible();
    });

    test('keeps the best-in-slot table inside its own scrolling box', async ({ page }) => {
      await page.goto(PAGE);
      const box = page.getByTestId('guide-bis-panel-alliance-60').getByRole('region');
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
      await expect(page.getByTestId('guide-race-row-gnome')).toContainText('Escape Artist');
      await expect(page.getByTestId('guide-race-row-gnome')).toContainText('Expansive Mind');
      await expect(page.getByTestId('guide-race-picks')).toContainText(
        'Gnome is the Alliance pick: Escape Artist and Expansive Mind.',
      );
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
        await expect(row.getByTestId('guide-leveling-row-talents')).toBeVisible();
        await expect(row.getByText('Load in planner')).toHaveAttribute(
          'href',
          /^\/planner\?spec=mage-frost&talents=/,
        );
        const chips = row.getByTestId('guide-leveling-row-talents').locator('li');
        expect(await chips.count()).toBeGreaterThan(3);
        await expect(chips.first().locator('img')).toHaveCount(1);
        await expect(chips.first()).toContainText(/\d\/\d/);
      }
      await expect(page.getByTestId('guide-leveling-legend')).toHaveText(
        'Outlined: points added in this level range.',
      );
      await expect(strip.getByTestId('guide-leveling-row-20')).toContainText('Level 20 to 29');
      await expect(strip.getByTestId('guide-leveling-row-20')).toContainText('11 to 20 points');
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

  test('desktop: the rail marks the section you are in', async ({ page }) => {
    await page.setViewportSize(DESKTOP);
    await page.goto(PAGE);
    const rail = page.getByTestId('guide-toc-rail');
    const scrollTo = (id: string) =>
      page.evaluate(
        (target) =>
          window.scrollTo(
            0,
            (document.getElementById(target)?.getBoundingClientRect().top ?? 0) + window.scrollY - 80,
          ),
        id,
      );
    await scrollTo('stat-priority');
    await expect(rail.locator('[aria-current="location"]')).toHaveText('Stat priority');
    await scrollTo('races');
    await expect(rail.locator('[aria-current="location"]')).toHaveText('Races');
    await expect(rail.locator('[aria-current="location"]')).toHaveCount(1);
  });

  test('phone: Rotation and Stat priority sit under the Build card with no disclosure', async ({ page }) => {
    await page.setViewportSize(PHONE);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-rail-fold-toggle')).toHaveCount(0);
    await expect(page.getByTestId('guide-rail-rotation')).toBeVisible();
    await expect(page.getByTestId('guide-rail-stat-priority')).toBeVisible();
  });

  test('desktop: the hero shows all three cards', async ({ page }) => {
    await page.setViewportSize(DESKTOP);
    await page.goto(PAGE);
    await expect(page.getByTestId('guide-rail-rotation')).toBeVisible();
    await expect(page.getByTestId('guide-rail-stat-priority')).toBeVisible();
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
