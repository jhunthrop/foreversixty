// web/tests/e2e/parity-guild.spec.ts
// The guild pages against their committed mocks (design/parity.md, "The guild states";
// design/specs/2026-11-05-web-layout.md section 7): every state at 1440 x 1000, each region's
// blocks and box compared (support/parity.ts), a side-by-side per state in
// test-results/parity/guild-<state>.png. Both sides draw the same fixtures
// (tests/e2e/fixtures/guild-parity). Desktop project only: the mocks are 1440. The tabs batch H
// redrew (Raids, the planner, Progression, Applications, Recruit, Settings) compare against the
// dense mocks (guild-dense.json, `parity`), drawn from the 40-raider world.
import { test, expect } from './support/test';
import { openGuildState } from './support/guild-parity';
import {
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  MOCKS_DIR,
  openMock,
  prepareParityPage,
  readParityConfig,
  settleParityPage,
  writeSideBySide,
} from './support/parity';
import { GUILD_PARITY_NOW, GUILD_PARITY_ZONE } from './support/guild-parity/clock';

const CONFIG = readParityConfig('guild-parity.json');
/** The guild tabs batch H redrew (design/mocks/ws_guild_dense_tabs.py) take their states' places;
 *  guild-dense.json marks them `parity`. Its Overview, Roster, Readiness and Loot states are review
 *  captures only. */
const DENSE = readParityConfig('guild-dense.json').states.filter((state) => state.parity === true);
const STATES = [...CONFIG.states, ...DENSE];

test.describe('Guild mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  // The mocks print the guild's raid times in its own zone; a viewer there sees no "your time" line.
  test.use({ viewport: CONFIG.viewport, timezoneId: GUILD_PARITY_ZONE });

  for (const state of STATES) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, MOCKS_DIR, state.mock);
      const mock = await captureRegions(mockPage);

      await page.clock.setFixedTime(GUILD_PARITY_NOW);
      await prepareParityPage(page);
      await openGuildState(page, state.id);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `guild-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
