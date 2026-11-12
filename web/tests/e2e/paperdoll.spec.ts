// web/tests/e2e/paperdoll.spec.ts
// The shared paperdoll (ui/PaperdollFrame) on the sim's Gear section (`worn`: the frame with the
// selected slot's client tooltip in its centre, owner ruling 2026-11-11; its list form on a phone,
// whose items open the page's tooltip on tap) and the report's Gear at the pull (`report`, the same
// frame), web layout spec 2026-11-05 §4.2. The character page's own coverage is
// character-gear.spec.ts.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { expectSimGearSlot, pasteAndRun } from './support/sim-page';

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
  await expectSimGearSlot(page, 'head');
}

/** The result's Gear doll, one of its frame sockets and one of its list rows. */
const doll = (page: Page) => page.getByTestId('sim-gear-doll');
const socket = (page: Page, name: string) => doll(page).getByTestId(`paperdoll-frame-slot-${name}`);
const slot = (page: Page, name: string) => doll(page).getByTestId(`paperdoll-list-slot-${name}`);

test.describe('the sim card at 1280', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  test('draws the frame and opens the slot’s client tooltip in its centre on hover and focus', async ({
    page,
  }) => {
    const errors: string[] = [];
    page.on('pageerror', (error) => errors.push(error.message));
    await loadSim(page);
    await expect(page.getByTestId('page-title')).toHaveText('Simfury');
    // The frame, as on BiS and the character page; the list form waits for a phone. Its title bar
    // is the spec and the level, since the page's head already names the character.
    await expect(doll(page).locator('.pdf-frame')).toBeVisible();
    await expect(doll(page).locator('.pdf-list')).toBeHidden();
    await expect(doll(page).locator('.pdf-title-text')).toHaveText('Fury Warrior');
    await expect(doll(page).locator('.pdf-meta')).toHaveText('Level 33');
    await expect(socket(page, 'neck')).toHaveCount(0);
    await expect(socket(page, 'off_hand')).toHaveAccessibleName(/Arcanite Reaper, two-handed$/);
    const head = socket(page, 'head');
    await expect(head).toHaveAccessibleName('Head: Lionheart Helm');
    // Beside the frame, the worn list (character/WornList): each piece over its slot and item level,
    // the off hand a two-hander covers saying so.
    const rows = page.getByTestId('sim-gear-card').getByTestId('sim-gear-rows');
    await expect(rows.getByTestId('sim-gear-row-head')).toContainText('Head · Item level 61');
    await expect(rows.getByTestId('sim-gear-row-off_hand')).toContainText('Two-handed');
    // The frame opens on the first item: its client tooltip in the centre, with the binding and
    // armour lines, and the floating host leaves the frame alone (one tooltip on the page).
    const centre = doll(page).locator('[data-region="paperdoll-centre"]');
    const tip = centre.getByTestId('item-tooltip');
    await expect(tip).toContainText('Lionheart Helm');
    await expect(tip.getByTestId('item-tooltip-bind')).toHaveText('Binds when equipped');
    await expect(tip).toContainText('Plate');
    await socket(page, 'main_hand').hover();
    await expect(tip).toContainText('Arcanite Reaper');
    await page.mouse.move(0, 0);
    await head.focus();
    await expect(tip).toContainText('Lionheart Helm');
    await expect(page.getByTestId('item-tooltip')).toHaveCount(1);
    expect(errors).toEqual([]);
  });
});

test.describe('the sim card on a 390 phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, hasTouch: true, isMobile: true });

  test('the list form: a tap opens the item, a second tap closes it', async ({ page }) => {
    await loadSim(page);
    // Under the frame's 423 px minimum the frame gives way to the list, and its image is never fetched.
    await expect(doll(page).locator('.pdf-frame')).toBeHidden();
    await expect(slot(page, 'off_hand')).toContainText('Two-handed');
    // The frame's list form says it all: the worn list beside the frame is hidden here.
    await expect(page.getByTestId('sim-gear-rows')).toBeHidden();
    // loadSim's Run is a mouse click, and the card can land under where that mouse rests: its
    // hover would open the panel before the first tap. Park the mouse off the card first.
    await page.mouse.move(0, 0);
    // The page's floating tooltip; the hidden frame's own centre tooltip is not it.
    const floating = page.getByTestId('item-tooltip').locator('visible=true');
    const head = slot(page, 'head');
    await head.tap();
    await expect(floating).toBeVisible();
    await expect(floating).toContainText('Lionheart Helm');
    await head.tap();
    await expect(floating).toHaveCount(0);
  });
});

test.describe('the report "At pull" at 1280', () => {
  test.use({ viewport: { width: 1280, height: 900 } });

  test('opens a combatant’s gear on the frame, each slot with its log item level', async ({ page }) => {
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
    // At pull's Gear link opens the combatant's own view, whose Gear at the pull is the paperdoll
    // (the report frame, 2026-11-07: GearList inside PlayerView).
    await page.getByTestId('at-pull-links').first().getByRole('link', { name: 'Gear ›' }).click();
    const frame = page.getByTestId('gear-list');
    await expect(frame.getByTestId('paperdoll-frame-slot-head')).toHaveAccessibleName('Head: Lionheart Helm');
    // Beside the frame, each worn piece over its slot, the log's own item level and the enchant
    // named from the build's table on its own green line; the off hand a two-hander covers says so; no gain.
    await expect(frame.getByTestId('gear-list-row-head')).toContainText('Head · Item level 66');
    await expect(frame.getByTestId('gear-list-row-head')).toContainText('Crusader');
    await expect(frame.getByTestId('gear-list-row-main_hand')).toContainText('Main hand · Item level 63');
    await expect(frame.getByTestId('gear-list-row-off_hand')).toContainText('Two-handed');
    await expect(frame.locator('[data-testid^="paperdoll-list-gain"]')).toHaveCount(0);
    // The frame opens on the head: its client tooltip in the centre, the enchant on its green line.
    const centre = frame.locator('[data-region="paperdoll-centre"]');
    await expect(centre).toContainText('Lionheart Helm');
    await expect(centre).toContainText('Crusader');
    await frame.getByTestId('paperdoll-frame-slot-main_hand').hover();
    await expect(centre).toContainText('Arcanite Reaper');
  });
});
