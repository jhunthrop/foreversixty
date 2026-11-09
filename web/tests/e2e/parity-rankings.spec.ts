// web/tests/e2e/parity-rankings.spec.ts
// /rankings and /rankings/<slug> against their mocks (design/parity.md; the look is the directions
// mock b1-rankings*): every state at 1440 x 1000, each region's blocks and box compared
// (support/parity.ts). The mocks are rendered before the tests (support/parity-mocks.ts:
// rankings_facts.mjs and ws_rankings.py) from the stubs this spec serves
// (support/rankings-parity.ts). A side-by-side per state lands in
// test-results/parity/rankings-<state>.png. Desktop only.
import { test, expect } from './support/test';
import {
  captureOutlined,
  captureRegions,
  compareRegions,
  formatFailures,
  openMock,
  prepareParityPage,
  readParityConfig,
  settleParityPage,
  writeSideBySide,
} from './support/parity';
import { renderRankingsMocks } from './support/parity-mocks';
import { openRankingsState, type RankingsParityState } from './support/rankings-parity';

const CONFIG = readParityConfig('rankings-parity.json');

test.describe('Rankings mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderRankingsMocks();
  });

  for (const state of CONFIG.states) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, mocks, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openRankingsState(page, state.id as RankingsParityState);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `rankings-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
