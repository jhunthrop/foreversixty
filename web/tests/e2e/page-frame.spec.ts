// web/tests/e2e/page-frame.spec.ts
// One page frame on every page (owner 2026-11-10: "why don't we have consistent width for every
// experience?"): at a given viewport the header bar, the page head band, the page body and the
// footer share one left and one right edge, whatever the page holds. The frame is .page-column
// (global.css, from tokens.css --page-max and --page-gutter); prose keeps its reading measure
// inside the body, which this does not measure.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { FRAME_PAGES, measureFrame, openFramePage } from './support/frame-pages';
import { openGuildState } from './support/guild-parity';
import { openTiersState } from './support/tiers-parity';

const WIDTH = 1920;
const HEIGHT = 1000;
/** At 1920 the wide frame is clamp(1600, 1920 - 160, 1792) = 1760 px with the 64 px gutter. */
const FRAME = { left: (WIDTH - 1760) / 2 + 64, right: WIDTH - (WIDTH - 1760) / 2 - 64 };

test.describe('one page frame', () => {
  test.skip(({ isMobile }) => isMobile, 'a desktop width');
  test.use({ viewport: { width: WIDTH, height: HEIGHT } });

  for (const framePage of FRAME_PAGES) {
    test(`${framePage.id}: the header, the bands, the body and the footer share the frame's edges`, async ({
      page,
    }) => {
      test.slow();
      await openFramePage(page, framePage);
      const frame = await measureFrame(page);
      expect(frame.header).toEqual(FRAME);
      expect(frame.columns.map((column) => column.where)).toEqual(expect.arrayContaining(['body', 'footer']));
      for (const column of frame.columns) {
        expect({ where: column.where, left: column.left, right: column.right }).toEqual({
          where: column.where,
          ...FRAME,
        });
      }
    });
  }

  /** The bar's width, and whether its door icons, "/" hint and short spec line show. */
  async function bar(page: Page) {
    const header = page.locator('.site-header');
    return {
      width: (await header.boundingBox())?.width,
      icons: await page.locator('.site-door-icon').first().isVisible(),
      hint: await page.locator('.site-search-kbd').isVisible(),
      spec: await page.locator('.csel-btn .csel-sub-short').isVisible(),
    };
  }

  // The owner's report (2026-11-10): loaded directly, a guild page's bar was narrower than the tier
  // list's, without the door icons, the "/" hint or the spec beside the name. Both now carry the
  // one bar.
  const BAR_PAGES: readonly { id: string; open: (page: Page) => Promise<void> }[] = [
    { id: 'a guild page', open: (page) => openGuildState(page, 'dense-officer') },
    { id: 'the tier list', open: (page) => openTiersState(page, 'dps-you') },
  ];
  for (const { id, open } of BAR_PAGES) {
    test(`${id} carries the full bar: the frame's width, door icons, "/" hint and spec`, async ({ page }) => {
      await open(page);
      await expect(page.locator('.csel-btn .csel-sub-short')).toBeVisible();
      expect(await bar(page)).toEqual({
        width: FRAME.right - FRAME.left,
        icons: true,
        hint: true,
        spec: true,
      });
    });
  }
});
