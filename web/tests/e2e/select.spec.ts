// web/tests/e2e/select.spec.ts
// The shared dropdown in a real browser (design/specs/2026-10-30-controls-charts.md §A, §C, §D
// captures 1 to 8): on a fine pointer our own listbox in the top layer, placed by Floating UI,
// with the keyboard and focus rules of §A6; on a phone the native picker behind the same face.
// The gallery on /design-sheet holds one live field of every variant; the report, sim and home
// pages are the representative call sites.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';
import { enhanced, fieldBox, listOf } from './support/select';

const SHEET = '/design-sheet';
const THREAT = '/reports/fixture2abcd?fight=3&tab=threat';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };
/** sim-settings.spec.ts's Fury export: the settings bar draws once a character is loaded. */
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

const isPhone = (page: Page): boolean => (page.viewportSize()?.width ?? 1280) < 768;

async function openSheet(page: Page): Promise<void> {
  // The dialog button stays visible whichever control a field shows, so it marks hydration.
  await gotoHydrated(page, SHEET, 'sheet-dialog-open');
}

const row = (list: Locator, label: string): Locator =>
  list.locator('[role="option"]').filter({ has: list.page().locator('.sel-row-label', { hasText: label }) });

async function activeLabel(page: Page, holder: Locator): Promise<string> {
  const id = await holder.getAttribute('aria-activedescendant');
  return id === null ? '' : ((await page.locator(`[id="${id}"] .sel-row-label`).textContent()) ?? '');
}

/** True when the topmost element at the centre of each of `targets` is inside `owner`. Rows the
 *  list has scrolled out of its own view are skipped: only the rows a person can see count. */
async function topmostInside(page: Page, targets: Locator, owner: string): Promise<boolean[]> {
  return targets.evaluateAll(
    (nodes, selector) =>
      nodes
        .map((node) => {
          const box = node.getBoundingClientRect();
          const scroller = node.closest('.sel-scroll')?.getBoundingClientRect();
          const y = box.y + box.height / 2;
          const shown =
            box.height > 0 &&
            y > 0 &&
            y < window.innerHeight &&
            (scroller === undefined || (y > scroller.top && y < scroller.bottom));
          return { x: box.x + box.width / 2, y, shown };
        })
        .filter((point) => point.shown)
        .map((point) => document.elementFromPoint(point.x, point.y)?.closest(selector) !== null),
    owner,
  );
}

test.describe('fine pointer: the listbox', () => {
  test.beforeEach(async ({ page }) => {
    test.skip(isPhone(page), 'the listbox is the fine-pointer control; phones keep the native picker');
    // The list rises 2 px as it fades in; still it, so boxes measure where it settles.
    await page.emulateMedia({ reducedMotion: 'reduce' });
  });

  test('the swap from the native field to the listbox moves the field by 0 px', async ({ page }) => {
    await page.addInitScript(() => {
      document.addEventListener('DOMContentLoaded', () => {
        const box = document.querySelector('#sheet-fight-style')?.closest('.sel')?.getBoundingClientRect();
        (window as unknown as { before: unknown }).before = box?.toJSON();
      });
    });
    await openSheet(page);
    const trigger = await enhanced(page, 'sheet-fight-style');
    const before = await page.evaluate(() => (window as unknown as { before: DOMRect }).before);
    const after = await fieldBox(page, 'sheet-fight-style').evaluate((node) =>
      node.getBoundingClientRect().toJSON(),
    );
    expect(after).toEqual(before);
    const face = await fieldBox(page, 'sheet-fight-style').locator('.sel-face').boundingBox();
    expect(await trigger.boundingBox()).toEqual(face);
  });

  test('a plain list: Down opens it, arrows and typeahead move, focus stays on the trigger, Enter commits', async ({
    page,
  }) => {
    await openSheet(page);
    const trigger = await enhanced(page, 'sheet-enemy');
    await trigger.focus();
    expect(await trigger.evaluate((node) => node.matches(':focus-visible'))).toBe(true);
    await page.keyboard.press('ArrowDown');
    await expect(trigger).toHaveAttribute('aria-expanded', 'true');
    await expect(trigger).toHaveAttribute('role', 'combobox');
    expect(await activeLabel(page, trigger)).toBe('Every enemy');
    await page.keyboard.press('ArrowDown');
    expect(await activeLabel(page, trigger)).toBe('Ragnaros');
    await page.keyboard.press('s');
    expect(await activeLabel(page, trigger)).toBe('Son of Flame');
    await expect(trigger).toBeFocused();
    await page.keyboard.press('Enter');
    await expect(trigger).toHaveAttribute('aria-expanded', 'false');
    await expect(page.getByTestId('sheet-enemy')).toHaveValue('sons');
    await expect(fieldBox(page, 'sheet-enemy').locator('.sel-label')).toHaveText('Son of Flame');
    await expect(trigger).toBeFocused();
  });

  test('a searchable list moves focus into its filter, the one combobox, and gives it back on close', async ({
    page,
  }) => {
    await openSheet(page);
    const trigger = await enhanced(page, 'sheet-where');
    await trigger.click();
    const list = listOf(page, 'sheet-where');
    const filter = list.locator('input[role="combobox"]');
    await expect(filter).toBeFocused();
    await expect(page.locator('[role="combobox"][aria-expanded="true"]')).toHaveCount(1);
    await expect(trigger).not.toHaveAttribute('role', 'combobox');
    await page.keyboard.type('black');
    await expect(list.locator('.sel-count')).toHaveText('4 matches');
    await expect(list.locator('.sel-grp-head')).toHaveText(['Raids', 'Dungeons']);
    await expect(list.locator('.sel-match').first()).toHaveText('Black');
    await page.keyboard.press('Escape');
    await expect(filter).toHaveValue('');
    await page.keyboard.press('Escape');
    await expect(list).toHaveCount(0);
    await expect(trigger).toBeFocused();
  });

  test('no match shows the no-match row', async ({ page }) => {
    await openSheet(page);
    await (await enhanced(page, 'sheet-where')).click();
    await page.keyboard.type('naxx');
    await expect(listOf(page, 'sheet-where').locator('.sel-empty')).toHaveText('Nothing matches “naxx”.');
  });

  test('a pointer parked on a row never commits it: Tab commits the keyboard-active row', async ({
    page,
  }) => {
    await openSheet(page);
    const trigger = await enhanced(page, 'sheet-enemy');
    // Open once to learn where Son of Flame draws, then park the mouse there and reopen by key.
    await trigger.click();
    const sons = await row(listOf(page, 'sheet-enemy'), 'Son of Flame').boundingBox();
    await page.keyboard.press('Escape');
    await page.mouse.move(sons!.x + sons!.width / 2, sons!.y + sons!.height / 2);
    await trigger.focus();
    await page.keyboard.press('ArrowDown');
    expect(await activeLabel(page, trigger)).toBe('Every enemy');
    await page.keyboard.press('Tab');
    await expect(page.getByTestId('sheet-enemy')).toHaveValue('all');
  });

  test('a click commits the clicked row; the multi-select toggles and stays open', async ({ page }) => {
    await openSheet(page);
    await (await enhanced(page, 'sheet-item')).click();
    await row(listOf(page, 'sheet-item'), 'Cruelty').click();
    await expect(page.getByTestId('sheet-item')).toHaveValue('c');

    const players = await enhanced(page, 'sheet-players');
    await expect(fieldBox(page, 'sheet-players').locator('.sel-label')).toHaveText(/^Top 8 of/);
    await players.click();
    const list = listOf(page, 'sheet-players');
    await expect(list.locator('[role="listbox"]')).toHaveAttribute('aria-multiselectable', 'true');
    await row(list, 'Thandril').click();
    await expect(list).toBeVisible();
    await expect(row(list, 'Thandril')).toHaveAttribute('aria-selected', 'true');
    await expect(list.locator('.sel-foot-count')).toHaveText('4 drawn');
    await expect(fieldBox(page, 'sheet-players').locator('.sel-label')).toHaveText(/^4 of/);
    await list.getByRole('button', { name: 'Top 8' }).click();
    await expect(list.locator('.sel-foot-count')).toHaveText('3 drawn');
    await page.keyboard.press('Escape');
    await expect(list).toHaveCount(0);
  });

  test('axe finds nothing with a plain, a searchable and the multi-select list open', async ({ page }) => {
    await openSheet(page);
    for (const testid of ['sheet-fight-style', 'sheet-source', 'sheet-players']) {
      await (await enhanced(page, testid)).click();
      await expect(listOf(page, testid)).toBeVisible();
      await expect(page.locator('[role="combobox"][aria-expanded="true"]')).toHaveCount(1);
      await expectNoAxeViolations(page, `.sel:has(> select[data-testid="${testid}"])`);
      await page.keyboard.press('Escape');
    }
  });

  test('the Source list opens scrolled to the picked player, crests and class colours on', async ({
    page,
  }) => {
    await openSheet(page);
    await (await enhanced(page, 'sheet-source')).click();
    const list = listOf(page, 'sheet-source');
    const picked = row(list, 'Morrigan');
    await expect(picked).toHaveAttribute('aria-selected', 'true');
    await expect(picked.locator('.sel-check')).toBeVisible();
    await expect(picked.locator('.sel-icon[data-kind="crest"] img')).toHaveAttribute('src', /crests\/mage/);
    await expect(row(list, 'Lightningbolltz').locator('.sel-you')).toHaveText('You');
    const scroller = list.locator('.sel-scroll');
    const inView = await picked.evaluate(
      (node, scrollBox) => {
        const a = node.getBoundingClientRect();
        const b = (scrollBox as Element).getBoundingClientRect();
        return a.top >= b.top && a.bottom <= b.bottom;
      },
      await scroller.elementHandle(),
    );
    expect(inView).toBe(true);
  });

  test('the list flips above a field near the viewport bottom; a right-slot field opens bottom-end', async ({
    page,
  }) => {
    await openSheet(page);
    const trigger = await enhanced(page, 'sheet-enemy');
    // Put the field 60 px above the viewport's bottom edge.
    await trigger.evaluate((node) => {
      const box = node.getBoundingClientRect();
      window.scrollBy(0, box.bottom - window.innerHeight + 60);
    });
    await trigger.click();
    const list = listOf(page, 'sheet-enemy');
    await expect(list).toHaveAttribute('data-side', 'top');
    const field = (await fieldBox(page, 'sheet-enemy').boundingBox())!;
    const above = (await list.boundingBox())!;
    expect(Math.abs(field.y - (above.y + above.height) - 4)).toBeLessThanOrEqual(1);
    await page.keyboard.press('Escape');

    await trigger.evaluate(() => window.scrollTo(0, 0));
    const players = await enhanced(page, 'sheet-players');
    await players.scrollIntoViewIfNeeded();
    await players.evaluate((node) => window.scrollBy(0, node.getBoundingClientRect().top - 120));
    await players.click();
    const playersList = listOf(page, 'sheet-players');
    await expect(playersList).toHaveAttribute('data-side', 'bottom');
    const box = (await fieldBox(page, 'sheet-players').boundingBox())!;
    const open = (await playersList.boundingBox())!;
    expect(Math.abs(open.x + open.width - (box.x + box.width))).toBeLessThanOrEqual(1);
    expect(Math.abs(open.y - (box.y + box.height) - 4)).toBeLessThanOrEqual(1);
  });

  test('inside a modal dialog the list draws above it, and Escape closes only the list', async ({ page }) => {
    await openSheet(page);
    await page.getByTestId('sheet-dialog-open').click();
    const dialog = page.getByTestId('sheet-dialog');
    await expect(dialog).toBeVisible();
    const trigger = await enhanced(dialog, 'sheet-dialog-style');
    await trigger.click();
    const list = listOf(page, 'sheet-dialog-style');
    await expect(list).toBeVisible();
    expect(await topmostInside(page, list.locator('[role="option"]'), '.sel-pop')).not.toContain(false);
    await row(list, 'Execute').hover();
    await page.keyboard.press('Escape');
    await expect(list).toHaveCount(0);
    await expect(dialog).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(dialog).toBeHidden();
  });

  test('one open overlay at a time: opening a list closes the other', async ({ page }) => {
    await openSheet(page);
    await (await enhanced(page, 'sheet-enemy')).click();
    await expect(listOf(page, 'sheet-enemy')).toBeVisible();
    await (await enhanced(page, 'sheet-item')).click();
    await expect(listOf(page, 'sheet-enemy')).toHaveCount(0);
    await expect(listOf(page, 'sheet-item')).toBeVisible();
  });
});

test.describe('fine pointer: call sites', () => {
  test.beforeEach(({ page }) => {
    test.skip(isPhone(page), 'the listbox is the fine-pointer control');
  });

  // The owner's bug, 2026-10-30: a menu disappeared under the threat chart. Every list is in the
  // top layer now, so the topmost element at each row's centre is that row.
  test('an open list is topmost over the report chart', async ({ page }) => {
    await page.goto(THREAT);
    await expect(page.getByTestId('threat-chart')).toBeVisible();
    const trigger = await enhanced(page, 'source-scope');
    await trigger.click();
    const list = listOf(page, 'source-scope');
    await expect(list).toBeVisible();
    const onTop = await topmostInside(page, list.locator('[role="option"]'), '.sel-pop');
    expect(onTop.length).toBeGreaterThan(0);
    expect(onTop).not.toContain(false);
  });

  test('the report Source picks a player through the list and the URL follows', async ({ page }) => {
    await page.goto(THREAT);
    const trigger = await enhanced(page, 'source-scope');
    await trigger.click();
    const list = listOf(page, 'source-scope');
    await expect(list.locator('.sel-grp-head').first()).toBeVisible();
    await row(list, 'Baelgrim').click();
    await expect(page).toHaveURL(/source=Player-4184-000000A1/);
    await expect(fieldBox(page, 'source-scope').locator('.sel-label')).toHaveText('Baelgrim');
  });

  test('the sim Fight style list carries each style’s description', async ({ page }) => {
    await page.goto('/sim');
    await page.getByTestId('sim-cold-paste-input').fill(FURY);
    await page.getByTestId('sim-cold-paste-run').click();
    await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
    const trigger = await enhanced(page, 'sim-style');
    await trigger.click();
    const list = listOf(page, 'sim-style');
    await expect(row(list, 'Execute').locator('.sel-row-desc')).toContainText('wider execute phase');
    await row(list, 'Execute').click();
    await expect(page.getByTestId('sim-style')).toHaveValue('execute');
  });
});

// Reviewer hold, 2026-10-30: the Fight style list of described rows ran off the bottom of a short
// screen. It flips above when it does not fit below and more room is above, and is capped to its
// side's room less 16 px either way, so every row is reachable inside the list.
test('on a short screen the Fight style list stays inside the viewport', async ({ page }) => {
  test.skip(isPhone(page), 'the listbox is the fine-pointer control');
  await page.setViewportSize({ width: 1440, height: 620 });
  await page.emulateMedia({ reducedMotion: 'reduce' });
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  const trigger = await enhanced(page, 'sim-style');
  for (const offset of [80, 260, 420]) {
    // The field `offset` px above the viewport's bottom edge.
    await trigger.evaluate(
      (node, by) => window.scrollBy(0, node.getBoundingClientRect().bottom - window.innerHeight + by),
      offset,
    );
    await trigger.click();
    const list = listOf(page, 'sim-style');
    await expect(list).toBeVisible();
    const box = (await list.boundingBox())!;
    expect(box.y, `top at ${offset}`).toBeGreaterThanOrEqual(0);
    expect(box.y + box.height, `bottom at ${offset}`).toBeLessThanOrEqual(620);
    // The last style is reachable by scrolling inside the list.
    await page.keyboard.press('End');
    await expect(list.locator('[role="option"]').last()).toBeInViewport();
    await page.keyboard.press('Escape');
    await expect(list).toHaveCount(0);
  }
});

test.describe('coarse pointer: the native picker behind the same face', () => {
  test.beforeEach(({ page }) => {
    test.skip(!isPhone(page), 'the phone project');
  });

  test('the native select stays, dense renders 44 px, and the face follows a pick', async ({ page }) => {
    await openSheet(page);
    const select = page.getByTestId('sheet-enemy');
    await expect(select).toBeVisible();
    await expect(page.getByTestId('sheet-enemy-trigger')).toHaveCount(0);
    expect((await fieldBox(page, 'sheet-enemy').boundingBox())?.height).toBe(44);
    await select.selectOption('rag');
    await expect(fieldBox(page, 'sheet-enemy').locator('.sel-label')).toHaveText('Ragnaros');
  });

  test('a static page’s field follows its select with no Svelte runtime', async ({ page }) => {
    await page.goto('/');
    const region = page.locator('#home-character-search-region');
    await expect(region).toBeVisible();
    await region.selectOption('eu');
    await expect(page.locator('.sel:has(> #home-character-search-region) .sel-label')).toHaveText('EU');
  });
});
