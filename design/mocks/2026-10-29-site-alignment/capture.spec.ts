// design/mocks/2026-10-29-site-alignment/capture.spec.ts
// Before captures of every user-facing route against the account pattern (spec
// design/specs/2026-10-29-site-alignment.md). Not a behaviour test: it stubs /v1/* (signed in as a
// Premium founder with two characters, Discord connected, or signed out), opens each route at 1440
// and 390, and writes full-page PNGs to design/reviews/2026-10-29-site-alignment/before-*.png.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test } from '@playwright/test';
import { isolate, signIn, signOut } from './stubs';
import { stubGuildAttendance } from '../../../web/tests/e2e/support/guild-attendance';
import { GUILD_URL } from '../../../web/tests/e2e/support/guild-raids';
import { SETTINGS_GUILD_PATH, guildChannel, stubGuildDiscord } from '../../../web/tests/e2e/support/guild-discord';

const OUT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../reviews/2026-10-29-site-alignment');
const SETTLE_MS = 1500;

interface Shot {
  name: string;
  url: string;
  widths: number[];
  session: 'in' | 'out' | 'guild' | 'settings';
}

const MAIN = [1440, 390];
const DESK = [1440];
const SHOTS: Shot[] = [
  { name: 'account', url: '/account', widths: MAIN, session: 'in' },
  { name: 'home-signed-in', url: '/', widths: MAIN, session: 'in' },
  { name: 'home-signed-out', url: '/', widths: DESK, session: 'out' },
  { name: 'bis', url: '/bis', widths: MAIN, session: 'in' },
  { name: 'bis-warrior-fury', url: '/bis/warrior/fury', widths: MAIN, session: 'in' },
  { name: 'guide-warrior-fury', url: '/guides/warrior/fury', widths: MAIN, session: 'in' },
  { name: 'guides', url: '/guides', widths: DESK, session: 'in' },
  { name: 'planner', url: '/planner', widths: MAIN, session: 'in' },
  { name: 'sim', url: '/sim', widths: MAIN, session: 'in' },
  { name: 'sim-weights', url: '/sim/weights', widths: DESK, session: 'in' },
  { name: 'sim-gear', url: '/sim/gear', widths: DESK, session: 'in' },
  { name: 'rankings', url: '/rankings', widths: MAIN, session: 'in' },
  { name: 'tiers', url: '/tiers', widths: MAIN, session: 'in' },
  { name: 'guild-landing', url: '/guild', widths: DESK, session: 'in' },
  { name: 'guild-overview', url: GUILD_URL, widths: MAIN, session: 'guild' },
  { name: 'guild-roster', url: `${GUILD_URL}#roster`, widths: MAIN, session: 'guild' },
  { name: 'guild-raids', url: `${GUILD_URL}#raids`, widths: DESK, session: 'guild' },
  { name: 'guild-loot', url: `${GUILD_URL}#loot`, widths: DESK, session: 'guild' },
  { name: 'guild-settings', url: SETTINGS_GUILD_PATH, widths: MAIN, session: 'settings' },
  { name: 'report', url: '/reports/fixture2abcd?fight=3', widths: MAIN, session: 'in' },
  { name: 'logs', url: '/logs', widths: MAIN, session: 'in' },
  { name: 'premium', url: '/premium', widths: MAIN, session: 'in' },
  { name: 'setup', url: '/setup', widths: MAIN, session: 'in' },
  { name: 'about', url: '/about', widths: DESK, session: 'in' },
  { name: 'terms', url: '/terms', widths: DESK, session: 'in' },
  { name: 'not-found', url: '/no-such-page', widths: DESK, session: 'in' },
  { name: 'character', url: '/character/us/hardcore/elyra-duskvale', widths: DESK, session: 'in' },
];

for (const shot of SHOTS) {
  for (const width of shot.widths) {
    test(`${shot.name} ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await isolate(page);
      if (shot.session === 'in') await signIn(page);
      if (shot.session === 'out') await signOut(page);
      if (shot.session === 'guild') await stubGuildAttendance(page, { viewer: 'officer', officerTools: true });
      if (shot.session === 'settings') await stubGuildDiscord(page, { discord: guildChannel(), officerTools: true });
      await page.goto(shot.url);
      await page.waitForLoadState('networkidle').catch(() => undefined);
      await page.waitForTimeout(SETTLE_MS);
      await page.screenshot({ path: path.join(OUT, `before-${shot.name}-${width}.png`), fullPage: true });
    });
  }
}
