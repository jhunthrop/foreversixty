// web/tests/e2e/parity-home.spec.ts
// Home against its committed mocks (design/parity.md, design/specs/2026-11-06-home.md): every
// state at 1440 x 1000, each region's blocks and box compared (support/parity.ts), a side-by-side
// per state in test-results/parity/home-<state>.png. Desktop project only: the mocks are 1440.
import { test, expect } from './support/test';
import { openHome, type HomeState } from './support/home-page';
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

const CONFIG = readParityConfig('home-parity.json');

test.describe('Home mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  for (const state of CONFIG.states as ((typeof CONFIG.states)[number] & HomeState)[]) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      // Home's mocks are the committed ones until its generator joins the test-time pattern
      // (design/parity.md); BiS and sim render theirs from the served data.
      await openMock(mockPage, MOCKS_DIR, state.mock);
      const mock = await captureRegions(mockPage);

      await prepareParityPage(page);
      await openHome(page, state);
      await settleParityPage(page);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `home-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
