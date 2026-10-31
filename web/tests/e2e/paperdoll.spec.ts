// web/tests/e2e/paperdoll.spec.ts
// The shared paperdoll on the sim card and the report's "At pull" (design/specs/2026-10-31-
// paperdoll.md, call sites 1 and 3): every item opens the shared tooltip on hover, focus and
// tap, and the tooltip carries the binding and armour lines the item row states. The character
// page's own coverage is character-gear.spec.ts.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

/** The fixture's Lionheart Helm (plate, binds when equipped) and two-handed Arcanite Reaper. */
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784|who=Simfury:Realm`;

async function loadSim(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('paperdoll-slot-head')).toHaveAttribute('data-state', 'item');
  // The paste runs a sim at once; wait it out so the card is at rest before it is touched.
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
}

test.describe('the sim card at 1280', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  test('draws the frame and opens the real item on hover and focus', async ({ page }) => {
    const errors: string[] = [];
    page.on('pageerror', (error) => errors.push(error.message));
    await loadSim(page);
    await expect(page.getByTestId('paperdoll-name')).toHaveText('Simfury');
    await expect(page.getByTestId('paperdoll-slot-neck')).toHaveAttribute('data-state', 'empty');
    const head = page.getByTestId('paperdoll-item-head');
    await expect(head).toHaveAttribute('aria-label', 'Head: Lionheart Helm');
    await head.hover();
    const tip = page.getByTestId('item-tooltip');
    await expect(tip).toBeVisible();
    await expect(tip).toContainText('Lionheart Helm');
    await expect(tip.getByTestId('item-tooltip-bind')).toHaveText('Binds when equipped');
    await expect(tip).toContainText('Plate');
    await page.mouse.move(0, 0);
    await page.getByTestId('paperdoll-item-main_hand').focus();
    await expect(tip).toContainText('Arcanite Reaper');
    expect(errors).toEqual([]);
  });
});

test.describe('the sim card on a 390 phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, hasTouch: true, isMobile: true });

  test('a tap opens the item, a second tap closes it', async ({ page }) => {
    await loadSim(page);
    const head = page.getByTestId('paperdoll-item-head');
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
    await expect(frame.getByTestId('paperdoll-source')).toHaveText('At pull');
    await expect(frame.getByTestId('paperdoll-slot-head').getByTestId('paperdoll-detail')).toHaveText(
      'Item level 66 · Enchanted',
    );
    await frame.getByTestId('paperdoll-item-head').hover();
    await expect(page.getByTestId('item-tooltip')).toContainText('Lionheart Helm');
  });
});
