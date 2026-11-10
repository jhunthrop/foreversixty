// web/tests/e2e/guild-captures.spec.ts
// Not a test: the guild lane's review captures. Skipped unless GUILD_CAPTURE=1:
//
//   GUILD_CAPTURE=1 GUILD_CAPTURE_DIR=design/reviews/2026-11-09-f-guild/r1 \
//     E2E_PORT=4420 npx playwright test tests/e2e/guild-captures.spec.ts --project=desktop --workers=1
//
// Every guild parity state (design/mocks/web/guild-parity.json: each role, tab and flow) opened as the
// parity spec opens it (the parity fixtures, their clock and zone), full page at each width, into
// <GUILD_CAPTURE_DIR>/build-<state>-<width>.png. GUILD_CAPTURE_WIDTHS (comma list, default 390,1440)
// and GUILD_CAPTURE_STATES (comma list, default every state) narrow the run.
import { mkdirSync } from 'node:fs';
import path from 'node:path';
import { test } from './support/test';
import { openGuildState } from './support/guild-parity';
import { GUILD_PARITY_NOW, GUILD_PARITY_ZONE } from './support/guild-parity/clock';
import { REPO_ROOT, prepareParityPage, readParityConfig, settleParityPage } from './support/parity';

const OUT = path.join(REPO_ROOT, process.env.GUILD_CAPTURE_DIR ?? 'design/reviews/2026-11-09-f-guild/captures');
const WIDTHS = (process.env.GUILD_CAPTURE_WIDTHS ?? '390,1440').split(',').map(Number);
const ONLY = process.env.GUILD_CAPTURE_STATES?.split(',');
const STATES = readParityConfig('guild-parity.json').states.filter(
  (state) => ONLY === undefined || ONLY.includes(state.id),
);

test.describe('Guild review captures', () => {
  test.skip(() => process.env.GUILD_CAPTURE !== '1', 'a capture tool, run on demand');
  test.skip(({ isMobile }) => isMobile, 'the widths are set per capture');
  test.use({ timezoneId: GUILD_PARITY_ZONE });

  for (const { id } of STATES) {
    test(id, async ({ page }) => {
      test.slow();
      mkdirSync(OUT, { recursive: true });
      for (const width of WIDTHS) {
        await page.setViewportSize({ width, height: 1000 });
        await page.clock.setFixedTime(GUILD_PARITY_NOW);
        await prepareParityPage(page);
        await openGuildState(page, id);
        await settleParityPage(page);
        await page.screenshot({ path: path.join(OUT, `build-${id}-${width}.png`), fullPage: true });
      }
    });
  }
});
