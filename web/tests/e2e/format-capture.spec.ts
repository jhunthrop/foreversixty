// web/tests/e2e/format-capture.spec.ts
// Not a test: the tool behind the house-format review (design/reviews/2026-11-10-unify-format).
// Skipped unless FORMAT_CAPTURE names the capture set:
//
//   FORMAT_CAPTURE=after npx playwright test tests/e2e/format-capture.spec.ts --project=desktop --workers=1
//
// Every page whose dates, clocks, numbers, durations, parses or countdowns changed, full page at
// 1440 and 390, to design/reviews/2026-11-10-unify-format/<set>/<page>-<width>.jpg.
import { mkdirSync } from 'node:fs';
import type { Page } from '@playwright/test';
import path from 'node:path';
import { test } from './support/test';
import { FRAME_PAGES, openFramePage, type FramePage } from './support/frame-pages';
import { openGuildState } from './support/guild-parity';
import { REPO_ROOT } from './support/parity';

const SET = process.env.FORMAT_CAPTURE ?? '';
const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-10-unify-format', SET);
const WIDTHS = [1440, 390] as const;
const HEIGHT = 1000;
const QUALITY = 70;
const FRAME_IDS = [
  'home',
  'tiers',
  'guides',
  'guides-class',
  'guides-spec',
  'bis-landing',
  'bis',
  'changelog',
  'rankings',
  'character',
  'logs',
  'reports',
  'report',
  'sim',
  'sim-result',
  'guild',
  'guild-roster',
  'guild-loot',
];
const GUILD_TABS = ['dense-raids', 'dense-planner', 'dense-member', 'dense-applications', 'dense-readiness'];

const PAGES: readonly FramePage[] = [
  ...FRAME_PAGES.filter((framePage) => FRAME_IDS.includes(framePage.id)),
  ...GUILD_TABS.map((id) => ({
    id: `guild-${id.replace('dense-', '')}`,
    open: (page: Page) => openGuildState(page, id),
  })),
];

test.describe('house format capture', () => {
  test.skip(SET === '', 'a tool, run by hand: FORMAT_CAPTURE=<set>');
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
