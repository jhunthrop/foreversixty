// web/tests/e2e/guild-captures.spec.ts
// Not a test: the guild lane's review captures (redesign plan, Wave 2: "capture at 390, 1024, 1440 and
// 2000 beside the mock"). Skipped unless GUILD_CAPTURE=1:
//
//   GUILD_CAPTURE=1 E2E_PORT=4420 npx playwright test tests/e2e/guild-captures.spec.ts --project=desktop
//
// For every guild sub-view it opens the built page as the parity spec does (the parity fixtures, the
// fixtures' clock and zone) and its batch 3 directions mock (design/mocks/directions/html), each full
// page at every width, into design/reviews/2026-11-09-p10-guild/{build,mock}-<view>-<width>.png.
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import { test } from './support/test';
import { openGuildState } from './support/guild-parity';
import { GUILD_PARITY_NOW, GUILD_PARITY_ZONE } from './support/guild-parity/clock';
import { REPO_ROOT, prepareParityPage, settleParityPage } from './support/parity';

const OUT = path.join(REPO_ROOT, 'design/reviews/2026-11-09-p10-guild');
const MOCKS = path.join(REPO_ROOT, 'design/mocks/directions/html');
const WIDTHS = [390, 1024, 1440, 2000] as const;

/** Each sub-view: its parity state and its directions mock. */
const VIEWS: ReadonlyArray<{ view: string; state: string; mock: string }> = [
  { view: 'overview', state: 'officer', mock: 'b3-guild-overview' },
  { view: 'roster', state: 'roster', mock: 'b3-guild-roster' },
  { view: 'raid', state: 'raid', mock: 'b3-guild-raid' },
  { view: 'progression', state: 'progression', mock: 'b3-guild-progression' },
  { view: 'readiness', state: 'readiness', mock: 'b3-guild-readiness' },
  { view: 'loot', state: 'loot', mock: 'b3-guild-loot' },
  { view: 'applications', state: 'applications', mock: 'b3-guild-applications' },
  { view: 'recruit', state: 'recruit', mock: 'b3-guild-recruit' },
  { view: 'settings', state: 'settings', mock: 'b3-guild-settings' },
];

test.describe('Guild review captures', () => {
  test.skip(() => process.env.GUILD_CAPTURE !== '1', 'a capture tool, run on demand');
  test.skip(({ isMobile }) => isMobile, 'the widths are set per capture');
  test.use({ timezoneId: GUILD_PARITY_ZONE });

  for (const { view, state, mock } of VIEWS) {
    test(`${view}`, async ({ page, browser }) => {
      test.slow();
      mkdirSync(OUT, { recursive: true });
      for (const width of WIDTHS) {
        await page.setViewportSize({ width, height: 1000 });
        await page.clock.setFixedTime(GUILD_PARITY_NOW);
        await prepareParityPage(page);
        await openGuildState(page, state);
        await settleParityPage(page);
        await page.screenshot({ path: path.join(OUT, `build-${view}-${width}.png`), fullPage: true });

        const mockPage = await browser.newPage({ viewport: { width, height: 1000 } });
        await mockPage.goto(`file://${path.join(MOCKS, `${mock}.html`)}`);
        await mockPage.waitForLoadState('networkidle');
        await mockPage.screenshot({ path: path.join(OUT, `mock-${view}-${width}.png`), fullPage: true });
        await mockPage.close();
      }
    });
  }
});
