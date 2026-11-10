// web/tests/e2e/frame-capture.spec.ts
// Not a test: the tool behind the before and after captures of a site-wide review (first the
// page-frame review, design/reviews/2026-11-10-header-width). Skipped unless FRAME_CAPTURE names
// the capture set:
//
//   FRAME_CAPTURE=after npx playwright test tests/e2e/frame-capture.spec.ts --project=desktop --workers=1
//
// For every capture page (support/frame-pages.ts CAPTURE_PAGES) at each width it writes, to
// <FRAME_CAPTURE_DIR>/<set>/: the first screen as <page>-<width>.jpg, the whole page as
// <page>-<width>-full.jpg at each width FRAME_CAPTURE_FULL names, and frame.json, the header bar's
// edges and every page column's content box per page and width. Optional settings:
//   FRAME_CAPTURE_DIR    the review folder, repo-relative (default the header-width review)
//   FRAME_CAPTURE_WIDTHS comma-separated widths (default 1440,1920,2000); below 600 a phone context
//   FRAME_CAPTURE_FULL   comma-separated widths that also get a full-page capture
//   FRAME_CAPTURE_PAGES  comma-separated page ids (default every capture page)
import { mkdirSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { test } from './support/test';
import { CAPTURE_PAGES, measureFrame, openFramePage, type FrameMeasure } from './support/frame-pages';
import { REPO_ROOT } from './support/parity';

const SET = process.env.FRAME_CAPTURE ?? '';
const OUT = path.join(
  REPO_ROOT,
  process.env.FRAME_CAPTURE_DIR ?? 'design/reviews/2026-11-10-header-width',
  SET,
);
const HEIGHT = 1000;
const PHONE_HEIGHT = 844;
/** Below this width the capture is a phone: touch, a mobile viewport and 2x density. */
const PHONE_MAX = 600;
const QUALITY = 70;

/** A comma-separated list from the environment, or the fallback when the variable is unset. */
const listFrom = (name: string, fallback: readonly string[]): readonly string[] =>
  (process.env[name] ?? fallback.join(',')).split(',').filter((item) => item !== '');

const WIDTHS = listFrom('FRAME_CAPTURE_WIDTHS', ['1440', '1920', '2000']).map(Number);
const FULL_WIDTHS = new Set(listFrom('FRAME_CAPTURE_FULL', []).map(Number));
const PAGE_IDS = new Set(
  listFrom(
    'FRAME_CAPTURE_PAGES',
    CAPTURE_PAGES.map((p) => p.id),
  ),
);
const PAGES = CAPTURE_PAGES.filter((p) => PAGE_IDS.has(p.id));

/** Scrolls the page through once so its lazy images load before a full-page capture. */
async function loadLazyImages(page: Page): Promise<void> {
  await page.evaluate(async () => {
    const step = window.innerHeight;
    for (let y = 0; y < document.body.scrollHeight; y += step) {
      window.scrollTo(0, y);
      await new Promise((resolve) => setTimeout(resolve, 60));
    }
    window.scrollTo(0, 0);
  });
  await page.waitForLoadState('networkidle');
}

const contextFor = (width: number) =>
  width < PHONE_MAX
    ? { viewport: { width, height: PHONE_HEIGHT }, isMobile: true, hasTouch: true, deviceScaleFactor: 2 }
    : { viewport: { width, height: HEIGHT } };

test.describe('page frame capture', () => {
  test.skip(SET === '', 'a tool, run by hand: FRAME_CAPTURE=<set>');
  test.skip(({ isMobile }) => isMobile, 'desktop widths only');

  test('every capture page at every width', async ({ browser, baseURL }) => {
    test.setTimeout(30 * 60_000);
    mkdirSync(OUT, { recursive: true });
    const frames: Record<string, Record<number, FrameMeasure>> = {};
    const failed: string[] = [];
    for (const framePage of PAGES) {
      for (const width of WIDTHS) {
        const context = await browser.newContext({ ...contextFor(width), baseURL });
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
        if (FULL_WIDTHS.has(width)) {
          await loadLazyImages(page);
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
