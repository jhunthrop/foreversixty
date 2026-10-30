// design/mocks/2026-10-29-site-alignment/capture-stage2-guild.spec.ts
// Stage 2 captures for the guild lane (design/specs/2026-10-29-site-alignment.md §5.9, §5.10): the
// landing, the guild page and each tab, the raid page, the standalone settings and the claim page.
// Writes design/reviews/2026-10-29-site-alignment/stage2-guild-<name>-<width>.png. Run as
// capture.spec.ts says, naming this file:
//   CAPTURE_PORT=4631 npx playwright test -c ../design/mocks/2026-10-29-site-alignment/playwright.capture.config.ts capture-stage2-guild
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { GUILD_ID } from '../../../web/src/fixtures/guild/mock-guild';
import { CLAIM_HOLDER_CAN, guildBilling } from '../../../web/src/test-support/billing-fixtures';
import { stubGuildAttendance } from '../../../web/tests/e2e/support/guild-attendance';
import { SETTINGS_GUILD_PATH, guildChannel, settingsBody, stubGuildDiscord } from '../../../web/tests/e2e/support/guild-discord';
import { GUILD_URL, ONYXIA_ID, raidUrl, raidsMe, stubGuildRaids } from '../../../web/tests/e2e/support/guild-raids';
import { envelope } from '../../../web/tests/e2e/support/officer-tools';
import { isolate, signIn } from './stubs';

const OUT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../reviews/2026-10-29-site-alignment');
const SETTLE_MS = 1500;

type Session = 'in' | 'guild' | 'settings' | 'tab';

interface Shot {
  name: string;
  url: string;
  widths: number[];
  session: Session;
}

const MAIN = [1440, 390];
const SHOTS: Shot[] = [
  { name: 'landing', url: '/guild', widths: MAIN, session: 'in' },
  { name: 'overview', url: GUILD_URL, widths: [...MAIN, 2000], session: 'guild' },
  { name: 'roster', url: `${GUILD_URL}#roster`, widths: MAIN, session: 'guild' },
  { name: 'raids', url: `${GUILD_URL}#raids`, widths: MAIN, session: 'guild' },
  { name: 'raid', url: raidUrl(ONYXIA_ID), widths: MAIN, session: 'guild' },
  { name: 'loot', url: `${GUILD_URL}#loot`, widths: MAIN, session: 'guild' },
  { name: 'settings', url: SETTINGS_GUILD_PATH, widths: [...MAIN, 2000], session: 'settings' },
  { name: 'settings-tab', url: `${GUILD_URL}#settings`, widths: MAIN, session: 'tab' },
];

async function stubSettingsTab(page: Page): Promise<void> {
  await stubGuildRaids(page, { viewer: 'officer', officerTools: true });
  const me = raidsMe('officer');
  const billingMe = { ...me, guilds: me.guilds.map((g) => ({ ...g, can: [...CLAIM_HOLDER_CAN], plan: guildBilling() })) };
  await page.route('**/v1/me', (route) => route.fulfill(envelope(billingMe)));
  await page.route(`**/v1/guilds/${GUILD_ID}/settings`, (route) =>
    route.fulfill(envelope(settingsBody(guildChannel(), false))),
  );
}

for (const shot of SHOTS) {
  for (const width of shot.widths) {
    test(`${shot.name} ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await isolate(page);
      if (shot.session === 'in') await signIn(page);
      if (shot.session === 'guild') await stubGuildAttendance(page, { viewer: 'officer', officerTools: true });
      if (shot.session === 'settings') await stubGuildDiscord(page, { discord: guildChannel(), officerTools: true });
      if (shot.session === 'tab') await stubSettingsTab(page);
      await page.goto(shot.url);
      await page.waitForLoadState('networkidle').catch(() => undefined);
      await page.waitForTimeout(SETTLE_MS);
      await page.screenshot({ path: path.join(OUT, `stage2-guild-${shot.name}-${width}.png`), fullPage: true });
    });
  }
}
