// design/mocks/2026-10-29-site-alignment/capture-stage2-planner-sim.spec.ts
// Stage 2 captures (planner and simulator lane) of the
// site-alignment spec. Same stubs and settle as capture.spec.ts; writes
// design/reviews/2026-10-29-site-alignment/stage2-*.png. /account must match before-account-*.
// Run as capture.spec.ts says, naming this file:
//   npx playwright test -c ../design/mocks/2026-10-29-site-alignment/playwright.capture.config.ts capture-stage2-planner-sim
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { isolate, signIn } from './stubs';

const OUT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../reviews/2026-10-29-site-alignment');
const SETTLE_MS = 1500;

interface Shot {
  name: string;
  url: string;
  widths: number[];
  session: 'in';
}

const MAIN = [1440, 390];
const SHOTS: Shot[] = [
  { name: 'planner', url: '/planner', widths: [...MAIN, 2000], session: 'in' },
  { name: 'sim', url: '/sim', widths: MAIN, session: 'in' },
  { name: 'sim-weights', url: '/sim/weights', widths: MAIN, session: 'in' },
  { name: 'sim-gear', url: '/sim/gear', widths: MAIN, session: 'in' },
];

for (const shot of SHOTS) {
  for (const width of shot.widths) {
    test(`${shot.name} ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await isolate(page);
      if (shot.session === 'in') await signIn(page);
      await page.goto(shot.url);
      await page.waitForLoadState('networkidle').catch(() => undefined);
      await page.waitForTimeout(SETTLE_MS);
      await page.screenshot({ path: path.join(OUT, `stage2-${shot.name}-${width}.png`), fullPage: true });
    });
  }
}
