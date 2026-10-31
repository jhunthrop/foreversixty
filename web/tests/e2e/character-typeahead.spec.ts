// web/tests/e2e/character-typeahead.spec.ts
// The suggestion list under both character lookup boxes (design/specs/2026-10-15-character-
// typeahead.md): the home page's full box, the header's compact one on the bar and in the
// phone Menu. GET /v1/characters/search is routed to a fixture of three rows -- two with a
// class and level, one with nulls -- and to nothing for "zz".
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { assertNoHorizontalScroll } from './support/phone-scroll';
import { CHARACTER_PATH, stubCharacter } from './support/character-fixture';

const ROWS = [
  {
    key: 'us/hardcore/elyra-duskvale',
    region: 'us',
    ruleset: 'hardcore',
    name: 'Elyra Duskvale',
    class: 'Priest',
    level: 60,
    spec: 'Shadow',
  },
  {
    key: 'us/normal/bolt',
    region: 'us',
    ruleset: 'normal',
    name: 'Bolt',
    class: 'Mage',
    level: 58,
    spec: 'Fire',
  },
  { key: 'eu/pvp/bo', region: 'eu', ruleset: 'pvp', name: 'Bo', class: null, level: null, spec: null },
];

async function stubSearch(page: Page): Promise<string[]> {
  const asked: string[] = [];
  await page.route('**/v1/characters/search*', (route) => {
    const url = new URL(route.request().url());
    asked.push(url.search);
    const rows = url.searchParams.get('q') === 'zz' ? [] : ROWS;
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { rows }, error: null, request_id: 'r' }),
    });
  });
  return asked;
}

/** Typing before the page's script has wired the form reaches no listener. */
async function wired(page: Page): Promise<void> {
  await expect(page.locator('form[data-character-search]:not([data-typeahead-wired])')).toHaveCount(0);
}

const options = (scope: Locator): Locator => scope.locator('a[role="option"]');

async function expectThreeRows(scope: Locator): Promise<void> {
  const rows = options(scope);
  await expect(rows).toHaveCount(3);
  await expect(rows.nth(0).locator('img')).toHaveAttribute('src', '/icons/hd/crests/priest.webp');
  await expect(rows.nth(0)).toContainText('Elyra Duskvale');
  await expect(rows.nth(0)).toContainText('Shadow Priest · 60');
  await expect(rows.nth(0)).toContainText('US Hardcore');
  await expect(rows.nth(1).locator('img')).toHaveAttribute('src', '/icons/hd/crests/mage.webp');
  await expect(rows.nth(2).locator('img')).toHaveCount(0);
  await expect(rows.nth(2)).toContainText('EU PvP');
  for (const row of await rows.all()) expect((await row.boundingBox())!.height).toBeGreaterThanOrEqual(44);
}

test.describe('the home page box', () => {
  test('lists the matches and ArrowDown + Enter opens the first', async ({ page }) => {
    const asked = await stubSearch(page);
    await stubCharacter(page);
    await page.goto('/');
    await wired(page);
    const box = page.getByTestId('home-character-search');
    const input = box.getByRole('combobox', { name: 'Character name' });
    await input.fill('bo');

    await expectThreeRows(box);
    await expect(input).toHaveAttribute('aria-expanded', 'true');
    expect(asked[0]).not.toContain('region=');
    expect(asked[0]).not.toContain('ruleset=');

    await input.press('ArrowDown');
    await expect(input).toHaveAttribute('aria-activedescendant', /.+/);
    await input.press('Enter');
    await expect(page).toHaveURL(new RegExp(`${CHARACTER_PATH}$`));
  });

  test('says so when nothing matches, and Enter still looks the typed name up', async ({ page }) => {
    await stubSearch(page);
    await page.goto('/');
    await wired(page);
    const box = page.getByTestId('home-character-search');
    const input = box.getByRole('combobox', { name: 'Character name' });
    await input.fill('zz');
    await expect(box.getByRole('listbox').getByRole('option')).toHaveText('No character named "zz" yet');
    await input.press('Enter');
    await expect(page).toHaveURL(/\/character\/us\/normal\/zz$/);
  });

  test('Escape closes the list', async ({ page }) => {
    await stubSearch(page);
    await page.goto('/');
    await wired(page);
    const box = page.getByTestId('home-character-search');
    await box.getByRole('combobox', { name: 'Character name' }).fill('bo');
    await expect(options(box)).toHaveCount(3);
    await box.getByRole('combobox', { name: 'Character name' }).press('Escape');
    await expect(box.getByRole('listbox')).toBeHidden();
  });
});

test.describe('the header lookup', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'bar layout only');

  test('opens from the bar, lists the matches and a click opens the second', async ({ page }) => {
    await stubSearch(page);
    await page.setViewportSize({ width: 1280, height: 900 });
    await page.goto('/guides');
    await wired(page);
    await page.getByTestId('site-search-toggle').click();
    const panel = page.locator('[data-header-search-panel]');
    const input = page.getByTestId('site-search-input');
    await input.fill('bo');

    await expectThreeRows(panel);
    const form = await panel.boundingBox();
    const list = await panel.locator('[role="listbox"]').boundingBox();
    expect(list!.y).toBeGreaterThan(form!.y);

    await options(panel).nth(1).click();
    await expect(page).toHaveURL(/\/character\/us\/normal\/bolt$/);
  });

  test('searches every realm: no region or ruleset filter', async ({ page }) => {
    const asked = await stubSearch(page);
    await page.setViewportSize({ width: 1280, height: 900 });
    await page.goto('/guides');
    await wired(page);
    await page.getByTestId('site-search-toggle').click();
    await page.getByTestId('site-search-input').fill('bo');
    await expect(options(page.locator('[data-header-search-panel]'))).toHaveCount(3);
    expect(asked[0]).not.toContain('region=');
    expect(asked[0]).not.toContain('ruleset=');
  });
});

test.describe('the phone Menu', () => {
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');
  test.use({ viewport: { width: 390, height: 844 } });

  test('lists the matches without scrolling sideways', async ({ page }) => {
    await stubSearch(page);
    await page.goto('/');
    await wired(page);
    await page.getByTestId('menu-button').click();
    const menuBox = page.locator('.site-menu-search');
    await page.getByTestId('site-menu-search-input').fill('bo');
    await expectThreeRows(menuBox);
    await assertNoHorizontalScroll(page, 390);
    const list = await menuBox.locator('[role="listbox"]').boundingBox();
    expect(list!.x + list!.width).toBeLessThanOrEqual(390);
  });
});
