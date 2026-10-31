// design/mocks/2026-10-30-leftovers/capture-b.spec.ts
// Lane B captures: home (the Your upgrades reserve), /setup (step tags, character row) and the
// report (Not credited tag, trash row button). Writes design/reviews/2026-10-30-leftovers/b-<route>-<width>.png.
// Run from web/ with the fixture site on CAPTURE_PORT (default 4672):
//   npx playwright test -c ../design/mocks/2026-10-30-leftovers/playwright.capture.config.ts
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { isolate, signIn } from '../2026-10-29-site-alignment/stubs';

const OUT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../reviews/2026-10-30-leftovers');
const SETTLE_MS = 1500;
const WIDTHS = [1440, 390];
const SHOTS = [
  { name: 'home', url: '/' },
  { name: 'setup', url: '/setup' },
  { name: 'report', url: '/reports/fixture2abcd?fight=3' },
];

for (const shot of SHOTS) {
  for (const width of WIDTHS) {
    test(`${shot.name} ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await isolate(page);
      await signIn(page);
      await page.goto(shot.url);
      await page.waitForLoadState('networkidle').catch(() => undefined);
      await page.waitForTimeout(SETTLE_MS);
      await page.screenshot({ path: path.join(OUT, `b-${shot.name}-${width}.png`), fullPage: true });
    });
  }
}
