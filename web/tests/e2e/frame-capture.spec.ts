// web/tests/e2e/frame-capture.spec.ts
// Not a test: the tool behind the page-frame review (design/reviews/2026-11-10-header-width).
// Skipped unless FRAME_CAPTURE names the capture set:
//
//   FRAME_CAPTURE=after npx playwright test tests/e2e/frame-capture.spec.ts --project=desktop --workers=1
//
// For every page type (support/frame-pages.ts) at 1440, 1920 and 2000 it writes, to
// design/reviews/2026-11-10-header-width/<set>/: the first screen as <page>-<width>.jpg,
// with FRAME_CAPTURE_FULL=1 the whole page at 1920 as <page>-1920-full.jpg, and frame.json, the
// header bar's edges and every page column's content box per page and width.
import { mkdirSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { test } from './support/test';
import { FRAME_PAGES, measureFrame, openFramePage, type FrameMeasure } from './support/frame-pages';
import { REPO_ROOT } from './support/parity';

const SET = process.env.FRAME_CAPTURE ?? '';
const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-10-header-width', SET);
const WIDTHS = [1440, 1920, 2000] as const;
const HEIGHT = 1000;
const QUALITY = 70;

test.describe('page frame capture', () => {
  test.skip(SET === '', 'a tool, run by hand: FRAME_CAPTURE=<set>');
  test.skip(({ isMobile }) => isMobile, 'desktop widths only');

  test('every page type at 1440, 1920 and 2000', async ({ browser, baseURL }) => {
    test.setTimeout(30 * 60_000);
    mkdirSync(OUT, { recursive: true });
    const frames: Record<string, Record<number, FrameMeasure>> = {};
    const failed: string[] = [];
    for (const framePage of FRAME_PAGES) {
      for (const width of WIDTHS) {
        const context = await browser.newContext({ viewport: { width, height: HEIGHT }, baseURL });
        const page = await context.newPage();
        try {
          await openFramePage(page, framePage);
        } catch (error: unknown) {
          // One page that will not settle must not cost the rest of the set; it is listed below.
          failed.push(`${framePage.id}-${width}: ${error instanceof Error ? error.message : String(error)}`);
          await context.close();
          continue;
        }
        await page.evaluate(() => window.scrollTo(0, 0));
        frames[framePage.id] = { ...frames[framePage.id], [width]: await measureFrame(page) };
        const file = `${framePage.id}-${width}`;
        await page.screenshot({ path: path.join(OUT, `${file}.jpg`), type: 'jpeg', quality: QUALITY });
        if (width === 1920 && process.env.FRAME_CAPTURE_FULL === '1') {
          await page.screenshot({
            path: path.join(OUT, `${file}-full.jpg`),
            type: 'jpeg',
            quality: QUALITY,
            fullPage: true,
          });
        }
        await context.close();
      }
    }
    writeFileSync(path.join(OUT, 'frame.json'), `${JSON.stringify(frames, null, 1)}\n`);
    if (failed.length > 0) throw new Error(`pages that did not settle:\n${failed.join('\n')}`);
  });
});
