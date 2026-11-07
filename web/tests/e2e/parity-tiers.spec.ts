// web/tests/e2e/parity-tiers.spec.ts
// /tiers, /tiers/healer and /tiers/tank against their mocks (design/parity.md,
// design/specs/2026-11-06-tiers.md): every state at 1440 x 1000, each region's blocks and box
// compared (support/parity.ts). The mocks are rendered before the tests (support/parity-mocks.ts:
// ws_tiers.py) from the published BiS files the build prerenders the lists from, so a re-rank
// changes both sides alike. A side-by-side per state lands in test-results/parity/tiers-<state>.png.
// Desktop only.
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
import { renderTiersMocks } from './support/parity-mocks';
import { openTiersState, type TiersParityState } from './support/tiers-parity';

const CONFIG = readParityConfig('tiers-parity.json');

test.describe('Tier list mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderTiersMocks();
  });

  for (const state of CONFIG.states) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, mocks, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openTiersState(page, state.id as TiersParityState);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `tiers-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
