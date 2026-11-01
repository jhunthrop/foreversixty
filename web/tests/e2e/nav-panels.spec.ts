// web/tests/e2e/nav-panels.spec.ts
// The grouped header (design/specs/2026-10-15-grouped-header.md): four doors, each opening a
// hover panel that lists its pages, with hover delays, the chevron toggle, the keyboard walk,
// the touch first tap, the Guides crest grid, the guild rows following the session, and the
// same four groups in the phone Menu. Every width keeps the no-sideways-scroll rule.
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { guildHref } from '../../src/lib/characters';
import { headerWired } from './support/header';
import { OBNOXIOUS, armoryPointer, signInWith, storePointer } from './support/selector';

const DOOR_KEYS = ['tools', 'logs', 'guild', 'guides'] as const;
const GUILD = { id: 1, region: 'us', ruleset: 'normal', name: 'OLYMPUS XXVII', verified: true };
const HOVER_OPEN_WAIT_MS = 400;
const LEAVE_CLOSE_BUDGET_MS = 300;
const BAR_HEIGHT_PX = 81;

const door = (page: Page, key: string): Locator => page.locator(`a.site-door[data-door="${key}"]`);
const panel = (page: Page, key: string): Locator => page.locator(`#site-panel-${key}`);

async function signedInOnGuides(page: Page, width: number, guilds: (typeof GUILD)[] = []) {
  // The Guild menu follows the current character's own guild line: Obnoxious is in the first.
  const guildLine =
    guilds[0] === undefined ? {} : { guild: { id: guilds[0].id, name: guilds[0].name, verified: true } };
  await signInWith(page, [{ ...OBNOXIOUS, ...guildLine }], guilds);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await page.setViewportSize({ width, height: 900 });
  await page.goto('/guides');
  await expect(page.getByTestId('character-selector')).toBeVisible();
}

const noSideScroll = (page: Page): Promise<boolean> =>
  page.evaluate(() => document.documentElement.scrollWidth === document.documentElement.clientWidth);

const doorBoxes = (page: Page) =>
  page
    .locator('a.site-door')
    .evaluateAll((links) => links.map((link) => link.getBoundingClientRect().toJSON()));

for (const width of [1280, 1440]) {
  test.describe(`hover panels at ${width}`, () => {
    test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

    test('each door opens its panel under it and nothing in the bar moves', async ({ page }) => {
      await signedInOnGuides(page, width);
      const barBefore = await page.getByRole('banner').boundingBox();
      const doorsBefore = await doorBoxes(page);
      for (const key of DOOR_KEYS) {
        await door(page, key).hover();
        await expect(panel(page, key)).toBeVisible();
        await expect(door(page, key)).toHaveAttribute('aria-expanded', 'true');
        expect(await noSideScroll(page)).toBe(true);
        expect(await page.getByRole('banner').boundingBox()).toEqual(barBefore);
        expect(await doorBoxes(page)).toEqual(doorsBefore);
        const box = (await panel(page, key).boundingBox())!;
        expect(box.y).toBeGreaterThanOrEqual(BAR_HEIGHT_PX - 1);
        expect(box.x).toBeGreaterThanOrEqual(0);
        expect(box.x + box.width).toBeLessThanOrEqual(width);
      }
    });

    test('the Tools panel lists its three rows in order, each 44px or taller', async ({ page }) => {
      await signedInOnGuides(page, width);
      await door(page, 'tools').hover();
      await expect(panel(page, 'tools')).toBeVisible();
      const rows = panel(page, 'tools').locator('a.site-panel-row');
      await expect(rows.locator('.site-panel-title')).toHaveText(['Planner', 'BiS', 'Simulator']);
      for (const row of await rows.all())
        expect((await row.boundingBox())!.height).toBeGreaterThanOrEqual(44);
    });

    test('moving to another door closes the first panel, leaving the header closes it', async ({ page }) => {
      await signedInOnGuides(page, width);
      await door(page, 'tools').hover();
      await expect(panel(page, 'tools')).toBeVisible();
      await door(page, 'logs').hover();
      await expect(panel(page, 'logs')).toBeVisible({ timeout: HOVER_OPEN_WAIT_MS });
      await expect(panel(page, 'tools')).toBeHidden();
      await page.mouse.move(width / 2, 600);
      await expect(panel(page, 'logs')).toBeHidden({ timeout: LEAVE_CLOSE_BUDGET_MS });
    });

    test('the chevron opens the panel without navigating and a second click closes it', async ({ page }) => {
      await signedInOnGuides(page, width);
      const url = page.url();
      const toggle = page.getByRole('button', { name: 'Open the Logs list' });
      await toggle.click();
      await expect(panel(page, 'logs')).toBeVisible();
      expect(page.url()).toBe(url);
      await toggle.click();
      await expect(panel(page, 'logs')).toBeHidden();
    });

    test('a click on the door label still goes to the door page', async ({ page }) => {
      await signedInOnGuides(page, width);
      await door(page, 'logs').click();
      await expect(page).toHaveURL(/\/logs$/);
    });

    test('the Guides panel shows the nine crests, Warrior first, and All guides', async ({ page }) => {
      await signedInOnGuides(page, width);
      await door(page, 'guides').hover();
      const tiles = panel(page, 'guides').locator('a.site-panel-class');
      await expect(tiles).toHaveCount(9);
      await expect(tiles.first()).toHaveAttribute('href', '/guides/warrior');
      await expect(tiles.first().locator('img')).toHaveAttribute('src', '/icons/hd/crests/warrior.webp');
      await expect(tiles.first()).toContainText('Arms · Fury · Protection');
      await expect(panel(page, 'guides').getByRole('menuitem', { name: 'All guides' })).toHaveAttribute(
        'href',
        '/guides',
      );
    });

    test('the Guild rows name the session guild and link to its tabs', async ({ page }) => {
      await signedInOnGuides(page, width, [GUILD]);
      const base = guildHref(GUILD.region, GUILD.ruleset, GUILD.name);
      await door(page, 'guild').hover();
      const rows = panel(page, 'guild').locator('a.site-panel-row');
      await expect(rows.first()).toContainText('OLYMPUS XXVII');
      await expect(rows.first()).toHaveAttribute('href', base);
      await expect(panel(page, 'guild').getByTestId('site-panel-row-roster')).toHaveAttribute(
        'href',
        `${base}#roster`,
      );
      await expect(panel(page, 'guild').getByTestId('site-panel-foot-guild')).toHaveAttribute(
        'href',
        `${base}/settings`,
      );
    });
  });
}

test.describe('keyboard', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('ArrowDown opens and focuses the first row, Escape returns to the door', async ({ page }) => {
    await signedInOnGuides(page, 1280);
    await door(page, 'logs').focus();
    await page.keyboard.press('ArrowDown');
    await expect(panel(page, 'logs')).toBeVisible();
    await expect(page.getByTestId('site-panel-row-upload-a-log')).toBeFocused();
    await page.keyboard.press('ArrowDown');
    await page.keyboard.press('ArrowDown');
    await expect(page.getByTestId('site-panel-row-rankings')).toBeFocused();
    await page.keyboard.press('Escape');
    await expect(panel(page, 'logs')).toBeHidden();
    await expect(door(page, 'logs')).toBeFocused();
  });

  test('Tab out of the panel closes it', async ({ page }) => {
    await signedInOnGuides(page, 1280);
    await door(page, 'tools').focus();
    await page.keyboard.press('ArrowDown');
    for (let step = 0; step < 4; step += 1) await page.keyboard.press('Tab');
    await expect(panel(page, 'tools')).toBeHidden();
  });
});

test.describe('the current page', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1280, height: 900 } });

  test('a signed-out visitor on /tiers sees Logs current and /guild linked everywhere', async ({ page }) => {
    await page.route('**/v1/me', (route) =>
      route.fulfill({ status: 401, contentType: 'application/json', body: '{"ok":false,"data":null}' }),
    );
    await page.goto('/tiers');
    await expect(door(page, 'logs')).toHaveAttribute('aria-current', 'page');
    await expect(page.getByTestId('site-panel-row-tier-list')).toHaveAttribute('href', '/tiers');
    await expect(page.locator('#site-panel-guild a[href^="/guild"]')).toHaveCount(5);
    for (const link of await page.locator('#site-panel-guild a:not([href^="https:"])').all()) {
      await expect(link).toHaveAttribute('href', /^\/guild(#\w+)?$/);
    }
  });
});

test.describe('touch', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.use({ viewport: { width: 1280, height: 900 }, hasTouch: true, isMobile: true });

  test('the first tap on Guild opens the panel and the second navigates', async ({ page }) => {
    await page.goto('/');
    await headerWired(page);
    expect(await page.evaluate(() => matchMedia('(hover: none)').matches)).toBe(true);
    await door(page, 'guild').tap();
    await expect(panel(page, 'guild')).toBeVisible();
    await door(page, 'guild').tap();
    await expect(page).toHaveURL(/\/guild$/);
  });
});

test.describe('the phone Menu', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');
  test.use({ viewport: { width: 390, height: 844 } });

  test('lists Tools, Logs, Guild and Guides in order with 44px rows', async ({ page }) => {
    await page.goto('/');
    await page.getByTestId('menu-button').click();
    await expect(page.locator('#site-menu h2.site-menu-heading')).toHaveText([
      'Tools',
      'Logs',
      'Guild',
      'Guides',
    ]);
    for (const row of await page.locator('#site-menu .site-menu-list a.site-panel-row').all()) {
      await expect(row).toBeVisible();
      expect((await row.boundingBox())!.height).toBeGreaterThanOrEqual(44);
    }
    await expect(page.locator('#site-menu .site-menu-list a.site-panel-class')).toHaveCount(9);
    expect(await noSideScroll(page)).toBe(true);
    await page.getByTestId('site-menu-row-bis').click();
    await expect(page).toHaveURL(/\/bis$/);
  });
});

test.describe('no sideways scroll at the six widths', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  for (const width of [1024, 1279, 1280, 1439, 1440, 2000]) {
    test(`${width}px with every panel open`, async ({ page }) => {
      await signedInOnGuides(page, width, [GUILD]);
      for (const key of DOOR_KEYS) {
        await door(page, key).hover();
        await expect(panel(page, key)).toBeVisible();
        expect(await noSideScroll(page)).toBe(true);
        const box = (await panel(page, key).boundingBox())!;
        expect(box.x).toBeGreaterThanOrEqual(0);
        expect(box.x + box.width).toBeLessThanOrEqual(width);
      }
    });
  }
});
