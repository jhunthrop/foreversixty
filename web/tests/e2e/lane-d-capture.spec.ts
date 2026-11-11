// web/tests/e2e/lane-d-capture.spec.ts
// Not a test: the tool behind lane D's review (design/reviews/2026-11-11-lane-d): one crest, the
// class-coloured name, the four-choice sign-up and the merged class picker, on the pages they show.
// Skipped unless LANE_D_CAPTURE names the capture set:
//
//   LANE_D_CAPTURE=after npx playwright test tests/e2e/lane-d-capture.spec.ts --project=desktop --workers=1
//
// Writes design/reviews/2026-11-11-lane-d/<set>/<page>-<width>.jpg, full page at 1440 and 390.
import { mkdirSync } from 'node:fs';
import type { Page } from '@playwright/test';
import path from 'node:path';
import { test } from './support/test';
import { FRAME_PAGES, openFramePage, type FramePage } from './support/frame-pages';
import { openGuildState } from './support/guild-parity';
import { openHome } from './support/home-page';
import { REPO_ROOT } from './support/parity';

const SET = process.env.LANE_D_CAPTURE ?? '';
const ONLY = process.env.LANE_D_PAGES?.split(',');
const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-11-lane-d', SET);
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const QUALITY = 72;
const FRAME_IDS = ['rankings', 'tiers', 'guides', 'character', '404'];
const SIXTY = { me: 'me-60.json', feed: 'feed-60.json', home: 'home-60.json', guild: 'guild-home-60.json' };

const PAGES: readonly FramePage[] = [
  { id: 'home-raid', open: (page: Page) => openHome(page, SIXTY) },
  { id: 'home-out', open: (page: Page) => openHome(page, { feed: 'feed-out.json' }) },
  { id: 'guild-raids', open: (page: Page) => openGuildState(page, 'dense-raids') },
  { id: 'guild-raid', open: (page: Page) => openGuildState(page, 'dense-raid-member') },
  { id: 'guild-roster', open: (page: Page) => openGuildState(page, 'dense-roster') },
  ...FRAME_PAGES.filter((framePage) => FRAME_IDS.includes(framePage.id)),
].filter((framePage) => ONLY === undefined || ONLY.includes(framePage.id));

test.describe('lane D capture', () => {
  test.skip(SET === '', 'a tool, run by hand: LANE_D_CAPTURE=<set>');
  test.skip(({ isMobile }) => isMobile, 'run on the desktop project; it sets its own widths');

  test('every changed page at 1440 and 390', async ({ browser, baseURL }) => {
    test.setTimeout(30 * 60_000);
    mkdirSync(OUT, { recursive: true });
    const failed: string[] = [];
    for (const framePage of PAGES) {
      for (const width of WIDTHS) {
        const context = await browser.newContext({ viewport: { width, height: HEIGHT }, baseURL });
        const page = await context.newPage();
        try {
          await openFramePage(page, framePage);
          await page.screenshot({
            path: path.join(OUT, `${framePage.id}-${width}.jpg`),
            type: 'jpeg',
            quality: QUALITY,
            fullPage: true,
          });
        } catch (error: unknown) {
          failed.push(`${framePage.id}-${width}: ${error instanceof Error ? error.message : String(error)}`);
        }
        await context.close();
      }
    }
    if (failed.length > 0) throw new Error(`pages that did not settle:\n${failed.join('\n')}`);
  });
});
