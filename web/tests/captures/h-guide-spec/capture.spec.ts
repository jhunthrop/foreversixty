// web/tests/captures/h-guide-spec/capture.spec.ts
// The 2026-11-10 h-guide-spec captures: four spec guides (a hunter and a healer among them) at
// 1440, 2000 and 390, and /sim/specs at 390 and 1440, opened through the parity gate's openers.
//   CAPTURE_OUT=<dir> CAPTURE_ONLY=<regex> CAPTURE_PORT=<port> \
//     npx playwright test -c tests/captures/h-guide-spec/playwright.capture.config.ts
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import { test, type Page } from '@playwright/test';
import { prepareParityPage, REPO_ROOT, settleParityPage } from '../../e2e/support/parity';
import { openToolState } from '../../e2e/support/sim-parity';
import { openGuidesState } from '../../e2e/support/guides-page';

const OUT =
  process.env.CAPTURE_OUT ?? path.join(REPO_ROOT, 'design/reviews/2026-11-10-h-guide-spec/captures');
const ONLY = process.env.CAPTURE_ONLY ? new RegExp(process.env.CAPTURE_ONLY) : null;
const HEIGHT = 1000;
const GUIDE_WIDTHS = [1440, 2000, 390] as const;
const GUIDES = ['hunter/survival', 'priest/holy', 'paladin/retribution', 'warlock/destruction'] as const;

type Opener = (page: Page) => Promise<void>;

const PAGES: readonly [name: string, widths: readonly number[], open: Opener][] = [
  ...GUIDES.map((guide): [string, readonly number[], Opener] => [
    `guide-${guide.replace('/', '-')}`,
    GUIDE_WIDTHS,
    (page) => openGuidesState(page, `/guides/${guide}`, 'out'),
  ]),
  ['sim-specs', [390, 1440], (page) => openToolState(page, 'specs')],
];

/** A full-page shot never scrolls, so a lazy image far below the fold stays an empty socket:
 *  walk the page once, as a reader's scroll would, and let the fetches finish. */
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

mkdirSync(OUT, { recursive: true });

for (const [name, widths, open] of PAGES) {
  if (ONLY && !ONLY.test(name)) continue;
  for (const width of widths) {
    test(`${name} at ${width}`, async ({ page }) => {
      test.slow();
      await page.setViewportSize({ width, height: HEIGHT });
      await prepareParityPage(page);
      await open(page);
      await loadLazyImages(page);
      await settleParityPage(page);
      await page.screenshot({ path: path.join(OUT, `${name}-${width}.png`), fullPage: true });
    });
  }
}
