// web/tests/e2e/batch-l-capture.spec.ts
// Not a test: the batch L integrator's merged-state captures (design/reviews/2026-11-11-batch-l): the
// pages where lanes D, E and G2 overlapped, home, rankings, planner and the guild raid pages, full page
// at 1440 and 390. Skipped unless BATCH_L_CAPTURE=1:
//
//   BATCH_L_CAPTURE=1 npx playwright test tests/e2e/batch-l-capture.spec.ts --project=desktop --workers=1
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { test } from './support/test';
import { FRAME_PAGES, openFramePage, type FramePage } from './support/frame-pages';
import { openGuildState } from './support/guild-parity';
import { openHome } from './support/home-page';
import { REPO_ROOT } from './support/parity';
import { routeRealIcons } from './support/real-icons';

const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-11-batch-l/captures');
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const QUALITY = 72;
const FRAME_IDS = ['rankings', 'planner'];
const SIXTY = { me: 'me-60.json', feed: 'feed-60.json', home: 'home-60.json', guild: 'guild-home-60.json' };

const PAGES: readonly FramePage[] = [
  { id: 'home-raid', open: (page: Page) => openHome(page, SIXTY) },
  { id: 'home-out', open: (page: Page) => openHome(page, { feed: 'feed-out.json' }) },
  { id: 'guild-raid', open: (page: Page) => openGuildState(page, 'dense-raid-member') },
  { id: 'guild-raid-officer', open: (page: Page) => openGuildState(page, 'dense-planner') },
  ...FRAME_PAGES.filter((framePage) => FRAME_IDS.includes(framePage.id)),
];

test.describe('batch L capture', () => {
  test.skip(process.env.BATCH_L_CAPTURE !== '1', 'a tool, run by hand: BATCH_L_CAPTURE=1');
  test.skip(({ isMobile }) => isMobile, 'run on the desktop project; it sets its own widths');

  test('the overlapped pages at 1440 and 390', async ({ browser, baseURL }) => {
    test.setTimeout(20 * 60_000);
    mkdirSync(OUT, { recursive: true });
    const failed: string[] = [];
    for (const framePage of PAGES) {
      for (const width of WIDTHS) {
        const context = await browser.newContext({ viewport: { width, height: HEIGHT }, baseURL });
        const page = await context.newPage();
        try {
          await routeRealIcons(page);
          await openFramePage(page, framePage);
          await page.screenshot({
            path: path.join(OUT, `${framePage.id}-${width}.jpg`),
            type: 'jpeg',
            quality: QUALITY,
            fullPage: true,
          });
        } catch (error: unknown) {
          failed.push(`${framePage.id}-${width}: ${error instanceof Error ? error.message : String(error)}`);
        } finally {
          await context.close();
        }
      }
    }
    if (failed.length > 0) throw new Error(`captures failed:\n${failed.join('\n')}`);
  });
});
