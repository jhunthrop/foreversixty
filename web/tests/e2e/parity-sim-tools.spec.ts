// web/tests/e2e/parity-sim-tools.spec.ts
// The simulator's tool pages against their mocks (design/parity.md): Top gear with a character and
// without one, Talents, the Droptimizer, Weights and Spec support at 1440 x 1000, each region's
// blocks and box compared (support/parity.ts). The mocks are rendered before the tests
// (support/parity-mocks.ts: ws_simtools.py) from data/builds/<active build>, the files this spec
// serves; the character arrives by ?code= (support/sim-parity.ts TOOL_CHARACTER). A side-by-side
// per state lands in test-results/parity/. Desktop only.
import type { Page } from '@playwright/test';
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
import { renderSimToolsMocks } from './support/parity-mocks';
import { openToolState, type ToolParityState } from './support/sim-parity';

const CONFIG = readParityConfig('simtools-parity.json');

// No known drift, so no EXPECTED_DRIFT: every region of every state is compared. A drift another
// lane is fixing takes parity-bis.spec.ts's EXPECTED_DRIFT pattern, with its date and lane.

async function settledBuild(page: Page, id: ToolParityState): Promise<void> {
  await prepareParityPage(page);
  await openToolState(page, id);
  await settleParityPage(page);
}

test.describe('Simulator tools mock parity at 1440', () => {
  test.skip(({ isMobile }) => isMobile, 'the parity mocks are desktop captures');
  test.use({ viewport: CONFIG.viewport });

  let mocks = '';
  // eslint-disable-next-line no-empty-pattern -- Playwright reads a hook's fixtures from its pattern
  test.beforeAll(({}, testInfo) => {
    test.skip(testInfo.project.name !== 'desktop', 'the parity mocks are desktop captures');
    mocks = renderSimToolsMocks();
  });

  for (const state of CONFIG.states) {
    test(`${state.id} matches ${state.mock}`, async ({ page, browser }) => {
      test.slow();
      const mockPage = await browser.newPage({ viewport: CONFIG.viewport });
      await openMock(mockPage, mocks, state.mock);
      const mock = await captureRegions(mockPage);

      await settledBuild(page, state.id as ToolParityState);
      const built = await captureRegions(page);

      await writeSideBySide(
        browser,
        `simtools-${state.id}`,
        await captureOutlined(mockPage, mock),
        await captureOutlined(page, built),
      );
      await mockPage.close();

      const failures = compareRegions(mock, built);
      expect(failures, `${state.id} differs from ${state.mock}:\n${formatFailures(failures)}`).toEqual([]);
    });
  }
});
