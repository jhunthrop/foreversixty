// web/tests/e2e/paperdoll.spec.ts
// The shared paperdoll on the sim's Gear section (`worn`, spec 2026-11-06-sim §4.5, open after
// a run) and the report's "At pull" (`report`), web layout spec 2026-11-05 §4.2: every item opens the shared tooltip on hover,
// focus and tap, and the tooltip carries the binding and armour lines the item row states. The
// character page's own coverage is character-gear.spec.ts.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { pasteAndRun } from './support/sim-page';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

/** The fixture's Lionheart Helm (plate, binds when equipped) and two-handed Arcanite Reaper. */
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784|who=Simfury:Realm`;

async function loadSim(page: Page): Promise<void> {
  // The paste runs a sim at once; wait it out so the page is at rest before it is touched.
  await pasteAndRun(page, FURY);
  // The result's Gear section loads as it nears the viewport.
  await page.getByTestId('sim-section-gear').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-gear-doll-slot-head')).toHaveAttribute('data-state', 'item');
}

/** The result's Gear doll. */
const strip = (page: Page) => page.getByTestId('sim-gear-doll');

test.describe('the sim card at 1280', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  test('draws the doll and opens the real item on hover and focus', async ({ page }) => {
    const errors: string[] = [];
    page.on('pageerror', (error) => errors.push(error.message));
    await loadSim(page);
    await expect(page.getByTestId('page-title')).toHaveText('Simfury');
    await expect(page.getByTestId('sim-gear-doll-slot-neck')).toHaveAttribute('data-state', 'empty');
    await expect(page.getByTestId('sim-gear-doll-slot-neck')).toContainText('Empty');
    await expect(page.getByTestId('sim-gear-doll-slot-off_hand')).toContainText('Two-hand in main hand');
    const head = strip(page).getByTestId('item-hover-12640');
    await expect(head).toHaveText('Lionheart Helm');
    await head.hover();
    const tip = page.getByTestId('item-tooltip');
    await expect(tip).toBeVisible();
    await expect(tip).toContainText('Lionheart Helm');
    await expect(tip.getByTestId('item-tooltip-bind')).toHaveText('Binds when equipped');
    await expect(tip).toContainText('Plate');
    await page.mouse.move(0, 0);
    await page.getByTestId('sim-gear-doll-slot-main_hand').focus();
    await expect(tip).toContainText('Arcanite Reaper');
    await expect(tip).toHaveCount(1);
    expect(errors).toEqual([]);
  });
});

test.describe('the sim card on a 390 phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, hasTouch: true, isMobile: true });

  test('a tap opens the item, a second tap closes it', async ({ page }) => {
    await loadSim(page);
    // loadSim's Run is a mouse click, and the card can land under where that mouse rests: its
    // hover would open the panel before the first tap. Park the mouse off the card first.
    await page.mouse.move(0, 0);
    const head = strip(page).getByTestId('item-hover-12640');
    await head.tap();
    await expect(page.getByTestId('item-tooltip')).toBeVisible();
    await head.tap();
    await expect(page.getByTestId('item-tooltip')).toHaveCount(0);
  });
});

test.describe('the report "At pull" at 1280', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  test('opens a combatant’s gear as the paperdoll, each slot with its log item level', async ({ page }) => {
    await page.route('**/fights/3/summary.json*', async (route) => {
      const response = await route.fetch();
      const summary = (await response.json()) as { combatants: { gear: unknown[] }[] };
      const piece = (ID: number, ItemLevel: number, Enchants: number[] | null = null) => ({
        ID,
        ItemLevel,
        Enchants,
        BonusIDs: null,
        Gems: null,
      });
      summary.combatants[0]!.gear = [
        piece(12640, 66, [1900]),
        ...Array.from({ length: 18 }, () => piece(0, 0)),
      ];
      summary.combatants[0]!.gear[15] = piece(12784, 63);
      await route.fulfill({ response, json: summary });
    });
    await page.goto('/reports/fixture2abcd?fight=3');
    await page.getByTestId('combatant-gear').first().click();
    const frame = page.getByTestId('gear-list');
    await expect(frame.getByTestId('paperdoll-slot-head')).toHaveAttribute('data-state', 'item');
    await expect(frame.getByTestId('paperdoll-line2-head')).toHaveText('Item level 66 · Enchanted');
    await expect(frame.getByTestId('paperdoll-slot-head').locator('.ui-doll-figure')).toHaveCount(0);
    // The worn set's average item level sits above the doll (owner 2026-11-06).
    await expect(page.getByTestId('gear-list-ilvl')).toHaveText(/^Item level \d+(\.\d)?$/);
    await frame.getByTestId('item-hover-12640').hover();
    await expect(page.getByTestId('item-tooltip')).toContainText('Lionheart Helm');
  });
});
