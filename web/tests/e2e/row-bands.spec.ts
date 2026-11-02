// web/tests/e2e/row-bands.spec.ts
// Row bands (design/specs/2026-11-01-row-bands.md): every lit row (the viewer's own row, an
// addressed fight) is a rounded band whose content sits at least 12 px inside both sides, at
// 1440 and 390; an unlit row keeps square corners so its hairline runs straight; a link row's
// hover band is a raised ground with a ring; and a roster row with no export keeps no blank slot.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { CHARACTER_PATH } from './support/character-fixture';
import { AFTER_END, DONE, stubGuildAttendance } from './support/guild-attendance';
import { CHICAGO, NOW, ONYXIA_ID, raidUrl, stubGuildRaids, threePlanned } from './support/guild-raids';
import { gotoHydrated } from './support/hydrated';
import { GUILD_PAGE, litBandRooms, stubCharacter, stubGuild } from './support/row-band-pages';

const ROOM_PX = 12;
const RADIUS = '8px';
const WIDTHS = [1440, 390] as const;

test.use({ timezoneId: CHICAGO });
test.skip(() => test.info().project.name !== 'desktop', 'each test sets its own widths');

async function expectRoomyBands(page: Page, count: number): Promise<void> {
  const bands = await litBandRooms(page);
  expect(bands, 'lit bands on the page').toHaveLength(count);
  for (const band of bands) {
    expect(band.left, `${band.classes} left`).toBeGreaterThanOrEqual(ROOM_PX - 0.5);
    expect(band.right, `${band.classes} right, at ${band.rightmost}`).toBeGreaterThanOrEqual(ROOM_PX - 0.5);
    expect(band.radius, band.classes).toBe(RADIUS);
  }
}

for (const width of WIDTHS) {
  test.describe(`at ${width}`, () => {
    test.use({ viewport: { width, height: 900 } });

    test('the roster own row is a rounded band holding its content', async ({ page }) => {
      await stubGuild(page, 'member');
      await gotoHydrated(page, `${GUILD_PAGE}#roster`, 'guild-roster-rows');
      await expectRoomyBands(page, 1);
    });

    test('a roster row with no export keeps no blank handoff slot', async ({ page }) => {
      await stubGuild(page, 'member');
      await gotoHydrated(page, `${GUILD_PAGE}#roster`, 'guild-roster-rows');
      const slots = page.getByTestId('guild-roster-handoff');
      await expect(slots.first()).toHaveAttribute('data-state', 'empty');
      await expect(slots.first()).toBeHidden();
      await expect(slots.first().locator('..')).toBeHidden();
    });

    test('the readiness own row is a rounded band holding its content', async ({ page }) => {
      await stubGuild(page, 'member');
      await gotoHydrated(page, `${GUILD_PAGE}#readiness`, 'guild-readiness-rows');
      await expectRoomyBands(page, 1);
    });

    test('the raid report own rows are rounded bands holding their content', async ({ page }) => {
      await stubGuildAttendance(page, { viewer: 'officer', nights: DONE });
      await page.clock.install({ time: AFTER_END });
      await gotoHydrated(page, raidUrl(ONYXIA_ID), 'raid-report');
      // Kraggor's own rows: deaths, loot, consumables, the attendance Present grid and, opened,
      // the sign-ups (an officer's row carries its state select).
      await expectRoomyBands(page, 4);
      await page.getByTestId('raid-roster-disclosure-summary').click();
      await expect(page.getByTestId('raid-roster')).toBeVisible();
      await expectRoomyBands(page, 5);
    });

    test('the addressed fight is a rounded band; the rows beside it keep square rules', async ({ page }) => {
      await stubCharacter(page);
      await page.goto(CHARACTER_PATH);
      const first = page.locator('li[id^="fight-"]').first();
      await expect(first).toBeVisible();
      const id = await first.getAttribute('id');
      expect(id).not.toBeNull();
      await page.evaluate((hash) => (window.location.hash = hash), id ?? '');
      await expectRoomyBands(page, 1);
      const unlit = page.locator('li[id^="fight-"]').nth(1);
      await expect(unlit).toHaveCSS('border-top-left-radius', '0px');
    });
  });
}

test('a planned raid row is square at rest and a ringed band on hover', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 900 });
  await stubGuildRaids(page, { viewer: 'member', raids: threePlanned() });
  await page.clock.install({ time: NOW });
  await gotoHydrated(page, `${GUILD_PAGE}#raids`, 'guild-raids-tab');
  const row = page.getByTestId('guild-planned-rows').locator('a').first();
  await expect(row).toHaveCSS('border-top-left-radius', '0px');
  await row.hover();
  await expect(row).toHaveCSS('border-top-left-radius', RADIUS);
  await expect(row).toHaveCSS('background-color', 'rgb(26, 33, 48)');
  await expect(row).toHaveCSS('box-shadow', /rgba\(229, 185, 85, 0\.22\) 0px 0px 0px 1px inset/);
});
