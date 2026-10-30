// design/mocks/2026-10-29-site-alignment/capture-stage2-home.spec.ts
// Stage 2 captures for the home / content lane. Writes design/reviews/2026-10-29-site-alignment/
// stage2-<route>-<width>.png. Run from web/ naming this file:
//   CAPTURE_PORT=4795 npx playwright test -c ../design/mocks/2026-10-29-site-alignment/playwright.capture.config.ts capture-stage2-home
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { isolate, signIn } from './stubs';

const OUT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../reviews/2026-10-29-site-alignment');
const SETTLE_MS = 1500;
const MAIN = [1440, 390];

interface Shot {
  name: string;
  url: string;
  widths: number[];
  session: 'in' | 'out';
}

const SHOTS: Shot[] = [
  { name: 'home-signed-in', url: '/', widths: [...MAIN, 2000], session: 'in' },
  { name: 'home-signed-out', url: '/', widths: MAIN, session: 'out' },
  { name: 'about', url: '/about', widths: MAIN, session: 'out' },
  { name: 'terms', url: '/terms', widths: [1440], session: 'out' },
  { name: 'not-found', url: '/no-such-page', widths: MAIN, session: 'out' },
  { name: 'premium', url: '/premium', widths: MAIN, session: 'in' },
  { name: 'setup', url: '/setup', widths: MAIN, session: 'in' },
  { name: 'rankings', url: '/rankings', widths: MAIN, session: 'in' },
  { name: 'logs', url: '/logs', widths: MAIN, session: 'in' },
  { name: 'report', url: '/reports/fixture2abcd?fight=3', widths: [...MAIN, 2000], session: 'in' },
  { name: 'reports', url: '/reports', widths: [1440], session: 'in' },
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
