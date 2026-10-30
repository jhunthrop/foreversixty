// design/mocks/2026-10-29-site-alignment/capture-stage1.spec.ts
// Stage 1 captures (the shared pieces: PageHead, Card, TabStrip, CharacterHeader surface) of the
// site-alignment spec. Same stubs and settle as capture.spec.ts; writes
// design/reviews/2026-10-29-site-alignment/stage1-*.png. /account must match before-account-*.
// Run as capture.spec.ts says, naming this file:
//   npx playwright test -c ../design/mocks/2026-10-29-site-alignment/playwright.capture.config.ts capture-stage1
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { isolate, signIn } from './stubs';
import { stubGuildAttendance } from '../../../web/tests/e2e/support/guild-attendance';
import { GUILD_URL } from '../../../web/tests/e2e/support/guild-raids';

const OUT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../reviews/2026-10-29-site-alignment');
const SETTLE_MS = 1500;

interface Shot {
  name: string;
  url: string;
  widths: number[];
  session: 'in' | 'guild';
}

const MAIN = [1440, 390];
const SHOTS: Shot[] = [
  { name: 'account', url: '/account', widths: [...MAIN, 2000], session: 'in' },
  { name: 'bis-warrior-fury', url: '/bis/warrior/fury', widths: MAIN, session: 'in' },
  { name: 'guide-warrior-fury', url: '/guides/warrior/fury', widths: MAIN, session: 'in' },
  { name: 'planner', url: '/planner', widths: MAIN, session: 'in' },
  { name: 'guild-overview', url: GUILD_URL, widths: MAIN, session: 'guild' },
  { name: 'sim', url: '/sim', widths: MAIN, session: 'in' },
  { name: 'premium', url: '/premium', widths: MAIN, session: 'in' },
  { name: 'setup', url: '/setup', widths: MAIN, session: 'in' },
];

for (const shot of SHOTS) {
  for (const width of shot.widths) {
    test(`${shot.name} ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await isolate(page);
      if (shot.session === 'in') await signIn(page);
      if (shot.session === 'guild') await stubGuildAttendance(page, { viewer: 'officer', officerTools: true });
      await page.goto(shot.url);
      await page.waitForLoadState('networkidle').catch(() => undefined);
      await page.waitForTimeout(SETTLE_MS);
      await page.screenshot({ path: path.join(OUT, `stage1-${shot.name}-${width}.png`), fullPage: true });
    });
  }
}
