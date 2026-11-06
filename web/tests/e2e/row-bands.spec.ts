// web/tests/e2e/row-bands.spec.ts
// Row bands (design/specs/2026-11-01-row-bands.md): every lit row (the viewer's own row) is a
// rounded band whose content sits at least 12 px inside both sides, at 1440 and 390; an unlit row keeps square corners so its hairline runs straight; a link row's
// hover band is a raised ground with a ring. (The roster's handoff slot: guild-roster-handoff.spec.ts.)
// The raid report is built from the ui/ primitives, whose own rows are Row's `selected` ground, not
// row bands (web layout spec 2026-11-05 §3A.2), so it is no longer measured here; nor is the
// character page, whose fight rows are Rows since design/specs/2026-11-06-character.md.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { CHICAGO, NOW, stubGuildRaids, threePlanned } from './support/guild-raids';
import { gotoHydrated } from './support/hydrated';
import { GUILD_PAGE, litBandRooms, stubGuild } from './support/row-band-pages';

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

    test('the readiness own row is a rounded band holding its content', async ({ page }) => {
      await stubGuild(page, 'member');
      await gotoHydrated(page, `${GUILD_PAGE}#readiness`, 'guild-readiness-rows');
      await expectRoomyBands(page, 1);
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
